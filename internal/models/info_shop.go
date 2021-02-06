package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _InfoShopMgr struct {
	*_BaseMgr
}

func InfoShopMgr(db *gorm.DB) *_InfoShopMgr {
	if db == nil {
		panic(fmt.Errorf("InfoShopMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_InfoShopMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_info_shop"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_InfoShopMgr) GetTableName() string {
	return "ps_info_shop"
}

func (obj *_InfoShopMgr) Get() (result InfoShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_InfoShopMgr) Gets() (results []*InfoShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_InfoShopMgr) WithIDInfo(idInfo uint32) Option {
	return optionFunc(func(o *options) { o.query["id_info"] = idInfo })
}

func (obj *_InfoShopMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

func (obj *_InfoShopMgr) GetByOption(opts ...Option) (result InfoShop, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_InfoShopMgr) GetByOptions(opts ...Option) (results []*InfoShop, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_InfoShopMgr) GetFromIDInfo(idInfo uint32) (results []*InfoShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_info = ?", idInfo).Find(&results).Error

	return
}

func (obj *_InfoShopMgr) GetBatchFromIDInfo(idInfos []uint32) (results []*InfoShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_info IN (?)", idInfos).Find(&results).Error

	return
}

func (obj *_InfoShopMgr) GetFromIDShop(idShop uint32) (results []*InfoShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

func (obj *_InfoShopMgr) GetBatchFromIDShop(idShops []uint32) (results []*InfoShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}


func (obj *_InfoShopMgr) FetchByPrimaryKey(idInfo uint32, idShop uint32) (result InfoShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_info = ? AND id_shop = ?", idInfo, idShop).Find(&result).Error

	return
}
