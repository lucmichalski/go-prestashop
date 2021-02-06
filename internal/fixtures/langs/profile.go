package fixtures

import (
	"encoding/xml"
)

// Profile was generated 2021-02-06 11:42:59 by evolutive on eg-cdn.gsi-network.com.
type Profile struct {
	XMLName xml.Name `xml:"profile"`
	Text    string   `xml:",chardata"`
	Profile []struct {
		Text string `xml:",chardata"`
		ID   string `xml:"id,attr"`
		Name string `xml:"name,attr"`
	} `xml:"profile"`
}
