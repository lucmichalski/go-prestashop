package fixtures

import (
	"encoding/xml"
)

// Cart was generated 2021-02-06 11:37:27 by evolutive on eg-cdn.gsi-network.com.
type Cart struct {
	XMLName xml.Name `xml:"cart"`
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
		Text string `xml:",chardata"`
		Cart []struct {
			Text                  string `xml:",chardata"`
			IDCustomer            string `xml:"id_customer,attr"`
			IDCarrier             string `xml:"id_carrier,attr"`
			IDAddressDelivery     string `xml:"id_address_delivery,attr"`
			IDGuest               string `xml:"id_guest,attr"`
			ID                    string `xml:"id,attr"`
			SecureKey             string `xml:"secure_key,attr"`
			DeliveryOption        string `xml:"delivery_option,attr"`
			IDCurrency            string `xml:"id_currency,attr"`
			Recyclable            string `xml:"recyclable,attr"`
			Gift                  string `xml:"gift,attr"`
			AllowSeparatedPackage string `xml:"allow_separated_package,attr"`
			IDShop                string `xml:"id_shop,attr"`
			IDShopGroup           string `xml:"id_shop_group,attr"`
		} `xml:"cart"`
	} `xml:"entities"`
}
