package netaffiliation

import (
	"strings"

	"gorm.io/gorm"

	"github.com/lucmichalski/go-prestashop/internal/models"
)

var ( 
	db *gorm.DB
	dbTablePrefix string
	Catalogs = []interface{}{
					
					&Flux2{},
					&Standard{},
					&Villatech{},
					&DcshoesFr{},
					&FannyChaussure{},
					&LentillesMoinsChere{},
					&Bazaravenue{},
					&Standard{},
					&LauraKent{},
					&Standard{},
					&Standard{},
					&Bricorama{},
					&Standard{},
					&Standard{},
					&GsellMaroquinerieEtArticle{},
					&Linvosge{},
					&PneusOnlineFr{},
					&TradeDiscount{},
					&Abix{},
					&FamilleMaryMielEtPlante{},
					&Roxy{},
					&Quiksilver{},
					&Meltonic{},
					&Standard{},
					&RougierAndPle{},
					&Standard{},
					&FouDePuzzle{},
					&CheminDeTable{},
					&Standard{},
					&Standard{},
					&EspacePlaisir{},
					&Condomz{},
					&Thalassoissambre{},
					&Standard{},
					&Vitalco{},
					&MonPuzzlePhoto{},
					&Homemaison{},
					&Standard{},
					&Layole{},
					&Hypnium{},
					&Standard{},
					&ShopMini{},
					&Craftine{},
					&ClefEnLigne{},
					&BoutiqueArthur{},
					&LeRoiDuMatelasFr{},
					&Deelux74{},
					&OriginalCup{},
					&Standard{},
					&ZafulFr{},
					&DresslilyFr{},
					&Standard{},
					&CapCadeau{},
					&LovelyDay{},
					&Zermara{},
					&Ripauste{},
					&Standard{},
					&MaisonFt{},
					&Standard{},
					&Dermoplant{},
					&MoulinDesSenteur{},
					&Standard{},
					&Newchic{},
					&Standard{},
					&CreaPlaisir{},
					&Standard{},
					&Bearco{},
					&VeniVici{},
				}
)

func initDB(client *gorm.DB, tablePrefix string) *gorm.DB {
	db = client
	dbTablePrefix = tablePrefix
	return db
}

func findOrCreateCategoryByName(db *gorm.DB, dbTablePrefix, name string) (*models.CategoryLang, error) {
	var cat models.CategoryLang
	db.Debug().Raw("SELECT * FROM " + dbTablePrefix + "category_lang WHERE name = '" + strings.ToLower(name) + "'").Scan(&cat)
	if cat.IDCategory == 0 {
		db.Debug().Raw("SELECT * FROM eg_lang WHERE iso_code = '" + strings.ToLower(iso_code) + "'").Scan(&lang)
		return &lang, nil // todo. fix this hacky return
	}
	return &cat, nil
}