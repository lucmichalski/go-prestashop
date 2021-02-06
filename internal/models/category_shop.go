package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _CategoryShopMgr struct {
	*_BaseMgr
}

func CategoryShopMgr(db *gorm.DB) *_CategoryShopMgr {
	if db == nil {
		panic(fmt.Errorf("CategoryShopMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CategoryShopMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_category_shop"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_CategoryShopMgr) GetTableName() string {
	return "ps_category_shop"
}

func (obj *_CategoryShopMgr) Get() (result CategoryShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_CategoryShopMgr) Gets() (results []*CategoryShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_CategoryShopMgr) WithIDCategory(idCategory int) Option {
	return optionFunc(func(o *options) { o.query["id_category"] = idCategory })
}

func (obj *_CategoryShopMgr) WithIDShop(idShop int) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

func (obj *_CategoryShopMgr) WithPosition(position uint32) Option {
	return optionFunc(func(o *options) { o.query["position"] = position })
}

func (obj *_CategoryShopMgr) GetByOption(opts ...Option) (result CategoryShop, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_CategoryShopMgr) GetByOptions(opts ...Option) (results []*CategoryShop, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_CategoryShopMgr) GetFromIDCategory(idCategory int) (results []*CategoryShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_category = ?", idCategory).Find(&results).Error

	return
}

func (obj *_CategoryShopMgr) GetBatchFromIDCategory(idCategorys []int) (results []*CategoryShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_category IN (?)", idCategorys).Find(&results).Error

	return
}

func (obj *_CategoryShopMgr) GetFromIDShop(idShop int) (results []*CategoryShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

func (obj *_CategoryShopMgr) GetBatchFromIDShop(idShops []int) (results []*CategoryShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

func (obj *_CategoryShopMgr) GetFromPosition(position uint32) (results []*CategoryShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("position = ?", position).Find(&results).Error

	return
}

func (obj *_CategoryShopMgr) GetBatchFromPosition(positions []uint32) (results []*CategoryShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("position IN (?)", positions).Find(&results).Error

	return
}


func (obj *_CategoryShopMgr) FetchByPrimaryKey(idCategory int, idShop int) (result CategoryShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_category = ? AND id_shop = ?", idCategory, idShop).Find(&result).Error

	return
}
