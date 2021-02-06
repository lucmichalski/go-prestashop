package fixtures

import (
	"encoding/xml"
)

// CategoryGroup was generated 2021-02-06 11:40:20 by evolutive on eg-cdn.gsi-network.com.
type CategoryGroup struct {
	XMLName xml.Name `xml:"category_group"`
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
		Text          string `xml:",chardata"`
		CategoryGroup []struct {
			Text       string `xml:",chardata"`
			IDCategory string `xml:"id_category,attr"`
			IDGroup    string `xml:"id_group,attr"`
			ID         string `xml:"id,attr"`
		} `xml:"category_group"`
	} `xml:"entities"`
}
