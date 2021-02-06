package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _ConnectionsMgr struct {
	*_BaseMgr
}

// ConnectionsMgr open func
func ConnectionsMgr(db *gorm.DB) *_ConnectionsMgr {
	if db == nil {
		panic(fmt.Errorf("ConnectionsMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ConnectionsMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_connections"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ConnectionsMgr) GetTableName() string {
	return "ps_connections"
}

// Get 获取
func (obj *_ConnectionsMgr) Get() (result Connections, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ConnectionsMgr) Gets() (results []*Connections, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDConnections id_connections获取
func (obj *_ConnectionsMgr) WithIDConnections(idConnections uint32) Option {
	return optionFunc(func(o *options) { o.query["id_connections"] = idConnections })
}

// WithIDShopGroup id_shop_group获取
func (obj *_ConnectionsMgr) WithIDShopGroup(idShopGroup uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop_group"] = idShopGroup })
}

// WithIDShop id_shop获取
func (obj *_ConnectionsMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

// WithIDGuest id_guest获取
func (obj *_ConnectionsMgr) WithIDGuest(idGuest uint32) Option {
	return optionFunc(func(o *options) { o.query["id_guest"] = idGuest })
}

// WithIDPage id_page获取
func (obj *_ConnectionsMgr) WithIDPage(idPage uint32) Option {
	return optionFunc(func(o *options) { o.query["id_page"] = idPage })
}

// WithIPAddress ip_address获取
func (obj *_ConnectionsMgr) WithIPAddress(ipAddress int64) Option {
	return optionFunc(func(o *options) { o.query["ip_address"] = ipAddress })
}

// WithDateAdd date_add获取
func (obj *_ConnectionsMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

// WithHTTPReferer http_referer获取
func (obj *_ConnectionsMgr) WithHTTPReferer(httpReferer string) Option {
	return optionFunc(func(o *options) { o.query["http_referer"] = httpReferer })
}

// GetByOption 功能选项模式获取
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

// GetByOptions 批量功能选项模式获取
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

//////////////////////////enume case ////////////////////////////////////////////

// GetFromIDConnections 通过id_connections获取内容
func (obj *_ConnectionsMgr) GetFromIDConnections(idConnections uint32) (result Connections, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_connections = ?", idConnections).Find(&result).Error

	return
}

// GetBatchFromIDConnections 批量唯一主键查找
func (obj *_ConnectionsMgr) GetBatchFromIDConnections(idConnectionss []uint32) (results []*Connections, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_connections IN (?)", idConnectionss).Find(&results).Error

	return
}

// GetFromIDShopGroup 通过id_shop_group获取内容
func (obj *_ConnectionsMgr) GetFromIDShopGroup(idShopGroup uint32) (results []*Connections, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_group = ?", idShopGroup).Find(&results).Error

	return
}

// GetBatchFromIDShopGroup 批量唯一主键查找
func (obj *_ConnectionsMgr) GetBatchFromIDShopGroup(idShopGroups []uint32) (results []*Connections, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_group IN (?)", idShopGroups).Find(&results).Error

	return
}

// GetFromIDShop 通过id_shop获取内容
func (obj *_ConnectionsMgr) GetFromIDShop(idShop uint32) (results []*Connections, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

// GetBatchFromIDShop 批量唯一主键查找
func (obj *_ConnectionsMgr) GetBatchFromIDShop(idShops []uint32) (results []*Connections, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

// GetFromIDGuest 通过id_guest获取内容
func (obj *_ConnectionsMgr) GetFromIDGuest(idGuest uint32) (results []*Connections, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_guest = ?", idGuest).Find(&results).Error

	return
}

// GetBatchFromIDGuest 批量唯一主键查找
func (obj *_ConnectionsMgr) GetBatchFromIDGuest(idGuests []uint32) (results []*Connections, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_guest IN (?)", idGuests).Find(&results).Error

	return
}

// GetFromIDPage 通过id_page获取内容
func (obj *_ConnectionsMgr) GetFromIDPage(idPage uint32) (results []*Connections, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_page = ?", idPage).Find(&results).Error

	return
}

// GetBatchFromIDPage 批量唯一主键查找
func (obj *_ConnectionsMgr) GetBatchFromIDPage(idPages []uint32) (results []*Connections, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_page IN (?)", idPages).Find(&results).Error

	return
}

// GetFromIPAddress 通过ip_address获取内容
func (obj *_ConnectionsMgr) GetFromIPAddress(ipAddress int64) (results []*Connections, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ip_address = ?", ipAddress).Find(&results).Error

	return
}

// GetBatchFromIPAddress 批量唯一主键查找
func (obj *_ConnectionsMgr) GetBatchFromIPAddress(ipAddresss []int64) (results []*Connections, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ip_address IN (?)", ipAddresss).Find(&results).Error

	return
}

// GetFromDateAdd 通过date_add获取内容
func (obj *_ConnectionsMgr) GetFromDateAdd(dateAdd time.Time) (results []*Connections, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error

	return
}

// GetBatchFromDateAdd 批量唯一主键查找
func (obj *_ConnectionsMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*Connections, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error

	return
}

// GetFromHTTPReferer 通过http_referer获取内容
func (obj *_ConnectionsMgr) GetFromHTTPReferer(httpReferer string) (results []*Connections, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("http_referer = ?", httpReferer).Find(&results).Error

	return
}

// GetBatchFromHTTPReferer 批量唯一主键查找
func (obj *_ConnectionsMgr) GetBatchFromHTTPReferer(httpReferers []string) (results []*Connections, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("http_referer IN (?)", httpReferers).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_ConnectionsMgr) FetchByPrimaryKey(idConnections uint32) (result Connections, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_connections = ?", idConnections).Find(&result).Error

	return
}

// FetchIndexByIDGuest  获取多个内容
func (obj *_ConnectionsMgr) FetchIndexByIDGuest(idGuest uint32) (results []*Connections, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_guest = ?", idGuest).Find(&results).Error

	return
}

// FetchIndexByIDPage  获取多个内容
func (obj *_ConnectionsMgr) FetchIndexByIDPage(idPage uint32) (results []*Connections, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_page = ?", idPage).Find(&results).Error

	return
}

// FetchIndexByDateAdd  获取多个内容
func (obj *_ConnectionsMgr) FetchIndexByDateAdd(dateAdd time.Time) (results []*Connections, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error

	return
}
