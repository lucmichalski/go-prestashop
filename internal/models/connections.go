package models

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _ConnectionsMgr struct {
	*_BaseMgr
}

func ConnectionsMgr(db *gorm.DB) *_ConnectionsMgr {
	if db == nil {
		panic(fmt.Errorf("ConnectionsMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ConnectionsMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_connections"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_ConnectionsMgr) GetTableName() string {
	return "ps_connections"
}

func (obj *_ConnectionsMgr) Get() (result Connections, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_ConnectionsMgr) Gets() (results []*Connections, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

func (obj *_ConnectionsMgr) WithIDConnections(idConnections uint32) Option {
	return optionFunc(func(o *options) { o.query["id_connections"] = idConnections })
}

func (obj *_ConnectionsMgr) WithIDShopGroup(idShopGroup uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop_group"] = idShopGroup })
}

func (obj *_ConnectionsMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

func (obj *_ConnectionsMgr) WithIDGuest(idGuest uint32) Option {
	return optionFunc(func(o *options) { o.query["id_guest"] = idGuest })
}

func (obj *_ConnectionsMgr) WithIDPage(idPage uint32) Option {
	return optionFunc(func(o *options) { o.query["id_page"] = idPage })
}

func (obj *_ConnectionsMgr) WithIPAddress(ipAddress int64) Option {
	return optionFunc(func(o *options) { o.query["ip_address"] = ipAddress })
}

func (obj *_ConnectionsMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

func (obj *_ConnectionsMgr) WithHTTPReferer(httpReferer string) Option {
	return optionFunc(func(o *options) { o.query["http_referer"] = httpReferer })
}

func (obj *_ConnectionsMgr) GetByOption(opts ...Option) (result Connections, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_ConnectionsMgr) GetByOptions(opts ...Option) (results []*Connections, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

func (obj *_ConnectionsMgr) GetFromIDConnections(idConnections uint32) (result Connections, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_connections = ?", idConnections).Find(&result).Error

	return
}

func (obj *_ConnectionsMgr) GetBatchFromIDConnections(idConnectionss []uint32) (results []*Connections, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_connections IN (?)", idConnectionss).Find(&results).Error

	return
}

func (obj *_ConnectionsMgr) GetFromIDShopGroup(idShopGroup uint32) (results []*Connections, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_group = ?", idShopGroup).Find(&results).Error

	return
}

func (obj *_ConnectionsMgr) GetBatchFromIDShopGroup(idShopGroups []uint32) (results []*Connections, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_group IN (?)", idShopGroups).Find(&results).Error

	return
}

func (obj *_ConnectionsMgr) GetFromIDShop(idShop uint32) (results []*Connections, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

func (obj *_ConnectionsMgr) GetBatchFromIDShop(idShops []uint32) (results []*Connections, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

func (obj *_ConnectionsMgr) GetFromIDGuest(idGuest uint32) (results []*Connections, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_guest = ?", idGuest).Find(&results).Error

	return
}

func (obj *_ConnectionsMgr) GetBatchFromIDGuest(idGuests []uint32) (results []*Connections, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_guest IN (?)", idGuests).Find(&results).Error

	return
}

func (obj *_ConnectionsMgr) GetFromIDPage(idPage uint32) (results []*Connections, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_page = ?", idPage).Find(&results).Error

	return
}

func (obj *_ConnectionsMgr) GetBatchFromIDPage(idPages []uint32) (results []*Connections, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_page IN (?)", idPages).Find(&results).Error

	return
}

func (obj *_ConnectionsMgr) GetFromIPAddress(ipAddress int64) (results []*Connections, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ip_address = ?", ipAddress).Find(&results).Error

	return
}

func (obj *_ConnectionsMgr) GetBatchFromIPAddress(ipAddresss []int64) (results []*Connections, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ip_address IN (?)", ipAddresss).Find(&results).Error

	return
}

func (obj *_ConnectionsMgr) GetFromDateAdd(dateAdd time.Time) (results []*Connections, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error

	return
}

func (obj *_ConnectionsMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*Connections, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error

	return
}

func (obj *_ConnectionsMgr) GetFromHTTPReferer(httpReferer string) (results []*Connections, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("http_referer = ?", httpReferer).Find(&results).Error

	return
}

func (obj *_ConnectionsMgr) GetBatchFromHTTPReferer(httpReferers []string) (results []*Connections, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("http_referer IN (?)", httpReferers).Find(&results).Error

	return
}

func (obj *_ConnectionsMgr) FetchByPrimaryKey(idConnections uint32) (result Connections, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_connections = ?", idConnections).Find(&result).Error

	return
}

func (obj *_ConnectionsMgr) FetchIndexByIDGuest(idGuest uint32) (results []*Connections, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_guest = ?", idGuest).Find(&results).Error

	return
}

func (obj *_ConnectionsMgr) FetchIndexByIDPage(idPage uint32) (results []*Connections, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_page = ?", idPage).Find(&results).Error

	return
}

func (obj *_ConnectionsMgr) FetchIndexByDateAdd(dateAdd time.Time) (results []*Connections, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error

	return
}
