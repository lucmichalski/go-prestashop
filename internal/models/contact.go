package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _ContactMgr struct {
	*_BaseMgr
}

func ContactMgr(db *gorm.DB) *_ContactMgr {
	if db == nil {
		panic(fmt.Errorf("ContactMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ContactMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_contact"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_ContactMgr) GetTableName() string {
	return "ps_contact"
}

func (obj *_ContactMgr) Get() (result Contact, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_ContactMgr) Gets() (results []*Contact, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

func (obj *_ContactMgr) WithIDContact(idContact uint32) Option {
	return optionFunc(func(o *options) { o.query["id_contact"] = idContact })
}

func (obj *_ContactMgr) WithEmail(email string) Option {
	return optionFunc(func(o *options) { o.query["email"] = email })
}

func (obj *_ContactMgr) WithCustomerService(customerService bool) Option {
	return optionFunc(func(o *options) { o.query["customer_service"] = customerService })
}

func (obj *_ContactMgr) WithPosition(position uint8) Option {
	return optionFunc(func(o *options) { o.query["position"] = position })
}

func (obj *_ContactMgr) GetByOption(opts ...Option) (result Contact, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_ContactMgr) GetByOptions(opts ...Option) (results []*Contact, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

func (obj *_ContactMgr) GetFromIDContact(idContact uint32) (result Contact, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_contact = ?", idContact).Find(&result).Error

	return
}

func (obj *_ContactMgr) GetBatchFromIDContact(idContacts []uint32) (results []*Contact, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_contact IN (?)", idContacts).Find(&results).Error

	return
}

func (obj *_ContactMgr) GetFromEmail(email string) (results []*Contact, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("email = ?", email).Find(&results).Error

	return
}

func (obj *_ContactMgr) GetBatchFromEmail(emails []string) (results []*Contact, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("email IN (?)", emails).Find(&results).Error

	return
}

func (obj *_ContactMgr) GetFromCustomerService(customerService bool) (results []*Contact, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("customer_service = ?", customerService).Find(&results).Error

	return
}

func (obj *_ContactMgr) GetBatchFromCustomerService(customerServices []bool) (results []*Contact, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("customer_service IN (?)", customerServices).Find(&results).Error

	return
}

func (obj *_ContactMgr) GetFromPosition(position uint8) (results []*Contact, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("position = ?", position).Find(&results).Error

	return
}

func (obj *_ContactMgr) GetBatchFromPosition(positions []uint8) (results []*Contact, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("position IN (?)", positions).Find(&results).Error

	return
}

func (obj *_ContactMgr) FetchByPrimaryKey(idContact uint32) (result Contact, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_contact = ?", idContact).Find(&result).Error

	return
}
