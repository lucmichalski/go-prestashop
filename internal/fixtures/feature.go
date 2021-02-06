package fixtures

import (
	"encoding/xml"
)

// Feature was generated 2021-02-06 11:40:25 by evolutive on eg-cdn.gsi-network.com.
type Feature struct {
	XMLName xml.Name `xml:"feature"`
	Text    string   `xml:",chardata"`
	Fields  struct {
		Text  string `xml:",chardata"`
		Class string `xml:"class,attr"`
		Field struct {
			Text string `xml:",chardata"`
			Name string `xml:"name,attr"`
		} `xml:"field"`
	} `xml:"fields"`
	Entities struct {
		Text    string `xml:",chardata"`
		Feature []struct {
			Text     string `xml:",chardata"`
			ID       string `xml:"id,attr"`
			Position string `xml:"position,attr"`
		} `xml:"feature"`
	} `xml:"entities"`
}
