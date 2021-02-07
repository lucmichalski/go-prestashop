package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _WarehouseProductLocationMgr struct {
	*_BaseMgr
}

func WarehouseProductLocationMgr(db *gorm.DB) *_WarehouseProductLocationMgr {
	if db == nil {
		panic(fmt.Errorf("WarehouseProductLocationMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_WarehouseProductLocationMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_warehouse_product_location"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_WarehouseProductLocationMgr) GetTableName() string {
	return "ps_warehouse_product_location"
}

func (obj *_WarehouseProductLocationMgr) Get() (result WarehouseProductLocation, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_WarehouseProductLocationMgr) Gets() (results []*WarehouseProductLocation, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

func (obj *_WarehouseProductLocationMgr) WithIDWarehouseProductLocation(idWarehouseProductLocation uint32) Option {
	return optionFunc(func(o *options) { o.query["id_warehouse_product_location"] = idWarehouseProductLocation })
}

func (obj *_WarehouseProductLocationMgr) WithIDProduct(idProduct uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product"] = idProduct })
}

func (obj *_WarehouseProductLocationMgr) WithIDProductAttribute(idProductAttribute uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product_attribute"] = idProductAttribute })
}

func (obj *_WarehouseProductLocationMgr) WithIDWarehouse(idWarehouse uint32) Option {
	return optionFunc(func(o *options) { o.query["id_warehouse"] = idWarehouse })
}

func (obj *_WarehouseProductLocationMgr) WithLocation(location string) Option {
	return optionFunc(func(o *options) { o.query["location"] = location })
}

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

func (obj *_WarehouseProductLocationMgr) GetFromIDWarehouseProductLocation(idWarehouseProductLocation uint32) (result WarehouseProductLocation, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_warehouse_product_location = ?", idWarehouseProductLocation).Find(&result).Error

	return
}

func (obj *_WarehouseProductLocationMgr) GetBatchFromIDWarehouseProductLocation(idWarehouseProductLocations []uint32) (results []*WarehouseProductLocation, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_warehouse_product_location IN (?)", idWarehouseProductLocations).Find(&results).Error

	return
}

func (obj *_WarehouseProductLocationMgr) GetFromIDProduct(idProduct uint32) (results []*WarehouseProductLocation, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ?", idProduct).Find(&results).Error

	return
}

func (obj *_WarehouseProductLocationMgr) GetBatchFromIDProduct(idProducts []uint32) (results []*WarehouseProductLocation, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product IN (?)", idProducts).Find(&results).Error

	return
}

func (obj *_WarehouseProductLocationMgr) GetFromIDProductAttribute(idProductAttribute uint32) (results []*WarehouseProductLocation, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_attribute = ?", idProductAttribute).Find(&results).Error

	return
}

func (obj *_WarehouseProductLocationMgr) GetBatchFromIDProductAttribute(idProductAttributes []uint32) (results []*WarehouseProductLocation, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_attribute IN (?)", idProductAttributes).Find(&results).Error

	return
}

func (obj *_WarehouseProductLocationMgr) GetFromIDWarehouse(idWarehouse uint32) (results []*WarehouseProductLocation, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_warehouse = ?", idWarehouse).Find(&results).Error

	return
}

func (obj *_WarehouseProductLocationMgr) GetBatchFromIDWarehouse(idWarehouses []uint32) (results []*WarehouseProductLocation, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_warehouse IN (?)", idWarehouses).Find(&results).Error

	return
}

func (obj *_WarehouseProductLocationMgr) GetFromLocation(location string) (results []*WarehouseProductLocation, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("location = ?", location).Find(&results).Error

	return
}

func (obj *_WarehouseProductLocationMgr) GetBatchFromLocation(locations []string) (results []*WarehouseProductLocation, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("location IN (?)", locations).Find(&results).Error

	return
}

func (obj *_WarehouseProductLocationMgr) FetchByPrimaryKey(idWarehouseProductLocation uint32) (result WarehouseProductLocation, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_warehouse_product_location = ?", idWarehouseProductLocation).Find(&result).Error

	return
}

func (obj *_WarehouseProductLocationMgr) FetchUniqueIndexByIDProduct(idProduct uint32, idProductAttribute uint32, idWarehouse uint32) (result WarehouseProductLocation, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ? AND id_product_attribute = ? AND id_warehouse = ?", idProduct, idProductAttribute, idWarehouse).Find(&result).Error

	return
}
