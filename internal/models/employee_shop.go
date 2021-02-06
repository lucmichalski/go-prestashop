package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _EmployeeShopMgr struct {
	*_BaseMgr
}

func EmployeeShopMgr(db *gorm.DB) *_EmployeeShopMgr {
	if db == nil {
		panic(fmt.Errorf("EmployeeShopMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EmployeeShopMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_employee_shop"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_EmployeeShopMgr) GetTableName() string {
	return "ps_employee_shop"
}

func (obj *_EmployeeShopMgr) Get() (result EmployeeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_EmployeeShopMgr) Gets() (results []*EmployeeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_EmployeeShopMgr) WithIDEmployee(idEmployee uint32) Option {
	return optionFunc(func(o *options) { o.query["id_employee"] = idEmployee })
}

func (obj *_EmployeeShopMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

func (obj *_EmployeeShopMgr) GetByOption(opts ...Option) (result EmployeeShop, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_EmployeeShopMgr) GetByOptions(opts ...Option) (results []*EmployeeShop, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_EmployeeShopMgr) GetFromIDEmployee(idEmployee uint32) (results []*EmployeeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee = ?", idEmployee).Find(&results).Error

	return
}

func (obj *_EmployeeShopMgr) GetBatchFromIDEmployee(idEmployees []uint32) (results []*EmployeeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee IN (?)", idEmployees).Find(&results).Error

	return
}

func (obj *_EmployeeShopMgr) GetFromIDShop(idShop uint32) (results []*EmployeeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

func (obj *_EmployeeShopMgr) GetBatchFromIDShop(idShops []uint32) (results []*EmployeeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}


func (obj *_EmployeeShopMgr) FetchByPrimaryKey(idEmployee uint32, idShop uint32) (result EmployeeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee = ? AND id_shop = ?", idEmployee, idShop).Find(&result).Error

	return
}

func (obj *_EmployeeShopMgr) FetchIndexByIDShop(idShop uint32) (results []*EmployeeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}
