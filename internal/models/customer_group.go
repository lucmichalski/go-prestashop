package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _CustomerGroupMgr struct {
	*_BaseMgr
}

func CustomerGroupMgr(db *gorm.DB) *_CustomerGroupMgr {
	if db == nil {
		panic(fmt.Errorf("CustomerGroupMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CustomerGroupMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_customer_group"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_CustomerGroupMgr) GetTableName() string {
	return "ps_customer_group"
}

func (obj *_CustomerGroupMgr) Get() (result CustomerGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_CustomerGroupMgr) Gets() (results []*CustomerGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_CustomerGroupMgr) WithIDCustomer(idCustomer uint32) Option {
	return optionFunc(func(o *options) { o.query["id_customer"] = idCustomer })
}

func (obj *_CustomerGroupMgr) WithIDGroup(idGroup uint32) Option {
	return optionFunc(func(o *options) { o.query["id_group"] = idGroup })
}

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


func (obj *_CustomerGroupMgr) GetFromIDCustomer(idCustomer uint32) (results []*CustomerGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer = ?", idCustomer).Find(&results).Error

	return
}

func (obj *_CustomerGroupMgr) GetBatchFromIDCustomer(idCustomers []uint32) (results []*CustomerGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer IN (?)", idCustomers).Find(&results).Error

	return
}

func (obj *_CustomerGroupMgr) GetFromIDGroup(idGroup uint32) (results []*CustomerGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_group = ?", idGroup).Find(&results).Error

	return
}

func (obj *_CustomerGroupMgr) GetBatchFromIDGroup(idGroups []uint32) (results []*CustomerGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_group IN (?)", idGroups).Find(&results).Error

	return
}


func (obj *_CustomerGroupMgr) FetchByPrimaryKey(idCustomer uint32, idGroup uint32) (result CustomerGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer = ? AND id_group = ?", idCustomer, idGroup).Find(&result).Error

	return
}

func (obj *_CustomerGroupMgr) FetchIndexByIDCustomer(idCustomer uint32) (results []*CustomerGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer = ?", idCustomer).Find(&results).Error

	return
}

func (obj *_CustomerGroupMgr) FetchIndexByCustomerLogin(idGroup uint32) (results []*CustomerGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_group = ?", idGroup).Find(&results).Error

	return
}
