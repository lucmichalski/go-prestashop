package	model	
import (	
"context"	
"gorm.io/gorm"	
"gorm.io/datatypes"	
"time"	
"fmt"	
)	

type _EgProductMgr struct {
	*_BaseMgr
}

// EgProductMgr open func
func EgProductMgr(db *gorm.DB) *_EgProductMgr {
	if db == nil {
		panic(fmt.Errorf("EgProductMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgProductMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_product"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgProductMgr) GetTableName() string {
	return "eg_product"
}

// Get 获取
func (obj *_EgProductMgr) Get() (result EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgProductMgr) Gets() (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDProduct id_product获取 
func (obj *_EgProductMgr) WithIDProduct(idProduct uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product"] = idProduct })
}

// WithIDSupplier id_supplier获取 
func (obj *_EgProductMgr) WithIDSupplier(idSupplier uint32) Option {
	return optionFunc(func(o *options) { o.query["id_supplier"] = idSupplier })
}

// WithIDManufacturer id_manufacturer获取 
func (obj *_EgProductMgr) WithIDManufacturer(idManufacturer uint32) Option {
	return optionFunc(func(o *options) { o.query["id_manufacturer"] = idManufacturer })
}

// WithIDCategoryDefault id_category_default获取 
func (obj *_EgProductMgr) WithIDCategoryDefault(idCategoryDefault uint32) Option {
	return optionFunc(func(o *options) { o.query["id_category_default"] = idCategoryDefault })
}

// WithIDShopDefault id_shop_default获取 
func (obj *_EgProductMgr) WithIDShopDefault(idShopDefault uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop_default"] = idShopDefault })
}

// WithIDTaxRulesGroup id_tax_rules_group获取 
func (obj *_EgProductMgr) WithIDTaxRulesGroup(idTaxRulesGroup uint32) Option {
	return optionFunc(func(o *options) { o.query["id_tax_rules_group"] = idTaxRulesGroup })
}

// WithOnSale on_sale获取 
func (obj *_EgProductMgr) WithOnSale(onSale bool) Option {
	return optionFunc(func(o *options) { o.query["on_sale"] = onSale })
}

// WithOnlineOnly online_only获取 
func (obj *_EgProductMgr) WithOnlineOnly(onlineOnly bool) Option {
	return optionFunc(func(o *options) { o.query["online_only"] = onlineOnly })
}

// WithEan13 ean13获取 
func (obj *_EgProductMgr) WithEan13(ean13 string) Option {
	return optionFunc(func(o *options) { o.query["ean13"] = ean13 })
}

// WithIsbn isbn获取 
func (obj *_EgProductMgr) WithIsbn(isbn string) Option {
	return optionFunc(func(o *options) { o.query["isbn"] = isbn })
}

// WithUpc upc获取 
func (obj *_EgProductMgr) WithUpc(upc string) Option {
	return optionFunc(func(o *options) { o.query["upc"] = upc })
}

// WithMpn mpn获取 
func (obj *_EgProductMgr) WithMpn(mpn string) Option {
	return optionFunc(func(o *options) { o.query["mpn"] = mpn })
}

// WithEcotax ecotax获取 
func (obj *_EgProductMgr) WithEcotax(ecotax float64) Option {
	return optionFunc(func(o *options) { o.query["ecotax"] = ecotax })
}

// WithQuantity quantity获取 
func (obj *_EgProductMgr) WithQuantity(quantity int) Option {
	return optionFunc(func(o *options) { o.query["quantity"] = quantity })
}

// WithMinimalQuantity minimal_quantity获取 
func (obj *_EgProductMgr) WithMinimalQuantity(minimalQuantity uint32) Option {
	return optionFunc(func(o *options) { o.query["minimal_quantity"] = minimalQuantity })
}

// WithLowStockThreshold low_stock_threshold获取 
func (obj *_EgProductMgr) WithLowStockThreshold(lowStockThreshold int) Option {
	return optionFunc(func(o *options) { o.query["low_stock_threshold"] = lowStockThreshold })
}

// WithLowStockAlert low_stock_alert获取 
func (obj *_EgProductMgr) WithLowStockAlert(lowStockAlert bool) Option {
	return optionFunc(func(o *options) { o.query["low_stock_alert"] = lowStockAlert })
}

// WithPrice price获取 
func (obj *_EgProductMgr) WithPrice(price float64) Option {
	return optionFunc(func(o *options) { o.query["price"] = price })
}

// WithWholesalePrice wholesale_price获取 
func (obj *_EgProductMgr) WithWholesalePrice(wholesalePrice float64) Option {
	return optionFunc(func(o *options) { o.query["wholesale_price"] = wholesalePrice })
}

// WithUnity unity获取 
func (obj *_EgProductMgr) WithUnity(unity string) Option {
	return optionFunc(func(o *options) { o.query["unity"] = unity })
}

// WithUnitPriceRatio unit_price_ratio获取 
func (obj *_EgProductMgr) WithUnitPriceRatio(unitPriceRatio float64) Option {
	return optionFunc(func(o *options) { o.query["unit_price_ratio"] = unitPriceRatio })
}

// WithAdditionalShippingCost additional_shipping_cost获取 
func (obj *_EgProductMgr) WithAdditionalShippingCost(additionalShippingCost float64) Option {
	return optionFunc(func(o *options) { o.query["additional_shipping_cost"] = additionalShippingCost })
}

// WithReference reference获取 
func (obj *_EgProductMgr) WithReference(reference string) Option {
	return optionFunc(func(o *options) { o.query["reference"] = reference })
}

// WithSupplierReference supplier_reference获取 
func (obj *_EgProductMgr) WithSupplierReference(supplierReference string) Option {
	return optionFunc(func(o *options) { o.query["supplier_reference"] = supplierReference })
}

// WithLocation location获取 
func (obj *_EgProductMgr) WithLocation(location string) Option {
	return optionFunc(func(o *options) { o.query["location"] = location })
}

// WithWidth width获取 
func (obj *_EgProductMgr) WithWidth(width float64) Option {
	return optionFunc(func(o *options) { o.query["width"] = width })
}

// WithHeight height获取 
func (obj *_EgProductMgr) WithHeight(height float64) Option {
	return optionFunc(func(o *options) { o.query["height"] = height })
}

// WithDepth depth获取 
func (obj *_EgProductMgr) WithDepth(depth float64) Option {
	return optionFunc(func(o *options) { o.query["depth"] = depth })
}

// WithWeight weight获取 
func (obj *_EgProductMgr) WithWeight(weight float64) Option {
	return optionFunc(func(o *options) { o.query["weight"] = weight })
}

// WithOutOfStock out_of_stock获取 
func (obj *_EgProductMgr) WithOutOfStock(outOfStock uint32) Option {
	return optionFunc(func(o *options) { o.query["out_of_stock"] = outOfStock })
}

// WithAdditionalDeliveryTimes additional_delivery_times获取 
func (obj *_EgProductMgr) WithAdditionalDeliveryTimes(additionalDeliveryTimes bool) Option {
	return optionFunc(func(o *options) { o.query["additional_delivery_times"] = additionalDeliveryTimes })
}

// WithQuantityDiscount quantity_discount获取 
func (obj *_EgProductMgr) WithQuantityDiscount(quantityDiscount bool) Option {
	return optionFunc(func(o *options) { o.query["quantity_discount"] = quantityDiscount })
}

// WithCustomizable customizable获取 
func (obj *_EgProductMgr) WithCustomizable(customizable int8) Option {
	return optionFunc(func(o *options) { o.query["customizable"] = customizable })
}

// WithUploadableFiles uploadable_files获取 
func (obj *_EgProductMgr) WithUploadableFiles(uploadableFiles int8) Option {
	return optionFunc(func(o *options) { o.query["uploadable_files"] = uploadableFiles })
}

// WithTextFields text_fields获取 
func (obj *_EgProductMgr) WithTextFields(textFields int8) Option {
	return optionFunc(func(o *options) { o.query["text_fields"] = textFields })
}

// WithActive active获取 
func (obj *_EgProductMgr) WithActive(active bool) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}

// WithRedirectType redirect_type获取 
func (obj *_EgProductMgr) WithRedirectType(redirectType string) Option {
	return optionFunc(func(o *options) { o.query["redirect_type"] = redirectType })
}

// WithIDTypeRedirected id_type_redirected获取 
func (obj *_EgProductMgr) WithIDTypeRedirected(idTypeRedirected uint32) Option {
	return optionFunc(func(o *options) { o.query["id_type_redirected"] = idTypeRedirected })
}

// WithAvailableForOrder available_for_order获取 
func (obj *_EgProductMgr) WithAvailableForOrder(availableForOrder bool) Option {
	return optionFunc(func(o *options) { o.query["available_for_order"] = availableForOrder })
}

// WithAvailableDate available_date获取 
func (obj *_EgProductMgr) WithAvailableDate(availableDate datatypes.Date) Option {
	return optionFunc(func(o *options) { o.query["available_date"] = availableDate })
}

// WithShowCondition show_condition获取 
func (obj *_EgProductMgr) WithShowCondition(showCondition bool) Option {
	return optionFunc(func(o *options) { o.query["show_condition"] = showCondition })
}

// WithCondition condition获取 
func (obj *_EgProductMgr) WithCondition(condition string) Option {
	return optionFunc(func(o *options) { o.query["condition"] = condition })
}

// WithShowPrice show_price获取 
func (obj *_EgProductMgr) WithShowPrice(showPrice bool) Option {
	return optionFunc(func(o *options) { o.query["show_price"] = showPrice })
}

// WithIndexed indexed获取 
func (obj *_EgProductMgr) WithIndexed(indexed bool) Option {
	return optionFunc(func(o *options) { o.query["indexed"] = indexed })
}

// WithVisibility visibility获取 
func (obj *_EgProductMgr) WithVisibility(visibility string) Option {
	return optionFunc(func(o *options) { o.query["visibility"] = visibility })
}

// WithCacheIsPack cache_is_pack获取 
func (obj *_EgProductMgr) WithCacheIsPack(cacheIsPack bool) Option {
	return optionFunc(func(o *options) { o.query["cache_is_pack"] = cacheIsPack })
}

// WithCacheHasAttachments cache_has_attachments获取 
func (obj *_EgProductMgr) WithCacheHasAttachments(cacheHasAttachments bool) Option {
	return optionFunc(func(o *options) { o.query["cache_has_attachments"] = cacheHasAttachments })
}

// WithIsVirtual is_virtual获取 
func (obj *_EgProductMgr) WithIsVirtual(isVirtual bool) Option {
	return optionFunc(func(o *options) { o.query["is_virtual"] = isVirtual })
}

// WithCacheDefaultAttribute cache_default_attribute获取 
func (obj *_EgProductMgr) WithCacheDefaultAttribute(cacheDefaultAttribute uint32) Option {
	return optionFunc(func(o *options) { o.query["cache_default_attribute"] = cacheDefaultAttribute })
}

// WithDateAdd date_add获取 
func (obj *_EgProductMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

// WithDateUpd date_upd获取 
func (obj *_EgProductMgr) WithDateUpd(dateUpd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_upd"] = dateUpd })
}

// WithAdvancedStockManagement advanced_stock_management获取 
func (obj *_EgProductMgr) WithAdvancedStockManagement(advancedStockManagement bool) Option {
	return optionFunc(func(o *options) { o.query["advanced_stock_management"] = advancedStockManagement })
}

// WithPackStockType pack_stock_type获取 
func (obj *_EgProductMgr) WithPackStockType(packStockType uint32) Option {
	return optionFunc(func(o *options) { o.query["pack_stock_type"] = packStockType })
}

// WithState state获取 
func (obj *_EgProductMgr) WithState(state uint32) Option {
	return optionFunc(func(o *options) { o.query["state"] = state })
}


// GetByOption 功能选项模式获取
func (obj *_EgProductMgr) GetByOption(opts ...Option) (result EgProduct, err error) {
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
func (obj *_EgProductMgr) GetByOptions(opts ...Option) (results []*EgProduct, err error) {
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
func (obj *_EgProductMgr)  GetFromIDProduct(idProduct uint32) (result EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ?", idProduct).Find(&result).Error
	
	return
}

// GetBatchFromIDProduct 批量唯一主键查找 
func (obj *_EgProductMgr) GetBatchFromIDProduct(idProducts []uint32) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product IN (?)", idProducts).Find(&results).Error
	
	return
}
 
// GetFromIDSupplier 通过id_supplier获取内容  
func (obj *_EgProductMgr) GetFromIDSupplier(idSupplier uint32) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supplier = ?", idSupplier).Find(&results).Error
	
	return
}

// GetBatchFromIDSupplier 批量唯一主键查找 
func (obj *_EgProductMgr) GetBatchFromIDSupplier(idSuppliers []uint32) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supplier IN (?)", idSuppliers).Find(&results).Error
	
	return
}
 
// GetFromIDManufacturer 通过id_manufacturer获取内容  
func (obj *_EgProductMgr) GetFromIDManufacturer(idManufacturer uint32) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_manufacturer = ?", idManufacturer).Find(&results).Error
	
	return
}

// GetBatchFromIDManufacturer 批量唯一主键查找 
func (obj *_EgProductMgr) GetBatchFromIDManufacturer(idManufacturers []uint32) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_manufacturer IN (?)", idManufacturers).Find(&results).Error
	
	return
}
 
// GetFromIDCategoryDefault 通过id_category_default获取内容  
func (obj *_EgProductMgr) GetFromIDCategoryDefault(idCategoryDefault uint32) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_category_default = ?", idCategoryDefault).Find(&results).Error
	
	return
}

// GetBatchFromIDCategoryDefault 批量唯一主键查找 
func (obj *_EgProductMgr) GetBatchFromIDCategoryDefault(idCategoryDefaults []uint32) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_category_default IN (?)", idCategoryDefaults).Find(&results).Error
	
	return
}
 
// GetFromIDShopDefault 通过id_shop_default获取内容  
func (obj *_EgProductMgr) GetFromIDShopDefault(idShopDefault uint32) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_default = ?", idShopDefault).Find(&results).Error
	
	return
}

// GetBatchFromIDShopDefault 批量唯一主键查找 
func (obj *_EgProductMgr) GetBatchFromIDShopDefault(idShopDefaults []uint32) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_default IN (?)", idShopDefaults).Find(&results).Error
	
	return
}
 
// GetFromIDTaxRulesGroup 通过id_tax_rules_group获取内容  
func (obj *_EgProductMgr) GetFromIDTaxRulesGroup(idTaxRulesGroup uint32) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tax_rules_group = ?", idTaxRulesGroup).Find(&results).Error
	
	return
}

// GetBatchFromIDTaxRulesGroup 批量唯一主键查找 
func (obj *_EgProductMgr) GetBatchFromIDTaxRulesGroup(idTaxRulesGroups []uint32) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tax_rules_group IN (?)", idTaxRulesGroups).Find(&results).Error
	
	return
}
 
// GetFromOnSale 通过on_sale获取内容  
func (obj *_EgProductMgr) GetFromOnSale(onSale bool) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("on_sale = ?", onSale).Find(&results).Error
	
	return
}

// GetBatchFromOnSale 批量唯一主键查找 
func (obj *_EgProductMgr) GetBatchFromOnSale(onSales []bool) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("on_sale IN (?)", onSales).Find(&results).Error
	
	return
}
 
// GetFromOnlineOnly 通过online_only获取内容  
func (obj *_EgProductMgr) GetFromOnlineOnly(onlineOnly bool) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("online_only = ?", onlineOnly).Find(&results).Error
	
	return
}

// GetBatchFromOnlineOnly 批量唯一主键查找 
func (obj *_EgProductMgr) GetBatchFromOnlineOnly(onlineOnlys []bool) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("online_only IN (?)", onlineOnlys).Find(&results).Error
	
	return
}
 
// GetFromEan13 通过ean13获取内容  
func (obj *_EgProductMgr) GetFromEan13(ean13 string) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ean13 = ?", ean13).Find(&results).Error
	
	return
}

// GetBatchFromEan13 批量唯一主键查找 
func (obj *_EgProductMgr) GetBatchFromEan13(ean13s []string) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ean13 IN (?)", ean13s).Find(&results).Error
	
	return
}
 
// GetFromIsbn 通过isbn获取内容  
func (obj *_EgProductMgr) GetFromIsbn(isbn string) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("isbn = ?", isbn).Find(&results).Error
	
	return
}

// GetBatchFromIsbn 批量唯一主键查找 
func (obj *_EgProductMgr) GetBatchFromIsbn(isbns []string) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("isbn IN (?)", isbns).Find(&results).Error
	
	return
}
 
// GetFromUpc 通过upc获取内容  
func (obj *_EgProductMgr) GetFromUpc(upc string) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("upc = ?", upc).Find(&results).Error
	
	return
}

// GetBatchFromUpc 批量唯一主键查找 
func (obj *_EgProductMgr) GetBatchFromUpc(upcs []string) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("upc IN (?)", upcs).Find(&results).Error
	
	return
}
 
// GetFromMpn 通过mpn获取内容  
func (obj *_EgProductMgr) GetFromMpn(mpn string) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("mpn = ?", mpn).Find(&results).Error
	
	return
}

// GetBatchFromMpn 批量唯一主键查找 
func (obj *_EgProductMgr) GetBatchFromMpn(mpns []string) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("mpn IN (?)", mpns).Find(&results).Error
	
	return
}
 
// GetFromEcotax 通过ecotax获取内容  
func (obj *_EgProductMgr) GetFromEcotax(ecotax float64) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ecotax = ?", ecotax).Find(&results).Error
	
	return
}

// GetBatchFromEcotax 批量唯一主键查找 
func (obj *_EgProductMgr) GetBatchFromEcotax(ecotaxs []float64) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ecotax IN (?)", ecotaxs).Find(&results).Error
	
	return
}
 
// GetFromQuantity 通过quantity获取内容  
func (obj *_EgProductMgr) GetFromQuantity(quantity int) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("quantity = ?", quantity).Find(&results).Error
	
	return
}

// GetBatchFromQuantity 批量唯一主键查找 
func (obj *_EgProductMgr) GetBatchFromQuantity(quantitys []int) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("quantity IN (?)", quantitys).Find(&results).Error
	
	return
}
 
// GetFromMinimalQuantity 通过minimal_quantity获取内容  
func (obj *_EgProductMgr) GetFromMinimalQuantity(minimalQuantity uint32) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("minimal_quantity = ?", minimalQuantity).Find(&results).Error
	
	return
}

// GetBatchFromMinimalQuantity 批量唯一主键查找 
func (obj *_EgProductMgr) GetBatchFromMinimalQuantity(minimalQuantitys []uint32) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("minimal_quantity IN (?)", minimalQuantitys).Find(&results).Error
	
	return
}
 
// GetFromLowStockThreshold 通过low_stock_threshold获取内容  
func (obj *_EgProductMgr) GetFromLowStockThreshold(lowStockThreshold int) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("low_stock_threshold = ?", lowStockThreshold).Find(&results).Error
	
	return
}

// GetBatchFromLowStockThreshold 批量唯一主键查找 
func (obj *_EgProductMgr) GetBatchFromLowStockThreshold(lowStockThresholds []int) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("low_stock_threshold IN (?)", lowStockThresholds).Find(&results).Error
	
	return
}
 
// GetFromLowStockAlert 通过low_stock_alert获取内容  
func (obj *_EgProductMgr) GetFromLowStockAlert(lowStockAlert bool) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("low_stock_alert = ?", lowStockAlert).Find(&results).Error
	
	return
}

// GetBatchFromLowStockAlert 批量唯一主键查找 
func (obj *_EgProductMgr) GetBatchFromLowStockAlert(lowStockAlerts []bool) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("low_stock_alert IN (?)", lowStockAlerts).Find(&results).Error
	
	return
}
 
// GetFromPrice 通过price获取内容  
func (obj *_EgProductMgr) GetFromPrice(price float64) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("price = ?", price).Find(&results).Error
	
	return
}

// GetBatchFromPrice 批量唯一主键查找 
func (obj *_EgProductMgr) GetBatchFromPrice(prices []float64) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("price IN (?)", prices).Find(&results).Error
	
	return
}
 
// GetFromWholesalePrice 通过wholesale_price获取内容  
func (obj *_EgProductMgr) GetFromWholesalePrice(wholesalePrice float64) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("wholesale_price = ?", wholesalePrice).Find(&results).Error
	
	return
}

// GetBatchFromWholesalePrice 批量唯一主键查找 
func (obj *_EgProductMgr) GetBatchFromWholesalePrice(wholesalePrices []float64) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("wholesale_price IN (?)", wholesalePrices).Find(&results).Error
	
	return
}
 
// GetFromUnity 通过unity获取内容  
func (obj *_EgProductMgr) GetFromUnity(unity string) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("unity = ?", unity).Find(&results).Error
	
	return
}

// GetBatchFromUnity 批量唯一主键查找 
func (obj *_EgProductMgr) GetBatchFromUnity(unitys []string) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("unity IN (?)", unitys).Find(&results).Error
	
	return
}
 
// GetFromUnitPriceRatio 通过unit_price_ratio获取内容  
func (obj *_EgProductMgr) GetFromUnitPriceRatio(unitPriceRatio float64) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("unit_price_ratio = ?", unitPriceRatio).Find(&results).Error
	
	return
}

// GetBatchFromUnitPriceRatio 批量唯一主键查找 
func (obj *_EgProductMgr) GetBatchFromUnitPriceRatio(unitPriceRatios []float64) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("unit_price_ratio IN (?)", unitPriceRatios).Find(&results).Error
	
	return
}
 
// GetFromAdditionalShippingCost 通过additional_shipping_cost获取内容  
func (obj *_EgProductMgr) GetFromAdditionalShippingCost(additionalShippingCost float64) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("additional_shipping_cost = ?", additionalShippingCost).Find(&results).Error
	
	return
}

// GetBatchFromAdditionalShippingCost 批量唯一主键查找 
func (obj *_EgProductMgr) GetBatchFromAdditionalShippingCost(additionalShippingCosts []float64) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("additional_shipping_cost IN (?)", additionalShippingCosts).Find(&results).Error
	
	return
}
 
// GetFromReference 通过reference获取内容  
func (obj *_EgProductMgr) GetFromReference(reference string) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reference = ?", reference).Find(&results).Error
	
	return
}

// GetBatchFromReference 批量唯一主键查找 
func (obj *_EgProductMgr) GetBatchFromReference(references []string) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reference IN (?)", references).Find(&results).Error
	
	return
}
 
// GetFromSupplierReference 通过supplier_reference获取内容  
func (obj *_EgProductMgr) GetFromSupplierReference(supplierReference string) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("supplier_reference = ?", supplierReference).Find(&results).Error
	
	return
}

// GetBatchFromSupplierReference 批量唯一主键查找 
func (obj *_EgProductMgr) GetBatchFromSupplierReference(supplierReferences []string) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("supplier_reference IN (?)", supplierReferences).Find(&results).Error
	
	return
}
 
// GetFromLocation 通过location获取内容  
func (obj *_EgProductMgr) GetFromLocation(location string) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("location = ?", location).Find(&results).Error
	
	return
}

// GetBatchFromLocation 批量唯一主键查找 
func (obj *_EgProductMgr) GetBatchFromLocation(locations []string) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("location IN (?)", locations).Find(&results).Error
	
	return
}
 
// GetFromWidth 通过width获取内容  
func (obj *_EgProductMgr) GetFromWidth(width float64) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("width = ?", width).Find(&results).Error
	
	return
}

// GetBatchFromWidth 批量唯一主键查找 
func (obj *_EgProductMgr) GetBatchFromWidth(widths []float64) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("width IN (?)", widths).Find(&results).Error
	
	return
}
 
// GetFromHeight 通过height获取内容  
func (obj *_EgProductMgr) GetFromHeight(height float64) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("height = ?", height).Find(&results).Error
	
	return
}

// GetBatchFromHeight 批量唯一主键查找 
func (obj *_EgProductMgr) GetBatchFromHeight(heights []float64) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("height IN (?)", heights).Find(&results).Error
	
	return
}
 
// GetFromDepth 通过depth获取内容  
func (obj *_EgProductMgr) GetFromDepth(depth float64) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("depth = ?", depth).Find(&results).Error
	
	return
}

// GetBatchFromDepth 批量唯一主键查找 
func (obj *_EgProductMgr) GetBatchFromDepth(depths []float64) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("depth IN (?)", depths).Find(&results).Error
	
	return
}
 
// GetFromWeight 通过weight获取内容  
func (obj *_EgProductMgr) GetFromWeight(weight float64) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("weight = ?", weight).Find(&results).Error
	
	return
}

// GetBatchFromWeight 批量唯一主键查找 
func (obj *_EgProductMgr) GetBatchFromWeight(weights []float64) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("weight IN (?)", weights).Find(&results).Error
	
	return
}
 
// GetFromOutOfStock 通过out_of_stock获取内容  
func (obj *_EgProductMgr) GetFromOutOfStock(outOfStock uint32) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("out_of_stock = ?", outOfStock).Find(&results).Error
	
	return
}

// GetBatchFromOutOfStock 批量唯一主键查找 
func (obj *_EgProductMgr) GetBatchFromOutOfStock(outOfStocks []uint32) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("out_of_stock IN (?)", outOfStocks).Find(&results).Error
	
	return
}
 
// GetFromAdditionalDeliveryTimes 通过additional_delivery_times获取内容  
func (obj *_EgProductMgr) GetFromAdditionalDeliveryTimes(additionalDeliveryTimes bool) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("additional_delivery_times = ?", additionalDeliveryTimes).Find(&results).Error
	
	return
}

// GetBatchFromAdditionalDeliveryTimes 批量唯一主键查找 
func (obj *_EgProductMgr) GetBatchFromAdditionalDeliveryTimes(additionalDeliveryTimess []bool) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("additional_delivery_times IN (?)", additionalDeliveryTimess).Find(&results).Error
	
	return
}
 
// GetFromQuantityDiscount 通过quantity_discount获取内容  
func (obj *_EgProductMgr) GetFromQuantityDiscount(quantityDiscount bool) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("quantity_discount = ?", quantityDiscount).Find(&results).Error
	
	return
}

// GetBatchFromQuantityDiscount 批量唯一主键查找 
func (obj *_EgProductMgr) GetBatchFromQuantityDiscount(quantityDiscounts []bool) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("quantity_discount IN (?)", quantityDiscounts).Find(&results).Error
	
	return
}
 
// GetFromCustomizable 通过customizable获取内容  
func (obj *_EgProductMgr) GetFromCustomizable(customizable int8) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("customizable = ?", customizable).Find(&results).Error
	
	return
}

// GetBatchFromCustomizable 批量唯一主键查找 
func (obj *_EgProductMgr) GetBatchFromCustomizable(customizables []int8) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("customizable IN (?)", customizables).Find(&results).Error
	
	return
}
 
// GetFromUploadableFiles 通过uploadable_files获取内容  
func (obj *_EgProductMgr) GetFromUploadableFiles(uploadableFiles int8) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("uploadable_files = ?", uploadableFiles).Find(&results).Error
	
	return
}

// GetBatchFromUploadableFiles 批量唯一主键查找 
func (obj *_EgProductMgr) GetBatchFromUploadableFiles(uploadableFiless []int8) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("uploadable_files IN (?)", uploadableFiless).Find(&results).Error
	
	return
}
 
// GetFromTextFields 通过text_fields获取内容  
func (obj *_EgProductMgr) GetFromTextFields(textFields int8) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("text_fields = ?", textFields).Find(&results).Error
	
	return
}

// GetBatchFromTextFields 批量唯一主键查找 
func (obj *_EgProductMgr) GetBatchFromTextFields(textFieldss []int8) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("text_fields IN (?)", textFieldss).Find(&results).Error
	
	return
}
 
// GetFromActive 通过active获取内容  
func (obj *_EgProductMgr) GetFromActive(active bool) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active = ?", active).Find(&results).Error
	
	return
}

// GetBatchFromActive 批量唯一主键查找 
func (obj *_EgProductMgr) GetBatchFromActive(actives []bool) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active IN (?)", actives).Find(&results).Error
	
	return
}
 
// GetFromRedirectType 通过redirect_type获取内容  
func (obj *_EgProductMgr) GetFromRedirectType(redirectType string) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("redirect_type = ?", redirectType).Find(&results).Error
	
	return
}

// GetBatchFromRedirectType 批量唯一主键查找 
func (obj *_EgProductMgr) GetBatchFromRedirectType(redirectTypes []string) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("redirect_type IN (?)", redirectTypes).Find(&results).Error
	
	return
}
 
// GetFromIDTypeRedirected 通过id_type_redirected获取内容  
func (obj *_EgProductMgr) GetFromIDTypeRedirected(idTypeRedirected uint32) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_type_redirected = ?", idTypeRedirected).Find(&results).Error
	
	return
}

// GetBatchFromIDTypeRedirected 批量唯一主键查找 
func (obj *_EgProductMgr) GetBatchFromIDTypeRedirected(idTypeRedirecteds []uint32) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_type_redirected IN (?)", idTypeRedirecteds).Find(&results).Error
	
	return
}
 
// GetFromAvailableForOrder 通过available_for_order获取内容  
func (obj *_EgProductMgr) GetFromAvailableForOrder(availableForOrder bool) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("available_for_order = ?", availableForOrder).Find(&results).Error
	
	return
}

// GetBatchFromAvailableForOrder 批量唯一主键查找 
func (obj *_EgProductMgr) GetBatchFromAvailableForOrder(availableForOrders []bool) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("available_for_order IN (?)", availableForOrders).Find(&results).Error
	
	return
}
 
// GetFromAvailableDate 通过available_date获取内容  
func (obj *_EgProductMgr) GetFromAvailableDate(availableDate datatypes.Date) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("available_date = ?", availableDate).Find(&results).Error
	
	return
}

// GetBatchFromAvailableDate 批量唯一主键查找 
func (obj *_EgProductMgr) GetBatchFromAvailableDate(availableDates []datatypes.Date) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("available_date IN (?)", availableDates).Find(&results).Error
	
	return
}
 
// GetFromShowCondition 通过show_condition获取内容  
func (obj *_EgProductMgr) GetFromShowCondition(showCondition bool) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("show_condition = ?", showCondition).Find(&results).Error
	
	return
}

// GetBatchFromShowCondition 批量唯一主键查找 
func (obj *_EgProductMgr) GetBatchFromShowCondition(showConditions []bool) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("show_condition IN (?)", showConditions).Find(&results).Error
	
	return
}
 
// GetFromCondition 通过condition获取内容  
func (obj *_EgProductMgr) GetFromCondition(condition string) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("condition = ?", condition).Find(&results).Error
	
	return
}

// GetBatchFromCondition 批量唯一主键查找 
func (obj *_EgProductMgr) GetBatchFromCondition(conditions []string) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("condition IN (?)", conditions).Find(&results).Error
	
	return
}
 
// GetFromShowPrice 通过show_price获取内容  
func (obj *_EgProductMgr) GetFromShowPrice(showPrice bool) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("show_price = ?", showPrice).Find(&results).Error
	
	return
}

// GetBatchFromShowPrice 批量唯一主键查找 
func (obj *_EgProductMgr) GetBatchFromShowPrice(showPrices []bool) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("show_price IN (?)", showPrices).Find(&results).Error
	
	return
}
 
// GetFromIndexed 通过indexed获取内容  
func (obj *_EgProductMgr) GetFromIndexed(indexed bool) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("indexed = ?", indexed).Find(&results).Error
	
	return
}

// GetBatchFromIndexed 批量唯一主键查找 
func (obj *_EgProductMgr) GetBatchFromIndexed(indexeds []bool) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("indexed IN (?)", indexeds).Find(&results).Error
	
	return
}
 
// GetFromVisibility 通过visibility获取内容  
func (obj *_EgProductMgr) GetFromVisibility(visibility string) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("visibility = ?", visibility).Find(&results).Error
	
	return
}

// GetBatchFromVisibility 批量唯一主键查找 
func (obj *_EgProductMgr) GetBatchFromVisibility(visibilitys []string) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("visibility IN (?)", visibilitys).Find(&results).Error
	
	return
}
 
// GetFromCacheIsPack 通过cache_is_pack获取内容  
func (obj *_EgProductMgr) GetFromCacheIsPack(cacheIsPack bool) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cache_is_pack = ?", cacheIsPack).Find(&results).Error
	
	return
}

// GetBatchFromCacheIsPack 批量唯一主键查找 
func (obj *_EgProductMgr) GetBatchFromCacheIsPack(cacheIsPacks []bool) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cache_is_pack IN (?)", cacheIsPacks).Find(&results).Error
	
	return
}
 
// GetFromCacheHasAttachments 通过cache_has_attachments获取内容  
func (obj *_EgProductMgr) GetFromCacheHasAttachments(cacheHasAttachments bool) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cache_has_attachments = ?", cacheHasAttachments).Find(&results).Error
	
	return
}

// GetBatchFromCacheHasAttachments 批量唯一主键查找 
func (obj *_EgProductMgr) GetBatchFromCacheHasAttachments(cacheHasAttachmentss []bool) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cache_has_attachments IN (?)", cacheHasAttachmentss).Find(&results).Error
	
	return
}
 
// GetFromIsVirtual 通过is_virtual获取内容  
func (obj *_EgProductMgr) GetFromIsVirtual(isVirtual bool) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_virtual = ?", isVirtual).Find(&results).Error
	
	return
}

// GetBatchFromIsVirtual 批量唯一主键查找 
func (obj *_EgProductMgr) GetBatchFromIsVirtual(isVirtuals []bool) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_virtual IN (?)", isVirtuals).Find(&results).Error
	
	return
}
 
// GetFromCacheDefaultAttribute 通过cache_default_attribute获取内容  
func (obj *_EgProductMgr) GetFromCacheDefaultAttribute(cacheDefaultAttribute uint32) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cache_default_attribute = ?", cacheDefaultAttribute).Find(&results).Error
	
	return
}

// GetBatchFromCacheDefaultAttribute 批量唯一主键查找 
func (obj *_EgProductMgr) GetBatchFromCacheDefaultAttribute(cacheDefaultAttributes []uint32) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cache_default_attribute IN (?)", cacheDefaultAttributes).Find(&results).Error
	
	return
}
 
// GetFromDateAdd 通过date_add获取内容  
func (obj *_EgProductMgr) GetFromDateAdd(dateAdd time.Time) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error
	
	return
}

// GetBatchFromDateAdd 批量唯一主键查找 
func (obj *_EgProductMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error
	
	return
}
 
// GetFromDateUpd 通过date_upd获取内容  
func (obj *_EgProductMgr) GetFromDateUpd(dateUpd time.Time) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd = ?", dateUpd).Find(&results).Error
	
	return
}

// GetBatchFromDateUpd 批量唯一主键查找 
func (obj *_EgProductMgr) GetBatchFromDateUpd(dateUpds []time.Time) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd IN (?)", dateUpds).Find(&results).Error
	
	return
}
 
// GetFromAdvancedStockManagement 通过advanced_stock_management获取内容  
func (obj *_EgProductMgr) GetFromAdvancedStockManagement(advancedStockManagement bool) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("advanced_stock_management = ?", advancedStockManagement).Find(&results).Error
	
	return
}

// GetBatchFromAdvancedStockManagement 批量唯一主键查找 
func (obj *_EgProductMgr) GetBatchFromAdvancedStockManagement(advancedStockManagements []bool) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("advanced_stock_management IN (?)", advancedStockManagements).Find(&results).Error
	
	return
}
 
// GetFromPackStockType 通过pack_stock_type获取内容  
func (obj *_EgProductMgr) GetFromPackStockType(packStockType uint32) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("pack_stock_type = ?", packStockType).Find(&results).Error
	
	return
}

// GetBatchFromPackStockType 批量唯一主键查找 
func (obj *_EgProductMgr) GetBatchFromPackStockType(packStockTypes []uint32) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("pack_stock_type IN (?)", packStockTypes).Find(&results).Error
	
	return
}
 
// GetFromState 通过state获取内容  
func (obj *_EgProductMgr) GetFromState(state uint32) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("state = ?", state).Find(&results).Error
	
	return
}

// GetBatchFromState 批量唯一主键查找 
func (obj *_EgProductMgr) GetBatchFromState(states []uint32) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("state IN (?)", states).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgProductMgr) FetchByPrimaryKey(idProduct uint32 ) (result EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ?", idProduct).Find(&result).Error
	
	return
}
 

 
 // FetchIndexByProductManufacturer  获取多个内容
 func (obj *_EgProductMgr) FetchIndexByProductManufacturer(idProduct uint32 ,idManufacturer uint32 ) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ? AND id_manufacturer = ?", idProduct , idManufacturer).Find(&results).Error
	
	return
}
 
 // FetchIndexByProductSupplier  获取多个内容
 func (obj *_EgProductMgr) FetchIndexByProductSupplier(idSupplier uint32 ) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supplier = ?", idSupplier).Find(&results).Error
	
	return
}
 
 // FetchIndexByIDCategoryDefault  获取多个内容
 func (obj *_EgProductMgr) FetchIndexByIDCategoryDefault(idCategoryDefault uint32 ) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_category_default = ?", idCategoryDefault).Find(&results).Error
	
	return
}
 
 // FetchIndexByReferenceIDx  获取多个内容
 func (obj *_EgProductMgr) FetchIndexByReferenceIDx(reference string ) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reference = ?", reference).Find(&results).Error
	
	return
}
 
 // FetchIndexBySupplierReferenceIDx  获取多个内容
 func (obj *_EgProductMgr) FetchIndexBySupplierReferenceIDx(supplierReference string ) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("supplier_reference = ?", supplierReference).Find(&results).Error
	
	return
}
 
 // FetchIndexByIndexed  获取多个内容
 func (obj *_EgProductMgr) FetchIndexByIndexed(indexed bool ) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("indexed = ?", indexed).Find(&results).Error
	
	return
}
 
 // FetchIndexByDateAdd  获取多个内容
 func (obj *_EgProductMgr) FetchIndexByDateAdd(dateAdd time.Time ) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error
	
	return
}
 
 // FetchIndexByState  获取多个内容
 func (obj *_EgProductMgr) FetchIndexByState(dateUpd time.Time ,state uint32 ) (results []*EgProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd = ? AND state = ?", dateUpd , state).Find(&results).Error
	
	return
}
 

	

