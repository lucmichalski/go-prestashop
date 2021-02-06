package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _ReferrerCacheMgr struct {
	*_BaseMgr
}

func ReferrerCacheMgr(db *gorm.DB) *_ReferrerCacheMgr {
	if db == nil {
		panic(fmt.Errorf("ReferrerCacheMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ReferrerCacheMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_referrer_cache"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_ReferrerCacheMgr) GetTableName() string {
	return "ps_referrer_cache"
}

func (obj *_ReferrerCacheMgr) Get() (result ReferrerCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_ReferrerCacheMgr) Gets() (results []*ReferrerCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_ReferrerCacheMgr) WithIDConnectionsSource(idConnectionsSource uint32) Option {
	return optionFunc(func(o *options) { o.query["id_connections_source"] = idConnectionsSource })
}

func (obj *_ReferrerCacheMgr) WithIDReferrer(idReferrer uint32) Option {
	return optionFunc(func(o *options) { o.query["id_referrer"] = idReferrer })
}

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


func (obj *_ReferrerCacheMgr) GetFromIDConnectionsSource(idConnectionsSource uint32) (results []*ReferrerCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_connections_source = ?", idConnectionsSource).Find(&results).Error

	return
}

func (obj *_ReferrerCacheMgr) GetBatchFromIDConnectionsSource(idConnectionsSources []uint32) (results []*ReferrerCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_connections_source IN (?)", idConnectionsSources).Find(&results).Error

	return
}

func (obj *_ReferrerCacheMgr) GetFromIDReferrer(idReferrer uint32) (results []*ReferrerCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_referrer = ?", idReferrer).Find(&results).Error

	return
}

func (obj *_ReferrerCacheMgr) GetBatchFromIDReferrer(idReferrers []uint32) (results []*ReferrerCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_referrer IN (?)", idReferrers).Find(&results).Error

	return
}


func (obj *_ReferrerCacheMgr) FetchByPrimaryKey(idConnectionsSource uint32, idReferrer uint32) (result ReferrerCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_connections_source = ? AND id_referrer = ?", idConnectionsSource, idReferrer).Find(&result).Error

	return
}
