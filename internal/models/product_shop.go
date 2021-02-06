package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
"gorm.io/datatypes"	
"time"	
)	

type _EgProductShopMgr struct {
	*_BaseMgr
}

// EgProductShopMgr open func
func EgProductShopMgr(db *gorm.DB) *_EgProductShopMgr {
	if db == nil {
		panic(fmt.Errorf("EgProductShopMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgProductShopMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_product_shop"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgProductShopMgr) GetTableName() string {
	return "eg_product_shop"
}

// Get 获取
func (obj *_EgProductShopMgr) Get() (result EgProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgProductShopMgr) Gets() (results []*EgProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDProduct id_product获取 
func (obj *_EgProductShopMgr) WithIDProduct(idProduct uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product"] = idProduct })
}

// WithIDShop id_shop获取 
func (obj *_EgProductShopMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

// WithIDCategoryDefault id_category_default获取 
func (obj *_EgProductShopMgr) WithIDCategoryDefault(idCategoryDefault uint32) Option {
	return optionFunc(func(o *options) { o.query["id_category_default"] = idCategoryDefault })
}

// WithIDTaxRulesGroup id_tax_rules_group获取 
func (obj *_EgProductShopMgr) WithIDTaxRulesGroup(idTaxRulesGroup uint32) Option {
	return optionFunc(func(o *options) { o.query["id_tax_rules_group"] = idTaxRulesGroup })
}

// WithOnSale on_sale获取 
func (obj *_EgProductShopMgr) WithOnSale(onSale bool) Option {
	return optionFunc(func(o *options) { o.query["on_sale"] = onSale })
}

// WithOnlineOnly online_only获取 
func (obj *_EgProductShopMgr) WithOnlineOnly(onlineOnly bool) Option {
	return optionFunc(func(o *options) { o.query["online_only"] = onlineOnly })
}

// WithEcotax ecotax获取 
func (obj *_EgProductShopMgr) WithEcotax(ecotax float64) Option {
	return optionFunc(func(o *options) { o.query["ecotax"] = ecotax })
}

// WithMinimalQuantity minimal_quantity获取 
func (obj *_EgProductShopMgr) WithMinimalQuantity(minimalQuantity uint32) Option {
	return optionFunc(func(o *options) { o.query["minimal_quantity"] = minimalQuantity })
}

// WithLowStockThreshold low_stock_threshold获取 
func (obj *_EgProductShopMgr) WithLowStockThreshold(lowStockThreshold int) Option {
	return optionFunc(func(o *options) { o.query["low_stock_threshold"] = lowStockThreshold })
}

// WithLowStockAlert low_stock_alert获取 
func (obj *_EgProductShopMgr) WithLowStockAlert(lowStockAlert bool) Option {
	return optionFunc(func(o *options) { o.query["low_stock_alert"] = lowStockAlert })
}

// WithPrice price获取 
func (obj *_EgProductShopMgr) WithPrice(price float64) Option {
	return optionFunc(func(o *options) { o.query["price"] = price })
}

// WithWholesalePrice wholesale_price获取 
func (obj *_EgProductShopMgr) WithWholesalePrice(wholesalePrice float64) Option {
	return optionFunc(func(o *options) { o.query["wholesale_price"] = wholesalePrice })
}

// WithUnity unity获取 
func (obj *_EgProductShopMgr) WithUnity(unity string) Option {
	return optionFunc(func(o *options) { o.query["unity"] = unity })
}

// WithUnitPriceRatio unit_price_ratio获取 
func (obj *_EgProductShopMgr) WithUnitPriceRatio(unitPriceRatio float64) Option {
	return optionFunc(func(o *options) { o.query["unit_price_ratio"] = unitPriceRatio })
}

// WithAdditionalShippingCost additional_shipping_cost获取 
func (obj *_EgProductShopMgr) WithAdditionalShippingCost(additionalShippingCost float64) Option {
	return optionFunc(func(o *options) { o.query["additional_shipping_cost"] = additionalShippingCost })
}

// WithCustomizable customizable获取 
func (obj *_EgProductShopMgr) WithCustomizable(customizable int8) Option {
	return optionFunc(func(o *options) { o.query["customizable"] = customizable })
}

// WithUploadableFiles uploadable_files获取 
func (obj *_EgProductShopMgr) WithUploadableFiles(uploadableFiles int8) Option {
	return optionFunc(func(o *options) { o.query["uploadable_files"] = uploadableFiles })
}

// WithTextFields text_fields获取 
func (obj *_EgProductShopMgr) WithTextFields(textFields int8) Option {
	return optionFunc(func(o *options) { o.query["text_fields"] = textFields })
}

// WithActive active获取 
func (obj *_EgProductShopMgr) WithActive(active bool) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}

// WithRedirectType redirect_type获取 
func (obj *_EgProductShopMgr) WithRedirectType(redirectType string) Option {
	return optionFunc(func(o *options) { o.query["redirect_type"] = redirectType })
}

// WithIDTypeRedirected id_type_redirected获取 
func (obj *_EgProductShopMgr) WithIDTypeRedirected(idTypeRedirected uint32) Option {
	return optionFunc(func(o *options) { o.query["id_type_redirected"] = idTypeRedirected })
}

// WithAvailableForOrder available_for_order获取 
func (obj *_EgProductShopMgr) WithAvailableForOrder(availableForOrder bool) Option {
	return optionFunc(func(o *options) { o.query["available_for_order"] = availableForOrder })
}

// WithAvailableDate available_date获取 
func (obj *_EgProductShopMgr) WithAvailableDate(availableDate datatypes.Date) Option {
	return optionFunc(func(o *options) { o.query["available_date"] = availableDate })
}

// WithShowCondition show_condition获取 
func (obj *_EgProductShopMgr) WithShowCondition(showCondition bool) Option {
	return optionFunc(func(o *options) { o.query["show_condition"] = showCondition })
}

// WithCondition condition获取 
func (obj *_EgProductShopMgr) WithCondition(condition string) Option {
	return optionFunc(func(o *options) { o.query["condition"] = condition })
}

// WithShowPrice show_price获取 
func (obj *_EgProductShopMgr) WithShowPrice(showPrice bool) Option {
	return optionFunc(func(o *options) { o.query["show_price"] = showPrice })
}

// WithIndexed indexed获取 
func (obj *_EgProductShopMgr) WithIndexed(indexed bool) Option {
	return optionFunc(func(o *options) { o.query["indexed"] = indexed })
}

// WithVisibility visibility获取 
func (obj *_EgProductShopMgr) WithVisibility(visibility string) Option {
	return optionFunc(func(o *options) { o.query["visibility"] = visibility })
}

// WithCacheDefaultAttribute cache_default_attribute获取 
func (obj *_EgProductShopMgr) WithCacheDefaultAttribute(cacheDefaultAttribute uint32) Option {
	return optionFunc(func(o *options) { o.query["cache_default_attribute"] = cacheDefaultAttribute })
}

// WithAdvancedStockManagement advanced_stock_management获取 
func (obj *_EgProductShopMgr) WithAdvancedStockManagement(advancedStockManagement bool) Option {
	return optionFunc(func(o *options) { o.query["advanced_stock_management"] = advancedStockManagement })
}

// WithDateAdd date_add获取 
func (obj *_EgProductShopMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

// WithDateUpd date_upd获取 
func (obj *_EgProductShopMgr) WithDateUpd(dateUpd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_upd"] = dateUpd })
}

// WithPackStockType pack_stock_type获取 
func (obj *_EgProductShopMgr) WithPackStockType(packStockType uint32) Option {
	return optionFunc(func(o *options) { o.query["pack_stock_type"] = packStockType })
}


// GetByOption 功能选项模式获取
func (obj *_EgProductShopMgr) GetByOption(opts ...Option) (result EgProductShop, err error) {
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
func (obj *_EgProductShopMgr) GetByOptions(opts ...Option) (results []*EgProductShop, err error) {
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


// GetFromIDProduct 通过id_product获取内容  
func (obj *_EgProductShopMgr) GetFromIDProduct(idProduct uint32) (results []*EgProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ?", idProduct).Find(&results).Error
	
	return
}

// GetBatchFromIDProduct 批量唯一主键查找 
func (obj *_EgProductShopMgr) GetBatchFromIDProduct(idProducts []uint32) (results []*EgProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product IN (?)", idProducts).Find(&results).Error
	
	return
}
 
// GetFromIDShop 通过id_shop获取内容  
func (obj *_EgProductShopMgr) GetFromIDShop(idShop uint32) (results []*EgProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error
	
	return
}

// GetBatchFromIDShop 批量唯一主键查找 
func (obj *_EgProductShopMgr) GetBatchFromIDShop(idShops []uint32) (results []*EgProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error
	
	return
}
 
// GetFromIDCategoryDefault 通过id_category_default获取内容  
func (obj *_EgProductShopMgr) GetFromIDCategoryDefault(idCategoryDefault uint32) (results []*EgProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_category_default = ?", idCategoryDefault).Find(&results).Error
	
	return
}

// GetBatchFromIDCategoryDefault 批量唯一主键查找 
func (obj *_EgProductShopMgr) GetBatchFromIDCategoryDefault(idCategoryDefaults []uint32) (results []*EgProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_category_default IN (?)", idCategoryDefaults).Find(&results).Error
	
	return
}
 
// GetFromIDTaxRulesGroup 通过id_tax_rules_group获取内容  
func (obj *_EgProductShopMgr) GetFromIDTaxRulesGroup(idTaxRulesGroup uint32) (results []*EgProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tax_rules_group = ?", idTaxRulesGroup).Find(&results).Error
	
	return
}

// GetBatchFromIDTaxRulesGroup 批量唯一主键查找 
func (obj *_EgProductShopMgr) GetBatchFromIDTaxRulesGroup(idTaxRulesGroups []uint32) (results []*EgProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tax_rules_group IN (?)", idTaxRulesGroups).Find(&results).Error
	
	return
}
 
// GetFromOnSale 通过on_sale获取内容  
func (obj *_EgProductShopMgr) GetFromOnSale(onSale bool) (results []*EgProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("on_sale = ?", onSale).Find(&results).Error
	
	return
}

// GetBatchFromOnSale 批量唯一主键查找 
func (obj *_EgProductShopMgr) GetBatchFromOnSale(onSales []bool) (results []*EgProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("on_sale IN (?)", onSales).Find(&results).Error
	
	return
}
 
// GetFromOnlineOnly 通过online_only获取内容  
func (obj *_EgProductShopMgr) GetFromOnlineOnly(onlineOnly bool) (results []*EgProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("online_only = ?", onlineOnly).Find(&results).Error
	
	return
}

// GetBatchFromOnlineOnly 批量唯一主键查找 
func (obj *_EgProductShopMgr) GetBatchFromOnlineOnly(onlineOnlys []bool) (results []*EgProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("online_only IN (?)", onlineOnlys).Find(&results).Error
	
	return
}
 
// GetFromEcotax 通过ecotax获取内容  
func (obj *_EgProductShopMgr) GetFromEcotax(ecotax float64) (results []*EgProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ecotax = ?", ecotax).Find(&results).Error
	
	return
}

// GetBatchFromEcotax 批量唯一主键查找 
func (obj *_EgProductShopMgr) GetBatchFromEcotax(ecotaxs []float64) (results []*EgProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ecotax IN (?)", ecotaxs).Find(&results).Error
	
	return
}
 
// GetFromMinimalQuantity 通过minimal_quantity获取内容  
func (obj *_EgProductShopMgr) GetFromMinimalQuantity(minimalQuantity uint32) (results []*EgProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("minimal_quantity = ?", minimalQuantity).Find(&results).Error
	
	return
}

// GetBatchFromMinimalQuantity 批量唯一主键查找 
func (obj *_EgProductShopMgr) GetBatchFromMinimalQuantity(minimalQuantitys []uint32) (results []*EgProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("minimal_quantity IN (?)", minimalQuantitys).Find(&results).Error
	
	return
}
 
// GetFromLowStockThreshold 通过low_stock_threshold获取内容  
func (obj *_EgProductShopMgr) GetFromLowStockThreshold(lowStockThreshold int) (results []*EgProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("low_stock_threshold = ?", lowStockThreshold).Find(&results).Error
	
	return
}

// GetBatchFromLowStockThreshold 批量唯一主键查找 
func (obj *_EgProductShopMgr) GetBatchFromLowStockThreshold(lowStockThresholds []int) (results []*EgProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("low_stock_threshold IN (?)", lowStockThresholds).Find(&results).Error
	
	return
}
 
// GetFromLowStockAlert 通过low_stock_alert获取内容  
func (obj *_EgProductShopMgr) GetFromLowStockAlert(lowStockAlert bool) (results []*EgProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("low_stock_alert = ?", lowStockAlert).Find(&results).Error
	
	return
}

// GetBatchFromLowStockAlert 批量唯一主键查找 
func (obj *_EgProductShopMgr) GetBatchFromLowStockAlert(lowStockAlerts []bool) (results []*EgProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("low_stock_alert IN (?)", lowStockAlerts).Find(&results).Error
	
	return
}
 
// GetFromPrice 通过price获取内容  
func (obj *_EgProductShopMgr) GetFromPrice(price float64) (results []*EgProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("price = ?", price).Find(&results).Error
	
	return
}

// GetBatchFromPrice 批量唯一主键查找 
func (obj *_EgProductShopMgr) GetBatchFromPrice(prices []float64) (results []*EgProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("price IN (?)", prices).Find(&results).Error
	
	return
}
 
// GetFromWholesalePrice 通过wholesale_price获取内容  
func (obj *_EgProductShopMgr) GetFromWholesalePrice(wholesalePrice float64) (results []*EgProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("wholesale_price = ?", wholesalePrice).Find(&results).Error
	
	return
}

// GetBatchFromWholesalePrice 批量唯一主键查找 
func (obj *_EgProductShopMgr) GetBatchFromWholesalePrice(wholesalePrices []float64) (results []*EgProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("wholesale_price IN (?)", wholesalePrices).Find(&results).Error
	
	return
}
 
// GetFromUnity 通过unity获取内容  
func (obj *_EgProductShopMgr) GetFromUnity(unity string) (results []*EgProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("unity = ?", unity).Find(&results).Error
	
	return
}

// GetBatchFromUnity 批量唯一主键查找 
func (obj *_EgProductShopMgr) GetBatchFromUnity(unitys []string) (results []*EgProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("unity IN (?)", unitys).Find(&results).Error
	
	return
}
 
// GetFromUnitPriceRatio 通过unit_price_ratio获取内容  
func (obj *_EgProductShopMgr) GetFromUnitPriceRatio(unitPriceRatio float64) (results []*EgProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("unit_price_ratio = ?", unitPriceRatio).Find(&results).Error
	
	return
}

// GetBatchFromUnitPriceRatio 批量唯一主键查找 
func (obj *_EgProductShopMgr) GetBatchFromUnitPriceRatio(unitPriceRatios []float64) (results []*EgProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("unit_price_ratio IN (?)", unitPriceRatios).Find(&results).Error
	
	return
}
 
// GetFromAdditionalShippingCost 通过additional_shipping_cost获取内容  
func (obj *_EgProductShopMgr) GetFromAdditionalShippingCost(additionalShippingCost float64) (results []*EgProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("additional_shipping_cost = ?", additionalShippingCost).Find(&results).Error
	
	return
}

// GetBatchFromAdditionalShippingCost 批量唯一主键查找 
func (obj *_EgProductShopMgr) GetBatchFromAdditionalShippingCost(additionalShippingCosts []float64) (results []*EgProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("additional_shipping_cost IN (?)", additionalShippingCosts).Find(&results).Error
	
	return
}
 
// GetFromCustomizable 通过customizable获取内容  
func (obj *_EgProductShopMgr) GetFromCustomizable(customizable int8) (results []*EgProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("customizable = ?", customizable).Find(&results).Error
	
	return
}

// GetBatchFromCustomizable 批量唯一主键查找 
func (obj *_EgProductShopMgr) GetBatchFromCustomizable(customizables []int8) (results []*EgProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("customizable IN (?)", customizables).Find(&results).Error
	
	return
}
 
// GetFromUploadableFiles 通过uploadable_files获取内容  
func (obj *_EgProductShopMgr) GetFromUploadableFiles(uploadableFiles int8) (results []*EgProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("uploadable_files = ?", uploadableFiles).Find(&results).Error
	
	return
}

// GetBatchFromUploadableFiles 批量唯一主键查找 
func (obj *_EgProductShopMgr) GetBatchFromUploadableFiles(uploadableFiless []int8) (results []*EgProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("uploadable_files IN (?)", uploadableFiless).Find(&results).Error
	
	return
}
 
// GetFromTextFields 通过text_fields获取内容  
func (obj *_EgProductShopMgr) GetFromTextFields(textFields int8) (results []*EgProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("text_fields = ?", textFields).Find(&results).Error
	
	return
}

// GetBatchFromTextFields 批量唯一主键查找 
func (obj *_EgProductShopMgr) GetBatchFromTextFields(textFieldss []int8) (results []*EgProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("text_fields IN (?)", textFieldss).Find(&results).Error
	
	return
}
 
// GetFromActive 通过active获取内容  
func (obj *_EgProductShopMgr) GetFromActive(active bool) (results []*EgProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active = ?", active).Find(&results).Error
	
	return
}

// GetBatchFromActive 批量唯一主键查找 
func (obj *_EgProductShopMgr) GetBatchFromActive(actives []bool) (results []*EgProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active IN (?)", actives).Find(&results).Error
	
	return
}
 
// GetFromRedirectType 通过redirect_type获取内容  
func (obj *_EgProductShopMgr) GetFromRedirectType(redirectType string) (results []*EgProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("redirect_type = ?", redirectType).Find(&results).Error
	
	return
}

// GetBatchFromRedirectType 批量唯一主键查找 
func (obj *_EgProductShopMgr) GetBatchFromRedirectType(redirectTypes []string) (results []*EgProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("redirect_type IN (?)", redirectTypes).Find(&results).Error
	
	return
}
 
// GetFromIDTypeRedirected 通过id_type_redirected获取内容  
func (obj *_EgProductShopMgr) GetFromIDTypeRedirected(idTypeRedirected uint32) (results []*EgProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_type_redirected = ?", idTypeRedirected).Find(&results).Error
	
	return
}

// GetBatchFromIDTypeRedirected 批量唯一主键查找 
func (obj *_EgProductShopMgr) GetBatchFromIDTypeRedirected(idTypeRedirecteds []uint32) (results []*EgProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_type_redirected IN (?)", idTypeRedirecteds).Find(&results).Error
	
	return
}
 
// GetFromAvailableForOrder 通过available_for_order获取内容  
func (obj *_EgProductShopMgr) GetFromAvailableForOrder(availableForOrder bool) (results []*EgProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("available_for_order = ?", availableForOrder).Find(&results).Error
	
	return
}

// GetBatchFromAvailableForOrder 批量唯一主键查找 
func (obj *_EgProductShopMgr) GetBatchFromAvailableForOrder(availableForOrders []bool) (results []*EgProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("available_for_order IN (?)", availableForOrders).Find(&results).Error
	
	return
}
 
// GetFromAvailableDate 通过available_date获取内容  
func (obj *_EgProductShopMgr) GetFromAvailableDate(availableDate datatypes.Date) (results []*EgProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("available_date = ?", availableDate).Find(&results).Error
	
	return
}

// GetBatchFromAvailableDate 批量唯一主键查找 
func (obj *_EgProductShopMgr) GetBatchFromAvailableDate(availableDates []datatypes.Date) (results []*EgProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("available_date IN (?)", availableDates).Find(&results).Error
	
	return
}
 
// GetFromShowCondition 通过show_condition获取内容  
func (obj *_EgProductShopMgr) GetFromShowCondition(showCondition bool) (results []*EgProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("show_condition = ?", showCondition).Find(&results).Error
	
	return
}

// GetBatchFromShowCondition 批量唯一主键查找 
func (obj *_EgProductShopMgr) GetBatchFromShowCondition(showConditions []bool) (results []*EgProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("show_condition IN (?)", showConditions).Find(&results).Error
	
	return
}
 
// GetFromCondition 通过condition获取内容  
func (obj *_EgProductShopMgr) GetFromCondition(condition string) (results []*EgProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("condition = ?", condition).Find(&results).Error
	
	return
}

// GetBatchFromCondition 批量唯一主键查找 
func (obj *_EgProductShopMgr) GetBatchFromCondition(conditions []string) (results []*EgProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("condition IN (?)", conditions).Find(&results).Error
	
	return
}
 
// GetFromShowPrice 通过show_price获取内容  
func (obj *_EgProductShopMgr) GetFromShowPrice(showPrice bool) (results []*EgProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("show_price = ?", showPrice).Find(&results).Error
	
	return
}

// GetBatchFromShowPrice 批量唯一主键查找 
func (obj *_EgProductShopMgr) GetBatchFromShowPrice(showPrices []bool) (results []*EgProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("show_price IN (?)", showPrices).Find(&results).Error
	
	return
}
 
// GetFromIndexed 通过indexed获取内容  
func (obj *_EgProductShopMgr) GetFromIndexed(indexed bool) (results []*EgProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("indexed = ?", indexed).Find(&results).Error
	
	return
}

// GetBatchFromIndexed 批量唯一主键查找 
func (obj *_EgProductShopMgr) GetBatchFromIndexed(indexeds []bool) (results []*EgProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("indexed IN (?)", indexeds).Find(&results).Error
	
	return
}
 
// GetFromVisibility 通过visibility获取内容  
func (obj *_EgProductShopMgr) GetFromVisibility(visibility string) (results []*EgProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("visibility = ?", visibility).Find(&results).Error
	
	return
}

// GetBatchFromVisibility 批量唯一主键查找 
func (obj *_EgProductShopMgr) GetBatchFromVisibility(visibilitys []string) (results []*EgProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("visibility IN (?)", visibilitys).Find(&results).Error
	
	return
}
 
// GetFromCacheDefaultAttribute 通过cache_default_attribute获取内容  
func (obj *_EgProductShopMgr) GetFromCacheDefaultAttribute(cacheDefaultAttribute uint32) (results []*EgProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cache_default_attribute = ?", cacheDefaultAttribute).Find(&results).Error
	
	return
}

// GetBatchFromCacheDefaultAttribute 批量唯一主键查找 
func (obj *_EgProductShopMgr) GetBatchFromCacheDefaultAttribute(cacheDefaultAttributes []uint32) (results []*EgProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cache_default_attribute IN (?)", cacheDefaultAttributes).Find(&results).Error
	
	return
}
 
// GetFromAdvancedStockManagement 通过advanced_stock_management获取内容  
func (obj *_EgProductShopMgr) GetFromAdvancedStockManagement(advancedStockManagement bool) (results []*EgProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("advanced_stock_management = ?", advancedStockManagement).Find(&results).Error
	
	return
}

// GetBatchFromAdvancedStockManagement 批量唯一主键查找 
func (obj *_EgProductShopMgr) GetBatchFromAdvancedStockManagement(advancedStockManagements []bool) (results []*EgProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("advanced_stock_management IN (?)", advancedStockManagements).Find(&results).Error
	
	return
}
 
// GetFromDateAdd 通过date_add获取内容  
func (obj *_EgProductShopMgr) GetFromDateAdd(dateAdd time.Time) (results []*EgProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error
	
	return
}

// GetBatchFromDateAdd 批量唯一主键查找 
func (obj *_EgProductShopMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*EgProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error
	
	return
}
 
// GetFromDateUpd 通过date_upd获取内容  
func (obj *_EgProductShopMgr) GetFromDateUpd(dateUpd time.Time) (results []*EgProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd = ?", dateUpd).Find(&results).Error
	
	return
}

// GetBatchFromDateUpd 批量唯一主键查找 
func (obj *_EgProductShopMgr) GetBatchFromDateUpd(dateUpds []time.Time) (results []*EgProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd IN (?)", dateUpds).Find(&results).Error
	
	return
}
 
// GetFromPackStockType 通过pack_stock_type获取内容  
func (obj *_EgProductShopMgr) GetFromPackStockType(packStockType uint32) (results []*EgProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("pack_stock_type = ?", packStockType).Find(&results).Error
	
	return
}

// GetBatchFromPackStockType 批量唯一主键查找 
func (obj *_EgProductShopMgr) GetBatchFromPackStockType(packStockTypes []uint32) (results []*EgProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("pack_stock_type IN (?)", packStockTypes).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgProductShopMgr) FetchByPrimaryKey(idProduct uint32 ,idShop uint32 ) (result EgProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ? AND id_shop = ?", idProduct , idShop).Find(&result).Error
	
	return
}
 

 
 // FetchIndexByIndexed  获取多个内容
 func (obj *_EgProductShopMgr) FetchIndexByIndexed(idProduct uint32 ,active bool ,indexed bool ) (results []*EgProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ? AND active = ? AND indexed = ?", idProduct , active , indexed).Find(&results).Error
	
	return
}
 
 // FetchIndexByIDCategoryDefault  获取多个内容
 func (obj *_EgProductShopMgr) FetchIndexByIDCategoryDefault(idCategoryDefault uint32 ) (results []*EgProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_category_default = ?", idCategoryDefault).Find(&results).Error
	
	return
}
 
 // FetchIndexByDateAdd  获取多个内容
 func (obj *_EgProductShopMgr) FetchIndexByDateAdd(active bool ,visibility string ,dateAdd time.Time ) (results []*EgProductShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active = ? AND visibility = ? AND date_add = ?", active , visibility , dateAdd).Find(&results).Error
	
	return
}
 

	

