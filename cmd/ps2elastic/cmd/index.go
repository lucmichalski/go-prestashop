package cmd

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/k0kubun/pp"
	"github.com/olivere/elastic/v7"
	"github.com/spf13/cobra"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

/*
	Refs:
	- curl -s -XGET 'http://127.0.0.1:9200/_nodes/http?pretty=1
	- curl -X GET 'http://localhost:9200/prestashop/_search'?pretty=true
	- go run main.go index --db-name YOUR_DATABASE --db-user root --db-pass YOUR_PASSWORD --db-table-prefix YOUR_DABASE_PREFIX
*/

var (
	dryRun        bool
	db            *gorm.DB
	dbName        string
	dbUser        string
	dbPass        string
	dbHost        string
	dbPort        string
	dbTablePrefix string
	elasticHost   string
	elasticPort   string
	elasticIndex  string
)

var IndexCmd = &cobra.Command{
	Use:     "index",
	Aliases: []string{"i"},
	Short:   "index prestashop products and categories to elasticsearch.",
	Long:    "index prestashop products and categories to elasticsearch",
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

		// Create a client
		esAddress := fmt.Sprintf("http://%s:%s", elasticHost, elasticPort)
		client, err := elastic.NewClient(elastic.SetURL(esAddress), elastic.SetSniff(false))
		checkErr("newClient", err)
		pp.Println("elastic new client connected")

		// Create an index
		_, _ = client.CreateIndex(elasticIndex).Do(context.Background())
		// checkErr("create Index", err)
		pp.Println("elastic CreateIndex ok")

		var products []*Product
		err = db.Debug().Raw(sqlProduct).Scan(&products).Error
		if err != nil {
			log.Fatal(err)
		}

		bulk := client.Bulk()
		ctx := context.Background()

		i := 0
		for _, product := range products {
			docinfo := fmt.Sprintf("%d", product.IdProduct)
			request := elastic.NewBulkIndexRequest().Index(elasticIndex).Type("doc").Id(docinfo).Doc(product)

			bulk.Add(request)
			pp.Println("bulk.Add:", docinfo)

			if i > 500 {
				policy := backoff.NewExponentialBackOff()
				policy.InitialInterval = time.Second
				operation := func() (err error) {
					_, err = bulk.Do(ctx)
					checkErr("es-bulk", err)
					// pp.Println("bulkResponse:", bulkResponse)
					return
				}
				if err := backoff.RetryNotify(operation, backoff.WithMaxRetries(policy, 5), func(err error, t time.Duration) { log.Info(err) }); err != nil {
					break
				} else {
					bulk = client.Bulk()
				}
				i = 0
			}
			i++
		}

		_, err = bulk.Do(ctx)
		checkErr("es-bulk", err)

	},
}

func init() {
	// todo. manage ssl connections for db
	IndexCmd.Flags().StringVarP(&dbTablePrefix, "db-table-prefix", "", "ps_", "database table prefix")
	IndexCmd.Flags().StringVarP(&dbName, "db-name", "", "prestashop", "database name")
	IndexCmd.Flags().StringVarP(&dbUser, "db-user", "", "root", "database username")
	IndexCmd.Flags().StringVarP(&dbPass, "db-pass", "", "", "database password")
	IndexCmd.Flags().StringVarP(&dbHost, "db-host", "", "127.0.0.1", "database host")
	IndexCmd.Flags().StringVarP(&dbPort, "db-port", "", "3306", "datbase port")
	// todo. manage https endpoints and doctypes
	IndexCmd.Flags().StringVarP(&elasticHost, "es-host", "", "127.0.0.1", "elastic host")
	IndexCmd.Flags().StringVarP(&elasticPort, "es-port", "", "9200", "elastic port")
	IndexCmd.Flags().StringVarP(&elasticIndex, "es-index", "", "prestashop", "elastic index name")
	IndexCmd.Flags().BoolVarP(&dryRun, "dry-run", "", false, "dry run")
	RootCmd.AddCommand(IndexCmd)
}
