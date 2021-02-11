package cmd

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/nozzle/throttler"
	"github.com/spf13/cobra"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// go run main.go index --db-name eg_prestanish --db-user root --db-pass "OvdZ5ZoXWgCWL4-hvZjg!" --db-table-prefix eg_
var (
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

		var products []*Product
		err = db.Raw(sqlAdminCatalogProducts).Scan(&products).Error
		if err != nil {
			log.Fatal(err)
		}

		t := throttler.New(6, len(products))

		for _, product := range products {

			go func(p *Product) error {
				// Let Throttler know when the goroutine completes
				// so it can dispatch another worker
				defer t.Done(nil)

				// create sub query for features

				query := fmt.Sprintf(`REPLACE into rt_products (
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
					name_category
					) VALUES ('%s','%d','%d','%d','%.2f','%d','%t','%s','%t','%s','%d','%d','%.2f','%d','%d','%t',(%s),(%s),'%s','%s','%s','%s','%s')`,
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
				)
				// pp.Println("query:", query)
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
		//'-', '!','=']
		//case '[':
		//	escape = '['
		//	break
		//case ']':
		//	escape = ']'
		//	break
		//case '^':
		//	escape = '^'
		//	break
		//case '~':
		//	escape = '~'
		//	break
		//case '%':
		//	escape = '%'
		//	break
		//case '@':
		//	escape = '@'
		//	break
		//case '|':
		//	escape = '|'
		//	break
		//case '&':
		//	escape = '&'
		//	break
		//case '/':
		//	escape = '/'
		//	break
		//case '$':
		//	escape = '$'
		//	break
		//case ')':
		//	escape = ')'
		//	break
		//case '(':
		//	escape = '('
		//	break
		//case '”':
		//	escape = '”'
		//	break
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
