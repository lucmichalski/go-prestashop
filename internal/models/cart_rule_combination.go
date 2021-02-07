package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _CartRuleCombinationMgr struct {
	*_BaseMgr
}

func CartRuleCombinationMgr(db *gorm.DB) *_CartRuleCombinationMgr {
	if db == nil {
		panic(fmt.Errorf("CartRuleCombinationMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CartRuleCombinationMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_cart_rule_combination"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_CartRuleCombinationMgr) GetTableName() string {
	return "ps_cart_rule_combination"
}

func (obj *_CartRuleCombinationMgr) Get() (result CartRuleCombination, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_CartRuleCombinationMgr) Gets() (results []*CartRuleCombination, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

func (obj *_CartRuleCombinationMgr) WithIDCartRule1(idCartRule1 uint32) Option {
	return optionFunc(func(o *options) { o.query["id_cart_rule_1"] = idCartRule1 })
}

func (obj *_CartRuleCombinationMgr) WithIDCartRule2(idCartRule2 uint32) Option {
	return optionFunc(func(o *options) { o.query["id_cart_rule_2"] = idCartRule2 })
}

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

func (obj *_CartRuleCombinationMgr) GetFromIDCartRule1(idCartRule1 uint32) (results []*CartRuleCombination, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart_rule_1 = ?", idCartRule1).Find(&results).Error

	return
}

func (obj *_CartRuleCombinationMgr) GetBatchFromIDCartRule1(idCartRule1s []uint32) (results []*CartRuleCombination, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart_rule_1 IN (?)", idCartRule1s).Find(&results).Error

	return
}

func (obj *_CartRuleCombinationMgr) GetFromIDCartRule2(idCartRule2 uint32) (results []*CartRuleCombination, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart_rule_2 = ?", idCartRule2).Find(&results).Error

	return
}

func (obj *_CartRuleCombinationMgr) GetBatchFromIDCartRule2(idCartRule2s []uint32) (results []*CartRuleCombination, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart_rule_2 IN (?)", idCartRule2s).Find(&results).Error

	return
}

func (obj *_CartRuleCombinationMgr) FetchByPrimaryKey(idCartRule1 uint32, idCartRule2 uint32) (result CartRuleCombination, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart_rule_1 = ? AND id_cart_rule_2 = ?", idCartRule1, idCartRule2).Find(&result).Error

	return
}

func (obj *_CartRuleCombinationMgr) FetchIndexByIDCartRule1(idCartRule1 uint32) (results []*CartRuleCombination, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart_rule_1 = ?", idCartRule1).Find(&results).Error

	return
}

func (obj *_CartRuleCombinationMgr) FetchIndexByIDCartRule2(idCartRule2 uint32) (results []*CartRuleCombination, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart_rule_2 = ?", idCartRule2).Find(&results).Error

	return
}
