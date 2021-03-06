package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _AccessoryMgr struct {
	*_BaseMgr
}

func AccessoryMgr(db *gorm.DB) *_AccessoryMgr {
	if db == nil {
		panic(fmt.Errorf("AccessoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AccessoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_accessory"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_AccessoryMgr) GetTableName() string {
	return "ps_accessory"
}

func (obj *_AccessoryMgr) Get() (result Accessory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_AccessoryMgr) Gets() (results []*Accessory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

func (obj *_AccessoryMgr) WithIDProduct1(idProduct1 uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product_1"] = idProduct1 })
}

func (obj *_AccessoryMgr) WithIDProduct2(idProduct2 uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product_2"] = idProduct2 })
}

func (obj *_AccessoryMgr) GetByOption(opts ...Option) (result Accessory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_AccessoryMgr) GetByOptions(opts ...Option) (results []*Accessory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

func (obj *_AccessoryMgr) GetFromIDProduct1(idProduct1 uint32) (results []*Accessory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_1 = ?", idProduct1).Find(&results).Error

	return
}

func (obj *_AccessoryMgr) GetBatchFromIDProduct1(idProduct1s []uint32) (results []*Accessory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_1 IN (?)", idProduct1s).Find(&results).Error

	return
}

func (obj *_AccessoryMgr) GetFromIDProduct2(idProduct2 uint32) (results []*Accessory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_2 = ?", idProduct2).Find(&results).Error

	return
}

func (obj *_AccessoryMgr) GetBatchFromIDProduct2(idProduct2s []uint32) (results []*Accessory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_2 IN (?)", idProduct2s).Find(&results).Error

	return
}

func (obj *_AccessoryMgr) FetchIndexByAccessoryProduct(idProduct1 uint32, idProduct2 uint32) (results []*Accessory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_1 = ? AND id_product_2 = ?", idProduct1, idProduct2).Find(&results).Error

	return
}
