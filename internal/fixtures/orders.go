package fixtures

import (
	"encoding/xml"
)

// Orders was generated 2021-02-06 11:40:15 by evolutive on eg-cdn.gsi-network.com.
type Orders struct {
	XMLName xml.Name `xml:"orders"`
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
		Text   string `xml:",chardata"`
		Orders []struct {
			Text                  string `xml:",chardata"`
			IDCart                string `xml:"id_cart,attr"`
			IDCarrier             string `xml:"id_carrier,attr"`
			IDCustomer            string `xml:"id_customer,attr"`
			IDAddressDelivery     string `xml:"id_address_delivery,attr"`
			IDAddressInvoice      string `xml:"id_address_invoice,attr"`
			ID                    string `xml:"id,attr"`
			IDOrder               string `xml:"id_order,attr"`
			Name                  string `xml:"name,attr"`
			CurrentState          string `xml:"current_state,attr"`
			SecureKey             string `xml:"secure_key,attr"`
			Gift                  string `xml:"gift,attr"`
			TotalProducts         string `xml:"total_products,attr"`
			TotalShipping         string `xml:"total_shipping,attr"`
			Payment               string `xml:"payment,attr"`
			Module                string `xml:"module,attr"`
			IDShopGroup           string `xml:"id_shop_group,attr"`
			IDShop                string `xml:"id_shop,attr"`
			IDCurrency            string `xml:"id_currency,attr"`
			ConversionRate        string `xml:"conversion_rate,attr"`
			Recyclable            string `xml:"recyclable,attr"`
			ShippingNumber        string `xml:"shipping_number,attr"`
			TotalDiscounts        string `xml:"total_discounts,attr"`
			TotalDiscountsTaxIncl string `xml:"total_discounts_tax_incl,attr"`
			TotalPaidReal         string `xml:"total_paid_real,attr"`
			TotalProductsWt       string `xml:"total_products_wt,attr"`
			TotalShippingTaxIncl  string `xml:"total_shipping_tax_incl,attr"`
			TotalShippingTaxExcl  string `xml:"total_shipping_tax_excl,attr"`
			TotalPaid             string `xml:"total_paid,attr"`
			TotalPaidTaxExcl      string `xml:"total_paid_tax_excl,attr"`
			TotalPaidTaxIncl      string `xml:"total_paid_tax_incl,attr"`
			CarrierTaxRate        string `xml:"carrier_tax_rate,attr"`
			TotalWrapping         string `xml:"total_wrapping,attr"`
			TotalWrappingTaxIncl  string `xml:"total_wrapping_tax_incl,attr"`
			TotalWrappingTaxExcl  string `xml:"total_wrapping_tax_excl,attr"`
			InvoiceNumber         string `xml:"invoice_number,attr"`
			DeliveryNumber        string `xml:"delivery_number,attr"`
			InvoiceDate           string `xml:"invoice_date,attr"`
			DeliveryDate          string `xml:"delivery_date,attr"`
			Valid                 string `xml:"valid,attr"`
			GiftMessage           string `xml:"gift_message,attr"`
		} `xml:"orders"`
	} `xml:"entities"`
}
