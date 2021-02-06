package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgPageMgr struct {
	*_BaseMgr
}

// EgPageMgr open func
func EgPageMgr(db *gorm.DB) *_EgPageMgr {
	if db == nil {
		panic(fmt.Errorf("EgPageMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgPageMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_page"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgPageMgr) GetTableName() string {
	return "eg_page"
}

// Get 获取
func (obj *_EgPageMgr) Get() (result EgPage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgPageMgr) Gets() (results []*EgPage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDPage id_page获取 
func (obj *_EgPageMgr) WithIDPage(idPage uint32) Option {
	return optionFunc(func(o *options) { o.query["id_page"] = idPage })
}

// WithIDPageType id_page_type获取 
func (obj *_EgPageMgr) WithIDPageType(idPageType uint32) Option {
	return optionFunc(func(o *options) { o.query["id_page_type"] = idPageType })
}

// WithIDObject id_object获取 
func (obj *_EgPageMgr) WithIDObject(idObject uint32) Option {
	return optionFunc(func(o *options) { o.query["id_object"] = idObject })
}


// GetByOption 功能选项模式获取
func (obj *_EgPageMgr) GetByOption(opts ...Option) (result EgPage, err error) {
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
func (obj *_EgPageMgr) GetByOptions(opts ...Option) (results []*EgPage, err error) {
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


// GetFromIDPage 通过id_page获取内容  
func (obj *_EgPageMgr)  GetFromIDPage(idPage uint32) (result EgPage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_page = ?", idPage).Find(&result).Error
	
	return
}

// GetBatchFromIDPage 批量唯一主键查找 
func (obj *_EgPageMgr) GetBatchFromIDPage(idPages []uint32) (results []*EgPage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_page IN (?)", idPages).Find(&results).Error
	
	return
}
 
// GetFromIDPageType 通过id_page_type获取内容  
func (obj *_EgPageMgr) GetFromIDPageType(idPageType uint32) (results []*EgPage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_page_type = ?", idPageType).Find(&results).Error
	
	return
}

// GetBatchFromIDPageType 批量唯一主键查找 
func (obj *_EgPageMgr) GetBatchFromIDPageType(idPageTypes []uint32) (results []*EgPage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_page_type IN (?)", idPageTypes).Find(&results).Error
	
	return
}
 
// GetFromIDObject 通过id_object获取内容  
func (obj *_EgPageMgr) GetFromIDObject(idObject uint32) (results []*EgPage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_object = ?", idObject).Find(&results).Error
	
	return
}

// GetBatchFromIDObject 批量唯一主键查找 
func (obj *_EgPageMgr) GetBatchFromIDObject(idObjects []uint32) (results []*EgPage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_object IN (?)", idObjects).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgPageMgr) FetchByPrimaryKey(idPage uint32 ) (result EgPage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_page = ?", idPage).Find(&result).Error
	
	return
}
 

 
 // FetchIndexByIDPageType  获取多个内容
 func (obj *_EgPageMgr) FetchIndexByIDPageType(idPageType uint32 ) (results []*EgPage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_page_type = ?", idPageType).Find(&results).Error
	
	return
}
 
 // FetchIndexByIDObject  获取多个内容
 func (obj *_EgPageMgr) FetchIndexByIDObject(idObject uint32 ) (results []*EgPage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_object = ?", idObject).Find(&results).Error
	
	return
}
 

	

