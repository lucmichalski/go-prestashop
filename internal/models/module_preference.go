package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _ModulePreferenceMgr struct {
	*_BaseMgr
}

// ModulePreferenceMgr open func
func ModulePreferenceMgr(db *gorm.DB) *_ModulePreferenceMgr {
	if db == nil {
		panic(fmt.Errorf("ModulePreferenceMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ModulePreferenceMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_module_preference"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ModulePreferenceMgr) GetTableName() string {
	return "ps_module_preference"
}

// Get 获取
func (obj *_ModulePreferenceMgr) Get() (result ModulePreference, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ModulePreferenceMgr) Gets() (results []*ModulePreference, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDModulePreference id_module_preference获取
func (obj *_ModulePreferenceMgr) WithIDModulePreference(idModulePreference int) Option {
	return optionFunc(func(o *options) { o.query["id_module_preference"] = idModulePreference })
}

// WithIDEmployee id_employee获取
func (obj *_ModulePreferenceMgr) WithIDEmployee(idEmployee int) Option {
	return optionFunc(func(o *options) { o.query["id_employee"] = idEmployee })
}

// WithModule module获取
func (obj *_ModulePreferenceMgr) WithModule(module string) Option {
	return optionFunc(func(o *options) { o.query["module"] = module })
}

// WithInterest interest获取
func (obj *_ModulePreferenceMgr) WithInterest(interest bool) Option {
	return optionFunc(func(o *options) { o.query["interest"] = interest })
}

// WithFavorite favorite获取
func (obj *_ModulePreferenceMgr) WithFavorite(favorite bool) Option {
	return optionFunc(func(o *options) { o.query["favorite"] = favorite })
}

// GetByOption 功能选项模式获取
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

// GetByOptions 批量功能选项模式获取
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

//////////////////////////enume case ////////////////////////////////////////////

// GetFromIDModulePreference 通过id_module_preference获取内容
func (obj *_ModulePreferenceMgr) GetFromIDModulePreference(idModulePreference int) (result ModulePreference, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_module_preference = ?", idModulePreference).Find(&result).Error

	return
}

// GetBatchFromIDModulePreference 批量唯一主键查找
func (obj *_ModulePreferenceMgr) GetBatchFromIDModulePreference(idModulePreferences []int) (results []*ModulePreference, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_module_preference IN (?)", idModulePreferences).Find(&results).Error

	return
}

// GetFromIDEmployee 通过id_employee获取内容
func (obj *_ModulePreferenceMgr) GetFromIDEmployee(idEmployee int) (results []*ModulePreference, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee = ?", idEmployee).Find(&results).Error

	return
}

// GetBatchFromIDEmployee 批量唯一主键查找
func (obj *_ModulePreferenceMgr) GetBatchFromIDEmployee(idEmployees []int) (results []*ModulePreference, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee IN (?)", idEmployees).Find(&results).Error

	return
}

// GetFromModule 通过module获取内容
func (obj *_ModulePreferenceMgr) GetFromModule(module string) (results []*ModulePreference, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("module = ?", module).Find(&results).Error

	return
}

// GetBatchFromModule 批量唯一主键查找
func (obj *_ModulePreferenceMgr) GetBatchFromModule(modules []string) (results []*ModulePreference, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("module IN (?)", modules).Find(&results).Error

	return
}

// GetFromInterest 通过interest获取内容
func (obj *_ModulePreferenceMgr) GetFromInterest(interest bool) (results []*ModulePreference, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("interest = ?", interest).Find(&results).Error

	return
}

// GetBatchFromInterest 批量唯一主键查找
func (obj *_ModulePreferenceMgr) GetBatchFromInterest(interests []bool) (results []*ModulePreference, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("interest IN (?)", interests).Find(&results).Error

	return
}

// GetFromFavorite 通过favorite获取内容
func (obj *_ModulePreferenceMgr) GetFromFavorite(favorite bool) (results []*ModulePreference, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("favorite = ?", favorite).Find(&results).Error

	return
}

// GetBatchFromFavorite 批量唯一主键查找
func (obj *_ModulePreferenceMgr) GetBatchFromFavorite(favorites []bool) (results []*ModulePreference, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("favorite IN (?)", favorites).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_ModulePreferenceMgr) FetchByPrimaryKey(idModulePreference int) (result ModulePreference, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_module_preference = ?", idModulePreference).Find(&result).Error

	return
}

// FetchUniqueIndexByEmployeeModule primay or index 获取唯一内容
func (obj *_ModulePreferenceMgr) FetchUniqueIndexByEmployeeModule(idEmployee int, module string) (result ModulePreference, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee = ? AND module = ?", idEmployee, module).Find(&result).Error

	return
}
