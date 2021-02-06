package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _CurrencyLangMgr struct {
	*_BaseMgr
}

// CurrencyLangMgr open func
func CurrencyLangMgr(db *gorm.DB) *_CurrencyLangMgr {
	if db == nil {
		panic(fmt.Errorf("CurrencyLangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CurrencyLangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_currency_lang"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CurrencyLangMgr) GetTableName() string {
	return "ps_currency_lang"
}

// Get 获取
func (obj *_CurrencyLangMgr) Get() (result CurrencyLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CurrencyLangMgr) Gets() (results []*CurrencyLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDCurrency id_currency获取
func (obj *_CurrencyLangMgr) WithIDCurrency(idCurrency uint32) Option {
	return optionFunc(func(o *options) { o.query["id_currency"] = idCurrency })
}

// WithIDLang id_lang获取
func (obj *_CurrencyLangMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

// WithName name获取
func (obj *_CurrencyLangMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithSymbol symbol获取
func (obj *_CurrencyLangMgr) WithSymbol(symbol string) Option {
	return optionFunc(func(o *options) { o.query["symbol"] = symbol })
}

// WithPattern pattern获取
func (obj *_CurrencyLangMgr) WithPattern(pattern string) Option {
	return optionFunc(func(o *options) { o.query["pattern"] = pattern })
}

// GetByOption 功能选项模式获取
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

// GetByOptions 批量功能选项模式获取
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

//////////////////////////enume case ////////////////////////////////////////////

// GetFromIDCurrency 通过id_currency获取内容
func (obj *_CurrencyLangMgr) GetFromIDCurrency(idCurrency uint32) (results []*CurrencyLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_currency = ?", idCurrency).Find(&results).Error

	return
}

// GetBatchFromIDCurrency 批量唯一主键查找
func (obj *_CurrencyLangMgr) GetBatchFromIDCurrency(idCurrencys []uint32) (results []*CurrencyLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_currency IN (?)", idCurrencys).Find(&results).Error

	return
}

// GetFromIDLang 通过id_lang获取内容
func (obj *_CurrencyLangMgr) GetFromIDLang(idLang uint32) (results []*CurrencyLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error

	return
}

// GetBatchFromIDLang 批量唯一主键查找
func (obj *_CurrencyLangMgr) GetBatchFromIDLang(idLangs []uint32) (results []*CurrencyLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_CurrencyLangMgr) GetFromName(name string) (results []*CurrencyLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量唯一主键查找
func (obj *_CurrencyLangMgr) GetBatchFromName(names []string) (results []*CurrencyLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

// GetFromSymbol 通过symbol获取内容
func (obj *_CurrencyLangMgr) GetFromSymbol(symbol string) (results []*CurrencyLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("symbol = ?", symbol).Find(&results).Error

	return
}

// GetBatchFromSymbol 批量唯一主键查找
func (obj *_CurrencyLangMgr) GetBatchFromSymbol(symbols []string) (results []*CurrencyLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("symbol IN (?)", symbols).Find(&results).Error

	return
}

// GetFromPattern 通过pattern获取内容
func (obj *_CurrencyLangMgr) GetFromPattern(pattern string) (results []*CurrencyLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("pattern = ?", pattern).Find(&results).Error

	return
}

// GetBatchFromPattern 批量唯一主键查找
func (obj *_CurrencyLangMgr) GetBatchFromPattern(patterns []string) (results []*CurrencyLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("pattern IN (?)", patterns).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_CurrencyLangMgr) FetchByPrimaryKey(idCurrency uint32, idLang uint32) (result CurrencyLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_currency = ? AND id_lang = ?", idCurrency, idLang).Find(&result).Error

	return
}
