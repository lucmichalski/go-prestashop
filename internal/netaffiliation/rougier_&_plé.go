package netaffiliation

import (
	"fmt"

	// "github.com/fatih/structs"
	// "github.com/k0kubun/pp"
)

/*

// Catalog Path: catalogs/rougier-and-ple.csv
// Catalog Separator: ;
// Catalog excerpt: #1

// ===================================================================================== \\
sku_unique: 102109,
marque: Vaessen Créative,
nom: Stylo Versamark à encre transparente,
description: Rougier & Plé, la grande maison des arts créatifs, vous propose cet articles parmi plus de 20 000 références de beaux-arts et de loisirs créatifs au meilleur prix. Livré en 6 jours par Colissimo, GLS domicile, Mondial Relay ou en magasin.,
url: https://afi.rougier-ple.fr/?P4826945983225S1UC4127d3026556V4,
image_grand_format: https://www.rougier-ple.fr/phproduct/20140430/P_102109_P_1_ZOOM.jpg,
montantlivraison: 3.99,
prix_ttc: 2.65,
ecotax: 0.00,
categorie: SCRAPBOOKING > TAMPONS ENCRES > TAMPON ENCREUR,
ean: 0712353210011,
stock: 0,
sku_produit: 9010172,
disponibilite: 0,
promotion: 2,
nouveaute: 0
\\ ===================================================================================== //

*/

type RougierAndPlé struct {
	Categorie        string `json:"categorie" struct2map:"key:categorie"`
	Description      string `json:"description" struct2map:"key:description"`
	Disponibilite    string `json:"disponibilite" struct2map:"key:disponibilite"`
	Ean              string `json:"ean" struct2map:"key:ean"`
	Ecotax           string `json:"ecotax" struct2map:"key:ecotax"`
	ImageGrandFormat string `json:"image_grand_format" struct2map:"key:image_grand_format"`
	Marque           string `json:"marque" struct2map:"key:marque"`
	Montantlivraison string `json:"montantlivraison" struct2map:"key:montantlivraison"`
	Nom              string `json:"nom" struct2map:"key:nom"`
	Nouveaute        string `json:"nouveaute" struct2map:"key:nouveaute"`
	PrixTtc          string `json:"prix_ttc" struct2map:"key:prix_ttc"`
	Promotion        string `json:"promotion" struct2map:"key:promotion"`
	SkuProduit       string `json:"sku_produit" struct2map:"key:sku_produit"`
	SkuUnique        string `json:"sku_unique" struct2map:"key:sku_unique"`
	Stock            string `json:"stock" struct2map:"key:stock"`
	URL              string `json:"url" struct2map:"key:url"`
}

/*
func (o *RougierAndPlé) Convert() ([]string, error) {

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
