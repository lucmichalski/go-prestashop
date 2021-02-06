package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _ContactShopMgr struct {
	*_BaseMgr
}

func ContactShopMgr(db *gorm.DB) *_ContactShopMgr {
	if db == nil {
		panic(fmt.Errorf("ContactShopMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ContactShopMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_contact_shop"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_ContactShopMgr) GetTableName() string {
	return "ps_contact_shop"
}

func (obj *_ContactShopMgr) Get() (result ContactShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_ContactShopMgr) Gets() (results []*ContactShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_ContactShopMgr) WithIDContact(idContact uint32) Option {
	return optionFunc(func(o *options) { o.query["id_contact"] = idContact })
}

func (obj *_ContactShopMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

func (obj *_ContactShopMgr) GetByOption(opts ...Option) (result ContactShop, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_ContactShopMgr) GetByOptions(opts ...Option) (results []*ContactShop, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_ContactShopMgr) GetFromIDContact(idContact uint32) (results []*ContactShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_contact = ?", idContact).Find(&results).Error

	return
}

func (obj *_ContactShopMgr) GetBatchFromIDContact(idContacts []uint32) (results []*ContactShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_contact IN (?)", idContacts).Find(&results).Error

	return
}

func (obj *_ContactShopMgr) GetFromIDShop(idShop uint32) (results []*ContactShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

func (obj *_ContactShopMgr) GetBatchFromIDShop(idShops []uint32) (results []*ContactShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}


func (obj *_ContactShopMgr) FetchByPrimaryKey(idContact uint32, idShop uint32) (result ContactShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_contact = ? AND id_shop = ?", idContact, idShop).Find(&result).Error

	return
}

func (obj *_ContactShopMgr) FetchIndexByIDShop(idShop uint32) (results []*ContactShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}
