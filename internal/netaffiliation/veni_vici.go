package netaffiliation

import (
	"fmt"

	// "github.com/fatih/structs"
	// "github.com/k0kubun/pp"
)

/*

// Catalog Path: catalogs/veni-vici.csv
// Catalog Separator: ;
// Catalog excerpt: #1

// ===================================================================================== \\
EAN13: ,
Nom: Miaou For You,
Langue: fr,
Id: 27,
Id declinaison: 16574,
URL Image: ,
ID Attribute: 22,
Nom Attribute: S,
Prix TTC: 29,
Prix barre TTC: 29,
Category: Homme => Homme => T Shirt Homme => T-Shirts,
URL Page Produit: https://action.metaffiliation.com/trk.php?mclic=P4AC1345983233S1U841277b01V4,
Reference Fabriquant: ,
Brand Name: Venivici,
Description: T-Shirt mode coupe droite &amp; col rond large100% Polycoton (tr&egrave;s confortable &agrave; porter)Fabriqu&eacute; en France
\\ ===================================================================================== //

*/

type VeniVici struct {
	BrandName           string `json:"brand_name" struct2map:"key:brand_name"`
	Category            string `json:"category" struct2map:"key:category"`
	Description         string `json:"description" struct2map:"key:description"`
	Ean13               string `json:"ean_13" struct2map:"key:ean_13"`
	ID                  string `json:"id" struct2map:"key:id"`
	IDAttribute         string `json:"id_attribute" struct2map:"key:id_attribute"`
	IDDeclinaison       string `json:"id_declinaison" struct2map:"key:id_declinaison"`
	Langue              string `json:"langue" struct2map:"key:langue"`
	Nom                 string `json:"nom" struct2map:"key:nom"`
	NomAttribute        string `json:"nom_attribute" struct2map:"key:nom_attribute"`
	PrixBarreTtc        string `json:"prix_barre_ttc" struct2map:"key:prix_barre_ttc"`
	PrixTtc             string `json:"prix_ttc" struct2map:"key:prix_ttc"`
	ReferenceFabriquant string `json:"reference_fabriquant" struct2map:"key:reference_fabriquant"`
	URLImage            string `json:"url_image" struct2map:"key:url_image"`
	URLPageProduit      string `json:"url_page_produit" struct2map:"key:url_page_produit"`
}

/*
func (o *VeniVici) Convert() ([]string, error) {

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
