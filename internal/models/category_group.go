package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _CategoryGroupMgr struct {
	*_BaseMgr
}

// CategoryGroupMgr open func
func CategoryGroupMgr(db *gorm.DB) *_CategoryGroupMgr {
	if db == nil {
		panic(fmt.Errorf("CategoryGroupMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CategoryGroupMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_category_group"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CategoryGroupMgr) GetTableName() string {
	return "eg_category_group"
}

// Get 获取
func (obj *_CategoryGroupMgr) Get() (result CategoryGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CategoryGroupMgr) Gets() (results []*CategoryGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDCategory id_category获取
func (obj *_CategoryGroupMgr) WithIDCategory(idCategory uint32) Option {
	return optionFunc(func(o *options) { o.query["id_category"] = idCategory })
}

// WithIDGroup id_group获取
func (obj *_CategoryGroupMgr) WithIDGroup(idGroup uint32) Option {
	return optionFunc(func(o *options) { o.query["id_group"] = idGroup })
}

// GetByOption 功能选项模式获取
func (obj *_CategoryGroupMgr) GetByOption(opts ...Option) (result CategoryGroup, err error) {
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
func (obj *_CategoryGroupMgr) GetByOptions(opts ...Option) (results []*CategoryGroup, err error) {
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

// GetFromIDCategory 通过id_category获取内容
func (obj *_CategoryGroupMgr) GetFromIDCategory(idCategory uint32) (results []*CategoryGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_category = ?", idCategory).Find(&results).Error

	return
}

// GetBatchFromIDCategory 批量唯一主键查找
func (obj *_CategoryGroupMgr) GetBatchFromIDCategory(idCategorys []uint32) (results []*CategoryGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_category IN (?)", idCategorys).Find(&results).Error

	return
}

// GetFromIDGroup 通过id_group获取内容
func (obj *_CategoryGroupMgr) GetFromIDGroup(idGroup uint32) (results []*CategoryGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_group = ?", idGroup).Find(&results).Error

	return
}

// GetBatchFromIDGroup 批量唯一主键查找
func (obj *_CategoryGroupMgr) GetBatchFromIDGroup(idGroups []uint32) (results []*CategoryGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_group IN (?)", idGroups).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_CategoryGroupMgr) FetchByPrimaryKey(idCategory uint32, idGroup uint32) (result CategoryGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_category = ? AND id_group = ?", idCategory, idGroup).Find(&result).Error

	return
}

// FetchIndexByIDCategory  获取多个内容
func (obj *_CategoryGroupMgr) FetchIndexByIDCategory(idCategory uint32) (results []*CategoryGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_category = ?", idCategory).Find(&results).Error

	return
}

// FetchIndexByIDGroup  获取多个内容
func (obj *_CategoryGroupMgr) FetchIndexByIDGroup(idGroup uint32) (results []*CategoryGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_group = ?", idGroup).Find(&results).Error

	return
}
