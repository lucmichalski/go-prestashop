package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _LayeredIndexableAttributeGroupLangValueMgr struct {
	*_BaseMgr
}

// LayeredIndexableAttributeGroupLangValueMgr open func
func LayeredIndexableAttributeGroupLangValueMgr(db *gorm.DB) *_LayeredIndexableAttributeGroupLangValueMgr {
	if db == nil {
		panic(fmt.Errorf("LayeredIndexableAttributeGroupLangValueMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_LayeredIndexableAttributeGroupLangValueMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_layered_indexable_attribute_group_lang_value"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_LayeredIndexableAttributeGroupLangValueMgr) GetTableName() string {
	return "ps_layered_indexable_attribute_group_lang_value"
}

// Get 获取
func (obj *_LayeredIndexableAttributeGroupLangValueMgr) Get() (result LayeredIndexableAttributeGroupLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_LayeredIndexableAttributeGroupLangValueMgr) Gets() (results []*LayeredIndexableAttributeGroupLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDAttributeGroup id_attribute_group获取
func (obj *_LayeredIndexableAttributeGroupLangValueMgr) WithIDAttributeGroup(idAttributeGroup int) Option {
	return optionFunc(func(o *options) { o.query["id_attribute_group"] = idAttributeGroup })
}

// WithIDLang id_lang获取
func (obj *_LayeredIndexableAttributeGroupLangValueMgr) WithIDLang(idLang int) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

// WithURLName url_name获取
func (obj *_LayeredIndexableAttributeGroupLangValueMgr) WithURLName(urlName string) Option {
	return optionFunc(func(o *options) { o.query["url_name"] = urlName })
}

// WithMetaTitle meta_title获取
func (obj *_LayeredIndexableAttributeGroupLangValueMgr) WithMetaTitle(metaTitle string) Option {
	return optionFunc(func(o *options) { o.query["meta_title"] = metaTitle })
}

// GetByOption 功能选项模式获取
func (obj *_LayeredIndexableAttributeGroupLangValueMgr) GetByOption(opts ...Option) (result LayeredIndexableAttributeGroupLangValue, err error) {
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
func (obj *_LayeredIndexableAttributeGroupLangValueMgr) GetByOptions(opts ...Option) (results []*LayeredIndexableAttributeGroupLangValue, err error) {
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

// GetFromIDAttributeGroup 通过id_attribute_group获取内容
func (obj *_LayeredIndexableAttributeGroupLangValueMgr) GetFromIDAttributeGroup(idAttributeGroup int) (results []*LayeredIndexableAttributeGroupLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attribute_group = ?", idAttributeGroup).Find(&results).Error

	return
}

// GetBatchFromIDAttributeGroup 批量唯一主键查找
func (obj *_LayeredIndexableAttributeGroupLangValueMgr) GetBatchFromIDAttributeGroup(idAttributeGroups []int) (results []*LayeredIndexableAttributeGroupLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attribute_group IN (?)", idAttributeGroups).Find(&results).Error

	return
}

// GetFromIDLang 通过id_lang获取内容
func (obj *_LayeredIndexableAttributeGroupLangValueMgr) GetFromIDLang(idLang int) (results []*LayeredIndexableAttributeGroupLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error

	return
}

// GetBatchFromIDLang 批量唯一主键查找
func (obj *_LayeredIndexableAttributeGroupLangValueMgr) GetBatchFromIDLang(idLangs []int) (results []*LayeredIndexableAttributeGroupLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error

	return
}

// GetFromURLName 通过url_name获取内容
func (obj *_LayeredIndexableAttributeGroupLangValueMgr) GetFromURLName(urlName string) (results []*LayeredIndexableAttributeGroupLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("url_name = ?", urlName).Find(&results).Error

	return
}

// GetBatchFromURLName 批量唯一主键查找
func (obj *_LayeredIndexableAttributeGroupLangValueMgr) GetBatchFromURLName(urlNames []string) (results []*LayeredIndexableAttributeGroupLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("url_name IN (?)", urlNames).Find(&results).Error

	return
}

// GetFromMetaTitle 通过meta_title获取内容
func (obj *_LayeredIndexableAttributeGroupLangValueMgr) GetFromMetaTitle(metaTitle string) (results []*LayeredIndexableAttributeGroupLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("meta_title = ?", metaTitle).Find(&results).Error

	return
}

// GetBatchFromMetaTitle 批量唯一主键查找
func (obj *_LayeredIndexableAttributeGroupLangValueMgr) GetBatchFromMetaTitle(metaTitles []string) (results []*LayeredIndexableAttributeGroupLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("meta_title IN (?)", metaTitles).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_LayeredIndexableAttributeGroupLangValueMgr) FetchByPrimaryKey(idAttributeGroup int, idLang int) (result LayeredIndexableAttributeGroupLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attribute_group = ? AND id_lang = ?", idAttributeGroup, idLang).Find(&result).Error

	return
}
