package models

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _StoreMgr struct {
	*_BaseMgr
}

func StoreMgr(db *gorm.DB) *_StoreMgr {
	if db == nil {
		panic(fmt.Errorf("StoreMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_StoreMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_store"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_StoreMgr) GetTableName() string {
	return "ps_store"
}

func (obj *_StoreMgr) Get() (result Store, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_StoreMgr) Gets() (results []*Store, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_StoreMgr) WithIDStore(idStore uint32) Option {
	return optionFunc(func(o *options) { o.query["id_store"] = idStore })
}

func (obj *_StoreMgr) WithIDCountry(idCountry uint32) Option {
	return optionFunc(func(o *options) { o.query["id_country"] = idCountry })
}

func (obj *_StoreMgr) WithIDState(idState uint32) Option {
	return optionFunc(func(o *options) { o.query["id_state"] = idState })
}

func (obj *_StoreMgr) WithCity(city string) Option {
	return optionFunc(func(o *options) { o.query["city"] = city })
}

func (obj *_StoreMgr) WithPostcode(postcode string) Option {
	return optionFunc(func(o *options) { o.query["postcode"] = postcode })
}

func (obj *_StoreMgr) WithLatitude(latitude float64) Option {
	return optionFunc(func(o *options) { o.query["latitude"] = latitude })
}

func (obj *_StoreMgr) WithLongitude(longitude float64) Option {
	return optionFunc(func(o *options) { o.query["longitude"] = longitude })
}

func (obj *_StoreMgr) WithPhone(phone string) Option {
	return optionFunc(func(o *options) { o.query["phone"] = phone })
}

func (obj *_StoreMgr) WithFax(fax string) Option {
	return optionFunc(func(o *options) { o.query["fax"] = fax })
}

func (obj *_StoreMgr) WithEmail(email string) Option {
	return optionFunc(func(o *options) { o.query["email"] = email })
}

func (obj *_StoreMgr) WithActive(active bool) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}

func (obj *_StoreMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

func (obj *_StoreMgr) WithDateUpd(dateUpd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_upd"] = dateUpd })
}

func (obj *_StoreMgr) GetByOption(opts ...Option) (result Store, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_StoreMgr) GetByOptions(opts ...Option) (results []*Store, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_StoreMgr) GetFromIDStore(idStore uint32) (result Store, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_store = ?", idStore).Find(&result).Error

	return
}

func (obj *_StoreMgr) GetBatchFromIDStore(idStores []uint32) (results []*Store, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_store IN (?)", idStores).Find(&results).Error

	return
}

func (obj *_StoreMgr) GetFromIDCountry(idCountry uint32) (results []*Store, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_country = ?", idCountry).Find(&results).Error

	return
}

func (obj *_StoreMgr) GetBatchFromIDCountry(idCountrys []uint32) (results []*Store, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_country IN (?)", idCountrys).Find(&results).Error

	return
}

func (obj *_StoreMgr) GetFromIDState(idState uint32) (results []*Store, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_state = ?", idState).Find(&results).Error

	return
}

func (obj *_StoreMgr) GetBatchFromIDState(idStates []uint32) (results []*Store, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_state IN (?)", idStates).Find(&results).Error

	return
}

func (obj *_StoreMgr) GetFromCity(city string) (results []*Store, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("city = ?", city).Find(&results).Error

	return
}

func (obj *_StoreMgr) GetBatchFromCity(citys []string) (results []*Store, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("city IN (?)", citys).Find(&results).Error

	return
}

func (obj *_StoreMgr) GetFromPostcode(postcode string) (results []*Store, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("postcode = ?", postcode).Find(&results).Error

	return
}

func (obj *_StoreMgr) GetBatchFromPostcode(postcodes []string) (results []*Store, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("postcode IN (?)", postcodes).Find(&results).Error

	return
}

func (obj *_StoreMgr) GetFromLatitude(latitude float64) (results []*Store, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("latitude = ?", latitude).Find(&results).Error

	return
}

func (obj *_StoreMgr) GetBatchFromLatitude(latitudes []float64) (results []*Store, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("latitude IN (?)", latitudes).Find(&results).Error

	return
}

func (obj *_StoreMgr) GetFromLongitude(longitude float64) (results []*Store, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("longitude = ?", longitude).Find(&results).Error

	return
}

func (obj *_StoreMgr) GetBatchFromLongitude(longitudes []float64) (results []*Store, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("longitude IN (?)", longitudes).Find(&results).Error

	return
}

func (obj *_StoreMgr) GetFromPhone(phone string) (results []*Store, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("phone = ?", phone).Find(&results).Error

	return
}

func (obj *_StoreMgr) GetBatchFromPhone(phones []string) (results []*Store, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("phone IN (?)", phones).Find(&results).Error

	return
}

func (obj *_StoreMgr) GetFromFax(fax string) (results []*Store, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("fax = ?", fax).Find(&results).Error

	return
}

func (obj *_StoreMgr) GetBatchFromFax(faxs []string) (results []*Store, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("fax IN (?)", faxs).Find(&results).Error

	return
}

func (obj *_StoreMgr) GetFromEmail(email string) (results []*Store, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("email = ?", email).Find(&results).Error

	return
}

func (obj *_StoreMgr) GetBatchFromEmail(emails []string) (results []*Store, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("email IN (?)", emails).Find(&results).Error

	return
}

func (obj *_StoreMgr) GetFromActive(active bool) (results []*Store, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active = ?", active).Find(&results).Error

	return
}

func (obj *_StoreMgr) GetBatchFromActive(actives []bool) (results []*Store, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active IN (?)", actives).Find(&results).Error

	return
}

func (obj *_StoreMgr) GetFromDateAdd(dateAdd time.Time) (results []*Store, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error

	return
}

func (obj *_StoreMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*Store, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error

	return
}

func (obj *_StoreMgr) GetFromDateUpd(dateUpd time.Time) (results []*Store, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd = ?", dateUpd).Find(&results).Error

	return
}

func (obj *_StoreMgr) GetBatchFromDateUpd(dateUpds []time.Time) (results []*Store, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd IN (?)", dateUpds).Find(&results).Error

	return
}


func (obj *_StoreMgr) FetchByPrimaryKey(idStore uint32) (result Store, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_store = ?", idStore).Find(&result).Error

	return
}
