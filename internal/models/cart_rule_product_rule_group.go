package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _CartRuleProductRuleGroupMgr struct {
	*_BaseMgr
}

// CartRuleProductRuleGroupMgr open func
func CartRuleProductRuleGroupMgr(db *gorm.DB) *_CartRuleProductRuleGroupMgr {
	if db == nil {
		panic(fmt.Errorf("CartRuleProductRuleGroupMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CartRuleProductRuleGroupMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_cart_rule_product_rule_group"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CartRuleProductRuleGroupMgr) GetTableName() string {
	return "ps_cart_rule_product_rule_group"
}

// Get 获取
func (obj *_CartRuleProductRuleGroupMgr) Get() (result CartRuleProductRuleGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CartRuleProductRuleGroupMgr) Gets() (results []*CartRuleProductRuleGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDProductRuleGroup id_product_rule_group获取
func (obj *_CartRuleProductRuleGroupMgr) WithIDProductRuleGroup(idProductRuleGroup uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product_rule_group"] = idProductRuleGroup })
}

// WithIDCartRule id_cart_rule获取
func (obj *_CartRuleProductRuleGroupMgr) WithIDCartRule(idCartRule uint32) Option {
	return optionFunc(func(o *options) { o.query["id_cart_rule"] = idCartRule })
}

// WithQuantity quantity获取
func (obj *_CartRuleProductRuleGroupMgr) WithQuantity(quantity uint32) Option {
	return optionFunc(func(o *options) { o.query["quantity"] = quantity })
}

// GetByOption 功能选项模式获取
func (obj *_CartRuleProductRuleGroupMgr) GetByOption(opts ...Option) (result CartRuleProductRuleGroup, err error) {
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
func (obj *_CartRuleProductRuleGroupMgr) GetByOptions(opts ...Option) (results []*CartRuleProductRuleGroup, err error) {
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

// GetFromIDProductRuleGroup 通过id_product_rule_group获取内容
func (obj *_CartRuleProductRuleGroupMgr) GetFromIDProductRuleGroup(idProductRuleGroup uint32) (result CartRuleProductRuleGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_rule_group = ?", idProductRuleGroup).Find(&result).Error

	return
}

// GetBatchFromIDProductRuleGroup 批量唯一主键查找
func (obj *_CartRuleProductRuleGroupMgr) GetBatchFromIDProductRuleGroup(idProductRuleGroups []uint32) (results []*CartRuleProductRuleGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_rule_group IN (?)", idProductRuleGroups).Find(&results).Error

	return
}

// GetFromIDCartRule 通过id_cart_rule获取内容
func (obj *_CartRuleProductRuleGroupMgr) GetFromIDCartRule(idCartRule uint32) (results []*CartRuleProductRuleGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart_rule = ?", idCartRule).Find(&results).Error

	return
}

// GetBatchFromIDCartRule 批量唯一主键查找
func (obj *_CartRuleProductRuleGroupMgr) GetBatchFromIDCartRule(idCartRules []uint32) (results []*CartRuleProductRuleGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart_rule IN (?)", idCartRules).Find(&results).Error

	return
}

// GetFromQuantity 通过quantity获取内容
func (obj *_CartRuleProductRuleGroupMgr) GetFromQuantity(quantity uint32) (results []*CartRuleProductRuleGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("quantity = ?", quantity).Find(&results).Error

	return
}

// GetBatchFromQuantity 批量唯一主键查找
func (obj *_CartRuleProductRuleGroupMgr) GetBatchFromQuantity(quantitys []uint32) (results []*CartRuleProductRuleGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("quantity IN (?)", quantitys).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_CartRuleProductRuleGroupMgr) FetchByPrimaryKey(idProductRuleGroup uint32) (result CartRuleProductRuleGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_rule_group = ?", idProductRuleGroup).Find(&result).Error

	return
}
