package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _EmailsubscriptionMgr struct {
	*_BaseMgr
}

// EmailsubscriptionMgr open func
func EmailsubscriptionMgr(db *gorm.DB) *_EmailsubscriptionMgr {
	if db == nil {
		panic(fmt.Errorf("EmailsubscriptionMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EmailsubscriptionMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_emailsubscription"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EmailsubscriptionMgr) GetTableName() string {
	return "ps_emailsubscription"
}

// Get 获取
func (obj *_EmailsubscriptionMgr) Get() (result Emailsubscription, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EmailsubscriptionMgr) Gets() (results []*Emailsubscription, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_EmailsubscriptionMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithIDShop id_shop获取
func (obj *_EmailsubscriptionMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

// WithIDShopGroup id_shop_group获取
func (obj *_EmailsubscriptionMgr) WithIDShopGroup(idShopGroup uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop_group"] = idShopGroup })
}

// WithEmail email获取
func (obj *_EmailsubscriptionMgr) WithEmail(email string) Option {
	return optionFunc(func(o *options) { o.query["email"] = email })
}

// WithNewsletterDateAdd newsletter_date_add获取
func (obj *_EmailsubscriptionMgr) WithNewsletterDateAdd(newsletterDateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["newsletter_date_add"] = newsletterDateAdd })
}

// WithIPRegistrationNewsletter ip_registration_newsletter获取
func (obj *_EmailsubscriptionMgr) WithIPRegistrationNewsletter(ipRegistrationNewsletter string) Option {
	return optionFunc(func(o *options) { o.query["ip_registration_newsletter"] = ipRegistrationNewsletter })
}

// WithHTTPReferer http_referer获取
func (obj *_EmailsubscriptionMgr) WithHTTPReferer(httpReferer string) Option {
	return optionFunc(func(o *options) { o.query["http_referer"] = httpReferer })
}

// WithActive active获取
func (obj *_EmailsubscriptionMgr) WithActive(active bool) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}

// WithIDLang id_lang获取
func (obj *_EmailsubscriptionMgr) WithIDLang(idLang int) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

// GetByOption 功能选项模式获取
func (obj *_EmailsubscriptionMgr) GetByOption(opts ...Option) (result Emailsubscription, err error) {
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
func (obj *_EmailsubscriptionMgr) GetByOptions(opts ...Option) (results []*Emailsubscription, err error) {
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
func (obj *_EmailsubscriptionMgr) GetFromID(id int) (result Emailsubscription, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_EmailsubscriptionMgr) GetBatchFromID(ids []int) (results []*Emailsubscription, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromIDShop 通过id_shop获取内容
func (obj *_EmailsubscriptionMgr) GetFromIDShop(idShop uint32) (results []*Emailsubscription, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

// GetBatchFromIDShop 批量唯一主键查找
func (obj *_EmailsubscriptionMgr) GetBatchFromIDShop(idShops []uint32) (results []*Emailsubscription, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

// GetFromIDShopGroup 通过id_shop_group获取内容
func (obj *_EmailsubscriptionMgr) GetFromIDShopGroup(idShopGroup uint32) (results []*Emailsubscription, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_group = ?", idShopGroup).Find(&results).Error

	return
}

// GetBatchFromIDShopGroup 批量唯一主键查找
func (obj *_EmailsubscriptionMgr) GetBatchFromIDShopGroup(idShopGroups []uint32) (results []*Emailsubscription, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_group IN (?)", idShopGroups).Find(&results).Error

	return
}

// GetFromEmail 通过email获取内容
func (obj *_EmailsubscriptionMgr) GetFromEmail(email string) (results []*Emailsubscription, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("email = ?", email).Find(&results).Error

	return
}

// GetBatchFromEmail 批量唯一主键查找
func (obj *_EmailsubscriptionMgr) GetBatchFromEmail(emails []string) (results []*Emailsubscription, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("email IN (?)", emails).Find(&results).Error

	return
}

// GetFromNewsletterDateAdd 通过newsletter_date_add获取内容
func (obj *_EmailsubscriptionMgr) GetFromNewsletterDateAdd(newsletterDateAdd time.Time) (results []*Emailsubscription, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("newsletter_date_add = ?", newsletterDateAdd).Find(&results).Error

	return
}

// GetBatchFromNewsletterDateAdd 批量唯一主键查找
func (obj *_EmailsubscriptionMgr) GetBatchFromNewsletterDateAdd(newsletterDateAdds []time.Time) (results []*Emailsubscription, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("newsletter_date_add IN (?)", newsletterDateAdds).Find(&results).Error

	return
}

// GetFromIPRegistrationNewsletter 通过ip_registration_newsletter获取内容
func (obj *_EmailsubscriptionMgr) GetFromIPRegistrationNewsletter(ipRegistrationNewsletter string) (results []*Emailsubscription, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ip_registration_newsletter = ?", ipRegistrationNewsletter).Find(&results).Error

	return
}

// GetBatchFromIPRegistrationNewsletter 批量唯一主键查找
func (obj *_EmailsubscriptionMgr) GetBatchFromIPRegistrationNewsletter(ipRegistrationNewsletters []string) (results []*Emailsubscription, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ip_registration_newsletter IN (?)", ipRegistrationNewsletters).Find(&results).Error

	return
}

// GetFromHTTPReferer 通过http_referer获取内容
func (obj *_EmailsubscriptionMgr) GetFromHTTPReferer(httpReferer string) (results []*Emailsubscription, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("http_referer = ?", httpReferer).Find(&results).Error

	return
}

// GetBatchFromHTTPReferer 批量唯一主键查找
func (obj *_EmailsubscriptionMgr) GetBatchFromHTTPReferer(httpReferers []string) (results []*Emailsubscription, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("http_referer IN (?)", httpReferers).Find(&results).Error

	return
}

// GetFromActive 通过active获取内容
func (obj *_EmailsubscriptionMgr) GetFromActive(active bool) (results []*Emailsubscription, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active = ?", active).Find(&results).Error

	return
}

// GetBatchFromActive 批量唯一主键查找
func (obj *_EmailsubscriptionMgr) GetBatchFromActive(actives []bool) (results []*Emailsubscription, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active IN (?)", actives).Find(&results).Error

	return
}

// GetFromIDLang 通过id_lang获取内容
func (obj *_EmailsubscriptionMgr) GetFromIDLang(idLang int) (results []*Emailsubscription, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error

	return
}

// GetBatchFromIDLang 批量唯一主键查找
func (obj *_EmailsubscriptionMgr) GetBatchFromIDLang(idLangs []int) (results []*Emailsubscription, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_EmailsubscriptionMgr) FetchByPrimaryKey(id int) (result Emailsubscription, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}
