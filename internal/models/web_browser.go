package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgWebBrowserMgr struct {
	*_BaseMgr
}

// EgWebBrowserMgr open func
func EgWebBrowserMgr(db *gorm.DB) *_EgWebBrowserMgr {
	if db == nil {
		panic(fmt.Errorf("EgWebBrowserMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgWebBrowserMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_web_browser"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgWebBrowserMgr) GetTableName() string {
	return "eg_web_browser"
}

// Get 获取
func (obj *_EgWebBrowserMgr) Get() (result EgWebBrowser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgWebBrowserMgr) Gets() (results []*EgWebBrowser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDWebBrowser id_web_browser获取 
func (obj *_EgWebBrowserMgr) WithIDWebBrowser(idWebBrowser uint32) Option {
	return optionFunc(func(o *options) { o.query["id_web_browser"] = idWebBrowser })
}

// WithName name获取 
func (obj *_EgWebBrowserMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}


// GetByOption 功能选项模式获取
func (obj *_EgWebBrowserMgr) GetByOption(opts ...Option) (result EgWebBrowser, err error) {
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
func (obj *_EgWebBrowserMgr) GetByOptions(opts ...Option) (results []*EgWebBrowser, err error) {
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


// GetFromIDWebBrowser 通过id_web_browser获取内容  
func (obj *_EgWebBrowserMgr)  GetFromIDWebBrowser(idWebBrowser uint32) (result EgWebBrowser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_web_browser = ?", idWebBrowser).Find(&result).Error
	
	return
}

// GetBatchFromIDWebBrowser 批量唯一主键查找 
func (obj *_EgWebBrowserMgr) GetBatchFromIDWebBrowser(idWebBrowsers []uint32) (results []*EgWebBrowser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_web_browser IN (?)", idWebBrowsers).Find(&results).Error
	
	return
}
 
// GetFromName 通过name获取内容  
func (obj *_EgWebBrowserMgr) GetFromName(name string) (results []*EgWebBrowser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error
	
	return
}

// GetBatchFromName 批量唯一主键查找 
func (obj *_EgWebBrowserMgr) GetBatchFromName(names []string) (results []*EgWebBrowser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgWebBrowserMgr) FetchByPrimaryKey(idWebBrowser uint32 ) (result EgWebBrowser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_web_browser = ?", idWebBrowser).Find(&result).Error
	
	return
}
 

 

	

