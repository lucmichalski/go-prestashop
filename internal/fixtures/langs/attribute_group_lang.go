package langs

import (
	"encoding/xml"
)

type AttributeGroupLang struct {
	XMLName        xml.Name `xml:"entity_attribute_group"`
	Text           string   `xml:",chardata"`
	AttributeGroup []struct {
		IDLang     string `xml:",chardata" struct2map:"key:id_lang"`
		ID         string `xml:"id,attr" struct2map:"key:id_attribute_group"`
		Name       string `xml:"name,attr" struct2map:"key:name"`
		PublicName string `xml:"public_name,attr" struct2map:"key:public_name"`
	} `xml:"attribute_group"`
}
