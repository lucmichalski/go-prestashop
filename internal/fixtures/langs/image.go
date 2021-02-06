package fixtures

import (
	"encoding/xml"
)

// Image was generated 2021-02-06 11:42:59 by evolutive on eg-cdn.gsi-network.com.
type Image struct {
	XMLName xml.Name `xml:"image"`
	Text    string   `xml:",chardata"`
	Image   []struct {
		Text   string `xml:",chardata"`
		ID     string `xml:"id,attr"`
		Legend string `xml:"legend,attr"`
	} `xml:"image"`
}
