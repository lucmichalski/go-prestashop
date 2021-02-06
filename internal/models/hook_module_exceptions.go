package	model	
import (	
"context"	
"gorm.io/gorm"	
"fmt"	
)	

type _EgHookModuleExceptionsMgr struct {
	*_BaseMgr
}

// EgHookModuleExceptionsMgr open func
func EgHookModuleExceptionsMgr(db *gorm.DB) *_EgHookModuleExceptionsMgr {
	if db == nil {
		panic(fmt.Errorf("EgHookModuleExceptionsMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgHookModuleExceptionsMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_hook_module_exceptions"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgHookModuleExceptionsMgr) GetTableName() string {
	return "eg_hook_module_exceptions"
}

// Get 获取
func (obj *_EgHookModuleExceptionsMgr) Get() (result EgHookModuleExceptions, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgHookModuleExceptionsMgr) Gets() (results []*EgHookModuleExceptions, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDHookModuleExceptions id_hook_module_exceptions获取 
func (obj *_EgHookModuleExceptionsMgr) WithIDHookModuleExceptions(idHookModuleExceptions uint32) Option {
	return optionFunc(func(o *options) { o.query["id_hook_module_exceptions"] = idHookModuleExceptions })
}

// WithIDShop id_shop获取 
func (obj *_EgHookModuleExceptionsMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

// WithIDModule id_module获取 
func (obj *_EgHookModuleExceptionsMgr) WithIDModule(idModule uint32) Option {
	return optionFunc(func(o *options) { o.query["id_module"] = idModule })
}

// WithIDHook id_hook获取 
func (obj *_EgHookModuleExceptionsMgr) WithIDHook(idHook uint32) Option {
	return optionFunc(func(o *options) { o.query["id_hook"] = idHook })
}

// WithFileName file_name获取 
func (obj *_EgHookModuleExceptionsMgr) WithFileName(fileName string) Option {
	return optionFunc(func(o *options) { o.query["file_name"] = fileName })
}


// GetByOption 功能选项模式获取
func (obj *_EgHookModuleExceptionsMgr) GetByOption(opts ...Option) (result EgHookModuleExceptions, err error) {
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
func (obj *_EgHookModuleExceptionsMgr) GetByOptions(opts ...Option) (results []*EgHookModuleExceptions, err error) {
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


// GetFromIDHookModuleExceptions 通过id_hook_module_exceptions获取内容  
func (obj *_EgHookModuleExceptionsMgr)  GetFromIDHookModuleExceptions(idHookModuleExceptions uint32) (result EgHookModuleExceptions, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_hook_module_exceptions = ?", idHookModuleExceptions).Find(&result).Error
	
	return
}

// GetBatchFromIDHookModuleExceptions 批量唯一主键查找 
func (obj *_EgHookModuleExceptionsMgr) GetBatchFromIDHookModuleExceptions(idHookModuleExceptionss []uint32) (results []*EgHookModuleExceptions, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_hook_module_exceptions IN (?)", idHookModuleExceptionss).Find(&results).Error
	
	return
}
 
// GetFromIDShop 通过id_shop获取内容  
func (obj *_EgHookModuleExceptionsMgr) GetFromIDShop(idShop uint32) (results []*EgHookModuleExceptions, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error
	
	return
}

// GetBatchFromIDShop 批量唯一主键查找 
func (obj *_EgHookModuleExceptionsMgr) GetBatchFromIDShop(idShops []uint32) (results []*EgHookModuleExceptions, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error
	
	return
}
 
// GetFromIDModule 通过id_module获取内容  
func (obj *_EgHookModuleExceptionsMgr) GetFromIDModule(idModule uint32) (results []*EgHookModuleExceptions, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_module = ?", idModule).Find(&results).Error
	
	return
}

// GetBatchFromIDModule 批量唯一主键查找 
func (obj *_EgHookModuleExceptionsMgr) GetBatchFromIDModule(idModules []uint32) (results []*EgHookModuleExceptions, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_module IN (?)", idModules).Find(&results).Error
	
	return
}
 
// GetFromIDHook 通过id_hook获取内容  
func (obj *_EgHookModuleExceptionsMgr) GetFromIDHook(idHook uint32) (results []*EgHookModuleExceptions, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_hook = ?", idHook).Find(&results).Error
	
	return
}

// GetBatchFromIDHook 批量唯一主键查找 
func (obj *_EgHookModuleExceptionsMgr) GetBatchFromIDHook(idHooks []uint32) (results []*EgHookModuleExceptions, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_hook IN (?)", idHooks).Find(&results).Error
	
	return
}
 
// GetFromFileName 通过file_name获取内容  
func (obj *_EgHookModuleExceptionsMgr) GetFromFileName(fileName string) (results []*EgHookModuleExceptions, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("file_name = ?", fileName).Find(&results).Error
	
	return
}

// GetBatchFromFileName 批量唯一主键查找 
func (obj *_EgHookModuleExceptionsMgr) GetBatchFromFileName(fileNames []string) (results []*EgHookModuleExceptions, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("file_name IN (?)", fileNames).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgHookModuleExceptionsMgr) FetchByPrimaryKey(idHookModuleExceptions uint32 ) (result EgHookModuleExceptions, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_hook_module_exceptions = ?", idHookModuleExceptions).Find(&result).Error
	
	return
}
 

 
 // FetchIndexByIDModule  获取多个内容
 func (obj *_EgHookModuleExceptionsMgr) FetchIndexByIDModule(idModule uint32 ) (results []*EgHookModuleExceptions, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_module = ?", idModule).Find(&results).Error
	
	return
}
 
 // FetchIndexByIDHook  获取多个内容
 func (obj *_EgHookModuleExceptionsMgr) FetchIndexByIDHook(idHook uint32 ) (results []*EgHookModuleExceptions, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_hook = ?", idHook).Find(&results).Error
	
	return
}
 

	

