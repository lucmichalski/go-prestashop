package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _SpecificPriceMgr struct {
	*_BaseMgr
}

func SpecificPriceMgr(db *gorm.DB) *_SpecificPriceMgr {
	if db == nil {
		panic(fmt.Errorf("SpecificPriceMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SpecificPriceMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_specific_price"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_SpecificPriceMgr) GetTableName() string {
	return "ps_specific_price"
}

func (obj *_SpecificPriceMgr) Get() (result SpecificPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_SpecificPriceMgr) Gets() (results []*SpecificPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_SpecificPriceMgr) WithIDSpecificPrice(idSpecificPrice uint32) Option {
	return optionFunc(func(o *options) { o.query["id_specific_price"] = idSpecificPrice })
}

func (obj *_SpecificPriceMgr) WithIDSpecificPriceRule(idSpecificPriceRule uint32) Option {
	return optionFunc(func(o *options) { o.query["id_specific_price_rule"] = idSpecificPriceRule })
}

func (obj *_SpecificPriceMgr) WithIDCart(idCart uint32) Option {
	return optionFunc(func(o *options) { o.query["id_cart"] = idCart })
}

func (obj *_SpecificPriceMgr) WithIDProduct(idProduct uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product"] = idProduct })
}

func (obj *_SpecificPriceMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

func (obj *_SpecificPriceMgr) WithIDShopGroup(idShopGroup uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop_group"] = idShopGroup })
}

func (obj *_SpecificPriceMgr) WithIDCurrency(idCurrency uint32) Option {
	return optionFunc(func(o *options) { o.query["id_currency"] = idCurrency })
}

func (obj *_SpecificPriceMgr) WithIDCountry(idCountry uint32) Option {
	return optionFunc(func(o *options) { o.query["id_country"] = idCountry })
}

func (obj *_SpecificPriceMgr) WithIDGroup(idGroup uint32) Option {
	return optionFunc(func(o *options) { o.query["id_group"] = idGroup })
}

func (obj *_SpecificPriceMgr) WithIDCustomer(idCustomer uint32) Option {
	return optionFunc(func(o *options) { o.query["id_customer"] = idCustomer })
}

func (obj *_SpecificPriceMgr) WithIDProductAttribute(idProductAttribute uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product_attribute"] = idProductAttribute })
}

func (obj *_SpecificPriceMgr) WithPrice(price float64) Option {
	return optionFunc(func(o *options) { o.query["price"] = price })
}

func (obj *_SpecificPriceMgr) WithFromQuantity(fromQuantity string) Option {
	return optionFunc(func(o *options) { o.query["from_quantity"] = fromQuantity })
}

func (obj *_SpecificPriceMgr) WithReduction(reduction float64) Option {
	return optionFunc(func(o *options) { o.query["reduction"] = reduction })
}

func (obj *_SpecificPriceMgr) WithReductionTax(reductionTax bool) Option {
	return optionFunc(func(o *options) { o.query["reduction_tax"] = reductionTax })
}

func (obj *_SpecificPriceMgr) WithReductionType(reductionType string) Option {
	return optionFunc(func(o *options) { o.query["reduction_type"] = reductionType })
}

func (obj *_SpecificPriceMgr) WithFrom(from time.Time) Option {
	return optionFunc(func(o *options) { o.query["from"] = from })
}

func (obj *_SpecificPriceMgr) WithTo(to time.Time) Option {
	return optionFunc(func(o *options) { o.query["to"] = to })
}

func (obj *_SpecificPriceMgr) GetByOption(opts ...Option) (result SpecificPrice, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_SpecificPriceMgr) GetByOptions(opts ...Option) (results []*SpecificPrice, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_SpecificPriceMgr) GetFromIDSpecificPrice(idSpecificPrice uint32) (result SpecificPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_specific_price = ?", idSpecificPrice).Find(&result).Error

	return
}

func (obj *_SpecificPriceMgr) GetBatchFromIDSpecificPrice(idSpecificPrices []uint32) (results []*SpecificPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_specific_price IN (?)", idSpecificPrices).Find(&results).Error

	return
}

func (obj *_SpecificPriceMgr) GetFromIDSpecificPriceRule(idSpecificPriceRule uint32) (results []*SpecificPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_specific_price_rule = ?", idSpecificPriceRule).Find(&results).Error

	return
}

func (obj *_SpecificPriceMgr) GetBatchFromIDSpecificPriceRule(idSpecificPriceRules []uint32) (results []*SpecificPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_specific_price_rule IN (?)", idSpecificPriceRules).Find(&results).Error

	return
}

func (obj *_SpecificPriceMgr) GetFromIDCart(idCart uint32) (results []*SpecificPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart = ?", idCart).Find(&results).Error

	return
}

func (obj *_SpecificPriceMgr) GetBatchFromIDCart(idCarts []uint32) (results []*SpecificPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart IN (?)", idCarts).Find(&results).Error

	return
}

func (obj *_SpecificPriceMgr) GetFromIDProduct(idProduct uint32) (results []*SpecificPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ?", idProduct).Find(&results).Error

	return
}

func (obj *_SpecificPriceMgr) GetBatchFromIDProduct(idProducts []uint32) (results []*SpecificPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product IN (?)", idProducts).Find(&results).Error

	return
}

func (obj *_SpecificPriceMgr) GetFromIDShop(idShop uint32) (results []*SpecificPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

func (obj *_SpecificPriceMgr) GetBatchFromIDShop(idShops []uint32) (results []*SpecificPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

func (obj *_SpecificPriceMgr) GetFromIDShopGroup(idShopGroup uint32) (results []*SpecificPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_group = ?", idShopGroup).Find(&results).Error

	return
}

func (obj *_SpecificPriceMgr) GetBatchFromIDShopGroup(idShopGroups []uint32) (results []*SpecificPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_group IN (?)", idShopGroups).Find(&results).Error

	return
}

func (obj *_SpecificPriceMgr) GetFromIDCurrency(idCurrency uint32) (results []*SpecificPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_currency = ?", idCurrency).Find(&results).Error

	return
}

func (obj *_SpecificPriceMgr) GetBatchFromIDCurrency(idCurrencys []uint32) (results []*SpecificPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_currency IN (?)", idCurrencys).Find(&results).Error

	return
}

func (obj *_SpecificPriceMgr) GetFromIDCountry(idCountry uint32) (results []*SpecificPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_country = ?", idCountry).Find(&results).Error

	return
}

func (obj *_SpecificPriceMgr) GetBatchFromIDCountry(idCountrys []uint32) (results []*SpecificPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_country IN (?)", idCountrys).Find(&results).Error

	return
}

func (obj *_SpecificPriceMgr) GetFromIDGroup(idGroup uint32) (results []*SpecificPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_group = ?", idGroup).Find(&results).Error

	return
}

func (obj *_SpecificPriceMgr) GetBatchFromIDGroup(idGroups []uint32) (results []*SpecificPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_group IN (?)", idGroups).Find(&results).Error

	return
}

func (obj *_SpecificPriceMgr) GetFromIDCustomer(idCustomer uint32) (results []*SpecificPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer = ?", idCustomer).Find(&results).Error

	return
}

func (obj *_SpecificPriceMgr) GetBatchFromIDCustomer(idCustomers []uint32) (results []*SpecificPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer IN (?)", idCustomers).Find(&results).Error

	return
}

func (obj *_SpecificPriceMgr) GetFromIDProductAttribute(idProductAttribute uint32) (results []*SpecificPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_attribute = ?", idProductAttribute).Find(&results).Error

	return
}

func (obj *_SpecificPriceMgr) GetBatchFromIDProductAttribute(idProductAttributes []uint32) (results []*SpecificPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_attribute IN (?)", idProductAttributes).Find(&results).Error

	return
}

func (obj *_SpecificPriceMgr) GetFromPrice(price float64) (results []*SpecificPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("price = ?", price).Find(&results).Error

	return
}

func (obj *_SpecificPriceMgr) GetBatchFromPrice(prices []float64) (results []*SpecificPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("price IN (?)", prices).Find(&results).Error

	return
}

func (obj *_SpecificPriceMgr) GetFromFromQuantity(fromQuantity string) (results []*SpecificPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("from_quantity = ?", fromQuantity).Find(&results).Error

	return
}

func (obj *_SpecificPriceMgr) GetBatchFromFromQuantity(fromQuantitys []string) (results []*SpecificPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("from_quantity IN (?)", fromQuantitys).Find(&results).Error

	return
}

func (obj *_SpecificPriceMgr) GetFromReduction(reduction float64) (results []*SpecificPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reduction = ?", reduction).Find(&results).Error

	return
}

func (obj *_SpecificPriceMgr) GetBatchFromReduction(reductions []float64) (results []*SpecificPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reduction IN (?)", reductions).Find(&results).Error

	return
}

func (obj *_SpecificPriceMgr) GetFromReductionTax(reductionTax bool) (results []*SpecificPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reduction_tax = ?", reductionTax).Find(&results).Error

	return
}

func (obj *_SpecificPriceMgr) GetBatchFromReductionTax(reductionTaxs []bool) (results []*SpecificPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reduction_tax IN (?)", reductionTaxs).Find(&results).Error

	return
}

func (obj *_SpecificPriceMgr) GetFromReductionType(reductionType string) (results []*SpecificPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reduction_type = ?", reductionType).Find(&results).Error

	return
}

func (obj *_SpecificPriceMgr) GetBatchFromReductionType(reductionTypes []string) (results []*SpecificPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reduction_type IN (?)", reductionTypes).Find(&results).Error

	return
}

func (obj *_SpecificPriceMgr) GetFromFrom(from time.Time) (results []*SpecificPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("from = ?", from).Find(&results).Error

	return
}

func (obj *_SpecificPriceMgr) GetBatchFromFrom(froms []time.Time) (results []*SpecificPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("from IN (?)", froms).Find(&results).Error

	return
}

func (obj *_SpecificPriceMgr) GetFromTo(to time.Time) (results []*SpecificPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("to = ?", to).Find(&results).Error

	return
}

func (obj *_SpecificPriceMgr) GetBatchFromTo(tos []time.Time) (results []*SpecificPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("to IN (?)", tos).Find(&results).Error

	return
}


func (obj *_SpecificPriceMgr) FetchByPrimaryKey(idSpecificPrice uint32) (result SpecificPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_specific_price = ?", idSpecificPrice).Find(&result).Error

	return
}

func (obj *_SpecificPriceMgr) FetchUniqueIndexByIDProduct2(idSpecificPriceRule uint32, idCart uint32, idProduct uint32, idShop uint32, idShopGroup uint32, idCurrency uint32, idCountry uint32, idGroup uint32, idCustomer uint32, idProductAttribute uint32, fromQuantity string, from time.Time, to time.Time) (result SpecificPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_specific_price_rule = ? AND id_cart = ? AND id_product = ? AND id_shop = ? AND id_shop_group = ? AND id_currency = ? AND id_country = ? AND id_group = ? AND id_customer = ? AND id_product_attribute = ? AND from_quantity = ? AND from = ? AND to = ?", idSpecificPriceRule, idCart, idProduct, idShop, idShopGroup, idCurrency, idCountry, idGroup, idCustomer, idProductAttribute, fromQuantity, from, to).Find(&result).Error

	return
}

func (obj *_SpecificPriceMgr) FetchIndexByIDSpecificPriceRule(idSpecificPriceRule uint32) (results []*SpecificPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_specific_price_rule = ?", idSpecificPriceRule).Find(&results).Error

	return
}

func (obj *_SpecificPriceMgr) FetchIndexByIDCart(idCart uint32) (results []*SpecificPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart = ?", idCart).Find(&results).Error

	return
}

func (obj *_SpecificPriceMgr) FetchIndexByIDProduct(idProduct uint32, idShop uint32, idCurrency uint32, idCountry uint32, idGroup uint32, idCustomer uint32, fromQuantity string, from time.Time, to time.Time) (results []*SpecificPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ? AND id_shop = ? AND id_currency = ? AND id_country = ? AND id_group = ? AND id_customer = ? AND from_quantity = ? AND from = ? AND to = ?", idProduct, idShop, idCurrency, idCountry, idGroup, idCustomer, fromQuantity, from, to).Find(&results).Error

	return
}

func (obj *_SpecificPriceMgr) FetchIndexByIDShop(idShop uint32) (results []*SpecificPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

func (obj *_SpecificPriceMgr) FetchIndexByIDCustomer(idCustomer uint32) (results []*SpecificPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer = ?", idCustomer).Find(&results).Error

	return
}

func (obj *_SpecificPriceMgr) FetchIndexByIDProductAttribute(idProductAttribute uint32) (results []*SpecificPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_attribute = ?", idProductAttribute).Find(&results).Error

	return
}

func (obj *_SpecificPriceMgr) FetchIndexByFromQuantity(fromQuantity string) (results []*SpecificPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("from_quantity = ?", fromQuantity).Find(&results).Error

	return
}

func (obj *_SpecificPriceMgr) FetchIndexByFrom(from time.Time) (results []*SpecificPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("from = ?", from).Find(&results).Error

	return
}

func (obj *_SpecificPriceMgr) FetchIndexByTo(to time.Time) (results []*SpecificPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("to = ?", to).Find(&results).Error

	return
}
