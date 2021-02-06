package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgProfileLangMgr struct {
	*_BaseMgr
}

// EgProfileLangMgr open func
func EgProfileLangMgr(db *gorm.DB) *_EgProfileLangMgr {
	if db == nil {
		panic(fmt.Errorf("EgProfileLangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgProfileLangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_profile_lang"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgProfileLangMgr) GetTableName() string {
	return "eg_profile_lang"
}

// Get 获取
func (obj *_EgProfileLangMgr) Get() (result EgProfileLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgProfileLangMgr) Gets() (results []*EgProfileLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDLang id_lang获取 
func (obj *_EgProfileLangMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

// WithIDProfile id_profile获取 
func (obj *_EgProfileLangMgr) WithIDProfile(idProfile uint32) Option {
	return optionFunc(func(o *options) { o.query["id_profile"] = idProfile })
}

// WithName name获取 
func (obj *_EgProfileLangMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}


// GetByOption 功能选项模式获取
func (obj *_EgProfileLangMgr) GetByOption(opts ...Option) (result EgProfileLang, err error) {
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
func (obj *_EgProfileLangMgr) GetByOptions(opts ...Option) (results []*EgProfileLang, err error) {
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


// GetFromIDLang 通过id_lang获取内容  
func (obj *_EgProfileLangMgr) GetFromIDLang(idLang uint32) (results []*EgProfileLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error
	
	return
}

// GetBatchFromIDLang 批量唯一主键查找 
func (obj *_EgProfileLangMgr) GetBatchFromIDLang(idLangs []uint32) (results []*EgProfileLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error
	
	return
}
 
// GetFromIDProfile 通过id_profile获取内容  
func (obj *_EgProfileLangMgr) GetFromIDProfile(idProfile uint32) (results []*EgProfileLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_profile = ?", idProfile).Find(&results).Error
	
	return
}

// GetBatchFromIDProfile 批量唯一主键查找 
func (obj *_EgProfileLangMgr) GetBatchFromIDProfile(idProfiles []uint32) (results []*EgProfileLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_profile IN (?)", idProfiles).Find(&results).Error
	
	return
}
 
// GetFromName 通过name获取内容  
func (obj *_EgProfileLangMgr) GetFromName(name string) (results []*EgProfileLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error
	
	return
}

// GetBatchFromName 批量唯一主键查找 
func (obj *_EgProfileLangMgr) GetBatchFromName(names []string) (results []*EgProfileLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgProfileLangMgr) FetchByPrimaryKey(idLang uint32 ,idProfile uint32 ) (result EgProfileLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ? AND id_profile = ?", idLang , idProfile).Find(&result).Error
	
	return
}
 

 

	

