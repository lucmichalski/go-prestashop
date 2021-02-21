package netaffiliation

import (
	"fmt"

	// "github.com/fatih/structs"
	// "github.com/k0kubun/pp"
)

/*

// Catalog Path: catalogs/condomz.csv
// Catalog Separator: ;
// Catalog excerpt: #1

// ===================================================================================== \\
categorie: Pour le prix§Pas cher 15 et 20 centimes,
identifiant: ID812,
titre: préservatif,
prix: 14.99,
url_produit: https://action.metaffiliation.com/trk.php?mclic=P4150E45983213S1UB41278d02131V4,
url_image: https://www.condomz.fr/images/products/condozone-100.jpg,
occasion: 0,
description: Condozone propose SON préservatif à 15 centimes depuis 2003, disponible uniquement sur Internet pour tous ceux qui veulent se protéger sans se ruiner. Les préservatifs Condozone sont testés et approuvés par notre équipe, leur qualité de fabrication est irréprochable (norme CE), et la finesse de leur latex en fait un super choix pour les afficionados du plaisir.,
reference: ,
livraison: 3.90,
disponibilite: En stock,
marque: ,
ean: 3700655300018,
garantie: 1,
prix_barre: 14.99
\\ ===================================================================================== //

*/

type Condomz struct {
	Categorie     string `json:"categorie" struct2map:"key:categorie"`
	Description   string `json:"description" struct2map:"key:description"`
	Disponibilite string `json:"disponibilite" struct2map:"key:disponibilite"`
	Ean           string `json:"ean" struct2map:"key:ean"`
	Garantie      string `json:"garantie" struct2map:"key:garantie"`
	Identifiant   string `json:"identifiant" struct2map:"key:identifiant"`
	Livraison     string `json:"livraison" struct2map:"key:livraison"`
	Marque        string `json:"marque" struct2map:"key:marque"`
	Occasion      string `json:"occasion" struct2map:"key:occasion"`
	Prix          string `json:"prix" struct2map:"key:prix"`
	PrixBarre     string `json:"prix_barre" struct2map:"key:prix_barre"`
	Reference     string `json:"reference" struct2map:"key:reference"`
	Titre         string `json:"titre" struct2map:"key:titre"`
	URLImage      string `json:"url_image" struct2map:"key:url_image"`
	URLProduit    string `json:"url_produit" struct2map:"key:url_produit"`
}

/*
func (o *Condomz) Convert() ([]string, error) {

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
