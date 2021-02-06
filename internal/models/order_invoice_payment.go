package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _OrderInvoicePaymentMgr struct {
	*_BaseMgr
}

func OrderInvoicePaymentMgr(db *gorm.DB) *_OrderInvoicePaymentMgr {
	if db == nil {
		panic(fmt.Errorf("OrderInvoicePaymentMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_OrderInvoicePaymentMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_order_invoice_payment"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_OrderInvoicePaymentMgr) GetTableName() string {
	return "ps_order_invoice_payment"
}

func (obj *_OrderInvoicePaymentMgr) Get() (result OrderInvoicePayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_OrderInvoicePaymentMgr) Gets() (results []*OrderInvoicePayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_OrderInvoicePaymentMgr) WithIDOrderInvoice(idOrderInvoice uint32) Option {
	return optionFunc(func(o *options) { o.query["id_order_invoice"] = idOrderInvoice })
}

func (obj *_OrderInvoicePaymentMgr) WithIDOrderPayment(idOrderPayment uint32) Option {
	return optionFunc(func(o *options) { o.query["id_order_payment"] = idOrderPayment })
}

func (obj *_OrderInvoicePaymentMgr) WithIDOrder(idOrder uint32) Option {
	return optionFunc(func(o *options) { o.query["id_order"] = idOrder })
}

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


func (obj *_OrderInvoicePaymentMgr) GetFromIDOrderInvoice(idOrderInvoice uint32) (results []*OrderInvoicePayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_invoice = ?", idOrderInvoice).Find(&results).Error

	return
}

func (obj *_OrderInvoicePaymentMgr) GetBatchFromIDOrderInvoice(idOrderInvoices []uint32) (results []*OrderInvoicePayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_invoice IN (?)", idOrderInvoices).Find(&results).Error

	return
}

func (obj *_OrderInvoicePaymentMgr) GetFromIDOrderPayment(idOrderPayment uint32) (results []*OrderInvoicePayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_payment = ?", idOrderPayment).Find(&results).Error

	return
}

func (obj *_OrderInvoicePaymentMgr) GetBatchFromIDOrderPayment(idOrderPayments []uint32) (results []*OrderInvoicePayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_payment IN (?)", idOrderPayments).Find(&results).Error

	return
}

func (obj *_OrderInvoicePaymentMgr) GetFromIDOrder(idOrder uint32) (results []*OrderInvoicePayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order = ?", idOrder).Find(&results).Error

	return
}

func (obj *_OrderInvoicePaymentMgr) GetBatchFromIDOrder(idOrders []uint32) (results []*OrderInvoicePayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order IN (?)", idOrders).Find(&results).Error

	return
}


func (obj *_OrderInvoicePaymentMgr) FetchByPrimaryKey(idOrderInvoice uint32, idOrderPayment uint32) (result OrderInvoicePayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_invoice = ? AND id_order_payment = ?", idOrderInvoice, idOrderPayment).Find(&result).Error

	return
}

func (obj *_OrderInvoicePaymentMgr) FetchIndexByOrderPayment(idOrderPayment uint32) (results []*OrderInvoicePayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_payment = ?", idOrderPayment).Find(&results).Error

	return
}

func (obj *_OrderInvoicePaymentMgr) FetchIndexByIDOrder(idOrder uint32) (results []*OrderInvoicePayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order = ?", idOrder).Find(&results).Error

	return
}
