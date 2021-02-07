package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _OrderCartRuleMgr struct {
	*_BaseMgr
}

func OrderCartRuleMgr(db *gorm.DB) *_OrderCartRuleMgr {
	if db == nil {
		panic(fmt.Errorf("OrderCartRuleMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_OrderCartRuleMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_order_cart_rule"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_OrderCartRuleMgr) GetTableName() string {
	return "ps_order_cart_rule"
}

func (obj *_OrderCartRuleMgr) Get() (result OrderCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_OrderCartRuleMgr) Gets() (results []*OrderCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

func (obj *_OrderCartRuleMgr) WithIDOrderCartRule(idOrderCartRule uint32) Option {
	return optionFunc(func(o *options) { o.query["id_order_cart_rule"] = idOrderCartRule })
}

func (obj *_OrderCartRuleMgr) WithIDOrder(idOrder uint32) Option {
	return optionFunc(func(o *options) { o.query["id_order"] = idOrder })
}

func (obj *_OrderCartRuleMgr) WithIDCartRule(idCartRule uint32) Option {
	return optionFunc(func(o *options) { o.query["id_cart_rule"] = idCartRule })
}

func (obj *_OrderCartRuleMgr) WithIDOrderInvoice(idOrderInvoice uint32) Option {
	return optionFunc(func(o *options) { o.query["id_order_invoice"] = idOrderInvoice })
}

func (obj *_OrderCartRuleMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

func (obj *_OrderCartRuleMgr) WithValue(value float64) Option {
	return optionFunc(func(o *options) { o.query["value"] = value })
}

func (obj *_OrderCartRuleMgr) WithValueTaxExcl(valueTaxExcl float64) Option {
	return optionFunc(func(o *options) { o.query["value_tax_excl"] = valueTaxExcl })
}

func (obj *_OrderCartRuleMgr) WithFreeShipping(freeShipping bool) Option {
	return optionFunc(func(o *options) { o.query["free_shipping"] = freeShipping })
}

func (obj *_OrderCartRuleMgr) WithDeleted(deleted bool) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

func (obj *_OrderCartRuleMgr) GetByOption(opts ...Option) (result OrderCartRule, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_OrderCartRuleMgr) GetByOptions(opts ...Option) (results []*OrderCartRule, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

func (obj *_OrderCartRuleMgr) GetFromIDOrderCartRule(idOrderCartRule uint32) (result OrderCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_cart_rule = ?", idOrderCartRule).Find(&result).Error

	return
}

func (obj *_OrderCartRuleMgr) GetBatchFromIDOrderCartRule(idOrderCartRules []uint32) (results []*OrderCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_cart_rule IN (?)", idOrderCartRules).Find(&results).Error

	return
}

func (obj *_OrderCartRuleMgr) GetFromIDOrder(idOrder uint32) (results []*OrderCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order = ?", idOrder).Find(&results).Error

	return
}

func (obj *_OrderCartRuleMgr) GetBatchFromIDOrder(idOrders []uint32) (results []*OrderCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order IN (?)", idOrders).Find(&results).Error

	return
}

func (obj *_OrderCartRuleMgr) GetFromIDCartRule(idCartRule uint32) (results []*OrderCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart_rule = ?", idCartRule).Find(&results).Error

	return
}

func (obj *_OrderCartRuleMgr) GetBatchFromIDCartRule(idCartRules []uint32) (results []*OrderCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart_rule IN (?)", idCartRules).Find(&results).Error

	return
}

func (obj *_OrderCartRuleMgr) GetFromIDOrderInvoice(idOrderInvoice uint32) (results []*OrderCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_invoice = ?", idOrderInvoice).Find(&results).Error

	return
}

func (obj *_OrderCartRuleMgr) GetBatchFromIDOrderInvoice(idOrderInvoices []uint32) (results []*OrderCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_invoice IN (?)", idOrderInvoices).Find(&results).Error

	return
}

func (obj *_OrderCartRuleMgr) GetFromName(name string) (results []*OrderCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

func (obj *_OrderCartRuleMgr) GetBatchFromName(names []string) (results []*OrderCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

func (obj *_OrderCartRuleMgr) GetFromValue(value float64) (results []*OrderCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("value = ?", value).Find(&results).Error

	return
}

func (obj *_OrderCartRuleMgr) GetBatchFromValue(values []float64) (results []*OrderCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("value IN (?)", values).Find(&results).Error

	return
}

func (obj *_OrderCartRuleMgr) GetFromValueTaxExcl(valueTaxExcl float64) (results []*OrderCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("value_tax_excl = ?", valueTaxExcl).Find(&results).Error

	return
}

func (obj *_OrderCartRuleMgr) GetBatchFromValueTaxExcl(valueTaxExcls []float64) (results []*OrderCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("value_tax_excl IN (?)", valueTaxExcls).Find(&results).Error

	return
}

func (obj *_OrderCartRuleMgr) GetFromFreeShipping(freeShipping bool) (results []*OrderCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("free_shipping = ?", freeShipping).Find(&results).Error

	return
}

func (obj *_OrderCartRuleMgr) GetBatchFromFreeShipping(freeShippings []bool) (results []*OrderCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("free_shipping IN (?)", freeShippings).Find(&results).Error

	return
}

func (obj *_OrderCartRuleMgr) GetFromDeleted(deleted bool) (results []*OrderCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("deleted = ?", deleted).Find(&results).Error

	return
}

func (obj *_OrderCartRuleMgr) GetBatchFromDeleted(deleteds []bool) (results []*OrderCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("deleted IN (?)", deleteds).Find(&results).Error

	return
}

func (obj *_OrderCartRuleMgr) FetchByPrimaryKey(idOrderCartRule uint32) (result OrderCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_cart_rule = ?", idOrderCartRule).Find(&result).Error

	return
}

func (obj *_OrderCartRuleMgr) FetchIndexByIDOrder(idOrder uint32) (results []*OrderCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order = ?", idOrder).Find(&results).Error

	return
}

func (obj *_OrderCartRuleMgr) FetchIndexByIDCartRule(idCartRule uint32) (results []*OrderCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart_rule = ?", idCartRule).Find(&results).Error

	return
}
