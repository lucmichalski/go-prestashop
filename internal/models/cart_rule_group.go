package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _CartRuleGroupMgr struct {
	*_BaseMgr
}

func CartRuleGroupMgr(db *gorm.DB) *_CartRuleGroupMgr {
	if db == nil {
		panic(fmt.Errorf("CartRuleGroupMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CartRuleGroupMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_cart_rule_group"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_CartRuleGroupMgr) GetTableName() string {
	return "ps_cart_rule_group"
}

func (obj *_CartRuleGroupMgr) Get() (result CartRuleGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_CartRuleGroupMgr) Gets() (results []*CartRuleGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

func (obj *_CartRuleGroupMgr) WithIDCartRule(idCartRule uint32) Option {
	return optionFunc(func(o *options) { o.query["id_cart_rule"] = idCartRule })
}

func (obj *_CartRuleGroupMgr) WithIDGroup(idGroup uint32) Option {
	return optionFunc(func(o *options) { o.query["id_group"] = idGroup })
}

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

func (obj *_CartRuleGroupMgr) GetFromIDCartRule(idCartRule uint32) (results []*CartRuleGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart_rule = ?", idCartRule).Find(&results).Error

	return
}

func (obj *_CartRuleGroupMgr) GetBatchFromIDCartRule(idCartRules []uint32) (results []*CartRuleGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart_rule IN (?)", idCartRules).Find(&results).Error

	return
}

func (obj *_CartRuleGroupMgr) GetFromIDGroup(idGroup uint32) (results []*CartRuleGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_group = ?", idGroup).Find(&results).Error

	return
}

func (obj *_CartRuleGroupMgr) GetBatchFromIDGroup(idGroups []uint32) (results []*CartRuleGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_group IN (?)", idGroups).Find(&results).Error

	return
}

func (obj *_CartRuleGroupMgr) FetchByPrimaryKey(idCartRule uint32, idGroup uint32) (result CartRuleGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart_rule = ? AND id_group = ?", idCartRule, idGroup).Find(&result).Error

	return
}
