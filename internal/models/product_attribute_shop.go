package model

import (
	"context"
	"fmt"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type _ProductAttributeShopMgr struct {
	*_BaseMgr
}

// ProductAttributeShopMgr open func
func ProductAttributeShopMgr(db *gorm.DB) *_ProductAttributeShopMgr {
	if db == nil {
		panic(fmt.Errorf("ProductAttributeShopMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ProductAttributeShopMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_product_attribute_shop"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ProductAttributeShopMgr) GetTableName() string {
	return "ps_product_attribute_shop"
}

// Get 获取
func (obj *_ProductAttributeShopMgr) Get() (result ProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ProductAttributeShopMgr) Gets() (results []*ProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDProduct id_product获取
func (obj *_ProductAttributeShopMgr) WithIDProduct(idProduct uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product"] = idProduct })
}

// WithIDProductAttribute id_product_attribute获取
func (obj *_ProductAttributeShopMgr) WithIDProductAttribute(idProductAttribute uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product_attribute"] = idProductAttribute })
}

// WithIDShop id_shop获取
func (obj *_ProductAttributeShopMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

// WithWholesalePrice wholesale_price获取
func (obj *_ProductAttributeShopMgr) WithWholesalePrice(wholesalePrice float64) Option {
	return optionFunc(func(o *options) { o.query["wholesale_price"] = wholesalePrice })
}

// WithPrice price获取
func (obj *_ProductAttributeShopMgr) WithPrice(price float64) Option {
	return optionFunc(func(o *options) { o.query["price"] = price })
}

// WithEcotax ecotax获取
func (obj *_ProductAttributeShopMgr) WithEcotax(ecotax float64) Option {
	return optionFunc(func(o *options) { o.query["ecotax"] = ecotax })
}

// WithWeight weight获取
func (obj *_ProductAttributeShopMgr) WithWeight(weight float64) Option {
	return optionFunc(func(o *options) { o.query["weight"] = weight })
}

// WithUnitPriceImpact unit_price_impact获取
func (obj *_ProductAttributeShopMgr) WithUnitPriceImpact(unitPriceImpact float64) Option {
	return optionFunc(func(o *options) { o.query["unit_price_impact"] = unitPriceImpact })
}

// WithDefaultOn default_on获取
func (obj *_ProductAttributeShopMgr) WithDefaultOn(defaultOn bool) Option {
	return optionFunc(func(o *options) { o.query["default_on"] = defaultOn })
}

// WithMinimalQuantity minimal_quantity获取
func (obj *_ProductAttributeShopMgr) WithMinimalQuantity(minimalQuantity uint32) Option {
	return optionFunc(func(o *options) { o.query["minimal_quantity"] = minimalQuantity })
}

// WithLowStockThreshold low_stock_threshold获取
func (obj *_ProductAttributeShopMgr) WithLowStockThreshold(lowStockThreshold int) Option {
	return optionFunc(func(o *options) { o.query["low_stock_threshold"] = lowStockThreshold })
}

// WithLowStockAlert low_stock_alert获取
func (obj *_ProductAttributeShopMgr) WithLowStockAlert(lowStockAlert bool) Option {
	return optionFunc(func(o *options) { o.query["low_stock_alert"] = lowStockAlert })
}

// WithAvailableDate available_date获取
func (obj *_ProductAttributeShopMgr) WithAvailableDate(availableDate datatypes.Date) Option {
	return optionFunc(func(o *options) { o.query["available_date"] = availableDate })
}

// GetByOption 功能选项模式获取
func (obj *_ProductAttributeShopMgr) GetByOption(opts ...Option) (result ProductAttributeShop, err error) {
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
func (obj *_ProductAttributeShopMgr) GetByOptions(opts ...Option) (results []*ProductAttributeShop, err error) {
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
func (obj *_ProductAttributeShopMgr) GetFromIDProduct(idProduct uint32) (results []*ProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ?", idProduct).Find(&results).Error

	return
}

// GetBatchFromIDProduct 批量唯一主键查找
func (obj *_ProductAttributeShopMgr) GetBatchFromIDProduct(idProducts []uint32) (results []*ProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product IN (?)", idProducts).Find(&results).Error

	return
}

// GetFromIDProductAttribute 通过id_product_attribute获取内容
func (obj *_ProductAttributeShopMgr) GetFromIDProductAttribute(idProductAttribute uint32) (results []*ProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_attribute = ?", idProductAttribute).Find(&results).Error

	return
}

// GetBatchFromIDProductAttribute 批量唯一主键查找
func (obj *_ProductAttributeShopMgr) GetBatchFromIDProductAttribute(idProductAttributes []uint32) (results []*ProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_attribute IN (?)", idProductAttributes).Find(&results).Error

	return
}

// GetFromIDShop 通过id_shop获取内容
func (obj *_ProductAttributeShopMgr) GetFromIDShop(idShop uint32) (results []*ProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

// GetBatchFromIDShop 批量唯一主键查找
func (obj *_ProductAttributeShopMgr) GetBatchFromIDShop(idShops []uint32) (results []*ProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

// GetFromWholesalePrice 通过wholesale_price获取内容
func (obj *_ProductAttributeShopMgr) GetFromWholesalePrice(wholesalePrice float64) (results []*ProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("wholesale_price = ?", wholesalePrice).Find(&results).Error

	return
}

// GetBatchFromWholesalePrice 批量唯一主键查找
func (obj *_ProductAttributeShopMgr) GetBatchFromWholesalePrice(wholesalePrices []float64) (results []*ProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("wholesale_price IN (?)", wholesalePrices).Find(&results).Error

	return
}

// GetFromPrice 通过price获取内容
func (obj *_ProductAttributeShopMgr) GetFromPrice(price float64) (results []*ProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("price = ?", price).Find(&results).Error

	return
}

// GetBatchFromPrice 批量唯一主键查找
func (obj *_ProductAttributeShopMgr) GetBatchFromPrice(prices []float64) (results []*ProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("price IN (?)", prices).Find(&results).Error

	return
}

// GetFromEcotax 通过ecotax获取内容
func (obj *_ProductAttributeShopMgr) GetFromEcotax(ecotax float64) (results []*ProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ecotax = ?", ecotax).Find(&results).Error

	return
}

// GetBatchFromEcotax 批量唯一主键查找
func (obj *_ProductAttributeShopMgr) GetBatchFromEcotax(ecotaxs []float64) (results []*ProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ecotax IN (?)", ecotaxs).Find(&results).Error

	return
}

// GetFromWeight 通过weight获取内容
func (obj *_ProductAttributeShopMgr) GetFromWeight(weight float64) (results []*ProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("weight = ?", weight).Find(&results).Error

	return
}

// GetBatchFromWeight 批量唯一主键查找
func (obj *_ProductAttributeShopMgr) GetBatchFromWeight(weights []float64) (results []*ProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("weight IN (?)", weights).Find(&results).Error

	return
}

// GetFromUnitPriceImpact 通过unit_price_impact获取内容
func (obj *_ProductAttributeShopMgr) GetFromUnitPriceImpact(unitPriceImpact float64) (results []*ProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("unit_price_impact = ?", unitPriceImpact).Find(&results).Error

	return
}

// GetBatchFromUnitPriceImpact 批量唯一主键查找
func (obj *_ProductAttributeShopMgr) GetBatchFromUnitPriceImpact(unitPriceImpacts []float64) (results []*ProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("unit_price_impact IN (?)", unitPriceImpacts).Find(&results).Error

	return
}

// GetFromDefaultOn 通过default_on获取内容
func (obj *_ProductAttributeShopMgr) GetFromDefaultOn(defaultOn bool) (results []*ProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("default_on = ?", defaultOn).Find(&results).Error

	return
}

// GetBatchFromDefaultOn 批量唯一主键查找
func (obj *_ProductAttributeShopMgr) GetBatchFromDefaultOn(defaultOns []bool) (results []*ProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("default_on IN (?)", defaultOns).Find(&results).Error

	return
}

// GetFromMinimalQuantity 通过minimal_quantity获取内容
func (obj *_ProductAttributeShopMgr) GetFromMinimalQuantity(minimalQuantity uint32) (results []*ProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("minimal_quantity = ?", minimalQuantity).Find(&results).Error

	return
}

// GetBatchFromMinimalQuantity 批量唯一主键查找
func (obj *_ProductAttributeShopMgr) GetBatchFromMinimalQuantity(minimalQuantitys []uint32) (results []*ProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("minimal_quantity IN (?)", minimalQuantitys).Find(&results).Error

	return
}

// GetFromLowStockThreshold 通过low_stock_threshold获取内容
func (obj *_ProductAttributeShopMgr) GetFromLowStockThreshold(lowStockThreshold int) (results []*ProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("low_stock_threshold = ?", lowStockThreshold).Find(&results).Error

	return
}

// GetBatchFromLowStockThreshold 批量唯一主键查找
func (obj *_ProductAttributeShopMgr) GetBatchFromLowStockThreshold(lowStockThresholds []int) (results []*ProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("low_stock_threshold IN (?)", lowStockThresholds).Find(&results).Error

	return
}

// GetFromLowStockAlert 通过low_stock_alert获取内容
func (obj *_ProductAttributeShopMgr) GetFromLowStockAlert(lowStockAlert bool) (results []*ProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("low_stock_alert = ?", lowStockAlert).Find(&results).Error

	return
}

// GetBatchFromLowStockAlert 批量唯一主键查找
func (obj *_ProductAttributeShopMgr) GetBatchFromLowStockAlert(lowStockAlerts []bool) (results []*ProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("low_stock_alert IN (?)", lowStockAlerts).Find(&results).Error

	return
}

// GetFromAvailableDate 通过available_date获取内容
func (obj *_ProductAttributeShopMgr) GetFromAvailableDate(availableDate datatypes.Date) (results []*ProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("available_date = ?", availableDate).Find(&results).Error

	return
}

// GetBatchFromAvailableDate 批量唯一主键查找
func (obj *_ProductAttributeShopMgr) GetBatchFromAvailableDate(availableDates []datatypes.Date) (results []*ProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("available_date IN (?)", availableDates).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_ProductAttributeShopMgr) FetchByPrimaryKey(idProductAttribute uint32, idShop uint32) (result ProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_attribute = ? AND id_shop = ?", idProductAttribute, idShop).Find(&result).Error

	return
}

// FetchUniqueIndexByIDProduct primay or index 获取唯一内容
func (obj *_ProductAttributeShopMgr) FetchUniqueIndexByIDProduct(idProduct uint32, idShop uint32, defaultOn bool) (result ProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ? AND id_shop = ? AND default_on = ?", idProduct, idShop, defaultOn).Find(&result).Error

	return
}
