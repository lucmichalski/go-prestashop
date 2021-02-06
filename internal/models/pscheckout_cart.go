package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _PscheckoutCartMgr struct {
	*_BaseMgr
}

// PscheckoutCartMgr open func
func PscheckoutCartMgr(db *gorm.DB) *_PscheckoutCartMgr {
	if db == nil {
		panic(fmt.Errorf("PscheckoutCartMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_PscheckoutCartMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_pscheckout_cart"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_PscheckoutCartMgr) GetTableName() string {
	return "ps_pscheckout_cart"
}

// Get 获取
func (obj *_PscheckoutCartMgr) Get() (result PscheckoutCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_PscheckoutCartMgr) Gets() (results []*PscheckoutCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDPscheckoutCart id_pscheckout_cart获取
func (obj *_PscheckoutCartMgr) WithIDPscheckoutCart(idPscheckoutCart uint32) Option {
	return optionFunc(func(o *options) { o.query["id_pscheckout_cart"] = idPscheckoutCart })
}

// WithIDCart id_cart获取
func (obj *_PscheckoutCartMgr) WithIDCart(idCart uint32) Option {
	return optionFunc(func(o *options) { o.query["id_cart"] = idCart })
}

// WithPaypalIntent paypal_intent获取
func (obj *_PscheckoutCartMgr) WithPaypalIntent(paypalIntent string) Option {
	return optionFunc(func(o *options) { o.query["paypal_intent"] = paypalIntent })
}

// WithPaypalOrder paypal_order获取
func (obj *_PscheckoutCartMgr) WithPaypalOrder(paypalOrder string) Option {
	return optionFunc(func(o *options) { o.query["paypal_order"] = paypalOrder })
}

// WithPaypalStatus paypal_status获取
func (obj *_PscheckoutCartMgr) WithPaypalStatus(paypalStatus string) Option {
	return optionFunc(func(o *options) { o.query["paypal_status"] = paypalStatus })
}

// WithPaypalFunding paypal_funding获取
func (obj *_PscheckoutCartMgr) WithPaypalFunding(paypalFunding string) Option {
	return optionFunc(func(o *options) { o.query["paypal_funding"] = paypalFunding })
}

// WithPaypalToken paypal_token获取
func (obj *_PscheckoutCartMgr) WithPaypalToken(paypalToken string) Option {
	return optionFunc(func(o *options) { o.query["paypal_token"] = paypalToken })
}

// WithPaypalTokenExpire paypal_token_expire获取
func (obj *_PscheckoutCartMgr) WithPaypalTokenExpire(paypalTokenExpire time.Time) Option {
	return optionFunc(func(o *options) { o.query["paypal_token_expire"] = paypalTokenExpire })
}

// WithPaypalAuthorizationExpire paypal_authorization_expire获取
func (obj *_PscheckoutCartMgr) WithPaypalAuthorizationExpire(paypalAuthorizationExpire time.Time) Option {
	return optionFunc(func(o *options) { o.query["paypal_authorization_expire"] = paypalAuthorizationExpire })
}

// WithIsExpressCheckout isExpressCheckout获取
func (obj *_PscheckoutCartMgr) WithIsExpressCheckout(isExpressCheckout bool) Option {
	return optionFunc(func(o *options) { o.query["isExpressCheckout"] = isExpressCheckout })
}

// WithIsHostedFields isHostedFields获取
func (obj *_PscheckoutCartMgr) WithIsHostedFields(isHostedFields bool) Option {
	return optionFunc(func(o *options) { o.query["isHostedFields"] = isHostedFields })
}

// WithDateAdd date_add获取
func (obj *_PscheckoutCartMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

// WithDateUpd date_upd获取
func (obj *_PscheckoutCartMgr) WithDateUpd(dateUpd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_upd"] = dateUpd })
}

// GetByOption 功能选项模式获取
func (obj *_PscheckoutCartMgr) GetByOption(opts ...Option) (result PscheckoutCart, err error) {
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
func (obj *_PscheckoutCartMgr) GetByOptions(opts ...Option) (results []*PscheckoutCart, err error) {
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

// GetFromIDPscheckoutCart 通过id_pscheckout_cart获取内容
func (obj *_PscheckoutCartMgr) GetFromIDPscheckoutCart(idPscheckoutCart uint32) (result PscheckoutCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_pscheckout_cart = ?", idPscheckoutCart).Find(&result).Error

	return
}

// GetBatchFromIDPscheckoutCart 批量唯一主键查找
func (obj *_PscheckoutCartMgr) GetBatchFromIDPscheckoutCart(idPscheckoutCarts []uint32) (results []*PscheckoutCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_pscheckout_cart IN (?)", idPscheckoutCarts).Find(&results).Error

	return
}

// GetFromIDCart 通过id_cart获取内容
func (obj *_PscheckoutCartMgr) GetFromIDCart(idCart uint32) (results []*PscheckoutCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart = ?", idCart).Find(&results).Error

	return
}

// GetBatchFromIDCart 批量唯一主键查找
func (obj *_PscheckoutCartMgr) GetBatchFromIDCart(idCarts []uint32) (results []*PscheckoutCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart IN (?)", idCarts).Find(&results).Error

	return
}

// GetFromPaypalIntent 通过paypal_intent获取内容
func (obj *_PscheckoutCartMgr) GetFromPaypalIntent(paypalIntent string) (results []*PscheckoutCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("paypal_intent = ?", paypalIntent).Find(&results).Error

	return
}

// GetBatchFromPaypalIntent 批量唯一主键查找
func (obj *_PscheckoutCartMgr) GetBatchFromPaypalIntent(paypalIntents []string) (results []*PscheckoutCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("paypal_intent IN (?)", paypalIntents).Find(&results).Error

	return
}

// GetFromPaypalOrder 通过paypal_order获取内容
func (obj *_PscheckoutCartMgr) GetFromPaypalOrder(paypalOrder string) (results []*PscheckoutCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("paypal_order = ?", paypalOrder).Find(&results).Error

	return
}

// GetBatchFromPaypalOrder 批量唯一主键查找
func (obj *_PscheckoutCartMgr) GetBatchFromPaypalOrder(paypalOrders []string) (results []*PscheckoutCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("paypal_order IN (?)", paypalOrders).Find(&results).Error

	return
}

// GetFromPaypalStatus 通过paypal_status获取内容
func (obj *_PscheckoutCartMgr) GetFromPaypalStatus(paypalStatus string) (results []*PscheckoutCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("paypal_status = ?", paypalStatus).Find(&results).Error

	return
}

// GetBatchFromPaypalStatus 批量唯一主键查找
func (obj *_PscheckoutCartMgr) GetBatchFromPaypalStatus(paypalStatuss []string) (results []*PscheckoutCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("paypal_status IN (?)", paypalStatuss).Find(&results).Error

	return
}

// GetFromPaypalFunding 通过paypal_funding获取内容
func (obj *_PscheckoutCartMgr) GetFromPaypalFunding(paypalFunding string) (results []*PscheckoutCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("paypal_funding = ?", paypalFunding).Find(&results).Error

	return
}

// GetBatchFromPaypalFunding 批量唯一主键查找
func (obj *_PscheckoutCartMgr) GetBatchFromPaypalFunding(paypalFundings []string) (results []*PscheckoutCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("paypal_funding IN (?)", paypalFundings).Find(&results).Error

	return
}

// GetFromPaypalToken 通过paypal_token获取内容
func (obj *_PscheckoutCartMgr) GetFromPaypalToken(paypalToken string) (results []*PscheckoutCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("paypal_token = ?", paypalToken).Find(&results).Error

	return
}

// GetBatchFromPaypalToken 批量唯一主键查找
func (obj *_PscheckoutCartMgr) GetBatchFromPaypalToken(paypalTokens []string) (results []*PscheckoutCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("paypal_token IN (?)", paypalTokens).Find(&results).Error

	return
}

// GetFromPaypalTokenExpire 通过paypal_token_expire获取内容
func (obj *_PscheckoutCartMgr) GetFromPaypalTokenExpire(paypalTokenExpire time.Time) (results []*PscheckoutCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("paypal_token_expire = ?", paypalTokenExpire).Find(&results).Error

	return
}

// GetBatchFromPaypalTokenExpire 批量唯一主键查找
func (obj *_PscheckoutCartMgr) GetBatchFromPaypalTokenExpire(paypalTokenExpires []time.Time) (results []*PscheckoutCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("paypal_token_expire IN (?)", paypalTokenExpires).Find(&results).Error

	return
}

// GetFromPaypalAuthorizationExpire 通过paypal_authorization_expire获取内容
func (obj *_PscheckoutCartMgr) GetFromPaypalAuthorizationExpire(paypalAuthorizationExpire time.Time) (results []*PscheckoutCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("paypal_authorization_expire = ?", paypalAuthorizationExpire).Find(&results).Error

	return
}

// GetBatchFromPaypalAuthorizationExpire 批量唯一主键查找
func (obj *_PscheckoutCartMgr) GetBatchFromPaypalAuthorizationExpire(paypalAuthorizationExpires []time.Time) (results []*PscheckoutCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("paypal_authorization_expire IN (?)", paypalAuthorizationExpires).Find(&results).Error

	return
}

// GetFromIsExpressCheckout 通过isExpressCheckout获取内容
func (obj *_PscheckoutCartMgr) GetFromIsExpressCheckout(isExpressCheckout bool) (results []*PscheckoutCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("isExpressCheckout = ?", isExpressCheckout).Find(&results).Error

	return
}

// GetBatchFromIsExpressCheckout 批量唯一主键查找
func (obj *_PscheckoutCartMgr) GetBatchFromIsExpressCheckout(isExpressCheckouts []bool) (results []*PscheckoutCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("isExpressCheckout IN (?)", isExpressCheckouts).Find(&results).Error

	return
}

// GetFromIsHostedFields 通过isHostedFields获取内容
func (obj *_PscheckoutCartMgr) GetFromIsHostedFields(isHostedFields bool) (results []*PscheckoutCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("isHostedFields = ?", isHostedFields).Find(&results).Error

	return
}

// GetBatchFromIsHostedFields 批量唯一主键查找
func (obj *_PscheckoutCartMgr) GetBatchFromIsHostedFields(isHostedFieldss []bool) (results []*PscheckoutCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("isHostedFields IN (?)", isHostedFieldss).Find(&results).Error

	return
}

// GetFromDateAdd 通过date_add获取内容
func (obj *_PscheckoutCartMgr) GetFromDateAdd(dateAdd time.Time) (results []*PscheckoutCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error

	return
}

// GetBatchFromDateAdd 批量唯一主键查找
func (obj *_PscheckoutCartMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*PscheckoutCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error

	return
}

// GetFromDateUpd 通过date_upd获取内容
func (obj *_PscheckoutCartMgr) GetFromDateUpd(dateUpd time.Time) (results []*PscheckoutCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd = ?", dateUpd).Find(&results).Error

	return
}

// GetBatchFromDateUpd 批量唯一主键查找
func (obj *_PscheckoutCartMgr) GetBatchFromDateUpd(dateUpds []time.Time) (results []*PscheckoutCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd IN (?)", dateUpds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_PscheckoutCartMgr) FetchByPrimaryKey(idPscheckoutCart uint32) (result PscheckoutCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_pscheckout_cart = ?", idPscheckoutCart).Find(&result).Error

	return
}
