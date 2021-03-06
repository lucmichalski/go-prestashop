package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _CmsLangMgr struct {
	*_BaseMgr
}

func CmsLangMgr(db *gorm.DB) *_CmsLangMgr {
	if db == nil {
		panic(fmt.Errorf("CmsLangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CmsLangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_cms_lang"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_CmsLangMgr) GetTableName() string {
	return "ps_cms_lang"
}

func (obj *_CmsLangMgr) Get() (result CmsLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_CmsLangMgr) Gets() (results []*CmsLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

func (obj *_CmsLangMgr) WithIDCms(idCms uint32) Option {
	return optionFunc(func(o *options) { o.query["id_cms"] = idCms })
}

func (obj *_CmsLangMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

func (obj *_CmsLangMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

func (obj *_CmsLangMgr) WithMetaTitle(metaTitle string) Option {
	return optionFunc(func(o *options) { o.query["meta_title"] = metaTitle })
}

func (obj *_CmsLangMgr) WithHeadSeoTitle(headSeoTitle string) Option {
	return optionFunc(func(o *options) { o.query["head_seo_title"] = headSeoTitle })
}

func (obj *_CmsLangMgr) WithMetaDescription(metaDescription string) Option {
	return optionFunc(func(o *options) { o.query["meta_description"] = metaDescription })
}

func (obj *_CmsLangMgr) WithMetaKeywords(metaKeywords string) Option {
	return optionFunc(func(o *options) { o.query["meta_keywords"] = metaKeywords })
}

func (obj *_CmsLangMgr) WithContent(content string) Option {
	return optionFunc(func(o *options) { o.query["content"] = content })
}

func (obj *_CmsLangMgr) WithLinkRewrite(linkRewrite string) Option {
	return optionFunc(func(o *options) { o.query["link_rewrite"] = linkRewrite })
}

func (obj *_CmsLangMgr) GetByOption(opts ...Option) (result CmsLang, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_CmsLangMgr) GetByOptions(opts ...Option) (results []*CmsLang, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

func (obj *_CmsLangMgr) GetFromIDCms(idCms uint32) (results []*CmsLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cms = ?", idCms).Find(&results).Error

	return
}

func (obj *_CmsLangMgr) GetBatchFromIDCms(idCmss []uint32) (results []*CmsLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cms IN (?)", idCmss).Find(&results).Error

	return
}

func (obj *_CmsLangMgr) GetFromIDLang(idLang uint32) (results []*CmsLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error

	return
}

func (obj *_CmsLangMgr) GetBatchFromIDLang(idLangs []uint32) (results []*CmsLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error

	return
}

func (obj *_CmsLangMgr) GetFromIDShop(idShop uint32) (results []*CmsLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

func (obj *_CmsLangMgr) GetBatchFromIDShop(idShops []uint32) (results []*CmsLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

func (obj *_CmsLangMgr) GetFromMetaTitle(metaTitle string) (results []*CmsLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("meta_title = ?", metaTitle).Find(&results).Error

	return
}

func (obj *_CmsLangMgr) GetBatchFromMetaTitle(metaTitles []string) (results []*CmsLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("meta_title IN (?)", metaTitles).Find(&results).Error

	return
}

func (obj *_CmsLangMgr) GetFromHeadSeoTitle(headSeoTitle string) (results []*CmsLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("head_seo_title = ?", headSeoTitle).Find(&results).Error

	return
}

func (obj *_CmsLangMgr) GetBatchFromHeadSeoTitle(headSeoTitles []string) (results []*CmsLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("head_seo_title IN (?)", headSeoTitles).Find(&results).Error

	return
}

func (obj *_CmsLangMgr) GetFromMetaDescription(metaDescription string) (results []*CmsLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("meta_description = ?", metaDescription).Find(&results).Error

	return
}

func (obj *_CmsLangMgr) GetBatchFromMetaDescription(metaDescriptions []string) (results []*CmsLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("meta_description IN (?)", metaDescriptions).Find(&results).Error

	return
}

func (obj *_CmsLangMgr) GetFromMetaKeywords(metaKeywords string) (results []*CmsLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("meta_keywords = ?", metaKeywords).Find(&results).Error

	return
}

func (obj *_CmsLangMgr) GetBatchFromMetaKeywords(metaKeywordss []string) (results []*CmsLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("meta_keywords IN (?)", metaKeywordss).Find(&results).Error

	return
}

func (obj *_CmsLangMgr) GetFromContent(content string) (results []*CmsLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("content = ?", content).Find(&results).Error

	return
}

func (obj *_CmsLangMgr) GetBatchFromContent(contents []string) (results []*CmsLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("content IN (?)", contents).Find(&results).Error

	return
}

func (obj *_CmsLangMgr) GetFromLinkRewrite(linkRewrite string) (results []*CmsLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("link_rewrite = ?", linkRewrite).Find(&results).Error

	return
}

func (obj *_CmsLangMgr) GetBatchFromLinkRewrite(linkRewrites []string) (results []*CmsLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("link_rewrite IN (?)", linkRewrites).Find(&results).Error

	return
}

func (obj *_CmsLangMgr) FetchByPrimaryKey(idCms uint32, idLang uint32, idShop uint32) (result CmsLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cms = ? AND id_lang = ? AND id_shop = ?", idCms, idLang, idShop).Find(&result).Error

	return
}
