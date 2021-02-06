package	model	
import (	
"context"	
"gorm.io/gorm"	
"time"	
"fmt"	
)	

type _EgStoreMgr struct {
	*_BaseMgr
}

// EgStoreMgr open func
func EgStoreMgr(db *gorm.DB) *_EgStoreMgr {
	if db == nil {
		panic(fmt.Errorf("EgStoreMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgStoreMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_store"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgStoreMgr) GetTableName() string {
	return "eg_store"
}

// Get 获取
func (obj *_EgStoreMgr) Get() (result EgStore, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgStoreMgr) Gets() (results []*EgStore, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDStore id_store获取 
func (obj *_EgStoreMgr) WithIDStore(idStore uint32) Option {
	return optionFunc(func(o *options) { o.query["id_store"] = idStore })
}

// WithIDCountry id_country获取 
func (obj *_EgStoreMgr) WithIDCountry(idCountry uint32) Option {
	return optionFunc(func(o *options) { o.query["id_country"] = idCountry })
}

// WithIDState id_state获取 
func (obj *_EgStoreMgr) WithIDState(idState uint32) Option {
	return optionFunc(func(o *options) { o.query["id_state"] = idState })
}

// WithCity city获取 
func (obj *_EgStoreMgr) WithCity(city string) Option {
	return optionFunc(func(o *options) { o.query["city"] = city })
}

// WithPostcode postcode获取 
func (obj *_EgStoreMgr) WithPostcode(postcode string) Option {
	return optionFunc(func(o *options) { o.query["postcode"] = postcode })
}

// WithLatitude latitude获取 
func (obj *_EgStoreMgr) WithLatitude(latitude float64) Option {
	return optionFunc(func(o *options) { o.query["latitude"] = latitude })
}

// WithLongitude longitude获取 
func (obj *_EgStoreMgr) WithLongitude(longitude float64) Option {
	return optionFunc(func(o *options) { o.query["longitude"] = longitude })
}

// WithPhone phone获取 
func (obj *_EgStoreMgr) WithPhone(phone string) Option {
	return optionFunc(func(o *options) { o.query["phone"] = phone })
}

// WithFax fax获取 
func (obj *_EgStoreMgr) WithFax(fax string) Option {
	return optionFunc(func(o *options) { o.query["fax"] = fax })
}

// WithEmail email获取 
func (obj *_EgStoreMgr) WithEmail(email string) Option {
	return optionFunc(func(o *options) { o.query["email"] = email })
}

// WithActive active获取 
func (obj *_EgStoreMgr) WithActive(active bool) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}

// WithDateAdd date_add获取 
func (obj *_EgStoreMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

// WithDateUpd date_upd获取 
func (obj *_EgStoreMgr) WithDateUpd(dateUpd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_upd"] = dateUpd })
}


// GetByOption 功能选项模式获取
func (obj *_EgStoreMgr) GetByOption(opts ...Option) (result EgStore, err error) {
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
func (obj *_EgStoreMgr) GetByOptions(opts ...Option) (results []*EgStore, err error) {
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
func (obj *_EgStoreMgr)  GetFromIDStore(idStore uint32) (result EgStore, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_store = ?", idStore).Find(&result).Error
	
	return
}

// GetBatchFromIDStore 批量唯一主键查找 
func (obj *_EgStoreMgr) GetBatchFromIDStore(idStores []uint32) (results []*EgStore, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_store IN (?)", idStores).Find(&results).Error
	
	return
}
 
// GetFromIDCountry 通过id_country获取内容  
func (obj *_EgStoreMgr) GetFromIDCountry(idCountry uint32) (results []*EgStore, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_country = ?", idCountry).Find(&results).Error
	
	return
}

// GetBatchFromIDCountry 批量唯一主键查找 
func (obj *_EgStoreMgr) GetBatchFromIDCountry(idCountrys []uint32) (results []*EgStore, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_country IN (?)", idCountrys).Find(&results).Error
	
	return
}
 
// GetFromIDState 通过id_state获取内容  
func (obj *_EgStoreMgr) GetFromIDState(idState uint32) (results []*EgStore, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_state = ?", idState).Find(&results).Error
	
	return
}

// GetBatchFromIDState 批量唯一主键查找 
func (obj *_EgStoreMgr) GetBatchFromIDState(idStates []uint32) (results []*EgStore, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_state IN (?)", idStates).Find(&results).Error
	
	return
}
 
// GetFromCity 通过city获取内容  
func (obj *_EgStoreMgr) GetFromCity(city string) (results []*EgStore, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("city = ?", city).Find(&results).Error
	
	return
}

// GetBatchFromCity 批量唯一主键查找 
func (obj *_EgStoreMgr) GetBatchFromCity(citys []string) (results []*EgStore, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("city IN (?)", citys).Find(&results).Error
	
	return
}
 
// GetFromPostcode 通过postcode获取内容  
func (obj *_EgStoreMgr) GetFromPostcode(postcode string) (results []*EgStore, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("postcode = ?", postcode).Find(&results).Error
	
	return
}

// GetBatchFromPostcode 批量唯一主键查找 
func (obj *_EgStoreMgr) GetBatchFromPostcode(postcodes []string) (results []*EgStore, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("postcode IN (?)", postcodes).Find(&results).Error
	
	return
}
 
// GetFromLatitude 通过latitude获取内容  
func (obj *_EgStoreMgr) GetFromLatitude(latitude float64) (results []*EgStore, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("latitude = ?", latitude).Find(&results).Error
	
	return
}

// GetBatchFromLatitude 批量唯一主键查找 
func (obj *_EgStoreMgr) GetBatchFromLatitude(latitudes []float64) (results []*EgStore, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("latitude IN (?)", latitudes).Find(&results).Error
	
	return
}
 
// GetFromLongitude 通过longitude获取内容  
func (obj *_EgStoreMgr) GetFromLongitude(longitude float64) (results []*EgStore, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("longitude = ?", longitude).Find(&results).Error
	
	return
}

// GetBatchFromLongitude 批量唯一主键查找 
func (obj *_EgStoreMgr) GetBatchFromLongitude(longitudes []float64) (results []*EgStore, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("longitude IN (?)", longitudes).Find(&results).Error
	
	return
}
 
// GetFromPhone 通过phone获取内容  
func (obj *_EgStoreMgr) GetFromPhone(phone string) (results []*EgStore, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("phone = ?", phone).Find(&results).Error
	
	return
}

// GetBatchFromPhone 批量唯一主键查找 
func (obj *_EgStoreMgr) GetBatchFromPhone(phones []string) (results []*EgStore, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("phone IN (?)", phones).Find(&results).Error
	
	return
}
 
// GetFromFax 通过fax获取内容  
func (obj *_EgStoreMgr) GetFromFax(fax string) (results []*EgStore, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("fax = ?", fax).Find(&results).Error
	
	return
}

// GetBatchFromFax 批量唯一主键查找 
func (obj *_EgStoreMgr) GetBatchFromFax(faxs []string) (results []*EgStore, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("fax IN (?)", faxs).Find(&results).Error
	
	return
}
 
// GetFromEmail 通过email获取内容  
func (obj *_EgStoreMgr) GetFromEmail(email string) (results []*EgStore, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("email = ?", email).Find(&results).Error
	
	return
}

// GetBatchFromEmail 批量唯一主键查找 
func (obj *_EgStoreMgr) GetBatchFromEmail(emails []string) (results []*EgStore, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("email IN (?)", emails).Find(&results).Error
	
	return
}
 
// GetFromActive 通过active获取内容  
func (obj *_EgStoreMgr) GetFromActive(active bool) (results []*EgStore, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active = ?", active).Find(&results).Error
	
	return
}

// GetBatchFromActive 批量唯一主键查找 
func (obj *_EgStoreMgr) GetBatchFromActive(actives []bool) (results []*EgStore, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active IN (?)", actives).Find(&results).Error
	
	return
}
 
// GetFromDateAdd 通过date_add获取内容  
func (obj *_EgStoreMgr) GetFromDateAdd(dateAdd time.Time) (results []*EgStore, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error
	
	return
}

// GetBatchFromDateAdd 批量唯一主键查找 
func (obj *_EgStoreMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*EgStore, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error
	
	return
}
 
// GetFromDateUpd 通过date_upd获取内容  
func (obj *_EgStoreMgr) GetFromDateUpd(dateUpd time.Time) (results []*EgStore, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd = ?", dateUpd).Find(&results).Error
	
	return
}

// GetBatchFromDateUpd 批量唯一主键查找 
func (obj *_EgStoreMgr) GetBatchFromDateUpd(dateUpds []time.Time) (results []*EgStore, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd IN (?)", dateUpds).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgStoreMgr) FetchByPrimaryKey(idStore uint32 ) (result EgStore, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_store = ?", idStore).Find(&result).Error
	
	return
}
 

 

	

