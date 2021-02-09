package langs

import (
	"encoding/xml"
)

// Manufacturer was generated 2021-02-06 11:42:58 by evolutive on eg-cdn.gsi-network.com.
type ManufacturerLang struct {
	XMLName      xml.Name `xml:"entity_manufacturer"`
	Text         string   `xml:",chardata"`
	Manufacturer []struct {
		IDLang           string `xml:",chardata" struct2map:"key:id_lang"`
		ID               string `xml:"id,attr" struct2map:"key:id_manufacturer"`
		Description      string `xml:"description,attr" struct2map:"key:description"`
		ShortDescription string `xml:"short_description,attr" struct2map:"key:short_description"`
		MetaTitle        string `xml:"meta_title,attr" struct2map:"key:meta_title"`
		MetaKeywords     string `xml:"meta_keywords,attr" struct2map:"key:meta_keywords"`
	} `xml:"manufacturer"`
}
