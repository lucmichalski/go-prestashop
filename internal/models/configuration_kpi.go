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

func ConfigurationKpiMgr(db *gorm.DB) *_ConfigurationKpiMgr {
	if db == nil {
		panic(fmt.Errorf("ConfigurationKpiMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ConfigurationKpiMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_configuration_kpi"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_ConfigurationKpiMgr) GetTableName() string {
	return "ps_configuration_kpi"
}

func (obj *_ConfigurationKpiMgr) Get() (result ConfigurationKpi, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_ConfigurationKpiMgr) Gets() (results []*ConfigurationKpi, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_ConfigurationKpiMgr) WithIDConfigurationKpi(idConfigurationKpi uint32) Option {
	return optionFunc(func(o *options) { o.query["id_configuration_kpi"] = idConfigurationKpi })
}

func (obj *_ConfigurationKpiMgr) WithIDShopGroup(idShopGroup uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop_group"] = idShopGroup })
}

func (obj *_ConfigurationKpiMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

func (obj *_ConfigurationKpiMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

func (obj *_ConfigurationKpiMgr) WithValue(value string) Option {
	return optionFunc(func(o *options) { o.query["value"] = value })
}

func (obj *_ConfigurationKpiMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

func (obj *_ConfigurationKpiMgr) WithDateUpd(dateUpd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_upd"] = dateUpd })
}

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


func (obj *_ConfigurationKpiMgr) GetFromIDConfigurationKpi(idConfigurationKpi uint32) (result ConfigurationKpi, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_configuration_kpi = ?", idConfigurationKpi).Find(&result).Error

	return
}

func (obj *_ConfigurationKpiMgr) GetBatchFromIDConfigurationKpi(idConfigurationKpis []uint32) (results []*ConfigurationKpi, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_configuration_kpi IN (?)", idConfigurationKpis).Find(&results).Error

	return
}

func (obj *_ConfigurationKpiMgr) GetFromIDShopGroup(idShopGroup uint32) (results []*ConfigurationKpi, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_group = ?", idShopGroup).Find(&results).Error

	return
}

func (obj *_ConfigurationKpiMgr) GetBatchFromIDShopGroup(idShopGroups []uint32) (results []*ConfigurationKpi, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_group IN (?)", idShopGroups).Find(&results).Error

	return
}

func (obj *_ConfigurationKpiMgr) GetFromIDShop(idShop uint32) (results []*ConfigurationKpi, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

func (obj *_ConfigurationKpiMgr) GetBatchFromIDShop(idShops []uint32) (results []*ConfigurationKpi, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

func (obj *_ConfigurationKpiMgr) GetFromName(name string) (results []*ConfigurationKpi, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

func (obj *_ConfigurationKpiMgr) GetBatchFromName(names []string) (results []*ConfigurationKpi, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

func (obj *_ConfigurationKpiMgr) GetFromValue(value string) (results []*ConfigurationKpi, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("value = ?", value).Find(&results).Error

	return
}

func (obj *_ConfigurationKpiMgr) GetBatchFromValue(values []string) (results []*ConfigurationKpi, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("value IN (?)", values).Find(&results).Error

	return
}

func (obj *_ConfigurationKpiMgr) GetFromDateAdd(dateAdd time.Time) (results []*ConfigurationKpi, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error

	return
}

func (obj *_ConfigurationKpiMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*ConfigurationKpi, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error

	return
}

func (obj *_ConfigurationKpiMgr) GetFromDateUpd(dateUpd time.Time) (results []*ConfigurationKpi, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd = ?", dateUpd).Find(&results).Error

	return
}

func (obj *_ConfigurationKpiMgr) GetBatchFromDateUpd(dateUpds []time.Time) (results []*ConfigurationKpi, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd IN (?)", dateUpds).Find(&results).Error

	return
}


func (obj *_ConfigurationKpiMgr) FetchByPrimaryKey(idConfigurationKpi uint32) (result ConfigurationKpi, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_configuration_kpi = ?", idConfigurationKpi).Find(&result).Error

	return
}

func (obj *_ConfigurationKpiMgr) FetchIndexByIDShopGroup(idShopGroup uint32) (results []*ConfigurationKpi, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_group = ?", idShopGroup).Find(&results).Error

	return
}

func (obj *_ConfigurationKpiMgr) FetchIndexByIDShop(idShop uint32) (results []*ConfigurationKpi, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

func (obj *_ConfigurationKpiMgr) FetchIndexByName(name string) (results []*ConfigurationKpi, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}
