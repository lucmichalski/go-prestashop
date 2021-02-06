package	model	
import (	
"time"	
"gorm.io/datatypes"	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgEmployeeMgr struct {
	*_BaseMgr
}

// EgEmployeeMgr open func
func EgEmployeeMgr(db *gorm.DB) *_EgEmployeeMgr {
	if db == nil {
		panic(fmt.Errorf("EgEmployeeMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgEmployeeMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_employee"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgEmployeeMgr) GetTableName() string {
	return "eg_employee"
}

// Get 获取
func (obj *_EgEmployeeMgr) Get() (result EgEmployee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgEmployeeMgr) Gets() (results []*EgEmployee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDEmployee id_employee获取 
func (obj *_EgEmployeeMgr) WithIDEmployee(idEmployee uint32) Option {
	return optionFunc(func(o *options) { o.query["id_employee"] = idEmployee })
}

// WithIDProfile id_profile获取 
func (obj *_EgEmployeeMgr) WithIDProfile(idProfile uint32) Option {
	return optionFunc(func(o *options) { o.query["id_profile"] = idProfile })
}

// WithIDLang id_lang获取 
func (obj *_EgEmployeeMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

// WithLastname lastname获取 
func (obj *_EgEmployeeMgr) WithLastname(lastname string) Option {
	return optionFunc(func(o *options) { o.query["lastname"] = lastname })
}

// WithFirstname firstname获取 
func (obj *_EgEmployeeMgr) WithFirstname(firstname string) Option {
	return optionFunc(func(o *options) { o.query["firstname"] = firstname })
}

// WithEmail email获取 
func (obj *_EgEmployeeMgr) WithEmail(email string) Option {
	return optionFunc(func(o *options) { o.query["email"] = email })
}

// WithPasswd passwd获取 
func (obj *_EgEmployeeMgr) WithPasswd(passwd string) Option {
	return optionFunc(func(o *options) { o.query["passwd"] = passwd })
}

// WithLastPasswdGen last_passwd_gen获取 
func (obj *_EgEmployeeMgr) WithLastPasswdGen(lastPasswdGen time.Time) Option {
	return optionFunc(func(o *options) { o.query["last_passwd_gen"] = lastPasswdGen })
}

// WithStatsDateFrom stats_date_from获取 
func (obj *_EgEmployeeMgr) WithStatsDateFrom(statsDateFrom datatypes.Date) Option {
	return optionFunc(func(o *options) { o.query["stats_date_from"] = statsDateFrom })
}

// WithStatsDateTo stats_date_to获取 
func (obj *_EgEmployeeMgr) WithStatsDateTo(statsDateTo datatypes.Date) Option {
	return optionFunc(func(o *options) { o.query["stats_date_to"] = statsDateTo })
}

// WithStatsCompareFrom stats_compare_from获取 
func (obj *_EgEmployeeMgr) WithStatsCompareFrom(statsCompareFrom datatypes.Date) Option {
	return optionFunc(func(o *options) { o.query["stats_compare_from"] = statsCompareFrom })
}

// WithStatsCompareTo stats_compare_to获取 
func (obj *_EgEmployeeMgr) WithStatsCompareTo(statsCompareTo datatypes.Date) Option {
	return optionFunc(func(o *options) { o.query["stats_compare_to"] = statsCompareTo })
}

// WithStatsCompareOption stats_compare_option获取 
func (obj *_EgEmployeeMgr) WithStatsCompareOption(statsCompareOption uint32) Option {
	return optionFunc(func(o *options) { o.query["stats_compare_option"] = statsCompareOption })
}

// WithPreselectDateRange preselect_date_range获取 
func (obj *_EgEmployeeMgr) WithPreselectDateRange(preselectDateRange string) Option {
	return optionFunc(func(o *options) { o.query["preselect_date_range"] = preselectDateRange })
}

// WithBoColor bo_color获取 
func (obj *_EgEmployeeMgr) WithBoColor(boColor string) Option {
	return optionFunc(func(o *options) { o.query["bo_color"] = boColor })
}

// WithBoTheme bo_theme获取 
func (obj *_EgEmployeeMgr) WithBoTheme(boTheme string) Option {
	return optionFunc(func(o *options) { o.query["bo_theme"] = boTheme })
}

// WithBoCSS bo_css获取 
func (obj *_EgEmployeeMgr) WithBoCSS(boCSS string) Option {
	return optionFunc(func(o *options) { o.query["bo_css"] = boCSS })
}

// WithDefaultTab default_tab获取 
func (obj *_EgEmployeeMgr) WithDefaultTab(defaultTab uint32) Option {
	return optionFunc(func(o *options) { o.query["default_tab"] = defaultTab })
}

// WithBoWidth bo_width获取 
func (obj *_EgEmployeeMgr) WithBoWidth(boWidth uint32) Option {
	return optionFunc(func(o *options) { o.query["bo_width"] = boWidth })
}

// WithBoMenu bo_menu获取 
func (obj *_EgEmployeeMgr) WithBoMenu(boMenu bool) Option {
	return optionFunc(func(o *options) { o.query["bo_menu"] = boMenu })
}

// WithActive active获取 
func (obj *_EgEmployeeMgr) WithActive(active bool) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}

// WithOptin optin获取 
func (obj *_EgEmployeeMgr) WithOptin(optin bool) Option {
	return optionFunc(func(o *options) { o.query["optin"] = optin })
}

// WithIDLastOrder id_last_order获取 
func (obj *_EgEmployeeMgr) WithIDLastOrder(idLastOrder uint32) Option {
	return optionFunc(func(o *options) { o.query["id_last_order"] = idLastOrder })
}

// WithIDLastCustomerMessage id_last_customer_message获取 
func (obj *_EgEmployeeMgr) WithIDLastCustomerMessage(idLastCustomerMessage uint32) Option {
	return optionFunc(func(o *options) { o.query["id_last_customer_message"] = idLastCustomerMessage })
}

// WithIDLastCustomer id_last_customer获取 
func (obj *_EgEmployeeMgr) WithIDLastCustomer(idLastCustomer uint32) Option {
	return optionFunc(func(o *options) { o.query["id_last_customer"] = idLastCustomer })
}

// WithLastConnectionDate last_connection_date获取 
func (obj *_EgEmployeeMgr) WithLastConnectionDate(lastConnectionDate datatypes.Date) Option {
	return optionFunc(func(o *options) { o.query["last_connection_date"] = lastConnectionDate })
}

// WithResetPasswordToken reset_password_token获取 
func (obj *_EgEmployeeMgr) WithResetPasswordToken(resetPasswordToken string) Option {
	return optionFunc(func(o *options) { o.query["reset_password_token"] = resetPasswordToken })
}

// WithResetPasswordValidity reset_password_validity获取 
func (obj *_EgEmployeeMgr) WithResetPasswordValidity(resetPasswordValidity time.Time) Option {
	return optionFunc(func(o *options) { o.query["reset_password_validity"] = resetPasswordValidity })
}


// GetByOption 功能选项模式获取
func (obj *_EgEmployeeMgr) GetByOption(opts ...Option) (result EgEmployee, err error) {
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
func (obj *_EgEmployeeMgr) GetByOptions(opts ...Option) (results []*EgEmployee, err error) {
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


// GetFromIDEmployee 通过id_employee获取内容  
func (obj *_EgEmployeeMgr)  GetFromIDEmployee(idEmployee uint32) (result EgEmployee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee = ?", idEmployee).Find(&result).Error
	
	return
}

// GetBatchFromIDEmployee 批量唯一主键查找 
func (obj *_EgEmployeeMgr) GetBatchFromIDEmployee(idEmployees []uint32) (results []*EgEmployee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee IN (?)", idEmployees).Find(&results).Error
	
	return
}
 
// GetFromIDProfile 通过id_profile获取内容  
func (obj *_EgEmployeeMgr) GetFromIDProfile(idProfile uint32) (results []*EgEmployee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_profile = ?", idProfile).Find(&results).Error
	
	return
}

// GetBatchFromIDProfile 批量唯一主键查找 
func (obj *_EgEmployeeMgr) GetBatchFromIDProfile(idProfiles []uint32) (results []*EgEmployee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_profile IN (?)", idProfiles).Find(&results).Error
	
	return
}
 
// GetFromIDLang 通过id_lang获取内容  
func (obj *_EgEmployeeMgr) GetFromIDLang(idLang uint32) (results []*EgEmployee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error
	
	return
}

// GetBatchFromIDLang 批量唯一主键查找 
func (obj *_EgEmployeeMgr) GetBatchFromIDLang(idLangs []uint32) (results []*EgEmployee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error
	
	return
}
 
// GetFromLastname 通过lastname获取内容  
func (obj *_EgEmployeeMgr) GetFromLastname(lastname string) (results []*EgEmployee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("lastname = ?", lastname).Find(&results).Error
	
	return
}

// GetBatchFromLastname 批量唯一主键查找 
func (obj *_EgEmployeeMgr) GetBatchFromLastname(lastnames []string) (results []*EgEmployee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("lastname IN (?)", lastnames).Find(&results).Error
	
	return
}
 
// GetFromFirstname 通过firstname获取内容  
func (obj *_EgEmployeeMgr) GetFromFirstname(firstname string) (results []*EgEmployee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("firstname = ?", firstname).Find(&results).Error
	
	return
}

// GetBatchFromFirstname 批量唯一主键查找 
func (obj *_EgEmployeeMgr) GetBatchFromFirstname(firstnames []string) (results []*EgEmployee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("firstname IN (?)", firstnames).Find(&results).Error
	
	return
}
 
// GetFromEmail 通过email获取内容  
func (obj *_EgEmployeeMgr) GetFromEmail(email string) (results []*EgEmployee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("email = ?", email).Find(&results).Error
	
	return
}

// GetBatchFromEmail 批量唯一主键查找 
func (obj *_EgEmployeeMgr) GetBatchFromEmail(emails []string) (results []*EgEmployee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("email IN (?)", emails).Find(&results).Error
	
	return
}
 
// GetFromPasswd 通过passwd获取内容  
func (obj *_EgEmployeeMgr) GetFromPasswd(passwd string) (results []*EgEmployee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("passwd = ?", passwd).Find(&results).Error
	
	return
}

// GetBatchFromPasswd 批量唯一主键查找 
func (obj *_EgEmployeeMgr) GetBatchFromPasswd(passwds []string) (results []*EgEmployee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("passwd IN (?)", passwds).Find(&results).Error
	
	return
}
 
// GetFromLastPasswdGen 通过last_passwd_gen获取内容  
func (obj *_EgEmployeeMgr) GetFromLastPasswdGen(lastPasswdGen time.Time) (results []*EgEmployee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("last_passwd_gen = ?", lastPasswdGen).Find(&results).Error
	
	return
}

// GetBatchFromLastPasswdGen 批量唯一主键查找 
func (obj *_EgEmployeeMgr) GetBatchFromLastPasswdGen(lastPasswdGens []time.Time) (results []*EgEmployee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("last_passwd_gen IN (?)", lastPasswdGens).Find(&results).Error
	
	return
}
 
// GetFromStatsDateFrom 通过stats_date_from获取内容  
func (obj *_EgEmployeeMgr) GetFromStatsDateFrom(statsDateFrom datatypes.Date) (results []*EgEmployee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("stats_date_from = ?", statsDateFrom).Find(&results).Error
	
	return
}

// GetBatchFromStatsDateFrom 批量唯一主键查找 
func (obj *_EgEmployeeMgr) GetBatchFromStatsDateFrom(statsDateFroms []datatypes.Date) (results []*EgEmployee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("stats_date_from IN (?)", statsDateFroms).Find(&results).Error
	
	return
}
 
// GetFromStatsDateTo 通过stats_date_to获取内容  
func (obj *_EgEmployeeMgr) GetFromStatsDateTo(statsDateTo datatypes.Date) (results []*EgEmployee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("stats_date_to = ?", statsDateTo).Find(&results).Error
	
	return
}

// GetBatchFromStatsDateTo 批量唯一主键查找 
func (obj *_EgEmployeeMgr) GetBatchFromStatsDateTo(statsDateTos []datatypes.Date) (results []*EgEmployee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("stats_date_to IN (?)", statsDateTos).Find(&results).Error
	
	return
}
 
// GetFromStatsCompareFrom 通过stats_compare_from获取内容  
func (obj *_EgEmployeeMgr) GetFromStatsCompareFrom(statsCompareFrom datatypes.Date) (results []*EgEmployee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("stats_compare_from = ?", statsCompareFrom).Find(&results).Error
	
	return
}

// GetBatchFromStatsCompareFrom 批量唯一主键查找 
func (obj *_EgEmployeeMgr) GetBatchFromStatsCompareFrom(statsCompareFroms []datatypes.Date) (results []*EgEmployee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("stats_compare_from IN (?)", statsCompareFroms).Find(&results).Error
	
	return
}
 
// GetFromStatsCompareTo 通过stats_compare_to获取内容  
func (obj *_EgEmployeeMgr) GetFromStatsCompareTo(statsCompareTo datatypes.Date) (results []*EgEmployee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("stats_compare_to = ?", statsCompareTo).Find(&results).Error
	
	return
}

// GetBatchFromStatsCompareTo 批量唯一主键查找 
func (obj *_EgEmployeeMgr) GetBatchFromStatsCompareTo(statsCompareTos []datatypes.Date) (results []*EgEmployee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("stats_compare_to IN (?)", statsCompareTos).Find(&results).Error
	
	return
}
 
// GetFromStatsCompareOption 通过stats_compare_option获取内容  
func (obj *_EgEmployeeMgr) GetFromStatsCompareOption(statsCompareOption uint32) (results []*EgEmployee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("stats_compare_option = ?", statsCompareOption).Find(&results).Error
	
	return
}

// GetBatchFromStatsCompareOption 批量唯一主键查找 
func (obj *_EgEmployeeMgr) GetBatchFromStatsCompareOption(statsCompareOptions []uint32) (results []*EgEmployee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("stats_compare_option IN (?)", statsCompareOptions).Find(&results).Error
	
	return
}
 
// GetFromPreselectDateRange 通过preselect_date_range获取内容  
func (obj *_EgEmployeeMgr) GetFromPreselectDateRange(preselectDateRange string) (results []*EgEmployee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("preselect_date_range = ?", preselectDateRange).Find(&results).Error
	
	return
}

// GetBatchFromPreselectDateRange 批量唯一主键查找 
func (obj *_EgEmployeeMgr) GetBatchFromPreselectDateRange(preselectDateRanges []string) (results []*EgEmployee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("preselect_date_range IN (?)", preselectDateRanges).Find(&results).Error
	
	return
}
 
// GetFromBoColor 通过bo_color获取内容  
func (obj *_EgEmployeeMgr) GetFromBoColor(boColor string) (results []*EgEmployee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("bo_color = ?", boColor).Find(&results).Error
	
	return
}

// GetBatchFromBoColor 批量唯一主键查找 
func (obj *_EgEmployeeMgr) GetBatchFromBoColor(boColors []string) (results []*EgEmployee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("bo_color IN (?)", boColors).Find(&results).Error
	
	return
}
 
// GetFromBoTheme 通过bo_theme获取内容  
func (obj *_EgEmployeeMgr) GetFromBoTheme(boTheme string) (results []*EgEmployee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("bo_theme = ?", boTheme).Find(&results).Error
	
	return
}

// GetBatchFromBoTheme 批量唯一主键查找 
func (obj *_EgEmployeeMgr) GetBatchFromBoTheme(boThemes []string) (results []*EgEmployee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("bo_theme IN (?)", boThemes).Find(&results).Error
	
	return
}
 
// GetFromBoCSS 通过bo_css获取内容  
func (obj *_EgEmployeeMgr) GetFromBoCSS(boCSS string) (results []*EgEmployee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("bo_css = ?", boCSS).Find(&results).Error
	
	return
}

// GetBatchFromBoCSS 批量唯一主键查找 
func (obj *_EgEmployeeMgr) GetBatchFromBoCSS(boCSSs []string) (results []*EgEmployee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("bo_css IN (?)", boCSSs).Find(&results).Error
	
	return
}
 
// GetFromDefaultTab 通过default_tab获取内容  
func (obj *_EgEmployeeMgr) GetFromDefaultTab(defaultTab uint32) (results []*EgEmployee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("default_tab = ?", defaultTab).Find(&results).Error
	
	return
}

// GetBatchFromDefaultTab 批量唯一主键查找 
func (obj *_EgEmployeeMgr) GetBatchFromDefaultTab(defaultTabs []uint32) (results []*EgEmployee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("default_tab IN (?)", defaultTabs).Find(&results).Error
	
	return
}
 
// GetFromBoWidth 通过bo_width获取内容  
func (obj *_EgEmployeeMgr) GetFromBoWidth(boWidth uint32) (results []*EgEmployee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("bo_width = ?", boWidth).Find(&results).Error
	
	return
}

// GetBatchFromBoWidth 批量唯一主键查找 
func (obj *_EgEmployeeMgr) GetBatchFromBoWidth(boWidths []uint32) (results []*EgEmployee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("bo_width IN (?)", boWidths).Find(&results).Error
	
	return
}
 
// GetFromBoMenu 通过bo_menu获取内容  
func (obj *_EgEmployeeMgr) GetFromBoMenu(boMenu bool) (results []*EgEmployee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("bo_menu = ?", boMenu).Find(&results).Error
	
	return
}

// GetBatchFromBoMenu 批量唯一主键查找 
func (obj *_EgEmployeeMgr) GetBatchFromBoMenu(boMenus []bool) (results []*EgEmployee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("bo_menu IN (?)", boMenus).Find(&results).Error
	
	return
}
 
// GetFromActive 通过active获取内容  
func (obj *_EgEmployeeMgr) GetFromActive(active bool) (results []*EgEmployee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active = ?", active).Find(&results).Error
	
	return
}

// GetBatchFromActive 批量唯一主键查找 
func (obj *_EgEmployeeMgr) GetBatchFromActive(actives []bool) (results []*EgEmployee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active IN (?)", actives).Find(&results).Error
	
	return
}
 
// GetFromOptin 通过optin获取内容  
func (obj *_EgEmployeeMgr) GetFromOptin(optin bool) (results []*EgEmployee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("optin = ?", optin).Find(&results).Error
	
	return
}

// GetBatchFromOptin 批量唯一主键查找 
func (obj *_EgEmployeeMgr) GetBatchFromOptin(optins []bool) (results []*EgEmployee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("optin IN (?)", optins).Find(&results).Error
	
	return
}
 
// GetFromIDLastOrder 通过id_last_order获取内容  
func (obj *_EgEmployeeMgr) GetFromIDLastOrder(idLastOrder uint32) (results []*EgEmployee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_last_order = ?", idLastOrder).Find(&results).Error
	
	return
}

// GetBatchFromIDLastOrder 批量唯一主键查找 
func (obj *_EgEmployeeMgr) GetBatchFromIDLastOrder(idLastOrders []uint32) (results []*EgEmployee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_last_order IN (?)", idLastOrders).Find(&results).Error
	
	return
}
 
// GetFromIDLastCustomerMessage 通过id_last_customer_message获取内容  
func (obj *_EgEmployeeMgr) GetFromIDLastCustomerMessage(idLastCustomerMessage uint32) (results []*EgEmployee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_last_customer_message = ?", idLastCustomerMessage).Find(&results).Error
	
	return
}

// GetBatchFromIDLastCustomerMessage 批量唯一主键查找 
func (obj *_EgEmployeeMgr) GetBatchFromIDLastCustomerMessage(idLastCustomerMessages []uint32) (results []*EgEmployee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_last_customer_message IN (?)", idLastCustomerMessages).Find(&results).Error
	
	return
}
 
// GetFromIDLastCustomer 通过id_last_customer获取内容  
func (obj *_EgEmployeeMgr) GetFromIDLastCustomer(idLastCustomer uint32) (results []*EgEmployee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_last_customer = ?", idLastCustomer).Find(&results).Error
	
	return
}

// GetBatchFromIDLastCustomer 批量唯一主键查找 
func (obj *_EgEmployeeMgr) GetBatchFromIDLastCustomer(idLastCustomers []uint32) (results []*EgEmployee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_last_customer IN (?)", idLastCustomers).Find(&results).Error
	
	return
}
 
// GetFromLastConnectionDate 通过last_connection_date获取内容  
func (obj *_EgEmployeeMgr) GetFromLastConnectionDate(lastConnectionDate datatypes.Date) (results []*EgEmployee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("last_connection_date = ?", lastConnectionDate).Find(&results).Error
	
	return
}

// GetBatchFromLastConnectionDate 批量唯一主键查找 
func (obj *_EgEmployeeMgr) GetBatchFromLastConnectionDate(lastConnectionDates []datatypes.Date) (results []*EgEmployee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("last_connection_date IN (?)", lastConnectionDates).Find(&results).Error
	
	return
}
 
// GetFromResetPasswordToken 通过reset_password_token获取内容  
func (obj *_EgEmployeeMgr) GetFromResetPasswordToken(resetPasswordToken string) (results []*EgEmployee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reset_password_token = ?", resetPasswordToken).Find(&results).Error
	
	return
}

// GetBatchFromResetPasswordToken 批量唯一主键查找 
func (obj *_EgEmployeeMgr) GetBatchFromResetPasswordToken(resetPasswordTokens []string) (results []*EgEmployee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reset_password_token IN (?)", resetPasswordTokens).Find(&results).Error
	
	return
}
 
// GetFromResetPasswordValidity 通过reset_password_validity获取内容  
func (obj *_EgEmployeeMgr) GetFromResetPasswordValidity(resetPasswordValidity time.Time) (results []*EgEmployee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reset_password_validity = ?", resetPasswordValidity).Find(&results).Error
	
	return
}

// GetBatchFromResetPasswordValidity 批量唯一主键查找 
func (obj *_EgEmployeeMgr) GetBatchFromResetPasswordValidity(resetPasswordValiditys []time.Time) (results []*EgEmployee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reset_password_validity IN (?)", resetPasswordValiditys).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgEmployeeMgr) FetchByPrimaryKey(idEmployee uint32 ) (result EgEmployee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee = ?", idEmployee).Find(&result).Error
	
	return
}
 

 
 // FetchIndexByIDEmployeePasswd  获取多个内容
 func (obj *_EgEmployeeMgr) FetchIndexByIDEmployeePasswd(idEmployee uint32 ,passwd string ) (results []*EgEmployee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee = ? AND passwd = ?", idEmployee , passwd).Find(&results).Error
	
	return
}
 
 // FetchIndexByIDProfile  获取多个内容
 func (obj *_EgEmployeeMgr) FetchIndexByIDProfile(idProfile uint32 ) (results []*EgEmployee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_profile = ?", idProfile).Find(&results).Error
	
	return
}
 
 // FetchIndexByEmployeeLogin  获取多个内容
 func (obj *_EgEmployeeMgr) FetchIndexByEmployeeLogin(email string ,passwd string ) (results []*EgEmployee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("email = ? AND passwd = ?", email , passwd).Find(&results).Error
	
	return
}
 

	
