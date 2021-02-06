package models

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _LayeredFilterMgr struct {
	*_BaseMgr
}

func LayeredFilterMgr(db *gorm.DB) *_LayeredFilterMgr {
	if db == nil {
		panic(fmt.Errorf("LayeredFilterMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_LayeredFilterMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_layered_filter"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_LayeredFilterMgr) GetTableName() string {
	return "ps_layered_filter"
}

func (obj *_LayeredFilterMgr) Get() (result LayeredFilter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_LayeredFilterMgr) Gets() (results []*LayeredFilter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_LayeredFilterMgr) WithIDLayeredFilter(idLayeredFilter uint32) Option {
	return optionFunc(func(o *options) { o.query["id_layered_filter"] = idLayeredFilter })
}

func (obj *_LayeredFilterMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

func (obj *_LayeredFilterMgr) WithFilters(filters string) Option {
	return optionFunc(func(o *options) { o.query["filters"] = filters })
}

func (obj *_LayeredFilterMgr) WithNCategories(nCategories uint32) Option {
	return optionFunc(func(o *options) { o.query["n_categories"] = nCategories })
}

func (obj *_LayeredFilterMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

func (obj *_LayeredFilterMgr) GetByOption(opts ...Option) (result LayeredFilter, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_LayeredFilterMgr) GetByOptions(opts ...Option) (results []*LayeredFilter, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_LayeredFilterMgr) GetFromIDLayeredFilter(idLayeredFilter uint32) (result LayeredFilter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_layered_filter = ?", idLayeredFilter).Find(&result).Error

	return
}

func (obj *_LayeredFilterMgr) GetBatchFromIDLayeredFilter(idLayeredFilters []uint32) (results []*LayeredFilter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_layered_filter IN (?)", idLayeredFilters).Find(&results).Error

	return
}

func (obj *_LayeredFilterMgr) GetFromName(name string) (results []*LayeredFilter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

func (obj *_LayeredFilterMgr) GetBatchFromName(names []string) (results []*LayeredFilter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

func (obj *_LayeredFilterMgr) GetFromFilters(filters string) (results []*LayeredFilter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("filters = ?", filters).Find(&results).Error

	return
}

func (obj *_LayeredFilterMgr) GetBatchFromFilters(filterss []string) (results []*LayeredFilter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("filters IN (?)", filterss).Find(&results).Error

	return
}

func (obj *_LayeredFilterMgr) GetFromNCategories(nCategories uint32) (results []*LayeredFilter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("n_categories = ?", nCategories).Find(&results).Error

	return
}

func (obj *_LayeredFilterMgr) GetBatchFromNCategories(nCategoriess []uint32) (results []*LayeredFilter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("n_categories IN (?)", nCategoriess).Find(&results).Error

	return
}

func (obj *_LayeredFilterMgr) GetFromDateAdd(dateAdd time.Time) (results []*LayeredFilter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error

	return
}

func (obj *_LayeredFilterMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*LayeredFilter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error

	return
}


func (obj *_LayeredFilterMgr) FetchByPrimaryKey(idLayeredFilter uint32) (result LayeredFilter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_layered_filter = ?", idLayeredFilter).Find(&result).Error

	return
}
