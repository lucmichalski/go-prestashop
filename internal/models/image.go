package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _ImageMgr struct {
	*_BaseMgr
}

// ImageMgr open func
func ImageMgr(db *gorm.DB) *_ImageMgr {
	if db == nil {
		panic(fmt.Errorf("ImageMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ImageMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_image"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ImageMgr) GetTableName() string {
	return "eg_image"
}

// Get 获取
func (obj *_ImageMgr) Get() (result Image, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ImageMgr) Gets() (results []*Image, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDImage id_image获取
func (obj *_ImageMgr) WithIDImage(idImage uint32) Option {
	return optionFunc(func(o *options) { o.query["id_image"] = idImage })
}

// WithIDProduct id_product获取
func (obj *_ImageMgr) WithIDProduct(idProduct uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product"] = idProduct })
}

// WithPosition position获取
func (obj *_ImageMgr) WithPosition(position uint16) Option {
	return optionFunc(func(o *options) { o.query["position"] = position })
}

// WithCover cover获取
func (obj *_ImageMgr) WithCover(cover bool) Option {
	return optionFunc(func(o *options) { o.query["cover"] = cover })
}

// GetByOption 功能选项模式获取
func (obj *_ImageMgr) GetByOption(opts ...Option) (result Image, err error) {
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
func (obj *_ImageMgr) GetByOptions(opts ...Option) (results []*Image, err error) {
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

// GetFromIDImage 通过id_image获取内容
func (obj *_ImageMgr) GetFromIDImage(idImage uint32) (result Image, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_image = ?", idImage).Find(&result).Error

	return
}

// GetBatchFromIDImage 批量唯一主键查找
func (obj *_ImageMgr) GetBatchFromIDImage(idImages []uint32) (results []*Image, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_image IN (?)", idImages).Find(&results).Error

	return
}

// GetFromIDProduct 通过id_product获取内容
func (obj *_ImageMgr) GetFromIDProduct(idProduct uint32) (results []*Image, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ?", idProduct).Find(&results).Error

	return
}

// GetBatchFromIDProduct 批量唯一主键查找
func (obj *_ImageMgr) GetBatchFromIDProduct(idProducts []uint32) (results []*Image, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product IN (?)", idProducts).Find(&results).Error

	return
}

// GetFromPosition 通过position获取内容
func (obj *_ImageMgr) GetFromPosition(position uint16) (results []*Image, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("position = ?", position).Find(&results).Error

	return
}

// GetBatchFromPosition 批量唯一主键查找
func (obj *_ImageMgr) GetBatchFromPosition(positions []uint16) (results []*Image, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("position IN (?)", positions).Find(&results).Error

	return
}

// GetFromCover 通过cover获取内容
func (obj *_ImageMgr) GetFromCover(cover bool) (results []*Image, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cover = ?", cover).Find(&results).Error

	return
}

// GetBatchFromCover 批量唯一主键查找
func (obj *_ImageMgr) GetBatchFromCover(covers []bool) (results []*Image, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cover IN (?)", covers).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_ImageMgr) FetchByPrimaryKey(idImage uint32) (result Image, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_image = ?", idImage).Find(&result).Error

	return
}

// FetchUniqueIndexByIDxProductImage primay or index 获取唯一内容
func (obj *_ImageMgr) FetchUniqueIndexByIDxProductImage(idImage uint32, idProduct uint32, cover bool) (result Image, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_image = ? AND id_product = ? AND cover = ?", idImage, idProduct, cover).Find(&result).Error

	return
}

// FetchUniqueIndexByIDProductCover primay or index 获取唯一内容
func (obj *_ImageMgr) FetchUniqueIndexByIDProductCover(idProduct uint32, cover bool) (result Image, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ? AND cover = ?", idProduct, cover).Find(&result).Error

	return
}

// FetchIndexByImageProduct  获取多个内容
func (obj *_ImageMgr) FetchIndexByImageProduct(idProduct uint32) (results []*Image, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ?", idProduct).Find(&results).Error

	return
}
