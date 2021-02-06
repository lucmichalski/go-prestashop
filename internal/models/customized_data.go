package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _CustomizedDataMgr struct {
	*_BaseMgr
}

func CustomizedDataMgr(db *gorm.DB) *_CustomizedDataMgr {
	if db == nil {
		panic(fmt.Errorf("CustomizedDataMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CustomizedDataMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_customized_data"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_CustomizedDataMgr) GetTableName() string {
	return "ps_customized_data"
}

func (obj *_CustomizedDataMgr) Get() (result CustomizedData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_CustomizedDataMgr) Gets() (results []*CustomizedData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_CustomizedDataMgr) WithIDCustomization(idCustomization uint32) Option {
	return optionFunc(func(o *options) { o.query["id_customization"] = idCustomization })
}

func (obj *_CustomizedDataMgr) WithType(_type bool) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

func (obj *_CustomizedDataMgr) WithIndex(index int) Option {
	return optionFunc(func(o *options) { o.query["index"] = index })
}

func (obj *_CustomizedDataMgr) WithValue(value string) Option {
	return optionFunc(func(o *options) { o.query["value"] = value })
}

func (obj *_CustomizedDataMgr) WithIDModule(idModule int) Option {
	return optionFunc(func(o *options) { o.query["id_module"] = idModule })
}

func (obj *_CustomizedDataMgr) WithPrice(price float64) Option {
	return optionFunc(func(o *options) { o.query["price"] = price })
}

func (obj *_CustomizedDataMgr) WithWeight(weight float64) Option {
	return optionFunc(func(o *options) { o.query["weight"] = weight })
}

func (obj *_CustomizedDataMgr) GetByOption(opts ...Option) (result CustomizedData, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_CustomizedDataMgr) GetByOptions(opts ...Option) (results []*CustomizedData, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_CustomizedDataMgr) GetFromIDCustomization(idCustomization uint32) (results []*CustomizedData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customization = ?", idCustomization).Find(&results).Error

	return
}

func (obj *_CustomizedDataMgr) GetBatchFromIDCustomization(idCustomizations []uint32) (results []*CustomizedData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customization IN (?)", idCustomizations).Find(&results).Error

	return
}

func (obj *_CustomizedDataMgr) GetFromType(_type bool) (results []*CustomizedData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type = ?", _type).Find(&results).Error

	return
}

func (obj *_CustomizedDataMgr) GetBatchFromType(_types []bool) (results []*CustomizedData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type IN (?)", _types).Find(&results).Error

	return
}

func (obj *_CustomizedDataMgr) GetFromIndex(index int) (results []*CustomizedData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("index = ?", index).Find(&results).Error

	return
}

func (obj *_CustomizedDataMgr) GetBatchFromIndex(indexs []int) (results []*CustomizedData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("index IN (?)", indexs).Find(&results).Error

	return
}

func (obj *_CustomizedDataMgr) GetFromValue(value string) (results []*CustomizedData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("value = ?", value).Find(&results).Error

	return
}

func (obj *_CustomizedDataMgr) GetBatchFromValue(values []string) (results []*CustomizedData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("value IN (?)", values).Find(&results).Error

	return
}

func (obj *_CustomizedDataMgr) GetFromIDModule(idModule int) (results []*CustomizedData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_module = ?", idModule).Find(&results).Error

	return
}

func (obj *_CustomizedDataMgr) GetBatchFromIDModule(idModules []int) (results []*CustomizedData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_module IN (?)", idModules).Find(&results).Error

	return
}

func (obj *_CustomizedDataMgr) GetFromPrice(price float64) (results []*CustomizedData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("price = ?", price).Find(&results).Error

	return
}

func (obj *_CustomizedDataMgr) GetBatchFromPrice(prices []float64) (results []*CustomizedData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("price IN (?)", prices).Find(&results).Error

	return
}

func (obj *_CustomizedDataMgr) GetFromWeight(weight float64) (results []*CustomizedData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("weight = ?", weight).Find(&results).Error

	return
}

func (obj *_CustomizedDataMgr) GetBatchFromWeight(weights []float64) (results []*CustomizedData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("weight IN (?)", weights).Find(&results).Error

	return
}


func (obj *_CustomizedDataMgr) FetchByPrimaryKey(idCustomization uint32, _type bool, index int) (result CustomizedData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customization = ? AND type = ? AND index = ?", idCustomization, _type, index).Find(&result).Error

	return
}
