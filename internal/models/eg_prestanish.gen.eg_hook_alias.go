package	model	
import (	
"gorm.io/gorm"	
"fmt"	
"context"	
)	

type _EgHookAliasMgr struct {
	*_BaseMgr
}

// EgHookAliasMgr open func
func EgHookAliasMgr(db *gorm.DB) *_EgHookAliasMgr {
	if db == nil {
		panic(fmt.Errorf("EgHookAliasMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgHookAliasMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_hook_alias"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgHookAliasMgr) GetTableName() string {
	return "eg_hook_alias"
}

// Get 获取
func (obj *_EgHookAliasMgr) Get() (result EgHookAlias, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgHookAliasMgr) Gets() (results []*EgHookAlias, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDHookAlias id_hook_alias获取 
func (obj *_EgHookAliasMgr) WithIDHookAlias(idHookAlias uint32) Option {
	return optionFunc(func(o *options) { o.query["id_hook_alias"] = idHookAlias })
}

// WithAlias alias获取 
func (obj *_EgHookAliasMgr) WithAlias(alias string) Option {
	return optionFunc(func(o *options) { o.query["alias"] = alias })
}

// WithName name获取 
func (obj *_EgHookAliasMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}


// GetByOption 功能选项模式获取
func (obj *_EgHookAliasMgr) GetByOption(opts ...Option) (result EgHookAlias, err error) {
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
func (obj *_EgHookAliasMgr) GetByOptions(opts ...Option) (results []*EgHookAlias, err error) {
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


// GetFromIDHookAlias 通过id_hook_alias获取内容  
func (obj *_EgHookAliasMgr)  GetFromIDHookAlias(idHookAlias uint32) (result EgHookAlias, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_hook_alias = ?", idHookAlias).Find(&result).Error
	
	return
}

// GetBatchFromIDHookAlias 批量唯一主键查找 
func (obj *_EgHookAliasMgr) GetBatchFromIDHookAlias(idHookAliass []uint32) (results []*EgHookAlias, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_hook_alias IN (?)", idHookAliass).Find(&results).Error
	
	return
}
 
// GetFromAlias 通过alias获取内容  
func (obj *_EgHookAliasMgr)  GetFromAlias(alias string) (result EgHookAlias, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("alias = ?", alias).Find(&result).Error
	
	return
}

// GetBatchFromAlias 批量唯一主键查找 
func (obj *_EgHookAliasMgr) GetBatchFromAlias(aliass []string) (results []*EgHookAlias, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("alias IN (?)", aliass).Find(&results).Error
	
	return
}
 
// GetFromName 通过name获取内容  
func (obj *_EgHookAliasMgr) GetFromName(name string) (results []*EgHookAlias, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error
	
	return
}

// GetBatchFromName 批量唯一主键查找 
func (obj *_EgHookAliasMgr) GetBatchFromName(names []string) (results []*EgHookAlias, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgHookAliasMgr) FetchByPrimaryKey(idHookAlias uint32 ) (result EgHookAlias, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_hook_alias = ?", idHookAlias).Find(&result).Error
	
	return
}
 
 // FetchUniqueByAlias primay or index 获取唯一内容
 func (obj *_EgHookAliasMgr) FetchUniqueByAlias(alias string ) (result EgHookAlias, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("alias = ?", alias).Find(&result).Error
	
	return
}
 

 

	

