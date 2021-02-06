package models

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _SmartyLazyCacheMgr struct {
	*_BaseMgr
}

func SmartyLazyCacheMgr(db *gorm.DB) *_SmartyLazyCacheMgr {
	if db == nil {
		panic(fmt.Errorf("SmartyLazyCacheMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SmartyLazyCacheMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_smarty_lazy_cache"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_SmartyLazyCacheMgr) GetTableName() string {
	return "ps_smarty_lazy_cache"
}

func (obj *_SmartyLazyCacheMgr) Get() (result SmartyLazyCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_SmartyLazyCacheMgr) Gets() (results []*SmartyLazyCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_SmartyLazyCacheMgr) WithTemplateHash(templateHash string) Option {
	return optionFunc(func(o *options) { o.query["template_hash"] = templateHash })
}

func (obj *_SmartyLazyCacheMgr) WithCacheID(cacheID string) Option {
	return optionFunc(func(o *options) { o.query["cache_id"] = cacheID })
}

func (obj *_SmartyLazyCacheMgr) WithCompileID(compileID string) Option {
	return optionFunc(func(o *options) { o.query["compile_id"] = compileID })
}

func (obj *_SmartyLazyCacheMgr) WithFilepath(filepath string) Option {
	return optionFunc(func(o *options) { o.query["filepath"] = filepath })
}

func (obj *_SmartyLazyCacheMgr) WithLastUpdate(lastUpdate time.Time) Option {
	return optionFunc(func(o *options) { o.query["last_update"] = lastUpdate })
}

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


func (obj *_SmartyLazyCacheMgr) GetFromTemplateHash(templateHash string) (results []*SmartyLazyCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("template_hash = ?", templateHash).Find(&results).Error

	return
}

func (obj *_SmartyLazyCacheMgr) GetBatchFromTemplateHash(templateHashs []string) (results []*SmartyLazyCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("template_hash IN (?)", templateHashs).Find(&results).Error

	return
}

func (obj *_SmartyLazyCacheMgr) GetFromCacheID(cacheID string) (results []*SmartyLazyCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cache_id = ?", cacheID).Find(&results).Error

	return
}

func (obj *_SmartyLazyCacheMgr) GetBatchFromCacheID(cacheIDs []string) (results []*SmartyLazyCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cache_id IN (?)", cacheIDs).Find(&results).Error

	return
}

func (obj *_SmartyLazyCacheMgr) GetFromCompileID(compileID string) (results []*SmartyLazyCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("compile_id = ?", compileID).Find(&results).Error

	return
}

func (obj *_SmartyLazyCacheMgr) GetBatchFromCompileID(compileIDs []string) (results []*SmartyLazyCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("compile_id IN (?)", compileIDs).Find(&results).Error

	return
}

func (obj *_SmartyLazyCacheMgr) GetFromFilepath(filepath string) (results []*SmartyLazyCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("filepath = ?", filepath).Find(&results).Error

	return
}

func (obj *_SmartyLazyCacheMgr) GetBatchFromFilepath(filepaths []string) (results []*SmartyLazyCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("filepath IN (?)", filepaths).Find(&results).Error

	return
}

func (obj *_SmartyLazyCacheMgr) GetFromLastUpdate(lastUpdate time.Time) (results []*SmartyLazyCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("last_update = ?", lastUpdate).Find(&results).Error

	return
}

func (obj *_SmartyLazyCacheMgr) GetBatchFromLastUpdate(lastUpdates []time.Time) (results []*SmartyLazyCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("last_update IN (?)", lastUpdates).Find(&results).Error

	return
}


func (obj *_SmartyLazyCacheMgr) FetchByPrimaryKey(templateHash string, cacheID string, compileID string) (result SmartyLazyCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("template_hash = ? AND cache_id = ? AND compile_id = ?", templateHash, cacheID, compileID).Find(&result).Error

	return
}
