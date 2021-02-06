package	model	
import (	
"context"	
"gorm.io/gorm"	
"fmt"	
)	

type _EgCurrencyLangMgr struct {
	*_BaseMgr
}

// EgCurrencyLangMgr open func
func EgCurrencyLangMgr(db *gorm.DB) *_EgCurrencyLangMgr {
	if db == nil {
		panic(fmt.Errorf("EgCurrencyLangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgCurrencyLangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_currency_lang"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgCurrencyLangMgr) GetTableName() string {
	return "eg_currency_lang"
}

// Get 获取
func (obj *_EgCurrencyLangMgr) Get() (result EgCurrencyLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgCurrencyLangMgr) Gets() (results []*EgCurrencyLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDCurrency id_currency获取 
func (obj *_EgCurrencyLangMgr) WithIDCurrency(idCurrency uint32) Option {
	return optionFunc(func(o *options) { o.query["id_currency"] = idCurrency })
}

// WithIDLang id_lang获取 
func (obj *_EgCurrencyLangMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

// WithName name获取 
func (obj *_EgCurrencyLangMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithSymbol symbol获取 
func (obj *_EgCurrencyLangMgr) WithSymbol(symbol string) Option {
	return optionFunc(func(o *options) { o.query["symbol"] = symbol })
}

// WithPattern pattern获取 
func (obj *_EgCurrencyLangMgr) WithPattern(pattern string) Option {
	return optionFunc(func(o *options) { o.query["pattern"] = pattern })
}


// GetByOption 功能选项模式获取
func (obj *_EgCurrencyLangMgr) GetByOption(opts ...Option) (result EgCurrencyLang, err error) {
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
func (obj *_EgCurrencyLangMgr) GetByOptions(opts ...Option) (results []*EgCurrencyLang, err error) {
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


// GetFromIDCurrency 通过id_currency获取内容  
func (obj *_EgCurrencyLangMgr) GetFromIDCurrency(idCurrency uint32) (results []*EgCurrencyLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_currency = ?", idCurrency).Find(&results).Error
	
	return
}

// GetBatchFromIDCurrency 批量唯一主键查找 
func (obj *_EgCurrencyLangMgr) GetBatchFromIDCurrency(idCurrencys []uint32) (results []*EgCurrencyLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_currency IN (?)", idCurrencys).Find(&results).Error
	
	return
}
 
// GetFromIDLang 通过id_lang获取内容  
func (obj *_EgCurrencyLangMgr) GetFromIDLang(idLang uint32) (results []*EgCurrencyLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error
	
	return
}

// GetBatchFromIDLang 批量唯一主键查找 
func (obj *_EgCurrencyLangMgr) GetBatchFromIDLang(idLangs []uint32) (results []*EgCurrencyLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error
	
	return
}
 
// GetFromName 通过name获取内容  
func (obj *_EgCurrencyLangMgr) GetFromName(name string) (results []*EgCurrencyLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error
	
	return
}

// GetBatchFromName 批量唯一主键查找 
func (obj *_EgCurrencyLangMgr) GetBatchFromName(names []string) (results []*EgCurrencyLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error
	
	return
}
 
// GetFromSymbol 通过symbol获取内容  
func (obj *_EgCurrencyLangMgr) GetFromSymbol(symbol string) (results []*EgCurrencyLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("symbol = ?", symbol).Find(&results).Error
	
	return
}

// GetBatchFromSymbol 批量唯一主键查找 
func (obj *_EgCurrencyLangMgr) GetBatchFromSymbol(symbols []string) (results []*EgCurrencyLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("symbol IN (?)", symbols).Find(&results).Error
	
	return
}
 
// GetFromPattern 通过pattern获取内容  
func (obj *_EgCurrencyLangMgr) GetFromPattern(pattern string) (results []*EgCurrencyLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("pattern = ?", pattern).Find(&results).Error
	
	return
}

// GetBatchFromPattern 批量唯一主键查找 
func (obj *_EgCurrencyLangMgr) GetBatchFromPattern(patterns []string) (results []*EgCurrencyLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("pattern IN (?)", patterns).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgCurrencyLangMgr) FetchByPrimaryKey(idCurrency uint32 ,idLang uint32 ) (result EgCurrencyLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_currency = ? AND id_lang = ?", idCurrency , idLang).Find(&result).Error
	
	return
}
 

 

	

