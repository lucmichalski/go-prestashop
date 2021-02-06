package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _OrderDetailTaxMgr struct {
	*_BaseMgr
}

// OrderDetailTaxMgr open func
func OrderDetailTaxMgr(db *gorm.DB) *_OrderDetailTaxMgr {
	if db == nil {
		panic(fmt.Errorf("OrderDetailTaxMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_OrderDetailTaxMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_order_detail_tax"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_OrderDetailTaxMgr) GetTableName() string {
	return "eg_order_detail_tax"
}

// Get 获取
func (obj *_OrderDetailTaxMgr) Get() (result OrderDetailTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_OrderDetailTaxMgr) Gets() (results []*OrderDetailTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDOrderDetail id_order_detail获取
func (obj *_OrderDetailTaxMgr) WithIDOrderDetail(idOrderDetail int) Option {
	return optionFunc(func(o *options) { o.query["id_order_detail"] = idOrderDetail })
}

// WithIDTax id_tax获取
func (obj *_OrderDetailTaxMgr) WithIDTax(idTax int) Option {
	return optionFunc(func(o *options) { o.query["id_tax"] = idTax })
}

// WithUnitAmount unit_amount获取
func (obj *_OrderDetailTaxMgr) WithUnitAmount(unitAmount float64) Option {
	return optionFunc(func(o *options) { o.query["unit_amount"] = unitAmount })
}

// WithTotalAmount total_amount获取
func (obj *_OrderDetailTaxMgr) WithTotalAmount(totalAmount float64) Option {
	return optionFunc(func(o *options) { o.query["total_amount"] = totalAmount })
}

// GetByOption 功能选项模式获取
func (obj *_OrderDetailTaxMgr) GetByOption(opts ...Option) (result OrderDetailTax, err error) {
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
func (obj *_OrderDetailTaxMgr) GetByOptions(opts ...Option) (results []*OrderDetailTax, err error) {
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

// GetFromIDOrderDetail 通过id_order_detail获取内容
func (obj *_OrderDetailTaxMgr) GetFromIDOrderDetail(idOrderDetail int) (results []*OrderDetailTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_detail = ?", idOrderDetail).Find(&results).Error

	return
}

// GetBatchFromIDOrderDetail 批量唯一主键查找
func (obj *_OrderDetailTaxMgr) GetBatchFromIDOrderDetail(idOrderDetails []int) (results []*OrderDetailTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_detail IN (?)", idOrderDetails).Find(&results).Error

	return
}

// GetFromIDTax 通过id_tax获取内容
func (obj *_OrderDetailTaxMgr) GetFromIDTax(idTax int) (results []*OrderDetailTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tax = ?", idTax).Find(&results).Error

	return
}

// GetBatchFromIDTax 批量唯一主键查找
func (obj *_OrderDetailTaxMgr) GetBatchFromIDTax(idTaxs []int) (results []*OrderDetailTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tax IN (?)", idTaxs).Find(&results).Error

	return
}

// GetFromUnitAmount 通过unit_amount获取内容
func (obj *_OrderDetailTaxMgr) GetFromUnitAmount(unitAmount float64) (results []*OrderDetailTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("unit_amount = ?", unitAmount).Find(&results).Error

	return
}

// GetBatchFromUnitAmount 批量唯一主键查找
func (obj *_OrderDetailTaxMgr) GetBatchFromUnitAmount(unitAmounts []float64) (results []*OrderDetailTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("unit_amount IN (?)", unitAmounts).Find(&results).Error

	return
}

// GetFromTotalAmount 通过total_amount获取内容
func (obj *_OrderDetailTaxMgr) GetFromTotalAmount(totalAmount float64) (results []*OrderDetailTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_amount = ?", totalAmount).Find(&results).Error

	return
}

// GetBatchFromTotalAmount 批量唯一主键查找
func (obj *_OrderDetailTaxMgr) GetBatchFromTotalAmount(totalAmounts []float64) (results []*OrderDetailTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_amount IN (?)", totalAmounts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchIndexByIDOrderDetail  获取多个内容
func (obj *_OrderDetailTaxMgr) FetchIndexByIDOrderDetail(idOrderDetail int) (results []*OrderDetailTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_detail = ?", idOrderDetail).Find(&results).Error

	return
}

// FetchIndexByIDTax  获取多个内容
func (obj *_OrderDetailTaxMgr) FetchIndexByIDTax(idTax int) (results []*OrderDetailTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tax = ?", idTax).Find(&results).Error

	return
}
