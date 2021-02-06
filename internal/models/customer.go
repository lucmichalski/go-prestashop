package	model	
import (	
"time"	
"gorm.io/datatypes"	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgCustomerMgr struct {
	*_BaseMgr
}

// EgCustomerMgr open func
func EgCustomerMgr(db *gorm.DB) *_EgCustomerMgr {
	if db == nil {
		panic(fmt.Errorf("EgCustomerMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgCustomerMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_customer"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgCustomerMgr) GetTableName() string {
	return "eg_customer"
}

// Get 获取
func (obj *_EgCustomerMgr) Get() (result EgCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgCustomerMgr) Gets() (results []*EgCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDCustomer id_customer获取 
func (obj *_EgCustomerMgr) WithIDCustomer(idCustomer uint32) Option {
	return optionFunc(func(o *options) { o.query["id_customer"] = idCustomer })
}

// WithIDShopGroup id_shop_group获取 
func (obj *_EgCustomerMgr) WithIDShopGroup(idShopGroup uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop_group"] = idShopGroup })
}

// WithIDShop id_shop获取 
func (obj *_EgCustomerMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

// WithIDGender id_gender获取 
func (obj *_EgCustomerMgr) WithIDGender(idGender uint32) Option {
	return optionFunc(func(o *options) { o.query["id_gender"] = idGender })
}

// WithIDDefaultGroup id_default_group获取 
func (obj *_EgCustomerMgr) WithIDDefaultGroup(idDefaultGroup uint32) Option {
	return optionFunc(func(o *options) { o.query["id_default_group"] = idDefaultGroup })
}

// WithIDLang id_lang获取 
func (obj *_EgCustomerMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

// WithIDRisk id_risk获取 
func (obj *_EgCustomerMgr) WithIDRisk(idRisk uint32) Option {
	return optionFunc(func(o *options) { o.query["id_risk"] = idRisk })
}

// WithCompany company获取 
func (obj *_EgCustomerMgr) WithCompany(company string) Option {
	return optionFunc(func(o *options) { o.query["company"] = company })
}

// WithSiret siret获取 
func (obj *_EgCustomerMgr) WithSiret(siret string) Option {
	return optionFunc(func(o *options) { o.query["siret"] = siret })
}

// WithApe ape获取 
func (obj *_EgCustomerMgr) WithApe(ape string) Option {
	return optionFunc(func(o *options) { o.query["ape"] = ape })
}

// WithFirstname firstname获取 
func (obj *_EgCustomerMgr) WithFirstname(firstname string) Option {
	return optionFunc(func(o *options) { o.query["firstname"] = firstname })
}

// WithLastname lastname获取 
func (obj *_EgCustomerMgr) WithLastname(lastname string) Option {
	return optionFunc(func(o *options) { o.query["lastname"] = lastname })
}

// WithEmail email获取 
func (obj *_EgCustomerMgr) WithEmail(email string) Option {
	return optionFunc(func(o *options) { o.query["email"] = email })
}

// WithPasswd passwd获取 
func (obj *_EgCustomerMgr) WithPasswd(passwd string) Option {
	return optionFunc(func(o *options) { o.query["passwd"] = passwd })
}

// WithLastPasswdGen last_passwd_gen获取 
func (obj *_EgCustomerMgr) WithLastPasswdGen(lastPasswdGen time.Time) Option {
	return optionFunc(func(o *options) { o.query["last_passwd_gen"] = lastPasswdGen })
}

// WithBirthday birthday获取 
func (obj *_EgCustomerMgr) WithBirthday(birthday datatypes.Date) Option {
	return optionFunc(func(o *options) { o.query["birthday"] = birthday })
}

// WithNewsletter newsletter获取 
func (obj *_EgCustomerMgr) WithNewsletter(newsletter bool) Option {
	return optionFunc(func(o *options) { o.query["newsletter"] = newsletter })
}

// WithIPRegistrationNewsletter ip_registration_newsletter获取 
func (obj *_EgCustomerMgr) WithIPRegistrationNewsletter(ipRegistrationNewsletter string) Option {
	return optionFunc(func(o *options) { o.query["ip_registration_newsletter"] = ipRegistrationNewsletter })
}

// WithNewsletterDateAdd newsletter_date_add获取 
func (obj *_EgCustomerMgr) WithNewsletterDateAdd(newsletterDateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["newsletter_date_add"] = newsletterDateAdd })
}

// WithOptin optin获取 
func (obj *_EgCustomerMgr) WithOptin(optin bool) Option {
	return optionFunc(func(o *options) { o.query["optin"] = optin })
}

// WithWebsite website获取 
func (obj *_EgCustomerMgr) WithWebsite(website string) Option {
	return optionFunc(func(o *options) { o.query["website"] = website })
}

// WithOutstandingAllowAmount outstanding_allow_amount获取 
func (obj *_EgCustomerMgr) WithOutstandingAllowAmount(outstandingAllowAmount float64) Option {
	return optionFunc(func(o *options) { o.query["outstanding_allow_amount"] = outstandingAllowAmount })
}

// WithShowPublicPrices show_public_prices获取 
func (obj *_EgCustomerMgr) WithShowPublicPrices(showPublicPrices bool) Option {
	return optionFunc(func(o *options) { o.query["show_public_prices"] = showPublicPrices })
}

// WithMaxPaymentDays max_payment_days获取 
func (obj *_EgCustomerMgr) WithMaxPaymentDays(maxPaymentDays uint32) Option {
	return optionFunc(func(o *options) { o.query["max_payment_days"] = maxPaymentDays })
}

// WithSecureKey secure_key获取 
func (obj *_EgCustomerMgr) WithSecureKey(secureKey string) Option {
	return optionFunc(func(o *options) { o.query["secure_key"] = secureKey })
}

// WithNote note获取 
func (obj *_EgCustomerMgr) WithNote(note string) Option {
	return optionFunc(func(o *options) { o.query["note"] = note })
}

// WithActive active获取 
func (obj *_EgCustomerMgr) WithActive(active bool) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}

// WithIsGuest is_guest获取 
func (obj *_EgCustomerMgr) WithIsGuest(isGuest bool) Option {
	return optionFunc(func(o *options) { o.query["is_guest"] = isGuest })
}

// WithDeleted deleted获取 
func (obj *_EgCustomerMgr) WithDeleted(deleted bool) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// WithDateAdd date_add获取 
func (obj *_EgCustomerMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

// WithDateUpd date_upd获取 
func (obj *_EgCustomerMgr) WithDateUpd(dateUpd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_upd"] = dateUpd })
}

// WithResetPasswordToken reset_password_token获取 
func (obj *_EgCustomerMgr) WithResetPasswordToken(resetPasswordToken string) Option {
	return optionFunc(func(o *options) { o.query["reset_password_token"] = resetPasswordToken })
}

// WithResetPasswordValidity reset_password_validity获取 
func (obj *_EgCustomerMgr) WithResetPasswordValidity(resetPasswordValidity time.Time) Option {
	return optionFunc(func(o *options) { o.query["reset_password_validity"] = resetPasswordValidity })
}


// GetByOption 功能选项模式获取
func (obj *_EgCustomerMgr) GetByOption(opts ...Option) (result EgCustomer, err error) {
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
func (obj *_EgCustomerMgr) GetByOptions(opts ...Option) (results []*EgCustomer, err error) {
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
func (obj *_EgCustomerMgr)  GetFromIDCustomer(idCustomer uint32) (result EgCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer = ?", idCustomer).Find(&result).Error
	
	return
}

// GetBatchFromIDCustomer 批量唯一主键查找 
func (obj *_EgCustomerMgr) GetBatchFromIDCustomer(idCustomers []uint32) (results []*EgCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer IN (?)", idCustomers).Find(&results).Error
	
	return
}
 
// GetFromIDShopGroup 通过id_shop_group获取内容  
func (obj *_EgCustomerMgr) GetFromIDShopGroup(idShopGroup uint32) (results []*EgCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_group = ?", idShopGroup).Find(&results).Error
	
	return
}

// GetBatchFromIDShopGroup 批量唯一主键查找 
func (obj *_EgCustomerMgr) GetBatchFromIDShopGroup(idShopGroups []uint32) (results []*EgCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_group IN (?)", idShopGroups).Find(&results).Error
	
	return
}
 
// GetFromIDShop 通过id_shop获取内容  
func (obj *_EgCustomerMgr) GetFromIDShop(idShop uint32) (results []*EgCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error
	
	return
}

// GetBatchFromIDShop 批量唯一主键查找 
func (obj *_EgCustomerMgr) GetBatchFromIDShop(idShops []uint32) (results []*EgCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error
	
	return
}
 
// GetFromIDGender 通过id_gender获取内容  
func (obj *_EgCustomerMgr) GetFromIDGender(idGender uint32) (results []*EgCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_gender = ?", idGender).Find(&results).Error
	
	return
}

// GetBatchFromIDGender 批量唯一主键查找 
func (obj *_EgCustomerMgr) GetBatchFromIDGender(idGenders []uint32) (results []*EgCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_gender IN (?)", idGenders).Find(&results).Error
	
	return
}
 
// GetFromIDDefaultGroup 通过id_default_group获取内容  
func (obj *_EgCustomerMgr) GetFromIDDefaultGroup(idDefaultGroup uint32) (results []*EgCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_default_group = ?", idDefaultGroup).Find(&results).Error
	
	return
}

// GetBatchFromIDDefaultGroup 批量唯一主键查找 
func (obj *_EgCustomerMgr) GetBatchFromIDDefaultGroup(idDefaultGroups []uint32) (results []*EgCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_default_group IN (?)", idDefaultGroups).Find(&results).Error
	
	return
}
 
// GetFromIDLang 通过id_lang获取内容  
func (obj *_EgCustomerMgr) GetFromIDLang(idLang uint32) (results []*EgCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error
	
	return
}

// GetBatchFromIDLang 批量唯一主键查找 
func (obj *_EgCustomerMgr) GetBatchFromIDLang(idLangs []uint32) (results []*EgCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error
	
	return
}
 
// GetFromIDRisk 通过id_risk获取内容  
func (obj *_EgCustomerMgr) GetFromIDRisk(idRisk uint32) (results []*EgCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_risk = ?", idRisk).Find(&results).Error
	
	return
}

// GetBatchFromIDRisk 批量唯一主键查找 
func (obj *_EgCustomerMgr) GetBatchFromIDRisk(idRisks []uint32) (results []*EgCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_risk IN (?)", idRisks).Find(&results).Error
	
	return
}
 
// GetFromCompany 通过company获取内容  
func (obj *_EgCustomerMgr) GetFromCompany(company string) (results []*EgCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("company = ?", company).Find(&results).Error
	
	return
}

// GetBatchFromCompany 批量唯一主键查找 
func (obj *_EgCustomerMgr) GetBatchFromCompany(companys []string) (results []*EgCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("company IN (?)", companys).Find(&results).Error
	
	return
}
 
// GetFromSiret 通过siret获取内容  
func (obj *_EgCustomerMgr) GetFromSiret(siret string) (results []*EgCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("siret = ?", siret).Find(&results).Error
	
	return
}

// GetBatchFromSiret 批量唯一主键查找 
func (obj *_EgCustomerMgr) GetBatchFromSiret(sirets []string) (results []*EgCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("siret IN (?)", sirets).Find(&results).Error
	
	return
}
 
// GetFromApe 通过ape获取内容  
func (obj *_EgCustomerMgr) GetFromApe(ape string) (results []*EgCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ape = ?", ape).Find(&results).Error
	
	return
}

// GetBatchFromApe 批量唯一主键查找 
func (obj *_EgCustomerMgr) GetBatchFromApe(apes []string) (results []*EgCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ape IN (?)", apes).Find(&results).Error
	
	return
}
 
// GetFromFirstname 通过firstname获取内容  
func (obj *_EgCustomerMgr) GetFromFirstname(firstname string) (results []*EgCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("firstname = ?", firstname).Find(&results).Error
	
	return
}

// GetBatchFromFirstname 批量唯一主键查找 
func (obj *_EgCustomerMgr) GetBatchFromFirstname(firstnames []string) (results []*EgCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("firstname IN (?)", firstnames).Find(&results).Error
	
	return
}
 
// GetFromLastname 通过lastname获取内容  
func (obj *_EgCustomerMgr) GetFromLastname(lastname string) (results []*EgCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("lastname = ?", lastname).Find(&results).Error
	
	return
}

// GetBatchFromLastname 批量唯一主键查找 
func (obj *_EgCustomerMgr) GetBatchFromLastname(lastnames []string) (results []*EgCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("lastname IN (?)", lastnames).Find(&results).Error
	
	return
}
 
// GetFromEmail 通过email获取内容  
func (obj *_EgCustomerMgr) GetFromEmail(email string) (results []*EgCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("email = ?", email).Find(&results).Error
	
	return
}

// GetBatchFromEmail 批量唯一主键查找 
func (obj *_EgCustomerMgr) GetBatchFromEmail(emails []string) (results []*EgCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("email IN (?)", emails).Find(&results).Error
	
	return
}
 
// GetFromPasswd 通过passwd获取内容  
func (obj *_EgCustomerMgr) GetFromPasswd(passwd string) (results []*EgCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("passwd = ?", passwd).Find(&results).Error
	
	return
}

// GetBatchFromPasswd 批量唯一主键查找 
func (obj *_EgCustomerMgr) GetBatchFromPasswd(passwds []string) (results []*EgCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("passwd IN (?)", passwds).Find(&results).Error
	
	return
}
 
// GetFromLastPasswdGen 通过last_passwd_gen获取内容  
func (obj *_EgCustomerMgr) GetFromLastPasswdGen(lastPasswdGen time.Time) (results []*EgCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("last_passwd_gen = ?", lastPasswdGen).Find(&results).Error
	
	return
}

// GetBatchFromLastPasswdGen 批量唯一主键查找 
func (obj *_EgCustomerMgr) GetBatchFromLastPasswdGen(lastPasswdGens []time.Time) (results []*EgCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("last_passwd_gen IN (?)", lastPasswdGens).Find(&results).Error
	
	return
}
 
// GetFromBirthday 通过birthday获取内容  
func (obj *_EgCustomerMgr) GetFromBirthday(birthday datatypes.Date) (results []*EgCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("birthday = ?", birthday).Find(&results).Error
	
	return
}

// GetBatchFromBirthday 批量唯一主键查找 
func (obj *_EgCustomerMgr) GetBatchFromBirthday(birthdays []datatypes.Date) (results []*EgCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("birthday IN (?)", birthdays).Find(&results).Error
	
	return
}
 
// GetFromNewsletter 通过newsletter获取内容  
func (obj *_EgCustomerMgr) GetFromNewsletter(newsletter bool) (results []*EgCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("newsletter = ?", newsletter).Find(&results).Error
	
	return
}

// GetBatchFromNewsletter 批量唯一主键查找 
func (obj *_EgCustomerMgr) GetBatchFromNewsletter(newsletters []bool) (results []*EgCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("newsletter IN (?)", newsletters).Find(&results).Error
	
	return
}
 
// GetFromIPRegistrationNewsletter 通过ip_registration_newsletter获取内容  
func (obj *_EgCustomerMgr) GetFromIPRegistrationNewsletter(ipRegistrationNewsletter string) (results []*EgCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ip_registration_newsletter = ?", ipRegistrationNewsletter).Find(&results).Error
	
	return
}

// GetBatchFromIPRegistrationNewsletter 批量唯一主键查找 
func (obj *_EgCustomerMgr) GetBatchFromIPRegistrationNewsletter(ipRegistrationNewsletters []string) (results []*EgCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ip_registration_newsletter IN (?)", ipRegistrationNewsletters).Find(&results).Error
	
	return
}
 
// GetFromNewsletterDateAdd 通过newsletter_date_add获取内容  
func (obj *_EgCustomerMgr) GetFromNewsletterDateAdd(newsletterDateAdd time.Time) (results []*EgCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("newsletter_date_add = ?", newsletterDateAdd).Find(&results).Error
	
	return
}

// GetBatchFromNewsletterDateAdd 批量唯一主键查找 
func (obj *_EgCustomerMgr) GetBatchFromNewsletterDateAdd(newsletterDateAdds []time.Time) (results []*EgCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("newsletter_date_add IN (?)", newsletterDateAdds).Find(&results).Error
	
	return
}
 
// GetFromOptin 通过optin获取内容  
func (obj *_EgCustomerMgr) GetFromOptin(optin bool) (results []*EgCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("optin = ?", optin).Find(&results).Error
	
	return
}

// GetBatchFromOptin 批量唯一主键查找 
func (obj *_EgCustomerMgr) GetBatchFromOptin(optins []bool) (results []*EgCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("optin IN (?)", optins).Find(&results).Error
	
	return
}
 
// GetFromWebsite 通过website获取内容  
func (obj *_EgCustomerMgr) GetFromWebsite(website string) (results []*EgCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("website = ?", website).Find(&results).Error
	
	return
}

// GetBatchFromWebsite 批量唯一主键查找 
func (obj *_EgCustomerMgr) GetBatchFromWebsite(websites []string) (results []*EgCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("website IN (?)", websites).Find(&results).Error
	
	return
}
 
// GetFromOutstandingAllowAmount 通过outstanding_allow_amount获取内容  
func (obj *_EgCustomerMgr) GetFromOutstandingAllowAmount(outstandingAllowAmount float64) (results []*EgCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("outstanding_allow_amount = ?", outstandingAllowAmount).Find(&results).Error
	
	return
}

// GetBatchFromOutstandingAllowAmount 批量唯一主键查找 
func (obj *_EgCustomerMgr) GetBatchFromOutstandingAllowAmount(outstandingAllowAmounts []float64) (results []*EgCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("outstanding_allow_amount IN (?)", outstandingAllowAmounts).Find(&results).Error
	
	return
}
 
// GetFromShowPublicPrices 通过show_public_prices获取内容  
func (obj *_EgCustomerMgr) GetFromShowPublicPrices(showPublicPrices bool) (results []*EgCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("show_public_prices = ?", showPublicPrices).Find(&results).Error
	
	return
}

// GetBatchFromShowPublicPrices 批量唯一主键查找 
func (obj *_EgCustomerMgr) GetBatchFromShowPublicPrices(showPublicPricess []bool) (results []*EgCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("show_public_prices IN (?)", showPublicPricess).Find(&results).Error
	
	return
}
 
// GetFromMaxPaymentDays 通过max_payment_days获取内容  
func (obj *_EgCustomerMgr) GetFromMaxPaymentDays(maxPaymentDays uint32) (results []*EgCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("max_payment_days = ?", maxPaymentDays).Find(&results).Error
	
	return
}

// GetBatchFromMaxPaymentDays 批量唯一主键查找 
func (obj *_EgCustomerMgr) GetBatchFromMaxPaymentDays(maxPaymentDayss []uint32) (results []*EgCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("max_payment_days IN (?)", maxPaymentDayss).Find(&results).Error
	
	return
}
 
// GetFromSecureKey 通过secure_key获取内容  
func (obj *_EgCustomerMgr) GetFromSecureKey(secureKey string) (results []*EgCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("secure_key = ?", secureKey).Find(&results).Error
	
	return
}

// GetBatchFromSecureKey 批量唯一主键查找 
func (obj *_EgCustomerMgr) GetBatchFromSecureKey(secureKeys []string) (results []*EgCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("secure_key IN (?)", secureKeys).Find(&results).Error
	
	return
}
 
// GetFromNote 通过note获取内容  
func (obj *_EgCustomerMgr) GetFromNote(note string) (results []*EgCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("note = ?", note).Find(&results).Error
	
	return
}

// GetBatchFromNote 批量唯一主键查找 
func (obj *_EgCustomerMgr) GetBatchFromNote(notes []string) (results []*EgCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("note IN (?)", notes).Find(&results).Error
	
	return
}
 
// GetFromActive 通过active获取内容  
func (obj *_EgCustomerMgr) GetFromActive(active bool) (results []*EgCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active = ?", active).Find(&results).Error
	
	return
}

// GetBatchFromActive 批量唯一主键查找 
func (obj *_EgCustomerMgr) GetBatchFromActive(actives []bool) (results []*EgCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active IN (?)", actives).Find(&results).Error
	
	return
}
 
// GetFromIsGuest 通过is_guest获取内容  
func (obj *_EgCustomerMgr) GetFromIsGuest(isGuest bool) (results []*EgCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_guest = ?", isGuest).Find(&results).Error
	
	return
}

// GetBatchFromIsGuest 批量唯一主键查找 
func (obj *_EgCustomerMgr) GetBatchFromIsGuest(isGuests []bool) (results []*EgCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_guest IN (?)", isGuests).Find(&results).Error
	
	return
}
 
// GetFromDeleted 通过deleted获取内容  
func (obj *_EgCustomerMgr) GetFromDeleted(deleted bool) (results []*EgCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("deleted = ?", deleted).Find(&results).Error
	
	return
}

// GetBatchFromDeleted 批量唯一主键查找 
func (obj *_EgCustomerMgr) GetBatchFromDeleted(deleteds []bool) (results []*EgCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("deleted IN (?)", deleteds).Find(&results).Error
	
	return
}
 
// GetFromDateAdd 通过date_add获取内容  
func (obj *_EgCustomerMgr) GetFromDateAdd(dateAdd time.Time) (results []*EgCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error
	
	return
}

// GetBatchFromDateAdd 批量唯一主键查找 
func (obj *_EgCustomerMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*EgCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error
	
	return
}
 
// GetFromDateUpd 通过date_upd获取内容  
func (obj *_EgCustomerMgr) GetFromDateUpd(dateUpd time.Time) (results []*EgCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd = ?", dateUpd).Find(&results).Error
	
	return
}

// GetBatchFromDateUpd 批量唯一主键查找 
func (obj *_EgCustomerMgr) GetBatchFromDateUpd(dateUpds []time.Time) (results []*EgCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd IN (?)", dateUpds).Find(&results).Error
	
	return
}
 
// GetFromResetPasswordToken 通过reset_password_token获取内容  
func (obj *_EgCustomerMgr) GetFromResetPasswordToken(resetPasswordToken string) (results []*EgCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reset_password_token = ?", resetPasswordToken).Find(&results).Error
	
	return
}

// GetBatchFromResetPasswordToken 批量唯一主键查找 
func (obj *_EgCustomerMgr) GetBatchFromResetPasswordToken(resetPasswordTokens []string) (results []*EgCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reset_password_token IN (?)", resetPasswordTokens).Find(&results).Error
	
	return
}
 
// GetFromResetPasswordValidity 通过reset_password_validity获取内容  
func (obj *_EgCustomerMgr) GetFromResetPasswordValidity(resetPasswordValidity time.Time) (results []*EgCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reset_password_validity = ?", resetPasswordValidity).Find(&results).Error
	
	return
}

// GetBatchFromResetPasswordValidity 批量唯一主键查找 
func (obj *_EgCustomerMgr) GetBatchFromResetPasswordValidity(resetPasswordValiditys []time.Time) (results []*EgCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reset_password_validity IN (?)", resetPasswordValiditys).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgCustomerMgr) FetchByPrimaryKey(idCustomer uint32 ) (result EgCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer = ?", idCustomer).Find(&result).Error
	
	return
}
 

 
 // FetchIndexByIDCustomerPasswd  获取多个内容
 func (obj *_EgCustomerMgr) FetchIndexByIDCustomerPasswd(idCustomer uint32 ,passwd string ) (results []*EgCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer = ? AND passwd = ?", idCustomer , passwd).Find(&results).Error
	
	return
}
 
 // FetchIndexByIDShopGroup  获取多个内容
 func (obj *_EgCustomerMgr) FetchIndexByIDShopGroup(idShopGroup uint32 ) (results []*EgCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_group = ?", idShopGroup).Find(&results).Error
	
	return
}
 
 // FetchIndexByIDShop  获取多个内容
 func (obj *_EgCustomerMgr) FetchIndexByIDShop(idShop uint32 ,dateAdd time.Time ) (results []*EgCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ? AND date_add = ?", idShop , dateAdd).Find(&results).Error
	
	return
}
 
 // FetchIndexByIDGender  获取多个内容
 func (obj *_EgCustomerMgr) FetchIndexByIDGender(idGender uint32 ) (results []*EgCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_gender = ?", idGender).Find(&results).Error
	
	return
}
 
 // FetchIndexByCustomerEmail  获取多个内容
 func (obj *_EgCustomerMgr) FetchIndexByCustomerEmail(email string ) (results []*EgCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("email = ?", email).Find(&results).Error
	
	return
}
 
 // FetchIndexByCustomerLogin  获取多个内容
 func (obj *_EgCustomerMgr) FetchIndexByCustomerLogin(email string ,passwd string ) (results []*EgCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("email = ? AND passwd = ?", email , passwd).Find(&results).Error
	
	return
}
 

	

