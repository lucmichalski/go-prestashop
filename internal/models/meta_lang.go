package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _MetaLangMgr struct {
	*_BaseMgr
}

func MetaLangMgr(db *gorm.DB) *_MetaLangMgr {
	if db == nil {
		panic(fmt.Errorf("MetaLangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_MetaLangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_meta_lang"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_MetaLangMgr) GetTableName() string {
	return "ps_meta_lang"
}

func (obj *_MetaLangMgr) Get() (result MetaLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_MetaLangMgr) Gets() (results []*MetaLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

func (obj *_MetaLangMgr) WithIDMeta(idMeta uint32) Option {
	return optionFunc(func(o *options) { o.query["id_meta"] = idMeta })
}

func (obj *_MetaLangMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

func (obj *_MetaLangMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

func (obj *_MetaLangMgr) WithTitle(title string) Option {
	return optionFunc(func(o *options) { o.query["title"] = title })
}

func (obj *_MetaLangMgr) WithDescription(description string) Option {
	return optionFunc(func(o *options) { o.query["description"] = description })
}

func (obj *_MetaLangMgr) WithKeywords(keywords string) Option {
	return optionFunc(func(o *options) { o.query["keywords"] = keywords })
}

func (obj *_MetaLangMgr) WithURLRewrite(urlRewrite string) Option {
	return optionFunc(func(o *options) { o.query["url_rewrite"] = urlRewrite })
}

func (obj *_MetaLangMgr) GetByOption(opts ...Option) (result MetaLang, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_MetaLangMgr) GetByOptions(opts ...Option) (results []*MetaLang, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

func (obj *_MetaLangMgr) GetFromIDMeta(idMeta uint32) (results []*MetaLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_meta = ?", idMeta).Find(&results).Error

	return
}

func (obj *_MetaLangMgr) GetBatchFromIDMeta(idMetas []uint32) (results []*MetaLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_meta IN (?)", idMetas).Find(&results).Error

	return
}

func (obj *_MetaLangMgr) GetFromIDShop(idShop uint32) (results []*MetaLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

func (obj *_MetaLangMgr) GetBatchFromIDShop(idShops []uint32) (results []*MetaLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

func (obj *_MetaLangMgr) GetFromIDLang(idLang uint32) (results []*MetaLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error

	return
}

func (obj *_MetaLangMgr) GetBatchFromIDLang(idLangs []uint32) (results []*MetaLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error

	return
}

func (obj *_MetaLangMgr) GetFromTitle(title string) (results []*MetaLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("title = ?", title).Find(&results).Error

	return
}

func (obj *_MetaLangMgr) GetBatchFromTitle(titles []string) (results []*MetaLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("title IN (?)", titles).Find(&results).Error

	return
}

func (obj *_MetaLangMgr) GetFromDescription(description string) (results []*MetaLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("description = ?", description).Find(&results).Error

	return
}

func (obj *_MetaLangMgr) GetBatchFromDescription(descriptions []string) (results []*MetaLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("description IN (?)", descriptions).Find(&results).Error

	return
}

func (obj *_MetaLangMgr) GetFromKeywords(keywords string) (results []*MetaLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("keywords = ?", keywords).Find(&results).Error

	return
}

func (obj *_MetaLangMgr) GetBatchFromKeywords(keywordss []string) (results []*MetaLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("keywords IN (?)", keywordss).Find(&results).Error

	return
}

func (obj *_MetaLangMgr) GetFromURLRewrite(urlRewrite string) (results []*MetaLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("url_rewrite = ?", urlRewrite).Find(&results).Error

	return
}

func (obj *_MetaLangMgr) GetBatchFromURLRewrite(urlRewrites []string) (results []*MetaLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("url_rewrite IN (?)", urlRewrites).Find(&results).Error

	return
}

func (obj *_MetaLangMgr) FetchByPrimaryKey(idMeta uint32, idShop uint32, idLang uint32) (result MetaLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_meta = ? AND id_shop = ? AND id_lang = ?", idMeta, idShop, idLang).Find(&result).Error

	return
}

func (obj *_MetaLangMgr) FetchIndexByIDShop(idShop uint32) (results []*MetaLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

func (obj *_MetaLangMgr) FetchIndexByIDLang(idLang uint32) (results []*MetaLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error

	return
}
