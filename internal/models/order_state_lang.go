package	model	
import (	
"context"	
"gorm.io/gorm"	
"fmt"	
)	

type _EgOrderStateLangMgr struct {
	*_BaseMgr
}

// EgOrderStateLangMgr open func
func EgOrderStateLangMgr(db *gorm.DB) *_EgOrderStateLangMgr {
	if db == nil {
		panic(fmt.Errorf("EgOrderStateLangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgOrderStateLangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_order_state_lang"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgOrderStateLangMgr) GetTableName() string {
	return "eg_order_state_lang"
}

// Get 获取
func (obj *_EgOrderStateLangMgr) Get() (result EgOrderStateLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgOrderStateLangMgr) Gets() (results []*EgOrderStateLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDOrderState id_order_state获取 
func (obj *_EgOrderStateLangMgr) WithIDOrderState(idOrderState uint32) Option {
	return optionFunc(func(o *options) { o.query["id_order_state"] = idOrderState })
}

// WithIDLang id_lang获取 
func (obj *_EgOrderStateLangMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

// WithName name获取 
func (obj *_EgOrderStateLangMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithTemplate template获取 
func (obj *_EgOrderStateLangMgr) WithTemplate(template string) Option {
	return optionFunc(func(o *options) { o.query["template"] = template })
}


// GetByOption 功能选项模式获取
func (obj *_EgOrderStateLangMgr) GetByOption(opts ...Option) (result EgOrderStateLang, err error) {
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
func (obj *_EgOrderStateLangMgr) GetByOptions(opts ...Option) (results []*EgOrderStateLang, err error) {
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


// GetFromIDOrderState 通过id_order_state获取内容  
func (obj *_EgOrderStateLangMgr) GetFromIDOrderState(idOrderState uint32) (results []*EgOrderStateLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_state = ?", idOrderState).Find(&results).Error
	
	return
}

// GetBatchFromIDOrderState 批量唯一主键查找 
func (obj *_EgOrderStateLangMgr) GetBatchFromIDOrderState(idOrderStates []uint32) (results []*EgOrderStateLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_state IN (?)", idOrderStates).Find(&results).Error
	
	return
}
 
// GetFromIDLang 通过id_lang获取内容  
func (obj *_EgOrderStateLangMgr) GetFromIDLang(idLang uint32) (results []*EgOrderStateLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error
	
	return
}

// GetBatchFromIDLang 批量唯一主键查找 
func (obj *_EgOrderStateLangMgr) GetBatchFromIDLang(idLangs []uint32) (results []*EgOrderStateLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error
	
	return
}
 
// GetFromName 通过name获取内容  
func (obj *_EgOrderStateLangMgr) GetFromName(name string) (results []*EgOrderStateLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error
	
	return
}

// GetBatchFromName 批量唯一主键查找 
func (obj *_EgOrderStateLangMgr) GetBatchFromName(names []string) (results []*EgOrderStateLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error
	
	return
}
 
// GetFromTemplate 通过template获取内容  
func (obj *_EgOrderStateLangMgr) GetFromTemplate(template string) (results []*EgOrderStateLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("template = ?", template).Find(&results).Error
	
	return
}

// GetBatchFromTemplate 批量唯一主键查找 
func (obj *_EgOrderStateLangMgr) GetBatchFromTemplate(templates []string) (results []*EgOrderStateLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("template IN (?)", templates).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgOrderStateLangMgr) FetchByPrimaryKey(idOrderState uint32 ,idLang uint32 ) (result EgOrderStateLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_state = ? AND id_lang = ?", idOrderState , idLang).Find(&result).Error
	
	return
}
 

 

	

