package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgStockMvtReasonLangMgr struct {
	*_BaseMgr
}

// EgStockMvtReasonLangMgr open func
func EgStockMvtReasonLangMgr(db *gorm.DB) *_EgStockMvtReasonLangMgr {
	if db == nil {
		panic(fmt.Errorf("EgStockMvtReasonLangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgStockMvtReasonLangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_stock_mvt_reason_lang"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgStockMvtReasonLangMgr) GetTableName() string {
	return "eg_stock_mvt_reason_lang"
}

// Get 获取
func (obj *_EgStockMvtReasonLangMgr) Get() (result EgStockMvtReasonLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgStockMvtReasonLangMgr) Gets() (results []*EgStockMvtReasonLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDStockMvtReason id_stock_mvt_reason获取 
func (obj *_EgStockMvtReasonLangMgr) WithIDStockMvtReason(idStockMvtReason uint32) Option {
	return optionFunc(func(o *options) { o.query["id_stock_mvt_reason"] = idStockMvtReason })
}

// WithIDLang id_lang获取 
func (obj *_EgStockMvtReasonLangMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

// WithName name获取 
func (obj *_EgStockMvtReasonLangMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}


// GetByOption 功能选项模式获取
func (obj *_EgStockMvtReasonLangMgr) GetByOption(opts ...Option) (result EgStockMvtReasonLang, err error) {
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
func (obj *_EgStockMvtReasonLangMgr) GetByOptions(opts ...Option) (results []*EgStockMvtReasonLang, err error) {
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


// GetFromIDStockMvtReason 通过id_stock_mvt_reason获取内容  
func (obj *_EgStockMvtReasonLangMgr) GetFromIDStockMvtReason(idStockMvtReason uint32) (results []*EgStockMvtReasonLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_stock_mvt_reason = ?", idStockMvtReason).Find(&results).Error
	
	return
}

// GetBatchFromIDStockMvtReason 批量唯一主键查找 
func (obj *_EgStockMvtReasonLangMgr) GetBatchFromIDStockMvtReason(idStockMvtReasons []uint32) (results []*EgStockMvtReasonLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_stock_mvt_reason IN (?)", idStockMvtReasons).Find(&results).Error
	
	return
}
 
// GetFromIDLang 通过id_lang获取内容  
func (obj *_EgStockMvtReasonLangMgr) GetFromIDLang(idLang uint32) (results []*EgStockMvtReasonLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error
	
	return
}

// GetBatchFromIDLang 批量唯一主键查找 
func (obj *_EgStockMvtReasonLangMgr) GetBatchFromIDLang(idLangs []uint32) (results []*EgStockMvtReasonLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error
	
	return
}
 
// GetFromName 通过name获取内容  
func (obj *_EgStockMvtReasonLangMgr) GetFromName(name string) (results []*EgStockMvtReasonLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error
	
	return
}

// GetBatchFromName 批量唯一主键查找 
func (obj *_EgStockMvtReasonLangMgr) GetBatchFromName(names []string) (results []*EgStockMvtReasonLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgStockMvtReasonLangMgr) FetchByPrimaryKey(idStockMvtReason uint32 ,idLang uint32 ) (result EgStockMvtReasonLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_stock_mvt_reason = ? AND id_lang = ?", idStockMvtReason , idLang).Find(&result).Error
	
	return
}
 

 

	

