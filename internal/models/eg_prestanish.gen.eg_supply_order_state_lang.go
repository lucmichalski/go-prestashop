package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgSupplyOrderStateLangMgr struct {
	*_BaseMgr
}

// EgSupplyOrderStateLangMgr open func
func EgSupplyOrderStateLangMgr(db *gorm.DB) *_EgSupplyOrderStateLangMgr {
	if db == nil {
		panic(fmt.Errorf("EgSupplyOrderStateLangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgSupplyOrderStateLangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_supply_order_state_lang"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgSupplyOrderStateLangMgr) GetTableName() string {
	return "eg_supply_order_state_lang"
}

// Get 获取
func (obj *_EgSupplyOrderStateLangMgr) Get() (result EgSupplyOrderStateLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgSupplyOrderStateLangMgr) Gets() (results []*EgSupplyOrderStateLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDSupplyOrderState id_supply_order_state获取 
func (obj *_EgSupplyOrderStateLangMgr) WithIDSupplyOrderState(idSupplyOrderState uint32) Option {
	return optionFunc(func(o *options) { o.query["id_supply_order_state"] = idSupplyOrderState })
}

// WithIDLang id_lang获取 
func (obj *_EgSupplyOrderStateLangMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

// WithName name获取 
func (obj *_EgSupplyOrderStateLangMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}


// GetByOption 功能选项模式获取
func (obj *_EgSupplyOrderStateLangMgr) GetByOption(opts ...Option) (result EgSupplyOrderStateLang, err error) {
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
func (obj *_EgSupplyOrderStateLangMgr) GetByOptions(opts ...Option) (results []*EgSupplyOrderStateLang, err error) {
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


// GetFromIDSupplyOrderState 通过id_supply_order_state获取内容  
func (obj *_EgSupplyOrderStateLangMgr) GetFromIDSupplyOrderState(idSupplyOrderState uint32) (results []*EgSupplyOrderStateLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supply_order_state = ?", idSupplyOrderState).Find(&results).Error
	
	return
}

// GetBatchFromIDSupplyOrderState 批量唯一主键查找 
func (obj *_EgSupplyOrderStateLangMgr) GetBatchFromIDSupplyOrderState(idSupplyOrderStates []uint32) (results []*EgSupplyOrderStateLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supply_order_state IN (?)", idSupplyOrderStates).Find(&results).Error
	
	return
}
 
// GetFromIDLang 通过id_lang获取内容  
func (obj *_EgSupplyOrderStateLangMgr) GetFromIDLang(idLang uint32) (results []*EgSupplyOrderStateLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error
	
	return
}

// GetBatchFromIDLang 批量唯一主键查找 
func (obj *_EgSupplyOrderStateLangMgr) GetBatchFromIDLang(idLangs []uint32) (results []*EgSupplyOrderStateLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error
	
	return
}
 
// GetFromName 通过name获取内容  
func (obj *_EgSupplyOrderStateLangMgr) GetFromName(name string) (results []*EgSupplyOrderStateLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error
	
	return
}

// GetBatchFromName 批量唯一主键查找 
func (obj *_EgSupplyOrderStateLangMgr) GetBatchFromName(names []string) (results []*EgSupplyOrderStateLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgSupplyOrderStateLangMgr) FetchByPrimaryKey(idSupplyOrderState uint32 ,idLang uint32 ) (result EgSupplyOrderStateLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supply_order_state = ? AND id_lang = ?", idSupplyOrderState , idLang).Find(&result).Error
	
	return
}
 

 

	

