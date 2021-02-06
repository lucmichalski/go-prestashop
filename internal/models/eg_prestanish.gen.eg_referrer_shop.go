package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgReferrerShopMgr struct {
	*_BaseMgr
}

// EgReferrerShopMgr open func
func EgReferrerShopMgr(db *gorm.DB) *_EgReferrerShopMgr {
	if db == nil {
		panic(fmt.Errorf("EgReferrerShopMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgReferrerShopMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_referrer_shop"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgReferrerShopMgr) GetTableName() string {
	return "eg_referrer_shop"
}

// Get 获取
func (obj *_EgReferrerShopMgr) Get() (result EgReferrerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgReferrerShopMgr) Gets() (results []*EgReferrerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDReferrer id_referrer获取 
func (obj *_EgReferrerShopMgr) WithIDReferrer(idReferrer uint32) Option {
	return optionFunc(func(o *options) { o.query["id_referrer"] = idReferrer })
}

// WithIDShop id_shop获取 
func (obj *_EgReferrerShopMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

// WithCacheVisitors cache_visitors获取 
func (obj *_EgReferrerShopMgr) WithCacheVisitors(cacheVisitors int) Option {
	return optionFunc(func(o *options) { o.query["cache_visitors"] = cacheVisitors })
}

// WithCacheVisits cache_visits获取 
func (obj *_EgReferrerShopMgr) WithCacheVisits(cacheVisits int) Option {
	return optionFunc(func(o *options) { o.query["cache_visits"] = cacheVisits })
}

// WithCachePages cache_pages获取 
func (obj *_EgReferrerShopMgr) WithCachePages(cachePages int) Option {
	return optionFunc(func(o *options) { o.query["cache_pages"] = cachePages })
}

// WithCacheRegistrations cache_registrations获取 
func (obj *_EgReferrerShopMgr) WithCacheRegistrations(cacheRegistrations int) Option {
	return optionFunc(func(o *options) { o.query["cache_registrations"] = cacheRegistrations })
}

// WithCacheOrders cache_orders获取 
func (obj *_EgReferrerShopMgr) WithCacheOrders(cacheOrders int) Option {
	return optionFunc(func(o *options) { o.query["cache_orders"] = cacheOrders })
}

// WithCacheSales cache_sales获取 
func (obj *_EgReferrerShopMgr) WithCacheSales(cacheSales float64) Option {
	return optionFunc(func(o *options) { o.query["cache_sales"] = cacheSales })
}

// WithCacheRegRate cache_reg_rate获取 
func (obj *_EgReferrerShopMgr) WithCacheRegRate(cacheRegRate float64) Option {
	return optionFunc(func(o *options) { o.query["cache_reg_rate"] = cacheRegRate })
}

// WithCacheOrderRate cache_order_rate获取 
func (obj *_EgReferrerShopMgr) WithCacheOrderRate(cacheOrderRate float64) Option {
	return optionFunc(func(o *options) { o.query["cache_order_rate"] = cacheOrderRate })
}


// GetByOption 功能选项模式获取
func (obj *_EgReferrerShopMgr) GetByOption(opts ...Option) (result EgReferrerShop, err error) {
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
func (obj *_EgReferrerShopMgr) GetByOptions(opts ...Option) (results []*EgReferrerShop, err error) {
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


// GetFromIDReferrer 通过id_referrer获取内容  
func (obj *_EgReferrerShopMgr) GetFromIDReferrer(idReferrer uint32) (results []*EgReferrerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_referrer = ?", idReferrer).Find(&results).Error
	
	return
}

// GetBatchFromIDReferrer 批量唯一主键查找 
func (obj *_EgReferrerShopMgr) GetBatchFromIDReferrer(idReferrers []uint32) (results []*EgReferrerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_referrer IN (?)", idReferrers).Find(&results).Error
	
	return
}
 
// GetFromIDShop 通过id_shop获取内容  
func (obj *_EgReferrerShopMgr) GetFromIDShop(idShop uint32) (results []*EgReferrerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error
	
	return
}

// GetBatchFromIDShop 批量唯一主键查找 
func (obj *_EgReferrerShopMgr) GetBatchFromIDShop(idShops []uint32) (results []*EgReferrerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error
	
	return
}
 
// GetFromCacheVisitors 通过cache_visitors获取内容  
func (obj *_EgReferrerShopMgr) GetFromCacheVisitors(cacheVisitors int) (results []*EgReferrerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cache_visitors = ?", cacheVisitors).Find(&results).Error
	
	return
}

// GetBatchFromCacheVisitors 批量唯一主键查找 
func (obj *_EgReferrerShopMgr) GetBatchFromCacheVisitors(cacheVisitorss []int) (results []*EgReferrerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cache_visitors IN (?)", cacheVisitorss).Find(&results).Error
	
	return
}
 
// GetFromCacheVisits 通过cache_visits获取内容  
func (obj *_EgReferrerShopMgr) GetFromCacheVisits(cacheVisits int) (results []*EgReferrerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cache_visits = ?", cacheVisits).Find(&results).Error
	
	return
}

// GetBatchFromCacheVisits 批量唯一主键查找 
func (obj *_EgReferrerShopMgr) GetBatchFromCacheVisits(cacheVisitss []int) (results []*EgReferrerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cache_visits IN (?)", cacheVisitss).Find(&results).Error
	
	return
}
 
// GetFromCachePages 通过cache_pages获取内容  
func (obj *_EgReferrerShopMgr) GetFromCachePages(cachePages int) (results []*EgReferrerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cache_pages = ?", cachePages).Find(&results).Error
	
	return
}

// GetBatchFromCachePages 批量唯一主键查找 
func (obj *_EgReferrerShopMgr) GetBatchFromCachePages(cachePagess []int) (results []*EgReferrerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cache_pages IN (?)", cachePagess).Find(&results).Error
	
	return
}
 
// GetFromCacheRegistrations 通过cache_registrations获取内容  
func (obj *_EgReferrerShopMgr) GetFromCacheRegistrations(cacheRegistrations int) (results []*EgReferrerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cache_registrations = ?", cacheRegistrations).Find(&results).Error
	
	return
}

// GetBatchFromCacheRegistrations 批量唯一主键查找 
func (obj *_EgReferrerShopMgr) GetBatchFromCacheRegistrations(cacheRegistrationss []int) (results []*EgReferrerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cache_registrations IN (?)", cacheRegistrationss).Find(&results).Error
	
	return
}
 
// GetFromCacheOrders 通过cache_orders获取内容  
func (obj *_EgReferrerShopMgr) GetFromCacheOrders(cacheOrders int) (results []*EgReferrerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cache_orders = ?", cacheOrders).Find(&results).Error
	
	return
}

// GetBatchFromCacheOrders 批量唯一主键查找 
func (obj *_EgReferrerShopMgr) GetBatchFromCacheOrders(cacheOrderss []int) (results []*EgReferrerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cache_orders IN (?)", cacheOrderss).Find(&results).Error
	
	return
}
 
// GetFromCacheSales 通过cache_sales获取内容  
func (obj *_EgReferrerShopMgr) GetFromCacheSales(cacheSales float64) (results []*EgReferrerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cache_sales = ?", cacheSales).Find(&results).Error
	
	return
}

// GetBatchFromCacheSales 批量唯一主键查找 
func (obj *_EgReferrerShopMgr) GetBatchFromCacheSales(cacheSaless []float64) (results []*EgReferrerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cache_sales IN (?)", cacheSaless).Find(&results).Error
	
	return
}
 
// GetFromCacheRegRate 通过cache_reg_rate获取内容  
func (obj *_EgReferrerShopMgr) GetFromCacheRegRate(cacheRegRate float64) (results []*EgReferrerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cache_reg_rate = ?", cacheRegRate).Find(&results).Error
	
	return
}

// GetBatchFromCacheRegRate 批量唯一主键查找 
func (obj *_EgReferrerShopMgr) GetBatchFromCacheRegRate(cacheRegRates []float64) (results []*EgReferrerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cache_reg_rate IN (?)", cacheRegRates).Find(&results).Error
	
	return
}
 
// GetFromCacheOrderRate 通过cache_order_rate获取内容  
func (obj *_EgReferrerShopMgr) GetFromCacheOrderRate(cacheOrderRate float64) (results []*EgReferrerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cache_order_rate = ?", cacheOrderRate).Find(&results).Error
	
	return
}

// GetBatchFromCacheOrderRate 批量唯一主键查找 
func (obj *_EgReferrerShopMgr) GetBatchFromCacheOrderRate(cacheOrderRates []float64) (results []*EgReferrerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cache_order_rate IN (?)", cacheOrderRates).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgReferrerShopMgr) FetchByPrimaryKey(idReferrer uint32 ,idShop uint32 ) (result EgReferrerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_referrer = ? AND id_shop = ?", idReferrer , idShop).Find(&result).Error
	
	return
}
 

 

	

