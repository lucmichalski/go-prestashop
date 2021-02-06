package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _LayeredCategoryMgr struct {
	*_BaseMgr
}

// LayeredCategoryMgr open func
func LayeredCategoryMgr(db *gorm.DB) *_LayeredCategoryMgr {
	if db == nil {
		panic(fmt.Errorf("LayeredCategoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_LayeredCategoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_layered_category"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_LayeredCategoryMgr) GetTableName() string {
	return "ps_layered_category"
}

// Get 获取
func (obj *_LayeredCategoryMgr) Get() (result LayeredCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_LayeredCategoryMgr) Gets() (results []*LayeredCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDLayeredCategory id_layered_category获取
func (obj *_LayeredCategoryMgr) WithIDLayeredCategory(idLayeredCategory uint32) Option {
	return optionFunc(func(o *options) { o.query["id_layered_category"] = idLayeredCategory })
}

// WithIDShop id_shop获取
func (obj *_LayeredCategoryMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

// WithIDCategory id_category获取
func (obj *_LayeredCategoryMgr) WithIDCategory(idCategory uint32) Option {
	return optionFunc(func(o *options) { o.query["id_category"] = idCategory })
}

// WithIDValue id_value获取
func (obj *_LayeredCategoryMgr) WithIDValue(idValue uint32) Option {
	return optionFunc(func(o *options) { o.query["id_value"] = idValue })
}

// WithType type获取
func (obj *_LayeredCategoryMgr) WithType(_type string) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithPosition position获取
func (obj *_LayeredCategoryMgr) WithPosition(position uint32) Option {
	return optionFunc(func(o *options) { o.query["position"] = position })
}

// WithFilterType filter_type获取
func (obj *_LayeredCategoryMgr) WithFilterType(filterType uint32) Option {
	return optionFunc(func(o *options) { o.query["filter_type"] = filterType })
}

// WithFilterShowLimit filter_show_limit获取
func (obj *_LayeredCategoryMgr) WithFilterShowLimit(filterShowLimit uint32) Option {
	return optionFunc(func(o *options) { o.query["filter_show_limit"] = filterShowLimit })
}

// GetByOption 功能选项模式获取
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

// GetByOptions 批量功能选项模式获取
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

//////////////////////////enume case ////////////////////////////////////////////

// GetFromIDLayeredCategory 通过id_layered_category获取内容
func (obj *_LayeredCategoryMgr) GetFromIDLayeredCategory(idLayeredCategory uint32) (result LayeredCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_layered_category = ?", idLayeredCategory).Find(&result).Error

	return
}

// GetBatchFromIDLayeredCategory 批量唯一主键查找
func (obj *_LayeredCategoryMgr) GetBatchFromIDLayeredCategory(idLayeredCategorys []uint32) (results []*LayeredCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_layered_category IN (?)", idLayeredCategorys).Find(&results).Error

	return
}

// GetFromIDShop 通过id_shop获取内容
func (obj *_LayeredCategoryMgr) GetFromIDShop(idShop uint32) (results []*LayeredCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

// GetBatchFromIDShop 批量唯一主键查找
func (obj *_LayeredCategoryMgr) GetBatchFromIDShop(idShops []uint32) (results []*LayeredCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

// GetFromIDCategory 通过id_category获取内容
func (obj *_LayeredCategoryMgr) GetFromIDCategory(idCategory uint32) (results []*LayeredCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_category = ?", idCategory).Find(&results).Error

	return
}

// GetBatchFromIDCategory 批量唯一主键查找
func (obj *_LayeredCategoryMgr) GetBatchFromIDCategory(idCategorys []uint32) (results []*LayeredCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_category IN (?)", idCategorys).Find(&results).Error

	return
}

// GetFromIDValue 通过id_value获取内容
func (obj *_LayeredCategoryMgr) GetFromIDValue(idValue uint32) (results []*LayeredCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_value = ?", idValue).Find(&results).Error

	return
}

// GetBatchFromIDValue 批量唯一主键查找
func (obj *_LayeredCategoryMgr) GetBatchFromIDValue(idValues []uint32) (results []*LayeredCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_value IN (?)", idValues).Find(&results).Error

	return
}

// GetFromType 通过type获取内容
func (obj *_LayeredCategoryMgr) GetFromType(_type string) (results []*LayeredCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量唯一主键查找
func (obj *_LayeredCategoryMgr) GetBatchFromType(_types []string) (results []*LayeredCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type IN (?)", _types).Find(&results).Error

	return
}

// GetFromPosition 通过position获取内容
func (obj *_LayeredCategoryMgr) GetFromPosition(position uint32) (results []*LayeredCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("position = ?", position).Find(&results).Error

	return
}

// GetBatchFromPosition 批量唯一主键查找
func (obj *_LayeredCategoryMgr) GetBatchFromPosition(positions []uint32) (results []*LayeredCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("position IN (?)", positions).Find(&results).Error

	return
}

// GetFromFilterType 通过filter_type获取内容
func (obj *_LayeredCategoryMgr) GetFromFilterType(filterType uint32) (results []*LayeredCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("filter_type = ?", filterType).Find(&results).Error

	return
}

// GetBatchFromFilterType 批量唯一主键查找
func (obj *_LayeredCategoryMgr) GetBatchFromFilterType(filterTypes []uint32) (results []*LayeredCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("filter_type IN (?)", filterTypes).Find(&results).Error

	return
}

// GetFromFilterShowLimit 通过filter_show_limit获取内容
func (obj *_LayeredCategoryMgr) GetFromFilterShowLimit(filterShowLimit uint32) (results []*LayeredCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("filter_show_limit = ?", filterShowLimit).Find(&results).Error

	return
}

// GetBatchFromFilterShowLimit 批量唯一主键查找
func (obj *_LayeredCategoryMgr) GetBatchFromFilterShowLimit(filterShowLimits []uint32) (results []*LayeredCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("filter_show_limit IN (?)", filterShowLimits).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_LayeredCategoryMgr) FetchByPrimaryKey(idLayeredCategory uint32) (result LayeredCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_layered_category = ?", idLayeredCategory).Find(&result).Error

	return
}

// FetchIndexByIDCategoryShop  获取多个内容
func (obj *_LayeredCategoryMgr) FetchIndexByIDCategoryShop(idShop uint32, idCategory uint32, idValue uint32, _type string, position uint32) (results []*LayeredCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ? AND id_category = ? AND id_value = ? AND type = ? AND position = ?", idShop, idCategory, idValue, _type, position).Find(&results).Error

	return
}

// FetchIndexByIDCategory  获取多个内容
func (obj *_LayeredCategoryMgr) FetchIndexByIDCategory(idCategory uint32, _type string) (results []*LayeredCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_category = ? AND type = ?", idCategory, _type).Find(&results).Error

	return
}
