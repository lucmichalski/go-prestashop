package	model	
import (	
"gorm.io/gorm"	
"fmt"	
"context"	
)	

type _EgPageViewedMgr struct {
	*_BaseMgr
}

// EgPageViewedMgr open func
func EgPageViewedMgr(db *gorm.DB) *_EgPageViewedMgr {
	if db == nil {
		panic(fmt.Errorf("EgPageViewedMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgPageViewedMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_page_viewed"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgPageViewedMgr) GetTableName() string {
	return "eg_page_viewed"
}

// Get 获取
func (obj *_EgPageViewedMgr) Get() (result EgPageViewed, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgPageViewedMgr) Gets() (results []*EgPageViewed, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDPage id_page获取 
func (obj *_EgPageViewedMgr) WithIDPage(idPage uint32) Option {
	return optionFunc(func(o *options) { o.query["id_page"] = idPage })
}

// WithIDShopGroup id_shop_group获取 
func (obj *_EgPageViewedMgr) WithIDShopGroup(idShopGroup uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop_group"] = idShopGroup })
}

// WithIDShop id_shop获取 
func (obj *_EgPageViewedMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

// WithIDDateRange id_date_range获取 
func (obj *_EgPageViewedMgr) WithIDDateRange(idDateRange uint32) Option {
	return optionFunc(func(o *options) { o.query["id_date_range"] = idDateRange })
}

// WithCounter counter获取 
func (obj *_EgPageViewedMgr) WithCounter(counter uint32) Option {
	return optionFunc(func(o *options) { o.query["counter"] = counter })
}


// GetByOption 功能选项模式获取
func (obj *_EgPageViewedMgr) GetByOption(opts ...Option) (result EgPageViewed, err error) {
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
func (obj *_EgPageViewedMgr) GetByOptions(opts ...Option) (results []*EgPageViewed, err error) {
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
func (obj *_EgPageViewedMgr) GetFromIDPage(idPage uint32) (results []*EgPageViewed, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_page = ?", idPage).Find(&results).Error
	
	return
}

// GetBatchFromIDPage 批量唯一主键查找 
func (obj *_EgPageViewedMgr) GetBatchFromIDPage(idPages []uint32) (results []*EgPageViewed, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_page IN (?)", idPages).Find(&results).Error
	
	return
}
 
// GetFromIDShopGroup 通过id_shop_group获取内容  
func (obj *_EgPageViewedMgr) GetFromIDShopGroup(idShopGroup uint32) (results []*EgPageViewed, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_group = ?", idShopGroup).Find(&results).Error
	
	return
}

// GetBatchFromIDShopGroup 批量唯一主键查找 
func (obj *_EgPageViewedMgr) GetBatchFromIDShopGroup(idShopGroups []uint32) (results []*EgPageViewed, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_group IN (?)", idShopGroups).Find(&results).Error
	
	return
}
 
// GetFromIDShop 通过id_shop获取内容  
func (obj *_EgPageViewedMgr) GetFromIDShop(idShop uint32) (results []*EgPageViewed, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error
	
	return
}

// GetBatchFromIDShop 批量唯一主键查找 
func (obj *_EgPageViewedMgr) GetBatchFromIDShop(idShops []uint32) (results []*EgPageViewed, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error
	
	return
}
 
// GetFromIDDateRange 通过id_date_range获取内容  
func (obj *_EgPageViewedMgr) GetFromIDDateRange(idDateRange uint32) (results []*EgPageViewed, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_date_range = ?", idDateRange).Find(&results).Error
	
	return
}

// GetBatchFromIDDateRange 批量唯一主键查找 
func (obj *_EgPageViewedMgr) GetBatchFromIDDateRange(idDateRanges []uint32) (results []*EgPageViewed, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_date_range IN (?)", idDateRanges).Find(&results).Error
	
	return
}
 
// GetFromCounter 通过counter获取内容  
func (obj *_EgPageViewedMgr) GetFromCounter(counter uint32) (results []*EgPageViewed, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("counter = ?", counter).Find(&results).Error
	
	return
}

// GetBatchFromCounter 批量唯一主键查找 
func (obj *_EgPageViewedMgr) GetBatchFromCounter(counters []uint32) (results []*EgPageViewed, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("counter IN (?)", counters).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgPageViewedMgr) FetchByPrimaryKey(idPage uint32 ,idShop uint32 ,idDateRange uint32 ) (result EgPageViewed, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_page = ? AND id_shop = ? AND id_date_range = ?", idPage , idShop , idDateRange).Find(&result).Error
	
	return
}
 

 

	

