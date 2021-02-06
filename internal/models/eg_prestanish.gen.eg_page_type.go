package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgPageTypeMgr struct {
	*_BaseMgr
}

// EgPageTypeMgr open func
func EgPageTypeMgr(db *gorm.DB) *_EgPageTypeMgr {
	if db == nil {
		panic(fmt.Errorf("EgPageTypeMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgPageTypeMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_page_type"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgPageTypeMgr) GetTableName() string {
	return "eg_page_type"
}

// Get 获取
func (obj *_EgPageTypeMgr) Get() (result EgPageType, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgPageTypeMgr) Gets() (results []*EgPageType, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDPageType id_page_type获取 
func (obj *_EgPageTypeMgr) WithIDPageType(idPageType uint32) Option {
	return optionFunc(func(o *options) { o.query["id_page_type"] = idPageType })
}

// WithName name获取 
func (obj *_EgPageTypeMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}


// GetByOption 功能选项模式获取
func (obj *_EgPageTypeMgr) GetByOption(opts ...Option) (result EgPageType, err error) {
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
func (obj *_EgPageTypeMgr) GetByOptions(opts ...Option) (results []*EgPageType, err error) {
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


// GetFromIDPageType 通过id_page_type获取内容  
func (obj *_EgPageTypeMgr)  GetFromIDPageType(idPageType uint32) (result EgPageType, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_page_type = ?", idPageType).Find(&result).Error
	
	return
}

// GetBatchFromIDPageType 批量唯一主键查找 
func (obj *_EgPageTypeMgr) GetBatchFromIDPageType(idPageTypes []uint32) (results []*EgPageType, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_page_type IN (?)", idPageTypes).Find(&results).Error
	
	return
}
 
// GetFromName 通过name获取内容  
func (obj *_EgPageTypeMgr) GetFromName(name string) (results []*EgPageType, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error
	
	return
}

// GetBatchFromName 批量唯一主键查找 
func (obj *_EgPageTypeMgr) GetBatchFromName(names []string) (results []*EgPageType, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgPageTypeMgr) FetchByPrimaryKey(idPageType uint32 ) (result EgPageType, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_page_type = ?", idPageType).Find(&result).Error
	
	return
}
 

 
 // FetchIndexByName  获取多个内容
 func (obj *_EgPageTypeMgr) FetchIndexByName(name string ) (results []*EgPageType, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error
	
	return
}
 

	

