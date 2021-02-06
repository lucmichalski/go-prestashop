package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _OrderStateMgr struct {
	*_BaseMgr
}

func OrderStateMgr(db *gorm.DB) *_OrderStateMgr {
	if db == nil {
		panic(fmt.Errorf("OrderStateMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_OrderStateMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_order_state"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_OrderStateMgr) GetTableName() string {
	return "ps_order_state"
}

func (obj *_OrderStateMgr) Get() (result OrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_OrderStateMgr) Gets() (results []*OrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_OrderStateMgr) WithIDOrderState(idOrderState uint32) Option {
	return optionFunc(func(o *options) { o.query["id_order_state"] = idOrderState })
}

func (obj *_OrderStateMgr) WithInvoice(invoice bool) Option {
	return optionFunc(func(o *options) { o.query["invoice"] = invoice })
}

func (obj *_OrderStateMgr) WithSendEmail(sendEmail bool) Option {
	return optionFunc(func(o *options) { o.query["send_email"] = sendEmail })
}

func (obj *_OrderStateMgr) WithModuleName(moduleName string) Option {
	return optionFunc(func(o *options) { o.query["module_name"] = moduleName })
}

func (obj *_OrderStateMgr) WithColor(color string) Option {
	return optionFunc(func(o *options) { o.query["color"] = color })
}

func (obj *_OrderStateMgr) WithUnremovable(unremovable bool) Option {
	return optionFunc(func(o *options) { o.query["unremovable"] = unremovable })
}

func (obj *_OrderStateMgr) WithHidden(hidden bool) Option {
	return optionFunc(func(o *options) { o.query["hidden"] = hidden })
}

func (obj *_OrderStateMgr) WithLogable(logable bool) Option {
	return optionFunc(func(o *options) { o.query["logable"] = logable })
}

func (obj *_OrderStateMgr) WithDelivery(delivery bool) Option {
	return optionFunc(func(o *options) { o.query["delivery"] = delivery })
}

func (obj *_OrderStateMgr) WithShipped(shipped bool) Option {
	return optionFunc(func(o *options) { o.query["shipped"] = shipped })
}

func (obj *_OrderStateMgr) WithPaid(paid bool) Option {
	return optionFunc(func(o *options) { o.query["paid"] = paid })
}

func (obj *_OrderStateMgr) WithPdfInvoice(pdfInvoice bool) Option {
	return optionFunc(func(o *options) { o.query["pdf_invoice"] = pdfInvoice })
}

func (obj *_OrderStateMgr) WithPdfDelivery(pdfDelivery bool) Option {
	return optionFunc(func(o *options) { o.query["pdf_delivery"] = pdfDelivery })
}

func (obj *_OrderStateMgr) WithDeleted(deleted bool) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

func (obj *_OrderStateMgr) GetByOption(opts ...Option) (result OrderState, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_OrderStateMgr) GetByOptions(opts ...Option) (results []*OrderState, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_OrderStateMgr) GetFromIDOrderState(idOrderState uint32) (result OrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_state = ?", idOrderState).Find(&result).Error

	return
}

func (obj *_OrderStateMgr) GetBatchFromIDOrderState(idOrderStates []uint32) (results []*OrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_state IN (?)", idOrderStates).Find(&results).Error

	return
}

func (obj *_OrderStateMgr) GetFromInvoice(invoice bool) (results []*OrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("invoice = ?", invoice).Find(&results).Error

	return
}

func (obj *_OrderStateMgr) GetBatchFromInvoice(invoices []bool) (results []*OrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("invoice IN (?)", invoices).Find(&results).Error

	return
}

func (obj *_OrderStateMgr) GetFromSendEmail(sendEmail bool) (results []*OrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("send_email = ?", sendEmail).Find(&results).Error

	return
}

func (obj *_OrderStateMgr) GetBatchFromSendEmail(sendEmails []bool) (results []*OrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("send_email IN (?)", sendEmails).Find(&results).Error

	return
}

func (obj *_OrderStateMgr) GetFromModuleName(moduleName string) (results []*OrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("module_name = ?", moduleName).Find(&results).Error

	return
}

func (obj *_OrderStateMgr) GetBatchFromModuleName(moduleNames []string) (results []*OrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("module_name IN (?)", moduleNames).Find(&results).Error

	return
}

func (obj *_OrderStateMgr) GetFromColor(color string) (results []*OrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("color = ?", color).Find(&results).Error

	return
}

func (obj *_OrderStateMgr) GetBatchFromColor(colors []string) (results []*OrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("color IN (?)", colors).Find(&results).Error

	return
}

func (obj *_OrderStateMgr) GetFromUnremovable(unremovable bool) (results []*OrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("unremovable = ?", unremovable).Find(&results).Error

	return
}

func (obj *_OrderStateMgr) GetBatchFromUnremovable(unremovables []bool) (results []*OrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("unremovable IN (?)", unremovables).Find(&results).Error

	return
}

func (obj *_OrderStateMgr) GetFromHidden(hidden bool) (results []*OrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("hidden = ?", hidden).Find(&results).Error

	return
}

func (obj *_OrderStateMgr) GetBatchFromHidden(hiddens []bool) (results []*OrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("hidden IN (?)", hiddens).Find(&results).Error

	return
}

func (obj *_OrderStateMgr) GetFromLogable(logable bool) (results []*OrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("logable = ?", logable).Find(&results).Error

	return
}

func (obj *_OrderStateMgr) GetBatchFromLogable(logables []bool) (results []*OrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("logable IN (?)", logables).Find(&results).Error

	return
}

func (obj *_OrderStateMgr) GetFromDelivery(delivery bool) (results []*OrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("delivery = ?", delivery).Find(&results).Error

	return
}

func (obj *_OrderStateMgr) GetBatchFromDelivery(deliverys []bool) (results []*OrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("delivery IN (?)", deliverys).Find(&results).Error

	return
}

func (obj *_OrderStateMgr) GetFromShipped(shipped bool) (results []*OrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("shipped = ?", shipped).Find(&results).Error

	return
}

func (obj *_OrderStateMgr) GetBatchFromShipped(shippeds []bool) (results []*OrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("shipped IN (?)", shippeds).Find(&results).Error

	return
}

func (obj *_OrderStateMgr) GetFromPaid(paid bool) (results []*OrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("paid = ?", paid).Find(&results).Error

	return
}

func (obj *_OrderStateMgr) GetBatchFromPaid(paids []bool) (results []*OrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("paid IN (?)", paids).Find(&results).Error

	return
}

func (obj *_OrderStateMgr) GetFromPdfInvoice(pdfInvoice bool) (results []*OrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("pdf_invoice = ?", pdfInvoice).Find(&results).Error

	return
}

func (obj *_OrderStateMgr) GetBatchFromPdfInvoice(pdfInvoices []bool) (results []*OrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("pdf_invoice IN (?)", pdfInvoices).Find(&results).Error

	return
}

func (obj *_OrderStateMgr) GetFromPdfDelivery(pdfDelivery bool) (results []*OrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("pdf_delivery = ?", pdfDelivery).Find(&results).Error

	return
}

func (obj *_OrderStateMgr) GetBatchFromPdfDelivery(pdfDeliverys []bool) (results []*OrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("pdf_delivery IN (?)", pdfDeliverys).Find(&results).Error

	return
}

func (obj *_OrderStateMgr) GetFromDeleted(deleted bool) (results []*OrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("deleted = ?", deleted).Find(&results).Error

	return
}

func (obj *_OrderStateMgr) GetBatchFromDeleted(deleteds []bool) (results []*OrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("deleted IN (?)", deleteds).Find(&results).Error

	return
}


func (obj *_OrderStateMgr) FetchByPrimaryKey(idOrderState uint32) (result OrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_state = ?", idOrderState).Find(&result).Error

	return
}

func (obj *_OrderStateMgr) FetchIndexByModuleName(moduleName string) (results []*OrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("module_name = ?", moduleName).Find(&results).Error

	return
}
