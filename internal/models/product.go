package models

import (
	"context"
	"fmt"
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type _ProductMgr struct {
	*_BaseMgr
}

func ProductMgr(db *gorm.DB) *_ProductMgr {
	if db == nil {
		panic(fmt.Errorf("ProductMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ProductMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_product"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_ProductMgr) GetTableName() string {
	return "ps_product"
}

func (obj *_ProductMgr) Get() (result Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_ProductMgr) Gets() (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

func (obj *_ProductMgr) WithIDProduct(idProduct uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product"] = idProduct })
}

func (obj *_ProductMgr) WithIDSupplier(idSupplier uint32) Option {
	return optionFunc(func(o *options) { o.query["id_supplier"] = idSupplier })
}

func (obj *_ProductMgr) WithIDManufacturer(idManufacturer uint32) Option {
	return optionFunc(func(o *options) { o.query["id_manufacturer"] = idManufacturer })
}

func (obj *_ProductMgr) WithIDCategoryDefault(idCategoryDefault uint32) Option {
	return optionFunc(func(o *options) { o.query["id_category_default"] = idCategoryDefault })
}

func (obj *_ProductMgr) WithIDShopDefault(idShopDefault uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop_default"] = idShopDefault })
}

func (obj *_ProductMgr) WithIDTaxRulesGroup(idTaxRulesGroup uint32) Option {
	return optionFunc(func(o *options) { o.query["id_tax_rules_group"] = idTaxRulesGroup })
}

func (obj *_ProductMgr) WithOnSale(onSale bool) Option {
	return optionFunc(func(o *options) { o.query["on_sale"] = onSale })
}

func (obj *_ProductMgr) WithOnlineOnly(onlineOnly bool) Option {
	return optionFunc(func(o *options) { o.query["online_only"] = onlineOnly })
}

func (obj *_ProductMgr) WithEan13(ean13 string) Option {
	return optionFunc(func(o *options) { o.query["ean13"] = ean13 })
}

func (obj *_ProductMgr) WithIsbn(isbn string) Option {
	return optionFunc(func(o *options) { o.query["isbn"] = isbn })
}

func (obj *_ProductMgr) WithUpc(upc string) Option {
	return optionFunc(func(o *options) { o.query["upc"] = upc })
}

func (obj *_ProductMgr) WithMpn(mpn string) Option {
	return optionFunc(func(o *options) { o.query["mpn"] = mpn })
}

func (obj *_ProductMgr) WithEcotax(ecotax float64) Option {
	return optionFunc(func(o *options) { o.query["ecotax"] = ecotax })
}

func (obj *_ProductMgr) WithQuantity(quantity int) Option {
	return optionFunc(func(o *options) { o.query["quantity"] = quantity })
}

func (obj *_ProductMgr) WithMinimalQuantity(minimalQuantity uint32) Option {
	return optionFunc(func(o *options) { o.query["minimal_quantity"] = minimalQuantity })
}

func (obj *_ProductMgr) WithLowStockThreshold(lowStockThreshold int) Option {
	return optionFunc(func(o *options) { o.query["low_stock_threshold"] = lowStockThreshold })
}

func (obj *_ProductMgr) WithLowStockAlert(lowStockAlert bool) Option {
	return optionFunc(func(o *options) { o.query["low_stock_alert"] = lowStockAlert })
}

func (obj *_ProductMgr) WithPrice(price float64) Option {
	return optionFunc(func(o *options) { o.query["price"] = price })
}

func (obj *_ProductMgr) WithWholesalePrice(wholesalePrice float64) Option {
	return optionFunc(func(o *options) { o.query["wholesale_price"] = wholesalePrice })
}

func (obj *_ProductMgr) WithUnity(unity string) Option {
	return optionFunc(func(o *options) { o.query["unity"] = unity })
}

func (obj *_ProductMgr) WithUnitPriceRatio(unitPriceRatio float64) Option {
	return optionFunc(func(o *options) { o.query["unit_price_ratio"] = unitPriceRatio })
}

func (obj *_ProductMgr) WithAdditionalShippingCost(additionalShippingCost float64) Option {
	return optionFunc(func(o *options) { o.query["additional_shipping_cost"] = additionalShippingCost })
}

func (obj *_ProductMgr) WithReference(reference string) Option {
	return optionFunc(func(o *options) { o.query["reference"] = reference })
}

func (obj *_ProductMgr) WithSupplierReference(supplierReference string) Option {
	return optionFunc(func(o *options) { o.query["supplier_reference"] = supplierReference })
}

func (obj *_ProductMgr) WithLocation(location string) Option {
	return optionFunc(func(o *options) { o.query["location"] = location })
}

func (obj *_ProductMgr) WithWidth(width float64) Option {
	return optionFunc(func(o *options) { o.query["width"] = width })
}

func (obj *_ProductMgr) WithHeight(height float64) Option {
	return optionFunc(func(o *options) { o.query["height"] = height })
}

func (obj *_ProductMgr) WithDepth(depth float64) Option {
	return optionFunc(func(o *options) { o.query["depth"] = depth })
}

func (obj *_ProductMgr) WithWeight(weight float64) Option {
	return optionFunc(func(o *options) { o.query["weight"] = weight })
}

func (obj *_ProductMgr) WithOutOfStock(outOfStock uint32) Option {
	return optionFunc(func(o *options) { o.query["out_of_stock"] = outOfStock })
}

func (obj *_ProductMgr) WithAdditionalDeliveryTimes(additionalDeliveryTimes bool) Option {
	return optionFunc(func(o *options) { o.query["additional_delivery_times"] = additionalDeliveryTimes })
}

func (obj *_ProductMgr) WithQuantityDiscount(quantityDiscount bool) Option {
	return optionFunc(func(o *options) { o.query["quantity_discount"] = quantityDiscount })
}

func (obj *_ProductMgr) WithCustomizable(customizable int8) Option {
	return optionFunc(func(o *options) { o.query["customizable"] = customizable })
}

func (obj *_ProductMgr) WithUploadableFiles(uploadableFiles int8) Option {
	return optionFunc(func(o *options) { o.query["uploadable_files"] = uploadableFiles })
}

func (obj *_ProductMgr) WithTextFields(textFields int8) Option {
	return optionFunc(func(o *options) { o.query["text_fields"] = textFields })
}

func (obj *_ProductMgr) WithActive(active bool) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}

func (obj *_ProductMgr) WithRedirectType(redirectType string) Option {
	return optionFunc(func(o *options) { o.query["redirect_type"] = redirectType })
}

func (obj *_ProductMgr) WithIDTypeRedirected(idTypeRedirected uint32) Option {
	return optionFunc(func(o *options) { o.query["id_type_redirected"] = idTypeRedirected })
}

func (obj *_ProductMgr) WithAvailableForOrder(availableForOrder bool) Option {
	return optionFunc(func(o *options) { o.query["available_for_order"] = availableForOrder })
}

func (obj *_ProductMgr) WithAvailableDate(availableDate datatypes.Date) Option {
	return optionFunc(func(o *options) { o.query["available_date"] = availableDate })
}

func (obj *_ProductMgr) WithShowCondition(showCondition bool) Option {
	return optionFunc(func(o *options) { o.query["show_condition"] = showCondition })
}

func (obj *_ProductMgr) WithCondition(condition string) Option {
	return optionFunc(func(o *options) { o.query["condition"] = condition })
}

func (obj *_ProductMgr) WithShowPrice(showPrice bool) Option {
	return optionFunc(func(o *options) { o.query["show_price"] = showPrice })
}

func (obj *_ProductMgr) WithIndexed(indexed bool) Option {
	return optionFunc(func(o *options) { o.query["indexed"] = indexed })
}

func (obj *_ProductMgr) WithVisibility(visibility string) Option {
	return optionFunc(func(o *options) { o.query["visibility"] = visibility })
}

func (obj *_ProductMgr) WithCacheIsPack(cacheIsPack bool) Option {
	return optionFunc(func(o *options) { o.query["cache_is_pack"] = cacheIsPack })
}

func (obj *_ProductMgr) WithCacheHasAttachments(cacheHasAttachments bool) Option {
	return optionFunc(func(o *options) { o.query["cache_has_attachments"] = cacheHasAttachments })
}

func (obj *_ProductMgr) WithIsVirtual(isVirtual bool) Option {
	return optionFunc(func(o *options) { o.query["is_virtual"] = isVirtual })
}

func (obj *_ProductMgr) WithCacheDefaultAttribute(cacheDefaultAttribute uint32) Option {
	return optionFunc(func(o *options) { o.query["cache_default_attribute"] = cacheDefaultAttribute })
}

func (obj *_ProductMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

func (obj *_ProductMgr) WithDateUpd(dateUpd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_upd"] = dateUpd })
}

func (obj *_ProductMgr) WithAdvancedStockManagement(advancedStockManagement bool) Option {
	return optionFunc(func(o *options) { o.query["advanced_stock_management"] = advancedStockManagement })
}

func (obj *_ProductMgr) WithPackStockType(packStockType uint32) Option {
	return optionFunc(func(o *options) { o.query["pack_stock_type"] = packStockType })
}

func (obj *_ProductMgr) WithState(state uint32) Option {
	return optionFunc(func(o *options) { o.query["state"] = state })
}

func (obj *_ProductMgr) GetByOption(opts ...Option) (result Product, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_ProductMgr) GetByOptions(opts ...Option) (results []*Product, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetFromIDProduct(idProduct uint32) (result Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ?", idProduct).Find(&result).Error

	return
}

func (obj *_ProductMgr) GetBatchFromIDProduct(idProducts []uint32) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product IN (?)", idProducts).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetFromIDSupplier(idSupplier uint32) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supplier = ?", idSupplier).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetBatchFromIDSupplier(idSuppliers []uint32) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supplier IN (?)", idSuppliers).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetFromIDManufacturer(idManufacturer uint32) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_manufacturer = ?", idManufacturer).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetBatchFromIDManufacturer(idManufacturers []uint32) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_manufacturer IN (?)", idManufacturers).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetFromIDCategoryDefault(idCategoryDefault uint32) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_category_default = ?", idCategoryDefault).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetBatchFromIDCategoryDefault(idCategoryDefaults []uint32) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_category_default IN (?)", idCategoryDefaults).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetFromIDShopDefault(idShopDefault uint32) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_default = ?", idShopDefault).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetBatchFromIDShopDefault(idShopDefaults []uint32) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_default IN (?)", idShopDefaults).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetFromIDTaxRulesGroup(idTaxRulesGroup uint32) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tax_rules_group = ?", idTaxRulesGroup).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetBatchFromIDTaxRulesGroup(idTaxRulesGroups []uint32) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tax_rules_group IN (?)", idTaxRulesGroups).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetFromOnSale(onSale bool) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("on_sale = ?", onSale).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetBatchFromOnSale(onSales []bool) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("on_sale IN (?)", onSales).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetFromOnlineOnly(onlineOnly bool) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("online_only = ?", onlineOnly).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetBatchFromOnlineOnly(onlineOnlys []bool) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("online_only IN (?)", onlineOnlys).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetFromEan13(ean13 string) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ean13 = ?", ean13).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetBatchFromEan13(ean13s []string) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ean13 IN (?)", ean13s).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetFromIsbn(isbn string) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("isbn = ?", isbn).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetBatchFromIsbn(isbns []string) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("isbn IN (?)", isbns).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetFromUpc(upc string) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("upc = ?", upc).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetBatchFromUpc(upcs []string) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("upc IN (?)", upcs).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetFromMpn(mpn string) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("mpn = ?", mpn).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetBatchFromMpn(mpns []string) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("mpn IN (?)", mpns).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetFromEcotax(ecotax float64) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ecotax = ?", ecotax).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetBatchFromEcotax(ecotaxs []float64) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ecotax IN (?)", ecotaxs).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetFromQuantity(quantity int) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("quantity = ?", quantity).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetBatchFromQuantity(quantitys []int) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("quantity IN (?)", quantitys).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetFromMinimalQuantity(minimalQuantity uint32) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("minimal_quantity = ?", minimalQuantity).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetBatchFromMinimalQuantity(minimalQuantitys []uint32) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("minimal_quantity IN (?)", minimalQuantitys).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetFromLowStockThreshold(lowStockThreshold int) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("low_stock_threshold = ?", lowStockThreshold).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetBatchFromLowStockThreshold(lowStockThresholds []int) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("low_stock_threshold IN (?)", lowStockThresholds).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetFromLowStockAlert(lowStockAlert bool) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("low_stock_alert = ?", lowStockAlert).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetBatchFromLowStockAlert(lowStockAlerts []bool) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("low_stock_alert IN (?)", lowStockAlerts).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetFromPrice(price float64) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("price = ?", price).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetBatchFromPrice(prices []float64) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("price IN (?)", prices).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetFromWholesalePrice(wholesalePrice float64) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("wholesale_price = ?", wholesalePrice).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetBatchFromWholesalePrice(wholesalePrices []float64) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("wholesale_price IN (?)", wholesalePrices).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetFromUnity(unity string) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("unity = ?", unity).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetBatchFromUnity(unitys []string) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("unity IN (?)", unitys).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetFromUnitPriceRatio(unitPriceRatio float64) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("unit_price_ratio = ?", unitPriceRatio).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetBatchFromUnitPriceRatio(unitPriceRatios []float64) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("unit_price_ratio IN (?)", unitPriceRatios).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetFromAdditionalShippingCost(additionalShippingCost float64) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("additional_shipping_cost = ?", additionalShippingCost).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetBatchFromAdditionalShippingCost(additionalShippingCosts []float64) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("additional_shipping_cost IN (?)", additionalShippingCosts).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetFromReference(reference string) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reference = ?", reference).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetBatchFromReference(references []string) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reference IN (?)", references).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetFromSupplierReference(supplierReference string) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("supplier_reference = ?", supplierReference).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetBatchFromSupplierReference(supplierReferences []string) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("supplier_reference IN (?)", supplierReferences).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetFromLocation(location string) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("location = ?", location).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetBatchFromLocation(locations []string) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("location IN (?)", locations).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetFromWidth(width float64) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("width = ?", width).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetBatchFromWidth(widths []float64) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("width IN (?)", widths).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetFromHeight(height float64) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("height = ?", height).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetBatchFromHeight(heights []float64) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("height IN (?)", heights).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetFromDepth(depth float64) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("depth = ?", depth).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetBatchFromDepth(depths []float64) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("depth IN (?)", depths).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetFromWeight(weight float64) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("weight = ?", weight).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetBatchFromWeight(weights []float64) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("weight IN (?)", weights).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetFromOutOfStock(outOfStock uint32) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("out_of_stock = ?", outOfStock).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetBatchFromOutOfStock(outOfStocks []uint32) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("out_of_stock IN (?)", outOfStocks).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetFromAdditionalDeliveryTimes(additionalDeliveryTimes bool) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("additional_delivery_times = ?", additionalDeliveryTimes).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetBatchFromAdditionalDeliveryTimes(additionalDeliveryTimess []bool) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("additional_delivery_times IN (?)", additionalDeliveryTimess).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetFromQuantityDiscount(quantityDiscount bool) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("quantity_discount = ?", quantityDiscount).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetBatchFromQuantityDiscount(quantityDiscounts []bool) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("quantity_discount IN (?)", quantityDiscounts).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetFromCustomizable(customizable int8) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("customizable = ?", customizable).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetBatchFromCustomizable(customizables []int8) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("customizable IN (?)", customizables).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetFromUploadableFiles(uploadableFiles int8) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("uploadable_files = ?", uploadableFiles).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetBatchFromUploadableFiles(uploadableFiless []int8) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("uploadable_files IN (?)", uploadableFiless).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetFromTextFields(textFields int8) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("text_fields = ?", textFields).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetBatchFromTextFields(textFieldss []int8) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("text_fields IN (?)", textFieldss).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetFromActive(active bool) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active = ?", active).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetBatchFromActive(actives []bool) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active IN (?)", actives).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetFromRedirectType(redirectType string) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("redirect_type = ?", redirectType).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetBatchFromRedirectType(redirectTypes []string) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("redirect_type IN (?)", redirectTypes).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetFromIDTypeRedirected(idTypeRedirected uint32) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_type_redirected = ?", idTypeRedirected).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetBatchFromIDTypeRedirected(idTypeRedirecteds []uint32) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_type_redirected IN (?)", idTypeRedirecteds).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetFromAvailableForOrder(availableForOrder bool) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("available_for_order = ?", availableForOrder).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetBatchFromAvailableForOrder(availableForOrders []bool) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("available_for_order IN (?)", availableForOrders).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetFromAvailableDate(availableDate datatypes.Date) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("available_date = ?", availableDate).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetBatchFromAvailableDate(availableDates []datatypes.Date) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("available_date IN (?)", availableDates).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetFromShowCondition(showCondition bool) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("show_condition = ?", showCondition).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetBatchFromShowCondition(showConditions []bool) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("show_condition IN (?)", showConditions).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetFromCondition(condition string) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("condition = ?", condition).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetBatchFromCondition(conditions []string) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("condition IN (?)", conditions).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetFromShowPrice(showPrice bool) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("show_price = ?", showPrice).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetBatchFromShowPrice(showPrices []bool) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("show_price IN (?)", showPrices).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetFromIndexed(indexed bool) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("indexed = ?", indexed).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetBatchFromIndexed(indexeds []bool) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("indexed IN (?)", indexeds).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetFromVisibility(visibility string) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("visibility = ?", visibility).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetBatchFromVisibility(visibilitys []string) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("visibility IN (?)", visibilitys).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetFromCacheIsPack(cacheIsPack bool) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cache_is_pack = ?", cacheIsPack).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetBatchFromCacheIsPack(cacheIsPacks []bool) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cache_is_pack IN (?)", cacheIsPacks).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetFromCacheHasAttachments(cacheHasAttachments bool) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cache_has_attachments = ?", cacheHasAttachments).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetBatchFromCacheHasAttachments(cacheHasAttachmentss []bool) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cache_has_attachments IN (?)", cacheHasAttachmentss).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetFromIsVirtual(isVirtual bool) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_virtual = ?", isVirtual).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetBatchFromIsVirtual(isVirtuals []bool) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_virtual IN (?)", isVirtuals).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetFromCacheDefaultAttribute(cacheDefaultAttribute uint32) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cache_default_attribute = ?", cacheDefaultAttribute).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetBatchFromCacheDefaultAttribute(cacheDefaultAttributes []uint32) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cache_default_attribute IN (?)", cacheDefaultAttributes).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetFromDateAdd(dateAdd time.Time) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetFromDateUpd(dateUpd time.Time) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd = ?", dateUpd).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetBatchFromDateUpd(dateUpds []time.Time) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd IN (?)", dateUpds).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetFromAdvancedStockManagement(advancedStockManagement bool) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("advanced_stock_management = ?", advancedStockManagement).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetBatchFromAdvancedStockManagement(advancedStockManagements []bool) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("advanced_stock_management IN (?)", advancedStockManagements).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetFromPackStockType(packStockType uint32) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("pack_stock_type = ?", packStockType).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetBatchFromPackStockType(packStockTypes []uint32) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("pack_stock_type IN (?)", packStockTypes).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetFromState(state uint32) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("state = ?", state).Find(&results).Error

	return
}

func (obj *_ProductMgr) GetBatchFromState(states []uint32) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("state IN (?)", states).Find(&results).Error

	return
}

func (obj *_ProductMgr) FetchByPrimaryKey(idProduct uint32) (result Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ?", idProduct).Find(&result).Error

	return
}

func (obj *_ProductMgr) FetchIndexByProductManufacturer(idProduct uint32, idManufacturer uint32) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ? AND id_manufacturer = ?", idProduct, idManufacturer).Find(&results).Error

	return
}

func (obj *_ProductMgr) FetchIndexByProductSupplier(idSupplier uint32) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supplier = ?", idSupplier).Find(&results).Error

	return
}

func (obj *_ProductMgr) FetchIndexByIDCategoryDefault(idCategoryDefault uint32) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_category_default = ?", idCategoryDefault).Find(&results).Error

	return
}

func (obj *_ProductMgr) FetchIndexByReferenceIDx(reference string) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reference = ?", reference).Find(&results).Error

	return
}

func (obj *_ProductMgr) FetchIndexBySupplierReferenceIDx(supplierReference string) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("supplier_reference = ?", supplierReference).Find(&results).Error

	return
}

func (obj *_ProductMgr) FetchIndexByIndexed(indexed bool) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("indexed = ?", indexed).Find(&results).Error

	return
}

func (obj *_ProductMgr) FetchIndexByDateAdd(dateAdd time.Time) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error

	return
}

func (obj *_ProductMgr) FetchIndexByState(dateUpd time.Time, state uint32) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd = ? AND state = ?", dateUpd, state).Find(&results).Error

	return
}
