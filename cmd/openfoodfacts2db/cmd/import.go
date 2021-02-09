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
	"path/filepath"
	"sort"
	"strings"

	"github.com/fatih/structs"
	mysql_driver "github.com/go-sql-driver/mysql"
	"github.com/iancoleman/strcase"
	"github.com/k0kubun/pp"
	"github.com/spf13/cobra"
	"golang.org/x/net/html/charset"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/lucmichalski/go-prestashop/internal/models"
)

/*
	Refs:
	go run main.go load --db-name go_prestashop --db-user root --db-pass "OvdZ5ZoXWgCWL4-hvZjg!" --db-table-prefix eg_
*/

var (
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
	languagesDef  = []string{"en", "fr"}
)

var ImportCmd = &cobra.Command{
	Use:     "import",
	Aliases: []string{"i"},
	Short:   "load prestashop-shop-creator xml to a mysql database.",
	Long:    "load prestashop-shop-creator xml to a mysql database",
	Run: func(cmd *cobra.Command, args []string) {

		if !dryRun {
			var err error
			// connect to database
			dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
			db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
				NamingStrategy: schema.NamingStrategy{
					SingularTable: true,
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

		// mapping for languages
		// languagesMap := make(map[string]int, 0)

		for _, fixture := range fixtures.Fixtures {

			n := structs.Name(fixture)
			osPathname := filepath.Join(workDir, "data", fmt.Sprintf("%s.xml", strings.ToLower(strcase.ToSnake(n))))
			if options.debug {
				pp.Println("osPathname:", osPathname)
			}

			file, err := os.Open(osPathname)
			if err != nil {
				log.Fatal(err)
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

		for _, language := range languages {

			for _, model := range langs.Fixtures {

				n := structs.Name(model)
				osPathname := filepath.Join(workDir, "langs", language, "data", fmt.Sprintf("%s.xml", strings.Replace(strings.ToLower(strcase.ToSnake(n)), "_lang", "", 1)))
				if options.debug {
					pp.Println("osPathname:", osPathname)
				}

				file, err := os.Open(osPathname)
				if err != nil {
					log.Fatal(err)
				}

				reader := bufio.NewReader(file)
				dec := xml.NewDecoder(reader)
				dec.CharsetReader = charset.NewReaderLabel
				dec.Strict = false
				if err := dec.Decode(&model); err != nil {
					log.Fatal(err)
				}

				// todo. create base dir
				csvFile, err := os.Create(filepath.Join(outputDir, "langs", language, "data", fmt.Sprintf("%s.csv", strings.ToLower(strcase.ToSnake(n)))))
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
				}

				columnsMap := make(map[string]bool, 0)
				rows := make([][]string, 0)
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
							row = append(row, fmt.Sprintf("%v", entity.(map[string]interface{})[k]))
							columnsMap[k] = true
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
				sort.Strings(columns)

				buf, err := encodeToCsv(columns, rows)
				csvFile.Write(buf)

				// load data into file
				// err := loadData(fp string, DB *gorm.DB)

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
	ImportCmd.Flags().BoolVarP(&dbEnv, "db-env", "", false, "use env variables to connect to the database.")
	ImportCmd.Flags().BoolVarP(&dbDrop, "db-drop", "", false, "drop/truncate database tables")
	ImportCmd.Flags().BoolVarP(&dbMigrate, "db-migrate", "", false, "create/update database tables")
	ImportCmd.Flags().StringVarP(&outputDir, "output-dir", "o", "../../shared/csv", "output directory")
	ImportCmd.Flags().StringVarP(&workDir, "workdir", "w", "../../shared/fixtures", "working directory")
	ImportCmd.Flags().StringSliceVarP(&languages, "language", "l", languagesDef, "languages to import")
	ImportCmd.Flags().BoolVarP(&dryRun, "dry-run", "", false, "dry run")
	RootCmd.AddCommand(ImportCmd)
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

	fmt.Printf("loading data from file %s with colums [%s]\n", fp, strings.Join(headers, ","))
	mysql_driver.RegisterLocalFile(fp)
	query := "LOAD DATA LOCAL INFILE '" + fp + "' INTO TABLE " + dbTablePrefix + table + " CHARACTER SET 'utf8mb4' FIELDS TERMINATED BY ',' ENCLOSED BY '\"' LINES TERMINATED BY '\n' IGNORE 1 LINES (`" + strings.Replace(strings.Join(headers, "`,`"), ",", "`,`", -1) + "`);" // SET created_at = NOW(), updated_at = NOW();`
	if options.debug {
		fmt.Println(query)
	}
	if err := db.Exec(query).Error; err != nil {
		return err
	}
	return nil
}
