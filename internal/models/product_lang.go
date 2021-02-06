package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _ProductLangMgr struct {
	*_BaseMgr
}

// ProductLangMgr open func
func ProductLangMgr(db *gorm.DB) *_ProductLangMgr {
	if db == nil {
		panic(fmt.Errorf("ProductLangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ProductLangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_product_lang"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ProductLangMgr) GetTableName() string {
	return "ps_product_lang"
}

// Get 获取
func (obj *_ProductLangMgr) Get() (result ProductLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ProductLangMgr) Gets() (results []*ProductLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDProduct id_product获取
func (obj *_ProductLangMgr) WithIDProduct(idProduct uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product"] = idProduct })
}

// WithIDShop id_shop获取
func (obj *_ProductLangMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

// WithIDLang id_lang获取
func (obj *_ProductLangMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

// WithDescription description获取
func (obj *_ProductLangMgr) WithDescription(description string) Option {
	return optionFunc(func(o *options) { o.query["description"] = description })
}

// WithDescriptionShort description_short获取
func (obj *_ProductLangMgr) WithDescriptionShort(descriptionShort string) Option {
	return optionFunc(func(o *options) { o.query["description_short"] = descriptionShort })
}

// WithLinkRewrite link_rewrite获取
func (obj *_ProductLangMgr) WithLinkRewrite(linkRewrite string) Option {
	return optionFunc(func(o *options) { o.query["link_rewrite"] = linkRewrite })
}

// WithMetaDescription meta_description获取
func (obj *_ProductLangMgr) WithMetaDescription(metaDescription string) Option {
	return optionFunc(func(o *options) { o.query["meta_description"] = metaDescription })
}

// WithMetaKeywords meta_keywords获取
func (obj *_ProductLangMgr) WithMetaKeywords(metaKeywords string) Option {
	return optionFunc(func(o *options) { o.query["meta_keywords"] = metaKeywords })
}

// WithMetaTitle meta_title获取
func (obj *_ProductLangMgr) WithMetaTitle(metaTitle string) Option {
	return optionFunc(func(o *options) { o.query["meta_title"] = metaTitle })
}

// WithName name获取
func (obj *_ProductLangMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithAvailableNow available_now获取
func (obj *_ProductLangMgr) WithAvailableNow(availableNow string) Option {
	return optionFunc(func(o *options) { o.query["available_now"] = availableNow })
}

// WithAvailableLater available_later获取
func (obj *_ProductLangMgr) WithAvailableLater(availableLater string) Option {
	return optionFunc(func(o *options) { o.query["available_later"] = availableLater })
}

// WithDeliveryInStock delivery_in_stock获取
func (obj *_ProductLangMgr) WithDeliveryInStock(deliveryInStock string) Option {
	return optionFunc(func(o *options) { o.query["delivery_in_stock"] = deliveryInStock })
}

// WithDeliveryOutStock delivery_out_stock获取
func (obj *_ProductLangMgr) WithDeliveryOutStock(deliveryOutStock string) Option {
	return optionFunc(func(o *options) { o.query["delivery_out_stock"] = deliveryOutStock })
}

// GetByOption 功能选项模式获取
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

// GetByOptions 批量功能选项模式获取
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

//////////////////////////enume case ////////////////////////////////////////////

// GetFromIDProduct 通过id_product获取内容
func (obj *_ProductLangMgr) GetFromIDProduct(idProduct uint32) (results []*ProductLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ?", idProduct).Find(&results).Error

	return
}

// GetBatchFromIDProduct 批量唯一主键查找
func (obj *_ProductLangMgr) GetBatchFromIDProduct(idProducts []uint32) (results []*ProductLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product IN (?)", idProducts).Find(&results).Error

	return
}

// GetFromIDShop 通过id_shop获取内容
func (obj *_ProductLangMgr) GetFromIDShop(idShop uint32) (results []*ProductLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

// GetBatchFromIDShop 批量唯一主键查找
func (obj *_ProductLangMgr) GetBatchFromIDShop(idShops []uint32) (results []*ProductLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

// GetFromIDLang 通过id_lang获取内容
func (obj *_ProductLangMgr) GetFromIDLang(idLang uint32) (results []*ProductLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error

	return
}

// GetBatchFromIDLang 批量唯一主键查找
func (obj *_ProductLangMgr) GetBatchFromIDLang(idLangs []uint32) (results []*ProductLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error

	return
}

// GetFromDescription 通过description获取内容
func (obj *_ProductLangMgr) GetFromDescription(description string) (results []*ProductLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("description = ?", description).Find(&results).Error

	return
}

// GetBatchFromDescription 批量唯一主键查找
func (obj *_ProductLangMgr) GetBatchFromDescription(descriptions []string) (results []*ProductLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("description IN (?)", descriptions).Find(&results).Error

	return
}

// GetFromDescriptionShort 通过description_short获取内容
func (obj *_ProductLangMgr) GetFromDescriptionShort(descriptionShort string) (results []*ProductLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("description_short = ?", descriptionShort).Find(&results).Error

	return
}

// GetBatchFromDescriptionShort 批量唯一主键查找
func (obj *_ProductLangMgr) GetBatchFromDescriptionShort(descriptionShorts []string) (results []*ProductLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("description_short IN (?)", descriptionShorts).Find(&results).Error

	return
}

// GetFromLinkRewrite 通过link_rewrite获取内容
func (obj *_ProductLangMgr) GetFromLinkRewrite(linkRewrite string) (results []*ProductLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("link_rewrite = ?", linkRewrite).Find(&results).Error

	return
}

// GetBatchFromLinkRewrite 批量唯一主键查找
func (obj *_ProductLangMgr) GetBatchFromLinkRewrite(linkRewrites []string) (results []*ProductLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("link_rewrite IN (?)", linkRewrites).Find(&results).Error

	return
}

// GetFromMetaDescription 通过meta_description获取内容
func (obj *_ProductLangMgr) GetFromMetaDescription(metaDescription string) (results []*ProductLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("meta_description = ?", metaDescription).Find(&results).Error

	return
}

// GetBatchFromMetaDescription 批量唯一主键查找
func (obj *_ProductLangMgr) GetBatchFromMetaDescription(metaDescriptions []string) (results []*ProductLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("meta_description IN (?)", metaDescriptions).Find(&results).Error

	return
}

// GetFromMetaKeywords 通过meta_keywords获取内容
func (obj *_ProductLangMgr) GetFromMetaKeywords(metaKeywords string) (results []*ProductLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("meta_keywords = ?", metaKeywords).Find(&results).Error

	return
}

// GetBatchFromMetaKeywords 批量唯一主键查找
func (obj *_ProductLangMgr) GetBatchFromMetaKeywords(metaKeywordss []string) (results []*ProductLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("meta_keywords IN (?)", metaKeywordss).Find(&results).Error

	return
}

// GetFromMetaTitle 通过meta_title获取内容
func (obj *_ProductLangMgr) GetFromMetaTitle(metaTitle string) (results []*ProductLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("meta_title = ?", metaTitle).Find(&results).Error

	return
}

// GetBatchFromMetaTitle 批量唯一主键查找
func (obj *_ProductLangMgr) GetBatchFromMetaTitle(metaTitles []string) (results []*ProductLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("meta_title IN (?)", metaTitles).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_ProductLangMgr) GetFromName(name string) (results []*ProductLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量唯一主键查找
func (obj *_ProductLangMgr) GetBatchFromName(names []string) (results []*ProductLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

// GetFromAvailableNow 通过available_now获取内容
func (obj *_ProductLangMgr) GetFromAvailableNow(availableNow string) (results []*ProductLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("available_now = ?", availableNow).Find(&results).Error

	return
}

// GetBatchFromAvailableNow 批量唯一主键查找
func (obj *_ProductLangMgr) GetBatchFromAvailableNow(availableNows []string) (results []*ProductLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("available_now IN (?)", availableNows).Find(&results).Error

	return
}

// GetFromAvailableLater 通过available_later获取内容
func (obj *_ProductLangMgr) GetFromAvailableLater(availableLater string) (results []*ProductLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("available_later = ?", availableLater).Find(&results).Error

	return
}

// GetBatchFromAvailableLater 批量唯一主键查找
func (obj *_ProductLangMgr) GetBatchFromAvailableLater(availableLaters []string) (results []*ProductLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("available_later IN (?)", availableLaters).Find(&results).Error

	return
}

// GetFromDeliveryInStock 通过delivery_in_stock获取内容
func (obj *_ProductLangMgr) GetFromDeliveryInStock(deliveryInStock string) (results []*ProductLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("delivery_in_stock = ?", deliveryInStock).Find(&results).Error

	return
}

// GetBatchFromDeliveryInStock 批量唯一主键查找
func (obj *_ProductLangMgr) GetBatchFromDeliveryInStock(deliveryInStocks []string) (results []*ProductLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("delivery_in_stock IN (?)", deliveryInStocks).Find(&results).Error

	return
}

// GetFromDeliveryOutStock 通过delivery_out_stock获取内容
func (obj *_ProductLangMgr) GetFromDeliveryOutStock(deliveryOutStock string) (results []*ProductLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("delivery_out_stock = ?", deliveryOutStock).Find(&results).Error

	return
}

// GetBatchFromDeliveryOutStock 批量唯一主键查找
func (obj *_ProductLangMgr) GetBatchFromDeliveryOutStock(deliveryOutStocks []string) (results []*ProductLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("delivery_out_stock IN (?)", deliveryOutStocks).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_ProductLangMgr) FetchByPrimaryKey(idProduct uint32, idShop uint32, idLang uint32) (result ProductLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ? AND id_shop = ? AND id_lang = ?", idProduct, idShop, idLang).Find(&result).Error

	return
}

// FetchIndexByIDLang  获取多个内容
func (obj *_ProductLangMgr) FetchIndexByIDLang(idLang uint32) (results []*ProductLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error

	return
}

// FetchIndexByName  获取多个内容
func (obj *_ProductLangMgr) FetchIndexByName(name string) (results []*ProductLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}
