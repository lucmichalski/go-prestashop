package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgQuickAccessMgr struct {
	*_BaseMgr
}

// EgQuickAccessMgr open func
func EgQuickAccessMgr(db *gorm.DB) *_EgQuickAccessMgr {
	if db == nil {
		panic(fmt.Errorf("EgQuickAccessMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgQuickAccessMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_quick_access"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgQuickAccessMgr) GetTableName() string {
	return "eg_quick_access"
}

// Get 获取
func (obj *_EgQuickAccessMgr) Get() (result EgQuickAccess, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgQuickAccessMgr) Gets() (results []*EgQuickAccess, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDQuickAccess id_quick_access获取 
func (obj *_EgQuickAccessMgr) WithIDQuickAccess(idQuickAccess uint32) Option {
	return optionFunc(func(o *options) { o.query["id_quick_access"] = idQuickAccess })
}

// WithNewWindow new_window获取 
func (obj *_EgQuickAccessMgr) WithNewWindow(newWindow bool) Option {
	return optionFunc(func(o *options) { o.query["new_window"] = newWindow })
}

// WithLink link获取 
func (obj *_EgQuickAccessMgr) WithLink(link string) Option {
	return optionFunc(func(o *options) { o.query["link"] = link })
}


// GetByOption 功能选项模式获取
func (obj *_EgQuickAccessMgr) GetByOption(opts ...Option) (result EgQuickAccess, err error) {
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
func (obj *_EgQuickAccessMgr) GetByOptions(opts ...Option) (results []*EgQuickAccess, err error) {
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


// GetFromIDQuickAccess 通过id_quick_access获取内容  
func (obj *_EgQuickAccessMgr)  GetFromIDQuickAccess(idQuickAccess uint32) (result EgQuickAccess, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_quick_access = ?", idQuickAccess).Find(&result).Error
	
	return
}

// GetBatchFromIDQuickAccess 批量唯一主键查找 
func (obj *_EgQuickAccessMgr) GetBatchFromIDQuickAccess(idQuickAccesss []uint32) (results []*EgQuickAccess, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_quick_access IN (?)", idQuickAccesss).Find(&results).Error
	
	return
}
 
// GetFromNewWindow 通过new_window获取内容  
func (obj *_EgQuickAccessMgr) GetFromNewWindow(newWindow bool) (results []*EgQuickAccess, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("new_window = ?", newWindow).Find(&results).Error
	
	return
}

// GetBatchFromNewWindow 批量唯一主键查找 
func (obj *_EgQuickAccessMgr) GetBatchFromNewWindow(newWindows []bool) (results []*EgQuickAccess, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("new_window IN (?)", newWindows).Find(&results).Error
	
	return
}
 
// GetFromLink 通过link获取内容  
func (obj *_EgQuickAccessMgr) GetFromLink(link string) (results []*EgQuickAccess, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("link = ?", link).Find(&results).Error
	
	return
}

// GetBatchFromLink 批量唯一主键查找 
func (obj *_EgQuickAccessMgr) GetBatchFromLink(links []string) (results []*EgQuickAccess, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("link IN (?)", links).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgQuickAccessMgr) FetchByPrimaryKey(idQuickAccess uint32 ) (result EgQuickAccess, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_quick_access = ?", idQuickAccess).Find(&result).Error
	
	return
}
 

 

	

