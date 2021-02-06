package fixtures

import (
	"encoding/xml"
)

// OrderDetail was generated 2021-02-06 11:17:08 by evolutive on eg-cdn.gsi-network.com.
type OrderDetail struct {
	XMLName xml.Name `xml:"order_detail"`
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
		OrderDetail []struct {
			Text                      string `xml:",chardata"`
			IDOrder                   string `xml:"id_order,attr"`
			ProductID                 string `xml:"product_id,attr"`
			ProductAttributeID        string `xml:"product_attribute_id,attr"`
			ProductName               string `xml:"product_name,attr"`
			ProductQuantity           string `xml:"product_quantity,attr"`
			ProductQuantityInStock    string `xml:"product_quantity_in_stock,attr"`
			ProductQuantityRefunded   string `xml:"product_quantity_refunded,attr"`
			ProductQuantityReturn     string `xml:"product_quantity_return,attr"`
			ProductQuantityReinjected string `xml:"product_quantity_reinjected,attr"`
			ProductPrice              string `xml:"product_price,attr"`
			ReductionPercent          string `xml:"reduction_percent,attr"`
			ProductQuantityDiscount   string `xml:"product_quantity_discount,attr"`
			ProductEan13              string `xml:"product_ean13,attr"`
			ProductUpc                string `xml:"product_upc,attr"`
			ProductReference          string `xml:"product_reference,attr"`
			ProductWeight             string `xml:"product_weight,attr"`
			IDOrderInvoice            string `xml:"id_order_invoice,attr"`
			IDWarehouse               string `xml:"id_warehouse,attr"`
			IDShop                    string `xml:"id_shop,attr"`
			ReductionAmount           string `xml:"reduction_amount,attr"`
			ReductionAmountTaxIncl    string `xml:"reduction_amount_tax_incl,attr"`
			ReductionAmountTaxExcl    string `xml:"reduction_amount_tax_excl,attr"`
			GroupReduction            string `xml:"group_reduction,attr"`
			ProductSupplierReference  string `xml:"product_supplier_reference,attr"`
			TaxComputationMethod      string `xml:"tax_computation_method,attr"`
			TaxName                   string `xml:"tax_name,attr"`
			TaxRate                   string `xml:"tax_rate,attr"`
			EcotaxTaxRate             string `xml:"ecotax_tax_rate,attr"`
			DiscountQuantityApplied   string `xml:"discount_quantity_applied,attr"`
			DownloadNb                string `xml:"download_nb,attr"`
			DownloadDeadline          string `xml:"download_deadline,attr"`
			TotalPriceTaxIncl         string `xml:"total_price_tax_incl,attr"`
			TotalPriceTaxExcl         string `xml:"total_price_tax_excl,attr"`
			UnitPriceTaxIncl          string `xml:"unit_price_tax_incl,attr"`
			UnitPriceTaxExcl          string `xml:"unit_price_tax_excl,attr"`
			TotalShippingPriceTaxIncl string `xml:"total_shipping_price_tax_incl,attr"`
			TotalShippingPriceTaxExcl string `xml:"total_shipping_price_tax_excl,attr"`
			PurchaseSupplierPrice     string `xml:"purchase_supplier_price,attr"`
			OriginalProductPrice      string `xml:"original_product_price,attr"`
			DownloadHash              string `xml:"download_hash,attr"`
		} `xml:"order_detail"`
	} `xml:"entities"`
}
