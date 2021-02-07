package models

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _EmailsubscriptionMgr struct {
	*_BaseMgr
}

func EmailsubscriptionMgr(db *gorm.DB) *_EmailsubscriptionMgr {
	if db == nil {
		panic(fmt.Errorf("EmailsubscriptionMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EmailsubscriptionMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_emailsubscription"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_EmailsubscriptionMgr) GetTableName() string {
	return "ps_emailsubscription"
}

func (obj *_EmailsubscriptionMgr) Get() (result Emailsubscription, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_EmailsubscriptionMgr) Gets() (results []*Emailsubscription, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

func (obj *_EmailsubscriptionMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

func (obj *_EmailsubscriptionMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

func (obj *_EmailsubscriptionMgr) WithIDShopGroup(idShopGroup uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop_group"] = idShopGroup })
}

func (obj *_EmailsubscriptionMgr) WithEmail(email string) Option {
	return optionFunc(func(o *options) { o.query["email"] = email })
}

func (obj *_EmailsubscriptionMgr) WithNewsletterDateAdd(newsletterDateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["newsletter_date_add"] = newsletterDateAdd })
}

func (obj *_EmailsubscriptionMgr) WithIPRegistrationNewsletter(ipRegistrationNewsletter string) Option {
	return optionFunc(func(o *options) { o.query["ip_registration_newsletter"] = ipRegistrationNewsletter })
}

func (obj *_EmailsubscriptionMgr) WithHTTPReferer(httpReferer string) Option {
	return optionFunc(func(o *options) { o.query["http_referer"] = httpReferer })
}

func (obj *_EmailsubscriptionMgr) WithActive(active bool) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}

func (obj *_EmailsubscriptionMgr) WithIDLang(idLang int) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

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

func (obj *_EmailsubscriptionMgr) GetFromID(id int) (result Emailsubscription, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

func (obj *_EmailsubscriptionMgr) GetBatchFromID(ids []int) (results []*Emailsubscription, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

func (obj *_EmailsubscriptionMgr) GetFromIDShop(idShop uint32) (results []*Emailsubscription, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

func (obj *_EmailsubscriptionMgr) GetBatchFromIDShop(idShops []uint32) (results []*Emailsubscription, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

func (obj *_EmailsubscriptionMgr) GetFromIDShopGroup(idShopGroup uint32) (results []*Emailsubscription, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_group = ?", idShopGroup).Find(&results).Error

	return
}

func (obj *_EmailsubscriptionMgr) GetBatchFromIDShopGroup(idShopGroups []uint32) (results []*Emailsubscription, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_group IN (?)", idShopGroups).Find(&results).Error

	return
}

func (obj *_EmailsubscriptionMgr) GetFromEmail(email string) (results []*Emailsubscription, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("email = ?", email).Find(&results).Error

	return
}

func (obj *_EmailsubscriptionMgr) GetBatchFromEmail(emails []string) (results []*Emailsubscription, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("email IN (?)", emails).Find(&results).Error

	return
}

func (obj *_EmailsubscriptionMgr) GetFromNewsletterDateAdd(newsletterDateAdd time.Time) (results []*Emailsubscription, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("newsletter_date_add = ?", newsletterDateAdd).Find(&results).Error

	return
}

func (obj *_EmailsubscriptionMgr) GetBatchFromNewsletterDateAdd(newsletterDateAdds []time.Time) (results []*Emailsubscription, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("newsletter_date_add IN (?)", newsletterDateAdds).Find(&results).Error

	return
}

func (obj *_EmailsubscriptionMgr) GetFromIPRegistrationNewsletter(ipRegistrationNewsletter string) (results []*Emailsubscription, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ip_registration_newsletter = ?", ipRegistrationNewsletter).Find(&results).Error

	return
}

func (obj *_EmailsubscriptionMgr) GetBatchFromIPRegistrationNewsletter(ipRegistrationNewsletters []string) (results []*Emailsubscription, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ip_registration_newsletter IN (?)", ipRegistrationNewsletters).Find(&results).Error

	return
}

func (obj *_EmailsubscriptionMgr) GetFromHTTPReferer(httpReferer string) (results []*Emailsubscription, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("http_referer = ?", httpReferer).Find(&results).Error

	return
}

func (obj *_EmailsubscriptionMgr) GetBatchFromHTTPReferer(httpReferers []string) (results []*Emailsubscription, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("http_referer IN (?)", httpReferers).Find(&results).Error

	return
}

func (obj *_EmailsubscriptionMgr) GetFromActive(active bool) (results []*Emailsubscription, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active = ?", active).Find(&results).Error

	return
}

func (obj *_EmailsubscriptionMgr) GetBatchFromActive(actives []bool) (results []*Emailsubscription, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active IN (?)", actives).Find(&results).Error

	return
}

func (obj *_EmailsubscriptionMgr) GetFromIDLang(idLang int) (results []*Emailsubscription, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error

	return
}

func (obj *_EmailsubscriptionMgr) GetBatchFromIDLang(idLangs []int) (results []*Emailsubscription, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error

	return
}

func (obj *_EmailsubscriptionMgr) FetchByPrimaryKey(id int) (result Emailsubscription, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}
