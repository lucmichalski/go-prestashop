package	model	
import (	
"time"	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgStatssearchMgr struct {
	*_BaseMgr
}

// EgStatssearchMgr open func
func EgStatssearchMgr(db *gorm.DB) *_EgStatssearchMgr {
	if db == nil {
		panic(fmt.Errorf("EgStatssearchMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgStatssearchMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_statssearch"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgStatssearchMgr) GetTableName() string {
	return "eg_statssearch"
}

// Get 获取
func (obj *_EgStatssearchMgr) Get() (result EgStatssearch, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgStatssearchMgr) Gets() (results []*EgStatssearch, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDStatssearch id_statssearch获取 
func (obj *_EgStatssearchMgr) WithIDStatssearch(idStatssearch uint32) Option {
	return optionFunc(func(o *options) { o.query["id_statssearch"] = idStatssearch })
}

// WithIDShop id_shop获取 
func (obj *_EgStatssearchMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

// WithIDShopGroup id_shop_group获取 
func (obj *_EgStatssearchMgr) WithIDShopGroup(idShopGroup uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop_group"] = idShopGroup })
}

// WithKeywords keywords获取 
func (obj *_EgStatssearchMgr) WithKeywords(keywords string) Option {
	return optionFunc(func(o *options) { o.query["keywords"] = keywords })
}

// WithResults results获取 
func (obj *_EgStatssearchMgr) WithResults(results int) Option {
	return optionFunc(func(o *options) { o.query["results"] = results })
}

// WithDateAdd date_add获取 
func (obj *_EgStatssearchMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}


// GetByOption 功能选项模式获取
func (obj *_EgStatssearchMgr) GetByOption(opts ...Option) (result EgStatssearch, err error) {
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
func (obj *_EgStatssearchMgr) GetByOptions(opts ...Option) (results []*EgStatssearch, err error) {
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


// GetFromIDStatssearch 通过id_statssearch获取内容  
func (obj *_EgStatssearchMgr)  GetFromIDStatssearch(idStatssearch uint32) (result EgStatssearch, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_statssearch = ?", idStatssearch).Find(&result).Error
	
	return
}

// GetBatchFromIDStatssearch 批量唯一主键查找 
func (obj *_EgStatssearchMgr) GetBatchFromIDStatssearch(idStatssearchs []uint32) (results []*EgStatssearch, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_statssearch IN (?)", idStatssearchs).Find(&results).Error
	
	return
}
 
// GetFromIDShop 通过id_shop获取内容  
func (obj *_EgStatssearchMgr) GetFromIDShop(idShop uint32) (results []*EgStatssearch, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error
	
	return
}

// GetBatchFromIDShop 批量唯一主键查找 
func (obj *_EgStatssearchMgr) GetBatchFromIDShop(idShops []uint32) (results []*EgStatssearch, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error
	
	return
}
 
// GetFromIDShopGroup 通过id_shop_group获取内容  
func (obj *_EgStatssearchMgr) GetFromIDShopGroup(idShopGroup uint32) (results []*EgStatssearch, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_group = ?", idShopGroup).Find(&results).Error
	
	return
}

// GetBatchFromIDShopGroup 批量唯一主键查找 
func (obj *_EgStatssearchMgr) GetBatchFromIDShopGroup(idShopGroups []uint32) (results []*EgStatssearch, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_group IN (?)", idShopGroups).Find(&results).Error
	
	return
}
 
// GetFromKeywords 通过keywords获取内容  
func (obj *_EgStatssearchMgr) GetFromKeywords(keywords string) (results []*EgStatssearch, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("keywords = ?", keywords).Find(&results).Error
	
	return
}

// GetBatchFromKeywords 批量唯一主键查找 
func (obj *_EgStatssearchMgr) GetBatchFromKeywords(keywordss []string) (results []*EgStatssearch, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("keywords IN (?)", keywordss).Find(&results).Error
	
	return
}
 
// GetFromResults 通过results获取内容  
func (obj *_EgStatssearchMgr) GetFromResults(results int) (results []*EgStatssearch, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("results = ?", results).Find(&results).Error
	
	return
}

// GetBatchFromResults 批量唯一主键查找 
func (obj *_EgStatssearchMgr) GetBatchFromResults(resultss []int) (results []*EgStatssearch, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("results IN (?)", resultss).Find(&results).Error
	
	return
}
 
// GetFromDateAdd 通过date_add获取内容  
func (obj *_EgStatssearchMgr) GetFromDateAdd(dateAdd time.Time) (results []*EgStatssearch, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error
	
	return
}

// GetBatchFromDateAdd 批量唯一主键查找 
func (obj *_EgStatssearchMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*EgStatssearch, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgStatssearchMgr) FetchByPrimaryKey(idStatssearch uint32 ) (result EgStatssearch, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_statssearch = ?", idStatssearch).Find(&result).Error
	
	return
}
 

 

	

