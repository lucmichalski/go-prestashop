package langs

import (
	"encoding/xml"
)

type CarrierLang struct {
	XMLName xml.Name `xml:"entity_carrier"`
	Text    string   `xml:",chardata"`
	Carrier []struct {
		IDLang string `xml:",chardata" struct2map:"key:id_lang"`
		ID     string `xml:"id,attr" struct2map:"key:id_carrier"`
		Delay  string `xml:"delay,attr" struct2map:"key:delay"`
	} `xml:"carrier"`
}
