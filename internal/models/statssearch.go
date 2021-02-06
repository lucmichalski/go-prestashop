package models

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _StatssearchMgr struct {
	*_BaseMgr
}

func StatssearchMgr(db *gorm.DB) *_StatssearchMgr {
	if db == nil {
		panic(fmt.Errorf("StatssearchMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_StatssearchMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_statssearch"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_StatssearchMgr) GetTableName() string {
	return "ps_statssearch"
}

func (obj *_StatssearchMgr) Get() (result Statssearch, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_StatssearchMgr) Gets() (results []*Statssearch, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_StatssearchMgr) WithIDStatssearch(idStatssearch uint32) Option {
	return optionFunc(func(o *options) { o.query["id_statssearch"] = idStatssearch })
}

func (obj *_StatssearchMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

func (obj *_StatssearchMgr) WithIDShopGroup(idShopGroup uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop_group"] = idShopGroup })
}

func (obj *_StatssearchMgr) WithKeywords(keywords string) Option {
	return optionFunc(func(o *options) { o.query["keywords"] = keywords })
}

func (obj *_StatssearchMgr) WithResults(results int) Option {
	return optionFunc(func(o *options) { o.query["results"] = results })
}

func (obj *_StatssearchMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

func (obj *_StatssearchMgr) GetByOption(opts ...Option) (result Statssearch, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_StatssearchMgr) GetByOptions(opts ...Option) (results []*Statssearch, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_StatssearchMgr) GetFromIDStatssearch(idStatssearch uint32) (result Statssearch, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_statssearch = ?", idStatssearch).Find(&result).Error

	return
}

func (obj *_StatssearchMgr) GetBatchFromIDStatssearch(idStatssearchs []uint32) (results []*Statssearch, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_statssearch IN (?)", idStatssearchs).Find(&results).Error

	return
}

func (obj *_StatssearchMgr) GetFromIDShop(idShop uint32) (results []*Statssearch, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

func (obj *_StatssearchMgr) GetBatchFromIDShop(idShops []uint32) (results []*Statssearch, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

func (obj *_StatssearchMgr) GetFromIDShopGroup(idShopGroup uint32) (results []*Statssearch, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_group = ?", idShopGroup).Find(&results).Error

	return
}

func (obj *_StatssearchMgr) GetBatchFromIDShopGroup(idShopGroups []uint32) (results []*Statssearch, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_group IN (?)", idShopGroups).Find(&results).Error

	return
}

func (obj *_StatssearchMgr) GetFromKeywords(keywords string) (results []*Statssearch, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("keywords = ?", keywords).Find(&results).Error

	return
}

func (obj *_StatssearchMgr) GetBatchFromKeywords(keywordss []string) (results []*Statssearch, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("keywords IN (?)", keywordss).Find(&results).Error

	return
}

func (obj *_StatssearchMgr) GetFromResults(results int) (results []*Statssearch, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("results = ?", results).Find(&results).Error

	return
}

func (obj *_StatssearchMgr) GetBatchFromResults(resultss []int) (results []*Statssearch, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("results IN (?)", resultss).Find(&results).Error

	return
}

func (obj *_StatssearchMgr) GetFromDateAdd(dateAdd time.Time) (results []*Statssearch, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error

	return
}

func (obj *_StatssearchMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*Statssearch, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error

	return
}


func (obj *_StatssearchMgr) FetchByPrimaryKey(idStatssearch uint32) (result Statssearch, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_statssearch = ?", idStatssearch).Find(&result).Error

	return
}
