package fixtures

import (
	"encoding/xml"
)

// Supplier was generated 2021-02-06 11:40:10 by evolutive on eg-cdn.gsi-network.com.
type Supplier struct {
	XMLName xml.Name `xml:"entity_supplier"`
	Text    string   `xml:",chardata"`
	Fields  struct {
		Text  string `xml:",chardata"`
		ID    string `xml:"id,attr"`
		Class string `xml:"class,attr"`
		Image string `xml:"image,attr"`
		Field []struct {
			Text string `xml:",chardata"`
			Name string `xml:"name,attr"`
		} `xml:"field"`
	} `xml:"fields"`
	Entities struct {
		Text     string `xml:",chardata"`
		Supplier []struct {
			Text   string `xml:",chardata"`
			ID     string `xml:"id,attr"`
			Name   string `xml:"name,attr"`
			Active string `xml:"active,attr"`
		} `xml:"supplier"`
	} `xml:"entities"`
}
