package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/beevik/etree"
	"github.com/k0kubun/pp"
	"gopkg.in/yaml.v2"
	// "github.com/nozzle/throttler"
	"github.com/spf13/cobra"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	psm "github.com/lucmichalski/go-prestashop/internal/models"
)

var (
	psDir          string
	workDir        string
	dryRun         bool
	db             *gorm.DB
	dbName         string
	dbUser         string
	dbPass         string
	dbHost         string
	dbPort         string
	dbTablePrefix  string
	reDatabaseHost = regexp.MustCompile(`database_host' => '(.*)'`)
	reDatabasePort = regexp.MustCompile(`database_port' => '(.*)'`)
	reDatabasePass = regexp.MustCompile(`database_password' => '(.*)'`)
	reDatabaseUser = regexp.MustCompile(`database_user' => '(.*)'`)
	reDatabaseName = regexp.MustCompile(`database_name' => '(.*)'`)
	reDatabasePref = regexp.MustCompile(`database_prefix' => '(.*)'`)
	moduleConfExts = []string{".xml"}
	defaultDbHosts = []string{"localhost", "127.0.0.1"}
	psLangs        []*psm.Lang
	psShops        []*psm.Shop
	psGroups       []*psm.Group
	// psModules      []*psm.Module
)

const (
	defaultDbHost  = `127.0.0.1`
	defaultDbPort  = `3306`
	parametersFile = `./app/config/parameters.php`
)

// type Config struct {
// }

type Module struct {
	*psm.Module
	DisplayName      string
	Description      string
	Author           string
	Tab              string
	IsConfigurable   bool
	NeedInstance     bool
	LimitedCountries string
}

/*
	- https://www.nulledfrm.com/threads/b2b-private-shop-premium-module-latest.6572/
*/

// go run main.go scan

var ScanCmd = &cobra.Command{
	Use:     "scan",
	Aliases: []string{"s"},
	Short:   "scan information about a prestashop deployment.",
	Long:    "scan information about a prestashop deployment (version, modules, shops...) into a yaml report.",
	Run: func(cmd *cobra.Command, args []string) {

		parametersFilePath := filepath.Join(psDir, parametersFile)

		if _, err := os.Stat(parametersFilePath); err == nil {

			log.Infof("Found '%s' config file, reading content from it...\n", parametersFilePath)
			dat, err := ioutil.ReadFile(parametersFilePath)
			checkErr("ReadFile.error", err)

			databaseHost := reDatabaseHost.FindSubmatch(dat)
			if len(databaseHost) > 0 {
				dbHost = string(databaseHost[1])
			}

			if !inSlice(dbHost, defaultDbHosts, false) {
				dbHost = defaultDbHost
			}

			databasePort := reDatabasePort.FindSubmatch(dat)
			if len(databasePort) > 0 {
				dbPort = string(databasePort[1])
			}
			if dbPort == "" {
				dbPort = defaultDbPort
			}

			databaseName := reDatabaseName.FindSubmatch(dat)
			if len(databaseName) > 0 {
				dbName = string(databaseName[1])
			}

			databaseUser := reDatabaseUser.FindSubmatch(dat)
			if len(databaseUser) > 0 {
				dbUser = string(databaseUser[1])
			}

			databasePass := reDatabasePass.FindSubmatch(dat)
			if len(databasePass) > 0 {
				dbPass = string(databasePass[1])
			}

			databasePref := reDatabasePref.FindSubmatch(dat)
			if len(databasePref) > 0 {
				dbTablePrefix = string(databasePref[1])
			}

		} else {
			log.Warnf("Could not find '%s' config file, using command cli arguments...\n", parametersFilePath)
		}

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

			// get
			err = db.Where("active = ?", 1).Find(&psLangs).Error
			pp.Println("psLangs:", psLangs)

			err = db.Where("active = ?", 1).Find(&psShops).Error
			if errors.Is(err, gorm.ErrRecordNotFound) {
				log.Fatal("shops not found")
			}
			pp.Println("psShops:", psShops)

			err = db.Find(&psGroups).Error
			if errors.Is(err, gorm.ErrRecordNotFound) {
				log.Fatal("groups not found")
			}
			pp.Println("psGroups:", psGroups)

			psModules := make([]*Module, 0)
			err = db.Find(&psModules).Error
			if errors.Is(err, gorm.ErrRecordNotFound) {
				log.Fatal("modules not found")
			}

			modulesPath := filepath.Join(psDir, "modules")
			files, err := ioutil.ReadDir(modulesPath)
			if err != nil {
				log.Fatal(err)
			}

			for _, f := range files {
				if (f.Name() != "." || f.Name() != "..") && f.IsDir() {

					// pp.Println(f.Name())
					modulePath := filepath.Join(modulesPath, f.Name())
					xmlConfigs, _ := filepath.Glob(fmt.Sprintf("%s/config*.xml", modulePath))
					for _, xmlConfig := range xmlConfigs {
						log.Infoln("module xml config found: ", xmlConfig)

						doc := etree.NewDocument()
						if err := doc.ReadFromFile(xmlConfig); err != nil {
							panic(err)
						}

						root := doc.SelectElement("module")
						if root != nil {
							fmt.Println("ROOT element:", root.Tag)
						} else {
							log.Warnln("could not parse root element from: ", xmlConfig)
						}

						var nameStr string
						name := root.SelectElement("name")
						if name != nil {
							nameStr = name.Text()
							fmt.Printf("  NAME: %s \n", name.Text())
						}

						var displayNameStr string
						displayName := root.SelectElement("displayName")
						if displayName != nil {
							displayNameStr = displayName.Text()
							fmt.Printf("  DISPLAY_NAME: %s \n", displayName.Text())
						}

						// version := root.SelectElement("version")
						// if version != nil {
						// 	fmt.Printf("  VERSION: %s \n", version.Text())
						// }

						var descriptionStr string
						description := root.SelectElement("description")
						if description != nil {
							descriptionStr = description.Text()
							fmt.Printf("  DESCRIPTION: %s \n", description.Text())
						}

						var authorStr string
						author := root.SelectElement("author")
						if author != nil {
							authorStr = author.Text()
							fmt.Printf("  AUTHOR: %s \n", author.Text())
						}

						var tabStr string
						tab := root.SelectElement("tab")
						if tab != nil {
							tabStr = tab.Text()
							fmt.Printf("  TAB: %s \n", tab.Text())
						}

						var isConfigurableBool bool
						isConfigurable := root.SelectElement("is_configurable")
						if isConfigurable != nil {
							var err error
							isConfigurableBool, err = strconv.ParseBool(isConfigurable.Text())
							checkErr("ParseBool.Error", err)
							fmt.Printf("  IS_CONFIGURABLE: %s \n", isConfigurable.Text())
						}

						var needInstanceBool bool
						needInstance := root.SelectElement("need_instance")
						if needInstance != nil {
							var err error
							needInstanceBool, err = strconv.ParseBool(needInstance.Text())
							checkErr("ParseBool.Error", err)
							fmt.Printf("  NEED_INSTANCE: %s \n", needInstance.Text())
						}

						var limitedCountriesStr string
						limitedCountries := root.SelectElement("limited_countries")
						if limitedCountries != nil {
							limitedCountriesStr = limitedCountries.Text()
							fmt.Printf("  LIMITED_COUNTRIES: %s \n", limitedCountries.Text())
						}

						for i, psModule := range psModules {
							if psModule.Module.Name == f.Name() && nameStr == f.Name() {
								psModules[i].DisplayName = displayNameStr
								psModules[i].Description = descriptionStr
								psModules[i].Author = authorStr
								psModules[i].Tab = tabStr
								psModules[i].IsConfigurable = isConfigurableBool
								psModules[i].NeedInstance = needInstanceBool
								psModules[i].LimitedCountries = limitedCountriesStr
							}
						}
					}

					// pp.Println("psModules:", psModules)
					d, err := yaml.Marshal(&psModules)
					if err != nil {
						log.Fatalf("error: %v", err)
					}
					fmt.Printf("--- t dump:\n%s\n\n", string(d))

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
