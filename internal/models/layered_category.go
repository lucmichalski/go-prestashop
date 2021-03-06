package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _LayeredCategoryMgr struct {
	*_BaseMgr
}

func LayeredCategoryMgr(db *gorm.DB) *_LayeredCategoryMgr {
	if db == nil {
		panic(fmt.Errorf("LayeredCategoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_LayeredCategoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_layered_category"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_LayeredCategoryMgr) GetTableName() string {
	return "ps_layered_category"
}

func (obj *_LayeredCategoryMgr) Get() (result LayeredCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_LayeredCategoryMgr) Gets() (results []*LayeredCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

func (obj *_LayeredCategoryMgr) WithIDLayeredCategory(idLayeredCategory uint32) Option {
	return optionFunc(func(o *options) { o.query["id_layered_category"] = idLayeredCategory })
}

func (obj *_LayeredCategoryMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

func (obj *_LayeredCategoryMgr) WithIDCategory(idCategory uint32) Option {
	return optionFunc(func(o *options) { o.query["id_category"] = idCategory })
}

func (obj *_LayeredCategoryMgr) WithIDValue(idValue uint32) Option {
	return optionFunc(func(o *options) { o.query["id_value"] = idValue })
}

func (obj *_LayeredCategoryMgr) WithType(_type string) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

func (obj *_LayeredCategoryMgr) WithPosition(position uint32) Option {
	return optionFunc(func(o *options) { o.query["position"] = position })
}

func (obj *_LayeredCategoryMgr) WithFilterType(filterType uint32) Option {
	return optionFunc(func(o *options) { o.query["filter_type"] = filterType })
}

func (obj *_LayeredCategoryMgr) WithFilterShowLimit(filterShowLimit uint32) Option {
	return optionFunc(func(o *options) { o.query["filter_show_limit"] = filterShowLimit })
}

func (obj *_LayeredCategoryMgr) GetByOption(opts ...Option) (result LayeredCategory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_LayeredCategoryMgr) GetByOptions(opts ...Option) (results []*LayeredCategory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

func (obj *_LayeredCategoryMgr) GetFromIDLayeredCategory(idLayeredCategory uint32) (result LayeredCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_layered_category = ?", idLayeredCategory).Find(&result).Error

	return
}

func (obj *_LayeredCategoryMgr) GetBatchFromIDLayeredCategory(idLayeredCategorys []uint32) (results []*LayeredCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_layered_category IN (?)", idLayeredCategorys).Find(&results).Error

	return
}

func (obj *_LayeredCategoryMgr) GetFromIDShop(idShop uint32) (results []*LayeredCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

func (obj *_LayeredCategoryMgr) GetBatchFromIDShop(idShops []uint32) (results []*LayeredCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

func (obj *_LayeredCategoryMgr) GetFromIDCategory(idCategory uint32) (results []*LayeredCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_category = ?", idCategory).Find(&results).Error

	return
}

func (obj *_LayeredCategoryMgr) GetBatchFromIDCategory(idCategorys []uint32) (results []*LayeredCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_category IN (?)", idCategorys).Find(&results).Error

	return
}

func (obj *_LayeredCategoryMgr) GetFromIDValue(idValue uint32) (results []*LayeredCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_value = ?", idValue).Find(&results).Error

	return
}

func (obj *_LayeredCategoryMgr) GetBatchFromIDValue(idValues []uint32) (results []*LayeredCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_value IN (?)", idValues).Find(&results).Error

	return
}

func (obj *_LayeredCategoryMgr) GetFromType(_type string) (results []*LayeredCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type = ?", _type).Find(&results).Error

	return
}

func (obj *_LayeredCategoryMgr) GetBatchFromType(_types []string) (results []*LayeredCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type IN (?)", _types).Find(&results).Error

	return
}

func (obj *_LayeredCategoryMgr) GetFromPosition(position uint32) (results []*LayeredCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("position = ?", position).Find(&results).Error

	return
}

func (obj *_LayeredCategoryMgr) GetBatchFromPosition(positions []uint32) (results []*LayeredCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("position IN (?)", positions).Find(&results).Error

	return
}

func (obj *_LayeredCategoryMgr) GetFromFilterType(filterType uint32) (results []*LayeredCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("filter_type = ?", filterType).Find(&results).Error

	return
}

func (obj *_LayeredCategoryMgr) GetBatchFromFilterType(filterTypes []uint32) (results []*LayeredCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("filter_type IN (?)", filterTypes).Find(&results).Error

	return
}

func (obj *_LayeredCategoryMgr) GetFromFilterShowLimit(filterShowLimit uint32) (results []*LayeredCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("filter_show_limit = ?", filterShowLimit).Find(&results).Error

	return
}

func (obj *_LayeredCategoryMgr) GetBatchFromFilterShowLimit(filterShowLimits []uint32) (results []*LayeredCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("filter_show_limit IN (?)", filterShowLimits).Find(&results).Error

	return
}

func (obj *_LayeredCategoryMgr) FetchByPrimaryKey(idLayeredCategory uint32) (result LayeredCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_layered_category = ?", idLayeredCategory).Find(&result).Error

	return
}

func (obj *_LayeredCategoryMgr) FetchIndexByIDCategoryShop(idShop uint32, idCategory uint32, idValue uint32, _type string, position uint32) (results []*LayeredCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ? AND id_category = ? AND id_value = ? AND type = ? AND position = ?", idShop, idCategory, idValue, _type, position).Find(&results).Error

	return
}

func (obj *_LayeredCategoryMgr) FetchIndexByIDCategory(idCategory uint32, _type string) (results []*LayeredCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_category = ? AND type = ?", idCategory, _type).Find(&results).Error

	return
}
