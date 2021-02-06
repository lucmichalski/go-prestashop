package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _TaxMgr struct {
	*_BaseMgr
}

// TaxMgr open func
func TaxMgr(db *gorm.DB) *_TaxMgr {
	if db == nil {
		panic(fmt.Errorf("TaxMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_TaxMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_tax"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_TaxMgr) GetTableName() string {
	return "eg_tax"
}

// Get 获取
func (obj *_TaxMgr) Get() (result Tax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_TaxMgr) Gets() (results []*Tax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDTax id_tax获取
func (obj *_TaxMgr) WithIDTax(idTax uint32) Option {
	return optionFunc(func(o *options) { o.query["id_tax"] = idTax })
}

// WithRate rate获取
func (obj *_TaxMgr) WithRate(rate float64) Option {
	return optionFunc(func(o *options) { o.query["rate"] = rate })
}

// WithActive active获取
func (obj *_TaxMgr) WithActive(active bool) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}

// WithDeleted deleted获取
func (obj *_TaxMgr) WithDeleted(deleted bool) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// GetByOption 功能选项模式获取
func (obj *_TaxMgr) GetByOption(opts ...Option) (result Tax, err error) {
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
func (obj *_TaxMgr) GetByOptions(opts ...Option) (results []*Tax, err error) {
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
func (obj *_TaxMgr) GetFromIDTax(idTax uint32) (result Tax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tax = ?", idTax).Find(&result).Error

	return
}

// GetBatchFromIDTax 批量唯一主键查找
func (obj *_TaxMgr) GetBatchFromIDTax(idTaxs []uint32) (results []*Tax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tax IN (?)", idTaxs).Find(&results).Error

	return
}

// GetFromRate 通过rate获取内容
func (obj *_TaxMgr) GetFromRate(rate float64) (results []*Tax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("rate = ?", rate).Find(&results).Error

	return
}

// GetBatchFromRate 批量唯一主键查找
func (obj *_TaxMgr) GetBatchFromRate(rates []float64) (results []*Tax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("rate IN (?)", rates).Find(&results).Error

	return
}

// GetFromActive 通过active获取内容
func (obj *_TaxMgr) GetFromActive(active bool) (results []*Tax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active = ?", active).Find(&results).Error

	return
}

// GetBatchFromActive 批量唯一主键查找
func (obj *_TaxMgr) GetBatchFromActive(actives []bool) (results []*Tax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active IN (?)", actives).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容
func (obj *_TaxMgr) GetFromDeleted(deleted bool) (results []*Tax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("deleted = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量唯一主键查找
func (obj *_TaxMgr) GetBatchFromDeleted(deleteds []bool) (results []*Tax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("deleted IN (?)", deleteds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_TaxMgr) FetchByPrimaryKey(idTax uint32) (result Tax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tax = ?", idTax).Find(&result).Error

	return
}
