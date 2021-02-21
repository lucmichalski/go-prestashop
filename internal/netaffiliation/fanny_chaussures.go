package netaffiliation

import (
	"fmt"

	// "github.com/fatih/structs"
	// "github.com/k0kubun/pp"
)

/*

// Catalog Path: catalogs/fanny-chaussures.csv
// Catalog Separator: |
// Catalog excerpt: #1

// ===================================================================================== \\
product_id: 20438,
sku: 000488-32-Noir,
qty: 1,
type_stock: 1,
name: BIRKENSTOCK Arizona Enfant-32-Noir,
product_url: https://action.metaffiliation.com/trk.php?mclic=P44EBB4598321BS1UD4127b10611388V4,
category_breadcrumb: Chaussures Enfant > Chaussures Fille > Sandales & Nu-pieds,
price_ttc: 50.00,
price_before_discount: ,
type_promo: 2,
image_url_1: http://media.fanny-chaussures.com/media/catalog/product/cache/1/image/1200x/9df78eab33525d08d6e5fb8d27136e95/b/i/birkenstock-arizona-enfant-noir-1_2.jpg,
description: Paire de sandales pour enfants Birkenstock Noir. Ces sandales BIRKENSTOCK Arizona Enfant Noir se présentent de la composition suivante : Tige en Synthétique , Doublure en Synthétique, Semelle intérieure en Cuir, Semelle extérieure en Caoutchouc. Conseil Chaussant : Choisissez 1 pointure au-dessous de votre pointure habituelle.. Faites vous plaisir avec ces sandales Noir au style propre à la marque Birkenstock. ,
matiere: Synthétique,
marque: Birkenstock,
gencode: 4052001676728,
taille: 32,
type_chaussure: Sandales & Nu-pieds,
genre: Enfants,
weight: 1,
idmodele: ,
couleur: Noir,
frais_port: 0
\\ ===================================================================================== //

*/

type FannyChaussure struct {
	CategoryBreadcrumb  string `json:"category_breadcrumb" struct2map:"key:category_breadcrumb"`
	Couleur             string `json:"couleur" struct2map:"key:couleur"`
	Description         string `json:"description" struct2map:"key:description"`
	FraisPort           string `json:"frais_port" struct2map:"key:frais_port"`
	Gencode             string `json:"gencode" struct2map:"key:gencode"`
	Genre               string `json:"genre" struct2map:"key:genre"`
	Idmodele            string `json:"idmodele" struct2map:"key:idmodele"`
	ImageURL1           string `json:"image_url_1" struct2map:"key:image_url_1"`
	Marque              string `json:"marque" struct2map:"key:marque"`
	Matiere             string `json:"matiere" struct2map:"key:matiere"`
	Name                string `json:"name" struct2map:"key:name"`
	PriceBeforeDiscount string `json:"price_before_discount" struct2map:"key:price_before_discount"`
	PriceTtc            string `json:"price_ttc" struct2map:"key:price_ttc"`
	ProductID           string `json:"product_id" struct2map:"key:product_id"`
	ProductURL          string `json:"product_url" struct2map:"key:product_url"`
	Qty                 string `json:"qty" struct2map:"key:qty"`
	Sku                 string `json:"sku" struct2map:"key:sku"`
	Taille              string `json:"taille" struct2map:"key:taille"`
	TypeChaussure       string `json:"type_chaussure" struct2map:"key:type_chaussure"`
	TypePromo           string `json:"type_promo" struct2map:"key:type_promo"`
	TypeStock           string `json:"type_stock" struct2map:"key:type_stock"`
	Weight              string `json:"weight" struct2map:"key:weight"`
}

/*
func (o *FannyChaussure) Convert() ([]string, error) {

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
