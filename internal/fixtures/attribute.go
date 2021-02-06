package fixtures

import (
	"encoding/xml"
)

// Attribute was generated 2021-02-06 11:40:15 by evolutive on eg-cdn.gsi-network.com.
type Attribute struct {
	XMLName xml.Name `xml:"attribute"`
	Text    string   `xml:",chardata"`
	Fields  struct {
		Text  string `xml:",chardata"`
		Class string `xml:"class,attr"`
		Field []struct {
			Text     string `xml:",chardata"`
			Name     string `xml:"name,attr"`
			Relation string `xml:"relation,attr"`
		} `xml:"field"`
	} `xml:"fields"`
	Entities struct {
		Text      string `xml:",chardata"`
		Attribute []struct {
			Text             string `xml:",chardata"`
			IDAttributeGroup string `xml:"id_attribute_group,attr"`
			ID               string `xml:"id,attr"`
			Color            string `xml:"color,attr"`
		} `xml:"attribute"`
	} `xml:"entities"`
}
