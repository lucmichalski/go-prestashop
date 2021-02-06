package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _EmployeeSessionMgr struct {
	*_BaseMgr
}

func EmployeeSessionMgr(db *gorm.DB) *_EmployeeSessionMgr {
	if db == nil {
		panic(fmt.Errorf("EmployeeSessionMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EmployeeSessionMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_employee_session"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_EmployeeSessionMgr) GetTableName() string {
	return "ps_employee_session"
}

func (obj *_EmployeeSessionMgr) Get() (result EmployeeSession, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_EmployeeSessionMgr) Gets() (results []*EmployeeSession, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_EmployeeSessionMgr) WithIDEmployeeSession(idEmployeeSession uint32) Option {
	return optionFunc(func(o *options) { o.query["id_employee_session"] = idEmployeeSession })
}

func (obj *_EmployeeSessionMgr) WithIDEmployee(idEmployee uint32) Option {
	return optionFunc(func(o *options) { o.query["id_employee"] = idEmployee })
}

func (obj *_EmployeeSessionMgr) WithToken(token string) Option {
	return optionFunc(func(o *options) { o.query["token"] = token })
}

func (obj *_EmployeeSessionMgr) GetByOption(opts ...Option) (result EmployeeSession, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_EmployeeSessionMgr) GetByOptions(opts ...Option) (results []*EmployeeSession, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_EmployeeSessionMgr) GetFromIDEmployeeSession(idEmployeeSession uint32) (result EmployeeSession, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee_session = ?", idEmployeeSession).Find(&result).Error

	return
}

func (obj *_EmployeeSessionMgr) GetBatchFromIDEmployeeSession(idEmployeeSessions []uint32) (results []*EmployeeSession, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee_session IN (?)", idEmployeeSessions).Find(&results).Error

	return
}

func (obj *_EmployeeSessionMgr) GetFromIDEmployee(idEmployee uint32) (results []*EmployeeSession, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee = ?", idEmployee).Find(&results).Error

	return
}

func (obj *_EmployeeSessionMgr) GetBatchFromIDEmployee(idEmployees []uint32) (results []*EmployeeSession, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee IN (?)", idEmployees).Find(&results).Error

	return
}

func (obj *_EmployeeSessionMgr) GetFromToken(token string) (results []*EmployeeSession, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("token = ?", token).Find(&results).Error

	return
}

func (obj *_EmployeeSessionMgr) GetBatchFromToken(tokens []string) (results []*EmployeeSession, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("token IN (?)", tokens).Find(&results).Error

	return
}


func (obj *_EmployeeSessionMgr) FetchByPrimaryKey(idEmployeeSession uint32) (result EmployeeSession, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee_session = ?", idEmployeeSession).Find(&result).Error

	return
}
