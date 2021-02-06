package models

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _StockMvtMgr struct {
	*_BaseMgr
}

func StockMvtMgr(db *gorm.DB) *_StockMvtMgr {
	if db == nil {
		panic(fmt.Errorf("StockMvtMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_StockMvtMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_stock_mvt"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_StockMvtMgr) GetTableName() string {
	return "ps_stock_mvt"
}

func (obj *_StockMvtMgr) Get() (result StockMvt, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_StockMvtMgr) Gets() (results []*StockMvt, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_StockMvtMgr) WithIDStockMvt(idStockMvt int64) Option {
	return optionFunc(func(o *options) { o.query["id_stock_mvt"] = idStockMvt })
}

func (obj *_StockMvtMgr) WithIDStock(idStock int) Option {
	return optionFunc(func(o *options) { o.query["id_stock"] = idStock })
}

func (obj *_StockMvtMgr) WithIDOrder(idOrder int) Option {
	return optionFunc(func(o *options) { o.query["id_order"] = idOrder })
}

func (obj *_StockMvtMgr) WithIDSupplyOrder(idSupplyOrder int) Option {
	return optionFunc(func(o *options) { o.query["id_supply_order"] = idSupplyOrder })
}

func (obj *_StockMvtMgr) WithIDStockMvtReason(idStockMvtReason int) Option {
	return optionFunc(func(o *options) { o.query["id_stock_mvt_reason"] = idStockMvtReason })
}

func (obj *_StockMvtMgr) WithIDEmployee(idEmployee int) Option {
	return optionFunc(func(o *options) { o.query["id_employee"] = idEmployee })
}

func (obj *_StockMvtMgr) WithEmployeeLastname(employeeLastname string) Option {
	return optionFunc(func(o *options) { o.query["employee_lastname"] = employeeLastname })
}

func (obj *_StockMvtMgr) WithEmployeeFirstname(employeeFirstname string) Option {
	return optionFunc(func(o *options) { o.query["employee_firstname"] = employeeFirstname })
}

func (obj *_StockMvtMgr) WithPhysicalQuantity(physicalQuantity int) Option {
	return optionFunc(func(o *options) { o.query["physical_quantity"] = physicalQuantity })
}

func (obj *_StockMvtMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

func (obj *_StockMvtMgr) WithSign(sign int16) Option {
	return optionFunc(func(o *options) { o.query["sign"] = sign })
}

func (obj *_StockMvtMgr) WithPriceTe(priceTe float64) Option {
	return optionFunc(func(o *options) { o.query["price_te"] = priceTe })
}

func (obj *_StockMvtMgr) WithLastWa(lastWa float64) Option {
	return optionFunc(func(o *options) { o.query["last_wa"] = lastWa })
}

func (obj *_StockMvtMgr) WithCurrentWa(currentWa float64) Option {
	return optionFunc(func(o *options) { o.query["current_wa"] = currentWa })
}

func (obj *_StockMvtMgr) WithReferer(referer int64) Option {
	return optionFunc(func(o *options) { o.query["referer"] = referer })
}

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


func (obj *_StockMvtMgr) GetFromIDStockMvt(idStockMvt int64) (result StockMvt, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_stock_mvt = ?", idStockMvt).Find(&result).Error

	return
}

func (obj *_StockMvtMgr) GetBatchFromIDStockMvt(idStockMvts []int64) (results []*StockMvt, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_stock_mvt IN (?)", idStockMvts).Find(&results).Error

	return
}

func (obj *_StockMvtMgr) GetFromIDStock(idStock int) (results []*StockMvt, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_stock = ?", idStock).Find(&results).Error

	return
}

func (obj *_StockMvtMgr) GetBatchFromIDStock(idStocks []int) (results []*StockMvt, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_stock IN (?)", idStocks).Find(&results).Error

	return
}

func (obj *_StockMvtMgr) GetFromIDOrder(idOrder int) (results []*StockMvt, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order = ?", idOrder).Find(&results).Error

	return
}

func (obj *_StockMvtMgr) GetBatchFromIDOrder(idOrders []int) (results []*StockMvt, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order IN (?)", idOrders).Find(&results).Error

	return
}

func (obj *_StockMvtMgr) GetFromIDSupplyOrder(idSupplyOrder int) (results []*StockMvt, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supply_order = ?", idSupplyOrder).Find(&results).Error

	return
}

func (obj *_StockMvtMgr) GetBatchFromIDSupplyOrder(idSupplyOrders []int) (results []*StockMvt, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supply_order IN (?)", idSupplyOrders).Find(&results).Error

	return
}

func (obj *_StockMvtMgr) GetFromIDStockMvtReason(idStockMvtReason int) (results []*StockMvt, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_stock_mvt_reason = ?", idStockMvtReason).Find(&results).Error

	return
}

func (obj *_StockMvtMgr) GetBatchFromIDStockMvtReason(idStockMvtReasons []int) (results []*StockMvt, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_stock_mvt_reason IN (?)", idStockMvtReasons).Find(&results).Error

	return
}

func (obj *_StockMvtMgr) GetFromIDEmployee(idEmployee int) (results []*StockMvt, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee = ?", idEmployee).Find(&results).Error

	return
}

func (obj *_StockMvtMgr) GetBatchFromIDEmployee(idEmployees []int) (results []*StockMvt, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee IN (?)", idEmployees).Find(&results).Error

	return
}

func (obj *_StockMvtMgr) GetFromEmployeeLastname(employeeLastname string) (results []*StockMvt, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("employee_lastname = ?", employeeLastname).Find(&results).Error

	return
}

func (obj *_StockMvtMgr) GetBatchFromEmployeeLastname(employeeLastnames []string) (results []*StockMvt, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("employee_lastname IN (?)", employeeLastnames).Find(&results).Error

	return
}

func (obj *_StockMvtMgr) GetFromEmployeeFirstname(employeeFirstname string) (results []*StockMvt, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("employee_firstname = ?", employeeFirstname).Find(&results).Error

	return
}

func (obj *_StockMvtMgr) GetBatchFromEmployeeFirstname(employeeFirstnames []string) (results []*StockMvt, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("employee_firstname IN (?)", employeeFirstnames).Find(&results).Error

	return
}

func (obj *_StockMvtMgr) GetFromPhysicalQuantity(physicalQuantity int) (results []*StockMvt, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("physical_quantity = ?", physicalQuantity).Find(&results).Error

	return
}

func (obj *_StockMvtMgr) GetBatchFromPhysicalQuantity(physicalQuantitys []int) (results []*StockMvt, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("physical_quantity IN (?)", physicalQuantitys).Find(&results).Error

	return
}

func (obj *_StockMvtMgr) GetFromDateAdd(dateAdd time.Time) (results []*StockMvt, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error

	return
}

func (obj *_StockMvtMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*StockMvt, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error

	return
}

func (obj *_StockMvtMgr) GetFromSign(sign int16) (results []*StockMvt, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("sign = ?", sign).Find(&results).Error

	return
}

func (obj *_StockMvtMgr) GetBatchFromSign(signs []int16) (results []*StockMvt, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("sign IN (?)", signs).Find(&results).Error

	return
}

func (obj *_StockMvtMgr) GetFromPriceTe(priceTe float64) (results []*StockMvt, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("price_te = ?", priceTe).Find(&results).Error

	return
}

func (obj *_StockMvtMgr) GetBatchFromPriceTe(priceTes []float64) (results []*StockMvt, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("price_te IN (?)", priceTes).Find(&results).Error

	return
}

func (obj *_StockMvtMgr) GetFromLastWa(lastWa float64) (results []*StockMvt, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("last_wa = ?", lastWa).Find(&results).Error

	return
}

func (obj *_StockMvtMgr) GetBatchFromLastWa(lastWas []float64) (results []*StockMvt, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("last_wa IN (?)", lastWas).Find(&results).Error

	return
}

func (obj *_StockMvtMgr) GetFromCurrentWa(currentWa float64) (results []*StockMvt, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("current_wa = ?", currentWa).Find(&results).Error

	return
}

func (obj *_StockMvtMgr) GetBatchFromCurrentWa(currentWas []float64) (results []*StockMvt, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("current_wa IN (?)", currentWas).Find(&results).Error

	return
}

func (obj *_StockMvtMgr) GetFromReferer(referer int64) (results []*StockMvt, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("referer = ?", referer).Find(&results).Error

	return
}

func (obj *_StockMvtMgr) GetBatchFromReferer(referers []int64) (results []*StockMvt, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("referer IN (?)", referers).Find(&results).Error

	return
}


func (obj *_StockMvtMgr) FetchByPrimaryKey(idStockMvt int64) (result StockMvt, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_stock_mvt = ?", idStockMvt).Find(&result).Error

	return
}

func (obj *_StockMvtMgr) FetchIndexByIDStock(idStock int) (results []*StockMvt, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_stock = ?", idStock).Find(&results).Error

	return
}

func (obj *_StockMvtMgr) FetchIndexByIDStockMvtReason(idStockMvtReason int) (results []*StockMvt, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_stock_mvt_reason = ?", idStockMvtReason).Find(&results).Error

	return
}
