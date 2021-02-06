package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _ProductGroupReductionCacheMgr struct {
	*_BaseMgr
}

func ProductGroupReductionCacheMgr(db *gorm.DB) *_ProductGroupReductionCacheMgr {
	if db == nil {
		panic(fmt.Errorf("ProductGroupReductionCacheMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ProductGroupReductionCacheMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_product_group_reduction_cache"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_ProductGroupReductionCacheMgr) GetTableName() string {
	return "ps_product_group_reduction_cache"
}

func (obj *_ProductGroupReductionCacheMgr) Get() (result ProductGroupReductionCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_ProductGroupReductionCacheMgr) Gets() (results []*ProductGroupReductionCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_ProductGroupReductionCacheMgr) WithIDProduct(idProduct uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product"] = idProduct })
}

func (obj *_ProductGroupReductionCacheMgr) WithIDGroup(idGroup uint32) Option {
	return optionFunc(func(o *options) { o.query["id_group"] = idGroup })
}

func (obj *_ProductGroupReductionCacheMgr) WithReduction(reduction float64) Option {
	return optionFunc(func(o *options) { o.query["reduction"] = reduction })
}

func (obj *_ProductGroupReductionCacheMgr) GetByOption(opts ...Option) (result ProductGroupReductionCache, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_ProductGroupReductionCacheMgr) GetByOptions(opts ...Option) (results []*ProductGroupReductionCache, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_ProductGroupReductionCacheMgr) GetFromIDProduct(idProduct uint32) (results []*ProductGroupReductionCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ?", idProduct).Find(&results).Error

	return
}

func (obj *_ProductGroupReductionCacheMgr) GetBatchFromIDProduct(idProducts []uint32) (results []*ProductGroupReductionCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product IN (?)", idProducts).Find(&results).Error

	return
}

func (obj *_ProductGroupReductionCacheMgr) GetFromIDGroup(idGroup uint32) (results []*ProductGroupReductionCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_group = ?", idGroup).Find(&results).Error

	return
}

func (obj *_ProductGroupReductionCacheMgr) GetBatchFromIDGroup(idGroups []uint32) (results []*ProductGroupReductionCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_group IN (?)", idGroups).Find(&results).Error

	return
}

func (obj *_ProductGroupReductionCacheMgr) GetFromReduction(reduction float64) (results []*ProductGroupReductionCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reduction = ?", reduction).Find(&results).Error

	return
}

func (obj *_ProductGroupReductionCacheMgr) GetBatchFromReduction(reductions []float64) (results []*ProductGroupReductionCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reduction IN (?)", reductions).Find(&results).Error

	return
}


func (obj *_ProductGroupReductionCacheMgr) FetchByPrimaryKey(idProduct uint32, idGroup uint32) (result ProductGroupReductionCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ? AND id_group = ?", idProduct, idGroup).Find(&result).Error

	return
}
