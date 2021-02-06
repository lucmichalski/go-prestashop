package fixtures

import (
	"encoding/xml"
)

// Profile was generated 2021-02-06 11:40:20 by evolutive on eg-cdn.gsi-network.com.
type Profile struct {
	XMLName xml.Name `xml:"profile"`
	Text    string   `xml:",chardata"`
	Fields  struct {
		Text  string `xml:",chardata"`
		ID    string `xml:"id,attr"`
		Class string `xml:"class,attr"`
		Sql   string `xml:"sql,attr"`
	} `xml:"fields"`
	Entities struct {
		Text    string `xml:",chardata"`
		Profile []struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
		} `xml:"profile"`
	} `xml:"entities"`
}
