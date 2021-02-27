package cmd

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"go/format"
	"io"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/cavaliercoder/grab"
	"github.com/k0kubun/pp"
	"github.com/miku/zek"
	"github.com/spf13/cobra"
)

var (
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
	tmplPackage = `package netaffiliation

import (
	"encoding/xml"
)

`
)

var ConvertCmd = &cobra.Command{
	Use:     "convert",
	Aliases: []string{"c"},
	Short:   "convert netaffiliation catalog to golang structs.",
	Long:    "convert netaffiliation catalog to golang structs",
	Run: func(cmd *cobra.Command, args []string) {

		client := grab.NewClient()

		// proxyURL
		catalogURL := "https://flux.netaffiliation.com/feed.php?lstv3=CD9B4220P1045983100L44D2F"

		tr := &http.Transport{
			// Proxy: http.ProxyURL(proxyURL),
			DialContext: (&net.Dialer{
				Timeout: 240 * time.Second,
				// KeepAlive: 30 * time.Second,
				DualStack: true,
			}).DialContext,
			MaxIdleConns:          100,
			IdleConnTimeout:       240 * time.Second,
			TLSHandshakeTimeout:   240 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
			DisableKeepAlives:     true,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		}

		client.UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.97 Safari/537.36"
		client.HTTPClient = &http.Client{
			Timeout:   240 * time.Second,
			Transport: tr,
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return nil
			},
		}

		catalogXmlFile := filepath.Join("catalogs", "netaffiliation.xml")

		if _, err := os.Stat(catalogXmlFile); os.IsNotExist(err) {

			req, err := grab.NewRequest(catalogXmlFile, catalogURL)
			if err != nil {
				log.Fatal(err)
				// continue
				//	return err
			}

			// start download
			fmt.Printf("Downloading %v...\n", req.URL())
			resp := client.Do(req)
			if resp.HTTPResponse == nil {
				// return fmt.Errorf("Could not download url=", req.URL())
				log.Fatal(err)
				// continue
				//	return err
			}
			fmt.Printf("  %v\n", resp.HTTPResponse.Status)

			// start UI loop
			t := time.NewTicker(500 * time.Millisecond)
			defer t.Stop()

		LoopRel:
			for {
				select {
				case <-t.C:
					fmt.Printf("  transferred %v / %v bytes (%.2f%%)\n",
						resp.BytesComplete(),
						resp.Size,
						100*resp.Progress())

				case <-resp.Done:
					// download is complete
					break LoopRel
				}
			}

			// check for errors
			if err := resp.Err(); err != nil {
				log.Fatalln("Download failed:", err, catalogURL)
				// continue
				//	return err
			}

			fmt.Printf("Download saved to %v \n", resp.Filename)
			catalogXmlFile = resp.Filename
		}

		file, err := os.Open(catalogXmlFile)
		if err != nil {
			log.Fatal(err)
			// continue
		}
		defer file.Close()

		root := new(zek.Node)
		root.MaxExamples = maxExamples

		reader := bufio.NewReader(file)
		if _, err := root.ReadFrom(reader); err != nil {
			log.Fatalf("could not read '%s', error: %s", catalogXmlFile, err)
			// return err
		}
		// Move root, if we have a tagName. Ignore unknown names.
		if tagName != "" {
			if n := root.ByName(tagName); n != nil {
				root = n
			}
		}
		root.Name = xml.Name{Space: "", Local: "feeds"}

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
		outputFile := filepath.Join(outputDir, fmt.Sprintf("%s.go", "feeds"))
		pp.Println("outputFile", outputFile)
		f, err := os.Create(outputFile)
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

			pp.Println(buf.String())
		}

	},
}

func init() {
	ConvertCmd.Flags().StringVarP(&outputDir, "output-dir", "", "../../internal/netaffiliation", "output package directory")
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
