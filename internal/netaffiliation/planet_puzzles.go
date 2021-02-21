package netaffiliation

import (
	"fmt"

	// "github.com/fatih/structs"
	// "github.com/k0kubun/pp"
)

/*

// Catalog Path: catalogs/planet-puzzles.csv
// Catalog Separator: |
// Catalog excerpt: #1

// ===================================================================================== \\
catégorie: Puzzles - Princes et Princesses,
marque: Djeco,
Nom: Djeco Puzzle Observation - Princesse,
productID: 1,
vendor_code: ,
refinterne: 1,
nb_pieces: 54,
designation: Puzzle Observation - Princesse,
Promotion: 2,
description: Puzzle Observation - Princesse,
ean13: ,
garantie: 1,
occasion: 0,
disponibilite: O,
prix: 13.95,
fp: 4.95,
url: https://msr.planet-puzzles.com/?P49E3945983251S1UC412889082354V4,
img90: ,
img70: https://data.planet-puzzles.com/m22/p1/p1_120x120.jpg,
Mots-clefs: ,
langue: ,
monnaie: ,
refinterne: Djeco-07556,
prix_barre: ,
reffabriquant: 07556,
img400: https://data.planet-puzzles.com/m22/p1/p1_700x700.jpg,
livraison: 2 jours,
image2: ,
image3: ,
categoryid2: 1 - Puzzles adultes
\\ ===================================================================================== //

*/

type PlanetPuzzle struct {
	Categorie     string `json:"categorie" struct2map:"key:categorie"`
	Categoryid2   string `json:"categoryid_2" struct2map:"key:categoryid_2"`
	Description   string `json:"description" struct2map:"key:description"`
	Designation   string `json:"designation" struct2map:"key:designation"`
	Disponibilite string `json:"disponibilite" struct2map:"key:disponibilite"`
	Ean13         string `json:"ean_13" struct2map:"key:ean_13"`
	Fp            string `json:"fp" struct2map:"key:fp"`
	Garantie      string `json:"garantie" struct2map:"key:garantie"`
	Image2        string `json:"image_2" struct2map:"key:image_2"`
	Image3        string `json:"image_3" struct2map:"key:image_3"`
	Img400        string `json:"img_400" struct2map:"key:img_400"`
	Img70         string `json:"img_70" struct2map:"key:img_70"`
	Img90         string `json:"img_90" struct2map:"key:img_90"`
	Langue        string `json:"langue" struct2map:"key:langue"`
	Livraison     string `json:"livraison" struct2map:"key:livraison"`
	Marque        string `json:"marque" struct2map:"key:marque"`
	Monnaie       string `json:"monnaie" struct2map:"key:monnaie"`
	MotsClefs     string `json:"mots_clefs" struct2map:"key:mots_clefs"`
	NbPieces      string `json:"nb_pieces" struct2map:"key:nb_pieces"`
	Nom           string `json:"nom" struct2map:"key:nom"`
	Occasion      string `json:"occasion" struct2map:"key:occasion"`
	Prix          string `json:"prix" struct2map:"key:prix"`
	PrixBarre     string `json:"prix_barre" struct2map:"key:prix_barre"`
	ProductID     string `json:"product_id" struct2map:"key:product_id"`
	Promotion     string `json:"promotion" struct2map:"key:promotion"`
	Reffabriquant string `json:"reffabriquant" struct2map:"key:reffabriquant"`
	Refinterne    string `json:"refinterne" struct2map:"key:refinterne"`
	URL           string `json:"url" struct2map:"key:url"`
	VendorCode    string `json:"vendor_code" struct2map:"key:vendor_code"`
}

/*
func (o *PlanetPuzzle) Convert() ([]string, error) {

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
