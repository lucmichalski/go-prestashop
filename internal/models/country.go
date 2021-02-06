package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _CountryMgr struct {
	*_BaseMgr
}

func CountryMgr(db *gorm.DB) *_CountryMgr {
	if db == nil {
		panic(fmt.Errorf("CountryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CountryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_country"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_CountryMgr) GetTableName() string {
	return "ps_country"
}

func (obj *_CountryMgr) Get() (result Country, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_CountryMgr) Gets() (results []*Country, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_CountryMgr) WithIDCountry(idCountry uint32) Option {
	return optionFunc(func(o *options) { o.query["id_country"] = idCountry })
}

func (obj *_CountryMgr) WithIDZone(idZone uint32) Option {
	return optionFunc(func(o *options) { o.query["id_zone"] = idZone })
}

func (obj *_CountryMgr) WithIDCurrency(idCurrency uint32) Option {
	return optionFunc(func(o *options) { o.query["id_currency"] = idCurrency })
}

func (obj *_CountryMgr) WithIsoCode(isoCode string) Option {
	return optionFunc(func(o *options) { o.query["iso_code"] = isoCode })
}

func (obj *_CountryMgr) WithCallPrefix(callPrefix int) Option {
	return optionFunc(func(o *options) { o.query["call_prefix"] = callPrefix })
}

func (obj *_CountryMgr) WithActive(active bool) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}

func (obj *_CountryMgr) WithContainsStates(containsStates bool) Option {
	return optionFunc(func(o *options) { o.query["contains_states"] = containsStates })
}

func (obj *_CountryMgr) WithNeedIDentificationNumber(needIDentificationNumber bool) Option {
	return optionFunc(func(o *options) { o.query["need_identification_number"] = needIDentificationNumber })
}

func (obj *_CountryMgr) WithNeedZipCode(needZipCode bool) Option {
	return optionFunc(func(o *options) { o.query["need_zip_code"] = needZipCode })
}

func (obj *_CountryMgr) WithZipCodeFormat(zipCodeFormat string) Option {
	return optionFunc(func(o *options) { o.query["zip_code_format"] = zipCodeFormat })
}

func (obj *_CountryMgr) WithDisplayTaxLabel(displayTaxLabel bool) Option {
	return optionFunc(func(o *options) { o.query["display_tax_label"] = displayTaxLabel })
}

func (obj *_CountryMgr) GetByOption(opts ...Option) (result Country, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_CountryMgr) GetByOptions(opts ...Option) (results []*Country, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_CountryMgr) GetFromIDCountry(idCountry uint32) (result Country, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_country = ?", idCountry).Find(&result).Error

	return
}

func (obj *_CountryMgr) GetBatchFromIDCountry(idCountrys []uint32) (results []*Country, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_country IN (?)", idCountrys).Find(&results).Error

	return
}

func (obj *_CountryMgr) GetFromIDZone(idZone uint32) (results []*Country, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_zone = ?", idZone).Find(&results).Error

	return
}

func (obj *_CountryMgr) GetBatchFromIDZone(idZones []uint32) (results []*Country, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_zone IN (?)", idZones).Find(&results).Error

	return
}

func (obj *_CountryMgr) GetFromIDCurrency(idCurrency uint32) (results []*Country, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_currency = ?", idCurrency).Find(&results).Error

	return
}

func (obj *_CountryMgr) GetBatchFromIDCurrency(idCurrencys []uint32) (results []*Country, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_currency IN (?)", idCurrencys).Find(&results).Error

	return
}

func (obj *_CountryMgr) GetFromIsoCode(isoCode string) (results []*Country, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("iso_code = ?", isoCode).Find(&results).Error

	return
}

func (obj *_CountryMgr) GetBatchFromIsoCode(isoCodes []string) (results []*Country, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("iso_code IN (?)", isoCodes).Find(&results).Error

	return
}

func (obj *_CountryMgr) GetFromCallPrefix(callPrefix int) (results []*Country, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("call_prefix = ?", callPrefix).Find(&results).Error

	return
}

func (obj *_CountryMgr) GetBatchFromCallPrefix(callPrefixs []int) (results []*Country, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("call_prefix IN (?)", callPrefixs).Find(&results).Error

	return
}

func (obj *_CountryMgr) GetFromActive(active bool) (results []*Country, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active = ?", active).Find(&results).Error

	return
}

func (obj *_CountryMgr) GetBatchFromActive(actives []bool) (results []*Country, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active IN (?)", actives).Find(&results).Error

	return
}

func (obj *_CountryMgr) GetFromContainsStates(containsStates bool) (results []*Country, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("contains_states = ?", containsStates).Find(&results).Error

	return
}

func (obj *_CountryMgr) GetBatchFromContainsStates(containsStatess []bool) (results []*Country, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("contains_states IN (?)", containsStatess).Find(&results).Error

	return
}

func (obj *_CountryMgr) GetFromNeedIDentificationNumber(needIDentificationNumber bool) (results []*Country, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("need_identification_number = ?", needIDentificationNumber).Find(&results).Error

	return
}

func (obj *_CountryMgr) GetBatchFromNeedIDentificationNumber(needIDentificationNumbers []bool) (results []*Country, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("need_identification_number IN (?)", needIDentificationNumbers).Find(&results).Error

	return
}

func (obj *_CountryMgr) GetFromNeedZipCode(needZipCode bool) (results []*Country, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("need_zip_code = ?", needZipCode).Find(&results).Error

	return
}

func (obj *_CountryMgr) GetBatchFromNeedZipCode(needZipCodes []bool) (results []*Country, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("need_zip_code IN (?)", needZipCodes).Find(&results).Error

	return
}

func (obj *_CountryMgr) GetFromZipCodeFormat(zipCodeFormat string) (results []*Country, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("zip_code_format = ?", zipCodeFormat).Find(&results).Error

	return
}

func (obj *_CountryMgr) GetBatchFromZipCodeFormat(zipCodeFormats []string) (results []*Country, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("zip_code_format IN (?)", zipCodeFormats).Find(&results).Error

	return
}

func (obj *_CountryMgr) GetFromDisplayTaxLabel(displayTaxLabel bool) (results []*Country, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("display_tax_label = ?", displayTaxLabel).Find(&results).Error

	return
}

func (obj *_CountryMgr) GetBatchFromDisplayTaxLabel(displayTaxLabels []bool) (results []*Country, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("display_tax_label IN (?)", displayTaxLabels).Find(&results).Error

	return
}


func (obj *_CountryMgr) FetchByPrimaryKey(idCountry uint32) (result Country, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_country = ?", idCountry).Find(&result).Error

	return
}

func (obj *_CountryMgr) FetchIndexByCountry(idZone uint32) (results []*Country, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_zone = ?", idZone).Find(&results).Error

	return
}

func (obj *_CountryMgr) FetchIndexByCountryIsoCode(isoCode string) (results []*Country, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("iso_code = ?", isoCode).Find(&results).Error

	return
}
