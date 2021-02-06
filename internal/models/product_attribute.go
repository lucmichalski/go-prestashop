package model

import (
	"context"
	"fmt"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type _ProductAttributeMgr struct {
	*_BaseMgr
}

func ProductAttributeMgr(db *gorm.DB) *_ProductAttributeMgr {
	if db == nil {
		panic(fmt.Errorf("ProductAttributeMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ProductAttributeMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_product_attribute"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_ProductAttributeMgr) GetTableName() string {
	return "ps_product_attribute"
}

func (obj *_ProductAttributeMgr) Get() (result ProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_ProductAttributeMgr) Gets() (results []*ProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_ProductAttributeMgr) WithIDProductAttribute(idProductAttribute uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product_attribute"] = idProductAttribute })
}

func (obj *_ProductAttributeMgr) WithIDProduct(idProduct uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product"] = idProduct })
}

func (obj *_ProductAttributeMgr) WithReference(reference string) Option {
	return optionFunc(func(o *options) { o.query["reference"] = reference })
}

func (obj *_ProductAttributeMgr) WithSupplierReference(supplierReference string) Option {
	return optionFunc(func(o *options) { o.query["supplier_reference"] = supplierReference })
}

func (obj *_ProductAttributeMgr) WithLocation(location string) Option {
	return optionFunc(func(o *options) { o.query["location"] = location })
}

func (obj *_ProductAttributeMgr) WithEan13(ean13 string) Option {
	return optionFunc(func(o *options) { o.query["ean13"] = ean13 })
}

func (obj *_ProductAttributeMgr) WithIsbn(isbn string) Option {
	return optionFunc(func(o *options) { o.query["isbn"] = isbn })
}

func (obj *_ProductAttributeMgr) WithUpc(upc string) Option {
	return optionFunc(func(o *options) { o.query["upc"] = upc })
}

func (obj *_ProductAttributeMgr) WithMpn(mpn string) Option {
	return optionFunc(func(o *options) { o.query["mpn"] = mpn })
}

func (obj *_ProductAttributeMgr) WithWholesalePrice(wholesalePrice float64) Option {
	return optionFunc(func(o *options) { o.query["wholesale_price"] = wholesalePrice })
}

func (obj *_ProductAttributeMgr) WithPrice(price float64) Option {
	return optionFunc(func(o *options) { o.query["price"] = price })
}

func (obj *_ProductAttributeMgr) WithEcotax(ecotax float64) Option {
	return optionFunc(func(o *options) { o.query["ecotax"] = ecotax })
}

func (obj *_ProductAttributeMgr) WithQuantity(quantity int) Option {
	return optionFunc(func(o *options) { o.query["quantity"] = quantity })
}

func (obj *_ProductAttributeMgr) WithWeight(weight float64) Option {
	return optionFunc(func(o *options) { o.query["weight"] = weight })
}

func (obj *_ProductAttributeMgr) WithUnitPriceImpact(unitPriceImpact float64) Option {
	return optionFunc(func(o *options) { o.query["unit_price_impact"] = unitPriceImpact })
}

func (obj *_ProductAttributeMgr) WithDefaultOn(defaultOn bool) Option {
	return optionFunc(func(o *options) { o.query["default_on"] = defaultOn })
}

func (obj *_ProductAttributeMgr) WithMinimalQuantity(minimalQuantity uint32) Option {
	return optionFunc(func(o *options) { o.query["minimal_quantity"] = minimalQuantity })
}

func (obj *_ProductAttributeMgr) WithLowStockThreshold(lowStockThreshold int) Option {
	return optionFunc(func(o *options) { o.query["low_stock_threshold"] = lowStockThreshold })
}

func (obj *_ProductAttributeMgr) WithLowStockAlert(lowStockAlert bool) Option {
	return optionFunc(func(o *options) { o.query["low_stock_alert"] = lowStockAlert })
}

func (obj *_ProductAttributeMgr) WithAvailableDate(availableDate datatypes.Date) Option {
	return optionFunc(func(o *options) { o.query["available_date"] = availableDate })
}

func (obj *_ProductAttributeMgr) GetByOption(opts ...Option) (result ProductAttribute, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_ProductAttributeMgr) GetByOptions(opts ...Option) (results []*ProductAttribute, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_ProductAttributeMgr) GetFromIDProductAttribute(idProductAttribute uint32) (result ProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_attribute = ?", idProductAttribute).Find(&result).Error

	return
}

func (obj *_ProductAttributeMgr) GetBatchFromIDProductAttribute(idProductAttributes []uint32) (results []*ProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_attribute IN (?)", idProductAttributes).Find(&results).Error

	return
}

func (obj *_ProductAttributeMgr) GetFromIDProduct(idProduct uint32) (results []*ProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ?", idProduct).Find(&results).Error

	return
}

func (obj *_ProductAttributeMgr) GetBatchFromIDProduct(idProducts []uint32) (results []*ProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product IN (?)", idProducts).Find(&results).Error

	return
}

func (obj *_ProductAttributeMgr) GetFromReference(reference string) (results []*ProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reference = ?", reference).Find(&results).Error

	return
}

func (obj *_ProductAttributeMgr) GetBatchFromReference(references []string) (results []*ProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reference IN (?)", references).Find(&results).Error

	return
}

func (obj *_ProductAttributeMgr) GetFromSupplierReference(supplierReference string) (results []*ProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("supplier_reference = ?", supplierReference).Find(&results).Error

	return
}

func (obj *_ProductAttributeMgr) GetBatchFromSupplierReference(supplierReferences []string) (results []*ProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("supplier_reference IN (?)", supplierReferences).Find(&results).Error

	return
}

func (obj *_ProductAttributeMgr) GetFromLocation(location string) (results []*ProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("location = ?", location).Find(&results).Error

	return
}

func (obj *_ProductAttributeMgr) GetBatchFromLocation(locations []string) (results []*ProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("location IN (?)", locations).Find(&results).Error

	return
}

func (obj *_ProductAttributeMgr) GetFromEan13(ean13 string) (results []*ProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ean13 = ?", ean13).Find(&results).Error

	return
}

func (obj *_ProductAttributeMgr) GetBatchFromEan13(ean13s []string) (results []*ProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ean13 IN (?)", ean13s).Find(&results).Error

	return
}

func (obj *_ProductAttributeMgr) GetFromIsbn(isbn string) (results []*ProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("isbn = ?", isbn).Find(&results).Error

	return
}

func (obj *_ProductAttributeMgr) GetBatchFromIsbn(isbns []string) (results []*ProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("isbn IN (?)", isbns).Find(&results).Error

	return
}

func (obj *_ProductAttributeMgr) GetFromUpc(upc string) (results []*ProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("upc = ?", upc).Find(&results).Error

	return
}

func (obj *_ProductAttributeMgr) GetBatchFromUpc(upcs []string) (results []*ProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("upc IN (?)", upcs).Find(&results).Error

	return
}

func (obj *_ProductAttributeMgr) GetFromMpn(mpn string) (results []*ProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("mpn = ?", mpn).Find(&results).Error

	return
}

func (obj *_ProductAttributeMgr) GetBatchFromMpn(mpns []string) (results []*ProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("mpn IN (?)", mpns).Find(&results).Error

	return
}

func (obj *_ProductAttributeMgr) GetFromWholesalePrice(wholesalePrice float64) (results []*ProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("wholesale_price = ?", wholesalePrice).Find(&results).Error

	return
}

func (obj *_ProductAttributeMgr) GetBatchFromWholesalePrice(wholesalePrices []float64) (results []*ProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("wholesale_price IN (?)", wholesalePrices).Find(&results).Error

	return
}

func (obj *_ProductAttributeMgr) GetFromPrice(price float64) (results []*ProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("price = ?", price).Find(&results).Error

	return
}

func (obj *_ProductAttributeMgr) GetBatchFromPrice(prices []float64) (results []*ProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("price IN (?)", prices).Find(&results).Error

	return
}

func (obj *_ProductAttributeMgr) GetFromEcotax(ecotax float64) (results []*ProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ecotax = ?", ecotax).Find(&results).Error

	return
}

func (obj *_ProductAttributeMgr) GetBatchFromEcotax(ecotaxs []float64) (results []*ProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ecotax IN (?)", ecotaxs).Find(&results).Error

	return
}

func (obj *_ProductAttributeMgr) GetFromQuantity(quantity int) (results []*ProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("quantity = ?", quantity).Find(&results).Error

	return
}

func (obj *_ProductAttributeMgr) GetBatchFromQuantity(quantitys []int) (results []*ProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("quantity IN (?)", quantitys).Find(&results).Error

	return
}

func (obj *_ProductAttributeMgr) GetFromWeight(weight float64) (results []*ProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("weight = ?", weight).Find(&results).Error

	return
}

func (obj *_ProductAttributeMgr) GetBatchFromWeight(weights []float64) (results []*ProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("weight IN (?)", weights).Find(&results).Error

	return
}

func (obj *_ProductAttributeMgr) GetFromUnitPriceImpact(unitPriceImpact float64) (results []*ProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("unit_price_impact = ?", unitPriceImpact).Find(&results).Error

	return
}

func (obj *_ProductAttributeMgr) GetBatchFromUnitPriceImpact(unitPriceImpacts []float64) (results []*ProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("unit_price_impact IN (?)", unitPriceImpacts).Find(&results).Error

	return
}

func (obj *_ProductAttributeMgr) GetFromDefaultOn(defaultOn bool) (results []*ProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("default_on = ?", defaultOn).Find(&results).Error

	return
}

func (obj *_ProductAttributeMgr) GetBatchFromDefaultOn(defaultOns []bool) (results []*ProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("default_on IN (?)", defaultOns).Find(&results).Error

	return
}

func (obj *_ProductAttributeMgr) GetFromMinimalQuantity(minimalQuantity uint32) (results []*ProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("minimal_quantity = ?", minimalQuantity).Find(&results).Error

	return
}

func (obj *_ProductAttributeMgr) GetBatchFromMinimalQuantity(minimalQuantitys []uint32) (results []*ProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("minimal_quantity IN (?)", minimalQuantitys).Find(&results).Error

	return
}

func (obj *_ProductAttributeMgr) GetFromLowStockThreshold(lowStockThreshold int) (results []*ProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("low_stock_threshold = ?", lowStockThreshold).Find(&results).Error

	return
}

func (obj *_ProductAttributeMgr) GetBatchFromLowStockThreshold(lowStockThresholds []int) (results []*ProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("low_stock_threshold IN (?)", lowStockThresholds).Find(&results).Error

	return
}

func (obj *_ProductAttributeMgr) GetFromLowStockAlert(lowStockAlert bool) (results []*ProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("low_stock_alert = ?", lowStockAlert).Find(&results).Error

	return
}

func (obj *_ProductAttributeMgr) GetBatchFromLowStockAlert(lowStockAlerts []bool) (results []*ProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("low_stock_alert IN (?)", lowStockAlerts).Find(&results).Error

	return
}

func (obj *_ProductAttributeMgr) GetFromAvailableDate(availableDate datatypes.Date) (results []*ProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("available_date = ?", availableDate).Find(&results).Error

	return
}

func (obj *_ProductAttributeMgr) GetBatchFromAvailableDate(availableDates []datatypes.Date) (results []*ProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("available_date IN (?)", availableDates).Find(&results).Error

	return
}


func (obj *_ProductAttributeMgr) FetchByPrimaryKey(idProductAttribute uint32) (result ProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_attribute = ?", idProductAttribute).Find(&result).Error

	return
}

func (obj *_ProductAttributeMgr) FetchUniqueIndexByProductDefault(idProduct uint32, defaultOn bool) (result ProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ? AND default_on = ?", idProduct, defaultOn).Find(&result).Error

	return
}

func (obj *_ProductAttributeMgr) FetchIndexByIDProductIDProductAttribute(idProductAttribute uint32, idProduct uint32) (results []*ProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_attribute = ? AND id_product = ?", idProductAttribute, idProduct).Find(&results).Error

	return
}

func (obj *_ProductAttributeMgr) FetchIndexByProductAttributeProduct(idProduct uint32) (results []*ProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ?", idProduct).Find(&results).Error

	return
}

func (obj *_ProductAttributeMgr) FetchIndexByReference(reference string) (results []*ProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reference = ?", reference).Find(&results).Error

	return
}

func (obj *_ProductAttributeMgr) FetchIndexBySupplierReference(supplierReference string) (results []*ProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("supplier_reference = ?", supplierReference).Find(&results).Error

	return
}
