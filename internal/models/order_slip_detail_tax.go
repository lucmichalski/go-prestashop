package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _OrderSlipDetailTaxMgr struct {
	*_BaseMgr
}

func OrderSlipDetailTaxMgr(db *gorm.DB) *_OrderSlipDetailTaxMgr {
	if db == nil {
		panic(fmt.Errorf("OrderSlipDetailTaxMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_OrderSlipDetailTaxMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_order_slip_detail_tax"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_OrderSlipDetailTaxMgr) GetTableName() string {
	return "ps_order_slip_detail_tax"
}

func (obj *_OrderSlipDetailTaxMgr) Get() (result OrderSlipDetailTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_OrderSlipDetailTaxMgr) Gets() (results []*OrderSlipDetailTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_OrderSlipDetailTaxMgr) WithIDOrderSlipDetail(idOrderSlipDetail uint32) Option {
	return optionFunc(func(o *options) { o.query["id_order_slip_detail"] = idOrderSlipDetail })
}

func (obj *_OrderSlipDetailTaxMgr) WithIDTax(idTax uint32) Option {
	return optionFunc(func(o *options) { o.query["id_tax"] = idTax })
}

func (obj *_OrderSlipDetailTaxMgr) WithUnitAmount(unitAmount float64) Option {
	return optionFunc(func(o *options) { o.query["unit_amount"] = unitAmount })
}

func (obj *_OrderSlipDetailTaxMgr) WithTotalAmount(totalAmount float64) Option {
	return optionFunc(func(o *options) { o.query["total_amount"] = totalAmount })
}

func (obj *_OrderSlipDetailTaxMgr) GetByOption(opts ...Option) (result OrderSlipDetailTax, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_OrderSlipDetailTaxMgr) GetByOptions(opts ...Option) (results []*OrderSlipDetailTax, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_OrderSlipDetailTaxMgr) GetFromIDOrderSlipDetail(idOrderSlipDetail uint32) (results []*OrderSlipDetailTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_slip_detail = ?", idOrderSlipDetail).Find(&results).Error

	return
}

func (obj *_OrderSlipDetailTaxMgr) GetBatchFromIDOrderSlipDetail(idOrderSlipDetails []uint32) (results []*OrderSlipDetailTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_slip_detail IN (?)", idOrderSlipDetails).Find(&results).Error

	return
}

func (obj *_OrderSlipDetailTaxMgr) GetFromIDTax(idTax uint32) (results []*OrderSlipDetailTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tax = ?", idTax).Find(&results).Error

	return
}

func (obj *_OrderSlipDetailTaxMgr) GetBatchFromIDTax(idTaxs []uint32) (results []*OrderSlipDetailTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tax IN (?)", idTaxs).Find(&results).Error

	return
}

func (obj *_OrderSlipDetailTaxMgr) GetFromUnitAmount(unitAmount float64) (results []*OrderSlipDetailTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("unit_amount = ?", unitAmount).Find(&results).Error

	return
}

func (obj *_OrderSlipDetailTaxMgr) GetBatchFromUnitAmount(unitAmounts []float64) (results []*OrderSlipDetailTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("unit_amount IN (?)", unitAmounts).Find(&results).Error

	return
}

func (obj *_OrderSlipDetailTaxMgr) GetFromTotalAmount(totalAmount float64) (results []*OrderSlipDetailTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_amount = ?", totalAmount).Find(&results).Error

	return
}

func (obj *_OrderSlipDetailTaxMgr) GetBatchFromTotalAmount(totalAmounts []float64) (results []*OrderSlipDetailTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_amount IN (?)", totalAmounts).Find(&results).Error

	return
}


func (obj *_OrderSlipDetailTaxMgr) FetchIndexByIDOrderSlipDetail(idOrderSlipDetail uint32) (results []*OrderSlipDetailTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_slip_detail = ?", idOrderSlipDetail).Find(&results).Error

	return
}

func (obj *_OrderSlipDetailTaxMgr) FetchIndexByIDTax(idTax uint32) (results []*OrderSlipDetailTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tax = ?", idTax).Find(&results).Error

	return
}
