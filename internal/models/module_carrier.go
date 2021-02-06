package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _ModuleCarrierMgr struct {
	*_BaseMgr
}

// ModuleCarrierMgr open func
func ModuleCarrierMgr(db *gorm.DB) *_ModuleCarrierMgr {
	if db == nil {
		panic(fmt.Errorf("ModuleCarrierMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ModuleCarrierMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_module_carrier"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ModuleCarrierMgr) GetTableName() string {
	return "ps_module_carrier"
}

// Get 获取
func (obj *_ModuleCarrierMgr) Get() (result ModuleCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ModuleCarrierMgr) Gets() (results []*ModuleCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDModule id_module获取
func (obj *_ModuleCarrierMgr) WithIDModule(idModule uint32) Option {
	return optionFunc(func(o *options) { o.query["id_module"] = idModule })
}

// WithIDShop id_shop获取
func (obj *_ModuleCarrierMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

// WithIDReference id_reference获取
func (obj *_ModuleCarrierMgr) WithIDReference(idReference int) Option {
	return optionFunc(func(o *options) { o.query["id_reference"] = idReference })
}

// GetByOption 功能选项模式获取
func (obj *_ModuleCarrierMgr) GetByOption(opts ...Option) (result ModuleCarrier, err error) {
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
func (obj *_ModuleCarrierMgr) GetByOptions(opts ...Option) (results []*ModuleCarrier, err error) {
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

// GetFromIDModule 通过id_module获取内容
func (obj *_ModuleCarrierMgr) GetFromIDModule(idModule uint32) (results []*ModuleCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_module = ?", idModule).Find(&results).Error

	return
}

// GetBatchFromIDModule 批量唯一主键查找
func (obj *_ModuleCarrierMgr) GetBatchFromIDModule(idModules []uint32) (results []*ModuleCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_module IN (?)", idModules).Find(&results).Error

	return
}

// GetFromIDShop 通过id_shop获取内容
func (obj *_ModuleCarrierMgr) GetFromIDShop(idShop uint32) (results []*ModuleCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

// GetBatchFromIDShop 批量唯一主键查找
func (obj *_ModuleCarrierMgr) GetBatchFromIDShop(idShops []uint32) (results []*ModuleCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

// GetFromIDReference 通过id_reference获取内容
func (obj *_ModuleCarrierMgr) GetFromIDReference(idReference int) (results []*ModuleCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_reference = ?", idReference).Find(&results).Error

	return
}

// GetBatchFromIDReference 批量唯一主键查找
func (obj *_ModuleCarrierMgr) GetBatchFromIDReference(idReferences []int) (results []*ModuleCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_reference IN (?)", idReferences).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_ModuleCarrierMgr) FetchByPrimaryKey(idModule uint32, idShop uint32, idReference int) (result ModuleCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_module = ? AND id_shop = ? AND id_reference = ?", idModule, idShop, idReference).Find(&result).Error

	return
}
