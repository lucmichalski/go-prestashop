package fixtures

import (
	"encoding/xml"
)

// AttributeGroup was generated 2021-02-06 11:42:58 by evolutive on eg-cdn.gsi-network.com.
type AttributeGroup struct {
	XMLName        xml.Name `xml:"attribute_group"`
	Text           string   `xml:",chardata"`
	AttributeGroup []struct {
		Text       string `xml:",chardata"`
		ID         string `xml:"id,attr"`
		Name       string `xml:"name,attr"`
		PublicName string `xml:"public_name,attr"`
	} `xml:"attribute_group"`
}
