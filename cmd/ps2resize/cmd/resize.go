package cmd

import (
	"bytes"
	"fmt"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/h2non/bimg"
	// "github.com/k0kubun/pp"
	"github.com/nozzle/throttler"
	"github.com/spf13/cobra"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	psDir         string
	workDir       string
	dryRun        bool
	db            *gorm.DB
	dbName        string
	dbUser        string
	dbPass        string
	dbHost        string
	dbPort        string
	dbTablePrefix string
	imgTypes      = map[string]string{ // todo. how to deal with other types of images ?
		"p": "products",
		//"c":  "categories",
		//"m":  "manufacturers",
		//"su": "suppliers",
		//"st": "stores",
	}
	imgQuality int
	imgHeight  int
	imgWidth   int
	imgExt     []string
	imgDefExt  = []string{".png", ".jpg", ".jpeg", ".webp"} // jpeg,png, webp,tiff,gif,pdf,svg
)

// go run main.go resize --db-name eg_prestanish --db-user root --db-pass "OvdZ5ZoXWgCWL4-hvZjg!" --db-table-prefix eg_ --img-width=1000

var ImportCmd = &cobra.Command{
	Use:     "resize",
	Aliases: []string{"r"},
	Short:   "resize too big images in prestashop images directory.",
	Long:    "resize too big images in prestashop images directory",
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
					// SkipDefaultTransaction: true,
				},
			})
			if err != nil {
				log.Fatal(err)
			}

		}

		type Result struct {
			IdImage uint
		}

		bimgOpts := bimg.Options{
			Lossless: true,
			Quality:  imgQuality,
		}

		for imgType := range imgTypes {

			imgDir := filepath.Join(psDir, "img", imgType)

			// todo. create cli argument for limit
			limit := 10000
			offset := 0

			for {

				var query string
				switch imgType {
				case "products", "p":
					query = `SELECT image_shop.id_image  AS id_image
			FROM  eg_product p
			 LEFT JOIN eg_shop shop ON (shop.id_shop = 1)
			 LEFT JOIN eg_image_shop image_shop ON (image_shop.id_product = p.id_product AND image_shop.cover = 1 AND image_shop.id_shop = 1)
			 LEFT JOIN eg_image i ON (i.id_image = image_shop.id_image)
			WHERE (1 AND state = 1)
			GROUP BY p.id_product
			ORDER BY p.id_product DESC
			LIMIT {{ .offset }}, {{ .limit }}`
				}

				sqlResult := bytes.NewBufferString("")
				sqlTemplate, _ := template.New("").Parse(query)
				sqlTemplate.Execute(sqlResult, map[string]string{"offset": fmt.Sprintf("%d", offset), "limit": fmt.Sprintf("%d", limit)})

				// todo. load image list or parse them by name ?
				var results []*Result
				db.Debug().Raw(sqlResult.String()).Scan(&results)

				offset = offset + limit

				if len(results) == 0 {
					break
				}

				t := throttler.New(10, len(results))

				for _, result := range results {

					go func(r *Result) error {
						// Let Throttler know when the goroutine completes
						// so it can dispatch another worker
						defer t.Done(nil)

						subDirectories := []rune(fmt.Sprintf("%d", r.IdImage))
						var prefixPath string
						for _, subDirectory := range subDirectories {
							prefixPath = filepath.Join(prefixPath, fmt.Sprintf("%c", subDirectory))
						}
						imagePrefixPath := filepath.Join(imgDir, prefixPath)
						imagePath := filepath.Join(imagePrefixPath, fmt.Sprintf("%d.jpg", r.IdImage))
						// extension := filepath.Ext(imagePath)

						log.Println("imagePath=", imagePath)

						// Read original image
						buffer, err := bimg.Read(imagePath)
						if err != nil {
							return err
						}

						destinationFilePath := imagePath // filepath.Join(imagePrefixPath, fmt.Sprintf("%d-%s%s", r.IdImage, imageType.Name, extension))
						log.Println("destinationFilePath=", destinationFilePath)

						// Import image buffer
						newImage := bimg.NewImage(buffer)

						// Get the image type
						mediaType := newImage.Type()

						switch mediaType {
						case "unknown":
							log.Println("mediaType=", mediaType)
							return nil // errors.New("Unsupported image format")
						}

						if imgHeight > 0 {
							bimgOpts.Height = imgHeight
						}
						if imgWidth > 0 {
							bimgOpts.Width = imgWidth
						}

						// Process image quality
						newBytes, err := newImage.Process(bimgOpts)
						if err != nil {
							log.Warnln("newBytes.err=", err)
							return err
						}

						err = bimg.Write(destinationFilePath, newBytes)
						if err != nil {
							log.Warnln("Write.err=", err)
							return err
						}

						return nil
					}(result)
					t.Throttle()
				}
			}
		}
	},
}

func init() {
	// todo. manage ssl connections for db
	ImportCmd.Flags().StringVarP(&dbTablePrefix, "db-table-prefix", "", "ps_", "database table prefix")
	ImportCmd.Flags().StringVarP(&dbName, "db-name", "", "prestashop", "database name")
	ImportCmd.Flags().StringVarP(&dbUser, "db-user", "", "root", "database username")
	ImportCmd.Flags().StringVarP(&dbPass, "db-pass", "", "", "database password")
	ImportCmd.Flags().StringVarP(&dbHost, "db-host", "", "127.0.0.1", "database host")
	ImportCmd.Flags().StringVarP(&dbPort, "db-port", "", "3306", "datbase port")
	ImportCmd.Flags().StringVarP(&psDir, "ps-dir", "p", "../../../evolutive-prestashop/shared/www", "prestashop directory")
	ImportCmd.Flags().IntVarP(&imgWidth, "img-width", "w", 1000, "Image format quality. (works only for JPEG format)")
	ImportCmd.Flags().IntVarP(&imgHeight, "img-height", "i", 0, "Image format quality. (works only for JPEG format)")
	ImportCmd.Flags().IntVarP(&imgQuality, "img-quality", "q", 90, "Image format quality. (works only for JPEG format)")
	ImportCmd.Flags().StringSliceVarP(&imgExt, "img-extension", "e", imgDefExt, "Process only image extensions")
	ImportCmd.Flags().BoolVarP(&dryRun, "dry-run", "", false, "dry run")
	RootCmd.AddCommand(ImportCmd)
}

func buildFolderForImage(basePath string, id int) string {
	idStr := fmt.Sprintf("%d", id)
	idLength := len(idStr)
	for i := 0; i < idLength; i++ {
		basePath = filepath.Join(basePath, string(idStr[i]))
	}
	return basePath
}
