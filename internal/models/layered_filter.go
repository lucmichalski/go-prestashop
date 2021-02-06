package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _LayeredFilterMgr struct {
	*_BaseMgr
}

// LayeredFilterMgr open func
func LayeredFilterMgr(db *gorm.DB) *_LayeredFilterMgr {
	if db == nil {
		panic(fmt.Errorf("LayeredFilterMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_LayeredFilterMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_layered_filter"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_LayeredFilterMgr) GetTableName() string {
	return "eg_layered_filter"
}

// Get 获取
func (obj *_LayeredFilterMgr) Get() (result LayeredFilter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_LayeredFilterMgr) Gets() (results []*LayeredFilter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDLayeredFilter id_layered_filter获取
func (obj *_LayeredFilterMgr) WithIDLayeredFilter(idLayeredFilter uint32) Option {
	return optionFunc(func(o *options) { o.query["id_layered_filter"] = idLayeredFilter })
}

// WithName name获取
func (obj *_LayeredFilterMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithFilters filters获取
func (obj *_LayeredFilterMgr) WithFilters(filters string) Option {
	return optionFunc(func(o *options) { o.query["filters"] = filters })
}

// WithNCategories n_categories获取
func (obj *_LayeredFilterMgr) WithNCategories(nCategories uint32) Option {
	return optionFunc(func(o *options) { o.query["n_categories"] = nCategories })
}

// WithDateAdd date_add获取
func (obj *_LayeredFilterMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

// GetByOption 功能选项模式获取
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

// GetByOptions 批量功能选项模式获取
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

//////////////////////////enume case ////////////////////////////////////////////

// GetFromIDLayeredFilter 通过id_layered_filter获取内容
func (obj *_LayeredFilterMgr) GetFromIDLayeredFilter(idLayeredFilter uint32) (result LayeredFilter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_layered_filter = ?", idLayeredFilter).Find(&result).Error

	return
}

// GetBatchFromIDLayeredFilter 批量唯一主键查找
func (obj *_LayeredFilterMgr) GetBatchFromIDLayeredFilter(idLayeredFilters []uint32) (results []*LayeredFilter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_layered_filter IN (?)", idLayeredFilters).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_LayeredFilterMgr) GetFromName(name string) (results []*LayeredFilter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量唯一主键查找
func (obj *_LayeredFilterMgr) GetBatchFromName(names []string) (results []*LayeredFilter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

// GetFromFilters 通过filters获取内容
func (obj *_LayeredFilterMgr) GetFromFilters(filters string) (results []*LayeredFilter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("filters = ?", filters).Find(&results).Error

	return
}

// GetBatchFromFilters 批量唯一主键查找
func (obj *_LayeredFilterMgr) GetBatchFromFilters(filterss []string) (results []*LayeredFilter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("filters IN (?)", filterss).Find(&results).Error

	return
}

// GetFromNCategories 通过n_categories获取内容
func (obj *_LayeredFilterMgr) GetFromNCategories(nCategories uint32) (results []*LayeredFilter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("n_categories = ?", nCategories).Find(&results).Error

	return
}

// GetBatchFromNCategories 批量唯一主键查找
func (obj *_LayeredFilterMgr) GetBatchFromNCategories(nCategoriess []uint32) (results []*LayeredFilter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("n_categories IN (?)", nCategoriess).Find(&results).Error

	return
}

// GetFromDateAdd 通过date_add获取内容
func (obj *_LayeredFilterMgr) GetFromDateAdd(dateAdd time.Time) (results []*LayeredFilter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error

	return
}

// GetBatchFromDateAdd 批量唯一主键查找
func (obj *_LayeredFilterMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*LayeredFilter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_LayeredFilterMgr) FetchByPrimaryKey(idLayeredFilter uint32) (result LayeredFilter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_layered_filter = ?", idLayeredFilter).Find(&result).Error

	return
}
