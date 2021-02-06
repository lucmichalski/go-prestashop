package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _CustomizationFieldLangMgr struct {
	*_BaseMgr
}

func CustomizationFieldLangMgr(db *gorm.DB) *_CustomizationFieldLangMgr {
	if db == nil {
		panic(fmt.Errorf("CustomizationFieldLangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CustomizationFieldLangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_customization_field_lang"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_CustomizationFieldLangMgr) GetTableName() string {
	return "ps_customization_field_lang"
}

func (obj *_CustomizationFieldLangMgr) Get() (result CustomizationFieldLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_CustomizationFieldLangMgr) Gets() (results []*CustomizationFieldLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_CustomizationFieldLangMgr) WithIDCustomizationField(idCustomizationField uint32) Option {
	return optionFunc(func(o *options) { o.query["id_customization_field"] = idCustomizationField })
}

func (obj *_CustomizationFieldLangMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

func (obj *_CustomizationFieldLangMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

func (obj *_CustomizationFieldLangMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

func (obj *_CustomizationFieldLangMgr) GetByOption(opts ...Option) (result CustomizationFieldLang, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_CustomizationFieldLangMgr) GetByOptions(opts ...Option) (results []*CustomizationFieldLang, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_CustomizationFieldLangMgr) GetFromIDCustomizationField(idCustomizationField uint32) (results []*CustomizationFieldLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customization_field = ?", idCustomizationField).Find(&results).Error

	return
}

func (obj *_CustomizationFieldLangMgr) GetBatchFromIDCustomizationField(idCustomizationFields []uint32) (results []*CustomizationFieldLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customization_field IN (?)", idCustomizationFields).Find(&results).Error

	return
}

func (obj *_CustomizationFieldLangMgr) GetFromIDLang(idLang uint32) (results []*CustomizationFieldLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error

	return
}

func (obj *_CustomizationFieldLangMgr) GetBatchFromIDLang(idLangs []uint32) (results []*CustomizationFieldLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error

	return
}

func (obj *_CustomizationFieldLangMgr) GetFromIDShop(idShop uint32) (results []*CustomizationFieldLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

func (obj *_CustomizationFieldLangMgr) GetBatchFromIDShop(idShops []uint32) (results []*CustomizationFieldLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

func (obj *_CustomizationFieldLangMgr) GetFromName(name string) (results []*CustomizationFieldLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

func (obj *_CustomizationFieldLangMgr) GetBatchFromName(names []string) (results []*CustomizationFieldLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}


func (obj *_CustomizationFieldLangMgr) FetchByPrimaryKey(idCustomizationField uint32, idLang uint32, idShop uint32) (result CustomizationFieldLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customization_field = ? AND id_lang = ? AND id_shop = ?", idCustomizationField, idLang, idShop).Find(&result).Error

	return
}
