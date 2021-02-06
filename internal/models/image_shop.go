package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _ImageShopMgr struct {
	*_BaseMgr
}

func ImageShopMgr(db *gorm.DB) *_ImageShopMgr {
	if db == nil {
		panic(fmt.Errorf("ImageShopMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ImageShopMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_image_shop"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_ImageShopMgr) GetTableName() string {
	return "ps_image_shop"
}

func (obj *_ImageShopMgr) Get() (result ImageShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_ImageShopMgr) Gets() (results []*ImageShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_ImageShopMgr) WithIDProduct(idProduct uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product"] = idProduct })
}

func (obj *_ImageShopMgr) WithIDImage(idImage uint32) Option {
	return optionFunc(func(o *options) { o.query["id_image"] = idImage })
}

func (obj *_ImageShopMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

func (obj *_ImageShopMgr) WithCover(cover bool) Option {
	return optionFunc(func(o *options) { o.query["cover"] = cover })
}

func (obj *_ImageShopMgr) GetByOption(opts ...Option) (result ImageShop, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_ImageShopMgr) GetByOptions(opts ...Option) (results []*ImageShop, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_ImageShopMgr) GetFromIDProduct(idProduct uint32) (results []*ImageShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ?", idProduct).Find(&results).Error

	return
}

func (obj *_ImageShopMgr) GetBatchFromIDProduct(idProducts []uint32) (results []*ImageShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product IN (?)", idProducts).Find(&results).Error

	return
}

func (obj *_ImageShopMgr) GetFromIDImage(idImage uint32) (results []*ImageShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_image = ?", idImage).Find(&results).Error

	return
}

func (obj *_ImageShopMgr) GetBatchFromIDImage(idImages []uint32) (results []*ImageShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_image IN (?)", idImages).Find(&results).Error

	return
}

func (obj *_ImageShopMgr) GetFromIDShop(idShop uint32) (results []*ImageShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

func (obj *_ImageShopMgr) GetBatchFromIDShop(idShops []uint32) (results []*ImageShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

func (obj *_ImageShopMgr) GetFromCover(cover bool) (results []*ImageShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cover = ?", cover).Find(&results).Error

	return
}

func (obj *_ImageShopMgr) GetBatchFromCover(covers []bool) (results []*ImageShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cover IN (?)", covers).Find(&results).Error

	return
}


func (obj *_ImageShopMgr) FetchByPrimaryKey(idImage uint32, idShop uint32) (result ImageShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_image = ? AND id_shop = ?", idImage, idShop).Find(&result).Error

	return
}

func (obj *_ImageShopMgr) FetchUniqueIndexByIDProduct(idProduct uint32, idShop uint32, cover bool) (result ImageShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ? AND id_shop = ? AND cover = ?", idProduct, idShop, cover).Find(&result).Error

	return
}

func (obj *_ImageShopMgr) FetchIndexByIDShop(idShop uint32) (results []*ImageShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}
