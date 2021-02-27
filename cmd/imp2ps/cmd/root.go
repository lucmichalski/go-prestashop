package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"

	"github.com/sirupsen/logrus"
)

var options struct {
	verbose      bool
	debug        bool
	logLevel     string
	generateDocs bool
}

var (
	log = logrus.New()
)

// RootCmd is the root command for ovh-qa
var RootCmd = &cobra.Command{
	Use:   "imp2ps",
	Short: "imp2ps is an helper to create products into webkul marketplace.",
	Long:  `imp2ps is an helper to create products into webkul marketplace.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	if options.generateDocs {
		err := doc.GenMarkdownTree(RootCmd, "../../docs")
		if err != nil {
			log.Fatalf("GenMarkdownTree: %s\n", err)
		}
	}
}

func init() {
	flags := RootCmd.PersistentFlags()
	flags.BoolVarP(&options.generateDocs, "generate-docs", "g", false, "generate documentation.")
	flags.BoolVarP(&options.debug, "debug", "b", false, "debug mode.")
	flags.BoolVarP(&options.verbose, "verbose", "v", false, "verbose output.")
}

func checkErr(message string, err error) {
	if err != nil {
		log.Fatalf("%s: %s\n", message, err)
	}
}

func inSlice(input string, dictionary []string, caseSensitive bool) bool {
	for _, entry := range dictionary {
		if !caseSensitive {
			input = strings.ToLower(input)
			entry = strings.ToLower(entry)
		}
		if entry == input {
			return true
		}
	}
	return false
}

func inSliceContains(input string, dictionary []string, caseSensitive bool) bool {
	for _, entry := range dictionary {
		if !caseSensitive {
			input = strings.ToLower(input)
			entry = strings.ToLower(entry)
		}
		if strings.Contains(input, entry) {
			return true
		}
	}
	return false
}

// fail fast on initialization
func must(i interface{}, err error) interface{} {
	if err != nil {
		panic(err)
	}
	return i
}
