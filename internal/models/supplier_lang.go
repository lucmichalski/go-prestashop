package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _SupplierLangMgr struct {
	*_BaseMgr
}

// SupplierLangMgr open func
func SupplierLangMgr(db *gorm.DB) *_SupplierLangMgr {
	if db == nil {
		panic(fmt.Errorf("SupplierLangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SupplierLangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_supplier_lang"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SupplierLangMgr) GetTableName() string {
	return "ps_supplier_lang"
}

// Get 获取
func (obj *_SupplierLangMgr) Get() (result SupplierLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SupplierLangMgr) Gets() (results []*SupplierLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDSupplier id_supplier获取
func (obj *_SupplierLangMgr) WithIDSupplier(idSupplier uint32) Option {
	return optionFunc(func(o *options) { o.query["id_supplier"] = idSupplier })
}

// WithIDLang id_lang获取
func (obj *_SupplierLangMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

// WithDescription description获取
func (obj *_SupplierLangMgr) WithDescription(description string) Option {
	return optionFunc(func(o *options) { o.query["description"] = description })
}

// WithMetaTitle meta_title获取
func (obj *_SupplierLangMgr) WithMetaTitle(metaTitle string) Option {
	return optionFunc(func(o *options) { o.query["meta_title"] = metaTitle })
}

// WithMetaKeywords meta_keywords获取
func (obj *_SupplierLangMgr) WithMetaKeywords(metaKeywords string) Option {
	return optionFunc(func(o *options) { o.query["meta_keywords"] = metaKeywords })
}

// WithMetaDescription meta_description获取
func (obj *_SupplierLangMgr) WithMetaDescription(metaDescription string) Option {
	return optionFunc(func(o *options) { o.query["meta_description"] = metaDescription })
}

// GetByOption 功能选项模式获取
func (obj *_SupplierLangMgr) GetByOption(opts ...Option) (result SupplierLang, err error) {
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
func (obj *_SupplierLangMgr) GetByOptions(opts ...Option) (results []*SupplierLang, err error) {
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

// GetFromIDSupplier 通过id_supplier获取内容
func (obj *_SupplierLangMgr) GetFromIDSupplier(idSupplier uint32) (results []*SupplierLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supplier = ?", idSupplier).Find(&results).Error

	return
}

// GetBatchFromIDSupplier 批量唯一主键查找
func (obj *_SupplierLangMgr) GetBatchFromIDSupplier(idSuppliers []uint32) (results []*SupplierLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supplier IN (?)", idSuppliers).Find(&results).Error

	return
}

// GetFromIDLang 通过id_lang获取内容
func (obj *_SupplierLangMgr) GetFromIDLang(idLang uint32) (results []*SupplierLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error

	return
}

// GetBatchFromIDLang 批量唯一主键查找
func (obj *_SupplierLangMgr) GetBatchFromIDLang(idLangs []uint32) (results []*SupplierLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error

	return
}

// GetFromDescription 通过description获取内容
func (obj *_SupplierLangMgr) GetFromDescription(description string) (results []*SupplierLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("description = ?", description).Find(&results).Error

	return
}

// GetBatchFromDescription 批量唯一主键查找
func (obj *_SupplierLangMgr) GetBatchFromDescription(descriptions []string) (results []*SupplierLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("description IN (?)", descriptions).Find(&results).Error

	return
}

// GetFromMetaTitle 通过meta_title获取内容
func (obj *_SupplierLangMgr) GetFromMetaTitle(metaTitle string) (results []*SupplierLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("meta_title = ?", metaTitle).Find(&results).Error

	return
}

// GetBatchFromMetaTitle 批量唯一主键查找
func (obj *_SupplierLangMgr) GetBatchFromMetaTitle(metaTitles []string) (results []*SupplierLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("meta_title IN (?)", metaTitles).Find(&results).Error

	return
}

// GetFromMetaKeywords 通过meta_keywords获取内容
func (obj *_SupplierLangMgr) GetFromMetaKeywords(metaKeywords string) (results []*SupplierLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("meta_keywords = ?", metaKeywords).Find(&results).Error

	return
}

// GetBatchFromMetaKeywords 批量唯一主键查找
func (obj *_SupplierLangMgr) GetBatchFromMetaKeywords(metaKeywordss []string) (results []*SupplierLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("meta_keywords IN (?)", metaKeywordss).Find(&results).Error

	return
}

// GetFromMetaDescription 通过meta_description获取内容
func (obj *_SupplierLangMgr) GetFromMetaDescription(metaDescription string) (results []*SupplierLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("meta_description = ?", metaDescription).Find(&results).Error

	return
}

// GetBatchFromMetaDescription 批量唯一主键查找
func (obj *_SupplierLangMgr) GetBatchFromMetaDescription(metaDescriptions []string) (results []*SupplierLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("meta_description IN (?)", metaDescriptions).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_SupplierLangMgr) FetchByPrimaryKey(idSupplier uint32, idLang uint32) (result SupplierLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supplier = ? AND id_lang = ?", idSupplier, idLang).Find(&result).Error

	return
}
