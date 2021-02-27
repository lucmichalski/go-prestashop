package netaffiliation

import (
	"fmt"

	// "github.com/fatih/structs"
	// "github.com/k0kubun/pp"
)

/*

// Catalog Path: catalogs/crea-plaisir.csv
// Catalog Separator: ;
// Catalog excerpt: #1

// ===================================================================================== \\
Sku_unique: 9993E,
EAN: 5060517149932,
HAN: 9993E,
Nom: Bobine de fil de jute Bakers Twine - 2mm x25m,
Prix_ttc: 2.62,
Marque: Craft Perfect,
Description: Bobine de fil de jute Bakers Twine - 2mm x25m de la marque Craft Perfect disponible sur Créaplaisir. Tout pour le scrapbooking à prix canons !,
Url: https://rqj.crea-plaisir.com/?P5116DD4598317S1U841327101V4,
Image_grand_format: https://cdn.shopify.com/s/files/1/0275/8308/6730/products/9993E.jpg?v=1611745257,
Categorie: Fil,
Montant_port: 4,9,
Ecotax: 0,
Garantie: 0,
Disponibilité: 1,
Promotion: 3,
Nouveaute: 1
\\ ===================================================================================== //

*/

type CreaPlaisir struct {
	Categorie        string `json:"categorie" struct2map:"key:categorie"`
	Description      string `json:"description" struct2map:"key:description"`
	Disponibilite    string `json:"disponibilite" struct2map:"key:disponibilite"`
	Ean              string `json:"ean" struct2map:"key:ean"`
	Ecotax           string `json:"ecotax" struct2map:"key:ecotax"`
	Garantie         string `json:"garantie" struct2map:"key:garantie"`
	Han              string `json:"han" struct2map:"key:han"`
	ImageGrandFormat string `json:"image_grand_format" struct2map:"key:image_grand_format"`
	Marque           string `json:"marque" struct2map:"key:marque"`
	MontantPort      string `json:"montant_port" struct2map:"key:montant_port"`
	Nom              string `json:"nom" struct2map:"key:nom"`
	Nouveaute        string `json:"nouveaute" struct2map:"key:nouveaute"`
	PrixTtc          string `json:"prix_ttc" struct2map:"key:prix_ttc"`
	Promotion        string `json:"promotion" struct2map:"key:promotion"`
	SkuUnique        string `json:"sku_unique" struct2map:"key:sku_unique"`
	URL              string `json:"url" struct2map:"key:url"`
}

/*
func (o *CreaPlaisir) Convert() ([]string, error) {

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
