package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _CarrierTaxRulesGroupShopMgr struct {
	*_BaseMgr
}

// CarrierTaxRulesGroupShopMgr open func
func CarrierTaxRulesGroupShopMgr(db *gorm.DB) *_CarrierTaxRulesGroupShopMgr {
	if db == nil {
		panic(fmt.Errorf("CarrierTaxRulesGroupShopMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CarrierTaxRulesGroupShopMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_carrier_tax_rules_group_shop"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CarrierTaxRulesGroupShopMgr) GetTableName() string {
	return "eg_carrier_tax_rules_group_shop"
}

// Get 获取
func (obj *_CarrierTaxRulesGroupShopMgr) Get() (result CarrierTaxRulesGroupShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CarrierTaxRulesGroupShopMgr) Gets() (results []*CarrierTaxRulesGroupShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDCarrier id_carrier获取
func (obj *_CarrierTaxRulesGroupShopMgr) WithIDCarrier(idCarrier uint32) Option {
	return optionFunc(func(o *options) { o.query["id_carrier"] = idCarrier })
}

// WithIDTaxRulesGroup id_tax_rules_group获取
func (obj *_CarrierTaxRulesGroupShopMgr) WithIDTaxRulesGroup(idTaxRulesGroup uint32) Option {
	return optionFunc(func(o *options) { o.query["id_tax_rules_group"] = idTaxRulesGroup })
}

// WithIDShop id_shop获取
func (obj *_CarrierTaxRulesGroupShopMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

// GetByOption 功能选项模式获取
func (obj *_CarrierTaxRulesGroupShopMgr) GetByOption(opts ...Option) (result CarrierTaxRulesGroupShop, err error) {
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
func (obj *_CarrierTaxRulesGroupShopMgr) GetByOptions(opts ...Option) (results []*CarrierTaxRulesGroupShop, err error) {
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

// GetFromIDCarrier 通过id_carrier获取内容
func (obj *_CarrierTaxRulesGroupShopMgr) GetFromIDCarrier(idCarrier uint32) (results []*CarrierTaxRulesGroupShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_carrier = ?", idCarrier).Find(&results).Error

	return
}

// GetBatchFromIDCarrier 批量唯一主键查找
func (obj *_CarrierTaxRulesGroupShopMgr) GetBatchFromIDCarrier(idCarriers []uint32) (results []*CarrierTaxRulesGroupShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_carrier IN (?)", idCarriers).Find(&results).Error

	return
}

// GetFromIDTaxRulesGroup 通过id_tax_rules_group获取内容
func (obj *_CarrierTaxRulesGroupShopMgr) GetFromIDTaxRulesGroup(idTaxRulesGroup uint32) (results []*CarrierTaxRulesGroupShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tax_rules_group = ?", idTaxRulesGroup).Find(&results).Error

	return
}

// GetBatchFromIDTaxRulesGroup 批量唯一主键查找
func (obj *_CarrierTaxRulesGroupShopMgr) GetBatchFromIDTaxRulesGroup(idTaxRulesGroups []uint32) (results []*CarrierTaxRulesGroupShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tax_rules_group IN (?)", idTaxRulesGroups).Find(&results).Error

	return
}

// GetFromIDShop 通过id_shop获取内容
func (obj *_CarrierTaxRulesGroupShopMgr) GetFromIDShop(idShop uint32) (results []*CarrierTaxRulesGroupShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

// GetBatchFromIDShop 批量唯一主键查找
func (obj *_CarrierTaxRulesGroupShopMgr) GetBatchFromIDShop(idShops []uint32) (results []*CarrierTaxRulesGroupShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_CarrierTaxRulesGroupShopMgr) FetchByPrimaryKey(idCarrier uint32, idTaxRulesGroup uint32, idShop uint32) (result CarrierTaxRulesGroupShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_carrier = ? AND id_tax_rules_group = ? AND id_shop = ?", idCarrier, idTaxRulesGroup, idShop).Find(&result).Error

	return
}
