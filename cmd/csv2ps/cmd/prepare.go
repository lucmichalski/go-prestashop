package cmd

import (
	"bytes"
	"crypto/tls"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"text/template"
	"time"
	"unicode"

	"github.com/beevik/etree"
	gofakeit "github.com/brianvoe/gofakeit/v6"
	"github.com/cavaliercoder/grab"
	"github.com/csimplestring/go-csv/detector"
	badger "github.com/dgraph-io/badger"
	"github.com/golang/snappy"
	"github.com/gosimple/slug"
	"github.com/iancoleman/strcase"
	"github.com/k0kubun/pp"
	"github.com/nozzle/throttler"
	"github.com/spf13/cobra"
	ccsv "github.com/tsak/concurrent-csv-writer"
	"github.com/yudppp/json2struct"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"gopkg.in/yaml.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	psm "github.com/lucmichalski/go-prestashop/internal/models"
)

var (
	psDir         string
	psMarketplace bool
	outputDir     string
	workDir       string
	proxyURLStr   string
	proxyURL      *url.URL
	dryRun        bool
	db            *gorm.DB
	kv            *badger.DB
	kvPath        = "./badger"
	dbName        string
	dbUser        string
	dbPass        string
	dbHost        string
	dbPort        string
	dbTablePrefix string
	dbMigrate     bool
	dbDrop        bool
	dbEnv         bool
	activeLangs   []*psm.Lang
	activeShops   []*psm.Shop
	activeGroups  []*psm.Group
	languages     []string
	languagesDef  = []string{"fr", "en"}
	imgTypes      = map[string]string{ // todo. how to deal with other types of images ?
		"p": "products",
		//"c":  "categories",
		//"m":  "manufacturers",
		//"su": "suppliers",
		//"st": "stores",
	}
	imgQuality int
	imgExt     []string
	imgDefExt  = []string{".png", ".jpg", ".jpeg", ".webp"} // jpeg,png, webp,tiff,gif,pdf,svg
	seed       = time.Now().UTC().UnixNano()
	noFixtures bool
)

func init() {
	gofakeit.Seed(0)
}

const (
	// not operational
	catSqlTemplate = `
INSERT IGNORE INTO {{.prefix }}category (id_category, id_parent, id_shop_default, level_depth, nleft, nright, active, date_add, date_upd, is_root_category)
VALUES ('33', '32', '1', '5', '0', '0', '1', NOW(), NOW(), '0');
INSERT IGNORE INTO {{.prefix }}category_lang (id_category, id_shop, id_lang, name, description, link_rewrite, meta_description, meta_keywords, meta_title)
VALUES ('33', '1', '1', 'Silencieux arrière', '', 'silencieux-arri-re', '', '', '');
INSERT IGNORE INTO {{.prefix }}category_shop (id_category, id_shop, position)
VALUES ('33', '1', '0');
INSERT IGNORE INTO {{.prefix }}category_group (id_category, id_group)
VALUES ('33', 1), ('33', 2), ('33', 3);
`

	baseTemplate = `package netaffiliation

import (
	"strings"

	"gorm.io/gorm"

	"github.com/lucmichalski/go-prestashop/internal/models"
)

var ( 
	db *gorm.DB
	dbTablePrefix string
	Catalogs = []interface{}{
					{{ range .Catalogs }}
					&{{ . }}{},{{ end }}
				}
)

func initDB(client *gorm.DB, tablePrefix string) *gorm.DB {
	db = client
	dbTablePrefix = tablePrefix
	return db
}

func findOrCreateCategoryByName(db *gorm.DB, dbTablePrefix, name string) (*models.CategoryLang, error) {
	var cat models.CategoryLang
	db.Debug().Raw("SELECT * FROM " + dbTablePrefix + "category_lang WHERE name = '" + strings.ToLower(name) + "'").Scan(&cat)
	if cat.IDCategory == 0 {
		db.Debug().Raw("SELECT * FROM eg_lang WHERE iso_code = '" + strings.ToLower(iso_code) + "'").Scan(&lang)
		return &lang, nil // todo. fix this hacky return
	}
	return &cat, nil
}`
)

const packageTemplate = `package netaffiliation

import (
	"fmt"

	// "github.com/fatih/structs"
	// "github.com/k0kubun/pp"
)

/*

// Catalog Path: {{ .CatalogPath }}
// Catalog Separator: {{ .CatalogSeparator }}
// Catalog excerpt: #{{ .Line }}

// ===================================================================================== \\
{{ .Row }}
\\ ===================================================================================== //

*/

{{ .Struct }}

/*
func (o *{{ .StructName }}) Convert() ([]string, error) {

	mp := make(map[string]string, 0) 

	mp["product_name"] = o.ProductName
	mp["reference"] = o.Reference

	categoryId, err := findOrCreateCategoryByNam(db, dbTablePrefix, o.Category)
	if err != nil {
		return nil, err
	}
	mp["category_id"] = fmt.Sprintf("%d", categoryId)

	defaultCategory, err := findOrCreateCategoryByName(db, dbTablePrefix, o.DefaultCategory)
	if err != nil {
		return nil, err
	}
	mp["default_category"] = fmt.Sprintf("%d", defaultCategory)

	mp["short_description"] = o.ShortDescription
	mp["description"] = o.Description
	mp["price"] = fo.Price
	mp["quantity"] = o.Quantity
	mp["image_ref"] = o.ImageRef

	// go ast ?
	// m := structs.Name(o)
	// pp.Println(m)
	// for _, feature := range o.Features {}
	mp["feature_name"] = o.FeatureNames
	mp["feature_values"] = o.FeatureValues

	row =  []string{
		mp["product_name"], 
		mp["reference"], 
		mp["category_id"],
		mp["default_category"],
		mp["short_description"],
		mp["description"],
		mp["price"],
		mp["quantity"],
		mp["image_ref"],
		mp["feature_name"],
		mp["feature_values"],
	}

	return row, nil
}
*/
`

var PrepareCmd = &cobra.Command{
	Use:     "prepare",
	Aliases: []string{"p"},
	Short:   "prepare csv catalog to prestashop import format (core or webkul's marketplace).",
	Long:    "prepare csv catalog to prestashop import format (core or webkul's marketplace).",
	Run: func(cmd *cobra.Command, args []string) {

		gtotal := 0

		if !dryRun {
			var err error
			// connect to database
			dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&collation=utf8mb4_general_ci&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
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

		var err error
		proxyURL, err = url.Parse(proxyURLStr)
		if err != nil {
			log.Fatal(err)
		}

		kv, err = badger.Open(badger.DefaultOptions(kvPath))
		if err != nil {
			log.Fatal(err)
		}
		defer kv.Close()

		osPathname := "catalogs/netaffiliation.xml"
		if options.debug {
			pp.Println("osPathname:", osPathname)
		}

		doc := etree.NewDocument()
		if err := doc.ReadFromFile(osPathname); err != nil {
			panic(err)
		}

		root := doc.SelectElement("campaigns")
		fmt.Println("ROOT element:", root.Tag)

		var catalogs []string
		reStruct := regexp.MustCompile(`type (.*) struct`)
		reTags := regexp.MustCompile(`json:"(.*)"`)

		err = db.Where("active = ?", 1).Find(&activeLangs).Error
		pp.Println("activeLangs:", activeLangs)

		err = db.Where("active = ?", 1).Find(&activeShops).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Fatal("active shops not found")
		}
		pp.Println("activeShops:", activeShops)

		err = db.Find(&activeGroups).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Fatal("groups not found")
		}
		pp.Println("activeGroups:", activeGroups)

		campaigns := root.SelectElements("campaign")
		// campaigns = shuffleXml(campaigns)

		t := throttler.New(3, len(campaigns))

		// shuffle
		for _, campaign := range campaigns {

			// todo. move throttler here, will avoid duplicate categories
			go func(c *etree.Element) error {
				// Let Throttler know when the goroutine completes
				// so it can dispatch another worker
				defer t.Done(nil)

				var feedURL string
				var feedName string
				var feedVersion string
				for _, pf := range c.SelectElements("product_feeds") {
					if product_feed_url := pf.SelectElement("product_feed"); product_feed_url != nil {
						feedVersion = product_feed_url.SelectAttrValue("version", "unknown")
						if feedVersion != "4" {
							continue
						}
						feedURL = product_feed_url.Text()
						feedName = product_feed_url.SelectAttrValue("name", "unknown")
					}
				}

				pp.Println("feedName:", feedName)
				pp.Println("feedURL:", feedURL)
				pp.Println("feedVersion:", feedVersion)

				// create a new catalog mapping
				cmp := &CatalogMap{
					Name: feedName,
					Feed: feedURL,
				}

				// create fake marketplace user
				if psMarketplace {

					var randomCountry psm.Country
					db.Order("RAND()").First(&randomCountry)

					var randomLanguage psm.Lang
					db.Order("RAND()").First(&randomLanguage)

					var randomCustomer psm.Customer
					db.Order("RAND()").First(&randomCustomer)

					if feedName == "" {
						feedName = gofakeit.Company()
					}

					seller := WkMpSeller{
						ShopNameUnique:          feedName,
						LinkRewrite:             slug.Make(feedName),
						SellerFirstname:         gofakeit.FirstName(),
						SellerLastname:          gofakeit.LastName(),
						BusinessEmail:           gofakeit.Email(),
						Phone:                   gofakeit.Phone(),
						Fax:                     "",
						Address:                 gofakeit.Street(),
						Postcode:                gofakeit.Zip(),
						City:                    gofakeit.City(),
						IdCountry:               int(randomCountry.IDCountry),
						IdState:                 0, // todo. left join on country select
						TaxIdentificationNumber: gofakeit.UUID(),
						DefaultLang:             int(randomLanguage.IDLang),
						FacebookId:              "",
						TwitterId:               "",
						GoogleId:                "",
						InstagramId:             "",
						ProfileImage:            "",
						ProfileBanner:           "",
						ShopImage:               "",
						ShopBanner:              "",
						Active:                  true,
						ShopApproved:            true,
						SellerCustomerId:        int(randomCustomer.IDCustomer), // randome user
						SellerDetailsAccess:     "",
						DateAdd:                 time.Now(),
						DateUpd:                 time.Now(),
					}

					var wkMpSeller WkMpSeller
					err := db.Where("shop_name_unique = ? ", feedName).First(&wkMpSeller).Error
					if errors.Is(err, gorm.ErrRecordNotFound) {
						err := db.Create(&seller).Error
						if err != nil {
							log.Fatal(err)
						}
						db.Where("shop_name_unique = ? ", feedName).First(&wkMpSeller)
						cmp.SellerId = wkMpSeller.IdSeller
						// WkMpSellerLang
						for _, activeLang := range activeLangs {
							wkMpSellerLang := &WkMpSellerLang{
								IdSeller:  wkMpSeller.IdSeller,
								IdLang:    activeLang.IDLang,
								ShopName:  feedName,
								AboutShop: gofakeit.HackerPhrase(),
							}
							err := db.Create(&wkMpSellerLang).Error
							if err != nil {
								return err
							}
						}
					} else {
						cmp.SellerId = wkMpSeller.IdSeller
					}

				}

				cmp.Mapping.LangSuffix = languagesDef
				cmp.Mapping.LangFields = []string{"name", "description", "description_short"}

				client := grab.NewClient()

				// proxyURL

				tr := &http.Transport{
					// Proxy: http.ProxyURL(proxyURL),
					DialContext: (&net.Dialer{
						Timeout: 240 * time.Second,
						// KeepAlive: 30 * time.Second,
						DualStack: true,
					}).DialContext,
					MaxIdleConns:          100,
					IdleConnTimeout:       240 * time.Second,
					TLSHandshakeTimeout:   240 * time.Second,
					ExpectContinueTimeout: 1 * time.Second,
					DisableKeepAlives:     true,
					TLSClientConfig: &tls.Config{
						InsecureSkipVerify: true,
					},
				}

				client.UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.97 Safari/537.36"
				client.HTTPClient = &http.Client{
					Timeout:   240 * time.Second,
					Transport: tr,
					CheckRedirect: func(req *http.Request, via []*http.Request) error {
						return nil
					},
				}

				feedName = stripCtlAndExtFromUnicode(feedName)

				localCatalogStruct := filepath.Join("..", "..", "internal", "netaffiliation", fmt.Sprintf("%s.go", strcase.ToSnake(feedName)))
				localCatalogJson := filepath.Join("catalogs", fmt.Sprintf("%s.json", slug.Make(feedName)))
				localCatalogCsv := filepath.Join("catalogs", fmt.Sprintf("%s.csv", slug.Make(feedName)))
				localCatalogMapYaml := filepath.Join("imports", fmt.Sprintf("%s.yml", slug.Make(feedName)))
				localCatalogFormattedCsv := filepath.Join("imports", fmt.Sprintf("%s.csv", slug.Make(feedName)))

				if _, err := os.Stat(localCatalogCsv); os.IsNotExist(err) {

					req, err := grab.NewRequest(localCatalogCsv, feedURL)
					if err != nil {
						log.Warnln(err)
						return err
					}

					// start download
					fmt.Printf("Downloading %v...\n", req.URL())
					resp := client.Do(req)
					if resp.HTTPResponse == nil {
						log.Warnln(err)
						return err
					}
					fmt.Printf("  %v\n", resp.HTTPResponse.Status)

					// start UI loop
					t := time.NewTicker(500 * time.Millisecond)
					defer t.Stop()

				LoopRel:
					for {
						select {
						case <-t.C:
							fmt.Printf("  transferred %v / %v bytes (%.2f%%)\n",
								resp.BytesComplete(),
								resp.Size,
								100*resp.Progress())

						case <-resp.Done:
							// download is complete
							break LoopRel
						}
					}

					// check for errors
					if err := resp.Err(); err != nil {
						log.Warnln("Download failed:", err, feedURL)
						return err
					}

					fmt.Printf("Download saved to %v \n", resp.Filename)
					localCatalogCsv = resp.Filename
				}

				inputFile, err := os.Open(localCatalogCsv)
				if err != nil {
					log.Warnln(err)
					return err
				}
				defer inputFile.Close()

				detect := detector.New()

				delimiters := detect.DetectDelimiter(inputFile, '"')
				//pp.Println("delimiters:", delimiters, "localCatalogCsv", localCatalogCsv)

				if len(delimiters) == 0 || len(delimiters) > 1 {
					return errors.New("No delimiter detected")
				}

				cmp.Separator = delimiters[0]

				cols, err := getHeaders(localCatalogCsv, delimiters[0])
				if err != nil {
					log.Warnln(err)
					return err
				}

				cmp.Fields = cols
				cmp.Mapping.Update = true
				//pp.Println("cols:", cols)

				for _, col := range cols {
					switch col {
					case "url", "product_url", "ur_lproduit", "product_page_url", "link", "url_de_la_page_produit", "url_produit", "prod_url", "product_link":
						cmp.Mapping.Product.Redirect = col
					case "name", "nom", "titre", "name_of_the_product", "title", "product_name", "prod_name", "nom_attribute", "nom_usuel_du_produit":
						cmp.Mapping.Product.Name = col
					case "reference", "ref", "reference_interne", "identifiant", "internal_reference", "reference_du_produit", "reference_fabriquant":
						cmp.Mapping.Product.Reference = col
					case "ean13", "ean", "ean_or_isbn", "ean_13", "reference_universelle", "universal_reference", "prod_ean", "id", "code_ean":
						cmp.Mapping.Product.Ean13 = col
					case "sku", "numero_modele_produit":
						cmp.Mapping.Product.Sku = col
					case "prix", "price", "current_price", "prix_de_vente", "prix_actuel", "prod_price", "prix_ttc", "prix_ttc_du_produit":
						cmp.Mapping.Product.Price = col
					case "mpn":
						cmp.Mapping.Product.Mpn = col
					case "description", "descriptif", "product_detail", "description_html", "discription_of_the_product", "prod_description_long":
						cmp.Mapping.Product.Description = col
					case "description_short", "product_highlight", "prod_description":
						cmp.Mapping.Product.DescriptionShort = col
					case "image":
						cmp.Mapping.Product.Image = col
					case "quantity", "quantite", "stock", "indicateur_de_stock", "stock_indicator", "qty": //, "disponibilite", "availability":
						cmp.Mapping.Product.Quantity = col
					case "img_large", "url_grande", "url_image_grande", "url_image", "big_image", "url_de_l'image_par_defaut", "image_link", "image_product", "url_of_the_big_image", "image_default", "image_url_1", "image_url":
						cmp.Mapping.Product.Image = col
					case "hauteur_cm":
						cmp.Mapping.Product.Height = col
					case "largeur_cm":
						cmp.Mapping.Product.Width = col
					case "poids_kg", "weight":
						cmp.Mapping.Product.Weight = col
					case "shipping_and_handling_cost", "frais_de_port", "frais_de_ports", "shipping":
						cmp.Mapping.Product.Shipping = col
					case "taille", "couleur", "matiere", "size", "color":
						cmp.Mapping.Product.Attributes = append(cmp.Mapping.Product.Attributes, col)
					case "marque", "brand", "genre", "garantie", "modele", "saison", "collection", "fabricant", "condition", "name_of_brand", "montage", "nom_marque", "type_peinture", "brand_name", "famille_de_produit", "largeur", "hauteur", "vitesse", "diametre", "charge", "runflat", "vehicule", "rechape":
						cmp.Mapping.Product.Features = append(cmp.Mapping.Product.Features, col)
					case "categorie_finale":
						cmp.Mapping.Category.Name = col
					case "category_breadcrumb", "categorie", "category", "breadcrumb", "product_category", "noms_de_toutes_les_categories":
						cmp.Mapping.Category.Breadcrumb = col
					}
				}

				// create a formatted version of the catalog
				cmp, total, err := csv2ps(db, localCatalogCsv, localCatalogFormattedCsv, cols, delimiters[0], cmp)
				checkErr("while writing formatted for prestashop csv file", err)

				cmp.Total = total
				gtotal += total

				outputBytes, err := csv2json(localCatalogCsv, cols, delimiters[0])
				if err != nil {
					log.Warnln(err)
					return err
				}

				cmpBytes, err := yaml.Marshal(cmp)
				if err != nil {
					fmt.Printf("err: %v\n", err)
					return err
				}
				err = ioutil.WriteFile(localCatalogMapYaml, cmpBytes, 0644)
				checkErr("while writing mapping yaml file", err)

				err = ioutil.WriteFile(localCatalogJson, outputBytes, 0644)
				checkErr("while writing json file", err)

				json2struct.SetDebug(options.debug)
				opt := json2struct.Options{
					UseOmitempty:   false,
					UseShortStruct: true,
					UseLocal:       false,
					UseExample:     false,
					Prefix:         "",
					Suffix:         "",
					Name:           strings.ToLower(feedName),
				}

				jsonFile, err := os.Open(localCatalogJson)
				if err != nil {
					log.Warnln(err)
					return err
				}
				defer jsonFile.Close()

				parsed, err := json2struct.Parse(jsonFile, opt)
				if err != nil {
					return err
				}

				// csvRow
				csvRowLine := 1
				csvRowExample, err := getExampleRow(localCatalogCsv, delimiters[0], cols, csvRowLine)
				// checkErr("while getting example row", err)
				if err != nil {
					return err
				}

				// catalogs
				structName := reStruct.FindStringSubmatch(parsed)

				parsed = reTags.ReplaceAllString(parsed, `json:"$1" struct2map:"key:$1"`)

				catalogs = append(catalogs, structName[1])

				structResult := bytes.NewBufferString("")
				structTemplate, _ := template.New("").Parse(packageTemplate)
				structTemplate.Execute(structResult, map[string]string{
					"CatalogSeparator": delimiters[0],
					"CatalogPath":      localCatalogCsv,
					"Line":             fmt.Sprintf("%d", csvRowLine),
					"Struct":           parsed,
					"StructName":       stripCtlAndExtFromUnicode(structName[1]),
					"Row":              strings.Join(csvRowExample, ",\n")},
				)

				err = ioutil.WriteFile(localCatalogStruct, structResult.Bytes(), 0644)
				// checkErr("while writing struct file", err)
				if err != nil {
					return err
				}

				return nil

			}(campaign)

			t.Throttle()

		}

		if t.Err() != nil {
			// Loop through the errors to see the details
			for i, err := range t.Errs() {
				log.Printf("error #%d: %s", i, err)
			}
			log.Fatal(t.Err())
		}

		localCatalogBase := filepath.Join("..", "..", "internal", "netaffiliation", fmt.Sprintf("%s.go", "base"))
		catalogResult := bytes.NewBufferString("")
		catalogTemplate, _ := template.New("").Parse(baseTemplate)
		catalogTemplate.Execute(catalogResult, map[string][]string{"Catalogs": catalogs})
		err = ioutil.WriteFile(localCatalogBase, catalogResult.Bytes(), 0644)
		checkErr("while writing struct file", err)

		pp.Println("total new entries:", gtotal)

	},
}

func init() {
	// todo. manage ssl connections for db
	PrepareCmd.Flags().StringVarP(&dbTablePrefix, "db-table-prefix", "", "ps_", "database table prefix")
	PrepareCmd.Flags().StringVarP(&dbName, "db-name", "", "prestashop", "database name")
	PrepareCmd.Flags().StringVarP(&dbUser, "db-user", "", "root", "database username")
	PrepareCmd.Flags().StringVarP(&dbPass, "db-pass", "", "", "database password")
	PrepareCmd.Flags().StringVarP(&dbHost, "db-host", "", "127.0.0.1", "database host")
	PrepareCmd.Flags().StringVarP(&dbPort, "db-port", "", "3306", "datbase port")
	PrepareCmd.Flags().BoolVarP(&dbEnv, "db-env", "", false, "use env variables to connect to the database.")
	PrepareCmd.Flags().BoolVarP(&dbDrop, "db-drop", "", false, "drop/truncate database tables")
	PrepareCmd.Flags().BoolVarP(&dbMigrate, "db-migrate", "", false, "create/update database tables")
	PrepareCmd.Flags().StringVarP(&proxyURLStr, "proxy", "x", "socks5://localhost:5566", "tor haproxy url")
	PrepareCmd.Flags().BoolVarP(&psMarketplace, "ps-market", "", false, "create wk-marketplace users and catalogs per feed.")
	PrepareCmd.Flags().StringVarP(&psDir, "ps-dir", "p", "../../../evolutive-prestashop/shared/www", "prestashop directory")
	PrepareCmd.Flags().IntVarP(&imgQuality, "img-quality", "q", 90, "Image format quality. (works only for JPEG format)")
	PrepareCmd.Flags().StringSliceVarP(&imgExt, "img-extension", "e", imgDefExt, "Process only image extensions")
	PrepareCmd.Flags().StringVarP(&outputDir, "output-dir", "o", "../../shared/csv", "output directory")
	PrepareCmd.Flags().StringVarP(&workDir, "workdir", "w", "../../shared/fixtures", "working directory")
	PrepareCmd.Flags().StringSliceVarP(&languages, "language", "l", languagesDef, "languages to import")
	PrepareCmd.Flags().BoolVarP(&noFixtures, "no-fixtures", "", false, "skip fixtures loading")
	PrepareCmd.Flags().BoolVarP(&dryRun, "dry-run", "", false, "dry run")
	RootCmd.AddCommand(PrepareCmd)
}

func buildFolderForImage(basePath string, id int) string {
	idStr := fmt.Sprintf("%d", id)
	idLength := len(idStr)
	for i := 0; i < idLength; i++ {
		basePath = filepath.Join(basePath, string(idStr[i]))
	}
	return basePath
}

func insert(ss []string, s string) []string {
	i := sort.SearchStrings(ss, s)
	ss = append(ss, "")
	copy(ss[i+1:], ss[i:])
	ss[i] = s
	return ss
}

func encodeToCsv(columns []string, rows [][]string) ([]byte, error) {
	var buf bytes.Buffer
	w := csv.NewWriter(&buf)
	sort.Strings(columns)
	if err := w.Write(columns); err != nil {
		return nil, err
	}
	for _, row := range rows {
		if err := w.Write(row); err != nil {
			return nil, err
		}
		w.Flush()
	}
	return buf.Bytes(), nil
}

func prepareTables(db *gorm.DB, tables ...interface{}) error {
	for _, table := range tables {
		if dbDrop {
			if db.Migrator().HasTable(table) {
				db.Migrator().DropTable(table)
			}
		}
		if dbMigrate {
			db.Migrator().AutoMigrate(table)
		}
	}
	return nil
}

func getExampleRow(fp, sep string, cols []string, line int) ([]string, error) {
	var headers, example, formated []string
	// read
	file, err := os.Open(fp)
	if err != nil {
		return headers, err
	}
	defer file.Close()
	r := csv.NewReader(file)
	if sep == "" {
		sep = ";"
	}
	r.Comma = []rune(sep)[0]
	r.LazyQuotes = true
	// r.Comment = '#'

	i := 0
	for {
		row, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			if perr, ok := err.(*csv.ParseError); ok && perr.Err == csv.ErrFieldCount {
				continue
			}
			return headers, err
		}
		if i == 0 {
			headers = row
		}
		if i == line {
			example = row
			break
		}
		i++
	}

	if len(headers) != len(example) {
		log.Fatal("mismtach betwen columns and records")
	}
	for i, n := range headers {
		formated = append(formated, fmt.Sprintf("%s: %s", n, example[i]))
	}
	return formated, nil
}

func csv2ps(db *gorm.DB, fp, fo string, columns []string, separator string, cmp *CatalogMap) (*CatalogMap, int, error) {
	rows := make([]map[string]string, 0)

	// read
	file, err := os.Open(fp)
	if err != nil {
		return cmp, 0, err
	}
	defer file.Close()

	// Csv Reader
	r := csv.NewReader(file)
	if separator == "" {
		separator = ";"
	}
	r.Comma = []rune(separator)[0]
	r.LazyQuotes = true
	r.Comment = '#'

	fmt.Println("separator=", separator)

	columns = normalize_headers(columns)

	z := 0
	for {

		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			if perr, ok := err.(*csv.ParseError); ok && perr.Err == csv.ErrFieldCount {
				continue
			}
			return cmp, 0, err
		}

		if z == 0 {
			z++
			continue
		}

		// pp.Println(z, "==", record)
		row := make(map[string]string, len(columns))

		// if len(record) != len(columns) {
		// 	continue
		// }

		for i, n := range columns {
			row[n] = record[i] // record[i]
		}
		rows = append(rows, row)
		z++
	}

	// CSV Writer
	// formatted, err := os.Create(fo)
	// if err != nil {
	// 	return cmp, 0, err
	// }
	// defer formatted.Close()

	// Create `sample.csv` in current directory
	writer, err := ccsv.NewCsvWriter(fo)
	if err != nil {
		panic("Could not open `sample.csv` for writing")
	}

	// Flush pending writes and close file upon exit of main()
	defer writer.Close()

	//writer := csv.NewWriter(formatted)
	//writer.Comma = ','
	//defer writer.Flush()

	headers := []string{"name_en", "name_fr", "reference", "category_id", "default_category", "short_description_en", "short_description_fr", "description_en", "description_fr", "price", "tax_id", "quantity", "image_ref", "feature_name", "feature_values", "minimum_garantie", "quantite_unitaire", "format_unitaire", "prix_total", "unite_total", "poids_piece", "ean13", "sku", "mpn"}
	err = writer.Write(headers)
	checkErr("Cannot write to file", err)

	detect := detector.New()

	t := throttler.New(4, len(rows))

	c := 0

	// shuffle rows
	// rows = shuffleMap(rows)

	for _, value := range rows {

		go func(v map[string]string) error {
			// Let Throttler know when the goroutine completes
			// so it can dispatch another worker
			defer t.Done(nil)

			var row []string

			var redirectLink string
			if cmp.Mapping.Product.Redirect != "" {
				productRedirect := cmp.Mapping.Product.Redirect
				redirectLink = v[productRedirect]
			} else {
				row = append(row, "")
			}
			_, ok := getFromBadger(redirectLink)
			if ok {
				log.Infoln("skipping link redirect as processed:", redirectLink)
				return nil
			}

			product := &Product{}

			// product_name
			if cmp.Mapping.Product.Name != "" {
				productName := cmp.Mapping.Product.Name
				name := v[productName]
				if len(name) > 127 {
					product.Name = name[0:127]
				} else {
					product.Name = name
				}
				row = append(row, product.Name)
			} else {
				row = append(row, "")
			}

			// product_name
			if cmp.Mapping.Product.Name != "" {
				productName := cmp.Mapping.Product.Name
				name := v[productName]
				if len(name) > 127 {
					product.Name = name[0:127]
				} else {
					product.Name = name
				}
				row = append(row, product.Name)
			} else {
				row = append(row, "")
			}

			// reference
			if cmp.Mapping.Product.Reference != "" {
				productRef := cmp.Mapping.Product.Reference
				product.Reference = v[productRef]
				row = append(row, v[productRef])
			} else {
				row = append(row, "")
			}

			// category_id / 200;212;253
			// reference
			var category_ids string
			var default_category string
			if cmp.Mapping.Category.Breadcrumb != "" {
				breadcrumb := cmp.Mapping.Category.Breadcrumb
				sanitizedBreadcrumb := strings.Replace(v[breadcrumb], ">", " > ", -1) // fix issue with lagiuole breadcrumb
				sanitizedBreadcrumb = strings.Replace(sanitizedBreadcrumb, "  >  ", " > ", -1)
				sanitizedBreadcrumb = strings.Replace(sanitizedBreadcrumb, "-", " ", -1)
				sanitizedBreadcrumb = strings.Replace(sanitizedBreadcrumb, "&", " and ", -1)
				sanitizedBreadcrumb = strings.Replace(sanitizedBreadcrumb, "_", " ", -1)
				sanitizedBreadcrumb = strings.Replace(sanitizedBreadcrumb, "'", " ", -1)
				sanitizedBreadcrumb = normalize_line(sanitizedBreadcrumb)
				if !strings.Contains(sanitizedBreadcrumb, "=>") {
					delimiters := detect.DetectDelimiter(strings.NewReader(sanitizedBreadcrumb), '"')
					// Do a search in the database for those categories
					if len(delimiters) > 0 && cmp.Mapping.Category.Separator == "" {
						cmp.Mapping.Category.Separator = delimiters[0]
						//pp.Println("cmp.Category.Separator = ", delimiters[0], "v=", v[breadcrumb])
						// row = append(row, "") //			row = append(row, v[breadcrumb])
					}
					// create nested category
					//pp.Println("cmp.Mapping.Category.Separator", cmp.Mapping.Category.Separator, "delimiters", delimiters)
					var breadcrumbSlice []string
					if cmp.Mapping.Category.Separator == "" {
						breadcrumbSlice = []string{v[breadcrumb]}
					} else {
						breadcrumbSlice = strings.Split(v[breadcrumb], cmp.Mapping.Category.Separator)
					}
					for i, category := range breadcrumbSlice {
						breadcrumbSlice[i] = strings.TrimSpace(category)
					}
					var err error
					category_ids, default_category, err = createNestedCategoryTree(db, cmp, breadcrumbSlice)
					checkErr("createNestedCategoryTreeError", err)
				} else {
					cmp.Mapping.Category.Separator = "=>"
				}
				// category_id
				row = append(row, category_ids)
			} else {
				category_ids, default_category, err = createNestedCategoryTree(db, cmp, nil)
				checkErr("createNestedCategoryTreeError", err)
				row = append(row, category_ids)
			}

			// row = append(row, "")

			// default_category / 253
			row = append(row, default_category)

			// short_description
			if cmp.Mapping.Product.DescriptionShort != "" {
				productDescShort := cmp.Mapping.Product.DescriptionShort
				product.DescriptionShort = v[productDescShort]
				row = append(row, v[productDescShort])
			} else {
				row = append(row, "")
			}

			// short_description
			if cmp.Mapping.Product.DescriptionShort != "" {
				productDescShort := cmp.Mapping.Product.DescriptionShort
				row = append(row, v[productDescShort])
			} else {
				row = append(row, "")
			}

			// description
			if cmp.Mapping.Product.Description != "" {
				productDesc := cmp.Mapping.Product.Description
				if v[productDesc] == "" {
					return nil
				}
				product.Description = v[productDesc]
				row = append(row, v[productDesc])
			} else {
				row = append(row, "")
			}

			// description
			if cmp.Mapping.Product.Description != "" {
				productDesc := cmp.Mapping.Product.Description
				row = append(row, v[productDesc])
			} else {
				row = append(row, "")
			}

			// price
			if cmp.Mapping.Product.Price != "" {
				productPrice := cmp.Mapping.Product.Price
				product.Price = strings.Replace(v[productPrice], ",", ".", -1)
				row = append(row, strings.Replace(v[productPrice], ",", ".", -1))
			} else {
				row = append(row, "")
			}

			// tax_id
			row = append(row, "1")

			// quantity
			row = append(row, "10")

			var imageRef string
			if cmp.Mapping.Product.Image != "" {
				productImage := cmp.Mapping.Product.Image
				imageRef = v[productImage]
				row = append(row, v[productImage])
			} else {
				row = append(row, "")
			}

			// feature_name
			// feature_values
			var features, featuresValues []string
			for _, feature := range cmp.Mapping.Product.Features {
				if v[feature] == "" {
					return nil
				}
				features = append(features, feature)
				featuresValues = append(featuresValues, v[feature])
			}
			row = append(row, strings.Join(features, ","))
			row = append(row, strings.Join(featuresValues, ","))

			// insert feature
			features_pairs, err := createOrFindFeatures(db, features, featuresValues)
			// pp.Println("features_pairs:", features_pairs)

			// minimum_garantie
			row = append(row, "")

			// quantite_unitaire
			if cmp.Mapping.Product.Quantity != "" {
				productQuantity := cmp.Mapping.Product.Quantity
				product.Quantity = v[productQuantity]
				row = append(row, v[productQuantity])
			} else {
				row = append(row, "")
			}

			// format_unitaire
			row = append(row, "")

			// prix_total
			if cmp.Mapping.Product.Price != "" {
				productPrice := cmp.Mapping.Product.Price
				product.Price = v[productPrice]
				row = append(row, v[productPrice])
			} else {
				row = append(row, "")
			}

			// unite_total
			row = append(row, "1")

			// poids_piece
			row = append(row, "1")

			if cmp.Mapping.Product.Ean13 != "" {
				productEan := cmp.Mapping.Product.Ean13
				ean13 := v[productEan]
				if len(ean13) > 12 {
					product.Ean13 = ean13[0:12]
				} else {
					product.Ean13 = ean13
				}
				row = append(row, product.Ean13)
			} else {
				row = append(row, "")
			}

			if cmp.Mapping.Product.Sku != "" {
				productSku := cmp.Mapping.Product.Sku
				product.Sku = v[productSku]
				row = append(row, v[productSku])
			} else {
				row = append(row, "")
			}

			if cmp.Mapping.Product.Mpn != "" {
				productMpn := cmp.Mapping.Product.Mpn
				product.Mpn = v[productMpn]
				row = append(row, v[productMpn])
			} else {
				row = append(row, "")
			}

			productId, err := createOrFindProduct(db, product, imageRef, default_category, features_pairs)
			if err != nil {
				// todo. log as an error
				return err
			}
			pp.Println("productId", productId)

			localImg, err := downloadImageToPs(imageRef, productId)
			log.Warnln("downloadImageToPs.Error", err)
			if err != nil {
				if strings.Contains(err.Error(), "bad content length") {
					log.Warnln("openFileByURL.Fallback")
					err := openFileByURL(imageRef, productId)
					if err != nil {
						log.Warnln("openFileByURL;", err)
						// todo. log as an error
						return err
					}
				}
			} // checkErr("downloadImageToPs.Error", err)
			pp.Println("localImg", localImg)
			// fallback

			/*
				var attributes, attributesValues []string
				for _, attribute := range cmp.Mapping.Product.Attributes {
					if value[attribute] != "" {
						attributes = append(attributes, attribute)
						attributesValues = append(attributesValues, value[attribute])
					}
				}
				row = append(row, strings.Join(attributes, ","))
				row = append(row, strings.Join(attributesValues, ","))
			*/

			if redirectLink != "" {
				err = addToBadger(redirectLink, "ok")
				if err != nil {
					return err
				}
			}

			err = writer.Write(row)
			// err = writer.Write(row)
			checkErr("Cannot write to file", err)
			c++
			return nil

		}(value)

		t.Throttle()

	}

	if t.Err() != nil {
		// Loop through the errors to see the details
		for i, err := range t.Errs() {
			log.Printf("error #%d: %s", i, err)
		}
		log.Fatal(t.Err())
	}

	return cmp, c, nil
}

func getHeaders(fp, sep string) ([]string, error) {
	var headers []string
	// read
	file, err := os.Open(fp)
	if err != nil {
		return headers, err
	}
	defer file.Close()
	r := csv.NewReader(file)
	if sep == "" {
		sep = ";"
	}
	r.Comma = []rune(sep)[0]
	r.LazyQuotes = true
	r.Comment = '#'

	i := 0
	for {
		row, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			if perr, ok := err.(*csv.ParseError); ok && perr.Err == csv.ErrFieldCount {
				continue
			}
			return headers, err
		}
		if i == 0 {
			headers = row
			break
		}
	}

	headers = delete_empty(headers)
	headers = normalize_headers(headers)

	return headers, nil
}

func normalize_line(str string) string {
	if str != "" {
		str = strings.Replace(str, "\"", "", -1)
		str = strings.Replace(str, "(", "", -1)
		str = strings.Replace(str, ")", "", -1)
		t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
		str, _, _ = transform.String(t, str)
	}
	return str
}

func normalize_headers(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			str = strings.Replace(str, "\"", "", -1)
			str = strings.Replace(str, "(", "", -1)
			str = strings.Replace(str, ")", "", -1)
			t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
			str, _, _ = transform.String(t, str)
			r = append(r, strcase.ToSnake(str))
		}
	}
	return r
}

func delete_empty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}

func createOrFindProduct(db *gorm.DB, product *Product, imageRef, default_category string, features_pairs []string) (int, error) {

	if default_category == "" {
		return 0, errors.New("default category is empty")
	}

	id_default_category, err := strconv.Atoi(default_category)
	if err != nil {
		return 0, err
	}

	var productExists psm.Product
	where := &psm.Product{
		IDCategoryDefault: uint32(id_default_category),
		Ean13:             product.Ean13,
		Reference:         product.Reference,
	}
	err = db.Where(where).First(&productExists).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {

		newProduct := &psm.Product{
			IDCategoryDefault: uint32(id_default_category),
			IDShopDefault:     1,
			Ean13:             product.Ean13,
			Mpn:               product.Mpn,
			Reference:         product.Reference,
			Active:            true,
			DateAdd:           time.Now(),
			DateUpd:           time.Now(),
		}

		if product.Quantity != "" {
			quantity, err := strconv.Atoi(product.Quantity)
			if err != nil {
				return 0, err
			}
			newProduct.Quantity = quantity
		}

		if product.Price != "" {
			product.Price = strings.Replace(product.Price, ",", ".", -1)
			product.Price = strings.Replace(product.Price, " EUR", "", -1)
			price, err := strconv.ParseFloat(product.Price, 64)
			if err != nil {
				return 0, err
			}
			newProduct.Price = price
		}

		err = db.Create(&newProduct).Error
		if err != nil {
			return 0, err
		}

		for _, shop := range activeShops {
			for _, activeLang := range activeLangs {

				linkRewrite := slug.Make(product.Name)
				if len(linkRewrite) > 127 {
					linkRewrite = linkRewrite[0:127]
				}

				productLang := &psm.ProductLang{
					IDProduct:        newProduct.IDProduct,
					IDLang:           uint32(activeLang.IDLang),
					IDShop:           uint32(shop.IDShop),
					Description:      stripCtlAndExtFromUnicode(product.Description),
					DescriptionShort: stripCtlAndExtFromUnicode(product.DescriptionShort),
					Name:             stripCtlAndExtFromUnicode(product.Name),
					LinkRewrite:      linkRewrite,
				}
				err := db.Create(&productLang).Error
				if err != nil {
					return 0, err
				}
			}
		}

		for _, shop := range activeShops {
			productShop := &psm.ProductShop{
				IDProduct:         newProduct.IDProduct,
				IDShop:            uint32(shop.IDShop),
				IDCategoryDefault: uint32(id_default_category),
				Active:            true,
				DateAdd:           time.Now(),
				DateUpd:           time.Now(),
			}

			if product.Price != "" {
				product.Price = strings.Replace(product.Price, ",", ".", -1)
				product.Price = strings.Replace(product.Price, " EUR", "", -1)
				price, err := strconv.ParseFloat(product.Price, 64)
				if err != nil {
					return 0, err
				}
				productShop.Price = price
			}

			err := db.Create(&productShop).Error
			if err != nil {
				return 0, err
			}
		}

		categoryProduct := &psm.CategoryProduct{
			IDCategory: uint32(id_default_category),
			IDProduct:  newProduct.IDProduct,
		}

		err = db.Create(&categoryProduct).Error
		if err != nil {
			return 0, err
		}

		for _, features_pair := range features_pairs {
			f := strings.Split(features_pair, ";")
			featureId, featureValueId := f[0], f[1]

			if featureId == "" || featureValueId == "" {
				continue
			}
			intFeatureId, err := strconv.Atoi(featureId)
			if err != nil {
				return 0, err
			}

			intFeatureValueId, err := strconv.Atoi(featureValueId)
			if err != nil {
				return 0, err
			}

			featureProduct := &psm.FeatureProduct{
				IDFeature:      uint32(intFeatureId),
				IDProduct:      newProduct.IDProduct,
				IDFeatureValue: uint32(intFeatureValueId),
			}
			err = db.Create(&featureProduct).Error
			if err != nil {
				return 0, err
			}
		}
		return int(newProduct.IDProduct), nil
	}
	return int(productExists.IDProduct), nil
}

func createOrFindFeatures(db *gorm.DB, featureNames, featureValues []string) ([]string, error) {

	var featureValuePairs []string
	for i, featureName := range featureNames {

		var featureId, featureValueId uint32

		var featureExists psm.Feature
		err := db.Joins("JOIN eg_feature_lang fl ON fl.id_feature = `eg_feature`.id_feature AND fl.name = ?", featureName).First(&featureExists).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {

			feature := &psm.Feature{
				Position: 0,
			}
			err := db.Create(&feature).Error
			if err != nil {
				return nil, err
			}
			for _, activeLang := range activeLangs {
				featureLang := &psm.FeatureLang{
					IDFeature: feature.IDFeature,
					IDLang:    uint32(activeLang.IDLang),
					Name:      featureName,
				}
				err := db.Create(&featureLang).Error
				if err != nil {
					return nil, err
				}
			}
			for _, shop := range activeShops {
				featureShop := &psm.FeatureShop{
					IDFeature: feature.IDFeature,
					IDShop:    uint32(shop.IDShop),
				}
				err := db.Create(&featureShop).Error
				if err != nil {
					return nil, err
				}
			}
			featureId = feature.IDFeature
		} else {
			featureId = featureExists.IDFeature
		}

		var featureValueExists psm.FeatureValueLang

		/*
			SELECT fv.`id_feature_value`,fv.`id_feature`,fv.`custom`
			FROM `eg_feature_value_lang` fvl
			LEFT JOIN eg_feature_value fv ON fv.id_feature_value = fvl.id_feature_value
			WHERE fvl.value = 'Rhum & Cachaca'
		*/

		err = db.Joins("LEFT JOIN eg_feature_value fv ON fv.id_feature_value = `eg_feature_value_lang`.id_feature_value").Where("`eg_feature_value_lang`.value = ?", featureValues[i]).First(&featureValueExists).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			featureValue := &psm.FeatureValue{
				IDFeature: featureId,
			}
			err = db.Create(&featureValue).Error
			if err != nil {
				return nil, err
			}

			for _, activeLang := range activeLangs {
				featureValueLang := &psm.FeatureValueLang{
					IDFeatureValue: featureValue.IDFeatureValue,
					IDLang:         uint32(activeLang.IDLang),
					Value:          featureValues[i],
				}
				err := db.Create(&featureValueLang).Error
				if err != nil {
					return nil, err
				}
			}
			featureValueId = featureValue.IDFeatureValue
		} else {
			featureValueId = featureValueExists.IDFeatureValue
		}
		featureValuePairs = append(featureValuePairs, fmt.Sprintf("%d;%d", featureId, featureValueId))

	}

	return featureValuePairs, nil
}

func createNestedCategoryTree(db *gorm.DB, cpm *CatalogMap, breadcrumbSlice []string) (string, string, error) {

	var category_ids []string

	parent_level := 2
	// shop name > categorie tree
	rootCategory := &psm.Category{
		IDParent:       2,
		IDShopDefault:  1,
		LevelDepth:     uint8(parent_level),
		Active:         true,
		IsRootCategory: true,
	}

	var id_category_root int
	var rootCategoryExists psm.Category
	err := db.Joins("JOIN eg_category_lang cl ON cl.id_category = `eg_category`.id_category AND cl.name = ?", cpm.Name).Where(rootCategory).First(&rootCategoryExists).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {

		//pp.Println("rootCategory.IDCategory", rootCategory.IDCategory)

		err := db.Create(&rootCategory).Error
		if err != nil {
			return "", "", err
		}
		//pp.Println("rootCategory.IDCategory", rootCategory.IDCategory)
		id_category_root = int(rootCategory.IDCategory)

		// category_lang
		for _, activeLang := range activeLangs {

			name := cpm.Name
			if len(name) > 127 {
				name = cpm.Name[0:127]
			}
			linkRewrite := slug.Make(cpm.Name)
			if len(linkRewrite) > 127 {
				linkRewrite = linkRewrite[0:127]
			}
			rootCategoryLang := &psm.CategoryLang{
				IDCategory:  rootCategory.IDCategory,
				IDShop:      1,
				IDLang:      uint32(activeLang.IDLang),
				Name:        cpm.Name,
				LinkRewrite: linkRewrite,
			}
			err := db.Create(&rootCategoryLang).Error
			if err != nil {
				return "", "", err
			}
		}

		// category_groups
		for _, group := range activeGroups {
			rootCategoryGroup := &psm.CategoryGroup{
				IDCategory: rootCategory.IDCategory,
				IDGroup:    group.IDGroup,
			}
			err := db.Create(&rootCategoryGroup).Error
			if err != nil {
				return "", "", err
			}
		}

		// category_shop
		for _, shop := range activeShops {
			rootCategoryShop := &psm.CategoryShop{
				IDCategory: int(rootCategory.IDCategory),
				IDShop:     int(shop.IDShop),
			}
			err := db.Create(&rootCategoryShop).Error
			if err != nil {
				return "", "", err
			}
		}
	} else {
		id_category_root = int(rootCategoryExists.IDCategory)
	}

	category_ids = append(category_ids, fmt.Sprintf("%d", id_category_root))

	for _, shop := range activeShops {
		var parentId uint32
		for _, category := range breadcrumbSlice {
			if parentId == 0 {
				parentId = uint32(id_category_root)
			}
			//pp.Println("category", category)
			parentId, category_ids, err = createCategoryTree(db, category, parentId, category_ids, shop.IDShop, parent_level+1)
		}
	}

	return strings.Join(category_ids, ";"), category_ids[len(category_ids)-1], nil
}

func createCategoryTree(db *gorm.DB, name string, id_parent uint32, category_ids []string, id_shop int, level int) (uint32, []string, error) {
	// shop name > categorie tree
	category := &psm.Category{
		IDParent:       id_parent,
		IDShopDefault:  uint32(id_shop),
		LevelDepth:     uint8(level),
		Active:         true,
		IsRootCategory: false,
	}

	var id_category_parent int
	var parentCategoryExists psm.Category
	err := db.Joins("JOIN eg_category_lang cl ON cl.id_category = `eg_category`.id_category AND cl.name = ? AND `eg_category`.id_parent = ?", name, id_parent).Where(category).First(&parentCategoryExists).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {

		err := db.Create(&category).Error
		if err != nil {
			return 0, category_ids, err
		}
		// pp.Println("category.IDCategory", category.IDCategory)

		id_category_parent = int(category.IDCategory)

		// category_lang
		for _, activeLang := range activeLangs {

			if len(name) > 127 {
				name = name[0:127]
			}

			linkRewrite := slug.Make(name)
			if len(linkRewrite) > 127 {
				linkRewrite = linkRewrite[0:127]
			}

			categoryLang := &psm.CategoryLang{
				IDCategory:  category.IDCategory,
				IDShop:      uint32(id_shop),
				IDLang:      uint32(activeLang.IDLang),
				Name:        name,
				LinkRewrite: linkRewrite,
			}
			err := db.Create(&categoryLang).Error
			if err != nil {
				return 0, category_ids, err
			}
		}

		// category_groups
		for _, group := range activeGroups {
			categoryGroup := &psm.CategoryGroup{
				IDCategory: category.IDCategory,
				IDGroup:    group.IDGroup,
			}
			err := db.Create(&categoryGroup).Error
			if err != nil {
				return 0, category_ids, err
			}
		}

		// category_shop
		for _, shop := range activeShops {
			categoryShop := &psm.CategoryShop{
				IDCategory: int(category.IDCategory),
				IDShop:     shop.IDShop,
			}
			err := db.Create(&categoryShop).Error
			if err != nil {
				return 0, category_ids, err
			}
		}
	} else {
		id_category_parent = int(parentCategoryExists.IDCategory)
	}

	category_ids = append(category_ids, fmt.Sprintf("%d", id_category_parent))

	return category.IDCategory, category_ids, nil
}

func downloadImageToPs(remoteUrl string, productId int) (string, error) {

	image := &psm.Image{
		IDProduct: uint32(productId),
		Cover:     true,
	}

	var imageExists psm.Image
	var imageId int
	err := db.Where(image).First(&imageExists).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {

		err := db.Create(&image).Error
		if err != nil {
			return "", err
		}

		for _, activeLang := range activeLangs {
			imageLang := &psm.ImageLang{
				IDImage: image.IDImage,
				IDLang:  uint32(activeLang.IDLang),
				Legend:  "",
			}
			err := db.Create(&imageLang).Error
			if err != nil {
				return "", err
			}
		}

		for _, shop := range activeShops {
			imageShop := &psm.ImageShop{
				IDProduct: uint32(productId),
				IDShop:    uint32(shop.IDShop),
				IDImage:   image.IDImage,
				Cover:     true,
			}
			err := db.Create(&imageShop).Error
			if err != nil {
				return "", err
			}
		}
		imageId = int(image.IDImage)
	} else {
		imageId = int(imageExists.IDImage)
	}

	localPath := buildFolderForImage(filepath.Join(psDir, "img", "p"), int(imageId))
	localImg := fmt.Sprintf("%s/%d.jpg", localPath, imageId)

	client := grab.NewClient()

	// proxyURL

	tr := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
		DialContext: (&net.Dialer{
			Timeout: 240 * time.Second,
			// KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       240 * time.Second,
		TLSHandshakeTimeout:   240 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		DisableKeepAlives:     true,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	client.UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.97 Safari/537.36"
	client.HTTPClient = &http.Client{
		Timeout:   240 * time.Second,
		Transport: tr,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return nil
		},
	}

	req, err := grab.NewRequest(localImg, remoteUrl)
	if err != nil {
		return "", err
	}

	// start download
	fmt.Printf("Downloading %v...\n", req.URL())
	resp := client.Do(req)
	if resp.HTTPResponse == nil {
		return "", err
	}
	fmt.Printf("  %v\n", resp.HTTPResponse.Status)

	// start UI loop
	t := time.NewTicker(500 * time.Millisecond)
	defer t.Stop()

LoopRel:
	for {
		select {
		case <-t.C:
			fmt.Printf("  transferred %v / %v bytes (%.2f%%)\n",
				resp.BytesComplete(),
				resp.Size,
				100*resp.Progress())

		case <-resp.Done:
			// download is complete
			break LoopRel
		}
	}

	// check for errors
	if err := resp.Err(); err != nil {
		return "", err
	}

	fmt.Printf("Download saved to %v \n", resp.Filename)
	localImg = resp.Filename

	return localImg, nil
}

func openFileByURL(rawURL string, productId int) error {

	image := &psm.Image{
		IDProduct: uint32(productId),
		Cover:     true,
	}

	var imageExists psm.Image
	var imageId int
	err := db.Where(image).First(&imageExists).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {

		err := db.Create(&image).Error
		if err != nil {
			return err
		}

		for _, activeLang := range activeLangs {
			imageLang := &psm.ImageLang{
				IDImage: image.IDImage,
				IDLang:  uint32(activeLang.IDLang),
				Legend:  "",
			}
			err := db.Create(&imageLang).Error
			if err != nil {
				return err
			}
		}

		for _, shop := range activeShops {
			imageShop := &psm.ImageShop{
				IDProduct: uint32(productId),
				IDShop:    uint32(shop.IDShop),
				IDImage:   image.IDImage,
				Cover:     true,
			}
			err := db.Create(&imageShop).Error
			if err != nil {
				return err
			}
		}
		imageId = int(image.IDImage)
	} else {
		imageId = int(imageExists.IDImage)
	}

	localPath := buildFolderForImage(filepath.Join(psDir, "img", "p"), int(imageId))
	localImg := fmt.Sprintf("%s/%d.jpg", localPath, imageId)

	if _, err := url.Parse(rawURL); err != nil {

		return err
	} else {

		file, err := os.Create(localImg)
		if err != nil {
			return err
		}

		tr := &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		}

		check := http.Client{
			Transport: tr,
			CheckRedirect: func(r *http.Request, via []*http.Request) error {
				r.URL.Opaque = r.URL.Path
				return nil
			},
		}
		resp, err := check.Get(rawURL) // add a filter to check redirect
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		fmt.Printf("----> Downloaded %v to %s\n", rawURL, localImg)

		_, err = io.Copy(file, resp.Body)
		if err != nil {
			return err
		}

		return nil
	}
}

func compress(data []byte) ([]byte, error) {
	return snappy.Encode([]byte{}, data), nil
}

func decompress(data []byte) ([]byte, error) {
	return snappy.Decode([]byte{}, data)
}

func getFromBadger(key string) (resp []byte, ok bool) {
	err := kv.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}
		err = item.Value(func(val []byte) error {
			// fmt.Printf("The answer is: %s\n", val)
			return nil
		})
		if err != nil {
			return err
		}
		return nil
	})
	return resp, err == nil
}

func addToBadger(key, value string) error {
	err := kv.Update(func(txn *badger.Txn) error {
		log.Println("indexing: ", key)
		cnt, err := compress([]byte(value))
		if err != nil {
			return err
		}
		err = txn.Set([]byte(key), cnt)
		return err
	})
	return err
}

// Advanced Unicode normalization and filtering,
// see http://blog.golang.org/normalization and
// http://godoc.org/golang.org/x/text/unicode/norm for more
// details.
func stripCtlAndExtFromUnicode(str string) string {
	isOk := func(r rune) bool {
		return r < 32 || r >= 127
	}
	// The isOk filter is such that there is no need to chain to norm.NFC
	t := transform.Chain(norm.NFKD, transform.RemoveFunc(isOk))
	// This Transformer could also trivially be applied as an io.Reader
	// or io.Writer filter to automatically do such filtering when reading
	// or writing data anywhere.
	str, _, _ = transform.String(t, str)
	return str
}

func shuffleMap(slc []map[string]string) []map[string]string {
	for i := 1; i < len(slc); i++ {
		r := rand.Intn(i + 1)
		if i != r {
			slc[r], slc[i] = slc[i], slc[r]
		}
	}
	return slc
}

// shuffleXml
func shuffleXml(slc []*etree.Element) []*etree.Element {
	for i := 1; i < len(slc); i++ {
		r := rand.Intn(i + 1)
		if i != r {
			slc[r], slc[i] = slc[i], slc[r]
		}
	}
	return slc
}
