package models

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _SmartyCacheMgr struct {
	*_BaseMgr
}

func SmartyCacheMgr(db *gorm.DB) *_SmartyCacheMgr {
	if db == nil {
		panic(fmt.Errorf("SmartyCacheMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SmartyCacheMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_smarty_cache"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_SmartyCacheMgr) GetTableName() string {
	return "ps_smarty_cache"
}

func (obj *_SmartyCacheMgr) Get() (result SmartyCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_SmartyCacheMgr) Gets() (results []*SmartyCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_SmartyCacheMgr) WithIDSmartyCache(idSmartyCache string) Option {
	return optionFunc(func(o *options) { o.query["id_smarty_cache"] = idSmartyCache })
}

func (obj *_SmartyCacheMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

func (obj *_SmartyCacheMgr) WithCacheID(cacheID string) Option {
	return optionFunc(func(o *options) { o.query["cache_id"] = cacheID })
}

func (obj *_SmartyCacheMgr) WithModified(modified time.Time) Option {
	return optionFunc(func(o *options) { o.query["modified"] = modified })
}

func (obj *_SmartyCacheMgr) WithContent(content string) Option {
	return optionFunc(func(o *options) { o.query["content"] = content })
}

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


func (obj *_SmartyCacheMgr) GetFromIDSmartyCache(idSmartyCache string) (result SmartyCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_smarty_cache = ?", idSmartyCache).Find(&result).Error

	return
}

func (obj *_SmartyCacheMgr) GetBatchFromIDSmartyCache(idSmartyCaches []string) (results []*SmartyCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_smarty_cache IN (?)", idSmartyCaches).Find(&results).Error

	return
}

func (obj *_SmartyCacheMgr) GetFromName(name string) (results []*SmartyCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

func (obj *_SmartyCacheMgr) GetBatchFromName(names []string) (results []*SmartyCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

func (obj *_SmartyCacheMgr) GetFromCacheID(cacheID string) (results []*SmartyCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cache_id = ?", cacheID).Find(&results).Error

	return
}

func (obj *_SmartyCacheMgr) GetBatchFromCacheID(cacheIDs []string) (results []*SmartyCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cache_id IN (?)", cacheIDs).Find(&results).Error

	return
}

func (obj *_SmartyCacheMgr) GetFromModified(modified time.Time) (results []*SmartyCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("modified = ?", modified).Find(&results).Error

	return
}

func (obj *_SmartyCacheMgr) GetBatchFromModified(modifieds []time.Time) (results []*SmartyCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("modified IN (?)", modifieds).Find(&results).Error

	return
}

func (obj *_SmartyCacheMgr) GetFromContent(content string) (results []*SmartyCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("content = ?", content).Find(&results).Error

	return
}

func (obj *_SmartyCacheMgr) GetBatchFromContent(contents []string) (results []*SmartyCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("content IN (?)", contents).Find(&results).Error

	return
}


func (obj *_SmartyCacheMgr) FetchByPrimaryKey(idSmartyCache string) (result SmartyCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_smarty_cache = ?", idSmartyCache).Find(&result).Error

	return
}

func (obj *_SmartyCacheMgr) FetchIndexByName(name string) (results []*SmartyCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

func (obj *_SmartyCacheMgr) FetchIndexByCacheID(cacheID string) (results []*SmartyCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cache_id = ?", cacheID).Find(&results).Error

	return
}

func (obj *_SmartyCacheMgr) FetchIndexByModified(modified time.Time) (results []*SmartyCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("modified = ?", modified).Find(&results).Error

	return
}
