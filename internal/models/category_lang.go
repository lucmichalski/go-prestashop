package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _CategoryLangMgr struct {
	*_BaseMgr
}

// CategoryLangMgr open func
func CategoryLangMgr(db *gorm.DB) *_CategoryLangMgr {
	if db == nil {
		panic(fmt.Errorf("CategoryLangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CategoryLangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_category_lang"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CategoryLangMgr) GetTableName() string {
	return "eg_category_lang"
}

// Get 获取
func (obj *_CategoryLangMgr) Get() (result CategoryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CategoryLangMgr) Gets() (results []*CategoryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDCategory id_category获取
func (obj *_CategoryLangMgr) WithIDCategory(idCategory uint32) Option {
	return optionFunc(func(o *options) { o.query["id_category"] = idCategory })
}

// WithIDShop id_shop获取
func (obj *_CategoryLangMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

// WithIDLang id_lang获取
func (obj *_CategoryLangMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

// WithName name获取
func (obj *_CategoryLangMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithDescription description获取
func (obj *_CategoryLangMgr) WithDescription(description string) Option {
	return optionFunc(func(o *options) { o.query["description"] = description })
}

// WithLinkRewrite link_rewrite获取
func (obj *_CategoryLangMgr) WithLinkRewrite(linkRewrite string) Option {
	return optionFunc(func(o *options) { o.query["link_rewrite"] = linkRewrite })
}

// WithMetaTitle meta_title获取
func (obj *_CategoryLangMgr) WithMetaTitle(metaTitle string) Option {
	return optionFunc(func(o *options) { o.query["meta_title"] = metaTitle })
}

// WithMetaKeywords meta_keywords获取
func (obj *_CategoryLangMgr) WithMetaKeywords(metaKeywords string) Option {
	return optionFunc(func(o *options) { o.query["meta_keywords"] = metaKeywords })
}

// WithMetaDescription meta_description获取
func (obj *_CategoryLangMgr) WithMetaDescription(metaDescription string) Option {
	return optionFunc(func(o *options) { o.query["meta_description"] = metaDescription })
}

// GetByOption 功能选项模式获取
func (obj *_CategoryLangMgr) GetByOption(opts ...Option) (result CategoryLang, err error) {
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
func (obj *_CategoryLangMgr) GetByOptions(opts ...Option) (results []*CategoryLang, err error) {
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

// GetFromIDCategory 通过id_category获取内容
func (obj *_CategoryLangMgr) GetFromIDCategory(idCategory uint32) (results []*CategoryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_category = ?", idCategory).Find(&results).Error

	return
}

// GetBatchFromIDCategory 批量唯一主键查找
func (obj *_CategoryLangMgr) GetBatchFromIDCategory(idCategorys []uint32) (results []*CategoryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_category IN (?)", idCategorys).Find(&results).Error

	return
}

// GetFromIDShop 通过id_shop获取内容
func (obj *_CategoryLangMgr) GetFromIDShop(idShop uint32) (results []*CategoryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

// GetBatchFromIDShop 批量唯一主键查找
func (obj *_CategoryLangMgr) GetBatchFromIDShop(idShops []uint32) (results []*CategoryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

// GetFromIDLang 通过id_lang获取内容
func (obj *_CategoryLangMgr) GetFromIDLang(idLang uint32) (results []*CategoryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error

	return
}

// GetBatchFromIDLang 批量唯一主键查找
func (obj *_CategoryLangMgr) GetBatchFromIDLang(idLangs []uint32) (results []*CategoryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_CategoryLangMgr) GetFromName(name string) (results []*CategoryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量唯一主键查找
func (obj *_CategoryLangMgr) GetBatchFromName(names []string) (results []*CategoryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

// GetFromDescription 通过description获取内容
func (obj *_CategoryLangMgr) GetFromDescription(description string) (results []*CategoryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("description = ?", description).Find(&results).Error

	return
}

// GetBatchFromDescription 批量唯一主键查找
func (obj *_CategoryLangMgr) GetBatchFromDescription(descriptions []string) (results []*CategoryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("description IN (?)", descriptions).Find(&results).Error

	return
}

// GetFromLinkRewrite 通过link_rewrite获取内容
func (obj *_CategoryLangMgr) GetFromLinkRewrite(linkRewrite string) (results []*CategoryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("link_rewrite = ?", linkRewrite).Find(&results).Error

	return
}

// GetBatchFromLinkRewrite 批量唯一主键查找
func (obj *_CategoryLangMgr) GetBatchFromLinkRewrite(linkRewrites []string) (results []*CategoryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("link_rewrite IN (?)", linkRewrites).Find(&results).Error

	return
}

// GetFromMetaTitle 通过meta_title获取内容
func (obj *_CategoryLangMgr) GetFromMetaTitle(metaTitle string) (results []*CategoryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("meta_title = ?", metaTitle).Find(&results).Error

	return
}

// GetBatchFromMetaTitle 批量唯一主键查找
func (obj *_CategoryLangMgr) GetBatchFromMetaTitle(metaTitles []string) (results []*CategoryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("meta_title IN (?)", metaTitles).Find(&results).Error

	return
}

// GetFromMetaKeywords 通过meta_keywords获取内容
func (obj *_CategoryLangMgr) GetFromMetaKeywords(metaKeywords string) (results []*CategoryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("meta_keywords = ?", metaKeywords).Find(&results).Error

	return
}

// GetBatchFromMetaKeywords 批量唯一主键查找
func (obj *_CategoryLangMgr) GetBatchFromMetaKeywords(metaKeywordss []string) (results []*CategoryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("meta_keywords IN (?)", metaKeywordss).Find(&results).Error

	return
}

// GetFromMetaDescription 通过meta_description获取内容
func (obj *_CategoryLangMgr) GetFromMetaDescription(metaDescription string) (results []*CategoryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("meta_description = ?", metaDescription).Find(&results).Error

	return
}

// GetBatchFromMetaDescription 批量唯一主键查找
func (obj *_CategoryLangMgr) GetBatchFromMetaDescription(metaDescriptions []string) (results []*CategoryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("meta_description IN (?)", metaDescriptions).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_CategoryLangMgr) FetchByPrimaryKey(idCategory uint32, idShop uint32, idLang uint32) (result CategoryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_category = ? AND id_shop = ? AND id_lang = ?", idCategory, idShop, idLang).Find(&result).Error

	return
}

// FetchIndexByCategoryName  获取多个内容
func (obj *_CategoryLangMgr) FetchIndexByCategoryName(name string) (results []*CategoryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}
