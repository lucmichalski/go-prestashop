package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _PageViewedMgr struct {
	*_BaseMgr
}

func PageViewedMgr(db *gorm.DB) *_PageViewedMgr {
	if db == nil {
		panic(fmt.Errorf("PageViewedMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_PageViewedMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_page_viewed"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_PageViewedMgr) GetTableName() string {
	return "ps_page_viewed"
}

func (obj *_PageViewedMgr) Get() (result PageViewed, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_PageViewedMgr) Gets() (results []*PageViewed, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_PageViewedMgr) WithIDPage(idPage uint32) Option {
	return optionFunc(func(o *options) { o.query["id_page"] = idPage })
}

func (obj *_PageViewedMgr) WithIDShopGroup(idShopGroup uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop_group"] = idShopGroup })
}

func (obj *_PageViewedMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

func (obj *_PageViewedMgr) WithIDDateRange(idDateRange uint32) Option {
	return optionFunc(func(o *options) { o.query["id_date_range"] = idDateRange })
}

func (obj *_PageViewedMgr) WithCounter(counter uint32) Option {
	return optionFunc(func(o *options) { o.query["counter"] = counter })
}

func (obj *_PageViewedMgr) GetByOption(opts ...Option) (result PageViewed, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_PageViewedMgr) GetByOptions(opts ...Option) (results []*PageViewed, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_PageViewedMgr) GetFromIDPage(idPage uint32) (results []*PageViewed, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_page = ?", idPage).Find(&results).Error

	return
}

func (obj *_PageViewedMgr) GetBatchFromIDPage(idPages []uint32) (results []*PageViewed, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_page IN (?)", idPages).Find(&results).Error

	return
}

func (obj *_PageViewedMgr) GetFromIDShopGroup(idShopGroup uint32) (results []*PageViewed, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_group = ?", idShopGroup).Find(&results).Error

	return
}

func (obj *_PageViewedMgr) GetBatchFromIDShopGroup(idShopGroups []uint32) (results []*PageViewed, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_group IN (?)", idShopGroups).Find(&results).Error

	return
}

func (obj *_PageViewedMgr) GetFromIDShop(idShop uint32) (results []*PageViewed, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

func (obj *_PageViewedMgr) GetBatchFromIDShop(idShops []uint32) (results []*PageViewed, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

func (obj *_PageViewedMgr) GetFromIDDateRange(idDateRange uint32) (results []*PageViewed, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_date_range = ?", idDateRange).Find(&results).Error

	return
}

func (obj *_PageViewedMgr) GetBatchFromIDDateRange(idDateRanges []uint32) (results []*PageViewed, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_date_range IN (?)", idDateRanges).Find(&results).Error

	return
}

func (obj *_PageViewedMgr) GetFromCounter(counter uint32) (results []*PageViewed, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("counter = ?", counter).Find(&results).Error

	return
}

func (obj *_PageViewedMgr) GetBatchFromCounter(counters []uint32) (results []*PageViewed, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("counter IN (?)", counters).Find(&results).Error

	return
}


func (obj *_PageViewedMgr) FetchByPrimaryKey(idPage uint32, idShop uint32, idDateRange uint32) (result PageViewed, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_page = ? AND id_shop = ? AND id_date_range = ?", idPage, idShop, idDateRange).Find(&result).Error

	return
}
