package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _LayeredIndexableFeatureValueLangValueMgr struct {
	*_BaseMgr
}

// LayeredIndexableFeatureValueLangValueMgr open func
func LayeredIndexableFeatureValueLangValueMgr(db *gorm.DB) *_LayeredIndexableFeatureValueLangValueMgr {
	if db == nil {
		panic(fmt.Errorf("LayeredIndexableFeatureValueLangValueMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_LayeredIndexableFeatureValueLangValueMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_layered_indexable_feature_value_lang_value"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_LayeredIndexableFeatureValueLangValueMgr) GetTableName() string {
	return "ps_layered_indexable_feature_value_lang_value"
}

// Get 获取
func (obj *_LayeredIndexableFeatureValueLangValueMgr) Get() (result LayeredIndexableFeatureValueLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_LayeredIndexableFeatureValueLangValueMgr) Gets() (results []*LayeredIndexableFeatureValueLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDFeatureValue id_feature_value获取
func (obj *_LayeredIndexableFeatureValueLangValueMgr) WithIDFeatureValue(idFeatureValue int) Option {
	return optionFunc(func(o *options) { o.query["id_feature_value"] = idFeatureValue })
}

// WithIDLang id_lang获取
func (obj *_LayeredIndexableFeatureValueLangValueMgr) WithIDLang(idLang int) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

// WithURLName url_name获取
func (obj *_LayeredIndexableFeatureValueLangValueMgr) WithURLName(urlName string) Option {
	return optionFunc(func(o *options) { o.query["url_name"] = urlName })
}

// WithMetaTitle meta_title获取
func (obj *_LayeredIndexableFeatureValueLangValueMgr) WithMetaTitle(metaTitle string) Option {
	return optionFunc(func(o *options) { o.query["meta_title"] = metaTitle })
}

// GetByOption 功能选项模式获取
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

// GetByOptions 批量功能选项模式获取
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

//////////////////////////enume case ////////////////////////////////////////////

// GetFromIDFeatureValue 通过id_feature_value获取内容
func (obj *_LayeredIndexableFeatureValueLangValueMgr) GetFromIDFeatureValue(idFeatureValue int) (results []*LayeredIndexableFeatureValueLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_feature_value = ?", idFeatureValue).Find(&results).Error

	return
}

// GetBatchFromIDFeatureValue 批量唯一主键查找
func (obj *_LayeredIndexableFeatureValueLangValueMgr) GetBatchFromIDFeatureValue(idFeatureValues []int) (results []*LayeredIndexableFeatureValueLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_feature_value IN (?)", idFeatureValues).Find(&results).Error

	return
}

// GetFromIDLang 通过id_lang获取内容
func (obj *_LayeredIndexableFeatureValueLangValueMgr) GetFromIDLang(idLang int) (results []*LayeredIndexableFeatureValueLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error

	return
}

// GetBatchFromIDLang 批量唯一主键查找
func (obj *_LayeredIndexableFeatureValueLangValueMgr) GetBatchFromIDLang(idLangs []int) (results []*LayeredIndexableFeatureValueLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error

	return
}

// GetFromURLName 通过url_name获取内容
func (obj *_LayeredIndexableFeatureValueLangValueMgr) GetFromURLName(urlName string) (results []*LayeredIndexableFeatureValueLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("url_name = ?", urlName).Find(&results).Error

	return
}

// GetBatchFromURLName 批量唯一主键查找
func (obj *_LayeredIndexableFeatureValueLangValueMgr) GetBatchFromURLName(urlNames []string) (results []*LayeredIndexableFeatureValueLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("url_name IN (?)", urlNames).Find(&results).Error

	return
}

// GetFromMetaTitle 通过meta_title获取内容
func (obj *_LayeredIndexableFeatureValueLangValueMgr) GetFromMetaTitle(metaTitle string) (results []*LayeredIndexableFeatureValueLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("meta_title = ?", metaTitle).Find(&results).Error

	return
}

// GetBatchFromMetaTitle 批量唯一主键查找
func (obj *_LayeredIndexableFeatureValueLangValueMgr) GetBatchFromMetaTitle(metaTitles []string) (results []*LayeredIndexableFeatureValueLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("meta_title IN (?)", metaTitles).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_LayeredIndexableFeatureValueLangValueMgr) FetchByPrimaryKey(idFeatureValue int, idLang int) (result LayeredIndexableFeatureValueLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_feature_value = ? AND id_lang = ?", idFeatureValue, idLang).Find(&result).Error

	return
}
