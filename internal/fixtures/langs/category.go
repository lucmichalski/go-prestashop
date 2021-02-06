package fixtures

import (
	"encoding/xml"
)

// Category was generated 2021-02-06 11:42:58 by evolutive on eg-cdn.gsi-network.com.
type Category struct {
	XMLName  xml.Name `xml:"category"`
	Text     string   `xml:",chardata"`
	Category []struct {
		Text        string `xml:",chardata"`
		IDShop      string `xml:"id_shop,attr"`
		ID          string `xml:"id,attr"`
		Name        string `xml:"name,attr"`
		Description string `xml:"description,attr"`
		LinkRewrite string `xml:"link_rewrite,attr"`
	} `xml:"category"`
}
