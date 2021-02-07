package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _LayeredIndexableAttributeGroupLangValueMgr struct {
	*_BaseMgr
}

func LayeredIndexableAttributeGroupLangValueMgr(db *gorm.DB) *_LayeredIndexableAttributeGroupLangValueMgr {
	if db == nil {
		panic(fmt.Errorf("LayeredIndexableAttributeGroupLangValueMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_LayeredIndexableAttributeGroupLangValueMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_layered_indexable_attribute_group_lang_value"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_LayeredIndexableAttributeGroupLangValueMgr) GetTableName() string {
	return "ps_layered_indexable_attribute_group_lang_value"
}

func (obj *_LayeredIndexableAttributeGroupLangValueMgr) Get() (result LayeredIndexableAttributeGroupLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_LayeredIndexableAttributeGroupLangValueMgr) Gets() (results []*LayeredIndexableAttributeGroupLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

func (obj *_LayeredIndexableAttributeGroupLangValueMgr) WithIDAttributeGroup(idAttributeGroup int) Option {
	return optionFunc(func(o *options) { o.query["id_attribute_group"] = idAttributeGroup })
}

func (obj *_LayeredIndexableAttributeGroupLangValueMgr) WithIDLang(idLang int) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

func (obj *_LayeredIndexableAttributeGroupLangValueMgr) WithURLName(urlName string) Option {
	return optionFunc(func(o *options) { o.query["url_name"] = urlName })
}

func (obj *_LayeredIndexableAttributeGroupLangValueMgr) WithMetaTitle(metaTitle string) Option {
	return optionFunc(func(o *options) { o.query["meta_title"] = metaTitle })
}

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

func (obj *_LayeredIndexableAttributeGroupLangValueMgr) GetFromIDAttributeGroup(idAttributeGroup int) (results []*LayeredIndexableAttributeGroupLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attribute_group = ?", idAttributeGroup).Find(&results).Error

	return
}

func (obj *_LayeredIndexableAttributeGroupLangValueMgr) GetBatchFromIDAttributeGroup(idAttributeGroups []int) (results []*LayeredIndexableAttributeGroupLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attribute_group IN (?)", idAttributeGroups).Find(&results).Error

	return
}

func (obj *_LayeredIndexableAttributeGroupLangValueMgr) GetFromIDLang(idLang int) (results []*LayeredIndexableAttributeGroupLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error

	return
}

func (obj *_LayeredIndexableAttributeGroupLangValueMgr) GetBatchFromIDLang(idLangs []int) (results []*LayeredIndexableAttributeGroupLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error

	return
}

func (obj *_LayeredIndexableAttributeGroupLangValueMgr) GetFromURLName(urlName string) (results []*LayeredIndexableAttributeGroupLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("url_name = ?", urlName).Find(&results).Error

	return
}

func (obj *_LayeredIndexableAttributeGroupLangValueMgr) GetBatchFromURLName(urlNames []string) (results []*LayeredIndexableAttributeGroupLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("url_name IN (?)", urlNames).Find(&results).Error

	return
}

func (obj *_LayeredIndexableAttributeGroupLangValueMgr) GetFromMetaTitle(metaTitle string) (results []*LayeredIndexableAttributeGroupLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("meta_title = ?", metaTitle).Find(&results).Error

	return
}

func (obj *_LayeredIndexableAttributeGroupLangValueMgr) GetBatchFromMetaTitle(metaTitles []string) (results []*LayeredIndexableAttributeGroupLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("meta_title IN (?)", metaTitles).Find(&results).Error

	return
}

func (obj *_LayeredIndexableAttributeGroupLangValueMgr) FetchByPrimaryKey(idAttributeGroup int, idLang int) (result LayeredIndexableAttributeGroupLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attribute_group = ? AND id_lang = ?", idAttributeGroup, idLang).Find(&result).Error

	return
}
