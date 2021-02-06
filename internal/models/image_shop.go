package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _ImageShopMgr struct {
	*_BaseMgr
}

// ImageShopMgr open func
func ImageShopMgr(db *gorm.DB) *_ImageShopMgr {
	if db == nil {
		panic(fmt.Errorf("ImageShopMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ImageShopMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_image_shop"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ImageShopMgr) GetTableName() string {
	return "eg_image_shop"
}

// Get 获取
func (obj *_ImageShopMgr) Get() (result ImageShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ImageShopMgr) Gets() (results []*ImageShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDProduct id_product获取
func (obj *_ImageShopMgr) WithIDProduct(idProduct uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product"] = idProduct })
}

// WithIDImage id_image获取
func (obj *_ImageShopMgr) WithIDImage(idImage uint32) Option {
	return optionFunc(func(o *options) { o.query["id_image"] = idImage })
}

// WithIDShop id_shop获取
func (obj *_ImageShopMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

// WithCover cover获取
func (obj *_ImageShopMgr) WithCover(cover bool) Option {
	return optionFunc(func(o *options) { o.query["cover"] = cover })
}

// GetByOption 功能选项模式获取
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

// GetByOptions 批量功能选项模式获取
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

//////////////////////////enume case ////////////////////////////////////////////

// GetFromIDProduct 通过id_product获取内容
func (obj *_ImageShopMgr) GetFromIDProduct(idProduct uint32) (results []*ImageShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ?", idProduct).Find(&results).Error

	return
}

// GetBatchFromIDProduct 批量唯一主键查找
func (obj *_ImageShopMgr) GetBatchFromIDProduct(idProducts []uint32) (results []*ImageShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product IN (?)", idProducts).Find(&results).Error

	return
}

// GetFromIDImage 通过id_image获取内容
func (obj *_ImageShopMgr) GetFromIDImage(idImage uint32) (results []*ImageShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_image = ?", idImage).Find(&results).Error

	return
}

// GetBatchFromIDImage 批量唯一主键查找
func (obj *_ImageShopMgr) GetBatchFromIDImage(idImages []uint32) (results []*ImageShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_image IN (?)", idImages).Find(&results).Error

	return
}

// GetFromIDShop 通过id_shop获取内容
func (obj *_ImageShopMgr) GetFromIDShop(idShop uint32) (results []*ImageShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

// GetBatchFromIDShop 批量唯一主键查找
func (obj *_ImageShopMgr) GetBatchFromIDShop(idShops []uint32) (results []*ImageShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

// GetFromCover 通过cover获取内容
func (obj *_ImageShopMgr) GetFromCover(cover bool) (results []*ImageShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cover = ?", cover).Find(&results).Error

	return
}

// GetBatchFromCover 批量唯一主键查找
func (obj *_ImageShopMgr) GetBatchFromCover(covers []bool) (results []*ImageShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cover IN (?)", covers).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_ImageShopMgr) FetchByPrimaryKey(idImage uint32, idShop uint32) (result ImageShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_image = ? AND id_shop = ?", idImage, idShop).Find(&result).Error

	return
}

// FetchUniqueIndexByIDProduct primay or index 获取唯一内容
func (obj *_ImageShopMgr) FetchUniqueIndexByIDProduct(idProduct uint32, idShop uint32, cover bool) (result ImageShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ? AND id_shop = ? AND cover = ?", idProduct, idShop, cover).Find(&result).Error

	return
}

// FetchIndexByIDShop  获取多个内容
func (obj *_ImageShopMgr) FetchIndexByIDShop(idShop uint32) (results []*ImageShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}
