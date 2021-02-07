package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _CustomizationMgr struct {
	*_BaseMgr
}

func CustomizationMgr(db *gorm.DB) *_CustomizationMgr {
	if db == nil {
		panic(fmt.Errorf("CustomizationMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CustomizationMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_customization"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_CustomizationMgr) GetTableName() string {
	return "ps_customization"
}

func (obj *_CustomizationMgr) Get() (result Customization, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_CustomizationMgr) Gets() (results []*Customization, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

func (obj *_CustomizationMgr) WithIDCustomization(idCustomization uint32) Option {
	return optionFunc(func(o *options) { o.query["id_customization"] = idCustomization })
}

func (obj *_CustomizationMgr) WithIDProductAttribute(idProductAttribute uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product_attribute"] = idProductAttribute })
}

func (obj *_CustomizationMgr) WithIDAddressDelivery(idAddressDelivery uint32) Option {
	return optionFunc(func(o *options) { o.query["id_address_delivery"] = idAddressDelivery })
}

func (obj *_CustomizationMgr) WithIDCart(idCart uint32) Option {
	return optionFunc(func(o *options) { o.query["id_cart"] = idCart })
}

func (obj *_CustomizationMgr) WithIDProduct(idProduct int) Option {
	return optionFunc(func(o *options) { o.query["id_product"] = idProduct })
}

func (obj *_CustomizationMgr) WithQuantity(quantity int) Option {
	return optionFunc(func(o *options) { o.query["quantity"] = quantity })
}

func (obj *_CustomizationMgr) WithQuantityRefunded(quantityRefunded int) Option {
	return optionFunc(func(o *options) { o.query["quantity_refunded"] = quantityRefunded })
}

func (obj *_CustomizationMgr) WithQuantityReturned(quantityReturned int) Option {
	return optionFunc(func(o *options) { o.query["quantity_returned"] = quantityReturned })
}

func (obj *_CustomizationMgr) WithInCart(inCart bool) Option {
	return optionFunc(func(o *options) { o.query["in_cart"] = inCart })
}

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

func (obj *_CustomizationMgr) GetFromIDCustomization(idCustomization uint32) (results []*Customization, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customization = ?", idCustomization).Find(&results).Error

	return
}

func (obj *_CustomizationMgr) GetBatchFromIDCustomization(idCustomizations []uint32) (results []*Customization, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customization IN (?)", idCustomizations).Find(&results).Error

	return
}

func (obj *_CustomizationMgr) GetFromIDProductAttribute(idProductAttribute uint32) (results []*Customization, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_attribute = ?", idProductAttribute).Find(&results).Error

	return
}

func (obj *_CustomizationMgr) GetBatchFromIDProductAttribute(idProductAttributes []uint32) (results []*Customization, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_attribute IN (?)", idProductAttributes).Find(&results).Error

	return
}

func (obj *_CustomizationMgr) GetFromIDAddressDelivery(idAddressDelivery uint32) (results []*Customization, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_address_delivery = ?", idAddressDelivery).Find(&results).Error

	return
}

func (obj *_CustomizationMgr) GetBatchFromIDAddressDelivery(idAddressDeliverys []uint32) (results []*Customization, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_address_delivery IN (?)", idAddressDeliverys).Find(&results).Error

	return
}

func (obj *_CustomizationMgr) GetFromIDCart(idCart uint32) (results []*Customization, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart = ?", idCart).Find(&results).Error

	return
}

func (obj *_CustomizationMgr) GetBatchFromIDCart(idCarts []uint32) (results []*Customization, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart IN (?)", idCarts).Find(&results).Error

	return
}

func (obj *_CustomizationMgr) GetFromIDProduct(idProduct int) (results []*Customization, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ?", idProduct).Find(&results).Error

	return
}

func (obj *_CustomizationMgr) GetBatchFromIDProduct(idProducts []int) (results []*Customization, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product IN (?)", idProducts).Find(&results).Error

	return
}

func (obj *_CustomizationMgr) GetFromQuantity(quantity int) (results []*Customization, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("quantity = ?", quantity).Find(&results).Error

	return
}

func (obj *_CustomizationMgr) GetBatchFromQuantity(quantitys []int) (results []*Customization, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("quantity IN (?)", quantitys).Find(&results).Error

	return
}

func (obj *_CustomizationMgr) GetFromQuantityRefunded(quantityRefunded int) (results []*Customization, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("quantity_refunded = ?", quantityRefunded).Find(&results).Error

	return
}

func (obj *_CustomizationMgr) GetBatchFromQuantityRefunded(quantityRefundeds []int) (results []*Customization, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("quantity_refunded IN (?)", quantityRefundeds).Find(&results).Error

	return
}

func (obj *_CustomizationMgr) GetFromQuantityReturned(quantityReturned int) (results []*Customization, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("quantity_returned = ?", quantityReturned).Find(&results).Error

	return
}

func (obj *_CustomizationMgr) GetBatchFromQuantityReturned(quantityReturneds []int) (results []*Customization, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("quantity_returned IN (?)", quantityReturneds).Find(&results).Error

	return
}

func (obj *_CustomizationMgr) GetFromInCart(inCart bool) (results []*Customization, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("in_cart = ?", inCart).Find(&results).Error

	return
}

func (obj *_CustomizationMgr) GetBatchFromInCart(inCarts []bool) (results []*Customization, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("in_cart IN (?)", inCarts).Find(&results).Error

	return
}

func (obj *_CustomizationMgr) FetchByPrimaryKey(idCustomization uint32, idAddressDelivery uint32, idCart uint32, idProduct int) (result Customization, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customization = ? AND id_address_delivery = ? AND id_cart = ? AND id_product = ?", idCustomization, idAddressDelivery, idCart, idProduct).Find(&result).Error

	return
}

func (obj *_CustomizationMgr) FetchIndexByIDProductAttribute(idProductAttribute uint32) (results []*Customization, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_attribute = ?", idProductAttribute).Find(&results).Error

	return
}

func (obj *_CustomizationMgr) FetchIndexByIDCartProduct(idProductAttribute uint32, idCart uint32, idProduct int) (results []*Customization, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_attribute = ? AND id_cart = ? AND id_product = ?", idProductAttribute, idCart, idProduct).Find(&results).Error

	return
}
