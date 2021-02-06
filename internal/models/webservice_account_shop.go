package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _WebserviceAccountShopMgr struct {
	*_BaseMgr
}

// WebserviceAccountShopMgr open func
func WebserviceAccountShopMgr(db *gorm.DB) *_WebserviceAccountShopMgr {
	if db == nil {
		panic(fmt.Errorf("WebserviceAccountShopMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_WebserviceAccountShopMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_webservice_account_shop"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_WebserviceAccountShopMgr) GetTableName() string {
	return "ps_webservice_account_shop"
}

// Get 获取
func (obj *_WebserviceAccountShopMgr) Get() (result WebserviceAccountShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_WebserviceAccountShopMgr) Gets() (results []*WebserviceAccountShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDWebserviceAccount id_webservice_account获取
func (obj *_WebserviceAccountShopMgr) WithIDWebserviceAccount(idWebserviceAccount uint32) Option {
	return optionFunc(func(o *options) { o.query["id_webservice_account"] = idWebserviceAccount })
}

// WithIDShop id_shop获取
func (obj *_WebserviceAccountShopMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

// GetByOption 功能选项模式获取
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

// GetByOptions 批量功能选项模式获取
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

//////////////////////////enume case ////////////////////////////////////////////

// GetFromIDWebserviceAccount 通过id_webservice_account获取内容
func (obj *_WebserviceAccountShopMgr) GetFromIDWebserviceAccount(idWebserviceAccount uint32) (results []*WebserviceAccountShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_webservice_account = ?", idWebserviceAccount).Find(&results).Error

	return
}

// GetBatchFromIDWebserviceAccount 批量唯一主键查找
func (obj *_WebserviceAccountShopMgr) GetBatchFromIDWebserviceAccount(idWebserviceAccounts []uint32) (results []*WebserviceAccountShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_webservice_account IN (?)", idWebserviceAccounts).Find(&results).Error

	return
}

// GetFromIDShop 通过id_shop获取内容
func (obj *_WebserviceAccountShopMgr) GetFromIDShop(idShop uint32) (results []*WebserviceAccountShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

// GetBatchFromIDShop 批量唯一主键查找
func (obj *_WebserviceAccountShopMgr) GetBatchFromIDShop(idShops []uint32) (results []*WebserviceAccountShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_WebserviceAccountShopMgr) FetchByPrimaryKey(idWebserviceAccount uint32, idShop uint32) (result WebserviceAccountShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_webservice_account = ? AND id_shop = ?", idWebserviceAccount, idShop).Find(&result).Error

	return
}

// FetchIndexByIDShop  获取多个内容
func (obj *_WebserviceAccountShopMgr) FetchIndexByIDShop(idShop uint32) (results []*WebserviceAccountShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}
