package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _RangeWeightMgr struct {
	*_BaseMgr
}

func RangeWeightMgr(db *gorm.DB) *_RangeWeightMgr {
	if db == nil {
		panic(fmt.Errorf("RangeWeightMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_RangeWeightMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_range_weight"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_RangeWeightMgr) GetTableName() string {
	return "ps_range_weight"
}

func (obj *_RangeWeightMgr) Get() (result RangeWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_RangeWeightMgr) Gets() (results []*RangeWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_RangeWeightMgr) WithIDRangeWeight(idRangeWeight uint32) Option {
	return optionFunc(func(o *options) { o.query["id_range_weight"] = idRangeWeight })
}

func (obj *_RangeWeightMgr) WithIDCarrier(idCarrier uint32) Option {
	return optionFunc(func(o *options) { o.query["id_carrier"] = idCarrier })
}

func (obj *_RangeWeightMgr) WithDelimiter1(delimiter1 float64) Option {
	return optionFunc(func(o *options) { o.query["delimiter1"] = delimiter1 })
}

func (obj *_RangeWeightMgr) WithDelimiter2(delimiter2 float64) Option {
	return optionFunc(func(o *options) { o.query["delimiter2"] = delimiter2 })
}

func (obj *_RangeWeightMgr) GetByOption(opts ...Option) (result RangeWeight, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_RangeWeightMgr) GetByOptions(opts ...Option) (results []*RangeWeight, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_RangeWeightMgr) GetFromIDRangeWeight(idRangeWeight uint32) (result RangeWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_range_weight = ?", idRangeWeight).Find(&result).Error

	return
}

func (obj *_RangeWeightMgr) GetBatchFromIDRangeWeight(idRangeWeights []uint32) (results []*RangeWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_range_weight IN (?)", idRangeWeights).Find(&results).Error

	return
}

func (obj *_RangeWeightMgr) GetFromIDCarrier(idCarrier uint32) (results []*RangeWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_carrier = ?", idCarrier).Find(&results).Error

	return
}

func (obj *_RangeWeightMgr) GetBatchFromIDCarrier(idCarriers []uint32) (results []*RangeWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_carrier IN (?)", idCarriers).Find(&results).Error

	return
}

func (obj *_RangeWeightMgr) GetFromDelimiter1(delimiter1 float64) (results []*RangeWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("delimiter1 = ?", delimiter1).Find(&results).Error

	return
}

func (obj *_RangeWeightMgr) GetBatchFromDelimiter1(delimiter1s []float64) (results []*RangeWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("delimiter1 IN (?)", delimiter1s).Find(&results).Error

	return
}

func (obj *_RangeWeightMgr) GetFromDelimiter2(delimiter2 float64) (results []*RangeWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("delimiter2 = ?", delimiter2).Find(&results).Error

	return
}

func (obj *_RangeWeightMgr) GetBatchFromDelimiter2(delimiter2s []float64) (results []*RangeWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("delimiter2 IN (?)", delimiter2s).Find(&results).Error

	return
}


func (obj *_RangeWeightMgr) FetchByPrimaryKey(idRangeWeight uint32) (result RangeWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_range_weight = ?", idRangeWeight).Find(&result).Error

	return
}

func (obj *_RangeWeightMgr) FetchUniqueIndexByIDCarrier(idCarrier uint32, delimiter1 float64, delimiter2 float64) (result RangeWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_carrier = ? AND delimiter1 = ? AND delimiter2 = ?", idCarrier, delimiter1, delimiter2).Find(&result).Error

	return
}
