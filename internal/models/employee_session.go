package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _EmployeeSessionMgr struct {
	*_BaseMgr
}

// EmployeeSessionMgr open func
func EmployeeSessionMgr(db *gorm.DB) *_EmployeeSessionMgr {
	if db == nil {
		panic(fmt.Errorf("EmployeeSessionMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EmployeeSessionMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_employee_session"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EmployeeSessionMgr) GetTableName() string {
	return "eg_employee_session"
}

// Get 获取
func (obj *_EmployeeSessionMgr) Get() (result EmployeeSession, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EmployeeSessionMgr) Gets() (results []*EmployeeSession, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDEmployeeSession id_employee_session获取
func (obj *_EmployeeSessionMgr) WithIDEmployeeSession(idEmployeeSession uint32) Option {
	return optionFunc(func(o *options) { o.query["id_employee_session"] = idEmployeeSession })
}

// WithIDEmployee id_employee获取
func (obj *_EmployeeSessionMgr) WithIDEmployee(idEmployee uint32) Option {
	return optionFunc(func(o *options) { o.query["id_employee"] = idEmployee })
}

// WithToken token获取
func (obj *_EmployeeSessionMgr) WithToken(token string) Option {
	return optionFunc(func(o *options) { o.query["token"] = token })
}

// GetByOption 功能选项模式获取
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

// GetByOptions 批量功能选项模式获取
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

//////////////////////////enume case ////////////////////////////////////////////

// GetFromIDEmployeeSession 通过id_employee_session获取内容
func (obj *_EmployeeSessionMgr) GetFromIDEmployeeSession(idEmployeeSession uint32) (result EmployeeSession, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee_session = ?", idEmployeeSession).Find(&result).Error

	return
}

// GetBatchFromIDEmployeeSession 批量唯一主键查找
func (obj *_EmployeeSessionMgr) GetBatchFromIDEmployeeSession(idEmployeeSessions []uint32) (results []*EmployeeSession, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee_session IN (?)", idEmployeeSessions).Find(&results).Error

	return
}

// GetFromIDEmployee 通过id_employee获取内容
func (obj *_EmployeeSessionMgr) GetFromIDEmployee(idEmployee uint32) (results []*EmployeeSession, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee = ?", idEmployee).Find(&results).Error

	return
}

// GetBatchFromIDEmployee 批量唯一主键查找
func (obj *_EmployeeSessionMgr) GetBatchFromIDEmployee(idEmployees []uint32) (results []*EmployeeSession, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee IN (?)", idEmployees).Find(&results).Error

	return
}

// GetFromToken 通过token获取内容
func (obj *_EmployeeSessionMgr) GetFromToken(token string) (results []*EmployeeSession, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("token = ?", token).Find(&results).Error

	return
}

// GetBatchFromToken 批量唯一主键查找
func (obj *_EmployeeSessionMgr) GetBatchFromToken(tokens []string) (results []*EmployeeSession, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("token IN (?)", tokens).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_EmployeeSessionMgr) FetchByPrimaryKey(idEmployeeSession uint32) (result EmployeeSession, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee_session = ?", idEmployeeSession).Find(&result).Error

	return
}
