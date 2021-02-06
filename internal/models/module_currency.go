package	model	
import (	
"context"	
"gorm.io/gorm"	
"fmt"	
)	

type _EgModuleCurrencyMgr struct {
	*_BaseMgr
}

// EgModuleCurrencyMgr open func
func EgModuleCurrencyMgr(db *gorm.DB) *_EgModuleCurrencyMgr {
	if db == nil {
		panic(fmt.Errorf("EgModuleCurrencyMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgModuleCurrencyMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_module_currency"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgModuleCurrencyMgr) GetTableName() string {
	return "eg_module_currency"
}

// Get 获取
func (obj *_EgModuleCurrencyMgr) Get() (result EgModuleCurrency, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgModuleCurrencyMgr) Gets() (results []*EgModuleCurrency, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDModule id_module获取 
func (obj *_EgModuleCurrencyMgr) WithIDModule(idModule uint32) Option {
	return optionFunc(func(o *options) { o.query["id_module"] = idModule })
}

// WithIDShop id_shop获取 
func (obj *_EgModuleCurrencyMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

// WithIDCurrency id_currency获取 
func (obj *_EgModuleCurrencyMgr) WithIDCurrency(idCurrency int) Option {
	return optionFunc(func(o *options) { o.query["id_currency"] = idCurrency })
}


// GetByOption 功能选项模式获取
func (obj *_EgModuleCurrencyMgr) GetByOption(opts ...Option) (result EgModuleCurrency, err error) {
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
func (obj *_EgModuleCurrencyMgr) GetByOptions(opts ...Option) (results []*EgModuleCurrency, err error) {
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


// GetFromIDModule 通过id_module获取内容  
func (obj *_EgModuleCurrencyMgr) GetFromIDModule(idModule uint32) (results []*EgModuleCurrency, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_module = ?", idModule).Find(&results).Error
	
	return
}

// GetBatchFromIDModule 批量唯一主键查找 
func (obj *_EgModuleCurrencyMgr) GetBatchFromIDModule(idModules []uint32) (results []*EgModuleCurrency, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_module IN (?)", idModules).Find(&results).Error
	
	return
}
 
// GetFromIDShop 通过id_shop获取内容  
func (obj *_EgModuleCurrencyMgr) GetFromIDShop(idShop uint32) (results []*EgModuleCurrency, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error
	
	return
}

// GetBatchFromIDShop 批量唯一主键查找 
func (obj *_EgModuleCurrencyMgr) GetBatchFromIDShop(idShops []uint32) (results []*EgModuleCurrency, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error
	
	return
}
 
// GetFromIDCurrency 通过id_currency获取内容  
func (obj *_EgModuleCurrencyMgr) GetFromIDCurrency(idCurrency int) (results []*EgModuleCurrency, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_currency = ?", idCurrency).Find(&results).Error
	
	return
}

// GetBatchFromIDCurrency 批量唯一主键查找 
func (obj *_EgModuleCurrencyMgr) GetBatchFromIDCurrency(idCurrencys []int) (results []*EgModuleCurrency, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_currency IN (?)", idCurrencys).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgModuleCurrencyMgr) FetchByPrimaryKey(idModule uint32 ,idShop uint32 ,idCurrency int ) (result EgModuleCurrency, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_module = ? AND id_shop = ? AND id_currency = ?", idModule , idShop , idCurrency).Find(&result).Error
	
	return
}
 

 
 // FetchIndexByIDModule  获取多个内容
 func (obj *_EgModuleCurrencyMgr) FetchIndexByIDModule(idModule uint32 ) (results []*EgModuleCurrency, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_module = ?", idModule).Find(&results).Error
	
	return
}
 

	

