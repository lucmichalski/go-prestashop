package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _CartRuleCountryMgr struct {
	*_BaseMgr
}

func CartRuleCountryMgr(db *gorm.DB) *_CartRuleCountryMgr {
	if db == nil {
		panic(fmt.Errorf("CartRuleCountryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CartRuleCountryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_cart_rule_country"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_CartRuleCountryMgr) GetTableName() string {
	return "ps_cart_rule_country"
}

func (obj *_CartRuleCountryMgr) Get() (result CartRuleCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_CartRuleCountryMgr) Gets() (results []*CartRuleCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

func (obj *_CartRuleCountryMgr) WithIDCartRule(idCartRule uint32) Option {
	return optionFunc(func(o *options) { o.query["id_cart_rule"] = idCartRule })
}

func (obj *_CartRuleCountryMgr) WithIDCountry(idCountry uint32) Option {
	return optionFunc(func(o *options) { o.query["id_country"] = idCountry })
}

func (obj *_CartRuleCountryMgr) GetByOption(opts ...Option) (result CartRuleCountry, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_CartRuleCountryMgr) GetByOptions(opts ...Option) (results []*CartRuleCountry, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

func (obj *_CartRuleCountryMgr) GetFromIDCartRule(idCartRule uint32) (results []*CartRuleCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart_rule = ?", idCartRule).Find(&results).Error

	return
}

func (obj *_CartRuleCountryMgr) GetBatchFromIDCartRule(idCartRules []uint32) (results []*CartRuleCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart_rule IN (?)", idCartRules).Find(&results).Error

	return
}

func (obj *_CartRuleCountryMgr) GetFromIDCountry(idCountry uint32) (results []*CartRuleCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_country = ?", idCountry).Find(&results).Error

	return
}

func (obj *_CartRuleCountryMgr) GetBatchFromIDCountry(idCountrys []uint32) (results []*CartRuleCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_country IN (?)", idCountrys).Find(&results).Error

	return
}

func (obj *_CartRuleCountryMgr) FetchByPrimaryKey(idCartRule uint32, idCountry uint32) (result CartRuleCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart_rule = ? AND id_country = ?", idCartRule, idCountry).Find(&result).Error

	return
}
