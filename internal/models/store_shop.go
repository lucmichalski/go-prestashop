package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _StoreShopMgr struct {
	*_BaseMgr
}

func StoreShopMgr(db *gorm.DB) *_StoreShopMgr {
	if db == nil {
		panic(fmt.Errorf("StoreShopMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_StoreShopMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_store_shop"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_StoreShopMgr) GetTableName() string {
	return "ps_store_shop"
}

func (obj *_StoreShopMgr) Get() (result StoreShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_StoreShopMgr) Gets() (results []*StoreShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

func (obj *_StoreShopMgr) WithIDStore(idStore uint32) Option {
	return optionFunc(func(o *options) { o.query["id_store"] = idStore })
}

func (obj *_StoreShopMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

func (obj *_StoreShopMgr) GetByOption(opts ...Option) (result StoreShop, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_StoreShopMgr) GetByOptions(opts ...Option) (results []*StoreShop, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

func (obj *_StoreShopMgr) GetFromIDStore(idStore uint32) (results []*StoreShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_store = ?", idStore).Find(&results).Error

	return
}

func (obj *_StoreShopMgr) GetBatchFromIDStore(idStores []uint32) (results []*StoreShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_store IN (?)", idStores).Find(&results).Error

	return
}

func (obj *_StoreShopMgr) GetFromIDShop(idShop uint32) (results []*StoreShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

func (obj *_StoreShopMgr) GetBatchFromIDShop(idShops []uint32) (results []*StoreShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

func (obj *_StoreShopMgr) FetchByPrimaryKey(idStore uint32, idShop uint32) (result StoreShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_store = ? AND id_shop = ?", idStore, idShop).Find(&result).Error

	return
}

func (obj *_StoreShopMgr) FetchIndexByIDShop(idShop uint32) (results []*StoreShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}
