package cmd

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"

	mysql_driver "github.com/go-sql-driver/mysql"
	// "github.com/k0kubun/pp"
	"github.com/karrick/godirwalk"
	"github.com/spf13/cobra"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// go run main.go insert --db-name eg_prestanish --db-user root --db-pass "OvdZ5ZoXWgCWL4-hvZjg!" --db-table-prefix eg_

var (
	psDir         string
	importDir     string
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
	importExt     = []string{".csv", ".txt"}
)

type WkMpImport struct {
	gorm.Model
	IdSeller         uint `gorm:"primaryKey;uniqueIndex:idx_mp_product;column:id_seller;type:int(10) unsigned;not null"`
	IDProduct        uint `gorm:"uniqueIndex:idx_mp_product;index:image_product;column:id_product;type:int(10) unsigned;not null"`
	IdImage          uint
	Name             string
	Reference        string
	CategoryId       uint
	DefaultCategory  uint
	ShortDescription string `gorm:"column:description;type:text"`
	Description      string `gorm:"column:short_description;type:text"`
	Price            float64
	TaxId            uint
	Quantity         uint
	ImageRef         string
	FeatureName      string
	FeatureValues    string
	FormatUnitaire   string
	MinimumGarantie  string
	QuantiteUnitaire uint
	PrixTotal        float64
	UniteTotal       uint
	PoidsPiece       float64
	Ean13            string
	Sku              string
	Mpn              string
}

type WkMpSellerProduct struct {
	IdMpProduct uint
	IdSeller    uint
	IdProduct   uint
	IdCategory  uint
}

var InsertCmd = &cobra.Command{
	Use:     "insert",
	Aliases: []string{"i"},
	Short:   "insert products to webkum's marketplace tables.",
	Long:    "insert products to webkum's marketplace tables",
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

			if dbDrop {
				if db.Migrator().HasTable(&WkMpImport{}) {
					db.Migrator().DropTable(&WkMpImport{})
				}
			}
			if dbMigrate {
				db.Migrator().AutoMigrate(&WkMpImport{})
			}

		}

		err := godirwalk.Walk(importDir, &godirwalk.Options{
			Callback: func(osPathname string, de *godirwalk.Dirent) error {
				if !de.IsDir() {
					// get the extension from the file path
					extension := filepath.Ext(osPathname)
					filename := path.Base(osPathname)
					basename := strings.Replace(filename, extension, "", -1)
					if inSlice(extension, importExt, false) {
						if options.verbose {
							log.Println("found=", osPathname, "filename=", filename, "extension=", extension, "basename=", basename)
						}
						loadData(db, osPathname, "wk_mp_import")
						// checkErr("loadData, error", err)
					}
				}
				return nil
			},
			Unsorted: true,
		})
		checkErr("Dir walk, error", err)

		var products []*WkMpImport
		db.Find(&products)

		for _, product := range products {

			if product.IDProduct == 0 {
				continue
			}

			// insert select products
			sqlQuery := `INSERT IGNORE INTO eg_wk_mp_seller_product (
			id_seller,
			id_ps_product,
			id_ps_shop,
			id_category,
			price,
			wholesale_price,
			unity,
			unit_price,
			id_tax_rules_group,
			on_sale,
			additional_shipping_cost,
			quantity,
			minimal_quantity,
			low_stock_threshold,
			low_stock_alert,
			active,
			status_before_deactivate,
			show_condition,
			` + "`" + `condition` + "`" + `,
			available_for_order,
			show_price,
			online_only,
			visibility,
			admin_assigned,
			width,
			height,
			depth,
			weight,
			reference,
			ean13,
			upc,
			isbn,
			out_of_stock,
			available_date,
			ps_id_carrier_reference,
			admin_approved,
			additional_delivery_times,
			date_add,
			date_upd,
			csv_request_no)
		SELECT 
		  ` + fmt.Sprintf("%d", product.IdSeller) + `,
		  	id_product,
		  	id_shop_default,
		  	id_category_default,
		  	price,
			wholesale_price,
			unity,
			unit_price_ratio,
			id_tax_rules_group,
			on_sale,
			additional_shipping_cost,
			quantity,
			minimal_quantity,
			low_stock_threshold,
			low_stock_alert,
			active,
			0,
			show_condition,
			` + "`" + `condition` + "`" + `,
			available_for_order,
			show_price,
			online_only,
			visibility,
			1,
			width,
			height,
			depth,
			weight,
			reference,
			ean13,
			upc,
			isbn,
			out_of_stock,
			available_date,		
			0,
			1,
			0,
			date_add,
			date_upd,	
			0
		FROM eg_product
		WHERE id_product=` + fmt.Sprintf("%d", product.IDProduct)
			db.Exec(sqlQuery)

			// retrieve id_seller_product
			var wkMpSellerProduct WkMpSellerProduct
			err := db.Where("id_seller = ? AND id_ps_product = ? and id_ps_shop = ?", product.IdSeller, product.IDProduct, 1).First(&wkMpSellerProduct).Error
			if errors.Is(err, gorm.ErrRecordNotFound) {
				log.Fatal("product not found")
			}

			// eg_wk_mp_seller_product_lang
			sqlQuery = `INSERT IGNORE INTO eg_wk_mp_seller_product_lang (id_mp_product, id_lang, product_name, short_description, description, available_now, available_later, meta_title, meta_description, link_rewrite, delivery_in_stock, delivery_out_stock)
						SELECT ` + fmt.Sprintf("%d", wkMpSellerProduct.IdMpProduct) + `, id_lang, name, description_short, description, available_now, available_later, meta_title, meta_description, link_rewrite, delivery_in_stock, delivery_out_stock FROM eg_product_lang WHERE
						id_product=` + fmt.Sprintf("%d", product.IDProduct)
			db.Exec(sqlQuery)

			// insert select categories
			sqlQuery = `INSERT IGNORE INTO eg_wk_mp_seller_product_category (
						  id_category,
						  id_seller_product,
						  is_default) 		
						VALUES ( 
						  ` + fmt.Sprintf("%d", product.DefaultCategory) + `,
						  ` + fmt.Sprintf("%d", wkMpSellerProduct.IdMpProduct) + `,
						  1)`
			db.Exec(sqlQuery)

			// eg_wk_mp_seller_product_image
			sqlQuery = `INSERT IGNORE INTO eg_wk_mp_seller_product_image (seller_product_id, seller_product_image_name, id_ps_image, position, cover, active) 
			  	VALUES (` + fmt.Sprintf("%d", wkMpSellerProduct.IdMpProduct) + `, "` + escape(product.Name) + `",` + fmt.Sprintf("%d", product.IdImage) + `, 0, 1, 1)`
			db.Exec(sqlQuery)

			// eg_wk_mp_product_feature
			sqlQuery = `INSERT IGNORE INTO eg_wk_mp_product_feature (ps_id_feature, mp_id_product, ps_id_feature_value, mp_id_feature_value)
							SELECT DISTINCT fv.id_feature, ` + fmt.Sprintf("%d", product.IDProduct) + `,` + fmt.Sprintf("%d", wkMpSellerProduct.IdMpProduct) + `, fv.id_feature_value, wmpfv.mp_id_feature_value FROM eg_feature_product fp
							LEFT JOIN eg_feature f ON f.id_feature=fp.id_feature
							LEFT JOIN eg_feature_value fv ON fv.id_feature=fp.id_feature
							LEFT JOIN eg_wk_mp_product_feature_value wmpfv ON wmpfv.ps_id_feature=fp.id_feature
							WHERE fp.id_product=` + fmt.Sprintf("%d", product.IDProduct)
			db.Debug().Exec(sqlQuery)

		}

		/*
			// to do. finish features import
			// insert select feature values
			sqlQuery = `INSERT IGNORE INTO eg_wk_mp_product_feature_value (ps_id_feature, is_custom)
							SELECT DISTINCT id_feature, custom FROM eg_feature_value`
			db.Debug().Exec(sqlQuery)

			// eg_wk_mp_product_feature_value_lang
			sqlQuery = `INSERT IGNORE INTO eg_wk_mp_product_feature_value_lang (ps_id_feature_value, mp_id_feature_value, id_lang, value)
								SELECT DISTINCT wmpfv.id_feature_value, wmpf.mp_id_feature_value,id_lang, fvl.value FROM eg_feature_product fp
								LEFT JOIN eg_feature f ON f.id_feature=fp.id_feature
								LEFT JOIN eg_feature_value fv ON fv.id_feature=fp.id_feature
								LEFT JOIN eg_feature_value_lang fvl ON fv.id_feature=fp.id_feature
								LEFT JOIN eg_wk_mp_product_feature wmpf ON wmpf.ps_id_feature=fp.id_feature
								LEFT JOIN eg_wk_mp_product_feature_value wmpfv ON wmpfv.ps_id_feature=fp.id_feature
								WHERE fp.id_product=` + fmt.Sprintf("%d", product.IDProduct)
			db.Debug().Exec(sqlQuery)
		*/

	},
}

func init() {
	// todo. manage ssl connections for db
	InsertCmd.Flags().StringVarP(&dbTablePrefix, "db-table-prefix", "", "ps_", "database table prefix")
	InsertCmd.Flags().StringVarP(&dbName, "db-name", "", "prestashop", "database name")
	InsertCmd.Flags().StringVarP(&dbUser, "db-user", "", "root", "database username")
	InsertCmd.Flags().StringVarP(&dbPass, "db-pass", "", "", "database password")
	InsertCmd.Flags().StringVarP(&dbHost, "db-host", "", "127.0.0.1", "database host")
	InsertCmd.Flags().StringVarP(&dbPort, "db-port", "", "3306", "datbase port")
	InsertCmd.Flags().BoolVarP(&dbEnv, "db-env", "", false, "use env variables to connect to the database.")
	InsertCmd.Flags().BoolVarP(&dbMigrate, "db-migrate", "", false, "migrate database tables")
	InsertCmd.Flags().BoolVarP(&dbDrop, "db-drop", "", false, "drop/truncate database tables")
	InsertCmd.Flags().StringVarP(&importDir, "import-dir", "o", "./catalogs", "output directory")
	InsertCmd.Flags().StringSliceVarP(&languages, "language", "l", languagesDef, "languages to import")
	InsertCmd.Flags().BoolVarP(&dryRun, "dry-run", "", false, "dry run")
	RootCmd.AddCommand(InsertCmd)
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
	r.Comma = ','
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

func escape(sql string) string {
	dest := make([]byte, 0, 2*len(sql))
	var escape byte
	for i := 0; i < len(sql); i++ {
		c := sql[i]

		escape = 0

		switch c {
		case 0: /* Must be escaped for 'mysql' */
			escape = '0'
			break
		case '\n': /* Must be escaped for logs */
			escape = 'n'
			break
		case '\r':
			escape = 'r'
			break
		case '\\':
			escape = '\\'
			break
		case '\'':
			escape = '\''
			break
		case '"': /* Better safe than sorry */
			escape = '"'
			break
		case '\032': /* This gives problems on Win32 */
			escape = 'Z'
		}

		if escape != 0 {
			dest = append(dest, '\\', escape)
		} else {
			dest = append(dest, c)
		}
	}

	return string(dest)
}

func loadData(db *gorm.DB, fp, table string) error {
	if db == nil {
		return errors.New("Database not connected")
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

	if len(headers) <= 1 {
		return errors.New("Invalid csv file")
	}

	fmt.Printf("loading data from file %s with colums [%s]\n", fp, strings.Join(headers, ","))
	mysql_driver.RegisterLocalFile(fp)
	query := "LOAD DATA LOCAL INFILE '" + fp + "' IGNORE INTO TABLE " + dbTablePrefix + table + " CHARACTER SET 'utf8mb4' FIELDS TERMINATED BY ',' ENCLOSED BY '\"' LINES TERMINATED BY '\n' IGNORE 1 LINES (`" + strings.Join(headers, "`,`") + "`);" // SET created_at = NOW(), updated_at = NOW();`
	if options.debug {
		fmt.Println(query)
	}
	if err := db.Debug().Exec(query).Error; err != nil {
		return err
	}
	return nil
}
