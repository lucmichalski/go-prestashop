package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _ModuleHistoryMgr struct {
	*_BaseMgr
}

// ModuleHistoryMgr open func
func ModuleHistoryMgr(db *gorm.DB) *_ModuleHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("ModuleHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ModuleHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_module_history"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ModuleHistoryMgr) GetTableName() string {
	return "eg_module_history"
}

// Get 获取
func (obj *_ModuleHistoryMgr) Get() (result ModuleHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ModuleHistoryMgr) Gets() (results []*ModuleHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_ModuleHistoryMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithIDEmployee id_employee获取
func (obj *_ModuleHistoryMgr) WithIDEmployee(idEmployee int) Option {
	return optionFunc(func(o *options) { o.query["id_employee"] = idEmployee })
}

// WithIDModule id_module获取
func (obj *_ModuleHistoryMgr) WithIDModule(idModule int) Option {
	return optionFunc(func(o *options) { o.query["id_module"] = idModule })
}

// WithDateAdd date_add获取
func (obj *_ModuleHistoryMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

// WithDateUpd date_upd获取
func (obj *_ModuleHistoryMgr) WithDateUpd(dateUpd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_upd"] = dateUpd })
}

// GetByOption 功能选项模式获取
func (obj *_ModuleHistoryMgr) GetByOption(opts ...Option) (result ModuleHistory, err error) {
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
func (obj *_ModuleHistoryMgr) GetByOptions(opts ...Option) (results []*ModuleHistory, err error) {
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

// GetFromID 通过id获取内容
func (obj *_ModuleHistoryMgr) GetFromID(id int) (result ModuleHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_ModuleHistoryMgr) GetBatchFromID(ids []int) (results []*ModuleHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromIDEmployee 通过id_employee获取内容
func (obj *_ModuleHistoryMgr) GetFromIDEmployee(idEmployee int) (results []*ModuleHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee = ?", idEmployee).Find(&results).Error

	return
}

// GetBatchFromIDEmployee 批量唯一主键查找
func (obj *_ModuleHistoryMgr) GetBatchFromIDEmployee(idEmployees []int) (results []*ModuleHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee IN (?)", idEmployees).Find(&results).Error

	return
}

// GetFromIDModule 通过id_module获取内容
func (obj *_ModuleHistoryMgr) GetFromIDModule(idModule int) (results []*ModuleHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_module = ?", idModule).Find(&results).Error

	return
}

// GetBatchFromIDModule 批量唯一主键查找
func (obj *_ModuleHistoryMgr) GetBatchFromIDModule(idModules []int) (results []*ModuleHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_module IN (?)", idModules).Find(&results).Error

	return
}

// GetFromDateAdd 通过date_add获取内容
func (obj *_ModuleHistoryMgr) GetFromDateAdd(dateAdd time.Time) (results []*ModuleHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error

	return
}

// GetBatchFromDateAdd 批量唯一主键查找
func (obj *_ModuleHistoryMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*ModuleHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error

	return
}

// GetFromDateUpd 通过date_upd获取内容
func (obj *_ModuleHistoryMgr) GetFromDateUpd(dateUpd time.Time) (results []*ModuleHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd = ?", dateUpd).Find(&results).Error

	return
}

// GetBatchFromDateUpd 批量唯一主键查找
func (obj *_ModuleHistoryMgr) GetBatchFromDateUpd(dateUpds []time.Time) (results []*ModuleHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd IN (?)", dateUpds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_ModuleHistoryMgr) FetchByPrimaryKey(id int) (result ModuleHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}
