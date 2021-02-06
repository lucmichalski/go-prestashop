package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _RangeWeightMgr struct {
	*_BaseMgr
}

// RangeWeightMgr open func
func RangeWeightMgr(db *gorm.DB) *_RangeWeightMgr {
	if db == nil {
		panic(fmt.Errorf("RangeWeightMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_RangeWeightMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_range_weight"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_RangeWeightMgr) GetTableName() string {
	return "eg_range_weight"
}

// Get 获取
func (obj *_RangeWeightMgr) Get() (result RangeWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_RangeWeightMgr) Gets() (results []*RangeWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDRangeWeight id_range_weight获取
func (obj *_RangeWeightMgr) WithIDRangeWeight(idRangeWeight uint32) Option {
	return optionFunc(func(o *options) { o.query["id_range_weight"] = idRangeWeight })
}

// WithIDCarrier id_carrier获取
func (obj *_RangeWeightMgr) WithIDCarrier(idCarrier uint32) Option {
	return optionFunc(func(o *options) { o.query["id_carrier"] = idCarrier })
}

// WithDelimiter1 delimiter1获取
func (obj *_RangeWeightMgr) WithDelimiter1(delimiter1 float64) Option {
	return optionFunc(func(o *options) { o.query["delimiter1"] = delimiter1 })
}

// WithDelimiter2 delimiter2获取
func (obj *_RangeWeightMgr) WithDelimiter2(delimiter2 float64) Option {
	return optionFunc(func(o *options) { o.query["delimiter2"] = delimiter2 })
}

// GetByOption 功能选项模式获取
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

// GetByOptions 批量功能选项模式获取
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

//////////////////////////enume case ////////////////////////////////////////////

// GetFromIDRangeWeight 通过id_range_weight获取内容
func (obj *_RangeWeightMgr) GetFromIDRangeWeight(idRangeWeight uint32) (result RangeWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_range_weight = ?", idRangeWeight).Find(&result).Error

	return
}

// GetBatchFromIDRangeWeight 批量唯一主键查找
func (obj *_RangeWeightMgr) GetBatchFromIDRangeWeight(idRangeWeights []uint32) (results []*RangeWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_range_weight IN (?)", idRangeWeights).Find(&results).Error

	return
}

// GetFromIDCarrier 通过id_carrier获取内容
func (obj *_RangeWeightMgr) GetFromIDCarrier(idCarrier uint32) (results []*RangeWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_carrier = ?", idCarrier).Find(&results).Error

	return
}

// GetBatchFromIDCarrier 批量唯一主键查找
func (obj *_RangeWeightMgr) GetBatchFromIDCarrier(idCarriers []uint32) (results []*RangeWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_carrier IN (?)", idCarriers).Find(&results).Error

	return
}

// GetFromDelimiter1 通过delimiter1获取内容
func (obj *_RangeWeightMgr) GetFromDelimiter1(delimiter1 float64) (results []*RangeWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("delimiter1 = ?", delimiter1).Find(&results).Error

	return
}

// GetBatchFromDelimiter1 批量唯一主键查找
func (obj *_RangeWeightMgr) GetBatchFromDelimiter1(delimiter1s []float64) (results []*RangeWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("delimiter1 IN (?)", delimiter1s).Find(&results).Error

	return
}

// GetFromDelimiter2 通过delimiter2获取内容
func (obj *_RangeWeightMgr) GetFromDelimiter2(delimiter2 float64) (results []*RangeWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("delimiter2 = ?", delimiter2).Find(&results).Error

	return
}

// GetBatchFromDelimiter2 批量唯一主键查找
func (obj *_RangeWeightMgr) GetBatchFromDelimiter2(delimiter2s []float64) (results []*RangeWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("delimiter2 IN (?)", delimiter2s).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_RangeWeightMgr) FetchByPrimaryKey(idRangeWeight uint32) (result RangeWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_range_weight = ?", idRangeWeight).Find(&result).Error

	return
}

// FetchUniqueIndexByIDCarrier primay or index 获取唯一内容
func (obj *_RangeWeightMgr) FetchUniqueIndexByIDCarrier(idCarrier uint32, delimiter1 float64, delimiter2 float64) (result RangeWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_carrier = ? AND delimiter1 = ? AND delimiter2 = ?", idCarrier, delimiter1, delimiter2).Find(&result).Error

	return
}
