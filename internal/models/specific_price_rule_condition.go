package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _SpecificPriceRuleConditionMgr struct {
	*_BaseMgr
}

// SpecificPriceRuleConditionMgr open func
func SpecificPriceRuleConditionMgr(db *gorm.DB) *_SpecificPriceRuleConditionMgr {
	if db == nil {
		panic(fmt.Errorf("SpecificPriceRuleConditionMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SpecificPriceRuleConditionMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_specific_price_rule_condition"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SpecificPriceRuleConditionMgr) GetTableName() string {
	return "eg_specific_price_rule_condition"
}

// Get 获取
func (obj *_SpecificPriceRuleConditionMgr) Get() (result SpecificPriceRuleCondition, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SpecificPriceRuleConditionMgr) Gets() (results []*SpecificPriceRuleCondition, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDSpecificPriceRuleCondition id_specific_price_rule_condition获取
func (obj *_SpecificPriceRuleConditionMgr) WithIDSpecificPriceRuleCondition(idSpecificPriceRuleCondition uint32) Option {
	return optionFunc(func(o *options) { o.query["id_specific_price_rule_condition"] = idSpecificPriceRuleCondition })
}

// WithIDSpecificPriceRuleConditionGroup id_specific_price_rule_condition_group获取
func (obj *_SpecificPriceRuleConditionMgr) WithIDSpecificPriceRuleConditionGroup(idSpecificPriceRuleConditionGroup uint32) Option {
	return optionFunc(func(o *options) {
		o.query["id_specific_price_rule_condition_group"] = idSpecificPriceRuleConditionGroup
	})
}

// WithType type获取
func (obj *_SpecificPriceRuleConditionMgr) WithType(_type string) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithValue value获取
func (obj *_SpecificPriceRuleConditionMgr) WithValue(value string) Option {
	return optionFunc(func(o *options) { o.query["value"] = value })
}

// GetByOption 功能选项模式获取
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

// GetByOptions 批量功能选项模式获取
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

//////////////////////////enume case ////////////////////////////////////////////

// GetFromIDSpecificPriceRuleCondition 通过id_specific_price_rule_condition获取内容
func (obj *_SpecificPriceRuleConditionMgr) GetFromIDSpecificPriceRuleCondition(idSpecificPriceRuleCondition uint32) (result SpecificPriceRuleCondition, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_specific_price_rule_condition = ?", idSpecificPriceRuleCondition).Find(&result).Error

	return
}

// GetBatchFromIDSpecificPriceRuleCondition 批量唯一主键查找
func (obj *_SpecificPriceRuleConditionMgr) GetBatchFromIDSpecificPriceRuleCondition(idSpecificPriceRuleConditions []uint32) (results []*SpecificPriceRuleCondition, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_specific_price_rule_condition IN (?)", idSpecificPriceRuleConditions).Find(&results).Error

	return
}

// GetFromIDSpecificPriceRuleConditionGroup 通过id_specific_price_rule_condition_group获取内容
func (obj *_SpecificPriceRuleConditionMgr) GetFromIDSpecificPriceRuleConditionGroup(idSpecificPriceRuleConditionGroup uint32) (results []*SpecificPriceRuleCondition, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_specific_price_rule_condition_group = ?", idSpecificPriceRuleConditionGroup).Find(&results).Error

	return
}

// GetBatchFromIDSpecificPriceRuleConditionGroup 批量唯一主键查找
func (obj *_SpecificPriceRuleConditionMgr) GetBatchFromIDSpecificPriceRuleConditionGroup(idSpecificPriceRuleConditionGroups []uint32) (results []*SpecificPriceRuleCondition, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_specific_price_rule_condition_group IN (?)", idSpecificPriceRuleConditionGroups).Find(&results).Error

	return
}

// GetFromType 通过type获取内容
func (obj *_SpecificPriceRuleConditionMgr) GetFromType(_type string) (results []*SpecificPriceRuleCondition, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量唯一主键查找
func (obj *_SpecificPriceRuleConditionMgr) GetBatchFromType(_types []string) (results []*SpecificPriceRuleCondition, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type IN (?)", _types).Find(&results).Error

	return
}

// GetFromValue 通过value获取内容
func (obj *_SpecificPriceRuleConditionMgr) GetFromValue(value string) (results []*SpecificPriceRuleCondition, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("value = ?", value).Find(&results).Error

	return
}

// GetBatchFromValue 批量唯一主键查找
func (obj *_SpecificPriceRuleConditionMgr) GetBatchFromValue(values []string) (results []*SpecificPriceRuleCondition, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("value IN (?)", values).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_SpecificPriceRuleConditionMgr) FetchByPrimaryKey(idSpecificPriceRuleCondition uint32) (result SpecificPriceRuleCondition, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_specific_price_rule_condition = ?", idSpecificPriceRuleCondition).Find(&result).Error

	return
}

// FetchIndexByIDSpecificPriceRuleConditionGroup  获取多个内容
func (obj *_SpecificPriceRuleConditionMgr) FetchIndexByIDSpecificPriceRuleConditionGroup(idSpecificPriceRuleConditionGroup uint32) (results []*SpecificPriceRuleCondition, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_specific_price_rule_condition_group = ?", idSpecificPriceRuleConditionGroup).Find(&results).Error

	return
}
