package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _CartRuleGroupMgr struct {
	*_BaseMgr
}

// CartRuleGroupMgr open func
func CartRuleGroupMgr(db *gorm.DB) *_CartRuleGroupMgr {
	if db == nil {
		panic(fmt.Errorf("CartRuleGroupMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CartRuleGroupMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_cart_rule_group"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CartRuleGroupMgr) GetTableName() string {
	return "ps_cart_rule_group"
}

// Get 获取
func (obj *_CartRuleGroupMgr) Get() (result CartRuleGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CartRuleGroupMgr) Gets() (results []*CartRuleGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDCartRule id_cart_rule获取
func (obj *_CartRuleGroupMgr) WithIDCartRule(idCartRule uint32) Option {
	return optionFunc(func(o *options) { o.query["id_cart_rule"] = idCartRule })
}

// WithIDGroup id_group获取
func (obj *_CartRuleGroupMgr) WithIDGroup(idGroup uint32) Option {
	return optionFunc(func(o *options) { o.query["id_group"] = idGroup })
}

// GetByOption 功能选项模式获取
func (obj *_CartRuleGroupMgr) GetByOption(opts ...Option) (result CartRuleGroup, err error) {
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
func (obj *_CartRuleGroupMgr) GetByOptions(opts ...Option) (results []*CartRuleGroup, err error) {
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

// GetFromIDCartRule 通过id_cart_rule获取内容
func (obj *_CartRuleGroupMgr) GetFromIDCartRule(idCartRule uint32) (results []*CartRuleGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart_rule = ?", idCartRule).Find(&results).Error

	return
}

// GetBatchFromIDCartRule 批量唯一主键查找
func (obj *_CartRuleGroupMgr) GetBatchFromIDCartRule(idCartRules []uint32) (results []*CartRuleGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart_rule IN (?)", idCartRules).Find(&results).Error

	return
}

// GetFromIDGroup 通过id_group获取内容
func (obj *_CartRuleGroupMgr) GetFromIDGroup(idGroup uint32) (results []*CartRuleGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_group = ?", idGroup).Find(&results).Error

	return
}

// GetBatchFromIDGroup 批量唯一主键查找
func (obj *_CartRuleGroupMgr) GetBatchFromIDGroup(idGroups []uint32) (results []*CartRuleGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_group IN (?)", idGroups).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_CartRuleGroupMgr) FetchByPrimaryKey(idCartRule uint32, idGroup uint32) (result CartRuleGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart_rule = ? AND id_group = ?", idCartRule, idGroup).Find(&result).Error

	return
}
