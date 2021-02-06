package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _CustomerGroupMgr struct {
	*_BaseMgr
}

// CustomerGroupMgr open func
func CustomerGroupMgr(db *gorm.DB) *_CustomerGroupMgr {
	if db == nil {
		panic(fmt.Errorf("CustomerGroupMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CustomerGroupMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_customer_group"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CustomerGroupMgr) GetTableName() string {
	return "eg_customer_group"
}

// Get 获取
func (obj *_CustomerGroupMgr) Get() (result CustomerGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CustomerGroupMgr) Gets() (results []*CustomerGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDCustomer id_customer获取
func (obj *_CustomerGroupMgr) WithIDCustomer(idCustomer uint32) Option {
	return optionFunc(func(o *options) { o.query["id_customer"] = idCustomer })
}

// WithIDGroup id_group获取
func (obj *_CustomerGroupMgr) WithIDGroup(idGroup uint32) Option {
	return optionFunc(func(o *options) { o.query["id_group"] = idGroup })
}

// GetByOption 功能选项模式获取
func (obj *_CustomerGroupMgr) GetByOption(opts ...Option) (result CustomerGroup, err error) {
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
func (obj *_CustomerGroupMgr) GetByOptions(opts ...Option) (results []*CustomerGroup, err error) {
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

// GetFromIDCustomer 通过id_customer获取内容
func (obj *_CustomerGroupMgr) GetFromIDCustomer(idCustomer uint32) (results []*CustomerGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer = ?", idCustomer).Find(&results).Error

	return
}

// GetBatchFromIDCustomer 批量唯一主键查找
func (obj *_CustomerGroupMgr) GetBatchFromIDCustomer(idCustomers []uint32) (results []*CustomerGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer IN (?)", idCustomers).Find(&results).Error

	return
}

// GetFromIDGroup 通过id_group获取内容
func (obj *_CustomerGroupMgr) GetFromIDGroup(idGroup uint32) (results []*CustomerGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_group = ?", idGroup).Find(&results).Error

	return
}

// GetBatchFromIDGroup 批量唯一主键查找
func (obj *_CustomerGroupMgr) GetBatchFromIDGroup(idGroups []uint32) (results []*CustomerGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_group IN (?)", idGroups).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_CustomerGroupMgr) FetchByPrimaryKey(idCustomer uint32, idGroup uint32) (result CustomerGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer = ? AND id_group = ?", idCustomer, idGroup).Find(&result).Error

	return
}

// FetchIndexByIDCustomer  获取多个内容
func (obj *_CustomerGroupMgr) FetchIndexByIDCustomer(idCustomer uint32) (results []*CustomerGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer = ?", idCustomer).Find(&results).Error

	return
}

// FetchIndexByCustomerLogin  获取多个内容
func (obj *_CustomerGroupMgr) FetchIndexByCustomerLogin(idGroup uint32) (results []*CustomerGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_group = ?", idGroup).Find(&results).Error

	return
}
