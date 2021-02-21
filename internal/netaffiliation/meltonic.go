package netaffiliation

import (
	"fmt"

	// "github.com/fatih/structs"
	// "github.com/k0kubun/pp"
)

/*

// Catalog Path: catalogs/meltonic.csv
// Catalog Separator: ,
// Catalog excerpt: #1

// ===================================================================================== \\
id: 1,
titre: Barre énergétique Bio salée - Miel, Pistaches & Fleur de sel de Guérande - MELTONIC,
description: Une recette originale et sportive par Meltonic qui va vous faire voyager sur la Côte Atlantique. Cette barre bio sport mélange le bon goût du sel de Guérande et de la pistache avec les vertus sportives du miel. Digestibilité optimum et performance sportive garantie pendant votre effort. Découvrez cette barre innovante aux ingredients naturels, bio et croquants ! ,
lien: https://action.metaffiliation.com/trk.php?mclic=P49A87459831DS1UA4129a70124V4,
état: new,
prix: 2.40 EUR,
prix ​soldé: ,
période ​de ​validité ​du ​prix ​soldé: ,
disponibilité: in stock,
lien image: https://cdn1.meltonic.fr/I-Grande-1230-barre-energetique-bio-salee-miel-pistaches-fleur-de-sel-de-guerande.net.jpg,
lien ​image supplémentaire: ,
gtin: 3760115097087,
référence fabricant: 140925,
marque: Meltonic,
catégorie de produits google: Santé et beauté > Santé > Fitness et nutrition > Barres nutritionnelles,
type de produit: ,
livraison: ,
étiquette personnalisée 0: ,
étiquette personnalisée 1: 
\\ ===================================================================================== //

*/

type Meltonic struct {
	CategorieDeProduitsGoogle    string `json:"categorie_de_produits_google" struct2map:"key:categorie_de_produits_google"`
	Description                  string `json:"description" struct2map:"key:description"`
	Disponibilite                string `json:"disponibilite" struct2map:"key:disponibilite"`
	Etat                         string `json:"etat" struct2map:"key:etat"`
	EtiquettePersonnalisee0      string `json:"etiquette_personnalisee_0" struct2map:"key:etiquette_personnalisee_0"`
	EtiquettePersonnalisee1      string `json:"etiquette_personnalisee_1" struct2map:"key:etiquette_personnalisee_1"`
	Gtin                         string `json:"gtin" struct2map:"key:gtin"`
	ID                           string `json:"id" struct2map:"key:id"`
	Lien                         string `json:"lien" struct2map:"key:lien"`
	LienImage                    string `json:"lien_image" struct2map:"key:lien_image"`
	LienImageSupplementaire      string `json:"lien_​image_supplementaire" struct2map:"key:lien_​image_supplementaire"`
	Livraison                    string `json:"livraison" struct2map:"key:livraison"`
	Marque                       string `json:"marque" struct2map:"key:marque"`
	PeriodeDeValiditeDuPrixSolde string `json:"periode_​de_​validite_​du_​prix_​solde" struct2map:"key:periode_​de_​validite_​du_​prix_​solde"`
	Prix                         string `json:"prix" struct2map:"key:prix"`
	PrixSolde                    string `json:"prix_​solde" struct2map:"key:prix_​solde"`
	ReferenceFabricant           string `json:"reference_fabricant" struct2map:"key:reference_fabricant"`
	Titre                        string `json:"titre" struct2map:"key:titre"`
	TypeDeProduit                string `json:"type_de_produit" struct2map:"key:type_de_produit"`
}

/*
func (o *Meltonic) Convert() ([]string, error) {

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
