package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _SpecificPriceRuleConditionGroupMgr struct {
	*_BaseMgr
}

// SpecificPriceRuleConditionGroupMgr open func
func SpecificPriceRuleConditionGroupMgr(db *gorm.DB) *_SpecificPriceRuleConditionGroupMgr {
	if db == nil {
		panic(fmt.Errorf("SpecificPriceRuleConditionGroupMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SpecificPriceRuleConditionGroupMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_specific_price_rule_condition_group"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SpecificPriceRuleConditionGroupMgr) GetTableName() string {
	return "ps_specific_price_rule_condition_group"
}

// Get 获取
func (obj *_SpecificPriceRuleConditionGroupMgr) Get() (result SpecificPriceRuleConditionGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SpecificPriceRuleConditionGroupMgr) Gets() (results []*SpecificPriceRuleConditionGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDSpecificPriceRuleConditionGroup id_specific_price_rule_condition_group获取
func (obj *_SpecificPriceRuleConditionGroupMgr) WithIDSpecificPriceRuleConditionGroup(idSpecificPriceRuleConditionGroup uint32) Option {
	return optionFunc(func(o *options) {
		o.query["id_specific_price_rule_condition_group"] = idSpecificPriceRuleConditionGroup
	})
}

// WithIDSpecificPriceRule id_specific_price_rule获取
func (obj *_SpecificPriceRuleConditionGroupMgr) WithIDSpecificPriceRule(idSpecificPriceRule uint32) Option {
	return optionFunc(func(o *options) { o.query["id_specific_price_rule"] = idSpecificPriceRule })
}

// GetByOption 功能选项模式获取
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

// GetByOptions 批量功能选项模式获取
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

//////////////////////////enume case ////////////////////////////////////////////

// GetFromIDSpecificPriceRuleConditionGroup 通过id_specific_price_rule_condition_group获取内容
func (obj *_SpecificPriceRuleConditionGroupMgr) GetFromIDSpecificPriceRuleConditionGroup(idSpecificPriceRuleConditionGroup uint32) (results []*SpecificPriceRuleConditionGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_specific_price_rule_condition_group = ?", idSpecificPriceRuleConditionGroup).Find(&results).Error

	return
}

// GetBatchFromIDSpecificPriceRuleConditionGroup 批量唯一主键查找
func (obj *_SpecificPriceRuleConditionGroupMgr) GetBatchFromIDSpecificPriceRuleConditionGroup(idSpecificPriceRuleConditionGroups []uint32) (results []*SpecificPriceRuleConditionGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_specific_price_rule_condition_group IN (?)", idSpecificPriceRuleConditionGroups).Find(&results).Error

	return
}

// GetFromIDSpecificPriceRule 通过id_specific_price_rule获取内容
func (obj *_SpecificPriceRuleConditionGroupMgr) GetFromIDSpecificPriceRule(idSpecificPriceRule uint32) (results []*SpecificPriceRuleConditionGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_specific_price_rule = ?", idSpecificPriceRule).Find(&results).Error

	return
}

// GetBatchFromIDSpecificPriceRule 批量唯一主键查找
func (obj *_SpecificPriceRuleConditionGroupMgr) GetBatchFromIDSpecificPriceRule(idSpecificPriceRules []uint32) (results []*SpecificPriceRuleConditionGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_specific_price_rule IN (?)", idSpecificPriceRules).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_SpecificPriceRuleConditionGroupMgr) FetchByPrimaryKey(idSpecificPriceRuleConditionGroup uint32, idSpecificPriceRule uint32) (result SpecificPriceRuleConditionGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_specific_price_rule_condition_group = ? AND id_specific_price_rule = ?", idSpecificPriceRuleConditionGroup, idSpecificPriceRule).Find(&result).Error

	return
}
