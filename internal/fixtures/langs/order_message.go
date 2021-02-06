package fixtures

import (
	"encoding/xml"
)

// OrderMessage was generated 2021-02-06 11:42:59 by evolutive on eg-cdn.gsi-network.com.
type OrderMessage struct {
	XMLName      xml.Name `xml:"order_message"`
	Text         string   `xml:",chardata"`
	OrderMessage []struct {
		Text    string `xml:",chardata"`
		ID      string `xml:"id,attr"`
		Name    string `xml:"name,attr"`
		Message string `xml:"message,attr"`
	} `xml:"order_message"`
}
