package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _GroupShopMgr struct {
	*_BaseMgr
}

func GroupShopMgr(db *gorm.DB) *_GroupShopMgr {
	if db == nil {
		panic(fmt.Errorf("GroupShopMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_GroupShopMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_group_shop"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_GroupShopMgr) GetTableName() string {
	return "ps_group_shop"
}

func (obj *_GroupShopMgr) Get() (result GroupShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_GroupShopMgr) Gets() (results []*GroupShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

func (obj *_GroupShopMgr) WithIDGroup(idGroup uint32) Option {
	return optionFunc(func(o *options) { o.query["id_group"] = idGroup })
}

func (obj *_GroupShopMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

func (obj *_GroupShopMgr) GetByOption(opts ...Option) (result GroupShop, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_GroupShopMgr) GetByOptions(opts ...Option) (results []*GroupShop, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

func (obj *_GroupShopMgr) GetFromIDGroup(idGroup uint32) (results []*GroupShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_group = ?", idGroup).Find(&results).Error

	return
}

func (obj *_GroupShopMgr) GetBatchFromIDGroup(idGroups []uint32) (results []*GroupShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_group IN (?)", idGroups).Find(&results).Error

	return
}

func (obj *_GroupShopMgr) GetFromIDShop(idShop uint32) (results []*GroupShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

func (obj *_GroupShopMgr) GetBatchFromIDShop(idShops []uint32) (results []*GroupShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

func (obj *_GroupShopMgr) FetchByPrimaryKey(idGroup uint32, idShop uint32) (result GroupShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_group = ? AND id_shop = ?", idGroup, idShop).Find(&result).Error

	return
}

func (obj *_GroupShopMgr) FetchIndexByIDShop(idShop uint32) (results []*GroupShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}
