package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _LayeredIndexableFeatureLangValueMgr struct {
	*_BaseMgr
}

func LayeredIndexableFeatureLangValueMgr(db *gorm.DB) *_LayeredIndexableFeatureLangValueMgr {
	if db == nil {
		panic(fmt.Errorf("LayeredIndexableFeatureLangValueMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_LayeredIndexableFeatureLangValueMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_layered_indexable_feature_lang_value"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_LayeredIndexableFeatureLangValueMgr) GetTableName() string {
	return "ps_layered_indexable_feature_lang_value"
}

func (obj *_LayeredIndexableFeatureLangValueMgr) Get() (result LayeredIndexableFeatureLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_LayeredIndexableFeatureLangValueMgr) Gets() (results []*LayeredIndexableFeatureLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_LayeredIndexableFeatureLangValueMgr) WithIDFeature(idFeature int) Option {
	return optionFunc(func(o *options) { o.query["id_feature"] = idFeature })
}

func (obj *_LayeredIndexableFeatureLangValueMgr) WithIDLang(idLang int) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

func (obj *_LayeredIndexableFeatureLangValueMgr) WithURLName(urlName string) Option {
	return optionFunc(func(o *options) { o.query["url_name"] = urlName })
}

func (obj *_LayeredIndexableFeatureLangValueMgr) WithMetaTitle(metaTitle string) Option {
	return optionFunc(func(o *options) { o.query["meta_title"] = metaTitle })
}

func (obj *_LayeredIndexableFeatureLangValueMgr) GetByOption(opts ...Option) (result LayeredIndexableFeatureLangValue, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_LayeredIndexableFeatureLangValueMgr) GetByOptions(opts ...Option) (results []*LayeredIndexableFeatureLangValue, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_LayeredIndexableFeatureLangValueMgr) GetFromIDFeature(idFeature int) (results []*LayeredIndexableFeatureLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_feature = ?", idFeature).Find(&results).Error

	return
}

func (obj *_LayeredIndexableFeatureLangValueMgr) GetBatchFromIDFeature(idFeatures []int) (results []*LayeredIndexableFeatureLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_feature IN (?)", idFeatures).Find(&results).Error

	return
}

func (obj *_LayeredIndexableFeatureLangValueMgr) GetFromIDLang(idLang int) (results []*LayeredIndexableFeatureLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error

	return
}

func (obj *_LayeredIndexableFeatureLangValueMgr) GetBatchFromIDLang(idLangs []int) (results []*LayeredIndexableFeatureLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error

	return
}

func (obj *_LayeredIndexableFeatureLangValueMgr) GetFromURLName(urlName string) (results []*LayeredIndexableFeatureLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("url_name = ?", urlName).Find(&results).Error

	return
}

func (obj *_LayeredIndexableFeatureLangValueMgr) GetBatchFromURLName(urlNames []string) (results []*LayeredIndexableFeatureLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("url_name IN (?)", urlNames).Find(&results).Error

	return
}

func (obj *_LayeredIndexableFeatureLangValueMgr) GetFromMetaTitle(metaTitle string) (results []*LayeredIndexableFeatureLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("meta_title = ?", metaTitle).Find(&results).Error

	return
}

func (obj *_LayeredIndexableFeatureLangValueMgr) GetBatchFromMetaTitle(metaTitles []string) (results []*LayeredIndexableFeatureLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("meta_title IN (?)", metaTitles).Find(&results).Error

	return
}


func (obj *_LayeredIndexableFeatureLangValueMgr) FetchByPrimaryKey(idFeature int, idLang int) (result LayeredIndexableFeatureLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_feature = ? AND id_lang = ?", idFeature, idLang).Find(&result).Error

	return
}
