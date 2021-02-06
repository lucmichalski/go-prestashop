package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _ProductAttachmentMgr struct {
	*_BaseMgr
}

func ProductAttachmentMgr(db *gorm.DB) *_ProductAttachmentMgr {
	if db == nil {
		panic(fmt.Errorf("ProductAttachmentMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ProductAttachmentMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_product_attachment"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_ProductAttachmentMgr) GetTableName() string {
	return "ps_product_attachment"
}

func (obj *_ProductAttachmentMgr) Get() (result ProductAttachment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_ProductAttachmentMgr) Gets() (results []*ProductAttachment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_ProductAttachmentMgr) WithIDProduct(idProduct uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product"] = idProduct })
}

func (obj *_ProductAttachmentMgr) WithIDAttachment(idAttachment uint32) Option {
	return optionFunc(func(o *options) { o.query["id_attachment"] = idAttachment })
}

func (obj *_ProductAttachmentMgr) GetByOption(opts ...Option) (result ProductAttachment, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_ProductAttachmentMgr) GetByOptions(opts ...Option) (results []*ProductAttachment, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_ProductAttachmentMgr) GetFromIDProduct(idProduct uint32) (results []*ProductAttachment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ?", idProduct).Find(&results).Error

	return
}

func (obj *_ProductAttachmentMgr) GetBatchFromIDProduct(idProducts []uint32) (results []*ProductAttachment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product IN (?)", idProducts).Find(&results).Error

	return
}

func (obj *_ProductAttachmentMgr) GetFromIDAttachment(idAttachment uint32) (results []*ProductAttachment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attachment = ?", idAttachment).Find(&results).Error

	return
}

func (obj *_ProductAttachmentMgr) GetBatchFromIDAttachment(idAttachments []uint32) (results []*ProductAttachment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attachment IN (?)", idAttachments).Find(&results).Error

	return
}


func (obj *_ProductAttachmentMgr) FetchByPrimaryKey(idProduct uint32, idAttachment uint32) (result ProductAttachment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ? AND id_attachment = ?", idProduct, idAttachment).Find(&result).Error

	return
}
