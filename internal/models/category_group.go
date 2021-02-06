package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _CategoryGroupMgr struct {
	*_BaseMgr
}

func CategoryGroupMgr(db *gorm.DB) *_CategoryGroupMgr {
	if db == nil {
		panic(fmt.Errorf("CategoryGroupMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CategoryGroupMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_category_group"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_CategoryGroupMgr) GetTableName() string {
	return "ps_category_group"
}

func (obj *_CategoryGroupMgr) Get() (result CategoryGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_CategoryGroupMgr) Gets() (results []*CategoryGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_CategoryGroupMgr) WithIDCategory(idCategory uint32) Option {
	return optionFunc(func(o *options) { o.query["id_category"] = idCategory })
}

func (obj *_CategoryGroupMgr) WithIDGroup(idGroup uint32) Option {
	return optionFunc(func(o *options) { o.query["id_group"] = idGroup })
}

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


func (obj *_CategoryGroupMgr) GetFromIDCategory(idCategory uint32) (results []*CategoryGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_category = ?", idCategory).Find(&results).Error

	return
}

func (obj *_CategoryGroupMgr) GetBatchFromIDCategory(idCategorys []uint32) (results []*CategoryGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_category IN (?)", idCategorys).Find(&results).Error

	return
}

func (obj *_CategoryGroupMgr) GetFromIDGroup(idGroup uint32) (results []*CategoryGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_group = ?", idGroup).Find(&results).Error

	return
}

func (obj *_CategoryGroupMgr) GetBatchFromIDGroup(idGroups []uint32) (results []*CategoryGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_group IN (?)", idGroups).Find(&results).Error

	return
}


func (obj *_CategoryGroupMgr) FetchByPrimaryKey(idCategory uint32, idGroup uint32) (result CategoryGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_category = ? AND id_group = ?", idCategory, idGroup).Find(&result).Error

	return
}

func (obj *_CategoryGroupMgr) FetchIndexByIDCategory(idCategory uint32) (results []*CategoryGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_category = ?", idCategory).Find(&results).Error

	return
}

func (obj *_CategoryGroupMgr) FetchIndexByIDGroup(idGroup uint32) (results []*CategoryGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_group = ?", idGroup).Find(&results).Error

	return
}
