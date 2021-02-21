package netaffiliation

import (
	"fmt"

	// "github.com/fatih/structs"
	// "github.com/k0kubun/pp"
)

/*

// Catalog Path: catalogs/vitalco.csv
// Catalog Separator: ,
// Catalog excerpt: #1

// ===================================================================================== \\
EAN: 3401597709814,
NOM: Activ'Foie,
ID: 11,
PRIX TTC: 28.9,
PRIX BARRES: 28.9,
CATEGORIE: Santé et beauté > Santé,
URL: https://action.metaffiliation.com/trk.php?mclic=P4B7EF4598321BS1UA4129870267V4,
IMAGE: https://ftp.vitalco.com/shopping2/active-foie-vitalco-new-nordic.png,
MARQUE: New Nordic,
DESCRIPTION: Favorise le bon fonctionnement du foie et contribue à une meilleure digestion des graisses.,
FRAIS DE PORT: 0,
STOCK: 1,
CONDITION: 1
\\ ===================================================================================== //

*/

type Vitalco struct {
	Categorie   string `json:"categorie" struct2map:"key:categorie"`
	Condition   string `json:"condition" struct2map:"key:condition"`
	Description string `json:"description" struct2map:"key:description"`
	Ean         string `json:"ean" struct2map:"key:ean"`
	FraisDePort string `json:"frais_de_port" struct2map:"key:frais_de_port"`
	ID          string `json:"id" struct2map:"key:id"`
	Image       string `json:"image" struct2map:"key:image"`
	Marque      string `json:"marque" struct2map:"key:marque"`
	Nom         string `json:"nom" struct2map:"key:nom"`
	PrixBarres  string `json:"prix_barres" struct2map:"key:prix_barres"`
	PrixTtc     string `json:"prix_ttc" struct2map:"key:prix_ttc"`
	Stock       string `json:"stock" struct2map:"key:stock"`
	URL         string `json:"url" struct2map:"key:url"`
}

/*
func (o *Vitalco) Convert() ([]string, error) {

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
