package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _StockAvailableMgr struct {
	*_BaseMgr
}

// StockAvailableMgr open func
func StockAvailableMgr(db *gorm.DB) *_StockAvailableMgr {
	if db == nil {
		panic(fmt.Errorf("StockAvailableMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_StockAvailableMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_stock_available"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_StockAvailableMgr) GetTableName() string {
	return "eg_stock_available"
}

// Get 获取
func (obj *_StockAvailableMgr) Get() (result StockAvailable, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_StockAvailableMgr) Gets() (results []*StockAvailable, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDStockAvailable id_stock_available获取
func (obj *_StockAvailableMgr) WithIDStockAvailable(idStockAvailable uint32) Option {
	return optionFunc(func(o *options) { o.query["id_stock_available"] = idStockAvailable })
}

// WithIDProduct id_product获取
func (obj *_StockAvailableMgr) WithIDProduct(idProduct uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product"] = idProduct })
}

// WithIDProductAttribute id_product_attribute获取
func (obj *_StockAvailableMgr) WithIDProductAttribute(idProductAttribute uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product_attribute"] = idProductAttribute })
}

// WithIDShop id_shop获取
func (obj *_StockAvailableMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

// WithIDShopGroup id_shop_group获取
func (obj *_StockAvailableMgr) WithIDShopGroup(idShopGroup uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop_group"] = idShopGroup })
}

// WithQuantity quantity获取
func (obj *_StockAvailableMgr) WithQuantity(quantity int) Option {
	return optionFunc(func(o *options) { o.query["quantity"] = quantity })
}

// WithPhysicalQuantity physical_quantity获取
func (obj *_StockAvailableMgr) WithPhysicalQuantity(physicalQuantity int) Option {
	return optionFunc(func(o *options) { o.query["physical_quantity"] = physicalQuantity })
}

// WithReservedQuantity reserved_quantity获取
func (obj *_StockAvailableMgr) WithReservedQuantity(reservedQuantity int) Option {
	return optionFunc(func(o *options) { o.query["reserved_quantity"] = reservedQuantity })
}

// WithDependsOnStock depends_on_stock获取
func (obj *_StockAvailableMgr) WithDependsOnStock(dependsOnStock bool) Option {
	return optionFunc(func(o *options) { o.query["depends_on_stock"] = dependsOnStock })
}

// WithOutOfStock out_of_stock获取
func (obj *_StockAvailableMgr) WithOutOfStock(outOfStock bool) Option {
	return optionFunc(func(o *options) { o.query["out_of_stock"] = outOfStock })
}

// WithLocation location获取
func (obj *_StockAvailableMgr) WithLocation(location string) Option {
	return optionFunc(func(o *options) { o.query["location"] = location })
}

// GetByOption 功能选项模式获取
func (obj *_StockAvailableMgr) GetByOption(opts ...Option) (result StockAvailable, err error) {
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
func (obj *_StockAvailableMgr) GetByOptions(opts ...Option) (results []*StockAvailable, err error) {
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

// GetFromIDStockAvailable 通过id_stock_available获取内容
func (obj *_StockAvailableMgr) GetFromIDStockAvailable(idStockAvailable uint32) (result StockAvailable, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_stock_available = ?", idStockAvailable).Find(&result).Error

	return
}

// GetBatchFromIDStockAvailable 批量唯一主键查找
func (obj *_StockAvailableMgr) GetBatchFromIDStockAvailable(idStockAvailables []uint32) (results []*StockAvailable, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_stock_available IN (?)", idStockAvailables).Find(&results).Error

	return
}

// GetFromIDProduct 通过id_product获取内容
func (obj *_StockAvailableMgr) GetFromIDProduct(idProduct uint32) (results []*StockAvailable, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ?", idProduct).Find(&results).Error

	return
}

// GetBatchFromIDProduct 批量唯一主键查找
func (obj *_StockAvailableMgr) GetBatchFromIDProduct(idProducts []uint32) (results []*StockAvailable, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product IN (?)", idProducts).Find(&results).Error

	return
}

// GetFromIDProductAttribute 通过id_product_attribute获取内容
func (obj *_StockAvailableMgr) GetFromIDProductAttribute(idProductAttribute uint32) (results []*StockAvailable, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_attribute = ?", idProductAttribute).Find(&results).Error

	return
}

// GetBatchFromIDProductAttribute 批量唯一主键查找
func (obj *_StockAvailableMgr) GetBatchFromIDProductAttribute(idProductAttributes []uint32) (results []*StockAvailable, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_attribute IN (?)", idProductAttributes).Find(&results).Error

	return
}

// GetFromIDShop 通过id_shop获取内容
func (obj *_StockAvailableMgr) GetFromIDShop(idShop uint32) (results []*StockAvailable, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

// GetBatchFromIDShop 批量唯一主键查找
func (obj *_StockAvailableMgr) GetBatchFromIDShop(idShops []uint32) (results []*StockAvailable, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

// GetFromIDShopGroup 通过id_shop_group获取内容
func (obj *_StockAvailableMgr) GetFromIDShopGroup(idShopGroup uint32) (results []*StockAvailable, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_group = ?", idShopGroup).Find(&results).Error

	return
}

// GetBatchFromIDShopGroup 批量唯一主键查找
func (obj *_StockAvailableMgr) GetBatchFromIDShopGroup(idShopGroups []uint32) (results []*StockAvailable, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_group IN (?)", idShopGroups).Find(&results).Error

	return
}

// GetFromQuantity 通过quantity获取内容
func (obj *_StockAvailableMgr) GetFromQuantity(quantity int) (results []*StockAvailable, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("quantity = ?", quantity).Find(&results).Error

	return
}

// GetBatchFromQuantity 批量唯一主键查找
func (obj *_StockAvailableMgr) GetBatchFromQuantity(quantitys []int) (results []*StockAvailable, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("quantity IN (?)", quantitys).Find(&results).Error

	return
}

// GetFromPhysicalQuantity 通过physical_quantity获取内容
func (obj *_StockAvailableMgr) GetFromPhysicalQuantity(physicalQuantity int) (results []*StockAvailable, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("physical_quantity = ?", physicalQuantity).Find(&results).Error

	return
}

// GetBatchFromPhysicalQuantity 批量唯一主键查找
func (obj *_StockAvailableMgr) GetBatchFromPhysicalQuantity(physicalQuantitys []int) (results []*StockAvailable, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("physical_quantity IN (?)", physicalQuantitys).Find(&results).Error

	return
}

// GetFromReservedQuantity 通过reserved_quantity获取内容
func (obj *_StockAvailableMgr) GetFromReservedQuantity(reservedQuantity int) (results []*StockAvailable, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reserved_quantity = ?", reservedQuantity).Find(&results).Error

	return
}

// GetBatchFromReservedQuantity 批量唯一主键查找
func (obj *_StockAvailableMgr) GetBatchFromReservedQuantity(reservedQuantitys []int) (results []*StockAvailable, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reserved_quantity IN (?)", reservedQuantitys).Find(&results).Error

	return
}

// GetFromDependsOnStock 通过depends_on_stock获取内容
func (obj *_StockAvailableMgr) GetFromDependsOnStock(dependsOnStock bool) (results []*StockAvailable, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("depends_on_stock = ?", dependsOnStock).Find(&results).Error

	return
}

// GetBatchFromDependsOnStock 批量唯一主键查找
func (obj *_StockAvailableMgr) GetBatchFromDependsOnStock(dependsOnStocks []bool) (results []*StockAvailable, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("depends_on_stock IN (?)", dependsOnStocks).Find(&results).Error

	return
}

// GetFromOutOfStock 通过out_of_stock获取内容
func (obj *_StockAvailableMgr) GetFromOutOfStock(outOfStock bool) (results []*StockAvailable, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("out_of_stock = ?", outOfStock).Find(&results).Error

	return
}

// GetBatchFromOutOfStock 批量唯一主键查找
func (obj *_StockAvailableMgr) GetBatchFromOutOfStock(outOfStocks []bool) (results []*StockAvailable, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("out_of_stock IN (?)", outOfStocks).Find(&results).Error

	return
}

// GetFromLocation 通过location获取内容
func (obj *_StockAvailableMgr) GetFromLocation(location string) (results []*StockAvailable, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("location = ?", location).Find(&results).Error

	return
}

// GetBatchFromLocation 批量唯一主键查找
func (obj *_StockAvailableMgr) GetBatchFromLocation(locations []string) (results []*StockAvailable, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("location IN (?)", locations).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_StockAvailableMgr) FetchByPrimaryKey(idStockAvailable uint32) (result StockAvailable, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_stock_available = ?", idStockAvailable).Find(&result).Error

	return
}

// FetchUniqueIndexByProductSQLstock primay or index 获取唯一内容
func (obj *_StockAvailableMgr) FetchUniqueIndexByProductSQLstock(idProduct uint32, idProductAttribute uint32, idShop uint32, idShopGroup uint32) (result StockAvailable, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ? AND id_product_attribute = ? AND id_shop = ? AND id_shop_group = ?", idProduct, idProductAttribute, idShop, idShopGroup).Find(&result).Error

	return
}

// FetchIndexByIDProduct  获取多个内容
func (obj *_StockAvailableMgr) FetchIndexByIDProduct(idProduct uint32) (results []*StockAvailable, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ?", idProduct).Find(&results).Error

	return
}

// FetchIndexByIDProductAttribute  获取多个内容
func (obj *_StockAvailableMgr) FetchIndexByIDProductAttribute(idProductAttribute uint32) (results []*StockAvailable, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_attribute = ?", idProductAttribute).Find(&results).Error

	return
}

// FetchIndexByIDShop  获取多个内容
func (obj *_StockAvailableMgr) FetchIndexByIDShop(idShop uint32) (results []*StockAvailable, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

// FetchIndexByIDShopGroup  获取多个内容
func (obj *_StockAvailableMgr) FetchIndexByIDShopGroup(idShopGroup uint32) (results []*StockAvailable, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_group = ?", idShopGroup).Find(&results).Error

	return
}
