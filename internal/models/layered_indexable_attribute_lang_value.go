package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _LayeredIndexableAttributeLangValueMgr struct {
	*_BaseMgr
}

// LayeredIndexableAttributeLangValueMgr open func
func LayeredIndexableAttributeLangValueMgr(db *gorm.DB) *_LayeredIndexableAttributeLangValueMgr {
	if db == nil {
		panic(fmt.Errorf("LayeredIndexableAttributeLangValueMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_LayeredIndexableAttributeLangValueMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_layered_indexable_attribute_lang_value"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_LayeredIndexableAttributeLangValueMgr) GetTableName() string {
	return "ps_layered_indexable_attribute_lang_value"
}

// Get 获取
func (obj *_LayeredIndexableAttributeLangValueMgr) Get() (result LayeredIndexableAttributeLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_LayeredIndexableAttributeLangValueMgr) Gets() (results []*LayeredIndexableAttributeLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDAttribute id_attribute获取
func (obj *_LayeredIndexableAttributeLangValueMgr) WithIDAttribute(idAttribute int) Option {
	return optionFunc(func(o *options) { o.query["id_attribute"] = idAttribute })
}

// WithIDLang id_lang获取
func (obj *_LayeredIndexableAttributeLangValueMgr) WithIDLang(idLang int) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

// WithURLName url_name获取
func (obj *_LayeredIndexableAttributeLangValueMgr) WithURLName(urlName string) Option {
	return optionFunc(func(o *options) { o.query["url_name"] = urlName })
}

// WithMetaTitle meta_title获取
func (obj *_LayeredIndexableAttributeLangValueMgr) WithMetaTitle(metaTitle string) Option {
	return optionFunc(func(o *options) { o.query["meta_title"] = metaTitle })
}

// GetByOption 功能选项模式获取
func (obj *_LayeredIndexableAttributeLangValueMgr) GetByOption(opts ...Option) (result LayeredIndexableAttributeLangValue, err error) {
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
func (obj *_LayeredIndexableAttributeLangValueMgr) GetByOptions(opts ...Option) (results []*LayeredIndexableAttributeLangValue, err error) {
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

// GetFromIDAttribute 通过id_attribute获取内容
func (obj *_LayeredIndexableAttributeLangValueMgr) GetFromIDAttribute(idAttribute int) (results []*LayeredIndexableAttributeLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attribute = ?", idAttribute).Find(&results).Error

	return
}

// GetBatchFromIDAttribute 批量唯一主键查找
func (obj *_LayeredIndexableAttributeLangValueMgr) GetBatchFromIDAttribute(idAttributes []int) (results []*LayeredIndexableAttributeLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attribute IN (?)", idAttributes).Find(&results).Error

	return
}

// GetFromIDLang 通过id_lang获取内容
func (obj *_LayeredIndexableAttributeLangValueMgr) GetFromIDLang(idLang int) (results []*LayeredIndexableAttributeLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error

	return
}

// GetBatchFromIDLang 批量唯一主键查找
func (obj *_LayeredIndexableAttributeLangValueMgr) GetBatchFromIDLang(idLangs []int) (results []*LayeredIndexableAttributeLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error

	return
}

// GetFromURLName 通过url_name获取内容
func (obj *_LayeredIndexableAttributeLangValueMgr) GetFromURLName(urlName string) (results []*LayeredIndexableAttributeLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("url_name = ?", urlName).Find(&results).Error

	return
}

// GetBatchFromURLName 批量唯一主键查找
func (obj *_LayeredIndexableAttributeLangValueMgr) GetBatchFromURLName(urlNames []string) (results []*LayeredIndexableAttributeLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("url_name IN (?)", urlNames).Find(&results).Error

	return
}

// GetFromMetaTitle 通过meta_title获取内容
func (obj *_LayeredIndexableAttributeLangValueMgr) GetFromMetaTitle(metaTitle string) (results []*LayeredIndexableAttributeLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("meta_title = ?", metaTitle).Find(&results).Error

	return
}

// GetBatchFromMetaTitle 批量唯一主键查找
func (obj *_LayeredIndexableAttributeLangValueMgr) GetBatchFromMetaTitle(metaTitles []string) (results []*LayeredIndexableAttributeLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("meta_title IN (?)", metaTitles).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_LayeredIndexableAttributeLangValueMgr) FetchByPrimaryKey(idAttribute int, idLang int) (result LayeredIndexableAttributeLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attribute = ? AND id_lang = ?", idAttribute, idLang).Find(&result).Error

	return
}
