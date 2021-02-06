package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgOrderReturnStateLangMgr struct {
	*_BaseMgr
}

// EgOrderReturnStateLangMgr open func
func EgOrderReturnStateLangMgr(db *gorm.DB) *_EgOrderReturnStateLangMgr {
	if db == nil {
		panic(fmt.Errorf("EgOrderReturnStateLangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgOrderReturnStateLangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_order_return_state_lang"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgOrderReturnStateLangMgr) GetTableName() string {
	return "eg_order_return_state_lang"
}

// Get 获取
func (obj *_EgOrderReturnStateLangMgr) Get() (result EgOrderReturnStateLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgOrderReturnStateLangMgr) Gets() (results []*EgOrderReturnStateLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDOrderReturnState id_order_return_state获取 
func (obj *_EgOrderReturnStateLangMgr) WithIDOrderReturnState(idOrderReturnState uint32) Option {
	return optionFunc(func(o *options) { o.query["id_order_return_state"] = idOrderReturnState })
}

// WithIDLang id_lang获取 
func (obj *_EgOrderReturnStateLangMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

// WithName name获取 
func (obj *_EgOrderReturnStateLangMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}


// GetByOption 功能选项模式获取
func (obj *_EgOrderReturnStateLangMgr) GetByOption(opts ...Option) (result EgOrderReturnStateLang, err error) {
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
func (obj *_EgOrderReturnStateLangMgr) GetByOptions(opts ...Option) (results []*EgOrderReturnStateLang, err error) {
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


// GetFromIDOrderReturnState 通过id_order_return_state获取内容  
func (obj *_EgOrderReturnStateLangMgr) GetFromIDOrderReturnState(idOrderReturnState uint32) (results []*EgOrderReturnStateLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_return_state = ?", idOrderReturnState).Find(&results).Error
	
	return
}

// GetBatchFromIDOrderReturnState 批量唯一主键查找 
func (obj *_EgOrderReturnStateLangMgr) GetBatchFromIDOrderReturnState(idOrderReturnStates []uint32) (results []*EgOrderReturnStateLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_return_state IN (?)", idOrderReturnStates).Find(&results).Error
	
	return
}
 
// GetFromIDLang 通过id_lang获取内容  
func (obj *_EgOrderReturnStateLangMgr) GetFromIDLang(idLang uint32) (results []*EgOrderReturnStateLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error
	
	return
}

// GetBatchFromIDLang 批量唯一主键查找 
func (obj *_EgOrderReturnStateLangMgr) GetBatchFromIDLang(idLangs []uint32) (results []*EgOrderReturnStateLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error
	
	return
}
 
// GetFromName 通过name获取内容  
func (obj *_EgOrderReturnStateLangMgr) GetFromName(name string) (results []*EgOrderReturnStateLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error
	
	return
}

// GetBatchFromName 批量唯一主键查找 
func (obj *_EgOrderReturnStateLangMgr) GetBatchFromName(names []string) (results []*EgOrderReturnStateLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgOrderReturnStateLangMgr) FetchByPrimaryKey(idOrderReturnState uint32 ,idLang uint32 ) (result EgOrderReturnStateLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_return_state = ? AND id_lang = ?", idOrderReturnState , idLang).Find(&result).Error
	
	return
}
 

 

	

