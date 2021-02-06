package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _PscheckoutOrderMatriceMgr struct {
	*_BaseMgr
}

// PscheckoutOrderMatriceMgr open func
func PscheckoutOrderMatriceMgr(db *gorm.DB) *_PscheckoutOrderMatriceMgr {
	if db == nil {
		panic(fmt.Errorf("PscheckoutOrderMatriceMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_PscheckoutOrderMatriceMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_pscheckout_order_matrice"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_PscheckoutOrderMatriceMgr) GetTableName() string {
	return "ps_pscheckout_order_matrice"
}

// Get 获取
func (obj *_PscheckoutOrderMatriceMgr) Get() (result PscheckoutOrderMatrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_PscheckoutOrderMatriceMgr) Gets() (results []*PscheckoutOrderMatrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDOrderMatrice id_order_matrice获取
func (obj *_PscheckoutOrderMatriceMgr) WithIDOrderMatrice(idOrderMatrice uint32) Option {
	return optionFunc(func(o *options) { o.query["id_order_matrice"] = idOrderMatrice })
}

// WithIDOrderPrestashop id_order_prestashop获取
func (obj *_PscheckoutOrderMatriceMgr) WithIDOrderPrestashop(idOrderPrestashop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_order_prestashop"] = idOrderPrestashop })
}

// WithIDOrderPaypal id_order_paypal获取
func (obj *_PscheckoutOrderMatriceMgr) WithIDOrderPaypal(idOrderPaypal string) Option {
	return optionFunc(func(o *options) { o.query["id_order_paypal"] = idOrderPaypal })
}

// GetByOption 功能选项模式获取
func (obj *_PscheckoutOrderMatriceMgr) GetByOption(opts ...Option) (result PscheckoutOrderMatrice, err error) {
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
func (obj *_PscheckoutOrderMatriceMgr) GetByOptions(opts ...Option) (results []*PscheckoutOrderMatrice, err error) {
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

// GetFromIDOrderMatrice 通过id_order_matrice获取内容
func (obj *_PscheckoutOrderMatriceMgr) GetFromIDOrderMatrice(idOrderMatrice uint32) (result PscheckoutOrderMatrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_matrice = ?", idOrderMatrice).Find(&result).Error

	return
}

// GetBatchFromIDOrderMatrice 批量唯一主键查找
func (obj *_PscheckoutOrderMatriceMgr) GetBatchFromIDOrderMatrice(idOrderMatrices []uint32) (results []*PscheckoutOrderMatrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_matrice IN (?)", idOrderMatrices).Find(&results).Error

	return
}

// GetFromIDOrderPrestashop 通过id_order_prestashop获取内容
func (obj *_PscheckoutOrderMatriceMgr) GetFromIDOrderPrestashop(idOrderPrestashop uint32) (results []*PscheckoutOrderMatrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_prestashop = ?", idOrderPrestashop).Find(&results).Error

	return
}

// GetBatchFromIDOrderPrestashop 批量唯一主键查找
func (obj *_PscheckoutOrderMatriceMgr) GetBatchFromIDOrderPrestashop(idOrderPrestashops []uint32) (results []*PscheckoutOrderMatrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_prestashop IN (?)", idOrderPrestashops).Find(&results).Error

	return
}

// GetFromIDOrderPaypal 通过id_order_paypal获取内容
func (obj *_PscheckoutOrderMatriceMgr) GetFromIDOrderPaypal(idOrderPaypal string) (results []*PscheckoutOrderMatrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_paypal = ?", idOrderPaypal).Find(&results).Error

	return
}

// GetBatchFromIDOrderPaypal 批量唯一主键查找
func (obj *_PscheckoutOrderMatriceMgr) GetBatchFromIDOrderPaypal(idOrderPaypals []string) (results []*PscheckoutOrderMatrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_paypal IN (?)", idOrderPaypals).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_PscheckoutOrderMatriceMgr) FetchByPrimaryKey(idOrderMatrice uint32) (result PscheckoutOrderMatrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_matrice = ?", idOrderMatrice).Find(&result).Error

	return
}
