package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _SupplierShopMgr struct {
	*_BaseMgr
}

// SupplierShopMgr open func
func SupplierShopMgr(db *gorm.DB) *_SupplierShopMgr {
	if db == nil {
		panic(fmt.Errorf("SupplierShopMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SupplierShopMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_supplier_shop"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SupplierShopMgr) GetTableName() string {
	return "ps_supplier_shop"
}

// Get 获取
func (obj *_SupplierShopMgr) Get() (result SupplierShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SupplierShopMgr) Gets() (results []*SupplierShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDSupplier id_supplier获取
func (obj *_SupplierShopMgr) WithIDSupplier(idSupplier uint32) Option {
	return optionFunc(func(o *options) { o.query["id_supplier"] = idSupplier })
}

// WithIDShop id_shop获取
func (obj *_SupplierShopMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

// GetByOption 功能选项模式获取
func (obj *_SupplierShopMgr) GetByOption(opts ...Option) (result SupplierShop, err error) {
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
func (obj *_SupplierShopMgr) GetByOptions(opts ...Option) (results []*SupplierShop, err error) {
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

// GetFromIDSupplier 通过id_supplier获取内容
func (obj *_SupplierShopMgr) GetFromIDSupplier(idSupplier uint32) (results []*SupplierShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supplier = ?", idSupplier).Find(&results).Error

	return
}

// GetBatchFromIDSupplier 批量唯一主键查找
func (obj *_SupplierShopMgr) GetBatchFromIDSupplier(idSuppliers []uint32) (results []*SupplierShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supplier IN (?)", idSuppliers).Find(&results).Error

	return
}

// GetFromIDShop 通过id_shop获取内容
func (obj *_SupplierShopMgr) GetFromIDShop(idShop uint32) (results []*SupplierShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

// GetBatchFromIDShop 批量唯一主键查找
func (obj *_SupplierShopMgr) GetBatchFromIDShop(idShops []uint32) (results []*SupplierShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_SupplierShopMgr) FetchByPrimaryKey(idSupplier uint32, idShop uint32) (result SupplierShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supplier = ? AND id_shop = ?", idSupplier, idShop).Find(&result).Error

	return
}

// FetchIndexByIDShop  获取多个内容
func (obj *_SupplierShopMgr) FetchIndexByIDShop(idShop uint32) (results []*SupplierShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}
