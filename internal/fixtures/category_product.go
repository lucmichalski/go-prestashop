package fixtures

import (
	"encoding/xml"
)

// CategoryProduct was generated 2021-02-06 11:40:10 by evolutive on eg-cdn.gsi-network.com.
type CategoryProduct struct {
	XMLName xml.Name `xml:"category_product"`
	Text    string   `xml:",chardata"`
	Fields  struct {
		Text    string `xml:",chardata"`
		Primary string `xml:"primary,attr"`
		Field   []struct {
			Text     string `xml:",chardata"`
			Name     string `xml:"name,attr"`
			Relation string `xml:"relation,attr"`
		} `xml:"field"`
	} `xml:"fields"`
	Entities struct {
		Text            string `xml:",chardata"`
		CategoryProduct []struct {
			Text       string `xml:",chardata"`
			IDCategory string `xml:"id_category,attr"`
			IDProduct  string `xml:"id_product,attr"`
			ID         string `xml:"id,attr"`
		} `xml:"category_product"`
	} `xml:"entities"`
}
