package langs

import (
	"encoding/xml"
)

// Supplier was generated 2021-02-06 11:42:59 by evolutive on eg-cdn.gsi-network.com.
type SupplierLang struct {
	XMLName  xml.Name `xml:"entity_supplier"`
	Text     string   `xml:",chardata"`
	Supplier []struct {
		IDLang string `xml:",chardata" struct2map:"key:id_lang"`
		ID     string `xml:"id,attr" struct2map:"key:id_supplier"`
		Name   string `xml:"name,attr" struct2map:"key:description"`
	} `xml:"supplier"`
}
