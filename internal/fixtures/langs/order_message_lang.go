package langs

import (
	"encoding/xml"
)

// OrderMessage was generated 2021-02-06 11:42:59 by evolutive on eg-cdn.gsi-network.com.
type OrderMessageLang struct {
	XMLName      xml.Name `xml:"entity_order_message"`
	Text         string   `xml:",chardata"`
	OrderMessage []struct {
		IDLang  string `xml:",chardata" struct2map:"key:id_lang"`
		ID      string `xml:"id,attr" struct2map:"key:id_order_message"`
		Name    string `xml:"name,attr" struct2map:"key:name"`
		Message string `xml:"message,attr" struct2map:"key:message"`
	} `xml:"order_message"`
}
