package models

import (
	"context"
	"fmt"
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type _EmployeeMgr struct {
	*_BaseMgr
}

func EmployeeMgr(db *gorm.DB) *_EmployeeMgr {
	if db == nil {
		panic(fmt.Errorf("EmployeeMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EmployeeMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_employee"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_EmployeeMgr) GetTableName() string {
	return "ps_employee"
}

func (obj *_EmployeeMgr) Get() (result Employee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_EmployeeMgr) Gets() (results []*Employee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

func (obj *_EmployeeMgr) WithIDEmployee(idEmployee uint32) Option {
	return optionFunc(func(o *options) { o.query["id_employee"] = idEmployee })
}

func (obj *_EmployeeMgr) WithIDProfile(idProfile uint32) Option {
	return optionFunc(func(o *options) { o.query["id_profile"] = idProfile })
}

func (obj *_EmployeeMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

func (obj *_EmployeeMgr) WithLastname(lastname string) Option {
	return optionFunc(func(o *options) { o.query["lastname"] = lastname })
}

func (obj *_EmployeeMgr) WithFirstname(firstname string) Option {
	return optionFunc(func(o *options) { o.query["firstname"] = firstname })
}

func (obj *_EmployeeMgr) WithEmail(email string) Option {
	return optionFunc(func(o *options) { o.query["email"] = email })
}

func (obj *_EmployeeMgr) WithPasswd(passwd string) Option {
	return optionFunc(func(o *options) { o.query["passwd"] = passwd })
}

func (obj *_EmployeeMgr) WithLastPasswdGen(lastPasswdGen time.Time) Option {
	return optionFunc(func(o *options) { o.query["last_passwd_gen"] = lastPasswdGen })
}

func (obj *_EmployeeMgr) WithStatsDateFrom(statsDateFrom datatypes.Date) Option {
	return optionFunc(func(o *options) { o.query["stats_date_from"] = statsDateFrom })
}

func (obj *_EmployeeMgr) WithStatsDateTo(statsDateTo datatypes.Date) Option {
	return optionFunc(func(o *options) { o.query["stats_date_to"] = statsDateTo })
}

func (obj *_EmployeeMgr) WithStatsCompareFrom(statsCompareFrom datatypes.Date) Option {
	return optionFunc(func(o *options) { o.query["stats_compare_from"] = statsCompareFrom })
}

func (obj *_EmployeeMgr) WithStatsCompareTo(statsCompareTo datatypes.Date) Option {
	return optionFunc(func(o *options) { o.query["stats_compare_to"] = statsCompareTo })
}

func (obj *_EmployeeMgr) WithStatsCompareOption(statsCompareOption uint32) Option {
	return optionFunc(func(o *options) { o.query["stats_compare_option"] = statsCompareOption })
}

func (obj *_EmployeeMgr) WithPreselectDateRange(preselectDateRange string) Option {
	return optionFunc(func(o *options) { o.query["preselect_date_range"] = preselectDateRange })
}

func (obj *_EmployeeMgr) WithBoColor(boColor string) Option {
	return optionFunc(func(o *options) { o.query["bo_color"] = boColor })
}

func (obj *_EmployeeMgr) WithBoTheme(boTheme string) Option {
	return optionFunc(func(o *options) { o.query["bo_theme"] = boTheme })
}

func (obj *_EmployeeMgr) WithBoCSS(boCSS string) Option {
	return optionFunc(func(o *options) { o.query["bo_css"] = boCSS })
}

func (obj *_EmployeeMgr) WithDefaultTab(defaultTab uint32) Option {
	return optionFunc(func(o *options) { o.query["default_tab"] = defaultTab })
}

func (obj *_EmployeeMgr) WithBoWidth(boWidth uint32) Option {
	return optionFunc(func(o *options) { o.query["bo_width"] = boWidth })
}

func (obj *_EmployeeMgr) WithBoMenu(boMenu bool) Option {
	return optionFunc(func(o *options) { o.query["bo_menu"] = boMenu })
}

func (obj *_EmployeeMgr) WithActive(active bool) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}

func (obj *_EmployeeMgr) WithOptin(optin bool) Option {
	return optionFunc(func(o *options) { o.query["optin"] = optin })
}

func (obj *_EmployeeMgr) WithIDLastOrder(idLastOrder uint32) Option {
	return optionFunc(func(o *options) { o.query["id_last_order"] = idLastOrder })
}

func (obj *_EmployeeMgr) WithIDLastCustomerMessage(idLastCustomerMessage uint32) Option {
	return optionFunc(func(o *options) { o.query["id_last_customer_message"] = idLastCustomerMessage })
}

func (obj *_EmployeeMgr) WithIDLastCustomer(idLastCustomer uint32) Option {
	return optionFunc(func(o *options) { o.query["id_last_customer"] = idLastCustomer })
}

func (obj *_EmployeeMgr) WithLastConnectionDate(lastConnectionDate datatypes.Date) Option {
	return optionFunc(func(o *options) { o.query["last_connection_date"] = lastConnectionDate })
}

func (obj *_EmployeeMgr) WithResetPasswordToken(resetPasswordToken string) Option {
	return optionFunc(func(o *options) { o.query["reset_password_token"] = resetPasswordToken })
}

func (obj *_EmployeeMgr) WithResetPasswordValidity(resetPasswordValidity time.Time) Option {
	return optionFunc(func(o *options) { o.query["reset_password_validity"] = resetPasswordValidity })
}

func (obj *_EmployeeMgr) GetByOption(opts ...Option) (result Employee, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_EmployeeMgr) GetByOptions(opts ...Option) (results []*Employee, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

func (obj *_EmployeeMgr) GetFromIDEmployee(idEmployee uint32) (result Employee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee = ?", idEmployee).Find(&result).Error

	return
}

func (obj *_EmployeeMgr) GetBatchFromIDEmployee(idEmployees []uint32) (results []*Employee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee IN (?)", idEmployees).Find(&results).Error

	return
}

func (obj *_EmployeeMgr) GetFromIDProfile(idProfile uint32) (results []*Employee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_profile = ?", idProfile).Find(&results).Error

	return
}

func (obj *_EmployeeMgr) GetBatchFromIDProfile(idProfiles []uint32) (results []*Employee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_profile IN (?)", idProfiles).Find(&results).Error

	return
}

func (obj *_EmployeeMgr) GetFromIDLang(idLang uint32) (results []*Employee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error

	return
}

func (obj *_EmployeeMgr) GetBatchFromIDLang(idLangs []uint32) (results []*Employee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error

	return
}

func (obj *_EmployeeMgr) GetFromLastname(lastname string) (results []*Employee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("lastname = ?", lastname).Find(&results).Error

	return
}

func (obj *_EmployeeMgr) GetBatchFromLastname(lastnames []string) (results []*Employee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("lastname IN (?)", lastnames).Find(&results).Error

	return
}

func (obj *_EmployeeMgr) GetFromFirstname(firstname string) (results []*Employee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("firstname = ?", firstname).Find(&results).Error

	return
}

func (obj *_EmployeeMgr) GetBatchFromFirstname(firstnames []string) (results []*Employee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("firstname IN (?)", firstnames).Find(&results).Error

	return
}

func (obj *_EmployeeMgr) GetFromEmail(email string) (results []*Employee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("email = ?", email).Find(&results).Error

	return
}

func (obj *_EmployeeMgr) GetBatchFromEmail(emails []string) (results []*Employee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("email IN (?)", emails).Find(&results).Error

	return
}

func (obj *_EmployeeMgr) GetFromPasswd(passwd string) (results []*Employee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("passwd = ?", passwd).Find(&results).Error

	return
}

func (obj *_EmployeeMgr) GetBatchFromPasswd(passwds []string) (results []*Employee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("passwd IN (?)", passwds).Find(&results).Error

	return
}

func (obj *_EmployeeMgr) GetFromLastPasswdGen(lastPasswdGen time.Time) (results []*Employee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("last_passwd_gen = ?", lastPasswdGen).Find(&results).Error

	return
}

func (obj *_EmployeeMgr) GetBatchFromLastPasswdGen(lastPasswdGens []time.Time) (results []*Employee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("last_passwd_gen IN (?)", lastPasswdGens).Find(&results).Error

	return
}

func (obj *_EmployeeMgr) GetFromStatsDateFrom(statsDateFrom datatypes.Date) (results []*Employee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("stats_date_from = ?", statsDateFrom).Find(&results).Error

	return
}

func (obj *_EmployeeMgr) GetBatchFromStatsDateFrom(statsDateFroms []datatypes.Date) (results []*Employee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("stats_date_from IN (?)", statsDateFroms).Find(&results).Error

	return
}

func (obj *_EmployeeMgr) GetFromStatsDateTo(statsDateTo datatypes.Date) (results []*Employee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("stats_date_to = ?", statsDateTo).Find(&results).Error

	return
}

func (obj *_EmployeeMgr) GetBatchFromStatsDateTo(statsDateTos []datatypes.Date) (results []*Employee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("stats_date_to IN (?)", statsDateTos).Find(&results).Error

	return
}

func (obj *_EmployeeMgr) GetFromStatsCompareFrom(statsCompareFrom datatypes.Date) (results []*Employee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("stats_compare_from = ?", statsCompareFrom).Find(&results).Error

	return
}

func (obj *_EmployeeMgr) GetBatchFromStatsCompareFrom(statsCompareFroms []datatypes.Date) (results []*Employee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("stats_compare_from IN (?)", statsCompareFroms).Find(&results).Error

	return
}

func (obj *_EmployeeMgr) GetFromStatsCompareTo(statsCompareTo datatypes.Date) (results []*Employee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("stats_compare_to = ?", statsCompareTo).Find(&results).Error

	return
}

func (obj *_EmployeeMgr) GetBatchFromStatsCompareTo(statsCompareTos []datatypes.Date) (results []*Employee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("stats_compare_to IN (?)", statsCompareTos).Find(&results).Error

	return
}

func (obj *_EmployeeMgr) GetFromStatsCompareOption(statsCompareOption uint32) (results []*Employee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("stats_compare_option = ?", statsCompareOption).Find(&results).Error

	return
}

func (obj *_EmployeeMgr) GetBatchFromStatsCompareOption(statsCompareOptions []uint32) (results []*Employee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("stats_compare_option IN (?)", statsCompareOptions).Find(&results).Error

	return
}

func (obj *_EmployeeMgr) GetFromPreselectDateRange(preselectDateRange string) (results []*Employee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("preselect_date_range = ?", preselectDateRange).Find(&results).Error

	return
}

func (obj *_EmployeeMgr) GetBatchFromPreselectDateRange(preselectDateRanges []string) (results []*Employee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("preselect_date_range IN (?)", preselectDateRanges).Find(&results).Error

	return
}

func (obj *_EmployeeMgr) GetFromBoColor(boColor string) (results []*Employee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("bo_color = ?", boColor).Find(&results).Error

	return
}

func (obj *_EmployeeMgr) GetBatchFromBoColor(boColors []string) (results []*Employee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("bo_color IN (?)", boColors).Find(&results).Error

	return
}

func (obj *_EmployeeMgr) GetFromBoTheme(boTheme string) (results []*Employee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("bo_theme = ?", boTheme).Find(&results).Error

	return
}

func (obj *_EmployeeMgr) GetBatchFromBoTheme(boThemes []string) (results []*Employee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("bo_theme IN (?)", boThemes).Find(&results).Error

	return
}

func (obj *_EmployeeMgr) GetFromBoCSS(boCSS string) (results []*Employee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("bo_css = ?", boCSS).Find(&results).Error

	return
}

func (obj *_EmployeeMgr) GetBatchFromBoCSS(boCSSs []string) (results []*Employee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("bo_css IN (?)", boCSSs).Find(&results).Error

	return
}

func (obj *_EmployeeMgr) GetFromDefaultTab(defaultTab uint32) (results []*Employee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("default_tab = ?", defaultTab).Find(&results).Error

	return
}

func (obj *_EmployeeMgr) GetBatchFromDefaultTab(defaultTabs []uint32) (results []*Employee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("default_tab IN (?)", defaultTabs).Find(&results).Error

	return
}

func (obj *_EmployeeMgr) GetFromBoWidth(boWidth uint32) (results []*Employee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("bo_width = ?", boWidth).Find(&results).Error

	return
}

func (obj *_EmployeeMgr) GetBatchFromBoWidth(boWidths []uint32) (results []*Employee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("bo_width IN (?)", boWidths).Find(&results).Error

	return
}

func (obj *_EmployeeMgr) GetFromBoMenu(boMenu bool) (results []*Employee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("bo_menu = ?", boMenu).Find(&results).Error

	return
}

func (obj *_EmployeeMgr) GetBatchFromBoMenu(boMenus []bool) (results []*Employee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("bo_menu IN (?)", boMenus).Find(&results).Error

	return
}

func (obj *_EmployeeMgr) GetFromActive(active bool) (results []*Employee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active = ?", active).Find(&results).Error

	return
}

func (obj *_EmployeeMgr) GetBatchFromActive(actives []bool) (results []*Employee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active IN (?)", actives).Find(&results).Error

	return
}

func (obj *_EmployeeMgr) GetFromOptin(optin bool) (results []*Employee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("optin = ?", optin).Find(&results).Error

	return
}

func (obj *_EmployeeMgr) GetBatchFromOptin(optins []bool) (results []*Employee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("optin IN (?)", optins).Find(&results).Error

	return
}

func (obj *_EmployeeMgr) GetFromIDLastOrder(idLastOrder uint32) (results []*Employee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_last_order = ?", idLastOrder).Find(&results).Error

	return
}

func (obj *_EmployeeMgr) GetBatchFromIDLastOrder(idLastOrders []uint32) (results []*Employee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_last_order IN (?)", idLastOrders).Find(&results).Error

	return
}

func (obj *_EmployeeMgr) GetFromIDLastCustomerMessage(idLastCustomerMessage uint32) (results []*Employee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_last_customer_message = ?", idLastCustomerMessage).Find(&results).Error

	return
}

func (obj *_EmployeeMgr) GetBatchFromIDLastCustomerMessage(idLastCustomerMessages []uint32) (results []*Employee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_last_customer_message IN (?)", idLastCustomerMessages).Find(&results).Error

	return
}

func (obj *_EmployeeMgr) GetFromIDLastCustomer(idLastCustomer uint32) (results []*Employee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_last_customer = ?", idLastCustomer).Find(&results).Error

	return
}

func (obj *_EmployeeMgr) GetBatchFromIDLastCustomer(idLastCustomers []uint32) (results []*Employee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_last_customer IN (?)", idLastCustomers).Find(&results).Error

	return
}

func (obj *_EmployeeMgr) GetFromLastConnectionDate(lastConnectionDate datatypes.Date) (results []*Employee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("last_connection_date = ?", lastConnectionDate).Find(&results).Error

	return
}

func (obj *_EmployeeMgr) GetBatchFromLastConnectionDate(lastConnectionDates []datatypes.Date) (results []*Employee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("last_connection_date IN (?)", lastConnectionDates).Find(&results).Error

	return
}

func (obj *_EmployeeMgr) GetFromResetPasswordToken(resetPasswordToken string) (results []*Employee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reset_password_token = ?", resetPasswordToken).Find(&results).Error

	return
}

func (obj *_EmployeeMgr) GetBatchFromResetPasswordToken(resetPasswordTokens []string) (results []*Employee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reset_password_token IN (?)", resetPasswordTokens).Find(&results).Error

	return
}

func (obj *_EmployeeMgr) GetFromResetPasswordValidity(resetPasswordValidity time.Time) (results []*Employee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reset_password_validity = ?", resetPasswordValidity).Find(&results).Error

	return
}

func (obj *_EmployeeMgr) GetBatchFromResetPasswordValidity(resetPasswordValiditys []time.Time) (results []*Employee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reset_password_validity IN (?)", resetPasswordValiditys).Find(&results).Error

	return
}

func (obj *_EmployeeMgr) FetchByPrimaryKey(idEmployee uint32) (result Employee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee = ?", idEmployee).Find(&result).Error

	return
}

func (obj *_EmployeeMgr) FetchIndexByIDEmployeePasswd(idEmployee uint32, passwd string) (results []*Employee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee = ? AND passwd = ?", idEmployee, passwd).Find(&results).Error

	return
}

func (obj *_EmployeeMgr) FetchIndexByIDProfile(idProfile uint32) (results []*Employee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_profile = ?", idProfile).Find(&results).Error

	return
}

func (obj *_EmployeeMgr) FetchIndexByEmployeeLogin(email string, passwd string) (results []*Employee, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("email = ? AND passwd = ?", email, passwd).Find(&results).Error

	return
}
