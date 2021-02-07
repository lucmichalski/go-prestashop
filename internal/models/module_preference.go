package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _ModulePreferenceMgr struct {
	*_BaseMgr
}

func ModulePreferenceMgr(db *gorm.DB) *_ModulePreferenceMgr {
	if db == nil {
		panic(fmt.Errorf("ModulePreferenceMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ModulePreferenceMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_module_preference"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_ModulePreferenceMgr) GetTableName() string {
	return "ps_module_preference"
}

func (obj *_ModulePreferenceMgr) Get() (result ModulePreference, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_ModulePreferenceMgr) Gets() (results []*ModulePreference, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

func (obj *_ModulePreferenceMgr) WithIDModulePreference(idModulePreference int) Option {
	return optionFunc(func(o *options) { o.query["id_module_preference"] = idModulePreference })
}

func (obj *_ModulePreferenceMgr) WithIDEmployee(idEmployee int) Option {
	return optionFunc(func(o *options) { o.query["id_employee"] = idEmployee })
}

func (obj *_ModulePreferenceMgr) WithModule(module string) Option {
	return optionFunc(func(o *options) { o.query["module"] = module })
}

func (obj *_ModulePreferenceMgr) WithInterest(interest bool) Option {
	return optionFunc(func(o *options) { o.query["interest"] = interest })
}

func (obj *_ModulePreferenceMgr) WithFavorite(favorite bool) Option {
	return optionFunc(func(o *options) { o.query["favorite"] = favorite })
}

func (obj *_ModulePreferenceMgr) GetByOption(opts ...Option) (result ModulePreference, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_ModulePreferenceMgr) GetByOptions(opts ...Option) (results []*ModulePreference, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

func (obj *_ModulePreferenceMgr) GetFromIDModulePreference(idModulePreference int) (result ModulePreference, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_module_preference = ?", idModulePreference).Find(&result).Error

	return
}

func (obj *_ModulePreferenceMgr) GetBatchFromIDModulePreference(idModulePreferences []int) (results []*ModulePreference, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_module_preference IN (?)", idModulePreferences).Find(&results).Error

	return
}

func (obj *_ModulePreferenceMgr) GetFromIDEmployee(idEmployee int) (results []*ModulePreference, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee = ?", idEmployee).Find(&results).Error

	return
}

func (obj *_ModulePreferenceMgr) GetBatchFromIDEmployee(idEmployees []int) (results []*ModulePreference, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee IN (?)", idEmployees).Find(&results).Error

	return
}

func (obj *_ModulePreferenceMgr) GetFromModule(module string) (results []*ModulePreference, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("module = ?", module).Find(&results).Error

	return
}

func (obj *_ModulePreferenceMgr) GetBatchFromModule(modules []string) (results []*ModulePreference, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("module IN (?)", modules).Find(&results).Error

	return
}

func (obj *_ModulePreferenceMgr) GetFromInterest(interest bool) (results []*ModulePreference, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("interest = ?", interest).Find(&results).Error

	return
}

func (obj *_ModulePreferenceMgr) GetBatchFromInterest(interests []bool) (results []*ModulePreference, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("interest IN (?)", interests).Find(&results).Error

	return
}

func (obj *_ModulePreferenceMgr) GetFromFavorite(favorite bool) (results []*ModulePreference, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("favorite = ?", favorite).Find(&results).Error

	return
}

func (obj *_ModulePreferenceMgr) GetBatchFromFavorite(favorites []bool) (results []*ModulePreference, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("favorite IN (?)", favorites).Find(&results).Error

	return
}

func (obj *_ModulePreferenceMgr) FetchByPrimaryKey(idModulePreference int) (result ModulePreference, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_module_preference = ?", idModulePreference).Find(&result).Error

	return
}

func (obj *_ModulePreferenceMgr) FetchUniqueIndexByEmployeeModule(idEmployee int, module string) (result ModulePreference, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee = ? AND module = ?", idEmployee, module).Find(&result).Error

	return
}
