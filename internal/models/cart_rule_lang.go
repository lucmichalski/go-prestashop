package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _CartRuleLangMgr struct {
	*_BaseMgr
}

func CartRuleLangMgr(db *gorm.DB) *_CartRuleLangMgr {
	if db == nil {
		panic(fmt.Errorf("CartRuleLangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CartRuleLangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_cart_rule_lang"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_CartRuleLangMgr) GetTableName() string {
	return "ps_cart_rule_lang"
}

func (obj *_CartRuleLangMgr) Get() (result CartRuleLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_CartRuleLangMgr) Gets() (results []*CartRuleLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

func (obj *_CartRuleLangMgr) WithIDCartRule(idCartRule uint32) Option {
	return optionFunc(func(o *options) { o.query["id_cart_rule"] = idCartRule })
}

func (obj *_CartRuleLangMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

func (obj *_CartRuleLangMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

func (obj *_CartRuleLangMgr) GetByOption(opts ...Option) (result CartRuleLang, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_CartRuleLangMgr) GetByOptions(opts ...Option) (results []*CartRuleLang, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

func (obj *_CartRuleLangMgr) GetFromIDCartRule(idCartRule uint32) (results []*CartRuleLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart_rule = ?", idCartRule).Find(&results).Error

	return
}

func (obj *_CartRuleLangMgr) GetBatchFromIDCartRule(idCartRules []uint32) (results []*CartRuleLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart_rule IN (?)", idCartRules).Find(&results).Error

	return
}

func (obj *_CartRuleLangMgr) GetFromIDLang(idLang uint32) (results []*CartRuleLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error

	return
}

func (obj *_CartRuleLangMgr) GetBatchFromIDLang(idLangs []uint32) (results []*CartRuleLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error

	return
}

func (obj *_CartRuleLangMgr) GetFromName(name string) (results []*CartRuleLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

func (obj *_CartRuleLangMgr) GetBatchFromName(names []string) (results []*CartRuleLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

func (obj *_CartRuleLangMgr) FetchByPrimaryKey(idCartRule uint32, idLang uint32) (result CartRuleLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart_rule = ? AND id_lang = ?", idCartRule, idLang).Find(&result).Error

	return
}
