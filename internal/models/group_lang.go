package	model	
import (	
"gorm.io/gorm"	
"fmt"	
"context"	
)	

type _EgGroupLangMgr struct {
	*_BaseMgr
}

// EgGroupLangMgr open func
func EgGroupLangMgr(db *gorm.DB) *_EgGroupLangMgr {
	if db == nil {
		panic(fmt.Errorf("EgGroupLangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgGroupLangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_group_lang"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgGroupLangMgr) GetTableName() string {
	return "eg_group_lang"
}

// Get 获取
func (obj *_EgGroupLangMgr) Get() (result EgGroupLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgGroupLangMgr) Gets() (results []*EgGroupLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDGroup id_group获取 
func (obj *_EgGroupLangMgr) WithIDGroup(idGroup uint32) Option {
	return optionFunc(func(o *options) { o.query["id_group"] = idGroup })
}

// WithIDLang id_lang获取 
func (obj *_EgGroupLangMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

// WithName name获取 
func (obj *_EgGroupLangMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}


// GetByOption 功能选项模式获取
func (obj *_EgGroupLangMgr) GetByOption(opts ...Option) (result EgGroupLang, err error) {
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
func (obj *_EgGroupLangMgr) GetByOptions(opts ...Option) (results []*EgGroupLang, err error) {
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


// GetFromIDGroup 通过id_group获取内容  
func (obj *_EgGroupLangMgr) GetFromIDGroup(idGroup uint32) (results []*EgGroupLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_group = ?", idGroup).Find(&results).Error
	
	return
}

// GetBatchFromIDGroup 批量唯一主键查找 
func (obj *_EgGroupLangMgr) GetBatchFromIDGroup(idGroups []uint32) (results []*EgGroupLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_group IN (?)", idGroups).Find(&results).Error
	
	return
}
 
// GetFromIDLang 通过id_lang获取内容  
func (obj *_EgGroupLangMgr) GetFromIDLang(idLang uint32) (results []*EgGroupLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error
	
	return
}

// GetBatchFromIDLang 批量唯一主键查找 
func (obj *_EgGroupLangMgr) GetBatchFromIDLang(idLangs []uint32) (results []*EgGroupLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error
	
	return
}
 
// GetFromName 通过name获取内容  
func (obj *_EgGroupLangMgr) GetFromName(name string) (results []*EgGroupLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error
	
	return
}

// GetBatchFromName 批量唯一主键查找 
func (obj *_EgGroupLangMgr) GetBatchFromName(names []string) (results []*EgGroupLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgGroupLangMgr) FetchByPrimaryKey(idGroup uint32 ,idLang uint32 ) (result EgGroupLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_group = ? AND id_lang = ?", idGroup , idLang).Find(&result).Error
	
	return
}
 

 

	
