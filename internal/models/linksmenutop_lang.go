package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _LinksmenutopLangMgr struct {
	*_BaseMgr
}

func LinksmenutopLangMgr(db *gorm.DB) *_LinksmenutopLangMgr {
	if db == nil {
		panic(fmt.Errorf("LinksmenutopLangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_LinksmenutopLangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_linksmenutop_lang"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_LinksmenutopLangMgr) GetTableName() string {
	return "ps_linksmenutop_lang"
}

func (obj *_LinksmenutopLangMgr) Get() (result LinksmenutopLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_LinksmenutopLangMgr) Gets() (results []*LinksmenutopLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

func (obj *_LinksmenutopLangMgr) WithIDLinksmenutop(idLinksmenutop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_linksmenutop"] = idLinksmenutop })
}

func (obj *_LinksmenutopLangMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

func (obj *_LinksmenutopLangMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

func (obj *_LinksmenutopLangMgr) WithLabel(label string) Option {
	return optionFunc(func(o *options) { o.query["label"] = label })
}

func (obj *_LinksmenutopLangMgr) WithLink(link string) Option {
	return optionFunc(func(o *options) { o.query["link"] = link })
}

func (obj *_LinksmenutopLangMgr) GetByOption(opts ...Option) (result LinksmenutopLang, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_LinksmenutopLangMgr) GetByOptions(opts ...Option) (results []*LinksmenutopLang, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

func (obj *_LinksmenutopLangMgr) GetFromIDLinksmenutop(idLinksmenutop uint32) (results []*LinksmenutopLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_linksmenutop = ?", idLinksmenutop).Find(&results).Error

	return
}

func (obj *_LinksmenutopLangMgr) GetBatchFromIDLinksmenutop(idLinksmenutops []uint32) (results []*LinksmenutopLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_linksmenutop IN (?)", idLinksmenutops).Find(&results).Error

	return
}

func (obj *_LinksmenutopLangMgr) GetFromIDLang(idLang uint32) (results []*LinksmenutopLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error

	return
}

func (obj *_LinksmenutopLangMgr) GetBatchFromIDLang(idLangs []uint32) (results []*LinksmenutopLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error

	return
}

func (obj *_LinksmenutopLangMgr) GetFromIDShop(idShop uint32) (results []*LinksmenutopLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

func (obj *_LinksmenutopLangMgr) GetBatchFromIDShop(idShops []uint32) (results []*LinksmenutopLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

func (obj *_LinksmenutopLangMgr) GetFromLabel(label string) (results []*LinksmenutopLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("label = ?", label).Find(&results).Error

	return
}

func (obj *_LinksmenutopLangMgr) GetBatchFromLabel(labels []string) (results []*LinksmenutopLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("label IN (?)", labels).Find(&results).Error

	return
}

func (obj *_LinksmenutopLangMgr) GetFromLink(link string) (results []*LinksmenutopLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("link = ?", link).Find(&results).Error

	return
}

func (obj *_LinksmenutopLangMgr) GetBatchFromLink(links []string) (results []*LinksmenutopLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("link IN (?)", links).Find(&results).Error

	return
}

func (obj *_LinksmenutopLangMgr) FetchIndexByIDLinksmenutop(idLinksmenutop uint32, idLang uint32, idShop uint32) (results []*LinksmenutopLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_linksmenutop = ? AND id_lang = ? AND id_shop = ?", idLinksmenutop, idLang, idShop).Find(&results).Error

	return
}
