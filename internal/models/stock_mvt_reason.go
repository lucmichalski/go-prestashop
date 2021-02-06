package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _StockMvtReasonMgr struct {
	*_BaseMgr
}

// StockMvtReasonMgr open func
func StockMvtReasonMgr(db *gorm.DB) *_StockMvtReasonMgr {
	if db == nil {
		panic(fmt.Errorf("StockMvtReasonMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_StockMvtReasonMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_stock_mvt_reason"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_StockMvtReasonMgr) GetTableName() string {
	return "ps_stock_mvt_reason"
}

// Get 获取
func (obj *_StockMvtReasonMgr) Get() (result StockMvtReason, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_StockMvtReasonMgr) Gets() (results []*StockMvtReason, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDStockMvtReason id_stock_mvt_reason获取
func (obj *_StockMvtReasonMgr) WithIDStockMvtReason(idStockMvtReason uint32) Option {
	return optionFunc(func(o *options) { o.query["id_stock_mvt_reason"] = idStockMvtReason })
}

// WithSign sign获取
func (obj *_StockMvtReasonMgr) WithSign(sign bool) Option {
	return optionFunc(func(o *options) { o.query["sign"] = sign })
}

// WithDateAdd date_add获取
func (obj *_StockMvtReasonMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

// WithDateUpd date_upd获取
func (obj *_StockMvtReasonMgr) WithDateUpd(dateUpd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_upd"] = dateUpd })
}

// WithDeleted deleted获取
func (obj *_StockMvtReasonMgr) WithDeleted(deleted bool) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// GetByOption 功能选项模式获取
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

// GetByOptions 批量功能选项模式获取
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

//////////////////////////enume case ////////////////////////////////////////////

// GetFromIDStockMvtReason 通过id_stock_mvt_reason获取内容
func (obj *_StockMvtReasonMgr) GetFromIDStockMvtReason(idStockMvtReason uint32) (result StockMvtReason, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_stock_mvt_reason = ?", idStockMvtReason).Find(&result).Error

	return
}

// GetBatchFromIDStockMvtReason 批量唯一主键查找
func (obj *_StockMvtReasonMgr) GetBatchFromIDStockMvtReason(idStockMvtReasons []uint32) (results []*StockMvtReason, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_stock_mvt_reason IN (?)", idStockMvtReasons).Find(&results).Error

	return
}

// GetFromSign 通过sign获取内容
func (obj *_StockMvtReasonMgr) GetFromSign(sign bool) (results []*StockMvtReason, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("sign = ?", sign).Find(&results).Error

	return
}

// GetBatchFromSign 批量唯一主键查找
func (obj *_StockMvtReasonMgr) GetBatchFromSign(signs []bool) (results []*StockMvtReason, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("sign IN (?)", signs).Find(&results).Error

	return
}

// GetFromDateAdd 通过date_add获取内容
func (obj *_StockMvtReasonMgr) GetFromDateAdd(dateAdd time.Time) (results []*StockMvtReason, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error

	return
}

// GetBatchFromDateAdd 批量唯一主键查找
func (obj *_StockMvtReasonMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*StockMvtReason, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error

	return
}

// GetFromDateUpd 通过date_upd获取内容
func (obj *_StockMvtReasonMgr) GetFromDateUpd(dateUpd time.Time) (results []*StockMvtReason, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd = ?", dateUpd).Find(&results).Error

	return
}

// GetBatchFromDateUpd 批量唯一主键查找
func (obj *_StockMvtReasonMgr) GetBatchFromDateUpd(dateUpds []time.Time) (results []*StockMvtReason, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd IN (?)", dateUpds).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容
func (obj *_StockMvtReasonMgr) GetFromDeleted(deleted bool) (results []*StockMvtReason, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("deleted = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量唯一主键查找
func (obj *_StockMvtReasonMgr) GetBatchFromDeleted(deleteds []bool) (results []*StockMvtReason, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("deleted IN (?)", deleteds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_StockMvtReasonMgr) FetchByPrimaryKey(idStockMvtReason uint32) (result StockMvtReason, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_stock_mvt_reason = ?", idStockMvtReason).Find(&result).Error

	return
}
