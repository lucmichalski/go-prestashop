package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _MailMgr struct {
	*_BaseMgr
}

// MailMgr open func
func MailMgr(db *gorm.DB) *_MailMgr {
	if db == nil {
		panic(fmt.Errorf("MailMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_MailMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_mail"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_MailMgr) GetTableName() string {
	return "ps_mail"
}

// Get 获取
func (obj *_MailMgr) Get() (result Mail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_MailMgr) Gets() (results []*Mail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDMail id_mail获取
func (obj *_MailMgr) WithIDMail(idMail uint32) Option {
	return optionFunc(func(o *options) { o.query["id_mail"] = idMail })
}

// WithRecipient recipient获取
func (obj *_MailMgr) WithRecipient(recipient string) Option {
	return optionFunc(func(o *options) { o.query["recipient"] = recipient })
}

// WithTemplate template获取
func (obj *_MailMgr) WithTemplate(template string) Option {
	return optionFunc(func(o *options) { o.query["template"] = template })
}

// WithSubject subject获取
func (obj *_MailMgr) WithSubject(subject string) Option {
	return optionFunc(func(o *options) { o.query["subject"] = subject })
}

// WithIDLang id_lang获取
func (obj *_MailMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

// WithDateAdd date_add获取
func (obj *_MailMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

// GetByOption 功能选项模式获取
func (obj *_MailMgr) GetByOption(opts ...Option) (result Mail, err error) {
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
func (obj *_MailMgr) GetByOptions(opts ...Option) (results []*Mail, err error) {
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

// GetFromIDMail 通过id_mail获取内容
func (obj *_MailMgr) GetFromIDMail(idMail uint32) (result Mail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_mail = ?", idMail).Find(&result).Error

	return
}

// GetBatchFromIDMail 批量唯一主键查找
func (obj *_MailMgr) GetBatchFromIDMail(idMails []uint32) (results []*Mail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_mail IN (?)", idMails).Find(&results).Error

	return
}

// GetFromRecipient 通过recipient获取内容
func (obj *_MailMgr) GetFromRecipient(recipient string) (results []*Mail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("recipient = ?", recipient).Find(&results).Error

	return
}

// GetBatchFromRecipient 批量唯一主键查找
func (obj *_MailMgr) GetBatchFromRecipient(recipients []string) (results []*Mail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("recipient IN (?)", recipients).Find(&results).Error

	return
}

// GetFromTemplate 通过template获取内容
func (obj *_MailMgr) GetFromTemplate(template string) (results []*Mail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("template = ?", template).Find(&results).Error

	return
}

// GetBatchFromTemplate 批量唯一主键查找
func (obj *_MailMgr) GetBatchFromTemplate(templates []string) (results []*Mail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("template IN (?)", templates).Find(&results).Error

	return
}

// GetFromSubject 通过subject获取内容
func (obj *_MailMgr) GetFromSubject(subject string) (results []*Mail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("subject = ?", subject).Find(&results).Error

	return
}

// GetBatchFromSubject 批量唯一主键查找
func (obj *_MailMgr) GetBatchFromSubject(subjects []string) (results []*Mail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("subject IN (?)", subjects).Find(&results).Error

	return
}

// GetFromIDLang 通过id_lang获取内容
func (obj *_MailMgr) GetFromIDLang(idLang uint32) (results []*Mail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error

	return
}

// GetBatchFromIDLang 批量唯一主键查找
func (obj *_MailMgr) GetBatchFromIDLang(idLangs []uint32) (results []*Mail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error

	return
}

// GetFromDateAdd 通过date_add获取内容
func (obj *_MailMgr) GetFromDateAdd(dateAdd time.Time) (results []*Mail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error

	return
}

// GetBatchFromDateAdd 批量唯一主键查找
func (obj *_MailMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*Mail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_MailMgr) FetchByPrimaryKey(idMail uint32) (result Mail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_mail = ?", idMail).Find(&result).Error

	return
}

// FetchIndexByRecipient  获取多个内容
func (obj *_MailMgr) FetchIndexByRecipient(recipient string) (results []*Mail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("recipient = ?", recipient).Find(&results).Error

	return
}
