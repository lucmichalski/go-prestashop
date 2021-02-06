package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _SpecificPriceRuleConditionMgr struct {
	*_BaseMgr
}

func SpecificPriceRuleConditionMgr(db *gorm.DB) *_SpecificPriceRuleConditionMgr {
	if db == nil {
		panic(fmt.Errorf("SpecificPriceRuleConditionMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SpecificPriceRuleConditionMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_specific_price_rule_condition"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_SpecificPriceRuleConditionMgr) GetTableName() string {
	return "ps_specific_price_rule_condition"
}

func (obj *_SpecificPriceRuleConditionMgr) Get() (result SpecificPriceRuleCondition, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_SpecificPriceRuleConditionMgr) Gets() (results []*SpecificPriceRuleCondition, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_SpecificPriceRuleConditionMgr) WithIDSpecificPriceRuleCondition(idSpecificPriceRuleCondition uint32) Option {
	return optionFunc(func(o *options) { o.query["id_specific_price_rule_condition"] = idSpecificPriceRuleCondition })
}

func (obj *_SpecificPriceRuleConditionMgr) WithIDSpecificPriceRuleConditionGroup(idSpecificPriceRuleConditionGroup uint32) Option {
	return optionFunc(func(o *options) {
		o.query["id_specific_price_rule_condition_group"] = idSpecificPriceRuleConditionGroup
	})
}

func (obj *_SpecificPriceRuleConditionMgr) WithType(_type string) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

func (obj *_SpecificPriceRuleConditionMgr) WithValue(value string) Option {
	return optionFunc(func(o *options) { o.query["value"] = value })
}

func (obj *_SpecificPriceRuleConditionMgr) GetByOption(opts ...Option) (result SpecificPriceRuleCondition, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_SpecificPriceRuleConditionMgr) GetByOptions(opts ...Option) (results []*SpecificPriceRuleCondition, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_SpecificPriceRuleConditionMgr) GetFromIDSpecificPriceRuleCondition(idSpecificPriceRuleCondition uint32) (result SpecificPriceRuleCondition, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_specific_price_rule_condition = ?", idSpecificPriceRuleCondition).Find(&result).Error

	return
}

func (obj *_SpecificPriceRuleConditionMgr) GetBatchFromIDSpecificPriceRuleCondition(idSpecificPriceRuleConditions []uint32) (results []*SpecificPriceRuleCondition, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_specific_price_rule_condition IN (?)", idSpecificPriceRuleConditions).Find(&results).Error

	return
}

func (obj *_SpecificPriceRuleConditionMgr) GetFromIDSpecificPriceRuleConditionGroup(idSpecificPriceRuleConditionGroup uint32) (results []*SpecificPriceRuleCondition, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_specific_price_rule_condition_group = ?", idSpecificPriceRuleConditionGroup).Find(&results).Error

	return
}

func (obj *_SpecificPriceRuleConditionMgr) GetBatchFromIDSpecificPriceRuleConditionGroup(idSpecificPriceRuleConditionGroups []uint32) (results []*SpecificPriceRuleCondition, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_specific_price_rule_condition_group IN (?)", idSpecificPriceRuleConditionGroups).Find(&results).Error

	return
}

func (obj *_SpecificPriceRuleConditionMgr) GetFromType(_type string) (results []*SpecificPriceRuleCondition, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type = ?", _type).Find(&results).Error

	return
}

func (obj *_SpecificPriceRuleConditionMgr) GetBatchFromType(_types []string) (results []*SpecificPriceRuleCondition, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type IN (?)", _types).Find(&results).Error

	return
}

func (obj *_SpecificPriceRuleConditionMgr) GetFromValue(value string) (results []*SpecificPriceRuleCondition, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("value = ?", value).Find(&results).Error

	return
}

func (obj *_SpecificPriceRuleConditionMgr) GetBatchFromValue(values []string) (results []*SpecificPriceRuleCondition, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("value IN (?)", values).Find(&results).Error

	return
}


func (obj *_SpecificPriceRuleConditionMgr) FetchByPrimaryKey(idSpecificPriceRuleCondition uint32) (result SpecificPriceRuleCondition, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_specific_price_rule_condition = ?", idSpecificPriceRuleCondition).Find(&result).Error

	return
}

func (obj *_SpecificPriceRuleConditionMgr) FetchIndexByIDSpecificPriceRuleConditionGroup(idSpecificPriceRuleConditionGroup uint32) (results []*SpecificPriceRuleCondition, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_specific_price_rule_condition_group = ?", idSpecificPriceRuleConditionGroup).Find(&results).Error

	return
}
