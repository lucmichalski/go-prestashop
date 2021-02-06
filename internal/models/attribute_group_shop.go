package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _AttributeGroupShopMgr struct {
	*_BaseMgr
}

// AttributeGroupShopMgr open func
func AttributeGroupShopMgr(db *gorm.DB) *_AttributeGroupShopMgr {
	if db == nil {
		panic(fmt.Errorf("AttributeGroupShopMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AttributeGroupShopMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_attribute_group_shop"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AttributeGroupShopMgr) GetTableName() string {
	return "eg_attribute_group_shop"
}

// Get 获取
func (obj *_AttributeGroupShopMgr) Get() (result AttributeGroupShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AttributeGroupShopMgr) Gets() (results []*AttributeGroupShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDAttributeGroup id_attribute_group获取
func (obj *_AttributeGroupShopMgr) WithIDAttributeGroup(idAttributeGroup int) Option {
	return optionFunc(func(o *options) { o.query["id_attribute_group"] = idAttributeGroup })
}

// WithIDShop id_shop获取
func (obj *_AttributeGroupShopMgr) WithIDShop(idShop int) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

// GetByOption 功能选项模式获取
func (obj *_AttributeGroupShopMgr) GetByOption(opts ...Option) (result AttributeGroupShop, err error) {
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
func (obj *_AttributeGroupShopMgr) GetByOptions(opts ...Option) (results []*AttributeGroupShop, err error) {
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

// GetFromIDAttributeGroup 通过id_attribute_group获取内容
func (obj *_AttributeGroupShopMgr) GetFromIDAttributeGroup(idAttributeGroup int) (results []*AttributeGroupShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attribute_group = ?", idAttributeGroup).Find(&results).Error

	return
}

// GetBatchFromIDAttributeGroup 批量唯一主键查找
func (obj *_AttributeGroupShopMgr) GetBatchFromIDAttributeGroup(idAttributeGroups []int) (results []*AttributeGroupShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attribute_group IN (?)", idAttributeGroups).Find(&results).Error

	return
}

// GetFromIDShop 通过id_shop获取内容
func (obj *_AttributeGroupShopMgr) GetFromIDShop(idShop int) (results []*AttributeGroupShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

// GetBatchFromIDShop 批量唯一主键查找
func (obj *_AttributeGroupShopMgr) GetBatchFromIDShop(idShops []int) (results []*AttributeGroupShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_AttributeGroupShopMgr) FetchByPrimaryKey(idAttributeGroup int, idShop int) (result AttributeGroupShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attribute_group = ? AND id_shop = ?", idAttributeGroup, idShop).Find(&result).Error

	return
}

// FetchIndexByIDXB43BAEE667A664FB  获取多个内容
func (obj *_AttributeGroupShopMgr) FetchIndexByIDXB43BAEE667A664FB(idAttributeGroup int) (results []*AttributeGroupShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attribute_group = ?", idAttributeGroup).Find(&results).Error

	return
}

// FetchIndexByIDXB43BAEE6274A50A0  获取多个内容
func (obj *_AttributeGroupShopMgr) FetchIndexByIDXB43BAEE6274A50A0(idShop int) (results []*AttributeGroupShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}
