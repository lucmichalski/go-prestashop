package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _CurrencyLangMgr struct {
	*_BaseMgr
}

func CurrencyLangMgr(db *gorm.DB) *_CurrencyLangMgr {
	if db == nil {
		panic(fmt.Errorf("CurrencyLangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CurrencyLangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_currency_lang"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_CurrencyLangMgr) GetTableName() string {
	return "ps_currency_lang"
}

func (obj *_CurrencyLangMgr) Get() (result CurrencyLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_CurrencyLangMgr) Gets() (results []*CurrencyLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_CurrencyLangMgr) WithIDCurrency(idCurrency uint32) Option {
	return optionFunc(func(o *options) { o.query["id_currency"] = idCurrency })
}

func (obj *_CurrencyLangMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

func (obj *_CurrencyLangMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

func (obj *_CurrencyLangMgr) WithSymbol(symbol string) Option {
	return optionFunc(func(o *options) { o.query["symbol"] = symbol })
}

func (obj *_CurrencyLangMgr) WithPattern(pattern string) Option {
	return optionFunc(func(o *options) { o.query["pattern"] = pattern })
}

func (obj *_CurrencyLangMgr) GetByOption(opts ...Option) (result CurrencyLang, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_CurrencyLangMgr) GetByOptions(opts ...Option) (results []*CurrencyLang, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_CurrencyLangMgr) GetFromIDCurrency(idCurrency uint32) (results []*CurrencyLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_currency = ?", idCurrency).Find(&results).Error

	return
}

func (obj *_CurrencyLangMgr) GetBatchFromIDCurrency(idCurrencys []uint32) (results []*CurrencyLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_currency IN (?)", idCurrencys).Find(&results).Error

	return
}

func (obj *_CurrencyLangMgr) GetFromIDLang(idLang uint32) (results []*CurrencyLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error

	return
}

func (obj *_CurrencyLangMgr) GetBatchFromIDLang(idLangs []uint32) (results []*CurrencyLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error

	return
}

func (obj *_CurrencyLangMgr) GetFromName(name string) (results []*CurrencyLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

func (obj *_CurrencyLangMgr) GetBatchFromName(names []string) (results []*CurrencyLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

func (obj *_CurrencyLangMgr) GetFromSymbol(symbol string) (results []*CurrencyLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("symbol = ?", symbol).Find(&results).Error

	return
}

func (obj *_CurrencyLangMgr) GetBatchFromSymbol(symbols []string) (results []*CurrencyLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("symbol IN (?)", symbols).Find(&results).Error

	return
}

func (obj *_CurrencyLangMgr) GetFromPattern(pattern string) (results []*CurrencyLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("pattern = ?", pattern).Find(&results).Error

	return
}

func (obj *_CurrencyLangMgr) GetBatchFromPattern(patterns []string) (results []*CurrencyLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("pattern IN (?)", patterns).Find(&results).Error

	return
}


func (obj *_CurrencyLangMgr) FetchByPrimaryKey(idCurrency uint32, idLang uint32) (result CurrencyLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_currency = ? AND id_lang = ?", idCurrency, idLang).Find(&result).Error

	return
}
