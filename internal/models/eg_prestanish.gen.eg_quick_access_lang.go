package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgQuickAccessLangMgr struct {
	*_BaseMgr
}

// EgQuickAccessLangMgr open func
func EgQuickAccessLangMgr(db *gorm.DB) *_EgQuickAccessLangMgr {
	if db == nil {
		panic(fmt.Errorf("EgQuickAccessLangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgQuickAccessLangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_quick_access_lang"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgQuickAccessLangMgr) GetTableName() string {
	return "eg_quick_access_lang"
}

// Get 获取
func (obj *_EgQuickAccessLangMgr) Get() (result EgQuickAccessLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgQuickAccessLangMgr) Gets() (results []*EgQuickAccessLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDQuickAccess id_quick_access获取 
func (obj *_EgQuickAccessLangMgr) WithIDQuickAccess(idQuickAccess uint32) Option {
	return optionFunc(func(o *options) { o.query["id_quick_access"] = idQuickAccess })
}

// WithIDLang id_lang获取 
func (obj *_EgQuickAccessLangMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

// WithName name获取 
func (obj *_EgQuickAccessLangMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}


// GetByOption 功能选项模式获取
func (obj *_EgQuickAccessLangMgr) GetByOption(opts ...Option) (result EgQuickAccessLang, err error) {
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
func (obj *_EgQuickAccessLangMgr) GetByOptions(opts ...Option) (results []*EgQuickAccessLang, err error) {
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
func (obj *_EgQuickAccessLangMgr) GetFromIDQuickAccess(idQuickAccess uint32) (results []*EgQuickAccessLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_quick_access = ?", idQuickAccess).Find(&results).Error
	
	return
}

// GetBatchFromIDQuickAccess 批量唯一主键查找 
func (obj *_EgQuickAccessLangMgr) GetBatchFromIDQuickAccess(idQuickAccesss []uint32) (results []*EgQuickAccessLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_quick_access IN (?)", idQuickAccesss).Find(&results).Error
	
	return
}
 
// GetFromIDLang 通过id_lang获取内容  
func (obj *_EgQuickAccessLangMgr) GetFromIDLang(idLang uint32) (results []*EgQuickAccessLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error
	
	return
}

// GetBatchFromIDLang 批量唯一主键查找 
func (obj *_EgQuickAccessLangMgr) GetBatchFromIDLang(idLangs []uint32) (results []*EgQuickAccessLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error
	
	return
}
 
// GetFromName 通过name获取内容  
func (obj *_EgQuickAccessLangMgr) GetFromName(name string) (results []*EgQuickAccessLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error
	
	return
}

// GetBatchFromName 批量唯一主键查找 
func (obj *_EgQuickAccessLangMgr) GetBatchFromName(names []string) (results []*EgQuickAccessLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgQuickAccessLangMgr) FetchByPrimaryKey(idQuickAccess uint32 ,idLang uint32 ) (result EgQuickAccessLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_quick_access = ? AND id_lang = ?", idQuickAccess , idLang).Find(&result).Error
	
	return
}
 

 

	

