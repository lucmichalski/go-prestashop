package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _CartCartRuleMgr struct {
	*_BaseMgr
}

func CartCartRuleMgr(db *gorm.DB) *_CartCartRuleMgr {
	if db == nil {
		panic(fmt.Errorf("CartCartRuleMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CartCartRuleMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_cart_cart_rule"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_CartCartRuleMgr) GetTableName() string {
	return "ps_cart_cart_rule"
}

func (obj *_CartCartRuleMgr) Get() (result CartCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_CartCartRuleMgr) Gets() (results []*CartCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

func (obj *_CartCartRuleMgr) WithIDCart(idCart uint32) Option {
	return optionFunc(func(o *options) { o.query["id_cart"] = idCart })
}

func (obj *_CartCartRuleMgr) WithIDCartRule(idCartRule uint32) Option {
	return optionFunc(func(o *options) { o.query["id_cart_rule"] = idCartRule })
}

func (obj *_CartCartRuleMgr) GetByOption(opts ...Option) (result CartCartRule, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_CartCartRuleMgr) GetByOptions(opts ...Option) (results []*CartCartRule, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

func (obj *_CartCartRuleMgr) GetFromIDCart(idCart uint32) (results []*CartCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart = ?", idCart).Find(&results).Error

	return
}

func (obj *_CartCartRuleMgr) GetBatchFromIDCart(idCarts []uint32) (results []*CartCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart IN (?)", idCarts).Find(&results).Error

	return
}

func (obj *_CartCartRuleMgr) GetFromIDCartRule(idCartRule uint32) (results []*CartCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart_rule = ?", idCartRule).Find(&results).Error

	return
}

func (obj *_CartCartRuleMgr) GetBatchFromIDCartRule(idCartRules []uint32) (results []*CartCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart_rule IN (?)", idCartRules).Find(&results).Error

	return
}

func (obj *_CartCartRuleMgr) FetchByPrimaryKey(idCart uint32, idCartRule uint32) (result CartCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart = ? AND id_cart_rule = ?", idCart, idCartRule).Find(&result).Error

	return
}

func (obj *_CartCartRuleMgr) FetchIndexByIDCartRule(idCartRule uint32) (results []*CartCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart_rule = ?", idCartRule).Find(&results).Error

	return
}
