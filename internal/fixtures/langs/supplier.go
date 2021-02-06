package fixtures

import (
	"encoding/xml"
)

// Supplier was generated 2021-02-06 11:42:59 by evolutive on eg-cdn.gsi-network.com.
type Supplier struct {
	XMLName  xml.Name `xml:"supplier"`
	Text     string   `xml:",chardata"`
	Supplier []struct {
		Text string `xml:",chardata"`
		ID   string `xml:"id,attr"`
		Name string `xml:"name,attr"`
	} `xml:"supplier"`
}
