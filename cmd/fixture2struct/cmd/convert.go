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
	"github.com/miku/zek"
	"github.com/spf13/cobra"
)

var (
	packageName          string
	outputDir            string
	workDir              string
	dryRun               bool
	withComments         bool
	withJSONTags         bool
	maxExamples          int
	debug                bool
	createExampleProgram bool
	tagName              string
	skipFormatting       bool
	strict               bool
	exampleMaxChars      int
	version              bool
	structName           string
	compact              bool
	nonCompact           bool
	uniqueExamples       bool
	omitEmptyText        bool
)

const (
	tmplPackage = `package fixtures

import (
	"encoding/xml"
)

`
)

var ConvertCmd = &cobra.Command{
	Use:     "convert",
	Aliases: []string{"c"},
	Short:   "convert prestashop-shop-creator xml to golang structs.",
	Long:    "convert prestashop-shop-creator xml to golang structs",
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

						root := new(zek.Node)
						root.MaxExamples = maxExamples

						reader := bufio.NewReader(file)
						if _, err := root.ReadFrom(reader); err != nil {
							log.Warnf("could not read '%s', error: %s", osPathname, err)
							return err
						}
						// Move root, if we have a tagName. Ignore unknown names.
						if tagName != "" {
							if n := root.ByName(tagName); n != nil {
								root = n
							}
						}
						root.Name = xml.Name{Space: "", Local: pkgName}

						var buf bytes.Buffer
						io.WriteString(&buf, tmplPackage)
						sw := zek.NewStructWriter(&buf)
						sw.WithComments = withComments
						sw.WithJSONTags = withJSONTags
						sw.Strict = strict
						sw.ExampleMaxChars = exampleMaxChars
						sw.Compact = !nonCompact
						sw.UniqueExamples = uniqueExamples
						sw.OmitEmptyText = omitEmptyText
						if err := sw.WriteNode(root); err != nil {
							log.Fatal(err)
						}

						f, err := os.Create(filepath.Join(outputDir, fmt.Sprintf("%s.go", pkgName)))
						if err != nil {
							log.Fatal(err)
						}
						defer f.Close()

						if !skipFormatting {
							b, err := format.Source(buf.Bytes())
							if err != nil {
								log.Fatal(err)
							}
							_, err = f.Write(b)
							if err != nil {
								log.Fatal(err)
							}
						} else {
							_, err := f.WriteString(buf.String())
							if err != nil {
								log.Fatal(err)
							}
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
	ConvertCmd.Flags().StringVarP(&packageName, "package-name", "", "", "package name")
	ConvertCmd.Flags().StringVarP(&outputDir, "output-dir", "", "../../internal/fixtures", "output package directory")
	ConvertCmd.Flags().StringVarP(&workDir, "work-dir", "w", "../../shared/fixtures/data", "working directory")
	ConvertCmd.Flags().BoolVarP(&dryRun, "dry-run", "", false, "dry run")
	ConvertCmd.Flags().BoolVarP(&withComments, "with-comments", "", false, "add comments with example")
	ConvertCmd.Flags().BoolVarP(&withJSONTags, "with-json-tags", "j", false, "add JSON tags")
	ConvertCmd.Flags().IntVarP(&maxExamples, "max-examples", "", 10, "limit number of examples")
	ConvertCmd.Flags().StringVarP(&tagName, "tag-name", "t", "", "emit struct for tag matching this name")
	ConvertCmd.Flags().BoolVarP(&skipFormatting, "skip-formatting", "", false, "skip formatting")
	ConvertCmd.Flags().BoolVarP(&strict, "strict", "", false, "strict parsing and writing")
	ConvertCmd.Flags().IntVarP(&exampleMaxChars, "", "x", 25, "max chars for example")
	ConvertCmd.Flags().StringVarP(&structName, "struct-name", "n", "fixtures", "use a different name for the top-level struct")
	ConvertCmd.Flags().BoolVarP(&compact, "compact", "", false, "emit more compact struct (noop, as this is the default since 0.1.7)")
	ConvertCmd.Flags().BoolVarP(&nonCompact, "non-compact", "", false, "emit less compact struct")
	ConvertCmd.Flags().BoolVarP(&uniqueExamples, "unique-examples", "", false, "filter out duplicated examples")
	ConvertCmd.Flags().BoolVarP(&omitEmptyText, "omit-empty-text", "", false, "omit empty Text fields")
	RootCmd.AddCommand(ConvertCmd)
}
