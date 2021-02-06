package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _StockMvtMgr struct {
	*_BaseMgr
}

// StockMvtMgr open func
func StockMvtMgr(db *gorm.DB) *_StockMvtMgr {
	if db == nil {
		panic(fmt.Errorf("StockMvtMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_StockMvtMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_stock_mvt"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_StockMvtMgr) GetTableName() string {
	return "ps_stock_mvt"
}

// Get 获取
func (obj *_StockMvtMgr) Get() (result StockMvt, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_StockMvtMgr) Gets() (results []*StockMvt, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDStockMvt id_stock_mvt获取
func (obj *_StockMvtMgr) WithIDStockMvt(idStockMvt int64) Option {
	return optionFunc(func(o *options) { o.query["id_stock_mvt"] = idStockMvt })
}

// WithIDStock id_stock获取
func (obj *_StockMvtMgr) WithIDStock(idStock int) Option {
	return optionFunc(func(o *options) { o.query["id_stock"] = idStock })
}

// WithIDOrder id_order获取
func (obj *_StockMvtMgr) WithIDOrder(idOrder int) Option {
	return optionFunc(func(o *options) { o.query["id_order"] = idOrder })
}

// WithIDSupplyOrder id_supply_order获取
func (obj *_StockMvtMgr) WithIDSupplyOrder(idSupplyOrder int) Option {
	return optionFunc(func(o *options) { o.query["id_supply_order"] = idSupplyOrder })
}

// WithIDStockMvtReason id_stock_mvt_reason获取
func (obj *_StockMvtMgr) WithIDStockMvtReason(idStockMvtReason int) Option {
	return optionFunc(func(o *options) { o.query["id_stock_mvt_reason"] = idStockMvtReason })
}

// WithIDEmployee id_employee获取
func (obj *_StockMvtMgr) WithIDEmployee(idEmployee int) Option {
	return optionFunc(func(o *options) { o.query["id_employee"] = idEmployee })
}

// WithEmployeeLastname employee_lastname获取
func (obj *_StockMvtMgr) WithEmployeeLastname(employeeLastname string) Option {
	return optionFunc(func(o *options) { o.query["employee_lastname"] = employeeLastname })
}

// WithEmployeeFirstname employee_firstname获取
func (obj *_StockMvtMgr) WithEmployeeFirstname(employeeFirstname string) Option {
	return optionFunc(func(o *options) { o.query["employee_firstname"] = employeeFirstname })
}

// WithPhysicalQuantity physical_quantity获取
func (obj *_StockMvtMgr) WithPhysicalQuantity(physicalQuantity int) Option {
	return optionFunc(func(o *options) { o.query["physical_quantity"] = physicalQuantity })
}

// WithDateAdd date_add获取
func (obj *_StockMvtMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

// WithSign sign获取
func (obj *_StockMvtMgr) WithSign(sign int16) Option {
	return optionFunc(func(o *options) { o.query["sign"] = sign })
}

// WithPriceTe price_te获取
func (obj *_StockMvtMgr) WithPriceTe(priceTe float64) Option {
	return optionFunc(func(o *options) { o.query["price_te"] = priceTe })
}

// WithLastWa last_wa获取
func (obj *_StockMvtMgr) WithLastWa(lastWa float64) Option {
	return optionFunc(func(o *options) { o.query["last_wa"] = lastWa })
}

// WithCurrentWa current_wa获取
func (obj *_StockMvtMgr) WithCurrentWa(currentWa float64) Option {
	return optionFunc(func(o *options) { o.query["current_wa"] = currentWa })
}

// WithReferer referer获取
func (obj *_StockMvtMgr) WithReferer(referer int64) Option {
	return optionFunc(func(o *options) { o.query["referer"] = referer })
}

// GetByOption 功能选项模式获取
func (obj *_StockMvtMgr) GetByOption(opts ...Option) (result StockMvt, err error) {
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
func (obj *_StockMvtMgr) GetByOptions(opts ...Option) (results []*StockMvt, err error) {
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

// GetFromIDStockMvt 通过id_stock_mvt获取内容
func (obj *_StockMvtMgr) GetFromIDStockMvt(idStockMvt int64) (result StockMvt, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_stock_mvt = ?", idStockMvt).Find(&result).Error

	return
}

// GetBatchFromIDStockMvt 批量唯一主键查找
func (obj *_StockMvtMgr) GetBatchFromIDStockMvt(idStockMvts []int64) (results []*StockMvt, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_stock_mvt IN (?)", idStockMvts).Find(&results).Error

	return
}

// GetFromIDStock 通过id_stock获取内容
func (obj *_StockMvtMgr) GetFromIDStock(idStock int) (results []*StockMvt, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_stock = ?", idStock).Find(&results).Error

	return
}

// GetBatchFromIDStock 批量唯一主键查找
func (obj *_StockMvtMgr) GetBatchFromIDStock(idStocks []int) (results []*StockMvt, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_stock IN (?)", idStocks).Find(&results).Error

	return
}

// GetFromIDOrder 通过id_order获取内容
func (obj *_StockMvtMgr) GetFromIDOrder(idOrder int) (results []*StockMvt, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order = ?", idOrder).Find(&results).Error

	return
}

// GetBatchFromIDOrder 批量唯一主键查找
func (obj *_StockMvtMgr) GetBatchFromIDOrder(idOrders []int) (results []*StockMvt, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order IN (?)", idOrders).Find(&results).Error

	return
}

// GetFromIDSupplyOrder 通过id_supply_order获取内容
func (obj *_StockMvtMgr) GetFromIDSupplyOrder(idSupplyOrder int) (results []*StockMvt, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supply_order = ?", idSupplyOrder).Find(&results).Error

	return
}

// GetBatchFromIDSupplyOrder 批量唯一主键查找
func (obj *_StockMvtMgr) GetBatchFromIDSupplyOrder(idSupplyOrders []int) (results []*StockMvt, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supply_order IN (?)", idSupplyOrders).Find(&results).Error

	return
}

// GetFromIDStockMvtReason 通过id_stock_mvt_reason获取内容
func (obj *_StockMvtMgr) GetFromIDStockMvtReason(idStockMvtReason int) (results []*StockMvt, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_stock_mvt_reason = ?", idStockMvtReason).Find(&results).Error

	return
}

// GetBatchFromIDStockMvtReason 批量唯一主键查找
func (obj *_StockMvtMgr) GetBatchFromIDStockMvtReason(idStockMvtReasons []int) (results []*StockMvt, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_stock_mvt_reason IN (?)", idStockMvtReasons).Find(&results).Error

	return
}

// GetFromIDEmployee 通过id_employee获取内容
func (obj *_StockMvtMgr) GetFromIDEmployee(idEmployee int) (results []*StockMvt, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee = ?", idEmployee).Find(&results).Error

	return
}

// GetBatchFromIDEmployee 批量唯一主键查找
func (obj *_StockMvtMgr) GetBatchFromIDEmployee(idEmployees []int) (results []*StockMvt, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee IN (?)", idEmployees).Find(&results).Error

	return
}

// GetFromEmployeeLastname 通过employee_lastname获取内容
func (obj *_StockMvtMgr) GetFromEmployeeLastname(employeeLastname string) (results []*StockMvt, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("employee_lastname = ?", employeeLastname).Find(&results).Error

	return
}

// GetBatchFromEmployeeLastname 批量唯一主键查找
func (obj *_StockMvtMgr) GetBatchFromEmployeeLastname(employeeLastnames []string) (results []*StockMvt, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("employee_lastname IN (?)", employeeLastnames).Find(&results).Error

	return
}

// GetFromEmployeeFirstname 通过employee_firstname获取内容
func (obj *_StockMvtMgr) GetFromEmployeeFirstname(employeeFirstname string) (results []*StockMvt, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("employee_firstname = ?", employeeFirstname).Find(&results).Error

	return
}

// GetBatchFromEmployeeFirstname 批量唯一主键查找
func (obj *_StockMvtMgr) GetBatchFromEmployeeFirstname(employeeFirstnames []string) (results []*StockMvt, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("employee_firstname IN (?)", employeeFirstnames).Find(&results).Error

	return
}

// GetFromPhysicalQuantity 通过physical_quantity获取内容
func (obj *_StockMvtMgr) GetFromPhysicalQuantity(physicalQuantity int) (results []*StockMvt, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("physical_quantity = ?", physicalQuantity).Find(&results).Error

	return
}

// GetBatchFromPhysicalQuantity 批量唯一主键查找
func (obj *_StockMvtMgr) GetBatchFromPhysicalQuantity(physicalQuantitys []int) (results []*StockMvt, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("physical_quantity IN (?)", physicalQuantitys).Find(&results).Error

	return
}

// GetFromDateAdd 通过date_add获取内容
func (obj *_StockMvtMgr) GetFromDateAdd(dateAdd time.Time) (results []*StockMvt, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error

	return
}

// GetBatchFromDateAdd 批量唯一主键查找
func (obj *_StockMvtMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*StockMvt, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error

	return
}

// GetFromSign 通过sign获取内容
func (obj *_StockMvtMgr) GetFromSign(sign int16) (results []*StockMvt, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("sign = ?", sign).Find(&results).Error

	return
}

// GetBatchFromSign 批量唯一主键查找
func (obj *_StockMvtMgr) GetBatchFromSign(signs []int16) (results []*StockMvt, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("sign IN (?)", signs).Find(&results).Error

	return
}

// GetFromPriceTe 通过price_te获取内容
func (obj *_StockMvtMgr) GetFromPriceTe(priceTe float64) (results []*StockMvt, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("price_te = ?", priceTe).Find(&results).Error

	return
}

// GetBatchFromPriceTe 批量唯一主键查找
func (obj *_StockMvtMgr) GetBatchFromPriceTe(priceTes []float64) (results []*StockMvt, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("price_te IN (?)", priceTes).Find(&results).Error

	return
}

// GetFromLastWa 通过last_wa获取内容
func (obj *_StockMvtMgr) GetFromLastWa(lastWa float64) (results []*StockMvt, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("last_wa = ?", lastWa).Find(&results).Error

	return
}

// GetBatchFromLastWa 批量唯一主键查找
func (obj *_StockMvtMgr) GetBatchFromLastWa(lastWas []float64) (results []*StockMvt, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("last_wa IN (?)", lastWas).Find(&results).Error

	return
}

// GetFromCurrentWa 通过current_wa获取内容
func (obj *_StockMvtMgr) GetFromCurrentWa(currentWa float64) (results []*StockMvt, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("current_wa = ?", currentWa).Find(&results).Error

	return
}

// GetBatchFromCurrentWa 批量唯一主键查找
func (obj *_StockMvtMgr) GetBatchFromCurrentWa(currentWas []float64) (results []*StockMvt, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("current_wa IN (?)", currentWas).Find(&results).Error

	return
}

// GetFromReferer 通过referer获取内容
func (obj *_StockMvtMgr) GetFromReferer(referer int64) (results []*StockMvt, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("referer = ?", referer).Find(&results).Error

	return
}

// GetBatchFromReferer 批量唯一主键查找
func (obj *_StockMvtMgr) GetBatchFromReferer(referers []int64) (results []*StockMvt, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("referer IN (?)", referers).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_StockMvtMgr) FetchByPrimaryKey(idStockMvt int64) (result StockMvt, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_stock_mvt = ?", idStockMvt).Find(&result).Error

	return
}

// FetchIndexByIDStock  获取多个内容
func (obj *_StockMvtMgr) FetchIndexByIDStock(idStock int) (results []*StockMvt, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_stock = ?", idStock).Find(&results).Error

	return
}

// FetchIndexByIDStockMvtReason  获取多个内容
func (obj *_StockMvtMgr) FetchIndexByIDStockMvtReason(idStockMvtReason int) (results []*StockMvt, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_stock_mvt_reason = ?", idStockMvtReason).Find(&results).Error

	return
}
