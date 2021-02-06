package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _CarrierShopMgr struct {
	*_BaseMgr
}

func CarrierShopMgr(db *gorm.DB) *_CarrierShopMgr {
	if db == nil {
		panic(fmt.Errorf("CarrierShopMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CarrierShopMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_carrier_shop"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_CarrierShopMgr) GetTableName() string {
	return "ps_carrier_shop"
}

func (obj *_CarrierShopMgr) Get() (result CarrierShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_CarrierShopMgr) Gets() (results []*CarrierShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_CarrierShopMgr) WithIDCarrier(idCarrier uint32) Option {
	return optionFunc(func(o *options) { o.query["id_carrier"] = idCarrier })
}

func (obj *_CarrierShopMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

func (obj *_CarrierShopMgr) GetByOption(opts ...Option) (result CarrierShop, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_CarrierShopMgr) GetByOptions(opts ...Option) (results []*CarrierShop, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_CarrierShopMgr) GetFromIDCarrier(idCarrier uint32) (results []*CarrierShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_carrier = ?", idCarrier).Find(&results).Error

	return
}

func (obj *_CarrierShopMgr) GetBatchFromIDCarrier(idCarriers []uint32) (results []*CarrierShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_carrier IN (?)", idCarriers).Find(&results).Error

	return
}

func (obj *_CarrierShopMgr) GetFromIDShop(idShop uint32) (results []*CarrierShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

func (obj *_CarrierShopMgr) GetBatchFromIDShop(idShops []uint32) (results []*CarrierShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}


func (obj *_CarrierShopMgr) FetchByPrimaryKey(idCarrier uint32, idShop uint32) (result CarrierShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_carrier = ? AND id_shop = ?", idCarrier, idShop).Find(&result).Error

	return
}

func (obj *_CarrierShopMgr) FetchIndexByIDShop(idShop uint32) (results []*CarrierShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}
