package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AddressMgr struct {
	*_BaseMgr
}

// AddressMgr open func
func AddressMgr(db *gorm.DB) *_AddressMgr {
	if db == nil {
		panic(fmt.Errorf("AddressMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AddressMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_address"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AddressMgr) GetTableName() string {
	return "ps_address"
}

// Get 获取
func (obj *_AddressMgr) Get() (result Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AddressMgr) Gets() (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDAddress id_address获取
func (obj *_AddressMgr) WithIDAddress(idAddress uint32) Option {
	return optionFunc(func(o *options) { o.query["id_address"] = idAddress })
}

// WithIDCountry id_country获取
func (obj *_AddressMgr) WithIDCountry(idCountry uint32) Option {
	return optionFunc(func(o *options) { o.query["id_country"] = idCountry })
}

// WithIDState id_state获取
func (obj *_AddressMgr) WithIDState(idState uint32) Option {
	return optionFunc(func(o *options) { o.query["id_state"] = idState })
}

// WithIDCustomer id_customer获取
func (obj *_AddressMgr) WithIDCustomer(idCustomer uint32) Option {
	return optionFunc(func(o *options) { o.query["id_customer"] = idCustomer })
}

// WithIDManufacturer id_manufacturer获取
func (obj *_AddressMgr) WithIDManufacturer(idManufacturer uint32) Option {
	return optionFunc(func(o *options) { o.query["id_manufacturer"] = idManufacturer })
}

// WithIDSupplier id_supplier获取
func (obj *_AddressMgr) WithIDSupplier(idSupplier uint32) Option {
	return optionFunc(func(o *options) { o.query["id_supplier"] = idSupplier })
}

// WithIDWarehouse id_warehouse获取
func (obj *_AddressMgr) WithIDWarehouse(idWarehouse uint32) Option {
	return optionFunc(func(o *options) { o.query["id_warehouse"] = idWarehouse })
}

// WithAlias alias获取
func (obj *_AddressMgr) WithAlias(alias string) Option {
	return optionFunc(func(o *options) { o.query["alias"] = alias })
}

// WithCompany company获取
func (obj *_AddressMgr) WithCompany(company string) Option {
	return optionFunc(func(o *options) { o.query["company"] = company })
}

// WithLastname lastname获取
func (obj *_AddressMgr) WithLastname(lastname string) Option {
	return optionFunc(func(o *options) { o.query["lastname"] = lastname })
}

// WithFirstname firstname获取
func (obj *_AddressMgr) WithFirstname(firstname string) Option {
	return optionFunc(func(o *options) { o.query["firstname"] = firstname })
}

// WithAddress1 address1获取
func (obj *_AddressMgr) WithAddress1(address1 string) Option {
	return optionFunc(func(o *options) { o.query["address1"] = address1 })
}

// WithAddress2 address2获取
func (obj *_AddressMgr) WithAddress2(address2 string) Option {
	return optionFunc(func(o *options) { o.query["address2"] = address2 })
}

// WithPostcode postcode获取
func (obj *_AddressMgr) WithPostcode(postcode string) Option {
	return optionFunc(func(o *options) { o.query["postcode"] = postcode })
}

// WithCity city获取
func (obj *_AddressMgr) WithCity(city string) Option {
	return optionFunc(func(o *options) { o.query["city"] = city })
}

// WithOther other获取
func (obj *_AddressMgr) WithOther(other string) Option {
	return optionFunc(func(o *options) { o.query["other"] = other })
}

// WithPhone phone获取
func (obj *_AddressMgr) WithPhone(phone string) Option {
	return optionFunc(func(o *options) { o.query["phone"] = phone })
}

// WithPhoneMobile phone_mobile获取
func (obj *_AddressMgr) WithPhoneMobile(phoneMobile string) Option {
	return optionFunc(func(o *options) { o.query["phone_mobile"] = phoneMobile })
}

// WithVatNumber vat_number获取
func (obj *_AddressMgr) WithVatNumber(vatNumber string) Option {
	return optionFunc(func(o *options) { o.query["vat_number"] = vatNumber })
}

// WithDni dni获取
func (obj *_AddressMgr) WithDni(dni string) Option {
	return optionFunc(func(o *options) { o.query["dni"] = dni })
}

// WithDateAdd date_add获取
func (obj *_AddressMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

// WithDateUpd date_upd获取
func (obj *_AddressMgr) WithDateUpd(dateUpd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_upd"] = dateUpd })
}

// WithActive active获取
func (obj *_AddressMgr) WithActive(active bool) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}

// WithDeleted deleted获取
func (obj *_AddressMgr) WithDeleted(deleted bool) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// GetByOption 功能选项模式获取
func (obj *_AddressMgr) GetByOption(opts ...Option) (result Address, err error) {
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
func (obj *_AddressMgr) GetByOptions(opts ...Option) (results []*Address, err error) {
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
func (obj *_AddressMgr) GetFromIDAddress(idAddress uint32) (result Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_address = ?", idAddress).Find(&result).Error

	return
}

// GetBatchFromIDAddress 批量唯一主键查找
func (obj *_AddressMgr) GetBatchFromIDAddress(idAddresss []uint32) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_address IN (?)", idAddresss).Find(&results).Error

	return
}

// GetFromIDCountry 通过id_country获取内容
func (obj *_AddressMgr) GetFromIDCountry(idCountry uint32) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_country = ?", idCountry).Find(&results).Error

	return
}

// GetBatchFromIDCountry 批量唯一主键查找
func (obj *_AddressMgr) GetBatchFromIDCountry(idCountrys []uint32) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_country IN (?)", idCountrys).Find(&results).Error

	return
}

// GetFromIDState 通过id_state获取内容
func (obj *_AddressMgr) GetFromIDState(idState uint32) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_state = ?", idState).Find(&results).Error

	return
}

// GetBatchFromIDState 批量唯一主键查找
func (obj *_AddressMgr) GetBatchFromIDState(idStates []uint32) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_state IN (?)", idStates).Find(&results).Error

	return
}

// GetFromIDCustomer 通过id_customer获取内容
func (obj *_AddressMgr) GetFromIDCustomer(idCustomer uint32) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer = ?", idCustomer).Find(&results).Error

	return
}

// GetBatchFromIDCustomer 批量唯一主键查找
func (obj *_AddressMgr) GetBatchFromIDCustomer(idCustomers []uint32) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer IN (?)", idCustomers).Find(&results).Error

	return
}

// GetFromIDManufacturer 通过id_manufacturer获取内容
func (obj *_AddressMgr) GetFromIDManufacturer(idManufacturer uint32) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_manufacturer = ?", idManufacturer).Find(&results).Error

	return
}

// GetBatchFromIDManufacturer 批量唯一主键查找
func (obj *_AddressMgr) GetBatchFromIDManufacturer(idManufacturers []uint32) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_manufacturer IN (?)", idManufacturers).Find(&results).Error

	return
}

// GetFromIDSupplier 通过id_supplier获取内容
func (obj *_AddressMgr) GetFromIDSupplier(idSupplier uint32) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supplier = ?", idSupplier).Find(&results).Error

	return
}

// GetBatchFromIDSupplier 批量唯一主键查找
func (obj *_AddressMgr) GetBatchFromIDSupplier(idSuppliers []uint32) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supplier IN (?)", idSuppliers).Find(&results).Error

	return
}

// GetFromIDWarehouse 通过id_warehouse获取内容
func (obj *_AddressMgr) GetFromIDWarehouse(idWarehouse uint32) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_warehouse = ?", idWarehouse).Find(&results).Error

	return
}

// GetBatchFromIDWarehouse 批量唯一主键查找
func (obj *_AddressMgr) GetBatchFromIDWarehouse(idWarehouses []uint32) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_warehouse IN (?)", idWarehouses).Find(&results).Error

	return
}

// GetFromAlias 通过alias获取内容
func (obj *_AddressMgr) GetFromAlias(alias string) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("alias = ?", alias).Find(&results).Error

	return
}

// GetBatchFromAlias 批量唯一主键查找
func (obj *_AddressMgr) GetBatchFromAlias(aliass []string) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("alias IN (?)", aliass).Find(&results).Error

	return
}

// GetFromCompany 通过company获取内容
func (obj *_AddressMgr) GetFromCompany(company string) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("company = ?", company).Find(&results).Error

	return
}

// GetBatchFromCompany 批量唯一主键查找
func (obj *_AddressMgr) GetBatchFromCompany(companys []string) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("company IN (?)", companys).Find(&results).Error

	return
}

// GetFromLastname 通过lastname获取内容
func (obj *_AddressMgr) GetFromLastname(lastname string) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("lastname = ?", lastname).Find(&results).Error

	return
}

// GetBatchFromLastname 批量唯一主键查找
func (obj *_AddressMgr) GetBatchFromLastname(lastnames []string) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("lastname IN (?)", lastnames).Find(&results).Error

	return
}

// GetFromFirstname 通过firstname获取内容
func (obj *_AddressMgr) GetFromFirstname(firstname string) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("firstname = ?", firstname).Find(&results).Error

	return
}

// GetBatchFromFirstname 批量唯一主键查找
func (obj *_AddressMgr) GetBatchFromFirstname(firstnames []string) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("firstname IN (?)", firstnames).Find(&results).Error

	return
}

// GetFromAddress1 通过address1获取内容
func (obj *_AddressMgr) GetFromAddress1(address1 string) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("address1 = ?", address1).Find(&results).Error

	return
}

// GetBatchFromAddress1 批量唯一主键查找
func (obj *_AddressMgr) GetBatchFromAddress1(address1s []string) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("address1 IN (?)", address1s).Find(&results).Error

	return
}

// GetFromAddress2 通过address2获取内容
func (obj *_AddressMgr) GetFromAddress2(address2 string) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("address2 = ?", address2).Find(&results).Error

	return
}

// GetBatchFromAddress2 批量唯一主键查找
func (obj *_AddressMgr) GetBatchFromAddress2(address2s []string) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("address2 IN (?)", address2s).Find(&results).Error

	return
}

// GetFromPostcode 通过postcode获取内容
func (obj *_AddressMgr) GetFromPostcode(postcode string) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("postcode = ?", postcode).Find(&results).Error

	return
}

// GetBatchFromPostcode 批量唯一主键查找
func (obj *_AddressMgr) GetBatchFromPostcode(postcodes []string) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("postcode IN (?)", postcodes).Find(&results).Error

	return
}

// GetFromCity 通过city获取内容
func (obj *_AddressMgr) GetFromCity(city string) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("city = ?", city).Find(&results).Error

	return
}

// GetBatchFromCity 批量唯一主键查找
func (obj *_AddressMgr) GetBatchFromCity(citys []string) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("city IN (?)", citys).Find(&results).Error

	return
}

// GetFromOther 通过other获取内容
func (obj *_AddressMgr) GetFromOther(other string) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("other = ?", other).Find(&results).Error

	return
}

// GetBatchFromOther 批量唯一主键查找
func (obj *_AddressMgr) GetBatchFromOther(others []string) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("other IN (?)", others).Find(&results).Error

	return
}

// GetFromPhone 通过phone获取内容
func (obj *_AddressMgr) GetFromPhone(phone string) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("phone = ?", phone).Find(&results).Error

	return
}

// GetBatchFromPhone 批量唯一主键查找
func (obj *_AddressMgr) GetBatchFromPhone(phones []string) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("phone IN (?)", phones).Find(&results).Error

	return
}

// GetFromPhoneMobile 通过phone_mobile获取内容
func (obj *_AddressMgr) GetFromPhoneMobile(phoneMobile string) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("phone_mobile = ?", phoneMobile).Find(&results).Error

	return
}

// GetBatchFromPhoneMobile 批量唯一主键查找
func (obj *_AddressMgr) GetBatchFromPhoneMobile(phoneMobiles []string) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("phone_mobile IN (?)", phoneMobiles).Find(&results).Error

	return
}

// GetFromVatNumber 通过vat_number获取内容
func (obj *_AddressMgr) GetFromVatNumber(vatNumber string) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("vat_number = ?", vatNumber).Find(&results).Error

	return
}

// GetBatchFromVatNumber 批量唯一主键查找
func (obj *_AddressMgr) GetBatchFromVatNumber(vatNumbers []string) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("vat_number IN (?)", vatNumbers).Find(&results).Error

	return
}

// GetFromDni 通过dni获取内容
func (obj *_AddressMgr) GetFromDni(dni string) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("dni = ?", dni).Find(&results).Error

	return
}

// GetBatchFromDni 批量唯一主键查找
func (obj *_AddressMgr) GetBatchFromDni(dnis []string) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("dni IN (?)", dnis).Find(&results).Error

	return
}

// GetFromDateAdd 通过date_add获取内容
func (obj *_AddressMgr) GetFromDateAdd(dateAdd time.Time) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error

	return
}

// GetBatchFromDateAdd 批量唯一主键查找
func (obj *_AddressMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error

	return
}

// GetFromDateUpd 通过date_upd获取内容
func (obj *_AddressMgr) GetFromDateUpd(dateUpd time.Time) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd = ?", dateUpd).Find(&results).Error

	return
}

// GetBatchFromDateUpd 批量唯一主键查找
func (obj *_AddressMgr) GetBatchFromDateUpd(dateUpds []time.Time) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd IN (?)", dateUpds).Find(&results).Error

	return
}

// GetFromActive 通过active获取内容
func (obj *_AddressMgr) GetFromActive(active bool) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active = ?", active).Find(&results).Error

	return
}

// GetBatchFromActive 批量唯一主键查找
func (obj *_AddressMgr) GetBatchFromActive(actives []bool) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active IN (?)", actives).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容
func (obj *_AddressMgr) GetFromDeleted(deleted bool) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("deleted = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量唯一主键查找
func (obj *_AddressMgr) GetBatchFromDeleted(deleteds []bool) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("deleted IN (?)", deleteds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_AddressMgr) FetchByPrimaryKey(idAddress uint32) (result Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_address = ?", idAddress).Find(&result).Error

	return
}

// FetchIndexByIDCountry  获取多个内容
func (obj *_AddressMgr) FetchIndexByIDCountry(idCountry uint32) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_country = ?", idCountry).Find(&results).Error

	return
}

// FetchIndexByIDState  获取多个内容
func (obj *_AddressMgr) FetchIndexByIDState(idState uint32) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_state = ?", idState).Find(&results).Error

	return
}

// FetchIndexByAddressCustomer  获取多个内容
func (obj *_AddressMgr) FetchIndexByAddressCustomer(idCustomer uint32) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer = ?", idCustomer).Find(&results).Error

	return
}

// FetchIndexByIDManufacturer  获取多个内容
func (obj *_AddressMgr) FetchIndexByIDManufacturer(idManufacturer uint32) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_manufacturer = ?", idManufacturer).Find(&results).Error

	return
}

// FetchIndexByIDSupplier  获取多个内容
func (obj *_AddressMgr) FetchIndexByIDSupplier(idSupplier uint32) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supplier = ?", idSupplier).Find(&results).Error

	return
}

// FetchIndexByIDWarehouse  获取多个内容
func (obj *_AddressMgr) FetchIndexByIDWarehouse(idWarehouse uint32) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_warehouse = ?", idWarehouse).Find(&results).Error

	return
}
