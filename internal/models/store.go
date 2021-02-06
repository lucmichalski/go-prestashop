package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _StoreMgr struct {
	*_BaseMgr
}

// StoreMgr open func
func StoreMgr(db *gorm.DB) *_StoreMgr {
	if db == nil {
		panic(fmt.Errorf("StoreMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_StoreMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_store"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_StoreMgr) GetTableName() string {
	return "ps_store"
}

// Get 获取
func (obj *_StoreMgr) Get() (result Store, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_StoreMgr) Gets() (results []*Store, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDStore id_store获取
func (obj *_StoreMgr) WithIDStore(idStore uint32) Option {
	return optionFunc(func(o *options) { o.query["id_store"] = idStore })
}

// WithIDCountry id_country获取
func (obj *_StoreMgr) WithIDCountry(idCountry uint32) Option {
	return optionFunc(func(o *options) { o.query["id_country"] = idCountry })
}

// WithIDState id_state获取
func (obj *_StoreMgr) WithIDState(idState uint32) Option {
	return optionFunc(func(o *options) { o.query["id_state"] = idState })
}

// WithCity city获取
func (obj *_StoreMgr) WithCity(city string) Option {
	return optionFunc(func(o *options) { o.query["city"] = city })
}

// WithPostcode postcode获取
func (obj *_StoreMgr) WithPostcode(postcode string) Option {
	return optionFunc(func(o *options) { o.query["postcode"] = postcode })
}

// WithLatitude latitude获取
func (obj *_StoreMgr) WithLatitude(latitude float64) Option {
	return optionFunc(func(o *options) { o.query["latitude"] = latitude })
}

// WithLongitude longitude获取
func (obj *_StoreMgr) WithLongitude(longitude float64) Option {
	return optionFunc(func(o *options) { o.query["longitude"] = longitude })
}

// WithPhone phone获取
func (obj *_StoreMgr) WithPhone(phone string) Option {
	return optionFunc(func(o *options) { o.query["phone"] = phone })
}

// WithFax fax获取
func (obj *_StoreMgr) WithFax(fax string) Option {
	return optionFunc(func(o *options) { o.query["fax"] = fax })
}

// WithEmail email获取
func (obj *_StoreMgr) WithEmail(email string) Option {
	return optionFunc(func(o *options) { o.query["email"] = email })
}

// WithActive active获取
func (obj *_StoreMgr) WithActive(active bool) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}

// WithDateAdd date_add获取
func (obj *_StoreMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

// WithDateUpd date_upd获取
func (obj *_StoreMgr) WithDateUpd(dateUpd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_upd"] = dateUpd })
}

// GetByOption 功能选项模式获取
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

// GetByOptions 批量功能选项模式获取
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

//////////////////////////enume case ////////////////////////////////////////////

// GetFromIDStore 通过id_store获取内容
func (obj *_StoreMgr) GetFromIDStore(idStore uint32) (result Store, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_store = ?", idStore).Find(&result).Error

	return
}

// GetBatchFromIDStore 批量唯一主键查找
func (obj *_StoreMgr) GetBatchFromIDStore(idStores []uint32) (results []*Store, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_store IN (?)", idStores).Find(&results).Error

	return
}

// GetFromIDCountry 通过id_country获取内容
func (obj *_StoreMgr) GetFromIDCountry(idCountry uint32) (results []*Store, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_country = ?", idCountry).Find(&results).Error

	return
}

// GetBatchFromIDCountry 批量唯一主键查找
func (obj *_StoreMgr) GetBatchFromIDCountry(idCountrys []uint32) (results []*Store, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_country IN (?)", idCountrys).Find(&results).Error

	return
}

// GetFromIDState 通过id_state获取内容
func (obj *_StoreMgr) GetFromIDState(idState uint32) (results []*Store, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_state = ?", idState).Find(&results).Error

	return
}

// GetBatchFromIDState 批量唯一主键查找
func (obj *_StoreMgr) GetBatchFromIDState(idStates []uint32) (results []*Store, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_state IN (?)", idStates).Find(&results).Error

	return
}

// GetFromCity 通过city获取内容
func (obj *_StoreMgr) GetFromCity(city string) (results []*Store, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("city = ?", city).Find(&results).Error

	return
}

// GetBatchFromCity 批量唯一主键查找
func (obj *_StoreMgr) GetBatchFromCity(citys []string) (results []*Store, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("city IN (?)", citys).Find(&results).Error

	return
}

// GetFromPostcode 通过postcode获取内容
func (obj *_StoreMgr) GetFromPostcode(postcode string) (results []*Store, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("postcode = ?", postcode).Find(&results).Error

	return
}

// GetBatchFromPostcode 批量唯一主键查找
func (obj *_StoreMgr) GetBatchFromPostcode(postcodes []string) (results []*Store, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("postcode IN (?)", postcodes).Find(&results).Error

	return
}

// GetFromLatitude 通过latitude获取内容
func (obj *_StoreMgr) GetFromLatitude(latitude float64) (results []*Store, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("latitude = ?", latitude).Find(&results).Error

	return
}

// GetBatchFromLatitude 批量唯一主键查找
func (obj *_StoreMgr) GetBatchFromLatitude(latitudes []float64) (results []*Store, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("latitude IN (?)", latitudes).Find(&results).Error

	return
}

// GetFromLongitude 通过longitude获取内容
func (obj *_StoreMgr) GetFromLongitude(longitude float64) (results []*Store, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("longitude = ?", longitude).Find(&results).Error

	return
}

// GetBatchFromLongitude 批量唯一主键查找
func (obj *_StoreMgr) GetBatchFromLongitude(longitudes []float64) (results []*Store, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("longitude IN (?)", longitudes).Find(&results).Error

	return
}

// GetFromPhone 通过phone获取内容
func (obj *_StoreMgr) GetFromPhone(phone string) (results []*Store, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("phone = ?", phone).Find(&results).Error

	return
}

// GetBatchFromPhone 批量唯一主键查找
func (obj *_StoreMgr) GetBatchFromPhone(phones []string) (results []*Store, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("phone IN (?)", phones).Find(&results).Error

	return
}

// GetFromFax 通过fax获取内容
func (obj *_StoreMgr) GetFromFax(fax string) (results []*Store, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("fax = ?", fax).Find(&results).Error

	return
}

// GetBatchFromFax 批量唯一主键查找
func (obj *_StoreMgr) GetBatchFromFax(faxs []string) (results []*Store, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("fax IN (?)", faxs).Find(&results).Error

	return
}

// GetFromEmail 通过email获取内容
func (obj *_StoreMgr) GetFromEmail(email string) (results []*Store, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("email = ?", email).Find(&results).Error

	return
}

// GetBatchFromEmail 批量唯一主键查找
func (obj *_StoreMgr) GetBatchFromEmail(emails []string) (results []*Store, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("email IN (?)", emails).Find(&results).Error

	return
}

// GetFromActive 通过active获取内容
func (obj *_StoreMgr) GetFromActive(active bool) (results []*Store, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active = ?", active).Find(&results).Error

	return
}

// GetBatchFromActive 批量唯一主键查找
func (obj *_StoreMgr) GetBatchFromActive(actives []bool) (results []*Store, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active IN (?)", actives).Find(&results).Error

	return
}

// GetFromDateAdd 通过date_add获取内容
func (obj *_StoreMgr) GetFromDateAdd(dateAdd time.Time) (results []*Store, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error

	return
}

// GetBatchFromDateAdd 批量唯一主键查找
func (obj *_StoreMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*Store, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error

	return
}

// GetFromDateUpd 通过date_upd获取内容
func (obj *_StoreMgr) GetFromDateUpd(dateUpd time.Time) (results []*Store, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd = ?", dateUpd).Find(&results).Error

	return
}

// GetBatchFromDateUpd 批量唯一主键查找
func (obj *_StoreMgr) GetBatchFromDateUpd(dateUpds []time.Time) (results []*Store, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd IN (?)", dateUpds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_StoreMgr) FetchByPrimaryKey(idStore uint32) (result Store, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_store = ?", idStore).Find(&result).Error

	return
}
