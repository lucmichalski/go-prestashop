package	model	
import (	
"context"	
"gorm.io/gorm"	
"fmt"	
)	

type _EgCurrencyMgr struct {
	*_BaseMgr
}

// EgCurrencyMgr open func
func EgCurrencyMgr(db *gorm.DB) *_EgCurrencyMgr {
	if db == nil {
		panic(fmt.Errorf("EgCurrencyMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgCurrencyMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_currency"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgCurrencyMgr) GetTableName() string {
	return "eg_currency"
}

// Get 获取
func (obj *_EgCurrencyMgr) Get() (result EgCurrency, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgCurrencyMgr) Gets() (results []*EgCurrency, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDCurrency id_currency获取 
func (obj *_EgCurrencyMgr) WithIDCurrency(idCurrency uint32) Option {
	return optionFunc(func(o *options) { o.query["id_currency"] = idCurrency })
}

// WithName name获取 
func (obj *_EgCurrencyMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithIsoCode iso_code获取 
func (obj *_EgCurrencyMgr) WithIsoCode(isoCode string) Option {
	return optionFunc(func(o *options) { o.query["iso_code"] = isoCode })
}

// WithNumericIsoCode numeric_iso_code获取 
func (obj *_EgCurrencyMgr) WithNumericIsoCode(numericIsoCode string) Option {
	return optionFunc(func(o *options) { o.query["numeric_iso_code"] = numericIsoCode })
}

// WithPrecision precision获取 
func (obj *_EgCurrencyMgr) WithPrecision(precision int) Option {
	return optionFunc(func(o *options) { o.query["precision"] = precision })
}

// WithConversionRate conversion_rate获取 
func (obj *_EgCurrencyMgr) WithConversionRate(conversionRate float64) Option {
	return optionFunc(func(o *options) { o.query["conversion_rate"] = conversionRate })
}

// WithDeleted deleted获取 
func (obj *_EgCurrencyMgr) WithDeleted(deleted bool) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// WithActive active获取 
func (obj *_EgCurrencyMgr) WithActive(active bool) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}

// WithUnofficial unofficial获取 
func (obj *_EgCurrencyMgr) WithUnofficial(unofficial bool) Option {
	return optionFunc(func(o *options) { o.query["unofficial"] = unofficial })
}

// WithModified modified获取 
func (obj *_EgCurrencyMgr) WithModified(modified bool) Option {
	return optionFunc(func(o *options) { o.query["modified"] = modified })
}


// GetByOption 功能选项模式获取
func (obj *_EgCurrencyMgr) GetByOption(opts ...Option) (result EgCurrency, err error) {
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
func (obj *_EgCurrencyMgr) GetByOptions(opts ...Option) (results []*EgCurrency, err error) {
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
func (obj *_EgCurrencyMgr)  GetFromIDCurrency(idCurrency uint32) (result EgCurrency, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_currency = ?", idCurrency).Find(&result).Error
	
	return
}

// GetBatchFromIDCurrency 批量唯一主键查找 
func (obj *_EgCurrencyMgr) GetBatchFromIDCurrency(idCurrencys []uint32) (results []*EgCurrency, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_currency IN (?)", idCurrencys).Find(&results).Error
	
	return
}
 
// GetFromName 通过name获取内容  
func (obj *_EgCurrencyMgr) GetFromName(name string) (results []*EgCurrency, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error
	
	return
}

// GetBatchFromName 批量唯一主键查找 
func (obj *_EgCurrencyMgr) GetBatchFromName(names []string) (results []*EgCurrency, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error
	
	return
}
 
// GetFromIsoCode 通过iso_code获取内容  
func (obj *_EgCurrencyMgr) GetFromIsoCode(isoCode string) (results []*EgCurrency, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("iso_code = ?", isoCode).Find(&results).Error
	
	return
}

// GetBatchFromIsoCode 批量唯一主键查找 
func (obj *_EgCurrencyMgr) GetBatchFromIsoCode(isoCodes []string) (results []*EgCurrency, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("iso_code IN (?)", isoCodes).Find(&results).Error
	
	return
}
 
// GetFromNumericIsoCode 通过numeric_iso_code获取内容  
func (obj *_EgCurrencyMgr) GetFromNumericIsoCode(numericIsoCode string) (results []*EgCurrency, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("numeric_iso_code = ?", numericIsoCode).Find(&results).Error
	
	return
}

// GetBatchFromNumericIsoCode 批量唯一主键查找 
func (obj *_EgCurrencyMgr) GetBatchFromNumericIsoCode(numericIsoCodes []string) (results []*EgCurrency, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("numeric_iso_code IN (?)", numericIsoCodes).Find(&results).Error
	
	return
}
 
// GetFromPrecision 通过precision获取内容  
func (obj *_EgCurrencyMgr) GetFromPrecision(precision int) (results []*EgCurrency, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("precision = ?", precision).Find(&results).Error
	
	return
}

// GetBatchFromPrecision 批量唯一主键查找 
func (obj *_EgCurrencyMgr) GetBatchFromPrecision(precisions []int) (results []*EgCurrency, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("precision IN (?)", precisions).Find(&results).Error
	
	return
}
 
// GetFromConversionRate 通过conversion_rate获取内容  
func (obj *_EgCurrencyMgr) GetFromConversionRate(conversionRate float64) (results []*EgCurrency, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("conversion_rate = ?", conversionRate).Find(&results).Error
	
	return
}

// GetBatchFromConversionRate 批量唯一主键查找 
func (obj *_EgCurrencyMgr) GetBatchFromConversionRate(conversionRates []float64) (results []*EgCurrency, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("conversion_rate IN (?)", conversionRates).Find(&results).Error
	
	return
}
 
// GetFromDeleted 通过deleted获取内容  
func (obj *_EgCurrencyMgr) GetFromDeleted(deleted bool) (results []*EgCurrency, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("deleted = ?", deleted).Find(&results).Error
	
	return
}

// GetBatchFromDeleted 批量唯一主键查找 
func (obj *_EgCurrencyMgr) GetBatchFromDeleted(deleteds []bool) (results []*EgCurrency, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("deleted IN (?)", deleteds).Find(&results).Error
	
	return
}
 
// GetFromActive 通过active获取内容  
func (obj *_EgCurrencyMgr) GetFromActive(active bool) (results []*EgCurrency, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active = ?", active).Find(&results).Error
	
	return
}

// GetBatchFromActive 批量唯一主键查找 
func (obj *_EgCurrencyMgr) GetBatchFromActive(actives []bool) (results []*EgCurrency, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active IN (?)", actives).Find(&results).Error
	
	return
}
 
// GetFromUnofficial 通过unofficial获取内容  
func (obj *_EgCurrencyMgr) GetFromUnofficial(unofficial bool) (results []*EgCurrency, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("unofficial = ?", unofficial).Find(&results).Error
	
	return
}

// GetBatchFromUnofficial 批量唯一主键查找 
func (obj *_EgCurrencyMgr) GetBatchFromUnofficial(unofficials []bool) (results []*EgCurrency, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("unofficial IN (?)", unofficials).Find(&results).Error
	
	return
}
 
// GetFromModified 通过modified获取内容  
func (obj *_EgCurrencyMgr) GetFromModified(modified bool) (results []*EgCurrency, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("modified = ?", modified).Find(&results).Error
	
	return
}

// GetBatchFromModified 批量唯一主键查找 
func (obj *_EgCurrencyMgr) GetBatchFromModified(modifieds []bool) (results []*EgCurrency, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("modified IN (?)", modifieds).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgCurrencyMgr) FetchByPrimaryKey(idCurrency uint32 ) (result EgCurrency, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_currency = ?", idCurrency).Find(&result).Error
	
	return
}
 

 
 // FetchIndexByCurrencyIsoCode  获取多个内容
 func (obj *_EgCurrencyMgr) FetchIndexByCurrencyIsoCode(isoCode string ) (results []*EgCurrency, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("iso_code = ?", isoCode).Find(&results).Error
	
	return
}
 

	

