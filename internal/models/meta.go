package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgMetaMgr struct {
	*_BaseMgr
}

// EgMetaMgr open func
func EgMetaMgr(db *gorm.DB) *_EgMetaMgr {
	if db == nil {
		panic(fmt.Errorf("EgMetaMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgMetaMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_meta"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgMetaMgr) GetTableName() string {
	return "eg_meta"
}

// Get 获取
func (obj *_EgMetaMgr) Get() (result EgMeta, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgMetaMgr) Gets() (results []*EgMeta, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDMeta id_meta获取 
func (obj *_EgMetaMgr) WithIDMeta(idMeta uint32) Option {
	return optionFunc(func(o *options) { o.query["id_meta"] = idMeta })
}

// WithPage page获取 
func (obj *_EgMetaMgr) WithPage(page string) Option {
	return optionFunc(func(o *options) { o.query["page"] = page })
}

// WithConfigurable configurable获取 
func (obj *_EgMetaMgr) WithConfigurable(configurable bool) Option {
	return optionFunc(func(o *options) { o.query["configurable"] = configurable })
}


// GetByOption 功能选项模式获取
func (obj *_EgMetaMgr) GetByOption(opts ...Option) (result EgMeta, err error) {
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
func (obj *_EgMetaMgr) GetByOptions(opts ...Option) (results []*EgMeta, err error) {
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


// GetFromIDMeta 通过id_meta获取内容  
func (obj *_EgMetaMgr)  GetFromIDMeta(idMeta uint32) (result EgMeta, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_meta = ?", idMeta).Find(&result).Error
	
	return
}

// GetBatchFromIDMeta 批量唯一主键查找 
func (obj *_EgMetaMgr) GetBatchFromIDMeta(idMetas []uint32) (results []*EgMeta, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_meta IN (?)", idMetas).Find(&results).Error
	
	return
}
 
// GetFromPage 通过page获取内容  
func (obj *_EgMetaMgr)  GetFromPage(page string) (result EgMeta, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("page = ?", page).Find(&result).Error
	
	return
}

// GetBatchFromPage 批量唯一主键查找 
func (obj *_EgMetaMgr) GetBatchFromPage(pages []string) (results []*EgMeta, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("page IN (?)", pages).Find(&results).Error
	
	return
}
 
// GetFromConfigurable 通过configurable获取内容  
func (obj *_EgMetaMgr) GetFromConfigurable(configurable bool) (results []*EgMeta, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("configurable = ?", configurable).Find(&results).Error
	
	return
}

// GetBatchFromConfigurable 批量唯一主键查找 
func (obj *_EgMetaMgr) GetBatchFromConfigurable(configurables []bool) (results []*EgMeta, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("configurable IN (?)", configurables).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgMetaMgr) FetchByPrimaryKey(idMeta uint32 ) (result EgMeta, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_meta = ?", idMeta).Find(&result).Error
	
	return
}
 
 // FetchUniqueByPage primay or index 获取唯一内容
 func (obj *_EgMetaMgr) FetchUniqueByPage(page string ) (result EgMeta, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("page = ?", page).Find(&result).Error
	
	return
}
 

 

	

