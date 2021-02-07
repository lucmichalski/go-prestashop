package langs

import (
	"encoding/xml"
)

type AttributeLang struct {
	XMLName   xml.Name `xml:"entity_attribute"`
	Text      string   `xml:",chardata"`
	Attribute []struct {
		Text string `xml:",chardata" struct2map:"key:text"`
		ID   string `xml:"id,attr" struct2map:"key:id"`
		Name string `xml:"name,attr" struct2map:"key:name"`
	} `xml:"attribute"`
}
