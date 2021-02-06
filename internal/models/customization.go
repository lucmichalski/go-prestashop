package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _CustomizationMgr struct {
	*_BaseMgr
}

// CustomizationMgr open func
func CustomizationMgr(db *gorm.DB) *_CustomizationMgr {
	if db == nil {
		panic(fmt.Errorf("CustomizationMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CustomizationMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_customization"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CustomizationMgr) GetTableName() string {
	return "ps_customization"
}

// Get 获取
func (obj *_CustomizationMgr) Get() (result Customization, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CustomizationMgr) Gets() (results []*Customization, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDCustomization id_customization获取
func (obj *_CustomizationMgr) WithIDCustomization(idCustomization uint32) Option {
	return optionFunc(func(o *options) { o.query["id_customization"] = idCustomization })
}

// WithIDProductAttribute id_product_attribute获取
func (obj *_CustomizationMgr) WithIDProductAttribute(idProductAttribute uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product_attribute"] = idProductAttribute })
}

// WithIDAddressDelivery id_address_delivery获取
func (obj *_CustomizationMgr) WithIDAddressDelivery(idAddressDelivery uint32) Option {
	return optionFunc(func(o *options) { o.query["id_address_delivery"] = idAddressDelivery })
}

// WithIDCart id_cart获取
func (obj *_CustomizationMgr) WithIDCart(idCart uint32) Option {
	return optionFunc(func(o *options) { o.query["id_cart"] = idCart })
}

// WithIDProduct id_product获取
func (obj *_CustomizationMgr) WithIDProduct(idProduct int) Option {
	return optionFunc(func(o *options) { o.query["id_product"] = idProduct })
}

// WithQuantity quantity获取
func (obj *_CustomizationMgr) WithQuantity(quantity int) Option {
	return optionFunc(func(o *options) { o.query["quantity"] = quantity })
}

// WithQuantityRefunded quantity_refunded获取
func (obj *_CustomizationMgr) WithQuantityRefunded(quantityRefunded int) Option {
	return optionFunc(func(o *options) { o.query["quantity_refunded"] = quantityRefunded })
}

// WithQuantityReturned quantity_returned获取
func (obj *_CustomizationMgr) WithQuantityReturned(quantityReturned int) Option {
	return optionFunc(func(o *options) { o.query["quantity_returned"] = quantityReturned })
}

// WithInCart in_cart获取
func (obj *_CustomizationMgr) WithInCart(inCart bool) Option {
	return optionFunc(func(o *options) { o.query["in_cart"] = inCart })
}

// GetByOption 功能选项模式获取
func (obj *_CustomizationMgr) GetByOption(opts ...Option) (result Customization, err error) {
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
func (obj *_CustomizationMgr) GetByOptions(opts ...Option) (results []*Customization, err error) {
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

// GetFromIDCustomization 通过id_customization获取内容
func (obj *_CustomizationMgr) GetFromIDCustomization(idCustomization uint32) (results []*Customization, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customization = ?", idCustomization).Find(&results).Error

	return
}

// GetBatchFromIDCustomization 批量唯一主键查找
func (obj *_CustomizationMgr) GetBatchFromIDCustomization(idCustomizations []uint32) (results []*Customization, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customization IN (?)", idCustomizations).Find(&results).Error

	return
}

// GetFromIDProductAttribute 通过id_product_attribute获取内容
func (obj *_CustomizationMgr) GetFromIDProductAttribute(idProductAttribute uint32) (results []*Customization, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_attribute = ?", idProductAttribute).Find(&results).Error

	return
}

// GetBatchFromIDProductAttribute 批量唯一主键查找
func (obj *_CustomizationMgr) GetBatchFromIDProductAttribute(idProductAttributes []uint32) (results []*Customization, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_attribute IN (?)", idProductAttributes).Find(&results).Error

	return
}

// GetFromIDAddressDelivery 通过id_address_delivery获取内容
func (obj *_CustomizationMgr) GetFromIDAddressDelivery(idAddressDelivery uint32) (results []*Customization, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_address_delivery = ?", idAddressDelivery).Find(&results).Error

	return
}

// GetBatchFromIDAddressDelivery 批量唯一主键查找
func (obj *_CustomizationMgr) GetBatchFromIDAddressDelivery(idAddressDeliverys []uint32) (results []*Customization, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_address_delivery IN (?)", idAddressDeliverys).Find(&results).Error

	return
}

// GetFromIDCart 通过id_cart获取内容
func (obj *_CustomizationMgr) GetFromIDCart(idCart uint32) (results []*Customization, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart = ?", idCart).Find(&results).Error

	return
}

// GetBatchFromIDCart 批量唯一主键查找
func (obj *_CustomizationMgr) GetBatchFromIDCart(idCarts []uint32) (results []*Customization, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart IN (?)", idCarts).Find(&results).Error

	return
}

// GetFromIDProduct 通过id_product获取内容
func (obj *_CustomizationMgr) GetFromIDProduct(idProduct int) (results []*Customization, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ?", idProduct).Find(&results).Error

	return
}

// GetBatchFromIDProduct 批量唯一主键查找
func (obj *_CustomizationMgr) GetBatchFromIDProduct(idProducts []int) (results []*Customization, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product IN (?)", idProducts).Find(&results).Error

	return
}

// GetFromQuantity 通过quantity获取内容
func (obj *_CustomizationMgr) GetFromQuantity(quantity int) (results []*Customization, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("quantity = ?", quantity).Find(&results).Error

	return
}

// GetBatchFromQuantity 批量唯一主键查找
func (obj *_CustomizationMgr) GetBatchFromQuantity(quantitys []int) (results []*Customization, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("quantity IN (?)", quantitys).Find(&results).Error

	return
}

// GetFromQuantityRefunded 通过quantity_refunded获取内容
func (obj *_CustomizationMgr) GetFromQuantityRefunded(quantityRefunded int) (results []*Customization, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("quantity_refunded = ?", quantityRefunded).Find(&results).Error

	return
}

// GetBatchFromQuantityRefunded 批量唯一主键查找
func (obj *_CustomizationMgr) GetBatchFromQuantityRefunded(quantityRefundeds []int) (results []*Customization, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("quantity_refunded IN (?)", quantityRefundeds).Find(&results).Error

	return
}

// GetFromQuantityReturned 通过quantity_returned获取内容
func (obj *_CustomizationMgr) GetFromQuantityReturned(quantityReturned int) (results []*Customization, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("quantity_returned = ?", quantityReturned).Find(&results).Error

	return
}

// GetBatchFromQuantityReturned 批量唯一主键查找
func (obj *_CustomizationMgr) GetBatchFromQuantityReturned(quantityReturneds []int) (results []*Customization, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("quantity_returned IN (?)", quantityReturneds).Find(&results).Error

	return
}

// GetFromInCart 通过in_cart获取内容
func (obj *_CustomizationMgr) GetFromInCart(inCart bool) (results []*Customization, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("in_cart = ?", inCart).Find(&results).Error

	return
}

// GetBatchFromInCart 批量唯一主键查找
func (obj *_CustomizationMgr) GetBatchFromInCart(inCarts []bool) (results []*Customization, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("in_cart IN (?)", inCarts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_CustomizationMgr) FetchByPrimaryKey(idCustomization uint32, idAddressDelivery uint32, idCart uint32, idProduct int) (result Customization, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customization = ? AND id_address_delivery = ? AND id_cart = ? AND id_product = ?", idCustomization, idAddressDelivery, idCart, idProduct).Find(&result).Error

	return
}

// FetchIndexByIDProductAttribute  获取多个内容
func (obj *_CustomizationMgr) FetchIndexByIDProductAttribute(idProductAttribute uint32) (results []*Customization, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_attribute = ?", idProductAttribute).Find(&results).Error

	return
}

// FetchIndexByIDCartProduct  获取多个内容
func (obj *_CustomizationMgr) FetchIndexByIDCartProduct(idProductAttribute uint32, idCart uint32, idProduct int) (results []*Customization, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_attribute = ? AND id_cart = ? AND id_product = ?", idProductAttribute, idCart, idProduct).Find(&results).Error

	return
}
