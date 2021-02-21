package netaffiliation

import (
	"fmt"

	// "github.com/fatih/structs"
	// "github.com/k0kubun/pp"
)

/*

// Catalog Path: catalogs/hypnia.csv
// Catalog Separator: |
// Catalog excerpt: #1

// ===================================================================================== \\
Nom: Matelas Bien-être Suprême-90 x 190 (cm),
reference interne: 59,
prix actuel: 424,
categorie: Hypnia > Matelas > Matelas ressort,
URLproduit: https://zlm.hypnia.fr/?P4BDC545983219S1UA4127cf0440V4,
MPN: ,
marque: Hypnia,
EAN: 3663295288415,
frais de port: 0,
descriptif: Ce matelas se distingue par la qualité des matériaux et ses performances exceptionnelles Un matelas haut de gamme de 30 cm d'épaisseur pour les plus exigeants en matière de confort et bien-être : profitez du meilleur de la mémoire de forme haute densité et des ressorts ensachés. D'autre part, la couche de latex naturel vous permettra de lutter contre les allergies et de bénéficier d'un confort de sommeil sans pareil. Ce matelas est à la fois sain, souple et indéformable ! L'âme du matelas, composée de ressorts ensachés et de 7 zones de confort optimisées (tête, épaules, dos, lombaires, fessiers, cuisses et jambes), vous permettra de bénéficier d'un confort sur-mesure. Son tissu doux et léger à base de bambou vous procurera un bien-être durable tout au long de la nuit. Profitez chaque nuit d'un confort exceptionnel et d'une qualité de couchage unique !,
garantie: Garantie 10 ans,
indicateur de stock: 83,
disponibilite: ,
indicateur de performance: ,
indicateur de promotion: 1,
indicateur de nouveaute: ,
URL image petite: https://www.hypnia.fr/media/catalog/product/m/a/matelas-bien-etre-suprem.jpg,
URL image moyenne: https://www.hypnia.fr/media/catalog/product/s/c/schema-bien-etre-supreme-hypnia.jpg,
URL image grande: https://www.hypnia.fr/media/catalog/product/p/u/pur-7-zones-photo1.jpg,
ecotaxe: 4.0000,
prix barre: 499,
genre: Unisexe,
Categorie_finale: Matelas ressort,
Debut de Remise: ,
Fin de remise: ,
Taille: 90 x 190 (cm),
Famille de produit: Matelas,
Couleur: Blanc,
Délai de livraison: 7 jours ,
Matière: Mousse à mémoire de forme,
: 
\\ ===================================================================================== //

*/

type Hypnium struct {
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
func (o *Hypnium) Convert() ([]string, error) {

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
