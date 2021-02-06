package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _CartRuleProductRuleMgr struct {
	*_BaseMgr
}

func CartRuleProductRuleMgr(db *gorm.DB) *_CartRuleProductRuleMgr {
	if db == nil {
		panic(fmt.Errorf("CartRuleProductRuleMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CartRuleProductRuleMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_cart_rule_product_rule"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_CartRuleProductRuleMgr) GetTableName() string {
	return "ps_cart_rule_product_rule"
}

func (obj *_CartRuleProductRuleMgr) Get() (result CartRuleProductRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_CartRuleProductRuleMgr) Gets() (results []*CartRuleProductRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_CartRuleProductRuleMgr) WithIDProductRule(idProductRule uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product_rule"] = idProductRule })
}

func (obj *_CartRuleProductRuleMgr) WithIDProductRuleGroup(idProductRuleGroup uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product_rule_group"] = idProductRuleGroup })
}

func (obj *_CartRuleProductRuleMgr) WithType(_type string) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

func (obj *_CartRuleProductRuleMgr) GetByOption(opts ...Option) (result CartRuleProductRule, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_CartRuleProductRuleMgr) GetByOptions(opts ...Option) (results []*CartRuleProductRule, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_CartRuleProductRuleMgr) GetFromIDProductRule(idProductRule uint32) (result CartRuleProductRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_rule = ?", idProductRule).Find(&result).Error

	return
}

func (obj *_CartRuleProductRuleMgr) GetBatchFromIDProductRule(idProductRules []uint32) (results []*CartRuleProductRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_rule IN (?)", idProductRules).Find(&results).Error

	return
}

func (obj *_CartRuleProductRuleMgr) GetFromIDProductRuleGroup(idProductRuleGroup uint32) (results []*CartRuleProductRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_rule_group = ?", idProductRuleGroup).Find(&results).Error

	return
}

func (obj *_CartRuleProductRuleMgr) GetBatchFromIDProductRuleGroup(idProductRuleGroups []uint32) (results []*CartRuleProductRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_rule_group IN (?)", idProductRuleGroups).Find(&results).Error

	return
}

func (obj *_CartRuleProductRuleMgr) GetFromType(_type string) (results []*CartRuleProductRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type = ?", _type).Find(&results).Error

	return
}

func (obj *_CartRuleProductRuleMgr) GetBatchFromType(_types []string) (results []*CartRuleProductRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type IN (?)", _types).Find(&results).Error

	return
}


func (obj *_CartRuleProductRuleMgr) FetchByPrimaryKey(idProductRule uint32) (result CartRuleProductRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_rule = ?", idProductRule).Find(&result).Error

	return
}
