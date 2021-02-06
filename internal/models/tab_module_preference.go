package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgTabModulePreferenceMgr struct {
	*_BaseMgr
}

// EgTabModulePreferenceMgr open func
func EgTabModulePreferenceMgr(db *gorm.DB) *_EgTabModulePreferenceMgr {
	if db == nil {
		panic(fmt.Errorf("EgTabModulePreferenceMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgTabModulePreferenceMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_tab_module_preference"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgTabModulePreferenceMgr) GetTableName() string {
	return "eg_tab_module_preference"
}

// Get 获取
func (obj *_EgTabModulePreferenceMgr) Get() (result EgTabModulePreference, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgTabModulePreferenceMgr) Gets() (results []*EgTabModulePreference, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDTabModulePreference id_tab_module_preference获取 
func (obj *_EgTabModulePreferenceMgr) WithIDTabModulePreference(idTabModulePreference int) Option {
	return optionFunc(func(o *options) { o.query["id_tab_module_preference"] = idTabModulePreference })
}

// WithIDEmployee id_employee获取 
func (obj *_EgTabModulePreferenceMgr) WithIDEmployee(idEmployee int) Option {
	return optionFunc(func(o *options) { o.query["id_employee"] = idEmployee })
}

// WithIDTab id_tab获取 
func (obj *_EgTabModulePreferenceMgr) WithIDTab(idTab int) Option {
	return optionFunc(func(o *options) { o.query["id_tab"] = idTab })
}

// WithModule module获取 
func (obj *_EgTabModulePreferenceMgr) WithModule(module string) Option {
	return optionFunc(func(o *options) { o.query["module"] = module })
}


// GetByOption 功能选项模式获取
func (obj *_EgTabModulePreferenceMgr) GetByOption(opts ...Option) (result EgTabModulePreference, err error) {
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
func (obj *_EgTabModulePreferenceMgr) GetByOptions(opts ...Option) (results []*EgTabModulePreference, err error) {
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
func (obj *_EgTabModulePreferenceMgr)  GetFromIDTabModulePreference(idTabModulePreference int) (result EgTabModulePreference, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tab_module_preference = ?", idTabModulePreference).Find(&result).Error
	
	return
}

// GetBatchFromIDTabModulePreference 批量唯一主键查找 
func (obj *_EgTabModulePreferenceMgr) GetBatchFromIDTabModulePreference(idTabModulePreferences []int) (results []*EgTabModulePreference, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tab_module_preference IN (?)", idTabModulePreferences).Find(&results).Error
	
	return
}
 
// GetFromIDEmployee 通过id_employee获取内容  
func (obj *_EgTabModulePreferenceMgr) GetFromIDEmployee(idEmployee int) (results []*EgTabModulePreference, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee = ?", idEmployee).Find(&results).Error
	
	return
}

// GetBatchFromIDEmployee 批量唯一主键查找 
func (obj *_EgTabModulePreferenceMgr) GetBatchFromIDEmployee(idEmployees []int) (results []*EgTabModulePreference, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee IN (?)", idEmployees).Find(&results).Error
	
	return
}
 
// GetFromIDTab 通过id_tab获取内容  
func (obj *_EgTabModulePreferenceMgr) GetFromIDTab(idTab int) (results []*EgTabModulePreference, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tab = ?", idTab).Find(&results).Error
	
	return
}

// GetBatchFromIDTab 批量唯一主键查找 
func (obj *_EgTabModulePreferenceMgr) GetBatchFromIDTab(idTabs []int) (results []*EgTabModulePreference, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tab IN (?)", idTabs).Find(&results).Error
	
	return
}
 
// GetFromModule 通过module获取内容  
func (obj *_EgTabModulePreferenceMgr) GetFromModule(module string) (results []*EgTabModulePreference, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("module = ?", module).Find(&results).Error
	
	return
}

// GetBatchFromModule 批量唯一主键查找 
func (obj *_EgTabModulePreferenceMgr) GetBatchFromModule(modules []string) (results []*EgTabModulePreference, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("module IN (?)", modules).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgTabModulePreferenceMgr) FetchByPrimaryKey(idTabModulePreference int ) (result EgTabModulePreference, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tab_module_preference = ?", idTabModulePreference).Find(&result).Error
	
	return
}
 
 // FetchUniqueIndexByEmployeeModule primay or index 获取唯一内容
 func (obj *_EgTabModulePreferenceMgr) FetchUniqueIndexByEmployeeModule(idEmployee int ,idTab int ,module string ) (result EgTabModulePreference, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee = ? AND id_tab = ? AND module = ?", idEmployee , idTab , module).Find(&result).Error
	
	return
}
 

 

	

