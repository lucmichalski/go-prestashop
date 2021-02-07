package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _CurrencyShopMgr struct {
	*_BaseMgr
}

func CurrencyShopMgr(db *gorm.DB) *_CurrencyShopMgr {
	if db == nil {
		panic(fmt.Errorf("CurrencyShopMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CurrencyShopMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_currency_shop"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_CurrencyShopMgr) GetTableName() string {
	return "ps_currency_shop"
}

func (obj *_CurrencyShopMgr) Get() (result CurrencyShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_CurrencyShopMgr) Gets() (results []*CurrencyShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

func (obj *_CurrencyShopMgr) WithIDCurrency(idCurrency uint32) Option {
	return optionFunc(func(o *options) { o.query["id_currency"] = idCurrency })
}

func (obj *_CurrencyShopMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

func (obj *_CurrencyShopMgr) WithConversionRate(conversionRate float64) Option {
	return optionFunc(func(o *options) { o.query["conversion_rate"] = conversionRate })
}

func (obj *_CurrencyShopMgr) GetByOption(opts ...Option) (result CurrencyShop, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_CurrencyShopMgr) GetByOptions(opts ...Option) (results []*CurrencyShop, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

func (obj *_CurrencyShopMgr) GetFromIDCurrency(idCurrency uint32) (results []*CurrencyShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_currency = ?", idCurrency).Find(&results).Error

	return
}

func (obj *_CurrencyShopMgr) GetBatchFromIDCurrency(idCurrencys []uint32) (results []*CurrencyShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_currency IN (?)", idCurrencys).Find(&results).Error

	return
}

func (obj *_CurrencyShopMgr) GetFromIDShop(idShop uint32) (results []*CurrencyShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

func (obj *_CurrencyShopMgr) GetBatchFromIDShop(idShops []uint32) (results []*CurrencyShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

func (obj *_CurrencyShopMgr) GetFromConversionRate(conversionRate float64) (results []*CurrencyShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("conversion_rate = ?", conversionRate).Find(&results).Error

	return
}

func (obj *_CurrencyShopMgr) GetBatchFromConversionRate(conversionRates []float64) (results []*CurrencyShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("conversion_rate IN (?)", conversionRates).Find(&results).Error

	return
}

func (obj *_CurrencyShopMgr) FetchByPrimaryKey(idCurrency uint32, idShop uint32) (result CurrencyShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_currency = ? AND id_shop = ?", idCurrency, idShop).Find(&result).Error

	return
}

func (obj *_CurrencyShopMgr) FetchIndexByIDShop(idShop uint32) (results []*CurrencyShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}
