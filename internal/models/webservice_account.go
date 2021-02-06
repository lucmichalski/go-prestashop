package	model	
import (	
"context"	
"gorm.io/gorm"	
"fmt"	
)	

type _EgWebserviceAccountMgr struct {
	*_BaseMgr
}

// EgWebserviceAccountMgr open func
func EgWebserviceAccountMgr(db *gorm.DB) *_EgWebserviceAccountMgr {
	if db == nil {
		panic(fmt.Errorf("EgWebserviceAccountMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgWebserviceAccountMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_webservice_account"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgWebserviceAccountMgr) GetTableName() string {
	return "eg_webservice_account"
}

// Get 获取
func (obj *_EgWebserviceAccountMgr) Get() (result EgWebserviceAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgWebserviceAccountMgr) Gets() (results []*EgWebserviceAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDWebserviceAccount id_webservice_account获取 
func (obj *_EgWebserviceAccountMgr) WithIDWebserviceAccount(idWebserviceAccount int) Option {
	return optionFunc(func(o *options) { o.query["id_webservice_account"] = idWebserviceAccount })
}

// WithKey key获取 
func (obj *_EgWebserviceAccountMgr) WithKey(key string) Option {
	return optionFunc(func(o *options) { o.query["key"] = key })
}

// WithDescription description获取 
func (obj *_EgWebserviceAccountMgr) WithDescription(description string) Option {
	return optionFunc(func(o *options) { o.query["description"] = description })
}

// WithClassName class_name获取 
func (obj *_EgWebserviceAccountMgr) WithClassName(className string) Option {
	return optionFunc(func(o *options) { o.query["class_name"] = className })
}

// WithIsModule is_module获取 
func (obj *_EgWebserviceAccountMgr) WithIsModule(isModule int8) Option {
	return optionFunc(func(o *options) { o.query["is_module"] = isModule })
}

// WithModuleName module_name获取 
func (obj *_EgWebserviceAccountMgr) WithModuleName(moduleName string) Option {
	return optionFunc(func(o *options) { o.query["module_name"] = moduleName })
}

// WithActive active获取 
func (obj *_EgWebserviceAccountMgr) WithActive(active int8) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}


// GetByOption 功能选项模式获取
func (obj *_EgWebserviceAccountMgr) GetByOption(opts ...Option) (result EgWebserviceAccount, err error) {
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
func (obj *_EgWebserviceAccountMgr) GetByOptions(opts ...Option) (results []*EgWebserviceAccount, err error) {
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


// GetFromIDWebserviceAccount 通过id_webservice_account获取内容  
func (obj *_EgWebserviceAccountMgr)  GetFromIDWebserviceAccount(idWebserviceAccount int) (result EgWebserviceAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_webservice_account = ?", idWebserviceAccount).Find(&result).Error
	
	return
}

// GetBatchFromIDWebserviceAccount 批量唯一主键查找 
func (obj *_EgWebserviceAccountMgr) GetBatchFromIDWebserviceAccount(idWebserviceAccounts []int) (results []*EgWebserviceAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_webservice_account IN (?)", idWebserviceAccounts).Find(&results).Error
	
	return
}
 
// GetFromKey 通过key获取内容  
func (obj *_EgWebserviceAccountMgr) GetFromKey(key string) (results []*EgWebserviceAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("key = ?", key).Find(&results).Error
	
	return
}

// GetBatchFromKey 批量唯一主键查找 
func (obj *_EgWebserviceAccountMgr) GetBatchFromKey(keys []string) (results []*EgWebserviceAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("key IN (?)", keys).Find(&results).Error
	
	return
}
 
// GetFromDescription 通过description获取内容  
func (obj *_EgWebserviceAccountMgr) GetFromDescription(description string) (results []*EgWebserviceAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("description = ?", description).Find(&results).Error
	
	return
}

// GetBatchFromDescription 批量唯一主键查找 
func (obj *_EgWebserviceAccountMgr) GetBatchFromDescription(descriptions []string) (results []*EgWebserviceAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("description IN (?)", descriptions).Find(&results).Error
	
	return
}
 
// GetFromClassName 通过class_name获取内容  
func (obj *_EgWebserviceAccountMgr) GetFromClassName(className string) (results []*EgWebserviceAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("class_name = ?", className).Find(&results).Error
	
	return
}

// GetBatchFromClassName 批量唯一主键查找 
func (obj *_EgWebserviceAccountMgr) GetBatchFromClassName(classNames []string) (results []*EgWebserviceAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("class_name IN (?)", classNames).Find(&results).Error
	
	return
}
 
// GetFromIsModule 通过is_module获取内容  
func (obj *_EgWebserviceAccountMgr) GetFromIsModule(isModule int8) (results []*EgWebserviceAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_module = ?", isModule).Find(&results).Error
	
	return
}

// GetBatchFromIsModule 批量唯一主键查找 
func (obj *_EgWebserviceAccountMgr) GetBatchFromIsModule(isModules []int8) (results []*EgWebserviceAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_module IN (?)", isModules).Find(&results).Error
	
	return
}
 
// GetFromModuleName 通过module_name获取内容  
func (obj *_EgWebserviceAccountMgr) GetFromModuleName(moduleName string) (results []*EgWebserviceAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("module_name = ?", moduleName).Find(&results).Error
	
	return
}

// GetBatchFromModuleName 批量唯一主键查找 
func (obj *_EgWebserviceAccountMgr) GetBatchFromModuleName(moduleNames []string) (results []*EgWebserviceAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("module_name IN (?)", moduleNames).Find(&results).Error
	
	return
}
 
// GetFromActive 通过active获取内容  
func (obj *_EgWebserviceAccountMgr) GetFromActive(active int8) (results []*EgWebserviceAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active = ?", active).Find(&results).Error
	
	return
}

// GetBatchFromActive 批量唯一主键查找 
func (obj *_EgWebserviceAccountMgr) GetBatchFromActive(actives []int8) (results []*EgWebserviceAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active IN (?)", actives).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgWebserviceAccountMgr) FetchByPrimaryKey(idWebserviceAccount int ) (result EgWebserviceAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_webservice_account = ?", idWebserviceAccount).Find(&result).Error
	
	return
}
 

 
 // FetchIndexByKey  获取多个内容
 func (obj *_EgWebserviceAccountMgr) FetchIndexByKey(key string ) (results []*EgWebserviceAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("key = ?", key).Find(&results).Error
	
	return
}
 

	

