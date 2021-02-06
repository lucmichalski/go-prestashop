package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _ProductLangMgr struct {
	*_BaseMgr
}

func ProductLangMgr(db *gorm.DB) *_ProductLangMgr {
	if db == nil {
		panic(fmt.Errorf("ProductLangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ProductLangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_product_lang"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_ProductLangMgr) GetTableName() string {
	return "ps_product_lang"
}

func (obj *_ProductLangMgr) Get() (result ProductLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_ProductLangMgr) Gets() (results []*ProductLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_ProductLangMgr) WithIDProduct(idProduct uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product"] = idProduct })
}

func (obj *_ProductLangMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

func (obj *_ProductLangMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

func (obj *_ProductLangMgr) WithDescription(description string) Option {
	return optionFunc(func(o *options) { o.query["description"] = description })
}

func (obj *_ProductLangMgr) WithDescriptionShort(descriptionShort string) Option {
	return optionFunc(func(o *options) { o.query["description_short"] = descriptionShort })
}

func (obj *_ProductLangMgr) WithLinkRewrite(linkRewrite string) Option {
	return optionFunc(func(o *options) { o.query["link_rewrite"] = linkRewrite })
}

func (obj *_ProductLangMgr) WithMetaDescription(metaDescription string) Option {
	return optionFunc(func(o *options) { o.query["meta_description"] = metaDescription })
}

func (obj *_ProductLangMgr) WithMetaKeywords(metaKeywords string) Option {
	return optionFunc(func(o *options) { o.query["meta_keywords"] = metaKeywords })
}

func (obj *_ProductLangMgr) WithMetaTitle(metaTitle string) Option {
	return optionFunc(func(o *options) { o.query["meta_title"] = metaTitle })
}

func (obj *_ProductLangMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

func (obj *_ProductLangMgr) WithAvailableNow(availableNow string) Option {
	return optionFunc(func(o *options) { o.query["available_now"] = availableNow })
}

func (obj *_ProductLangMgr) WithAvailableLater(availableLater string) Option {
	return optionFunc(func(o *options) { o.query["available_later"] = availableLater })
}

func (obj *_ProductLangMgr) WithDeliveryInStock(deliveryInStock string) Option {
	return optionFunc(func(o *options) { o.query["delivery_in_stock"] = deliveryInStock })
}

func (obj *_ProductLangMgr) WithDeliveryOutStock(deliveryOutStock string) Option {
	return optionFunc(func(o *options) { o.query["delivery_out_stock"] = deliveryOutStock })
}

func (obj *_ProductLangMgr) GetByOption(opts ...Option) (result ProductLang, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_ProductLangMgr) GetByOptions(opts ...Option) (results []*ProductLang, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_ProductLangMgr) GetFromIDProduct(idProduct uint32) (results []*ProductLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ?", idProduct).Find(&results).Error

	return
}

func (obj *_ProductLangMgr) GetBatchFromIDProduct(idProducts []uint32) (results []*ProductLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product IN (?)", idProducts).Find(&results).Error

	return
}

func (obj *_ProductLangMgr) GetFromIDShop(idShop uint32) (results []*ProductLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

func (obj *_ProductLangMgr) GetBatchFromIDShop(idShops []uint32) (results []*ProductLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

func (obj *_ProductLangMgr) GetFromIDLang(idLang uint32) (results []*ProductLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error

	return
}

func (obj *_ProductLangMgr) GetBatchFromIDLang(idLangs []uint32) (results []*ProductLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error

	return
}

func (obj *_ProductLangMgr) GetFromDescription(description string) (results []*ProductLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("description = ?", description).Find(&results).Error

	return
}

func (obj *_ProductLangMgr) GetBatchFromDescription(descriptions []string) (results []*ProductLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("description IN (?)", descriptions).Find(&results).Error

	return
}

func (obj *_ProductLangMgr) GetFromDescriptionShort(descriptionShort string) (results []*ProductLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("description_short = ?", descriptionShort).Find(&results).Error

	return
}

func (obj *_ProductLangMgr) GetBatchFromDescriptionShort(descriptionShorts []string) (results []*ProductLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("description_short IN (?)", descriptionShorts).Find(&results).Error

	return
}

func (obj *_ProductLangMgr) GetFromLinkRewrite(linkRewrite string) (results []*ProductLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("link_rewrite = ?", linkRewrite).Find(&results).Error

	return
}

func (obj *_ProductLangMgr) GetBatchFromLinkRewrite(linkRewrites []string) (results []*ProductLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("link_rewrite IN (?)", linkRewrites).Find(&results).Error

	return
}

func (obj *_ProductLangMgr) GetFromMetaDescription(metaDescription string) (results []*ProductLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("meta_description = ?", metaDescription).Find(&results).Error

	return
}

func (obj *_ProductLangMgr) GetBatchFromMetaDescription(metaDescriptions []string) (results []*ProductLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("meta_description IN (?)", metaDescriptions).Find(&results).Error

	return
}

func (obj *_ProductLangMgr) GetFromMetaKeywords(metaKeywords string) (results []*ProductLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("meta_keywords = ?", metaKeywords).Find(&results).Error

	return
}

func (obj *_ProductLangMgr) GetBatchFromMetaKeywords(metaKeywordss []string) (results []*ProductLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("meta_keywords IN (?)", metaKeywordss).Find(&results).Error

	return
}

func (obj *_ProductLangMgr) GetFromMetaTitle(metaTitle string) (results []*ProductLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("meta_title = ?", metaTitle).Find(&results).Error

	return
}

func (obj *_ProductLangMgr) GetBatchFromMetaTitle(metaTitles []string) (results []*ProductLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("meta_title IN (?)", metaTitles).Find(&results).Error

	return
}

func (obj *_ProductLangMgr) GetFromName(name string) (results []*ProductLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

func (obj *_ProductLangMgr) GetBatchFromName(names []string) (results []*ProductLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

func (obj *_ProductLangMgr) GetFromAvailableNow(availableNow string) (results []*ProductLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("available_now = ?", availableNow).Find(&results).Error

	return
}

func (obj *_ProductLangMgr) GetBatchFromAvailableNow(availableNows []string) (results []*ProductLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("available_now IN (?)", availableNows).Find(&results).Error

	return
}

func (obj *_ProductLangMgr) GetFromAvailableLater(availableLater string) (results []*ProductLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("available_later = ?", availableLater).Find(&results).Error

	return
}

func (obj *_ProductLangMgr) GetBatchFromAvailableLater(availableLaters []string) (results []*ProductLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("available_later IN (?)", availableLaters).Find(&results).Error

	return
}

func (obj *_ProductLangMgr) GetFromDeliveryInStock(deliveryInStock string) (results []*ProductLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("delivery_in_stock = ?", deliveryInStock).Find(&results).Error

	return
}

func (obj *_ProductLangMgr) GetBatchFromDeliveryInStock(deliveryInStocks []string) (results []*ProductLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("delivery_in_stock IN (?)", deliveryInStocks).Find(&results).Error

	return
}

func (obj *_ProductLangMgr) GetFromDeliveryOutStock(deliveryOutStock string) (results []*ProductLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("delivery_out_stock = ?", deliveryOutStock).Find(&results).Error

	return
}

func (obj *_ProductLangMgr) GetBatchFromDeliveryOutStock(deliveryOutStocks []string) (results []*ProductLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("delivery_out_stock IN (?)", deliveryOutStocks).Find(&results).Error

	return
}


func (obj *_ProductLangMgr) FetchByPrimaryKey(idProduct uint32, idShop uint32, idLang uint32) (result ProductLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ? AND id_shop = ? AND id_lang = ?", idProduct, idShop, idLang).Find(&result).Error

	return
}

func (obj *_ProductLangMgr) FetchIndexByIDLang(idLang uint32) (results []*ProductLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error

	return
}

func (obj *_ProductLangMgr) FetchIndexByName(name string) (results []*ProductLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}
