package langs

import (
	"encoding/xml"
)

// Category was generated 2021-02-06 11:42:58 by evolutive on eg-cdn.gsi-network.com.
type CategoryLang struct {
	XMLName  xml.Name `xml:"entity_category"`
	Text     string   `xml:",chardata"`
	Category []struct {
		IDLang      string `xml:",chardata" struct2map:"key:id_lang"`
		IDShop      string `xml:"id_shop,attr" struct2map:"key:id_shop"`
		ID          string `xml:"id,attr" struct2map:"key:id_category"`
		Name        string `xml:"name,attr" struct2map:"key:name"`
		Description string `xml:"description,attr" struct2map:"key:description"`
		LinkRewrite string `xml:"link_rewrite,attr" struct2map:"key:link_rewrite"`
	} `xml:"category"`
}
