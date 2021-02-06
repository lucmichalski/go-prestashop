package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
"time"	
)	

type _EgOrdersMgr struct {
	*_BaseMgr
}

// EgOrdersMgr open func
func EgOrdersMgr(db *gorm.DB) *_EgOrdersMgr {
	if db == nil {
		panic(fmt.Errorf("EgOrdersMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgOrdersMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_orders"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgOrdersMgr) GetTableName() string {
	return "eg_orders"
}

// Get 获取
func (obj *_EgOrdersMgr) Get() (result EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgOrdersMgr) Gets() (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDOrder id_order获取 
func (obj *_EgOrdersMgr) WithIDOrder(idOrder uint32) Option {
	return optionFunc(func(o *options) { o.query["id_order"] = idOrder })
}

// WithReference reference获取 
func (obj *_EgOrdersMgr) WithReference(reference string) Option {
	return optionFunc(func(o *options) { o.query["reference"] = reference })
}

// WithIDShopGroup id_shop_group获取 
func (obj *_EgOrdersMgr) WithIDShopGroup(idShopGroup uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop_group"] = idShopGroup })
}

// WithIDShop id_shop获取 
func (obj *_EgOrdersMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

// WithIDCarrier id_carrier获取 
func (obj *_EgOrdersMgr) WithIDCarrier(idCarrier uint32) Option {
	return optionFunc(func(o *options) { o.query["id_carrier"] = idCarrier })
}

// WithIDLang id_lang获取 
func (obj *_EgOrdersMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

// WithIDCustomer id_customer获取 
func (obj *_EgOrdersMgr) WithIDCustomer(idCustomer uint32) Option {
	return optionFunc(func(o *options) { o.query["id_customer"] = idCustomer })
}

// WithIDCart id_cart获取 
func (obj *_EgOrdersMgr) WithIDCart(idCart uint32) Option {
	return optionFunc(func(o *options) { o.query["id_cart"] = idCart })
}

// WithIDCurrency id_currency获取 
func (obj *_EgOrdersMgr) WithIDCurrency(idCurrency uint32) Option {
	return optionFunc(func(o *options) { o.query["id_currency"] = idCurrency })
}

// WithIDAddressDelivery id_address_delivery获取 
func (obj *_EgOrdersMgr) WithIDAddressDelivery(idAddressDelivery uint32) Option {
	return optionFunc(func(o *options) { o.query["id_address_delivery"] = idAddressDelivery })
}

// WithIDAddressInvoice id_address_invoice获取 
func (obj *_EgOrdersMgr) WithIDAddressInvoice(idAddressInvoice uint32) Option {
	return optionFunc(func(o *options) { o.query["id_address_invoice"] = idAddressInvoice })
}

// WithCurrentState current_state获取 
func (obj *_EgOrdersMgr) WithCurrentState(currentState uint32) Option {
	return optionFunc(func(o *options) { o.query["current_state"] = currentState })
}

// WithSecureKey secure_key获取 
func (obj *_EgOrdersMgr) WithSecureKey(secureKey string) Option {
	return optionFunc(func(o *options) { o.query["secure_key"] = secureKey })
}

// WithPayment payment获取 
func (obj *_EgOrdersMgr) WithPayment(payment string) Option {
	return optionFunc(func(o *options) { o.query["payment"] = payment })
}

// WithConversionRate conversion_rate获取 
func (obj *_EgOrdersMgr) WithConversionRate(conversionRate float64) Option {
	return optionFunc(func(o *options) { o.query["conversion_rate"] = conversionRate })
}

// WithModule module获取 
func (obj *_EgOrdersMgr) WithModule(module string) Option {
	return optionFunc(func(o *options) { o.query["module"] = module })
}

// WithRecyclable recyclable获取 
func (obj *_EgOrdersMgr) WithRecyclable(recyclable bool) Option {
	return optionFunc(func(o *options) { o.query["recyclable"] = recyclable })
}

// WithGift gift获取 
func (obj *_EgOrdersMgr) WithGift(gift bool) Option {
	return optionFunc(func(o *options) { o.query["gift"] = gift })
}

// WithGiftMessage gift_message获取 
func (obj *_EgOrdersMgr) WithGiftMessage(giftMessage string) Option {
	return optionFunc(func(o *options) { o.query["gift_message"] = giftMessage })
}

// WithMobileTheme mobile_theme获取 
func (obj *_EgOrdersMgr) WithMobileTheme(mobileTheme bool) Option {
	return optionFunc(func(o *options) { o.query["mobile_theme"] = mobileTheme })
}

// WithShippingNumber shipping_number获取 
func (obj *_EgOrdersMgr) WithShippingNumber(shippingNumber string) Option {
	return optionFunc(func(o *options) { o.query["shipping_number"] = shippingNumber })
}

// WithTotalDiscounts total_discounts获取 
func (obj *_EgOrdersMgr) WithTotalDiscounts(totalDiscounts float64) Option {
	return optionFunc(func(o *options) { o.query["total_discounts"] = totalDiscounts })
}

// WithTotalDiscountsTaxIncl total_discounts_tax_incl获取 
func (obj *_EgOrdersMgr) WithTotalDiscountsTaxIncl(totalDiscountsTaxIncl float64) Option {
	return optionFunc(func(o *options) { o.query["total_discounts_tax_incl"] = totalDiscountsTaxIncl })
}

// WithTotalDiscountsTaxExcl total_discounts_tax_excl获取 
func (obj *_EgOrdersMgr) WithTotalDiscountsTaxExcl(totalDiscountsTaxExcl float64) Option {
	return optionFunc(func(o *options) { o.query["total_discounts_tax_excl"] = totalDiscountsTaxExcl })
}

// WithTotalPaid total_paid获取 
func (obj *_EgOrdersMgr) WithTotalPaid(totalPaid float64) Option {
	return optionFunc(func(o *options) { o.query["total_paid"] = totalPaid })
}

// WithTotalPaidTaxIncl total_paid_tax_incl获取 
func (obj *_EgOrdersMgr) WithTotalPaidTaxIncl(totalPaidTaxIncl float64) Option {
	return optionFunc(func(o *options) { o.query["total_paid_tax_incl"] = totalPaidTaxIncl })
}

// WithTotalPaidTaxExcl total_paid_tax_excl获取 
func (obj *_EgOrdersMgr) WithTotalPaidTaxExcl(totalPaidTaxExcl float64) Option {
	return optionFunc(func(o *options) { o.query["total_paid_tax_excl"] = totalPaidTaxExcl })
}

// WithTotalPaidReal total_paid_real获取 
func (obj *_EgOrdersMgr) WithTotalPaidReal(totalPaidReal float64) Option {
	return optionFunc(func(o *options) { o.query["total_paid_real"] = totalPaidReal })
}

// WithTotalProducts total_products获取 
func (obj *_EgOrdersMgr) WithTotalProducts(totalProducts float64) Option {
	return optionFunc(func(o *options) { o.query["total_products"] = totalProducts })
}

// WithTotalProductsWt total_products_wt获取 
func (obj *_EgOrdersMgr) WithTotalProductsWt(totalProductsWt float64) Option {
	return optionFunc(func(o *options) { o.query["total_products_wt"] = totalProductsWt })
}

// WithTotalShipping total_shipping获取 
func (obj *_EgOrdersMgr) WithTotalShipping(totalShipping float64) Option {
	return optionFunc(func(o *options) { o.query["total_shipping"] = totalShipping })
}

// WithTotalShippingTaxIncl total_shipping_tax_incl获取 
func (obj *_EgOrdersMgr) WithTotalShippingTaxIncl(totalShippingTaxIncl float64) Option {
	return optionFunc(func(o *options) { o.query["total_shipping_tax_incl"] = totalShippingTaxIncl })
}

// WithTotalShippingTaxExcl total_shipping_tax_excl获取 
func (obj *_EgOrdersMgr) WithTotalShippingTaxExcl(totalShippingTaxExcl float64) Option {
	return optionFunc(func(o *options) { o.query["total_shipping_tax_excl"] = totalShippingTaxExcl })
}

// WithCarrierTaxRate carrier_tax_rate获取 
func (obj *_EgOrdersMgr) WithCarrierTaxRate(carrierTaxRate float64) Option {
	return optionFunc(func(o *options) { o.query["carrier_tax_rate"] = carrierTaxRate })
}

// WithTotalWrapping total_wrapping获取 
func (obj *_EgOrdersMgr) WithTotalWrapping(totalWrapping float64) Option {
	return optionFunc(func(o *options) { o.query["total_wrapping"] = totalWrapping })
}

// WithTotalWrappingTaxIncl total_wrapping_tax_incl获取 
func (obj *_EgOrdersMgr) WithTotalWrappingTaxIncl(totalWrappingTaxIncl float64) Option {
	return optionFunc(func(o *options) { o.query["total_wrapping_tax_incl"] = totalWrappingTaxIncl })
}

// WithTotalWrappingTaxExcl total_wrapping_tax_excl获取 
func (obj *_EgOrdersMgr) WithTotalWrappingTaxExcl(totalWrappingTaxExcl float64) Option {
	return optionFunc(func(o *options) { o.query["total_wrapping_tax_excl"] = totalWrappingTaxExcl })
}

// WithRoundMode round_mode获取 
func (obj *_EgOrdersMgr) WithRoundMode(roundMode bool) Option {
	return optionFunc(func(o *options) { o.query["round_mode"] = roundMode })
}

// WithRoundType round_type获取 
func (obj *_EgOrdersMgr) WithRoundType(roundType bool) Option {
	return optionFunc(func(o *options) { o.query["round_type"] = roundType })
}

// WithInvoiceNumber invoice_number获取 
func (obj *_EgOrdersMgr) WithInvoiceNumber(invoiceNumber uint32) Option {
	return optionFunc(func(o *options) { o.query["invoice_number"] = invoiceNumber })
}

// WithDeliveryNumber delivery_number获取 
func (obj *_EgOrdersMgr) WithDeliveryNumber(deliveryNumber uint32) Option {
	return optionFunc(func(o *options) { o.query["delivery_number"] = deliveryNumber })
}

// WithInvoiceDate invoice_date获取 
func (obj *_EgOrdersMgr) WithInvoiceDate(invoiceDate time.Time) Option {
	return optionFunc(func(o *options) { o.query["invoice_date"] = invoiceDate })
}

// WithDeliveryDate delivery_date获取 
func (obj *_EgOrdersMgr) WithDeliveryDate(deliveryDate time.Time) Option {
	return optionFunc(func(o *options) { o.query["delivery_date"] = deliveryDate })
}

// WithValid valid获取 
func (obj *_EgOrdersMgr) WithValid(valid uint32) Option {
	return optionFunc(func(o *options) { o.query["valid"] = valid })
}

// WithDateAdd date_add获取 
func (obj *_EgOrdersMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

// WithDateUpd date_upd获取 
func (obj *_EgOrdersMgr) WithDateUpd(dateUpd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_upd"] = dateUpd })
}


// GetByOption 功能选项模式获取
func (obj *_EgOrdersMgr) GetByOption(opts ...Option) (result EgOrders, err error) {
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
func (obj *_EgOrdersMgr) GetByOptions(opts ...Option) (results []*EgOrders, err error) {
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
func (obj *_EgOrdersMgr)  GetFromIDOrder(idOrder uint32) (result EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order = ?", idOrder).Find(&result).Error
	
	return
}

// GetBatchFromIDOrder 批量唯一主键查找 
func (obj *_EgOrdersMgr) GetBatchFromIDOrder(idOrders []uint32) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order IN (?)", idOrders).Find(&results).Error
	
	return
}
 
// GetFromReference 通过reference获取内容  
func (obj *_EgOrdersMgr) GetFromReference(reference string) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reference = ?", reference).Find(&results).Error
	
	return
}

// GetBatchFromReference 批量唯一主键查找 
func (obj *_EgOrdersMgr) GetBatchFromReference(references []string) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reference IN (?)", references).Find(&results).Error
	
	return
}
 
// GetFromIDShopGroup 通过id_shop_group获取内容  
func (obj *_EgOrdersMgr) GetFromIDShopGroup(idShopGroup uint32) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_group = ?", idShopGroup).Find(&results).Error
	
	return
}

// GetBatchFromIDShopGroup 批量唯一主键查找 
func (obj *_EgOrdersMgr) GetBatchFromIDShopGroup(idShopGroups []uint32) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_group IN (?)", idShopGroups).Find(&results).Error
	
	return
}
 
// GetFromIDShop 通过id_shop获取内容  
func (obj *_EgOrdersMgr) GetFromIDShop(idShop uint32) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error
	
	return
}

// GetBatchFromIDShop 批量唯一主键查找 
func (obj *_EgOrdersMgr) GetBatchFromIDShop(idShops []uint32) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error
	
	return
}
 
// GetFromIDCarrier 通过id_carrier获取内容  
func (obj *_EgOrdersMgr) GetFromIDCarrier(idCarrier uint32) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_carrier = ?", idCarrier).Find(&results).Error
	
	return
}

// GetBatchFromIDCarrier 批量唯一主键查找 
func (obj *_EgOrdersMgr) GetBatchFromIDCarrier(idCarriers []uint32) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_carrier IN (?)", idCarriers).Find(&results).Error
	
	return
}
 
// GetFromIDLang 通过id_lang获取内容  
func (obj *_EgOrdersMgr) GetFromIDLang(idLang uint32) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error
	
	return
}

// GetBatchFromIDLang 批量唯一主键查找 
func (obj *_EgOrdersMgr) GetBatchFromIDLang(idLangs []uint32) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error
	
	return
}
 
// GetFromIDCustomer 通过id_customer获取内容  
func (obj *_EgOrdersMgr) GetFromIDCustomer(idCustomer uint32) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer = ?", idCustomer).Find(&results).Error
	
	return
}

// GetBatchFromIDCustomer 批量唯一主键查找 
func (obj *_EgOrdersMgr) GetBatchFromIDCustomer(idCustomers []uint32) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer IN (?)", idCustomers).Find(&results).Error
	
	return
}
 
// GetFromIDCart 通过id_cart获取内容  
func (obj *_EgOrdersMgr) GetFromIDCart(idCart uint32) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart = ?", idCart).Find(&results).Error
	
	return
}

// GetBatchFromIDCart 批量唯一主键查找 
func (obj *_EgOrdersMgr) GetBatchFromIDCart(idCarts []uint32) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart IN (?)", idCarts).Find(&results).Error
	
	return
}
 
// GetFromIDCurrency 通过id_currency获取内容  
func (obj *_EgOrdersMgr) GetFromIDCurrency(idCurrency uint32) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_currency = ?", idCurrency).Find(&results).Error
	
	return
}

// GetBatchFromIDCurrency 批量唯一主键查找 
func (obj *_EgOrdersMgr) GetBatchFromIDCurrency(idCurrencys []uint32) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_currency IN (?)", idCurrencys).Find(&results).Error
	
	return
}
 
// GetFromIDAddressDelivery 通过id_address_delivery获取内容  
func (obj *_EgOrdersMgr) GetFromIDAddressDelivery(idAddressDelivery uint32) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_address_delivery = ?", idAddressDelivery).Find(&results).Error
	
	return
}

// GetBatchFromIDAddressDelivery 批量唯一主键查找 
func (obj *_EgOrdersMgr) GetBatchFromIDAddressDelivery(idAddressDeliverys []uint32) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_address_delivery IN (?)", idAddressDeliverys).Find(&results).Error
	
	return
}
 
// GetFromIDAddressInvoice 通过id_address_invoice获取内容  
func (obj *_EgOrdersMgr) GetFromIDAddressInvoice(idAddressInvoice uint32) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_address_invoice = ?", idAddressInvoice).Find(&results).Error
	
	return
}

// GetBatchFromIDAddressInvoice 批量唯一主键查找 
func (obj *_EgOrdersMgr) GetBatchFromIDAddressInvoice(idAddressInvoices []uint32) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_address_invoice IN (?)", idAddressInvoices).Find(&results).Error
	
	return
}
 
// GetFromCurrentState 通过current_state获取内容  
func (obj *_EgOrdersMgr) GetFromCurrentState(currentState uint32) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("current_state = ?", currentState).Find(&results).Error
	
	return
}

// GetBatchFromCurrentState 批量唯一主键查找 
func (obj *_EgOrdersMgr) GetBatchFromCurrentState(currentStates []uint32) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("current_state IN (?)", currentStates).Find(&results).Error
	
	return
}
 
// GetFromSecureKey 通过secure_key获取内容  
func (obj *_EgOrdersMgr) GetFromSecureKey(secureKey string) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("secure_key = ?", secureKey).Find(&results).Error
	
	return
}

// GetBatchFromSecureKey 批量唯一主键查找 
func (obj *_EgOrdersMgr) GetBatchFromSecureKey(secureKeys []string) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("secure_key IN (?)", secureKeys).Find(&results).Error
	
	return
}
 
// GetFromPayment 通过payment获取内容  
func (obj *_EgOrdersMgr) GetFromPayment(payment string) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("payment = ?", payment).Find(&results).Error
	
	return
}

// GetBatchFromPayment 批量唯一主键查找 
func (obj *_EgOrdersMgr) GetBatchFromPayment(payments []string) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("payment IN (?)", payments).Find(&results).Error
	
	return
}
 
// GetFromConversionRate 通过conversion_rate获取内容  
func (obj *_EgOrdersMgr) GetFromConversionRate(conversionRate float64) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("conversion_rate = ?", conversionRate).Find(&results).Error
	
	return
}

// GetBatchFromConversionRate 批量唯一主键查找 
func (obj *_EgOrdersMgr) GetBatchFromConversionRate(conversionRates []float64) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("conversion_rate IN (?)", conversionRates).Find(&results).Error
	
	return
}
 
// GetFromModule 通过module获取内容  
func (obj *_EgOrdersMgr) GetFromModule(module string) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("module = ?", module).Find(&results).Error
	
	return
}

// GetBatchFromModule 批量唯一主键查找 
func (obj *_EgOrdersMgr) GetBatchFromModule(modules []string) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("module IN (?)", modules).Find(&results).Error
	
	return
}
 
// GetFromRecyclable 通过recyclable获取内容  
func (obj *_EgOrdersMgr) GetFromRecyclable(recyclable bool) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("recyclable = ?", recyclable).Find(&results).Error
	
	return
}

// GetBatchFromRecyclable 批量唯一主键查找 
func (obj *_EgOrdersMgr) GetBatchFromRecyclable(recyclables []bool) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("recyclable IN (?)", recyclables).Find(&results).Error
	
	return
}
 
// GetFromGift 通过gift获取内容  
func (obj *_EgOrdersMgr) GetFromGift(gift bool) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("gift = ?", gift).Find(&results).Error
	
	return
}

// GetBatchFromGift 批量唯一主键查找 
func (obj *_EgOrdersMgr) GetBatchFromGift(gifts []bool) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("gift IN (?)", gifts).Find(&results).Error
	
	return
}
 
// GetFromGiftMessage 通过gift_message获取内容  
func (obj *_EgOrdersMgr) GetFromGiftMessage(giftMessage string) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("gift_message = ?", giftMessage).Find(&results).Error
	
	return
}

// GetBatchFromGiftMessage 批量唯一主键查找 
func (obj *_EgOrdersMgr) GetBatchFromGiftMessage(giftMessages []string) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("gift_message IN (?)", giftMessages).Find(&results).Error
	
	return
}
 
// GetFromMobileTheme 通过mobile_theme获取内容  
func (obj *_EgOrdersMgr) GetFromMobileTheme(mobileTheme bool) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("mobile_theme = ?", mobileTheme).Find(&results).Error
	
	return
}

// GetBatchFromMobileTheme 批量唯一主键查找 
func (obj *_EgOrdersMgr) GetBatchFromMobileTheme(mobileThemes []bool) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("mobile_theme IN (?)", mobileThemes).Find(&results).Error
	
	return
}
 
// GetFromShippingNumber 通过shipping_number获取内容  
func (obj *_EgOrdersMgr) GetFromShippingNumber(shippingNumber string) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("shipping_number = ?", shippingNumber).Find(&results).Error
	
	return
}

// GetBatchFromShippingNumber 批量唯一主键查找 
func (obj *_EgOrdersMgr) GetBatchFromShippingNumber(shippingNumbers []string) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("shipping_number IN (?)", shippingNumbers).Find(&results).Error
	
	return
}
 
// GetFromTotalDiscounts 通过total_discounts获取内容  
func (obj *_EgOrdersMgr) GetFromTotalDiscounts(totalDiscounts float64) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_discounts = ?", totalDiscounts).Find(&results).Error
	
	return
}

// GetBatchFromTotalDiscounts 批量唯一主键查找 
func (obj *_EgOrdersMgr) GetBatchFromTotalDiscounts(totalDiscountss []float64) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_discounts IN (?)", totalDiscountss).Find(&results).Error
	
	return
}
 
// GetFromTotalDiscountsTaxIncl 通过total_discounts_tax_incl获取内容  
func (obj *_EgOrdersMgr) GetFromTotalDiscountsTaxIncl(totalDiscountsTaxIncl float64) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_discounts_tax_incl = ?", totalDiscountsTaxIncl).Find(&results).Error
	
	return
}

// GetBatchFromTotalDiscountsTaxIncl 批量唯一主键查找 
func (obj *_EgOrdersMgr) GetBatchFromTotalDiscountsTaxIncl(totalDiscountsTaxIncls []float64) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_discounts_tax_incl IN (?)", totalDiscountsTaxIncls).Find(&results).Error
	
	return
}
 
// GetFromTotalDiscountsTaxExcl 通过total_discounts_tax_excl获取内容  
func (obj *_EgOrdersMgr) GetFromTotalDiscountsTaxExcl(totalDiscountsTaxExcl float64) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_discounts_tax_excl = ?", totalDiscountsTaxExcl).Find(&results).Error
	
	return
}

// GetBatchFromTotalDiscountsTaxExcl 批量唯一主键查找 
func (obj *_EgOrdersMgr) GetBatchFromTotalDiscountsTaxExcl(totalDiscountsTaxExcls []float64) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_discounts_tax_excl IN (?)", totalDiscountsTaxExcls).Find(&results).Error
	
	return
}
 
// GetFromTotalPaid 通过total_paid获取内容  
func (obj *_EgOrdersMgr) GetFromTotalPaid(totalPaid float64) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_paid = ?", totalPaid).Find(&results).Error
	
	return
}

// GetBatchFromTotalPaid 批量唯一主键查找 
func (obj *_EgOrdersMgr) GetBatchFromTotalPaid(totalPaids []float64) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_paid IN (?)", totalPaids).Find(&results).Error
	
	return
}
 
// GetFromTotalPaidTaxIncl 通过total_paid_tax_incl获取内容  
func (obj *_EgOrdersMgr) GetFromTotalPaidTaxIncl(totalPaidTaxIncl float64) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_paid_tax_incl = ?", totalPaidTaxIncl).Find(&results).Error
	
	return
}

// GetBatchFromTotalPaidTaxIncl 批量唯一主键查找 
func (obj *_EgOrdersMgr) GetBatchFromTotalPaidTaxIncl(totalPaidTaxIncls []float64) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_paid_tax_incl IN (?)", totalPaidTaxIncls).Find(&results).Error
	
	return
}
 
// GetFromTotalPaidTaxExcl 通过total_paid_tax_excl获取内容  
func (obj *_EgOrdersMgr) GetFromTotalPaidTaxExcl(totalPaidTaxExcl float64) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_paid_tax_excl = ?", totalPaidTaxExcl).Find(&results).Error
	
	return
}

// GetBatchFromTotalPaidTaxExcl 批量唯一主键查找 
func (obj *_EgOrdersMgr) GetBatchFromTotalPaidTaxExcl(totalPaidTaxExcls []float64) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_paid_tax_excl IN (?)", totalPaidTaxExcls).Find(&results).Error
	
	return
}
 
// GetFromTotalPaidReal 通过total_paid_real获取内容  
func (obj *_EgOrdersMgr) GetFromTotalPaidReal(totalPaidReal float64) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_paid_real = ?", totalPaidReal).Find(&results).Error
	
	return
}

// GetBatchFromTotalPaidReal 批量唯一主键查找 
func (obj *_EgOrdersMgr) GetBatchFromTotalPaidReal(totalPaidReals []float64) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_paid_real IN (?)", totalPaidReals).Find(&results).Error
	
	return
}
 
// GetFromTotalProducts 通过total_products获取内容  
func (obj *_EgOrdersMgr) GetFromTotalProducts(totalProducts float64) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_products = ?", totalProducts).Find(&results).Error
	
	return
}

// GetBatchFromTotalProducts 批量唯一主键查找 
func (obj *_EgOrdersMgr) GetBatchFromTotalProducts(totalProductss []float64) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_products IN (?)", totalProductss).Find(&results).Error
	
	return
}
 
// GetFromTotalProductsWt 通过total_products_wt获取内容  
func (obj *_EgOrdersMgr) GetFromTotalProductsWt(totalProductsWt float64) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_products_wt = ?", totalProductsWt).Find(&results).Error
	
	return
}

// GetBatchFromTotalProductsWt 批量唯一主键查找 
func (obj *_EgOrdersMgr) GetBatchFromTotalProductsWt(totalProductsWts []float64) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_products_wt IN (?)", totalProductsWts).Find(&results).Error
	
	return
}
 
// GetFromTotalShipping 通过total_shipping获取内容  
func (obj *_EgOrdersMgr) GetFromTotalShipping(totalShipping float64) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_shipping = ?", totalShipping).Find(&results).Error
	
	return
}

// GetBatchFromTotalShipping 批量唯一主键查找 
func (obj *_EgOrdersMgr) GetBatchFromTotalShipping(totalShippings []float64) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_shipping IN (?)", totalShippings).Find(&results).Error
	
	return
}
 
// GetFromTotalShippingTaxIncl 通过total_shipping_tax_incl获取内容  
func (obj *_EgOrdersMgr) GetFromTotalShippingTaxIncl(totalShippingTaxIncl float64) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_shipping_tax_incl = ?", totalShippingTaxIncl).Find(&results).Error
	
	return
}

// GetBatchFromTotalShippingTaxIncl 批量唯一主键查找 
func (obj *_EgOrdersMgr) GetBatchFromTotalShippingTaxIncl(totalShippingTaxIncls []float64) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_shipping_tax_incl IN (?)", totalShippingTaxIncls).Find(&results).Error
	
	return
}
 
// GetFromTotalShippingTaxExcl 通过total_shipping_tax_excl获取内容  
func (obj *_EgOrdersMgr) GetFromTotalShippingTaxExcl(totalShippingTaxExcl float64) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_shipping_tax_excl = ?", totalShippingTaxExcl).Find(&results).Error
	
	return
}

// GetBatchFromTotalShippingTaxExcl 批量唯一主键查找 
func (obj *_EgOrdersMgr) GetBatchFromTotalShippingTaxExcl(totalShippingTaxExcls []float64) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_shipping_tax_excl IN (?)", totalShippingTaxExcls).Find(&results).Error
	
	return
}
 
// GetFromCarrierTaxRate 通过carrier_tax_rate获取内容  
func (obj *_EgOrdersMgr) GetFromCarrierTaxRate(carrierTaxRate float64) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("carrier_tax_rate = ?", carrierTaxRate).Find(&results).Error
	
	return
}

// GetBatchFromCarrierTaxRate 批量唯一主键查找 
func (obj *_EgOrdersMgr) GetBatchFromCarrierTaxRate(carrierTaxRates []float64) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("carrier_tax_rate IN (?)", carrierTaxRates).Find(&results).Error
	
	return
}
 
// GetFromTotalWrapping 通过total_wrapping获取内容  
func (obj *_EgOrdersMgr) GetFromTotalWrapping(totalWrapping float64) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_wrapping = ?", totalWrapping).Find(&results).Error
	
	return
}

// GetBatchFromTotalWrapping 批量唯一主键查找 
func (obj *_EgOrdersMgr) GetBatchFromTotalWrapping(totalWrappings []float64) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_wrapping IN (?)", totalWrappings).Find(&results).Error
	
	return
}
 
// GetFromTotalWrappingTaxIncl 通过total_wrapping_tax_incl获取内容  
func (obj *_EgOrdersMgr) GetFromTotalWrappingTaxIncl(totalWrappingTaxIncl float64) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_wrapping_tax_incl = ?", totalWrappingTaxIncl).Find(&results).Error
	
	return
}

// GetBatchFromTotalWrappingTaxIncl 批量唯一主键查找 
func (obj *_EgOrdersMgr) GetBatchFromTotalWrappingTaxIncl(totalWrappingTaxIncls []float64) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_wrapping_tax_incl IN (?)", totalWrappingTaxIncls).Find(&results).Error
	
	return
}
 
// GetFromTotalWrappingTaxExcl 通过total_wrapping_tax_excl获取内容  
func (obj *_EgOrdersMgr) GetFromTotalWrappingTaxExcl(totalWrappingTaxExcl float64) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_wrapping_tax_excl = ?", totalWrappingTaxExcl).Find(&results).Error
	
	return
}

// GetBatchFromTotalWrappingTaxExcl 批量唯一主键查找 
func (obj *_EgOrdersMgr) GetBatchFromTotalWrappingTaxExcl(totalWrappingTaxExcls []float64) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_wrapping_tax_excl IN (?)", totalWrappingTaxExcls).Find(&results).Error
	
	return
}
 
// GetFromRoundMode 通过round_mode获取内容  
func (obj *_EgOrdersMgr) GetFromRoundMode(roundMode bool) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("round_mode = ?", roundMode).Find(&results).Error
	
	return
}

// GetBatchFromRoundMode 批量唯一主键查找 
func (obj *_EgOrdersMgr) GetBatchFromRoundMode(roundModes []bool) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("round_mode IN (?)", roundModes).Find(&results).Error
	
	return
}
 
// GetFromRoundType 通过round_type获取内容  
func (obj *_EgOrdersMgr) GetFromRoundType(roundType bool) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("round_type = ?", roundType).Find(&results).Error
	
	return
}

// GetBatchFromRoundType 批量唯一主键查找 
func (obj *_EgOrdersMgr) GetBatchFromRoundType(roundTypes []bool) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("round_type IN (?)", roundTypes).Find(&results).Error
	
	return
}
 
// GetFromInvoiceNumber 通过invoice_number获取内容  
func (obj *_EgOrdersMgr) GetFromInvoiceNumber(invoiceNumber uint32) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("invoice_number = ?", invoiceNumber).Find(&results).Error
	
	return
}

// GetBatchFromInvoiceNumber 批量唯一主键查找 
func (obj *_EgOrdersMgr) GetBatchFromInvoiceNumber(invoiceNumbers []uint32) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("invoice_number IN (?)", invoiceNumbers).Find(&results).Error
	
	return
}
 
// GetFromDeliveryNumber 通过delivery_number获取内容  
func (obj *_EgOrdersMgr) GetFromDeliveryNumber(deliveryNumber uint32) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("delivery_number = ?", deliveryNumber).Find(&results).Error
	
	return
}

// GetBatchFromDeliveryNumber 批量唯一主键查找 
func (obj *_EgOrdersMgr) GetBatchFromDeliveryNumber(deliveryNumbers []uint32) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("delivery_number IN (?)", deliveryNumbers).Find(&results).Error
	
	return
}
 
// GetFromInvoiceDate 通过invoice_date获取内容  
func (obj *_EgOrdersMgr) GetFromInvoiceDate(invoiceDate time.Time) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("invoice_date = ?", invoiceDate).Find(&results).Error
	
	return
}

// GetBatchFromInvoiceDate 批量唯一主键查找 
func (obj *_EgOrdersMgr) GetBatchFromInvoiceDate(invoiceDates []time.Time) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("invoice_date IN (?)", invoiceDates).Find(&results).Error
	
	return
}
 
// GetFromDeliveryDate 通过delivery_date获取内容  
func (obj *_EgOrdersMgr) GetFromDeliveryDate(deliveryDate time.Time) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("delivery_date = ?", deliveryDate).Find(&results).Error
	
	return
}

// GetBatchFromDeliveryDate 批量唯一主键查找 
func (obj *_EgOrdersMgr) GetBatchFromDeliveryDate(deliveryDates []time.Time) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("delivery_date IN (?)", deliveryDates).Find(&results).Error
	
	return
}
 
// GetFromValid 通过valid获取内容  
func (obj *_EgOrdersMgr) GetFromValid(valid uint32) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("valid = ?", valid).Find(&results).Error
	
	return
}

// GetBatchFromValid 批量唯一主键查找 
func (obj *_EgOrdersMgr) GetBatchFromValid(valids []uint32) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("valid IN (?)", valids).Find(&results).Error
	
	return
}
 
// GetFromDateAdd 通过date_add获取内容  
func (obj *_EgOrdersMgr) GetFromDateAdd(dateAdd time.Time) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error
	
	return
}

// GetBatchFromDateAdd 批量唯一主键查找 
func (obj *_EgOrdersMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error
	
	return
}
 
// GetFromDateUpd 通过date_upd获取内容  
func (obj *_EgOrdersMgr) GetFromDateUpd(dateUpd time.Time) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd = ?", dateUpd).Find(&results).Error
	
	return
}

// GetBatchFromDateUpd 批量唯一主键查找 
func (obj *_EgOrdersMgr) GetBatchFromDateUpd(dateUpds []time.Time) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd IN (?)", dateUpds).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgOrdersMgr) FetchByPrimaryKey(idOrder uint32 ) (result EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order = ?", idOrder).Find(&result).Error
	
	return
}
 

 
 // FetchIndexByReference  获取多个内容
 func (obj *_EgOrdersMgr) FetchIndexByReference(reference string ) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reference = ?", reference).Find(&results).Error
	
	return
}
 
 // FetchIndexByIDShopGroup  获取多个内容
 func (obj *_EgOrdersMgr) FetchIndexByIDShopGroup(idShopGroup uint32 ) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_group = ?", idShopGroup).Find(&results).Error
	
	return
}
 
 // FetchIndexByIDShop  获取多个内容
 func (obj *_EgOrdersMgr) FetchIndexByIDShop(idShop uint32 ) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error
	
	return
}
 
 // FetchIndexByIDCarrier  获取多个内容
 func (obj *_EgOrdersMgr) FetchIndexByIDCarrier(idCarrier uint32 ) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_carrier = ?", idCarrier).Find(&results).Error
	
	return
}
 
 // FetchIndexByIDLang  获取多个内容
 func (obj *_EgOrdersMgr) FetchIndexByIDLang(idLang uint32 ) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error
	
	return
}
 
 // FetchIndexByIDCustomer  获取多个内容
 func (obj *_EgOrdersMgr) FetchIndexByIDCustomer(idCustomer uint32 ) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer = ?", idCustomer).Find(&results).Error
	
	return
}
 
 // FetchIndexByIDCart  获取多个内容
 func (obj *_EgOrdersMgr) FetchIndexByIDCart(idCart uint32 ) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart = ?", idCart).Find(&results).Error
	
	return
}
 
 // FetchIndexByIDCurrency  获取多个内容
 func (obj *_EgOrdersMgr) FetchIndexByIDCurrency(idCurrency uint32 ) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_currency = ?", idCurrency).Find(&results).Error
	
	return
}
 
 // FetchIndexByIDAddressDelivery  获取多个内容
 func (obj *_EgOrdersMgr) FetchIndexByIDAddressDelivery(idAddressDelivery uint32 ) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_address_delivery = ?", idAddressDelivery).Find(&results).Error
	
	return
}
 
 // FetchIndexByIDAddressInvoice  获取多个内容
 func (obj *_EgOrdersMgr) FetchIndexByIDAddressInvoice(idAddressInvoice uint32 ) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_address_invoice = ?", idAddressInvoice).Find(&results).Error
	
	return
}
 
 // FetchIndexByCurrentState  获取多个内容
 func (obj *_EgOrdersMgr) FetchIndexByCurrentState(currentState uint32 ) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("current_state = ?", currentState).Find(&results).Error
	
	return
}
 
 // FetchIndexByInvoiceNumber  获取多个内容
 func (obj *_EgOrdersMgr) FetchIndexByInvoiceNumber(invoiceNumber uint32 ) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("invoice_number = ?", invoiceNumber).Find(&results).Error
	
	return
}
 
 // FetchIndexByDateAdd  获取多个内容
 func (obj *_EgOrdersMgr) FetchIndexByDateAdd(dateAdd time.Time ) (results []*EgOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error
	
	return
}
 

	

