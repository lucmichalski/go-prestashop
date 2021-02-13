package cmd

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/dghubble/sling"
	"github.com/k0kubun/pp"
	"github.com/spf13/cobra"
)

var (
	matchImageUrl        string
	matchImagePath       string
	matchAllOrientations bool
)

var SearchCmd = &cobra.Command{
	Use:     "search",
	Aliases: []string{"s"},
	Short:   "Searches for a similar image in the database. Scores range from 0 to 100, with 100 being a perfect match.",
	Long:    "Searches for a similar image in the database. Scores range from 0 to 100, with 100 being a perfect match.",
	Run: func(cmd *cobra.Command, args []string) {

		if matchImageUrl != "" && matchImagePath != "" {
			log.Fatal("Error, you cannot use url and image argument at the same time.")
		}

		params := &Search{
			AllOrientations: matchAllOrientations,
		}

		if matchImageUrl != "" {
			params.Url = matchImageUrl
		}

		// todo. multipart form
		if matchImagePath != "" {
			params.Image = matchImagePath
		}

		pp.Println("params", params)

		// todo: add timeout
		httpClient := &http.Client{}

		// Preparing request
		searchBase := sling.New().Base(matchAddr).Client(httpClient)
		req, err := searchBase.New().Post("search").BodyForm(params).Request()
		if err != nil {
			log.Fatal(err)
		}

		// get response
		resp, err := httpClient.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		// parsing response
		var response MatchSearch
		err = json.Unmarshal(body, &response)
		if err != nil {
			log.Fatal(err)
		}

		// Results from dsys match
		pp.Println(response)

	},
}

func init() {
	SearchCmd.Flags().StringVarP(&matchAddr, "match-addr", "", "http://localhost:8888", "image match service address")
	SearchCmd.Flags().StringVarP(&matchImageUrl, "url", "", "", "url of the image to compare.")
	SearchCmd.Flags().StringVarP(&matchImagePath, "image", "", "", "filepath of the image to compare.")
	SearchCmd.Flags().BoolVarP(&matchAllOrientations, "all-orientations", "", false, "whether or not to search for similar 90 degree rotations of the image.")
	RootCmd.AddCommand(SearchCmd)
}
