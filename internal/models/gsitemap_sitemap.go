package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgGsitemapSitemapMgr struct {
	*_BaseMgr
}

// EgGsitemapSitemapMgr open func
func EgGsitemapSitemapMgr(db *gorm.DB) *_EgGsitemapSitemapMgr {
	if db == nil {
		panic(fmt.Errorf("EgGsitemapSitemapMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgGsitemapSitemapMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_gsitemap_sitemap"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgGsitemapSitemapMgr) GetTableName() string {
	return "eg_gsitemap_sitemap"
}

// Get 获取
func (obj *_EgGsitemapSitemapMgr) Get() (result EgGsitemapSitemap, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgGsitemapSitemapMgr) Gets() (results []*EgGsitemapSitemap, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithLink link获取 
func (obj *_EgGsitemapSitemapMgr) WithLink(link string) Option {
	return optionFunc(func(o *options) { o.query["link"] = link })
}

// WithIDShop id_shop获取 
func (obj *_EgGsitemapSitemapMgr) WithIDShop(idShop int) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}


// GetByOption 功能选项模式获取
func (obj *_EgGsitemapSitemapMgr) GetByOption(opts ...Option) (result EgGsitemapSitemap, err error) {
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
func (obj *_EgGsitemapSitemapMgr) GetByOptions(opts ...Option) (results []*EgGsitemapSitemap, err error) {
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
func (obj *_EgGsitemapSitemapMgr) GetFromLink(link string) (results []*EgGsitemapSitemap, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("link = ?", link).Find(&results).Error
	
	return
}

// GetBatchFromLink 批量唯一主键查找 
func (obj *_EgGsitemapSitemapMgr) GetBatchFromLink(links []string) (results []*EgGsitemapSitemap, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("link IN (?)", links).Find(&results).Error
	
	return
}
 
// GetFromIDShop 通过id_shop获取内容  
func (obj *_EgGsitemapSitemapMgr) GetFromIDShop(idShop int) (results []*EgGsitemapSitemap, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error
	
	return
}

// GetBatchFromIDShop 批量唯一主键查找 
func (obj *_EgGsitemapSitemapMgr) GetBatchFromIDShop(idShops []int) (results []*EgGsitemapSitemap, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 

 

	

