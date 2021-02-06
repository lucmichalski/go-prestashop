package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _GroupShopMgr struct {
	*_BaseMgr
}

// GroupShopMgr open func
func GroupShopMgr(db *gorm.DB) *_GroupShopMgr {
	if db == nil {
		panic(fmt.Errorf("GroupShopMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_GroupShopMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_group_shop"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_GroupShopMgr) GetTableName() string {
	return "eg_group_shop"
}

// Get 获取
func (obj *_GroupShopMgr) Get() (result GroupShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_GroupShopMgr) Gets() (results []*GroupShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDGroup id_group获取
func (obj *_GroupShopMgr) WithIDGroup(idGroup uint32) Option {
	return optionFunc(func(o *options) { o.query["id_group"] = idGroup })
}

// WithIDShop id_shop获取
func (obj *_GroupShopMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

// GetByOption 功能选项模式获取
func (obj *_GroupShopMgr) GetByOption(opts ...Option) (result GroupShop, err error) {
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
func (obj *_GroupShopMgr) GetByOptions(opts ...Option) (results []*GroupShop, err error) {
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

// GetFromIDGroup 通过id_group获取内容
func (obj *_GroupShopMgr) GetFromIDGroup(idGroup uint32) (results []*GroupShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_group = ?", idGroup).Find(&results).Error

	return
}

// GetBatchFromIDGroup 批量唯一主键查找
func (obj *_GroupShopMgr) GetBatchFromIDGroup(idGroups []uint32) (results []*GroupShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_group IN (?)", idGroups).Find(&results).Error

	return
}

// GetFromIDShop 通过id_shop获取内容
func (obj *_GroupShopMgr) GetFromIDShop(idShop uint32) (results []*GroupShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

// GetBatchFromIDShop 批量唯一主键查找
func (obj *_GroupShopMgr) GetBatchFromIDShop(idShops []uint32) (results []*GroupShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_GroupShopMgr) FetchByPrimaryKey(idGroup uint32, idShop uint32) (result GroupShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_group = ? AND id_shop = ?", idGroup, idShop).Find(&result).Error

	return
}

// FetchIndexByIDShop  获取多个内容
func (obj *_GroupShopMgr) FetchIndexByIDShop(idShop uint32) (results []*GroupShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}
