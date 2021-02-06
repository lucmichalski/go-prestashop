package	model	
import (	
"context"	
"gorm.io/gorm"	
"fmt"	
)	

type _EgContactLangMgr struct {
	*_BaseMgr
}

// EgContactLangMgr open func
func EgContactLangMgr(db *gorm.DB) *_EgContactLangMgr {
	if db == nil {
		panic(fmt.Errorf("EgContactLangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgContactLangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_contact_lang"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgContactLangMgr) GetTableName() string {
	return "eg_contact_lang"
}

// Get 获取
func (obj *_EgContactLangMgr) Get() (result EgContactLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgContactLangMgr) Gets() (results []*EgContactLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDContact id_contact获取 
func (obj *_EgContactLangMgr) WithIDContact(idContact uint32) Option {
	return optionFunc(func(o *options) { o.query["id_contact"] = idContact })
}

// WithIDLang id_lang获取 
func (obj *_EgContactLangMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

// WithName name获取 
func (obj *_EgContactLangMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithDescription description获取 
func (obj *_EgContactLangMgr) WithDescription(description string) Option {
	return optionFunc(func(o *options) { o.query["description"] = description })
}


// GetByOption 功能选项模式获取
func (obj *_EgContactLangMgr) GetByOption(opts ...Option) (result EgContactLang, err error) {
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
func (obj *_EgContactLangMgr) GetByOptions(opts ...Option) (results []*EgContactLang, err error) {
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


// GetFromIDContact 通过id_contact获取内容  
func (obj *_EgContactLangMgr) GetFromIDContact(idContact uint32) (results []*EgContactLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_contact = ?", idContact).Find(&results).Error
	
	return
}

// GetBatchFromIDContact 批量唯一主键查找 
func (obj *_EgContactLangMgr) GetBatchFromIDContact(idContacts []uint32) (results []*EgContactLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_contact IN (?)", idContacts).Find(&results).Error
	
	return
}
 
// GetFromIDLang 通过id_lang获取内容  
func (obj *_EgContactLangMgr) GetFromIDLang(idLang uint32) (results []*EgContactLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error
	
	return
}

// GetBatchFromIDLang 批量唯一主键查找 
func (obj *_EgContactLangMgr) GetBatchFromIDLang(idLangs []uint32) (results []*EgContactLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error
	
	return
}
 
// GetFromName 通过name获取内容  
func (obj *_EgContactLangMgr) GetFromName(name string) (results []*EgContactLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error
	
	return
}

// GetBatchFromName 批量唯一主键查找 
func (obj *_EgContactLangMgr) GetBatchFromName(names []string) (results []*EgContactLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error
	
	return
}
 
// GetFromDescription 通过description获取内容  
func (obj *_EgContactLangMgr) GetFromDescription(description string) (results []*EgContactLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("description = ?", description).Find(&results).Error
	
	return
}

// GetBatchFromDescription 批量唯一主键查找 
func (obj *_EgContactLangMgr) GetBatchFromDescription(descriptions []string) (results []*EgContactLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("description IN (?)", descriptions).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgContactLangMgr) FetchByPrimaryKey(idContact uint32 ,idLang uint32 ) (result EgContactLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_contact = ? AND id_lang = ?", idContact , idLang).Find(&result).Error
	
	return
}
 

 

	

