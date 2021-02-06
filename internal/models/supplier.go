package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _SupplierMgr struct {
	*_BaseMgr
}

// SupplierMgr open func
func SupplierMgr(db *gorm.DB) *_SupplierMgr {
	if db == nil {
		panic(fmt.Errorf("SupplierMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SupplierMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_supplier"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SupplierMgr) GetTableName() string {
	return "ps_supplier"
}

// Get 获取
func (obj *_SupplierMgr) Get() (result Supplier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SupplierMgr) Gets() (results []*Supplier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDSupplier id_supplier获取
func (obj *_SupplierMgr) WithIDSupplier(idSupplier uint32) Option {
	return optionFunc(func(o *options) { o.query["id_supplier"] = idSupplier })
}

// WithName name获取
func (obj *_SupplierMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithDateAdd date_add获取
func (obj *_SupplierMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

// WithDateUpd date_upd获取
func (obj *_SupplierMgr) WithDateUpd(dateUpd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_upd"] = dateUpd })
}

// WithActive active获取
func (obj *_SupplierMgr) WithActive(active bool) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}

// GetByOption 功能选项模式获取
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

// GetByOptions 批量功能选项模式获取
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

//////////////////////////enume case ////////////////////////////////////////////

// GetFromIDSupplier 通过id_supplier获取内容
func (obj *_SupplierMgr) GetFromIDSupplier(idSupplier uint32) (result Supplier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supplier = ?", idSupplier).Find(&result).Error

	return
}

// GetBatchFromIDSupplier 批量唯一主键查找
func (obj *_SupplierMgr) GetBatchFromIDSupplier(idSuppliers []uint32) (results []*Supplier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supplier IN (?)", idSuppliers).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_SupplierMgr) GetFromName(name string) (results []*Supplier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量唯一主键查找
func (obj *_SupplierMgr) GetBatchFromName(names []string) (results []*Supplier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

// GetFromDateAdd 通过date_add获取内容
func (obj *_SupplierMgr) GetFromDateAdd(dateAdd time.Time) (results []*Supplier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error

	return
}

// GetBatchFromDateAdd 批量唯一主键查找
func (obj *_SupplierMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*Supplier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error

	return
}

// GetFromDateUpd 通过date_upd获取内容
func (obj *_SupplierMgr) GetFromDateUpd(dateUpd time.Time) (results []*Supplier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd = ?", dateUpd).Find(&results).Error

	return
}

// GetBatchFromDateUpd 批量唯一主键查找
func (obj *_SupplierMgr) GetBatchFromDateUpd(dateUpds []time.Time) (results []*Supplier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd IN (?)", dateUpds).Find(&results).Error

	return
}

// GetFromActive 通过active获取内容
func (obj *_SupplierMgr) GetFromActive(active bool) (results []*Supplier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active = ?", active).Find(&results).Error

	return
}

// GetBatchFromActive 批量唯一主键查找
func (obj *_SupplierMgr) GetBatchFromActive(actives []bool) (results []*Supplier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active IN (?)", actives).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_SupplierMgr) FetchByPrimaryKey(idSupplier uint32) (result Supplier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supplier = ?", idSupplier).Find(&result).Error

	return
}
