package fixtures

import (
	"encoding/xml"
)

// ProductAttributeCombination was generated 2021-02-06 11:37:24 by evolutive on eg-cdn.gsi-network.com.
type ProductAttributeCombination struct {
	XMLName xml.Name `xml:"entity_product_attribute_combination"`
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
		Text                        string `xml:",chardata"`
		ProductAttributeCombination []struct {
			Text               string `xml:",chardata"`
			IDAttribute        string `xml:"id_attribute,attr"`
			IDProductAttribute string `xml:"id_product_attribute,attr"`
		} `xml:"product_attribute_combination"`
	} `xml:"entities"`
}
