package	model	
import (	
"context"	
"gorm.io/gorm"	
"time"	
"fmt"	
)	

type _EgAddressMgr struct {
	*_BaseMgr
}

// EgAddressMgr open func
func EgAddressMgr(db *gorm.DB) *_EgAddressMgr {
	if db == nil {
		panic(fmt.Errorf("EgAddressMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgAddressMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_address"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgAddressMgr) GetTableName() string {
	return "eg_address"
}

// Get 获取
func (obj *_EgAddressMgr) Get() (result EgAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgAddressMgr) Gets() (results []*EgAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDAddress id_address获取 
func (obj *_EgAddressMgr) WithIDAddress(idAddress uint32) Option {
	return optionFunc(func(o *options) { o.query["id_address"] = idAddress })
}

// WithIDCountry id_country获取 
func (obj *_EgAddressMgr) WithIDCountry(idCountry uint32) Option {
	return optionFunc(func(o *options) { o.query["id_country"] = idCountry })
}

// WithIDState id_state获取 
func (obj *_EgAddressMgr) WithIDState(idState uint32) Option {
	return optionFunc(func(o *options) { o.query["id_state"] = idState })
}

// WithIDCustomer id_customer获取 
func (obj *_EgAddressMgr) WithIDCustomer(idCustomer uint32) Option {
	return optionFunc(func(o *options) { o.query["id_customer"] = idCustomer })
}

// WithIDManufacturer id_manufacturer获取 
func (obj *_EgAddressMgr) WithIDManufacturer(idManufacturer uint32) Option {
	return optionFunc(func(o *options) { o.query["id_manufacturer"] = idManufacturer })
}

// WithIDSupplier id_supplier获取 
func (obj *_EgAddressMgr) WithIDSupplier(idSupplier uint32) Option {
	return optionFunc(func(o *options) { o.query["id_supplier"] = idSupplier })
}

// WithIDWarehouse id_warehouse获取 
func (obj *_EgAddressMgr) WithIDWarehouse(idWarehouse uint32) Option {
	return optionFunc(func(o *options) { o.query["id_warehouse"] = idWarehouse })
}

// WithAlias alias获取 
func (obj *_EgAddressMgr) WithAlias(alias string) Option {
	return optionFunc(func(o *options) { o.query["alias"] = alias })
}

// WithCompany company获取 
func (obj *_EgAddressMgr) WithCompany(company string) Option {
	return optionFunc(func(o *options) { o.query["company"] = company })
}

// WithLastname lastname获取 
func (obj *_EgAddressMgr) WithLastname(lastname string) Option {
	return optionFunc(func(o *options) { o.query["lastname"] = lastname })
}

// WithFirstname firstname获取 
func (obj *_EgAddressMgr) WithFirstname(firstname string) Option {
	return optionFunc(func(o *options) { o.query["firstname"] = firstname })
}

// WithAddress1 address1获取 
func (obj *_EgAddressMgr) WithAddress1(address1 string) Option {
	return optionFunc(func(o *options) { o.query["address1"] = address1 })
}

// WithAddress2 address2获取 
func (obj *_EgAddressMgr) WithAddress2(address2 string) Option {
	return optionFunc(func(o *options) { o.query["address2"] = address2 })
}

// WithPostcode postcode获取 
func (obj *_EgAddressMgr) WithPostcode(postcode string) Option {
	return optionFunc(func(o *options) { o.query["postcode"] = postcode })
}

// WithCity city获取 
func (obj *_EgAddressMgr) WithCity(city string) Option {
	return optionFunc(func(o *options) { o.query["city"] = city })
}

// WithOther other获取 
func (obj *_EgAddressMgr) WithOther(other string) Option {
	return optionFunc(func(o *options) { o.query["other"] = other })
}

// WithPhone phone获取 
func (obj *_EgAddressMgr) WithPhone(phone string) Option {
	return optionFunc(func(o *options) { o.query["phone"] = phone })
}

// WithPhoneMobile phone_mobile获取 
func (obj *_EgAddressMgr) WithPhoneMobile(phoneMobile string) Option {
	return optionFunc(func(o *options) { o.query["phone_mobile"] = phoneMobile })
}

// WithVatNumber vat_number获取 
func (obj *_EgAddressMgr) WithVatNumber(vatNumber string) Option {
	return optionFunc(func(o *options) { o.query["vat_number"] = vatNumber })
}

// WithDni dni获取 
func (obj *_EgAddressMgr) WithDni(dni string) Option {
	return optionFunc(func(o *options) { o.query["dni"] = dni })
}

// WithDateAdd date_add获取 
func (obj *_EgAddressMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

// WithDateUpd date_upd获取 
func (obj *_EgAddressMgr) WithDateUpd(dateUpd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_upd"] = dateUpd })
}

// WithActive active获取 
func (obj *_EgAddressMgr) WithActive(active bool) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}

// WithDeleted deleted获取 
func (obj *_EgAddressMgr) WithDeleted(deleted bool) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}


// GetByOption 功能选项模式获取
func (obj *_EgAddressMgr) GetByOption(opts ...Option) (result EgAddress, err error) {
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
func (obj *_EgAddressMgr) GetByOptions(opts ...Option) (results []*EgAddress, err error) {
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


// GetFromIDAddress 通过id_address获取内容  
func (obj *_EgAddressMgr)  GetFromIDAddress(idAddress uint32) (result EgAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_address = ?", idAddress).Find(&result).Error
	
	return
}

// GetBatchFromIDAddress 批量唯一主键查找 
func (obj *_EgAddressMgr) GetBatchFromIDAddress(idAddresss []uint32) (results []*EgAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_address IN (?)", idAddresss).Find(&results).Error
	
	return
}
 
// GetFromIDCountry 通过id_country获取内容  
func (obj *_EgAddressMgr) GetFromIDCountry(idCountry uint32) (results []*EgAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_country = ?", idCountry).Find(&results).Error
	
	return
}

// GetBatchFromIDCountry 批量唯一主键查找 
func (obj *_EgAddressMgr) GetBatchFromIDCountry(idCountrys []uint32) (results []*EgAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_country IN (?)", idCountrys).Find(&results).Error
	
	return
}
 
// GetFromIDState 通过id_state获取内容  
func (obj *_EgAddressMgr) GetFromIDState(idState uint32) (results []*EgAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_state = ?", idState).Find(&results).Error
	
	return
}

// GetBatchFromIDState 批量唯一主键查找 
func (obj *_EgAddressMgr) GetBatchFromIDState(idStates []uint32) (results []*EgAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_state IN (?)", idStates).Find(&results).Error
	
	return
}
 
// GetFromIDCustomer 通过id_customer获取内容  
func (obj *_EgAddressMgr) GetFromIDCustomer(idCustomer uint32) (results []*EgAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer = ?", idCustomer).Find(&results).Error
	
	return
}

// GetBatchFromIDCustomer 批量唯一主键查找 
func (obj *_EgAddressMgr) GetBatchFromIDCustomer(idCustomers []uint32) (results []*EgAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer IN (?)", idCustomers).Find(&results).Error
	
	return
}
 
// GetFromIDManufacturer 通过id_manufacturer获取内容  
func (obj *_EgAddressMgr) GetFromIDManufacturer(idManufacturer uint32) (results []*EgAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_manufacturer = ?", idManufacturer).Find(&results).Error
	
	return
}

// GetBatchFromIDManufacturer 批量唯一主键查找 
func (obj *_EgAddressMgr) GetBatchFromIDManufacturer(idManufacturers []uint32) (results []*EgAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_manufacturer IN (?)", idManufacturers).Find(&results).Error
	
	return
}
 
// GetFromIDSupplier 通过id_supplier获取内容  
func (obj *_EgAddressMgr) GetFromIDSupplier(idSupplier uint32) (results []*EgAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supplier = ?", idSupplier).Find(&results).Error
	
	return
}

// GetBatchFromIDSupplier 批量唯一主键查找 
func (obj *_EgAddressMgr) GetBatchFromIDSupplier(idSuppliers []uint32) (results []*EgAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supplier IN (?)", idSuppliers).Find(&results).Error
	
	return
}
 
// GetFromIDWarehouse 通过id_warehouse获取内容  
func (obj *_EgAddressMgr) GetFromIDWarehouse(idWarehouse uint32) (results []*EgAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_warehouse = ?", idWarehouse).Find(&results).Error
	
	return
}

// GetBatchFromIDWarehouse 批量唯一主键查找 
func (obj *_EgAddressMgr) GetBatchFromIDWarehouse(idWarehouses []uint32) (results []*EgAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_warehouse IN (?)", idWarehouses).Find(&results).Error
	
	return
}
 
// GetFromAlias 通过alias获取内容  
func (obj *_EgAddressMgr) GetFromAlias(alias string) (results []*EgAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("alias = ?", alias).Find(&results).Error
	
	return
}

// GetBatchFromAlias 批量唯一主键查找 
func (obj *_EgAddressMgr) GetBatchFromAlias(aliass []string) (results []*EgAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("alias IN (?)", aliass).Find(&results).Error
	
	return
}
 
// GetFromCompany 通过company获取内容  
func (obj *_EgAddressMgr) GetFromCompany(company string) (results []*EgAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("company = ?", company).Find(&results).Error
	
	return
}

// GetBatchFromCompany 批量唯一主键查找 
func (obj *_EgAddressMgr) GetBatchFromCompany(companys []string) (results []*EgAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("company IN (?)", companys).Find(&results).Error
	
	return
}
 
// GetFromLastname 通过lastname获取内容  
func (obj *_EgAddressMgr) GetFromLastname(lastname string) (results []*EgAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("lastname = ?", lastname).Find(&results).Error
	
	return
}

// GetBatchFromLastname 批量唯一主键查找 
func (obj *_EgAddressMgr) GetBatchFromLastname(lastnames []string) (results []*EgAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("lastname IN (?)", lastnames).Find(&results).Error
	
	return
}
 
// GetFromFirstname 通过firstname获取内容  
func (obj *_EgAddressMgr) GetFromFirstname(firstname string) (results []*EgAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("firstname = ?", firstname).Find(&results).Error
	
	return
}

// GetBatchFromFirstname 批量唯一主键查找 
func (obj *_EgAddressMgr) GetBatchFromFirstname(firstnames []string) (results []*EgAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("firstname IN (?)", firstnames).Find(&results).Error
	
	return
}
 
// GetFromAddress1 通过address1获取内容  
func (obj *_EgAddressMgr) GetFromAddress1(address1 string) (results []*EgAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("address1 = ?", address1).Find(&results).Error
	
	return
}

// GetBatchFromAddress1 批量唯一主键查找 
func (obj *_EgAddressMgr) GetBatchFromAddress1(address1s []string) (results []*EgAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("address1 IN (?)", address1s).Find(&results).Error
	
	return
}
 
// GetFromAddress2 通过address2获取内容  
func (obj *_EgAddressMgr) GetFromAddress2(address2 string) (results []*EgAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("address2 = ?", address2).Find(&results).Error
	
	return
}

// GetBatchFromAddress2 批量唯一主键查找 
func (obj *_EgAddressMgr) GetBatchFromAddress2(address2s []string) (results []*EgAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("address2 IN (?)", address2s).Find(&results).Error
	
	return
}
 
// GetFromPostcode 通过postcode获取内容  
func (obj *_EgAddressMgr) GetFromPostcode(postcode string) (results []*EgAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("postcode = ?", postcode).Find(&results).Error
	
	return
}

// GetBatchFromPostcode 批量唯一主键查找 
func (obj *_EgAddressMgr) GetBatchFromPostcode(postcodes []string) (results []*EgAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("postcode IN (?)", postcodes).Find(&results).Error
	
	return
}
 
// GetFromCity 通过city获取内容  
func (obj *_EgAddressMgr) GetFromCity(city string) (results []*EgAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("city = ?", city).Find(&results).Error
	
	return
}

// GetBatchFromCity 批量唯一主键查找 
func (obj *_EgAddressMgr) GetBatchFromCity(citys []string) (results []*EgAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("city IN (?)", citys).Find(&results).Error
	
	return
}
 
// GetFromOther 通过other获取内容  
func (obj *_EgAddressMgr) GetFromOther(other string) (results []*EgAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("other = ?", other).Find(&results).Error
	
	return
}

// GetBatchFromOther 批量唯一主键查找 
func (obj *_EgAddressMgr) GetBatchFromOther(others []string) (results []*EgAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("other IN (?)", others).Find(&results).Error
	
	return
}
 
// GetFromPhone 通过phone获取内容  
func (obj *_EgAddressMgr) GetFromPhone(phone string) (results []*EgAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("phone = ?", phone).Find(&results).Error
	
	return
}

// GetBatchFromPhone 批量唯一主键查找 
func (obj *_EgAddressMgr) GetBatchFromPhone(phones []string) (results []*EgAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("phone IN (?)", phones).Find(&results).Error
	
	return
}
 
// GetFromPhoneMobile 通过phone_mobile获取内容  
func (obj *_EgAddressMgr) GetFromPhoneMobile(phoneMobile string) (results []*EgAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("phone_mobile = ?", phoneMobile).Find(&results).Error
	
	return
}

// GetBatchFromPhoneMobile 批量唯一主键查找 
func (obj *_EgAddressMgr) GetBatchFromPhoneMobile(phoneMobiles []string) (results []*EgAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("phone_mobile IN (?)", phoneMobiles).Find(&results).Error
	
	return
}
 
// GetFromVatNumber 通过vat_number获取内容  
func (obj *_EgAddressMgr) GetFromVatNumber(vatNumber string) (results []*EgAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("vat_number = ?", vatNumber).Find(&results).Error
	
	return
}

// GetBatchFromVatNumber 批量唯一主键查找 
func (obj *_EgAddressMgr) GetBatchFromVatNumber(vatNumbers []string) (results []*EgAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("vat_number IN (?)", vatNumbers).Find(&results).Error
	
	return
}
 
// GetFromDni 通过dni获取内容  
func (obj *_EgAddressMgr) GetFromDni(dni string) (results []*EgAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("dni = ?", dni).Find(&results).Error
	
	return
}

// GetBatchFromDni 批量唯一主键查找 
func (obj *_EgAddressMgr) GetBatchFromDni(dnis []string) (results []*EgAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("dni IN (?)", dnis).Find(&results).Error
	
	return
}
 
// GetFromDateAdd 通过date_add获取内容  
func (obj *_EgAddressMgr) GetFromDateAdd(dateAdd time.Time) (results []*EgAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error
	
	return
}

// GetBatchFromDateAdd 批量唯一主键查找 
func (obj *_EgAddressMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*EgAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error
	
	return
}
 
// GetFromDateUpd 通过date_upd获取内容  
func (obj *_EgAddressMgr) GetFromDateUpd(dateUpd time.Time) (results []*EgAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd = ?", dateUpd).Find(&results).Error
	
	return
}

// GetBatchFromDateUpd 批量唯一主键查找 
func (obj *_EgAddressMgr) GetBatchFromDateUpd(dateUpds []time.Time) (results []*EgAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd IN (?)", dateUpds).Find(&results).Error
	
	return
}
 
// GetFromActive 通过active获取内容  
func (obj *_EgAddressMgr) GetFromActive(active bool) (results []*EgAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active = ?", active).Find(&results).Error
	
	return
}

// GetBatchFromActive 批量唯一主键查找 
func (obj *_EgAddressMgr) GetBatchFromActive(actives []bool) (results []*EgAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active IN (?)", actives).Find(&results).Error
	
	return
}
 
// GetFromDeleted 通过deleted获取内容  
func (obj *_EgAddressMgr) GetFromDeleted(deleted bool) (results []*EgAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("deleted = ?", deleted).Find(&results).Error
	
	return
}

// GetBatchFromDeleted 批量唯一主键查找 
func (obj *_EgAddressMgr) GetBatchFromDeleted(deleteds []bool) (results []*EgAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("deleted IN (?)", deleteds).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgAddressMgr) FetchByPrimaryKey(idAddress uint32 ) (result EgAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_address = ?", idAddress).Find(&result).Error
	
	return
}
 

 
 // FetchIndexByIDCountry  获取多个内容
 func (obj *_EgAddressMgr) FetchIndexByIDCountry(idCountry uint32 ) (results []*EgAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_country = ?", idCountry).Find(&results).Error
	
	return
}
 
 // FetchIndexByIDState  获取多个内容
 func (obj *_EgAddressMgr) FetchIndexByIDState(idState uint32 ) (results []*EgAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_state = ?", idState).Find(&results).Error
	
	return
}
 
 // FetchIndexByAddressCustomer  获取多个内容
 func (obj *_EgAddressMgr) FetchIndexByAddressCustomer(idCustomer uint32 ) (results []*EgAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer = ?", idCustomer).Find(&results).Error
	
	return
}
 
 // FetchIndexByIDManufacturer  获取多个内容
 func (obj *_EgAddressMgr) FetchIndexByIDManufacturer(idManufacturer uint32 ) (results []*EgAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_manufacturer = ?", idManufacturer).Find(&results).Error
	
	return
}
 
 // FetchIndexByIDSupplier  获取多个内容
 func (obj *_EgAddressMgr) FetchIndexByIDSupplier(idSupplier uint32 ) (results []*EgAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supplier = ?", idSupplier).Find(&results).Error
	
	return
}
 
 // FetchIndexByIDWarehouse  获取多个内容
 func (obj *_EgAddressMgr) FetchIndexByIDWarehouse(idWarehouse uint32 ) (results []*EgAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_warehouse = ?", idWarehouse).Find(&results).Error
	
	return
}
 

	

