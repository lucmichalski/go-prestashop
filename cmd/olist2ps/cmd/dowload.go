package cmd

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os/exec"
	"strings"
)

func install() {
	cmd := exec.Command("pip3", "install", "-U", "kaggle")
	cmb, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Println(string(cmb))
}

// $ kaggle datasets list --help
// usage: kaggle datasets list [-h] [--sort-by SORT_BY] [--size SIZE]
//                             [--file-type FILE_TYPE] [--license LICENSE_NAME]
//                             [--tags TAG_IDS] [-s SEARCH] [-m] [--user USER]
//                             [-p PAGE] [-v] [--max-size MAX_SIZE]
//                             [--min-size MIN_SIZE]
//
// optional arguments:
//   -h, --help            show this help message and exit
//   --sort-by SORT_BY     Sort list results. Default is 'hottest'. Valid options are 'hottest', 'votes', 'updated', and 'active'
//   --size SIZE           DEPRECATED. Please use --max-size and --min-size to filter dataset sizes.
//   --file-type FILE_TYPE
//                         Search for datasets with a specific file type. Default is 'all'. Valid options are 'all', 'csv', 'sqlite', 'json', and 'bigQuery'. Please note that bigQuery datasets cannot be downloaded
//   --license LICENSE_NAME
//                         Search for datasets with a specific license. Default is 'all'. Valid options are 'all', 'cc', 'gpl', 'odb', and 'other'
//   --tags TAG_IDS        Search for datasets that have specific tags. Tag list should be comma separated
//   -s SEARCH, --search SEARCH
//                         Term(s) to search for
//   -m, --mine            Display only my items
//   --user USER           Find public datasets owned by a specific user or organization
//   -p PAGE, --page PAGE  Page number for results paging. Page size is 20 by default
//   -v, --csv             Print results in CSV format (if not set print in table format)
//   --max-size MAX_SIZE   Specify the maximum size of the dataset to return (bytes)
//   --min-size MIN_SIZE   Specify the minimum size of the dataset to return (bytes)
//
func listByVotesPopularity(minSize, maxSize, maxEntries int) (entries [][]string) {
	// kaggle datasets list --csv --file-type csv --min-size 1024000000 --max-size 1096000000 -p 1
	args := make([]string, 0, 16)
	args = append(args, "datasets")
	args = append(args, "list")
	if true {
		args = append(args, "--csv")
	}
	args = append(args, "--sort-by")
	args = append(args, "votes")
	args = append(args, "--file-type")
	args = append(args, "csv")
	args = append(args, "--min-size")
	args = append(args, fmt.Sprintf("%d", minSize))
	args = append(args, "--max-size")
	args = append(args, fmt.Sprintf("%d", maxSize))
	args = append(args, "-p")

	entries = make([][]string, 0)

	for page := 1; ; page++ {

		args = append(args, fmt.Sprintf("%d", page))
		cmd := exec.Command("kaggle", args...)
		args = args[:len(args)-1]

		cmb, err := cmd.CombinedOutput()
		if err != nil {
			log.Fatalf("%v", err)
		}
		lines := strings.Split(string(cmb), "\n")

		if len(lines) >= 1 && lines[0] == "No datasets found" {
			break
		}
		if len(lines) >= 2 {
			r := csv.NewReader(strings.NewReader(strings.Join(lines[1:], "\n")))

			for {
				record, err := r.Read()
				if err == io.EOF {
					break
				}
				if err != nil {
					log.Fatal(err)
				}

				// fmt.Println(record)
				entries = append(entries, record)
			}
		}

		if len(entries) >= maxEntries {
			break
		}
	}

	return
}

// $ kaggle datasets files --help
// usage: kaggle datasets files [-h] [-v] [dataset]
//
// optional arguments:
//   -h, --help  show this help message and exit
//   dataset     Dataset URL suffix in format <owner>/<dataset-name> (use "kaggle datasets list" to show options)
//   -v, --csv   Print results in CSV format (if not set print in table format)
//
func files(dataset string) {
	// kaggle datasets files --csv kentonnlp/2014-new-york-city-taxi-trips
	cmd := exec.Command("kaggle", "datasets", "files", "--csv", dataset)
	cmb, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Println(string(cmb))
}

// $ kaggle datasets download --help
// usage: kaggle datasets download [-h] [-f FILE_NAME] [-p PATH] [-w] [--unzip]
//                                 [-o] [-q]
//                                 [dataset]
//
// optional arguments:
//   -h, --help            show this help message and exit
//   dataset               Dataset URL suffix in format <owner>/<dataset-name> (use "kaggle datasets list" to show options)
//   -f FILE_NAME, --file FILE_NAME
//                         File name, all files downloaded if not provided
//                         (use "kaggle datasets files -d <dataset>" to show options)
//   -p PATH, --path PATH  Folder where file(s) will be downloaded, defaults to current working directory
//   -w, --wp              Download files to current working path
//   --unzip               Unzip the downloaded file. Will delete the zip file when completed.
//   -o, --force           Skip check whether local version of file is up to date, force file download
//   -q, --quiet           Suppress printing information about the upload/download progress
func download(dataset, path string) {
	// kaggle datasets download --unzip
	cmd := exec.Command("kaggle", "datasets", "download", "--path", path, "--unzip", dataset)
	cmb, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Println(string(cmb))
}
