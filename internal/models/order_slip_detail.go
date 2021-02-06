package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _OrderSlipDetailMgr struct {
	*_BaseMgr
}

func OrderSlipDetailMgr(db *gorm.DB) *_OrderSlipDetailMgr {
	if db == nil {
		panic(fmt.Errorf("OrderSlipDetailMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_OrderSlipDetailMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_order_slip_detail"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_OrderSlipDetailMgr) GetTableName() string {
	return "ps_order_slip_detail"
}

func (obj *_OrderSlipDetailMgr) Get() (result OrderSlipDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_OrderSlipDetailMgr) Gets() (results []*OrderSlipDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_OrderSlipDetailMgr) WithIDOrderSlip(idOrderSlip uint32) Option {
	return optionFunc(func(o *options) { o.query["id_order_slip"] = idOrderSlip })
}

func (obj *_OrderSlipDetailMgr) WithIDOrderDetail(idOrderDetail uint32) Option {
	return optionFunc(func(o *options) { o.query["id_order_detail"] = idOrderDetail })
}

func (obj *_OrderSlipDetailMgr) WithProductQuantity(productQuantity uint32) Option {
	return optionFunc(func(o *options) { o.query["product_quantity"] = productQuantity })
}

func (obj *_OrderSlipDetailMgr) WithUnitPriceTaxExcl(unitPriceTaxExcl float64) Option {
	return optionFunc(func(o *options) { o.query["unit_price_tax_excl"] = unitPriceTaxExcl })
}

func (obj *_OrderSlipDetailMgr) WithUnitPriceTaxIncl(unitPriceTaxIncl float64) Option {
	return optionFunc(func(o *options) { o.query["unit_price_tax_incl"] = unitPriceTaxIncl })
}

func (obj *_OrderSlipDetailMgr) WithTotalPriceTaxExcl(totalPriceTaxExcl float64) Option {
	return optionFunc(func(o *options) { o.query["total_price_tax_excl"] = totalPriceTaxExcl })
}

func (obj *_OrderSlipDetailMgr) WithTotalPriceTaxIncl(totalPriceTaxIncl float64) Option {
	return optionFunc(func(o *options) { o.query["total_price_tax_incl"] = totalPriceTaxIncl })
}

func (obj *_OrderSlipDetailMgr) WithAmountTaxExcl(amountTaxExcl float64) Option {
	return optionFunc(func(o *options) { o.query["amount_tax_excl"] = amountTaxExcl })
}

func (obj *_OrderSlipDetailMgr) WithAmountTaxIncl(amountTaxIncl float64) Option {
	return optionFunc(func(o *options) { o.query["amount_tax_incl"] = amountTaxIncl })
}

func (obj *_OrderSlipDetailMgr) GetByOption(opts ...Option) (result OrderSlipDetail, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_OrderSlipDetailMgr) GetByOptions(opts ...Option) (results []*OrderSlipDetail, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_OrderSlipDetailMgr) GetFromIDOrderSlip(idOrderSlip uint32) (results []*OrderSlipDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_slip = ?", idOrderSlip).Find(&results).Error

	return
}

func (obj *_OrderSlipDetailMgr) GetBatchFromIDOrderSlip(idOrderSlips []uint32) (results []*OrderSlipDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_slip IN (?)", idOrderSlips).Find(&results).Error

	return
}

func (obj *_OrderSlipDetailMgr) GetFromIDOrderDetail(idOrderDetail uint32) (results []*OrderSlipDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_detail = ?", idOrderDetail).Find(&results).Error

	return
}

func (obj *_OrderSlipDetailMgr) GetBatchFromIDOrderDetail(idOrderDetails []uint32) (results []*OrderSlipDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_detail IN (?)", idOrderDetails).Find(&results).Error

	return
}

func (obj *_OrderSlipDetailMgr) GetFromProductQuantity(productQuantity uint32) (results []*OrderSlipDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_quantity = ?", productQuantity).Find(&results).Error

	return
}

func (obj *_OrderSlipDetailMgr) GetBatchFromProductQuantity(productQuantitys []uint32) (results []*OrderSlipDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_quantity IN (?)", productQuantitys).Find(&results).Error

	return
}

func (obj *_OrderSlipDetailMgr) GetFromUnitPriceTaxExcl(unitPriceTaxExcl float64) (results []*OrderSlipDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("unit_price_tax_excl = ?", unitPriceTaxExcl).Find(&results).Error

	return
}

func (obj *_OrderSlipDetailMgr) GetBatchFromUnitPriceTaxExcl(unitPriceTaxExcls []float64) (results []*OrderSlipDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("unit_price_tax_excl IN (?)", unitPriceTaxExcls).Find(&results).Error

	return
}

func (obj *_OrderSlipDetailMgr) GetFromUnitPriceTaxIncl(unitPriceTaxIncl float64) (results []*OrderSlipDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("unit_price_tax_incl = ?", unitPriceTaxIncl).Find(&results).Error

	return
}

func (obj *_OrderSlipDetailMgr) GetBatchFromUnitPriceTaxIncl(unitPriceTaxIncls []float64) (results []*OrderSlipDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("unit_price_tax_incl IN (?)", unitPriceTaxIncls).Find(&results).Error

	return
}

func (obj *_OrderSlipDetailMgr) GetFromTotalPriceTaxExcl(totalPriceTaxExcl float64) (results []*OrderSlipDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_price_tax_excl = ?", totalPriceTaxExcl).Find(&results).Error

	return
}

func (obj *_OrderSlipDetailMgr) GetBatchFromTotalPriceTaxExcl(totalPriceTaxExcls []float64) (results []*OrderSlipDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_price_tax_excl IN (?)", totalPriceTaxExcls).Find(&results).Error

	return
}

func (obj *_OrderSlipDetailMgr) GetFromTotalPriceTaxIncl(totalPriceTaxIncl float64) (results []*OrderSlipDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_price_tax_incl = ?", totalPriceTaxIncl).Find(&results).Error

	return
}

func (obj *_OrderSlipDetailMgr) GetBatchFromTotalPriceTaxIncl(totalPriceTaxIncls []float64) (results []*OrderSlipDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_price_tax_incl IN (?)", totalPriceTaxIncls).Find(&results).Error

	return
}

func (obj *_OrderSlipDetailMgr) GetFromAmountTaxExcl(amountTaxExcl float64) (results []*OrderSlipDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("amount_tax_excl = ?", amountTaxExcl).Find(&results).Error

	return
}

func (obj *_OrderSlipDetailMgr) GetBatchFromAmountTaxExcl(amountTaxExcls []float64) (results []*OrderSlipDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("amount_tax_excl IN (?)", amountTaxExcls).Find(&results).Error

	return
}

func (obj *_OrderSlipDetailMgr) GetFromAmountTaxIncl(amountTaxIncl float64) (results []*OrderSlipDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("amount_tax_incl = ?", amountTaxIncl).Find(&results).Error

	return
}

func (obj *_OrderSlipDetailMgr) GetBatchFromAmountTaxIncl(amountTaxIncls []float64) (results []*OrderSlipDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("amount_tax_incl IN (?)", amountTaxIncls).Find(&results).Error

	return
}


func (obj *_OrderSlipDetailMgr) FetchByPrimaryKey(idOrderSlip uint32, idOrderDetail uint32) (result OrderSlipDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_slip = ? AND id_order_detail = ?", idOrderSlip, idOrderDetail).Find(&result).Error

	return
}
