package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _PageViewedMgr struct {
	*_BaseMgr
}

// PageViewedMgr open func
func PageViewedMgr(db *gorm.DB) *_PageViewedMgr {
	if db == nil {
		panic(fmt.Errorf("PageViewedMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_PageViewedMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_page_viewed"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_PageViewedMgr) GetTableName() string {
	return "eg_page_viewed"
}

// Get 获取
func (obj *_PageViewedMgr) Get() (result PageViewed, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_PageViewedMgr) Gets() (results []*PageViewed, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDPage id_page获取
func (obj *_PageViewedMgr) WithIDPage(idPage uint32) Option {
	return optionFunc(func(o *options) { o.query["id_page"] = idPage })
}

// WithIDShopGroup id_shop_group获取
func (obj *_PageViewedMgr) WithIDShopGroup(idShopGroup uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop_group"] = idShopGroup })
}

// WithIDShop id_shop获取
func (obj *_PageViewedMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

// WithIDDateRange id_date_range获取
func (obj *_PageViewedMgr) WithIDDateRange(idDateRange uint32) Option {
	return optionFunc(func(o *options) { o.query["id_date_range"] = idDateRange })
}

// WithCounter counter获取
func (obj *_PageViewedMgr) WithCounter(counter uint32) Option {
	return optionFunc(func(o *options) { o.query["counter"] = counter })
}

// GetByOption 功能选项模式获取
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

// GetByOptions 批量功能选项模式获取
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

//////////////////////////enume case ////////////////////////////////////////////

// GetFromIDPage 通过id_page获取内容
func (obj *_PageViewedMgr) GetFromIDPage(idPage uint32) (results []*PageViewed, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_page = ?", idPage).Find(&results).Error

	return
}

// GetBatchFromIDPage 批量唯一主键查找
func (obj *_PageViewedMgr) GetBatchFromIDPage(idPages []uint32) (results []*PageViewed, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_page IN (?)", idPages).Find(&results).Error

	return
}

// GetFromIDShopGroup 通过id_shop_group获取内容
func (obj *_PageViewedMgr) GetFromIDShopGroup(idShopGroup uint32) (results []*PageViewed, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_group = ?", idShopGroup).Find(&results).Error

	return
}

// GetBatchFromIDShopGroup 批量唯一主键查找
func (obj *_PageViewedMgr) GetBatchFromIDShopGroup(idShopGroups []uint32) (results []*PageViewed, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_group IN (?)", idShopGroups).Find(&results).Error

	return
}

// GetFromIDShop 通过id_shop获取内容
func (obj *_PageViewedMgr) GetFromIDShop(idShop uint32) (results []*PageViewed, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

// GetBatchFromIDShop 批量唯一主键查找
func (obj *_PageViewedMgr) GetBatchFromIDShop(idShops []uint32) (results []*PageViewed, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

// GetFromIDDateRange 通过id_date_range获取内容
func (obj *_PageViewedMgr) GetFromIDDateRange(idDateRange uint32) (results []*PageViewed, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_date_range = ?", idDateRange).Find(&results).Error

	return
}

// GetBatchFromIDDateRange 批量唯一主键查找
func (obj *_PageViewedMgr) GetBatchFromIDDateRange(idDateRanges []uint32) (results []*PageViewed, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_date_range IN (?)", idDateRanges).Find(&results).Error

	return
}

// GetFromCounter 通过counter获取内容
func (obj *_PageViewedMgr) GetFromCounter(counter uint32) (results []*PageViewed, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("counter = ?", counter).Find(&results).Error

	return
}

// GetBatchFromCounter 批量唯一主键查找
func (obj *_PageViewedMgr) GetBatchFromCounter(counters []uint32) (results []*PageViewed, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("counter IN (?)", counters).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_PageViewedMgr) FetchByPrimaryKey(idPage uint32, idShop uint32, idDateRange uint32) (result PageViewed, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_page = ? AND id_shop = ? AND id_date_range = ?", idPage, idShop, idDateRange).Find(&result).Error

	return
}
