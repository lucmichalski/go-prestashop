package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _OrderInvoicePaymentMgr struct {
	*_BaseMgr
}

// OrderInvoicePaymentMgr open func
func OrderInvoicePaymentMgr(db *gorm.DB) *_OrderInvoicePaymentMgr {
	if db == nil {
		panic(fmt.Errorf("OrderInvoicePaymentMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_OrderInvoicePaymentMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_order_invoice_payment"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_OrderInvoicePaymentMgr) GetTableName() string {
	return "ps_order_invoice_payment"
}

// Get 获取
func (obj *_OrderInvoicePaymentMgr) Get() (result OrderInvoicePayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_OrderInvoicePaymentMgr) Gets() (results []*OrderInvoicePayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDOrderInvoice id_order_invoice获取
func (obj *_OrderInvoicePaymentMgr) WithIDOrderInvoice(idOrderInvoice uint32) Option {
	return optionFunc(func(o *options) { o.query["id_order_invoice"] = idOrderInvoice })
}

// WithIDOrderPayment id_order_payment获取
func (obj *_OrderInvoicePaymentMgr) WithIDOrderPayment(idOrderPayment uint32) Option {
	return optionFunc(func(o *options) { o.query["id_order_payment"] = idOrderPayment })
}

// WithIDOrder id_order获取
func (obj *_OrderInvoicePaymentMgr) WithIDOrder(idOrder uint32) Option {
	return optionFunc(func(o *options) { o.query["id_order"] = idOrder })
}

// GetByOption 功能选项模式获取
func (obj *_OrderInvoicePaymentMgr) GetByOption(opts ...Option) (result OrderInvoicePayment, err error) {
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
func (obj *_OrderInvoicePaymentMgr) GetByOptions(opts ...Option) (results []*OrderInvoicePayment, err error) {
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

// GetFromIDOrderInvoice 通过id_order_invoice获取内容
func (obj *_OrderInvoicePaymentMgr) GetFromIDOrderInvoice(idOrderInvoice uint32) (results []*OrderInvoicePayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_invoice = ?", idOrderInvoice).Find(&results).Error

	return
}

// GetBatchFromIDOrderInvoice 批量唯一主键查找
func (obj *_OrderInvoicePaymentMgr) GetBatchFromIDOrderInvoice(idOrderInvoices []uint32) (results []*OrderInvoicePayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_invoice IN (?)", idOrderInvoices).Find(&results).Error

	return
}

// GetFromIDOrderPayment 通过id_order_payment获取内容
func (obj *_OrderInvoicePaymentMgr) GetFromIDOrderPayment(idOrderPayment uint32) (results []*OrderInvoicePayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_payment = ?", idOrderPayment).Find(&results).Error

	return
}

// GetBatchFromIDOrderPayment 批量唯一主键查找
func (obj *_OrderInvoicePaymentMgr) GetBatchFromIDOrderPayment(idOrderPayments []uint32) (results []*OrderInvoicePayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_payment IN (?)", idOrderPayments).Find(&results).Error

	return
}

// GetFromIDOrder 通过id_order获取内容
func (obj *_OrderInvoicePaymentMgr) GetFromIDOrder(idOrder uint32) (results []*OrderInvoicePayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order = ?", idOrder).Find(&results).Error

	return
}

// GetBatchFromIDOrder 批量唯一主键查找
func (obj *_OrderInvoicePaymentMgr) GetBatchFromIDOrder(idOrders []uint32) (results []*OrderInvoicePayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order IN (?)", idOrders).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_OrderInvoicePaymentMgr) FetchByPrimaryKey(idOrderInvoice uint32, idOrderPayment uint32) (result OrderInvoicePayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_invoice = ? AND id_order_payment = ?", idOrderInvoice, idOrderPayment).Find(&result).Error

	return
}

// FetchIndexByOrderPayment  获取多个内容
func (obj *_OrderInvoicePaymentMgr) FetchIndexByOrderPayment(idOrderPayment uint32) (results []*OrderInvoicePayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_payment = ?", idOrderPayment).Find(&results).Error

	return
}

// FetchIndexByIDOrder  获取多个内容
func (obj *_OrderInvoicePaymentMgr) FetchIndexByIDOrder(idOrder uint32) (results []*OrderInvoicePayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order = ?", idOrder).Find(&results).Error

	return
}
