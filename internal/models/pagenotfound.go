package	model	
import (	
"time"	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgPagenotfoundMgr struct {
	*_BaseMgr
}

// EgPagenotfoundMgr open func
func EgPagenotfoundMgr(db *gorm.DB) *_EgPagenotfoundMgr {
	if db == nil {
		panic(fmt.Errorf("EgPagenotfoundMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgPagenotfoundMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_pagenotfound"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgPagenotfoundMgr) GetTableName() string {
	return "eg_pagenotfound"
}

// Get 获取
func (obj *_EgPagenotfoundMgr) Get() (result EgPagenotfound, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgPagenotfoundMgr) Gets() (results []*EgPagenotfound, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDPagenotfound id_pagenotfound获取 
func (obj *_EgPagenotfoundMgr) WithIDPagenotfound(idPagenotfound uint32) Option {
	return optionFunc(func(o *options) { o.query["id_pagenotfound"] = idPagenotfound })
}

// WithIDShop id_shop获取 
func (obj *_EgPagenotfoundMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

// WithIDShopGroup id_shop_group获取 
func (obj *_EgPagenotfoundMgr) WithIDShopGroup(idShopGroup uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop_group"] = idShopGroup })
}

// WithRequestURI request_uri获取 
func (obj *_EgPagenotfoundMgr) WithRequestURI(requestURI string) Option {
	return optionFunc(func(o *options) { o.query["request_uri"] = requestURI })
}

// WithHTTPReferer http_referer获取 
func (obj *_EgPagenotfoundMgr) WithHTTPReferer(httpReferer string) Option {
	return optionFunc(func(o *options) { o.query["http_referer"] = httpReferer })
}

// WithDateAdd date_add获取 
func (obj *_EgPagenotfoundMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}


// GetByOption 功能选项模式获取
func (obj *_EgPagenotfoundMgr) GetByOption(opts ...Option) (result EgPagenotfound, err error) {
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
func (obj *_EgPagenotfoundMgr) GetByOptions(opts ...Option) (results []*EgPagenotfound, err error) {
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


// GetFromIDPagenotfound 通过id_pagenotfound获取内容  
func (obj *_EgPagenotfoundMgr)  GetFromIDPagenotfound(idPagenotfound uint32) (result EgPagenotfound, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_pagenotfound = ?", idPagenotfound).Find(&result).Error
	
	return
}

// GetBatchFromIDPagenotfound 批量唯一主键查找 
func (obj *_EgPagenotfoundMgr) GetBatchFromIDPagenotfound(idPagenotfounds []uint32) (results []*EgPagenotfound, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_pagenotfound IN (?)", idPagenotfounds).Find(&results).Error
	
	return
}
 
// GetFromIDShop 通过id_shop获取内容  
func (obj *_EgPagenotfoundMgr) GetFromIDShop(idShop uint32) (results []*EgPagenotfound, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error
	
	return
}

// GetBatchFromIDShop 批量唯一主键查找 
func (obj *_EgPagenotfoundMgr) GetBatchFromIDShop(idShops []uint32) (results []*EgPagenotfound, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error
	
	return
}
 
// GetFromIDShopGroup 通过id_shop_group获取内容  
func (obj *_EgPagenotfoundMgr) GetFromIDShopGroup(idShopGroup uint32) (results []*EgPagenotfound, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_group = ?", idShopGroup).Find(&results).Error
	
	return
}

// GetBatchFromIDShopGroup 批量唯一主键查找 
func (obj *_EgPagenotfoundMgr) GetBatchFromIDShopGroup(idShopGroups []uint32) (results []*EgPagenotfound, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_group IN (?)", idShopGroups).Find(&results).Error
	
	return
}
 
// GetFromRequestURI 通过request_uri获取内容  
func (obj *_EgPagenotfoundMgr) GetFromRequestURI(requestURI string) (results []*EgPagenotfound, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("request_uri = ?", requestURI).Find(&results).Error
	
	return
}

// GetBatchFromRequestURI 批量唯一主键查找 
func (obj *_EgPagenotfoundMgr) GetBatchFromRequestURI(requestURIs []string) (results []*EgPagenotfound, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("request_uri IN (?)", requestURIs).Find(&results).Error
	
	return
}
 
// GetFromHTTPReferer 通过http_referer获取内容  
func (obj *_EgPagenotfoundMgr) GetFromHTTPReferer(httpReferer string) (results []*EgPagenotfound, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("http_referer = ?", httpReferer).Find(&results).Error
	
	return
}

// GetBatchFromHTTPReferer 批量唯一主键查找 
func (obj *_EgPagenotfoundMgr) GetBatchFromHTTPReferer(httpReferers []string) (results []*EgPagenotfound, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("http_referer IN (?)", httpReferers).Find(&results).Error
	
	return
}
 
// GetFromDateAdd 通过date_add获取内容  
func (obj *_EgPagenotfoundMgr) GetFromDateAdd(dateAdd time.Time) (results []*EgPagenotfound, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error
	
	return
}

// GetBatchFromDateAdd 批量唯一主键查找 
func (obj *_EgPagenotfoundMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*EgPagenotfound, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgPagenotfoundMgr) FetchByPrimaryKey(idPagenotfound uint32 ) (result EgPagenotfound, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_pagenotfound = ?", idPagenotfound).Find(&result).Error
	
	return
}
 

 
 // FetchIndexByDateAdd  获取多个内容
 func (obj *_EgPagenotfoundMgr) FetchIndexByDateAdd(dateAdd time.Time ) (results []*EgPagenotfound, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error
	
	return
}
 

	

