package fixtures

import (
	"encoding/xml"
)

// OrderHistory was generated 2021-02-06 11:37:30 by evolutive on eg-cdn.gsi-network.com.
type OrderHistory struct {
	XMLName xml.Name `xml:"entity_order_history"`
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
		Text         string `xml:",chardata"`
		OrderHistory []struct {
			Text         string `xml:",chardata"`
			IDOrder      string `xml:"id_order,attr"`
			IDOrderState string `xml:"id_order_state,attr"`
			IDEmployee   string `xml:"id_employee,attr"`
		} `xml:"order_history"`
	} `xml:"entities"`
}
