package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _StockMgr struct {
	*_BaseMgr
}

func StockMgr(db *gorm.DB) *_StockMgr {
	if db == nil {
		panic(fmt.Errorf("StockMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_StockMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_stock"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_StockMgr) GetTableName() string {
	return "ps_stock"
}

func (obj *_StockMgr) Get() (result Stock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_StockMgr) Gets() (results []*Stock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

func (obj *_StockMgr) WithIDStock(idStock uint32) Option {
	return optionFunc(func(o *options) { o.query["id_stock"] = idStock })
}

func (obj *_StockMgr) WithIDWarehouse(idWarehouse uint32) Option {
	return optionFunc(func(o *options) { o.query["id_warehouse"] = idWarehouse })
}

func (obj *_StockMgr) WithIDProduct(idProduct uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product"] = idProduct })
}

func (obj *_StockMgr) WithIDProductAttribute(idProductAttribute uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product_attribute"] = idProductAttribute })
}

func (obj *_StockMgr) WithReference(reference string) Option {
	return optionFunc(func(o *options) { o.query["reference"] = reference })
}

func (obj *_StockMgr) WithEan13(ean13 string) Option {
	return optionFunc(func(o *options) { o.query["ean13"] = ean13 })
}

func (obj *_StockMgr) WithIsbn(isbn string) Option {
	return optionFunc(func(o *options) { o.query["isbn"] = isbn })
}

func (obj *_StockMgr) WithUpc(upc string) Option {
	return optionFunc(func(o *options) { o.query["upc"] = upc })
}

func (obj *_StockMgr) WithMpn(mpn string) Option {
	return optionFunc(func(o *options) { o.query["mpn"] = mpn })
}

func (obj *_StockMgr) WithPhysicalQuantity(physicalQuantity uint32) Option {
	return optionFunc(func(o *options) { o.query["physical_quantity"] = physicalQuantity })
}

func (obj *_StockMgr) WithUsableQuantity(usableQuantity uint32) Option {
	return optionFunc(func(o *options) { o.query["usable_quantity"] = usableQuantity })
}

func (obj *_StockMgr) WithPriceTe(priceTe float64) Option {
	return optionFunc(func(o *options) { o.query["price_te"] = priceTe })
}

func (obj *_StockMgr) GetByOption(opts ...Option) (result Stock, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_StockMgr) GetByOptions(opts ...Option) (results []*Stock, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

func (obj *_StockMgr) GetFromIDStock(idStock uint32) (result Stock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_stock = ?", idStock).Find(&result).Error

	return
}

func (obj *_StockMgr) GetBatchFromIDStock(idStocks []uint32) (results []*Stock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_stock IN (?)", idStocks).Find(&results).Error

	return
}

func (obj *_StockMgr) GetFromIDWarehouse(idWarehouse uint32) (results []*Stock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_warehouse = ?", idWarehouse).Find(&results).Error

	return
}

func (obj *_StockMgr) GetBatchFromIDWarehouse(idWarehouses []uint32) (results []*Stock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_warehouse IN (?)", idWarehouses).Find(&results).Error

	return
}

func (obj *_StockMgr) GetFromIDProduct(idProduct uint32) (results []*Stock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ?", idProduct).Find(&results).Error

	return
}

func (obj *_StockMgr) GetBatchFromIDProduct(idProducts []uint32) (results []*Stock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product IN (?)", idProducts).Find(&results).Error

	return
}

func (obj *_StockMgr) GetFromIDProductAttribute(idProductAttribute uint32) (results []*Stock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_attribute = ?", idProductAttribute).Find(&results).Error

	return
}

func (obj *_StockMgr) GetBatchFromIDProductAttribute(idProductAttributes []uint32) (results []*Stock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_attribute IN (?)", idProductAttributes).Find(&results).Error

	return
}

func (obj *_StockMgr) GetFromReference(reference string) (results []*Stock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reference = ?", reference).Find(&results).Error

	return
}

func (obj *_StockMgr) GetBatchFromReference(references []string) (results []*Stock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reference IN (?)", references).Find(&results).Error

	return
}

func (obj *_StockMgr) GetFromEan13(ean13 string) (results []*Stock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ean13 = ?", ean13).Find(&results).Error

	return
}

func (obj *_StockMgr) GetBatchFromEan13(ean13s []string) (results []*Stock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ean13 IN (?)", ean13s).Find(&results).Error

	return
}

func (obj *_StockMgr) GetFromIsbn(isbn string) (results []*Stock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("isbn = ?", isbn).Find(&results).Error

	return
}

func (obj *_StockMgr) GetBatchFromIsbn(isbns []string) (results []*Stock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("isbn IN (?)", isbns).Find(&results).Error

	return
}

func (obj *_StockMgr) GetFromUpc(upc string) (results []*Stock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("upc = ?", upc).Find(&results).Error

	return
}

func (obj *_StockMgr) GetBatchFromUpc(upcs []string) (results []*Stock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("upc IN (?)", upcs).Find(&results).Error

	return
}

func (obj *_StockMgr) GetFromMpn(mpn string) (results []*Stock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("mpn = ?", mpn).Find(&results).Error

	return
}

func (obj *_StockMgr) GetBatchFromMpn(mpns []string) (results []*Stock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("mpn IN (?)", mpns).Find(&results).Error

	return
}

func (obj *_StockMgr) GetFromPhysicalQuantity(physicalQuantity uint32) (results []*Stock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("physical_quantity = ?", physicalQuantity).Find(&results).Error

	return
}

func (obj *_StockMgr) GetBatchFromPhysicalQuantity(physicalQuantitys []uint32) (results []*Stock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("physical_quantity IN (?)", physicalQuantitys).Find(&results).Error

	return
}

func (obj *_StockMgr) GetFromUsableQuantity(usableQuantity uint32) (results []*Stock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("usable_quantity = ?", usableQuantity).Find(&results).Error

	return
}

func (obj *_StockMgr) GetBatchFromUsableQuantity(usableQuantitys []uint32) (results []*Stock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("usable_quantity IN (?)", usableQuantitys).Find(&results).Error

	return
}

func (obj *_StockMgr) GetFromPriceTe(priceTe float64) (results []*Stock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("price_te = ?", priceTe).Find(&results).Error

	return
}

func (obj *_StockMgr) GetBatchFromPriceTe(priceTes []float64) (results []*Stock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("price_te IN (?)", priceTes).Find(&results).Error

	return
}

func (obj *_StockMgr) FetchByPrimaryKey(idStock uint32) (result Stock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_stock = ?", idStock).Find(&result).Error

	return
}

func (obj *_StockMgr) FetchIndexByIDWarehouse(idWarehouse uint32) (results []*Stock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_warehouse = ?", idWarehouse).Find(&results).Error

	return
}

func (obj *_StockMgr) FetchIndexByIDProduct(idProduct uint32) (results []*Stock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ?", idProduct).Find(&results).Error

	return
}

func (obj *_StockMgr) FetchIndexByIDProductAttribute(idProductAttribute uint32) (results []*Stock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_attribute = ?", idProductAttribute).Find(&results).Error

	return
}
