package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _ConfigurationKpiLangMgr struct {
	*_BaseMgr
}

func ConfigurationKpiLangMgr(db *gorm.DB) *_ConfigurationKpiLangMgr {
	if db == nil {
		panic(fmt.Errorf("ConfigurationKpiLangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ConfigurationKpiLangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_configuration_kpi_lang"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_ConfigurationKpiLangMgr) GetTableName() string {
	return "ps_configuration_kpi_lang"
}

func (obj *_ConfigurationKpiLangMgr) Get() (result ConfigurationKpiLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_ConfigurationKpiLangMgr) Gets() (results []*ConfigurationKpiLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_ConfigurationKpiLangMgr) WithIDConfigurationKpi(idConfigurationKpi uint32) Option {
	return optionFunc(func(o *options) { o.query["id_configuration_kpi"] = idConfigurationKpi })
}

func (obj *_ConfigurationKpiLangMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

func (obj *_ConfigurationKpiLangMgr) WithValue(value string) Option {
	return optionFunc(func(o *options) { o.query["value"] = value })
}

func (obj *_ConfigurationKpiLangMgr) WithDateUpd(dateUpd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_upd"] = dateUpd })
}

func (obj *_ConfigurationKpiLangMgr) GetByOption(opts ...Option) (result ConfigurationKpiLang, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_ConfigurationKpiLangMgr) GetByOptions(opts ...Option) (results []*ConfigurationKpiLang, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_ConfigurationKpiLangMgr) GetFromIDConfigurationKpi(idConfigurationKpi uint32) (results []*ConfigurationKpiLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_configuration_kpi = ?", idConfigurationKpi).Find(&results).Error

	return
}

func (obj *_ConfigurationKpiLangMgr) GetBatchFromIDConfigurationKpi(idConfigurationKpis []uint32) (results []*ConfigurationKpiLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_configuration_kpi IN (?)", idConfigurationKpis).Find(&results).Error

	return
}

func (obj *_ConfigurationKpiLangMgr) GetFromIDLang(idLang uint32) (results []*ConfigurationKpiLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error

	return
}

func (obj *_ConfigurationKpiLangMgr) GetBatchFromIDLang(idLangs []uint32) (results []*ConfigurationKpiLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error

	return
}

func (obj *_ConfigurationKpiLangMgr) GetFromValue(value string) (results []*ConfigurationKpiLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("value = ?", value).Find(&results).Error

	return
}

func (obj *_ConfigurationKpiLangMgr) GetBatchFromValue(values []string) (results []*ConfigurationKpiLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("value IN (?)", values).Find(&results).Error

	return
}

func (obj *_ConfigurationKpiLangMgr) GetFromDateUpd(dateUpd time.Time) (results []*ConfigurationKpiLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd = ?", dateUpd).Find(&results).Error

	return
}

func (obj *_ConfigurationKpiLangMgr) GetBatchFromDateUpd(dateUpds []time.Time) (results []*ConfigurationKpiLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd IN (?)", dateUpds).Find(&results).Error

	return
}


func (obj *_ConfigurationKpiLangMgr) FetchByPrimaryKey(idConfigurationKpi uint32, idLang uint32) (result ConfigurationKpiLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_configuration_kpi = ? AND id_lang = ?", idConfigurationKpi, idLang).Find(&result).Error

	return
}
