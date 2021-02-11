package cmd

import (
	"fmt"
	"os"

	"github.com/k0kubun/pp"
	"github.com/spf13/cobra"
	"github.com/ziutek/mymysql/autorc"
	mmysql "github.com/ziutek/mymysql/mysql"
	_ "github.com/ziutek/mymysql/thrsafe"
)

// go run main.go search --term lorem
var (
	searchTerm string
)

var SearchCmd = &cobra.Command{
	Use:     "search",
	Aliases: []string{"s"},
	Short:   "search query on manticoresearch real-time index.",
	Long:    "search query on manticoresearch real-time index with facets",
	Run: func(cmd *cobra.Command, args []string) {

		query := "SELECT id,date_add,date_upd,shop_id,lang_id,category_id,category_link,manufacturer_id,afo,available_for_order,price,manufacturer,id_product,link_rewrite,HIGHLIGHT({passage_boundary='sentence'},'description') FROM rt_products WHERE MATCH('" + escape(searchTerm) + "') LIMIT 0,20 FACET shop_id FACET lang_id FACET price FACET manufacturer_id;"

		scan := func(rcount int, rows []mmysql.Row, res mmysql.Result) error {
			if res.StatusOnly() {
				return nil
			}

			pp.Println("rcount=", rcount)

			for idx, row := range rows {
				fmt.Print("#", rcount, "-", idx, ") ")
				for i := range row {
					fmt.Print(row.Str(i))
					fmt.Print(" ")
				}
				fmt.Println("")
			}
			return nil
		}

		pp.Println("query:", query)

		if err := RunSQL("127.0.0.1:9306", "root", "", "rt_products", query, scan); err != nil {
			fmt.Println(err)
		}

		os.Exit(1)

	},
}

func init() {
	SearchCmd.Flags().StringVarP(&searchTerm, "term", "", "", "dry run")
	RootCmd.AddCommand(SearchCmd)
}

type ScanFun func(int, []mmysql.Row, mmysql.Result) error

func RunSQL(hostport, user, pass, db, cmd string, scan ScanFun) error {

	conn := autorc.New("tcp", "", hostport, user, "", db)

	err := conn.Reconnect()
	if err != nil {
		return err
	}

	res, err := conn.Raw.Start(cmd)
	if err != nil {
		return err
	}

	rows, err := res.GetRows()
	if err != nil {
		return err
	}

	RScount := 0
	scanErr := error(nil)

	for {
		if scanErr == nil {
			func() {
				defer func() {
					if x := recover(); x != nil {
						scanErr = fmt.Errorf("%v", x)
					}
				}()
				scanErr = scan(RScount, rows, res)
			}()
		}

		if res.MoreResults() {
			res, err = res.NextResult()
			if err != nil {
				return err
			}
			rows, err = res.GetRows()
			if err != nil {
				return err
			}
		} else {
			break
		}

		RScount++
	}
	return scanErr
}
