package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _RangePriceMgr struct {
	*_BaseMgr
}

func RangePriceMgr(db *gorm.DB) *_RangePriceMgr {
	if db == nil {
		panic(fmt.Errorf("RangePriceMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_RangePriceMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_range_price"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_RangePriceMgr) GetTableName() string {
	return "ps_range_price"
}

func (obj *_RangePriceMgr) Get() (result RangePrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_RangePriceMgr) Gets() (results []*RangePrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_RangePriceMgr) WithIDRangePrice(idRangePrice uint32) Option {
	return optionFunc(func(o *options) { o.query["id_range_price"] = idRangePrice })
}

func (obj *_RangePriceMgr) WithIDCarrier(idCarrier uint32) Option {
	return optionFunc(func(o *options) { o.query["id_carrier"] = idCarrier })
}

func (obj *_RangePriceMgr) WithDelimiter1(delimiter1 float64) Option {
	return optionFunc(func(o *options) { o.query["delimiter1"] = delimiter1 })
}

func (obj *_RangePriceMgr) WithDelimiter2(delimiter2 float64) Option {
	return optionFunc(func(o *options) { o.query["delimiter2"] = delimiter2 })
}

func (obj *_RangePriceMgr) GetByOption(opts ...Option) (result RangePrice, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_RangePriceMgr) GetByOptions(opts ...Option) (results []*RangePrice, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_RangePriceMgr) GetFromIDRangePrice(idRangePrice uint32) (result RangePrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_range_price = ?", idRangePrice).Find(&result).Error

	return
}

func (obj *_RangePriceMgr) GetBatchFromIDRangePrice(idRangePrices []uint32) (results []*RangePrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_range_price IN (?)", idRangePrices).Find(&results).Error

	return
}

func (obj *_RangePriceMgr) GetFromIDCarrier(idCarrier uint32) (results []*RangePrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_carrier = ?", idCarrier).Find(&results).Error

	return
}

func (obj *_RangePriceMgr) GetBatchFromIDCarrier(idCarriers []uint32) (results []*RangePrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_carrier IN (?)", idCarriers).Find(&results).Error

	return
}

func (obj *_RangePriceMgr) GetFromDelimiter1(delimiter1 float64) (results []*RangePrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("delimiter1 = ?", delimiter1).Find(&results).Error

	return
}

func (obj *_RangePriceMgr) GetBatchFromDelimiter1(delimiter1s []float64) (results []*RangePrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("delimiter1 IN (?)", delimiter1s).Find(&results).Error

	return
}

func (obj *_RangePriceMgr) GetFromDelimiter2(delimiter2 float64) (results []*RangePrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("delimiter2 = ?", delimiter2).Find(&results).Error

	return
}

func (obj *_RangePriceMgr) GetBatchFromDelimiter2(delimiter2s []float64) (results []*RangePrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("delimiter2 IN (?)", delimiter2s).Find(&results).Error

	return
}


func (obj *_RangePriceMgr) FetchByPrimaryKey(idRangePrice uint32) (result RangePrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_range_price = ?", idRangePrice).Find(&result).Error

	return
}

func (obj *_RangePriceMgr) FetchUniqueIndexByIDCarrier(idCarrier uint32, delimiter1 float64, delimiter2 float64) (result RangePrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_carrier = ? AND delimiter1 = ? AND delimiter2 = ?", idCarrier, delimiter1, delimiter2).Find(&result).Error

	return
}
