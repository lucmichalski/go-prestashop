package fixtures

import (
	"encoding/xml"
)

// Customer was generated 2021-02-06 11:37:30 by evolutive on eg-cdn.gsi-network.com.
type Customer struct {
	XMLName xml.Name `xml:"entity_customer"`
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
		Text     string `xml:",chardata"`
		Customer []struct {
			Text                     string `xml:",chardata"`
			IDGender                 string `xml:"id_gender,attr"`
			ID                       string `xml:"id,attr"`
			IDDefaultGroup           string `xml:"id_default_group,attr"`
			Firstname                string `xml:"firstname,attr"`
			Lastname                 string `xml:"lastname,attr"`
			Email                    string `xml:"email,attr"`
			Passwd                   string `xml:"passwd,attr"`
			LastPasswdGen            string `xml:"last_passwd_gen,attr"`
			Birthday                 string `xml:"birthday,attr"`
			Newsletter               string `xml:"newsletter,attr"`
			IpRegistrationNewsletter string `xml:"ip_registration_newsletter,attr"`
			NewsletterDateAdd        string `xml:"newsletter_date_add,attr"`
			Optin                    string `xml:"optin,attr"`
			SecureKey                string `xml:"secure_key,attr"`
			Note                     string `xml:"note,attr"`
			Active                   string `xml:"active,attr"`
			IsGuest                  string `xml:"is_guest,attr"`
		} `xml:"customer"`
	} `xml:"entities"`
}
