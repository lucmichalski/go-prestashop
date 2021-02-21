package netaffiliation

import (
	"fmt"

	// "github.com/fatih/structs"
	// "github.com/k0kubun/pp"
)

/*

// Catalog Path: catalogs/shop-mini.csv
// Catalog Separator: ;
// Catalog excerpt: #1

// ===================================================================================== \\
id: 409,
réference ean: ,
nom: T-shirt Femme monogramme MINI Colour Block Blanc/aquamarine,
référence: 80142445547,
prix TTC: 26.10,
prix barré: 29.00,
catégories: Accueil > Lifestyle > Pour elle > Vêtements,
url produit: https://cbl.shop-mini.fr/?P4C77B45983211S1UA4129730304V4,
url image: https://shop-mini.fr/img/p/1/3/6/8/1368.jpg,
réference HAN: ,
marque: ,
description: <p>Basique MINI à porter toute l'année. En jersey simple doux pour une coupe mi-ample confortable, avec encolure ronde, épaules tombantes et ourlet arrondi dans le dos. Un bloc de couleur contrastant souligne le monogramme MINI en impression caoutchouc noir.</p><p>100% coton.</p>,
disponible à l'achat: 1
\\ ===================================================================================== //

*/

type ShopMini struct {
	Categories        string `json:"categories" struct2map:"key:categories"`
	Description       string `json:"description" struct2map:"key:description"`
	DisponiblealAchat string `json:"disponible_a_l'achat" struct2map:"key:disponible_a_l'achat"`
	ID                string `json:"id" struct2map:"key:id"`
	Marque            string `json:"marque" struct2map:"key:marque"`
	Nom               string `json:"nom" struct2map:"key:nom"`
	PrixBarre         string `json:"prix_barre" struct2map:"key:prix_barre"`
	PrixTtc           string `json:"prix_ttc" struct2map:"key:prix_ttc"`
	Reference         string `json:"reference" struct2map:"key:reference"`
	ReferenceEan      string `json:"reference_ean" struct2map:"key:reference_ean"`
	ReferenceHan      string `json:"reference_han" struct2map:"key:reference_han"`
	URLImage          string `json:"url_image" struct2map:"key:url_image"`
	URLProduit        string `json:"url_produit" struct2map:"key:url_produit"`
}

/*
func (o *ShopMini) Convert() ([]string, error) {

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
