package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _SpecificPricePriorityMgr struct {
	*_BaseMgr
}

func SpecificPricePriorityMgr(db *gorm.DB) *_SpecificPricePriorityMgr {
	if db == nil {
		panic(fmt.Errorf("SpecificPricePriorityMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SpecificPricePriorityMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_specific_price_priority"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_SpecificPricePriorityMgr) GetTableName() string {
	return "ps_specific_price_priority"
}

func (obj *_SpecificPricePriorityMgr) Get() (result SpecificPricePriority, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_SpecificPricePriorityMgr) Gets() (results []*SpecificPricePriority, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_SpecificPricePriorityMgr) WithIDSpecificPricePriority(idSpecificPricePriority int) Option {
	return optionFunc(func(o *options) { o.query["id_specific_price_priority"] = idSpecificPricePriority })
}

func (obj *_SpecificPricePriorityMgr) WithIDProduct(idProduct int) Option {
	return optionFunc(func(o *options) { o.query["id_product"] = idProduct })
}

func (obj *_SpecificPricePriorityMgr) WithPriority(priority string) Option {
	return optionFunc(func(o *options) { o.query["priority"] = priority })
}

func (obj *_SpecificPricePriorityMgr) GetByOption(opts ...Option) (result SpecificPricePriority, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_SpecificPricePriorityMgr) GetByOptions(opts ...Option) (results []*SpecificPricePriority, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_SpecificPricePriorityMgr) GetFromIDSpecificPricePriority(idSpecificPricePriority int) (results []*SpecificPricePriority, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_specific_price_priority = ?", idSpecificPricePriority).Find(&results).Error

	return
}

func (obj *_SpecificPricePriorityMgr) GetBatchFromIDSpecificPricePriority(idSpecificPricePrioritys []int) (results []*SpecificPricePriority, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_specific_price_priority IN (?)", idSpecificPricePrioritys).Find(&results).Error

	return
}

func (obj *_SpecificPricePriorityMgr) GetFromIDProduct(idProduct int) (results []*SpecificPricePriority, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ?", idProduct).Find(&results).Error

	return
}

func (obj *_SpecificPricePriorityMgr) GetBatchFromIDProduct(idProducts []int) (results []*SpecificPricePriority, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product IN (?)", idProducts).Find(&results).Error

	return
}

func (obj *_SpecificPricePriorityMgr) GetFromPriority(priority string) (results []*SpecificPricePriority, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("priority = ?", priority).Find(&results).Error

	return
}

func (obj *_SpecificPricePriorityMgr) GetBatchFromPriority(prioritys []string) (results []*SpecificPricePriority, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("priority IN (?)", prioritys).Find(&results).Error

	return
}


func (obj *_SpecificPricePriorityMgr) FetchByPrimaryKey(idSpecificPricePriority int, idProduct int) (result SpecificPricePriority, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_specific_price_priority = ? AND id_product = ?", idSpecificPricePriority, idProduct).Find(&result).Error

	return
}

func (obj *_SpecificPricePriorityMgr) FetchUniqueByIDProduct(idProduct int) (result SpecificPricePriority, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ?", idProduct).Find(&result).Error

	return
}
