package models

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _PagenotfoundMgr struct {
	*_BaseMgr
}

func PagenotfoundMgr(db *gorm.DB) *_PagenotfoundMgr {
	if db == nil {
		panic(fmt.Errorf("PagenotfoundMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_PagenotfoundMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_pagenotfound"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_PagenotfoundMgr) GetTableName() string {
	return "ps_pagenotfound"
}

func (obj *_PagenotfoundMgr) Get() (result Pagenotfound, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_PagenotfoundMgr) Gets() (results []*Pagenotfound, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

func (obj *_PagenotfoundMgr) WithIDPagenotfound(idPagenotfound uint32) Option {
	return optionFunc(func(o *options) { o.query["id_pagenotfound"] = idPagenotfound })
}

func (obj *_PagenotfoundMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

func (obj *_PagenotfoundMgr) WithIDShopGroup(idShopGroup uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop_group"] = idShopGroup })
}

func (obj *_PagenotfoundMgr) WithRequestURI(requestURI string) Option {
	return optionFunc(func(o *options) { o.query["request_uri"] = requestURI })
}

func (obj *_PagenotfoundMgr) WithHTTPReferer(httpReferer string) Option {
	return optionFunc(func(o *options) { o.query["http_referer"] = httpReferer })
}

func (obj *_PagenotfoundMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

func (obj *_PagenotfoundMgr) GetByOption(opts ...Option) (result Pagenotfound, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_PagenotfoundMgr) GetByOptions(opts ...Option) (results []*Pagenotfound, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

func (obj *_PagenotfoundMgr) GetFromIDPagenotfound(idPagenotfound uint32) (result Pagenotfound, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_pagenotfound = ?", idPagenotfound).Find(&result).Error

	return
}

func (obj *_PagenotfoundMgr) GetBatchFromIDPagenotfound(idPagenotfounds []uint32) (results []*Pagenotfound, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_pagenotfound IN (?)", idPagenotfounds).Find(&results).Error

	return
}

func (obj *_PagenotfoundMgr) GetFromIDShop(idShop uint32) (results []*Pagenotfound, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

func (obj *_PagenotfoundMgr) GetBatchFromIDShop(idShops []uint32) (results []*Pagenotfound, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

func (obj *_PagenotfoundMgr) GetFromIDShopGroup(idShopGroup uint32) (results []*Pagenotfound, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_group = ?", idShopGroup).Find(&results).Error

	return
}

func (obj *_PagenotfoundMgr) GetBatchFromIDShopGroup(idShopGroups []uint32) (results []*Pagenotfound, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_group IN (?)", idShopGroups).Find(&results).Error

	return
}

func (obj *_PagenotfoundMgr) GetFromRequestURI(requestURI string) (results []*Pagenotfound, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("request_uri = ?", requestURI).Find(&results).Error

	return
}

func (obj *_PagenotfoundMgr) GetBatchFromRequestURI(requestURIs []string) (results []*Pagenotfound, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("request_uri IN (?)", requestURIs).Find(&results).Error

	return
}

func (obj *_PagenotfoundMgr) GetFromHTTPReferer(httpReferer string) (results []*Pagenotfound, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("http_referer = ?", httpReferer).Find(&results).Error

	return
}

func (obj *_PagenotfoundMgr) GetBatchFromHTTPReferer(httpReferers []string) (results []*Pagenotfound, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("http_referer IN (?)", httpReferers).Find(&results).Error

	return
}

func (obj *_PagenotfoundMgr) GetFromDateAdd(dateAdd time.Time) (results []*Pagenotfound, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error

	return
}

func (obj *_PagenotfoundMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*Pagenotfound, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error

	return
}

func (obj *_PagenotfoundMgr) FetchByPrimaryKey(idPagenotfound uint32) (result Pagenotfound, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_pagenotfound = ?", idPagenotfound).Find(&result).Error

	return
}

func (obj *_PagenotfoundMgr) FetchIndexByDateAdd(dateAdd time.Time) (results []*Pagenotfound, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error

	return
}
