package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _ConfigurationKpiMgr struct {
	*_BaseMgr
}

// ConfigurationKpiMgr open func
func ConfigurationKpiMgr(db *gorm.DB) *_ConfigurationKpiMgr {
	if db == nil {
		panic(fmt.Errorf("ConfigurationKpiMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ConfigurationKpiMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_configuration_kpi"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ConfigurationKpiMgr) GetTableName() string {
	return "eg_configuration_kpi"
}

// Get 获取
func (obj *_ConfigurationKpiMgr) Get() (result ConfigurationKpi, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ConfigurationKpiMgr) Gets() (results []*ConfigurationKpi, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDConfigurationKpi id_configuration_kpi获取
func (obj *_ConfigurationKpiMgr) WithIDConfigurationKpi(idConfigurationKpi uint32) Option {
	return optionFunc(func(o *options) { o.query["id_configuration_kpi"] = idConfigurationKpi })
}

// WithIDShopGroup id_shop_group获取
func (obj *_ConfigurationKpiMgr) WithIDShopGroup(idShopGroup uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop_group"] = idShopGroup })
}

// WithIDShop id_shop获取
func (obj *_ConfigurationKpiMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

// WithName name获取
func (obj *_ConfigurationKpiMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithValue value获取
func (obj *_ConfigurationKpiMgr) WithValue(value string) Option {
	return optionFunc(func(o *options) { o.query["value"] = value })
}

// WithDateAdd date_add获取
func (obj *_ConfigurationKpiMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

// WithDateUpd date_upd获取
func (obj *_ConfigurationKpiMgr) WithDateUpd(dateUpd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_upd"] = dateUpd })
}

// GetByOption 功能选项模式获取
func (obj *_ConfigurationKpiMgr) GetByOption(opts ...Option) (result ConfigurationKpi, err error) {
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
func (obj *_ConfigurationKpiMgr) GetByOptions(opts ...Option) (results []*ConfigurationKpi, err error) {
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

// GetFromIDConfigurationKpi 通过id_configuration_kpi获取内容
func (obj *_ConfigurationKpiMgr) GetFromIDConfigurationKpi(idConfigurationKpi uint32) (result ConfigurationKpi, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_configuration_kpi = ?", idConfigurationKpi).Find(&result).Error

	return
}

// GetBatchFromIDConfigurationKpi 批量唯一主键查找
func (obj *_ConfigurationKpiMgr) GetBatchFromIDConfigurationKpi(idConfigurationKpis []uint32) (results []*ConfigurationKpi, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_configuration_kpi IN (?)", idConfigurationKpis).Find(&results).Error

	return
}

// GetFromIDShopGroup 通过id_shop_group获取内容
func (obj *_ConfigurationKpiMgr) GetFromIDShopGroup(idShopGroup uint32) (results []*ConfigurationKpi, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_group = ?", idShopGroup).Find(&results).Error

	return
}

// GetBatchFromIDShopGroup 批量唯一主键查找
func (obj *_ConfigurationKpiMgr) GetBatchFromIDShopGroup(idShopGroups []uint32) (results []*ConfigurationKpi, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_group IN (?)", idShopGroups).Find(&results).Error

	return
}

// GetFromIDShop 通过id_shop获取内容
func (obj *_ConfigurationKpiMgr) GetFromIDShop(idShop uint32) (results []*ConfigurationKpi, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

// GetBatchFromIDShop 批量唯一主键查找
func (obj *_ConfigurationKpiMgr) GetBatchFromIDShop(idShops []uint32) (results []*ConfigurationKpi, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_ConfigurationKpiMgr) GetFromName(name string) (results []*ConfigurationKpi, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量唯一主键查找
func (obj *_ConfigurationKpiMgr) GetBatchFromName(names []string) (results []*ConfigurationKpi, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

// GetFromValue 通过value获取内容
func (obj *_ConfigurationKpiMgr) GetFromValue(value string) (results []*ConfigurationKpi, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("value = ?", value).Find(&results).Error

	return
}

// GetBatchFromValue 批量唯一主键查找
func (obj *_ConfigurationKpiMgr) GetBatchFromValue(values []string) (results []*ConfigurationKpi, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("value IN (?)", values).Find(&results).Error

	return
}

// GetFromDateAdd 通过date_add获取内容
func (obj *_ConfigurationKpiMgr) GetFromDateAdd(dateAdd time.Time) (results []*ConfigurationKpi, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error

	return
}

// GetBatchFromDateAdd 批量唯一主键查找
func (obj *_ConfigurationKpiMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*ConfigurationKpi, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error

	return
}

// GetFromDateUpd 通过date_upd获取内容
func (obj *_ConfigurationKpiMgr) GetFromDateUpd(dateUpd time.Time) (results []*ConfigurationKpi, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd = ?", dateUpd).Find(&results).Error

	return
}

// GetBatchFromDateUpd 批量唯一主键查找
func (obj *_ConfigurationKpiMgr) GetBatchFromDateUpd(dateUpds []time.Time) (results []*ConfigurationKpi, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd IN (?)", dateUpds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_ConfigurationKpiMgr) FetchByPrimaryKey(idConfigurationKpi uint32) (result ConfigurationKpi, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_configuration_kpi = ?", idConfigurationKpi).Find(&result).Error

	return
}

// FetchIndexByIDShopGroup  获取多个内容
func (obj *_ConfigurationKpiMgr) FetchIndexByIDShopGroup(idShopGroup uint32) (results []*ConfigurationKpi, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_group = ?", idShopGroup).Find(&results).Error

	return
}

// FetchIndexByIDShop  获取多个内容
func (obj *_ConfigurationKpiMgr) FetchIndexByIDShop(idShop uint32) (results []*ConfigurationKpi, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

// FetchIndexByName  获取多个内容
func (obj *_ConfigurationKpiMgr) FetchIndexByName(name string) (results []*ConfigurationKpi, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}
