package fixtures

import (
	"encoding/xml"
)

// Feature was generated 2021-02-06 11:43:01 by evolutive on eg-cdn.gsi-network.com.
type Feature struct {
	XMLName xml.Name `xml:"feature"`
	Text    string   `xml:",chardata"`
	Feature []struct {
		Text string `xml:",chardata"`
		ID   string `xml:"id,attr"`
		Name string `xml:"name,attr"`
	} `xml:"feature"`
}
