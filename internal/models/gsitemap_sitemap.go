package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _GsitemapSitemapMgr struct {
	*_BaseMgr
}

// GsitemapSitemapMgr open func
func GsitemapSitemapMgr(db *gorm.DB) *_GsitemapSitemapMgr {
	if db == nil {
		panic(fmt.Errorf("GsitemapSitemapMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_GsitemapSitemapMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_gsitemap_sitemap"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_GsitemapSitemapMgr) GetTableName() string {
	return "eg_gsitemap_sitemap"
}

// Get 获取
func (obj *_GsitemapSitemapMgr) Get() (result GsitemapSitemap, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_GsitemapSitemapMgr) Gets() (results []*GsitemapSitemap, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithLink link获取
func (obj *_GsitemapSitemapMgr) WithLink(link string) Option {
	return optionFunc(func(o *options) { o.query["link"] = link })
}

// WithIDShop id_shop获取
func (obj *_GsitemapSitemapMgr) WithIDShop(idShop int) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

// GetByOption 功能选项模式获取
func (obj *_GsitemapSitemapMgr) GetByOption(opts ...Option) (result GsitemapSitemap, err error) {
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
func (obj *_GsitemapSitemapMgr) GetByOptions(opts ...Option) (results []*GsitemapSitemap, err error) {
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

// GetFromLink 通过link获取内容
func (obj *_GsitemapSitemapMgr) GetFromLink(link string) (results []*GsitemapSitemap, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("link = ?", link).Find(&results).Error

	return
}

// GetBatchFromLink 批量唯一主键查找
func (obj *_GsitemapSitemapMgr) GetBatchFromLink(links []string) (results []*GsitemapSitemap, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("link IN (?)", links).Find(&results).Error

	return
}

// GetFromIDShop 通过id_shop获取内容
func (obj *_GsitemapSitemapMgr) GetFromIDShop(idShop int) (results []*GsitemapSitemap, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

// GetBatchFromIDShop 批量唯一主键查找
func (obj *_GsitemapSitemapMgr) GetBatchFromIDShop(idShops []int) (results []*GsitemapSitemap, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
