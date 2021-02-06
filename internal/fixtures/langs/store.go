package fixtures

import (
	"encoding/xml"
)

// Store was generated 2021-02-06 11:43:01 by evolutive on eg-cdn.gsi-network.com.
type Store struct {
	XMLName xml.Name `xml:"store"`
	Text    string   `xml:",chardata"`
	Store   []struct {
		Text     string `xml:",chardata"`
		ID       string `xml:"id,attr"`
		Name     string `xml:"name,attr"`
		Address1 string `xml:"address1,attr"`
		Address2 string `xml:"address2,attr"`
		Hours    string `xml:"hours,attr"`
	} `xml:"store"`
}
