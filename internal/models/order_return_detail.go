package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _OrderReturnDetailMgr struct {
	*_BaseMgr
}

// OrderReturnDetailMgr open func
func OrderReturnDetailMgr(db *gorm.DB) *_OrderReturnDetailMgr {
	if db == nil {
		panic(fmt.Errorf("OrderReturnDetailMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_OrderReturnDetailMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_order_return_detail"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_OrderReturnDetailMgr) GetTableName() string {
	return "eg_order_return_detail"
}

// Get 获取
func (obj *_OrderReturnDetailMgr) Get() (result OrderReturnDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_OrderReturnDetailMgr) Gets() (results []*OrderReturnDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDOrderReturn id_order_return获取
func (obj *_OrderReturnDetailMgr) WithIDOrderReturn(idOrderReturn uint32) Option {
	return optionFunc(func(o *options) { o.query["id_order_return"] = idOrderReturn })
}

// WithIDOrderDetail id_order_detail获取
func (obj *_OrderReturnDetailMgr) WithIDOrderDetail(idOrderDetail uint32) Option {
	return optionFunc(func(o *options) { o.query["id_order_detail"] = idOrderDetail })
}

// WithIDCustomization id_customization获取
func (obj *_OrderReturnDetailMgr) WithIDCustomization(idCustomization uint32) Option {
	return optionFunc(func(o *options) { o.query["id_customization"] = idCustomization })
}

// WithProductQuantity product_quantity获取
func (obj *_OrderReturnDetailMgr) WithProductQuantity(productQuantity uint32) Option {
	return optionFunc(func(o *options) { o.query["product_quantity"] = productQuantity })
}

// GetByOption 功能选项模式获取
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

// GetByOptions 批量功能选项模式获取
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

//////////////////////////enume case ////////////////////////////////////////////

// GetFromIDOrderReturn 通过id_order_return获取内容
func (obj *_OrderReturnDetailMgr) GetFromIDOrderReturn(idOrderReturn uint32) (results []*OrderReturnDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_return = ?", idOrderReturn).Find(&results).Error

	return
}

// GetBatchFromIDOrderReturn 批量唯一主键查找
func (obj *_OrderReturnDetailMgr) GetBatchFromIDOrderReturn(idOrderReturns []uint32) (results []*OrderReturnDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_return IN (?)", idOrderReturns).Find(&results).Error

	return
}

// GetFromIDOrderDetail 通过id_order_detail获取内容
func (obj *_OrderReturnDetailMgr) GetFromIDOrderDetail(idOrderDetail uint32) (results []*OrderReturnDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_detail = ?", idOrderDetail).Find(&results).Error

	return
}

// GetBatchFromIDOrderDetail 批量唯一主键查找
func (obj *_OrderReturnDetailMgr) GetBatchFromIDOrderDetail(idOrderDetails []uint32) (results []*OrderReturnDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_detail IN (?)", idOrderDetails).Find(&results).Error

	return
}

// GetFromIDCustomization 通过id_customization获取内容
func (obj *_OrderReturnDetailMgr) GetFromIDCustomization(idCustomization uint32) (results []*OrderReturnDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customization = ?", idCustomization).Find(&results).Error

	return
}

// GetBatchFromIDCustomization 批量唯一主键查找
func (obj *_OrderReturnDetailMgr) GetBatchFromIDCustomization(idCustomizations []uint32) (results []*OrderReturnDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customization IN (?)", idCustomizations).Find(&results).Error

	return
}

// GetFromProductQuantity 通过product_quantity获取内容
func (obj *_OrderReturnDetailMgr) GetFromProductQuantity(productQuantity uint32) (results []*OrderReturnDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_quantity = ?", productQuantity).Find(&results).Error

	return
}

// GetBatchFromProductQuantity 批量唯一主键查找
func (obj *_OrderReturnDetailMgr) GetBatchFromProductQuantity(productQuantitys []uint32) (results []*OrderReturnDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_quantity IN (?)", productQuantitys).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_OrderReturnDetailMgr) FetchByPrimaryKey(idOrderReturn uint32, idOrderDetail uint32, idCustomization uint32) (result OrderReturnDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_return = ? AND id_order_detail = ? AND id_customization = ?", idOrderReturn, idOrderDetail, idCustomization).Find(&result).Error

	return
}
