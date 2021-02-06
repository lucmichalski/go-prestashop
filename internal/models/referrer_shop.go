package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _ReferrerShopMgr struct {
	*_BaseMgr
}

func ReferrerShopMgr(db *gorm.DB) *_ReferrerShopMgr {
	if db == nil {
		panic(fmt.Errorf("ReferrerShopMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ReferrerShopMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_referrer_shop"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_ReferrerShopMgr) GetTableName() string {
	return "ps_referrer_shop"
}

func (obj *_ReferrerShopMgr) Get() (result ReferrerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_ReferrerShopMgr) Gets() (results []*ReferrerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_ReferrerShopMgr) WithIDReferrer(idReferrer uint32) Option {
	return optionFunc(func(o *options) { o.query["id_referrer"] = idReferrer })
}

func (obj *_ReferrerShopMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

func (obj *_ReferrerShopMgr) WithCacheVisitors(cacheVisitors int) Option {
	return optionFunc(func(o *options) { o.query["cache_visitors"] = cacheVisitors })
}

func (obj *_ReferrerShopMgr) WithCacheVisits(cacheVisits int) Option {
	return optionFunc(func(o *options) { o.query["cache_visits"] = cacheVisits })
}

func (obj *_ReferrerShopMgr) WithCachePages(cachePages int) Option {
	return optionFunc(func(o *options) { o.query["cache_pages"] = cachePages })
}

func (obj *_ReferrerShopMgr) WithCacheRegistrations(cacheRegistrations int) Option {
	return optionFunc(func(o *options) { o.query["cache_registrations"] = cacheRegistrations })
}

func (obj *_ReferrerShopMgr) WithCacheOrders(cacheOrders int) Option {
	return optionFunc(func(o *options) { o.query["cache_orders"] = cacheOrders })
}

func (obj *_ReferrerShopMgr) WithCacheSales(cacheSales float64) Option {
	return optionFunc(func(o *options) { o.query["cache_sales"] = cacheSales })
}

func (obj *_ReferrerShopMgr) WithCacheRegRate(cacheRegRate float64) Option {
	return optionFunc(func(o *options) { o.query["cache_rps_rate"] = cacheRegRate })
}

func (obj *_ReferrerShopMgr) WithCacheOrderRate(cacheOrderRate float64) Option {
	return optionFunc(func(o *options) { o.query["cache_order_rate"] = cacheOrderRate })
}

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


func (obj *_ReferrerShopMgr) GetFromIDReferrer(idReferrer uint32) (results []*ReferrerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_referrer = ?", idReferrer).Find(&results).Error

	return
}

func (obj *_ReferrerShopMgr) GetBatchFromIDReferrer(idReferrers []uint32) (results []*ReferrerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_referrer IN (?)", idReferrers).Find(&results).Error

	return
}

func (obj *_ReferrerShopMgr) GetFromIDShop(idShop uint32) (results []*ReferrerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

func (obj *_ReferrerShopMgr) GetBatchFromIDShop(idShops []uint32) (results []*ReferrerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

func (obj *_ReferrerShopMgr) GetFromCacheVisitors(cacheVisitors int) (results []*ReferrerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cache_visitors = ?", cacheVisitors).Find(&results).Error

	return
}

func (obj *_ReferrerShopMgr) GetBatchFromCacheVisitors(cacheVisitorss []int) (results []*ReferrerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cache_visitors IN (?)", cacheVisitorss).Find(&results).Error

	return
}

func (obj *_ReferrerShopMgr) GetFromCacheVisits(cacheVisits int) (results []*ReferrerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cache_visits = ?", cacheVisits).Find(&results).Error

	return
}

func (obj *_ReferrerShopMgr) GetBatchFromCacheVisits(cacheVisitss []int) (results []*ReferrerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cache_visits IN (?)", cacheVisitss).Find(&results).Error

	return
}

func (obj *_ReferrerShopMgr) GetFromCachePages(cachePages int) (results []*ReferrerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cache_pages = ?", cachePages).Find(&results).Error

	return
}

func (obj *_ReferrerShopMgr) GetBatchFromCachePages(cachePagess []int) (results []*ReferrerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cache_pages IN (?)", cachePagess).Find(&results).Error

	return
}

func (obj *_ReferrerShopMgr) GetFromCacheRegistrations(cacheRegistrations int) (results []*ReferrerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cache_registrations = ?", cacheRegistrations).Find(&results).Error

	return
}

func (obj *_ReferrerShopMgr) GetBatchFromCacheRegistrations(cacheRegistrationss []int) (results []*ReferrerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cache_registrations IN (?)", cacheRegistrationss).Find(&results).Error

	return
}

func (obj *_ReferrerShopMgr) GetFromCacheOrders(cacheOrders int) (results []*ReferrerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cache_orders = ?", cacheOrders).Find(&results).Error

	return
}

func (obj *_ReferrerShopMgr) GetBatchFromCacheOrders(cacheOrderss []int) (results []*ReferrerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cache_orders IN (?)", cacheOrderss).Find(&results).Error

	return
}

func (obj *_ReferrerShopMgr) GetFromCacheSales(cacheSales float64) (results []*ReferrerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cache_sales = ?", cacheSales).Find(&results).Error

	return
}

func (obj *_ReferrerShopMgr) GetBatchFromCacheSales(cacheSaless []float64) (results []*ReferrerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cache_sales IN (?)", cacheSaless).Find(&results).Error

	return
}

func (obj *_ReferrerShopMgr) GetFromCacheRegRate(cacheRegRate float64) (results []*ReferrerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cache_rps_rate = ?", cacheRegRate).Find(&results).Error

	return
}

func (obj *_ReferrerShopMgr) GetBatchFromCacheRegRate(cacheRegRates []float64) (results []*ReferrerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cache_rps_rate IN (?)", cacheRegRates).Find(&results).Error

	return
}

func (obj *_ReferrerShopMgr) GetFromCacheOrderRate(cacheOrderRate float64) (results []*ReferrerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cache_order_rate = ?", cacheOrderRate).Find(&results).Error

	return
}

func (obj *_ReferrerShopMgr) GetBatchFromCacheOrderRate(cacheOrderRates []float64) (results []*ReferrerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cache_order_rate IN (?)", cacheOrderRates).Find(&results).Error

	return
}


func (obj *_ReferrerShopMgr) FetchByPrimaryKey(idReferrer uint32, idShop uint32) (result ReferrerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_referrer = ? AND id_shop = ?", idReferrer, idShop).Find(&result).Error

	return
}
