package models

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _StockMvtReasonMgr struct {
	*_BaseMgr
}

func StockMvtReasonMgr(db *gorm.DB) *_StockMvtReasonMgr {
	if db == nil {
		panic(fmt.Errorf("StockMvtReasonMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_StockMvtReasonMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_stock_mvt_reason"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_StockMvtReasonMgr) GetTableName() string {
	return "ps_stock_mvt_reason"
}

func (obj *_StockMvtReasonMgr) Get() (result StockMvtReason, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_StockMvtReasonMgr) Gets() (results []*StockMvtReason, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_StockMvtReasonMgr) WithIDStockMvtReason(idStockMvtReason uint32) Option {
	return optionFunc(func(o *options) { o.query["id_stock_mvt_reason"] = idStockMvtReason })
}

func (obj *_StockMvtReasonMgr) WithSign(sign bool) Option {
	return optionFunc(func(o *options) { o.query["sign"] = sign })
}

func (obj *_StockMvtReasonMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

func (obj *_StockMvtReasonMgr) WithDateUpd(dateUpd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_upd"] = dateUpd })
}

func (obj *_StockMvtReasonMgr) WithDeleted(deleted bool) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

func (obj *_StockMvtReasonMgr) GetByOption(opts ...Option) (result StockMvtReason, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_StockMvtReasonMgr) GetByOptions(opts ...Option) (results []*StockMvtReason, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_StockMvtReasonMgr) GetFromIDStockMvtReason(idStockMvtReason uint32) (result StockMvtReason, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_stock_mvt_reason = ?", idStockMvtReason).Find(&result).Error

	return
}

func (obj *_StockMvtReasonMgr) GetBatchFromIDStockMvtReason(idStockMvtReasons []uint32) (results []*StockMvtReason, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_stock_mvt_reason IN (?)", idStockMvtReasons).Find(&results).Error

	return
}

func (obj *_StockMvtReasonMgr) GetFromSign(sign bool) (results []*StockMvtReason, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("sign = ?", sign).Find(&results).Error

	return
}

func (obj *_StockMvtReasonMgr) GetBatchFromSign(signs []bool) (results []*StockMvtReason, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("sign IN (?)", signs).Find(&results).Error

	return
}

func (obj *_StockMvtReasonMgr) GetFromDateAdd(dateAdd time.Time) (results []*StockMvtReason, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error

	return
}

func (obj *_StockMvtReasonMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*StockMvtReason, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error

	return
}

func (obj *_StockMvtReasonMgr) GetFromDateUpd(dateUpd time.Time) (results []*StockMvtReason, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd = ?", dateUpd).Find(&results).Error

	return
}

func (obj *_StockMvtReasonMgr) GetBatchFromDateUpd(dateUpds []time.Time) (results []*StockMvtReason, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd IN (?)", dateUpds).Find(&results).Error

	return
}

func (obj *_StockMvtReasonMgr) GetFromDeleted(deleted bool) (results []*StockMvtReason, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("deleted = ?", deleted).Find(&results).Error

	return
}

func (obj *_StockMvtReasonMgr) GetBatchFromDeleted(deleteds []bool) (results []*StockMvtReason, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("deleted IN (?)", deleteds).Find(&results).Error

	return
}


func (obj *_StockMvtReasonMgr) FetchByPrimaryKey(idStockMvtReason uint32) (result StockMvtReason, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_stock_mvt_reason = ?", idStockMvtReason).Find(&result).Error

	return
}
