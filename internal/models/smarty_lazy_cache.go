package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _SmartyLazyCacheMgr struct {
	*_BaseMgr
}

// SmartyLazyCacheMgr open func
func SmartyLazyCacheMgr(db *gorm.DB) *_SmartyLazyCacheMgr {
	if db == nil {
		panic(fmt.Errorf("SmartyLazyCacheMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SmartyLazyCacheMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_smarty_lazy_cache"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SmartyLazyCacheMgr) GetTableName() string {
	return "eg_smarty_lazy_cache"
}

// Get 获取
func (obj *_SmartyLazyCacheMgr) Get() (result SmartyLazyCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SmartyLazyCacheMgr) Gets() (results []*SmartyLazyCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithTemplateHash template_hash获取
func (obj *_SmartyLazyCacheMgr) WithTemplateHash(templateHash string) Option {
	return optionFunc(func(o *options) { o.query["template_hash"] = templateHash })
}

// WithCacheID cache_id获取
func (obj *_SmartyLazyCacheMgr) WithCacheID(cacheID string) Option {
	return optionFunc(func(o *options) { o.query["cache_id"] = cacheID })
}

// WithCompileID compile_id获取
func (obj *_SmartyLazyCacheMgr) WithCompileID(compileID string) Option {
	return optionFunc(func(o *options) { o.query["compile_id"] = compileID })
}

// WithFilepath filepath获取
func (obj *_SmartyLazyCacheMgr) WithFilepath(filepath string) Option {
	return optionFunc(func(o *options) { o.query["filepath"] = filepath })
}

// WithLastUpdate last_update获取
func (obj *_SmartyLazyCacheMgr) WithLastUpdate(lastUpdate time.Time) Option {
	return optionFunc(func(o *options) { o.query["last_update"] = lastUpdate })
}

// GetByOption 功能选项模式获取
func (obj *_SmartyLazyCacheMgr) GetByOption(opts ...Option) (result SmartyLazyCache, err error) {
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
func (obj *_SmartyLazyCacheMgr) GetByOptions(opts ...Option) (results []*SmartyLazyCache, err error) {
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

// GetFromTemplateHash 通过template_hash获取内容
func (obj *_SmartyLazyCacheMgr) GetFromTemplateHash(templateHash string) (results []*SmartyLazyCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("template_hash = ?", templateHash).Find(&results).Error

	return
}

// GetBatchFromTemplateHash 批量唯一主键查找
func (obj *_SmartyLazyCacheMgr) GetBatchFromTemplateHash(templateHashs []string) (results []*SmartyLazyCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("template_hash IN (?)", templateHashs).Find(&results).Error

	return
}

// GetFromCacheID 通过cache_id获取内容
func (obj *_SmartyLazyCacheMgr) GetFromCacheID(cacheID string) (results []*SmartyLazyCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cache_id = ?", cacheID).Find(&results).Error

	return
}

// GetBatchFromCacheID 批量唯一主键查找
func (obj *_SmartyLazyCacheMgr) GetBatchFromCacheID(cacheIDs []string) (results []*SmartyLazyCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cache_id IN (?)", cacheIDs).Find(&results).Error

	return
}

// GetFromCompileID 通过compile_id获取内容
func (obj *_SmartyLazyCacheMgr) GetFromCompileID(compileID string) (results []*SmartyLazyCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("compile_id = ?", compileID).Find(&results).Error

	return
}

// GetBatchFromCompileID 批量唯一主键查找
func (obj *_SmartyLazyCacheMgr) GetBatchFromCompileID(compileIDs []string) (results []*SmartyLazyCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("compile_id IN (?)", compileIDs).Find(&results).Error

	return
}

// GetFromFilepath 通过filepath获取内容
func (obj *_SmartyLazyCacheMgr) GetFromFilepath(filepath string) (results []*SmartyLazyCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("filepath = ?", filepath).Find(&results).Error

	return
}

// GetBatchFromFilepath 批量唯一主键查找
func (obj *_SmartyLazyCacheMgr) GetBatchFromFilepath(filepaths []string) (results []*SmartyLazyCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("filepath IN (?)", filepaths).Find(&results).Error

	return
}

// GetFromLastUpdate 通过last_update获取内容
func (obj *_SmartyLazyCacheMgr) GetFromLastUpdate(lastUpdate time.Time) (results []*SmartyLazyCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("last_update = ?", lastUpdate).Find(&results).Error

	return
}

// GetBatchFromLastUpdate 批量唯一主键查找
func (obj *_SmartyLazyCacheMgr) GetBatchFromLastUpdate(lastUpdates []time.Time) (results []*SmartyLazyCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("last_update IN (?)", lastUpdates).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_SmartyLazyCacheMgr) FetchByPrimaryKey(templateHash string, cacheID string, compileID string) (result SmartyLazyCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("template_hash = ? AND cache_id = ? AND compile_id = ?", templateHash, cacheID, compileID).Find(&result).Error

	return
}
