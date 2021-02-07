package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _LayeredIndexableFeatureValueLangValueMgr struct {
	*_BaseMgr
}

func LayeredIndexableFeatureValueLangValueMgr(db *gorm.DB) *_LayeredIndexableFeatureValueLangValueMgr {
	if db == nil {
		panic(fmt.Errorf("LayeredIndexableFeatureValueLangValueMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_LayeredIndexableFeatureValueLangValueMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_layered_indexable_feature_value_lang_value"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_LayeredIndexableFeatureValueLangValueMgr) GetTableName() string {
	return "ps_layered_indexable_feature_value_lang_value"
}

func (obj *_LayeredIndexableFeatureValueLangValueMgr) Get() (result LayeredIndexableFeatureValueLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_LayeredIndexableFeatureValueLangValueMgr) Gets() (results []*LayeredIndexableFeatureValueLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

func (obj *_LayeredIndexableFeatureValueLangValueMgr) WithIDFeatureValue(idFeatureValue int) Option {
	return optionFunc(func(o *options) { o.query["id_feature_value"] = idFeatureValue })
}

func (obj *_LayeredIndexableFeatureValueLangValueMgr) WithIDLang(idLang int) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

func (obj *_LayeredIndexableFeatureValueLangValueMgr) WithURLName(urlName string) Option {
	return optionFunc(func(o *options) { o.query["url_name"] = urlName })
}

func (obj *_LayeredIndexableFeatureValueLangValueMgr) WithMetaTitle(metaTitle string) Option {
	return optionFunc(func(o *options) { o.query["meta_title"] = metaTitle })
}

func (obj *_LayeredIndexableFeatureValueLangValueMgr) GetByOption(opts ...Option) (result LayeredIndexableFeatureValueLangValue, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_LayeredIndexableFeatureValueLangValueMgr) GetByOptions(opts ...Option) (results []*LayeredIndexableFeatureValueLangValue, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

func (obj *_LayeredIndexableFeatureValueLangValueMgr) GetFromIDFeatureValue(idFeatureValue int) (results []*LayeredIndexableFeatureValueLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_feature_value = ?", idFeatureValue).Find(&results).Error

	return
}

func (obj *_LayeredIndexableFeatureValueLangValueMgr) GetBatchFromIDFeatureValue(idFeatureValues []int) (results []*LayeredIndexableFeatureValueLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_feature_value IN (?)", idFeatureValues).Find(&results).Error

	return
}

func (obj *_LayeredIndexableFeatureValueLangValueMgr) GetFromIDLang(idLang int) (results []*LayeredIndexableFeatureValueLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error

	return
}

func (obj *_LayeredIndexableFeatureValueLangValueMgr) GetBatchFromIDLang(idLangs []int) (results []*LayeredIndexableFeatureValueLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error

	return
}

func (obj *_LayeredIndexableFeatureValueLangValueMgr) GetFromURLName(urlName string) (results []*LayeredIndexableFeatureValueLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("url_name = ?", urlName).Find(&results).Error

	return
}

func (obj *_LayeredIndexableFeatureValueLangValueMgr) GetBatchFromURLName(urlNames []string) (results []*LayeredIndexableFeatureValueLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("url_name IN (?)", urlNames).Find(&results).Error

	return
}

func (obj *_LayeredIndexableFeatureValueLangValueMgr) GetFromMetaTitle(metaTitle string) (results []*LayeredIndexableFeatureValueLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("meta_title = ?", metaTitle).Find(&results).Error

	return
}

func (obj *_LayeredIndexableFeatureValueLangValueMgr) GetBatchFromMetaTitle(metaTitles []string) (results []*LayeredIndexableFeatureValueLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("meta_title IN (?)", metaTitles).Find(&results).Error

	return
}

func (obj *_LayeredIndexableFeatureValueLangValueMgr) FetchByPrimaryKey(idFeatureValue int, idLang int) (result LayeredIndexableFeatureValueLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_feature_value = ? AND id_lang = ?", idFeatureValue, idLang).Find(&result).Error

	return
}
