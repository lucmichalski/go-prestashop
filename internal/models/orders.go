package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _OrdersMgr struct {
	*_BaseMgr
}

// OrdersMgr open func
func OrdersMgr(db *gorm.DB) *_OrdersMgr {
	if db == nil {
		panic(fmt.Errorf("OrdersMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_OrdersMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_orders"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_OrdersMgr) GetTableName() string {
	return "eg_orders"
}

// Get 获取
func (obj *_OrdersMgr) Get() (result Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_OrdersMgr) Gets() (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDOrder id_order获取
func (obj *_OrdersMgr) WithIDOrder(idOrder uint32) Option {
	return optionFunc(func(o *options) { o.query["id_order"] = idOrder })
}

// WithReference reference获取
func (obj *_OrdersMgr) WithReference(reference string) Option {
	return optionFunc(func(o *options) { o.query["reference"] = reference })
}

// WithIDShopGroup id_shop_group获取
func (obj *_OrdersMgr) WithIDShopGroup(idShopGroup uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop_group"] = idShopGroup })
}

// WithIDShop id_shop获取
func (obj *_OrdersMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

// WithIDCarrier id_carrier获取
func (obj *_OrdersMgr) WithIDCarrier(idCarrier uint32) Option {
	return optionFunc(func(o *options) { o.query["id_carrier"] = idCarrier })
}

// WithIDLang id_lang获取
func (obj *_OrdersMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

// WithIDCustomer id_customer获取
func (obj *_OrdersMgr) WithIDCustomer(idCustomer uint32) Option {
	return optionFunc(func(o *options) { o.query["id_customer"] = idCustomer })
}

// WithIDCart id_cart获取
func (obj *_OrdersMgr) WithIDCart(idCart uint32) Option {
	return optionFunc(func(o *options) { o.query["id_cart"] = idCart })
}

// WithIDCurrency id_currency获取
func (obj *_OrdersMgr) WithIDCurrency(idCurrency uint32) Option {
	return optionFunc(func(o *options) { o.query["id_currency"] = idCurrency })
}

// WithIDAddressDelivery id_address_delivery获取
func (obj *_OrdersMgr) WithIDAddressDelivery(idAddressDelivery uint32) Option {
	return optionFunc(func(o *options) { o.query["id_address_delivery"] = idAddressDelivery })
}

// WithIDAddressInvoice id_address_invoice获取
func (obj *_OrdersMgr) WithIDAddressInvoice(idAddressInvoice uint32) Option {
	return optionFunc(func(o *options) { o.query["id_address_invoice"] = idAddressInvoice })
}

// WithCurrentState current_state获取
func (obj *_OrdersMgr) WithCurrentState(currentState uint32) Option {
	return optionFunc(func(o *options) { o.query["current_state"] = currentState })
}

// WithSecureKey secure_key获取
func (obj *_OrdersMgr) WithSecureKey(secureKey string) Option {
	return optionFunc(func(o *options) { o.query["secure_key"] = secureKey })
}

// WithPayment payment获取
func (obj *_OrdersMgr) WithPayment(payment string) Option {
	return optionFunc(func(o *options) { o.query["payment"] = payment })
}

// WithConversionRate conversion_rate获取
func (obj *_OrdersMgr) WithConversionRate(conversionRate float64) Option {
	return optionFunc(func(o *options) { o.query["conversion_rate"] = conversionRate })
}

// WithModule module获取
func (obj *_OrdersMgr) WithModule(module string) Option {
	return optionFunc(func(o *options) { o.query["module"] = module })
}

// WithRecyclable recyclable获取
func (obj *_OrdersMgr) WithRecyclable(recyclable bool) Option {
	return optionFunc(func(o *options) { o.query["recyclable"] = recyclable })
}

// WithGift gift获取
func (obj *_OrdersMgr) WithGift(gift bool) Option {
	return optionFunc(func(o *options) { o.query["gift"] = gift })
}

// WithGiftMessage gift_message获取
func (obj *_OrdersMgr) WithGiftMessage(giftMessage string) Option {
	return optionFunc(func(o *options) { o.query["gift_message"] = giftMessage })
}

// WithMobileTheme mobile_theme获取
func (obj *_OrdersMgr) WithMobileTheme(mobileTheme bool) Option {
	return optionFunc(func(o *options) { o.query["mobile_theme"] = mobileTheme })
}

// WithShippingNumber shipping_number获取
func (obj *_OrdersMgr) WithShippingNumber(shippingNumber string) Option {
	return optionFunc(func(o *options) { o.query["shipping_number"] = shippingNumber })
}

// WithTotalDiscounts total_discounts获取
func (obj *_OrdersMgr) WithTotalDiscounts(totalDiscounts float64) Option {
	return optionFunc(func(o *options) { o.query["total_discounts"] = totalDiscounts })
}

// WithTotalDiscountsTaxIncl total_discounts_tax_incl获取
func (obj *_OrdersMgr) WithTotalDiscountsTaxIncl(totalDiscountsTaxIncl float64) Option {
	return optionFunc(func(o *options) { o.query["total_discounts_tax_incl"] = totalDiscountsTaxIncl })
}

// WithTotalDiscountsTaxExcl total_discounts_tax_excl获取
func (obj *_OrdersMgr) WithTotalDiscountsTaxExcl(totalDiscountsTaxExcl float64) Option {
	return optionFunc(func(o *options) { o.query["total_discounts_tax_excl"] = totalDiscountsTaxExcl })
}

// WithTotalPaid total_paid获取
func (obj *_OrdersMgr) WithTotalPaid(totalPaid float64) Option {
	return optionFunc(func(o *options) { o.query["total_paid"] = totalPaid })
}

// WithTotalPaidTaxIncl total_paid_tax_incl获取
func (obj *_OrdersMgr) WithTotalPaidTaxIncl(totalPaidTaxIncl float64) Option {
	return optionFunc(func(o *options) { o.query["total_paid_tax_incl"] = totalPaidTaxIncl })
}

// WithTotalPaidTaxExcl total_paid_tax_excl获取
func (obj *_OrdersMgr) WithTotalPaidTaxExcl(totalPaidTaxExcl float64) Option {
	return optionFunc(func(o *options) { o.query["total_paid_tax_excl"] = totalPaidTaxExcl })
}

// WithTotalPaidReal total_paid_real获取
func (obj *_OrdersMgr) WithTotalPaidReal(totalPaidReal float64) Option {
	return optionFunc(func(o *options) { o.query["total_paid_real"] = totalPaidReal })
}

// WithTotalProducts total_products获取
func (obj *_OrdersMgr) WithTotalProducts(totalProducts float64) Option {
	return optionFunc(func(o *options) { o.query["total_products"] = totalProducts })
}

// WithTotalProductsWt total_products_wt获取
func (obj *_OrdersMgr) WithTotalProductsWt(totalProductsWt float64) Option {
	return optionFunc(func(o *options) { o.query["total_products_wt"] = totalProductsWt })
}

// WithTotalShipping total_shipping获取
func (obj *_OrdersMgr) WithTotalShipping(totalShipping float64) Option {
	return optionFunc(func(o *options) { o.query["total_shipping"] = totalShipping })
}

// WithTotalShippingTaxIncl total_shipping_tax_incl获取
func (obj *_OrdersMgr) WithTotalShippingTaxIncl(totalShippingTaxIncl float64) Option {
	return optionFunc(func(o *options) { o.query["total_shipping_tax_incl"] = totalShippingTaxIncl })
}

// WithTotalShippingTaxExcl total_shipping_tax_excl获取
func (obj *_OrdersMgr) WithTotalShippingTaxExcl(totalShippingTaxExcl float64) Option {
	return optionFunc(func(o *options) { o.query["total_shipping_tax_excl"] = totalShippingTaxExcl })
}

// WithCarrierTaxRate carrier_tax_rate获取
func (obj *_OrdersMgr) WithCarrierTaxRate(carrierTaxRate float64) Option {
	return optionFunc(func(o *options) { o.query["carrier_tax_rate"] = carrierTaxRate })
}

// WithTotalWrapping total_wrapping获取
func (obj *_OrdersMgr) WithTotalWrapping(totalWrapping float64) Option {
	return optionFunc(func(o *options) { o.query["total_wrapping"] = totalWrapping })
}

// WithTotalWrappingTaxIncl total_wrapping_tax_incl获取
func (obj *_OrdersMgr) WithTotalWrappingTaxIncl(totalWrappingTaxIncl float64) Option {
	return optionFunc(func(o *options) { o.query["total_wrapping_tax_incl"] = totalWrappingTaxIncl })
}

// WithTotalWrappingTaxExcl total_wrapping_tax_excl获取
func (obj *_OrdersMgr) WithTotalWrappingTaxExcl(totalWrappingTaxExcl float64) Option {
	return optionFunc(func(o *options) { o.query["total_wrapping_tax_excl"] = totalWrappingTaxExcl })
}

// WithRoundMode round_mode获取
func (obj *_OrdersMgr) WithRoundMode(roundMode bool) Option {
	return optionFunc(func(o *options) { o.query["round_mode"] = roundMode })
}

// WithRoundType round_type获取
func (obj *_OrdersMgr) WithRoundType(roundType bool) Option {
	return optionFunc(func(o *options) { o.query["round_type"] = roundType })
}

// WithInvoiceNumber invoice_number获取
func (obj *_OrdersMgr) WithInvoiceNumber(invoiceNumber uint32) Option {
	return optionFunc(func(o *options) { o.query["invoice_number"] = invoiceNumber })
}

// WithDeliveryNumber delivery_number获取
func (obj *_OrdersMgr) WithDeliveryNumber(deliveryNumber uint32) Option {
	return optionFunc(func(o *options) { o.query["delivery_number"] = deliveryNumber })
}

// WithInvoiceDate invoice_date获取
func (obj *_OrdersMgr) WithInvoiceDate(invoiceDate time.Time) Option {
	return optionFunc(func(o *options) { o.query["invoice_date"] = invoiceDate })
}

// WithDeliveryDate delivery_date获取
func (obj *_OrdersMgr) WithDeliveryDate(deliveryDate time.Time) Option {
	return optionFunc(func(o *options) { o.query["delivery_date"] = deliveryDate })
}

// WithValid valid获取
func (obj *_OrdersMgr) WithValid(valid uint32) Option {
	return optionFunc(func(o *options) { o.query["valid"] = valid })
}

// WithDateAdd date_add获取
func (obj *_OrdersMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

// WithDateUpd date_upd获取
func (obj *_OrdersMgr) WithDateUpd(dateUpd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_upd"] = dateUpd })
}

// GetByOption 功能选项模式获取
func (obj *_OrdersMgr) GetByOption(opts ...Option) (result Orders, err error) {
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
func (obj *_OrdersMgr) GetByOptions(opts ...Option) (results []*Orders, err error) {
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

// GetFromIDOrder 通过id_order获取内容
func (obj *_OrdersMgr) GetFromIDOrder(idOrder uint32) (result Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order = ?", idOrder).Find(&result).Error

	return
}

// GetBatchFromIDOrder 批量唯一主键查找
func (obj *_OrdersMgr) GetBatchFromIDOrder(idOrders []uint32) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order IN (?)", idOrders).Find(&results).Error

	return
}

// GetFromReference 通过reference获取内容
func (obj *_OrdersMgr) GetFromReference(reference string) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reference = ?", reference).Find(&results).Error

	return
}

// GetBatchFromReference 批量唯一主键查找
func (obj *_OrdersMgr) GetBatchFromReference(references []string) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reference IN (?)", references).Find(&results).Error

	return
}

// GetFromIDShopGroup 通过id_shop_group获取内容
func (obj *_OrdersMgr) GetFromIDShopGroup(idShopGroup uint32) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_group = ?", idShopGroup).Find(&results).Error

	return
}

// GetBatchFromIDShopGroup 批量唯一主键查找
func (obj *_OrdersMgr) GetBatchFromIDShopGroup(idShopGroups []uint32) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_group IN (?)", idShopGroups).Find(&results).Error

	return
}

// GetFromIDShop 通过id_shop获取内容
func (obj *_OrdersMgr) GetFromIDShop(idShop uint32) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

// GetBatchFromIDShop 批量唯一主键查找
func (obj *_OrdersMgr) GetBatchFromIDShop(idShops []uint32) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

// GetFromIDCarrier 通过id_carrier获取内容
func (obj *_OrdersMgr) GetFromIDCarrier(idCarrier uint32) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_carrier = ?", idCarrier).Find(&results).Error

	return
}

// GetBatchFromIDCarrier 批量唯一主键查找
func (obj *_OrdersMgr) GetBatchFromIDCarrier(idCarriers []uint32) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_carrier IN (?)", idCarriers).Find(&results).Error

	return
}

// GetFromIDLang 通过id_lang获取内容
func (obj *_OrdersMgr) GetFromIDLang(idLang uint32) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error

	return
}

// GetBatchFromIDLang 批量唯一主键查找
func (obj *_OrdersMgr) GetBatchFromIDLang(idLangs []uint32) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error

	return
}

// GetFromIDCustomer 通过id_customer获取内容
func (obj *_OrdersMgr) GetFromIDCustomer(idCustomer uint32) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer = ?", idCustomer).Find(&results).Error

	return
}

// GetBatchFromIDCustomer 批量唯一主键查找
func (obj *_OrdersMgr) GetBatchFromIDCustomer(idCustomers []uint32) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer IN (?)", idCustomers).Find(&results).Error

	return
}

// GetFromIDCart 通过id_cart获取内容
func (obj *_OrdersMgr) GetFromIDCart(idCart uint32) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart = ?", idCart).Find(&results).Error

	return
}

// GetBatchFromIDCart 批量唯一主键查找
func (obj *_OrdersMgr) GetBatchFromIDCart(idCarts []uint32) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart IN (?)", idCarts).Find(&results).Error

	return
}

// GetFromIDCurrency 通过id_currency获取内容
func (obj *_OrdersMgr) GetFromIDCurrency(idCurrency uint32) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_currency = ?", idCurrency).Find(&results).Error

	return
}

// GetBatchFromIDCurrency 批量唯一主键查找
func (obj *_OrdersMgr) GetBatchFromIDCurrency(idCurrencys []uint32) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_currency IN (?)", idCurrencys).Find(&results).Error

	return
}

// GetFromIDAddressDelivery 通过id_address_delivery获取内容
func (obj *_OrdersMgr) GetFromIDAddressDelivery(idAddressDelivery uint32) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_address_delivery = ?", idAddressDelivery).Find(&results).Error

	return
}

// GetBatchFromIDAddressDelivery 批量唯一主键查找
func (obj *_OrdersMgr) GetBatchFromIDAddressDelivery(idAddressDeliverys []uint32) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_address_delivery IN (?)", idAddressDeliverys).Find(&results).Error

	return
}

// GetFromIDAddressInvoice 通过id_address_invoice获取内容
func (obj *_OrdersMgr) GetFromIDAddressInvoice(idAddressInvoice uint32) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_address_invoice = ?", idAddressInvoice).Find(&results).Error

	return
}

// GetBatchFromIDAddressInvoice 批量唯一主键查找
func (obj *_OrdersMgr) GetBatchFromIDAddressInvoice(idAddressInvoices []uint32) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_address_invoice IN (?)", idAddressInvoices).Find(&results).Error

	return
}

// GetFromCurrentState 通过current_state获取内容
func (obj *_OrdersMgr) GetFromCurrentState(currentState uint32) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("current_state = ?", currentState).Find(&results).Error

	return
}

// GetBatchFromCurrentState 批量唯一主键查找
func (obj *_OrdersMgr) GetBatchFromCurrentState(currentStates []uint32) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("current_state IN (?)", currentStates).Find(&results).Error

	return
}

// GetFromSecureKey 通过secure_key获取内容
func (obj *_OrdersMgr) GetFromSecureKey(secureKey string) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("secure_key = ?", secureKey).Find(&results).Error

	return
}

// GetBatchFromSecureKey 批量唯一主键查找
func (obj *_OrdersMgr) GetBatchFromSecureKey(secureKeys []string) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("secure_key IN (?)", secureKeys).Find(&results).Error

	return
}

// GetFromPayment 通过payment获取内容
func (obj *_OrdersMgr) GetFromPayment(payment string) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("payment = ?", payment).Find(&results).Error

	return
}

// GetBatchFromPayment 批量唯一主键查找
func (obj *_OrdersMgr) GetBatchFromPayment(payments []string) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("payment IN (?)", payments).Find(&results).Error

	return
}

// GetFromConversionRate 通过conversion_rate获取内容
func (obj *_OrdersMgr) GetFromConversionRate(conversionRate float64) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("conversion_rate = ?", conversionRate).Find(&results).Error

	return
}

// GetBatchFromConversionRate 批量唯一主键查找
func (obj *_OrdersMgr) GetBatchFromConversionRate(conversionRates []float64) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("conversion_rate IN (?)", conversionRates).Find(&results).Error

	return
}

// GetFromModule 通过module获取内容
func (obj *_OrdersMgr) GetFromModule(module string) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("module = ?", module).Find(&results).Error

	return
}

// GetBatchFromModule 批量唯一主键查找
func (obj *_OrdersMgr) GetBatchFromModule(modules []string) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("module IN (?)", modules).Find(&results).Error

	return
}

// GetFromRecyclable 通过recyclable获取内容
func (obj *_OrdersMgr) GetFromRecyclable(recyclable bool) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("recyclable = ?", recyclable).Find(&results).Error

	return
}

// GetBatchFromRecyclable 批量唯一主键查找
func (obj *_OrdersMgr) GetBatchFromRecyclable(recyclables []bool) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("recyclable IN (?)", recyclables).Find(&results).Error

	return
}

// GetFromGift 通过gift获取内容
func (obj *_OrdersMgr) GetFromGift(gift bool) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("gift = ?", gift).Find(&results).Error

	return
}

// GetBatchFromGift 批量唯一主键查找
func (obj *_OrdersMgr) GetBatchFromGift(gifts []bool) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("gift IN (?)", gifts).Find(&results).Error

	return
}

// GetFromGiftMessage 通过gift_message获取内容
func (obj *_OrdersMgr) GetFromGiftMessage(giftMessage string) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("gift_message = ?", giftMessage).Find(&results).Error

	return
}

// GetBatchFromGiftMessage 批量唯一主键查找
func (obj *_OrdersMgr) GetBatchFromGiftMessage(giftMessages []string) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("gift_message IN (?)", giftMessages).Find(&results).Error

	return
}

// GetFromMobileTheme 通过mobile_theme获取内容
func (obj *_OrdersMgr) GetFromMobileTheme(mobileTheme bool) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("mobile_theme = ?", mobileTheme).Find(&results).Error

	return
}

// GetBatchFromMobileTheme 批量唯一主键查找
func (obj *_OrdersMgr) GetBatchFromMobileTheme(mobileThemes []bool) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("mobile_theme IN (?)", mobileThemes).Find(&results).Error

	return
}

// GetFromShippingNumber 通过shipping_number获取内容
func (obj *_OrdersMgr) GetFromShippingNumber(shippingNumber string) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("shipping_number = ?", shippingNumber).Find(&results).Error

	return
}

// GetBatchFromShippingNumber 批量唯一主键查找
func (obj *_OrdersMgr) GetBatchFromShippingNumber(shippingNumbers []string) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("shipping_number IN (?)", shippingNumbers).Find(&results).Error

	return
}

// GetFromTotalDiscounts 通过total_discounts获取内容
func (obj *_OrdersMgr) GetFromTotalDiscounts(totalDiscounts float64) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_discounts = ?", totalDiscounts).Find(&results).Error

	return
}

// GetBatchFromTotalDiscounts 批量唯一主键查找
func (obj *_OrdersMgr) GetBatchFromTotalDiscounts(totalDiscountss []float64) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_discounts IN (?)", totalDiscountss).Find(&results).Error

	return
}

// GetFromTotalDiscountsTaxIncl 通过total_discounts_tax_incl获取内容
func (obj *_OrdersMgr) GetFromTotalDiscountsTaxIncl(totalDiscountsTaxIncl float64) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_discounts_tax_incl = ?", totalDiscountsTaxIncl).Find(&results).Error

	return
}

// GetBatchFromTotalDiscountsTaxIncl 批量唯一主键查找
func (obj *_OrdersMgr) GetBatchFromTotalDiscountsTaxIncl(totalDiscountsTaxIncls []float64) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_discounts_tax_incl IN (?)", totalDiscountsTaxIncls).Find(&results).Error

	return
}

// GetFromTotalDiscountsTaxExcl 通过total_discounts_tax_excl获取内容
func (obj *_OrdersMgr) GetFromTotalDiscountsTaxExcl(totalDiscountsTaxExcl float64) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_discounts_tax_excl = ?", totalDiscountsTaxExcl).Find(&results).Error

	return
}

// GetBatchFromTotalDiscountsTaxExcl 批量唯一主键查找
func (obj *_OrdersMgr) GetBatchFromTotalDiscountsTaxExcl(totalDiscountsTaxExcls []float64) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_discounts_tax_excl IN (?)", totalDiscountsTaxExcls).Find(&results).Error

	return
}

// GetFromTotalPaid 通过total_paid获取内容
func (obj *_OrdersMgr) GetFromTotalPaid(totalPaid float64) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_paid = ?", totalPaid).Find(&results).Error

	return
}

// GetBatchFromTotalPaid 批量唯一主键查找
func (obj *_OrdersMgr) GetBatchFromTotalPaid(totalPaids []float64) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_paid IN (?)", totalPaids).Find(&results).Error

	return
}

// GetFromTotalPaidTaxIncl 通过total_paid_tax_incl获取内容
func (obj *_OrdersMgr) GetFromTotalPaidTaxIncl(totalPaidTaxIncl float64) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_paid_tax_incl = ?", totalPaidTaxIncl).Find(&results).Error

	return
}

// GetBatchFromTotalPaidTaxIncl 批量唯一主键查找
func (obj *_OrdersMgr) GetBatchFromTotalPaidTaxIncl(totalPaidTaxIncls []float64) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_paid_tax_incl IN (?)", totalPaidTaxIncls).Find(&results).Error

	return
}

// GetFromTotalPaidTaxExcl 通过total_paid_tax_excl获取内容
func (obj *_OrdersMgr) GetFromTotalPaidTaxExcl(totalPaidTaxExcl float64) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_paid_tax_excl = ?", totalPaidTaxExcl).Find(&results).Error

	return
}

// GetBatchFromTotalPaidTaxExcl 批量唯一主键查找
func (obj *_OrdersMgr) GetBatchFromTotalPaidTaxExcl(totalPaidTaxExcls []float64) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_paid_tax_excl IN (?)", totalPaidTaxExcls).Find(&results).Error

	return
}

// GetFromTotalPaidReal 通过total_paid_real获取内容
func (obj *_OrdersMgr) GetFromTotalPaidReal(totalPaidReal float64) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_paid_real = ?", totalPaidReal).Find(&results).Error

	return
}

// GetBatchFromTotalPaidReal 批量唯一主键查找
func (obj *_OrdersMgr) GetBatchFromTotalPaidReal(totalPaidReals []float64) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_paid_real IN (?)", totalPaidReals).Find(&results).Error

	return
}

// GetFromTotalProducts 通过total_products获取内容
func (obj *_OrdersMgr) GetFromTotalProducts(totalProducts float64) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_products = ?", totalProducts).Find(&results).Error

	return
}

// GetBatchFromTotalProducts 批量唯一主键查找
func (obj *_OrdersMgr) GetBatchFromTotalProducts(totalProductss []float64) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_products IN (?)", totalProductss).Find(&results).Error

	return
}

// GetFromTotalProductsWt 通过total_products_wt获取内容
func (obj *_OrdersMgr) GetFromTotalProductsWt(totalProductsWt float64) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_products_wt = ?", totalProductsWt).Find(&results).Error

	return
}

// GetBatchFromTotalProductsWt 批量唯一主键查找
func (obj *_OrdersMgr) GetBatchFromTotalProductsWt(totalProductsWts []float64) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_products_wt IN (?)", totalProductsWts).Find(&results).Error

	return
}

// GetFromTotalShipping 通过total_shipping获取内容
func (obj *_OrdersMgr) GetFromTotalShipping(totalShipping float64) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_shipping = ?", totalShipping).Find(&results).Error

	return
}

// GetBatchFromTotalShipping 批量唯一主键查找
func (obj *_OrdersMgr) GetBatchFromTotalShipping(totalShippings []float64) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_shipping IN (?)", totalShippings).Find(&results).Error

	return
}

// GetFromTotalShippingTaxIncl 通过total_shipping_tax_incl获取内容
func (obj *_OrdersMgr) GetFromTotalShippingTaxIncl(totalShippingTaxIncl float64) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_shipping_tax_incl = ?", totalShippingTaxIncl).Find(&results).Error

	return
}

// GetBatchFromTotalShippingTaxIncl 批量唯一主键查找
func (obj *_OrdersMgr) GetBatchFromTotalShippingTaxIncl(totalShippingTaxIncls []float64) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_shipping_tax_incl IN (?)", totalShippingTaxIncls).Find(&results).Error

	return
}

// GetFromTotalShippingTaxExcl 通过total_shipping_tax_excl获取内容
func (obj *_OrdersMgr) GetFromTotalShippingTaxExcl(totalShippingTaxExcl float64) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_shipping_tax_excl = ?", totalShippingTaxExcl).Find(&results).Error

	return
}

// GetBatchFromTotalShippingTaxExcl 批量唯一主键查找
func (obj *_OrdersMgr) GetBatchFromTotalShippingTaxExcl(totalShippingTaxExcls []float64) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_shipping_tax_excl IN (?)", totalShippingTaxExcls).Find(&results).Error

	return
}

// GetFromCarrierTaxRate 通过carrier_tax_rate获取内容
func (obj *_OrdersMgr) GetFromCarrierTaxRate(carrierTaxRate float64) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("carrier_tax_rate = ?", carrierTaxRate).Find(&results).Error

	return
}

// GetBatchFromCarrierTaxRate 批量唯一主键查找
func (obj *_OrdersMgr) GetBatchFromCarrierTaxRate(carrierTaxRates []float64) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("carrier_tax_rate IN (?)", carrierTaxRates).Find(&results).Error

	return
}

// GetFromTotalWrapping 通过total_wrapping获取内容
func (obj *_OrdersMgr) GetFromTotalWrapping(totalWrapping float64) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_wrapping = ?", totalWrapping).Find(&results).Error

	return
}

// GetBatchFromTotalWrapping 批量唯一主键查找
func (obj *_OrdersMgr) GetBatchFromTotalWrapping(totalWrappings []float64) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_wrapping IN (?)", totalWrappings).Find(&results).Error

	return
}

// GetFromTotalWrappingTaxIncl 通过total_wrapping_tax_incl获取内容
func (obj *_OrdersMgr) GetFromTotalWrappingTaxIncl(totalWrappingTaxIncl float64) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_wrapping_tax_incl = ?", totalWrappingTaxIncl).Find(&results).Error

	return
}

// GetBatchFromTotalWrappingTaxIncl 批量唯一主键查找
func (obj *_OrdersMgr) GetBatchFromTotalWrappingTaxIncl(totalWrappingTaxIncls []float64) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_wrapping_tax_incl IN (?)", totalWrappingTaxIncls).Find(&results).Error

	return
}

// GetFromTotalWrappingTaxExcl 通过total_wrapping_tax_excl获取内容
func (obj *_OrdersMgr) GetFromTotalWrappingTaxExcl(totalWrappingTaxExcl float64) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_wrapping_tax_excl = ?", totalWrappingTaxExcl).Find(&results).Error

	return
}

// GetBatchFromTotalWrappingTaxExcl 批量唯一主键查找
func (obj *_OrdersMgr) GetBatchFromTotalWrappingTaxExcl(totalWrappingTaxExcls []float64) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_wrapping_tax_excl IN (?)", totalWrappingTaxExcls).Find(&results).Error

	return
}

// GetFromRoundMode 通过round_mode获取内容
func (obj *_OrdersMgr) GetFromRoundMode(roundMode bool) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("round_mode = ?", roundMode).Find(&results).Error

	return
}

// GetBatchFromRoundMode 批量唯一主键查找
func (obj *_OrdersMgr) GetBatchFromRoundMode(roundModes []bool) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("round_mode IN (?)", roundModes).Find(&results).Error

	return
}

// GetFromRoundType 通过round_type获取内容
func (obj *_OrdersMgr) GetFromRoundType(roundType bool) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("round_type = ?", roundType).Find(&results).Error

	return
}

// GetBatchFromRoundType 批量唯一主键查找
func (obj *_OrdersMgr) GetBatchFromRoundType(roundTypes []bool) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("round_type IN (?)", roundTypes).Find(&results).Error

	return
}

// GetFromInvoiceNumber 通过invoice_number获取内容
func (obj *_OrdersMgr) GetFromInvoiceNumber(invoiceNumber uint32) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("invoice_number = ?", invoiceNumber).Find(&results).Error

	return
}

// GetBatchFromInvoiceNumber 批量唯一主键查找
func (obj *_OrdersMgr) GetBatchFromInvoiceNumber(invoiceNumbers []uint32) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("invoice_number IN (?)", invoiceNumbers).Find(&results).Error

	return
}

// GetFromDeliveryNumber 通过delivery_number获取内容
func (obj *_OrdersMgr) GetFromDeliveryNumber(deliveryNumber uint32) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("delivery_number = ?", deliveryNumber).Find(&results).Error

	return
}

// GetBatchFromDeliveryNumber 批量唯一主键查找
func (obj *_OrdersMgr) GetBatchFromDeliveryNumber(deliveryNumbers []uint32) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("delivery_number IN (?)", deliveryNumbers).Find(&results).Error

	return
}

// GetFromInvoiceDate 通过invoice_date获取内容
func (obj *_OrdersMgr) GetFromInvoiceDate(invoiceDate time.Time) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("invoice_date = ?", invoiceDate).Find(&results).Error

	return
}

// GetBatchFromInvoiceDate 批量唯一主键查找
func (obj *_OrdersMgr) GetBatchFromInvoiceDate(invoiceDates []time.Time) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("invoice_date IN (?)", invoiceDates).Find(&results).Error

	return
}

// GetFromDeliveryDate 通过delivery_date获取内容
func (obj *_OrdersMgr) GetFromDeliveryDate(deliveryDate time.Time) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("delivery_date = ?", deliveryDate).Find(&results).Error

	return
}

// GetBatchFromDeliveryDate 批量唯一主键查找
func (obj *_OrdersMgr) GetBatchFromDeliveryDate(deliveryDates []time.Time) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("delivery_date IN (?)", deliveryDates).Find(&results).Error

	return
}

// GetFromValid 通过valid获取内容
func (obj *_OrdersMgr) GetFromValid(valid uint32) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("valid = ?", valid).Find(&results).Error

	return
}

// GetBatchFromValid 批量唯一主键查找
func (obj *_OrdersMgr) GetBatchFromValid(valids []uint32) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("valid IN (?)", valids).Find(&results).Error

	return
}

// GetFromDateAdd 通过date_add获取内容
func (obj *_OrdersMgr) GetFromDateAdd(dateAdd time.Time) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error

	return
}

// GetBatchFromDateAdd 批量唯一主键查找
func (obj *_OrdersMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error

	return
}

// GetFromDateUpd 通过date_upd获取内容
func (obj *_OrdersMgr) GetFromDateUpd(dateUpd time.Time) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd = ?", dateUpd).Find(&results).Error

	return
}

// GetBatchFromDateUpd 批量唯一主键查找
func (obj *_OrdersMgr) GetBatchFromDateUpd(dateUpds []time.Time) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd IN (?)", dateUpds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_OrdersMgr) FetchByPrimaryKey(idOrder uint32) (result Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order = ?", idOrder).Find(&result).Error

	return
}

// FetchIndexByReference  获取多个内容
func (obj *_OrdersMgr) FetchIndexByReference(reference string) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reference = ?", reference).Find(&results).Error

	return
}

// FetchIndexByIDShopGroup  获取多个内容
func (obj *_OrdersMgr) FetchIndexByIDShopGroup(idShopGroup uint32) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_group = ?", idShopGroup).Find(&results).Error

	return
}

// FetchIndexByIDShop  获取多个内容
func (obj *_OrdersMgr) FetchIndexByIDShop(idShop uint32) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

// FetchIndexByIDCarrier  获取多个内容
func (obj *_OrdersMgr) FetchIndexByIDCarrier(idCarrier uint32) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_carrier = ?", idCarrier).Find(&results).Error

	return
}

// FetchIndexByIDLang  获取多个内容
func (obj *_OrdersMgr) FetchIndexByIDLang(idLang uint32) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error

	return
}

// FetchIndexByIDCustomer  获取多个内容
func (obj *_OrdersMgr) FetchIndexByIDCustomer(idCustomer uint32) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer = ?", idCustomer).Find(&results).Error

	return
}

// FetchIndexByIDCart  获取多个内容
func (obj *_OrdersMgr) FetchIndexByIDCart(idCart uint32) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart = ?", idCart).Find(&results).Error

	return
}

// FetchIndexByIDCurrency  获取多个内容
func (obj *_OrdersMgr) FetchIndexByIDCurrency(idCurrency uint32) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_currency = ?", idCurrency).Find(&results).Error

	return
}

// FetchIndexByIDAddressDelivery  获取多个内容
func (obj *_OrdersMgr) FetchIndexByIDAddressDelivery(idAddressDelivery uint32) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_address_delivery = ?", idAddressDelivery).Find(&results).Error

	return
}

// FetchIndexByIDAddressInvoice  获取多个内容
func (obj *_OrdersMgr) FetchIndexByIDAddressInvoice(idAddressInvoice uint32) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_address_invoice = ?", idAddressInvoice).Find(&results).Error

	return
}

// FetchIndexByCurrentState  获取多个内容
func (obj *_OrdersMgr) FetchIndexByCurrentState(currentState uint32) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("current_state = ?", currentState).Find(&results).Error

	return
}

// FetchIndexByInvoiceNumber  获取多个内容
func (obj *_OrdersMgr) FetchIndexByInvoiceNumber(invoiceNumber uint32) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("invoice_number = ?", invoiceNumber).Find(&results).Error

	return
}

// FetchIndexByDateAdd  获取多个内容
func (obj *_OrdersMgr) FetchIndexByDateAdd(dateAdd time.Time) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error

	return
}
