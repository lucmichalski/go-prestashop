package fixtures

import (
	"encoding/xml"
)

// Product was generated 2021-02-06 11:43:01 by evolutive on eg-cdn.gsi-network.com.
type Product struct {
	XMLName xml.Name `xml:"product"`
	Text    string   `xml:",chardata"`
	Product []struct {
		Text             string `xml:",chardata"`
		IDShop           string `xml:"id_shop,attr"`
		ID               string `xml:"id,attr"`
		Name             string `xml:"name,attr"`
		Description      string `xml:"description,attr"`
		DescriptionShort string `xml:"description_short,attr"`
		LinkRewrite      string `xml:"link_rewrite,attr"`
		AvailableNow     string `xml:"available_now,attr"`
	} `xml:"product"`
}
