package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _CmsShopMgr struct {
	*_BaseMgr
}

func CmsShopMgr(db *gorm.DB) *_CmsShopMgr {
	if db == nil {
		panic(fmt.Errorf("CmsShopMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CmsShopMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_cms_shop"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_CmsShopMgr) GetTableName() string {
	return "ps_cms_shop"
}

func (obj *_CmsShopMgr) Get() (result CmsShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_CmsShopMgr) Gets() (results []*CmsShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_CmsShopMgr) WithIDCms(idCms uint32) Option {
	return optionFunc(func(o *options) { o.query["id_cms"] = idCms })
}

func (obj *_CmsShopMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

func (obj *_CmsShopMgr) GetByOption(opts ...Option) (result CmsShop, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_CmsShopMgr) GetByOptions(opts ...Option) (results []*CmsShop, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_CmsShopMgr) GetFromIDCms(idCms uint32) (results []*CmsShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cms = ?", idCms).Find(&results).Error

	return
}

func (obj *_CmsShopMgr) GetBatchFromIDCms(idCmss []uint32) (results []*CmsShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cms IN (?)", idCmss).Find(&results).Error

	return
}

func (obj *_CmsShopMgr) GetFromIDShop(idShop uint32) (results []*CmsShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

func (obj *_CmsShopMgr) GetBatchFromIDShop(idShops []uint32) (results []*CmsShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}


func (obj *_CmsShopMgr) FetchByPrimaryKey(idCms uint32, idShop uint32) (result CmsShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cms = ? AND id_shop = ?", idCms, idShop).Find(&result).Error

	return
}

func (obj *_CmsShopMgr) FetchIndexByIDShop(idShop uint32) (results []*CmsShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}
