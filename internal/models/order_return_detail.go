package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _OrderReturnDetailMgr struct {
	*_BaseMgr
}

func OrderReturnDetailMgr(db *gorm.DB) *_OrderReturnDetailMgr {
	if db == nil {
		panic(fmt.Errorf("OrderReturnDetailMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_OrderReturnDetailMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_order_return_detail"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_OrderReturnDetailMgr) GetTableName() string {
	return "ps_order_return_detail"
}

func (obj *_OrderReturnDetailMgr) Get() (result OrderReturnDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_OrderReturnDetailMgr) Gets() (results []*OrderReturnDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

func (obj *_OrderReturnDetailMgr) WithIDOrderReturn(idOrderReturn uint32) Option {
	return optionFunc(func(o *options) { o.query["id_order_return"] = idOrderReturn })
}

func (obj *_OrderReturnDetailMgr) WithIDOrderDetail(idOrderDetail uint32) Option {
	return optionFunc(func(o *options) { o.query["id_order_detail"] = idOrderDetail })
}

func (obj *_OrderReturnDetailMgr) WithIDCustomization(idCustomization uint32) Option {
	return optionFunc(func(o *options) { o.query["id_customization"] = idCustomization })
}

func (obj *_OrderReturnDetailMgr) WithProductQuantity(productQuantity uint32) Option {
	return optionFunc(func(o *options) { o.query["product_quantity"] = productQuantity })
}

func (obj *_OrderReturnDetailMgr) GetByOption(opts ...Option) (result OrderReturnDetail, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_OrderReturnDetailMgr) GetByOptions(opts ...Option) (results []*OrderReturnDetail, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

func (obj *_OrderReturnDetailMgr) GetFromIDOrderReturn(idOrderReturn uint32) (results []*OrderReturnDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_return = ?", idOrderReturn).Find(&results).Error

	return
}

func (obj *_OrderReturnDetailMgr) GetBatchFromIDOrderReturn(idOrderReturns []uint32) (results []*OrderReturnDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_return IN (?)", idOrderReturns).Find(&results).Error

	return
}

func (obj *_OrderReturnDetailMgr) GetFromIDOrderDetail(idOrderDetail uint32) (results []*OrderReturnDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_detail = ?", idOrderDetail).Find(&results).Error

	return
}

func (obj *_OrderReturnDetailMgr) GetBatchFromIDOrderDetail(idOrderDetails []uint32) (results []*OrderReturnDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_detail IN (?)", idOrderDetails).Find(&results).Error

	return
}

func (obj *_OrderReturnDetailMgr) GetFromIDCustomization(idCustomization uint32) (results []*OrderReturnDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customization = ?", idCustomization).Find(&results).Error

	return
}

func (obj *_OrderReturnDetailMgr) GetBatchFromIDCustomization(idCustomizations []uint32) (results []*OrderReturnDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customization IN (?)", idCustomizations).Find(&results).Error

	return
}

func (obj *_OrderReturnDetailMgr) GetFromProductQuantity(productQuantity uint32) (results []*OrderReturnDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_quantity = ?", productQuantity).Find(&results).Error

	return
}

func (obj *_OrderReturnDetailMgr) GetBatchFromProductQuantity(productQuantitys []uint32) (results []*OrderReturnDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_quantity IN (?)", productQuantitys).Find(&results).Error

	return
}

func (obj *_OrderReturnDetailMgr) FetchByPrimaryKey(idOrderReturn uint32, idOrderDetail uint32, idCustomization uint32) (result OrderReturnDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_return = ? AND id_order_detail = ? AND id_customization = ?", idOrderReturn, idOrderDetail, idCustomization).Find(&result).Error

	return
}
