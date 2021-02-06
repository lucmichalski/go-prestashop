package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgHookMgr struct {
	*_BaseMgr
}

// EgHookMgr open func
func EgHookMgr(db *gorm.DB) *_EgHookMgr {
	if db == nil {
		panic(fmt.Errorf("EgHookMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgHookMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_hook"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgHookMgr) GetTableName() string {
	return "eg_hook"
}

// Get 获取
func (obj *_EgHookMgr) Get() (result EgHook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgHookMgr) Gets() (results []*EgHook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDHook id_hook获取 
func (obj *_EgHookMgr) WithIDHook(idHook uint32) Option {
	return optionFunc(func(o *options) { o.query["id_hook"] = idHook })
}

// WithName name获取 
func (obj *_EgHookMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithTitle title获取 
func (obj *_EgHookMgr) WithTitle(title string) Option {
	return optionFunc(func(o *options) { o.query["title"] = title })
}

// WithDescription description获取 
func (obj *_EgHookMgr) WithDescription(description string) Option {
	return optionFunc(func(o *options) { o.query["description"] = description })
}

// WithPosition position获取 
func (obj *_EgHookMgr) WithPosition(position bool) Option {
	return optionFunc(func(o *options) { o.query["position"] = position })
}


// GetByOption 功能选项模式获取
func (obj *_EgHookMgr) GetByOption(opts ...Option) (result EgHook, err error) {
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
func (obj *_EgHookMgr) GetByOptions(opts ...Option) (results []*EgHook, err error) {
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


// GetFromIDHook 通过id_hook获取内容  
func (obj *_EgHookMgr)  GetFromIDHook(idHook uint32) (result EgHook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_hook = ?", idHook).Find(&result).Error
	
	return
}

// GetBatchFromIDHook 批量唯一主键查找 
func (obj *_EgHookMgr) GetBatchFromIDHook(idHooks []uint32) (results []*EgHook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_hook IN (?)", idHooks).Find(&results).Error
	
	return
}
 
// GetFromName 通过name获取内容  
func (obj *_EgHookMgr)  GetFromName(name string) (result EgHook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&result).Error
	
	return
}

// GetBatchFromName 批量唯一主键查找 
func (obj *_EgHookMgr) GetBatchFromName(names []string) (results []*EgHook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error
	
	return
}
 
// GetFromTitle 通过title获取内容  
func (obj *_EgHookMgr) GetFromTitle(title string) (results []*EgHook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("title = ?", title).Find(&results).Error
	
	return
}

// GetBatchFromTitle 批量唯一主键查找 
func (obj *_EgHookMgr) GetBatchFromTitle(titles []string) (results []*EgHook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("title IN (?)", titles).Find(&results).Error
	
	return
}
 
// GetFromDescription 通过description获取内容  
func (obj *_EgHookMgr) GetFromDescription(description string) (results []*EgHook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("description = ?", description).Find(&results).Error
	
	return
}

// GetBatchFromDescription 批量唯一主键查找 
func (obj *_EgHookMgr) GetBatchFromDescription(descriptions []string) (results []*EgHook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("description IN (?)", descriptions).Find(&results).Error
	
	return
}
 
// GetFromPosition 通过position获取内容  
func (obj *_EgHookMgr) GetFromPosition(position bool) (results []*EgHook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("position = ?", position).Find(&results).Error
	
	return
}

// GetBatchFromPosition 批量唯一主键查找 
func (obj *_EgHookMgr) GetBatchFromPosition(positions []bool) (results []*EgHook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("position IN (?)", positions).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgHookMgr) FetchByPrimaryKey(idHook uint32 ) (result EgHook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_hook = ?", idHook).Find(&result).Error
	
	return
}
 
 // FetchUniqueByHookName primay or index 获取唯一内容
 func (obj *_EgHookMgr) FetchUniqueByHookName(name string ) (result EgHook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&result).Error
	
	return
}
 

 

	

