package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/dghubble/sling"
	"github.com/k0kubun/pp"
	"github.com/nozzle/throttler"
	"github.com/spf13/cobra"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	dryRun        bool
	psDomain      string
	psDir         string
	matchAddr     string
	db            *gorm.DB
	dbName        string
	dbUser        string
	dbPass        string
	dbHost        string
	dbPort        string
	dbTablePrefix string
	dbDrop        bool
)

var IndexCmd = &cobra.Command{
	Use:     "index",
	Aliases: []string{"i"},
	Short:   "index prestashop images into image match.",
	Long:    "index prestashop images into image match.",
	Run: func(cmd *cobra.Command, args []string) {

		if !dryRun {
			var err error
			// connect to database
			dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
			db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
				NamingStrategy: schema.NamingStrategy{
					TablePrefix:   dbTablePrefix,                   // table name prefix, table for `Lang` would be `eg_lang`
					SingularTable: true,                            // use singular table name, table for `User` would be `user` with this option enabled
					NameReplacer:  strings.NewReplacer("ID", "Id"), // use name replacer to change struct/field name before convert it to db name
				},
			})
			if err != nil {
				log.Fatal(err)
			}
		}

		if dbDrop {
			if db.Migrator().HasTable(&Match{}) {
				db.Migrator().DropTable(&Match{})
			}
		}

		db.Migrator().AutoMigrate(&Match{})

		query := `SELECT   i.id_image as id_image,
						   il.id_lang as id_lang,
						   i.id_product as id_product,
					       i.id_shop as id_shop,
					       il.legend as legend,
					       pl.link_rewrite as link_rewrite
					FROM ` + dbTablePrefix + `image_shop i 
					LEFT JOIN ` + dbTablePrefix + `product_lang pl ON pl.id_product=i.id_product AND pl.id_lang=1
					LEFT JOIN ` + dbTablePrefix + `image_lang il ON il.id_image=i.id_image AND il.id_lang=1`

		var images []*Image
		db.Raw(query).Scan(&images)

		t := throttler.New(2, len(images))

		for _, image := range images {

			go func(i *Image) error {
				// Let Throttler know when the goroutine completes
				// so it can dispatch another worker
				defer t.Done(nil)

				// Prepare metadata
				imageJson, err := json.Marshal(image)
				if err != nil {
					log.Warn(err)
					return err
				}

				imgPrefixPath := buildFolderForImage(filepath.Join("img", "p"), i.IdImage)
				mediaUrl := filepath.Join(imgPrefixPath, fmt.Sprintf("%d.jpg", i.IdImage))

				// BodyForm parameters
				params := &Add{
					Url:      psDomain + "/" + mediaUrl,
					Filepath: mediaUrl,
					Metadata: string(imageJson),
				}

				pp.Println("params", params)

				var existsImage Match
				err = db.Debug().Where("id_image = ? ", i.IdImage, true).First(&existsImage).Error
				if !errors.Is(err, gorm.ErrRecordNotFound) {
					log.Infoln("Skipping: ", i.IdImage)
					return nil
				}

				httpClient := &http.Client{}

				// Preparing request
				summaryBase := sling.New().Base(matchAddr).Client(httpClient)
				req, err := summaryBase.New().Post("add").BodyForm(params).Request()
				if err != nil {
					log.Warn(err)
					return err
				}

				// get response
				resp, err := httpClient.Do(req)
				if err != nil {
					log.Warn(err)
					return err
				}
				defer resp.Body.Close()
				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					log.Warn(err)
					return err
				}

				// parsing response
				var response Response
				err = json.Unmarshal(body, &response)
				if err != nil {
					log.Warn(err)
					return err
				}

				// Index to dsys match
				pp.Println(response)
				if response.Status != "ok" {
					log.Infoln("params:", params)
					log.Warn("Something went wrong while indexing the media.")
				} else {
					log.Infoln("Submitted media >> ", i.IdImage)
					img := &Match{
						IdProduct: i.IdProduct,
						IdImage:   i.IdImage,
						IdShop:    i.IdShop,
						Filepath:  params.Filepath,
						// Fingerprint: fingerprint,
						// Mimetype:    mimeType,
					}
					if err := db.Create(&img).Error; err != nil {
						log.Fatal(err)
					}
				}

				return nil

			}(image)
			t.Throttle()

		}

		if t.Err() != nil {
			// Loop through the errors to see the details
			for i, err := range t.Errs() {
				log.Printf("error #%d: %s", i, err)
			}
			log.Fatal(t.Err())
		}

	},
}

func init() {
	IndexCmd.Flags().StringVarP(&dbTablePrefix, "db-table-prefix", "", "ps_", "database table prefix")
	IndexCmd.Flags().StringVarP(&dbName, "db-name", "", "prestashop", "database name")
	IndexCmd.Flags().StringVarP(&dbUser, "db-user", "", "root", "database username")
	IndexCmd.Flags().StringVarP(&dbPass, "db-pass", "", "", "database password")
	IndexCmd.Flags().StringVarP(&dbHost, "db-host", "", "127.0.0.1", "database host")
	IndexCmd.Flags().StringVarP(&dbPort, "db-port", "", "3306", "datbase port")
	IndexCmd.Flags().StringVarP(&matchAddr, "match-addr", "", "http://localhost:8888", "image match service address")
	IndexCmd.Flags().StringVarP(&psDomain, "ps-domain", "", "https://prestanish.evolutive.group", "prestashop domain with protocol")
	IndexCmd.Flags().StringVarP(&psDir, "ps-dir", "", "../../../shared/www", "prestashop directory")
	IndexCmd.Flags().BoolVarP(&dryRun, "dry-run", "", false, "dry run")
	RootCmd.AddCommand(IndexCmd)
}

func buildFolderForImage(basePath string, id uint) string {
	idStr := fmt.Sprintf("%d", id)
	idLength := len(idStr)
	for i := 0; i < idLength; i++ {
		basePath = filepath.Join(basePath, string(idStr[i]))
	}
	return basePath
}
