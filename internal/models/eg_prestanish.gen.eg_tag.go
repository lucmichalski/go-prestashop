package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgTagMgr struct {
	*_BaseMgr
}

// EgTagMgr open func
func EgTagMgr(db *gorm.DB) *_EgTagMgr {
	if db == nil {
		panic(fmt.Errorf("EgTagMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgTagMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_tag"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgTagMgr) GetTableName() string {
	return "eg_tag"
}

// Get 获取
func (obj *_EgTagMgr) Get() (result EgTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgTagMgr) Gets() (results []*EgTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDTag id_tag获取 
func (obj *_EgTagMgr) WithIDTag(idTag uint32) Option {
	return optionFunc(func(o *options) { o.query["id_tag"] = idTag })
}

// WithIDLang id_lang获取 
func (obj *_EgTagMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

// WithName name获取 
func (obj *_EgTagMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}


// GetByOption 功能选项模式获取
func (obj *_EgTagMgr) GetByOption(opts ...Option) (result EgTag, err error) {
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
func (obj *_EgTagMgr) GetByOptions(opts ...Option) (results []*EgTag, err error) {
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


// GetFromIDTag 通过id_tag获取内容  
func (obj *_EgTagMgr)  GetFromIDTag(idTag uint32) (result EgTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tag = ?", idTag).Find(&result).Error
	
	return
}

// GetBatchFromIDTag 批量唯一主键查找 
func (obj *_EgTagMgr) GetBatchFromIDTag(idTags []uint32) (results []*EgTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tag IN (?)", idTags).Find(&results).Error
	
	return
}
 
// GetFromIDLang 通过id_lang获取内容  
func (obj *_EgTagMgr) GetFromIDLang(idLang uint32) (results []*EgTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error
	
	return
}

// GetBatchFromIDLang 批量唯一主键查找 
func (obj *_EgTagMgr) GetBatchFromIDLang(idLangs []uint32) (results []*EgTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error
	
	return
}
 
// GetFromName 通过name获取内容  
func (obj *_EgTagMgr) GetFromName(name string) (results []*EgTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error
	
	return
}

// GetBatchFromName 批量唯一主键查找 
func (obj *_EgTagMgr) GetBatchFromName(names []string) (results []*EgTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgTagMgr) FetchByPrimaryKey(idTag uint32 ) (result EgTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tag = ?", idTag).Find(&result).Error
	
	return
}
 

 
 // FetchIndexByIDLang  获取多个内容
 func (obj *_EgTagMgr) FetchIndexByIDLang(idLang uint32 ) (results []*EgTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error
	
	return
}
 
 // FetchIndexByTagName  获取多个内容
 func (obj *_EgTagMgr) FetchIndexByTagName(name string ) (results []*EgTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error
	
	return
}
 

	

