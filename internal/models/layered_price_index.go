package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _LayeredPriceIndexMgr struct {
	*_BaseMgr
}

// LayeredPriceIndexMgr open func
func LayeredPriceIndexMgr(db *gorm.DB) *_LayeredPriceIndexMgr {
	if db == nil {
		panic(fmt.Errorf("LayeredPriceIndexMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_LayeredPriceIndexMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_layered_price_index"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_LayeredPriceIndexMgr) GetTableName() string {
	return "eg_layered_price_index"
}

// Get 获取
func (obj *_LayeredPriceIndexMgr) Get() (result LayeredPriceIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_LayeredPriceIndexMgr) Gets() (results []*LayeredPriceIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDProduct id_product获取
func (obj *_LayeredPriceIndexMgr) WithIDProduct(idProduct int) Option {
	return optionFunc(func(o *options) { o.query["id_product"] = idProduct })
}

// WithIDCurrency id_currency获取
func (obj *_LayeredPriceIndexMgr) WithIDCurrency(idCurrency int) Option {
	return optionFunc(func(o *options) { o.query["id_currency"] = idCurrency })
}

// WithIDShop id_shop获取
func (obj *_LayeredPriceIndexMgr) WithIDShop(idShop int) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

// WithPriceMin price_min获取
func (obj *_LayeredPriceIndexMgr) WithPriceMin(priceMin float64) Option {
	return optionFunc(func(o *options) { o.query["price_min"] = priceMin })
}

// WithPriceMax price_max获取
func (obj *_LayeredPriceIndexMgr) WithPriceMax(priceMax float64) Option {
	return optionFunc(func(o *options) { o.query["price_max"] = priceMax })
}

// WithIDCountry id_country获取
func (obj *_LayeredPriceIndexMgr) WithIDCountry(idCountry int) Option {
	return optionFunc(func(o *options) { o.query["id_country"] = idCountry })
}

// GetByOption 功能选项模式获取
func (obj *_LayeredPriceIndexMgr) GetByOption(opts ...Option) (result LayeredPriceIndex, err error) {
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
func (obj *_LayeredPriceIndexMgr) GetByOptions(opts ...Option) (results []*LayeredPriceIndex, err error) {
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
func (obj *_LayeredPriceIndexMgr) GetFromIDProduct(idProduct int) (results []*LayeredPriceIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ?", idProduct).Find(&results).Error

	return
}

// GetBatchFromIDProduct 批量唯一主键查找
func (obj *_LayeredPriceIndexMgr) GetBatchFromIDProduct(idProducts []int) (results []*LayeredPriceIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product IN (?)", idProducts).Find(&results).Error

	return
}

// GetFromIDCurrency 通过id_currency获取内容
func (obj *_LayeredPriceIndexMgr) GetFromIDCurrency(idCurrency int) (results []*LayeredPriceIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_currency = ?", idCurrency).Find(&results).Error

	return
}

// GetBatchFromIDCurrency 批量唯一主键查找
func (obj *_LayeredPriceIndexMgr) GetBatchFromIDCurrency(idCurrencys []int) (results []*LayeredPriceIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_currency IN (?)", idCurrencys).Find(&results).Error

	return
}

// GetFromIDShop 通过id_shop获取内容
func (obj *_LayeredPriceIndexMgr) GetFromIDShop(idShop int) (results []*LayeredPriceIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

// GetBatchFromIDShop 批量唯一主键查找
func (obj *_LayeredPriceIndexMgr) GetBatchFromIDShop(idShops []int) (results []*LayeredPriceIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

// GetFromPriceMin 通过price_min获取内容
func (obj *_LayeredPriceIndexMgr) GetFromPriceMin(priceMin float64) (results []*LayeredPriceIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("price_min = ?", priceMin).Find(&results).Error

	return
}

// GetBatchFromPriceMin 批量唯一主键查找
func (obj *_LayeredPriceIndexMgr) GetBatchFromPriceMin(priceMins []float64) (results []*LayeredPriceIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("price_min IN (?)", priceMins).Find(&results).Error

	return
}

// GetFromPriceMax 通过price_max获取内容
func (obj *_LayeredPriceIndexMgr) GetFromPriceMax(priceMax float64) (results []*LayeredPriceIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("price_max = ?", priceMax).Find(&results).Error

	return
}

// GetBatchFromPriceMax 批量唯一主键查找
func (obj *_LayeredPriceIndexMgr) GetBatchFromPriceMax(priceMaxs []float64) (results []*LayeredPriceIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("price_max IN (?)", priceMaxs).Find(&results).Error

	return
}

// GetFromIDCountry 通过id_country获取内容
func (obj *_LayeredPriceIndexMgr) GetFromIDCountry(idCountry int) (results []*LayeredPriceIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_country = ?", idCountry).Find(&results).Error

	return
}

// GetBatchFromIDCountry 批量唯一主键查找
func (obj *_LayeredPriceIndexMgr) GetBatchFromIDCountry(idCountrys []int) (results []*LayeredPriceIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_country IN (?)", idCountrys).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_LayeredPriceIndexMgr) FetchByPrimaryKey(idProduct int, idCurrency int, idShop int, idCountry int) (result LayeredPriceIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ? AND id_currency = ? AND id_shop = ? AND id_country = ?", idProduct, idCurrency, idShop, idCountry).Find(&result).Error

	return
}

// FetchIndexByIDCurrency  获取多个内容
func (obj *_LayeredPriceIndexMgr) FetchIndexByIDCurrency(idCurrency int) (results []*LayeredPriceIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_currency = ?", idCurrency).Find(&results).Error

	return
}

// FetchIndexByPriceMin  获取多个内容
func (obj *_LayeredPriceIndexMgr) FetchIndexByPriceMin(priceMin float64) (results []*LayeredPriceIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("price_min = ?", priceMin).Find(&results).Error

	return
}

// FetchIndexByPriceMax  获取多个内容
func (obj *_LayeredPriceIndexMgr) FetchIndexByPriceMax(priceMax float64) (results []*LayeredPriceIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("price_max = ?", priceMax).Find(&results).Error

	return
}
