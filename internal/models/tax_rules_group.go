package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _TaxRulesGroupMgr struct {
	*_BaseMgr
}

// TaxRulesGroupMgr open func
func TaxRulesGroupMgr(db *gorm.DB) *_TaxRulesGroupMgr {
	if db == nil {
		panic(fmt.Errorf("TaxRulesGroupMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_TaxRulesGroupMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_tax_rules_group"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_TaxRulesGroupMgr) GetTableName() string {
	return "eg_tax_rules_group"
}

// Get 获取
func (obj *_TaxRulesGroupMgr) Get() (result TaxRulesGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_TaxRulesGroupMgr) Gets() (results []*TaxRulesGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDTaxRulesGroup id_tax_rules_group获取
func (obj *_TaxRulesGroupMgr) WithIDTaxRulesGroup(idTaxRulesGroup int) Option {
	return optionFunc(func(o *options) { o.query["id_tax_rules_group"] = idTaxRulesGroup })
}

// WithName name获取
func (obj *_TaxRulesGroupMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithActive active获取
func (obj *_TaxRulesGroupMgr) WithActive(active int) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}

// WithDeleted deleted获取
func (obj *_TaxRulesGroupMgr) WithDeleted(deleted bool) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// WithDateAdd date_add获取
func (obj *_TaxRulesGroupMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

// WithDateUpd date_upd获取
func (obj *_TaxRulesGroupMgr) WithDateUpd(dateUpd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_upd"] = dateUpd })
}

// GetByOption 功能选项模式获取
func (obj *_TaxRulesGroupMgr) GetByOption(opts ...Option) (result TaxRulesGroup, err error) {
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
func (obj *_TaxRulesGroupMgr) GetByOptions(opts ...Option) (results []*TaxRulesGroup, err error) {
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

// GetFromIDTaxRulesGroup 通过id_tax_rules_group获取内容
func (obj *_TaxRulesGroupMgr) GetFromIDTaxRulesGroup(idTaxRulesGroup int) (result TaxRulesGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tax_rules_group = ?", idTaxRulesGroup).Find(&result).Error

	return
}

// GetBatchFromIDTaxRulesGroup 批量唯一主键查找
func (obj *_TaxRulesGroupMgr) GetBatchFromIDTaxRulesGroup(idTaxRulesGroups []int) (results []*TaxRulesGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tax_rules_group IN (?)", idTaxRulesGroups).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_TaxRulesGroupMgr) GetFromName(name string) (results []*TaxRulesGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量唯一主键查找
func (obj *_TaxRulesGroupMgr) GetBatchFromName(names []string) (results []*TaxRulesGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

// GetFromActive 通过active获取内容
func (obj *_TaxRulesGroupMgr) GetFromActive(active int) (results []*TaxRulesGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active = ?", active).Find(&results).Error

	return
}

// GetBatchFromActive 批量唯一主键查找
func (obj *_TaxRulesGroupMgr) GetBatchFromActive(actives []int) (results []*TaxRulesGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active IN (?)", actives).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容
func (obj *_TaxRulesGroupMgr) GetFromDeleted(deleted bool) (results []*TaxRulesGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("deleted = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量唯一主键查找
func (obj *_TaxRulesGroupMgr) GetBatchFromDeleted(deleteds []bool) (results []*TaxRulesGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("deleted IN (?)", deleteds).Find(&results).Error

	return
}

// GetFromDateAdd 通过date_add获取内容
func (obj *_TaxRulesGroupMgr) GetFromDateAdd(dateAdd time.Time) (results []*TaxRulesGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error

	return
}

// GetBatchFromDateAdd 批量唯一主键查找
func (obj *_TaxRulesGroupMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*TaxRulesGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error

	return
}

// GetFromDateUpd 通过date_upd获取内容
func (obj *_TaxRulesGroupMgr) GetFromDateUpd(dateUpd time.Time) (results []*TaxRulesGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd = ?", dateUpd).Find(&results).Error

	return
}

// GetBatchFromDateUpd 批量唯一主键查找
func (obj *_TaxRulesGroupMgr) GetBatchFromDateUpd(dateUpds []time.Time) (results []*TaxRulesGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd IN (?)", dateUpds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_TaxRulesGroupMgr) FetchByPrimaryKey(idTaxRulesGroup int) (result TaxRulesGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tax_rules_group = ?", idTaxRulesGroup).Find(&result).Error

	return
}
