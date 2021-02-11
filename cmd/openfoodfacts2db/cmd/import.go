package cmd

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"os"
	"strings"

	// "github.com/fatih/structs"
	mysql_driver "github.com/go-sql-driver/mysql"
	"github.com/k0kubun/pp"
	"github.com/spf13/cobra"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	// "github.com/lucmichalski/go-prestashop/internal/models"
)

/*
	Refs:
	- https://github.com/lucmichalski/presta-faker/blob/master/src/PrestaFaker/Webservice/Sql/Category.php
*/

var (
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
	csvLoad       bool
	csvFile       string
	outputDir     string
	dryRun        bool
)

var ImportCmd = &cobra.Command{
	Use:     "import",
	Aliases: []string{"i"},
	Short:   "import openfoodfcats dataset to prestashop database.",
	Long:    "import openfoodfcats dataset to prestashop database.",
	Run: func(cmd *cobra.Command, args []string) {

		if !dryRun {
			var err error
			// connect to database
			dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
			db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
				NamingStrategy: schema.NamingStrategy{
					// TablePrefix:   dbTablePrefix,                   // table name prefix, table for `Lang` would be `eg_lang`
					SingularTable: true,                            // use singular table name, table for `User` would be `user` with this option enabled
					NameReplacer:  strings.NewReplacer("ID", "Id"), // use name replacer to change struct/field name before convert it to db name
					// SkipDefaultTransaction: true,
				},
			})
			if err != nil {
				log.Fatal(err)
			}

			// create and/or truncate tables
			err = prepareTables(db, Tables...)
			if err != nil {
				log.Fatal(err)
			}

			// load dataset into a temporary table
			// todo: split file
			// refs: https://github.com/uk0/go-split-file/blob/master/command.go
			if csvLoad {
				err := loadData(db, csvFile, "openfoodfact_tmp")
				if err != nil {
					log.Fatal(err)
				}
			}

			// extract products
			var productCnt int64
			db.Model(&OpenFoodFact{}).Count(&productCnt)
			offset := 0
			limit := 10000
			maxIter := math.Round(float64(int(productCnt) / limit))
			for i := 0; i < int(maxIter); i++ {
				var products []*OpenFoodFact
				offset := offset + limit
				db.Debug().Model(&OpenFoodFact{}).Limit(limit).Offset(offset).Find(&products)
				for _, product := range products {

					/*
									   return [
									        // Product
									        'id_supplier' => 0,
									        'id_manufacturer' => 0,
									        'id_category_default' => 0, // required
									        'id_shop_default' => 1,
									        'id_tax_rules_group' => 0,
									        'on_sale' => 0,
									        'online_only' => 0,
									        'ean13' => '',
									        'upc' => '',
									        'ecotax' => 0,
									        'quantity' => 0,
									        'minimal_quantity' => 0,
									        'price' => 0, // required
									        'wholesale_price' => 0,
									        'unity' => '',
									        'unit_price_ratio' => 0,
									        'additional_shipping_cost' => 0,
									        'reference' => '0',
									        'supplier_reference' => '',
									        'location' => '',
									        'width' => 0,
									        'height' => 0,
									        'depth' => 0,
									        'weight' => 0,
									        'out_of_stock' => 2,
									        'quantity_discount' => 0,
									        'customizable' => 0,
									        'uploadable_files' => 0,
									        'text_fields' => 0,
									        'active' => 1,
									        'redirect_type' => '',
									        'id_product_redirected' => 0,
									        'available_for_order' => 1,
									        'available_date' => '0000-00-00',
									        'condition' => 'new',
									        'show_price' => 1,
									        'indexed' => 1,
									        'visibility' => 'both',
									        'cache_is_pack' => 0,
									        'cache_has_attachments' => 0,
									        'is_virtual' => 0,
									        'cache_default_attribute' => 0,
									        'advanced_stock_management' => 0,

									        // Lang
									        'id_shop' => 1,
									        'id_lang' => 1,
									        'description' => '',
									        'description_short' => '',
									        'link_rewrite' => '', // required
									        'meta_description' => '',
									        'meta_keywords' => '',
									        'meta_title' => '',
									        'name' => '', // required
									        'available_now' => '',
									        'available_later' => '',


						INSERT INTO :prefix:product (`id_product`, `id_supplier`, `id_manufacturer`, `id_category_default`, `id_shop_default`, `id_tax_rules_group`, `on_sale`, `online_only`, `ean13`, `upc`, `ecotax`, `quantity`, `minimal_quantity`, `price`, `wholesale_price`, `unity`, `unit_price_ratio`, `additional_shipping_cost`, `reference`, `supplier_reference`, `location`, `width`, `height`, `depth`, `weight`, `out_of_stock`, `quantity_discount`, `customizable`, `uploadable_files`, `text_fields`, `active`, `redirect_type`, `id_product_redirected`, `available_for_order`, `available_date`, `condition`, `show_price`, `indexed`, `visibility`, `cache_is_pack`, `cache_has_attachments`, `is_virtual`, `cache_default_attribute`, `date_add`, `date_upd`, `advanced_stock_management`)
						VALUES (:id:, :id_supplier:, :id_manufacturer:, :id_category_default:, :id_shop_default:, :id_tax_rules_group:, :on_sale:, :online_only:, :ean13:, :upc:, :ecotax:, :quantity:, :minimal_quantity:, :price:, :wholesale_price:, :unity:, :unit_price_ratio:, :additional_shipping_cost:, :reference:, :supplier_reference:, :location:, :width:, :height:, :depth:, :weight:, :out_of_stock:, :quantity_discount:, :customizable:, :uploadable_files:, :text_fields:, :active:, :redirect_type:, :id_product_redirected:, :available_for_order:, :available_date:, :condition:, :show_price:, :indexed:, :visibility:, :cache_is_pack:, :cache_has_attachments:, :is_virtual:, :cache_default_attribute:, NOW(), NOW(), :advanced_stock_management:);
						INSERT INTO :prefix:product_lang (`id_product`, `id_shop`, `id_lang`, `description`, `description_short`, `link_rewrite`, `meta_description`, `meta_keywords`, `meta_title`, `name`, `available_now`, `available_later`)
						VALUES (:id:, :id_shop:, :id_lang:, :description:, :description_short:, :link_rewrite:, :meta_description:, :meta_keywords:, :meta_title:, :name:, :available_now:, :available_later:);
						INSERT INTO :prefix:product_shop (`id_product`, `id_shop`, `id_category_default`, `id_tax_rules_group`, `on_sale`, `online_only`, `ecotax`, `minimal_quantity`, `price`, `wholesale_price`, `unity`, `unit_price_ratio`, `additional_shipping_cost`, `customizable`, `uploadable_files`, `text_fields`, `active`, `redirect_type`, `id_product_redirected`, `available_for_order`, `available_date`, `condition`, `show_price`, `indexed`, `visibility`, `cache_default_attribute`, `advanced_stock_management`, `date_add`, `date_upd`)
						VALUES (:id:, :id_shop:, :id_category_default:, :id_tax_rules_group:, :on_sale:, :online_only:, :ecotax:, :minimal_quantity:, :price:, :wholesale_price:, :unity:, :unit_price_ratio:, :additional_shipping_cost:, :customizable:, :uploadable_files:, :text_fields:, :active:, :redirect_type:, :id_product_redirected:, :available_for_order:, :available_date:, :condition:, :show_price:, :indexed:, :visibility:, :cache_default_attribute:, :advanced_stock_management:, NOW(), NOW());
						INSERT INTO :prefix:category_product (`id_product`, `id_category`, `position`)
						VALUES (:id:, :id_category_default:, 0);
						EOF;
						    const PRODUCT_FEATURE_SQL = <<<EOF
						INSERT INTO :prefix:feature_product (`id_product`, `id_feature`, `id_feature_value`)
						VALUES (:id:, :id_feature:, :id_feature_value:);
						EOF;

					*/

					// extract product category
					/*
									   return [
									        // Category
									        'id_parent' => 0,
									        'id_shop_default' => 1,
									        'level_depth' => 0,
									        'nleft' => 0,
									        'nright' => 0,
									        'active' => 1,

									        // Lang
									        'id_shop' => 1,
									        'id_lang' => 1,
									        'description' => '',
									        'link_rewrite' => '', // required
									        'meta_description' => '',
									        'meta_keywords' => '',
									        'meta_title' => '',
									        'name' => '', // required

									        // Shop
									        'position' => 0,
									    ];

						INSERT INTO :prefix:category (`id_category`, `id_parent`, `id_shop_default`, `level_depth`, `nleft`, `nright`, `active`, `date_add`, `date_upd`, `is_root_category`)
						VALUES (:id:, :id_parent:, :id_shop_default:, :level_depth:, :nleft:, :nright:, :active:, NOW(), NOW(), :is_root_category:);
						INSERT INTO :prefix:category_lang (`id_category`, `id_shop`, `id_lang`, `name`, `description`, `link_rewrite`, `meta_description`, `meta_keywords`, `meta_title`)
						VALUES (:id:, :id_shop:, :id_lang:, :name:, :description:, :link_rewrite:, :meta_description:, :meta_keywords:, :meta_title:);
						INSERT INTO :prefix:category_shop (`id_category`, `id_shop`, `position`)
						VALUES (:id:, :id_shop:, :position:);
						INSERT INTO :prefix:category_group (`id_category`, `id_group`)
						VALUES (:id:, 1), (:id:, 2), (:id:, 3);

					*/

					// extract products image/cover
					/*

						        return [
						            // Image
						            'id_product' => 0, // required
						            'position' => 0,
						            'cover' => 0,

						            // Lang
						            'id_lang' => 1,
						            'legend' => '',

						            // Shop
						            'id_shop' => 1,
						        ];

								    const IMAGE_SQL = <<<EOF
								INSERT INTO :prefix:image (`id_image`, `id_product`, `position`, `cover`)
								VALUES (:id:, :id_product:, :position:, :cover:);
								INSERT INTO :prefix:image_lang (`id_image`, `id_lang`, `legend`)
								VALUES (:id:, :id_lang:, :legend:);
								INSERT INTO :prefix:image_shop (`id_image`, `id_shop`, `cover`)
								VALUES (:id:, :id_shop:, :cover:);

					*/

					// extract features
					/*
								return [
						            // Feature
						            'position' => 0,

						            // Lang
						            'id_lang' => 1,
						            'name' => '', // required

						            // Shop
						            'id_shop' => 1,
						        ];
								INSERT INTO :prefix:feature (`id_feature`, `position`)
								VALUES (:id:, :position:);
								INSERT INTO :prefix:feature_lang (`id_feature`, `id_lang`, `name`)
								VALUES (:id:, :id_lang:, :name:);
								INSERT INTO :prefix:feature_shop (`id_feature`, `id_shop`)
								VALUES (:id:, :id_shop:);
					*/

					/*
						        return [
						            // Feature
						            'id_feature' => 0, // required

						            // Lang
						            'id_lang' => 1,
						            'value' => '', // required
						        ];

								INSERT INTO :prefix:feature_value (`id_feature_value`, `id_feature`, `custom`)
								VALUES (:id:, :id_feature:, :custom:);
								INSERT INTO :prefix:feature_value_lang (`id_feature_value`, `id_lang`, `value`)
								VALUES (:id:, :id_lang:, :value:);

					*/

				}
			}

		}

	},
}

func init() {
	// todo. manage ssl connections for db
	ImportCmd.Flags().StringVarP(&dbTablePrefix, "db-table-prefix", "", "", "database table prefix")
	ImportCmd.Flags().StringVarP(&dbName, "db-name", "", "openfoodfacts", "database name")
	ImportCmd.Flags().StringVarP(&dbUser, "db-user", "", "root", "database username")
	ImportCmd.Flags().StringVarP(&dbPass, "db-pass", "", "", "database password")
	ImportCmd.Flags().StringVarP(&dbHost, "db-host", "", "127.0.0.1", "database host")
	ImportCmd.Flags().StringVarP(&dbPort, "db-port", "", "3306", "datbase port")
	ImportCmd.Flags().BoolVarP(&dbEnv, "db-env", "", false, "use env variables to connect to the database.")
	ImportCmd.Flags().BoolVarP(&dbDrop, "db-drop", "", false, "drop/truncate database tables")
	ImportCmd.Flags().BoolVarP(&dbMigrate, "db-migrate", "", false, "create/update database tables")
	ImportCmd.Flags().StringVarP(&outputDir, "output-dir", "o", "../../shared/datasets/openfoodfacts.org/prestashop", "csv dataset input file")
	ImportCmd.Flags().StringVarP(&csvFile, "csv-file", "f", "../../shared/datasets/en.openfoodfacts.org.products.csv", "csv dataset input file")
	ImportCmd.Flags().BoolVarP(&csvLoad, "csv-load", "", false, "load csv file into db")
	ImportCmd.Flags().BoolVarP(&dryRun, "dry-run", "", false, "dry run")
	RootCmd.AddCommand(ImportCmd)
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

func countFileLine(name string) (count int64, err error) {
	data, err := ioutil.ReadFile(name)
	if err != nil {
		return
	}
	count = 0
	for i := 0; i < len(data); i++ {
		if data[i] == '\n' {
			count++
		}
	}
	return
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
	r.Comma = '\t'
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

func loadData(db *gorm.DB, fp, table string) error {
	if db == nil {
		return errors.New("Database not connected")
	}

	// get the first line
	headers, err := getHeaders(fp)
	if err != nil {
		return err
	}

	var cleanHdrs []string
	for _, hdr := range headers {
		if strings.HasPrefix(hdr, "-") {
			cleanHdrs = append(cleanHdrs, string(hdr[1:]))
		} else {
			cleanHdrs = append(cleanHdrs, hdr)
		}
	}

	fmt.Printf("loading data from file %s with colums [%s]\n", fp, strings.Join(cleanHdrs, ","))
	mysql_driver.RegisterLocalFile(fp)
	query := "LOAD DATA LOCAL INFILE '" + fp + "' INTO TABLE " + dbTablePrefix + table + " CHARACTER SET 'utf8mb4' FIELDS TERMINATED BY '\t' ENCLOSED BY '\"' LINES TERMINATED BY '\n' IGNORE 1 LINES (`" + strings.Join(cleanHdrs, "`,`") + "`) SET created_at = NOW(), updated_at = NOW();"
	if options.debug {
		fmt.Println(query)
	}
	if err := db.Exec(query).Error; err != nil {
		return err
	}
	return nil
}
