package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _ProductCarrierMgr struct {
	*_BaseMgr
}

// ProductCarrierMgr open func
func ProductCarrierMgr(db *gorm.DB) *_ProductCarrierMgr {
	if db == nil {
		panic(fmt.Errorf("ProductCarrierMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ProductCarrierMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_product_carrier"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ProductCarrierMgr) GetTableName() string {
	return "eg_product_carrier"
}

// Get 获取
func (obj *_ProductCarrierMgr) Get() (result ProductCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ProductCarrierMgr) Gets() (results []*ProductCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDProduct id_product获取
func (obj *_ProductCarrierMgr) WithIDProduct(idProduct uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product"] = idProduct })
}

// WithIDCarrierReference id_carrier_reference获取
func (obj *_ProductCarrierMgr) WithIDCarrierReference(idCarrierReference uint32) Option {
	return optionFunc(func(o *options) { o.query["id_carrier_reference"] = idCarrierReference })
}

// WithIDShop id_shop获取
func (obj *_ProductCarrierMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

// GetByOption 功能选项模式获取
func (obj *_ProductCarrierMgr) GetByOption(opts ...Option) (result ProductCarrier, err error) {
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
func (obj *_ProductCarrierMgr) GetByOptions(opts ...Option) (results []*ProductCarrier, err error) {
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

// GetFromIDProduct 通过id_product获取内容
func (obj *_ProductCarrierMgr) GetFromIDProduct(idProduct uint32) (results []*ProductCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ?", idProduct).Find(&results).Error

	return
}

// GetBatchFromIDProduct 批量唯一主键查找
func (obj *_ProductCarrierMgr) GetBatchFromIDProduct(idProducts []uint32) (results []*ProductCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product IN (?)", idProducts).Find(&results).Error

	return
}

// GetFromIDCarrierReference 通过id_carrier_reference获取内容
func (obj *_ProductCarrierMgr) GetFromIDCarrierReference(idCarrierReference uint32) (results []*ProductCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_carrier_reference = ?", idCarrierReference).Find(&results).Error

	return
}

// GetBatchFromIDCarrierReference 批量唯一主键查找
func (obj *_ProductCarrierMgr) GetBatchFromIDCarrierReference(idCarrierReferences []uint32) (results []*ProductCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_carrier_reference IN (?)", idCarrierReferences).Find(&results).Error

	return
}

// GetFromIDShop 通过id_shop获取内容
func (obj *_ProductCarrierMgr) GetFromIDShop(idShop uint32) (results []*ProductCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

// GetBatchFromIDShop 批量唯一主键查找
func (obj *_ProductCarrierMgr) GetBatchFromIDShop(idShops []uint32) (results []*ProductCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_ProductCarrierMgr) FetchByPrimaryKey(idProduct uint32, idCarrierReference uint32, idShop uint32) (result ProductCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ? AND id_carrier_reference = ? AND id_shop = ?", idProduct, idCarrierReference, idShop).Find(&result).Error

	return
}
