package fixtures

import (
	"encoding/xml"
)

// OrderMessage was generated 2021-02-06 11:17:19 by evolutive on eg-cdn.gsi-network.com.
type OrderMessage struct {
	XMLName xml.Name `xml:"order_message"`
	Text    string   `xml:",chardata"`
	Fields  struct {
		Text  string `xml:",chardata"`
		ID    string `xml:"id,attr"`
		Class string `xml:"class,attr"`
	} `xml:"fields"`
	Entities struct {
		Text         string `xml:",chardata"`
		OrderMessage []struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
		} `xml:"order_message"`
	} `xml:"entities"`
}
