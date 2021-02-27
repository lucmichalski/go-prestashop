package cmd

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	// "github.com/beevik/etree"
	"github.com/k0kubun/pp"
	// "github.com/nozzle/throttler"
	"github.com/spf13/cobra"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	// "github.com/lucmichalski/go-prestashop/internal/pluck"
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
)

// go run main.go scan --db-name eg_prestanish --db-user root --db-pass "OvdZ5ZoXWgCWL4-hvZjg!" --db-table-prefix eg_

var ScanCmd = &cobra.Command{
	Use:     "scan",
	Aliases: []string{"s"},
	Short:   "scan information about a prestashop deployment.",
	Long:    "scan information about a prestashop deployment",
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

			modulePath := filepath.Join(psDir, "modules")
			files, err := ioutil.ReadDir(modulePath)
			if err != nil {
				log.Fatal(err)
			}

			for _, f := range files {
				if (f.Name() != "." || f.Name() != "..") && f.IsDir() {
					pp.Println(f.Name())
					// check if confix.xml exists

				}
			}

		}

	},
}

func init() {
	// todo. manage ssl connections for db
	ScanCmd.Flags().StringVarP(&dbTablePrefix, "db-table-prefix", "", "ps_", "database table prefix")
	ScanCmd.Flags().StringVarP(&dbName, "db-name", "", "prestashop", "database name")
	ScanCmd.Flags().StringVarP(&dbUser, "db-user", "", "root", "database username")
	ScanCmd.Flags().StringVarP(&dbPass, "db-pass", "", "", "database password")
	ScanCmd.Flags().StringVarP(&dbHost, "db-host", "", "127.0.0.1", "database host")
	ScanCmd.Flags().StringVarP(&dbPort, "db-port", "", "3306", "datbase port")
	ScanCmd.Flags().StringVarP(&psDir, "ps-dir", "p", "../../../evolutive-prestashop/shared/www", "prestashop directory")
	ScanCmd.Flags().BoolVarP(&dryRun, "dry-run", "", false, "dry run")
	RootCmd.AddCommand(ScanCmd)
}
