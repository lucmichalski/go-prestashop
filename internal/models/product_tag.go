package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _ProductTagMgr struct {
	*_BaseMgr
}

func ProductTagMgr(db *gorm.DB) *_ProductTagMgr {
	if db == nil {
		panic(fmt.Errorf("ProductTagMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ProductTagMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_product_tag"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_ProductTagMgr) GetTableName() string {
	return "ps_product_tag"
}

func (obj *_ProductTagMgr) Get() (result ProductTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_ProductTagMgr) Gets() (results []*ProductTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

func (obj *_ProductTagMgr) WithIDProduct(idProduct uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product"] = idProduct })
}

func (obj *_ProductTagMgr) WithIDTag(idTag uint32) Option {
	return optionFunc(func(o *options) { o.query["id_tag"] = idTag })
}

func (obj *_ProductTagMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

func (obj *_ProductTagMgr) GetByOption(opts ...Option) (result ProductTag, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_ProductTagMgr) GetByOptions(opts ...Option) (results []*ProductTag, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

func (obj *_ProductTagMgr) GetFromIDProduct(idProduct uint32) (results []*ProductTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ?", idProduct).Find(&results).Error

	return
}

func (obj *_ProductTagMgr) GetBatchFromIDProduct(idProducts []uint32) (results []*ProductTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product IN (?)", idProducts).Find(&results).Error

	return
}

func (obj *_ProductTagMgr) GetFromIDTag(idTag uint32) (results []*ProductTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tag = ?", idTag).Find(&results).Error

	return
}

func (obj *_ProductTagMgr) GetBatchFromIDTag(idTags []uint32) (results []*ProductTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tag IN (?)", idTags).Find(&results).Error

	return
}

func (obj *_ProductTagMgr) GetFromIDLang(idLang uint32) (results []*ProductTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error

	return
}

func (obj *_ProductTagMgr) GetBatchFromIDLang(idLangs []uint32) (results []*ProductTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error

	return
}

func (obj *_ProductTagMgr) FetchByPrimaryKey(idProduct uint32, idTag uint32) (result ProductTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ? AND id_tag = ?", idProduct, idTag).Find(&result).Error

	return
}

func (obj *_ProductTagMgr) FetchIndexByIDTag(idTag uint32) (results []*ProductTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tag = ?", idTag).Find(&results).Error

	return
}

func (obj *_ProductTagMgr) FetchIndexByIDLang(idTag uint32, idLang uint32) (results []*ProductTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tag = ? AND id_lang = ?", idTag, idLang).Find(&results).Error

	return
}
