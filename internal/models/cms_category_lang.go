package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _CmsCategoryLangMgr struct {
	*_BaseMgr
}

func CmsCategoryLangMgr(db *gorm.DB) *_CmsCategoryLangMgr {
	if db == nil {
		panic(fmt.Errorf("CmsCategoryLangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CmsCategoryLangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_cms_category_lang"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_CmsCategoryLangMgr) GetTableName() string {
	return "ps_cms_category_lang"
}

func (obj *_CmsCategoryLangMgr) Get() (result CmsCategoryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_CmsCategoryLangMgr) Gets() (results []*CmsCategoryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_CmsCategoryLangMgr) WithIDCmsCategory(idCmsCategory uint32) Option {
	return optionFunc(func(o *options) { o.query["id_cms_category"] = idCmsCategory })
}

func (obj *_CmsCategoryLangMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

func (obj *_CmsCategoryLangMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

func (obj *_CmsCategoryLangMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

func (obj *_CmsCategoryLangMgr) WithDescription(description string) Option {
	return optionFunc(func(o *options) { o.query["description"] = description })
}

func (obj *_CmsCategoryLangMgr) WithLinkRewrite(linkRewrite string) Option {
	return optionFunc(func(o *options) { o.query["link_rewrite"] = linkRewrite })
}

func (obj *_CmsCategoryLangMgr) WithMetaTitle(metaTitle string) Option {
	return optionFunc(func(o *options) { o.query["meta_title"] = metaTitle })
}

func (obj *_CmsCategoryLangMgr) WithMetaKeywords(metaKeywords string) Option {
	return optionFunc(func(o *options) { o.query["meta_keywords"] = metaKeywords })
}

func (obj *_CmsCategoryLangMgr) WithMetaDescription(metaDescription string) Option {
	return optionFunc(func(o *options) { o.query["meta_description"] = metaDescription })
}

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


func (obj *_CmsCategoryLangMgr) GetFromIDCmsCategory(idCmsCategory uint32) (results []*CmsCategoryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cms_category = ?", idCmsCategory).Find(&results).Error

	return
}

func (obj *_CmsCategoryLangMgr) GetBatchFromIDCmsCategory(idCmsCategorys []uint32) (results []*CmsCategoryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cms_category IN (?)", idCmsCategorys).Find(&results).Error

	return
}

func (obj *_CmsCategoryLangMgr) GetFromIDLang(idLang uint32) (results []*CmsCategoryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error

	return
}

func (obj *_CmsCategoryLangMgr) GetBatchFromIDLang(idLangs []uint32) (results []*CmsCategoryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error

	return
}

func (obj *_CmsCategoryLangMgr) GetFromIDShop(idShop uint32) (results []*CmsCategoryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

func (obj *_CmsCategoryLangMgr) GetBatchFromIDShop(idShops []uint32) (results []*CmsCategoryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

func (obj *_CmsCategoryLangMgr) GetFromName(name string) (results []*CmsCategoryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

func (obj *_CmsCategoryLangMgr) GetBatchFromName(names []string) (results []*CmsCategoryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

func (obj *_CmsCategoryLangMgr) GetFromDescription(description string) (results []*CmsCategoryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("description = ?", description).Find(&results).Error

	return
}

func (obj *_CmsCategoryLangMgr) GetBatchFromDescription(descriptions []string) (results []*CmsCategoryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("description IN (?)", descriptions).Find(&results).Error

	return
}

func (obj *_CmsCategoryLangMgr) GetFromLinkRewrite(linkRewrite string) (results []*CmsCategoryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("link_rewrite = ?", linkRewrite).Find(&results).Error

	return
}

func (obj *_CmsCategoryLangMgr) GetBatchFromLinkRewrite(linkRewrites []string) (results []*CmsCategoryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("link_rewrite IN (?)", linkRewrites).Find(&results).Error

	return
}

func (obj *_CmsCategoryLangMgr) GetFromMetaTitle(metaTitle string) (results []*CmsCategoryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("meta_title = ?", metaTitle).Find(&results).Error

	return
}

func (obj *_CmsCategoryLangMgr) GetBatchFromMetaTitle(metaTitles []string) (results []*CmsCategoryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("meta_title IN (?)", metaTitles).Find(&results).Error

	return
}

func (obj *_CmsCategoryLangMgr) GetFromMetaKeywords(metaKeywords string) (results []*CmsCategoryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("meta_keywords = ?", metaKeywords).Find(&results).Error

	return
}

func (obj *_CmsCategoryLangMgr) GetBatchFromMetaKeywords(metaKeywordss []string) (results []*CmsCategoryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("meta_keywords IN (?)", metaKeywordss).Find(&results).Error

	return
}

func (obj *_CmsCategoryLangMgr) GetFromMetaDescription(metaDescription string) (results []*CmsCategoryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("meta_description = ?", metaDescription).Find(&results).Error

	return
}

func (obj *_CmsCategoryLangMgr) GetBatchFromMetaDescription(metaDescriptions []string) (results []*CmsCategoryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("meta_description IN (?)", metaDescriptions).Find(&results).Error

	return
}


func (obj *_CmsCategoryLangMgr) FetchByPrimaryKey(idCmsCategory uint32, idLang uint32, idShop uint32) (result CmsCategoryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cms_category = ? AND id_lang = ? AND id_shop = ?", idCmsCategory, idLang, idShop).Find(&result).Error

	return
}

func (obj *_CmsCategoryLangMgr) FetchIndexByCategoryName(name string) (results []*CmsCategoryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}
