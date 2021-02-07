package models

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _ConfigurationMgr struct {
	*_BaseMgr
}

func ConfigurationMgr(db *gorm.DB) *_ConfigurationMgr {
	if db == nil {
		panic(fmt.Errorf("ConfigurationMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ConfigurationMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_configuration"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_ConfigurationMgr) GetTableName() string {
	return "ps_configuration"
}

func (obj *_ConfigurationMgr) Get() (result Configuration, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_ConfigurationMgr) Gets() (results []*Configuration, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

func (obj *_ConfigurationMgr) WithIDConfiguration(idConfiguration uint32) Option {
	return optionFunc(func(o *options) { o.query["id_configuration"] = idConfiguration })
}

func (obj *_ConfigurationMgr) WithIDShopGroup(idShopGroup uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop_group"] = idShopGroup })
}

func (obj *_ConfigurationMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

func (obj *_ConfigurationMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

func (obj *_ConfigurationMgr) WithValue(value string) Option {
	return optionFunc(func(o *options) { o.query["value"] = value })
}

func (obj *_ConfigurationMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

func (obj *_ConfigurationMgr) WithDateUpd(dateUpd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_upd"] = dateUpd })
}

func (obj *_ConfigurationMgr) GetByOption(opts ...Option) (result Configuration, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_ConfigurationMgr) GetByOptions(opts ...Option) (results []*Configuration, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

func (obj *_ConfigurationMgr) GetFromIDConfiguration(idConfiguration uint32) (result Configuration, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_configuration = ?", idConfiguration).Find(&result).Error

	return
}

func (obj *_ConfigurationMgr) GetBatchFromIDConfiguration(idConfigurations []uint32) (results []*Configuration, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_configuration IN (?)", idConfigurations).Find(&results).Error

	return
}

func (obj *_ConfigurationMgr) GetFromIDShopGroup(idShopGroup uint32) (results []*Configuration, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_group = ?", idShopGroup).Find(&results).Error

	return
}

func (obj *_ConfigurationMgr) GetBatchFromIDShopGroup(idShopGroups []uint32) (results []*Configuration, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_group IN (?)", idShopGroups).Find(&results).Error

	return
}

func (obj *_ConfigurationMgr) GetFromIDShop(idShop uint32) (results []*Configuration, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

func (obj *_ConfigurationMgr) GetBatchFromIDShop(idShops []uint32) (results []*Configuration, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

func (obj *_ConfigurationMgr) GetFromName(name string) (results []*Configuration, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

func (obj *_ConfigurationMgr) GetBatchFromName(names []string) (results []*Configuration, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

func (obj *_ConfigurationMgr) GetFromValue(value string) (results []*Configuration, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("value = ?", value).Find(&results).Error

	return
}

func (obj *_ConfigurationMgr) GetBatchFromValue(values []string) (results []*Configuration, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("value IN (?)", values).Find(&results).Error

	return
}

func (obj *_ConfigurationMgr) GetFromDateAdd(dateAdd time.Time) (results []*Configuration, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error

	return
}

func (obj *_ConfigurationMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*Configuration, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error

	return
}

func (obj *_ConfigurationMgr) GetFromDateUpd(dateUpd time.Time) (results []*Configuration, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd = ?", dateUpd).Find(&results).Error

	return
}

func (obj *_ConfigurationMgr) GetBatchFromDateUpd(dateUpds []time.Time) (results []*Configuration, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd IN (?)", dateUpds).Find(&results).Error

	return
}

func (obj *_ConfigurationMgr) FetchByPrimaryKey(idConfiguration uint32) (result Configuration, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_configuration = ?", idConfiguration).Find(&result).Error

	return
}

func (obj *_ConfigurationMgr) FetchIndexByIDShopGroup(idShopGroup uint32) (results []*Configuration, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_group = ?", idShopGroup).Find(&results).Error

	return
}

func (obj *_ConfigurationMgr) FetchIndexByIDShop(idShop uint32) (results []*Configuration, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

func (obj *_ConfigurationMgr) FetchIndexByName(name string) (results []*Configuration, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}
