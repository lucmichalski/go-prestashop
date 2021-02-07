package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _OrderDetailTaxMgr struct {
	*_BaseMgr
}

func OrderDetailTaxMgr(db *gorm.DB) *_OrderDetailTaxMgr {
	if db == nil {
		panic(fmt.Errorf("OrderDetailTaxMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_OrderDetailTaxMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_order_detail_tax"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_OrderDetailTaxMgr) GetTableName() string {
	return "ps_order_detail_tax"
}

func (obj *_OrderDetailTaxMgr) Get() (result OrderDetailTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_OrderDetailTaxMgr) Gets() (results []*OrderDetailTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

func (obj *_OrderDetailTaxMgr) WithIDOrderDetail(idOrderDetail int) Option {
	return optionFunc(func(o *options) { o.query["id_order_detail"] = idOrderDetail })
}

func (obj *_OrderDetailTaxMgr) WithIDTax(idTax int) Option {
	return optionFunc(func(o *options) { o.query["id_tax"] = idTax })
}

func (obj *_OrderDetailTaxMgr) WithUnitAmount(unitAmount float64) Option {
	return optionFunc(func(o *options) { o.query["unit_amount"] = unitAmount })
}

func (obj *_OrderDetailTaxMgr) WithTotalAmount(totalAmount float64) Option {
	return optionFunc(func(o *options) { o.query["total_amount"] = totalAmount })
}

func (obj *_OrderDetailTaxMgr) GetByOption(opts ...Option) (result OrderDetailTax, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_OrderDetailTaxMgr) GetByOptions(opts ...Option) (results []*OrderDetailTax, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

func (obj *_OrderDetailTaxMgr) GetFromIDOrderDetail(idOrderDetail int) (results []*OrderDetailTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_detail = ?", idOrderDetail).Find(&results).Error

	return
}

func (obj *_OrderDetailTaxMgr) GetBatchFromIDOrderDetail(idOrderDetails []int) (results []*OrderDetailTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_detail IN (?)", idOrderDetails).Find(&results).Error

	return
}

func (obj *_OrderDetailTaxMgr) GetFromIDTax(idTax int) (results []*OrderDetailTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tax = ?", idTax).Find(&results).Error

	return
}

func (obj *_OrderDetailTaxMgr) GetBatchFromIDTax(idTaxs []int) (results []*OrderDetailTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tax IN (?)", idTaxs).Find(&results).Error

	return
}

func (obj *_OrderDetailTaxMgr) GetFromUnitAmount(unitAmount float64) (results []*OrderDetailTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("unit_amount = ?", unitAmount).Find(&results).Error

	return
}

func (obj *_OrderDetailTaxMgr) GetBatchFromUnitAmount(unitAmounts []float64) (results []*OrderDetailTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("unit_amount IN (?)", unitAmounts).Find(&results).Error

	return
}

func (obj *_OrderDetailTaxMgr) GetFromTotalAmount(totalAmount float64) (results []*OrderDetailTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_amount = ?", totalAmount).Find(&results).Error

	return
}

func (obj *_OrderDetailTaxMgr) GetBatchFromTotalAmount(totalAmounts []float64) (results []*OrderDetailTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_amount IN (?)", totalAmounts).Find(&results).Error

	return
}

func (obj *_OrderDetailTaxMgr) FetchIndexByIDOrderDetail(idOrderDetail int) (results []*OrderDetailTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_detail = ?", idOrderDetail).Find(&results).Error

	return
}

func (obj *_OrderDetailTaxMgr) FetchIndexByIDTax(idTax int) (results []*OrderDetailTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tax = ?", idTax).Find(&results).Error

	return
}
