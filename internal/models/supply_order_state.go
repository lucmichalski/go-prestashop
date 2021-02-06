package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _SupplyOrderStateMgr struct {
	*_BaseMgr
}

func SupplyOrderStateMgr(db *gorm.DB) *_SupplyOrderStateMgr {
	if db == nil {
		panic(fmt.Errorf("SupplyOrderStateMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SupplyOrderStateMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_supply_order_state"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_SupplyOrderStateMgr) GetTableName() string {
	return "ps_supply_order_state"
}

func (obj *_SupplyOrderStateMgr) Get() (result SupplyOrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_SupplyOrderStateMgr) Gets() (results []*SupplyOrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_SupplyOrderStateMgr) WithIDSupplyOrderState(idSupplyOrderState uint32) Option {
	return optionFunc(func(o *options) { o.query["id_supply_order_state"] = idSupplyOrderState })
}

func (obj *_SupplyOrderStateMgr) WithDeliveryNote(deliveryNote bool) Option {
	return optionFunc(func(o *options) { o.query["delivery_note"] = deliveryNote })
}

func (obj *_SupplyOrderStateMgr) WithEditable(editable bool) Option {
	return optionFunc(func(o *options) { o.query["editable"] = editable })
}

func (obj *_SupplyOrderStateMgr) WithReceiptState(receiptState bool) Option {
	return optionFunc(func(o *options) { o.query["receipt_state"] = receiptState })
}

func (obj *_SupplyOrderStateMgr) WithPendingReceipt(pendingReceipt bool) Option {
	return optionFunc(func(o *options) { o.query["pending_receipt"] = pendingReceipt })
}

func (obj *_SupplyOrderStateMgr) WithEnclosed(enclosed bool) Option {
	return optionFunc(func(o *options) { o.query["enclosed"] = enclosed })
}

func (obj *_SupplyOrderStateMgr) WithColor(color string) Option {
	return optionFunc(func(o *options) { o.query["color"] = color })
}

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


func (obj *_SupplyOrderStateMgr) GetFromIDSupplyOrderState(idSupplyOrderState uint32) (result SupplyOrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supply_order_state = ?", idSupplyOrderState).Find(&result).Error

	return
}

func (obj *_SupplyOrderStateMgr) GetBatchFromIDSupplyOrderState(idSupplyOrderStates []uint32) (results []*SupplyOrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supply_order_state IN (?)", idSupplyOrderStates).Find(&results).Error

	return
}

func (obj *_SupplyOrderStateMgr) GetFromDeliveryNote(deliveryNote bool) (results []*SupplyOrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("delivery_note = ?", deliveryNote).Find(&results).Error

	return
}

func (obj *_SupplyOrderStateMgr) GetBatchFromDeliveryNote(deliveryNotes []bool) (results []*SupplyOrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("delivery_note IN (?)", deliveryNotes).Find(&results).Error

	return
}

func (obj *_SupplyOrderStateMgr) GetFromEditable(editable bool) (results []*SupplyOrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("editable = ?", editable).Find(&results).Error

	return
}

func (obj *_SupplyOrderStateMgr) GetBatchFromEditable(editables []bool) (results []*SupplyOrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("editable IN (?)", editables).Find(&results).Error

	return
}

func (obj *_SupplyOrderStateMgr) GetFromReceiptState(receiptState bool) (results []*SupplyOrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("receipt_state = ?", receiptState).Find(&results).Error

	return
}

func (obj *_SupplyOrderStateMgr) GetBatchFromReceiptState(receiptStates []bool) (results []*SupplyOrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("receipt_state IN (?)", receiptStates).Find(&results).Error

	return
}

func (obj *_SupplyOrderStateMgr) GetFromPendingReceipt(pendingReceipt bool) (results []*SupplyOrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("pending_receipt = ?", pendingReceipt).Find(&results).Error

	return
}

func (obj *_SupplyOrderStateMgr) GetBatchFromPendingReceipt(pendingReceipts []bool) (results []*SupplyOrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("pending_receipt IN (?)", pendingReceipts).Find(&results).Error

	return
}

func (obj *_SupplyOrderStateMgr) GetFromEnclosed(enclosed bool) (results []*SupplyOrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("enclosed = ?", enclosed).Find(&results).Error

	return
}

func (obj *_SupplyOrderStateMgr) GetBatchFromEnclosed(encloseds []bool) (results []*SupplyOrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("enclosed IN (?)", encloseds).Find(&results).Error

	return
}

func (obj *_SupplyOrderStateMgr) GetFromColor(color string) (results []*SupplyOrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("color = ?", color).Find(&results).Error

	return
}

func (obj *_SupplyOrderStateMgr) GetBatchFromColor(colors []string) (results []*SupplyOrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("color IN (?)", colors).Find(&results).Error

	return
}


func (obj *_SupplyOrderStateMgr) FetchByPrimaryKey(idSupplyOrderState uint32) (result SupplyOrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supply_order_state = ?", idSupplyOrderState).Find(&result).Error

	return
}
