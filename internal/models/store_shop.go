package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _StoreShopMgr struct {
	*_BaseMgr
}

// StoreShopMgr open func
func StoreShopMgr(db *gorm.DB) *_StoreShopMgr {
	if db == nil {
		panic(fmt.Errorf("StoreShopMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_StoreShopMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_store_shop"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_StoreShopMgr) GetTableName() string {
	return "ps_store_shop"
}

// Get 获取
func (obj *_StoreShopMgr) Get() (result StoreShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_StoreShopMgr) Gets() (results []*StoreShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDStore id_store获取
func (obj *_StoreShopMgr) WithIDStore(idStore uint32) Option {
	return optionFunc(func(o *options) { o.query["id_store"] = idStore })
}

// WithIDShop id_shop获取
func (obj *_StoreShopMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

// GetByOption 功能选项模式获取
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

// GetByOptions 批量功能选项模式获取
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

//////////////////////////enume case ////////////////////////////////////////////

// GetFromIDStore 通过id_store获取内容
func (obj *_StoreShopMgr) GetFromIDStore(idStore uint32) (results []*StoreShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_store = ?", idStore).Find(&results).Error

	return
}

// GetBatchFromIDStore 批量唯一主键查找
func (obj *_StoreShopMgr) GetBatchFromIDStore(idStores []uint32) (results []*StoreShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_store IN (?)", idStores).Find(&results).Error

	return
}

// GetFromIDShop 通过id_shop获取内容
func (obj *_StoreShopMgr) GetFromIDShop(idShop uint32) (results []*StoreShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

// GetBatchFromIDShop 批量唯一主键查找
func (obj *_StoreShopMgr) GetBatchFromIDShop(idShops []uint32) (results []*StoreShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_StoreShopMgr) FetchByPrimaryKey(idStore uint32, idShop uint32) (result StoreShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_store = ? AND id_shop = ?", idStore, idShop).Find(&result).Error

	return
}

// FetchIndexByIDShop  获取多个内容
func (obj *_StoreShopMgr) FetchIndexByIDShop(idShop uint32) (results []*StoreShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}
