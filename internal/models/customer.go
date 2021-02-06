package model

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

// CustomerMgr open func
func CustomerMgr(db *gorm.DB) *_CustomerMgr {
	if db == nil {
		panic(fmt.Errorf("CustomerMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CustomerMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_customer"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CustomerMgr) GetTableName() string {
	return "ps_customer"
}

// Get 获取
func (obj *_CustomerMgr) Get() (result Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CustomerMgr) Gets() (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDCustomer id_customer获取
func (obj *_CustomerMgr) WithIDCustomer(idCustomer uint32) Option {
	return optionFunc(func(o *options) { o.query["id_customer"] = idCustomer })
}

// WithIDShopGroup id_shop_group获取
func (obj *_CustomerMgr) WithIDShopGroup(idShopGroup uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop_group"] = idShopGroup })
}

// WithIDShop id_shop获取
func (obj *_CustomerMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

// WithIDGender id_gender获取
func (obj *_CustomerMgr) WithIDGender(idGender uint32) Option {
	return optionFunc(func(o *options) { o.query["id_gender"] = idGender })
}

// WithIDDefaultGroup id_default_group获取
func (obj *_CustomerMgr) WithIDDefaultGroup(idDefaultGroup uint32) Option {
	return optionFunc(func(o *options) { o.query["id_default_group"] = idDefaultGroup })
}

// WithIDLang id_lang获取
func (obj *_CustomerMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

// WithIDRisk id_risk获取
func (obj *_CustomerMgr) WithIDRisk(idRisk uint32) Option {
	return optionFunc(func(o *options) { o.query["id_risk"] = idRisk })
}

// WithCompany company获取
func (obj *_CustomerMgr) WithCompany(company string) Option {
	return optionFunc(func(o *options) { o.query["company"] = company })
}

// WithSiret siret获取
func (obj *_CustomerMgr) WithSiret(siret string) Option {
	return optionFunc(func(o *options) { o.query["siret"] = siret })
}

// WithApe ape获取
func (obj *_CustomerMgr) WithApe(ape string) Option {
	return optionFunc(func(o *options) { o.query["ape"] = ape })
}

// WithFirstname firstname获取
func (obj *_CustomerMgr) WithFirstname(firstname string) Option {
	return optionFunc(func(o *options) { o.query["firstname"] = firstname })
}

// WithLastname lastname获取
func (obj *_CustomerMgr) WithLastname(lastname string) Option {
	return optionFunc(func(o *options) { o.query["lastname"] = lastname })
}

// WithEmail email获取
func (obj *_CustomerMgr) WithEmail(email string) Option {
	return optionFunc(func(o *options) { o.query["email"] = email })
}

// WithPasswd passwd获取
func (obj *_CustomerMgr) WithPasswd(passwd string) Option {
	return optionFunc(func(o *options) { o.query["passwd"] = passwd })
}

// WithLastPasswdGen last_passwd_gen获取
func (obj *_CustomerMgr) WithLastPasswdGen(lastPasswdGen time.Time) Option {
	return optionFunc(func(o *options) { o.query["last_passwd_gen"] = lastPasswdGen })
}

// WithBirthday birthday获取
func (obj *_CustomerMgr) WithBirthday(birthday datatypes.Date) Option {
	return optionFunc(func(o *options) { o.query["birthday"] = birthday })
}

// WithNewsletter newsletter获取
func (obj *_CustomerMgr) WithNewsletter(newsletter bool) Option {
	return optionFunc(func(o *options) { o.query["newsletter"] = newsletter })
}

// WithIPRegistrationNewsletter ip_registration_newsletter获取
func (obj *_CustomerMgr) WithIPRegistrationNewsletter(ipRegistrationNewsletter string) Option {
	return optionFunc(func(o *options) { o.query["ip_registration_newsletter"] = ipRegistrationNewsletter })
}

// WithNewsletterDateAdd newsletter_date_add获取
func (obj *_CustomerMgr) WithNewsletterDateAdd(newsletterDateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["newsletter_date_add"] = newsletterDateAdd })
}

// WithOptin optin获取
func (obj *_CustomerMgr) WithOptin(optin bool) Option {
	return optionFunc(func(o *options) { o.query["optin"] = optin })
}

// WithWebsite website获取
func (obj *_CustomerMgr) WithWebsite(website string) Option {
	return optionFunc(func(o *options) { o.query["website"] = website })
}

// WithOutstandingAllowAmount outstanding_allow_amount获取
func (obj *_CustomerMgr) WithOutstandingAllowAmount(outstandingAllowAmount float64) Option {
	return optionFunc(func(o *options) { o.query["outstanding_allow_amount"] = outstandingAllowAmount })
}

// WithShowPublicPrices show_public_prices获取
func (obj *_CustomerMgr) WithShowPublicPrices(showPublicPrices bool) Option {
	return optionFunc(func(o *options) { o.query["show_public_prices"] = showPublicPrices })
}

// WithMaxPaymentDays max_payment_days获取
func (obj *_CustomerMgr) WithMaxPaymentDays(maxPaymentDays uint32) Option {
	return optionFunc(func(o *options) { o.query["max_payment_days"] = maxPaymentDays })
}

// WithSecureKey secure_key获取
func (obj *_CustomerMgr) WithSecureKey(secureKey string) Option {
	return optionFunc(func(o *options) { o.query["secure_key"] = secureKey })
}

// WithNote note获取
func (obj *_CustomerMgr) WithNote(note string) Option {
	return optionFunc(func(o *options) { o.query["note"] = note })
}

// WithActive active获取
func (obj *_CustomerMgr) WithActive(active bool) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}

// WithIsGuest is_guest获取
func (obj *_CustomerMgr) WithIsGuest(isGuest bool) Option {
	return optionFunc(func(o *options) { o.query["is_guest"] = isGuest })
}

// WithDeleted deleted获取
func (obj *_CustomerMgr) WithDeleted(deleted bool) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// WithDateAdd date_add获取
func (obj *_CustomerMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

// WithDateUpd date_upd获取
func (obj *_CustomerMgr) WithDateUpd(dateUpd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_upd"] = dateUpd })
}

// WithResetPasswordToken reset_password_token获取
func (obj *_CustomerMgr) WithResetPasswordToken(resetPasswordToken string) Option {
	return optionFunc(func(o *options) { o.query["reset_password_token"] = resetPasswordToken })
}

// WithResetPasswordValidity reset_password_validity获取
func (obj *_CustomerMgr) WithResetPasswordValidity(resetPasswordValidity time.Time) Option {
	return optionFunc(func(o *options) { o.query["reset_password_validity"] = resetPasswordValidity })
}

// GetByOption 功能选项模式获取
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

// GetByOptions 批量功能选项模式获取
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

//////////////////////////enume case ////////////////////////////////////////////

// GetFromIDCustomer 通过id_customer获取内容
func (obj *_CustomerMgr) GetFromIDCustomer(idCustomer uint32) (result Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer = ?", idCustomer).Find(&result).Error

	return
}

// GetBatchFromIDCustomer 批量唯一主键查找
func (obj *_CustomerMgr) GetBatchFromIDCustomer(idCustomers []uint32) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer IN (?)", idCustomers).Find(&results).Error

	return
}

// GetFromIDShopGroup 通过id_shop_group获取内容
func (obj *_CustomerMgr) GetFromIDShopGroup(idShopGroup uint32) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_group = ?", idShopGroup).Find(&results).Error

	return
}

// GetBatchFromIDShopGroup 批量唯一主键查找
func (obj *_CustomerMgr) GetBatchFromIDShopGroup(idShopGroups []uint32) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_group IN (?)", idShopGroups).Find(&results).Error

	return
}

// GetFromIDShop 通过id_shop获取内容
func (obj *_CustomerMgr) GetFromIDShop(idShop uint32) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

// GetBatchFromIDShop 批量唯一主键查找
func (obj *_CustomerMgr) GetBatchFromIDShop(idShops []uint32) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

// GetFromIDGender 通过id_gender获取内容
func (obj *_CustomerMgr) GetFromIDGender(idGender uint32) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_gender = ?", idGender).Find(&results).Error

	return
}

// GetBatchFromIDGender 批量唯一主键查找
func (obj *_CustomerMgr) GetBatchFromIDGender(idGenders []uint32) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_gender IN (?)", idGenders).Find(&results).Error

	return
}

// GetFromIDDefaultGroup 通过id_default_group获取内容
func (obj *_CustomerMgr) GetFromIDDefaultGroup(idDefaultGroup uint32) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_default_group = ?", idDefaultGroup).Find(&results).Error

	return
}

// GetBatchFromIDDefaultGroup 批量唯一主键查找
func (obj *_CustomerMgr) GetBatchFromIDDefaultGroup(idDefaultGroups []uint32) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_default_group IN (?)", idDefaultGroups).Find(&results).Error

	return
}

// GetFromIDLang 通过id_lang获取内容
func (obj *_CustomerMgr) GetFromIDLang(idLang uint32) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error

	return
}

// GetBatchFromIDLang 批量唯一主键查找
func (obj *_CustomerMgr) GetBatchFromIDLang(idLangs []uint32) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error

	return
}

// GetFromIDRisk 通过id_risk获取内容
func (obj *_CustomerMgr) GetFromIDRisk(idRisk uint32) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_risk = ?", idRisk).Find(&results).Error

	return
}

// GetBatchFromIDRisk 批量唯一主键查找
func (obj *_CustomerMgr) GetBatchFromIDRisk(idRisks []uint32) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_risk IN (?)", idRisks).Find(&results).Error

	return
}

// GetFromCompany 通过company获取内容
func (obj *_CustomerMgr) GetFromCompany(company string) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("company = ?", company).Find(&results).Error

	return
}

// GetBatchFromCompany 批量唯一主键查找
func (obj *_CustomerMgr) GetBatchFromCompany(companys []string) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("company IN (?)", companys).Find(&results).Error

	return
}

// GetFromSiret 通过siret获取内容
func (obj *_CustomerMgr) GetFromSiret(siret string) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("siret = ?", siret).Find(&results).Error

	return
}

// GetBatchFromSiret 批量唯一主键查找
func (obj *_CustomerMgr) GetBatchFromSiret(sirets []string) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("siret IN (?)", sirets).Find(&results).Error

	return
}

// GetFromApe 通过ape获取内容
func (obj *_CustomerMgr) GetFromApe(ape string) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ape = ?", ape).Find(&results).Error

	return
}

// GetBatchFromApe 批量唯一主键查找
func (obj *_CustomerMgr) GetBatchFromApe(apes []string) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ape IN (?)", apes).Find(&results).Error

	return
}

// GetFromFirstname 通过firstname获取内容
func (obj *_CustomerMgr) GetFromFirstname(firstname string) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("firstname = ?", firstname).Find(&results).Error

	return
}

// GetBatchFromFirstname 批量唯一主键查找
func (obj *_CustomerMgr) GetBatchFromFirstname(firstnames []string) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("firstname IN (?)", firstnames).Find(&results).Error

	return
}

// GetFromLastname 通过lastname获取内容
func (obj *_CustomerMgr) GetFromLastname(lastname string) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("lastname = ?", lastname).Find(&results).Error

	return
}

// GetBatchFromLastname 批量唯一主键查找
func (obj *_CustomerMgr) GetBatchFromLastname(lastnames []string) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("lastname IN (?)", lastnames).Find(&results).Error

	return
}

// GetFromEmail 通过email获取内容
func (obj *_CustomerMgr) GetFromEmail(email string) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("email = ?", email).Find(&results).Error

	return
}

// GetBatchFromEmail 批量唯一主键查找
func (obj *_CustomerMgr) GetBatchFromEmail(emails []string) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("email IN (?)", emails).Find(&results).Error

	return
}

// GetFromPasswd 通过passwd获取内容
func (obj *_CustomerMgr) GetFromPasswd(passwd string) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("passwd = ?", passwd).Find(&results).Error

	return
}

// GetBatchFromPasswd 批量唯一主键查找
func (obj *_CustomerMgr) GetBatchFromPasswd(passwds []string) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("passwd IN (?)", passwds).Find(&results).Error

	return
}

// GetFromLastPasswdGen 通过last_passwd_gen获取内容
func (obj *_CustomerMgr) GetFromLastPasswdGen(lastPasswdGen time.Time) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("last_passwd_gen = ?", lastPasswdGen).Find(&results).Error

	return
}

// GetBatchFromLastPasswdGen 批量唯一主键查找
func (obj *_CustomerMgr) GetBatchFromLastPasswdGen(lastPasswdGens []time.Time) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("last_passwd_gen IN (?)", lastPasswdGens).Find(&results).Error

	return
}

// GetFromBirthday 通过birthday获取内容
func (obj *_CustomerMgr) GetFromBirthday(birthday datatypes.Date) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("birthday = ?", birthday).Find(&results).Error

	return
}

// GetBatchFromBirthday 批量唯一主键查找
func (obj *_CustomerMgr) GetBatchFromBirthday(birthdays []datatypes.Date) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("birthday IN (?)", birthdays).Find(&results).Error

	return
}

// GetFromNewsletter 通过newsletter获取内容
func (obj *_CustomerMgr) GetFromNewsletter(newsletter bool) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("newsletter = ?", newsletter).Find(&results).Error

	return
}

// GetBatchFromNewsletter 批量唯一主键查找
func (obj *_CustomerMgr) GetBatchFromNewsletter(newsletters []bool) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("newsletter IN (?)", newsletters).Find(&results).Error

	return
}

// GetFromIPRegistrationNewsletter 通过ip_registration_newsletter获取内容
func (obj *_CustomerMgr) GetFromIPRegistrationNewsletter(ipRegistrationNewsletter string) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ip_registration_newsletter = ?", ipRegistrationNewsletter).Find(&results).Error

	return
}

// GetBatchFromIPRegistrationNewsletter 批量唯一主键查找
func (obj *_CustomerMgr) GetBatchFromIPRegistrationNewsletter(ipRegistrationNewsletters []string) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ip_registration_newsletter IN (?)", ipRegistrationNewsletters).Find(&results).Error

	return
}

// GetFromNewsletterDateAdd 通过newsletter_date_add获取内容
func (obj *_CustomerMgr) GetFromNewsletterDateAdd(newsletterDateAdd time.Time) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("newsletter_date_add = ?", newsletterDateAdd).Find(&results).Error

	return
}

// GetBatchFromNewsletterDateAdd 批量唯一主键查找
func (obj *_CustomerMgr) GetBatchFromNewsletterDateAdd(newsletterDateAdds []time.Time) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("newsletter_date_add IN (?)", newsletterDateAdds).Find(&results).Error

	return
}

// GetFromOptin 通过optin获取内容
func (obj *_CustomerMgr) GetFromOptin(optin bool) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("optin = ?", optin).Find(&results).Error

	return
}

// GetBatchFromOptin 批量唯一主键查找
func (obj *_CustomerMgr) GetBatchFromOptin(optins []bool) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("optin IN (?)", optins).Find(&results).Error

	return
}

// GetFromWebsite 通过website获取内容
func (obj *_CustomerMgr) GetFromWebsite(website string) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("website = ?", website).Find(&results).Error

	return
}

// GetBatchFromWebsite 批量唯一主键查找
func (obj *_CustomerMgr) GetBatchFromWebsite(websites []string) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("website IN (?)", websites).Find(&results).Error

	return
}

// GetFromOutstandingAllowAmount 通过outstanding_allow_amount获取内容
func (obj *_CustomerMgr) GetFromOutstandingAllowAmount(outstandingAllowAmount float64) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("outstanding_allow_amount = ?", outstandingAllowAmount).Find(&results).Error

	return
}

// GetBatchFromOutstandingAllowAmount 批量唯一主键查找
func (obj *_CustomerMgr) GetBatchFromOutstandingAllowAmount(outstandingAllowAmounts []float64) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("outstanding_allow_amount IN (?)", outstandingAllowAmounts).Find(&results).Error

	return
}

// GetFromShowPublicPrices 通过show_public_prices获取内容
func (obj *_CustomerMgr) GetFromShowPublicPrices(showPublicPrices bool) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("show_public_prices = ?", showPublicPrices).Find(&results).Error

	return
}

// GetBatchFromShowPublicPrices 批量唯一主键查找
func (obj *_CustomerMgr) GetBatchFromShowPublicPrices(showPublicPricess []bool) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("show_public_prices IN (?)", showPublicPricess).Find(&results).Error

	return
}

// GetFromMaxPaymentDays 通过max_payment_days获取内容
func (obj *_CustomerMgr) GetFromMaxPaymentDays(maxPaymentDays uint32) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("max_payment_days = ?", maxPaymentDays).Find(&results).Error

	return
}

// GetBatchFromMaxPaymentDays 批量唯一主键查找
func (obj *_CustomerMgr) GetBatchFromMaxPaymentDays(maxPaymentDayss []uint32) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("max_payment_days IN (?)", maxPaymentDayss).Find(&results).Error

	return
}

// GetFromSecureKey 通过secure_key获取内容
func (obj *_CustomerMgr) GetFromSecureKey(secureKey string) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("secure_key = ?", secureKey).Find(&results).Error

	return
}

// GetBatchFromSecureKey 批量唯一主键查找
func (obj *_CustomerMgr) GetBatchFromSecureKey(secureKeys []string) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("secure_key IN (?)", secureKeys).Find(&results).Error

	return
}

// GetFromNote 通过note获取内容
func (obj *_CustomerMgr) GetFromNote(note string) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("note = ?", note).Find(&results).Error

	return
}

// GetBatchFromNote 批量唯一主键查找
func (obj *_CustomerMgr) GetBatchFromNote(notes []string) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("note IN (?)", notes).Find(&results).Error

	return
}

// GetFromActive 通过active获取内容
func (obj *_CustomerMgr) GetFromActive(active bool) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active = ?", active).Find(&results).Error

	return
}

// GetBatchFromActive 批量唯一主键查找
func (obj *_CustomerMgr) GetBatchFromActive(actives []bool) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active IN (?)", actives).Find(&results).Error

	return
}

// GetFromIsGuest 通过is_guest获取内容
func (obj *_CustomerMgr) GetFromIsGuest(isGuest bool) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_guest = ?", isGuest).Find(&results).Error

	return
}

// GetBatchFromIsGuest 批量唯一主键查找
func (obj *_CustomerMgr) GetBatchFromIsGuest(isGuests []bool) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_guest IN (?)", isGuests).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容
func (obj *_CustomerMgr) GetFromDeleted(deleted bool) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("deleted = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量唯一主键查找
func (obj *_CustomerMgr) GetBatchFromDeleted(deleteds []bool) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("deleted IN (?)", deleteds).Find(&results).Error

	return
}

// GetFromDateAdd 通过date_add获取内容
func (obj *_CustomerMgr) GetFromDateAdd(dateAdd time.Time) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error

	return
}

// GetBatchFromDateAdd 批量唯一主键查找
func (obj *_CustomerMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error

	return
}

// GetFromDateUpd 通过date_upd获取内容
func (obj *_CustomerMgr) GetFromDateUpd(dateUpd time.Time) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd = ?", dateUpd).Find(&results).Error

	return
}

// GetBatchFromDateUpd 批量唯一主键查找
func (obj *_CustomerMgr) GetBatchFromDateUpd(dateUpds []time.Time) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd IN (?)", dateUpds).Find(&results).Error

	return
}

// GetFromResetPasswordToken 通过reset_password_token获取内容
func (obj *_CustomerMgr) GetFromResetPasswordToken(resetPasswordToken string) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reset_password_token = ?", resetPasswordToken).Find(&results).Error

	return
}

// GetBatchFromResetPasswordToken 批量唯一主键查找
func (obj *_CustomerMgr) GetBatchFromResetPasswordToken(resetPasswordTokens []string) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reset_password_token IN (?)", resetPasswordTokens).Find(&results).Error

	return
}

// GetFromResetPasswordValidity 通过reset_password_validity获取内容
func (obj *_CustomerMgr) GetFromResetPasswordValidity(resetPasswordValidity time.Time) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reset_password_validity = ?", resetPasswordValidity).Find(&results).Error

	return
}

// GetBatchFromResetPasswordValidity 批量唯一主键查找
func (obj *_CustomerMgr) GetBatchFromResetPasswordValidity(resetPasswordValiditys []time.Time) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reset_password_validity IN (?)", resetPasswordValiditys).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_CustomerMgr) FetchByPrimaryKey(idCustomer uint32) (result Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer = ?", idCustomer).Find(&result).Error

	return
}

// FetchIndexByIDCustomerPasswd  获取多个内容
func (obj *_CustomerMgr) FetchIndexByIDCustomerPasswd(idCustomer uint32, passwd string) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer = ? AND passwd = ?", idCustomer, passwd).Find(&results).Error

	return
}

// FetchIndexByIDShopGroup  获取多个内容
func (obj *_CustomerMgr) FetchIndexByIDShopGroup(idShopGroup uint32) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_group = ?", idShopGroup).Find(&results).Error

	return
}

// FetchIndexByIDShop  获取多个内容
func (obj *_CustomerMgr) FetchIndexByIDShop(idShop uint32, dateAdd time.Time) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ? AND date_add = ?", idShop, dateAdd).Find(&results).Error

	return
}

// FetchIndexByIDGender  获取多个内容
func (obj *_CustomerMgr) FetchIndexByIDGender(idGender uint32) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_gender = ?", idGender).Find(&results).Error

	return
}

// FetchIndexByCustomerEmail  获取多个内容
func (obj *_CustomerMgr) FetchIndexByCustomerEmail(email string) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("email = ?", email).Find(&results).Error

	return
}

// FetchIndexByCustomerLogin  获取多个内容
func (obj *_CustomerMgr) FetchIndexByCustomerLogin(email string, passwd string) (results []*Customer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("email = ? AND passwd = ?", email, passwd).Find(&results).Error

	return
}
