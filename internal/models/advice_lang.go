package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _AdviceLangMgr struct {
	*_BaseMgr
}

func AdviceLangMgr(db *gorm.DB) *_AdviceLangMgr {
	if db == nil {
		panic(fmt.Errorf("AdviceLangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AdviceLangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_advice_lang"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_AdviceLangMgr) GetTableName() string {
	return "ps_advice_lang"
}

func (obj *_AdviceLangMgr) Get() (result AdviceLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_AdviceLangMgr) Gets() (results []*AdviceLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_AdviceLangMgr) WithIDAdvice(idAdvice int) Option {
	return optionFunc(func(o *options) { o.query["id_advice"] = idAdvice })
}

func (obj *_AdviceLangMgr) WithIDLang(idLang int) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

func (obj *_AdviceLangMgr) WithHTML(html string) Option {
	return optionFunc(func(o *options) { o.query["html"] = html })
}

func (obj *_AdviceLangMgr) GetByOption(opts ...Option) (result AdviceLang, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_AdviceLangMgr) GetByOptions(opts ...Option) (results []*AdviceLang, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_AdviceLangMgr) GetFromIDAdvice(idAdvice int) (results []*AdviceLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_advice = ?", idAdvice).Find(&results).Error

	return
}

func (obj *_AdviceLangMgr) GetBatchFromIDAdvice(idAdvices []int) (results []*AdviceLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_advice IN (?)", idAdvices).Find(&results).Error

	return
}

func (obj *_AdviceLangMgr) GetFromIDLang(idLang int) (results []*AdviceLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error

	return
}

func (obj *_AdviceLangMgr) GetBatchFromIDLang(idLangs []int) (results []*AdviceLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error

	return
}

func (obj *_AdviceLangMgr) GetFromHTML(html string) (results []*AdviceLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("html = ?", html).Find(&results).Error

	return
}

func (obj *_AdviceLangMgr) GetBatchFromHTML(htmls []string) (results []*AdviceLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("html IN (?)", htmls).Find(&results).Error

	return
}


func (obj *_AdviceLangMgr) FetchByPrimaryKey(idAdvice int, idLang int) (result AdviceLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_advice = ? AND id_lang = ?", idAdvice, idLang).Find(&result).Error

	return
}
