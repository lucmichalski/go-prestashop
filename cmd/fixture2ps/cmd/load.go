package cmd

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/emvi/iso-639-1"
	"github.com/fatih/structs"
	mysql_driver "github.com/go-sql-driver/mysql"
	"github.com/goombaio/namegenerator"
	"github.com/h2non/bimg"
	"github.com/iancoleman/strcase"
	"github.com/k0kubun/pp"
	"github.com/karrick/godirwalk"
	"github.com/spf13/cobra"
	"golang.org/x/net/html/charset"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/lucmichalski/go-prestashop/internal/fixtures"
	"github.com/lucmichalski/go-prestashop/internal/fixtures/langs"
	"github.com/lucmichalski/go-prestashop/internal/models"
	"github.com/lucmichalski/go-prestashop/internal/struct2map"
)

var (
	psDir         string
	outputDir     string
	workDir       string
	dryRun        bool
	db            *gorm.DB
	dbName        string
	dbUser        string
	dbPass        string
	dbHost        string
	dbPort        string
	dbTablePrefix string
	dbMigrate     bool
	dbDrop        bool
	dbEnv         bool
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

var ImportCmd = &cobra.Command{
	Use:     "load",
	Aliases: []string{"l"},
	Short:   "load prestashop-shop-creator xml to a mysql database.",
	Long:    "load prestashop-shop-creator xml to a mysql database",
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

			err = prepareTables(db, models.Tables...)
			if err != nil {
				log.Fatal(err)
			}
		}

		nameGenerator := namegenerator.NewNameGenerator(seed)

		if !noFixtures {

			// load fixtures
			for _, fixture := range fixtures.Fixtures {

				n := structs.Name(fixture)
				osPathname := filepath.Join(workDir, "data", fmt.Sprintf("%s.xml", strings.ToLower(strcase.ToSnake(n))))
				if options.debug {
					pp.Println("osPathname:", osPathname)
				}

				file, err := os.Open(osPathname)
				if err != nil {
					log.Warn(err)
					continue
				}

				reader := bufio.NewReader(file)
				dec := xml.NewDecoder(reader)
				dec.CharsetReader = charset.NewReaderLabel
				dec.Strict = false
				if err := dec.Decode(&fixture); err != nil {
					log.Fatal(err)
				}

				csvDirPath := filepath.Join(outputDir, "data")
				csvBaseName := strings.ToLower(strcase.ToSnake(n))
				csvFileName := fmt.Sprintf("%s.csv", csvBaseName)
				csvFilePath := filepath.Join(csvDirPath, csvFileName)

				csvFile, err := os.Create(csvFilePath)
				if err != nil {
					log.Fatal("Unable to open output")
				}
				defer csvFile.Close()

				entities := make([]interface{}, 0)
				switch t := fixture.(type) {
				case *fixtures.Guest:
					m := structs.Map(fixture.(*fixtures.Guest).Entities)
					entities = append(entities, m["Guest"].([]interface{})...)
				case *fixtures.OrderMessage:
					m := structs.Map(fixture.(*fixtures.OrderMessage).Entities)
					entities = append(entities, m["OrderMessage"].([]interface{})...)
				case *fixtures.RangeWeight:
					m := structs.Map(fixture.(*fixtures.RangeWeight).Entities)
					entities = append(entities, m["RangeWeight"].([]interface{})...)
				case *fixtures.Profile:
					m := structs.Map(fixture.(*fixtures.Profile).Entities)
					entities = append(entities, m["Profile"].([]interface{})...)
				case *fixtures.Image:
					m := structs.Map(fixture.(*fixtures.Image).Entities)
					entities = append(entities, m["Image"].([]interface{})...)
				case *fixtures.OrderHistory:
					m := structs.Map(fixture.(*fixtures.OrderHistory).Entities)
					entities = append(entities, m["OrderHistory"].([]interface{})...)
				case *fixtures.ProductSupplier:
					m := structs.Map(fixture.(*fixtures.ProductSupplier).Entities)
					entities = append(entities, m["ProductSupplier"].([]interface{})...)
				case *fixtures.CategoryGroup:
					m := structs.Map(fixture.(*fixtures.CategoryGroup).Entities)
					entities = append(entities, m["CategoryGroup"].([]interface{})...)
				case *fixtures.Carrier:
					m := structs.Map(fixture.(*fixtures.Carrier).Entities)
					entities = append(entities, m["Carrier"].([]interface{})...)
				case *fixtures.Product:
					m := structs.Map(fixture.(*fixtures.Product).Entities)
					entities = append(entities, m["Product"].([]interface{})...)
				case *fixtures.Cart:
					m := structs.Map(fixture.(*fixtures.Cart).Entities)
					entities = append(entities, m["Cart"].([]interface{})...)
				case *fixtures.CategoryProduct:
					m := structs.Map(fixture.(*fixtures.CategoryProduct).Entities)
					entities = append(entities, m["CategoryProduct"].([]interface{})...)
				case *fixtures.CarrierGroup:
					m := structs.Map(fixture.(*fixtures.CarrierGroup).Entities)
					entities = append(entities, m["CarrierGroup"].([]interface{})...)
				case *fixtures.Orders:
					m := structs.Map(fixture.(*fixtures.Orders).Entities)
					entities = append(entities, m["Orders"].([]interface{})...)
				case *fixtures.RangePrice:
					m := structs.Map(fixture.(*fixtures.RangePrice).Entities)
					entities = append(entities, m["RangePrice"].([]interface{})...)
				case *fixtures.Alias:
					m := structs.Map(fixture.(*fixtures.Alias).Entities)
					entities = append(entities, m["Alias"].([]interface{})...)
				case *fixtures.StockAvailable:
					m := structs.Map(fixture.(*fixtures.StockAvailable).Entities)
					entities = append(entities, m["StockAvailable"].([]interface{})...)
				case *fixtures.Manufacturer:
					m := structs.Map(fixture.(*fixtures.Manufacturer).Entities)
					entities = append(entities, m["Manufacturer"].([]interface{})...)
				case *fixtures.FeatureValue:
					m := structs.Map(fixture.(*fixtures.FeatureValue).Entities)
					entities = append(entities, m["FeatureValue"].([]interface{})...)
				case *fixtures.ProductAttribute:
					m := structs.Map(fixture.(*fixtures.ProductAttribute).Entities)
					entities = append(entities, m["ProductAttribute"].([]interface{})...)
				case *fixtures.AttributeGroup:
					m := structs.Map(fixture.(*fixtures.AttributeGroup).Entities)
					entities = append(entities, m["AttributeGroup"].([]interface{})...)
				case *fixtures.Store:
					m := structs.Map(fixture.(*fixtures.Store).Entities)
					entities = append(entities, m["Store"].([]interface{})...)
				case *fixtures.ProductAttributeImage:
					m := structs.Map(fixture.(*fixtures.ProductAttributeImage).Entities)
					entities = append(entities, m["ProductAttributeImage"].([]interface{})...)
				case *fixtures.Category:
					m := structs.Map(fixture.(*fixtures.Category).Entities)
					entities = append(entities, m["Category"].([]interface{})...)
				case *fixtures.Connections:
					m := structs.Map(fixture.(*fixtures.Connections).Entities)
					entities = append(entities, m["Connections"].([]interface{})...)
				case *fixtures.Delivery:
					m := structs.Map(fixture.(*fixtures.Delivery).Entities)
					entities = append(entities, m["Delivery"].([]interface{})...)
				case *fixtures.Feature:
					m := structs.Map(fixture.(*fixtures.Feature).Entities)
					entities = append(entities, m["Feature"].([]interface{})...)
				case *fixtures.ProductAttributeCombination:
					m := structs.Map(fixture.(*fixtures.ProductAttributeCombination).Entities)
					entities = append(entities, m["ProductAttributeCombination"].([]interface{})...)
				case *fixtures.Attribute:
					m := structs.Map(fixture.(*fixtures.Attribute).Entities)
					entities = append(entities, m["Attribute"].([]interface{})...)
				case *fixtures.FeatureProduct:
					m := structs.Map(fixture.(*fixtures.FeatureProduct).Entities)
					entities = append(entities, m["FeatureProduct"].([]interface{})...)
				case *fixtures.OrderCarrier:
					m := structs.Map(fixture.(*fixtures.OrderCarrier).Entities)
					entities = append(entities, m["OrderCarrier"].([]interface{})...)
				case *fixtures.Customer:
					m := structs.Map(fixture.(*fixtures.Customer).Entities)
					entities = append(entities, m["Customer"].([]interface{})...)
				case *fixtures.OrderDetail:
					m := structs.Map(fixture.(*fixtures.OrderDetail).Entities)
					entities = append(entities, m["OrderDetail"].([]interface{})...)
				case *fixtures.Supplier:
					m := structs.Map(fixture.(*fixtures.Supplier).Entities)
					entities = append(entities, m["Supplier"].([]interface{})...)
				case *fixtures.Address:
					m := structs.Map(fixture.(*fixtures.Address).Entities)
					entities = append(entities, m["Address"].([]interface{})...)
				case *fixtures.SpecificPrice:
					m := structs.Map(fixture.(*fixtures.SpecificPrice).Entities)
					entities = append(entities, m["SpecificPrice"].([]interface{})...)
				case *fixtures.CarrierZone:
					m := structs.Map(fixture.(*fixtures.CarrierZone).Entities)
					entities = append(entities, m["CarrierZone"].([]interface{})...)
				default:
					fmt.Printf("case %T:\n", t) // %T prints whatever type t has
					log.Fatal("missing")
				}

				columnsMap := make(map[string]bool, 0)
				rows := make([][]string, 0)
				for _, entity := range entities {
					var row []string
					switch t := entity.(type) {
					case map[string]interface{}:
						keys := make([]string, 0, len(entity.(map[string]interface{})))
						for k := range entity.(map[string]interface{}) {
							if strings.ToLower(k) != "text" && strings.ToLower(k) != "id" {
								keys = append(keys, k)
							}
						}
						sort.Strings(keys)
						for _, k := range keys {
							if strings.ToLower(k) != "text" && strings.ToLower(k) != "id" {
								row = append(row, fmt.Sprintf("%v", entity.(map[string]interface{})[k]))
								columnsMap[k] = true
							}
						}
					default:
						fmt.Printf("case %T:\n", t) // %T prints whatever type t has
					}
					rows = append(rows, row)
				}

				var columns []string
				for column := range columnsMap {
					column = strings.Replace(column, "ID", "Id", 1)
					column = strcase.ToSnake(column)
					column = strings.Replace(column, "_1", "1", 1)
					column = strings.Replace(column, "_2", "2", 1)
					columns = append(columns, column)
				}

				buf, err := encodeToCsv(columns, rows)
				csvFile.Write(buf)

				if !dryRun {
					// load data into file
					err := loadData(db, csvFilePath, csvBaseName, columns...)
					if err != nil {
						log.Fatal(err)
					}
				}
			}

			// load fixtures per language
			for _, language := range languages {

				// create or get language_id
				lang, err := findOrCreateLangByIso(db, language)
				if err != nil {
					log.Fatal(err)
				}

				pp.Println("lang:", lang)

				for _, model := range langs.Fixtures {

					n := structs.Name(model)
					osPathname := filepath.Join(workDir, "langs", language, "data", fmt.Sprintf("%s.xml", strings.Replace(strings.ToLower(strcase.ToSnake(n)), "_lang", "", 1)))
					if options.debug {
						pp.Println("osPathname:", osPathname)
					}

					file, err := os.Open(osPathname)
					if err != nil {
						log.Warn(err)
						continue
					}

					reader := bufio.NewReader(file)
					dec := xml.NewDecoder(reader)
					dec.CharsetReader = charset.NewReaderLabel
					dec.Strict = false
					if err := dec.Decode(&model); err != nil {
						log.Fatal(err)
					}

					csvDirPath := filepath.Join(outputDir, "langs", language, "data")
					csvBaseName := strings.ToLower(strcase.ToSnake(n))
					csvFileName := fmt.Sprintf("%s.csv", csvBaseName)
					csvFilePath := filepath.Join(csvDirPath, csvFileName)

					// todo. create base dir
					csvFile, err := os.Create(csvFilePath)
					if err != nil {
						log.Fatal("Unable to open output")
					}
					defer csvFile.Close()

					entities := make([]interface{}, 0)
					switch t := model.(type) {
					case *langs.FeatureLang:
						for _, entity := range model.(*langs.FeatureLang).Feature {
							m, err := struct2map.Decode(&entity)
							if err != nil {
								log.Fatal(err)
							}
							entities = append(entities, m)
						}
					case *langs.SupplierLang:
						for _, entity := range model.(*langs.SupplierLang).Supplier {
							m, err := struct2map.Decode(&entity)
							if err != nil {
								log.Fatal(err)
							}
							entities = append(entities, m)
						}
					case *langs.StoreLang:
						for _, entity := range model.(*langs.StoreLang).Store {
							m, err := struct2map.Decode(&entity)
							if err != nil {
								log.Fatal(err)
							}
							entities = append(entities, m)
						}
					case *langs.AttributeLang:
						for _, entity := range model.(*langs.AttributeLang).Attribute {
							m, err := struct2map.Decode(&entity)
							if err != nil {
								log.Fatal(err)
							}
							entities = append(entities, m)
						}
					case *langs.FeatureValueLang:
						for _, entity := range model.(*langs.FeatureValueLang).FeatureValue {
							m, err := struct2map.Decode(&entity)
							if err != nil {
								log.Fatal(err)
							}
							entities = append(entities, m)
						}
					case *langs.AttributeGroupLang:
						for _, entity := range model.(*langs.AttributeGroupLang).AttributeGroup {
							m, err := struct2map.Decode(&entity)
							if err != nil {
								log.Fatal(err)
							}
							entities = append(entities, m)
						}
					case *langs.OrderMessageLang:
						for _, entity := range model.(*langs.OrderMessageLang).OrderMessage {
							m, err := struct2map.Decode(&entity)
							if err != nil {
								log.Fatal(err)
							}
							entities = append(entities, m)
						}
					case *langs.CategoryLang:
						for _, entity := range model.(*langs.CategoryLang).Category {
							m, err := struct2map.Decode(&entity)
							if err != nil {
								log.Fatal(err)
							}
							entities = append(entities, m)
						}
					case *langs.ProfileLang:
						for _, entity := range model.(*langs.ProfileLang).Profile {
							m, err := struct2map.Decode(&entity)
							if err != nil {
								log.Fatal(err)
							}
							entities = append(entities, m)
						}
					case *langs.ProductLang:
						for _, entity := range model.(*langs.ProductLang).Product {
							m, err := struct2map.Decode(&entity)
							if err != nil {
								log.Fatal(err)
							}
							entities = append(entities, m)
						}
					case *langs.ManufacturerLang:
						for _, entity := range model.(*langs.ManufacturerLang).Manufacturer {
							m, err := struct2map.Decode(&entity)
							if err != nil {
								log.Fatal(err)
							}
							entities = append(entities, m)
						}
					case *langs.CarrierLang:
						for _, entity := range model.(*langs.CarrierLang).Carrier {
							m, err := struct2map.Decode(&entity)
							if err != nil {
								log.Fatal(err)
							}
							entities = append(entities, m)
						}
					case *langs.ImageLang:
						for _, entity := range model.(*langs.ImageLang).Image {
							m, err := struct2map.Decode(&entity)
							if err != nil {
								log.Fatal(err)
							}
							entities = append(entities, m)
						}
					default:
						fmt.Printf("case %T:\n", t) // %T prints whatever type t has
						log.Fatal("missing")
					}

					columnsMap := make(map[string]bool, 0)
					rows := make([][]string, 0)
					fakeID := 1
					for _, entity := range entities {
						var row []string
						switch t := entity.(type) {
						case map[string]interface{}:

							keys := make([]string, 0, len(entity.(map[string]interface{})))
							for k := range entity.(map[string]interface{}) {
								keys = append(keys, k)
							}
							sort.Strings(keys)
							for _, k := range keys {
								v := entity.(map[string]interface{})[k]
								if k == "id_shop" {
									v = 1
								}
								if k == "id_lang" {
									v = lang.IDLang
								}
								if k == "id_category" {
									v = fakeID
									fakeID++
								}
								if k == "id_parent" {
									v = 2
								}
								if k == "name" && fmt.Sprintf("%v", v) == "" {
									v = nameGenerator.Generate()
								}
								if k == "reference" && fmt.Sprintf("%v", v) == "" {
									v = faker.Username()
								}
								if fmt.Sprintf("%v", v) == "0000-00-00" {
									v = time.Now().Format("YYYY-mm-dd")
								}
								if fmt.Sprintf("%v", v) == "0000-00-00 00:00:00" {
									v = time.Now().Format("YYYY-mm-dd hh:nn:ss")
								}
								row = append(row, fmt.Sprintf("%v", v))
								columnsMap[k] = true
							}
						default:
							fmt.Printf("case %T:\n", t) // %T prints whatever type t has
						}
						rows = append(rows, row)
					}

					////////////////////////
					// pre processing

					// csvBaseName
					// if csvBaseName == "category_lang" {
					// }

					var columns []string
					for column := range columnsMap {
						column = strings.Replace(column, "ID", "Id", 1)
						column = strcase.ToSnake(column)
						column = strings.Replace(column, "_1", "1", 1)
						column = strings.Replace(column, "_2", "2", 1)
						columns = append(columns, column)
					}
					sort.Strings(columns)

					buf, err := encodeToCsv(columns, rows)
					csvFile.Write(buf)

					if !dryRun {
						// load data into file
						err := loadDataLang(db, csvFilePath, csvBaseName, lang.IDLang, columns...)
						if err != nil {
							log.Fatal(err)
						}
					}

					////////////////////////
					// post processing

				}
			}

			// todo. find faster alternative to rnad() function
			// ref. - https://stackoverflow.com/questions/1823306/mysql-alternatives-to-order-by-rand
			//      - https://stackoverflow.com/questions/4329396/mysql-select-10-random-rows-from-600k-rows-fast
			// SELECT id_address FROM eg_address AS r1 JOIN (SELECT CEIL(RAND() * (SELECT MAX(id_address) FROM eg_address)) AS id) AS r2 WHERE r1.id_address >= r2.id ORDER BY r1.id_address ASC LIMIT 1
			// fix missing relationships

			// manufacturer
			// UPDATE eg_product SET id_manufacturer=(SELECT id_manufacturer FROM eg_manufacturer WHERE active=1 ORDER BY RAND() LIMIT 1);
			err := db.Debug().Exec("UPDATE " + dbTablePrefix + "product SET id_manufacturer=(SELECT id_manufacturer FROM eg_manufacturer WHERE active=1 ORDER BY RAND() LIMIT 1)").Error
			if err != nil {
				log.Fatal(err)
			}

			// supplier
			// UPDATE eg_product SET id_supplier=(SELECT id_supplier FROM eg_supplier WHERE active=1 ORDER BY RAND() LIMIT 1);
			err = db.Debug().Exec("UPDATE " + dbTablePrefix + "product SET id_supplier=(SELECT id_supplier FROM eg_supplier WHERE active=1 ORDER BY RAND() LIMIT 1)").Error
			if err != nil {
				log.Fatal(err)
			}

			// *** eg_product_shop ***
			err = db.Debug().Exec("INSERT IGNORE INTO " + dbTablePrefix + "product_shop ( id_product, id_shop, id_category_default, id_tax_rules_group, on_sale, online_only, ecotax, minimal_quantity, low_stock_threshold, low_stock_alert, price, wholesale_price, unity, unit_price_ratio, additional_shipping_cost, customizable, uploadable_files, text_fields, active, redirect_type, id_type_redirected, available_for_order, available_date, show_condition, `condition`, show_price, indexed, visibility, cache_default_attribute, advanced_stock_management, date_add, date_upd, pack_stock_type) SELECT id_product, id_shop_default, id_category_default, id_tax_rules_group, on_sale, online_only, ecotax, minimal_quantity, low_stock_threshold, low_stock_alert, price, wholesale_price, unity, unit_price_ratio, additional_shipping_cost, customizable, uploadable_files, text_fields, active, redirect_type, id_type_redirected, available_for_order, available_date, show_condition, `condition`, show_price, indexed, visibility, cache_default_attribute, advanced_stock_management, date_add, date_upd, pack_stock_type FROM " + dbTablePrefix + "product").Error
			if err != nil {
				log.Fatal(err)
			}

			// eg_orders
			err = db.Debug().Exec("UPDATE " + dbTablePrefix + "orders SET id_lang=1, id_carrier=1, date_add=NOW(), valid=1;").Error
			if err != nil {
				log.Fatal(err)
			}

			// NOT WORKING as you need as much addresses as you have orders but keep it for later ^^
			// UPDATE eg_orders o INNER JOIN eg_address b ON b.id_customer = o.id_customer SET o.id_address_delivery = b.id_address;
			err = db.Debug().Exec("UPDATE " + dbTablePrefix + "orders o INNER JOIN eg_address b ON b.id_customer = o.id_customer SET o.id_address_delivery = b.id_address;").Error
			if err != nil {
				log.Fatal(err)
			}

			err = db.Debug().Exec("UPDATE " + dbTablePrefix + "orders o INNER JOIN eg_address b ON b.id_customer = o.id_customer SET o.id_address_invoice = b.id_address;").Error
			if err != nil {
				log.Fatal(err)
			}

			randomAdressSql := `SELECT id_address FROM eg_address AS r1 JOIN (SELECT CEIL(RAND() * (SELECT MAX(id_address) FROM eg_address)) AS id) AS r2 WHERE r1.id_address >= r2.id ORDER BY r1.id_address ASC LIMIT 1`

			// UPDATE eg_orders o SET id_address_delivery=(SELECT id_address FROM eg_address a WHERE o.id_customer=a.id_customer ORDER BY RAND() LIMIT 1);
			err = db.Debug().Exec("UPDATE " + dbTablePrefix + "orders o SET id_address_delivery=(" + randomAdressSql + ") WHERE id_address_delivery=0").Error
			if err != nil {
				log.Fatal(err)
			}

			// UPDATE eg_orders o SET id_address_invoice=(SELECT id_address FROM eg_address a WHERE o.id_customer=a.id_customer ORDER BY RAND() LIMIT 1);
			err = db.Debug().Exec("UPDATE " + dbTablePrefix + "orders o SET id_address_invoice=(" + randomAdressSql + ") WHERE id_address_invoice=0").Error
			if err != nil {
				log.Fatal(err)
			}

			// for now, we deal with only one level as we have some troubles to rebuild the nested category tree
			err = db.Debug().Exec("UPDATE " + dbTablePrefix + "category SET id_parent=2, level_depth=1 WHERE id_category>2;").Error
			if err != nil {
				log.Fatal(err)
			}

			// INSERT IGNORE INTO eg_category_shop ( id_category, id_shop, position) SELECT id_category, 1, 0 FROM eg_category;
			err = db.Debug().Exec("INSERT IGNORE INTO " + dbTablePrefix + "category_shop ( id_category, id_shop, position) SELECT id_category, 1, 0 FROM " + dbTablePrefix + "category").Error
			if err != nil {
				log.Fatal(err)
			}

			// images
			// INSERT IGNORE INTO eg_image_shop ( id_image, id_product, id_shop) SELECT id_image, id_product, 1, 0 FROM eg_image;
			err = db.Debug().Exec("INSERT IGNORE INTO " + dbTablePrefix + "image_shop ( id_image, id_product, id_shop) SELECT id_image, id_product, 1 FROM " + dbTablePrefix + "image").Error
			if err != nil {
				log.Fatal(err)
			}

			// customers
			err = db.Debug().Exec("INSERT IGNORE INTO " + dbTablePrefix + "customer_group ( id_customer, id_group) SELECT id_customer, id_default_group FROM " + dbTablePrefix + "customer").Error
			if err != nil {
				log.Fatal(err)
			}

			// *** eg_product ***
			// set default category
			/*
				err = db.Debug().Exec("UPDATE " + dbTablePrefix + "product SET id_category_default=(SELECT id_category FROM eg_category ORDER BY RAND() LIMIT 1)").Error
				if err != nil {
					log.Fatal(err)
				}

				err = db.Debug().Exec("UPDATE " + dbTablePrefix + "category_product cp SET id_category=(SELECT id_category_default FROM eg_product p WHERE cp.id_product=p.id_product)").Error
				if err != nil {
					log.Fatal(err)
				}
			*/

			// UPDATE eg_product SET id_category_default=(SELECT id_category FROM eg_category ORDER BY RAND() LIMIT 1);
			// UPDATE eg_category_product ecp SET id_category=(SELECT id_category_default FROM eg_product ep WHERE ecp.id_product=ep.id_product);

			// UPDATE eg_category SET id_parent=2

			err = db.Debug().Exec("TRUNCATE TABLE eg_category_group").Error
			if err != nil {
				log.Fatal(err)
			}

			var groups = []string{"1", "2", "3"}
			for _, group := range groups {
				err = db.Debug().Exec("INSERT INTO eg_category_group (`id_category`, `id_group`) SELECT id_category, " + group + " FROM eg_category;").Error
				if err != nil {
					log.Fatal(err)
				}
			}

			var procedureRepairNestedTree = `
DROP PROCEDURE IF EXISTS repair_nested_tree;
DELIMITER //
CREATE PROCEDURE repair_nested_tree ()
MODIFIES SQL DATA
BEGIN
    DECLARE currentId, currentParentId  INT;
    DECLARE currentLeft                 INT;
    DECLARE startId                     INT DEFAULT 1;
    # Determines the max size for MEMORY tables.
    SET max_heap_table_size = 1024 * 1024 * 512;
    START TRANSACTION;
    # Temporary MEMORY table to do all the heavy lifting in,
    # otherwise performance is simply abysmal.
    CREATE TABLE ` +
				"`" +
				`tmp_category_tree` +
				"`" +
				` (
        ` +
				"`" +
				`id_category` +
				"`" +
				` int(10) unsigned NOT NULL DEFAULT 0,
        ` +
				"`" +
				`id_parent` +
				"`" +
				`   int(10)          DEFAULT NULL,
        ` +
				"`" +
				`nleft` +
				"`" +
				`       int(10) unsigned DEFAULT NULL,
        ` +
				"`" +
				`nright` +
				"`" +
				`      int(10) unsigned DEFAULT NULL,
        PRIMARY KEY      (` +
				"`" +
				`id_category` +
				"`" +
				`),
        INDEX USING HASH (` +
				"`" +
				`id_parent` +
				"`" +
				`),
        INDEX USING HASH (` +
				"`" +
				`nleft` +
				"`" +
				`),
        INDEX USING HASH (` +
				"`" +
				`nright` +
				"`" +
				`)
    ) ENGINE = MEMORY
    SELECT ` +
				"`" +
				`id_category` +
				"`" +
				`,
           ` +
				"`" +
				`id_parent` +
				"`" +
				`,
           ` +
				"`" +
				`nleft` +
				"`" +
				`,
           ` +
				"`" +
				`nright` +
				"`" +
				`
    FROM   ` +
				"`" +
				`eg_category` +
				"`" +
				`;
    # Leveling the playing field.
    UPDATE  ` +
				"`" +
				`tmp_category_tree` +
				"`" +
				`
    SET     ` +
				"`" +
				`nleft` +
				"`" +
				`  = NULL,
            ` +
				"`" +
				`nright` +
				"`" +
				` = NULL;
    # Establishing starting numbers for all root elements.
    WHILE EXISTS (SELECT * FROM ` +
				"`" +
				`tmp_category_tree` +
				"`" +
				` WHERE ` +
				"`" +
				`id_parent` +
				"`" +
				` = 0 AND ` +
				"`" +
				`nleft` +
				"`" +
				` IS NULL AND ` +
				"`" +
				`nright` +
				"`" +
				` IS NULL LIMIT 1) DO
        UPDATE ` +
				"`" +
				`tmp_category_tree` +
				"`" +
				`
        SET    ` +
				"`" +
				`nleft` +
				"`" +
				`  = startId,
               ` +
				"`" +
				`nright` +
				"`" +
				` = startId + 1
        WHERE  ` +
				"`" +
				`id_parent` +
				"`" +
				` = 0
          AND  ` +
				"`" +
				`nleft` +
				"`" +
				`       IS NULL
          AND  ` +
				"`" +
				`nright` +
				"`" +
				`      IS NULL
        LIMIT  1;
        SET startId = startId + 2;
    END WHILE;
    # Switching the indexes for the lft/rght columns to B-Trees to speed up the next section, which uses range queries.
    DROP INDEX ` +
				"`" +
				`nleft` +
				"`" +
				`  ON ` +
				"`" +
				`tmp_category_tree` +
				"`" +
				`;
    DROP INDEX ` +
				"`" +
				`nright` +
				"`" +
				` ON ` +
				"`" +
				`tmp_category_tree` +
				"`" +
				`;
    CREATE INDEX ` +
				"`" +
				`nleft` +
				"`" +
				`  USING BTREE ON ` +
				"`" +
				`tmp_category_tree` +
				"`" +
				` (` +
				"`" +
				`nleft` +
				"`" +
				`);
    CREATE INDEX ` +
				"`" +
				`nright` +
				"`" +
				` USING BTREE ON ` +
				"`" +
				`tmp_category_tree` +
				"`" +
				` (` +
				"`" +
				`nright` +
				"`" +
				`);
    # Numbering all child elements
    WHILE EXISTS (SELECT * FROM ` +
				"`" +
				`tmp_category_tree` +
				"`" +
				` WHERE ` +
				"`" +
				`nleft` +
				"`" +
				` IS NULL LIMIT 1) DO
        # Picking an unprocessed element which has a processed parent.
        SELECT     ` +
				"`" +
				`tmp_category_tree` +
				"`" +
				`.` +
				"`" +
				`id_category` +
				"`" +
				`
          INTO     currentId
        FROM       ` +
				"`" +
				`tmp_category_tree` +
				"`" +
				`
        INNER JOIN ` +
				"`" +
				`tmp_category_tree` +
				"`" +
				` AS ` +
				"`" +
				`parents` +
				"`" +
				`
                ON ` +
				"`" +
				`tmp_category_tree` +
				"`" +
				`.` +
				"`" +
				`id_parent` +
				"`" +
				` = ` +
				"`" +
				`parents` +
				"`" +
				`.` +
				"`" +
				`id_category` +
				"`" +
				`
        WHERE      ` +
				"`" +
				`tmp_category_tree` +
				"`" +
				`.` +
				"`" +
				`nleft` +
				"`" +
				` IS NULL
          AND      ` +
				"`" +
				`parents` +
				"`" +
				`.` +
				"`" +
				`nleft` +
				"`" +
				`  IS NOT NULL
        LIMIT      1;
        # Finding the element's parent.
        SELECT  ` +
				"`" +
				`id_parent` +
				"`" +
				`
          INTO  currentParentId
        FROM    ` +
				"`" +
				`tmp_category_tree` +
				"`" +
				`
        WHERE   ` +
				"`" +
				`id_category` +
				"`" +
				` = currentId;
        # Finding the parent's nleft value.
        SELECT  ` +
				"`" +
				`nleft` +
				"`" +
				`
          INTO  currentLeft
        FROM    ` +
				"`" +
				`tmp_category_tree` +
				"`" +
				`
        WHERE   ` +
				"`" +
				`id_category` +
				"`" +
				` = currentParentId;
        # Shifting all elements to the right of the current element 2 to the right.
        UPDATE ` +
				"`" +
				`tmp_category_tree` +
				"`" +
				`
        SET    ` +
				"`" +
				`nright` +
				"`" +
				` = ` +
				"`" +
				`nright` +
				"`" +
				` + 2
        WHERE  ` +
				"`" +
				`nright` +
				"`" +
				` > currentLeft;
        UPDATE ` +
				"`" +
				`tmp_category_tree` +
				"`" +
				`
        SET    ` +
				"`" +
				`nleft` +
				"`" +
				` = ` +
				"`" +
				`nleft` +
				"`" +
				` + 2
        WHERE  ` +
				"`" +
				`nleft` +
				"`" +
				` > currentLeft;
        # Setting nleft and nright values for current element.
        UPDATE ` +
				"`" +
				`tmp_category_tree` +
				"`" +
				`
        SET    ` +
				"`" +
				`nleft` +
				"`" +
				`  = currentLeft + 1,
               ` +
				"`" +
				`nright` +
				"`" +
				` = currentLeft + 2
        WHERE  ` +
				"`" +
				`id_category` +
				"`" +
				`   = currentId;
    END WHILE;
    # Writing calculated values back to physical table.
    UPDATE ` +
				"`" +
				`eg_category` +
				"`" +
				`, ` +
				"`" +
				`tmp_category_tree` +
				"`" +
				`
    SET    ` +
				"`" +
				`eg_category` +
				"`" +
				`.` +
				"`" +
				`nleft` +
				"`" +
				`  = ` +
				"`" +
				`tmp_category_tree` +
				"`" +
				`.` +
				"`" +
				`nleft` +
				"`" +
				`,
           ` +
				"`" +
				`eg_category` +
				"`" +
				`.` +
				"`" +
				`nright` +
				"`" +
				` = ` +
				"`" +
				`tmp_category_tree` +
				"`" +
				`.` +
				"`" +
				`nright` +
				"`" +
				`
    WHERE  ` +
				"`" +
				`eg_category` +
				"`" +
				`.` +
				"`" +
				`id_category` +
				"`" +
				`   = ` +
				"`" +
				`tmp_category_tree` +
				"`" +
				`.` +
				"`" +
				`id_category` +
				"`" +
				`;
    COMMIT;
    DROP TABLE ` +
				"`" +
				`tmp_category_tree` +
				"`" +
				`;
END//
DELIMITER ;
`
			err = db.Debug().Raw(procedureRepairNestedTree).Error
			if err != nil {
				log.Fatal(err)
			}

			err = db.Debug().Exec("CALL repair_nested_tree();").Error
			if err != nil {
				log.Fatal(err)
			}

		}

		// copy and process images
		// workDir
		bimgOpts := bimg.Options{
			Lossless: true,
			Quality:  imgQuality,
		}

		for imgType, imgTypeLabel := range imgTypes {

			imgDir := filepath.Join(workDir, "img", imgType)

			// load image types
			var imageTypes []*models.ImageType
			db.Where(imgTypeLabel + " = 1").Find(&imageTypes)

			i := 1
			err := godirwalk.Walk(imgDir, &godirwalk.Options{
				Callback: func(osPathname string, de *godirwalk.Dirent) error {
					if !de.IsDir() {
						// get the extension from the file path

						extension := filepath.Ext(osPathname)
						filename := path.Base(osPathname)
						basename := strings.Replace(filename, extension, "", -1)
						if inSlice(extension, imgExt, false) {
							if options.verbose {
								log.Println("found=", osPathname, "filename=", filename, "extension=", extension, "basename=", basename)
							}

							// Read original image
							buffer, err := bimg.Read(osPathname)
							checkErr("bimg.Read, error", err)

							// Import image buffer
							newImage := bimg.NewImage(buffer)

							// Get the image dimension
							imgDim, err := newImage.Size()

							// Get the image type
							mediaType := newImage.Type()

							switch mediaType {
							case "unknown":
								return nil // errors.New("Unsupported image format")
							}

							bimgOpts.Height = imgDim.Height
							bimgOpts.Width = imgDim.Width

							// Process image quality
							newBytes, err := newImage.Process(bimgOpts)
							checkErr("bimg.Process, error", err)

							// Write final output
							if !dryRun {
								destinationPrepfixPath := filepath.Join(psDir, "img", imgType)
								destinationFinalPath := buildFolderForImage(destinationPrepfixPath, i)
								if err := os.MkdirAll(destinationFinalPath, 0755); err != nil {
									log.Fatal(err)
								}
								destinationFilePath := filepath.Join(destinationFinalPath, fmt.Sprintf("%d%s", i, extension))
								log.Println("destinationFilePath=", destinationFilePath)
								err = bimg.Write(destinationFilePath, newBytes)
								checkErr("bimg.Write, error", err)

								// todo: product_mini_ ?!
								// https://prestanish.evolutive.group/img/tmp/product_mini_21.jpg?time=1612942719

								for _, imageType := range imageTypes {

									// Import image buffer
									newImage := bimg.NewImage(buffer)

									// Get the image dimension
									// imgDim, err := newImage.Size()

									// Get the image type
									mediaType := newImage.Type()

									switch mediaType {
									case "unknown":
										return nil // errors.New("Unsupported image format")
									}

									bimgOpts.Height = int(imageType.Height)
									bimgOpts.Width = int(imageType.Width)

									// Process image quality
									newBytes, err := newImage.Process(bimgOpts)
									checkErr("bimg.Process, error", err)

									// Write final output
									destinationPrepfixPath := filepath.Join(psDir, "img", imgType)
									destinationFinalPath := buildFolderForImage(destinationPrepfixPath, i)
									if err := os.MkdirAll(destinationFinalPath, 0755); err != nil {
										log.Fatal(err)
									}
									destinationFilePath := filepath.Join(destinationFinalPath, fmt.Sprintf("%d-%s%s", i, imageType.Name, extension))
									log.Println("destinationFilePath=", destinationFilePath)
									err = bimg.Write(destinationFilePath, newBytes)
									checkErr("bimg.Write, error", err)

								}

								i++
							}
						}
					}
					return nil
				},
				Unsorted: true,
			})
			checkErr("Dir walk, error", err)

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
	ImportCmd.Flags().BoolVarP(&dbEnv, "db-env", "", false, "use env variables to connect to the database.")
	ImportCmd.Flags().BoolVarP(&dbDrop, "db-drop", "", false, "drop/truncate database tables")
	ImportCmd.Flags().BoolVarP(&dbMigrate, "db-migrate", "", false, "create/update database tables")
	ImportCmd.Flags().StringVarP(&psDir, "ps-dir", "p", "../../shared/www", "prestashop directory")
	ImportCmd.Flags().IntVarP(&imgQuality, "img-quality", "q", 90, "Image format quality. (works only for JPEG format)")
	ImportCmd.Flags().StringSliceVarP(&imgExt, "img-extension", "e", imgDefExt, "Process only image extensions")
	ImportCmd.Flags().StringVarP(&outputDir, "output-dir", "o", "../../shared/csv", "output directory")
	ImportCmd.Flags().StringVarP(&workDir, "workdir", "w", "../../shared/fixtures", "working directory")
	ImportCmd.Flags().StringSliceVarP(&languages, "language", "l", languagesDef, "languages to import")
	ImportCmd.Flags().BoolVarP(&noFixtures, "no-fixtures", "", false, "skip fixtures loading")
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

func findCategoryByName(db *gorm.DB, name string) (*models.CategoryLang, error) {
	var cat models.CategoryLang
	db.Debug().Raw("SELECT * FROM `" + dbTablePrefix + "category_lang` WHERE name = '" + strings.ToLower(name) + "'").Scan(&cat)
	return &cat, nil
}

func findOrCreateLangByIso(db *gorm.DB, iso_code string) (*models.Lang, error) {
	var lang models.Lang
	db.Debug().Raw("SELECT * FROM eg_lang WHERE iso_code = '" + strings.ToLower(iso_code) + "'").Scan(&lang)
	if lang.IDLang == 0 {
		db.Exec("INSERT INTO `" + dbTablePrefix + "lang` (`name`,`active`,`iso_code`,`language_code`,`locale`,`date_format_lite`,`date_format_full`,`is_rtl`) VALUES ('" + iso6391.Name(iso_code) + "',true,'" + iso_code + "','" + iso_code + "','','','',false)")
		db.Debug().Raw("SELECT * FROM eg_lang WHERE iso_code = '" + strings.ToLower(iso_code) + "'").Scan(&lang)
		return &lang, nil // todo. fix this hacky return
	}
	return &lang, nil
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

func getHeaders(fp string) ([]string, error) {
	var headers []string
	// read
	file, err := os.Open(fp)
	if err != nil {
		return headers, err
	}
	defer file.Close()
	r := csv.NewReader(file)
	r.Comma = ';'
	r.LazyQuotes = true

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

	return headers, nil
}

func loadData(db *gorm.DB, fp, table string, columns ...string) error {
	if db == nil {
		return errors.New("Database not connected")
	}
	if len(columns) == 0 {
		return nil // errors.New("No columns are defined")
	}

	// get the first line
	headers, err := getHeaders(fp)
	if err != nil {
		return err
	}

	if dbDrop {
		if err := db.Exec("TRUNCATE TABLE " + dbTablePrefix + table).Error; err != nil {
			return err
		}
	}

	fmt.Printf("loading data from file %s with colums [%s]\n", fp, strings.Join(headers, ","))
	mysql_driver.RegisterLocalFile(fp)
	query := "LOAD DATA LOCAL INFILE '" + fp + "' IGNORE INTO TABLE " + dbTablePrefix + table + " CHARACTER SET 'utf8mb4' FIELDS TERMINATED BY ',' ENCLOSED BY '\"' LINES TERMINATED BY '\n' IGNORE 1 LINES (`" + strings.Replace(strings.Join(headers, "`,`"), ",", "`,`", -1) + "`);" // SET created_at = NOW(), updated_at = NOW();`
	if options.debug {
		fmt.Println(query)
	}
	if err := db.Debug().Exec(query).Error; err != nil {
		return err
	}
	return nil
}

func loadDataLang(db *gorm.DB, fp, table string, id_lang int, columns ...string) error {
	if db == nil {
		return errors.New("Database not connected")
	}
	if len(columns) == 0 {
		return nil // errors.New("No columns are defined")
	}

	// get the first line
	headers, err := getHeaders(fp)
	if err != nil {
		return err
	}

	if dbDrop {
		if err := db.Debug().Exec("DELETE FROM " + dbTablePrefix + table + " WHERE id_lang='" + fmt.Sprintf("%d", id_lang) + "'").Error; err != nil {
			return err
		}
	}

	fmt.Printf("loading data from file %s with colums [%s]\n", fp, strings.Join(headers, ","))
	mysql_driver.RegisterLocalFile(fp)
	query := "LOAD DATA LOCAL INFILE '" + fp + "' IGNORE INTO TABLE " + dbTablePrefix + table + " CHARACTER SET 'utf8mb4' FIELDS TERMINATED BY ',' ENCLOSED BY '\"' LINES TERMINATED BY '\n' IGNORE 1 LINES (`" + strings.Replace(strings.Join(headers, "`,`"), ",", "`,`", -1) + "`);" // SET created_at = NOW(), updated_at = NOW();`
	if options.debug {
		fmt.Println(query)
	}
	if err := db.Debug().Exec(query).Error; err != nil {
		return err
	}
	return nil
}

/*

SELECT SQL_CALC_FOUND_ROWS p.`id_product`  AS `id_product`,
 p.`reference`  AS `reference`,
 sa.`price`  AS `price`,
 p.`id_shop_default`  AS `id_shop_default`,
 p.`is_virtual`  AS `is_virtual`,
 pl.`name`  AS `name`,
 pl.`link_rewrite`  AS `link_rewrite`,
 sa.`active`  AS `active`,
 shop.`name`  AS `shopname`,
 image_shop.`id_image`  AS `id_image`,
 cl.`name`  AS `name_category`,
 0 AS `price_final`,
 pd.`nb_downloadable`  AS `nb_downloadable`,
 sav.`quantity`  AS `sav_quantity`,
 IF(sav.`quantity`<=0, 1, 0) AS `badge_danger`
FROM  `eg_product` p
 LEFT JOIN `eg_product_lang` pl ON (pl.`id_product` = p.`id_product` AND pl.`id_lang` = 1 AND pl.`id_shop` = 1)
 LEFT JOIN `eg_stock_available` sav ON (sav.`id_product` = p.`id_product` AND sav.`id_product_attribute` = 0 AND sav.id_shop = 1  AND sav.id_shop_group = 0 )
 JOIN `eg_product_shop` sa ON (p.`id_product` = sa.`id_product` AND sa.id_shop = 1)
 LEFT JOIN `eg_category_lang` cl ON (sa.`id_category_default` = cl.`id_category` AND cl.`id_lang` = 1 AND cl.id_shop = 1)
 LEFT JOIN `eg_category` c ON (c.`id_category` = cl.`id_category`)
 LEFT JOIN `eg_shop` shop ON (shop.id_shop = 1)
 LEFT JOIN `eg_image_shop` image_shop ON (image_shop.`id_product` = p.`id_product` AND image_shop.`cover` = 1 AND image_shop.id_shop = 1)
 LEFT JOIN `eg_image` i ON (i.`id_image` = image_shop.`id_image`)
 LEFT JOIN `eg_product_download` pd ON (pd.`id_product` = p.`id_product`)
WHERE (1 AND state = 1)

ORDER BY  `id_product` asc

LIMIT 0, 20
;

*/
