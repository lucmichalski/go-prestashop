package netaffiliation

import (
	"fmt"

	// "github.com/fatih/structs"
	// "github.com/k0kubun/pp"
)

/*

// Catalog Path: catalogs/bazaravenue.csv
// Catalog Separator: |
// Catalog excerpt: #1

// ===================================================================================== \\
Nom: Hamac Headdemock, Fatboy,
reference interne: 21,
prix actuel: 419.00,
categorie: Accueil > Meuble > Meuble d'extérieur > Bains de soleil et transats,
URLproduit: https://action.metaffiliation.com/trk.php?mclic=P44F9F459831FS1U841286501V4,
MPN: ,
marque: FATBOY,
EAN: 0,
frais de port: 4.90,
descriptif: Seul ou à deux, offrez-vous un moment de pure détente en toute simplicité avec le hamac Headdemock de Fatboy. Pas besoin d'arbres, avec son support il s'installe partout en un clin d'il! Avec ce hamac au style ultra tendance, Fatboy innove et vous propose de créer un coin de détente à l'extérieur, pour savourer un moment de repos seul ou avec vos proches. Ce hamac grand format peut en effet accueillir deux personnes très confortablement sur sa toile rembourrée. Robuste et fonctionnel, il possède une structure unique qui vous permet de le monter et démonter rapidement pour le déplacer au gré de vos envies. Vous partez en weekend ou en vacances ? A l'aide du sac de transport, votre hamac vous accompagne en voyage! Polyvalent, Headdemock peut être utilisé à l'intérieur comme à l'extérieur grâce à son revêtement en polyester très résistant, que vous nettoierez aisément à l'aide d'un chiffon doux et humide. Chez vous, à la plage ou pour un pique-nique, Headdemock se déploiera élégamment et il fera le bonheur de tous les amateurs de sieste! Caractéristiques : Dimensions (L x l x h) : 330 x 127 x 110 cm Contient 1 structure de soutien, 1 toile Vendu dans un sac muni de poignées pour un transport facile Toile en polyester imperméable et imputrescible Rembourrage 100% fibres de polyester Structure en acier avec patins antidérapants en caoutchouc Poids maximal supporté 150 Kg Nettoyage à l'eau tiède et au savon,
garantie: ,
indicateur de stock: ,
disponibilite: ,
indicateur de performance: ,
indicateur de promotion: ,
indicateur de nouveaute: ,
URL image petite: http://www.bazaravenue.com/17-large_default/hamac-headdemock-fatboy.jpg,
URL image moyenne: http://www.bazaravenue.com/18-large_default/hamac-headdemock-fatboy.jpg,
URL image grande: http://www.bazaravenue.com/11906-large_default/hamac-headdemock-fatboy.jpg,
ecotaxe: ,
prix barre: 419.00,
genre: ,
Categorie_finale: ,
Debut de Remise: ,
Fin de remise: ,
Taille: ,
Famille de produit: ,
Couleur: ,
Délai de livraison: ,
Matière: ,
: 
\\ ===================================================================================== //

*/

type Bazaravenue struct {
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
func (o *Bazaravenue) Convert() ([]string, error) {

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
