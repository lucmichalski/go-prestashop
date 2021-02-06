package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _CustomerSessionMgr struct {
	*_BaseMgr
}

// CustomerSessionMgr open func
func CustomerSessionMgr(db *gorm.DB) *_CustomerSessionMgr {
	if db == nil {
		panic(fmt.Errorf("CustomerSessionMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CustomerSessionMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_customer_session"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CustomerSessionMgr) GetTableName() string {
	return "eg_customer_session"
}

// Get 获取
func (obj *_CustomerSessionMgr) Get() (result CustomerSession, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CustomerSessionMgr) Gets() (results []*CustomerSession, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDCustomerSession id_customer_session获取
func (obj *_CustomerSessionMgr) WithIDCustomerSession(idCustomerSession uint32) Option {
	return optionFunc(func(o *options) { o.query["id_customer_session"] = idCustomerSession })
}

// WithIDCustomer id_customer获取
func (obj *_CustomerSessionMgr) WithIDCustomer(idCustomer uint32) Option {
	return optionFunc(func(o *options) { o.query["id_customer"] = idCustomer })
}

// WithToken token获取
func (obj *_CustomerSessionMgr) WithToken(token string) Option {
	return optionFunc(func(o *options) { o.query["token"] = token })
}

// GetByOption 功能选项模式获取
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

// GetByOptions 批量功能选项模式获取
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

//////////////////////////enume case ////////////////////////////////////////////

// GetFromIDCustomerSession 通过id_customer_session获取内容
func (obj *_CustomerSessionMgr) GetFromIDCustomerSession(idCustomerSession uint32) (result CustomerSession, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer_session = ?", idCustomerSession).Find(&result).Error

	return
}

// GetBatchFromIDCustomerSession 批量唯一主键查找
func (obj *_CustomerSessionMgr) GetBatchFromIDCustomerSession(idCustomerSessions []uint32) (results []*CustomerSession, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer_session IN (?)", idCustomerSessions).Find(&results).Error

	return
}

// GetFromIDCustomer 通过id_customer获取内容
func (obj *_CustomerSessionMgr) GetFromIDCustomer(idCustomer uint32) (results []*CustomerSession, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer = ?", idCustomer).Find(&results).Error

	return
}

// GetBatchFromIDCustomer 批量唯一主键查找
func (obj *_CustomerSessionMgr) GetBatchFromIDCustomer(idCustomers []uint32) (results []*CustomerSession, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer IN (?)", idCustomers).Find(&results).Error

	return
}

// GetFromToken 通过token获取内容
func (obj *_CustomerSessionMgr) GetFromToken(token string) (results []*CustomerSession, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("token = ?", token).Find(&results).Error

	return
}

// GetBatchFromToken 批量唯一主键查找
func (obj *_CustomerSessionMgr) GetBatchFromToken(tokens []string) (results []*CustomerSession, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("token IN (?)", tokens).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_CustomerSessionMgr) FetchByPrimaryKey(idCustomerSession uint32) (result CustomerSession, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer_session = ?", idCustomerSession).Find(&result).Error

	return
}
