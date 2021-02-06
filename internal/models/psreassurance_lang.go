package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _PsreassuranceLangMgr struct {
	*_BaseMgr
}

func PsreassuranceLangMgr(db *gorm.DB) *_PsreassuranceLangMgr {
	if db == nil {
		panic(fmt.Errorf("PsreassuranceLangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_PsreassuranceLangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_psreassurance_lang"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_PsreassuranceLangMgr) GetTableName() string {
	return "ps_psreassurance_lang"
}

func (obj *_PsreassuranceLangMgr) Get() (result PsreassuranceLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_PsreassuranceLangMgr) Gets() (results []*PsreassuranceLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_PsreassuranceLangMgr) WithIDPsreassurance(idPsreassurance uint32) Option {
	return optionFunc(func(o *options) { o.query["id_psreassurance"] = idPsreassurance })
}

func (obj *_PsreassuranceLangMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

func (obj *_PsreassuranceLangMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

func (obj *_PsreassuranceLangMgr) WithTitle(title string) Option {
	return optionFunc(func(o *options) { o.query["title"] = title })
}

func (obj *_PsreassuranceLangMgr) WithDescription(description string) Option {
	return optionFunc(func(o *options) { o.query["description"] = description })
}

func (obj *_PsreassuranceLangMgr) WithLink(link string) Option {
	return optionFunc(func(o *options) { o.query["link"] = link })
}

func (obj *_PsreassuranceLangMgr) GetByOption(opts ...Option) (result PsreassuranceLang, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_PsreassuranceLangMgr) GetByOptions(opts ...Option) (results []*PsreassuranceLang, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_PsreassuranceLangMgr) GetFromIDPsreassurance(idPsreassurance uint32) (results []*PsreassuranceLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_psreassurance = ?", idPsreassurance).Find(&results).Error

	return
}

func (obj *_PsreassuranceLangMgr) GetBatchFromIDPsreassurance(idPsreassurances []uint32) (results []*PsreassuranceLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_psreassurance IN (?)", idPsreassurances).Find(&results).Error

	return
}

func (obj *_PsreassuranceLangMgr) GetFromIDLang(idLang uint32) (results []*PsreassuranceLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error

	return
}

func (obj *_PsreassuranceLangMgr) GetBatchFromIDLang(idLangs []uint32) (results []*PsreassuranceLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error

	return
}

func (obj *_PsreassuranceLangMgr) GetFromIDShop(idShop uint32) (results []*PsreassuranceLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

func (obj *_PsreassuranceLangMgr) GetBatchFromIDShop(idShops []uint32) (results []*PsreassuranceLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

func (obj *_PsreassuranceLangMgr) GetFromTitle(title string) (results []*PsreassuranceLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("title = ?", title).Find(&results).Error

	return
}

func (obj *_PsreassuranceLangMgr) GetBatchFromTitle(titles []string) (results []*PsreassuranceLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("title IN (?)", titles).Find(&results).Error

	return
}

func (obj *_PsreassuranceLangMgr) GetFromDescription(description string) (results []*PsreassuranceLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("description = ?", description).Find(&results).Error

	return
}

func (obj *_PsreassuranceLangMgr) GetBatchFromDescription(descriptions []string) (results []*PsreassuranceLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("description IN (?)", descriptions).Find(&results).Error

	return
}

func (obj *_PsreassuranceLangMgr) GetFromLink(link string) (results []*PsreassuranceLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("link = ?", link).Find(&results).Error

	return
}

func (obj *_PsreassuranceLangMgr) GetBatchFromLink(links []string) (results []*PsreassuranceLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("link IN (?)", links).Find(&results).Error

	return
}


func (obj *_PsreassuranceLangMgr) FetchByPrimaryKey(idPsreassurance uint32, idLang uint32, idShop uint32) (result PsreassuranceLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_psreassurance = ? AND id_lang = ? AND id_shop = ?", idPsreassurance, idLang, idShop).Find(&result).Error

	return
}
