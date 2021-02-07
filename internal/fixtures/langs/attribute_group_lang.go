package langs

import (
	"encoding/xml"
)

type AttributeGroupLang struct {
	XMLName        xml.Name `xml:"entity_attribute_group"`
	Text           string   `xml:",chardata"`
	AttributeGroup []struct {
		Text       string `xml:",chardata" struct2map:"-"`
		ID         string `xml:"id,attr" struct2map:"key:id"`
		Name       string `xml:"name,attr" struct2map:"key:name"`
		PublicName string `xml:"public_name,attr" struct2map:"key:public_name"`
	} `xml:"attribute_group"`
}
