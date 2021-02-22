package netaffiliation

import (
	"fmt"

	// "github.com/fatih/structs"
	// "github.com/k0kubun/pp"
)

/*

// Catalog Path: catalogs/moulin-des-senteurs.csv
// Catalog Separator: 	
// Catalog excerpt: #1

// ===================================================================================== \\
titre: Savon au Lait d'Anesse BIO 100g,
identifiant: 969223635001,
prix: 2.50 EUR,
clics: 0,
clics gratuits: 0,
état: neuf,
disponibilité: en stock,
Annonces Shopping - active: GB,DE,FR,IT,ES,NL,AT,BE,CH,IE,PL,SE,DK,CZ,TR,PT,RO,GR,FI,HU,SK,
Annonces Shopping - refusée ou non valide: ,
Fiches gratuites - active: GB,DE,FR,IT,ES,NL,AT,BE,CH,IE,PL,SE,DK,CZ,TR,PT,RO,GR,FI,HU,SK,
Fiches gratuites - refusée ou non valide: ,
Remarketing dynamique  - active: FR,
Remarketing dynamique  - refusée ou non valide: ,
Canal: Magasin en ligne,
pays: FR,
langue: fr,
description: Savon naturel au ait d'ânesse certifié bio (10%). Fabriqué en France. Livraison 24h-48h. Achetez dès maintenant ! ,
identifiant existe: oui,
lien: https://jod.moulin-des-senteurs.com/?P5115E5459831BS1U941322f083V4,
lien image: https://cdn.shopify.com/s/files/1/0057/9587/2825/products/Savon_Anesse_125gr_A_-min.png?v=1572428956,
lien image supplémentaire: ,
marque: Moulin Des Senteurs,
poids du colis: 0.1 kg
\\ ===================================================================================== //

*/

type MoulinDesSenteur struct {
	AnnoncesShoppingActive                 string `json:"annonces_shopping___active" struct2map:"key:annonces_shopping___active"`
	AnnoncesShoppingRefuseeOuNonValide     string `json:"annonces_shopping___refusee_ou_non_valide" struct2map:"key:annonces_shopping___refusee_ou_non_valide"`
	Canal                                  string `json:"canal" struct2map:"key:canal"`
	Clics                                  string `json:"clics" struct2map:"key:clics"`
	ClicsGratuits                          string `json:"clics_gratuits" struct2map:"key:clics_gratuits"`
	Description                            string `json:"description" struct2map:"key:description"`
	Disponibilite                          string `json:"disponibilite" struct2map:"key:disponibilite"`
	Etat                                   string `json:"etat" struct2map:"key:etat"`
	FichesGratuitesActive                  string `json:"fiches_gratuites___active" struct2map:"key:fiches_gratuites___active"`
	FichesGratuitesRefuseeOuNonValide      string `json:"fiches_gratuites___refusee_ou_non_valide" struct2map:"key:fiches_gratuites___refusee_ou_non_valide"`
	Identifiant                            string `json:"identifiant" struct2map:"key:identifiant"`
	IdentifiantExiste                      string `json:"identifiant_existe" struct2map:"key:identifiant_existe"`
	Langue                                 string `json:"langue" struct2map:"key:langue"`
	Lien                                   string `json:"lien" struct2map:"key:lien"`
	LienImage                              string `json:"lien_image" struct2map:"key:lien_image"`
	LienImageSupplementaire                string `json:"lien_image_supplementaire" struct2map:"key:lien_image_supplementaire"`
	Marque                                 string `json:"marque" struct2map:"key:marque"`
	Pays                                   string `json:"pays" struct2map:"key:pays"`
	PoidsDuColis                           string `json:"poids_du_colis" struct2map:"key:poids_du_colis"`
	Prix                                   string `json:"prix" struct2map:"key:prix"`
	RemarketingDynamiqueActive             string `json:"remarketing_dynamique____active" struct2map:"key:remarketing_dynamique____active"`
	RemarketingDynamiqueRefuseeOuNonValide string `json:"remarketing_dynamique____refusee_ou_non_valide" struct2map:"key:remarketing_dynamique____refusee_ou_non_valide"`
	Titre                                  string `json:"titre" struct2map:"key:titre"`
}

/*
func (o *MoulinDesSenteur) Convert() ([]string, error) {

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
