package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _SupplyOrderDetailMgr struct {
	*_BaseMgr
}

func SupplyOrderDetailMgr(db *gorm.DB) *_SupplyOrderDetailMgr {
	if db == nil {
		panic(fmt.Errorf("SupplyOrderDetailMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SupplyOrderDetailMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_supply_order_detail"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_SupplyOrderDetailMgr) GetTableName() string {
	return "ps_supply_order_detail"
}

func (obj *_SupplyOrderDetailMgr) Get() (result SupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_SupplyOrderDetailMgr) Gets() (results []*SupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

func (obj *_SupplyOrderDetailMgr) WithIDSupplyOrderDetail(idSupplyOrderDetail uint32) Option {
	return optionFunc(func(o *options) { o.query["id_supply_order_detail"] = idSupplyOrderDetail })
}

func (obj *_SupplyOrderDetailMgr) WithIDSupplyOrder(idSupplyOrder uint32) Option {
	return optionFunc(func(o *options) { o.query["id_supply_order"] = idSupplyOrder })
}

func (obj *_SupplyOrderDetailMgr) WithIDCurrency(idCurrency uint32) Option {
	return optionFunc(func(o *options) { o.query["id_currency"] = idCurrency })
}

func (obj *_SupplyOrderDetailMgr) WithIDProduct(idProduct uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product"] = idProduct })
}

func (obj *_SupplyOrderDetailMgr) WithIDProductAttribute(idProductAttribute uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product_attribute"] = idProductAttribute })
}

func (obj *_SupplyOrderDetailMgr) WithReference(reference string) Option {
	return optionFunc(func(o *options) { o.query["reference"] = reference })
}

func (obj *_SupplyOrderDetailMgr) WithSupplierReference(supplierReference string) Option {
	return optionFunc(func(o *options) { o.query["supplier_reference"] = supplierReference })
}

func (obj *_SupplyOrderDetailMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

func (obj *_SupplyOrderDetailMgr) WithEan13(ean13 string) Option {
	return optionFunc(func(o *options) { o.query["ean13"] = ean13 })
}

func (obj *_SupplyOrderDetailMgr) WithIsbn(isbn string) Option {
	return optionFunc(func(o *options) { o.query["isbn"] = isbn })
}

func (obj *_SupplyOrderDetailMgr) WithUpc(upc string) Option {
	return optionFunc(func(o *options) { o.query["upc"] = upc })
}

func (obj *_SupplyOrderDetailMgr) WithMpn(mpn string) Option {
	return optionFunc(func(o *options) { o.query["mpn"] = mpn })
}

func (obj *_SupplyOrderDetailMgr) WithExchangeRate(exchangeRate float64) Option {
	return optionFunc(func(o *options) { o.query["exchange_rate"] = exchangeRate })
}

func (obj *_SupplyOrderDetailMgr) WithUnitPriceTe(unitPriceTe float64) Option {
	return optionFunc(func(o *options) { o.query["unit_price_te"] = unitPriceTe })
}

func (obj *_SupplyOrderDetailMgr) WithQuantityExpected(quantityExpected uint32) Option {
	return optionFunc(func(o *options) { o.query["quantity_expected"] = quantityExpected })
}

func (obj *_SupplyOrderDetailMgr) WithQuantityReceived(quantityReceived uint32) Option {
	return optionFunc(func(o *options) { o.query["quantity_received"] = quantityReceived })
}

func (obj *_SupplyOrderDetailMgr) WithPriceTe(priceTe float64) Option {
	return optionFunc(func(o *options) { o.query["price_te"] = priceTe })
}

func (obj *_SupplyOrderDetailMgr) WithDiscountRate(discountRate float64) Option {
	return optionFunc(func(o *options) { o.query["discount_rate"] = discountRate })
}

func (obj *_SupplyOrderDetailMgr) WithDiscountValueTe(discountValueTe float64) Option {
	return optionFunc(func(o *options) { o.query["discount_value_te"] = discountValueTe })
}

func (obj *_SupplyOrderDetailMgr) WithPriceWithDiscountTe(priceWithDiscountTe float64) Option {
	return optionFunc(func(o *options) { o.query["price_with_discount_te"] = priceWithDiscountTe })
}

func (obj *_SupplyOrderDetailMgr) WithTaxRate(taxRate float64) Option {
	return optionFunc(func(o *options) { o.query["tax_rate"] = taxRate })
}

func (obj *_SupplyOrderDetailMgr) WithTaxValue(taxValue float64) Option {
	return optionFunc(func(o *options) { o.query["tax_value"] = taxValue })
}

func (obj *_SupplyOrderDetailMgr) WithPriceTi(priceTi float64) Option {
	return optionFunc(func(o *options) { o.query["price_ti"] = priceTi })
}

func (obj *_SupplyOrderDetailMgr) WithTaxValueWithOrderDiscount(taxValueWithOrderDiscount float64) Option {
	return optionFunc(func(o *options) { o.query["tax_value_with_order_discount"] = taxValueWithOrderDiscount })
}

func (obj *_SupplyOrderDetailMgr) WithPriceWithOrderDiscountTe(priceWithOrderDiscountTe float64) Option {
	return optionFunc(func(o *options) { o.query["price_with_order_discount_te"] = priceWithOrderDiscountTe })
}

func (obj *_SupplyOrderDetailMgr) GetByOption(opts ...Option) (result SupplyOrderDetail, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_SupplyOrderDetailMgr) GetByOptions(opts ...Option) (results []*SupplyOrderDetail, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

func (obj *_SupplyOrderDetailMgr) GetFromIDSupplyOrderDetail(idSupplyOrderDetail uint32) (result SupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supply_order_detail = ?", idSupplyOrderDetail).Find(&result).Error

	return
}

func (obj *_SupplyOrderDetailMgr) GetBatchFromIDSupplyOrderDetail(idSupplyOrderDetails []uint32) (results []*SupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supply_order_detail IN (?)", idSupplyOrderDetails).Find(&results).Error

	return
}

func (obj *_SupplyOrderDetailMgr) GetFromIDSupplyOrder(idSupplyOrder uint32) (results []*SupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supply_order = ?", idSupplyOrder).Find(&results).Error

	return
}

func (obj *_SupplyOrderDetailMgr) GetBatchFromIDSupplyOrder(idSupplyOrders []uint32) (results []*SupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supply_order IN (?)", idSupplyOrders).Find(&results).Error

	return
}

func (obj *_SupplyOrderDetailMgr) GetFromIDCurrency(idCurrency uint32) (results []*SupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_currency = ?", idCurrency).Find(&results).Error

	return
}

func (obj *_SupplyOrderDetailMgr) GetBatchFromIDCurrency(idCurrencys []uint32) (results []*SupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_currency IN (?)", idCurrencys).Find(&results).Error

	return
}

func (obj *_SupplyOrderDetailMgr) GetFromIDProduct(idProduct uint32) (results []*SupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ?", idProduct).Find(&results).Error

	return
}

func (obj *_SupplyOrderDetailMgr) GetBatchFromIDProduct(idProducts []uint32) (results []*SupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product IN (?)", idProducts).Find(&results).Error

	return
}

func (obj *_SupplyOrderDetailMgr) GetFromIDProductAttribute(idProductAttribute uint32) (results []*SupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_attribute = ?", idProductAttribute).Find(&results).Error

	return
}

func (obj *_SupplyOrderDetailMgr) GetBatchFromIDProductAttribute(idProductAttributes []uint32) (results []*SupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_attribute IN (?)", idProductAttributes).Find(&results).Error

	return
}

func (obj *_SupplyOrderDetailMgr) GetFromReference(reference string) (results []*SupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reference = ?", reference).Find(&results).Error

	return
}

func (obj *_SupplyOrderDetailMgr) GetBatchFromReference(references []string) (results []*SupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reference IN (?)", references).Find(&results).Error

	return
}

func (obj *_SupplyOrderDetailMgr) GetFromSupplierReference(supplierReference string) (results []*SupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("supplier_reference = ?", supplierReference).Find(&results).Error

	return
}

func (obj *_SupplyOrderDetailMgr) GetBatchFromSupplierReference(supplierReferences []string) (results []*SupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("supplier_reference IN (?)", supplierReferences).Find(&results).Error

	return
}

func (obj *_SupplyOrderDetailMgr) GetFromName(name string) (results []*SupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

func (obj *_SupplyOrderDetailMgr) GetBatchFromName(names []string) (results []*SupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

func (obj *_SupplyOrderDetailMgr) GetFromEan13(ean13 string) (results []*SupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ean13 = ?", ean13).Find(&results).Error

	return
}

func (obj *_SupplyOrderDetailMgr) GetBatchFromEan13(ean13s []string) (results []*SupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ean13 IN (?)", ean13s).Find(&results).Error

	return
}

func (obj *_SupplyOrderDetailMgr) GetFromIsbn(isbn string) (results []*SupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("isbn = ?", isbn).Find(&results).Error

	return
}

func (obj *_SupplyOrderDetailMgr) GetBatchFromIsbn(isbns []string) (results []*SupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("isbn IN (?)", isbns).Find(&results).Error

	return
}

func (obj *_SupplyOrderDetailMgr) GetFromUpc(upc string) (results []*SupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("upc = ?", upc).Find(&results).Error

	return
}

func (obj *_SupplyOrderDetailMgr) GetBatchFromUpc(upcs []string) (results []*SupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("upc IN (?)", upcs).Find(&results).Error

	return
}

func (obj *_SupplyOrderDetailMgr) GetFromMpn(mpn string) (results []*SupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("mpn = ?", mpn).Find(&results).Error

	return
}

func (obj *_SupplyOrderDetailMgr) GetBatchFromMpn(mpns []string) (results []*SupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("mpn IN (?)", mpns).Find(&results).Error

	return
}

func (obj *_SupplyOrderDetailMgr) GetFromExchangeRate(exchangeRate float64) (results []*SupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("exchange_rate = ?", exchangeRate).Find(&results).Error

	return
}

func (obj *_SupplyOrderDetailMgr) GetBatchFromExchangeRate(exchangeRates []float64) (results []*SupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("exchange_rate IN (?)", exchangeRates).Find(&results).Error

	return
}

func (obj *_SupplyOrderDetailMgr) GetFromUnitPriceTe(unitPriceTe float64) (results []*SupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("unit_price_te = ?", unitPriceTe).Find(&results).Error

	return
}

func (obj *_SupplyOrderDetailMgr) GetBatchFromUnitPriceTe(unitPriceTes []float64) (results []*SupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("unit_price_te IN (?)", unitPriceTes).Find(&results).Error

	return
}

func (obj *_SupplyOrderDetailMgr) GetFromQuantityExpected(quantityExpected uint32) (results []*SupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("quantity_expected = ?", quantityExpected).Find(&results).Error

	return
}

func (obj *_SupplyOrderDetailMgr) GetBatchFromQuantityExpected(quantityExpecteds []uint32) (results []*SupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("quantity_expected IN (?)", quantityExpecteds).Find(&results).Error

	return
}

func (obj *_SupplyOrderDetailMgr) GetFromQuantityReceived(quantityReceived uint32) (results []*SupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("quantity_received = ?", quantityReceived).Find(&results).Error

	return
}

func (obj *_SupplyOrderDetailMgr) GetBatchFromQuantityReceived(quantityReceiveds []uint32) (results []*SupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("quantity_received IN (?)", quantityReceiveds).Find(&results).Error

	return
}

func (obj *_SupplyOrderDetailMgr) GetFromPriceTe(priceTe float64) (results []*SupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("price_te = ?", priceTe).Find(&results).Error

	return
}

func (obj *_SupplyOrderDetailMgr) GetBatchFromPriceTe(priceTes []float64) (results []*SupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("price_te IN (?)", priceTes).Find(&results).Error

	return
}

func (obj *_SupplyOrderDetailMgr) GetFromDiscountRate(discountRate float64) (results []*SupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("discount_rate = ?", discountRate).Find(&results).Error

	return
}

func (obj *_SupplyOrderDetailMgr) GetBatchFromDiscountRate(discountRates []float64) (results []*SupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("discount_rate IN (?)", discountRates).Find(&results).Error

	return
}

func (obj *_SupplyOrderDetailMgr) GetFromDiscountValueTe(discountValueTe float64) (results []*SupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("discount_value_te = ?", discountValueTe).Find(&results).Error

	return
}

func (obj *_SupplyOrderDetailMgr) GetBatchFromDiscountValueTe(discountValueTes []float64) (results []*SupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("discount_value_te IN (?)", discountValueTes).Find(&results).Error

	return
}

func (obj *_SupplyOrderDetailMgr) GetFromPriceWithDiscountTe(priceWithDiscountTe float64) (results []*SupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("price_with_discount_te = ?", priceWithDiscountTe).Find(&results).Error

	return
}

func (obj *_SupplyOrderDetailMgr) GetBatchFromPriceWithDiscountTe(priceWithDiscountTes []float64) (results []*SupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("price_with_discount_te IN (?)", priceWithDiscountTes).Find(&results).Error

	return
}

func (obj *_SupplyOrderDetailMgr) GetFromTaxRate(taxRate float64) (results []*SupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("tax_rate = ?", taxRate).Find(&results).Error

	return
}

func (obj *_SupplyOrderDetailMgr) GetBatchFromTaxRate(taxRates []float64) (results []*SupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("tax_rate IN (?)", taxRates).Find(&results).Error

	return
}

func (obj *_SupplyOrderDetailMgr) GetFromTaxValue(taxValue float64) (results []*SupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("tax_value = ?", taxValue).Find(&results).Error

	return
}

func (obj *_SupplyOrderDetailMgr) GetBatchFromTaxValue(taxValues []float64) (results []*SupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("tax_value IN (?)", taxValues).Find(&results).Error

	return
}

func (obj *_SupplyOrderDetailMgr) GetFromPriceTi(priceTi float64) (results []*SupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("price_ti = ?", priceTi).Find(&results).Error

	return
}

func (obj *_SupplyOrderDetailMgr) GetBatchFromPriceTi(priceTis []float64) (results []*SupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("price_ti IN (?)", priceTis).Find(&results).Error

	return
}

func (obj *_SupplyOrderDetailMgr) GetFromTaxValueWithOrderDiscount(taxValueWithOrderDiscount float64) (results []*SupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("tax_value_with_order_discount = ?", taxValueWithOrderDiscount).Find(&results).Error

	return
}

func (obj *_SupplyOrderDetailMgr) GetBatchFromTaxValueWithOrderDiscount(taxValueWithOrderDiscounts []float64) (results []*SupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("tax_value_with_order_discount IN (?)", taxValueWithOrderDiscounts).Find(&results).Error

	return
}

func (obj *_SupplyOrderDetailMgr) GetFromPriceWithOrderDiscountTe(priceWithOrderDiscountTe float64) (results []*SupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("price_with_order_discount_te = ?", priceWithOrderDiscountTe).Find(&results).Error

	return
}

func (obj *_SupplyOrderDetailMgr) GetBatchFromPriceWithOrderDiscountTe(priceWithOrderDiscountTes []float64) (results []*SupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("price_with_order_discount_te IN (?)", priceWithOrderDiscountTes).Find(&results).Error

	return
}

func (obj *_SupplyOrderDetailMgr) FetchByPrimaryKey(idSupplyOrderDetail uint32) (result SupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supply_order_detail = ?", idSupplyOrderDetail).Find(&result).Error

	return
}

func (obj *_SupplyOrderDetailMgr) FetchIndexByIDSupplyOrder(idSupplyOrder uint32, idProduct uint32) (results []*SupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supply_order = ? AND id_product = ?", idSupplyOrder, idProduct).Find(&results).Error

	return
}

func (obj *_SupplyOrderDetailMgr) FetchIndexByIDProductProductAttribute(idProduct uint32, idProductAttribute uint32) (results []*SupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ? AND id_product_attribute = ?", idProduct, idProductAttribute).Find(&results).Error

	return
}

func (obj *_SupplyOrderDetailMgr) FetchIndexByIDProductAttribute(idProductAttribute uint32) (results []*SupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_attribute = ?", idProductAttribute).Find(&results).Error

	return
}
