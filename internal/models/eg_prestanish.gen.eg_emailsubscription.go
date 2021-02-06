package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
"time"	
)	

type _EgEmailsubscriptionMgr struct {
	*_BaseMgr
}

// EgEmailsubscriptionMgr open func
func EgEmailsubscriptionMgr(db *gorm.DB) *_EgEmailsubscriptionMgr {
	if db == nil {
		panic(fmt.Errorf("EgEmailsubscriptionMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgEmailsubscriptionMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_emailsubscription"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgEmailsubscriptionMgr) GetTableName() string {
	return "eg_emailsubscription"
}

// Get 获取
func (obj *_EgEmailsubscriptionMgr) Get() (result EgEmailsubscription, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgEmailsubscriptionMgr) Gets() (results []*EgEmailsubscription, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 
func (obj *_EgEmailsubscriptionMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithIDShop id_shop获取 
func (obj *_EgEmailsubscriptionMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

// WithIDShopGroup id_shop_group获取 
func (obj *_EgEmailsubscriptionMgr) WithIDShopGroup(idShopGroup uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop_group"] = idShopGroup })
}

// WithEmail email获取 
func (obj *_EgEmailsubscriptionMgr) WithEmail(email string) Option {
	return optionFunc(func(o *options) { o.query["email"] = email })
}

// WithNewsletterDateAdd newsletter_date_add获取 
func (obj *_EgEmailsubscriptionMgr) WithNewsletterDateAdd(newsletterDateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["newsletter_date_add"] = newsletterDateAdd })
}

// WithIPRegistrationNewsletter ip_registration_newsletter获取 
func (obj *_EgEmailsubscriptionMgr) WithIPRegistrationNewsletter(ipRegistrationNewsletter string) Option {
	return optionFunc(func(o *options) { o.query["ip_registration_newsletter"] = ipRegistrationNewsletter })
}

// WithHTTPReferer http_referer获取 
func (obj *_EgEmailsubscriptionMgr) WithHTTPReferer(httpReferer string) Option {
	return optionFunc(func(o *options) { o.query["http_referer"] = httpReferer })
}

// WithActive active获取 
func (obj *_EgEmailsubscriptionMgr) WithActive(active bool) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}

// WithIDLang id_lang获取 
func (obj *_EgEmailsubscriptionMgr) WithIDLang(idLang int) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}


// GetByOption 功能选项模式获取
func (obj *_EgEmailsubscriptionMgr) GetByOption(opts ...Option) (result EgEmailsubscription, err error) {
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
func (obj *_EgEmailsubscriptionMgr) GetByOptions(opts ...Option) (results []*EgEmailsubscription, err error) {
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


// GetFromID 通过id获取内容  
func (obj *_EgEmailsubscriptionMgr)  GetFromID(id int) (result EgEmailsubscription, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error
	
	return
}

// GetBatchFromID 批量唯一主键查找 
func (obj *_EgEmailsubscriptionMgr) GetBatchFromID(ids []int) (results []*EgEmailsubscription, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error
	
	return
}
 
// GetFromIDShop 通过id_shop获取内容  
func (obj *_EgEmailsubscriptionMgr) GetFromIDShop(idShop uint32) (results []*EgEmailsubscription, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error
	
	return
}

// GetBatchFromIDShop 批量唯一主键查找 
func (obj *_EgEmailsubscriptionMgr) GetBatchFromIDShop(idShops []uint32) (results []*EgEmailsubscription, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error
	
	return
}
 
// GetFromIDShopGroup 通过id_shop_group获取内容  
func (obj *_EgEmailsubscriptionMgr) GetFromIDShopGroup(idShopGroup uint32) (results []*EgEmailsubscription, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_group = ?", idShopGroup).Find(&results).Error
	
	return
}

// GetBatchFromIDShopGroup 批量唯一主键查找 
func (obj *_EgEmailsubscriptionMgr) GetBatchFromIDShopGroup(idShopGroups []uint32) (results []*EgEmailsubscription, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_group IN (?)", idShopGroups).Find(&results).Error
	
	return
}
 
// GetFromEmail 通过email获取内容  
func (obj *_EgEmailsubscriptionMgr) GetFromEmail(email string) (results []*EgEmailsubscription, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("email = ?", email).Find(&results).Error
	
	return
}

// GetBatchFromEmail 批量唯一主键查找 
func (obj *_EgEmailsubscriptionMgr) GetBatchFromEmail(emails []string) (results []*EgEmailsubscription, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("email IN (?)", emails).Find(&results).Error
	
	return
}
 
// GetFromNewsletterDateAdd 通过newsletter_date_add获取内容  
func (obj *_EgEmailsubscriptionMgr) GetFromNewsletterDateAdd(newsletterDateAdd time.Time) (results []*EgEmailsubscription, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("newsletter_date_add = ?", newsletterDateAdd).Find(&results).Error
	
	return
}

// GetBatchFromNewsletterDateAdd 批量唯一主键查找 
func (obj *_EgEmailsubscriptionMgr) GetBatchFromNewsletterDateAdd(newsletterDateAdds []time.Time) (results []*EgEmailsubscription, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("newsletter_date_add IN (?)", newsletterDateAdds).Find(&results).Error
	
	return
}
 
// GetFromIPRegistrationNewsletter 通过ip_registration_newsletter获取内容  
func (obj *_EgEmailsubscriptionMgr) GetFromIPRegistrationNewsletter(ipRegistrationNewsletter string) (results []*EgEmailsubscription, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ip_registration_newsletter = ?", ipRegistrationNewsletter).Find(&results).Error
	
	return
}

// GetBatchFromIPRegistrationNewsletter 批量唯一主键查找 
func (obj *_EgEmailsubscriptionMgr) GetBatchFromIPRegistrationNewsletter(ipRegistrationNewsletters []string) (results []*EgEmailsubscription, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ip_registration_newsletter IN (?)", ipRegistrationNewsletters).Find(&results).Error
	
	return
}
 
// GetFromHTTPReferer 通过http_referer获取内容  
func (obj *_EgEmailsubscriptionMgr) GetFromHTTPReferer(httpReferer string) (results []*EgEmailsubscription, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("http_referer = ?", httpReferer).Find(&results).Error
	
	return
}

// GetBatchFromHTTPReferer 批量唯一主键查找 
func (obj *_EgEmailsubscriptionMgr) GetBatchFromHTTPReferer(httpReferers []string) (results []*EgEmailsubscription, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("http_referer IN (?)", httpReferers).Find(&results).Error
	
	return
}
 
// GetFromActive 通过active获取内容  
func (obj *_EgEmailsubscriptionMgr) GetFromActive(active bool) (results []*EgEmailsubscription, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active = ?", active).Find(&results).Error
	
	return
}

// GetBatchFromActive 批量唯一主键查找 
func (obj *_EgEmailsubscriptionMgr) GetBatchFromActive(actives []bool) (results []*EgEmailsubscription, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active IN (?)", actives).Find(&results).Error
	
	return
}
 
// GetFromIDLang 通过id_lang获取内容  
func (obj *_EgEmailsubscriptionMgr) GetFromIDLang(idLang int) (results []*EgEmailsubscription, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error
	
	return
}

// GetBatchFromIDLang 批量唯一主键查找 
func (obj *_EgEmailsubscriptionMgr) GetBatchFromIDLang(idLangs []int) (results []*EgEmailsubscription, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgEmailsubscriptionMgr) FetchByPrimaryKey(id int ) (result EgEmailsubscription, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error
	
	return
}
 

 

	

