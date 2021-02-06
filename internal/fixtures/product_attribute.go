package fixtures

import (
	"encoding/xml"
)

// ProductAttribute was generated 2021-02-06 11:19:17 by evolutive on eg-cdn.gsi-network.com.
type ProductAttribute struct {
	XMLName xml.Name `xml:"product_attribute"`
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
		Text             string `xml:",chardata"`
		ProductAttribute []struct {
			Text              string `xml:",chardata"`
			IDProduct         string `xml:"id_product,attr"`
			ID                string `xml:"id,attr"`
			Ean13             string `xml:"ean13,attr"`
			Upc               string `xml:"upc,attr"`
			Quantity          string `xml:"quantity,attr"`
			DefaultOn         string `xml:"default_on,attr"`
			Reference         string `xml:"reference,attr"`
			SupplierReference string `xml:"supplier_reference,attr"`
			WholesalePrice    string `xml:"wholesale_price,attr"`
			Price             string `xml:"price,attr"`
			Ecotax            string `xml:"ecotax,attr"`
			Weight            string `xml:"weight,attr"`
			UnitPriceImpact   string `xml:"unit_price_impact,attr"`
			MinimalQuantity   string `xml:"minimal_quantity,attr"`
			AvailableDate     string `xml:"available_date,attr"`
		} `xml:"product_attribute"`
	} `xml:"entities"`
}
