package netaffiliation

import (
	"fmt"

	// "github.com/fatih/structs"
	// "github.com/k0kubun/pp"
)

/*

// Catalog Path: catalogs/linvosges.csv
// Catalog Separator: _
// Catalog excerpt: #1

// ===================================================================================== \\
<?xml version="1.0" encoding="UTF-8"?><products><product><Nom><![CDATA[Cigales en concerto]]></Nom><reference: <product><Nom><![CDATA[Feuille à feuille]]></Nom><reference,
interne>998336101</reference: interne>999928801</reference,
interne><prix: interne><prix,
actuel>115.00</prix: actuel>34.00</prix,
actuel><categorie><![CDATA[Drap]]></categorie><URLproduit><![CDATA[https://xej.linvosges.com/?P497474598322BS1UD4128510192706V4]]></URLproduit><MPN /><marque><![CDATA[Linvosges]]></marque><EAN /><frais: actuel><categorie><![CDATA[Linge de lit fantaisie]]></categorie><URLproduit><![CDATA[https://xej.linvosges.com/?P497474598322BS1UC412851040611V4]]></URLproduit><MPN /><marque><![CDATA[Linvosges]]></marque><EAN /><frais,
de: de,
port>7.50</frais: port>7.50</frais,
de: de,
port><descriptif><![CDATA[  Caractéristiques Drap 1 à 2 pers. Cigales en concerto   Jeu de rayures, rehaussée d'un passepoil en velours. Impression coeur de fibre qui permet de fixer la couleur au coeur des fibres pour un toucher et une douceur exceptionnels. Percale 100% coton 80 fils/cm2. Lavage à 60°.  Dimensions du drap Cigales en concerto :   180x300 cm : 1 personne  240x310 cm : 2 personnes  270x310 cm : 2 personnes  ]]></descriptif><garantie /><indicateur: port><descriptif><![CDATA[ Un satin imprimé de délicates couleurs pastel au motif végétal associé à un uni rose poudre et rehaussé d'un pimpant passepoil en velours turquoise. Satin 100% coton 108 fils/cm2. Lavage à 60°. ]]></descriptif><garantie /><indicateur,
de: de,
stock>001</indicateur: stock>001</indicateur,
de: de,
stock><disponibilite /><indicateur: stock><disponibilite /><indicateur,
de: de,
performance /><indicateur: performance /><indicateur,
de: de,
promotion /><indicateur: promotion /><indicateur,
de: de,
nouveaute /><URL: nouveaute /><URL,
image: image,
petite /><URL: petite /><URL,
image: image,
moyenne /><URL: moyenne /><URL,
image: image,
grande><![CDATA[https://cdn.linvosges.com/mediatheque/images-categories/gammes/drap-cigales-en-concerto/zoom/drap-cigales-en-concerto.jpg]]></URL: grande><![CDATA[https://cdn.linvosges.com/mediatheque/images-categories/gammes/feuille-a-feuille-271/zoom/feuille-a-feuille-271.jpg]]></URL,
image: image,
grande><ecotaxe /><prix: grande><ecotaxe /><prix,
barre>115.00</prix: barre>34.00</prix,
barre><genre /><Categorie: barre><genre /><Categorie,
finale /><Debut: finale /><Debut,
de: de,
Remise /><Fin: Remise /><Fin,
de: de,
remise /><Taille /><Famille: remise /><Taille /><Famille,
de: de,
produit /><Couleur /><Délai: produit /><Couleur /><Délai,
de: de,
livraison /><Matière /></product>: livraison /><Matière /></product>
\\ ===================================================================================== //

*/

type Linvosge struct {
	XMLVersion10EncodingUtf8ProductsProductNomBangCdataCigalesEnConcertoNomReference                                                                                                                                                                                                                                                                                         string `json:"<?xml_version=1.0_encoding=utf_8?><products><product><nom><![cdata[cigales_en_concerto]]></nom><reference" struct2map:"key:<?xml_version=1.0_encoding=utf_8?><products><product><nom><![cdata[cigales_en_concerto]]></nom><reference"`
	Actuel11500Prix                                                                                                                                                                                                                                                                                                                                                          string `json:"actuel>115.00</prix" struct2map:"key:actuel>115.00</prix"`
	ActuelCategorieBangCdataDrapCategorieUrLproduitBangCdataHTTPSXejLinvosgesComp497474598322Bs1Ud4128510192706v4UrLproduitMpnMarqueBangCdataLinvosgesMarqueEanFrais                                                                                                                                                                                                         string `json:"actuel><categorie><![cdata[drap]]></categorie><ur_lproduit><![cdata[https://xej.linvosges.com/?p_497474598322_bs_1_ud_4128510192706_v_4]]></ur_lproduit><mpn_/><marque><![cdata[linvosges]]></marque><ean_/><frais" struct2map:"key:actuel><categorie><![cdata[drap]]></categorie><ur_lproduit><![cdata[https://xej.linvosges.com/?p_497474598322_bs_1_ud_4128510192706_v_4]]></ur_lproduit><mpn_/><marque><![cdata[linvosges]]></marque><ean_/><frais"`
	Barre11500Prix                                                                                                                                                                                                                                                                                                                                                           string `json:"barre>115.00</prix" struct2map:"key:barre>115.00</prix"`
	BarreGenreCategorie                                                                                                                                                                                                                                                                                                                                                      string `json:"barre><genre_/><categorie" struct2map:"key:barre><genre_/><categorie"`
	De                                                                                                                                                                                                                                                                                                                                                                       string `json:"de" struct2map:"key:de"`
	FinaleDebut                                                                                                                                                                                                                                                                                                                                                              string `json:"finale_/><debut" struct2map:"key:finale_/><debut"`
	GrandeBangCdataHTTPSCdnLinvosgesComMediathequeImagesCategoriesGammesDrapCigalesEnConcertoZoomDrapCigalesEnConcertoJpgURL                                                                                                                                                                                                                                                 string `json:"grande><![cdata[https://cdn.linvosges.com/mediatheque/images_categories/gammes/drap_cigales_en_concerto/zoom/drap_cigales_en_concerto.jpg]]></url" struct2map:"key:grande><![cdata[https://cdn.linvosges.com/mediatheque/images_categories/gammes/drap_cigales_en_concerto/zoom/drap_cigales_en_concerto.jpg]]></url"`
	GrandeEcotaxePrix                                                                                                                                                                                                                                                                                                                                                        string `json:"grande><ecotaxe_/><prix" struct2map:"key:grande><ecotaxe_/><prix"`
	Image                                                                                                                                                                                                                                                                                                                                                                    string `json:"image" struct2map:"key:image"`
	Interne998336101Reference                                                                                                                                                                                                                                                                                                                                                string `json:"interne>998336101</reference" struct2map:"key:interne>998336101</reference"`
	InternePrix                                                                                                                                                                                                                                                                                                                                                              string `json:"interne><prix" struct2map:"key:interne><prix"`
	LivraisonMatiereProduct                                                                                                                                                                                                                                                                                                                                                  string `json:"livraison_/><matiere_/></product>" struct2map:"key:livraison_/><matiere_/></product>"`
	MoyenneURL                                                                                                                                                                                                                                                                                                                                                               string `json:"moyenne_/><url" struct2map:"key:moyenne_/><url"`
	NouveauteURL                                                                                                                                                                                                                                                                                                                                                             string `json:"nouveaute_/><url" struct2map:"key:nouveaute_/><url"`
	PerformanceIndicateur                                                                                                                                                                                                                                                                                                                                                    string `json:"performance_/><indicateur" struct2map:"key:performance_/><indicateur"`
	PetiteURL                                                                                                                                                                                                                                                                                                                                                                string `json:"petite_/><url" struct2map:"key:petite_/><url"`
	Port750Frais                                                                                                                                                                                                                                                                                                                                                             string `json:"port>7.50</frais" struct2map:"key:port>7.50</frais"`
	PortDescriptifBangCdataCaracteristiquesDrap1a2PersCigalesEnConcertoJeuDeRayuresRehausseedUnPassepoilEnVeloursImpressionCoeurDeFibreQuiPermetDeFixerLaCouleurAuCoeurDesFibresPourUnToucherEtUneDouceurExceptionnelsPercale100Coton80FilsCm2Lavagea60DimensionsDuDrapCigalesEnConcerto180x300Cm1Personne240x310Cm2Personnes270x310Cm2PersonnesDescriptifGarantieIndicateur string `json:"port><descriptif><![cdata[__caracteristiques_drap_1_a_2_pers._cigales_en_concerto___jeu_de_rayures,_rehaussee_d'un_passepoil_en_velours._impression_coeur_de_fibre_qui_permet_de_fixer_la_couleur_au_coeur_des_fibres_pour_un_toucher_et_une_douceur_exceptionnels._percale_100%_coton_80_fils/cm_2._lavage_a_60°.__dimensions_du_drap_cigales_en_concerto_:___180_x_300_cm_:_1_personne__240_x_310_cm_:_2_personnes__270_x_310_cm_:_2_personnes__]]></descriptif><garantie_/><indicateur" struct2map:"key:port><descriptif><![cdata[__caracteristiques_drap_1_a_2_pers._cigales_en_concerto___jeu_de_rayures,_rehaussee_d'un_passepoil_en_velours._impression_coeur_de_fibre_qui_permet_de_fixer_la_couleur_au_coeur_des_fibres_pour_un_toucher_et_une_douceur_exceptionnels._percale_100%_coton_80_fils/cm_2._lavage_a_60°.__dimensions_du_drap_cigales_en_concerto_:___180_x_300_cm_:_1_personne__240_x_310_cm_:_2_personnes__270_x_310_cm_:_2_personnes__]]></descriptif><garantie_/><indicateur"`
	ProduitCouleurDelai                                                                                                                                                                                                                                                                                                                                                      string `json:"produit_/><couleur_/><delai" struct2map:"key:produit_/><couleur_/><delai"`
	PromotionIndicateur                                                                                                                                                                                                                                                                                                                                                      string `json:"promotion_/><indicateur" struct2map:"key:promotion_/><indicateur"`
	RemiseFin                                                                                                                                                                                                                                                                                                                                                                string `json:"remise_/><fin" struct2map:"key:remise_/><fin"`
	RemiseTailleFamille                                                                                                                                                                                                                                                                                                                                                      string `json:"remise_/><taille_/><famille" struct2map:"key:remise_/><taille_/><famille"`
	Stock001Indicateur                                                                                                                                                                                                                                                                                                                                                       string `json:"stock>001</indicateur" struct2map:"key:stock>001</indicateur"`
	StockDisponibiliteIndicateur                                                                                                                                                                                                                                                                                                                                             string `json:"stock><disponibilite_/><indicateur" struct2map:"key:stock><disponibilite_/><indicateur"`
}

/*
func (o *Linvosge) Convert() ([]string, error) {

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
