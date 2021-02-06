package fixtures

import (
	"encoding/xml"
)

// ProductAttributeImage was generated 2021-02-06 11:37:00 by evolutive on eg-cdn.gsi-network.com.
type ProductAttributeImage struct {
	XMLName xml.Name `xml:"product_attribute_image"`
	Text    string   `xml:",chardata"`
	Fields  struct {
		Text    string `xml:",chardata"`
		Primary string `xml:"primary,attr"`
		Field   []struct {
			Text     string `xml:",chardata"`
			Name     string `xml:"name,attr"`
			Relation string `xml:"relation,attr"`
		} `xml:"field"`
	} `xml:"fields"`
	Entities struct {
		Text                  string `xml:",chardata"`
		ProductAttributeImage []struct {
			Text               string `xml:",chardata"`
			IDProductAttribute string `xml:"id_product_attribute,attr"`
			IDImage            string `xml:"id_image,attr"`
		} `xml:"product_attribute_image"`
	} `xml:"entities"`
}
