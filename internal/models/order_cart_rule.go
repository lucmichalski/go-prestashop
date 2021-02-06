package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _OrderCartRuleMgr struct {
	*_BaseMgr
}

// OrderCartRuleMgr open func
func OrderCartRuleMgr(db *gorm.DB) *_OrderCartRuleMgr {
	if db == nil {
		panic(fmt.Errorf("OrderCartRuleMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_OrderCartRuleMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_order_cart_rule"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_OrderCartRuleMgr) GetTableName() string {
	return "ps_order_cart_rule"
}

// Get 获取
func (obj *_OrderCartRuleMgr) Get() (result OrderCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_OrderCartRuleMgr) Gets() (results []*OrderCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDOrderCartRule id_order_cart_rule获取
func (obj *_OrderCartRuleMgr) WithIDOrderCartRule(idOrderCartRule uint32) Option {
	return optionFunc(func(o *options) { o.query["id_order_cart_rule"] = idOrderCartRule })
}

// WithIDOrder id_order获取
func (obj *_OrderCartRuleMgr) WithIDOrder(idOrder uint32) Option {
	return optionFunc(func(o *options) { o.query["id_order"] = idOrder })
}

// WithIDCartRule id_cart_rule获取
func (obj *_OrderCartRuleMgr) WithIDCartRule(idCartRule uint32) Option {
	return optionFunc(func(o *options) { o.query["id_cart_rule"] = idCartRule })
}

// WithIDOrderInvoice id_order_invoice获取
func (obj *_OrderCartRuleMgr) WithIDOrderInvoice(idOrderInvoice uint32) Option {
	return optionFunc(func(o *options) { o.query["id_order_invoice"] = idOrderInvoice })
}

// WithName name获取
func (obj *_OrderCartRuleMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithValue value获取
func (obj *_OrderCartRuleMgr) WithValue(value float64) Option {
	return optionFunc(func(o *options) { o.query["value"] = value })
}

// WithValueTaxExcl value_tax_excl获取
func (obj *_OrderCartRuleMgr) WithValueTaxExcl(valueTaxExcl float64) Option {
	return optionFunc(func(o *options) { o.query["value_tax_excl"] = valueTaxExcl })
}

// WithFreeShipping free_shipping获取
func (obj *_OrderCartRuleMgr) WithFreeShipping(freeShipping bool) Option {
	return optionFunc(func(o *options) { o.query["free_shipping"] = freeShipping })
}

// WithDeleted deleted获取
func (obj *_OrderCartRuleMgr) WithDeleted(deleted bool) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// GetByOption 功能选项模式获取
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

// GetByOptions 批量功能选项模式获取
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

//////////////////////////enume case ////////////////////////////////////////////

// GetFromIDOrderCartRule 通过id_order_cart_rule获取内容
func (obj *_OrderCartRuleMgr) GetFromIDOrderCartRule(idOrderCartRule uint32) (result OrderCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_cart_rule = ?", idOrderCartRule).Find(&result).Error

	return
}

// GetBatchFromIDOrderCartRule 批量唯一主键查找
func (obj *_OrderCartRuleMgr) GetBatchFromIDOrderCartRule(idOrderCartRules []uint32) (results []*OrderCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_cart_rule IN (?)", idOrderCartRules).Find(&results).Error

	return
}

// GetFromIDOrder 通过id_order获取内容
func (obj *_OrderCartRuleMgr) GetFromIDOrder(idOrder uint32) (results []*OrderCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order = ?", idOrder).Find(&results).Error

	return
}

// GetBatchFromIDOrder 批量唯一主键查找
func (obj *_OrderCartRuleMgr) GetBatchFromIDOrder(idOrders []uint32) (results []*OrderCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order IN (?)", idOrders).Find(&results).Error

	return
}

// GetFromIDCartRule 通过id_cart_rule获取内容
func (obj *_OrderCartRuleMgr) GetFromIDCartRule(idCartRule uint32) (results []*OrderCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart_rule = ?", idCartRule).Find(&results).Error

	return
}

// GetBatchFromIDCartRule 批量唯一主键查找
func (obj *_OrderCartRuleMgr) GetBatchFromIDCartRule(idCartRules []uint32) (results []*OrderCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart_rule IN (?)", idCartRules).Find(&results).Error

	return
}

// GetFromIDOrderInvoice 通过id_order_invoice获取内容
func (obj *_OrderCartRuleMgr) GetFromIDOrderInvoice(idOrderInvoice uint32) (results []*OrderCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_invoice = ?", idOrderInvoice).Find(&results).Error

	return
}

// GetBatchFromIDOrderInvoice 批量唯一主键查找
func (obj *_OrderCartRuleMgr) GetBatchFromIDOrderInvoice(idOrderInvoices []uint32) (results []*OrderCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_invoice IN (?)", idOrderInvoices).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_OrderCartRuleMgr) GetFromName(name string) (results []*OrderCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量唯一主键查找
func (obj *_OrderCartRuleMgr) GetBatchFromName(names []string) (results []*OrderCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

// GetFromValue 通过value获取内容
func (obj *_OrderCartRuleMgr) GetFromValue(value float64) (results []*OrderCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("value = ?", value).Find(&results).Error

	return
}

// GetBatchFromValue 批量唯一主键查找
func (obj *_OrderCartRuleMgr) GetBatchFromValue(values []float64) (results []*OrderCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("value IN (?)", values).Find(&results).Error

	return
}

// GetFromValueTaxExcl 通过value_tax_excl获取内容
func (obj *_OrderCartRuleMgr) GetFromValueTaxExcl(valueTaxExcl float64) (results []*OrderCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("value_tax_excl = ?", valueTaxExcl).Find(&results).Error

	return
}

// GetBatchFromValueTaxExcl 批量唯一主键查找
func (obj *_OrderCartRuleMgr) GetBatchFromValueTaxExcl(valueTaxExcls []float64) (results []*OrderCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("value_tax_excl IN (?)", valueTaxExcls).Find(&results).Error

	return
}

// GetFromFreeShipping 通过free_shipping获取内容
func (obj *_OrderCartRuleMgr) GetFromFreeShipping(freeShipping bool) (results []*OrderCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("free_shipping = ?", freeShipping).Find(&results).Error

	return
}

// GetBatchFromFreeShipping 批量唯一主键查找
func (obj *_OrderCartRuleMgr) GetBatchFromFreeShipping(freeShippings []bool) (results []*OrderCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("free_shipping IN (?)", freeShippings).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容
func (obj *_OrderCartRuleMgr) GetFromDeleted(deleted bool) (results []*OrderCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("deleted = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量唯一主键查找
func (obj *_OrderCartRuleMgr) GetBatchFromDeleted(deleteds []bool) (results []*OrderCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("deleted IN (?)", deleteds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_OrderCartRuleMgr) FetchByPrimaryKey(idOrderCartRule uint32) (result OrderCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_cart_rule = ?", idOrderCartRule).Find(&result).Error

	return
}

// FetchIndexByIDOrder  获取多个内容
func (obj *_OrderCartRuleMgr) FetchIndexByIDOrder(idOrder uint32) (results []*OrderCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order = ?", idOrder).Find(&results).Error

	return
}

// FetchIndexByIDCartRule  获取多个内容
func (obj *_OrderCartRuleMgr) FetchIndexByIDCartRule(idCartRule uint32) (results []*OrderCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart_rule = ?", idCartRule).Find(&results).Error

	return
}
