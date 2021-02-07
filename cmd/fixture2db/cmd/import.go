package cmd

import (
	"bufio"
	"bytes"
	"encoding/xml"
	"fmt"
	"go/format"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/karrick/godirwalk"
	"github.com/spf13/cobra"

	"github.com/lucmichalski/go-prestashop/internal/fixtures"
	"github.com/lucmichalski/go-prestashop/internal/fixtures/langs"
	"github.com/lucmichalski/go-prestashop/internal/models"
)

var (
	workDir string
	dryRun  bool
)

var ImportCmd = &cobra.Command{
	Use:     "import",
	Aliases: []string{"i"},
	Short:   "import prestashop-shop-creator xml to mysql database.",
	Long:    "import prestashop-shop-creator xml to mysql database",
	Run: func(cmd *cobra.Command, args []string) {

		err := godirwalk.Walk(workDir, &godirwalk.Options{
			Callback: func(osPathname string, de *godirwalk.Dirent) error {
				if !de.IsDir() {
					// get the extension from the file path
					extension := filepath.Ext(osPathname)
					filename := path.Base(osPathname)
					pkgName := strings.Replace(filename, extension, "", -1)
					if extension == ".xml" {
						if options.verbose {
							log.Println("found ", osPathname, "extension", extension)
						}

						file, err := os.Open(osPathname)
						if err != nil {
							return err
						}

					}
				}
				return nil
			},
			Unsorted: true,
		})
		checkErr("Dir walk, error", err)

	},
}

func init() {
	ImportCmd.Flags().StringVarP(&workDir, "work-dir", "w", "../../shared/fixtures/data", "working directory")
	ImportCmd.Flags().BoolVarP(&dryRun, "dry-run", "", false, "dry run")
	RootCmd.AddCommand(ImportCmd)
}
