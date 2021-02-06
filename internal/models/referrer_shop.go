package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _ReferrerShopMgr struct {
	*_BaseMgr
}

// ReferrerShopMgr open func
func ReferrerShopMgr(db *gorm.DB) *_ReferrerShopMgr {
	if db == nil {
		panic(fmt.Errorf("ReferrerShopMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ReferrerShopMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_referrer_shop"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ReferrerShopMgr) GetTableName() string {
	return "ps_referrer_shop"
}

// Get 获取
func (obj *_ReferrerShopMgr) Get() (result ReferrerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ReferrerShopMgr) Gets() (results []*ReferrerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDReferrer id_referrer获取
func (obj *_ReferrerShopMgr) WithIDReferrer(idReferrer uint32) Option {
	return optionFunc(func(o *options) { o.query["id_referrer"] = idReferrer })
}

// WithIDShop id_shop获取
func (obj *_ReferrerShopMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

// WithCacheVisitors cache_visitors获取
func (obj *_ReferrerShopMgr) WithCacheVisitors(cacheVisitors int) Option {
	return optionFunc(func(o *options) { o.query["cache_visitors"] = cacheVisitors })
}

// WithCacheVisits cache_visits获取
func (obj *_ReferrerShopMgr) WithCacheVisits(cacheVisits int) Option {
	return optionFunc(func(o *options) { o.query["cache_visits"] = cacheVisits })
}

// WithCachePages cache_pages获取
func (obj *_ReferrerShopMgr) WithCachePages(cachePages int) Option {
	return optionFunc(func(o *options) { o.query["cache_pages"] = cachePages })
}

// WithCacheRegistrations cache_registrations获取
func (obj *_ReferrerShopMgr) WithCacheRegistrations(cacheRegistrations int) Option {
	return optionFunc(func(o *options) { o.query["cache_registrations"] = cacheRegistrations })
}

// WithCacheOrders cache_orders获取
func (obj *_ReferrerShopMgr) WithCacheOrders(cacheOrders int) Option {
	return optionFunc(func(o *options) { o.query["cache_orders"] = cacheOrders })
}

// WithCacheSales cache_sales获取
func (obj *_ReferrerShopMgr) WithCacheSales(cacheSales float64) Option {
	return optionFunc(func(o *options) { o.query["cache_sales"] = cacheSales })
}

// WithCacheRegRate cache_rps_rate获取
func (obj *_ReferrerShopMgr) WithCacheRegRate(cacheRegRate float64) Option {
	return optionFunc(func(o *options) { o.query["cache_rps_rate"] = cacheRegRate })
}

// WithCacheOrderRate cache_order_rate获取
func (obj *_ReferrerShopMgr) WithCacheOrderRate(cacheOrderRate float64) Option {
	return optionFunc(func(o *options) { o.query["cache_order_rate"] = cacheOrderRate })
}

// GetByOption 功能选项模式获取
func (obj *_ReferrerShopMgr) GetByOption(opts ...Option) (result ReferrerShop, err error) {
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
func (obj *_ReferrerShopMgr) GetByOptions(opts ...Option) (results []*ReferrerShop, err error) {
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
func (obj *_ReferrerShopMgr) GetFromIDReferrer(idReferrer uint32) (results []*ReferrerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_referrer = ?", idReferrer).Find(&results).Error

	return
}

// GetBatchFromIDReferrer 批量唯一主键查找
func (obj *_ReferrerShopMgr) GetBatchFromIDReferrer(idReferrers []uint32) (results []*ReferrerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_referrer IN (?)", idReferrers).Find(&results).Error

	return
}

// GetFromIDShop 通过id_shop获取内容
func (obj *_ReferrerShopMgr) GetFromIDShop(idShop uint32) (results []*ReferrerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

// GetBatchFromIDShop 批量唯一主键查找
func (obj *_ReferrerShopMgr) GetBatchFromIDShop(idShops []uint32) (results []*ReferrerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

// GetFromCacheVisitors 通过cache_visitors获取内容
func (obj *_ReferrerShopMgr) GetFromCacheVisitors(cacheVisitors int) (results []*ReferrerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cache_visitors = ?", cacheVisitors).Find(&results).Error

	return
}

// GetBatchFromCacheVisitors 批量唯一主键查找
func (obj *_ReferrerShopMgr) GetBatchFromCacheVisitors(cacheVisitorss []int) (results []*ReferrerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cache_visitors IN (?)", cacheVisitorss).Find(&results).Error

	return
}

// GetFromCacheVisits 通过cache_visits获取内容
func (obj *_ReferrerShopMgr) GetFromCacheVisits(cacheVisits int) (results []*ReferrerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cache_visits = ?", cacheVisits).Find(&results).Error

	return
}

// GetBatchFromCacheVisits 批量唯一主键查找
func (obj *_ReferrerShopMgr) GetBatchFromCacheVisits(cacheVisitss []int) (results []*ReferrerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cache_visits IN (?)", cacheVisitss).Find(&results).Error

	return
}

// GetFromCachePages 通过cache_pages获取内容
func (obj *_ReferrerShopMgr) GetFromCachePages(cachePages int) (results []*ReferrerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cache_pages = ?", cachePages).Find(&results).Error

	return
}

// GetBatchFromCachePages 批量唯一主键查找
func (obj *_ReferrerShopMgr) GetBatchFromCachePages(cachePagess []int) (results []*ReferrerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cache_pages IN (?)", cachePagess).Find(&results).Error

	return
}

// GetFromCacheRegistrations 通过cache_registrations获取内容
func (obj *_ReferrerShopMgr) GetFromCacheRegistrations(cacheRegistrations int) (results []*ReferrerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cache_registrations = ?", cacheRegistrations).Find(&results).Error

	return
}

// GetBatchFromCacheRegistrations 批量唯一主键查找
func (obj *_ReferrerShopMgr) GetBatchFromCacheRegistrations(cacheRegistrationss []int) (results []*ReferrerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cache_registrations IN (?)", cacheRegistrationss).Find(&results).Error

	return
}

// GetFromCacheOrders 通过cache_orders获取内容
func (obj *_ReferrerShopMgr) GetFromCacheOrders(cacheOrders int) (results []*ReferrerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cache_orders = ?", cacheOrders).Find(&results).Error

	return
}

// GetBatchFromCacheOrders 批量唯一主键查找
func (obj *_ReferrerShopMgr) GetBatchFromCacheOrders(cacheOrderss []int) (results []*ReferrerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cache_orders IN (?)", cacheOrderss).Find(&results).Error

	return
}

// GetFromCacheSales 通过cache_sales获取内容
func (obj *_ReferrerShopMgr) GetFromCacheSales(cacheSales float64) (results []*ReferrerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cache_sales = ?", cacheSales).Find(&results).Error

	return
}

// GetBatchFromCacheSales 批量唯一主键查找
func (obj *_ReferrerShopMgr) GetBatchFromCacheSales(cacheSaless []float64) (results []*ReferrerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cache_sales IN (?)", cacheSaless).Find(&results).Error

	return
}

// GetFromCacheRegRate 通过cache_rps_rate获取内容
func (obj *_ReferrerShopMgr) GetFromCacheRegRate(cacheRegRate float64) (results []*ReferrerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cache_rps_rate = ?", cacheRegRate).Find(&results).Error

	return
}

// GetBatchFromCacheRegRate 批量唯一主键查找
func (obj *_ReferrerShopMgr) GetBatchFromCacheRegRate(cacheRegRates []float64) (results []*ReferrerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cache_rps_rate IN (?)", cacheRegRates).Find(&results).Error

	return
}

// GetFromCacheOrderRate 通过cache_order_rate获取内容
func (obj *_ReferrerShopMgr) GetFromCacheOrderRate(cacheOrderRate float64) (results []*ReferrerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cache_order_rate = ?", cacheOrderRate).Find(&results).Error

	return
}

// GetBatchFromCacheOrderRate 批量唯一主键查找
func (obj *_ReferrerShopMgr) GetBatchFromCacheOrderRate(cacheOrderRates []float64) (results []*ReferrerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cache_order_rate IN (?)", cacheOrderRates).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_ReferrerShopMgr) FetchByPrimaryKey(idReferrer uint32, idShop uint32) (result ReferrerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_referrer = ? AND id_shop = ?", idReferrer, idShop).Find(&result).Error

	return
}
