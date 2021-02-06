package	model	
import (	
"context"	
"gorm.io/gorm"	
"fmt"	
)	

type _EgStoreLangMgr struct {
	*_BaseMgr
}

// EgStoreLangMgr open func
func EgStoreLangMgr(db *gorm.DB) *_EgStoreLangMgr {
	if db == nil {
		panic(fmt.Errorf("EgStoreLangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgStoreLangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_store_lang"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgStoreLangMgr) GetTableName() string {
	return "eg_store_lang"
}

// Get 获取
func (obj *_EgStoreLangMgr) Get() (result EgStoreLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgStoreLangMgr) Gets() (results []*EgStoreLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDStore id_store获取 
func (obj *_EgStoreLangMgr) WithIDStore(idStore uint32) Option {
	return optionFunc(func(o *options) { o.query["id_store"] = idStore })
}

// WithIDLang id_lang获取 
func (obj *_EgStoreLangMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

// WithName name获取 
func (obj *_EgStoreLangMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithAddress1 address1获取 
func (obj *_EgStoreLangMgr) WithAddress1(address1 string) Option {
	return optionFunc(func(o *options) { o.query["address1"] = address1 })
}

// WithAddress2 address2获取 
func (obj *_EgStoreLangMgr) WithAddress2(address2 string) Option {
	return optionFunc(func(o *options) { o.query["address2"] = address2 })
}

// WithHours hours获取 
func (obj *_EgStoreLangMgr) WithHours(hours string) Option {
	return optionFunc(func(o *options) { o.query["hours"] = hours })
}

// WithNote note获取 
func (obj *_EgStoreLangMgr) WithNote(note string) Option {
	return optionFunc(func(o *options) { o.query["note"] = note })
}


// GetByOption 功能选项模式获取
func (obj *_EgStoreLangMgr) GetByOption(opts ...Option) (result EgStoreLang, err error) {
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
func (obj *_EgStoreLangMgr) GetByOptions(opts ...Option) (results []*EgStoreLang, err error) {
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


// GetFromIDStore 通过id_store获取内容  
func (obj *_EgStoreLangMgr) GetFromIDStore(idStore uint32) (results []*EgStoreLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_store = ?", idStore).Find(&results).Error
	
	return
}

// GetBatchFromIDStore 批量唯一主键查找 
func (obj *_EgStoreLangMgr) GetBatchFromIDStore(idStores []uint32) (results []*EgStoreLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_store IN (?)", idStores).Find(&results).Error
	
	return
}
 
// GetFromIDLang 通过id_lang获取内容  
func (obj *_EgStoreLangMgr) GetFromIDLang(idLang uint32) (results []*EgStoreLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error
	
	return
}

// GetBatchFromIDLang 批量唯一主键查找 
func (obj *_EgStoreLangMgr) GetBatchFromIDLang(idLangs []uint32) (results []*EgStoreLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error
	
	return
}
 
// GetFromName 通过name获取内容  
func (obj *_EgStoreLangMgr) GetFromName(name string) (results []*EgStoreLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error
	
	return
}

// GetBatchFromName 批量唯一主键查找 
func (obj *_EgStoreLangMgr) GetBatchFromName(names []string) (results []*EgStoreLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error
	
	return
}
 
// GetFromAddress1 通过address1获取内容  
func (obj *_EgStoreLangMgr) GetFromAddress1(address1 string) (results []*EgStoreLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("address1 = ?", address1).Find(&results).Error
	
	return
}

// GetBatchFromAddress1 批量唯一主键查找 
func (obj *_EgStoreLangMgr) GetBatchFromAddress1(address1s []string) (results []*EgStoreLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("address1 IN (?)", address1s).Find(&results).Error
	
	return
}
 
// GetFromAddress2 通过address2获取内容  
func (obj *_EgStoreLangMgr) GetFromAddress2(address2 string) (results []*EgStoreLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("address2 = ?", address2).Find(&results).Error
	
	return
}

// GetBatchFromAddress2 批量唯一主键查找 
func (obj *_EgStoreLangMgr) GetBatchFromAddress2(address2s []string) (results []*EgStoreLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("address2 IN (?)", address2s).Find(&results).Error
	
	return
}
 
// GetFromHours 通过hours获取内容  
func (obj *_EgStoreLangMgr) GetFromHours(hours string) (results []*EgStoreLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("hours = ?", hours).Find(&results).Error
	
	return
}

// GetBatchFromHours 批量唯一主键查找 
func (obj *_EgStoreLangMgr) GetBatchFromHours(hourss []string) (results []*EgStoreLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("hours IN (?)", hourss).Find(&results).Error
	
	return
}
 
// GetFromNote 通过note获取内容  
func (obj *_EgStoreLangMgr) GetFromNote(note string) (results []*EgStoreLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("note = ?", note).Find(&results).Error
	
	return
}

// GetBatchFromNote 批量唯一主键查找 
func (obj *_EgStoreLangMgr) GetBatchFromNote(notes []string) (results []*EgStoreLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("note IN (?)", notes).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgStoreLangMgr) FetchByPrimaryKey(idStore uint32 ,idLang uint32 ) (result EgStoreLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_store = ? AND id_lang = ?", idStore , idLang).Find(&result).Error
	
	return
}
 

 

	

