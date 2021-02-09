package langs

import (
	"encoding/xml"
)

type AttributeLang struct {
	XMLName   xml.Name `xml:"entity_attribute"`
	Text      string   `xml:",chardata"`
	Attribute []struct {
		IDLang string `xml:",chardata" struct2map:"key:id_lang"`
		ID     string `xml:"id,attr" struct2map:"key:id_attribute"`
		Name   string `xml:"name,attr" struct2map:"key:name"`
	} `xml:"attribute"`
}
