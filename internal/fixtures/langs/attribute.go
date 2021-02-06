package fixtures

import (
	"encoding/xml"
)

// Attribute was generated 2021-02-06 11:42:59 by evolutive on eg-cdn.gsi-network.com.
type Attribute struct {
	XMLName   xml.Name `xml:"attribute"`
	Text      string   `xml:",chardata"`
	Attribute []struct {
		Text string `xml:",chardata"`
		ID   string `xml:"id,attr"`
		Name string `xml:"name,attr"`
	} `xml:"attribute"`
}
