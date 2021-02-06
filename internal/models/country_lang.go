package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _CountryLangMgr struct {
	*_BaseMgr
}

func CountryLangMgr(db *gorm.DB) *_CountryLangMgr {
	if db == nil {
		panic(fmt.Errorf("CountryLangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CountryLangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_country_lang"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_CountryLangMgr) GetTableName() string {
	return "ps_country_lang"
}

func (obj *_CountryLangMgr) Get() (result CountryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_CountryLangMgr) Gets() (results []*CountryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_CountryLangMgr) WithIDCountry(idCountry uint32) Option {
	return optionFunc(func(o *options) { o.query["id_country"] = idCountry })
}

func (obj *_CountryLangMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

func (obj *_CountryLangMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

func (obj *_CountryLangMgr) GetByOption(opts ...Option) (result CountryLang, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_CountryLangMgr) GetByOptions(opts ...Option) (results []*CountryLang, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_CountryLangMgr) GetFromIDCountry(idCountry uint32) (results []*CountryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_country = ?", idCountry).Find(&results).Error

	return
}

func (obj *_CountryLangMgr) GetBatchFromIDCountry(idCountrys []uint32) (results []*CountryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_country IN (?)", idCountrys).Find(&results).Error

	return
}

func (obj *_CountryLangMgr) GetFromIDLang(idLang uint32) (results []*CountryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error

	return
}

func (obj *_CountryLangMgr) GetBatchFromIDLang(idLangs []uint32) (results []*CountryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error

	return
}

func (obj *_CountryLangMgr) GetFromName(name string) (results []*CountryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

func (obj *_CountryLangMgr) GetBatchFromName(names []string) (results []*CountryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}


func (obj *_CountryLangMgr) FetchByPrimaryKey(idCountry uint32, idLang uint32) (result CountryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_country = ? AND id_lang = ?", idCountry, idLang).Find(&result).Error

	return
}
