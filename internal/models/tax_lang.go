package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _TaxLangMgr struct {
	*_BaseMgr
}

// TaxLangMgr open func
func TaxLangMgr(db *gorm.DB) *_TaxLangMgr {
	if db == nil {
		panic(fmt.Errorf("TaxLangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_TaxLangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_tax_lang"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_TaxLangMgr) GetTableName() string {
	return "ps_tax_lang"
}

// Get 获取
func (obj *_TaxLangMgr) Get() (result TaxLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_TaxLangMgr) Gets() (results []*TaxLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDTax id_tax获取
func (obj *_TaxLangMgr) WithIDTax(idTax uint32) Option {
	return optionFunc(func(o *options) { o.query["id_tax"] = idTax })
}

// WithIDLang id_lang获取
func (obj *_TaxLangMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

// WithName name获取
func (obj *_TaxLangMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// GetByOption 功能选项模式获取
func (obj *_TaxLangMgr) GetByOption(opts ...Option) (result TaxLang, err error) {
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
func (obj *_TaxLangMgr) GetByOptions(opts ...Option) (results []*TaxLang, err error) {
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

// GetFromIDTax 通过id_tax获取内容
func (obj *_TaxLangMgr) GetFromIDTax(idTax uint32) (results []*TaxLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tax = ?", idTax).Find(&results).Error

	return
}

// GetBatchFromIDTax 批量唯一主键查找
func (obj *_TaxLangMgr) GetBatchFromIDTax(idTaxs []uint32) (results []*TaxLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tax IN (?)", idTaxs).Find(&results).Error

	return
}

// GetFromIDLang 通过id_lang获取内容
func (obj *_TaxLangMgr) GetFromIDLang(idLang uint32) (results []*TaxLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error

	return
}

// GetBatchFromIDLang 批量唯一主键查找
func (obj *_TaxLangMgr) GetBatchFromIDLang(idLangs []uint32) (results []*TaxLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_TaxLangMgr) GetFromName(name string) (results []*TaxLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量唯一主键查找
func (obj *_TaxLangMgr) GetBatchFromName(names []string) (results []*TaxLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_TaxLangMgr) FetchByPrimaryKey(idTax uint32, idLang uint32) (result TaxLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tax = ? AND id_lang = ?", idTax, idLang).Find(&result).Error

	return
}
