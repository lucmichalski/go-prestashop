package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _WarehouseProductLocationMgr struct {
	*_BaseMgr
}

// WarehouseProductLocationMgr open func
func WarehouseProductLocationMgr(db *gorm.DB) *_WarehouseProductLocationMgr {
	if db == nil {
		panic(fmt.Errorf("WarehouseProductLocationMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_WarehouseProductLocationMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_warehouse_product_location"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_WarehouseProductLocationMgr) GetTableName() string {
	return "eg_warehouse_product_location"
}

// Get 获取
func (obj *_WarehouseProductLocationMgr) Get() (result WarehouseProductLocation, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_WarehouseProductLocationMgr) Gets() (results []*WarehouseProductLocation, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDWarehouseProductLocation id_warehouse_product_location获取
func (obj *_WarehouseProductLocationMgr) WithIDWarehouseProductLocation(idWarehouseProductLocation uint32) Option {
	return optionFunc(func(o *options) { o.query["id_warehouse_product_location"] = idWarehouseProductLocation })
}

// WithIDProduct id_product获取
func (obj *_WarehouseProductLocationMgr) WithIDProduct(idProduct uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product"] = idProduct })
}

// WithIDProductAttribute id_product_attribute获取
func (obj *_WarehouseProductLocationMgr) WithIDProductAttribute(idProductAttribute uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product_attribute"] = idProductAttribute })
}

// WithIDWarehouse id_warehouse获取
func (obj *_WarehouseProductLocationMgr) WithIDWarehouse(idWarehouse uint32) Option {
	return optionFunc(func(o *options) { o.query["id_warehouse"] = idWarehouse })
}

// WithLocation location获取
func (obj *_WarehouseProductLocationMgr) WithLocation(location string) Option {
	return optionFunc(func(o *options) { o.query["location"] = location })
}

// GetByOption 功能选项模式获取
func (obj *_WarehouseProductLocationMgr) GetByOption(opts ...Option) (result WarehouseProductLocation, err error) {
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
func (obj *_WarehouseProductLocationMgr) GetByOptions(opts ...Option) (results []*WarehouseProductLocation, err error) {
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

// GetFromIDWarehouseProductLocation 通过id_warehouse_product_location获取内容
func (obj *_WarehouseProductLocationMgr) GetFromIDWarehouseProductLocation(idWarehouseProductLocation uint32) (result WarehouseProductLocation, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_warehouse_product_location = ?", idWarehouseProductLocation).Find(&result).Error

	return
}

// GetBatchFromIDWarehouseProductLocation 批量唯一主键查找
func (obj *_WarehouseProductLocationMgr) GetBatchFromIDWarehouseProductLocation(idWarehouseProductLocations []uint32) (results []*WarehouseProductLocation, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_warehouse_product_location IN (?)", idWarehouseProductLocations).Find(&results).Error

	return
}

// GetFromIDProduct 通过id_product获取内容
func (obj *_WarehouseProductLocationMgr) GetFromIDProduct(idProduct uint32) (results []*WarehouseProductLocation, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ?", idProduct).Find(&results).Error

	return
}

// GetBatchFromIDProduct 批量唯一主键查找
func (obj *_WarehouseProductLocationMgr) GetBatchFromIDProduct(idProducts []uint32) (results []*WarehouseProductLocation, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product IN (?)", idProducts).Find(&results).Error

	return
}

// GetFromIDProductAttribute 通过id_product_attribute获取内容
func (obj *_WarehouseProductLocationMgr) GetFromIDProductAttribute(idProductAttribute uint32) (results []*WarehouseProductLocation, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_attribute = ?", idProductAttribute).Find(&results).Error

	return
}

// GetBatchFromIDProductAttribute 批量唯一主键查找
func (obj *_WarehouseProductLocationMgr) GetBatchFromIDProductAttribute(idProductAttributes []uint32) (results []*WarehouseProductLocation, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_attribute IN (?)", idProductAttributes).Find(&results).Error

	return
}

// GetFromIDWarehouse 通过id_warehouse获取内容
func (obj *_WarehouseProductLocationMgr) GetFromIDWarehouse(idWarehouse uint32) (results []*WarehouseProductLocation, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_warehouse = ?", idWarehouse).Find(&results).Error

	return
}

// GetBatchFromIDWarehouse 批量唯一主键查找
func (obj *_WarehouseProductLocationMgr) GetBatchFromIDWarehouse(idWarehouses []uint32) (results []*WarehouseProductLocation, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_warehouse IN (?)", idWarehouses).Find(&results).Error

	return
}

// GetFromLocation 通过location获取内容
func (obj *_WarehouseProductLocationMgr) GetFromLocation(location string) (results []*WarehouseProductLocation, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("location = ?", location).Find(&results).Error

	return
}

// GetBatchFromLocation 批量唯一主键查找
func (obj *_WarehouseProductLocationMgr) GetBatchFromLocation(locations []string) (results []*WarehouseProductLocation, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("location IN (?)", locations).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_WarehouseProductLocationMgr) FetchByPrimaryKey(idWarehouseProductLocation uint32) (result WarehouseProductLocation, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_warehouse_product_location = ?", idWarehouseProductLocation).Find(&result).Error

	return
}

// FetchUniqueIndexByIDProduct primay or index 获取唯一内容
func (obj *_WarehouseProductLocationMgr) FetchUniqueIndexByIDProduct(idProduct uint32, idProductAttribute uint32, idWarehouse uint32) (result WarehouseProductLocation, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ? AND id_product_attribute = ? AND id_warehouse = ?", idProduct, idProductAttribute, idWarehouse).Find(&result).Error

	return
}
