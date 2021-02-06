package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _LayeredFilterShopMgr struct {
	*_BaseMgr
}

func LayeredFilterShopMgr(db *gorm.DB) *_LayeredFilterShopMgr {
	if db == nil {
		panic(fmt.Errorf("LayeredFilterShopMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_LayeredFilterShopMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_layered_filter_shop"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_LayeredFilterShopMgr) GetTableName() string {
	return "ps_layered_filter_shop"
}

func (obj *_LayeredFilterShopMgr) Get() (result LayeredFilterShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_LayeredFilterShopMgr) Gets() (results []*LayeredFilterShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_LayeredFilterShopMgr) WithIDLayeredFilter(idLayeredFilter uint32) Option {
	return optionFunc(func(o *options) { o.query["id_layered_filter"] = idLayeredFilter })
}

func (obj *_LayeredFilterShopMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

func (obj *_LayeredFilterShopMgr) GetByOption(opts ...Option) (result LayeredFilterShop, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_LayeredFilterShopMgr) GetByOptions(opts ...Option) (results []*LayeredFilterShop, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_LayeredFilterShopMgr) GetFromIDLayeredFilter(idLayeredFilter uint32) (results []*LayeredFilterShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_layered_filter = ?", idLayeredFilter).Find(&results).Error

	return
}

func (obj *_LayeredFilterShopMgr) GetBatchFromIDLayeredFilter(idLayeredFilters []uint32) (results []*LayeredFilterShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_layered_filter IN (?)", idLayeredFilters).Find(&results).Error

	return
}

func (obj *_LayeredFilterShopMgr) GetFromIDShop(idShop uint32) (results []*LayeredFilterShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

func (obj *_LayeredFilterShopMgr) GetBatchFromIDShop(idShops []uint32) (results []*LayeredFilterShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}


func (obj *_LayeredFilterShopMgr) FetchByPrimaryKey(idLayeredFilter uint32, idShop uint32) (result LayeredFilterShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_layered_filter = ? AND id_shop = ?", idLayeredFilter, idShop).Find(&result).Error

	return
}

func (obj *_LayeredFilterShopMgr) FetchIndexByIDShop(idShop uint32) (results []*LayeredFilterShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}
