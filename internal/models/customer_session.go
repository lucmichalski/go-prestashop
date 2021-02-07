package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _CustomerSessionMgr struct {
	*_BaseMgr
}

func CustomerSessionMgr(db *gorm.DB) *_CustomerSessionMgr {
	if db == nil {
		panic(fmt.Errorf("CustomerSessionMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CustomerSessionMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_customer_session"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_CustomerSessionMgr) GetTableName() string {
	return "ps_customer_session"
}

func (obj *_CustomerSessionMgr) Get() (result CustomerSession, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_CustomerSessionMgr) Gets() (results []*CustomerSession, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

func (obj *_CustomerSessionMgr) WithIDCustomerSession(idCustomerSession uint32) Option {
	return optionFunc(func(o *options) { o.query["id_customer_session"] = idCustomerSession })
}

func (obj *_CustomerSessionMgr) WithIDCustomer(idCustomer uint32) Option {
	return optionFunc(func(o *options) { o.query["id_customer"] = idCustomer })
}

func (obj *_CustomerSessionMgr) WithToken(token string) Option {
	return optionFunc(func(o *options) { o.query["token"] = token })
}

func (obj *_CustomerSessionMgr) GetByOption(opts ...Option) (result CustomerSession, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_CustomerSessionMgr) GetByOptions(opts ...Option) (results []*CustomerSession, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

func (obj *_CustomerSessionMgr) GetFromIDCustomerSession(idCustomerSession uint32) (result CustomerSession, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer_session = ?", idCustomerSession).Find(&result).Error

	return
}

func (obj *_CustomerSessionMgr) GetBatchFromIDCustomerSession(idCustomerSessions []uint32) (results []*CustomerSession, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer_session IN (?)", idCustomerSessions).Find(&results).Error

	return
}

func (obj *_CustomerSessionMgr) GetFromIDCustomer(idCustomer uint32) (results []*CustomerSession, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer = ?", idCustomer).Find(&results).Error

	return
}

func (obj *_CustomerSessionMgr) GetBatchFromIDCustomer(idCustomers []uint32) (results []*CustomerSession, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer IN (?)", idCustomers).Find(&results).Error

	return
}

func (obj *_CustomerSessionMgr) GetFromToken(token string) (results []*CustomerSession, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("token = ?", token).Find(&results).Error

	return
}

func (obj *_CustomerSessionMgr) GetBatchFromToken(tokens []string) (results []*CustomerSession, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("token IN (?)", tokens).Find(&results).Error

	return
}

func (obj *_CustomerSessionMgr) FetchByPrimaryKey(idCustomerSession uint32) (result CustomerSession, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer_session = ?", idCustomerSession).Find(&result).Error

	return
}
