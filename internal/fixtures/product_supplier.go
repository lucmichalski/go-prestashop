package fixtures

import (
	"encoding/xml"
)

// ProductSupplier was generated 2021-02-06 11:17:19 by evolutive on eg-cdn.gsi-network.com.
type ProductSupplier struct {
	XMLName xml.Name `xml:"product_supplier"`
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
		Text            string `xml:",chardata"`
		ProductSupplier []struct {
			Text                     string `xml:",chardata"`
			IDProduct                string `xml:"id_product,attr"`
			IDProductAttribute       string `xml:"id_product_attribute,attr"`
			IDSupplier               string `xml:"id_supplier,attr"`
			ProductSupplierReference string `xml:"product_supplier_reference,attr"`
			ProductSupplierPriceTe   string `xml:"product_supplier_price_te,attr"`
		} `xml:"product_supplier"`
	} `xml:"entities"`
}
