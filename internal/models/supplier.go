package models

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _SupplierMgr struct {
	*_BaseMgr
}

func SupplierMgr(db *gorm.DB) *_SupplierMgr {
	if db == nil {
		panic(fmt.Errorf("SupplierMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SupplierMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_supplier"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_SupplierMgr) GetTableName() string {
	return "ps_supplier"
}

func (obj *_SupplierMgr) Get() (result Supplier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_SupplierMgr) Gets() (results []*Supplier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_SupplierMgr) WithIDSupplier(idSupplier uint32) Option {
	return optionFunc(func(o *options) { o.query["id_supplier"] = idSupplier })
}

func (obj *_SupplierMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

func (obj *_SupplierMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

func (obj *_SupplierMgr) WithDateUpd(dateUpd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_upd"] = dateUpd })
}

func (obj *_SupplierMgr) WithActive(active bool) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}

func (obj *_SupplierMgr) GetByOption(opts ...Option) (result Supplier, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_SupplierMgr) GetByOptions(opts ...Option) (results []*Supplier, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_SupplierMgr) GetFromIDSupplier(idSupplier uint32) (result Supplier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supplier = ?", idSupplier).Find(&result).Error

	return
}

func (obj *_SupplierMgr) GetBatchFromIDSupplier(idSuppliers []uint32) (results []*Supplier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supplier IN (?)", idSuppliers).Find(&results).Error

	return
}

func (obj *_SupplierMgr) GetFromName(name string) (results []*Supplier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

func (obj *_SupplierMgr) GetBatchFromName(names []string) (results []*Supplier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

func (obj *_SupplierMgr) GetFromDateAdd(dateAdd time.Time) (results []*Supplier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error

	return
}

func (obj *_SupplierMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*Supplier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error

	return
}

func (obj *_SupplierMgr) GetFromDateUpd(dateUpd time.Time) (results []*Supplier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd = ?", dateUpd).Find(&results).Error

	return
}

func (obj *_SupplierMgr) GetBatchFromDateUpd(dateUpds []time.Time) (results []*Supplier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd IN (?)", dateUpds).Find(&results).Error

	return
}

func (obj *_SupplierMgr) GetFromActive(active bool) (results []*Supplier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active = ?", active).Find(&results).Error

	return
}

func (obj *_SupplierMgr) GetBatchFromActive(actives []bool) (results []*Supplier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active IN (?)", actives).Find(&results).Error

	return
}


func (obj *_SupplierMgr) FetchByPrimaryKey(idSupplier uint32) (result Supplier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supplier = ?", idSupplier).Find(&result).Error

	return
}
