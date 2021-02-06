package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _ContactMgr struct {
	*_BaseMgr
}

// ContactMgr open func
func ContactMgr(db *gorm.DB) *_ContactMgr {
	if db == nil {
		panic(fmt.Errorf("ContactMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ContactMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_contact"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ContactMgr) GetTableName() string {
	return "eg_contact"
}

// Get 获取
func (obj *_ContactMgr) Get() (result Contact, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ContactMgr) Gets() (results []*Contact, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDContact id_contact获取
func (obj *_ContactMgr) WithIDContact(idContact uint32) Option {
	return optionFunc(func(o *options) { o.query["id_contact"] = idContact })
}

// WithEmail email获取
func (obj *_ContactMgr) WithEmail(email string) Option {
	return optionFunc(func(o *options) { o.query["email"] = email })
}

// WithCustomerService customer_service获取
func (obj *_ContactMgr) WithCustomerService(customerService bool) Option {
	return optionFunc(func(o *options) { o.query["customer_service"] = customerService })
}

// WithPosition position获取
func (obj *_ContactMgr) WithPosition(position uint8) Option {
	return optionFunc(func(o *options) { o.query["position"] = position })
}

// GetByOption 功能选项模式获取
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

// GetByOptions 批量功能选项模式获取
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

//////////////////////////enume case ////////////////////////////////////////////

// GetFromIDContact 通过id_contact获取内容
func (obj *_ContactMgr) GetFromIDContact(idContact uint32) (result Contact, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_contact = ?", idContact).Find(&result).Error

	return
}

// GetBatchFromIDContact 批量唯一主键查找
func (obj *_ContactMgr) GetBatchFromIDContact(idContacts []uint32) (results []*Contact, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_contact IN (?)", idContacts).Find(&results).Error

	return
}

// GetFromEmail 通过email获取内容
func (obj *_ContactMgr) GetFromEmail(email string) (results []*Contact, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("email = ?", email).Find(&results).Error

	return
}

// GetBatchFromEmail 批量唯一主键查找
func (obj *_ContactMgr) GetBatchFromEmail(emails []string) (results []*Contact, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("email IN (?)", emails).Find(&results).Error

	return
}

// GetFromCustomerService 通过customer_service获取内容
func (obj *_ContactMgr) GetFromCustomerService(customerService bool) (results []*Contact, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("customer_service = ?", customerService).Find(&results).Error

	return
}

// GetBatchFromCustomerService 批量唯一主键查找
func (obj *_ContactMgr) GetBatchFromCustomerService(customerServices []bool) (results []*Contact, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("customer_service IN (?)", customerServices).Find(&results).Error

	return
}

// GetFromPosition 通过position获取内容
func (obj *_ContactMgr) GetFromPosition(position uint8) (results []*Contact, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("position = ?", position).Find(&results).Error

	return
}

// GetBatchFromPosition 批量唯一主键查找
func (obj *_ContactMgr) GetBatchFromPosition(positions []uint8) (results []*Contact, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("position IN (?)", positions).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_ContactMgr) FetchByPrimaryKey(idContact uint32) (result Contact, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_contact = ?", idContact).Find(&result).Error

	return
}
