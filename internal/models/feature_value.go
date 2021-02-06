package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _FeatureValueMgr struct {
	*_BaseMgr
}

func FeatureValueMgr(db *gorm.DB) *_FeatureValueMgr {
	if db == nil {
		panic(fmt.Errorf("FeatureValueMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_FeatureValueMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_feature_value"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_FeatureValueMgr) GetTableName() string {
	return "ps_feature_value"
}

func (obj *_FeatureValueMgr) Get() (result FeatureValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_FeatureValueMgr) Gets() (results []*FeatureValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_FeatureValueMgr) WithIDFeatureValue(idFeatureValue uint32) Option {
	return optionFunc(func(o *options) { o.query["id_feature_value"] = idFeatureValue })
}

func (obj *_FeatureValueMgr) WithIDFeature(idFeature uint32) Option {
	return optionFunc(func(o *options) { o.query["id_feature"] = idFeature })
}

func (obj *_FeatureValueMgr) WithCustom(custom uint8) Option {
	return optionFunc(func(o *options) { o.query["custom"] = custom })
}

func (obj *_FeatureValueMgr) GetByOption(opts ...Option) (result FeatureValue, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_FeatureValueMgr) GetByOptions(opts ...Option) (results []*FeatureValue, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_FeatureValueMgr) GetFromIDFeatureValue(idFeatureValue uint32) (result FeatureValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_feature_value = ?", idFeatureValue).Find(&result).Error

	return
}

func (obj *_FeatureValueMgr) GetBatchFromIDFeatureValue(idFeatureValues []uint32) (results []*FeatureValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_feature_value IN (?)", idFeatureValues).Find(&results).Error

	return
}

func (obj *_FeatureValueMgr) GetFromIDFeature(idFeature uint32) (results []*FeatureValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_feature = ?", idFeature).Find(&results).Error

	return
}

func (obj *_FeatureValueMgr) GetBatchFromIDFeature(idFeatures []uint32) (results []*FeatureValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_feature IN (?)", idFeatures).Find(&results).Error

	return
}

func (obj *_FeatureValueMgr) GetFromCustom(custom uint8) (results []*FeatureValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("custom = ?", custom).Find(&results).Error

	return
}

func (obj *_FeatureValueMgr) GetBatchFromCustom(customs []uint8) (results []*FeatureValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("custom IN (?)", customs).Find(&results).Error

	return
}


func (obj *_FeatureValueMgr) FetchByPrimaryKey(idFeatureValue uint32) (result FeatureValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_feature_value = ?", idFeatureValue).Find(&result).Error

	return
}

func (obj *_FeatureValueMgr) FetchIndexByFeature(idFeature uint32) (results []*FeatureValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_feature = ?", idFeature).Find(&results).Error

	return
}
