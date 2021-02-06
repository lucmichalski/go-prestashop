package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgAuthorizationRoleMgr struct {
	*_BaseMgr
}

// EgAuthorizationRoleMgr open func
func EgAuthorizationRoleMgr(db *gorm.DB) *_EgAuthorizationRoleMgr {
	if db == nil {
		panic(fmt.Errorf("EgAuthorizationRoleMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgAuthorizationRoleMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_authorization_role"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgAuthorizationRoleMgr) GetTableName() string {
	return "eg_authorization_role"
}

// Get 获取
func (obj *_EgAuthorizationRoleMgr) Get() (result EgAuthorizationRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgAuthorizationRoleMgr) Gets() (results []*EgAuthorizationRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDAuthorizationRole id_authorization_role获取 
func (obj *_EgAuthorizationRoleMgr) WithIDAuthorizationRole(idAuthorizationRole uint32) Option {
	return optionFunc(func(o *options) { o.query["id_authorization_role"] = idAuthorizationRole })
}

// WithSlug slug获取 
func (obj *_EgAuthorizationRoleMgr) WithSlug(slug string) Option {
	return optionFunc(func(o *options) { o.query["slug"] = slug })
}


// GetByOption 功能选项模式获取
func (obj *_EgAuthorizationRoleMgr) GetByOption(opts ...Option) (result EgAuthorizationRole, err error) {
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
func (obj *_EgAuthorizationRoleMgr) GetByOptions(opts ...Option) (results []*EgAuthorizationRole, err error) {
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


// GetFromIDAuthorizationRole 通过id_authorization_role获取内容  
func (obj *_EgAuthorizationRoleMgr)  GetFromIDAuthorizationRole(idAuthorizationRole uint32) (result EgAuthorizationRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_authorization_role = ?", idAuthorizationRole).Find(&result).Error
	
	return
}

// GetBatchFromIDAuthorizationRole 批量唯一主键查找 
func (obj *_EgAuthorizationRoleMgr) GetBatchFromIDAuthorizationRole(idAuthorizationRoles []uint32) (results []*EgAuthorizationRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_authorization_role IN (?)", idAuthorizationRoles).Find(&results).Error
	
	return
}
 
// GetFromSlug 通过slug获取内容  
func (obj *_EgAuthorizationRoleMgr)  GetFromSlug(slug string) (result EgAuthorizationRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("slug = ?", slug).Find(&result).Error
	
	return
}

// GetBatchFromSlug 批量唯一主键查找 
func (obj *_EgAuthorizationRoleMgr) GetBatchFromSlug(slugs []string) (results []*EgAuthorizationRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("slug IN (?)", slugs).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgAuthorizationRoleMgr) FetchByPrimaryKey(idAuthorizationRole uint32 ) (result EgAuthorizationRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_authorization_role = ?", idAuthorizationRole).Find(&result).Error
	
	return
}
 
 // FetchUniqueBySlug primay or index 获取唯一内容
 func (obj *_EgAuthorizationRoleMgr) FetchUniqueBySlug(slug string ) (result EgAuthorizationRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("slug = ?", slug).Find(&result).Error
	
	return
}
 

 

	

