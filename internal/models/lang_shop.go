package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _LangShopMgr struct {
	*_BaseMgr
}

func LangShopMgr(db *gorm.DB) *_LangShopMgr {
	if db == nil {
		panic(fmt.Errorf("LangShopMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_LangShopMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_lang_shop"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_LangShopMgr) GetTableName() string {
	return "ps_lang_shop"
}

func (obj *_LangShopMgr) Get() (result LangShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_LangShopMgr) Gets() (results []*LangShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_LangShopMgr) WithIDLang(idLang int) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

func (obj *_LangShopMgr) WithIDShop(idShop int) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

func (obj *_LangShopMgr) GetByOption(opts ...Option) (result LangShop, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_LangShopMgr) GetByOptions(opts ...Option) (results []*LangShop, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_LangShopMgr) GetFromIDLang(idLang int) (results []*LangShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error

	return
}

func (obj *_LangShopMgr) GetBatchFromIDLang(idLangs []int) (results []*LangShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error

	return
}

func (obj *_LangShopMgr) GetFromIDShop(idShop int) (results []*LangShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

func (obj *_LangShopMgr) GetBatchFromIDShop(idShops []int) (results []*LangShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}


func (obj *_LangShopMgr) FetchByPrimaryKey(idLang int, idShop int) (result LangShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ? AND id_shop = ?", idLang, idShop).Find(&result).Error

	return
}

func (obj *_LangShopMgr) FetchIndexByIDXA5D79262BA299860(idLang int) (results []*LangShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error

	return
}

func (obj *_LangShopMgr) FetchIndexByIDXA5D79262274A50A0(idShop int) (results []*LangShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}
