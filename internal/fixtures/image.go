package fixtures

import (
	"encoding/xml"
)

// Image was generated 2021-02-06 11:17:10 by evolutive on eg-cdn.gsi-network.com.
type Image struct {
	XMLName xml.Name `xml:"image"`
	Text    string   `xml:",chardata"`
	Fields  struct {
		Text    string `xml:",chardata"`
		Primary string `xml:"primary,attr"`
		Class   string `xml:"class,attr"`
		Image   string `xml:"image,attr"`
		Field   []struct {
			Text     string `xml:",chardata"`
			Name     string `xml:"name,attr"`
			Relation string `xml:"relation,attr"`
		} `xml:"field"`
	} `xml:"fields"`
	Entities struct {
		Text  string `xml:",chardata"`
		Image []struct {
			Text      string `xml:",chardata"`
			IDProduct string `xml:"id_product,attr"`
			ID        string `xml:"id,attr"`
			Cover     string `xml:"cover,attr"`
		} `xml:"image"`
	} `xml:"entities"`
}
