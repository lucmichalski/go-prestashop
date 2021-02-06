package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _TabModulePreferenceMgr struct {
	*_BaseMgr
}

func TabModulePreferenceMgr(db *gorm.DB) *_TabModulePreferenceMgr {
	if db == nil {
		panic(fmt.Errorf("TabModulePreferenceMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_TabModulePreferenceMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_tab_module_preference"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_TabModulePreferenceMgr) GetTableName() string {
	return "ps_tab_module_preference"
}

func (obj *_TabModulePreferenceMgr) Get() (result TabModulePreference, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_TabModulePreferenceMgr) Gets() (results []*TabModulePreference, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_TabModulePreferenceMgr) WithIDTabModulePreference(idTabModulePreference int) Option {
	return optionFunc(func(o *options) { o.query["id_tab_module_preference"] = idTabModulePreference })
}

func (obj *_TabModulePreferenceMgr) WithIDEmployee(idEmployee int) Option {
	return optionFunc(func(o *options) { o.query["id_employee"] = idEmployee })
}

func (obj *_TabModulePreferenceMgr) WithIDTab(idTab int) Option {
	return optionFunc(func(o *options) { o.query["id_tab"] = idTab })
}

func (obj *_TabModulePreferenceMgr) WithModule(module string) Option {
	return optionFunc(func(o *options) { o.query["module"] = module })
}

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


func (obj *_TabModulePreferenceMgr) GetFromIDTabModulePreference(idTabModulePreference int) (result TabModulePreference, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tab_module_preference = ?", idTabModulePreference).Find(&result).Error

	return
}

func (obj *_TabModulePreferenceMgr) GetBatchFromIDTabModulePreference(idTabModulePreferences []int) (results []*TabModulePreference, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tab_module_preference IN (?)", idTabModulePreferences).Find(&results).Error

	return
}

func (obj *_TabModulePreferenceMgr) GetFromIDEmployee(idEmployee int) (results []*TabModulePreference, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee = ?", idEmployee).Find(&results).Error

	return
}

func (obj *_TabModulePreferenceMgr) GetBatchFromIDEmployee(idEmployees []int) (results []*TabModulePreference, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee IN (?)", idEmployees).Find(&results).Error

	return
}

func (obj *_TabModulePreferenceMgr) GetFromIDTab(idTab int) (results []*TabModulePreference, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tab = ?", idTab).Find(&results).Error

	return
}

func (obj *_TabModulePreferenceMgr) GetBatchFromIDTab(idTabs []int) (results []*TabModulePreference, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tab IN (?)", idTabs).Find(&results).Error

	return
}

func (obj *_TabModulePreferenceMgr) GetFromModule(module string) (results []*TabModulePreference, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("module = ?", module).Find(&results).Error

	return
}

func (obj *_TabModulePreferenceMgr) GetBatchFromModule(modules []string) (results []*TabModulePreference, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("module IN (?)", modules).Find(&results).Error

	return
}


func (obj *_TabModulePreferenceMgr) FetchByPrimaryKey(idTabModulePreference int) (result TabModulePreference, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tab_module_preference = ?", idTabModulePreference).Find(&result).Error

	return
}

func (obj *_TabModulePreferenceMgr) FetchUniqueIndexByEmployeeModule(idEmployee int, idTab int, module string) (result TabModulePreference, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee = ? AND id_tab = ? AND module = ?", idEmployee, idTab, module).Find(&result).Error

	return
}
