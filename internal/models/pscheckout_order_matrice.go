package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _PscheckoutOrderMatriceMgr struct {
	*_BaseMgr
}

func PscheckoutOrderMatriceMgr(db *gorm.DB) *_PscheckoutOrderMatriceMgr {
	if db == nil {
		panic(fmt.Errorf("PscheckoutOrderMatriceMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_PscheckoutOrderMatriceMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_pscheckout_order_matrice"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_PscheckoutOrderMatriceMgr) GetTableName() string {
	return "ps_pscheckout_order_matrice"
}

func (obj *_PscheckoutOrderMatriceMgr) Get() (result PscheckoutOrderMatrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_PscheckoutOrderMatriceMgr) Gets() (results []*PscheckoutOrderMatrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_PscheckoutOrderMatriceMgr) WithIDOrderMatrice(idOrderMatrice uint32) Option {
	return optionFunc(func(o *options) { o.query["id_order_matrice"] = idOrderMatrice })
}

func (obj *_PscheckoutOrderMatriceMgr) WithIDOrderPrestashop(idOrderPrestashop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_order_prestashop"] = idOrderPrestashop })
}

func (obj *_PscheckoutOrderMatriceMgr) WithIDOrderPaypal(idOrderPaypal string) Option {
	return optionFunc(func(o *options) { o.query["id_order_paypal"] = idOrderPaypal })
}

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


func (obj *_PscheckoutOrderMatriceMgr) GetFromIDOrderMatrice(idOrderMatrice uint32) (result PscheckoutOrderMatrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_matrice = ?", idOrderMatrice).Find(&result).Error

	return
}

func (obj *_PscheckoutOrderMatriceMgr) GetBatchFromIDOrderMatrice(idOrderMatrices []uint32) (results []*PscheckoutOrderMatrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_matrice IN (?)", idOrderMatrices).Find(&results).Error

	return
}

func (obj *_PscheckoutOrderMatriceMgr) GetFromIDOrderPrestashop(idOrderPrestashop uint32) (results []*PscheckoutOrderMatrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_prestashop = ?", idOrderPrestashop).Find(&results).Error

	return
}

func (obj *_PscheckoutOrderMatriceMgr) GetBatchFromIDOrderPrestashop(idOrderPrestashops []uint32) (results []*PscheckoutOrderMatrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_prestashop IN (?)", idOrderPrestashops).Find(&results).Error

	return
}

func (obj *_PscheckoutOrderMatriceMgr) GetFromIDOrderPaypal(idOrderPaypal string) (results []*PscheckoutOrderMatrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_paypal = ?", idOrderPaypal).Find(&results).Error

	return
}

func (obj *_PscheckoutOrderMatriceMgr) GetBatchFromIDOrderPaypal(idOrderPaypals []string) (results []*PscheckoutOrderMatrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_paypal IN (?)", idOrderPaypals).Find(&results).Error

	return
}


func (obj *_PscheckoutOrderMatriceMgr) FetchByPrimaryKey(idOrderMatrice uint32) (result PscheckoutOrderMatrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_matrice = ?", idOrderMatrice).Find(&result).Error

	return
}
