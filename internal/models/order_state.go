package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _OrderStateMgr struct {
	*_BaseMgr
}

// OrderStateMgr open func
func OrderStateMgr(db *gorm.DB) *_OrderStateMgr {
	if db == nil {
		panic(fmt.Errorf("OrderStateMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_OrderStateMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_order_state"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_OrderStateMgr) GetTableName() string {
	return "ps_order_state"
}

// Get 获取
func (obj *_OrderStateMgr) Get() (result OrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_OrderStateMgr) Gets() (results []*OrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDOrderState id_order_state获取
func (obj *_OrderStateMgr) WithIDOrderState(idOrderState uint32) Option {
	return optionFunc(func(o *options) { o.query["id_order_state"] = idOrderState })
}

// WithInvoice invoice获取
func (obj *_OrderStateMgr) WithInvoice(invoice bool) Option {
	return optionFunc(func(o *options) { o.query["invoice"] = invoice })
}

// WithSendEmail send_email获取
func (obj *_OrderStateMgr) WithSendEmail(sendEmail bool) Option {
	return optionFunc(func(o *options) { o.query["send_email"] = sendEmail })
}

// WithModuleName module_name获取
func (obj *_OrderStateMgr) WithModuleName(moduleName string) Option {
	return optionFunc(func(o *options) { o.query["module_name"] = moduleName })
}

// WithColor color获取
func (obj *_OrderStateMgr) WithColor(color string) Option {
	return optionFunc(func(o *options) { o.query["color"] = color })
}

// WithUnremovable unremovable获取
func (obj *_OrderStateMgr) WithUnremovable(unremovable bool) Option {
	return optionFunc(func(o *options) { o.query["unremovable"] = unremovable })
}

// WithHidden hidden获取
func (obj *_OrderStateMgr) WithHidden(hidden bool) Option {
	return optionFunc(func(o *options) { o.query["hidden"] = hidden })
}

// WithLogable logable获取
func (obj *_OrderStateMgr) WithLogable(logable bool) Option {
	return optionFunc(func(o *options) { o.query["logable"] = logable })
}

// WithDelivery delivery获取
func (obj *_OrderStateMgr) WithDelivery(delivery bool) Option {
	return optionFunc(func(o *options) { o.query["delivery"] = delivery })
}

// WithShipped shipped获取
func (obj *_OrderStateMgr) WithShipped(shipped bool) Option {
	return optionFunc(func(o *options) { o.query["shipped"] = shipped })
}

// WithPaid paid获取
func (obj *_OrderStateMgr) WithPaid(paid bool) Option {
	return optionFunc(func(o *options) { o.query["paid"] = paid })
}

// WithPdfInvoice pdf_invoice获取
func (obj *_OrderStateMgr) WithPdfInvoice(pdfInvoice bool) Option {
	return optionFunc(func(o *options) { o.query["pdf_invoice"] = pdfInvoice })
}

// WithPdfDelivery pdf_delivery获取
func (obj *_OrderStateMgr) WithPdfDelivery(pdfDelivery bool) Option {
	return optionFunc(func(o *options) { o.query["pdf_delivery"] = pdfDelivery })
}

// WithDeleted deleted获取
func (obj *_OrderStateMgr) WithDeleted(deleted bool) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// GetByOption 功能选项模式获取
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

// GetByOptions 批量功能选项模式获取
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

//////////////////////////enume case ////////////////////////////////////////////

// GetFromIDOrderState 通过id_order_state获取内容
func (obj *_OrderStateMgr) GetFromIDOrderState(idOrderState uint32) (result OrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_state = ?", idOrderState).Find(&result).Error

	return
}

// GetBatchFromIDOrderState 批量唯一主键查找
func (obj *_OrderStateMgr) GetBatchFromIDOrderState(idOrderStates []uint32) (results []*OrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_state IN (?)", idOrderStates).Find(&results).Error

	return
}

// GetFromInvoice 通过invoice获取内容
func (obj *_OrderStateMgr) GetFromInvoice(invoice bool) (results []*OrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("invoice = ?", invoice).Find(&results).Error

	return
}

// GetBatchFromInvoice 批量唯一主键查找
func (obj *_OrderStateMgr) GetBatchFromInvoice(invoices []bool) (results []*OrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("invoice IN (?)", invoices).Find(&results).Error

	return
}

// GetFromSendEmail 通过send_email获取内容
func (obj *_OrderStateMgr) GetFromSendEmail(sendEmail bool) (results []*OrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("send_email = ?", sendEmail).Find(&results).Error

	return
}

// GetBatchFromSendEmail 批量唯一主键查找
func (obj *_OrderStateMgr) GetBatchFromSendEmail(sendEmails []bool) (results []*OrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("send_email IN (?)", sendEmails).Find(&results).Error

	return
}

// GetFromModuleName 通过module_name获取内容
func (obj *_OrderStateMgr) GetFromModuleName(moduleName string) (results []*OrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("module_name = ?", moduleName).Find(&results).Error

	return
}

// GetBatchFromModuleName 批量唯一主键查找
func (obj *_OrderStateMgr) GetBatchFromModuleName(moduleNames []string) (results []*OrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("module_name IN (?)", moduleNames).Find(&results).Error

	return
}

// GetFromColor 通过color获取内容
func (obj *_OrderStateMgr) GetFromColor(color string) (results []*OrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("color = ?", color).Find(&results).Error

	return
}

// GetBatchFromColor 批量唯一主键查找
func (obj *_OrderStateMgr) GetBatchFromColor(colors []string) (results []*OrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("color IN (?)", colors).Find(&results).Error

	return
}

// GetFromUnremovable 通过unremovable获取内容
func (obj *_OrderStateMgr) GetFromUnremovable(unremovable bool) (results []*OrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("unremovable = ?", unremovable).Find(&results).Error

	return
}

// GetBatchFromUnremovable 批量唯一主键查找
func (obj *_OrderStateMgr) GetBatchFromUnremovable(unremovables []bool) (results []*OrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("unremovable IN (?)", unremovables).Find(&results).Error

	return
}

// GetFromHidden 通过hidden获取内容
func (obj *_OrderStateMgr) GetFromHidden(hidden bool) (results []*OrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("hidden = ?", hidden).Find(&results).Error

	return
}

// GetBatchFromHidden 批量唯一主键查找
func (obj *_OrderStateMgr) GetBatchFromHidden(hiddens []bool) (results []*OrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("hidden IN (?)", hiddens).Find(&results).Error

	return
}

// GetFromLogable 通过logable获取内容
func (obj *_OrderStateMgr) GetFromLogable(logable bool) (results []*OrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("logable = ?", logable).Find(&results).Error

	return
}

// GetBatchFromLogable 批量唯一主键查找
func (obj *_OrderStateMgr) GetBatchFromLogable(logables []bool) (results []*OrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("logable IN (?)", logables).Find(&results).Error

	return
}

// GetFromDelivery 通过delivery获取内容
func (obj *_OrderStateMgr) GetFromDelivery(delivery bool) (results []*OrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("delivery = ?", delivery).Find(&results).Error

	return
}

// GetBatchFromDelivery 批量唯一主键查找
func (obj *_OrderStateMgr) GetBatchFromDelivery(deliverys []bool) (results []*OrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("delivery IN (?)", deliverys).Find(&results).Error

	return
}

// GetFromShipped 通过shipped获取内容
func (obj *_OrderStateMgr) GetFromShipped(shipped bool) (results []*OrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("shipped = ?", shipped).Find(&results).Error

	return
}

// GetBatchFromShipped 批量唯一主键查找
func (obj *_OrderStateMgr) GetBatchFromShipped(shippeds []bool) (results []*OrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("shipped IN (?)", shippeds).Find(&results).Error

	return
}

// GetFromPaid 通过paid获取内容
func (obj *_OrderStateMgr) GetFromPaid(paid bool) (results []*OrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("paid = ?", paid).Find(&results).Error

	return
}

// GetBatchFromPaid 批量唯一主键查找
func (obj *_OrderStateMgr) GetBatchFromPaid(paids []bool) (results []*OrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("paid IN (?)", paids).Find(&results).Error

	return
}

// GetFromPdfInvoice 通过pdf_invoice获取内容
func (obj *_OrderStateMgr) GetFromPdfInvoice(pdfInvoice bool) (results []*OrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("pdf_invoice = ?", pdfInvoice).Find(&results).Error

	return
}

// GetBatchFromPdfInvoice 批量唯一主键查找
func (obj *_OrderStateMgr) GetBatchFromPdfInvoice(pdfInvoices []bool) (results []*OrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("pdf_invoice IN (?)", pdfInvoices).Find(&results).Error

	return
}

// GetFromPdfDelivery 通过pdf_delivery获取内容
func (obj *_OrderStateMgr) GetFromPdfDelivery(pdfDelivery bool) (results []*OrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("pdf_delivery = ?", pdfDelivery).Find(&results).Error

	return
}

// GetBatchFromPdfDelivery 批量唯一主键查找
func (obj *_OrderStateMgr) GetBatchFromPdfDelivery(pdfDeliverys []bool) (results []*OrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("pdf_delivery IN (?)", pdfDeliverys).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容
func (obj *_OrderStateMgr) GetFromDeleted(deleted bool) (results []*OrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("deleted = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量唯一主键查找
func (obj *_OrderStateMgr) GetBatchFromDeleted(deleteds []bool) (results []*OrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("deleted IN (?)", deleteds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_OrderStateMgr) FetchByPrimaryKey(idOrderState uint32) (result OrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_state = ?", idOrderState).Find(&result).Error

	return
}

// FetchIndexByModuleName  获取多个内容
func (obj *_OrderStateMgr) FetchIndexByModuleName(moduleName string) (results []*OrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("module_name = ?", moduleName).Find(&results).Error

	return
}
