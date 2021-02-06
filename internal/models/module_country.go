package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _ModuleCountryMgr struct {
	*_BaseMgr
}

// ModuleCountryMgr open func
func ModuleCountryMgr(db *gorm.DB) *_ModuleCountryMgr {
	if db == nil {
		panic(fmt.Errorf("ModuleCountryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ModuleCountryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_module_country"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ModuleCountryMgr) GetTableName() string {
	return "ps_module_country"
}

// Get 获取
func (obj *_ModuleCountryMgr) Get() (result ModuleCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ModuleCountryMgr) Gets() (results []*ModuleCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDModule id_module获取
func (obj *_ModuleCountryMgr) WithIDModule(idModule uint32) Option {
	return optionFunc(func(o *options) { o.query["id_module"] = idModule })
}

// WithIDShop id_shop获取
func (obj *_ModuleCountryMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

// WithIDCountry id_country获取
func (obj *_ModuleCountryMgr) WithIDCountry(idCountry uint32) Option {
	return optionFunc(func(o *options) { o.query["id_country"] = idCountry })
}

// GetByOption 功能选项模式获取
func (obj *_ModuleCountryMgr) GetByOption(opts ...Option) (result ModuleCountry, err error) {
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
func (obj *_ModuleCountryMgr) GetByOptions(opts ...Option) (results []*ModuleCountry, err error) {
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
func (obj *_ModuleCountryMgr) GetFromIDModule(idModule uint32) (results []*ModuleCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_module = ?", idModule).Find(&results).Error

	return
}

// GetBatchFromIDModule 批量唯一主键查找
func (obj *_ModuleCountryMgr) GetBatchFromIDModule(idModules []uint32) (results []*ModuleCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_module IN (?)", idModules).Find(&results).Error

	return
}

// GetFromIDShop 通过id_shop获取内容
func (obj *_ModuleCountryMgr) GetFromIDShop(idShop uint32) (results []*ModuleCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

// GetBatchFromIDShop 批量唯一主键查找
func (obj *_ModuleCountryMgr) GetBatchFromIDShop(idShops []uint32) (results []*ModuleCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

// GetFromIDCountry 通过id_country获取内容
func (obj *_ModuleCountryMgr) GetFromIDCountry(idCountry uint32) (results []*ModuleCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_country = ?", idCountry).Find(&results).Error

	return
}

// GetBatchFromIDCountry 批量唯一主键查找
func (obj *_ModuleCountryMgr) GetBatchFromIDCountry(idCountrys []uint32) (results []*ModuleCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_country IN (?)", idCountrys).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_ModuleCountryMgr) FetchByPrimaryKey(idModule uint32, idShop uint32, idCountry uint32) (result ModuleCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_module = ? AND id_shop = ? AND id_country = ?", idModule, idShop, idCountry).Find(&result).Error

	return
}
