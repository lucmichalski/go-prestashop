package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgCmsRoleMgr struct {
	*_BaseMgr
}

// EgCmsRoleMgr open func
func EgCmsRoleMgr(db *gorm.DB) *_EgCmsRoleMgr {
	if db == nil {
		panic(fmt.Errorf("EgCmsRoleMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgCmsRoleMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_cms_role"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgCmsRoleMgr) GetTableName() string {
	return "eg_cms_role"
}

// Get 获取
func (obj *_EgCmsRoleMgr) Get() (result EgCmsRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgCmsRoleMgr) Gets() (results []*EgCmsRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDCmsRole id_cms_role获取 
func (obj *_EgCmsRoleMgr) WithIDCmsRole(idCmsRole uint32) Option {
	return optionFunc(func(o *options) { o.query["id_cms_role"] = idCmsRole })
}

// WithName name获取 
func (obj *_EgCmsRoleMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithIDCms id_cms获取 
func (obj *_EgCmsRoleMgr) WithIDCms(idCms uint32) Option {
	return optionFunc(func(o *options) { o.query["id_cms"] = idCms })
}


// GetByOption 功能选项模式获取
func (obj *_EgCmsRoleMgr) GetByOption(opts ...Option) (result EgCmsRole, err error) {
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
func (obj *_EgCmsRoleMgr) GetByOptions(opts ...Option) (results []*EgCmsRole, err error) {
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


// GetFromIDCmsRole 通过id_cms_role获取内容  
func (obj *_EgCmsRoleMgr) GetFromIDCmsRole(idCmsRole uint32) (results []*EgCmsRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cms_role = ?", idCmsRole).Find(&results).Error
	
	return
}

// GetBatchFromIDCmsRole 批量唯一主键查找 
func (obj *_EgCmsRoleMgr) GetBatchFromIDCmsRole(idCmsRoles []uint32) (results []*EgCmsRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cms_role IN (?)", idCmsRoles).Find(&results).Error
	
	return
}
 
// GetFromName 通过name获取内容  
func (obj *_EgCmsRoleMgr)  GetFromName(name string) (result EgCmsRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&result).Error
	
	return
}

// GetBatchFromName 批量唯一主键查找 
func (obj *_EgCmsRoleMgr) GetBatchFromName(names []string) (results []*EgCmsRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error
	
	return
}
 
// GetFromIDCms 通过id_cms获取内容  
func (obj *_EgCmsRoleMgr) GetFromIDCms(idCms uint32) (results []*EgCmsRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cms = ?", idCms).Find(&results).Error
	
	return
}

// GetBatchFromIDCms 批量唯一主键查找 
func (obj *_EgCmsRoleMgr) GetBatchFromIDCms(idCmss []uint32) (results []*EgCmsRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cms IN (?)", idCmss).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgCmsRoleMgr) FetchByPrimaryKey(idCmsRole uint32 ,idCms uint32 ) (result EgCmsRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cms_role = ? AND id_cms = ?", idCmsRole , idCms).Find(&result).Error
	
	return
}
 
 // FetchUniqueByName primay or index 获取唯一内容
 func (obj *_EgCmsRoleMgr) FetchUniqueByName(name string ) (result EgCmsRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&result).Error
	
	return
}
 

 

	
