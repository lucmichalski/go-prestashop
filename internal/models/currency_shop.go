package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgCurrencyShopMgr struct {
	*_BaseMgr
}

// EgCurrencyShopMgr open func
func EgCurrencyShopMgr(db *gorm.DB) *_EgCurrencyShopMgr {
	if db == nil {
		panic(fmt.Errorf("EgCurrencyShopMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgCurrencyShopMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_currency_shop"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgCurrencyShopMgr) GetTableName() string {
	return "eg_currency_shop"
}

// Get 获取
func (obj *_EgCurrencyShopMgr) Get() (result EgCurrencyShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgCurrencyShopMgr) Gets() (results []*EgCurrencyShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDCurrency id_currency获取 
func (obj *_EgCurrencyShopMgr) WithIDCurrency(idCurrency uint32) Option {
	return optionFunc(func(o *options) { o.query["id_currency"] = idCurrency })
}

// WithIDShop id_shop获取 
func (obj *_EgCurrencyShopMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

// WithConversionRate conversion_rate获取 
func (obj *_EgCurrencyShopMgr) WithConversionRate(conversionRate float64) Option {
	return optionFunc(func(o *options) { o.query["conversion_rate"] = conversionRate })
}


// GetByOption 功能选项模式获取
func (obj *_EgCurrencyShopMgr) GetByOption(opts ...Option) (result EgCurrencyShop, err error) {
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
func (obj *_EgCurrencyShopMgr) GetByOptions(opts ...Option) (results []*EgCurrencyShop, err error) {
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
func (obj *_EgCurrencyShopMgr) GetFromIDCurrency(idCurrency uint32) (results []*EgCurrencyShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_currency = ?", idCurrency).Find(&results).Error
	
	return
}

// GetBatchFromIDCurrency 批量唯一主键查找 
func (obj *_EgCurrencyShopMgr) GetBatchFromIDCurrency(idCurrencys []uint32) (results []*EgCurrencyShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_currency IN (?)", idCurrencys).Find(&results).Error
	
	return
}
 
// GetFromIDShop 通过id_shop获取内容  
func (obj *_EgCurrencyShopMgr) GetFromIDShop(idShop uint32) (results []*EgCurrencyShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error
	
	return
}

// GetBatchFromIDShop 批量唯一主键查找 
func (obj *_EgCurrencyShopMgr) GetBatchFromIDShop(idShops []uint32) (results []*EgCurrencyShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error
	
	return
}
 
// GetFromConversionRate 通过conversion_rate获取内容  
func (obj *_EgCurrencyShopMgr) GetFromConversionRate(conversionRate float64) (results []*EgCurrencyShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("conversion_rate = ?", conversionRate).Find(&results).Error
	
	return
}

// GetBatchFromConversionRate 批量唯一主键查找 
func (obj *_EgCurrencyShopMgr) GetBatchFromConversionRate(conversionRates []float64) (results []*EgCurrencyShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("conversion_rate IN (?)", conversionRates).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgCurrencyShopMgr) FetchByPrimaryKey(idCurrency uint32 ,idShop uint32 ) (result EgCurrencyShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_currency = ? AND id_shop = ?", idCurrency , idShop).Find(&result).Error
	
	return
}
 

 
 // FetchIndexByIDShop  获取多个内容
 func (obj *_EgCurrencyShopMgr) FetchIndexByIDShop(idShop uint32 ) (results []*EgCurrencyShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error
	
	return
}
 

	

