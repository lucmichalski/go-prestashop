package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _WarehouseMgr struct {
	*_BaseMgr
}

func WarehouseMgr(db *gorm.DB) *_WarehouseMgr {
	if db == nil {
		panic(fmt.Errorf("WarehouseMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_WarehouseMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_warehouse"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_WarehouseMgr) GetTableName() string {
	return "ps_warehouse"
}

func (obj *_WarehouseMgr) Get() (result Warehouse, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_WarehouseMgr) Gets() (results []*Warehouse, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_WarehouseMgr) WithIDWarehouse(idWarehouse uint32) Option {
	return optionFunc(func(o *options) { o.query["id_warehouse"] = idWarehouse })
}

func (obj *_WarehouseMgr) WithIDCurrency(idCurrency uint32) Option {
	return optionFunc(func(o *options) { o.query["id_currency"] = idCurrency })
}

func (obj *_WarehouseMgr) WithIDAddress(idAddress uint32) Option {
	return optionFunc(func(o *options) { o.query["id_address"] = idAddress })
}

func (obj *_WarehouseMgr) WithIDEmployee(idEmployee uint32) Option {
	return optionFunc(func(o *options) { o.query["id_employee"] = idEmployee })
}

func (obj *_WarehouseMgr) WithReference(reference string) Option {
	return optionFunc(func(o *options) { o.query["reference"] = reference })
}

func (obj *_WarehouseMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

func (obj *_WarehouseMgr) WithManagementType(managementType string) Option {
	return optionFunc(func(o *options) { o.query["management_type"] = managementType })
}

func (obj *_WarehouseMgr) WithDeleted(deleted bool) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

func (obj *_WarehouseMgr) GetByOption(opts ...Option) (result Warehouse, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_WarehouseMgr) GetByOptions(opts ...Option) (results []*Warehouse, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_WarehouseMgr) GetFromIDWarehouse(idWarehouse uint32) (result Warehouse, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_warehouse = ?", idWarehouse).Find(&result).Error

	return
}

func (obj *_WarehouseMgr) GetBatchFromIDWarehouse(idWarehouses []uint32) (results []*Warehouse, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_warehouse IN (?)", idWarehouses).Find(&results).Error

	return
}

func (obj *_WarehouseMgr) GetFromIDCurrency(idCurrency uint32) (results []*Warehouse, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_currency = ?", idCurrency).Find(&results).Error

	return
}

func (obj *_WarehouseMgr) GetBatchFromIDCurrency(idCurrencys []uint32) (results []*Warehouse, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_currency IN (?)", idCurrencys).Find(&results).Error

	return
}

func (obj *_WarehouseMgr) GetFromIDAddress(idAddress uint32) (results []*Warehouse, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_address = ?", idAddress).Find(&results).Error

	return
}

func (obj *_WarehouseMgr) GetBatchFromIDAddress(idAddresss []uint32) (results []*Warehouse, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_address IN (?)", idAddresss).Find(&results).Error

	return
}

func (obj *_WarehouseMgr) GetFromIDEmployee(idEmployee uint32) (results []*Warehouse, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee = ?", idEmployee).Find(&results).Error

	return
}

func (obj *_WarehouseMgr) GetBatchFromIDEmployee(idEmployees []uint32) (results []*Warehouse, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee IN (?)", idEmployees).Find(&results).Error

	return
}

func (obj *_WarehouseMgr) GetFromReference(reference string) (results []*Warehouse, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reference = ?", reference).Find(&results).Error

	return
}

func (obj *_WarehouseMgr) GetBatchFromReference(references []string) (results []*Warehouse, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reference IN (?)", references).Find(&results).Error

	return
}

func (obj *_WarehouseMgr) GetFromName(name string) (results []*Warehouse, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

func (obj *_WarehouseMgr) GetBatchFromName(names []string) (results []*Warehouse, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

func (obj *_WarehouseMgr) GetFromManagementType(managementType string) (results []*Warehouse, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("management_type = ?", managementType).Find(&results).Error

	return
}

func (obj *_WarehouseMgr) GetBatchFromManagementType(managementTypes []string) (results []*Warehouse, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("management_type IN (?)", managementTypes).Find(&results).Error

	return
}

func (obj *_WarehouseMgr) GetFromDeleted(deleted bool) (results []*Warehouse, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("deleted = ?", deleted).Find(&results).Error

	return
}

func (obj *_WarehouseMgr) GetBatchFromDeleted(deleteds []bool) (results []*Warehouse, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("deleted IN (?)", deleteds).Find(&results).Error

	return
}


func (obj *_WarehouseMgr) FetchByPrimaryKey(idWarehouse uint32) (result Warehouse, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_warehouse = ?", idWarehouse).Find(&result).Error

	return
}
