package fixtures

import (
	"encoding/xml"
)

// AttributeGroup was generated 2021-02-06 11:37:27 by evolutive on eg-cdn.gsi-network.com.
type AttributeGroup struct {
	XMLName xml.Name `xml:"entity_attribute_group"`
	Text    string   `xml:",chardata"`
	Fields  struct {
		Text  string `xml:",chardata"`
		Class string `xml:"class,attr"`
		Field []struct {
			Text string `xml:",chardata"`
			Name string `xml:"name,attr"`
		} `xml:"field"`
	} `xml:"fields"`
	Entities struct {
		Text           string `xml:",chardata"`
		AttributeGroup []struct {
			Text         string `xml:",chardata"`
			ID           string `xml:"id,attr"`
			GroupType    string `xml:"group_type,attr"`
			IsColorGroup string `xml:"is_color_group,attr"`
		} `xml:"attribute_group"`
	} `xml:"entities"`
}
