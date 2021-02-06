package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _StockMgr struct {
	*_BaseMgr
}

// StockMgr open func
func StockMgr(db *gorm.DB) *_StockMgr {
	if db == nil {
		panic(fmt.Errorf("StockMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_StockMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_stock"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_StockMgr) GetTableName() string {
	return "eg_stock"
}

// Get 获取
func (obj *_StockMgr) Get() (result Stock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_StockMgr) Gets() (results []*Stock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDStock id_stock获取
func (obj *_StockMgr) WithIDStock(idStock uint32) Option {
	return optionFunc(func(o *options) { o.query["id_stock"] = idStock })
}

// WithIDWarehouse id_warehouse获取
func (obj *_StockMgr) WithIDWarehouse(idWarehouse uint32) Option {
	return optionFunc(func(o *options) { o.query["id_warehouse"] = idWarehouse })
}

// WithIDProduct id_product获取
func (obj *_StockMgr) WithIDProduct(idProduct uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product"] = idProduct })
}

// WithIDProductAttribute id_product_attribute获取
func (obj *_StockMgr) WithIDProductAttribute(idProductAttribute uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product_attribute"] = idProductAttribute })
}

// WithReference reference获取
func (obj *_StockMgr) WithReference(reference string) Option {
	return optionFunc(func(o *options) { o.query["reference"] = reference })
}

// WithEan13 ean13获取
func (obj *_StockMgr) WithEan13(ean13 string) Option {
	return optionFunc(func(o *options) { o.query["ean13"] = ean13 })
}

// WithIsbn isbn获取
func (obj *_StockMgr) WithIsbn(isbn string) Option {
	return optionFunc(func(o *options) { o.query["isbn"] = isbn })
}

// WithUpc upc获取
func (obj *_StockMgr) WithUpc(upc string) Option {
	return optionFunc(func(o *options) { o.query["upc"] = upc })
}

// WithMpn mpn获取
func (obj *_StockMgr) WithMpn(mpn string) Option {
	return optionFunc(func(o *options) { o.query["mpn"] = mpn })
}

// WithPhysicalQuantity physical_quantity获取
func (obj *_StockMgr) WithPhysicalQuantity(physicalQuantity uint32) Option {
	return optionFunc(func(o *options) { o.query["physical_quantity"] = physicalQuantity })
}

// WithUsableQuantity usable_quantity获取
func (obj *_StockMgr) WithUsableQuantity(usableQuantity uint32) Option {
	return optionFunc(func(o *options) { o.query["usable_quantity"] = usableQuantity })
}

// WithPriceTe price_te获取
func (obj *_StockMgr) WithPriceTe(priceTe float64) Option {
	return optionFunc(func(o *options) { o.query["price_te"] = priceTe })
}

// GetByOption 功能选项模式获取
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

// GetByOptions 批量功能选项模式获取
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

//////////////////////////enume case ////////////////////////////////////////////

// GetFromIDStock 通过id_stock获取内容
func (obj *_StockMgr) GetFromIDStock(idStock uint32) (result Stock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_stock = ?", idStock).Find(&result).Error

	return
}

// GetBatchFromIDStock 批量唯一主键查找
func (obj *_StockMgr) GetBatchFromIDStock(idStocks []uint32) (results []*Stock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_stock IN (?)", idStocks).Find(&results).Error

	return
}

// GetFromIDWarehouse 通过id_warehouse获取内容
func (obj *_StockMgr) GetFromIDWarehouse(idWarehouse uint32) (results []*Stock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_warehouse = ?", idWarehouse).Find(&results).Error

	return
}

// GetBatchFromIDWarehouse 批量唯一主键查找
func (obj *_StockMgr) GetBatchFromIDWarehouse(idWarehouses []uint32) (results []*Stock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_warehouse IN (?)", idWarehouses).Find(&results).Error

	return
}

// GetFromIDProduct 通过id_product获取内容
func (obj *_StockMgr) GetFromIDProduct(idProduct uint32) (results []*Stock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ?", idProduct).Find(&results).Error

	return
}

// GetBatchFromIDProduct 批量唯一主键查找
func (obj *_StockMgr) GetBatchFromIDProduct(idProducts []uint32) (results []*Stock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product IN (?)", idProducts).Find(&results).Error

	return
}

// GetFromIDProductAttribute 通过id_product_attribute获取内容
func (obj *_StockMgr) GetFromIDProductAttribute(idProductAttribute uint32) (results []*Stock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_attribute = ?", idProductAttribute).Find(&results).Error

	return
}

// GetBatchFromIDProductAttribute 批量唯一主键查找
func (obj *_StockMgr) GetBatchFromIDProductAttribute(idProductAttributes []uint32) (results []*Stock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_attribute IN (?)", idProductAttributes).Find(&results).Error

	return
}

// GetFromReference 通过reference获取内容
func (obj *_StockMgr) GetFromReference(reference string) (results []*Stock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reference = ?", reference).Find(&results).Error

	return
}

// GetBatchFromReference 批量唯一主键查找
func (obj *_StockMgr) GetBatchFromReference(references []string) (results []*Stock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reference IN (?)", references).Find(&results).Error

	return
}

// GetFromEan13 通过ean13获取内容
func (obj *_StockMgr) GetFromEan13(ean13 string) (results []*Stock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ean13 = ?", ean13).Find(&results).Error

	return
}

// GetBatchFromEan13 批量唯一主键查找
func (obj *_StockMgr) GetBatchFromEan13(ean13s []string) (results []*Stock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ean13 IN (?)", ean13s).Find(&results).Error

	return
}

// GetFromIsbn 通过isbn获取内容
func (obj *_StockMgr) GetFromIsbn(isbn string) (results []*Stock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("isbn = ?", isbn).Find(&results).Error

	return
}

// GetBatchFromIsbn 批量唯一主键查找
func (obj *_StockMgr) GetBatchFromIsbn(isbns []string) (results []*Stock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("isbn IN (?)", isbns).Find(&results).Error

	return
}

// GetFromUpc 通过upc获取内容
func (obj *_StockMgr) GetFromUpc(upc string) (results []*Stock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("upc = ?", upc).Find(&results).Error

	return
}

// GetBatchFromUpc 批量唯一主键查找
func (obj *_StockMgr) GetBatchFromUpc(upcs []string) (results []*Stock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("upc IN (?)", upcs).Find(&results).Error

	return
}

// GetFromMpn 通过mpn获取内容
func (obj *_StockMgr) GetFromMpn(mpn string) (results []*Stock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("mpn = ?", mpn).Find(&results).Error

	return
}

// GetBatchFromMpn 批量唯一主键查找
func (obj *_StockMgr) GetBatchFromMpn(mpns []string) (results []*Stock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("mpn IN (?)", mpns).Find(&results).Error

	return
}

// GetFromPhysicalQuantity 通过physical_quantity获取内容
func (obj *_StockMgr) GetFromPhysicalQuantity(physicalQuantity uint32) (results []*Stock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("physical_quantity = ?", physicalQuantity).Find(&results).Error

	return
}

// GetBatchFromPhysicalQuantity 批量唯一主键查找
func (obj *_StockMgr) GetBatchFromPhysicalQuantity(physicalQuantitys []uint32) (results []*Stock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("physical_quantity IN (?)", physicalQuantitys).Find(&results).Error

	return
}

// GetFromUsableQuantity 通过usable_quantity获取内容
func (obj *_StockMgr) GetFromUsableQuantity(usableQuantity uint32) (results []*Stock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("usable_quantity = ?", usableQuantity).Find(&results).Error

	return
}

// GetBatchFromUsableQuantity 批量唯一主键查找
func (obj *_StockMgr) GetBatchFromUsableQuantity(usableQuantitys []uint32) (results []*Stock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("usable_quantity IN (?)", usableQuantitys).Find(&results).Error

	return
}

// GetFromPriceTe 通过price_te获取内容
func (obj *_StockMgr) GetFromPriceTe(priceTe float64) (results []*Stock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("price_te = ?", priceTe).Find(&results).Error

	return
}

// GetBatchFromPriceTe 批量唯一主键查找
func (obj *_StockMgr) GetBatchFromPriceTe(priceTes []float64) (results []*Stock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("price_te IN (?)", priceTes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_StockMgr) FetchByPrimaryKey(idStock uint32) (result Stock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_stock = ?", idStock).Find(&result).Error

	return
}

// FetchIndexByIDWarehouse  获取多个内容
func (obj *_StockMgr) FetchIndexByIDWarehouse(idWarehouse uint32) (results []*Stock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_warehouse = ?", idWarehouse).Find(&results).Error

	return
}

// FetchIndexByIDProduct  获取多个内容
func (obj *_StockMgr) FetchIndexByIDProduct(idProduct uint32) (results []*Stock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ?", idProduct).Find(&results).Error

	return
}

// FetchIndexByIDProductAttribute  获取多个内容
func (obj *_StockMgr) FetchIndexByIDProductAttribute(idProductAttribute uint32) (results []*Stock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_attribute = ?", idProductAttribute).Find(&results).Error

	return
}
