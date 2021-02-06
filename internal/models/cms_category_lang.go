package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _CmsCategoryLangMgr struct {
	*_BaseMgr
}

// CmsCategoryLangMgr open func
func CmsCategoryLangMgr(db *gorm.DB) *_CmsCategoryLangMgr {
	if db == nil {
		panic(fmt.Errorf("CmsCategoryLangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CmsCategoryLangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_cms_category_lang"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CmsCategoryLangMgr) GetTableName() string {
	return "ps_cms_category_lang"
}

// Get 获取
func (obj *_CmsCategoryLangMgr) Get() (result CmsCategoryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CmsCategoryLangMgr) Gets() (results []*CmsCategoryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDCmsCategory id_cms_category获取
func (obj *_CmsCategoryLangMgr) WithIDCmsCategory(idCmsCategory uint32) Option {
	return optionFunc(func(o *options) { o.query["id_cms_category"] = idCmsCategory })
}

// WithIDLang id_lang获取
func (obj *_CmsCategoryLangMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

// WithIDShop id_shop获取
func (obj *_CmsCategoryLangMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

// WithName name获取
func (obj *_CmsCategoryLangMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithDescription description获取
func (obj *_CmsCategoryLangMgr) WithDescription(description string) Option {
	return optionFunc(func(o *options) { o.query["description"] = description })
}

// WithLinkRewrite link_rewrite获取
func (obj *_CmsCategoryLangMgr) WithLinkRewrite(linkRewrite string) Option {
	return optionFunc(func(o *options) { o.query["link_rewrite"] = linkRewrite })
}

// WithMetaTitle meta_title获取
func (obj *_CmsCategoryLangMgr) WithMetaTitle(metaTitle string) Option {
	return optionFunc(func(o *options) { o.query["meta_title"] = metaTitle })
}

// WithMetaKeywords meta_keywords获取
func (obj *_CmsCategoryLangMgr) WithMetaKeywords(metaKeywords string) Option {
	return optionFunc(func(o *options) { o.query["meta_keywords"] = metaKeywords })
}

// WithMetaDescription meta_description获取
func (obj *_CmsCategoryLangMgr) WithMetaDescription(metaDescription string) Option {
	return optionFunc(func(o *options) { o.query["meta_description"] = metaDescription })
}

// GetByOption 功能选项模式获取
func (obj *_CmsCategoryLangMgr) GetByOption(opts ...Option) (result CmsCategoryLang, err error) {
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
func (obj *_CmsCategoryLangMgr) GetByOptions(opts ...Option) (results []*CmsCategoryLang, err error) {
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

// GetFromIDCmsCategory 通过id_cms_category获取内容
func (obj *_CmsCategoryLangMgr) GetFromIDCmsCategory(idCmsCategory uint32) (results []*CmsCategoryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cms_category = ?", idCmsCategory).Find(&results).Error

	return
}

// GetBatchFromIDCmsCategory 批量唯一主键查找
func (obj *_CmsCategoryLangMgr) GetBatchFromIDCmsCategory(idCmsCategorys []uint32) (results []*CmsCategoryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cms_category IN (?)", idCmsCategorys).Find(&results).Error

	return
}

// GetFromIDLang 通过id_lang获取内容
func (obj *_CmsCategoryLangMgr) GetFromIDLang(idLang uint32) (results []*CmsCategoryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error

	return
}

// GetBatchFromIDLang 批量唯一主键查找
func (obj *_CmsCategoryLangMgr) GetBatchFromIDLang(idLangs []uint32) (results []*CmsCategoryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error

	return
}

// GetFromIDShop 通过id_shop获取内容
func (obj *_CmsCategoryLangMgr) GetFromIDShop(idShop uint32) (results []*CmsCategoryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

// GetBatchFromIDShop 批量唯一主键查找
func (obj *_CmsCategoryLangMgr) GetBatchFromIDShop(idShops []uint32) (results []*CmsCategoryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_CmsCategoryLangMgr) GetFromName(name string) (results []*CmsCategoryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量唯一主键查找
func (obj *_CmsCategoryLangMgr) GetBatchFromName(names []string) (results []*CmsCategoryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

// GetFromDescription 通过description获取内容
func (obj *_CmsCategoryLangMgr) GetFromDescription(description string) (results []*CmsCategoryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("description = ?", description).Find(&results).Error

	return
}

// GetBatchFromDescription 批量唯一主键查找
func (obj *_CmsCategoryLangMgr) GetBatchFromDescription(descriptions []string) (results []*CmsCategoryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("description IN (?)", descriptions).Find(&results).Error

	return
}

// GetFromLinkRewrite 通过link_rewrite获取内容
func (obj *_CmsCategoryLangMgr) GetFromLinkRewrite(linkRewrite string) (results []*CmsCategoryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("link_rewrite = ?", linkRewrite).Find(&results).Error

	return
}

// GetBatchFromLinkRewrite 批量唯一主键查找
func (obj *_CmsCategoryLangMgr) GetBatchFromLinkRewrite(linkRewrites []string) (results []*CmsCategoryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("link_rewrite IN (?)", linkRewrites).Find(&results).Error

	return
}

// GetFromMetaTitle 通过meta_title获取内容
func (obj *_CmsCategoryLangMgr) GetFromMetaTitle(metaTitle string) (results []*CmsCategoryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("meta_title = ?", metaTitle).Find(&results).Error

	return
}

// GetBatchFromMetaTitle 批量唯一主键查找
func (obj *_CmsCategoryLangMgr) GetBatchFromMetaTitle(metaTitles []string) (results []*CmsCategoryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("meta_title IN (?)", metaTitles).Find(&results).Error

	return
}

// GetFromMetaKeywords 通过meta_keywords获取内容
func (obj *_CmsCategoryLangMgr) GetFromMetaKeywords(metaKeywords string) (results []*CmsCategoryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("meta_keywords = ?", metaKeywords).Find(&results).Error

	return
}

// GetBatchFromMetaKeywords 批量唯一主键查找
func (obj *_CmsCategoryLangMgr) GetBatchFromMetaKeywords(metaKeywordss []string) (results []*CmsCategoryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("meta_keywords IN (?)", metaKeywordss).Find(&results).Error

	return
}

// GetFromMetaDescription 通过meta_description获取内容
func (obj *_CmsCategoryLangMgr) GetFromMetaDescription(metaDescription string) (results []*CmsCategoryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("meta_description = ?", metaDescription).Find(&results).Error

	return
}

// GetBatchFromMetaDescription 批量唯一主键查找
func (obj *_CmsCategoryLangMgr) GetBatchFromMetaDescription(metaDescriptions []string) (results []*CmsCategoryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("meta_description IN (?)", metaDescriptions).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_CmsCategoryLangMgr) FetchByPrimaryKey(idCmsCategory uint32, idLang uint32, idShop uint32) (result CmsCategoryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cms_category = ? AND id_lang = ? AND id_shop = ?", idCmsCategory, idLang, idShop).Find(&result).Error

	return
}

// FetchIndexByCategoryName  获取多个内容
func (obj *_CmsCategoryLangMgr) FetchIndexByCategoryName(name string) (results []*CmsCategoryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}
