package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _GroupReductionMgr struct {
	*_BaseMgr
}

// GroupReductionMgr open func
func GroupReductionMgr(db *gorm.DB) *_GroupReductionMgr {
	if db == nil {
		panic(fmt.Errorf("GroupReductionMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_GroupReductionMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_group_reduction"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_GroupReductionMgr) GetTableName() string {
	return "eg_group_reduction"
}

// Get 获取
func (obj *_GroupReductionMgr) Get() (result GroupReduction, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_GroupReductionMgr) Gets() (results []*GroupReduction, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDGroupReduction id_group_reduction获取
func (obj *_GroupReductionMgr) WithIDGroupReduction(idGroupReduction string) Option {
	return optionFunc(func(o *options) { o.query["id_group_reduction"] = idGroupReduction })
}

// WithIDGroup id_group获取
func (obj *_GroupReductionMgr) WithIDGroup(idGroup uint32) Option {
	return optionFunc(func(o *options) { o.query["id_group"] = idGroup })
}

// WithIDCategory id_category获取
func (obj *_GroupReductionMgr) WithIDCategory(idCategory uint32) Option {
	return optionFunc(func(o *options) { o.query["id_category"] = idCategory })
}

// WithReduction reduction获取
func (obj *_GroupReductionMgr) WithReduction(reduction float64) Option {
	return optionFunc(func(o *options) { o.query["reduction"] = reduction })
}

// GetByOption 功能选项模式获取
func (obj *_GroupReductionMgr) GetByOption(opts ...Option) (result GroupReduction, err error) {
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
func (obj *_GroupReductionMgr) GetByOptions(opts ...Option) (results []*GroupReduction, err error) {
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

// GetFromIDGroupReduction 通过id_group_reduction获取内容
func (obj *_GroupReductionMgr) GetFromIDGroupReduction(idGroupReduction string) (result GroupReduction, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_group_reduction = ?", idGroupReduction).Find(&result).Error

	return
}

// GetBatchFromIDGroupReduction 批量唯一主键查找
func (obj *_GroupReductionMgr) GetBatchFromIDGroupReduction(idGroupReductions []string) (results []*GroupReduction, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_group_reduction IN (?)", idGroupReductions).Find(&results).Error

	return
}

// GetFromIDGroup 通过id_group获取内容
func (obj *_GroupReductionMgr) GetFromIDGroup(idGroup uint32) (results []*GroupReduction, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_group = ?", idGroup).Find(&results).Error

	return
}

// GetBatchFromIDGroup 批量唯一主键查找
func (obj *_GroupReductionMgr) GetBatchFromIDGroup(idGroups []uint32) (results []*GroupReduction, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_group IN (?)", idGroups).Find(&results).Error

	return
}

// GetFromIDCategory 通过id_category获取内容
func (obj *_GroupReductionMgr) GetFromIDCategory(idCategory uint32) (results []*GroupReduction, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_category = ?", idCategory).Find(&results).Error

	return
}

// GetBatchFromIDCategory 批量唯一主键查找
func (obj *_GroupReductionMgr) GetBatchFromIDCategory(idCategorys []uint32) (results []*GroupReduction, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_category IN (?)", idCategorys).Find(&results).Error

	return
}

// GetFromReduction 通过reduction获取内容
func (obj *_GroupReductionMgr) GetFromReduction(reduction float64) (results []*GroupReduction, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reduction = ?", reduction).Find(&results).Error

	return
}

// GetBatchFromReduction 批量唯一主键查找
func (obj *_GroupReductionMgr) GetBatchFromReduction(reductions []float64) (results []*GroupReduction, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reduction IN (?)", reductions).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_GroupReductionMgr) FetchByPrimaryKey(idGroupReduction string) (result GroupReduction, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_group_reduction = ?", idGroupReduction).Find(&result).Error

	return
}

// FetchUniqueIndexByIDGroup primay or index 获取唯一内容
func (obj *_GroupReductionMgr) FetchUniqueIndexByIDGroup(idGroup uint32, idCategory uint32) (result GroupReduction, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_group = ? AND id_category = ?", idGroup, idCategory).Find(&result).Error

	return
}
