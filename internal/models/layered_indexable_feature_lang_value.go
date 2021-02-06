package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _LayeredIndexableFeatureLangValueMgr struct {
	*_BaseMgr
}

// LayeredIndexableFeatureLangValueMgr open func
func LayeredIndexableFeatureLangValueMgr(db *gorm.DB) *_LayeredIndexableFeatureLangValueMgr {
	if db == nil {
		panic(fmt.Errorf("LayeredIndexableFeatureLangValueMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_LayeredIndexableFeatureLangValueMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_layered_indexable_feature_lang_value"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_LayeredIndexableFeatureLangValueMgr) GetTableName() string {
	return "ps_layered_indexable_feature_lang_value"
}

// Get 获取
func (obj *_LayeredIndexableFeatureLangValueMgr) Get() (result LayeredIndexableFeatureLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_LayeredIndexableFeatureLangValueMgr) Gets() (results []*LayeredIndexableFeatureLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDFeature id_feature获取
func (obj *_LayeredIndexableFeatureLangValueMgr) WithIDFeature(idFeature int) Option {
	return optionFunc(func(o *options) { o.query["id_feature"] = idFeature })
}

// WithIDLang id_lang获取
func (obj *_LayeredIndexableFeatureLangValueMgr) WithIDLang(idLang int) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

// WithURLName url_name获取
func (obj *_LayeredIndexableFeatureLangValueMgr) WithURLName(urlName string) Option {
	return optionFunc(func(o *options) { o.query["url_name"] = urlName })
}

// WithMetaTitle meta_title获取
func (obj *_LayeredIndexableFeatureLangValueMgr) WithMetaTitle(metaTitle string) Option {
	return optionFunc(func(o *options) { o.query["meta_title"] = metaTitle })
}

// GetByOption 功能选项模式获取
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

// GetByOptions 批量功能选项模式获取
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

//////////////////////////enume case ////////////////////////////////////////////

// GetFromIDFeature 通过id_feature获取内容
func (obj *_LayeredIndexableFeatureLangValueMgr) GetFromIDFeature(idFeature int) (results []*LayeredIndexableFeatureLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_feature = ?", idFeature).Find(&results).Error

	return
}

// GetBatchFromIDFeature 批量唯一主键查找
func (obj *_LayeredIndexableFeatureLangValueMgr) GetBatchFromIDFeature(idFeatures []int) (results []*LayeredIndexableFeatureLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_feature IN (?)", idFeatures).Find(&results).Error

	return
}

// GetFromIDLang 通过id_lang获取内容
func (obj *_LayeredIndexableFeatureLangValueMgr) GetFromIDLang(idLang int) (results []*LayeredIndexableFeatureLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error

	return
}

// GetBatchFromIDLang 批量唯一主键查找
func (obj *_LayeredIndexableFeatureLangValueMgr) GetBatchFromIDLang(idLangs []int) (results []*LayeredIndexableFeatureLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error

	return
}

// GetFromURLName 通过url_name获取内容
func (obj *_LayeredIndexableFeatureLangValueMgr) GetFromURLName(urlName string) (results []*LayeredIndexableFeatureLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("url_name = ?", urlName).Find(&results).Error

	return
}

// GetBatchFromURLName 批量唯一主键查找
func (obj *_LayeredIndexableFeatureLangValueMgr) GetBatchFromURLName(urlNames []string) (results []*LayeredIndexableFeatureLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("url_name IN (?)", urlNames).Find(&results).Error

	return
}

// GetFromMetaTitle 通过meta_title获取内容
func (obj *_LayeredIndexableFeatureLangValueMgr) GetFromMetaTitle(metaTitle string) (results []*LayeredIndexableFeatureLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("meta_title = ?", metaTitle).Find(&results).Error

	return
}

// GetBatchFromMetaTitle 批量唯一主键查找
func (obj *_LayeredIndexableFeatureLangValueMgr) GetBatchFromMetaTitle(metaTitles []string) (results []*LayeredIndexableFeatureLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("meta_title IN (?)", metaTitles).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_LayeredIndexableFeatureLangValueMgr) FetchByPrimaryKey(idFeature int, idLang int) (result LayeredIndexableFeatureLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_feature = ? AND id_lang = ?", idFeature, idLang).Find(&result).Error

	return
}
