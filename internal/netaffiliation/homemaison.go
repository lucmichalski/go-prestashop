package netaffiliation

import (
	"fmt"

	// "github.com/fatih/structs"
	// "github.com/k0kubun/pp"
)

/*

// Catalog Path: catalogs/homemaison.csv
// Catalog Separator: |
// Catalog excerpt: #1

// ===================================================================================== \\
Nom: Tube nu en Fer Forgé Rouille 2 m diam 28mm - Rouille - 2 m,
reference interne: 1071910HM1267,
prix actuel: 29.90,
categorie: Tubes tringles,
URLproduit: https://action.metaffiliation.com/trk.php?mclic=P455C745983233S1U841291308V4,
MPN: HM1071910HM1267,
marque: HomeMaison,
EAN: 8420938405432,
frais de port: 4.9,
descriptif: Vendu à l'unitéTube nu en fer forgé.Diamètre : 28 mmLongueur :2 mColoris :  rouilleAttention, les embouts et supports sont vendus séparéments.Différents embouts possibles dans la catégorie  embouts Différents supports ( mural ou plafond ) disponibles dans la catégorie  supports ,
garantie: 0,
indicateur de stock: 1,
disponibilite: 770,
indicateur de performance: ,
indicateur de promotion: 2,
indicateur de nouveaute: ,
URL image petite: ,
URL image moyenne: ,
URL image grande: https://media.homemaison.com/Produit/img4/tube-nu-en-fer-forge-rouille-2-m-rouille.jpg,
ecotaxe: ,
prix barre: 29.90,
genre: H,
Categorie_finale: ,
Debut de Remise: ,
Fin de remise: ,
Taille: 2 m,
Famille de produit: Tubes tringles,
Couleur: Rouille,
Délai de livraison: 5 jours,
Matière: ,
: 
\\ ===================================================================================== //

*/

type Homemaison struct {
	Categorie               string `json:"categorie" struct2map:"key:categorie"`
	CategorieFinale         string `json:"categorie_finale" struct2map:"key:categorie_finale"`
	Couleur                 string `json:"couleur" struct2map:"key:couleur"`
	DebutDeRemise           string `json:"debut_de_remise" struct2map:"key:debut_de_remise"`
	DelaiDeLivraison        string `json:"delai_de_livraison" struct2map:"key:delai_de_livraison"`
	Descriptif              string `json:"descriptif" struct2map:"key:descriptif"`
	Disponibilite           string `json:"disponibilite" struct2map:"key:disponibilite"`
	Ean                     string `json:"ean" struct2map:"key:ean"`
	Ecotaxe                 string `json:"ecotaxe" struct2map:"key:ecotaxe"`
	FamilleDeProduit        string `json:"famille_de_produit" struct2map:"key:famille_de_produit"`
	FinDeRemise             string `json:"fin_de_remise" struct2map:"key:fin_de_remise"`
	FraisDePort             string `json:"frais_de_port" struct2map:"key:frais_de_port"`
	Garantie                string `json:"garantie" struct2map:"key:garantie"`
	Genre                   string `json:"genre" struct2map:"key:genre"`
	IndicateurDeNouveaute   string `json:"indicateur_de_nouveaute" struct2map:"key:indicateur_de_nouveaute"`
	IndicateurDePerformance string `json:"indicateur_de_performance" struct2map:"key:indicateur_de_performance"`
	IndicateurDePromotion   string `json:"indicateur_de_promotion" struct2map:"key:indicateur_de_promotion"`
	IndicateurDeStock       string `json:"indicateur_de_stock" struct2map:"key:indicateur_de_stock"`
	Marque                  string `json:"marque" struct2map:"key:marque"`
	Matiere                 string `json:"matiere" struct2map:"key:matiere"`
	Mpn                     string `json:"mpn" struct2map:"key:mpn"`
	Nom                     string `json:"nom" struct2map:"key:nom"`
	PrixActuel              string `json:"prix_actuel" struct2map:"key:prix_actuel"`
	PrixBarre               string `json:"prix_barre" struct2map:"key:prix_barre"`
	ReferenceInterne        string `json:"reference_interne" struct2map:"key:reference_interne"`
	Taille                  string `json:"taille" struct2map:"key:taille"`
	UrLproduit              string `json:"ur_lproduit" struct2map:"key:ur_lproduit"`
	URLImageGrande          string `json:"url_image_grande" struct2map:"key:url_image_grande"`
	URLImageMoyenne         string `json:"url_image_moyenne" struct2map:"key:url_image_moyenne"`
	URLImagePetite          string `json:"url_image_petite" struct2map:"key:url_image_petite"`
}

/*
func (o *Homemaison) Convert() ([]string, error) {

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
