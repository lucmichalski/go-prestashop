package netaffiliation

import (
	"fmt"

	// "github.com/fatih/structs"
	// "github.com/k0kubun/pp"
)

/*

// Catalog Path: catalogs/thalassoissambres.csv
// Catalog Separator: ;
// Catalog excerpt: #1

// ===================================================================================== \\
Id: 181,
Nom: <p>Thalgo Soin Haute Correction Cellulite</p>,
Prix Actuel: 36,
Catégorie du produit: Santé et beauté > Hygiène personnelle > Cosmétiques > Soin de la peau > Crèmes et lotions,
Url du Produit: https://action.metaffiliation.com/trk.php?mclic=P43251459831DS1U84129b101V4,
N modèle du produit: ,
Marque du produit: Thalgo,
Code EAN: 3,53E+12,
Référence Constructeur: VT13015,
Frais de port: 0,
 : <p style="text-align: justify;">Ce soin corrige fondamentalement l&#39;alt&eacute;ration structurelle de la cellulite install&eacute;e. La technologie Oxy Active favorise la r&eacute;duction de l&#39;hypoxie tissulaire pour lutter contre l&#39;agglom&eacute;ration cellulitique naissante. Liposlim, polysaccharide biotechnologique, permet de r&eacute;duire significativement le volume et le nombre d&#39;adipocytes, tout en freinant leur future cr&eacute;ation.</p>,
Garantie: 0,
Indicateur de Stock: 1,
Disponibilité du produit: 48h,
Indicateur de performance: 15,
Promotion: 0,
Nouveauté: 0,
URL Image Petite: https://www.thalassoissambres.com/stream/index.html?image=%2Fuploaded%2Fphotos%2Fphotos_fichier_9.jpg&width=325,
URL Image moyenne: https://www.thalassoissambres.com/stream/index.html?image=%2Fuploaded%2Fphotos%2Fphotos_fichier_9.jpg&width=325,
URL Grande: https://www.thalassoissambres.com/stream/index.html?image=%2Fuploaded%2Fphotos%2Fphotos_fichier_9.jpg&width=325,
Ecotaxe: 0
\\ ===================================================================================== //

*/

type Thalassoissambre struct {
	string                  `json:"" struct2map:"key:"`
	CategorieDuProduit      string `json:"categorie_du_produit" struct2map:"key:categorie_du_produit"`
	CodeEan                 string `json:"code_ean" struct2map:"key:code_ean"`
	DisponibiliteDuProduit  string `json:"disponibilite_du_produit" struct2map:"key:disponibilite_du_produit"`
	Ecotaxe                 string `json:"ecotaxe" struct2map:"key:ecotaxe"`
	FraisDePort             string `json:"frais_de_port" struct2map:"key:frais_de_port"`
	Garantie                string `json:"garantie" struct2map:"key:garantie"`
	ID                      string `json:"id" struct2map:"key:id"`
	IndicateurDePerformance string `json:"indicateur_de_performance" struct2map:"key:indicateur_de_performance"`
	IndicateurDeStock       string `json:"indicateur_de_stock" struct2map:"key:indicateur_de_stock"`
	MarqueDuProduit         string `json:"marque_du_produit" struct2map:"key:marque_du_produit"`
	NModeleDuProduit        string `json:"n_modele_du_produit" struct2map:"key:n_modele_du_produit"`
	Nom                     string `json:"nom" struct2map:"key:nom"`
	Nouveaute               string `json:"nouveaute" struct2map:"key:nouveaute"`
	PrixActuel              string `json:"prix_actuel" struct2map:"key:prix_actuel"`
	Promotion               string `json:"promotion" struct2map:"key:promotion"`
	ReferenceConstructeur   string `json:"reference_constructeur" struct2map:"key:reference_constructeur"`
	URLDuProduit            string `json:"url_du_produit" struct2map:"key:url_du_produit"`
	URLGrande               string `json:"url_grande" struct2map:"key:url_grande"`
	URLImageMoyenne         string `json:"url_image_moyenne" struct2map:"key:url_image_moyenne"`
	URLImagePetite          string `json:"url_image_petite" struct2map:"key:url_image_petite"`
}

/*
func (o *Thalassoissambre) Convert() ([]string, error) {

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
