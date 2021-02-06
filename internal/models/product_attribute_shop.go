package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
"gorm.io/datatypes"	
)	

type _EgProductAttributeShopMgr struct {
	*_BaseMgr
}

// EgProductAttributeShopMgr open func
func EgProductAttributeShopMgr(db *gorm.DB) *_EgProductAttributeShopMgr {
	if db == nil {
		panic(fmt.Errorf("EgProductAttributeShopMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgProductAttributeShopMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_product_attribute_shop"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgProductAttributeShopMgr) GetTableName() string {
	return "eg_product_attribute_shop"
}

// Get 获取
func (obj *_EgProductAttributeShopMgr) Get() (result EgProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgProductAttributeShopMgr) Gets() (results []*EgProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDProduct id_product获取 
func (obj *_EgProductAttributeShopMgr) WithIDProduct(idProduct uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product"] = idProduct })
}

// WithIDProductAttribute id_product_attribute获取 
func (obj *_EgProductAttributeShopMgr) WithIDProductAttribute(idProductAttribute uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product_attribute"] = idProductAttribute })
}

// WithIDShop id_shop获取 
func (obj *_EgProductAttributeShopMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

// WithWholesalePrice wholesale_price获取 
func (obj *_EgProductAttributeShopMgr) WithWholesalePrice(wholesalePrice float64) Option {
	return optionFunc(func(o *options) { o.query["wholesale_price"] = wholesalePrice })
}

// WithPrice price获取 
func (obj *_EgProductAttributeShopMgr) WithPrice(price float64) Option {
	return optionFunc(func(o *options) { o.query["price"] = price })
}

// WithEcotax ecotax获取 
func (obj *_EgProductAttributeShopMgr) WithEcotax(ecotax float64) Option {
	return optionFunc(func(o *options) { o.query["ecotax"] = ecotax })
}

// WithWeight weight获取 
func (obj *_EgProductAttributeShopMgr) WithWeight(weight float64) Option {
	return optionFunc(func(o *options) { o.query["weight"] = weight })
}

// WithUnitPriceImpact unit_price_impact获取 
func (obj *_EgProductAttributeShopMgr) WithUnitPriceImpact(unitPriceImpact float64) Option {
	return optionFunc(func(o *options) { o.query["unit_price_impact"] = unitPriceImpact })
}

// WithDefaultOn default_on获取 
func (obj *_EgProductAttributeShopMgr) WithDefaultOn(defaultOn bool) Option {
	return optionFunc(func(o *options) { o.query["default_on"] = defaultOn })
}

// WithMinimalQuantity minimal_quantity获取 
func (obj *_EgProductAttributeShopMgr) WithMinimalQuantity(minimalQuantity uint32) Option {
	return optionFunc(func(o *options) { o.query["minimal_quantity"] = minimalQuantity })
}

// WithLowStockThreshold low_stock_threshold获取 
func (obj *_EgProductAttributeShopMgr) WithLowStockThreshold(lowStockThreshold int) Option {
	return optionFunc(func(o *options) { o.query["low_stock_threshold"] = lowStockThreshold })
}

// WithLowStockAlert low_stock_alert获取 
func (obj *_EgProductAttributeShopMgr) WithLowStockAlert(lowStockAlert bool) Option {
	return optionFunc(func(o *options) { o.query["low_stock_alert"] = lowStockAlert })
}

// WithAvailableDate available_date获取 
func (obj *_EgProductAttributeShopMgr) WithAvailableDate(availableDate datatypes.Date) Option {
	return optionFunc(func(o *options) { o.query["available_date"] = availableDate })
}


// GetByOption 功能选项模式获取
func (obj *_EgProductAttributeShopMgr) GetByOption(opts ...Option) (result EgProductAttributeShop, err error) {
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
func (obj *_EgProductAttributeShopMgr) GetByOptions(opts ...Option) (results []*EgProductAttributeShop, err error) {
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
func (obj *_EgProductAttributeShopMgr) GetFromIDProduct(idProduct uint32) (results []*EgProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ?", idProduct).Find(&results).Error
	
	return
}

// GetBatchFromIDProduct 批量唯一主键查找 
func (obj *_EgProductAttributeShopMgr) GetBatchFromIDProduct(idProducts []uint32) (results []*EgProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product IN (?)", idProducts).Find(&results).Error
	
	return
}
 
// GetFromIDProductAttribute 通过id_product_attribute获取内容  
func (obj *_EgProductAttributeShopMgr) GetFromIDProductAttribute(idProductAttribute uint32) (results []*EgProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_attribute = ?", idProductAttribute).Find(&results).Error
	
	return
}

// GetBatchFromIDProductAttribute 批量唯一主键查找 
func (obj *_EgProductAttributeShopMgr) GetBatchFromIDProductAttribute(idProductAttributes []uint32) (results []*EgProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_attribute IN (?)", idProductAttributes).Find(&results).Error
	
	return
}
 
// GetFromIDShop 通过id_shop获取内容  
func (obj *_EgProductAttributeShopMgr) GetFromIDShop(idShop uint32) (results []*EgProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error
	
	return
}

// GetBatchFromIDShop 批量唯一主键查找 
func (obj *_EgProductAttributeShopMgr) GetBatchFromIDShop(idShops []uint32) (results []*EgProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error
	
	return
}
 
// GetFromWholesalePrice 通过wholesale_price获取内容  
func (obj *_EgProductAttributeShopMgr) GetFromWholesalePrice(wholesalePrice float64) (results []*EgProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("wholesale_price = ?", wholesalePrice).Find(&results).Error
	
	return
}

// GetBatchFromWholesalePrice 批量唯一主键查找 
func (obj *_EgProductAttributeShopMgr) GetBatchFromWholesalePrice(wholesalePrices []float64) (results []*EgProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("wholesale_price IN (?)", wholesalePrices).Find(&results).Error
	
	return
}
 
// GetFromPrice 通过price获取内容  
func (obj *_EgProductAttributeShopMgr) GetFromPrice(price float64) (results []*EgProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("price = ?", price).Find(&results).Error
	
	return
}

// GetBatchFromPrice 批量唯一主键查找 
func (obj *_EgProductAttributeShopMgr) GetBatchFromPrice(prices []float64) (results []*EgProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("price IN (?)", prices).Find(&results).Error
	
	return
}
 
// GetFromEcotax 通过ecotax获取内容  
func (obj *_EgProductAttributeShopMgr) GetFromEcotax(ecotax float64) (results []*EgProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ecotax = ?", ecotax).Find(&results).Error
	
	return
}

// GetBatchFromEcotax 批量唯一主键查找 
func (obj *_EgProductAttributeShopMgr) GetBatchFromEcotax(ecotaxs []float64) (results []*EgProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ecotax IN (?)", ecotaxs).Find(&results).Error
	
	return
}
 
// GetFromWeight 通过weight获取内容  
func (obj *_EgProductAttributeShopMgr) GetFromWeight(weight float64) (results []*EgProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("weight = ?", weight).Find(&results).Error
	
	return
}

// GetBatchFromWeight 批量唯一主键查找 
func (obj *_EgProductAttributeShopMgr) GetBatchFromWeight(weights []float64) (results []*EgProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("weight IN (?)", weights).Find(&results).Error
	
	return
}
 
// GetFromUnitPriceImpact 通过unit_price_impact获取内容  
func (obj *_EgProductAttributeShopMgr) GetFromUnitPriceImpact(unitPriceImpact float64) (results []*EgProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("unit_price_impact = ?", unitPriceImpact).Find(&results).Error
	
	return
}

// GetBatchFromUnitPriceImpact 批量唯一主键查找 
func (obj *_EgProductAttributeShopMgr) GetBatchFromUnitPriceImpact(unitPriceImpacts []float64) (results []*EgProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("unit_price_impact IN (?)", unitPriceImpacts).Find(&results).Error
	
	return
}
 
// GetFromDefaultOn 通过default_on获取内容  
func (obj *_EgProductAttributeShopMgr) GetFromDefaultOn(defaultOn bool) (results []*EgProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("default_on = ?", defaultOn).Find(&results).Error
	
	return
}

// GetBatchFromDefaultOn 批量唯一主键查找 
func (obj *_EgProductAttributeShopMgr) GetBatchFromDefaultOn(defaultOns []bool) (results []*EgProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("default_on IN (?)", defaultOns).Find(&results).Error
	
	return
}
 
// GetFromMinimalQuantity 通过minimal_quantity获取内容  
func (obj *_EgProductAttributeShopMgr) GetFromMinimalQuantity(minimalQuantity uint32) (results []*EgProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("minimal_quantity = ?", minimalQuantity).Find(&results).Error
	
	return
}

// GetBatchFromMinimalQuantity 批量唯一主键查找 
func (obj *_EgProductAttributeShopMgr) GetBatchFromMinimalQuantity(minimalQuantitys []uint32) (results []*EgProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("minimal_quantity IN (?)", minimalQuantitys).Find(&results).Error
	
	return
}
 
// GetFromLowStockThreshold 通过low_stock_threshold获取内容  
func (obj *_EgProductAttributeShopMgr) GetFromLowStockThreshold(lowStockThreshold int) (results []*EgProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("low_stock_threshold = ?", lowStockThreshold).Find(&results).Error
	
	return
}

// GetBatchFromLowStockThreshold 批量唯一主键查找 
func (obj *_EgProductAttributeShopMgr) GetBatchFromLowStockThreshold(lowStockThresholds []int) (results []*EgProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("low_stock_threshold IN (?)", lowStockThresholds).Find(&results).Error
	
	return
}
 
// GetFromLowStockAlert 通过low_stock_alert获取内容  
func (obj *_EgProductAttributeShopMgr) GetFromLowStockAlert(lowStockAlert bool) (results []*EgProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("low_stock_alert = ?", lowStockAlert).Find(&results).Error
	
	return
}

// GetBatchFromLowStockAlert 批量唯一主键查找 
func (obj *_EgProductAttributeShopMgr) GetBatchFromLowStockAlert(lowStockAlerts []bool) (results []*EgProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("low_stock_alert IN (?)", lowStockAlerts).Find(&results).Error
	
	return
}
 
// GetFromAvailableDate 通过available_date获取内容  
func (obj *_EgProductAttributeShopMgr) GetFromAvailableDate(availableDate datatypes.Date) (results []*EgProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("available_date = ?", availableDate).Find(&results).Error
	
	return
}

// GetBatchFromAvailableDate 批量唯一主键查找 
func (obj *_EgProductAttributeShopMgr) GetBatchFromAvailableDate(availableDates []datatypes.Date) (results []*EgProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("available_date IN (?)", availableDates).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgProductAttributeShopMgr) FetchByPrimaryKey(idProductAttribute uint32 ,idShop uint32 ) (result EgProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_attribute = ? AND id_shop = ?", idProductAttribute , idShop).Find(&result).Error
	
	return
}
 
 // FetchUniqueIndexByIDProduct primay or index 获取唯一内容
 func (obj *_EgProductAttributeShopMgr) FetchUniqueIndexByIDProduct(idProduct uint32 ,idShop uint32 ,defaultOn bool ) (result EgProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ? AND id_shop = ? AND default_on = ?", idProduct , idShop , defaultOn).Find(&result).Error
	
	return
}
 

 

	

