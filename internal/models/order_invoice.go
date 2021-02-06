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

// OrderInvoiceMgr open func
func OrderInvoiceMgr(db *gorm.DB) *_OrderInvoiceMgr {
	if db == nil {
		panic(fmt.Errorf("OrderInvoiceMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_OrderInvoiceMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_order_invoice"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_OrderInvoiceMgr) GetTableName() string {
	return "ps_order_invoice"
}

// Get 获取
func (obj *_OrderInvoiceMgr) Get() (result OrderInvoice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_OrderInvoiceMgr) Gets() (results []*OrderInvoice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDOrderInvoice id_order_invoice获取
func (obj *_OrderInvoiceMgr) WithIDOrderInvoice(idOrderInvoice uint32) Option {
	return optionFunc(func(o *options) { o.query["id_order_invoice"] = idOrderInvoice })
}

// WithIDOrder id_order获取
func (obj *_OrderInvoiceMgr) WithIDOrder(idOrder int) Option {
	return optionFunc(func(o *options) { o.query["id_order"] = idOrder })
}

// WithNumber number获取
func (obj *_OrderInvoiceMgr) WithNumber(number int) Option {
	return optionFunc(func(o *options) { o.query["number"] = number })
}

// WithDeliveryNumber delivery_number获取
func (obj *_OrderInvoiceMgr) WithDeliveryNumber(deliveryNumber int) Option {
	return optionFunc(func(o *options) { o.query["delivery_number"] = deliveryNumber })
}

// WithDeliveryDate delivery_date获取
func (obj *_OrderInvoiceMgr) WithDeliveryDate(deliveryDate time.Time) Option {
	return optionFunc(func(o *options) { o.query["delivery_date"] = deliveryDate })
}

// WithTotalDiscountTaxExcl total_discount_tax_excl获取
func (obj *_OrderInvoiceMgr) WithTotalDiscountTaxExcl(totalDiscountTaxExcl float64) Option {
	return optionFunc(func(o *options) { o.query["total_discount_tax_excl"] = totalDiscountTaxExcl })
}

// WithTotalDiscountTaxIncl total_discount_tax_incl获取
func (obj *_OrderInvoiceMgr) WithTotalDiscountTaxIncl(totalDiscountTaxIncl float64) Option {
	return optionFunc(func(o *options) { o.query["total_discount_tax_incl"] = totalDiscountTaxIncl })
}

// WithTotalPaidTaxExcl total_paid_tax_excl获取
func (obj *_OrderInvoiceMgr) WithTotalPaidTaxExcl(totalPaidTaxExcl float64) Option {
	return optionFunc(func(o *options) { o.query["total_paid_tax_excl"] = totalPaidTaxExcl })
}

// WithTotalPaidTaxIncl total_paid_tax_incl获取
func (obj *_OrderInvoiceMgr) WithTotalPaidTaxIncl(totalPaidTaxIncl float64) Option {
	return optionFunc(func(o *options) { o.query["total_paid_tax_incl"] = totalPaidTaxIncl })
}

// WithTotalProducts total_products获取
func (obj *_OrderInvoiceMgr) WithTotalProducts(totalProducts float64) Option {
	return optionFunc(func(o *options) { o.query["total_products"] = totalProducts })
}

// WithTotalProductsWt total_products_wt获取
func (obj *_OrderInvoiceMgr) WithTotalProductsWt(totalProductsWt float64) Option {
	return optionFunc(func(o *options) { o.query["total_products_wt"] = totalProductsWt })
}

// WithTotalShippingTaxExcl total_shipping_tax_excl获取
func (obj *_OrderInvoiceMgr) WithTotalShippingTaxExcl(totalShippingTaxExcl float64) Option {
	return optionFunc(func(o *options) { o.query["total_shipping_tax_excl"] = totalShippingTaxExcl })
}

// WithTotalShippingTaxIncl total_shipping_tax_incl获取
func (obj *_OrderInvoiceMgr) WithTotalShippingTaxIncl(totalShippingTaxIncl float64) Option {
	return optionFunc(func(o *options) { o.query["total_shipping_tax_incl"] = totalShippingTaxIncl })
}

// WithShippingTaxComputationMethod shipping_tax_computation_method获取
func (obj *_OrderInvoiceMgr) WithShippingTaxComputationMethod(shippingTaxComputationMethod uint32) Option {
	return optionFunc(func(o *options) { o.query["shipping_tax_computation_method"] = shippingTaxComputationMethod })
}

// WithTotalWrappingTaxExcl total_wrapping_tax_excl获取
func (obj *_OrderInvoiceMgr) WithTotalWrappingTaxExcl(totalWrappingTaxExcl float64) Option {
	return optionFunc(func(o *options) { o.query["total_wrapping_tax_excl"] = totalWrappingTaxExcl })
}

// WithTotalWrappingTaxIncl total_wrapping_tax_incl获取
func (obj *_OrderInvoiceMgr) WithTotalWrappingTaxIncl(totalWrappingTaxIncl float64) Option {
	return optionFunc(func(o *options) { o.query["total_wrapping_tax_incl"] = totalWrappingTaxIncl })
}

// WithShopAddress shop_address获取
func (obj *_OrderInvoiceMgr) WithShopAddress(shopAddress string) Option {
	return optionFunc(func(o *options) { o.query["shop_address"] = shopAddress })
}

// WithNote note获取
func (obj *_OrderInvoiceMgr) WithNote(note string) Option {
	return optionFunc(func(o *options) { o.query["note"] = note })
}

// WithDateAdd date_add获取
func (obj *_OrderInvoiceMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

// GetByOption 功能选项模式获取
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

// GetByOptions 批量功能选项模式获取
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

//////////////////////////enume case ////////////////////////////////////////////

// GetFromIDOrderInvoice 通过id_order_invoice获取内容
func (obj *_OrderInvoiceMgr) GetFromIDOrderInvoice(idOrderInvoice uint32) (result OrderInvoice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_invoice = ?", idOrderInvoice).Find(&result).Error

	return
}

// GetBatchFromIDOrderInvoice 批量唯一主键查找
func (obj *_OrderInvoiceMgr) GetBatchFromIDOrderInvoice(idOrderInvoices []uint32) (results []*OrderInvoice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_invoice IN (?)", idOrderInvoices).Find(&results).Error

	return
}

// GetFromIDOrder 通过id_order获取内容
func (obj *_OrderInvoiceMgr) GetFromIDOrder(idOrder int) (results []*OrderInvoice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order = ?", idOrder).Find(&results).Error

	return
}

// GetBatchFromIDOrder 批量唯一主键查找
func (obj *_OrderInvoiceMgr) GetBatchFromIDOrder(idOrders []int) (results []*OrderInvoice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order IN (?)", idOrders).Find(&results).Error

	return
}

// GetFromNumber 通过number获取内容
func (obj *_OrderInvoiceMgr) GetFromNumber(number int) (results []*OrderInvoice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("number = ?", number).Find(&results).Error

	return
}

// GetBatchFromNumber 批量唯一主键查找
func (obj *_OrderInvoiceMgr) GetBatchFromNumber(numbers []int) (results []*OrderInvoice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("number IN (?)", numbers).Find(&results).Error

	return
}

// GetFromDeliveryNumber 通过delivery_number获取内容
func (obj *_OrderInvoiceMgr) GetFromDeliveryNumber(deliveryNumber int) (results []*OrderInvoice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("delivery_number = ?", deliveryNumber).Find(&results).Error

	return
}

// GetBatchFromDeliveryNumber 批量唯一主键查找
func (obj *_OrderInvoiceMgr) GetBatchFromDeliveryNumber(deliveryNumbers []int) (results []*OrderInvoice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("delivery_number IN (?)", deliveryNumbers).Find(&results).Error

	return
}

// GetFromDeliveryDate 通过delivery_date获取内容
func (obj *_OrderInvoiceMgr) GetFromDeliveryDate(deliveryDate time.Time) (results []*OrderInvoice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("delivery_date = ?", deliveryDate).Find(&results).Error

	return
}

// GetBatchFromDeliveryDate 批量唯一主键查找
func (obj *_OrderInvoiceMgr) GetBatchFromDeliveryDate(deliveryDates []time.Time) (results []*OrderInvoice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("delivery_date IN (?)", deliveryDates).Find(&results).Error

	return
}

// GetFromTotalDiscountTaxExcl 通过total_discount_tax_excl获取内容
func (obj *_OrderInvoiceMgr) GetFromTotalDiscountTaxExcl(totalDiscountTaxExcl float64) (results []*OrderInvoice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_discount_tax_excl = ?", totalDiscountTaxExcl).Find(&results).Error

	return
}

// GetBatchFromTotalDiscountTaxExcl 批量唯一主键查找
func (obj *_OrderInvoiceMgr) GetBatchFromTotalDiscountTaxExcl(totalDiscountTaxExcls []float64) (results []*OrderInvoice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_discount_tax_excl IN (?)", totalDiscountTaxExcls).Find(&results).Error

	return
}

// GetFromTotalDiscountTaxIncl 通过total_discount_tax_incl获取内容
func (obj *_OrderInvoiceMgr) GetFromTotalDiscountTaxIncl(totalDiscountTaxIncl float64) (results []*OrderInvoice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_discount_tax_incl = ?", totalDiscountTaxIncl).Find(&results).Error

	return
}

// GetBatchFromTotalDiscountTaxIncl 批量唯一主键查找
func (obj *_OrderInvoiceMgr) GetBatchFromTotalDiscountTaxIncl(totalDiscountTaxIncls []float64) (results []*OrderInvoice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_discount_tax_incl IN (?)", totalDiscountTaxIncls).Find(&results).Error

	return
}

// GetFromTotalPaidTaxExcl 通过total_paid_tax_excl获取内容
func (obj *_OrderInvoiceMgr) GetFromTotalPaidTaxExcl(totalPaidTaxExcl float64) (results []*OrderInvoice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_paid_tax_excl = ?", totalPaidTaxExcl).Find(&results).Error

	return
}

// GetBatchFromTotalPaidTaxExcl 批量唯一主键查找
func (obj *_OrderInvoiceMgr) GetBatchFromTotalPaidTaxExcl(totalPaidTaxExcls []float64) (results []*OrderInvoice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_paid_tax_excl IN (?)", totalPaidTaxExcls).Find(&results).Error

	return
}

// GetFromTotalPaidTaxIncl 通过total_paid_tax_incl获取内容
func (obj *_OrderInvoiceMgr) GetFromTotalPaidTaxIncl(totalPaidTaxIncl float64) (results []*OrderInvoice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_paid_tax_incl = ?", totalPaidTaxIncl).Find(&results).Error

	return
}

// GetBatchFromTotalPaidTaxIncl 批量唯一主键查找
func (obj *_OrderInvoiceMgr) GetBatchFromTotalPaidTaxIncl(totalPaidTaxIncls []float64) (results []*OrderInvoice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_paid_tax_incl IN (?)", totalPaidTaxIncls).Find(&results).Error

	return
}

// GetFromTotalProducts 通过total_products获取内容
func (obj *_OrderInvoiceMgr) GetFromTotalProducts(totalProducts float64) (results []*OrderInvoice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_products = ?", totalProducts).Find(&results).Error

	return
}

// GetBatchFromTotalProducts 批量唯一主键查找
func (obj *_OrderInvoiceMgr) GetBatchFromTotalProducts(totalProductss []float64) (results []*OrderInvoice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_products IN (?)", totalProductss).Find(&results).Error

	return
}

// GetFromTotalProductsWt 通过total_products_wt获取内容
func (obj *_OrderInvoiceMgr) GetFromTotalProductsWt(totalProductsWt float64) (results []*OrderInvoice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_products_wt = ?", totalProductsWt).Find(&results).Error

	return
}

// GetBatchFromTotalProductsWt 批量唯一主键查找
func (obj *_OrderInvoiceMgr) GetBatchFromTotalProductsWt(totalProductsWts []float64) (results []*OrderInvoice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_products_wt IN (?)", totalProductsWts).Find(&results).Error

	return
}

// GetFromTotalShippingTaxExcl 通过total_shipping_tax_excl获取内容
func (obj *_OrderInvoiceMgr) GetFromTotalShippingTaxExcl(totalShippingTaxExcl float64) (results []*OrderInvoice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_shipping_tax_excl = ?", totalShippingTaxExcl).Find(&results).Error

	return
}

// GetBatchFromTotalShippingTaxExcl 批量唯一主键查找
func (obj *_OrderInvoiceMgr) GetBatchFromTotalShippingTaxExcl(totalShippingTaxExcls []float64) (results []*OrderInvoice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_shipping_tax_excl IN (?)", totalShippingTaxExcls).Find(&results).Error

	return
}

// GetFromTotalShippingTaxIncl 通过total_shipping_tax_incl获取内容
func (obj *_OrderInvoiceMgr) GetFromTotalShippingTaxIncl(totalShippingTaxIncl float64) (results []*OrderInvoice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_shipping_tax_incl = ?", totalShippingTaxIncl).Find(&results).Error

	return
}

// GetBatchFromTotalShippingTaxIncl 批量唯一主键查找
func (obj *_OrderInvoiceMgr) GetBatchFromTotalShippingTaxIncl(totalShippingTaxIncls []float64) (results []*OrderInvoice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_shipping_tax_incl IN (?)", totalShippingTaxIncls).Find(&results).Error

	return
}

// GetFromShippingTaxComputationMethod 通过shipping_tax_computation_method获取内容
func (obj *_OrderInvoiceMgr) GetFromShippingTaxComputationMethod(shippingTaxComputationMethod uint32) (results []*OrderInvoice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("shipping_tax_computation_method = ?", shippingTaxComputationMethod).Find(&results).Error

	return
}

// GetBatchFromShippingTaxComputationMethod 批量唯一主键查找
func (obj *_OrderInvoiceMgr) GetBatchFromShippingTaxComputationMethod(shippingTaxComputationMethods []uint32) (results []*OrderInvoice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("shipping_tax_computation_method IN (?)", shippingTaxComputationMethods).Find(&results).Error

	return
}

// GetFromTotalWrappingTaxExcl 通过total_wrapping_tax_excl获取内容
func (obj *_OrderInvoiceMgr) GetFromTotalWrappingTaxExcl(totalWrappingTaxExcl float64) (results []*OrderInvoice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_wrapping_tax_excl = ?", totalWrappingTaxExcl).Find(&results).Error

	return
}

// GetBatchFromTotalWrappingTaxExcl 批量唯一主键查找
func (obj *_OrderInvoiceMgr) GetBatchFromTotalWrappingTaxExcl(totalWrappingTaxExcls []float64) (results []*OrderInvoice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_wrapping_tax_excl IN (?)", totalWrappingTaxExcls).Find(&results).Error

	return
}

// GetFromTotalWrappingTaxIncl 通过total_wrapping_tax_incl获取内容
func (obj *_OrderInvoiceMgr) GetFromTotalWrappingTaxIncl(totalWrappingTaxIncl float64) (results []*OrderInvoice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_wrapping_tax_incl = ?", totalWrappingTaxIncl).Find(&results).Error

	return
}

// GetBatchFromTotalWrappingTaxIncl 批量唯一主键查找
func (obj *_OrderInvoiceMgr) GetBatchFromTotalWrappingTaxIncl(totalWrappingTaxIncls []float64) (results []*OrderInvoice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_wrapping_tax_incl IN (?)", totalWrappingTaxIncls).Find(&results).Error

	return
}

// GetFromShopAddress 通过shop_address获取内容
func (obj *_OrderInvoiceMgr) GetFromShopAddress(shopAddress string) (results []*OrderInvoice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("shop_address = ?", shopAddress).Find(&results).Error

	return
}

// GetBatchFromShopAddress 批量唯一主键查找
func (obj *_OrderInvoiceMgr) GetBatchFromShopAddress(shopAddresss []string) (results []*OrderInvoice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("shop_address IN (?)", shopAddresss).Find(&results).Error

	return
}

// GetFromNote 通过note获取内容
func (obj *_OrderInvoiceMgr) GetFromNote(note string) (results []*OrderInvoice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("note = ?", note).Find(&results).Error

	return
}

// GetBatchFromNote 批量唯一主键查找
func (obj *_OrderInvoiceMgr) GetBatchFromNote(notes []string) (results []*OrderInvoice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("note IN (?)", notes).Find(&results).Error

	return
}

// GetFromDateAdd 通过date_add获取内容
func (obj *_OrderInvoiceMgr) GetFromDateAdd(dateAdd time.Time) (results []*OrderInvoice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error

	return
}

// GetBatchFromDateAdd 批量唯一主键查找
func (obj *_OrderInvoiceMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*OrderInvoice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_OrderInvoiceMgr) FetchByPrimaryKey(idOrderInvoice uint32) (result OrderInvoice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_invoice = ?", idOrderInvoice).Find(&result).Error

	return
}

// FetchIndexByIDOrder  获取多个内容
func (obj *_OrderInvoiceMgr) FetchIndexByIDOrder(idOrder int) (results []*OrderInvoice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order = ?", idOrder).Find(&results).Error

	return
}
