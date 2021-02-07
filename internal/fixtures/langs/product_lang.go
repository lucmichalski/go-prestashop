package langs

import (
	"encoding/xml"
)

// Product was generated 2021-02-06 11:43:01 by evolutive on eg-cdn.gsi-network.com.
type ProductLang struct {
	XMLName xml.Name `xml:"entity_product"`
	Text    string   `xml:",chardata"`
	Product []struct {
		Text             string `xml:",chardata" struct2map:"key:name"`
		IDShop           string `xml:"id_shop,attr" struct2map:"key:id_shop"`
		ID               string `xml:"id,attr" struct2map:"key:id"`
		Name             string `xml:"name,attr" struct2map:"key:name"`
		Description      string `xml:"description,attr" struct2map:"key:description"`
		DescriptionShort string `xml:"description_short,attr" struct2map:"key:description_short"`
		LinkRewrite      string `xml:"link_rewrite,attr" struct2map:"key:link_rewrite"`
		AvailableNow     string `xml:"available_now,attr" struct2map:"key:available_now"`
	} `xml:"product"`
}
