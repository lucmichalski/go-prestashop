package fixtures

import (
	"encoding/xml"
)

// Manufacturer was generated 2021-02-06 11:42:58 by evolutive on eg-cdn.gsi-network.com.
type Manufacturer struct {
	XMLName      xml.Name `xml:"manufacturer"`
	Text         string   `xml:",chardata"`
	Manufacturer []struct {
		Text             string `xml:",chardata"`
		ID               string `xml:"id,attr"`
		Description      string `xml:"description,attr"`
		ShortDescription string `xml:"short_description,attr"`
		MetaTitle        string `xml:"meta_title,attr"`
		MetaKeywords     string `xml:"meta_keywords,attr"`
	} `xml:"manufacturer"`
}
