package models

import (
	"context"
	"fmt"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type _ProductAttributeShopMgr struct {
	*_BaseMgr
}

func ProductAttributeShopMgr(db *gorm.DB) *_ProductAttributeShopMgr {
	if db == nil {
		panic(fmt.Errorf("ProductAttributeShopMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ProductAttributeShopMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_product_attribute_shop"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_ProductAttributeShopMgr) GetTableName() string {
	return "ps_product_attribute_shop"
}

func (obj *_ProductAttributeShopMgr) Get() (result ProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_ProductAttributeShopMgr) Gets() (results []*ProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_ProductAttributeShopMgr) WithIDProduct(idProduct uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product"] = idProduct })
}

func (obj *_ProductAttributeShopMgr) WithIDProductAttribute(idProductAttribute uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product_attribute"] = idProductAttribute })
}

func (obj *_ProductAttributeShopMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

func (obj *_ProductAttributeShopMgr) WithWholesalePrice(wholesalePrice float64) Option {
	return optionFunc(func(o *options) { o.query["wholesale_price"] = wholesalePrice })
}

func (obj *_ProductAttributeShopMgr) WithPrice(price float64) Option {
	return optionFunc(func(o *options) { o.query["price"] = price })
}

func (obj *_ProductAttributeShopMgr) WithEcotax(ecotax float64) Option {
	return optionFunc(func(o *options) { o.query["ecotax"] = ecotax })
}

func (obj *_ProductAttributeShopMgr) WithWeight(weight float64) Option {
	return optionFunc(func(o *options) { o.query["weight"] = weight })
}

func (obj *_ProductAttributeShopMgr) WithUnitPriceImpact(unitPriceImpact float64) Option {
	return optionFunc(func(o *options) { o.query["unit_price_impact"] = unitPriceImpact })
}

func (obj *_ProductAttributeShopMgr) WithDefaultOn(defaultOn bool) Option {
	return optionFunc(func(o *options) { o.query["default_on"] = defaultOn })
}

func (obj *_ProductAttributeShopMgr) WithMinimalQuantity(minimalQuantity uint32) Option {
	return optionFunc(func(o *options) { o.query["minimal_quantity"] = minimalQuantity })
}

func (obj *_ProductAttributeShopMgr) WithLowStockThreshold(lowStockThreshold int) Option {
	return optionFunc(func(o *options) { o.query["low_stock_threshold"] = lowStockThreshold })
}

func (obj *_ProductAttributeShopMgr) WithLowStockAlert(lowStockAlert bool) Option {
	return optionFunc(func(o *options) { o.query["low_stock_alert"] = lowStockAlert })
}

func (obj *_ProductAttributeShopMgr) WithAvailableDate(availableDate datatypes.Date) Option {
	return optionFunc(func(o *options) { o.query["available_date"] = availableDate })
}

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


func (obj *_ProductAttributeShopMgr) GetFromIDProduct(idProduct uint32) (results []*ProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ?", idProduct).Find(&results).Error

	return
}

func (obj *_ProductAttributeShopMgr) GetBatchFromIDProduct(idProducts []uint32) (results []*ProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product IN (?)", idProducts).Find(&results).Error

	return
}

func (obj *_ProductAttributeShopMgr) GetFromIDProductAttribute(idProductAttribute uint32) (results []*ProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_attribute = ?", idProductAttribute).Find(&results).Error

	return
}

func (obj *_ProductAttributeShopMgr) GetBatchFromIDProductAttribute(idProductAttributes []uint32) (results []*ProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_attribute IN (?)", idProductAttributes).Find(&results).Error

	return
}

func (obj *_ProductAttributeShopMgr) GetFromIDShop(idShop uint32) (results []*ProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

func (obj *_ProductAttributeShopMgr) GetBatchFromIDShop(idShops []uint32) (results []*ProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

func (obj *_ProductAttributeShopMgr) GetFromWholesalePrice(wholesalePrice float64) (results []*ProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("wholesale_price = ?", wholesalePrice).Find(&results).Error

	return
}

func (obj *_ProductAttributeShopMgr) GetBatchFromWholesalePrice(wholesalePrices []float64) (results []*ProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("wholesale_price IN (?)", wholesalePrices).Find(&results).Error

	return
}

func (obj *_ProductAttributeShopMgr) GetFromPrice(price float64) (results []*ProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("price = ?", price).Find(&results).Error

	return
}

func (obj *_ProductAttributeShopMgr) GetBatchFromPrice(prices []float64) (results []*ProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("price IN (?)", prices).Find(&results).Error

	return
}

func (obj *_ProductAttributeShopMgr) GetFromEcotax(ecotax float64) (results []*ProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ecotax = ?", ecotax).Find(&results).Error

	return
}

func (obj *_ProductAttributeShopMgr) GetBatchFromEcotax(ecotaxs []float64) (results []*ProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ecotax IN (?)", ecotaxs).Find(&results).Error

	return
}

func (obj *_ProductAttributeShopMgr) GetFromWeight(weight float64) (results []*ProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("weight = ?", weight).Find(&results).Error

	return
}

func (obj *_ProductAttributeShopMgr) GetBatchFromWeight(weights []float64) (results []*ProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("weight IN (?)", weights).Find(&results).Error

	return
}

func (obj *_ProductAttributeShopMgr) GetFromUnitPriceImpact(unitPriceImpact float64) (results []*ProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("unit_price_impact = ?", unitPriceImpact).Find(&results).Error

	return
}

func (obj *_ProductAttributeShopMgr) GetBatchFromUnitPriceImpact(unitPriceImpacts []float64) (results []*ProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("unit_price_impact IN (?)", unitPriceImpacts).Find(&results).Error

	return
}

func (obj *_ProductAttributeShopMgr) GetFromDefaultOn(defaultOn bool) (results []*ProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("default_on = ?", defaultOn).Find(&results).Error

	return
}

func (obj *_ProductAttributeShopMgr) GetBatchFromDefaultOn(defaultOns []bool) (results []*ProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("default_on IN (?)", defaultOns).Find(&results).Error

	return
}

func (obj *_ProductAttributeShopMgr) GetFromMinimalQuantity(minimalQuantity uint32) (results []*ProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("minimal_quantity = ?", minimalQuantity).Find(&results).Error

	return
}

func (obj *_ProductAttributeShopMgr) GetBatchFromMinimalQuantity(minimalQuantitys []uint32) (results []*ProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("minimal_quantity IN (?)", minimalQuantitys).Find(&results).Error

	return
}

func (obj *_ProductAttributeShopMgr) GetFromLowStockThreshold(lowStockThreshold int) (results []*ProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("low_stock_threshold = ?", lowStockThreshold).Find(&results).Error

	return
}

func (obj *_ProductAttributeShopMgr) GetBatchFromLowStockThreshold(lowStockThresholds []int) (results []*ProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("low_stock_threshold IN (?)", lowStockThresholds).Find(&results).Error

	return
}

func (obj *_ProductAttributeShopMgr) GetFromLowStockAlert(lowStockAlert bool) (results []*ProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("low_stock_alert = ?", lowStockAlert).Find(&results).Error

	return
}

func (obj *_ProductAttributeShopMgr) GetBatchFromLowStockAlert(lowStockAlerts []bool) (results []*ProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("low_stock_alert IN (?)", lowStockAlerts).Find(&results).Error

	return
}

func (obj *_ProductAttributeShopMgr) GetFromAvailableDate(availableDate datatypes.Date) (results []*ProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("available_date = ?", availableDate).Find(&results).Error

	return
}

func (obj *_ProductAttributeShopMgr) GetBatchFromAvailableDate(availableDates []datatypes.Date) (results []*ProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("available_date IN (?)", availableDates).Find(&results).Error

	return
}


func (obj *_ProductAttributeShopMgr) FetchByPrimaryKey(idProductAttribute uint32, idShop uint32) (result ProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_attribute = ? AND id_shop = ?", idProductAttribute, idShop).Find(&result).Error

	return
}

func (obj *_ProductAttributeShopMgr) FetchUniqueIndexByIDProduct(idProduct uint32, idShop uint32, defaultOn bool) (result ProductAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ? AND id_shop = ? AND default_on = ?", idProduct, idShop, defaultOn).Find(&result).Error

	return
}
