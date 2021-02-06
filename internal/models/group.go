package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _GroupMgr struct {
	*_BaseMgr
}

// GroupMgr open func
func GroupMgr(db *gorm.DB) *_GroupMgr {
	if db == nil {
		panic(fmt.Errorf("GroupMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_GroupMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_group"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_GroupMgr) GetTableName() string {
	return "ps_group"
}

// Get 获取
func (obj *_GroupMgr) Get() (result Group, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_GroupMgr) Gets() (results []*Group, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDGroup id_group获取
func (obj *_GroupMgr) WithIDGroup(idGroup uint32) Option {
	return optionFunc(func(o *options) { o.query["id_group"] = idGroup })
}

// WithReduction reduction获取
func (obj *_GroupMgr) WithReduction(reduction float64) Option {
	return optionFunc(func(o *options) { o.query["reduction"] = reduction })
}

// WithPriceDisplayMethod price_display_method获取
func (obj *_GroupMgr) WithPriceDisplayMethod(priceDisplayMethod int8) Option {
	return optionFunc(func(o *options) { o.query["price_display_method"] = priceDisplayMethod })
}

// WithShowPrices show_prices获取
func (obj *_GroupMgr) WithShowPrices(showPrices bool) Option {
	return optionFunc(func(o *options) { o.query["show_prices"] = showPrices })
}

// WithDateAdd date_add获取
func (obj *_GroupMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

// WithDateUpd date_upd获取
func (obj *_GroupMgr) WithDateUpd(dateUpd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_upd"] = dateUpd })
}

// GetByOption 功能选项模式获取
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

// GetByOptions 批量功能选项模式获取
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

//////////////////////////enume case ////////////////////////////////////////////

// GetFromIDGroup 通过id_group获取内容
func (obj *_GroupMgr) GetFromIDGroup(idGroup uint32) (result Group, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_group = ?", idGroup).Find(&result).Error

	return
}

// GetBatchFromIDGroup 批量唯一主键查找
func (obj *_GroupMgr) GetBatchFromIDGroup(idGroups []uint32) (results []*Group, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_group IN (?)", idGroups).Find(&results).Error

	return
}

// GetFromReduction 通过reduction获取内容
func (obj *_GroupMgr) GetFromReduction(reduction float64) (results []*Group, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reduction = ?", reduction).Find(&results).Error

	return
}

// GetBatchFromReduction 批量唯一主键查找
func (obj *_GroupMgr) GetBatchFromReduction(reductions []float64) (results []*Group, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reduction IN (?)", reductions).Find(&results).Error

	return
}

// GetFromPriceDisplayMethod 通过price_display_method获取内容
func (obj *_GroupMgr) GetFromPriceDisplayMethod(priceDisplayMethod int8) (results []*Group, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("price_display_method = ?", priceDisplayMethod).Find(&results).Error

	return
}

// GetBatchFromPriceDisplayMethod 批量唯一主键查找
func (obj *_GroupMgr) GetBatchFromPriceDisplayMethod(priceDisplayMethods []int8) (results []*Group, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("price_display_method IN (?)", priceDisplayMethods).Find(&results).Error

	return
}

// GetFromShowPrices 通过show_prices获取内容
func (obj *_GroupMgr) GetFromShowPrices(showPrices bool) (results []*Group, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("show_prices = ?", showPrices).Find(&results).Error

	return
}

// GetBatchFromShowPrices 批量唯一主键查找
func (obj *_GroupMgr) GetBatchFromShowPrices(showPricess []bool) (results []*Group, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("show_prices IN (?)", showPricess).Find(&results).Error

	return
}

// GetFromDateAdd 通过date_add获取内容
func (obj *_GroupMgr) GetFromDateAdd(dateAdd time.Time) (results []*Group, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error

	return
}

// GetBatchFromDateAdd 批量唯一主键查找
func (obj *_GroupMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*Group, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error

	return
}

// GetFromDateUpd 通过date_upd获取内容
func (obj *_GroupMgr) GetFromDateUpd(dateUpd time.Time) (results []*Group, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd = ?", dateUpd).Find(&results).Error

	return
}

// GetBatchFromDateUpd 批量唯一主键查找
func (obj *_GroupMgr) GetBatchFromDateUpd(dateUpds []time.Time) (results []*Group, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd IN (?)", dateUpds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_GroupMgr) FetchByPrimaryKey(idGroup uint32) (result Group, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_group = ?", idGroup).Find(&result).Error

	return
}
