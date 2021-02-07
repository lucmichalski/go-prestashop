package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _HomesliderMgr struct {
	*_BaseMgr
}

func HomesliderMgr(db *gorm.DB) *_HomesliderMgr {
	if db == nil {
		panic(fmt.Errorf("HomesliderMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_HomesliderMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_homeslider"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_HomesliderMgr) GetTableName() string {
	return "ps_homeslider"
}

func (obj *_HomesliderMgr) Get() (result Homeslider, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_HomesliderMgr) Gets() (results []*Homeslider, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

func (obj *_HomesliderMgr) WithIDHomesliderSlides(idHomesliderSlides uint32) Option {
	return optionFunc(func(o *options) { o.query["id_homeslider_slides"] = idHomesliderSlides })
}

func (obj *_HomesliderMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

func (obj *_HomesliderMgr) GetByOption(opts ...Option) (result Homeslider, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_HomesliderMgr) GetByOptions(opts ...Option) (results []*Homeslider, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

func (obj *_HomesliderMgr) GetFromIDHomesliderSlides(idHomesliderSlides uint32) (results []*Homeslider, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_homeslider_slides = ?", idHomesliderSlides).Find(&results).Error

	return
}

func (obj *_HomesliderMgr) GetBatchFromIDHomesliderSlides(idHomesliderSlidess []uint32) (results []*Homeslider, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_homeslider_slides IN (?)", idHomesliderSlidess).Find(&results).Error

	return
}

func (obj *_HomesliderMgr) GetFromIDShop(idShop uint32) (results []*Homeslider, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

func (obj *_HomesliderMgr) GetBatchFromIDShop(idShops []uint32) (results []*Homeslider, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

func (obj *_HomesliderMgr) FetchByPrimaryKey(idHomesliderSlides uint32, idShop uint32) (result Homeslider, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_homeslider_slides = ? AND id_shop = ?", idHomesliderSlides, idShop).Find(&result).Error

	return
}
