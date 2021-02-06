package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _CartRuleProductRuleValueMgr struct {
	*_BaseMgr
}

// CartRuleProductRuleValueMgr open func
func CartRuleProductRuleValueMgr(db *gorm.DB) *_CartRuleProductRuleValueMgr {
	if db == nil {
		panic(fmt.Errorf("CartRuleProductRuleValueMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CartRuleProductRuleValueMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_cart_rule_product_rule_value"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CartRuleProductRuleValueMgr) GetTableName() string {
	return "ps_cart_rule_product_rule_value"
}

// Get 获取
func (obj *_CartRuleProductRuleValueMgr) Get() (result CartRuleProductRuleValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CartRuleProductRuleValueMgr) Gets() (results []*CartRuleProductRuleValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDProductRule id_product_rule获取
func (obj *_CartRuleProductRuleValueMgr) WithIDProductRule(idProductRule uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product_rule"] = idProductRule })
}

// WithIDItem id_item获取
func (obj *_CartRuleProductRuleValueMgr) WithIDItem(idItem uint32) Option {
	return optionFunc(func(o *options) { o.query["id_item"] = idItem })
}

// GetByOption 功能选项模式获取
func (obj *_CartRuleProductRuleValueMgr) GetByOption(opts ...Option) (result CartRuleProductRuleValue, err error) {
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
func (obj *_CartRuleProductRuleValueMgr) GetByOptions(opts ...Option) (results []*CartRuleProductRuleValue, err error) {
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

// GetFromIDProductRule 通过id_product_rule获取内容
func (obj *_CartRuleProductRuleValueMgr) GetFromIDProductRule(idProductRule uint32) (results []*CartRuleProductRuleValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_rule = ?", idProductRule).Find(&results).Error

	return
}

// GetBatchFromIDProductRule 批量唯一主键查找
func (obj *_CartRuleProductRuleValueMgr) GetBatchFromIDProductRule(idProductRules []uint32) (results []*CartRuleProductRuleValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_rule IN (?)", idProductRules).Find(&results).Error

	return
}

// GetFromIDItem 通过id_item获取内容
func (obj *_CartRuleProductRuleValueMgr) GetFromIDItem(idItem uint32) (results []*CartRuleProductRuleValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_item = ?", idItem).Find(&results).Error

	return
}

// GetBatchFromIDItem 批量唯一主键查找
func (obj *_CartRuleProductRuleValueMgr) GetBatchFromIDItem(idItems []uint32) (results []*CartRuleProductRuleValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_item IN (?)", idItems).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_CartRuleProductRuleValueMgr) FetchByPrimaryKey(idProductRule uint32, idItem uint32) (result CartRuleProductRuleValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_rule = ? AND id_item = ?", idProductRule, idItem).Find(&result).Error

	return
}
