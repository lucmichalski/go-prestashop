package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _ShopURLMgr struct {
	*_BaseMgr
}

func ShopURLMgr(db *gorm.DB) *_ShopURLMgr {
	if db == nil {
		panic(fmt.Errorf("ShopURLMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ShopURLMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_shop_url"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_ShopURLMgr) GetTableName() string {
	return "ps_shop_url"
}

func (obj *_ShopURLMgr) Get() (result ShopURL, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_ShopURLMgr) Gets() (results []*ShopURL, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_ShopURLMgr) WithIDShopURL(idShopURL uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop_url"] = idShopURL })
}

func (obj *_ShopURLMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

func (obj *_ShopURLMgr) WithDomain(domain string) Option {
	return optionFunc(func(o *options) { o.query["domain"] = domain })
}

func (obj *_ShopURLMgr) WithDomainSsl(domainSsl string) Option {
	return optionFunc(func(o *options) { o.query["domain_ssl"] = domainSsl })
}

func (obj *_ShopURLMgr) WithPhysicalURI(physicalURI string) Option {
	return optionFunc(func(o *options) { o.query["physical_uri"] = physicalURI })
}

func (obj *_ShopURLMgr) WithVirtualURI(virtualURI string) Option {
	return optionFunc(func(o *options) { o.query["virtual_uri"] = virtualURI })
}

func (obj *_ShopURLMgr) WithMain(main bool) Option {
	return optionFunc(func(o *options) { o.query["main"] = main })
}

func (obj *_ShopURLMgr) WithActive(active bool) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}

func (obj *_ShopURLMgr) GetByOption(opts ...Option) (result ShopURL, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_ShopURLMgr) GetByOptions(opts ...Option) (results []*ShopURL, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_ShopURLMgr) GetFromIDShopURL(idShopURL uint32) (result ShopURL, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_url = ?", idShopURL).Find(&result).Error

	return
}

func (obj *_ShopURLMgr) GetBatchFromIDShopURL(idShopURLs []uint32) (results []*ShopURL, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_url IN (?)", idShopURLs).Find(&results).Error

	return
}

func (obj *_ShopURLMgr) GetFromIDShop(idShop uint32) (results []*ShopURL, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

func (obj *_ShopURLMgr) GetBatchFromIDShop(idShops []uint32) (results []*ShopURL, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

func (obj *_ShopURLMgr) GetFromDomain(domain string) (results []*ShopURL, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("domain = ?", domain).Find(&results).Error

	return
}

func (obj *_ShopURLMgr) GetBatchFromDomain(domains []string) (results []*ShopURL, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("domain IN (?)", domains).Find(&results).Error

	return
}

func (obj *_ShopURLMgr) GetFromDomainSsl(domainSsl string) (results []*ShopURL, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("domain_ssl = ?", domainSsl).Find(&results).Error

	return
}

func (obj *_ShopURLMgr) GetBatchFromDomainSsl(domainSsls []string) (results []*ShopURL, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("domain_ssl IN (?)", domainSsls).Find(&results).Error

	return
}

func (obj *_ShopURLMgr) GetFromPhysicalURI(physicalURI string) (results []*ShopURL, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("physical_uri = ?", physicalURI).Find(&results).Error

	return
}

func (obj *_ShopURLMgr) GetBatchFromPhysicalURI(physicalURIs []string) (results []*ShopURL, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("physical_uri IN (?)", physicalURIs).Find(&results).Error

	return
}

func (obj *_ShopURLMgr) GetFromVirtualURI(virtualURI string) (results []*ShopURL, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("virtual_uri = ?", virtualURI).Find(&results).Error

	return
}

func (obj *_ShopURLMgr) GetBatchFromVirtualURI(virtualURIs []string) (results []*ShopURL, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("virtual_uri IN (?)", virtualURIs).Find(&results).Error

	return
}

func (obj *_ShopURLMgr) GetFromMain(main bool) (results []*ShopURL, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("main = ?", main).Find(&results).Error

	return
}

func (obj *_ShopURLMgr) GetBatchFromMain(mains []bool) (results []*ShopURL, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("main IN (?)", mains).Find(&results).Error

	return
}

func (obj *_ShopURLMgr) GetFromActive(active bool) (results []*ShopURL, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active = ?", active).Find(&results).Error

	return
}

func (obj *_ShopURLMgr) GetBatchFromActive(actives []bool) (results []*ShopURL, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active IN (?)", actives).Find(&results).Error

	return
}


func (obj *_ShopURLMgr) FetchByPrimaryKey(idShopURL uint32) (result ShopURL, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_url = ?", idShopURL).Find(&result).Error

	return
}

func (obj *_ShopURLMgr) FetchUniqueIndexByFullShopURL(domain string, physicalURI string, virtualURI string) (result ShopURL, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("domain = ? AND physical_uri = ? AND virtual_uri = ?", domain, physicalURI, virtualURI).Find(&result).Error

	return
}

func (obj *_ShopURLMgr) FetchUniqueIndexByFullShopURLSsl(domainSsl string, physicalURI string, virtualURI string) (result ShopURL, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("domain_ssl = ? AND physical_uri = ? AND virtual_uri = ?", domainSsl, physicalURI, virtualURI).Find(&result).Error

	return
}

func (obj *_ShopURLMgr) FetchIndexByIDShop(idShop uint32, main bool) (results []*ShopURL, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ? AND main = ?", idShop, main).Find(&results).Error

	return
}
