package netaffiliation

import (
	"fmt"

	// "github.com/fatih/structs"
	// "github.com/k0kubun/pp"
)

/*

// Catalog Path: catalogs/zermara.csv
// Catalog Separator: ,
// Catalog excerpt: #1

// ===================================================================================== \\
Handle: helga,
Title: Helga,
URL: https://axl.zenmara.fr/?P5113A34598317S1U84131bb01V4,
Body (HTML): <p><strong>Un ensemble de yoga confort</strong></p><p><br>Trouver les bonnes tenues pour se rendre au yoga, ce n’est pas toujours très simple. Enfin, cela ne l’était pas toujours. Désormais, vous avez Yooga-Pants. Grâce à nous, vous n’aurez plus jamais à vous rendre à vo séances en ayant l’impression d’être mal habillée. Vous trouverez à votre disposition des vêtements tous plus qualitatifs les uns que les autres, comme cette magnifique tenue 2 en 1 assortie qui, en plus d’avoir une apparence superbe, est très confortable et ne vous dérangera pas même lors de vos exercices les plus exigeants.</p>
<p><strong>Caractéristiques</strong></p>
<ul>
<li>Matière : Coton</li>
<li>Taille S, M, L</li>
<li>Couleurs disponibles : 4</li>
<li>Livraison en 5 à 7 jours ouvrés en France</li>
<li>Extensible</li>
<li>Protection solaire</li>
<li>Anti-trace</li>
<li>Collant</li>
</ul>
<p>Composition et entretien<br>Yooga Pants ®<br>Yooga Flex, notre célèbre matière de haute compression, s&amp;#39;ajuste<br>parfaitement à votre silhouette. Elle vous offre un maintien optimal et est</p>
<p>idéale pour toutes vos activités sportives : yoga, boxe, gym et bien plus<br>encore.</p>,
Vendor: Yooga-Pants,
Type: ,
Tags: Expédié Depuis la France,
Published: TRUE,
Option1 Name: Couleur,
Option1 Value: YP0079-2,
Option2 Name: Taille,
Option2 Value: S,
Option3 Name: Expédié de,
Option3 Value: France,
Variant SKU: 38275416-yp0079-2-s-france,
Variant Grams: 0,
Variant Inventory Tracker: shopify,
Variant Inventory Policy: deny,
Variant Fulfillment Service: oberlo,
Variant Price: 59,
Variant Compare At Price: ,
Variant Requires Shipping: TRUE,
Variant Taxable: FALSE,
Variant Barcode: ,
Image Src: https://cdn.shopify.com/s/files/1/0017/6842/3471/products/Untitleddesign_3.png?v=1598031026,
Image Position: 1,
Image Alt Text: ,
Gift Card: FALSE,
SEO Title: ,
SEO Description: ,
Google Shopping / Google Product Category: ,
Google Shopping / Gender: ,
Google Shopping / Age Group: ,
Google Shopping / MPN: ,
Google Shopping / AdWords Grouping: ,
Google Shopping / AdWords Labels: ,
Google Shopping / Condition: ,
Google Shopping / Custom Product: ,
Google Shopping / Custom Label 0: ,
Google Shopping / Custom Label 1: ,
Google Shopping / Custom Label 2: ,
Google Shopping / Custom Label 3: ,
Google Shopping / Custom Label 4: ,
Variant Image: https://cdn.shopify.com/s/files/1/0017/6842/3471/products/product-image-1458864597.jpg?v=1598031026,
Variant Weight Unit: kg,
Variant Tax Code: ,
Cost per item: 
\\ ===================================================================================== //

*/

type Zermara struct {
	BodyHTML                            string `json:"body_html" struct2map:"key:body_html"`
	CostPerItem                         string `json:"cost_per_item" struct2map:"key:cost_per_item"`
	GiftCard                            string `json:"gift_card" struct2map:"key:gift_card"`
	GoogleShoppingAdWordsGrouping       string `json:"google_shopping_/_ad_words_grouping" struct2map:"key:google_shopping_/_ad_words_grouping"`
	GoogleShoppingAdWordsLabels         string `json:"google_shopping_/_ad_words_labels" struct2map:"key:google_shopping_/_ad_words_labels"`
	GoogleShoppingAgeGroup              string `json:"google_shopping_/_age_group" struct2map:"key:google_shopping_/_age_group"`
	GoogleShoppingCondition             string `json:"google_shopping_/_condition" struct2map:"key:google_shopping_/_condition"`
	GoogleShoppingCustomLabel0          string `json:"google_shopping_/_custom_label_0" struct2map:"key:google_shopping_/_custom_label_0"`
	GoogleShoppingCustomLabel1          string `json:"google_shopping_/_custom_label_1" struct2map:"key:google_shopping_/_custom_label_1"`
	GoogleShoppingCustomLabel2          string `json:"google_shopping_/_custom_label_2" struct2map:"key:google_shopping_/_custom_label_2"`
	GoogleShoppingCustomLabel3          string `json:"google_shopping_/_custom_label_3" struct2map:"key:google_shopping_/_custom_label_3"`
	GoogleShoppingCustomLabel4          string `json:"google_shopping_/_custom_label_4" struct2map:"key:google_shopping_/_custom_label_4"`
	GoogleShoppingCustomProduct         string `json:"google_shopping_/_custom_product" struct2map:"key:google_shopping_/_custom_product"`
	GoogleShoppingGender                string `json:"google_shopping_/_gender" struct2map:"key:google_shopping_/_gender"`
	GoogleShoppingGoogleProductCategory string `json:"google_shopping_/_google_product_category" struct2map:"key:google_shopping_/_google_product_category"`
	GoogleShoppingMpn                   string `json:"google_shopping_/_mpn" struct2map:"key:google_shopping_/_mpn"`
	Handle                              string `json:"handle" struct2map:"key:handle"`
	ImageAltText                        string `json:"image_alt_text" struct2map:"key:image_alt_text"`
	ImagePosition                       string `json:"image_position" struct2map:"key:image_position"`
	ImageSrc                            string `json:"image_src" struct2map:"key:image_src"`
	Option1Name                         string `json:"option_1_name" struct2map:"key:option_1_name"`
	Option1Value                        string `json:"option_1_value" struct2map:"key:option_1_value"`
	Option2Name                         string `json:"option_2_name" struct2map:"key:option_2_name"`
	Option2Value                        string `json:"option_2_value" struct2map:"key:option_2_value"`
	Option3Name                         string `json:"option_3_name" struct2map:"key:option_3_name"`
	Option3Value                        string `json:"option_3_value" struct2map:"key:option_3_value"`
	Published                           string `json:"published" struct2map:"key:published"`
	SeoDescription                      string `json:"seo_description" struct2map:"key:seo_description"`
	SeoTitle                            string `json:"seo_title" struct2map:"key:seo_title"`
	Tags                                string `json:"tags" struct2map:"key:tags"`
	Title                               string `json:"title" struct2map:"key:title"`
	Type                                string `json:"type" struct2map:"key:type"`
	URL                                 string `json:"url" struct2map:"key:url"`
	VariantBarcode                      string `json:"variant_barcode" struct2map:"key:variant_barcode"`
	VariantCompareAtPrice               string `json:"variant_compare_at_price" struct2map:"key:variant_compare_at_price"`
	VariantFulfillmentService           string `json:"variant_fulfillment_service" struct2map:"key:variant_fulfillment_service"`
	VariantGrams                        string `json:"variant_grams" struct2map:"key:variant_grams"`
	VariantImage                        string `json:"variant_image" struct2map:"key:variant_image"`
	VariantInventoryPolicy              string `json:"variant_inventory_policy" struct2map:"key:variant_inventory_policy"`
	VariantInventoryTracker             string `json:"variant_inventory_tracker" struct2map:"key:variant_inventory_tracker"`
	VariantPrice                        string `json:"variant_price" struct2map:"key:variant_price"`
	VariantRequiresShipping             string `json:"variant_requires_shipping" struct2map:"key:variant_requires_shipping"`
	VariantSku                          string `json:"variant_sku" struct2map:"key:variant_sku"`
	VariantTaxCode                      string `json:"variant_tax_code" struct2map:"key:variant_tax_code"`
	VariantTaxable                      string `json:"variant_taxable" struct2map:"key:variant_taxable"`
	VariantWeightUnit                   string `json:"variant_weight_unit" struct2map:"key:variant_weight_unit"`
	Vendor                              string `json:"vendor" struct2map:"key:vendor"`
}

/*
func (o *Zermara) Convert() ([]string, error) {

	mp := make(map[string]string, 0) 

	mp["product_name"] = o.ProductName
	mp["reference"] = o.Reference

	categoryId, err := findOrCreateCategoryByNam(db, dbTablePrefix, o.Category)
	if err != nil {
		return nil, err
	}
	mp["category_id"] = fmt.Sprintf("%d", categoryId)

	defaultCategory, err := findOrCreateCategoryByName(db, dbTablePrefix, o.DefaultCategory)
	if err != nil {
		return nil, err
	}
	mp["default_category"] = fmt.Sprintf("%d", defaultCategory)

	mp["short_description"] = o.ShortDescription
	mp["description"] = o.Description
	mp["price"] = fo.Price
	mp["quantity"] = o.Quantity
	mp["image_ref"] = o.ImageRef

	// go ast ?
	// m := structs.Name(o)
	// pp.Println(m)
	// for _, feature := range o.Features {}
	mp["feature_name"] = o.FeatureNames
	mp["feature_values"] = o.FeatureValues

	row =  []string{
		mp["product_name"], 
		mp["reference"], 
		mp["category_id"],
		mp["default_category"],
		mp["short_description"],
		mp["description"],
		mp["price"],
		mp["quantity"],
		mp["image_ref"],
		mp["feature_name"],
		mp["feature_values"],
	}

	return row, nil
}
*/
