package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _ImageMgr struct {
	*_BaseMgr
}

func ImageMgr(db *gorm.DB) *_ImageMgr {
	if db == nil {
		panic(fmt.Errorf("ImageMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ImageMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_image"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_ImageMgr) GetTableName() string {
	return "ps_image"
}

func (obj *_ImageMgr) Get() (result Image, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_ImageMgr) Gets() (results []*Image, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

func (obj *_ImageMgr) WithIDImage(idImage uint32) Option {
	return optionFunc(func(o *options) { o.query["id_image"] = idImage })
}

func (obj *_ImageMgr) WithIDProduct(idProduct uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product"] = idProduct })
}

func (obj *_ImageMgr) WithPosition(position uint16) Option {
	return optionFunc(func(o *options) { o.query["position"] = position })
}

func (obj *_ImageMgr) WithCover(cover bool) Option {
	return optionFunc(func(o *options) { o.query["cover"] = cover })
}

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

func (obj *_ImageMgr) GetFromIDImage(idImage uint32) (result Image, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_image = ?", idImage).Find(&result).Error

	return
}

func (obj *_ImageMgr) GetBatchFromIDImage(idImages []uint32) (results []*Image, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_image IN (?)", idImages).Find(&results).Error

	return
}

func (obj *_ImageMgr) GetFromIDProduct(idProduct uint32) (results []*Image, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ?", idProduct).Find(&results).Error

	return
}

func (obj *_ImageMgr) GetBatchFromIDProduct(idProducts []uint32) (results []*Image, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product IN (?)", idProducts).Find(&results).Error

	return
}

func (obj *_ImageMgr) GetFromPosition(position uint16) (results []*Image, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("position = ?", position).Find(&results).Error

	return
}

func (obj *_ImageMgr) GetBatchFromPosition(positions []uint16) (results []*Image, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("position IN (?)", positions).Find(&results).Error

	return
}

func (obj *_ImageMgr) GetFromCover(cover bool) (results []*Image, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cover = ?", cover).Find(&results).Error

	return
}

func (obj *_ImageMgr) GetBatchFromCover(covers []bool) (results []*Image, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cover IN (?)", covers).Find(&results).Error

	return
}

func (obj *_ImageMgr) FetchByPrimaryKey(idImage uint32) (result Image, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_image = ?", idImage).Find(&result).Error

	return
}

func (obj *_ImageMgr) FetchUniqueIndexByIDxProductImage(idImage uint32, idProduct uint32, cover bool) (result Image, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_image = ? AND id_product = ? AND cover = ?", idImage, idProduct, cover).Find(&result).Error

	return
}

func (obj *_ImageMgr) FetchUniqueIndexByIDProductCover(idProduct uint32, cover bool) (result Image, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ? AND cover = ?", idProduct, cover).Find(&result).Error

	return
}

func (obj *_ImageMgr) FetchIndexByImageProduct(idProduct uint32) (results []*Image, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ?", idProduct).Find(&results).Error

	return
}
