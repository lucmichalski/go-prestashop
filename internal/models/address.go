package models

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AddressMgr struct {
	*_BaseMgr
}

func AddressMgr(db *gorm.DB) *_AddressMgr {
	if db == nil {
		panic(fmt.Errorf("AddressMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AddressMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_address"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_AddressMgr) GetTableName() string {
	return "ps_address"
}

func (obj *_AddressMgr) Get() (result Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_AddressMgr) Gets() (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

func (obj *_AddressMgr) WithIDAddress(idAddress uint32) Option {
	return optionFunc(func(o *options) { o.query["id_address"] = idAddress })
}

func (obj *_AddressMgr) WithIDCountry(idCountry uint32) Option {
	return optionFunc(func(o *options) { o.query["id_country"] = idCountry })
}

func (obj *_AddressMgr) WithIDState(idState uint32) Option {
	return optionFunc(func(o *options) { o.query["id_state"] = idState })
}

func (obj *_AddressMgr) WithIDCustomer(idCustomer uint32) Option {
	return optionFunc(func(o *options) { o.query["id_customer"] = idCustomer })
}

func (obj *_AddressMgr) WithIDManufacturer(idManufacturer uint32) Option {
	return optionFunc(func(o *options) { o.query["id_manufacturer"] = idManufacturer })
}

func (obj *_AddressMgr) WithIDSupplier(idSupplier uint32) Option {
	return optionFunc(func(o *options) { o.query["id_supplier"] = idSupplier })
}

func (obj *_AddressMgr) WithIDWarehouse(idWarehouse uint32) Option {
	return optionFunc(func(o *options) { o.query["id_warehouse"] = idWarehouse })
}

func (obj *_AddressMgr) WithAlias(alias string) Option {
	return optionFunc(func(o *options) { o.query["alias"] = alias })
}

func (obj *_AddressMgr) WithCompany(company string) Option {
	return optionFunc(func(o *options) { o.query["company"] = company })
}

func (obj *_AddressMgr) WithLastname(lastname string) Option {
	return optionFunc(func(o *options) { o.query["lastname"] = lastname })
}

func (obj *_AddressMgr) WithFirstname(firstname string) Option {
	return optionFunc(func(o *options) { o.query["firstname"] = firstname })
}

func (obj *_AddressMgr) WithAddress1(address1 string) Option {
	return optionFunc(func(o *options) { o.query["address1"] = address1 })
}

func (obj *_AddressMgr) WithAddress2(address2 string) Option {
	return optionFunc(func(o *options) { o.query["address2"] = address2 })
}

func (obj *_AddressMgr) WithPostcode(postcode string) Option {
	return optionFunc(func(o *options) { o.query["postcode"] = postcode })
}

func (obj *_AddressMgr) WithCity(city string) Option {
	return optionFunc(func(o *options) { o.query["city"] = city })
}

func (obj *_AddressMgr) WithOther(other string) Option {
	return optionFunc(func(o *options) { o.query["other"] = other })
}

func (obj *_AddressMgr) WithPhone(phone string) Option {
	return optionFunc(func(o *options) { o.query["phone"] = phone })
}

func (obj *_AddressMgr) WithPhoneMobile(phoneMobile string) Option {
	return optionFunc(func(o *options) { o.query["phone_mobile"] = phoneMobile })
}

func (obj *_AddressMgr) WithVatNumber(vatNumber string) Option {
	return optionFunc(func(o *options) { o.query["vat_number"] = vatNumber })
}

func (obj *_AddressMgr) WithDni(dni string) Option {
	return optionFunc(func(o *options) { o.query["dni"] = dni })
}

func (obj *_AddressMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

func (obj *_AddressMgr) WithDateUpd(dateUpd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_upd"] = dateUpd })
}

func (obj *_AddressMgr) WithActive(active bool) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}

func (obj *_AddressMgr) WithDeleted(deleted bool) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

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

func (obj *_AddressMgr) GetFromIDAddress(idAddress uint32) (result Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_address = ?", idAddress).Find(&result).Error

	return
}

func (obj *_AddressMgr) GetBatchFromIDAddress(idAddresss []uint32) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_address IN (?)", idAddresss).Find(&results).Error

	return
}

func (obj *_AddressMgr) GetFromIDCountry(idCountry uint32) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_country = ?", idCountry).Find(&results).Error

	return
}

func (obj *_AddressMgr) GetBatchFromIDCountry(idCountrys []uint32) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_country IN (?)", idCountrys).Find(&results).Error

	return
}

func (obj *_AddressMgr) GetFromIDState(idState uint32) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_state = ?", idState).Find(&results).Error

	return
}

func (obj *_AddressMgr) GetBatchFromIDState(idStates []uint32) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_state IN (?)", idStates).Find(&results).Error

	return
}

func (obj *_AddressMgr) GetFromIDCustomer(idCustomer uint32) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer = ?", idCustomer).Find(&results).Error

	return
}

func (obj *_AddressMgr) GetBatchFromIDCustomer(idCustomers []uint32) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer IN (?)", idCustomers).Find(&results).Error

	return
}

func (obj *_AddressMgr) GetFromIDManufacturer(idManufacturer uint32) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_manufacturer = ?", idManufacturer).Find(&results).Error

	return
}

func (obj *_AddressMgr) GetBatchFromIDManufacturer(idManufacturers []uint32) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_manufacturer IN (?)", idManufacturers).Find(&results).Error

	return
}

func (obj *_AddressMgr) GetFromIDSupplier(idSupplier uint32) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supplier = ?", idSupplier).Find(&results).Error

	return
}

func (obj *_AddressMgr) GetBatchFromIDSupplier(idSuppliers []uint32) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supplier IN (?)", idSuppliers).Find(&results).Error

	return
}

func (obj *_AddressMgr) GetFromIDWarehouse(idWarehouse uint32) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_warehouse = ?", idWarehouse).Find(&results).Error

	return
}

func (obj *_AddressMgr) GetBatchFromIDWarehouse(idWarehouses []uint32) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_warehouse IN (?)", idWarehouses).Find(&results).Error

	return
}

func (obj *_AddressMgr) GetFromAlias(alias string) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("alias = ?", alias).Find(&results).Error

	return
}

func (obj *_AddressMgr) GetBatchFromAlias(aliass []string) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("alias IN (?)", aliass).Find(&results).Error

	return
}

func (obj *_AddressMgr) GetFromCompany(company string) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("company = ?", company).Find(&results).Error

	return
}

func (obj *_AddressMgr) GetBatchFromCompany(companys []string) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("company IN (?)", companys).Find(&results).Error

	return
}

func (obj *_AddressMgr) GetFromLastname(lastname string) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("lastname = ?", lastname).Find(&results).Error

	return
}

func (obj *_AddressMgr) GetBatchFromLastname(lastnames []string) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("lastname IN (?)", lastnames).Find(&results).Error

	return
}

func (obj *_AddressMgr) GetFromFirstname(firstname string) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("firstname = ?", firstname).Find(&results).Error

	return
}

func (obj *_AddressMgr) GetBatchFromFirstname(firstnames []string) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("firstname IN (?)", firstnames).Find(&results).Error

	return
}

func (obj *_AddressMgr) GetFromAddress1(address1 string) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("address1 = ?", address1).Find(&results).Error

	return
}

func (obj *_AddressMgr) GetBatchFromAddress1(address1s []string) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("address1 IN (?)", address1s).Find(&results).Error

	return
}

func (obj *_AddressMgr) GetFromAddress2(address2 string) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("address2 = ?", address2).Find(&results).Error

	return
}

func (obj *_AddressMgr) GetBatchFromAddress2(address2s []string) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("address2 IN (?)", address2s).Find(&results).Error

	return
}

func (obj *_AddressMgr) GetFromPostcode(postcode string) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("postcode = ?", postcode).Find(&results).Error

	return
}

func (obj *_AddressMgr) GetBatchFromPostcode(postcodes []string) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("postcode IN (?)", postcodes).Find(&results).Error

	return
}

func (obj *_AddressMgr) GetFromCity(city string) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("city = ?", city).Find(&results).Error

	return
}

func (obj *_AddressMgr) GetBatchFromCity(citys []string) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("city IN (?)", citys).Find(&results).Error

	return
}

func (obj *_AddressMgr) GetFromOther(other string) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("other = ?", other).Find(&results).Error

	return
}

func (obj *_AddressMgr) GetBatchFromOther(others []string) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("other IN (?)", others).Find(&results).Error

	return
}

func (obj *_AddressMgr) GetFromPhone(phone string) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("phone = ?", phone).Find(&results).Error

	return
}

func (obj *_AddressMgr) GetBatchFromPhone(phones []string) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("phone IN (?)", phones).Find(&results).Error

	return
}

func (obj *_AddressMgr) GetFromPhoneMobile(phoneMobile string) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("phone_mobile = ?", phoneMobile).Find(&results).Error

	return
}

func (obj *_AddressMgr) GetBatchFromPhoneMobile(phoneMobiles []string) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("phone_mobile IN (?)", phoneMobiles).Find(&results).Error

	return
}

func (obj *_AddressMgr) GetFromVatNumber(vatNumber string) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("vat_number = ?", vatNumber).Find(&results).Error

	return
}

func (obj *_AddressMgr) GetBatchFromVatNumber(vatNumbers []string) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("vat_number IN (?)", vatNumbers).Find(&results).Error

	return
}

func (obj *_AddressMgr) GetFromDni(dni string) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("dni = ?", dni).Find(&results).Error

	return
}

func (obj *_AddressMgr) GetBatchFromDni(dnis []string) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("dni IN (?)", dnis).Find(&results).Error

	return
}

func (obj *_AddressMgr) GetFromDateAdd(dateAdd time.Time) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error

	return
}

func (obj *_AddressMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error

	return
}

func (obj *_AddressMgr) GetFromDateUpd(dateUpd time.Time) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd = ?", dateUpd).Find(&results).Error

	return
}

func (obj *_AddressMgr) GetBatchFromDateUpd(dateUpds []time.Time) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd IN (?)", dateUpds).Find(&results).Error

	return
}

func (obj *_AddressMgr) GetFromActive(active bool) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active = ?", active).Find(&results).Error

	return
}

func (obj *_AddressMgr) GetBatchFromActive(actives []bool) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active IN (?)", actives).Find(&results).Error

	return
}

func (obj *_AddressMgr) GetFromDeleted(deleted bool) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("deleted = ?", deleted).Find(&results).Error

	return
}

func (obj *_AddressMgr) GetBatchFromDeleted(deleteds []bool) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("deleted IN (?)", deleteds).Find(&results).Error

	return
}

func (obj *_AddressMgr) FetchByPrimaryKey(idAddress uint32) (result Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_address = ?", idAddress).Find(&result).Error

	return
}

func (obj *_AddressMgr) FetchIndexByIDCountry(idCountry uint32) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_country = ?", idCountry).Find(&results).Error

	return
}

func (obj *_AddressMgr) FetchIndexByIDState(idState uint32) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_state = ?", idState).Find(&results).Error

	return
}

func (obj *_AddressMgr) FetchIndexByAddressCustomer(idCustomer uint32) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer = ?", idCustomer).Find(&results).Error

	return
}

func (obj *_AddressMgr) FetchIndexByIDManufacturer(idManufacturer uint32) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_manufacturer = ?", idManufacturer).Find(&results).Error

	return
}

func (obj *_AddressMgr) FetchIndexByIDSupplier(idSupplier uint32) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supplier = ?", idSupplier).Find(&results).Error

	return
}

func (obj *_AddressMgr) FetchIndexByIDWarehouse(idWarehouse uint32) (results []*Address, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_warehouse = ?", idWarehouse).Find(&results).Error

	return
}
