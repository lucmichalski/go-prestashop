package	model	
import (	
"gorm.io/gorm"	
"fmt"	
"context"	
)	

type _EgShopURLMgr struct {
	*_BaseMgr
}

// EgShopURLMgr open func
func EgShopURLMgr(db *gorm.DB) *_EgShopURLMgr {
	if db == nil {
		panic(fmt.Errorf("EgShopURLMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgShopURLMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_shop_url"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgShopURLMgr) GetTableName() string {
	return "eg_shop_url"
}

// Get 获取
func (obj *_EgShopURLMgr) Get() (result EgShopURL, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgShopURLMgr) Gets() (results []*EgShopURL, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDShopURL id_shop_url获取 
func (obj *_EgShopURLMgr) WithIDShopURL(idShopURL uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop_url"] = idShopURL })
}

// WithIDShop id_shop获取 
func (obj *_EgShopURLMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

// WithDomain domain获取 
func (obj *_EgShopURLMgr) WithDomain(domain string) Option {
	return optionFunc(func(o *options) { o.query["domain"] = domain })
}

// WithDomainSsl domain_ssl获取 
func (obj *_EgShopURLMgr) WithDomainSsl(domainSsl string) Option {
	return optionFunc(func(o *options) { o.query["domain_ssl"] = domainSsl })
}

// WithPhysicalURI physical_uri获取 
func (obj *_EgShopURLMgr) WithPhysicalURI(physicalURI string) Option {
	return optionFunc(func(o *options) { o.query["physical_uri"] = physicalURI })
}

// WithVirtualURI virtual_uri获取 
func (obj *_EgShopURLMgr) WithVirtualURI(virtualURI string) Option {
	return optionFunc(func(o *options) { o.query["virtual_uri"] = virtualURI })
}

// WithMain main获取 
func (obj *_EgShopURLMgr) WithMain(main bool) Option {
	return optionFunc(func(o *options) { o.query["main"] = main })
}

// WithActive active获取 
func (obj *_EgShopURLMgr) WithActive(active bool) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}


// GetByOption 功能选项模式获取
func (obj *_EgShopURLMgr) GetByOption(opts ...Option) (result EgShopURL, err error) {
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
func (obj *_EgShopURLMgr) GetByOptions(opts ...Option) (results []*EgShopURL, err error) {
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


// GetFromIDShopURL 通过id_shop_url获取内容  
func (obj *_EgShopURLMgr)  GetFromIDShopURL(idShopURL uint32) (result EgShopURL, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_url = ?", idShopURL).Find(&result).Error
	
	return
}

// GetBatchFromIDShopURL 批量唯一主键查找 
func (obj *_EgShopURLMgr) GetBatchFromIDShopURL(idShopURLs []uint32) (results []*EgShopURL, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_url IN (?)", idShopURLs).Find(&results).Error
	
	return
}
 
// GetFromIDShop 通过id_shop获取内容  
func (obj *_EgShopURLMgr) GetFromIDShop(idShop uint32) (results []*EgShopURL, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error
	
	return
}

// GetBatchFromIDShop 批量唯一主键查找 
func (obj *_EgShopURLMgr) GetBatchFromIDShop(idShops []uint32) (results []*EgShopURL, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error
	
	return
}
 
// GetFromDomain 通过domain获取内容  
func (obj *_EgShopURLMgr) GetFromDomain(domain string) (results []*EgShopURL, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("domain = ?", domain).Find(&results).Error
	
	return
}

// GetBatchFromDomain 批量唯一主键查找 
func (obj *_EgShopURLMgr) GetBatchFromDomain(domains []string) (results []*EgShopURL, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("domain IN (?)", domains).Find(&results).Error
	
	return
}
 
// GetFromDomainSsl 通过domain_ssl获取内容  
func (obj *_EgShopURLMgr) GetFromDomainSsl(domainSsl string) (results []*EgShopURL, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("domain_ssl = ?", domainSsl).Find(&results).Error
	
	return
}

// GetBatchFromDomainSsl 批量唯一主键查找 
func (obj *_EgShopURLMgr) GetBatchFromDomainSsl(domainSsls []string) (results []*EgShopURL, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("domain_ssl IN (?)", domainSsls).Find(&results).Error
	
	return
}
 
// GetFromPhysicalURI 通过physical_uri获取内容  
func (obj *_EgShopURLMgr) GetFromPhysicalURI(physicalURI string) (results []*EgShopURL, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("physical_uri = ?", physicalURI).Find(&results).Error
	
	return
}

// GetBatchFromPhysicalURI 批量唯一主键查找 
func (obj *_EgShopURLMgr) GetBatchFromPhysicalURI(physicalURIs []string) (results []*EgShopURL, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("physical_uri IN (?)", physicalURIs).Find(&results).Error
	
	return
}
 
// GetFromVirtualURI 通过virtual_uri获取内容  
func (obj *_EgShopURLMgr) GetFromVirtualURI(virtualURI string) (results []*EgShopURL, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("virtual_uri = ?", virtualURI).Find(&results).Error
	
	return
}

// GetBatchFromVirtualURI 批量唯一主键查找 
func (obj *_EgShopURLMgr) GetBatchFromVirtualURI(virtualURIs []string) (results []*EgShopURL, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("virtual_uri IN (?)", virtualURIs).Find(&results).Error
	
	return
}
 
// GetFromMain 通过main获取内容  
func (obj *_EgShopURLMgr) GetFromMain(main bool) (results []*EgShopURL, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("main = ?", main).Find(&results).Error
	
	return
}

// GetBatchFromMain 批量唯一主键查找 
func (obj *_EgShopURLMgr) GetBatchFromMain(mains []bool) (results []*EgShopURL, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("main IN (?)", mains).Find(&results).Error
	
	return
}
 
// GetFromActive 通过active获取内容  
func (obj *_EgShopURLMgr) GetFromActive(active bool) (results []*EgShopURL, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active = ?", active).Find(&results).Error
	
	return
}

// GetBatchFromActive 批量唯一主键查找 
func (obj *_EgShopURLMgr) GetBatchFromActive(actives []bool) (results []*EgShopURL, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active IN (?)", actives).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgShopURLMgr) FetchByPrimaryKey(idShopURL uint32 ) (result EgShopURL, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_url = ?", idShopURL).Find(&result).Error
	
	return
}
 
 // FetchUniqueIndexByFullShopURL primay or index 获取唯一内容
 func (obj *_EgShopURLMgr) FetchUniqueIndexByFullShopURL(domain string ,physicalURI string ,virtualURI string ) (result EgShopURL, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("domain = ? AND physical_uri = ? AND virtual_uri = ?", domain , physicalURI , virtualURI).Find(&result).Error
	
	return
}
 
 // FetchUniqueIndexByFullShopURLSsl primay or index 获取唯一内容
 func (obj *_EgShopURLMgr) FetchUniqueIndexByFullShopURLSsl(domainSsl string ,physicalURI string ,virtualURI string ) (result EgShopURL, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("domain_ssl = ? AND physical_uri = ? AND virtual_uri = ?", domainSsl , physicalURI , virtualURI).Find(&result).Error
	
	return
}
 

 
 // FetchIndexByIDShop  获取多个内容
 func (obj *_EgShopURLMgr) FetchIndexByIDShop(idShop uint32 ,main bool ) (results []*EgShopURL, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ? AND main = ?", idShop , main).Find(&results).Error
	
	return
}
 

	

