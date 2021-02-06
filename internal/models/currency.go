package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _CurrencyMgr struct {
	*_BaseMgr
}

func CurrencyMgr(db *gorm.DB) *_CurrencyMgr {
	if db == nil {
		panic(fmt.Errorf("CurrencyMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CurrencyMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_currency"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_CurrencyMgr) GetTableName() string {
	return "ps_currency"
}

func (obj *_CurrencyMgr) Get() (result Currency, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_CurrencyMgr) Gets() (results []*Currency, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_CurrencyMgr) WithIDCurrency(idCurrency uint32) Option {
	return optionFunc(func(o *options) { o.query["id_currency"] = idCurrency })
}

func (obj *_CurrencyMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

func (obj *_CurrencyMgr) WithIsoCode(isoCode string) Option {
	return optionFunc(func(o *options) { o.query["iso_code"] = isoCode })
}

func (obj *_CurrencyMgr) WithNumericIsoCode(numericIsoCode string) Option {
	return optionFunc(func(o *options) { o.query["numeric_iso_code"] = numericIsoCode })
}

func (obj *_CurrencyMgr) WithPrecision(precision int) Option {
	return optionFunc(func(o *options) { o.query["precision"] = precision })
}

func (obj *_CurrencyMgr) WithConversionRate(conversionRate float64) Option {
	return optionFunc(func(o *options) { o.query["conversion_rate"] = conversionRate })
}

func (obj *_CurrencyMgr) WithDeleted(deleted bool) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

func (obj *_CurrencyMgr) WithActive(active bool) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}

func (obj *_CurrencyMgr) WithUnofficial(unofficial bool) Option {
	return optionFunc(func(o *options) { o.query["unofficial"] = unofficial })
}

func (obj *_CurrencyMgr) WithModified(modified bool) Option {
	return optionFunc(func(o *options) { o.query["modified"] = modified })
}

func (obj *_CurrencyMgr) GetByOption(opts ...Option) (result Currency, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_CurrencyMgr) GetByOptions(opts ...Option) (results []*Currency, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_CurrencyMgr) GetFromIDCurrency(idCurrency uint32) (result Currency, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_currency = ?", idCurrency).Find(&result).Error

	return
}

func (obj *_CurrencyMgr) GetBatchFromIDCurrency(idCurrencys []uint32) (results []*Currency, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_currency IN (?)", idCurrencys).Find(&results).Error

	return
}

func (obj *_CurrencyMgr) GetFromName(name string) (results []*Currency, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

func (obj *_CurrencyMgr) GetBatchFromName(names []string) (results []*Currency, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

func (obj *_CurrencyMgr) GetFromIsoCode(isoCode string) (results []*Currency, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("iso_code = ?", isoCode).Find(&results).Error

	return
}

func (obj *_CurrencyMgr) GetBatchFromIsoCode(isoCodes []string) (results []*Currency, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("iso_code IN (?)", isoCodes).Find(&results).Error

	return
}

func (obj *_CurrencyMgr) GetFromNumericIsoCode(numericIsoCode string) (results []*Currency, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("numeric_iso_code = ?", numericIsoCode).Find(&results).Error

	return
}

func (obj *_CurrencyMgr) GetBatchFromNumericIsoCode(numericIsoCodes []string) (results []*Currency, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("numeric_iso_code IN (?)", numericIsoCodes).Find(&results).Error

	return
}

func (obj *_CurrencyMgr) GetFromPrecision(precision int) (results []*Currency, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("precision = ?", precision).Find(&results).Error

	return
}

func (obj *_CurrencyMgr) GetBatchFromPrecision(precisions []int) (results []*Currency, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("precision IN (?)", precisions).Find(&results).Error

	return
}

func (obj *_CurrencyMgr) GetFromConversionRate(conversionRate float64) (results []*Currency, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("conversion_rate = ?", conversionRate).Find(&results).Error

	return
}

func (obj *_CurrencyMgr) GetBatchFromConversionRate(conversionRates []float64) (results []*Currency, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("conversion_rate IN (?)", conversionRates).Find(&results).Error

	return
}

func (obj *_CurrencyMgr) GetFromDeleted(deleted bool) (results []*Currency, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("deleted = ?", deleted).Find(&results).Error

	return
}

func (obj *_CurrencyMgr) GetBatchFromDeleted(deleteds []bool) (results []*Currency, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("deleted IN (?)", deleteds).Find(&results).Error

	return
}

func (obj *_CurrencyMgr) GetFromActive(active bool) (results []*Currency, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active = ?", active).Find(&results).Error

	return
}

func (obj *_CurrencyMgr) GetBatchFromActive(actives []bool) (results []*Currency, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active IN (?)", actives).Find(&results).Error

	return
}

func (obj *_CurrencyMgr) GetFromUnofficial(unofficial bool) (results []*Currency, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("unofficial = ?", unofficial).Find(&results).Error

	return
}

func (obj *_CurrencyMgr) GetBatchFromUnofficial(unofficials []bool) (results []*Currency, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("unofficial IN (?)", unofficials).Find(&results).Error

	return
}

func (obj *_CurrencyMgr) GetFromModified(modified bool) (results []*Currency, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("modified = ?", modified).Find(&results).Error

	return
}

func (obj *_CurrencyMgr) GetBatchFromModified(modifieds []bool) (results []*Currency, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("modified IN (?)", modifieds).Find(&results).Error

	return
}


func (obj *_CurrencyMgr) FetchByPrimaryKey(idCurrency uint32) (result Currency, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_currency = ?", idCurrency).Find(&result).Error

	return
}

func (obj *_CurrencyMgr) FetchIndexByCurrencyIsoCode(isoCode string) (results []*Currency, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("iso_code = ?", isoCode).Find(&results).Error

	return
}
