package	model	
import (	
"context"	
"gorm.io/gorm"	
"time"	
"fmt"	
)	

type _EgCartMgr struct {
	*_BaseMgr
}

// EgCartMgr open func
func EgCartMgr(db *gorm.DB) *_EgCartMgr {
	if db == nil {
		panic(fmt.Errorf("EgCartMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgCartMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_cart"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgCartMgr) GetTableName() string {
	return "eg_cart"
}

// Get 获取
func (obj *_EgCartMgr) Get() (result EgCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgCartMgr) Gets() (results []*EgCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDCart id_cart获取 
func (obj *_EgCartMgr) WithIDCart(idCart uint32) Option {
	return optionFunc(func(o *options) { o.query["id_cart"] = idCart })
}

// WithIDShopGroup id_shop_group获取 
func (obj *_EgCartMgr) WithIDShopGroup(idShopGroup uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop_group"] = idShopGroup })
}

// WithIDShop id_shop获取 
func (obj *_EgCartMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

// WithIDCarrier id_carrier获取 
func (obj *_EgCartMgr) WithIDCarrier(idCarrier uint32) Option {
	return optionFunc(func(o *options) { o.query["id_carrier"] = idCarrier })
}

// WithDeliveryOption delivery_option获取 
func (obj *_EgCartMgr) WithDeliveryOption(deliveryOption string) Option {
	return optionFunc(func(o *options) { o.query["delivery_option"] = deliveryOption })
}

// WithIDLang id_lang获取 
func (obj *_EgCartMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

// WithIDAddressDelivery id_address_delivery获取 
func (obj *_EgCartMgr) WithIDAddressDelivery(idAddressDelivery uint32) Option {
	return optionFunc(func(o *options) { o.query["id_address_delivery"] = idAddressDelivery })
}

// WithIDAddressInvoice id_address_invoice获取 
func (obj *_EgCartMgr) WithIDAddressInvoice(idAddressInvoice uint32) Option {
	return optionFunc(func(o *options) { o.query["id_address_invoice"] = idAddressInvoice })
}

// WithIDCurrency id_currency获取 
func (obj *_EgCartMgr) WithIDCurrency(idCurrency uint32) Option {
	return optionFunc(func(o *options) { o.query["id_currency"] = idCurrency })
}

// WithIDCustomer id_customer获取 
func (obj *_EgCartMgr) WithIDCustomer(idCustomer uint32) Option {
	return optionFunc(func(o *options) { o.query["id_customer"] = idCustomer })
}

// WithIDGuest id_guest获取 
func (obj *_EgCartMgr) WithIDGuest(idGuest uint32) Option {
	return optionFunc(func(o *options) { o.query["id_guest"] = idGuest })
}

// WithSecureKey secure_key获取 
func (obj *_EgCartMgr) WithSecureKey(secureKey string) Option {
	return optionFunc(func(o *options) { o.query["secure_key"] = secureKey })
}

// WithRecyclable recyclable获取 
func (obj *_EgCartMgr) WithRecyclable(recyclable bool) Option {
	return optionFunc(func(o *options) { o.query["recyclable"] = recyclable })
}

// WithGift gift获取 
func (obj *_EgCartMgr) WithGift(gift bool) Option {
	return optionFunc(func(o *options) { o.query["gift"] = gift })
}

// WithGiftMessage gift_message获取 
func (obj *_EgCartMgr) WithGiftMessage(giftMessage string) Option {
	return optionFunc(func(o *options) { o.query["gift_message"] = giftMessage })
}

// WithMobileTheme mobile_theme获取 
func (obj *_EgCartMgr) WithMobileTheme(mobileTheme bool) Option {
	return optionFunc(func(o *options) { o.query["mobile_theme"] = mobileTheme })
}

// WithAllowSeperatedPackage allow_seperated_package获取 
func (obj *_EgCartMgr) WithAllowSeperatedPackage(allowSeperatedPackage bool) Option {
	return optionFunc(func(o *options) { o.query["allow_seperated_package"] = allowSeperatedPackage })
}

// WithDateAdd date_add获取 
func (obj *_EgCartMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

// WithDateUpd date_upd获取 
func (obj *_EgCartMgr) WithDateUpd(dateUpd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_upd"] = dateUpd })
}

// WithCheckoutSessionData checkout_session_data获取 
func (obj *_EgCartMgr) WithCheckoutSessionData(checkoutSessionData string) Option {
	return optionFunc(func(o *options) { o.query["checkout_session_data"] = checkoutSessionData })
}


// GetByOption 功能选项模式获取
func (obj *_EgCartMgr) GetByOption(opts ...Option) (result EgCart, err error) {
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
func (obj *_EgCartMgr) GetByOptions(opts ...Option) (results []*EgCart, err error) {
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


// GetFromIDCart 通过id_cart获取内容  
func (obj *_EgCartMgr)  GetFromIDCart(idCart uint32) (result EgCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart = ?", idCart).Find(&result).Error
	
	return
}

// GetBatchFromIDCart 批量唯一主键查找 
func (obj *_EgCartMgr) GetBatchFromIDCart(idCarts []uint32) (results []*EgCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart IN (?)", idCarts).Find(&results).Error
	
	return
}
 
// GetFromIDShopGroup 通过id_shop_group获取内容  
func (obj *_EgCartMgr) GetFromIDShopGroup(idShopGroup uint32) (results []*EgCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_group = ?", idShopGroup).Find(&results).Error
	
	return
}

// GetBatchFromIDShopGroup 批量唯一主键查找 
func (obj *_EgCartMgr) GetBatchFromIDShopGroup(idShopGroups []uint32) (results []*EgCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_group IN (?)", idShopGroups).Find(&results).Error
	
	return
}
 
// GetFromIDShop 通过id_shop获取内容  
func (obj *_EgCartMgr) GetFromIDShop(idShop uint32) (results []*EgCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error
	
	return
}

// GetBatchFromIDShop 批量唯一主键查找 
func (obj *_EgCartMgr) GetBatchFromIDShop(idShops []uint32) (results []*EgCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error
	
	return
}
 
// GetFromIDCarrier 通过id_carrier获取内容  
func (obj *_EgCartMgr) GetFromIDCarrier(idCarrier uint32) (results []*EgCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_carrier = ?", idCarrier).Find(&results).Error
	
	return
}

// GetBatchFromIDCarrier 批量唯一主键查找 
func (obj *_EgCartMgr) GetBatchFromIDCarrier(idCarriers []uint32) (results []*EgCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_carrier IN (?)", idCarriers).Find(&results).Error
	
	return
}
 
// GetFromDeliveryOption 通过delivery_option获取内容  
func (obj *_EgCartMgr) GetFromDeliveryOption(deliveryOption string) (results []*EgCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("delivery_option = ?", deliveryOption).Find(&results).Error
	
	return
}

// GetBatchFromDeliveryOption 批量唯一主键查找 
func (obj *_EgCartMgr) GetBatchFromDeliveryOption(deliveryOptions []string) (results []*EgCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("delivery_option IN (?)", deliveryOptions).Find(&results).Error
	
	return
}
 
// GetFromIDLang 通过id_lang获取内容  
func (obj *_EgCartMgr) GetFromIDLang(idLang uint32) (results []*EgCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error
	
	return
}

// GetBatchFromIDLang 批量唯一主键查找 
func (obj *_EgCartMgr) GetBatchFromIDLang(idLangs []uint32) (results []*EgCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error
	
	return
}
 
// GetFromIDAddressDelivery 通过id_address_delivery获取内容  
func (obj *_EgCartMgr) GetFromIDAddressDelivery(idAddressDelivery uint32) (results []*EgCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_address_delivery = ?", idAddressDelivery).Find(&results).Error
	
	return
}

// GetBatchFromIDAddressDelivery 批量唯一主键查找 
func (obj *_EgCartMgr) GetBatchFromIDAddressDelivery(idAddressDeliverys []uint32) (results []*EgCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_address_delivery IN (?)", idAddressDeliverys).Find(&results).Error
	
	return
}
 
// GetFromIDAddressInvoice 通过id_address_invoice获取内容  
func (obj *_EgCartMgr) GetFromIDAddressInvoice(idAddressInvoice uint32) (results []*EgCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_address_invoice = ?", idAddressInvoice).Find(&results).Error
	
	return
}

// GetBatchFromIDAddressInvoice 批量唯一主键查找 
func (obj *_EgCartMgr) GetBatchFromIDAddressInvoice(idAddressInvoices []uint32) (results []*EgCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_address_invoice IN (?)", idAddressInvoices).Find(&results).Error
	
	return
}
 
// GetFromIDCurrency 通过id_currency获取内容  
func (obj *_EgCartMgr) GetFromIDCurrency(idCurrency uint32) (results []*EgCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_currency = ?", idCurrency).Find(&results).Error
	
	return
}

// GetBatchFromIDCurrency 批量唯一主键查找 
func (obj *_EgCartMgr) GetBatchFromIDCurrency(idCurrencys []uint32) (results []*EgCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_currency IN (?)", idCurrencys).Find(&results).Error
	
	return
}
 
// GetFromIDCustomer 通过id_customer获取内容  
func (obj *_EgCartMgr) GetFromIDCustomer(idCustomer uint32) (results []*EgCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer = ?", idCustomer).Find(&results).Error
	
	return
}

// GetBatchFromIDCustomer 批量唯一主键查找 
func (obj *_EgCartMgr) GetBatchFromIDCustomer(idCustomers []uint32) (results []*EgCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer IN (?)", idCustomers).Find(&results).Error
	
	return
}
 
// GetFromIDGuest 通过id_guest获取内容  
func (obj *_EgCartMgr) GetFromIDGuest(idGuest uint32) (results []*EgCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_guest = ?", idGuest).Find(&results).Error
	
	return
}

// GetBatchFromIDGuest 批量唯一主键查找 
func (obj *_EgCartMgr) GetBatchFromIDGuest(idGuests []uint32) (results []*EgCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_guest IN (?)", idGuests).Find(&results).Error
	
	return
}
 
// GetFromSecureKey 通过secure_key获取内容  
func (obj *_EgCartMgr) GetFromSecureKey(secureKey string) (results []*EgCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("secure_key = ?", secureKey).Find(&results).Error
	
	return
}

// GetBatchFromSecureKey 批量唯一主键查找 
func (obj *_EgCartMgr) GetBatchFromSecureKey(secureKeys []string) (results []*EgCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("secure_key IN (?)", secureKeys).Find(&results).Error
	
	return
}
 
// GetFromRecyclable 通过recyclable获取内容  
func (obj *_EgCartMgr) GetFromRecyclable(recyclable bool) (results []*EgCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("recyclable = ?", recyclable).Find(&results).Error
	
	return
}

// GetBatchFromRecyclable 批量唯一主键查找 
func (obj *_EgCartMgr) GetBatchFromRecyclable(recyclables []bool) (results []*EgCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("recyclable IN (?)", recyclables).Find(&results).Error
	
	return
}
 
// GetFromGift 通过gift获取内容  
func (obj *_EgCartMgr) GetFromGift(gift bool) (results []*EgCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("gift = ?", gift).Find(&results).Error
	
	return
}

// GetBatchFromGift 批量唯一主键查找 
func (obj *_EgCartMgr) GetBatchFromGift(gifts []bool) (results []*EgCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("gift IN (?)", gifts).Find(&results).Error
	
	return
}
 
// GetFromGiftMessage 通过gift_message获取内容  
func (obj *_EgCartMgr) GetFromGiftMessage(giftMessage string) (results []*EgCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("gift_message = ?", giftMessage).Find(&results).Error
	
	return
}

// GetBatchFromGiftMessage 批量唯一主键查找 
func (obj *_EgCartMgr) GetBatchFromGiftMessage(giftMessages []string) (results []*EgCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("gift_message IN (?)", giftMessages).Find(&results).Error
	
	return
}
 
// GetFromMobileTheme 通过mobile_theme获取内容  
func (obj *_EgCartMgr) GetFromMobileTheme(mobileTheme bool) (results []*EgCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("mobile_theme = ?", mobileTheme).Find(&results).Error
	
	return
}

// GetBatchFromMobileTheme 批量唯一主键查找 
func (obj *_EgCartMgr) GetBatchFromMobileTheme(mobileThemes []bool) (results []*EgCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("mobile_theme IN (?)", mobileThemes).Find(&results).Error
	
	return
}
 
// GetFromAllowSeperatedPackage 通过allow_seperated_package获取内容  
func (obj *_EgCartMgr) GetFromAllowSeperatedPackage(allowSeperatedPackage bool) (results []*EgCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("allow_seperated_package = ?", allowSeperatedPackage).Find(&results).Error
	
	return
}

// GetBatchFromAllowSeperatedPackage 批量唯一主键查找 
func (obj *_EgCartMgr) GetBatchFromAllowSeperatedPackage(allowSeperatedPackages []bool) (results []*EgCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("allow_seperated_package IN (?)", allowSeperatedPackages).Find(&results).Error
	
	return
}
 
// GetFromDateAdd 通过date_add获取内容  
func (obj *_EgCartMgr) GetFromDateAdd(dateAdd time.Time) (results []*EgCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error
	
	return
}

// GetBatchFromDateAdd 批量唯一主键查找 
func (obj *_EgCartMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*EgCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error
	
	return
}
 
// GetFromDateUpd 通过date_upd获取内容  
func (obj *_EgCartMgr) GetFromDateUpd(dateUpd time.Time) (results []*EgCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd = ?", dateUpd).Find(&results).Error
	
	return
}

// GetBatchFromDateUpd 批量唯一主键查找 
func (obj *_EgCartMgr) GetBatchFromDateUpd(dateUpds []time.Time) (results []*EgCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd IN (?)", dateUpds).Find(&results).Error
	
	return
}
 
// GetFromCheckoutSessionData 通过checkout_session_data获取内容  
func (obj *_EgCartMgr) GetFromCheckoutSessionData(checkoutSessionData string) (results []*EgCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("checkout_session_data = ?", checkoutSessionData).Find(&results).Error
	
	return
}

// GetBatchFromCheckoutSessionData 批量唯一主键查找 
func (obj *_EgCartMgr) GetBatchFromCheckoutSessionData(checkoutSessionDatas []string) (results []*EgCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("checkout_session_data IN (?)", checkoutSessionDatas).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgCartMgr) FetchByPrimaryKey(idCart uint32 ) (result EgCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart = ?", idCart).Find(&result).Error
	
	return
}
 

 
 // FetchIndexByIDShopGroup  获取多个内容
 func (obj *_EgCartMgr) FetchIndexByIDShopGroup(idShopGroup uint32 ) (results []*EgCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_group = ?", idShopGroup).Find(&results).Error
	
	return
}
 
 // FetchIndexByIDShop2  获取多个内容
 func (obj *_EgCartMgr) FetchIndexByIDShop2(idShop uint32 ,dateUpd time.Time ) (results []*EgCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ? AND date_upd = ?", idShop , dateUpd).Find(&results).Error
	
	return
}
 
 // FetchIndexByIDShop  获取多个内容
 func (obj *_EgCartMgr) FetchIndexByIDShop(idShop uint32 ,dateAdd time.Time ) (results []*EgCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ? AND date_add = ?", idShop , dateAdd).Find(&results).Error
	
	return
}
 
 // FetchIndexByIDCarrier  获取多个内容
 func (obj *_EgCartMgr) FetchIndexByIDCarrier(idCarrier uint32 ) (results []*EgCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_carrier = ?", idCarrier).Find(&results).Error
	
	return
}
 
 // FetchIndexByIDLang  获取多个内容
 func (obj *_EgCartMgr) FetchIndexByIDLang(idLang uint32 ) (results []*EgCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error
	
	return
}
 
 // FetchIndexByIDAddressDelivery  获取多个内容
 func (obj *_EgCartMgr) FetchIndexByIDAddressDelivery(idAddressDelivery uint32 ) (results []*EgCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_address_delivery = ?", idAddressDelivery).Find(&results).Error
	
	return
}
 
 // FetchIndexByIDAddressInvoice  获取多个内容
 func (obj *_EgCartMgr) FetchIndexByIDAddressInvoice(idAddressInvoice uint32 ) (results []*EgCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_address_invoice = ?", idAddressInvoice).Find(&results).Error
	
	return
}
 
 // FetchIndexByIDCurrency  获取多个内容
 func (obj *_EgCartMgr) FetchIndexByIDCurrency(idCurrency uint32 ) (results []*EgCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_currency = ?", idCurrency).Find(&results).Error
	
	return
}
 
 // FetchIndexByCartCustomer  获取多个内容
 func (obj *_EgCartMgr) FetchIndexByCartCustomer(idCustomer uint32 ) (results []*EgCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer = ?", idCustomer).Find(&results).Error
	
	return
}
 
 // FetchIndexByIDGuest  获取多个内容
 func (obj *_EgCartMgr) FetchIndexByIDGuest(idGuest uint32 ) (results []*EgCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_guest = ?", idGuest).Find(&results).Error
	
	return
}
 

	

