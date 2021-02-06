package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _ManufacturerLangMgr struct {
	*_BaseMgr
}

func ManufacturerLangMgr(db *gorm.DB) *_ManufacturerLangMgr {
	if db == nil {
		panic(fmt.Errorf("ManufacturerLangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ManufacturerLangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_manufacturer_lang"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_ManufacturerLangMgr) GetTableName() string {
	return "ps_manufacturer_lang"
}

func (obj *_ManufacturerLangMgr) Get() (result ManufacturerLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_ManufacturerLangMgr) Gets() (results []*ManufacturerLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_ManufacturerLangMgr) WithIDManufacturer(idManufacturer uint32) Option {
	return optionFunc(func(o *options) { o.query["id_manufacturer"] = idManufacturer })
}

func (obj *_ManufacturerLangMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

func (obj *_ManufacturerLangMgr) WithDescription(description string) Option {
	return optionFunc(func(o *options) { o.query["description"] = description })
}

func (obj *_ManufacturerLangMgr) WithShortDescription(shortDescription string) Option {
	return optionFunc(func(o *options) { o.query["short_description"] = shortDescription })
}

func (obj *_ManufacturerLangMgr) WithMetaTitle(metaTitle string) Option {
	return optionFunc(func(o *options) { o.query["meta_title"] = metaTitle })
}

func (obj *_ManufacturerLangMgr) WithMetaKeywords(metaKeywords string) Option {
	return optionFunc(func(o *options) { o.query["meta_keywords"] = metaKeywords })
}

func (obj *_ManufacturerLangMgr) WithMetaDescription(metaDescription string) Option {
	return optionFunc(func(o *options) { o.query["meta_description"] = metaDescription })
}

func (obj *_ManufacturerLangMgr) GetByOption(opts ...Option) (result ManufacturerLang, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_ManufacturerLangMgr) GetByOptions(opts ...Option) (results []*ManufacturerLang, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_ManufacturerLangMgr) GetFromIDManufacturer(idManufacturer uint32) (results []*ManufacturerLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_manufacturer = ?", idManufacturer).Find(&results).Error

	return
}

func (obj *_ManufacturerLangMgr) GetBatchFromIDManufacturer(idManufacturers []uint32) (results []*ManufacturerLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_manufacturer IN (?)", idManufacturers).Find(&results).Error

	return
}

func (obj *_ManufacturerLangMgr) GetFromIDLang(idLang uint32) (results []*ManufacturerLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error

	return
}

func (obj *_ManufacturerLangMgr) GetBatchFromIDLang(idLangs []uint32) (results []*ManufacturerLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error

	return
}

func (obj *_ManufacturerLangMgr) GetFromDescription(description string) (results []*ManufacturerLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("description = ?", description).Find(&results).Error

	return
}

func (obj *_ManufacturerLangMgr) GetBatchFromDescription(descriptions []string) (results []*ManufacturerLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("description IN (?)", descriptions).Find(&results).Error

	return
}

func (obj *_ManufacturerLangMgr) GetFromShortDescription(shortDescription string) (results []*ManufacturerLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("short_description = ?", shortDescription).Find(&results).Error

	return
}

func (obj *_ManufacturerLangMgr) GetBatchFromShortDescription(shortDescriptions []string) (results []*ManufacturerLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("short_description IN (?)", shortDescriptions).Find(&results).Error

	return
}

func (obj *_ManufacturerLangMgr) GetFromMetaTitle(metaTitle string) (results []*ManufacturerLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("meta_title = ?", metaTitle).Find(&results).Error

	return
}

func (obj *_ManufacturerLangMgr) GetBatchFromMetaTitle(metaTitles []string) (results []*ManufacturerLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("meta_title IN (?)", metaTitles).Find(&results).Error

	return
}

func (obj *_ManufacturerLangMgr) GetFromMetaKeywords(metaKeywords string) (results []*ManufacturerLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("meta_keywords = ?", metaKeywords).Find(&results).Error

	return
}

func (obj *_ManufacturerLangMgr) GetBatchFromMetaKeywords(metaKeywordss []string) (results []*ManufacturerLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("meta_keywords IN (?)", metaKeywordss).Find(&results).Error

	return
}

func (obj *_ManufacturerLangMgr) GetFromMetaDescription(metaDescription string) (results []*ManufacturerLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("meta_description = ?", metaDescription).Find(&results).Error

	return
}

func (obj *_ManufacturerLangMgr) GetBatchFromMetaDescription(metaDescriptions []string) (results []*ManufacturerLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("meta_description IN (?)", metaDescriptions).Find(&results).Error

	return
}


func (obj *_ManufacturerLangMgr) FetchByPrimaryKey(idManufacturer uint32, idLang uint32) (result ManufacturerLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_manufacturer = ? AND id_lang = ?", idManufacturer, idLang).Find(&result).Error

	return
}
