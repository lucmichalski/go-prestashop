package fixtures

import (
	"encoding/xml"
)

// Product was generated 2021-02-06 11:40:25 by evolutive on eg-cdn.gsi-network.com.
type Product struct {
	XMLName xml.Name `xml:"product"`
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
		Text    string `xml:",chardata"`
		Product []struct {
			Text                    string `xml:",chardata"`
			IDSupplier              string `xml:"id_supplier,attr"`
			IDManufacturer          string `xml:"id_manufacturer,attr"`
			IDCategoryDefault       string `xml:"id_category_default,attr"`
			ID                      string `xml:"id,attr"`
			Ean13                   string `xml:"ean13,attr"`
			Upc                     string `xml:"upc,attr"`
			Quantity                string `xml:"quantity,attr"`
			Price                   string `xml:"price,attr"`
			Reference               string `xml:"reference,attr"`
			Width                   string `xml:"width,attr"`
			Height                  string `xml:"height,attr"`
			Depth                   string `xml:"depth,attr"`
			OutOfStock              string `xml:"out_of_stock,attr"`
			OnSale                  string `xml:"on_sale,attr"`
			OnlineOnly              string `xml:"online_only,attr"`
			Ecotax                  string `xml:"ecotax,attr"`
			MinimalQuantity         string `xml:"minimal_quantity,attr"`
			WholesalePrice          string `xml:"wholesale_price,attr"`
			UnitPriceRatio          string `xml:"unit_price_ratio,attr"`
			AdditionalShippingCost  string `xml:"additional_shipping_cost,attr"`
			SupplierReference       string `xml:"supplier_reference,attr"`
			QuantityDiscount        string `xml:"quantity_discount,attr"`
			Customizable            string `xml:"customizable,attr"`
			UploadableFiles         string `xml:"uploadable_files,attr"`
			TextFields              string `xml:"text_fields,attr"`
			Active                  string `xml:"active,attr"`
			RedirectType            string `xml:"redirect_type,attr"`
			IDTypeRedirected        string `xml:"id_type_redirected,attr"`
			AvailableForOrder       string `xml:"available_for_order,attr"`
			AvailableDate           string `xml:"available_date,attr"`
			Condition               string `xml:"condition,attr"`
			ShowPrice               string `xml:"show_price,attr"`
			Indexed                 string `xml:"indexed,attr"`
			CacheIsPack             string `xml:"cache_is_pack,attr"`
			CacheHasAttachments     string `xml:"cache_has_attachments,attr"`
			IsVirtual               string `xml:"is_virtual,attr"`
			CacheDefaultAttribute   string `xml:"cache_default_attribute,attr"`
			AdvancedStockManagement string `xml:"advanced_stock_management,attr"`
		} `xml:"product"`
	} `xml:"entities"`
}
