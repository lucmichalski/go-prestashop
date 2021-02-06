package models

import (
	"context"
	"fmt"
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type _CustomerMgr struct {
	*_BaseMgr
}

func CustomerMgr(db *gorm.DB) *_CustomerMgr {
	if db == nil {
		panic(fmt.Errorf("CustomerMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CustomerMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_customer"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_CustomerMgr) GetTableName() string {
	return "ps_customer"
}

func (obj *_CustomerMgr) Get() (result Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_CustomerMgr) Gets() (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_CustomerMgr) WithIDCustomer(idCustomer uint32) Option {
	return optionFunc(func(o *options) { o.query["id_customer"] = idCustomer })
}

func (obj *_CustomerMgr) WithIDShopGroup(idShopGroup uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop_group"] = idShopGroup })
}

func (obj *_CustomerMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

func (obj *_CustomerMgr) WithIDGender(idGender uint32) Option {
	return optionFunc(func(o *options) { o.query["id_gender"] = idGender })
}

func (obj *_CustomerMgr) WithIDDefaultGroup(idDefaultGroup uint32) Option {
	return optionFunc(func(o *options) { o.query["id_default_group"] = idDefaultGroup })
}

func (obj *_CustomerMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

func (obj *_CustomerMgr) WithIDRisk(idRisk uint32) Option {
	return optionFunc(func(o *options) { o.query["id_risk"] = idRisk })
}

func (obj *_CustomerMgr) WithCompany(company string) Option {
	return optionFunc(func(o *options) { o.query["company"] = company })
}

func (obj *_CustomerMgr) WithSiret(siret string) Option {
	return optionFunc(func(o *options) { o.query["siret"] = siret })
}

func (obj *_CustomerMgr) WithApe(ape string) Option {
	return optionFunc(func(o *options) { o.query["ape"] = ape })
}

func (obj *_CustomerMgr) WithFirstname(firstname string) Option {
	return optionFunc(func(o *options) { o.query["firstname"] = firstname })
}

func (obj *_CustomerMgr) WithLastname(lastname string) Option {
	return optionFunc(func(o *options) { o.query["lastname"] = lastname })
}

func (obj *_CustomerMgr) WithEmail(email string) Option {
	return optionFunc(func(o *options) { o.query["email"] = email })
}

func (obj *_CustomerMgr) WithPasswd(passwd string) Option {
	return optionFunc(func(o *options) { o.query["passwd"] = passwd })
}

func (obj *_CustomerMgr) WithLastPasswdGen(lastPasswdGen time.Time) Option {
	return optionFunc(func(o *options) { o.query["last_passwd_gen"] = lastPasswdGen })
}

func (obj *_CustomerMgr) WithBirthday(birthday datatypes.Date) Option {
	return optionFunc(func(o *options) { o.query["birthday"] = birthday })
}

func (obj *_CustomerMgr) WithNewsletter(newsletter bool) Option {
	return optionFunc(func(o *options) { o.query["newsletter"] = newsletter })
}

func (obj *_CustomerMgr) WithIPRegistrationNewsletter(ipRegistrationNewsletter string) Option {
	return optionFunc(func(o *options) { o.query["ip_registration_newsletter"] = ipRegistrationNewsletter })
}

func (obj *_CustomerMgr) WithNewsletterDateAdd(newsletterDateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["newsletter_date_add"] = newsletterDateAdd })
}

func (obj *_CustomerMgr) WithOptin(optin bool) Option {
	return optionFunc(func(o *options) { o.query["optin"] = optin })
}

func (obj *_CustomerMgr) WithWebsite(website string) Option {
	return optionFunc(func(o *options) { o.query["website"] = website })
}

func (obj *_CustomerMgr) WithOutstandingAllowAmount(outstandingAllowAmount float64) Option {
	return optionFunc(func(o *options) { o.query["outstanding_allow_amount"] = outstandingAllowAmount })
}

func (obj *_CustomerMgr) WithShowPublicPrices(showPublicPrices bool) Option {
	return optionFunc(func(o *options) { o.query["show_public_prices"] = showPublicPrices })
}

func (obj *_CustomerMgr) WithMaxPaymentDays(maxPaymentDays uint32) Option {
	return optionFunc(func(o *options) { o.query["max_payment_days"] = maxPaymentDays })
}

func (obj *_CustomerMgr) WithSecureKey(secureKey string) Option {
	return optionFunc(func(o *options) { o.query["secure_key"] = secureKey })
}

func (obj *_CustomerMgr) WithNote(note string) Option {
	return optionFunc(func(o *options) { o.query["note"] = note })
}

func (obj *_CustomerMgr) WithActive(active bool) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}

func (obj *_CustomerMgr) WithIsGuest(isGuest bool) Option {
	return optionFunc(func(o *options) { o.query["is_guest"] = isGuest })
}

func (obj *_CustomerMgr) WithDeleted(deleted bool) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

func (obj *_CustomerMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

func (obj *_CustomerMgr) WithDateUpd(dateUpd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_upd"] = dateUpd })
}

func (obj *_CustomerMgr) WithResetPasswordToken(resetPasswordToken string) Option {
	return optionFunc(func(o *options) { o.query["reset_password_token"] = resetPasswordToken })
}

func (obj *_CustomerMgr) WithResetPasswordValidity(resetPasswordValidity time.Time) Option {
	return optionFunc(func(o *options) { o.query["reset_password_validity"] = resetPasswordValidity })
}

func (obj *_CustomerMgr) GetByOption(opts ...Option) (result Customer, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_CustomerMgr) GetByOptions(opts ...Option) (results []*Customer, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_CustomerMgr) GetFromIDCustomer(idCustomer uint32) (result Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer = ?", idCustomer).Find(&result).Error

	return
}

func (obj *_CustomerMgr) GetBatchFromIDCustomer(idCustomers []uint32) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer IN (?)", idCustomers).Find(&results).Error

	return
}

func (obj *_CustomerMgr) GetFromIDShopGroup(idShopGroup uint32) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_group = ?", idShopGroup).Find(&results).Error

	return
}

func (obj *_CustomerMgr) GetBatchFromIDShopGroup(idShopGroups []uint32) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_group IN (?)", idShopGroups).Find(&results).Error

	return
}

func (obj *_CustomerMgr) GetFromIDShop(idShop uint32) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

func (obj *_CustomerMgr) GetBatchFromIDShop(idShops []uint32) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

func (obj *_CustomerMgr) GetFromIDGender(idGender uint32) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_gender = ?", idGender).Find(&results).Error

	return
}

func (obj *_CustomerMgr) GetBatchFromIDGender(idGenders []uint32) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_gender IN (?)", idGenders).Find(&results).Error

	return
}

func (obj *_CustomerMgr) GetFromIDDefaultGroup(idDefaultGroup uint32) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_default_group = ?", idDefaultGroup).Find(&results).Error

	return
}

func (obj *_CustomerMgr) GetBatchFromIDDefaultGroup(idDefaultGroups []uint32) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_default_group IN (?)", idDefaultGroups).Find(&results).Error

	return
}

func (obj *_CustomerMgr) GetFromIDLang(idLang uint32) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error

	return
}

func (obj *_CustomerMgr) GetBatchFromIDLang(idLangs []uint32) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error

	return
}

func (obj *_CustomerMgr) GetFromIDRisk(idRisk uint32) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_risk = ?", idRisk).Find(&results).Error

	return
}

func (obj *_CustomerMgr) GetBatchFromIDRisk(idRisks []uint32) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_risk IN (?)", idRisks).Find(&results).Error

	return
}

func (obj *_CustomerMgr) GetFromCompany(company string) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("company = ?", company).Find(&results).Error

	return
}

func (obj *_CustomerMgr) GetBatchFromCompany(companys []string) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("company IN (?)", companys).Find(&results).Error

	return
}

func (obj *_CustomerMgr) GetFromSiret(siret string) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("siret = ?", siret).Find(&results).Error

	return
}

func (obj *_CustomerMgr) GetBatchFromSiret(sirets []string) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("siret IN (?)", sirets).Find(&results).Error

	return
}

func (obj *_CustomerMgr) GetFromApe(ape string) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ape = ?", ape).Find(&results).Error

	return
}

func (obj *_CustomerMgr) GetBatchFromApe(apes []string) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ape IN (?)", apes).Find(&results).Error

	return
}

func (obj *_CustomerMgr) GetFromFirstname(firstname string) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("firstname = ?", firstname).Find(&results).Error

	return
}

func (obj *_CustomerMgr) GetBatchFromFirstname(firstnames []string) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("firstname IN (?)", firstnames).Find(&results).Error

	return
}

func (obj *_CustomerMgr) GetFromLastname(lastname string) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("lastname = ?", lastname).Find(&results).Error

	return
}

func (obj *_CustomerMgr) GetBatchFromLastname(lastnames []string) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("lastname IN (?)", lastnames).Find(&results).Error

	return
}

func (obj *_CustomerMgr) GetFromEmail(email string) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("email = ?", email).Find(&results).Error

	return
}

func (obj *_CustomerMgr) GetBatchFromEmail(emails []string) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("email IN (?)", emails).Find(&results).Error

	return
}

func (obj *_CustomerMgr) GetFromPasswd(passwd string) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("passwd = ?", passwd).Find(&results).Error

	return
}

func (obj *_CustomerMgr) GetBatchFromPasswd(passwds []string) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("passwd IN (?)", passwds).Find(&results).Error

	return
}

func (obj *_CustomerMgr) GetFromLastPasswdGen(lastPasswdGen time.Time) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("last_passwd_gen = ?", lastPasswdGen).Find(&results).Error

	return
}

func (obj *_CustomerMgr) GetBatchFromLastPasswdGen(lastPasswdGens []time.Time) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("last_passwd_gen IN (?)", lastPasswdGens).Find(&results).Error

	return
}

func (obj *_CustomerMgr) GetFromBirthday(birthday datatypes.Date) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("birthday = ?", birthday).Find(&results).Error

	return
}

func (obj *_CustomerMgr) GetBatchFromBirthday(birthdays []datatypes.Date) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("birthday IN (?)", birthdays).Find(&results).Error

	return
}

func (obj *_CustomerMgr) GetFromNewsletter(newsletter bool) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("newsletter = ?", newsletter).Find(&results).Error

	return
}

func (obj *_CustomerMgr) GetBatchFromNewsletter(newsletters []bool) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("newsletter IN (?)", newsletters).Find(&results).Error

	return
}

func (obj *_CustomerMgr) GetFromIPRegistrationNewsletter(ipRegistrationNewsletter string) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ip_registration_newsletter = ?", ipRegistrationNewsletter).Find(&results).Error

	return
}

func (obj *_CustomerMgr) GetBatchFromIPRegistrationNewsletter(ipRegistrationNewsletters []string) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ip_registration_newsletter IN (?)", ipRegistrationNewsletters).Find(&results).Error

	return
}

func (obj *_CustomerMgr) GetFromNewsletterDateAdd(newsletterDateAdd time.Time) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("newsletter_date_add = ?", newsletterDateAdd).Find(&results).Error

	return
}

func (obj *_CustomerMgr) GetBatchFromNewsletterDateAdd(newsletterDateAdds []time.Time) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("newsletter_date_add IN (?)", newsletterDateAdds).Find(&results).Error

	return
}

func (obj *_CustomerMgr) GetFromOptin(optin bool) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("optin = ?", optin).Find(&results).Error

	return
}

func (obj *_CustomerMgr) GetBatchFromOptin(optins []bool) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("optin IN (?)", optins).Find(&results).Error

	return
}

func (obj *_CustomerMgr) GetFromWebsite(website string) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("website = ?", website).Find(&results).Error

	return
}

func (obj *_CustomerMgr) GetBatchFromWebsite(websites []string) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("website IN (?)", websites).Find(&results).Error

	return
}

func (obj *_CustomerMgr) GetFromOutstandingAllowAmount(outstandingAllowAmount float64) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("outstanding_allow_amount = ?", outstandingAllowAmount).Find(&results).Error

	return
}

func (obj *_CustomerMgr) GetBatchFromOutstandingAllowAmount(outstandingAllowAmounts []float64) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("outstanding_allow_amount IN (?)", outstandingAllowAmounts).Find(&results).Error

	return
}

func (obj *_CustomerMgr) GetFromShowPublicPrices(showPublicPrices bool) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("show_public_prices = ?", showPublicPrices).Find(&results).Error

	return
}

func (obj *_CustomerMgr) GetBatchFromShowPublicPrices(showPublicPricess []bool) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("show_public_prices IN (?)", showPublicPricess).Find(&results).Error

	return
}

func (obj *_CustomerMgr) GetFromMaxPaymentDays(maxPaymentDays uint32) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("max_payment_days = ?", maxPaymentDays).Find(&results).Error

	return
}

func (obj *_CustomerMgr) GetBatchFromMaxPaymentDays(maxPaymentDayss []uint32) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("max_payment_days IN (?)", maxPaymentDayss).Find(&results).Error

	return
}

func (obj *_CustomerMgr) GetFromSecureKey(secureKey string) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("secure_key = ?", secureKey).Find(&results).Error

	return
}

func (obj *_CustomerMgr) GetBatchFromSecureKey(secureKeys []string) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("secure_key IN (?)", secureKeys).Find(&results).Error

	return
}

func (obj *_CustomerMgr) GetFromNote(note string) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("note = ?", note).Find(&results).Error

	return
}

func (obj *_CustomerMgr) GetBatchFromNote(notes []string) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("note IN (?)", notes).Find(&results).Error

	return
}

func (obj *_CustomerMgr) GetFromActive(active bool) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active = ?", active).Find(&results).Error

	return
}

func (obj *_CustomerMgr) GetBatchFromActive(actives []bool) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active IN (?)", actives).Find(&results).Error

	return
}

func (obj *_CustomerMgr) GetFromIsGuest(isGuest bool) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_guest = ?", isGuest).Find(&results).Error

	return
}

func (obj *_CustomerMgr) GetBatchFromIsGuest(isGuests []bool) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_guest IN (?)", isGuests).Find(&results).Error

	return
}

func (obj *_CustomerMgr) GetFromDeleted(deleted bool) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("deleted = ?", deleted).Find(&results).Error

	return
}

func (obj *_CustomerMgr) GetBatchFromDeleted(deleteds []bool) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("deleted IN (?)", deleteds).Find(&results).Error

	return
}

func (obj *_CustomerMgr) GetFromDateAdd(dateAdd time.Time) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error

	return
}

func (obj *_CustomerMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error

	return
}

func (obj *_CustomerMgr) GetFromDateUpd(dateUpd time.Time) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd = ?", dateUpd).Find(&results).Error

	return
}

func (obj *_CustomerMgr) GetBatchFromDateUpd(dateUpds []time.Time) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd IN (?)", dateUpds).Find(&results).Error

	return
}

func (obj *_CustomerMgr) GetFromResetPasswordToken(resetPasswordToken string) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reset_password_token = ?", resetPasswordToken).Find(&results).Error

	return
}

func (obj *_CustomerMgr) GetBatchFromResetPasswordToken(resetPasswordTokens []string) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reset_password_token IN (?)", resetPasswordTokens).Find(&results).Error

	return
}

func (obj *_CustomerMgr) GetFromResetPasswordValidity(resetPasswordValidity time.Time) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reset_password_validity = ?", resetPasswordValidity).Find(&results).Error

	return
}

func (obj *_CustomerMgr) GetBatchFromResetPasswordValidity(resetPasswordValiditys []time.Time) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reset_password_validity IN (?)", resetPasswordValiditys).Find(&results).Error

	return
}


func (obj *_CustomerMgr) FetchByPrimaryKey(idCustomer uint32) (result Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer = ?", idCustomer).Find(&result).Error

	return
}

func (obj *_CustomerMgr) FetchIndexByIDCustomerPasswd(idCustomer uint32, passwd string) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer = ? AND passwd = ?", idCustomer, passwd).Find(&results).Error

	return
}

func (obj *_CustomerMgr) FetchIndexByIDShopGroup(idShopGroup uint32) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_group = ?", idShopGroup).Find(&results).Error

	return
}

func (obj *_CustomerMgr) FetchIndexByIDShop(idShop uint32, dateAdd time.Time) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ? AND date_add = ?", idShop, dateAdd).Find(&results).Error

	return
}

func (obj *_CustomerMgr) FetchIndexByIDGender(idGender uint32) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_gender = ?", idGender).Find(&results).Error

	return
}

func (obj *_CustomerMgr) FetchIndexByCustomerEmail(email string) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("email = ?", email).Find(&results).Error

	return
}

func (obj *_CustomerMgr) FetchIndexByCustomerLogin(email string, passwd string) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("email = ? AND passwd = ?", email, passwd).Find(&results).Error

	return
}
