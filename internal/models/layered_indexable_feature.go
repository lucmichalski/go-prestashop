package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _LayeredIndexableFeatureMgr struct {
	*_BaseMgr
}

func LayeredIndexableFeatureMgr(db *gorm.DB) *_LayeredIndexableFeatureMgr {
	if db == nil {
		panic(fmt.Errorf("LayeredIndexableFeatureMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_LayeredIndexableFeatureMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_layered_indexable_feature"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_LayeredIndexableFeatureMgr) GetTableName() string {
	return "ps_layered_indexable_feature"
}

func (obj *_LayeredIndexableFeatureMgr) Get() (result LayeredIndexableFeature, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_LayeredIndexableFeatureMgr) Gets() (results []*LayeredIndexableFeature, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_LayeredIndexableFeatureMgr) WithIDFeature(idFeature int) Option {
	return optionFunc(func(o *options) { o.query["id_feature"] = idFeature })
}

func (obj *_LayeredIndexableFeatureMgr) WithIndexable(indexable bool) Option {
	return optionFunc(func(o *options) { o.query["indexable"] = indexable })
}

func (obj *_LayeredIndexableFeatureMgr) GetByOption(opts ...Option) (result LayeredIndexableFeature, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_LayeredIndexableFeatureMgr) GetByOptions(opts ...Option) (results []*LayeredIndexableFeature, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_LayeredIndexableFeatureMgr) GetFromIDFeature(idFeature int) (result LayeredIndexableFeature, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_feature = ?", idFeature).Find(&result).Error

	return
}

func (obj *_LayeredIndexableFeatureMgr) GetBatchFromIDFeature(idFeatures []int) (results []*LayeredIndexableFeature, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_feature IN (?)", idFeatures).Find(&results).Error

	return
}

func (obj *_LayeredIndexableFeatureMgr) GetFromIndexable(indexable bool) (results []*LayeredIndexableFeature, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("indexable = ?", indexable).Find(&results).Error

	return
}

func (obj *_LayeredIndexableFeatureMgr) GetBatchFromIndexable(indexables []bool) (results []*LayeredIndexableFeature, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("indexable IN (?)", indexables).Find(&results).Error

	return
}


func (obj *_LayeredIndexableFeatureMgr) FetchByPrimaryKey(idFeature int) (result LayeredIndexableFeature, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_feature = ?", idFeature).Find(&result).Error

	return
}
