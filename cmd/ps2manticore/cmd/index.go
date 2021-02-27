package cmd

import (
	"bytes"
	"database/sql"
	"fmt"
	"strings"
	"text/template"

	"github.com/nozzle/throttler"
	"github.com/spf13/cobra"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

/*
	Todo. select data with ranges (offset,limit...)
*/

// go run main.go index --db-name eg_prestanish --db-user root --db-pass "OvdZ5ZoXWgCWL4-hvZjg!" --db-table-prefix eg_ --index-name products
var (
	indexName       string
	psDir           string
	workDir         string
	dryRun          bool
	db              *gorm.DB
	dbName          string
	dbUser          string
	dbPass          string
	dbHost          string
	dbPort          string
	dbTablePrefix   string
	manticoreDriver string
	manticoreHost   string
	manticorePort   string
)

var ImportCmd = &cobra.Command{
	Use:     "index",
	Aliases: []string{"i"},
	Short:   "index prestashop products and categories to manticoresearch.",
	Long:    "index prestashop products and categories to manticoresearch real-time indices",
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

		cl, err := sql.Open("mysql", "@tcp(127.0.0.1:9306)/")
		if err != nil {
			panic(err)
		}

		switch indexName {
		case "products", "product":

			// todo. create cli argument for limit
			limit := 1000000
			offset := 0

			for {

				var products []*Product

				sqlResult := bytes.NewBufferString("")
				sqlTemplate, _ := template.New("").Parse(sqlAdminCatalogProducts)
				sqlTemplate.Execute(sqlResult, map[string]string{"prefix": dbTablePrefix, "offset": fmt.Sprintf("%d", offset), "limit": fmt.Sprintf("%d", limit)})

				err = db.Raw(sqlResult.String()).Scan(&products).Error
				if err != nil {
					log.Fatal(err)
				}

				offset = offset + limit

				if len(products) == 0 {
					break
				}

				t := throttler.New(6, len(products))

				for _, product := range products {

					go func(p *Product) error {
						// Let Throttler know when the goroutine completes
						// so it can dispatch another worker
						defer t.Done(nil)

						query := `SELECT eg_feature_product.id_product, eg_feature_value_lang.id_feature_value as id_feature_value, eg_feature_value_lang.value as value, eg_feature_lang.name as name, eg_feature_lang.id_feature as id_feature
						FROM eg_feature_product
						INNER JOIN eg_feature_value ON eg_feature_product.id_feature_value = eg_feature_value.id_feature_value
						INNER JOIN eg_feature_value_lang ON eg_feature_value_lang.id_feature_value = eg_feature_value.id_feature_value
						INNER JOIN eg_feature ON eg_feature_value.id_feature = eg_feature.id_feature
						INNER JOIN eg_feature_lang ON eg_feature.id_feature = eg_feature_lang.id_feature
						WHERE eg_feature_lang.id_lang = 1 
						AND eg_feature_value_lang.id_lang = 1 
						AND eg_feature_product.id_product = ` + fmt.Sprintf("%d", p.IdProduct)

						// create sub query for features
						var features []*Feature
						db.Raw(query).Scan(&features)

						var engine, make, model, size, color, season string
						var engine_id, make_id, model_id, size_id, season_id, color_id int
						for _, feat := range features {
							switch feat.Name {
							case "Marque", "brand", "brand_name", "marque", "fabricant", "name_of_brand", "nom_marque", "make", "manufacturer":
								make = feat.Value
								engine_id = feat.IdFeatureValue
							case "Modèle", "model", "modele", "vehicule":
								model = feat.Value
								model_id = feat.IdFeatureValue
							case "Motorisation", "engine", "motor", "moteur":
								engine = feat.Value
								engine_id = feat.IdFeatureValue
							case "size", "taille":
								size = feat.Value
								size_id = feat.IdFeatureValue
							case "color", "couleur", "matiere":
								color = feat.Value
								color_id = feat.IdFeatureValue
							// case "genre", "kind":
							// 	genre = feat.Value
							// 	genre_id = feat.IdFeatureValue
							case "saison", "season":
								season = feat.Value
								season_id = feat.IdFeatureValue
							}
						}

						query = fmt.Sprintf(`REPLACE into rt_products (
					id,
					date_add,
					date_upd,
					id_product,
					price,
					id_shop_default,
					is_virtual,
					link_rewrite,
					active,
					shop_name,
					shop_id,
					id_image,
					price_final,
					nb_downloadable,
					sav_quantity,
					badge_danger,
					features,
					feature_values,
					name,
					reference,
					description,
					description_short,
					name_category,
					engine,
					id_engine,
					model,
					id_model,
					make,
					id_make,
					size,
					id_size,
					color,
					id_color,
					season,
					id_season,
					ft_engine,
					ft_model,
					ft_make,
					id_category
					) VALUES ('%s','%d','%d','%d','%.2f','%d','%t','%s','%t','%s','%d','%d','%.2f','%d','%d','%t',(%s),(%s),'%s','%s','%s','%s','%s','%s','%d','%s','%d','%s','%d','%s','%d','%s','%d','%s','%d','%s','%s','%s','%d')`,
							fmt.Sprintf("%d-%d", p.IdProduct, p.ShopId),
							p.DateAdd.Unix(),
							p.DateUpd.Unix(),
							p.IdProduct,
							p.Price, //
							p.IdShopDefault,
							p.IsVirtual,
							p.LinkRewrite,
							p.Active,
							escape(p.ShopName),
							p.ShopId,
							p.IdImage,
							p.PriceFinal,
							p.NbDownloadable,
							p.SavQuantity,
							p.BadgeDanger,
							p.Features,
							p.FeatureValues,
							escape(p.Name),
							escape(p.Reference),
							escape(p.Description),
							escape(p.DescriptionShort),
							escape(p.NameCategory),
							escape(make),
							make_id,
							escape(engine),
							engine_id,
							escape(model),
							model_id,
							escape(size),
							size_id,
							escape(color),
							color_id,
							escape(season),
							season_id,
							escape(engine),
							escape(model),
							escape(make),
							p.IdCategory,
						)

						_, err = cl.Exec(query)
						if err != nil {
							log.Infoln(query)
							log.Fatal(err)
							return err
						}
						log.Infoln("Index product >> ", p.IdProduct, "==", p.Name)
						return nil
					}(product)

					t.Throttle()

				}

				if t.Err() != nil {
					// Loop through the errors to see the details
					for i, err := range t.Errs() {
						log.Printf("error #%d: %s", i, err)
					}
					log.Fatal(t.Err())
				}
			}

		case "orders", "order":

			// todo. create cli argument for limit
			limit := 10000
			offset := 0

			for {

				sqlResult := bytes.NewBufferString("")
				sqlTemplate, _ := template.New("").Parse(sqlOrdersExtended)
				sqlTemplate.Execute(sqlResult, map[string]string{"prefix": dbTablePrefix, "offset": fmt.Sprintf("%d", offset), "limit": fmt.Sprintf("%d", limit)})

				var orders []*Order
				err = db.Raw(sqlResult.String()).Scan(&orders).Error
				if err != nil {
					log.Fatal(err)
				}

				offset = offset + limit

				if len(orders) == 0 {
					break
				}

				t := throttler.New(1, len(orders))

				for _, order := range orders {

					go func(o *Order) error {
						// Let Throttler know when the goroutine completes
						// so it can dispatch another worker
						defer t.Done(nil)

						query := fmt.Sprintf(`REPLACE into rt_orders (
											id,
											id_order,
											id_currency,
											id_pdf,
											id_shop,
											id_lang,
											order_payment,
											module_payment,
											total_paid,
											total_discounts,
											total_discounts_tax_incl,
											total_discounts_tax_excl,
											total_paid_tax_incl,
											total_paid_tax_excl,
											total_paid_real,
											total_products,
											total_products_wt,
											total_shipping,
											total_shipping_tax_incl,
											total_shipping_tax_excl,
											carrier_tax_rate,
											total_wrapping,
											total_wrapping_tax_incl,
											total_wrapping_tax_excl,
											round_mode,
											round_type,
											invoice_number,
											delivery_number,
											invoice_date,
											delivery_date,
											valid,
											date_add,
											date_upd,
											id_address_delivery,
											customer_name,
											order_status_name,
											order_status_color,
											order_new,
											id_country,
											country_name,
											company,
											vat_number,
											phone_mobile,
											phone,
											delivery_lastname,
											delivery_firstname,
											address_1,
											address_2,
											postcode,
											city,
											badge_success				
							) VALUES ('%d','%d','%d','%d','%d','%d','%s','%s','%.2f','%.2f','%.2f','%.2f','%.2f','%.2f','%.2f','%s','%s','%.2f','%.2f','%.2f','%.2f','%.2f','%.2f','%d','%d','%d','%s','%d','%d','%d','%d','%d','%.2f','%d','%s','%s','%s','%d','%d','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%d')`,
							o.IdOrder,                              // uint
							o.IdOrder,                              // uint
							o.IdCurrency,                           // uint
							o.IdPdf,                                // uint
							o.IdShop,                               // uint
							o.IdLang,                               // uint
							escape(o.OrderPayment),                 // string
							escape(o.ModulePayment),                // string
							o.TotalPaid,                            // float64
							o.TotalDiscounts,                       // float64
							o.TotalDiscountsTaxIncl,                // float64
							o.TotalDiscountsTaxExcl,                // float64
							o.TotalPaidTaxIncl,                     // float64
							o.TotalPaidTaxExcl,                     // float64
							o.TotalPaidReal,                        // float64
							fmt.Sprintf(string(o.TotalProducts)),   // uint
							fmt.Sprintf(string(o.TotalProductsWt)), // uint
							o.TotalShipping,                        // float64
							o.TotalShippingTaxIncl,                 // float64
							o.TotalShippingTaxExcl,                 // float64
							o.TotalWrapping,                        // float64
							o.TotalWrappingTaxIncl,                 // float64
							o.TotalWrappingTaxExcl,                 // float64
							o.RoundMode,                            // uint
							o.RoundType,                            // uint
							o.InvoiceNumber,                        // uint
							escape(o.DeliveryNumber),               // string
							o.InvoiceDate.Unix(),                   // time.Time
							o.DeliveryDate.Unix(),                  // time.Time
							o.Valid,                                // uint
							o.DateAdd.Unix(),                       // time.Time
							o.DateUpd.Unix(),                       // time.Time
							o.CarrierTaxRate,                       // float64
							o.IdAddressDelivery,                    // uint
							escape(o.CustomerName),                 // string
							escape(o.OrderStatusName),              // string
							escape(o.OrderStatusColor),             // string
							o.OrderNew,                             // uint
							o.IdCountry,                            // uint
							escape(o.CountryName),                  // string
							escape(o.Company),                      // string
							escape(o.VatNumber),                    // string
							escape(o.PhoneMobile),                  // string
							escape(o.Phone),                        // string
							escape(o.DeliveryLastname),             // string
							escape(o.DeliveryFirstname),            // string
							escape(o.Address1),                     // string
							escape(o.Address2),                     // string
							escape(o.Postcode),                     // string
							escape(o.City),                         // string
							o.BadgeSuccess,                         // uint
						)

						_, err = cl.Exec(query)
						if err != nil {
							log.Infoln(query)
							log.Fatal(err)
							return err
						}
						log.Infoln("Index order >> ", o.IdOrder, "==", o.CustomerName)
						return nil
					}(order)

					t.Throttle()

				}

				if t.Err() != nil {
					// Loop through the errors to see the details
					for i, err := range t.Errs() {
						log.Printf("error #%d: %s", i, err)
					}
					log.Fatal(t.Err())
				}

			}
		case "customers", "customer":

			// todo. create cli argument for limit
			limit := 10000
			offset := 0

			for {

				sqlResult := bytes.NewBufferString("")
				sqlTemplate, _ := template.New("").Parse(sqlCustomerExtended)
				sqlTemplate.Execute(sqlResult, map[string]string{"prefix": dbTablePrefix, "offset": fmt.Sprintf("%d", offset), "limit": fmt.Sprintf("%d", limit)})

				var customers []*Customer
				err = db.Raw(sqlResult.String()).Scan(&customers).Error
				if err != nil {
					log.Fatal(err)
				}

				offset = offset + limit

				if len(customers) == 0 {
					break
				}

				t := throttler.New(1, len(customers))

				for _, customer := range customers {

					go func(c *Customer) error {
						// Let Throttler know when the goroutine completes
						// so it can dispatch another worker
						defer t.Done(nil)

						query := fmt.Sprintf(`REPLACE into rt_customers (
							id,
							id_customer, 
							firstname, 
							lastname, 
							fullname,
							note,				
							email, 
							active, 
							deleted,
							newsletter, 
							optin, 
							ip_registration_newsletter,
							birthday, 
							date_add, 
							social_title,
							id_group,
							group_name,
							shop_name, 
							company, 
							total_spent, 
							connect				
							) VALUES ('%d','%d','%s','%s','%s','%s','%s','%d','%d','%d','%d','%s','%d','%d','%s','%d','%s','%s','%s','%.2f','%d')`,
							c.IdCustomer,                       //  uint
							c.IdCustomer,                       //  uint
							escape(c.Firstname),                //  string
							escape(c.Lastname),                 //   string
							escape(c.Fullname),                 //   string
							escape(c.Note),                     //   string
							escape(c.Email),                    //   string
							c.Active,                           //    uint
							c.Deleted,                          //   uint
							c.Newsletter,                       //   uint
							c.Optin,                            // uint
							escape(c.IpRegistrationNewsletter), // string
							c.Birthday.Unix(),                  // string
							c.DateAdd.Unix(),                   //  time.Time
							escape(c.SocialTitle),              //  string
							c.IdGroup,                          //  uint
							escape(c.GroupName),                //   string
							escape(c.ShopeName),                //    string
							escape(c.Company),                  //    string
							c.TotalSpent,                       //    float64
							c.Connect.Unix(),                   //    time.Time
						)

						_, err = cl.Exec(query)
						if err != nil {
							log.Infoln(query)
							log.Fatal(err)
							return err
						}
						log.Infoln("Index customer >> ", c.IdCustomer, "==", c.Fullname)
						return nil
					}(customer)

					t.Throttle()

				}

				if t.Err() != nil {
					// Loop through the errors to see the details
					for i, err := range t.Errs() {
						log.Printf("error #%d: %s", i, err)
					}
					log.Fatal(t.Err())
				}
			}

		}

	},
}

func init() {
	// todo. manage ssl connections for db
	ImportCmd.Flags().StringVarP(&indexName, "index-name", "", "products", "realt-time index name to index")
	ImportCmd.Flags().StringVarP(&dbTablePrefix, "db-table-prefix", "", "ps_", "database table prefix")
	ImportCmd.Flags().StringVarP(&dbName, "db-name", "", "prestashop", "database name")
	ImportCmd.Flags().StringVarP(&dbUser, "db-user", "", "root", "database username")
	ImportCmd.Flags().StringVarP(&dbPass, "db-pass", "", "", "database password")
	ImportCmd.Flags().StringVarP(&dbHost, "db-host", "", "127.0.0.1", "database host")
	ImportCmd.Flags().StringVarP(&dbPort, "db-port", "", "3306", "datbase port")
	ImportCmd.Flags().StringVarP(&psDir, "ps-dir", "p", "../../shared/www", "prestashop directory")
	ImportCmd.Flags().BoolVarP(&dryRun, "dry-run", "", false, "dry run")
	RootCmd.AddCommand(ImportCmd)
}

func mysqlRealEscapeString(value string) string {
	replace := map[string]string{"\\": "\\\\", "'": `\'`, "\\0": "\\\\0", "\n": "\\n", "\r": "\\r", `"`: `\"`, "\x1a": "\\Z"}

	for b, a := range replace {
		value = strings.Replace(value, b, a, -1)
	}

	return value
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

func stripslashes(str string) string {
	var dstRune []rune
	strRune := []rune(str)
	strLenth := len(strRune)
	for i := 0; i < strLenth; i++ {
		if strRune[i] == []rune{'\\'}[0] {
			i++
		}
		dstRune = append(dstRune, strRune[i])
	}
	return string(dstRune)
}
