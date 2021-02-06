package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgWebservicePermissionMgr struct {
	*_BaseMgr
}

// EgWebservicePermissionMgr open func
func EgWebservicePermissionMgr(db *gorm.DB) *_EgWebservicePermissionMgr {
	if db == nil {
		panic(fmt.Errorf("EgWebservicePermissionMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgWebservicePermissionMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_webservice_permission"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgWebservicePermissionMgr) GetTableName() string {
	return "eg_webservice_permission"
}

// Get 获取
func (obj *_EgWebservicePermissionMgr) Get() (result EgWebservicePermission, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgWebservicePermissionMgr) Gets() (results []*EgWebservicePermission, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDWebservicePermission id_webservice_permission获取 
func (obj *_EgWebservicePermissionMgr) WithIDWebservicePermission(idWebservicePermission int) Option {
	return optionFunc(func(o *options) { o.query["id_webservice_permission"] = idWebservicePermission })
}

// WithResource resource获取 
func (obj *_EgWebservicePermissionMgr) WithResource(resource string) Option {
	return optionFunc(func(o *options) { o.query["resource"] = resource })
}

// WithMethod method获取 
func (obj *_EgWebservicePermissionMgr) WithMethod(method string) Option {
	return optionFunc(func(o *options) { o.query["method"] = method })
}

// WithIDWebserviceAccount id_webservice_account获取 
func (obj *_EgWebservicePermissionMgr) WithIDWebserviceAccount(idWebserviceAccount int) Option {
	return optionFunc(func(o *options) { o.query["id_webservice_account"] = idWebserviceAccount })
}


// GetByOption 功能选项模式获取
func (obj *_EgWebservicePermissionMgr) GetByOption(opts ...Option) (result EgWebservicePermission, err error) {
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
func (obj *_EgWebservicePermissionMgr) GetByOptions(opts ...Option) (results []*EgWebservicePermission, err error) {
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


// GetFromIDWebservicePermission 通过id_webservice_permission获取内容  
func (obj *_EgWebservicePermissionMgr)  GetFromIDWebservicePermission(idWebservicePermission int) (result EgWebservicePermission, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_webservice_permission = ?", idWebservicePermission).Find(&result).Error
	
	return
}

// GetBatchFromIDWebservicePermission 批量唯一主键查找 
func (obj *_EgWebservicePermissionMgr) GetBatchFromIDWebservicePermission(idWebservicePermissions []int) (results []*EgWebservicePermission, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_webservice_permission IN (?)", idWebservicePermissions).Find(&results).Error
	
	return
}
 
// GetFromResource 通过resource获取内容  
func (obj *_EgWebservicePermissionMgr) GetFromResource(resource string) (results []*EgWebservicePermission, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("resource = ?", resource).Find(&results).Error
	
	return
}

// GetBatchFromResource 批量唯一主键查找 
func (obj *_EgWebservicePermissionMgr) GetBatchFromResource(resources []string) (results []*EgWebservicePermission, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("resource IN (?)", resources).Find(&results).Error
	
	return
}
 
// GetFromMethod 通过method获取内容  
func (obj *_EgWebservicePermissionMgr) GetFromMethod(method string) (results []*EgWebservicePermission, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("method = ?", method).Find(&results).Error
	
	return
}

// GetBatchFromMethod 批量唯一主键查找 
func (obj *_EgWebservicePermissionMgr) GetBatchFromMethod(methods []string) (results []*EgWebservicePermission, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("method IN (?)", methods).Find(&results).Error
	
	return
}
 
// GetFromIDWebserviceAccount 通过id_webservice_account获取内容  
func (obj *_EgWebservicePermissionMgr) GetFromIDWebserviceAccount(idWebserviceAccount int) (results []*EgWebservicePermission, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_webservice_account = ?", idWebserviceAccount).Find(&results).Error
	
	return
}

// GetBatchFromIDWebserviceAccount 批量唯一主键查找 
func (obj *_EgWebservicePermissionMgr) GetBatchFromIDWebserviceAccount(idWebserviceAccounts []int) (results []*EgWebservicePermission, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_webservice_account IN (?)", idWebserviceAccounts).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgWebservicePermissionMgr) FetchByPrimaryKey(idWebservicePermission int ) (result EgWebservicePermission, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_webservice_permission = ?", idWebservicePermission).Find(&result).Error
	
	return
}
 
 // FetchUniqueIndexByResource2 primay or index 获取唯一内容
 func (obj *_EgWebservicePermissionMgr) FetchUniqueIndexByResource2(resource string ,method string ,idWebserviceAccount int ) (result EgWebservicePermission, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("resource = ? AND method = ? AND id_webservice_account = ?", resource , method , idWebserviceAccount).Find(&result).Error
	
	return
}
 

 
 // FetchIndexByResource  获取多个内容
 func (obj *_EgWebservicePermissionMgr) FetchIndexByResource(resource string ) (results []*EgWebservicePermission, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("resource = ?", resource).Find(&results).Error
	
	return
}
 
 // FetchIndexByMethod  获取多个内容
 func (obj *_EgWebservicePermissionMgr) FetchIndexByMethod(method string ) (results []*EgWebservicePermission, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("method = ?", method).Find(&results).Error
	
	return
}
 
 // FetchIndexByIDWebserviceAccount  获取多个内容
 func (obj *_EgWebservicePermissionMgr) FetchIndexByIDWebserviceAccount(idWebserviceAccount int ) (results []*EgWebservicePermission, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_webservice_account = ?", idWebserviceAccount).Find(&results).Error
	
	return
}
 

	

