package cmd

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	// "errors"
	"fmt"
	"io"
	"os"
	"strings"

	// "github.com/k0kubun/pp"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
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
	Use:   "csv2ps",
	Short: "csv2ps is an helper to load csv catalog into a prestashop database.",
	Long:  `csv2ps is an helper to load csv catalog into a prestashop database.`,
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

// fail fast on initialization
func must(i interface{}, err error) interface{} {
	if err != nil {
		panic(err)
	}
	return i
}

func csv2json(fp string, columns []string, separator string) ([]byte, error) {
	rows := make([]map[string]string, 0)

	// read
	file, err := os.Open(fp)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	r := csv.NewReader(file)
	if separator == "" {
		separator = ";"
	}
	r.Comma = []rune(separator)[0]
	r.LazyQuotes = true
	r.Comment = '#'

	// csvReader := csv.NewReader(ReplaceSoloCarriageReturns(r))
	// csvReader.TrimLeadingSpace = true
	// csvReader.LazyQuotes = true
	// csvReader.Comment = '#'

	fmt.Println("separator=", separator)

	z := 0
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			if perr, ok := err.(*csv.ParseError); ok && perr.Err == csv.ErrFieldCount {
				continue
			}
			return nil, err
		}
		// pp.Println(z, "==", record)
		row := make(map[string]string, len(columns))

		// if len(record) != len(columns) {
		// 	continue
		// }

		for i, n := range columns {
			row[n] = record[i] // record[i]
		}
		rows = append(rows, row)
		z++
	}

	data, err := json.MarshalIndent(&rows, "", "  ")
	if err != nil {
		return nil, err
	}
	return data, nil
}

// ReplaceSoloCarriageReturns wraps an io.Reader, on every call of Read it
// for instances of lonely \r replacing them with \r\n before returning to the end customer
// lots of files in the wild will come without "proper" line breaks, which irritates go's
// standard csv package. This'll fix by wrapping the reader passed to csv.NewReader:
//    rdr, err := csv.NewReader(ReplaceSoloCarriageReturns(r))
//
func ReplaceSoloCarriageReturns(data io.Reader) io.Reader {
	return crlfReplaceReader{
		rdr: bufio.NewReader(data),
	}
}

// crlfReplaceReader wraps a reader
type crlfReplaceReader struct {
	rdr *bufio.Reader
}

// Read implements io.Reader for crlfReplaceReader
func (c crlfReplaceReader) Read(p []byte) (n int, err error) {
	if len(p) == 0 {
		return
	}

	for {
		if n == len(p) {
			return
		}

		p[n], err = c.rdr.ReadByte()
		if err != nil {
			return
		}

		// any time we encounter \r & still have space, check to see if \n follows
		// if next char is not \n, add it in manually
		if p[n] == '\r' && n < len(p) {
			if pk, err := c.rdr.Peek(1); (err == nil && pk[0] != '\n') || (err != nil && err.Error() == io.EOF.Error()) {
				n++
				p[n] = '\n'
			}
		}

		n++
	}
	return
}
