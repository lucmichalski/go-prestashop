package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _SupplyOrderStateMgr struct {
	*_BaseMgr
}

// SupplyOrderStateMgr open func
func SupplyOrderStateMgr(db *gorm.DB) *_SupplyOrderStateMgr {
	if db == nil {
		panic(fmt.Errorf("SupplyOrderStateMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SupplyOrderStateMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_supply_order_state"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SupplyOrderStateMgr) GetTableName() string {
	return "ps_supply_order_state"
}

// Get 获取
func (obj *_SupplyOrderStateMgr) Get() (result SupplyOrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SupplyOrderStateMgr) Gets() (results []*SupplyOrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDSupplyOrderState id_supply_order_state获取
func (obj *_SupplyOrderStateMgr) WithIDSupplyOrderState(idSupplyOrderState uint32) Option {
	return optionFunc(func(o *options) { o.query["id_supply_order_state"] = idSupplyOrderState })
}

// WithDeliveryNote delivery_note获取
func (obj *_SupplyOrderStateMgr) WithDeliveryNote(deliveryNote bool) Option {
	return optionFunc(func(o *options) { o.query["delivery_note"] = deliveryNote })
}

// WithEditable editable获取
func (obj *_SupplyOrderStateMgr) WithEditable(editable bool) Option {
	return optionFunc(func(o *options) { o.query["editable"] = editable })
}

// WithReceiptState receipt_state获取
func (obj *_SupplyOrderStateMgr) WithReceiptState(receiptState bool) Option {
	return optionFunc(func(o *options) { o.query["receipt_state"] = receiptState })
}

// WithPendingReceipt pending_receipt获取
func (obj *_SupplyOrderStateMgr) WithPendingReceipt(pendingReceipt bool) Option {
	return optionFunc(func(o *options) { o.query["pending_receipt"] = pendingReceipt })
}

// WithEnclosed enclosed获取
func (obj *_SupplyOrderStateMgr) WithEnclosed(enclosed bool) Option {
	return optionFunc(func(o *options) { o.query["enclosed"] = enclosed })
}

// WithColor color获取
func (obj *_SupplyOrderStateMgr) WithColor(color string) Option {
	return optionFunc(func(o *options) { o.query["color"] = color })
}

// GetByOption 功能选项模式获取
func (obj *_SupplyOrderStateMgr) GetByOption(opts ...Option) (result SupplyOrderState, err error) {
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
func (obj *_SupplyOrderStateMgr) GetByOptions(opts ...Option) (results []*SupplyOrderState, err error) {
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

// GetFromIDSupplyOrderState 通过id_supply_order_state获取内容
func (obj *_SupplyOrderStateMgr) GetFromIDSupplyOrderState(idSupplyOrderState uint32) (result SupplyOrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supply_order_state = ?", idSupplyOrderState).Find(&result).Error

	return
}

// GetBatchFromIDSupplyOrderState 批量唯一主键查找
func (obj *_SupplyOrderStateMgr) GetBatchFromIDSupplyOrderState(idSupplyOrderStates []uint32) (results []*SupplyOrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supply_order_state IN (?)", idSupplyOrderStates).Find(&results).Error

	return
}

// GetFromDeliveryNote 通过delivery_note获取内容
func (obj *_SupplyOrderStateMgr) GetFromDeliveryNote(deliveryNote bool) (results []*SupplyOrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("delivery_note = ?", deliveryNote).Find(&results).Error

	return
}

// GetBatchFromDeliveryNote 批量唯一主键查找
func (obj *_SupplyOrderStateMgr) GetBatchFromDeliveryNote(deliveryNotes []bool) (results []*SupplyOrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("delivery_note IN (?)", deliveryNotes).Find(&results).Error

	return
}

// GetFromEditable 通过editable获取内容
func (obj *_SupplyOrderStateMgr) GetFromEditable(editable bool) (results []*SupplyOrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("editable = ?", editable).Find(&results).Error

	return
}

// GetBatchFromEditable 批量唯一主键查找
func (obj *_SupplyOrderStateMgr) GetBatchFromEditable(editables []bool) (results []*SupplyOrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("editable IN (?)", editables).Find(&results).Error

	return
}

// GetFromReceiptState 通过receipt_state获取内容
func (obj *_SupplyOrderStateMgr) GetFromReceiptState(receiptState bool) (results []*SupplyOrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("receipt_state = ?", receiptState).Find(&results).Error

	return
}

// GetBatchFromReceiptState 批量唯一主键查找
func (obj *_SupplyOrderStateMgr) GetBatchFromReceiptState(receiptStates []bool) (results []*SupplyOrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("receipt_state IN (?)", receiptStates).Find(&results).Error

	return
}

// GetFromPendingReceipt 通过pending_receipt获取内容
func (obj *_SupplyOrderStateMgr) GetFromPendingReceipt(pendingReceipt bool) (results []*SupplyOrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("pending_receipt = ?", pendingReceipt).Find(&results).Error

	return
}

// GetBatchFromPendingReceipt 批量唯一主键查找
func (obj *_SupplyOrderStateMgr) GetBatchFromPendingReceipt(pendingReceipts []bool) (results []*SupplyOrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("pending_receipt IN (?)", pendingReceipts).Find(&results).Error

	return
}

// GetFromEnclosed 通过enclosed获取内容
func (obj *_SupplyOrderStateMgr) GetFromEnclosed(enclosed bool) (results []*SupplyOrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("enclosed = ?", enclosed).Find(&results).Error

	return
}

// GetBatchFromEnclosed 批量唯一主键查找
func (obj *_SupplyOrderStateMgr) GetBatchFromEnclosed(encloseds []bool) (results []*SupplyOrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("enclosed IN (?)", encloseds).Find(&results).Error

	return
}

// GetFromColor 通过color获取内容
func (obj *_SupplyOrderStateMgr) GetFromColor(color string) (results []*SupplyOrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("color = ?", color).Find(&results).Error

	return
}

// GetBatchFromColor 批量唯一主键查找
func (obj *_SupplyOrderStateMgr) GetBatchFromColor(colors []string) (results []*SupplyOrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("color IN (?)", colors).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_SupplyOrderStateMgr) FetchByPrimaryKey(idSupplyOrderState uint32) (result SupplyOrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supply_order_state = ?", idSupplyOrderState).Find(&result).Error

	return
}
