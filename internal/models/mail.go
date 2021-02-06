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

func MailMgr(db *gorm.DB) *_MailMgr {
	if db == nil {
		panic(fmt.Errorf("MailMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_MailMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_mail"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_MailMgr) GetTableName() string {
	return "ps_mail"
}

func (obj *_MailMgr) Get() (result Mail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_MailMgr) Gets() (results []*Mail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_MailMgr) WithIDMail(idMail uint32) Option {
	return optionFunc(func(o *options) { o.query["id_mail"] = idMail })
}

func (obj *_MailMgr) WithRecipient(recipient string) Option {
	return optionFunc(func(o *options) { o.query["recipient"] = recipient })
}

func (obj *_MailMgr) WithTemplate(template string) Option {
	return optionFunc(func(o *options) { o.query["template"] = template })
}

func (obj *_MailMgr) WithSubject(subject string) Option {
	return optionFunc(func(o *options) { o.query["subject"] = subject })
}

func (obj *_MailMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

func (obj *_MailMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

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


func (obj *_MailMgr) GetFromIDMail(idMail uint32) (result Mail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_mail = ?", idMail).Find(&result).Error

	return
}

func (obj *_MailMgr) GetBatchFromIDMail(idMails []uint32) (results []*Mail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_mail IN (?)", idMails).Find(&results).Error

	return
}

func (obj *_MailMgr) GetFromRecipient(recipient string) (results []*Mail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("recipient = ?", recipient).Find(&results).Error

	return
}

func (obj *_MailMgr) GetBatchFromRecipient(recipients []string) (results []*Mail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("recipient IN (?)", recipients).Find(&results).Error

	return
}

func (obj *_MailMgr) GetFromTemplate(template string) (results []*Mail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("template = ?", template).Find(&results).Error

	return
}

func (obj *_MailMgr) GetBatchFromTemplate(templates []string) (results []*Mail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("template IN (?)", templates).Find(&results).Error

	return
}

func (obj *_MailMgr) GetFromSubject(subject string) (results []*Mail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("subject = ?", subject).Find(&results).Error

	return
}

func (obj *_MailMgr) GetBatchFromSubject(subjects []string) (results []*Mail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("subject IN (?)", subjects).Find(&results).Error

	return
}

func (obj *_MailMgr) GetFromIDLang(idLang uint32) (results []*Mail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error

	return
}

func (obj *_MailMgr) GetBatchFromIDLang(idLangs []uint32) (results []*Mail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error

	return
}

func (obj *_MailMgr) GetFromDateAdd(dateAdd time.Time) (results []*Mail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error

	return
}

func (obj *_MailMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*Mail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error

	return
}


func (obj *_MailMgr) FetchByPrimaryKey(idMail uint32) (result Mail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_mail = ?", idMail).Find(&result).Error

	return
}

func (obj *_MailMgr) FetchIndexByRecipient(recipient string) (results []*Mail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("recipient = ?", recipient).Find(&results).Error

	return
}
