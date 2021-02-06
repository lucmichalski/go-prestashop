package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _SpecificPriceRuleConditionGroupMgr struct {
	*_BaseMgr
}

func SpecificPriceRuleConditionGroupMgr(db *gorm.DB) *_SpecificPriceRuleConditionGroupMgr {
	if db == nil {
		panic(fmt.Errorf("SpecificPriceRuleConditionGroupMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SpecificPriceRuleConditionGroupMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_specific_price_rule_condition_group"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_SpecificPriceRuleConditionGroupMgr) GetTableName() string {
	return "ps_specific_price_rule_condition_group"
}

func (obj *_SpecificPriceRuleConditionGroupMgr) Get() (result SpecificPriceRuleConditionGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_SpecificPriceRuleConditionGroupMgr) Gets() (results []*SpecificPriceRuleConditionGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_SpecificPriceRuleConditionGroupMgr) WithIDSpecificPriceRuleConditionGroup(idSpecificPriceRuleConditionGroup uint32) Option {
	return optionFunc(func(o *options) {
		o.query["id_specific_price_rule_condition_group"] = idSpecificPriceRuleConditionGroup
	})
}

func (obj *_SpecificPriceRuleConditionGroupMgr) WithIDSpecificPriceRule(idSpecificPriceRule uint32) Option {
	return optionFunc(func(o *options) { o.query["id_specific_price_rule"] = idSpecificPriceRule })
}

func (obj *_SpecificPriceRuleConditionGroupMgr) GetByOption(opts ...Option) (result SpecificPriceRuleConditionGroup, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_SpecificPriceRuleConditionGroupMgr) GetByOptions(opts ...Option) (results []*SpecificPriceRuleConditionGroup, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_SpecificPriceRuleConditionGroupMgr) GetFromIDSpecificPriceRuleConditionGroup(idSpecificPriceRuleConditionGroup uint32) (results []*SpecificPriceRuleConditionGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_specific_price_rule_condition_group = ?", idSpecificPriceRuleConditionGroup).Find(&results).Error

	return
}

func (obj *_SpecificPriceRuleConditionGroupMgr) GetBatchFromIDSpecificPriceRuleConditionGroup(idSpecificPriceRuleConditionGroups []uint32) (results []*SpecificPriceRuleConditionGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_specific_price_rule_condition_group IN (?)", idSpecificPriceRuleConditionGroups).Find(&results).Error

	return
}

func (obj *_SpecificPriceRuleConditionGroupMgr) GetFromIDSpecificPriceRule(idSpecificPriceRule uint32) (results []*SpecificPriceRuleConditionGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_specific_price_rule = ?", idSpecificPriceRule).Find(&results).Error

	return
}

func (obj *_SpecificPriceRuleConditionGroupMgr) GetBatchFromIDSpecificPriceRule(idSpecificPriceRules []uint32) (results []*SpecificPriceRuleConditionGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_specific_price_rule IN (?)", idSpecificPriceRules).Find(&results).Error

	return
}


func (obj *_SpecificPriceRuleConditionGroupMgr) FetchByPrimaryKey(idSpecificPriceRuleConditionGroup uint32, idSpecificPriceRule uint32) (result SpecificPriceRuleConditionGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_specific_price_rule_condition_group = ? AND id_specific_price_rule = ?", idSpecificPriceRuleConditionGroup, idSpecificPriceRule).Find(&result).Error

	return
}
