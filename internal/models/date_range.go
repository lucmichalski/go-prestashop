package models

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _DateRangeMgr struct {
	*_BaseMgr
}

func DateRangeMgr(db *gorm.DB) *_DateRangeMgr {
	if db == nil {
		panic(fmt.Errorf("DateRangeMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_DateRangeMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_date_range"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_DateRangeMgr) GetTableName() string {
	return "ps_date_range"
}

func (obj *_DateRangeMgr) Get() (result DateRange, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_DateRangeMgr) Gets() (results []*DateRange, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

func (obj *_DateRangeMgr) WithIDDateRange(idDateRange uint32) Option {
	return optionFunc(func(o *options) { o.query["id_date_range"] = idDateRange })
}

func (obj *_DateRangeMgr) WithTimeStart(timeStart time.Time) Option {
	return optionFunc(func(o *options) { o.query["time_start"] = timeStart })
}

func (obj *_DateRangeMgr) WithTimeEnd(timeEnd time.Time) Option {
	return optionFunc(func(o *options) { o.query["time_end"] = timeEnd })
}

func (obj *_DateRangeMgr) GetByOption(opts ...Option) (result DateRange, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_DateRangeMgr) GetByOptions(opts ...Option) (results []*DateRange, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

func (obj *_DateRangeMgr) GetFromIDDateRange(idDateRange uint32) (result DateRange, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_date_range = ?", idDateRange).Find(&result).Error

	return
}

func (obj *_DateRangeMgr) GetBatchFromIDDateRange(idDateRanges []uint32) (results []*DateRange, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_date_range IN (?)", idDateRanges).Find(&results).Error

	return
}

func (obj *_DateRangeMgr) GetFromTimeStart(timeStart time.Time) (results []*DateRange, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("time_start = ?", timeStart).Find(&results).Error

	return
}

func (obj *_DateRangeMgr) GetBatchFromTimeStart(timeStarts []time.Time) (results []*DateRange, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("time_start IN (?)", timeStarts).Find(&results).Error

	return
}

func (obj *_DateRangeMgr) GetFromTimeEnd(timeEnd time.Time) (results []*DateRange, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("time_end = ?", timeEnd).Find(&results).Error

	return
}

func (obj *_DateRangeMgr) GetBatchFromTimeEnd(timeEnds []time.Time) (results []*DateRange, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("time_end IN (?)", timeEnds).Find(&results).Error

	return
}

func (obj *_DateRangeMgr) FetchByPrimaryKey(idDateRange uint32) (result DateRange, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_date_range = ?", idDateRange).Find(&result).Error

	return
}
