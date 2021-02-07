package fixtures

import (
	"encoding/xml"
)

// Alias was generated 2021-02-06 11:40:20 by evolutive on eg-cdn.gsi-network.com.
type Alias struct {
	XMLName xml.Name `xml:"entity_alias"`
	Text    string   `xml:",chardata"`
	Fields  struct {
		Text  string `xml:",chardata"`
		ID    string `xml:"id,attr"`
		Class string `xml:"class,attr"`
		Field []struct {
			Text string `xml:",chardata"`
			Name string `xml:"name,attr"`
		} `xml:"field"`
	} `xml:"fields"`
	Entities struct {
		Text  string `xml:",chardata"`
		Alias []struct {
			Text   string `xml:",chardata"`
			ID     string `xml:"id,attr"`
			Alias  string `xml:"alias,attr"`
			Search string `xml:"search,attr"`
			Active string `xml:"active,attr"`
		} `xml:"alias"`
	} `xml:"entities"`
}
