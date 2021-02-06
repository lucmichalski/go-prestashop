package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _CartProductMgr struct {
	*_BaseMgr
}

// CartProductMgr open func
func CartProductMgr(db *gorm.DB) *_CartProductMgr {
	if db == nil {
		panic(fmt.Errorf("CartProductMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CartProductMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_cart_product"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CartProductMgr) GetTableName() string {
	return "eg_cart_product"
}

// Get 获取
func (obj *_CartProductMgr) Get() (result CartProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CartProductMgr) Gets() (results []*CartProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDCart id_cart获取
func (obj *_CartProductMgr) WithIDCart(idCart uint32) Option {
	return optionFunc(func(o *options) { o.query["id_cart"] = idCart })
}

// WithIDProduct id_product获取
func (obj *_CartProductMgr) WithIDProduct(idProduct uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product"] = idProduct })
}

// WithIDAddressDelivery id_address_delivery获取
func (obj *_CartProductMgr) WithIDAddressDelivery(idAddressDelivery uint32) Option {
	return optionFunc(func(o *options) { o.query["id_address_delivery"] = idAddressDelivery })
}

// WithIDShop id_shop获取
func (obj *_CartProductMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

// WithIDProductAttribute id_product_attribute获取
func (obj *_CartProductMgr) WithIDProductAttribute(idProductAttribute uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product_attribute"] = idProductAttribute })
}

// WithIDCustomization id_customization获取
func (obj *_CartProductMgr) WithIDCustomization(idCustomization uint32) Option {
	return optionFunc(func(o *options) { o.query["id_customization"] = idCustomization })
}

// WithQuantity quantity获取
func (obj *_CartProductMgr) WithQuantity(quantity uint32) Option {
	return optionFunc(func(o *options) { o.query["quantity"] = quantity })
}

// WithDateAdd date_add获取
func (obj *_CartProductMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

// GetByOption 功能选项模式获取
func (obj *_CartProductMgr) GetByOption(opts ...Option) (result CartProduct, err error) {
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
func (obj *_CartProductMgr) GetByOptions(opts ...Option) (results []*CartProduct, err error) {
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

// GetFromIDCart 通过id_cart获取内容
func (obj *_CartProductMgr) GetFromIDCart(idCart uint32) (results []*CartProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart = ?", idCart).Find(&results).Error

	return
}

// GetBatchFromIDCart 批量唯一主键查找
func (obj *_CartProductMgr) GetBatchFromIDCart(idCarts []uint32) (results []*CartProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart IN (?)", idCarts).Find(&results).Error

	return
}

// GetFromIDProduct 通过id_product获取内容
func (obj *_CartProductMgr) GetFromIDProduct(idProduct uint32) (results []*CartProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ?", idProduct).Find(&results).Error

	return
}

// GetBatchFromIDProduct 批量唯一主键查找
func (obj *_CartProductMgr) GetBatchFromIDProduct(idProducts []uint32) (results []*CartProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product IN (?)", idProducts).Find(&results).Error

	return
}

// GetFromIDAddressDelivery 通过id_address_delivery获取内容
func (obj *_CartProductMgr) GetFromIDAddressDelivery(idAddressDelivery uint32) (results []*CartProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_address_delivery = ?", idAddressDelivery).Find(&results).Error

	return
}

// GetBatchFromIDAddressDelivery 批量唯一主键查找
func (obj *_CartProductMgr) GetBatchFromIDAddressDelivery(idAddressDeliverys []uint32) (results []*CartProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_address_delivery IN (?)", idAddressDeliverys).Find(&results).Error

	return
}

// GetFromIDShop 通过id_shop获取内容
func (obj *_CartProductMgr) GetFromIDShop(idShop uint32) (results []*CartProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

// GetBatchFromIDShop 批量唯一主键查找
func (obj *_CartProductMgr) GetBatchFromIDShop(idShops []uint32) (results []*CartProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

// GetFromIDProductAttribute 通过id_product_attribute获取内容
func (obj *_CartProductMgr) GetFromIDProductAttribute(idProductAttribute uint32) (results []*CartProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_attribute = ?", idProductAttribute).Find(&results).Error

	return
}

// GetBatchFromIDProductAttribute 批量唯一主键查找
func (obj *_CartProductMgr) GetBatchFromIDProductAttribute(idProductAttributes []uint32) (results []*CartProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_attribute IN (?)", idProductAttributes).Find(&results).Error

	return
}

// GetFromIDCustomization 通过id_customization获取内容
func (obj *_CartProductMgr) GetFromIDCustomization(idCustomization uint32) (results []*CartProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customization = ?", idCustomization).Find(&results).Error

	return
}

// GetBatchFromIDCustomization 批量唯一主键查找
func (obj *_CartProductMgr) GetBatchFromIDCustomization(idCustomizations []uint32) (results []*CartProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customization IN (?)", idCustomizations).Find(&results).Error

	return
}

// GetFromQuantity 通过quantity获取内容
func (obj *_CartProductMgr) GetFromQuantity(quantity uint32) (results []*CartProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("quantity = ?", quantity).Find(&results).Error

	return
}

// GetBatchFromQuantity 批量唯一主键查找
func (obj *_CartProductMgr) GetBatchFromQuantity(quantitys []uint32) (results []*CartProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("quantity IN (?)", quantitys).Find(&results).Error

	return
}

// GetFromDateAdd 通过date_add获取内容
func (obj *_CartProductMgr) GetFromDateAdd(dateAdd time.Time) (results []*CartProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error

	return
}

// GetBatchFromDateAdd 批量唯一主键查找
func (obj *_CartProductMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*CartProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_CartProductMgr) FetchByPrimaryKey(idCart uint32, idProduct uint32, idAddressDelivery uint32, idProductAttribute uint32, idCustomization uint32) (result CartProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart = ? AND id_product = ? AND id_address_delivery = ? AND id_product_attribute = ? AND id_customization = ?", idCart, idProduct, idAddressDelivery, idProductAttribute, idCustomization).Find(&result).Error

	return
}

// FetchIndexByIDCartOrder  获取多个内容
func (obj *_CartProductMgr) FetchIndexByIDCartOrder(idCart uint32, idProduct uint32, idProductAttribute uint32, dateAdd time.Time) (results []*CartProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart = ? AND id_product = ? AND id_product_attribute = ? AND date_add = ?", idCart, idProduct, idProductAttribute, dateAdd).Find(&results).Error

	return
}

// FetchIndexByIDProductAttribute  获取多个内容
func (obj *_CartProductMgr) FetchIndexByIDProductAttribute(idProductAttribute uint32) (results []*CartProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_attribute = ?", idProductAttribute).Find(&results).Error

	return
}
