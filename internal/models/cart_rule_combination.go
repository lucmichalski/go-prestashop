package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _CartRuleCombinationMgr struct {
	*_BaseMgr
}

// CartRuleCombinationMgr open func
func CartRuleCombinationMgr(db *gorm.DB) *_CartRuleCombinationMgr {
	if db == nil {
		panic(fmt.Errorf("CartRuleCombinationMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CartRuleCombinationMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_cart_rule_combination"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CartRuleCombinationMgr) GetTableName() string {
	return "ps_cart_rule_combination"
}

// Get 获取
func (obj *_CartRuleCombinationMgr) Get() (result CartRuleCombination, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CartRuleCombinationMgr) Gets() (results []*CartRuleCombination, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDCartRule1 id_cart_rule_1获取
func (obj *_CartRuleCombinationMgr) WithIDCartRule1(idCartRule1 uint32) Option {
	return optionFunc(func(o *options) { o.query["id_cart_rule_1"] = idCartRule1 })
}

// WithIDCartRule2 id_cart_rule_2获取
func (obj *_CartRuleCombinationMgr) WithIDCartRule2(idCartRule2 uint32) Option {
	return optionFunc(func(o *options) { o.query["id_cart_rule_2"] = idCartRule2 })
}

// GetByOption 功能选项模式获取
func (obj *_CartRuleCombinationMgr) GetByOption(opts ...Option) (result CartRuleCombination, err error) {
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
func (obj *_CartRuleCombinationMgr) GetByOptions(opts ...Option) (results []*CartRuleCombination, err error) {
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

// GetFromIDCartRule1 通过id_cart_rule_1获取内容
func (obj *_CartRuleCombinationMgr) GetFromIDCartRule1(idCartRule1 uint32) (results []*CartRuleCombination, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart_rule_1 = ?", idCartRule1).Find(&results).Error

	return
}

// GetBatchFromIDCartRule1 批量唯一主键查找
func (obj *_CartRuleCombinationMgr) GetBatchFromIDCartRule1(idCartRule1s []uint32) (results []*CartRuleCombination, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart_rule_1 IN (?)", idCartRule1s).Find(&results).Error

	return
}

// GetFromIDCartRule2 通过id_cart_rule_2获取内容
func (obj *_CartRuleCombinationMgr) GetFromIDCartRule2(idCartRule2 uint32) (results []*CartRuleCombination, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart_rule_2 = ?", idCartRule2).Find(&results).Error

	return
}

// GetBatchFromIDCartRule2 批量唯一主键查找
func (obj *_CartRuleCombinationMgr) GetBatchFromIDCartRule2(idCartRule2s []uint32) (results []*CartRuleCombination, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart_rule_2 IN (?)", idCartRule2s).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_CartRuleCombinationMgr) FetchByPrimaryKey(idCartRule1 uint32, idCartRule2 uint32) (result CartRuleCombination, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart_rule_1 = ? AND id_cart_rule_2 = ?", idCartRule1, idCartRule2).Find(&result).Error

	return
}

// FetchIndexByIDCartRule1  获取多个内容
func (obj *_CartRuleCombinationMgr) FetchIndexByIDCartRule1(idCartRule1 uint32) (results []*CartRuleCombination, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart_rule_1 = ?", idCartRule1).Find(&results).Error

	return
}

// FetchIndexByIDCartRule2  获取多个内容
func (obj *_CartRuleCombinationMgr) FetchIndexByIDCartRule2(idCartRule2 uint32) (results []*CartRuleCombination, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart_rule_2 = ?", idCartRule2).Find(&results).Error

	return
}
