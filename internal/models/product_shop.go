package models

import (
	"context"
	"fmt"
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type _ProductShopMgr struct {
	*_BaseMgr
}

func ProductShopMgr(db *gorm.DB) *_ProductShopMgr {
	if db == nil {
		panic(fmt.Errorf("ProductShopMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ProductShopMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_product_shop"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_ProductShopMgr) GetTableName() string {
	return "ps_product_shop"
}

func (obj *_ProductShopMgr) Get() (result ProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_ProductShopMgr) Gets() (results []*ProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_ProductShopMgr) WithIDProduct(idProduct uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product"] = idProduct })
}

func (obj *_ProductShopMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

func (obj *_ProductShopMgr) WithIDCategoryDefault(idCategoryDefault uint32) Option {
	return optionFunc(func(o *options) { o.query["id_category_default"] = idCategoryDefault })
}

func (obj *_ProductShopMgr) WithIDTaxRulesGroup(idTaxRulesGroup uint32) Option {
	return optionFunc(func(o *options) { o.query["id_tax_rules_group"] = idTaxRulesGroup })
}

func (obj *_ProductShopMgr) WithOnSale(onSale bool) Option {
	return optionFunc(func(o *options) { o.query["on_sale"] = onSale })
}

func (obj *_ProductShopMgr) WithOnlineOnly(onlineOnly bool) Option {
	return optionFunc(func(o *options) { o.query["online_only"] = onlineOnly })
}

func (obj *_ProductShopMgr) WithEcotax(ecotax float64) Option {
	return optionFunc(func(o *options) { o.query["ecotax"] = ecotax })
}

func (obj *_ProductShopMgr) WithMinimalQuantity(minimalQuantity uint32) Option {
	return optionFunc(func(o *options) { o.query["minimal_quantity"] = minimalQuantity })
}

func (obj *_ProductShopMgr) WithLowStockThreshold(lowStockThreshold int) Option {
	return optionFunc(func(o *options) { o.query["low_stock_threshold"] = lowStockThreshold })
}

func (obj *_ProductShopMgr) WithLowStockAlert(lowStockAlert bool) Option {
	return optionFunc(func(o *options) { o.query["low_stock_alert"] = lowStockAlert })
}

func (obj *_ProductShopMgr) WithPrice(price float64) Option {
	return optionFunc(func(o *options) { o.query["price"] = price })
}

func (obj *_ProductShopMgr) WithWholesalePrice(wholesalePrice float64) Option {
	return optionFunc(func(o *options) { o.query["wholesale_price"] = wholesalePrice })
}

func (obj *_ProductShopMgr) WithUnity(unity string) Option {
	return optionFunc(func(o *options) { o.query["unity"] = unity })
}

func (obj *_ProductShopMgr) WithUnitPriceRatio(unitPriceRatio float64) Option {
	return optionFunc(func(o *options) { o.query["unit_price_ratio"] = unitPriceRatio })
}

func (obj *_ProductShopMgr) WithAdditionalShippingCost(additionalShippingCost float64) Option {
	return optionFunc(func(o *options) { o.query["additional_shipping_cost"] = additionalShippingCost })
}

func (obj *_ProductShopMgr) WithCustomizable(customizable int8) Option {
	return optionFunc(func(o *options) { o.query["customizable"] = customizable })
}

func (obj *_ProductShopMgr) WithUploadableFiles(uploadableFiles int8) Option {
	return optionFunc(func(o *options) { o.query["uploadable_files"] = uploadableFiles })
}

func (obj *_ProductShopMgr) WithTextFields(textFields int8) Option {
	return optionFunc(func(o *options) { o.query["text_fields"] = textFields })
}

func (obj *_ProductShopMgr) WithActive(active bool) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}

func (obj *_ProductShopMgr) WithRedirectType(redirectType string) Option {
	return optionFunc(func(o *options) { o.query["redirect_type"] = redirectType })
}

func (obj *_ProductShopMgr) WithIDTypeRedirected(idTypeRedirected uint32) Option {
	return optionFunc(func(o *options) { o.query["id_type_redirected"] = idTypeRedirected })
}

func (obj *_ProductShopMgr) WithAvailableForOrder(availableForOrder bool) Option {
	return optionFunc(func(o *options) { o.query["available_for_order"] = availableForOrder })
}

func (obj *_ProductShopMgr) WithAvailableDate(availableDate datatypes.Date) Option {
	return optionFunc(func(o *options) { o.query["available_date"] = availableDate })
}

func (obj *_ProductShopMgr) WithShowCondition(showCondition bool) Option {
	return optionFunc(func(o *options) { o.query["show_condition"] = showCondition })
}

func (obj *_ProductShopMgr) WithCondition(condition string) Option {
	return optionFunc(func(o *options) { o.query["condition"] = condition })
}

func (obj *_ProductShopMgr) WithShowPrice(showPrice bool) Option {
	return optionFunc(func(o *options) { o.query["show_price"] = showPrice })
}

func (obj *_ProductShopMgr) WithIndexed(indexed bool) Option {
	return optionFunc(func(o *options) { o.query["indexed"] = indexed })
}

func (obj *_ProductShopMgr) WithVisibility(visibility string) Option {
	return optionFunc(func(o *options) { o.query["visibility"] = visibility })
}

func (obj *_ProductShopMgr) WithCacheDefaultAttribute(cacheDefaultAttribute uint32) Option {
	return optionFunc(func(o *options) { o.query["cache_default_attribute"] = cacheDefaultAttribute })
}

func (obj *_ProductShopMgr) WithAdvancedStockManagement(advancedStockManagement bool) Option {
	return optionFunc(func(o *options) { o.query["advanced_stock_management"] = advancedStockManagement })
}

func (obj *_ProductShopMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

func (obj *_ProductShopMgr) WithDateUpd(dateUpd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_upd"] = dateUpd })
}

func (obj *_ProductShopMgr) WithPackStockType(packStockType uint32) Option {
	return optionFunc(func(o *options) { o.query["pack_stock_type"] = packStockType })
}

func (obj *_ProductShopMgr) GetByOption(opts ...Option) (result ProductShop, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_ProductShopMgr) GetByOptions(opts ...Option) (results []*ProductShop, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_ProductShopMgr) GetFromIDProduct(idProduct uint32) (results []*ProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ?", idProduct).Find(&results).Error

	return
}

func (obj *_ProductShopMgr) GetBatchFromIDProduct(idProducts []uint32) (results []*ProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product IN (?)", idProducts).Find(&results).Error

	return
}

func (obj *_ProductShopMgr) GetFromIDShop(idShop uint32) (results []*ProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

func (obj *_ProductShopMgr) GetBatchFromIDShop(idShops []uint32) (results []*ProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

func (obj *_ProductShopMgr) GetFromIDCategoryDefault(idCategoryDefault uint32) (results []*ProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_category_default = ?", idCategoryDefault).Find(&results).Error

	return
}

func (obj *_ProductShopMgr) GetBatchFromIDCategoryDefault(idCategoryDefaults []uint32) (results []*ProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_category_default IN (?)", idCategoryDefaults).Find(&results).Error

	return
}

func (obj *_ProductShopMgr) GetFromIDTaxRulesGroup(idTaxRulesGroup uint32) (results []*ProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tax_rules_group = ?", idTaxRulesGroup).Find(&results).Error

	return
}

func (obj *_ProductShopMgr) GetBatchFromIDTaxRulesGroup(idTaxRulesGroups []uint32) (results []*ProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tax_rules_group IN (?)", idTaxRulesGroups).Find(&results).Error

	return
}

func (obj *_ProductShopMgr) GetFromOnSale(onSale bool) (results []*ProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("on_sale = ?", onSale).Find(&results).Error

	return
}

func (obj *_ProductShopMgr) GetBatchFromOnSale(onSales []bool) (results []*ProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("on_sale IN (?)", onSales).Find(&results).Error

	return
}

func (obj *_ProductShopMgr) GetFromOnlineOnly(onlineOnly bool) (results []*ProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("online_only = ?", onlineOnly).Find(&results).Error

	return
}

func (obj *_ProductShopMgr) GetBatchFromOnlineOnly(onlineOnlys []bool) (results []*ProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("online_only IN (?)", onlineOnlys).Find(&results).Error

	return
}

func (obj *_ProductShopMgr) GetFromEcotax(ecotax float64) (results []*ProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ecotax = ?", ecotax).Find(&results).Error

	return
}

func (obj *_ProductShopMgr) GetBatchFromEcotax(ecotaxs []float64) (results []*ProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ecotax IN (?)", ecotaxs).Find(&results).Error

	return
}

func (obj *_ProductShopMgr) GetFromMinimalQuantity(minimalQuantity uint32) (results []*ProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("minimal_quantity = ?", minimalQuantity).Find(&results).Error

	return
}

func (obj *_ProductShopMgr) GetBatchFromMinimalQuantity(minimalQuantitys []uint32) (results []*ProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("minimal_quantity IN (?)", minimalQuantitys).Find(&results).Error

	return
}

func (obj *_ProductShopMgr) GetFromLowStockThreshold(lowStockThreshold int) (results []*ProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("low_stock_threshold = ?", lowStockThreshold).Find(&results).Error

	return
}

func (obj *_ProductShopMgr) GetBatchFromLowStockThreshold(lowStockThresholds []int) (results []*ProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("low_stock_threshold IN (?)", lowStockThresholds).Find(&results).Error

	return
}

func (obj *_ProductShopMgr) GetFromLowStockAlert(lowStockAlert bool) (results []*ProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("low_stock_alert = ?", lowStockAlert).Find(&results).Error

	return
}

func (obj *_ProductShopMgr) GetBatchFromLowStockAlert(lowStockAlerts []bool) (results []*ProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("low_stock_alert IN (?)", lowStockAlerts).Find(&results).Error

	return
}

func (obj *_ProductShopMgr) GetFromPrice(price float64) (results []*ProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("price = ?", price).Find(&results).Error

	return
}

func (obj *_ProductShopMgr) GetBatchFromPrice(prices []float64) (results []*ProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("price IN (?)", prices).Find(&results).Error

	return
}

func (obj *_ProductShopMgr) GetFromWholesalePrice(wholesalePrice float64) (results []*ProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("wholesale_price = ?", wholesalePrice).Find(&results).Error

	return
}

func (obj *_ProductShopMgr) GetBatchFromWholesalePrice(wholesalePrices []float64) (results []*ProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("wholesale_price IN (?)", wholesalePrices).Find(&results).Error

	return
}

func (obj *_ProductShopMgr) GetFromUnity(unity string) (results []*ProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("unity = ?", unity).Find(&results).Error

	return
}

func (obj *_ProductShopMgr) GetBatchFromUnity(unitys []string) (results []*ProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("unity IN (?)", unitys).Find(&results).Error

	return
}

func (obj *_ProductShopMgr) GetFromUnitPriceRatio(unitPriceRatio float64) (results []*ProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("unit_price_ratio = ?", unitPriceRatio).Find(&results).Error

	return
}

func (obj *_ProductShopMgr) GetBatchFromUnitPriceRatio(unitPriceRatios []float64) (results []*ProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("unit_price_ratio IN (?)", unitPriceRatios).Find(&results).Error

	return
}

func (obj *_ProductShopMgr) GetFromAdditionalShippingCost(additionalShippingCost float64) (results []*ProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("additional_shipping_cost = ?", additionalShippingCost).Find(&results).Error

	return
}

func (obj *_ProductShopMgr) GetBatchFromAdditionalShippingCost(additionalShippingCosts []float64) (results []*ProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("additional_shipping_cost IN (?)", additionalShippingCosts).Find(&results).Error

	return
}

func (obj *_ProductShopMgr) GetFromCustomizable(customizable int8) (results []*ProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("customizable = ?", customizable).Find(&results).Error

	return
}

func (obj *_ProductShopMgr) GetBatchFromCustomizable(customizables []int8) (results []*ProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("customizable IN (?)", customizables).Find(&results).Error

	return
}

func (obj *_ProductShopMgr) GetFromUploadableFiles(uploadableFiles int8) (results []*ProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("uploadable_files = ?", uploadableFiles).Find(&results).Error

	return
}

func (obj *_ProductShopMgr) GetBatchFromUploadableFiles(uploadableFiless []int8) (results []*ProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("uploadable_files IN (?)", uploadableFiless).Find(&results).Error

	return
}

func (obj *_ProductShopMgr) GetFromTextFields(textFields int8) (results []*ProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("text_fields = ?", textFields).Find(&results).Error

	return
}

func (obj *_ProductShopMgr) GetBatchFromTextFields(textFieldss []int8) (results []*ProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("text_fields IN (?)", textFieldss).Find(&results).Error

	return
}

func (obj *_ProductShopMgr) GetFromActive(active bool) (results []*ProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active = ?", active).Find(&results).Error

	return
}

func (obj *_ProductShopMgr) GetBatchFromActive(actives []bool) (results []*ProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active IN (?)", actives).Find(&results).Error

	return
}

func (obj *_ProductShopMgr) GetFromRedirectType(redirectType string) (results []*ProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("redirect_type = ?", redirectType).Find(&results).Error

	return
}

func (obj *_ProductShopMgr) GetBatchFromRedirectType(redirectTypes []string) (results []*ProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("redirect_type IN (?)", redirectTypes).Find(&results).Error

	return
}

func (obj *_ProductShopMgr) GetFromIDTypeRedirected(idTypeRedirected uint32) (results []*ProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_type_redirected = ?", idTypeRedirected).Find(&results).Error

	return
}

func (obj *_ProductShopMgr) GetBatchFromIDTypeRedirected(idTypeRedirecteds []uint32) (results []*ProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_type_redirected IN (?)", idTypeRedirecteds).Find(&results).Error

	return
}

func (obj *_ProductShopMgr) GetFromAvailableForOrder(availableForOrder bool) (results []*ProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("available_for_order = ?", availableForOrder).Find(&results).Error

	return
}

func (obj *_ProductShopMgr) GetBatchFromAvailableForOrder(availableForOrders []bool) (results []*ProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("available_for_order IN (?)", availableForOrders).Find(&results).Error

	return
}

func (obj *_ProductShopMgr) GetFromAvailableDate(availableDate datatypes.Date) (results []*ProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("available_date = ?", availableDate).Find(&results).Error

	return
}

func (obj *_ProductShopMgr) GetBatchFromAvailableDate(availableDates []datatypes.Date) (results []*ProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("available_date IN (?)", availableDates).Find(&results).Error

	return
}

func (obj *_ProductShopMgr) GetFromShowCondition(showCondition bool) (results []*ProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("show_condition = ?", showCondition).Find(&results).Error

	return
}

func (obj *_ProductShopMgr) GetBatchFromShowCondition(showConditions []bool) (results []*ProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("show_condition IN (?)", showConditions).Find(&results).Error

	return
}

func (obj *_ProductShopMgr) GetFromCondition(condition string) (results []*ProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("condition = ?", condition).Find(&results).Error

	return
}

func (obj *_ProductShopMgr) GetBatchFromCondition(conditions []string) (results []*ProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("condition IN (?)", conditions).Find(&results).Error

	return
}

func (obj *_ProductShopMgr) GetFromShowPrice(showPrice bool) (results []*ProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("show_price = ?", showPrice).Find(&results).Error

	return
}

func (obj *_ProductShopMgr) GetBatchFromShowPrice(showPrices []bool) (results []*ProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("show_price IN (?)", showPrices).Find(&results).Error

	return
}

func (obj *_ProductShopMgr) GetFromIndexed(indexed bool) (results []*ProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("indexed = ?", indexed).Find(&results).Error

	return
}

func (obj *_ProductShopMgr) GetBatchFromIndexed(indexeds []bool) (results []*ProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("indexed IN (?)", indexeds).Find(&results).Error

	return
}

func (obj *_ProductShopMgr) GetFromVisibility(visibility string) (results []*ProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("visibility = ?", visibility).Find(&results).Error

	return
}

func (obj *_ProductShopMgr) GetBatchFromVisibility(visibilitys []string) (results []*ProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("visibility IN (?)", visibilitys).Find(&results).Error

	return
}

func (obj *_ProductShopMgr) GetFromCacheDefaultAttribute(cacheDefaultAttribute uint32) (results []*ProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cache_default_attribute = ?", cacheDefaultAttribute).Find(&results).Error

	return
}

func (obj *_ProductShopMgr) GetBatchFromCacheDefaultAttribute(cacheDefaultAttributes []uint32) (results []*ProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cache_default_attribute IN (?)", cacheDefaultAttributes).Find(&results).Error

	return
}

func (obj *_ProductShopMgr) GetFromAdvancedStockManagement(advancedStockManagement bool) (results []*ProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("advanced_stock_management = ?", advancedStockManagement).Find(&results).Error

	return
}

func (obj *_ProductShopMgr) GetBatchFromAdvancedStockManagement(advancedStockManagements []bool) (results []*ProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("advanced_stock_management IN (?)", advancedStockManagements).Find(&results).Error

	return
}

func (obj *_ProductShopMgr) GetFromDateAdd(dateAdd time.Time) (results []*ProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error

	return
}

func (obj *_ProductShopMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*ProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error

	return
}

func (obj *_ProductShopMgr) GetFromDateUpd(dateUpd time.Time) (results []*ProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd = ?", dateUpd).Find(&results).Error

	return
}

func (obj *_ProductShopMgr) GetBatchFromDateUpd(dateUpds []time.Time) (results []*ProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd IN (?)", dateUpds).Find(&results).Error

	return
}

func (obj *_ProductShopMgr) GetFromPackStockType(packStockType uint32) (results []*ProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("pack_stock_type = ?", packStockType).Find(&results).Error

	return
}

func (obj *_ProductShopMgr) GetBatchFromPackStockType(packStockTypes []uint32) (results []*ProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("pack_stock_type IN (?)", packStockTypes).Find(&results).Error

	return
}


func (obj *_ProductShopMgr) FetchByPrimaryKey(idProduct uint32, idShop uint32) (result ProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ? AND id_shop = ?", idProduct, idShop).Find(&result).Error

	return
}

func (obj *_ProductShopMgr) FetchIndexByIndexed(idProduct uint32, active bool, indexed bool) (results []*ProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ? AND active = ? AND indexed = ?", idProduct, active, indexed).Find(&results).Error

	return
}

func (obj *_ProductShopMgr) FetchIndexByIDCategoryDefault(idCategoryDefault uint32) (results []*ProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_category_default = ?", idCategoryDefault).Find(&results).Error

	return
}

func (obj *_ProductShopMgr) FetchIndexByDateAdd(active bool, visibility string, dateAdd time.Time) (results []*ProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active = ? AND visibility = ? AND date_add = ?", active, visibility, dateAdd).Find(&results).Error

	return
}
