package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _StockMvtReasonLangMgr struct {
	*_BaseMgr
}

func StockMvtReasonLangMgr(db *gorm.DB) *_StockMvtReasonLangMgr {
	if db == nil {
		panic(fmt.Errorf("StockMvtReasonLangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_StockMvtReasonLangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_stock_mvt_reason_lang"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_StockMvtReasonLangMgr) GetTableName() string {
	return "ps_stock_mvt_reason_lang"
}

func (obj *_StockMvtReasonLangMgr) Get() (result StockMvtReasonLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_StockMvtReasonLangMgr) Gets() (results []*StockMvtReasonLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_StockMvtReasonLangMgr) WithIDStockMvtReason(idStockMvtReason uint32) Option {
	return optionFunc(func(o *options) { o.query["id_stock_mvt_reason"] = idStockMvtReason })
}

func (obj *_StockMvtReasonLangMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

func (obj *_StockMvtReasonLangMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

func (obj *_StockMvtReasonLangMgr) GetByOption(opts ...Option) (result StockMvtReasonLang, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_StockMvtReasonLangMgr) GetByOptions(opts ...Option) (results []*StockMvtReasonLang, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_StockMvtReasonLangMgr) GetFromIDStockMvtReason(idStockMvtReason uint32) (results []*StockMvtReasonLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_stock_mvt_reason = ?", idStockMvtReason).Find(&results).Error

	return
}

func (obj *_StockMvtReasonLangMgr) GetBatchFromIDStockMvtReason(idStockMvtReasons []uint32) (results []*StockMvtReasonLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_stock_mvt_reason IN (?)", idStockMvtReasons).Find(&results).Error

	return
}

func (obj *_StockMvtReasonLangMgr) GetFromIDLang(idLang uint32) (results []*StockMvtReasonLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error

	return
}

func (obj *_StockMvtReasonLangMgr) GetBatchFromIDLang(idLangs []uint32) (results []*StockMvtReasonLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error

	return
}

func (obj *_StockMvtReasonLangMgr) GetFromName(name string) (results []*StockMvtReasonLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

func (obj *_StockMvtReasonLangMgr) GetBatchFromName(names []string) (results []*StockMvtReasonLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}


func (obj *_StockMvtReasonLangMgr) FetchByPrimaryKey(idStockMvtReason uint32, idLang uint32) (result StockMvtReasonLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_stock_mvt_reason = ? AND id_lang = ?", idStockMvtReason, idLang).Find(&result).Error

	return
}
