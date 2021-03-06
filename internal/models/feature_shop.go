package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _FeatureShopMgr struct {
	*_BaseMgr
}

func FeatureShopMgr(db *gorm.DB) *_FeatureShopMgr {
	if db == nil {
		panic(fmt.Errorf("FeatureShopMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_FeatureShopMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_feature_shop"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_FeatureShopMgr) GetTableName() string {
	return "ps_feature_shop"
}

func (obj *_FeatureShopMgr) Get() (result FeatureShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_FeatureShopMgr) Gets() (results []*FeatureShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

func (obj *_FeatureShopMgr) WithIDFeature(idFeature uint32) Option {
	return optionFunc(func(o *options) { o.query["id_feature"] = idFeature })
}

func (obj *_FeatureShopMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

func (obj *_FeatureShopMgr) GetByOption(opts ...Option) (result FeatureShop, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_FeatureShopMgr) GetByOptions(opts ...Option) (results []*FeatureShop, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

func (obj *_FeatureShopMgr) GetFromIDFeature(idFeature uint32) (results []*FeatureShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_feature = ?", idFeature).Find(&results).Error

	return
}

func (obj *_FeatureShopMgr) GetBatchFromIDFeature(idFeatures []uint32) (results []*FeatureShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_feature IN (?)", idFeatures).Find(&results).Error

	return
}

func (obj *_FeatureShopMgr) GetFromIDShop(idShop uint32) (results []*FeatureShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

func (obj *_FeatureShopMgr) GetBatchFromIDShop(idShops []uint32) (results []*FeatureShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

func (obj *_FeatureShopMgr) FetchByPrimaryKey(idFeature uint32, idShop uint32) (result FeatureShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_feature = ? AND id_shop = ?", idFeature, idShop).Find(&result).Error

	return
}

func (obj *_FeatureShopMgr) FetchIndexByIDShop(idShop uint32) (results []*FeatureShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}
