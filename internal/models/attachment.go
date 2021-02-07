package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _AttachmentMgr struct {
	*_BaseMgr
}

func AttachmentMgr(db *gorm.DB) *_AttachmentMgr {
	if db == nil {
		panic(fmt.Errorf("AttachmentMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AttachmentMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_attachment"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_AttachmentMgr) GetTableName() string {
	return "ps_attachment"
}

func (obj *_AttachmentMgr) Get() (result Attachment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_AttachmentMgr) Gets() (results []*Attachment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

func (obj *_AttachmentMgr) WithIDAttachment(idAttachment uint32) Option {
	return optionFunc(func(o *options) { o.query["id_attachment"] = idAttachment })
}

func (obj *_AttachmentMgr) WithFile(file string) Option {
	return optionFunc(func(o *options) { o.query["file"] = file })
}

func (obj *_AttachmentMgr) WithFileName(fileName string) Option {
	return optionFunc(func(o *options) { o.query["file_name"] = fileName })
}

func (obj *_AttachmentMgr) WithFileSize(fileSize uint64) Option {
	return optionFunc(func(o *options) { o.query["file_size"] = fileSize })
}

func (obj *_AttachmentMgr) WithMime(mime string) Option {
	return optionFunc(func(o *options) { o.query["mime"] = mime })
}

func (obj *_AttachmentMgr) GetByOption(opts ...Option) (result Attachment, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_AttachmentMgr) GetByOptions(opts ...Option) (results []*Attachment, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

func (obj *_AttachmentMgr) GetFromIDAttachment(idAttachment uint32) (result Attachment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attachment = ?", idAttachment).Find(&result).Error

	return
}

func (obj *_AttachmentMgr) GetBatchFromIDAttachment(idAttachments []uint32) (results []*Attachment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attachment IN (?)", idAttachments).Find(&results).Error

	return
}

func (obj *_AttachmentMgr) GetFromFile(file string) (results []*Attachment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("file = ?", file).Find(&results).Error

	return
}

func (obj *_AttachmentMgr) GetBatchFromFile(files []string) (results []*Attachment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("file IN (?)", files).Find(&results).Error

	return
}

func (obj *_AttachmentMgr) GetFromFileName(fileName string) (results []*Attachment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("file_name = ?", fileName).Find(&results).Error

	return
}

func (obj *_AttachmentMgr) GetBatchFromFileName(fileNames []string) (results []*Attachment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("file_name IN (?)", fileNames).Find(&results).Error

	return
}

func (obj *_AttachmentMgr) GetFromFileSize(fileSize uint64) (results []*Attachment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("file_size = ?", fileSize).Find(&results).Error

	return
}

func (obj *_AttachmentMgr) GetBatchFromFileSize(fileSizes []uint64) (results []*Attachment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("file_size IN (?)", fileSizes).Find(&results).Error

	return
}

func (obj *_AttachmentMgr) GetFromMime(mime string) (results []*Attachment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("mime = ?", mime).Find(&results).Error

	return
}

func (obj *_AttachmentMgr) GetBatchFromMime(mimes []string) (results []*Attachment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("mime IN (?)", mimes).Find(&results).Error

	return
}

func (obj *_AttachmentMgr) FetchByPrimaryKey(idAttachment uint32) (result Attachment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attachment = ?", idAttachment).Find(&result).Error

	return
}
