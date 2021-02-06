package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _CountryShopMgr struct {
	*_BaseMgr
}

// CountryShopMgr open func
func CountryShopMgr(db *gorm.DB) *_CountryShopMgr {
	if db == nil {
		panic(fmt.Errorf("CountryShopMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CountryShopMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_country_shop"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CountryShopMgr) GetTableName() string {
	return "eg_country_shop"
}

// Get 获取
func (obj *_CountryShopMgr) Get() (result CountryShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CountryShopMgr) Gets() (results []*CountryShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDCountry id_country获取
func (obj *_CountryShopMgr) WithIDCountry(idCountry uint32) Option {
	return optionFunc(func(o *options) { o.query["id_country"] = idCountry })
}

// WithIDShop id_shop获取
func (obj *_CountryShopMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

// GetByOption 功能选项模式获取
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

// GetByOptions 批量功能选项模式获取
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

//////////////////////////enume case ////////////////////////////////////////////

// GetFromIDCountry 通过id_country获取内容
func (obj *_CountryShopMgr) GetFromIDCountry(idCountry uint32) (results []*CountryShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_country = ?", idCountry).Find(&results).Error

	return
}

// GetBatchFromIDCountry 批量唯一主键查找
func (obj *_CountryShopMgr) GetBatchFromIDCountry(idCountrys []uint32) (results []*CountryShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_country IN (?)", idCountrys).Find(&results).Error

	return
}

// GetFromIDShop 通过id_shop获取内容
func (obj *_CountryShopMgr) GetFromIDShop(idShop uint32) (results []*CountryShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

// GetBatchFromIDShop 批量唯一主键查找
func (obj *_CountryShopMgr) GetBatchFromIDShop(idShops []uint32) (results []*CountryShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_CountryShopMgr) FetchByPrimaryKey(idCountry uint32, idShop uint32) (result CountryShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_country = ? AND id_shop = ?", idCountry, idShop).Find(&result).Error

	return
}

// FetchIndexByIDShop  获取多个内容
func (obj *_CountryShopMgr) FetchIndexByIDShop(idShop uint32) (results []*CountryShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}
