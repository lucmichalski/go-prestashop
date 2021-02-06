package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _WarehouseCarrierMgr struct {
	*_BaseMgr
}

func WarehouseCarrierMgr(db *gorm.DB) *_WarehouseCarrierMgr {
	if db == nil {
		panic(fmt.Errorf("WarehouseCarrierMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_WarehouseCarrierMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_warehouse_carrier"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_WarehouseCarrierMgr) GetTableName() string {
	return "ps_warehouse_carrier"
}

func (obj *_WarehouseCarrierMgr) Get() (result WarehouseCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_WarehouseCarrierMgr) Gets() (results []*WarehouseCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_WarehouseCarrierMgr) WithIDCarrier(idCarrier uint32) Option {
	return optionFunc(func(o *options) { o.query["id_carrier"] = idCarrier })
}

func (obj *_WarehouseCarrierMgr) WithIDWarehouse(idWarehouse uint32) Option {
	return optionFunc(func(o *options) { o.query["id_warehouse"] = idWarehouse })
}

func (obj *_WarehouseCarrierMgr) GetByOption(opts ...Option) (result WarehouseCarrier, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_WarehouseCarrierMgr) GetByOptions(opts ...Option) (results []*WarehouseCarrier, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_WarehouseCarrierMgr) GetFromIDCarrier(idCarrier uint32) (results []*WarehouseCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_carrier = ?", idCarrier).Find(&results).Error

	return
}

func (obj *_WarehouseCarrierMgr) GetBatchFromIDCarrier(idCarriers []uint32) (results []*WarehouseCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_carrier IN (?)", idCarriers).Find(&results).Error

	return
}

func (obj *_WarehouseCarrierMgr) GetFromIDWarehouse(idWarehouse uint32) (results []*WarehouseCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_warehouse = ?", idWarehouse).Find(&results).Error

	return
}

func (obj *_WarehouseCarrierMgr) GetBatchFromIDWarehouse(idWarehouses []uint32) (results []*WarehouseCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_warehouse IN (?)", idWarehouses).Find(&results).Error

	return
}


func (obj *_WarehouseCarrierMgr) FetchByPrimaryKey(idCarrier uint32, idWarehouse uint32) (result WarehouseCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_carrier = ? AND id_warehouse = ?", idCarrier, idWarehouse).Find(&result).Error

	return
}

func (obj *_WarehouseCarrierMgr) FetchIndexByIDCarrier(idCarrier uint32) (results []*WarehouseCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_carrier = ?", idCarrier).Find(&results).Error

	return
}

func (obj *_WarehouseCarrierMgr) FetchIndexByIDWarehouse(idWarehouse uint32) (results []*WarehouseCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_warehouse = ?", idWarehouse).Find(&results).Error

	return
}
