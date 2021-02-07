package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _CartRuleShopMgr struct {
	*_BaseMgr
}

func CartRuleShopMgr(db *gorm.DB) *_CartRuleShopMgr {
	if db == nil {
		panic(fmt.Errorf("CartRuleShopMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CartRuleShopMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_cart_rule_shop"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_CartRuleShopMgr) GetTableName() string {
	return "ps_cart_rule_shop"
}

func (obj *_CartRuleShopMgr) Get() (result CartRuleShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_CartRuleShopMgr) Gets() (results []*CartRuleShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

func (obj *_CartRuleShopMgr) WithIDCartRule(idCartRule uint32) Option {
	return optionFunc(func(o *options) { o.query["id_cart_rule"] = idCartRule })
}

func (obj *_CartRuleShopMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

func (obj *_CartRuleShopMgr) GetByOption(opts ...Option) (result CartRuleShop, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_CartRuleShopMgr) GetByOptions(opts ...Option) (results []*CartRuleShop, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

func (obj *_CartRuleShopMgr) GetFromIDCartRule(idCartRule uint32) (results []*CartRuleShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart_rule = ?", idCartRule).Find(&results).Error

	return
}

func (obj *_CartRuleShopMgr) GetBatchFromIDCartRule(idCartRules []uint32) (results []*CartRuleShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart_rule IN (?)", idCartRules).Find(&results).Error

	return
}

func (obj *_CartRuleShopMgr) GetFromIDShop(idShop uint32) (results []*CartRuleShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

func (obj *_CartRuleShopMgr) GetBatchFromIDShop(idShops []uint32) (results []*CartRuleShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

func (obj *_CartRuleShopMgr) FetchByPrimaryKey(idCartRule uint32, idShop uint32) (result CartRuleShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart_rule = ? AND id_shop = ?", idCartRule, idShop).Find(&result).Error

	return
}
