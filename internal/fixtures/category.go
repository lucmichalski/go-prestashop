package fixtures

import (
	"encoding/xml"
)

// Category was generated 2021-02-06 11:37:00 by evolutive on eg-cdn.gsi-network.com.
type Category struct {
	XMLName xml.Name `xml:"category"`
	Text    string   `xml:",chardata"`
	Fields  struct {
		Text  string `xml:",chardata"`
		ID    string `xml:"id,attr"`
		Class string `xml:"class,attr"`
		Sql   string `xml:"sql,attr"`
		Image string `xml:"image,attr"`
		Field []struct {
			Text     string `xml:",chardata"`
			Name     string `xml:"name,attr"`
			Relation string `xml:"relation,attr"`
		} `xml:"field"`
	} `xml:"fields"`
	Entities struct {
		Text     string `xml:",chardata"`
		Category []struct {
			Text           string `xml:",chardata"`
			IDParent       string `xml:"id_parent,attr"`
			ID             string `xml:"id,attr"`
			Active         string `xml:"active,attr"`
			IsRootCategory string `xml:"is_root_category,attr"`
		} `xml:"category"`
	} `xml:"entities"`
}
