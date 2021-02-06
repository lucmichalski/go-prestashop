package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
"gorm.io/datatypes"	
)	

type _EgProductAttributeMgr struct {
	*_BaseMgr
}

// EgProductAttributeMgr open func
func EgProductAttributeMgr(db *gorm.DB) *_EgProductAttributeMgr {
	if db == nil {
		panic(fmt.Errorf("EgProductAttributeMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgProductAttributeMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_product_attribute"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgProductAttributeMgr) GetTableName() string {
	return "eg_product_attribute"
}

// Get 获取
func (obj *_EgProductAttributeMgr) Get() (result EgProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgProductAttributeMgr) Gets() (results []*EgProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDProductAttribute id_product_attribute获取 
func (obj *_EgProductAttributeMgr) WithIDProductAttribute(idProductAttribute uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product_attribute"] = idProductAttribute })
}

// WithIDProduct id_product获取 
func (obj *_EgProductAttributeMgr) WithIDProduct(idProduct uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product"] = idProduct })
}

// WithReference reference获取 
func (obj *_EgProductAttributeMgr) WithReference(reference string) Option {
	return optionFunc(func(o *options) { o.query["reference"] = reference })
}

// WithSupplierReference supplier_reference获取 
func (obj *_EgProductAttributeMgr) WithSupplierReference(supplierReference string) Option {
	return optionFunc(func(o *options) { o.query["supplier_reference"] = supplierReference })
}

// WithLocation location获取 
func (obj *_EgProductAttributeMgr) WithLocation(location string) Option {
	return optionFunc(func(o *options) { o.query["location"] = location })
}

// WithEan13 ean13获取 
func (obj *_EgProductAttributeMgr) WithEan13(ean13 string) Option {
	return optionFunc(func(o *options) { o.query["ean13"] = ean13 })
}

// WithIsbn isbn获取 
func (obj *_EgProductAttributeMgr) WithIsbn(isbn string) Option {
	return optionFunc(func(o *options) { o.query["isbn"] = isbn })
}

// WithUpc upc获取 
func (obj *_EgProductAttributeMgr) WithUpc(upc string) Option {
	return optionFunc(func(o *options) { o.query["upc"] = upc })
}

// WithMpn mpn获取 
func (obj *_EgProductAttributeMgr) WithMpn(mpn string) Option {
	return optionFunc(func(o *options) { o.query["mpn"] = mpn })
}

// WithWholesalePrice wholesale_price获取 
func (obj *_EgProductAttributeMgr) WithWholesalePrice(wholesalePrice float64) Option {
	return optionFunc(func(o *options) { o.query["wholesale_price"] = wholesalePrice })
}

// WithPrice price获取 
func (obj *_EgProductAttributeMgr) WithPrice(price float64) Option {
	return optionFunc(func(o *options) { o.query["price"] = price })
}

// WithEcotax ecotax获取 
func (obj *_EgProductAttributeMgr) WithEcotax(ecotax float64) Option {
	return optionFunc(func(o *options) { o.query["ecotax"] = ecotax })
}

// WithQuantity quantity获取 
func (obj *_EgProductAttributeMgr) WithQuantity(quantity int) Option {
	return optionFunc(func(o *options) { o.query["quantity"] = quantity })
}

// WithWeight weight获取 
func (obj *_EgProductAttributeMgr) WithWeight(weight float64) Option {
	return optionFunc(func(o *options) { o.query["weight"] = weight })
}

// WithUnitPriceImpact unit_price_impact获取 
func (obj *_EgProductAttributeMgr) WithUnitPriceImpact(unitPriceImpact float64) Option {
	return optionFunc(func(o *options) { o.query["unit_price_impact"] = unitPriceImpact })
}

// WithDefaultOn default_on获取 
func (obj *_EgProductAttributeMgr) WithDefaultOn(defaultOn bool) Option {
	return optionFunc(func(o *options) { o.query["default_on"] = defaultOn })
}

// WithMinimalQuantity minimal_quantity获取 
func (obj *_EgProductAttributeMgr) WithMinimalQuantity(minimalQuantity uint32) Option {
	return optionFunc(func(o *options) { o.query["minimal_quantity"] = minimalQuantity })
}

// WithLowStockThreshold low_stock_threshold获取 
func (obj *_EgProductAttributeMgr) WithLowStockThreshold(lowStockThreshold int) Option {
	return optionFunc(func(o *options) { o.query["low_stock_threshold"] = lowStockThreshold })
}

// WithLowStockAlert low_stock_alert获取 
func (obj *_EgProductAttributeMgr) WithLowStockAlert(lowStockAlert bool) Option {
	return optionFunc(func(o *options) { o.query["low_stock_alert"] = lowStockAlert })
}

// WithAvailableDate available_date获取 
func (obj *_EgProductAttributeMgr) WithAvailableDate(availableDate datatypes.Date) Option {
	return optionFunc(func(o *options) { o.query["available_date"] = availableDate })
}


// GetByOption 功能选项模式获取
func (obj *_EgProductAttributeMgr) GetByOption(opts ...Option) (result EgProductAttribute, err error) {
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
func (obj *_EgProductAttributeMgr) GetByOptions(opts ...Option) (results []*EgProductAttribute, err error) {
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


// GetFromIDProductAttribute 通过id_product_attribute获取内容  
func (obj *_EgProductAttributeMgr)  GetFromIDProductAttribute(idProductAttribute uint32) (result EgProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_attribute = ?", idProductAttribute).Find(&result).Error
	
	return
}

// GetBatchFromIDProductAttribute 批量唯一主键查找 
func (obj *_EgProductAttributeMgr) GetBatchFromIDProductAttribute(idProductAttributes []uint32) (results []*EgProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_attribute IN (?)", idProductAttributes).Find(&results).Error
	
	return
}
 
// GetFromIDProduct 通过id_product获取内容  
func (obj *_EgProductAttributeMgr) GetFromIDProduct(idProduct uint32) (results []*EgProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ?", idProduct).Find(&results).Error
	
	return
}

// GetBatchFromIDProduct 批量唯一主键查找 
func (obj *_EgProductAttributeMgr) GetBatchFromIDProduct(idProducts []uint32) (results []*EgProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product IN (?)", idProducts).Find(&results).Error
	
	return
}
 
// GetFromReference 通过reference获取内容  
func (obj *_EgProductAttributeMgr) GetFromReference(reference string) (results []*EgProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reference = ?", reference).Find(&results).Error
	
	return
}

// GetBatchFromReference 批量唯一主键查找 
func (obj *_EgProductAttributeMgr) GetBatchFromReference(references []string) (results []*EgProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reference IN (?)", references).Find(&results).Error
	
	return
}
 
// GetFromSupplierReference 通过supplier_reference获取内容  
func (obj *_EgProductAttributeMgr) GetFromSupplierReference(supplierReference string) (results []*EgProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("supplier_reference = ?", supplierReference).Find(&results).Error
	
	return
}

// GetBatchFromSupplierReference 批量唯一主键查找 
func (obj *_EgProductAttributeMgr) GetBatchFromSupplierReference(supplierReferences []string) (results []*EgProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("supplier_reference IN (?)", supplierReferences).Find(&results).Error
	
	return
}
 
// GetFromLocation 通过location获取内容  
func (obj *_EgProductAttributeMgr) GetFromLocation(location string) (results []*EgProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("location = ?", location).Find(&results).Error
	
	return
}

// GetBatchFromLocation 批量唯一主键查找 
func (obj *_EgProductAttributeMgr) GetBatchFromLocation(locations []string) (results []*EgProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("location IN (?)", locations).Find(&results).Error
	
	return
}
 
// GetFromEan13 通过ean13获取内容  
func (obj *_EgProductAttributeMgr) GetFromEan13(ean13 string) (results []*EgProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ean13 = ?", ean13).Find(&results).Error
	
	return
}

// GetBatchFromEan13 批量唯一主键查找 
func (obj *_EgProductAttributeMgr) GetBatchFromEan13(ean13s []string) (results []*EgProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ean13 IN (?)", ean13s).Find(&results).Error
	
	return
}
 
// GetFromIsbn 通过isbn获取内容  
func (obj *_EgProductAttributeMgr) GetFromIsbn(isbn string) (results []*EgProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("isbn = ?", isbn).Find(&results).Error
	
	return
}

// GetBatchFromIsbn 批量唯一主键查找 
func (obj *_EgProductAttributeMgr) GetBatchFromIsbn(isbns []string) (results []*EgProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("isbn IN (?)", isbns).Find(&results).Error
	
	return
}
 
// GetFromUpc 通过upc获取内容  
func (obj *_EgProductAttributeMgr) GetFromUpc(upc string) (results []*EgProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("upc = ?", upc).Find(&results).Error
	
	return
}

// GetBatchFromUpc 批量唯一主键查找 
func (obj *_EgProductAttributeMgr) GetBatchFromUpc(upcs []string) (results []*EgProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("upc IN (?)", upcs).Find(&results).Error
	
	return
}
 
// GetFromMpn 通过mpn获取内容  
func (obj *_EgProductAttributeMgr) GetFromMpn(mpn string) (results []*EgProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("mpn = ?", mpn).Find(&results).Error
	
	return
}

// GetBatchFromMpn 批量唯一主键查找 
func (obj *_EgProductAttributeMgr) GetBatchFromMpn(mpns []string) (results []*EgProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("mpn IN (?)", mpns).Find(&results).Error
	
	return
}
 
// GetFromWholesalePrice 通过wholesale_price获取内容  
func (obj *_EgProductAttributeMgr) GetFromWholesalePrice(wholesalePrice float64) (results []*EgProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("wholesale_price = ?", wholesalePrice).Find(&results).Error
	
	return
}

// GetBatchFromWholesalePrice 批量唯一主键查找 
func (obj *_EgProductAttributeMgr) GetBatchFromWholesalePrice(wholesalePrices []float64) (results []*EgProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("wholesale_price IN (?)", wholesalePrices).Find(&results).Error
	
	return
}
 
// GetFromPrice 通过price获取内容  
func (obj *_EgProductAttributeMgr) GetFromPrice(price float64) (results []*EgProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("price = ?", price).Find(&results).Error
	
	return
}

// GetBatchFromPrice 批量唯一主键查找 
func (obj *_EgProductAttributeMgr) GetBatchFromPrice(prices []float64) (results []*EgProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("price IN (?)", prices).Find(&results).Error
	
	return
}
 
// GetFromEcotax 通过ecotax获取内容  
func (obj *_EgProductAttributeMgr) GetFromEcotax(ecotax float64) (results []*EgProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ecotax = ?", ecotax).Find(&results).Error
	
	return
}

// GetBatchFromEcotax 批量唯一主键查找 
func (obj *_EgProductAttributeMgr) GetBatchFromEcotax(ecotaxs []float64) (results []*EgProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ecotax IN (?)", ecotaxs).Find(&results).Error
	
	return
}
 
// GetFromQuantity 通过quantity获取内容  
func (obj *_EgProductAttributeMgr) GetFromQuantity(quantity int) (results []*EgProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("quantity = ?", quantity).Find(&results).Error
	
	return
}

// GetBatchFromQuantity 批量唯一主键查找 
func (obj *_EgProductAttributeMgr) GetBatchFromQuantity(quantitys []int) (results []*EgProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("quantity IN (?)", quantitys).Find(&results).Error
	
	return
}
 
// GetFromWeight 通过weight获取内容  
func (obj *_EgProductAttributeMgr) GetFromWeight(weight float64) (results []*EgProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("weight = ?", weight).Find(&results).Error
	
	return
}

// GetBatchFromWeight 批量唯一主键查找 
func (obj *_EgProductAttributeMgr) GetBatchFromWeight(weights []float64) (results []*EgProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("weight IN (?)", weights).Find(&results).Error
	
	return
}
 
// GetFromUnitPriceImpact 通过unit_price_impact获取内容  
func (obj *_EgProductAttributeMgr) GetFromUnitPriceImpact(unitPriceImpact float64) (results []*EgProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("unit_price_impact = ?", unitPriceImpact).Find(&results).Error
	
	return
}

// GetBatchFromUnitPriceImpact 批量唯一主键查找 
func (obj *_EgProductAttributeMgr) GetBatchFromUnitPriceImpact(unitPriceImpacts []float64) (results []*EgProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("unit_price_impact IN (?)", unitPriceImpacts).Find(&results).Error
	
	return
}
 
// GetFromDefaultOn 通过default_on获取内容  
func (obj *_EgProductAttributeMgr) GetFromDefaultOn(defaultOn bool) (results []*EgProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("default_on = ?", defaultOn).Find(&results).Error
	
	return
}

// GetBatchFromDefaultOn 批量唯一主键查找 
func (obj *_EgProductAttributeMgr) GetBatchFromDefaultOn(defaultOns []bool) (results []*EgProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("default_on IN (?)", defaultOns).Find(&results).Error
	
	return
}
 
// GetFromMinimalQuantity 通过minimal_quantity获取内容  
func (obj *_EgProductAttributeMgr) GetFromMinimalQuantity(minimalQuantity uint32) (results []*EgProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("minimal_quantity = ?", minimalQuantity).Find(&results).Error
	
	return
}

// GetBatchFromMinimalQuantity 批量唯一主键查找 
func (obj *_EgProductAttributeMgr) GetBatchFromMinimalQuantity(minimalQuantitys []uint32) (results []*EgProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("minimal_quantity IN (?)", minimalQuantitys).Find(&results).Error
	
	return
}
 
// GetFromLowStockThreshold 通过low_stock_threshold获取内容  
func (obj *_EgProductAttributeMgr) GetFromLowStockThreshold(lowStockThreshold int) (results []*EgProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("low_stock_threshold = ?", lowStockThreshold).Find(&results).Error
	
	return
}

// GetBatchFromLowStockThreshold 批量唯一主键查找 
func (obj *_EgProductAttributeMgr) GetBatchFromLowStockThreshold(lowStockThresholds []int) (results []*EgProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("low_stock_threshold IN (?)", lowStockThresholds).Find(&results).Error
	
	return
}
 
// GetFromLowStockAlert 通过low_stock_alert获取内容  
func (obj *_EgProductAttributeMgr) GetFromLowStockAlert(lowStockAlert bool) (results []*EgProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("low_stock_alert = ?", lowStockAlert).Find(&results).Error
	
	return
}

// GetBatchFromLowStockAlert 批量唯一主键查找 
func (obj *_EgProductAttributeMgr) GetBatchFromLowStockAlert(lowStockAlerts []bool) (results []*EgProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("low_stock_alert IN (?)", lowStockAlerts).Find(&results).Error
	
	return
}
 
// GetFromAvailableDate 通过available_date获取内容  
func (obj *_EgProductAttributeMgr) GetFromAvailableDate(availableDate datatypes.Date) (results []*EgProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("available_date = ?", availableDate).Find(&results).Error
	
	return
}

// GetBatchFromAvailableDate 批量唯一主键查找 
func (obj *_EgProductAttributeMgr) GetBatchFromAvailableDate(availableDates []datatypes.Date) (results []*EgProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("available_date IN (?)", availableDates).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgProductAttributeMgr) FetchByPrimaryKey(idProductAttribute uint32 ) (result EgProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_attribute = ?", idProductAttribute).Find(&result).Error
	
	return
}
 
 // FetchUniqueIndexByProductDefault primay or index 获取唯一内容
 func (obj *_EgProductAttributeMgr) FetchUniqueIndexByProductDefault(idProduct uint32 ,defaultOn bool ) (result EgProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ? AND default_on = ?", idProduct , defaultOn).Find(&result).Error
	
	return
}
 

 
 // FetchIndexByIDProductIDProductAttribute  获取多个内容
 func (obj *_EgProductAttributeMgr) FetchIndexByIDProductIDProductAttribute(idProductAttribute uint32 ,idProduct uint32 ) (results []*EgProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_attribute = ? AND id_product = ?", idProductAttribute , idProduct).Find(&results).Error
	
	return
}
 
 // FetchIndexByProductAttributeProduct  获取多个内容
 func (obj *_EgProductAttributeMgr) FetchIndexByProductAttributeProduct(idProduct uint32 ) (results []*EgProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ?", idProduct).Find(&results).Error
	
	return
}
 
 // FetchIndexByReference  获取多个内容
 func (obj *_EgProductAttributeMgr) FetchIndexByReference(reference string ) (results []*EgProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reference = ?", reference).Find(&results).Error
	
	return
}
 
 // FetchIndexBySupplierReference  获取多个内容
 func (obj *_EgProductAttributeMgr) FetchIndexBySupplierReference(supplierReference string ) (results []*EgProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("supplier_reference = ?", supplierReference).Find(&results).Error
	
	return
}
 

	

