package netaffiliation

import (
	"fmt"

	// "github.com/fatih/structs"
	// "github.com/k0kubun/pp"
)

/*

// Catalog Path: catalogs/lentilles-moins-cheres.csv
// Catalog Separator: |
// Catalog excerpt: #1

// ===================================================================================== \\
ref_universelle: 033362008728,
nom_usuel_produit: Air Optix Colors,
ref_interne: AIR-OPTIX-COLORS,
Prix_actuel: 21.00,
Prix_barre: ,
categorie_produit: Lentilles de contact>Mensuelles,
url_produit: https://rfg.lentillesmoinscheres.com/?P45DD345983217S1U94128eb080V4,
url_image: https://assets.lentillesmoinscheres.com/prod/site_fr/assets/images/product/ALCON/AIR-OPTIX-COLORS.jpg?201801081828FR,
ref_fabricant: Air Optix Colors,
nom_marque_fabricant: ALCON,
descr_produit: Boîte de 2 lentilles Air Optix Colors, mensuelles sphériques et souples, du laboratoire Alcon. Port journalier.  Photos non contractuelles : le rendu n'est pas le même si vous avez les yeux foncés. ,
frais_ports: 4.80,
eco_taxe: 0,
garantie: 1   
\\ ===================================================================================== //

*/

type LentillesMoinsChere struct {
	CategorieProduit   string `json:"categorie_produit" struct2map:"key:categorie_produit"`
	DescrProduit       string `json:"descr_produit" struct2map:"key:descr_produit"`
	EcoTaxe            string `json:"eco_taxe" struct2map:"key:eco_taxe"`
	FraisPorts         string `json:"frais_ports" struct2map:"key:frais_ports"`
	Garantie           string `json:"garantie" struct2map:"key:garantie"`
	NomMarqueFabricant string `json:"nom_marque_fabricant" struct2map:"key:nom_marque_fabricant"`
	NomUsuelProduit    string `json:"nom_usuel_produit" struct2map:"key:nom_usuel_produit"`
	PrixActuel         string `json:"prix_actuel" struct2map:"key:prix_actuel"`
	PrixBarre          string `json:"prix_barre" struct2map:"key:prix_barre"`
	RefFabricant       string `json:"ref_fabricant" struct2map:"key:ref_fabricant"`
	RefInterne         string `json:"ref_interne" struct2map:"key:ref_interne"`
	RefUniverselle     string `json:"ref_universelle" struct2map:"key:ref_universelle"`
	URLImage           string `json:"url_image" struct2map:"key:url_image"`
	URLProduit         string `json:"url_produit" struct2map:"key:url_produit"`
}

/*
func (o *LentillesMoinsChere) Convert() ([]string, error) {

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
