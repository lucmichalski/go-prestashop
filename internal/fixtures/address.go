package fixtures

import (
	"encoding/xml"
)

// Address was generated 2021-02-06 11:16:01 by evolutive on eg-cdn.gsi-network.com.
type Address struct {
	XMLName xml.Name `xml:"address"`
	Text    string   `xml:",chardata"`
	Fields  struct {
		Text  string `xml:",chardata"`
		ID    string `xml:"id,attr"`
		Class string `xml:"class,attr"`
		Field []struct {
			Text     string `xml:",chardata"`
			Name     string `xml:"name,attr"`
			Relation string `xml:"relation,attr"`
		} `xml:"field"`
	} `xml:"fields"`
	Entities struct {
		Text    string `xml:",chardata"`
		Address []struct {
			Text           string `xml:",chardata"`
			IDCustomer     string `xml:"id_customer,attr"`
			IDSupplier     string `xml:"id_supplier,attr"`
			IDCountry      string `xml:"id_country,attr"`
			IDState        string `xml:"id_state,attr"`
			IDManufacturer string `xml:"id_manufacturer,attr"`
			ID             string `xml:"id,attr"`
			Alias          string `xml:"alias,attr"`
			Company        string `xml:"company,attr"`
			Lastname       string `xml:"lastname,attr"`
			Firstname      string `xml:"firstname,attr"`
			Address1       string `xml:"address1,attr"`
			Address2       string `xml:"address2,attr"`
			Postcode       string `xml:"postcode,attr"`
			City           string `xml:"city,attr"`
			Other          string `xml:"other,attr"`
			PhoneMobile    string `xml:"phone_mobile,attr"`
			VatNumber      string `xml:"vat_number,attr"`
			Active         string `xml:"active,attr"`
			IDWarehouse    string `xml:"id_warehouse,attr"`
			Dni            string `xml:"dni,attr"`
		} `xml:"address"`
	} `xml:"entities"`
}
