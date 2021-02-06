package fixtures

import (
	"encoding/xml"
)

// Carrier was generated 2021-02-06 11:42:59 by evolutive on eg-cdn.gsi-network.com.
type Carrier struct {
	XMLName xml.Name `xml:"carrier"`
	Text    string   `xml:",chardata"`
	Carrier []struct {
		Text  string `xml:",chardata"`
		ID    string `xml:"id,attr"`
		Delay string `xml:"delay,attr"`
	} `xml:"carrier"`
}
