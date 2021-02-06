package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _AttributeImpactMgr struct {
	*_BaseMgr
}

func AttributeImpactMgr(db *gorm.DB) *_AttributeImpactMgr {
	if db == nil {
		panic(fmt.Errorf("AttributeImpactMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AttributeImpactMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_attribute_impact"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_AttributeImpactMgr) GetTableName() string {
	return "ps_attribute_impact"
}

func (obj *_AttributeImpactMgr) Get() (result AttributeImpact, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_AttributeImpactMgr) Gets() (results []*AttributeImpact, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_AttributeImpactMgr) WithIDAttributeImpact(idAttributeImpact uint32) Option {
	return optionFunc(func(o *options) { o.query["id_attribute_impact"] = idAttributeImpact })
}

func (obj *_AttributeImpactMgr) WithIDProduct(idProduct uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product"] = idProduct })
}

func (obj *_AttributeImpactMgr) WithIDAttribute(idAttribute uint32) Option {
	return optionFunc(func(o *options) { o.query["id_attribute"] = idAttribute })
}

func (obj *_AttributeImpactMgr) WithWeight(weight float64) Option {
	return optionFunc(func(o *options) { o.query["weight"] = weight })
}

func (obj *_AttributeImpactMgr) WithPrice(price float64) Option {
	return optionFunc(func(o *options) { o.query["price"] = price })
}

func (obj *_AttributeImpactMgr) GetByOption(opts ...Option) (result AttributeImpact, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_AttributeImpactMgr) GetByOptions(opts ...Option) (results []*AttributeImpact, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_AttributeImpactMgr) GetFromIDAttributeImpact(idAttributeImpact uint32) (result AttributeImpact, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attribute_impact = ?", idAttributeImpact).Find(&result).Error

	return
}

func (obj *_AttributeImpactMgr) GetBatchFromIDAttributeImpact(idAttributeImpacts []uint32) (results []*AttributeImpact, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attribute_impact IN (?)", idAttributeImpacts).Find(&results).Error

	return
}

func (obj *_AttributeImpactMgr) GetFromIDProduct(idProduct uint32) (results []*AttributeImpact, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ?", idProduct).Find(&results).Error

	return
}

func (obj *_AttributeImpactMgr) GetBatchFromIDProduct(idProducts []uint32) (results []*AttributeImpact, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product IN (?)", idProducts).Find(&results).Error

	return
}

func (obj *_AttributeImpactMgr) GetFromIDAttribute(idAttribute uint32) (results []*AttributeImpact, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attribute = ?", idAttribute).Find(&results).Error

	return
}

func (obj *_AttributeImpactMgr) GetBatchFromIDAttribute(idAttributes []uint32) (results []*AttributeImpact, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attribute IN (?)", idAttributes).Find(&results).Error

	return
}

func (obj *_AttributeImpactMgr) GetFromWeight(weight float64) (results []*AttributeImpact, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("weight = ?", weight).Find(&results).Error

	return
}

func (obj *_AttributeImpactMgr) GetBatchFromWeight(weights []float64) (results []*AttributeImpact, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("weight IN (?)", weights).Find(&results).Error

	return
}

func (obj *_AttributeImpactMgr) GetFromPrice(price float64) (results []*AttributeImpact, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("price = ?", price).Find(&results).Error

	return
}

func (obj *_AttributeImpactMgr) GetBatchFromPrice(prices []float64) (results []*AttributeImpact, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("price IN (?)", prices).Find(&results).Error

	return
}


func (obj *_AttributeImpactMgr) FetchByPrimaryKey(idAttributeImpact uint32) (result AttributeImpact, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attribute_impact = ?", idAttributeImpact).Find(&result).Error

	return
}

func (obj *_AttributeImpactMgr) FetchUniqueIndexByIDProduct(idProduct uint32, idAttribute uint32) (result AttributeImpact, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ? AND id_attribute = ?", idProduct, idAttribute).Find(&result).Error

	return
}
