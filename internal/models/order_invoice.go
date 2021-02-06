package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _OrderInvoiceMgr struct {
	*_BaseMgr
}

func OrderInvoiceMgr(db *gorm.DB) *_OrderInvoiceMgr {
	if db == nil {
		panic(fmt.Errorf("OrderInvoiceMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_OrderInvoiceMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_order_invoice"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_OrderInvoiceMgr) GetTableName() string {
	return "ps_order_invoice"
}

func (obj *_OrderInvoiceMgr) Get() (result OrderInvoice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_OrderInvoiceMgr) Gets() (results []*OrderInvoice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_OrderInvoiceMgr) WithIDOrderInvoice(idOrderInvoice uint32) Option {
	return optionFunc(func(o *options) { o.query["id_order_invoice"] = idOrderInvoice })
}

func (obj *_OrderInvoiceMgr) WithIDOrder(idOrder int) Option {
	return optionFunc(func(o *options) { o.query["id_order"] = idOrder })
}

func (obj *_OrderInvoiceMgr) WithNumber(number int) Option {
	return optionFunc(func(o *options) { o.query["number"] = number })
}

func (obj *_OrderInvoiceMgr) WithDeliveryNumber(deliveryNumber int) Option {
	return optionFunc(func(o *options) { o.query["delivery_number"] = deliveryNumber })
}

func (obj *_OrderInvoiceMgr) WithDeliveryDate(deliveryDate time.Time) Option {
	return optionFunc(func(o *options) { o.query["delivery_date"] = deliveryDate })
}

func (obj *_OrderInvoiceMgr) WithTotalDiscountTaxExcl(totalDiscountTaxExcl float64) Option {
	return optionFunc(func(o *options) { o.query["total_discount_tax_excl"] = totalDiscountTaxExcl })
}

func (obj *_OrderInvoiceMgr) WithTotalDiscountTaxIncl(totalDiscountTaxIncl float64) Option {
	return optionFunc(func(o *options) { o.query["total_discount_tax_incl"] = totalDiscountTaxIncl })
}

func (obj *_OrderInvoiceMgr) WithTotalPaidTaxExcl(totalPaidTaxExcl float64) Option {
	return optionFunc(func(o *options) { o.query["total_paid_tax_excl"] = totalPaidTaxExcl })
}

func (obj *_OrderInvoiceMgr) WithTotalPaidTaxIncl(totalPaidTaxIncl float64) Option {
	return optionFunc(func(o *options) { o.query["total_paid_tax_incl"] = totalPaidTaxIncl })
}

func (obj *_OrderInvoiceMgr) WithTotalProducts(totalProducts float64) Option {
	return optionFunc(func(o *options) { o.query["total_products"] = totalProducts })
}

func (obj *_OrderInvoiceMgr) WithTotalProductsWt(totalProductsWt float64) Option {
	return optionFunc(func(o *options) { o.query["total_products_wt"] = totalProductsWt })
}

func (obj *_OrderInvoiceMgr) WithTotalShippingTaxExcl(totalShippingTaxExcl float64) Option {
	return optionFunc(func(o *options) { o.query["total_shipping_tax_excl"] = totalShippingTaxExcl })
}

func (obj *_OrderInvoiceMgr) WithTotalShippingTaxIncl(totalShippingTaxIncl float64) Option {
	return optionFunc(func(o *options) { o.query["total_shipping_tax_incl"] = totalShippingTaxIncl })
}

func (obj *_OrderInvoiceMgr) WithShippingTaxComputationMethod(shippingTaxComputationMethod uint32) Option {
	return optionFunc(func(o *options) { o.query["shipping_tax_computation_method"] = shippingTaxComputationMethod })
}

func (obj *_OrderInvoiceMgr) WithTotalWrappingTaxExcl(totalWrappingTaxExcl float64) Option {
	return optionFunc(func(o *options) { o.query["total_wrapping_tax_excl"] = totalWrappingTaxExcl })
}

func (obj *_OrderInvoiceMgr) WithTotalWrappingTaxIncl(totalWrappingTaxIncl float64) Option {
	return optionFunc(func(o *options) { o.query["total_wrapping_tax_incl"] = totalWrappingTaxIncl })
}

func (obj *_OrderInvoiceMgr) WithShopAddress(shopAddress string) Option {
	return optionFunc(func(o *options) { o.query["shop_address"] = shopAddress })
}

func (obj *_OrderInvoiceMgr) WithNote(note string) Option {
	return optionFunc(func(o *options) { o.query["note"] = note })
}

func (obj *_OrderInvoiceMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

func (obj *_OrderInvoiceMgr) GetByOption(opts ...Option) (result OrderInvoice, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_OrderInvoiceMgr) GetByOptions(opts ...Option) (results []*OrderInvoice, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_OrderInvoiceMgr) GetFromIDOrderInvoice(idOrderInvoice uint32) (result OrderInvoice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_invoice = ?", idOrderInvoice).Find(&result).Error

	return
}

func (obj *_OrderInvoiceMgr) GetBatchFromIDOrderInvoice(idOrderInvoices []uint32) (results []*OrderInvoice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_invoice IN (?)", idOrderInvoices).Find(&results).Error

	return
}

func (obj *_OrderInvoiceMgr) GetFromIDOrder(idOrder int) (results []*OrderInvoice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order = ?", idOrder).Find(&results).Error

	return
}

func (obj *_OrderInvoiceMgr) GetBatchFromIDOrder(idOrders []int) (results []*OrderInvoice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order IN (?)", idOrders).Find(&results).Error

	return
}

func (obj *_OrderInvoiceMgr) GetFromNumber(number int) (results []*OrderInvoice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("number = ?", number).Find(&results).Error

	return
}

func (obj *_OrderInvoiceMgr) GetBatchFromNumber(numbers []int) (results []*OrderInvoice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("number IN (?)", numbers).Find(&results).Error

	return
}

func (obj *_OrderInvoiceMgr) GetFromDeliveryNumber(deliveryNumber int) (results []*OrderInvoice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("delivery_number = ?", deliveryNumber).Find(&results).Error

	return
}

func (obj *_OrderInvoiceMgr) GetBatchFromDeliveryNumber(deliveryNumbers []int) (results []*OrderInvoice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("delivery_number IN (?)", deliveryNumbers).Find(&results).Error

	return
}

func (obj *_OrderInvoiceMgr) GetFromDeliveryDate(deliveryDate time.Time) (results []*OrderInvoice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("delivery_date = ?", deliveryDate).Find(&results).Error

	return
}

func (obj *_OrderInvoiceMgr) GetBatchFromDeliveryDate(deliveryDates []time.Time) (results []*OrderInvoice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("delivery_date IN (?)", deliveryDates).Find(&results).Error

	return
}

func (obj *_OrderInvoiceMgr) GetFromTotalDiscountTaxExcl(totalDiscountTaxExcl float64) (results []*OrderInvoice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_discount_tax_excl = ?", totalDiscountTaxExcl).Find(&results).Error

	return
}

func (obj *_OrderInvoiceMgr) GetBatchFromTotalDiscountTaxExcl(totalDiscountTaxExcls []float64) (results []*OrderInvoice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_discount_tax_excl IN (?)", totalDiscountTaxExcls).Find(&results).Error

	return
}

func (obj *_OrderInvoiceMgr) GetFromTotalDiscountTaxIncl(totalDiscountTaxIncl float64) (results []*OrderInvoice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_discount_tax_incl = ?", totalDiscountTaxIncl).Find(&results).Error

	return
}

func (obj *_OrderInvoiceMgr) GetBatchFromTotalDiscountTaxIncl(totalDiscountTaxIncls []float64) (results []*OrderInvoice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_discount_tax_incl IN (?)", totalDiscountTaxIncls).Find(&results).Error

	return
}

func (obj *_OrderInvoiceMgr) GetFromTotalPaidTaxExcl(totalPaidTaxExcl float64) (results []*OrderInvoice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_paid_tax_excl = ?", totalPaidTaxExcl).Find(&results).Error

	return
}

func (obj *_OrderInvoiceMgr) GetBatchFromTotalPaidTaxExcl(totalPaidTaxExcls []float64) (results []*OrderInvoice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_paid_tax_excl IN (?)", totalPaidTaxExcls).Find(&results).Error

	return
}

func (obj *_OrderInvoiceMgr) GetFromTotalPaidTaxIncl(totalPaidTaxIncl float64) (results []*OrderInvoice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_paid_tax_incl = ?", totalPaidTaxIncl).Find(&results).Error

	return
}

func (obj *_OrderInvoiceMgr) GetBatchFromTotalPaidTaxIncl(totalPaidTaxIncls []float64) (results []*OrderInvoice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_paid_tax_incl IN (?)", totalPaidTaxIncls).Find(&results).Error

	return
}

func (obj *_OrderInvoiceMgr) GetFromTotalProducts(totalProducts float64) (results []*OrderInvoice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_products = ?", totalProducts).Find(&results).Error

	return
}

func (obj *_OrderInvoiceMgr) GetBatchFromTotalProducts(totalProductss []float64) (results []*OrderInvoice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_products IN (?)", totalProductss).Find(&results).Error

	return
}

func (obj *_OrderInvoiceMgr) GetFromTotalProductsWt(totalProductsWt float64) (results []*OrderInvoice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_products_wt = ?", totalProductsWt).Find(&results).Error

	return
}

func (obj *_OrderInvoiceMgr) GetBatchFromTotalProductsWt(totalProductsWts []float64) (results []*OrderInvoice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_products_wt IN (?)", totalProductsWts).Find(&results).Error

	return
}

func (obj *_OrderInvoiceMgr) GetFromTotalShippingTaxExcl(totalShippingTaxExcl float64) (results []*OrderInvoice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_shipping_tax_excl = ?", totalShippingTaxExcl).Find(&results).Error

	return
}

func (obj *_OrderInvoiceMgr) GetBatchFromTotalShippingTaxExcl(totalShippingTaxExcls []float64) (results []*OrderInvoice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_shipping_tax_excl IN (?)", totalShippingTaxExcls).Find(&results).Error

	return
}

func (obj *_OrderInvoiceMgr) GetFromTotalShippingTaxIncl(totalShippingTaxIncl float64) (results []*OrderInvoice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_shipping_tax_incl = ?", totalShippingTaxIncl).Find(&results).Error

	return
}

func (obj *_OrderInvoiceMgr) GetBatchFromTotalShippingTaxIncl(totalShippingTaxIncls []float64) (results []*OrderInvoice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_shipping_tax_incl IN (?)", totalShippingTaxIncls).Find(&results).Error

	return
}

func (obj *_OrderInvoiceMgr) GetFromShippingTaxComputationMethod(shippingTaxComputationMethod uint32) (results []*OrderInvoice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("shipping_tax_computation_method = ?", shippingTaxComputationMethod).Find(&results).Error

	return
}

func (obj *_OrderInvoiceMgr) GetBatchFromShippingTaxComputationMethod(shippingTaxComputationMethods []uint32) (results []*OrderInvoice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("shipping_tax_computation_method IN (?)", shippingTaxComputationMethods).Find(&results).Error

	return
}

func (obj *_OrderInvoiceMgr) GetFromTotalWrappingTaxExcl(totalWrappingTaxExcl float64) (results []*OrderInvoice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_wrapping_tax_excl = ?", totalWrappingTaxExcl).Find(&results).Error

	return
}

func (obj *_OrderInvoiceMgr) GetBatchFromTotalWrappingTaxExcl(totalWrappingTaxExcls []float64) (results []*OrderInvoice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_wrapping_tax_excl IN (?)", totalWrappingTaxExcls).Find(&results).Error

	return
}

func (obj *_OrderInvoiceMgr) GetFromTotalWrappingTaxIncl(totalWrappingTaxIncl float64) (results []*OrderInvoice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_wrapping_tax_incl = ?", totalWrappingTaxIncl).Find(&results).Error

	return
}

func (obj *_OrderInvoiceMgr) GetBatchFromTotalWrappingTaxIncl(totalWrappingTaxIncls []float64) (results []*OrderInvoice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_wrapping_tax_incl IN (?)", totalWrappingTaxIncls).Find(&results).Error

	return
}

func (obj *_OrderInvoiceMgr) GetFromShopAddress(shopAddress string) (results []*OrderInvoice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("shop_address = ?", shopAddress).Find(&results).Error

	return
}

func (obj *_OrderInvoiceMgr) GetBatchFromShopAddress(shopAddresss []string) (results []*OrderInvoice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("shop_address IN (?)", shopAddresss).Find(&results).Error

	return
}

func (obj *_OrderInvoiceMgr) GetFromNote(note string) (results []*OrderInvoice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("note = ?", note).Find(&results).Error

	return
}

func (obj *_OrderInvoiceMgr) GetBatchFromNote(notes []string) (results []*OrderInvoice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("note IN (?)", notes).Find(&results).Error

	return
}

func (obj *_OrderInvoiceMgr) GetFromDateAdd(dateAdd time.Time) (results []*OrderInvoice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error

	return
}

func (obj *_OrderInvoiceMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*OrderInvoice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error

	return
}


func (obj *_OrderInvoiceMgr) FetchByPrimaryKey(idOrderInvoice uint32) (result OrderInvoice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_invoice = ?", idOrderInvoice).Find(&result).Error

	return
}

func (obj *_OrderInvoiceMgr) FetchIndexByIDOrder(idOrder int) (results []*OrderInvoice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order = ?", idOrder).Find(&results).Error

	return
}
