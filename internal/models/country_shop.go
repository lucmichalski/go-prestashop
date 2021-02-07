package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _CountryShopMgr struct {
	*_BaseMgr
}

func CountryShopMgr(db *gorm.DB) *_CountryShopMgr {
	if db == nil {
		panic(fmt.Errorf("CountryShopMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CountryShopMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_country_shop"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_CountryShopMgr) GetTableName() string {
	return "ps_country_shop"
}

func (obj *_CountryShopMgr) Get() (result CountryShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_CountryShopMgr) Gets() (results []*CountryShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

func (obj *_CountryShopMgr) WithIDCountry(idCountry uint32) Option {
	return optionFunc(func(o *options) { o.query["id_country"] = idCountry })
}

func (obj *_CountryShopMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

func (obj *_CountryShopMgr) GetByOption(opts ...Option) (result CountryShop, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_CountryShopMgr) GetByOptions(opts ...Option) (results []*CountryShop, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

func (obj *_CountryShopMgr) GetFromIDCountry(idCountry uint32) (results []*CountryShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_country = ?", idCountry).Find(&results).Error

	return
}

func (obj *_CountryShopMgr) GetBatchFromIDCountry(idCountrys []uint32) (results []*CountryShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_country IN (?)", idCountrys).Find(&results).Error

	return
}

func (obj *_CountryShopMgr) GetFromIDShop(idShop uint32) (results []*CountryShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

func (obj *_CountryShopMgr) GetBatchFromIDShop(idShops []uint32) (results []*CountryShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

func (obj *_CountryShopMgr) FetchByPrimaryKey(idCountry uint32, idShop uint32) (result CountryShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_country = ? AND id_shop = ?", idCountry, idShop).Find(&result).Error

	return
}

func (obj *_CountryShopMgr) FetchIndexByIDShop(idShop uint32) (results []*CountryShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}
