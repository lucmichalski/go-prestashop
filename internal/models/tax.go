package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _TaxMgr struct {
	*_BaseMgr
}

func TaxMgr(db *gorm.DB) *_TaxMgr {
	if db == nil {
		panic(fmt.Errorf("TaxMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_TaxMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_tax"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_TaxMgr) GetTableName() string {
	return "ps_tax"
}

func (obj *_TaxMgr) Get() (result Tax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_TaxMgr) Gets() (results []*Tax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_TaxMgr) WithIDTax(idTax uint32) Option {
	return optionFunc(func(o *options) { o.query["id_tax"] = idTax })
}

func (obj *_TaxMgr) WithRate(rate float64) Option {
	return optionFunc(func(o *options) { o.query["rate"] = rate })
}

func (obj *_TaxMgr) WithActive(active bool) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}

func (obj *_TaxMgr) WithDeleted(deleted bool) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

func (obj *_TaxMgr) GetByOption(opts ...Option) (result Tax, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_TaxMgr) GetByOptions(opts ...Option) (results []*Tax, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_TaxMgr) GetFromIDTax(idTax uint32) (result Tax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tax = ?", idTax).Find(&result).Error

	return
}

func (obj *_TaxMgr) GetBatchFromIDTax(idTaxs []uint32) (results []*Tax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tax IN (?)", idTaxs).Find(&results).Error

	return
}

func (obj *_TaxMgr) GetFromRate(rate float64) (results []*Tax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("rate = ?", rate).Find(&results).Error

	return
}

func (obj *_TaxMgr) GetBatchFromRate(rates []float64) (results []*Tax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("rate IN (?)", rates).Find(&results).Error

	return
}

func (obj *_TaxMgr) GetFromActive(active bool) (results []*Tax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active = ?", active).Find(&results).Error

	return
}

func (obj *_TaxMgr) GetBatchFromActive(actives []bool) (results []*Tax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active IN (?)", actives).Find(&results).Error

	return
}

func (obj *_TaxMgr) GetFromDeleted(deleted bool) (results []*Tax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("deleted = ?", deleted).Find(&results).Error

	return
}

func (obj *_TaxMgr) GetBatchFromDeleted(deleteds []bool) (results []*Tax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("deleted IN (?)", deleteds).Find(&results).Error

	return
}


func (obj *_TaxMgr) FetchByPrimaryKey(idTax uint32) (result Tax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tax = ?", idTax).Find(&result).Error

	return
}
