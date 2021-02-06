package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _TabModulePreferenceMgr struct {
	*_BaseMgr
}

// TabModulePreferenceMgr open func
func TabModulePreferenceMgr(db *gorm.DB) *_TabModulePreferenceMgr {
	if db == nil {
		panic(fmt.Errorf("TabModulePreferenceMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_TabModulePreferenceMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_tab_module_preference"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_TabModulePreferenceMgr) GetTableName() string {
	return "eg_tab_module_preference"
}

// Get 获取
func (obj *_TabModulePreferenceMgr) Get() (result TabModulePreference, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_TabModulePreferenceMgr) Gets() (results []*TabModulePreference, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDTabModulePreference id_tab_module_preference获取
func (obj *_TabModulePreferenceMgr) WithIDTabModulePreference(idTabModulePreference int) Option {
	return optionFunc(func(o *options) { o.query["id_tab_module_preference"] = idTabModulePreference })
}

// WithIDEmployee id_employee获取
func (obj *_TabModulePreferenceMgr) WithIDEmployee(idEmployee int) Option {
	return optionFunc(func(o *options) { o.query["id_employee"] = idEmployee })
}

// WithIDTab id_tab获取
func (obj *_TabModulePreferenceMgr) WithIDTab(idTab int) Option {
	return optionFunc(func(o *options) { o.query["id_tab"] = idTab })
}

// WithModule module获取
func (obj *_TabModulePreferenceMgr) WithModule(module string) Option {
	return optionFunc(func(o *options) { o.query["module"] = module })
}

// GetByOption 功能选项模式获取
func (obj *_TabModulePreferenceMgr) GetByOption(opts ...Option) (result TabModulePreference, err error) {
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
func (obj *_TabModulePreferenceMgr) GetByOptions(opts ...Option) (results []*TabModulePreference, err error) {
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

// GetFromIDTabModulePreference 通过id_tab_module_preference获取内容
func (obj *_TabModulePreferenceMgr) GetFromIDTabModulePreference(idTabModulePreference int) (result TabModulePreference, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tab_module_preference = ?", idTabModulePreference).Find(&result).Error

	return
}

// GetBatchFromIDTabModulePreference 批量唯一主键查找
func (obj *_TabModulePreferenceMgr) GetBatchFromIDTabModulePreference(idTabModulePreferences []int) (results []*TabModulePreference, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tab_module_preference IN (?)", idTabModulePreferences).Find(&results).Error

	return
}

// GetFromIDEmployee 通过id_employee获取内容
func (obj *_TabModulePreferenceMgr) GetFromIDEmployee(idEmployee int) (results []*TabModulePreference, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee = ?", idEmployee).Find(&results).Error

	return
}

// GetBatchFromIDEmployee 批量唯一主键查找
func (obj *_TabModulePreferenceMgr) GetBatchFromIDEmployee(idEmployees []int) (results []*TabModulePreference, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee IN (?)", idEmployees).Find(&results).Error

	return
}

// GetFromIDTab 通过id_tab获取内容
func (obj *_TabModulePreferenceMgr) GetFromIDTab(idTab int) (results []*TabModulePreference, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tab = ?", idTab).Find(&results).Error

	return
}

// GetBatchFromIDTab 批量唯一主键查找
func (obj *_TabModulePreferenceMgr) GetBatchFromIDTab(idTabs []int) (results []*TabModulePreference, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tab IN (?)", idTabs).Find(&results).Error

	return
}

// GetFromModule 通过module获取内容
func (obj *_TabModulePreferenceMgr) GetFromModule(module string) (results []*TabModulePreference, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("module = ?", module).Find(&results).Error

	return
}

// GetBatchFromModule 批量唯一主键查找
func (obj *_TabModulePreferenceMgr) GetBatchFromModule(modules []string) (results []*TabModulePreference, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("module IN (?)", modules).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_TabModulePreferenceMgr) FetchByPrimaryKey(idTabModulePreference int) (result TabModulePreference, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tab_module_preference = ?", idTabModulePreference).Find(&result).Error

	return
}

// FetchUniqueIndexByEmployeeModule primay or index 获取唯一内容
func (obj *_TabModulePreferenceMgr) FetchUniqueIndexByEmployeeModule(idEmployee int, idTab int, module string) (result TabModulePreference, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee = ? AND id_tab = ? AND module = ?", idEmployee, idTab, module).Find(&result).Error

	return
}
