package netaffiliation

import (
	"fmt"

	// "github.com/fatih/structs"
	// "github.com/k0kubun/pp"
)

/*

// Catalog Path: catalogs/original-cup.csv
// Catalog Separator: |
// Catalog excerpt: #1

// ===================================================================================== \\
ID_PRODUCT: 10,
NAME_PRODUCT: 20 Gobelets Rouges 53cl,
REFERENCE_PRODUCT: GR53,
SUPPLIER_REFERENCE: ,
MANUFACTURER: ,
CATEGORY: Bar et Fête > Gobelets,
DESCRIPTION: BEER PONG : Chaque équipe (1 ou 2 joueurs) joue chacun son tour et doit viser les gobelets adverses. Quand le tir est réussi, le gobelet doit être bu et retiré. L'équipe qui n'a plus de gobelets a perdu la partie. QUALITÉ PREMIUM : Verres réutilisables de la marque Original Cup. En plastique rigide et très résistant. Les gobelets sont lavables et passent au lave-vaisselle plus de 20 fois. INDISPENSABLES : Des gobelets uniques pour tes soirées, barbecues, pique-niques, anniversaires, EVJF/EVG et bien sûr quand tu joues au Beer Pong !,
DESCRIPTION_SHORT: L'incontournable Gobelet Rouge Américain XXL pour animer tes soirées et jouer au Beer Pong avec tes potes !,
PRICE_PRODUCT: 4.5,
WHOLESALE_PRICE: 0.98,
PRICE_HT: 3.75,
PRICE_REDUCTION: 4.5,
POURCENTAGE_REDUCTION: 0,
QUANTITY: 16808,
WEIGHT: 0.280000,
EAN: 3760191910003,
UPC: ,
ECOTAX: 0.00,
ACTIVE: 1,
AVAILABLE_PRODUCT: En stock,
URL_PRODUCT: https://qrj.originalcup.fr/?P4D9454598319S1U84125a101V4,
IMAGE_PRODUCT: https://www.originalcup.fr/5953-large/20-gobelets-rouges-53cl.jpg,
FDP: 0.00,
ID_MERE: 10,
DELAIS_LIVRAISON: 24/48 hrs,
IMAGE_PRODUCT_2: https://www.originalcup.fr/5954-large/20-gobelets-rouges-53cl.jpg,
IMAGE_PRODUCT_3: https://www.originalcup.fr/5955-large/20-gobelets-rouges-53cl.jpg,
REDUCTION_FROM: ,
REDUCTION_TO: ,
META_KEYWORDS: '20 gobelet rouge','20 gobelets rouges','20 red cup','20 red cups','20 gobelet americain','20 gobelets americains','20 gobelets américains','20 gobelet américain',
META_DESCRIPTION: Pack 20 gobelets rouges 53 cl à l'américaine ! Bénéficiez de prix imbattables et d'une livraison rapide !,
URL_REWRITE: https://qrj.originalcup.fr/?P4D9454598319S1U84125a101V4,
PRODUCT_TYPE: simple,
PRODUCT_VARIATION: ,
CURRENCY: EUR,
CONDITION: new,
SUPPLIER: ,
MINIMAL_QUANTITY: 1,
IS_VIRTUAL: 0,
AVAILABLE_FOR_ORDER: 1,
AVAILABLE_DATE: 0000-00-00,
SHOW_PRICE: 1,
VISIBILITY: both,
AVAILABLE_NOW: En stock,
AVAILABLE_LATER: Revient bientôt,
STOCK_AVAILABLES: ,
DESCRIPTION_HTML: <ul class='a-unordered-list a-vertical a-spacing-mini'> <li><span class='a-list-item'>BEER PONG : Chaque équipe (1 ou 2 joueurs) joue chacun son tour et doit viser les gobelets adverses. Quand le tir est réussi, le gobelet doit être bu et retiré. L'équipe qui n'a plus de gobelets a perdu la partie.</span></li> </ul> <ul class='a-unordered-list a-vertical a-spacing-mini'> <li><span class='a-list-item'>QUALITÉ PREMIUM : Verres réutilisables de la marque Original Cup. En plastique rigide et très résistant. Les gobelets sont lavables et passent au lave-vaisselle plus de 20 fois.</span></li> </ul> <ul class='a-unordered-list a-vertical a-spacing-mini'> <li><span class='a-list-item'>INDISPENSABLES : Des gobelets uniques pour tes soirées, barbecues, pique-niques, anniversaires, EVJF/EVG et bien sûr quand tu joues au Beer Pong !</span></li> </ul>,
AVAILABILITY: 1,
COULEUR: ,
LOT: ,
TAILLE: 
\\ ===================================================================================== //

*/

type OriginalCup struct {
	Active               string `json:"active" struct2map:"key:active"`
	Availability         string `json:"availability" struct2map:"key:availability"`
	AvailableDate        string `json:"available_date" struct2map:"key:available_date"`
	AvailableForOrder    string `json:"available_for_order" struct2map:"key:available_for_order"`
	AvailableLater       string `json:"available_later" struct2map:"key:available_later"`
	AvailableNow         string `json:"available_now" struct2map:"key:available_now"`
	AvailableProduct     string `json:"available_product" struct2map:"key:available_product"`
	Category             string `json:"category" struct2map:"key:category"`
	Condition            string `json:"condition" struct2map:"key:condition"`
	Couleur              string `json:"couleur" struct2map:"key:couleur"`
	Currency             string `json:"currency" struct2map:"key:currency"`
	DelaisLivraison      string `json:"delais_livraison" struct2map:"key:delais_livraison"`
	Description          string `json:"description" struct2map:"key:description"`
	DescriptionHTML      string `json:"description_html" struct2map:"key:description_html"`
	DescriptionShort     string `json:"description_short" struct2map:"key:description_short"`
	Ean                  string `json:"ean" struct2map:"key:ean"`
	Ecotax               string `json:"ecotax" struct2map:"key:ecotax"`
	Fdp                  string `json:"fdp" struct2map:"key:fdp"`
	IDMere               string `json:"id_mere" struct2map:"key:id_mere"`
	IDProduct            string `json:"id_product" struct2map:"key:id_product"`
	ImageProduct         string `json:"image_product" struct2map:"key:image_product"`
	ImageProduct2        string `json:"image_product_2" struct2map:"key:image_product_2"`
	ImageProduct3        string `json:"image_product_3" struct2map:"key:image_product_3"`
	IsVirtual            string `json:"is_virtual" struct2map:"key:is_virtual"`
	Lot                  string `json:"lot" struct2map:"key:lot"`
	Manufacturer         string `json:"manufacturer" struct2map:"key:manufacturer"`
	MetaDescription      string `json:"meta_description" struct2map:"key:meta_description"`
	MetaKeywords         string `json:"meta_keywords" struct2map:"key:meta_keywords"`
	MinimalQuantity      string `json:"minimal_quantity" struct2map:"key:minimal_quantity"`
	NameProduct          string `json:"name_product" struct2map:"key:name_product"`
	PourcentageReduction string `json:"pourcentage_reduction" struct2map:"key:pourcentage_reduction"`
	PriceHt              string `json:"price_ht" struct2map:"key:price_ht"`
	PriceProduct         string `json:"price_product" struct2map:"key:price_product"`
	PriceReduction       string `json:"price_reduction" struct2map:"key:price_reduction"`
	ProductType          string `json:"product_type" struct2map:"key:product_type"`
	ProductVariation     string `json:"product_variation" struct2map:"key:product_variation"`
	Quantity             string `json:"quantity" struct2map:"key:quantity"`
	ReductionFrom        string `json:"reduction_from" struct2map:"key:reduction_from"`
	ReductionTo          string `json:"reduction_to" struct2map:"key:reduction_to"`
	ReferenceProduct     string `json:"reference_product" struct2map:"key:reference_product"`
	ShowPrice            string `json:"show_price" struct2map:"key:show_price"`
	StockAvailables      string `json:"stock_availables" struct2map:"key:stock_availables"`
	Supplier             string `json:"supplier" struct2map:"key:supplier"`
	SupplierReference    string `json:"supplier_reference" struct2map:"key:supplier_reference"`
	Taille               string `json:"taille" struct2map:"key:taille"`
	Upc                  string `json:"upc" struct2map:"key:upc"`
	URLProduct           string `json:"url_product" struct2map:"key:url_product"`
	URLRewrite           string `json:"url_rewrite" struct2map:"key:url_rewrite"`
	Visibility           string `json:"visibility" struct2map:"key:visibility"`
	Weight               string `json:"weight" struct2map:"key:weight"`
	WholesalePrice       string `json:"wholesale_price" struct2map:"key:wholesale_price"`
}

/*
func (o *OriginalCup) Convert() ([]string, error) {

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
