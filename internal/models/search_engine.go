package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgSearchEngineMgr struct {
	*_BaseMgr
}

// EgSearchEngineMgr open func
func EgSearchEngineMgr(db *gorm.DB) *_EgSearchEngineMgr {
	if db == nil {
		panic(fmt.Errorf("EgSearchEngineMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgSearchEngineMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_search_engine"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgSearchEngineMgr) GetTableName() string {
	return "eg_search_engine"
}

// Get 获取
func (obj *_EgSearchEngineMgr) Get() (result EgSearchEngine, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgSearchEngineMgr) Gets() (results []*EgSearchEngine, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDSearchEngine id_search_engine获取 
func (obj *_EgSearchEngineMgr) WithIDSearchEngine(idSearchEngine uint32) Option {
	return optionFunc(func(o *options) { o.query["id_search_engine"] = idSearchEngine })
}

// WithServer server获取 
func (obj *_EgSearchEngineMgr) WithServer(server string) Option {
	return optionFunc(func(o *options) { o.query["server"] = server })
}

// WithGetvar getvar获取 
func (obj *_EgSearchEngineMgr) WithGetvar(getvar string) Option {
	return optionFunc(func(o *options) { o.query["getvar"] = getvar })
}


// GetByOption 功能选项模式获取
func (obj *_EgSearchEngineMgr) GetByOption(opts ...Option) (result EgSearchEngine, err error) {
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
func (obj *_EgSearchEngineMgr) GetByOptions(opts ...Option) (results []*EgSearchEngine, err error) {
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


// GetFromIDSearchEngine 通过id_search_engine获取内容  
func (obj *_EgSearchEngineMgr)  GetFromIDSearchEngine(idSearchEngine uint32) (result EgSearchEngine, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_search_engine = ?", idSearchEngine).Find(&result).Error
	
	return
}

// GetBatchFromIDSearchEngine 批量唯一主键查找 
func (obj *_EgSearchEngineMgr) GetBatchFromIDSearchEngine(idSearchEngines []uint32) (results []*EgSearchEngine, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_search_engine IN (?)", idSearchEngines).Find(&results).Error
	
	return
}
 
// GetFromServer 通过server获取内容  
func (obj *_EgSearchEngineMgr) GetFromServer(server string) (results []*EgSearchEngine, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("server = ?", server).Find(&results).Error
	
	return
}

// GetBatchFromServer 批量唯一主键查找 
func (obj *_EgSearchEngineMgr) GetBatchFromServer(servers []string) (results []*EgSearchEngine, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("server IN (?)", servers).Find(&results).Error
	
	return
}
 
// GetFromGetvar 通过getvar获取内容  
func (obj *_EgSearchEngineMgr) GetFromGetvar(getvar string) (results []*EgSearchEngine, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("getvar = ?", getvar).Find(&results).Error
	
	return
}

// GetBatchFromGetvar 批量唯一主键查找 
func (obj *_EgSearchEngineMgr) GetBatchFromGetvar(getvars []string) (results []*EgSearchEngine, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("getvar IN (?)", getvars).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgSearchEngineMgr) FetchByPrimaryKey(idSearchEngine uint32 ) (result EgSearchEngine, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_search_engine = ?", idSearchEngine).Find(&result).Error
	
	return
}
 

 

	

