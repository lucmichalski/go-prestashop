package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _OrderCarrierMgr struct {
	*_BaseMgr
}

// OrderCarrierMgr open func
func OrderCarrierMgr(db *gorm.DB) *_OrderCarrierMgr {
	if db == nil {
		panic(fmt.Errorf("OrderCarrierMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_OrderCarrierMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_order_carrier"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_OrderCarrierMgr) GetTableName() string {
	return "ps_order_carrier"
}

// Get 获取
func (obj *_OrderCarrierMgr) Get() (result OrderCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_OrderCarrierMgr) Gets() (results []*OrderCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDOrderCarrier id_order_carrier获取
func (obj *_OrderCarrierMgr) WithIDOrderCarrier(idOrderCarrier int) Option {
	return optionFunc(func(o *options) { o.query["id_order_carrier"] = idOrderCarrier })
}

// WithIDOrder id_order获取
func (obj *_OrderCarrierMgr) WithIDOrder(idOrder uint32) Option {
	return optionFunc(func(o *options) { o.query["id_order"] = idOrder })
}

// WithIDCarrier id_carrier获取
func (obj *_OrderCarrierMgr) WithIDCarrier(idCarrier uint32) Option {
	return optionFunc(func(o *options) { o.query["id_carrier"] = idCarrier })
}

// WithIDOrderInvoice id_order_invoice获取
func (obj *_OrderCarrierMgr) WithIDOrderInvoice(idOrderInvoice uint32) Option {
	return optionFunc(func(o *options) { o.query["id_order_invoice"] = idOrderInvoice })
}

// WithWeight weight获取
func (obj *_OrderCarrierMgr) WithWeight(weight float64) Option {
	return optionFunc(func(o *options) { o.query["weight"] = weight })
}

// WithShippingCostTaxExcl shipping_cost_tax_excl获取
func (obj *_OrderCarrierMgr) WithShippingCostTaxExcl(shippingCostTaxExcl float64) Option {
	return optionFunc(func(o *options) { o.query["shipping_cost_tax_excl"] = shippingCostTaxExcl })
}

// WithShippingCostTaxIncl shipping_cost_tax_incl获取
func (obj *_OrderCarrierMgr) WithShippingCostTaxIncl(shippingCostTaxIncl float64) Option {
	return optionFunc(func(o *options) { o.query["shipping_cost_tax_incl"] = shippingCostTaxIncl })
}

// WithTrackingNumber tracking_number获取
func (obj *_OrderCarrierMgr) WithTrackingNumber(trackingNumber string) Option {
	return optionFunc(func(o *options) { o.query["tracking_number"] = trackingNumber })
}

// WithDateAdd date_add获取
func (obj *_OrderCarrierMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

// GetByOption 功能选项模式获取
func (obj *_OrderCarrierMgr) GetByOption(opts ...Option) (result OrderCarrier, err error) {
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
func (obj *_OrderCarrierMgr) GetByOptions(opts ...Option) (results []*OrderCarrier, err error) {
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

// GetFromIDOrderCarrier 通过id_order_carrier获取内容
func (obj *_OrderCarrierMgr) GetFromIDOrderCarrier(idOrderCarrier int) (result OrderCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_carrier = ?", idOrderCarrier).Find(&result).Error

	return
}

// GetBatchFromIDOrderCarrier 批量唯一主键查找
func (obj *_OrderCarrierMgr) GetBatchFromIDOrderCarrier(idOrderCarriers []int) (results []*OrderCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_carrier IN (?)", idOrderCarriers).Find(&results).Error

	return
}

// GetFromIDOrder 通过id_order获取内容
func (obj *_OrderCarrierMgr) GetFromIDOrder(idOrder uint32) (results []*OrderCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order = ?", idOrder).Find(&results).Error

	return
}

// GetBatchFromIDOrder 批量唯一主键查找
func (obj *_OrderCarrierMgr) GetBatchFromIDOrder(idOrders []uint32) (results []*OrderCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order IN (?)", idOrders).Find(&results).Error

	return
}

// GetFromIDCarrier 通过id_carrier获取内容
func (obj *_OrderCarrierMgr) GetFromIDCarrier(idCarrier uint32) (results []*OrderCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_carrier = ?", idCarrier).Find(&results).Error

	return
}

// GetBatchFromIDCarrier 批量唯一主键查找
func (obj *_OrderCarrierMgr) GetBatchFromIDCarrier(idCarriers []uint32) (results []*OrderCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_carrier IN (?)", idCarriers).Find(&results).Error

	return
}

// GetFromIDOrderInvoice 通过id_order_invoice获取内容
func (obj *_OrderCarrierMgr) GetFromIDOrderInvoice(idOrderInvoice uint32) (results []*OrderCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_invoice = ?", idOrderInvoice).Find(&results).Error

	return
}

// GetBatchFromIDOrderInvoice 批量唯一主键查找
func (obj *_OrderCarrierMgr) GetBatchFromIDOrderInvoice(idOrderInvoices []uint32) (results []*OrderCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_invoice IN (?)", idOrderInvoices).Find(&results).Error

	return
}

// GetFromWeight 通过weight获取内容
func (obj *_OrderCarrierMgr) GetFromWeight(weight float64) (results []*OrderCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("weight = ?", weight).Find(&results).Error

	return
}

// GetBatchFromWeight 批量唯一主键查找
func (obj *_OrderCarrierMgr) GetBatchFromWeight(weights []float64) (results []*OrderCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("weight IN (?)", weights).Find(&results).Error

	return
}

// GetFromShippingCostTaxExcl 通过shipping_cost_tax_excl获取内容
func (obj *_OrderCarrierMgr) GetFromShippingCostTaxExcl(shippingCostTaxExcl float64) (results []*OrderCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("shipping_cost_tax_excl = ?", shippingCostTaxExcl).Find(&results).Error

	return
}

// GetBatchFromShippingCostTaxExcl 批量唯一主键查找
func (obj *_OrderCarrierMgr) GetBatchFromShippingCostTaxExcl(shippingCostTaxExcls []float64) (results []*OrderCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("shipping_cost_tax_excl IN (?)", shippingCostTaxExcls).Find(&results).Error

	return
}

// GetFromShippingCostTaxIncl 通过shipping_cost_tax_incl获取内容
func (obj *_OrderCarrierMgr) GetFromShippingCostTaxIncl(shippingCostTaxIncl float64) (results []*OrderCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("shipping_cost_tax_incl = ?", shippingCostTaxIncl).Find(&results).Error

	return
}

// GetBatchFromShippingCostTaxIncl 批量唯一主键查找
func (obj *_OrderCarrierMgr) GetBatchFromShippingCostTaxIncl(shippingCostTaxIncls []float64) (results []*OrderCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("shipping_cost_tax_incl IN (?)", shippingCostTaxIncls).Find(&results).Error

	return
}

// GetFromTrackingNumber 通过tracking_number获取内容
func (obj *_OrderCarrierMgr) GetFromTrackingNumber(trackingNumber string) (results []*OrderCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("tracking_number = ?", trackingNumber).Find(&results).Error

	return
}

// GetBatchFromTrackingNumber 批量唯一主键查找
func (obj *_OrderCarrierMgr) GetBatchFromTrackingNumber(trackingNumbers []string) (results []*OrderCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("tracking_number IN (?)", trackingNumbers).Find(&results).Error

	return
}

// GetFromDateAdd 通过date_add获取内容
func (obj *_OrderCarrierMgr) GetFromDateAdd(dateAdd time.Time) (results []*OrderCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error

	return
}

// GetBatchFromDateAdd 批量唯一主键查找
func (obj *_OrderCarrierMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*OrderCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_OrderCarrierMgr) FetchByPrimaryKey(idOrderCarrier int) (result OrderCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_carrier = ?", idOrderCarrier).Find(&result).Error

	return
}

// FetchIndexByIDOrder  获取多个内容
func (obj *_OrderCarrierMgr) FetchIndexByIDOrder(idOrder uint32) (results []*OrderCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order = ?", idOrder).Find(&results).Error

	return
}

// FetchIndexByIDCarrier  获取多个内容
func (obj *_OrderCarrierMgr) FetchIndexByIDCarrier(idCarrier uint32) (results []*OrderCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_carrier = ?", idCarrier).Find(&results).Error

	return
}

// FetchIndexByIDOrderInvoice  获取多个内容
func (obj *_OrderCarrierMgr) FetchIndexByIDOrderInvoice(idOrderInvoice uint32) (results []*OrderCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_invoice = ?", idOrderInvoice).Find(&results).Error

	return
}
