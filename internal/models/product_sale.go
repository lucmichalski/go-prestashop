package model

import (
	"context"
	"fmt"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type _ProductSaleMgr struct {
	*_BaseMgr
}

func ProductSaleMgr(db *gorm.DB) *_ProductSaleMgr {
	if db == nil {
		panic(fmt.Errorf("ProductSaleMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ProductSaleMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_product_sale"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_ProductSaleMgr) GetTableName() string {
	return "ps_product_sale"
}

func (obj *_ProductSaleMgr) Get() (result ProductSale, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_ProductSaleMgr) Gets() (results []*ProductSale, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_ProductSaleMgr) WithIDProduct(idProduct uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product"] = idProduct })
}

func (obj *_ProductSaleMgr) WithQuantity(quantity uint32) Option {
	return optionFunc(func(o *options) { o.query["quantity"] = quantity })
}

func (obj *_ProductSaleMgr) WithSaleNbr(saleNbr uint32) Option {
	return optionFunc(func(o *options) { o.query["sale_nbr"] = saleNbr })
}

func (obj *_ProductSaleMgr) WithDateUpd(dateUpd datatypes.Date) Option {
	return optionFunc(func(o *options) { o.query["date_upd"] = dateUpd })
}

func (obj *_ProductSaleMgr) GetByOption(opts ...Option) (result ProductSale, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_ProductSaleMgr) GetByOptions(opts ...Option) (results []*ProductSale, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_ProductSaleMgr) GetFromIDProduct(idProduct uint32) (result ProductSale, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ?", idProduct).Find(&result).Error

	return
}

func (obj *_ProductSaleMgr) GetBatchFromIDProduct(idProducts []uint32) (results []*ProductSale, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product IN (?)", idProducts).Find(&results).Error

	return
}

func (obj *_ProductSaleMgr) GetFromQuantity(quantity uint32) (results []*ProductSale, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("quantity = ?", quantity).Find(&results).Error

	return
}

func (obj *_ProductSaleMgr) GetBatchFromQuantity(quantitys []uint32) (results []*ProductSale, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("quantity IN (?)", quantitys).Find(&results).Error

	return
}

func (obj *_ProductSaleMgr) GetFromSaleNbr(saleNbr uint32) (results []*ProductSale, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("sale_nbr = ?", saleNbr).Find(&results).Error

	return
}

func (obj *_ProductSaleMgr) GetBatchFromSaleNbr(saleNbrs []uint32) (results []*ProductSale, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("sale_nbr IN (?)", saleNbrs).Find(&results).Error

	return
}

func (obj *_ProductSaleMgr) GetFromDateUpd(dateUpd datatypes.Date) (results []*ProductSale, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd = ?", dateUpd).Find(&results).Error

	return
}

func (obj *_ProductSaleMgr) GetBatchFromDateUpd(dateUpds []datatypes.Date) (results []*ProductSale, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd IN (?)", dateUpds).Find(&results).Error

	return
}


func (obj *_ProductSaleMgr) FetchByPrimaryKey(idProduct uint32) (result ProductSale, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ?", idProduct).Find(&result).Error

	return
}

func (obj *_ProductSaleMgr) FetchIndexByQuantity(quantity uint32) (results []*ProductSale, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("quantity = ?", quantity).Find(&results).Error

	return
}
