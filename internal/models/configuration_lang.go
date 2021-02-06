package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _ConfigurationLangMgr struct {
	*_BaseMgr
}

// ConfigurationLangMgr open func
func ConfigurationLangMgr(db *gorm.DB) *_ConfigurationLangMgr {
	if db == nil {
		panic(fmt.Errorf("ConfigurationLangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ConfigurationLangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_configuration_lang"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ConfigurationLangMgr) GetTableName() string {
	return "eg_configuration_lang"
}

// Get 获取
func (obj *_ConfigurationLangMgr) Get() (result ConfigurationLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ConfigurationLangMgr) Gets() (results []*ConfigurationLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDConfiguration id_configuration获取
func (obj *_ConfigurationLangMgr) WithIDConfiguration(idConfiguration uint32) Option {
	return optionFunc(func(o *options) { o.query["id_configuration"] = idConfiguration })
}

// WithIDLang id_lang获取
func (obj *_ConfigurationLangMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

// WithValue value获取
func (obj *_ConfigurationLangMgr) WithValue(value string) Option {
	return optionFunc(func(o *options) { o.query["value"] = value })
}

// WithDateUpd date_upd获取
func (obj *_ConfigurationLangMgr) WithDateUpd(dateUpd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_upd"] = dateUpd })
}

// GetByOption 功能选项模式获取
func (obj *_ConfigurationLangMgr) GetByOption(opts ...Option) (result ConfigurationLang, err error) {
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
func (obj *_ConfigurationLangMgr) GetByOptions(opts ...Option) (results []*ConfigurationLang, err error) {
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

// GetFromIDConfiguration 通过id_configuration获取内容
func (obj *_ConfigurationLangMgr) GetFromIDConfiguration(idConfiguration uint32) (results []*ConfigurationLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_configuration = ?", idConfiguration).Find(&results).Error

	return
}

// GetBatchFromIDConfiguration 批量唯一主键查找
func (obj *_ConfigurationLangMgr) GetBatchFromIDConfiguration(idConfigurations []uint32) (results []*ConfigurationLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_configuration IN (?)", idConfigurations).Find(&results).Error

	return
}

// GetFromIDLang 通过id_lang获取内容
func (obj *_ConfigurationLangMgr) GetFromIDLang(idLang uint32) (results []*ConfigurationLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error

	return
}

// GetBatchFromIDLang 批量唯一主键查找
func (obj *_ConfigurationLangMgr) GetBatchFromIDLang(idLangs []uint32) (results []*ConfigurationLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error

	return
}

// GetFromValue 通过value获取内容
func (obj *_ConfigurationLangMgr) GetFromValue(value string) (results []*ConfigurationLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("value = ?", value).Find(&results).Error

	return
}

// GetBatchFromValue 批量唯一主键查找
func (obj *_ConfigurationLangMgr) GetBatchFromValue(values []string) (results []*ConfigurationLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("value IN (?)", values).Find(&results).Error

	return
}

// GetFromDateUpd 通过date_upd获取内容
func (obj *_ConfigurationLangMgr) GetFromDateUpd(dateUpd time.Time) (results []*ConfigurationLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd = ?", dateUpd).Find(&results).Error

	return
}

// GetBatchFromDateUpd 批量唯一主键查找
func (obj *_ConfigurationLangMgr) GetBatchFromDateUpd(dateUpds []time.Time) (results []*ConfigurationLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd IN (?)", dateUpds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_ConfigurationLangMgr) FetchByPrimaryKey(idConfiguration uint32, idLang uint32) (result ConfigurationLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_configuration = ? AND id_lang = ?", idConfiguration, idLang).Find(&result).Error

	return
}
