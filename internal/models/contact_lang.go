package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _ContactLangMgr struct {
	*_BaseMgr
}

func ContactLangMgr(db *gorm.DB) *_ContactLangMgr {
	if db == nil {
		panic(fmt.Errorf("ContactLangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ContactLangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_contact_lang"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_ContactLangMgr) GetTableName() string {
	return "ps_contact_lang"
}

func (obj *_ContactLangMgr) Get() (result ContactLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_ContactLangMgr) Gets() (results []*ContactLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_ContactLangMgr) WithIDContact(idContact uint32) Option {
	return optionFunc(func(o *options) { o.query["id_contact"] = idContact })
}

func (obj *_ContactLangMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

func (obj *_ContactLangMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

func (obj *_ContactLangMgr) WithDescription(description string) Option {
	return optionFunc(func(o *options) { o.query["description"] = description })
}

func (obj *_ContactLangMgr) GetByOption(opts ...Option) (result ContactLang, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_ContactLangMgr) GetByOptions(opts ...Option) (results []*ContactLang, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_ContactLangMgr) GetFromIDContact(idContact uint32) (results []*ContactLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_contact = ?", idContact).Find(&results).Error

	return
}

func (obj *_ContactLangMgr) GetBatchFromIDContact(idContacts []uint32) (results []*ContactLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_contact IN (?)", idContacts).Find(&results).Error

	return
}

func (obj *_ContactLangMgr) GetFromIDLang(idLang uint32) (results []*ContactLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error

	return
}

func (obj *_ContactLangMgr) GetBatchFromIDLang(idLangs []uint32) (results []*ContactLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error

	return
}

func (obj *_ContactLangMgr) GetFromName(name string) (results []*ContactLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

func (obj *_ContactLangMgr) GetBatchFromName(names []string) (results []*ContactLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

func (obj *_ContactLangMgr) GetFromDescription(description string) (results []*ContactLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("description = ?", description).Find(&results).Error

	return
}

func (obj *_ContactLangMgr) GetBatchFromDescription(descriptions []string) (results []*ContactLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("description IN (?)", descriptions).Find(&results).Error

	return
}


func (obj *_ContactLangMgr) FetchByPrimaryKey(idContact uint32, idLang uint32) (result ContactLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_contact = ? AND id_lang = ?", idContact, idLang).Find(&result).Error

	return
}
