package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _GuestMgr struct {
	*_BaseMgr
}

func GuestMgr(db *gorm.DB) *_GuestMgr {
	if db == nil {
		panic(fmt.Errorf("GuestMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_GuestMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_guest"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_GuestMgr) GetTableName() string {
	return "ps_guest"
}

func (obj *_GuestMgr) Get() (result Guest, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_GuestMgr) Gets() (results []*Guest, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_GuestMgr) WithIDGuest(idGuest uint32) Option {
	return optionFunc(func(o *options) { o.query["id_guest"] = idGuest })
}

func (obj *_GuestMgr) WithIDOperatingSystem(idOperatingSystem uint32) Option {
	return optionFunc(func(o *options) { o.query["id_operating_system"] = idOperatingSystem })
}

func (obj *_GuestMgr) WithIDWebBrowser(idWebBrowser uint32) Option {
	return optionFunc(func(o *options) { o.query["id_web_browser"] = idWebBrowser })
}

func (obj *_GuestMgr) WithIDCustomer(idCustomer uint32) Option {
	return optionFunc(func(o *options) { o.query["id_customer"] = idCustomer })
}

func (obj *_GuestMgr) WithJavascript(javascript bool) Option {
	return optionFunc(func(o *options) { o.query["javascript"] = javascript })
}

func (obj *_GuestMgr) WithScreenResolutionX(screenResolutionX uint16) Option {
	return optionFunc(func(o *options) { o.query["screen_resolution_x"] = screenResolutionX })
}

func (obj *_GuestMgr) WithScreenResolutionY(screenResolutionY uint16) Option {
	return optionFunc(func(o *options) { o.query["screen_resolution_y"] = screenResolutionY })
}

func (obj *_GuestMgr) WithScreenColor(screenColor uint8) Option {
	return optionFunc(func(o *options) { o.query["screen_color"] = screenColor })
}

func (obj *_GuestMgr) WithSunJava(sunJava bool) Option {
	return optionFunc(func(o *options) { o.query["sun_java"] = sunJava })
}

func (obj *_GuestMgr) WithAdobeFlash(adobeFlash bool) Option {
	return optionFunc(func(o *options) { o.query["adobe_flash"] = adobeFlash })
}

func (obj *_GuestMgr) WithAdobeDirector(adobeDirector bool) Option {
	return optionFunc(func(o *options) { o.query["adobe_director"] = adobeDirector })
}

func (obj *_GuestMgr) WithAppleQuicktime(appleQuicktime bool) Option {
	return optionFunc(func(o *options) { o.query["apple_quicktime"] = appleQuicktime })
}

func (obj *_GuestMgr) WithRealPlayer(realPlayer bool) Option {
	return optionFunc(func(o *options) { o.query["real_player"] = realPlayer })
}

func (obj *_GuestMgr) WithWindowsMedia(windowsMedia bool) Option {
	return optionFunc(func(o *options) { o.query["windows_media"] = windowsMedia })
}

func (obj *_GuestMgr) WithAcceptLanguage(acceptLanguage string) Option {
	return optionFunc(func(o *options) { o.query["accept_language"] = acceptLanguage })
}

func (obj *_GuestMgr) WithMobileTheme(mobileTheme bool) Option {
	return optionFunc(func(o *options) { o.query["mobile_theme"] = mobileTheme })
}

func (obj *_GuestMgr) GetByOption(opts ...Option) (result Guest, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_GuestMgr) GetByOptions(opts ...Option) (results []*Guest, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_GuestMgr) GetFromIDGuest(idGuest uint32) (result Guest, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_guest = ?", idGuest).Find(&result).Error

	return
}

func (obj *_GuestMgr) GetBatchFromIDGuest(idGuests []uint32) (results []*Guest, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_guest IN (?)", idGuests).Find(&results).Error

	return
}

func (obj *_GuestMgr) GetFromIDOperatingSystem(idOperatingSystem uint32) (results []*Guest, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_operating_system = ?", idOperatingSystem).Find(&results).Error

	return
}

func (obj *_GuestMgr) GetBatchFromIDOperatingSystem(idOperatingSystems []uint32) (results []*Guest, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_operating_system IN (?)", idOperatingSystems).Find(&results).Error

	return
}

func (obj *_GuestMgr) GetFromIDWebBrowser(idWebBrowser uint32) (results []*Guest, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_web_browser = ?", idWebBrowser).Find(&results).Error

	return
}

func (obj *_GuestMgr) GetBatchFromIDWebBrowser(idWebBrowsers []uint32) (results []*Guest, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_web_browser IN (?)", idWebBrowsers).Find(&results).Error

	return
}

func (obj *_GuestMgr) GetFromIDCustomer(idCustomer uint32) (results []*Guest, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer = ?", idCustomer).Find(&results).Error

	return
}

func (obj *_GuestMgr) GetBatchFromIDCustomer(idCustomers []uint32) (results []*Guest, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer IN (?)", idCustomers).Find(&results).Error

	return
}

func (obj *_GuestMgr) GetFromJavascript(javascript bool) (results []*Guest, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("javascript = ?", javascript).Find(&results).Error

	return
}

func (obj *_GuestMgr) GetBatchFromJavascript(javascripts []bool) (results []*Guest, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("javascript IN (?)", javascripts).Find(&results).Error

	return
}

func (obj *_GuestMgr) GetFromScreenResolutionX(screenResolutionX uint16) (results []*Guest, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("screen_resolution_x = ?", screenResolutionX).Find(&results).Error

	return
}

func (obj *_GuestMgr) GetBatchFromScreenResolutionX(screenResolutionXs []uint16) (results []*Guest, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("screen_resolution_x IN (?)", screenResolutionXs).Find(&results).Error

	return
}

func (obj *_GuestMgr) GetFromScreenResolutionY(screenResolutionY uint16) (results []*Guest, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("screen_resolution_y = ?", screenResolutionY).Find(&results).Error

	return
}

func (obj *_GuestMgr) GetBatchFromScreenResolutionY(screenResolutionYs []uint16) (results []*Guest, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("screen_resolution_y IN (?)", screenResolutionYs).Find(&results).Error

	return
}

func (obj *_GuestMgr) GetFromScreenColor(screenColor uint8) (results []*Guest, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("screen_color = ?", screenColor).Find(&results).Error

	return
}

func (obj *_GuestMgr) GetBatchFromScreenColor(screenColors []uint8) (results []*Guest, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("screen_color IN (?)", screenColors).Find(&results).Error

	return
}

func (obj *_GuestMgr) GetFromSunJava(sunJava bool) (results []*Guest, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("sun_java = ?", sunJava).Find(&results).Error

	return
}

func (obj *_GuestMgr) GetBatchFromSunJava(sunJavas []bool) (results []*Guest, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("sun_java IN (?)", sunJavas).Find(&results).Error

	return
}

func (obj *_GuestMgr) GetFromAdobeFlash(adobeFlash bool) (results []*Guest, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("adobe_flash = ?", adobeFlash).Find(&results).Error

	return
}

func (obj *_GuestMgr) GetBatchFromAdobeFlash(adobeFlashs []bool) (results []*Guest, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("adobe_flash IN (?)", adobeFlashs).Find(&results).Error

	return
}

func (obj *_GuestMgr) GetFromAdobeDirector(adobeDirector bool) (results []*Guest, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("adobe_director = ?", adobeDirector).Find(&results).Error

	return
}

func (obj *_GuestMgr) GetBatchFromAdobeDirector(adobeDirectors []bool) (results []*Guest, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("adobe_director IN (?)", adobeDirectors).Find(&results).Error

	return
}

func (obj *_GuestMgr) GetFromAppleQuicktime(appleQuicktime bool) (results []*Guest, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("apple_quicktime = ?", appleQuicktime).Find(&results).Error

	return
}

func (obj *_GuestMgr) GetBatchFromAppleQuicktime(appleQuicktimes []bool) (results []*Guest, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("apple_quicktime IN (?)", appleQuicktimes).Find(&results).Error

	return
}

func (obj *_GuestMgr) GetFromRealPlayer(realPlayer bool) (results []*Guest, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("real_player = ?", realPlayer).Find(&results).Error

	return
}

func (obj *_GuestMgr) GetBatchFromRealPlayer(realPlayers []bool) (results []*Guest, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("real_player IN (?)", realPlayers).Find(&results).Error

	return
}

func (obj *_GuestMgr) GetFromWindowsMedia(windowsMedia bool) (results []*Guest, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("windows_media = ?", windowsMedia).Find(&results).Error

	return
}

func (obj *_GuestMgr) GetBatchFromWindowsMedia(windowsMedias []bool) (results []*Guest, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("windows_media IN (?)", windowsMedias).Find(&results).Error

	return
}

func (obj *_GuestMgr) GetFromAcceptLanguage(acceptLanguage string) (results []*Guest, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("accept_language = ?", acceptLanguage).Find(&results).Error

	return
}

func (obj *_GuestMgr) GetBatchFromAcceptLanguage(acceptLanguages []string) (results []*Guest, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("accept_language IN (?)", acceptLanguages).Find(&results).Error

	return
}

func (obj *_GuestMgr) GetFromMobileTheme(mobileTheme bool) (results []*Guest, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("mobile_theme = ?", mobileTheme).Find(&results).Error

	return
}

func (obj *_GuestMgr) GetBatchFromMobileTheme(mobileThemes []bool) (results []*Guest, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("mobile_theme IN (?)", mobileThemes).Find(&results).Error

	return
}


func (obj *_GuestMgr) FetchByPrimaryKey(idGuest uint32) (result Guest, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_guest = ?", idGuest).Find(&result).Error

	return
}

func (obj *_GuestMgr) FetchIndexByIDOperatingSystem(idOperatingSystem uint32) (results []*Guest, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_operating_system = ?", idOperatingSystem).Find(&results).Error

	return
}

func (obj *_GuestMgr) FetchIndexByIDWebBrowser(idWebBrowser uint32) (results []*Guest, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_web_browser = ?", idWebBrowser).Find(&results).Error

	return
}

func (obj *_GuestMgr) FetchIndexByIDCustomer(idCustomer uint32) (results []*Guest, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer = ?", idCustomer).Find(&results).Error

	return
}
