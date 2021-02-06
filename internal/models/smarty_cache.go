package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _SmartyCacheMgr struct {
	*_BaseMgr
}

// SmartyCacheMgr open func
func SmartyCacheMgr(db *gorm.DB) *_SmartyCacheMgr {
	if db == nil {
		panic(fmt.Errorf("SmartyCacheMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SmartyCacheMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_smarty_cache"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SmartyCacheMgr) GetTableName() string {
	return "eg_smarty_cache"
}

// Get 获取
func (obj *_SmartyCacheMgr) Get() (result SmartyCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SmartyCacheMgr) Gets() (results []*SmartyCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDSmartyCache id_smarty_cache获取
func (obj *_SmartyCacheMgr) WithIDSmartyCache(idSmartyCache string) Option {
	return optionFunc(func(o *options) { o.query["id_smarty_cache"] = idSmartyCache })
}

// WithName name获取
func (obj *_SmartyCacheMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithCacheID cache_id获取
func (obj *_SmartyCacheMgr) WithCacheID(cacheID string) Option {
	return optionFunc(func(o *options) { o.query["cache_id"] = cacheID })
}

// WithModified modified获取
func (obj *_SmartyCacheMgr) WithModified(modified time.Time) Option {
	return optionFunc(func(o *options) { o.query["modified"] = modified })
}

// WithContent content获取
func (obj *_SmartyCacheMgr) WithContent(content string) Option {
	return optionFunc(func(o *options) { o.query["content"] = content })
}

// GetByOption 功能选项模式获取
func (obj *_SmartyCacheMgr) GetByOption(opts ...Option) (result SmartyCache, err error) {
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
func (obj *_SmartyCacheMgr) GetByOptions(opts ...Option) (results []*SmartyCache, err error) {
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

// GetFromIDSmartyCache 通过id_smarty_cache获取内容
func (obj *_SmartyCacheMgr) GetFromIDSmartyCache(idSmartyCache string) (result SmartyCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_smarty_cache = ?", idSmartyCache).Find(&result).Error

	return
}

// GetBatchFromIDSmartyCache 批量唯一主键查找
func (obj *_SmartyCacheMgr) GetBatchFromIDSmartyCache(idSmartyCaches []string) (results []*SmartyCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_smarty_cache IN (?)", idSmartyCaches).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_SmartyCacheMgr) GetFromName(name string) (results []*SmartyCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量唯一主键查找
func (obj *_SmartyCacheMgr) GetBatchFromName(names []string) (results []*SmartyCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

// GetFromCacheID 通过cache_id获取内容
func (obj *_SmartyCacheMgr) GetFromCacheID(cacheID string) (results []*SmartyCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cache_id = ?", cacheID).Find(&results).Error

	return
}

// GetBatchFromCacheID 批量唯一主键查找
func (obj *_SmartyCacheMgr) GetBatchFromCacheID(cacheIDs []string) (results []*SmartyCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cache_id IN (?)", cacheIDs).Find(&results).Error

	return
}

// GetFromModified 通过modified获取内容
func (obj *_SmartyCacheMgr) GetFromModified(modified time.Time) (results []*SmartyCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("modified = ?", modified).Find(&results).Error

	return
}

// GetBatchFromModified 批量唯一主键查找
func (obj *_SmartyCacheMgr) GetBatchFromModified(modifieds []time.Time) (results []*SmartyCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("modified IN (?)", modifieds).Find(&results).Error

	return
}

// GetFromContent 通过content获取内容
func (obj *_SmartyCacheMgr) GetFromContent(content string) (results []*SmartyCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("content = ?", content).Find(&results).Error

	return
}

// GetBatchFromContent 批量唯一主键查找
func (obj *_SmartyCacheMgr) GetBatchFromContent(contents []string) (results []*SmartyCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("content IN (?)", contents).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_SmartyCacheMgr) FetchByPrimaryKey(idSmartyCache string) (result SmartyCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_smarty_cache = ?", idSmartyCache).Find(&result).Error

	return
}

// FetchIndexByName  获取多个内容
func (obj *_SmartyCacheMgr) FetchIndexByName(name string) (results []*SmartyCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

// FetchIndexByCacheID  获取多个内容
func (obj *_SmartyCacheMgr) FetchIndexByCacheID(cacheID string) (results []*SmartyCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cache_id = ?", cacheID).Find(&results).Error

	return
}

// FetchIndexByModified  获取多个内容
func (obj *_SmartyCacheMgr) FetchIndexByModified(modified time.Time) (results []*SmartyCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("modified = ?", modified).Find(&results).Error

	return
}
