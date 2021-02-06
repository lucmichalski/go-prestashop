package models

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _GroupMgr struct {
	*_BaseMgr
}

func GroupMgr(db *gorm.DB) *_GroupMgr {
	if db == nil {
		panic(fmt.Errorf("GroupMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_GroupMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_group"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_GroupMgr) GetTableName() string {
	return "ps_group"
}

func (obj *_GroupMgr) Get() (result Group, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_GroupMgr) Gets() (results []*Group, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_GroupMgr) WithIDGroup(idGroup uint32) Option {
	return optionFunc(func(o *options) { o.query["id_group"] = idGroup })
}

func (obj *_GroupMgr) WithReduction(reduction float64) Option {
	return optionFunc(func(o *options) { o.query["reduction"] = reduction })
}

func (obj *_GroupMgr) WithPriceDisplayMethod(priceDisplayMethod int8) Option {
	return optionFunc(func(o *options) { o.query["price_display_method"] = priceDisplayMethod })
}

func (obj *_GroupMgr) WithShowPrices(showPrices bool) Option {
	return optionFunc(func(o *options) { o.query["show_prices"] = showPrices })
}

func (obj *_GroupMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

func (obj *_GroupMgr) WithDateUpd(dateUpd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_upd"] = dateUpd })
}

func (obj *_GroupMgr) GetByOption(opts ...Option) (result Group, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_GroupMgr) GetByOptions(opts ...Option) (results []*Group, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_GroupMgr) GetFromIDGroup(idGroup uint32) (result Group, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_group = ?", idGroup).Find(&result).Error

	return
}

func (obj *_GroupMgr) GetBatchFromIDGroup(idGroups []uint32) (results []*Group, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_group IN (?)", idGroups).Find(&results).Error

	return
}

func (obj *_GroupMgr) GetFromReduction(reduction float64) (results []*Group, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reduction = ?", reduction).Find(&results).Error

	return
}

func (obj *_GroupMgr) GetBatchFromReduction(reductions []float64) (results []*Group, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reduction IN (?)", reductions).Find(&results).Error

	return
}

func (obj *_GroupMgr) GetFromPriceDisplayMethod(priceDisplayMethod int8) (results []*Group, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("price_display_method = ?", priceDisplayMethod).Find(&results).Error

	return
}

func (obj *_GroupMgr) GetBatchFromPriceDisplayMethod(priceDisplayMethods []int8) (results []*Group, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("price_display_method IN (?)", priceDisplayMethods).Find(&results).Error

	return
}

func (obj *_GroupMgr) GetFromShowPrices(showPrices bool) (results []*Group, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("show_prices = ?", showPrices).Find(&results).Error

	return
}

func (obj *_GroupMgr) GetBatchFromShowPrices(showPricess []bool) (results []*Group, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("show_prices IN (?)", showPricess).Find(&results).Error

	return
}

func (obj *_GroupMgr) GetFromDateAdd(dateAdd time.Time) (results []*Group, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error

	return
}

func (obj *_GroupMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*Group, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error

	return
}

func (obj *_GroupMgr) GetFromDateUpd(dateUpd time.Time) (results []*Group, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd = ?", dateUpd).Find(&results).Error

	return
}

func (obj *_GroupMgr) GetBatchFromDateUpd(dateUpds []time.Time) (results []*Group, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd IN (?)", dateUpds).Find(&results).Error

	return
}


func (obj *_GroupMgr) FetchByPrimaryKey(idGroup uint32) (result Group, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_group = ?", idGroup).Find(&result).Error

	return
}
