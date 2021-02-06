package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _ReferrerCacheMgr struct {
	*_BaseMgr
}

// ReferrerCacheMgr open func
func ReferrerCacheMgr(db *gorm.DB) *_ReferrerCacheMgr {
	if db == nil {
		panic(fmt.Errorf("ReferrerCacheMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ReferrerCacheMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_referrer_cache"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ReferrerCacheMgr) GetTableName() string {
	return "eg_referrer_cache"
}

// Get 获取
func (obj *_ReferrerCacheMgr) Get() (result ReferrerCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ReferrerCacheMgr) Gets() (results []*ReferrerCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDConnectionsSource id_connections_source获取
func (obj *_ReferrerCacheMgr) WithIDConnectionsSource(idConnectionsSource uint32) Option {
	return optionFunc(func(o *options) { o.query["id_connections_source"] = idConnectionsSource })
}

// WithIDReferrer id_referrer获取
func (obj *_ReferrerCacheMgr) WithIDReferrer(idReferrer uint32) Option {
	return optionFunc(func(o *options) { o.query["id_referrer"] = idReferrer })
}

// GetByOption 功能选项模式获取
func (obj *_ReferrerCacheMgr) GetByOption(opts ...Option) (result ReferrerCache, err error) {
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
func (obj *_ReferrerCacheMgr) GetByOptions(opts ...Option) (results []*ReferrerCache, err error) {
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

// GetFromIDConnectionsSource 通过id_connections_source获取内容
func (obj *_ReferrerCacheMgr) GetFromIDConnectionsSource(idConnectionsSource uint32) (results []*ReferrerCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_connections_source = ?", idConnectionsSource).Find(&results).Error

	return
}

// GetBatchFromIDConnectionsSource 批量唯一主键查找
func (obj *_ReferrerCacheMgr) GetBatchFromIDConnectionsSource(idConnectionsSources []uint32) (results []*ReferrerCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_connections_source IN (?)", idConnectionsSources).Find(&results).Error

	return
}

// GetFromIDReferrer 通过id_referrer获取内容
func (obj *_ReferrerCacheMgr) GetFromIDReferrer(idReferrer uint32) (results []*ReferrerCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_referrer = ?", idReferrer).Find(&results).Error

	return
}

// GetBatchFromIDReferrer 批量唯一主键查找
func (obj *_ReferrerCacheMgr) GetBatchFromIDReferrer(idReferrers []uint32) (results []*ReferrerCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_referrer IN (?)", idReferrers).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_ReferrerCacheMgr) FetchByPrimaryKey(idConnectionsSource uint32, idReferrer uint32) (result ReferrerCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_connections_source = ? AND id_referrer = ?", idConnectionsSource, idReferrer).Find(&result).Error

	return
}
