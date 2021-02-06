package fixtures

import (
	"encoding/xml"
)

// Connections was generated 2021-02-06 11:40:35 by evolutive on eg-cdn.gsi-network.com.
type Connections struct {
	XMLName xml.Name `xml:"connections"`
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
		Text        string `xml:",chardata"`
		Connections []struct {
			Text        string `xml:",chardata"`
			IDGuest     string `xml:"id_guest,attr"`
			ID          string `xml:"id,attr"`
			IDPage      string `xml:"id_page,attr"`
			IpAddress   string `xml:"ip_address,attr"`
			HTTPReferer string `xml:"http_referer,attr"`
			IDShopGroup string `xml:"id_shop_group,attr"`
			IDShop      string `xml:"id_shop,attr"`
		} `xml:"connections"`
	} `xml:"entities"`
}
