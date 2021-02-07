package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _WebserviceAccountShopMgr struct {
	*_BaseMgr
}

func WebserviceAccountShopMgr(db *gorm.DB) *_WebserviceAccountShopMgr {
	if db == nil {
		panic(fmt.Errorf("WebserviceAccountShopMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_WebserviceAccountShopMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_webservice_account_shop"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_WebserviceAccountShopMgr) GetTableName() string {
	return "ps_webservice_account_shop"
}

func (obj *_WebserviceAccountShopMgr) Get() (result WebserviceAccountShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_WebserviceAccountShopMgr) Gets() (results []*WebserviceAccountShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

func (obj *_WebserviceAccountShopMgr) WithIDWebserviceAccount(idWebserviceAccount uint32) Option {
	return optionFunc(func(o *options) { o.query["id_webservice_account"] = idWebserviceAccount })
}

func (obj *_WebserviceAccountShopMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

func (obj *_WebserviceAccountShopMgr) GetByOption(opts ...Option) (result WebserviceAccountShop, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_WebserviceAccountShopMgr) GetByOptions(opts ...Option) (results []*WebserviceAccountShop, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

func (obj *_WebserviceAccountShopMgr) GetFromIDWebserviceAccount(idWebserviceAccount uint32) (results []*WebserviceAccountShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_webservice_account = ?", idWebserviceAccount).Find(&results).Error

	return
}

func (obj *_WebserviceAccountShopMgr) GetBatchFromIDWebserviceAccount(idWebserviceAccounts []uint32) (results []*WebserviceAccountShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_webservice_account IN (?)", idWebserviceAccounts).Find(&results).Error

	return
}

func (obj *_WebserviceAccountShopMgr) GetFromIDShop(idShop uint32) (results []*WebserviceAccountShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

func (obj *_WebserviceAccountShopMgr) GetBatchFromIDShop(idShops []uint32) (results []*WebserviceAccountShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

func (obj *_WebserviceAccountShopMgr) FetchByPrimaryKey(idWebserviceAccount uint32, idShop uint32) (result WebserviceAccountShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_webservice_account = ? AND id_shop = ?", idWebserviceAccount, idShop).Find(&result).Error

	return
}

func (obj *_WebserviceAccountShopMgr) FetchIndexByIDShop(idShop uint32) (results []*WebserviceAccountShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}
