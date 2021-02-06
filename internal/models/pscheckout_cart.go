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

func PscheckoutCartMgr(db *gorm.DB) *_PscheckoutCartMgr {
	if db == nil {
		panic(fmt.Errorf("PscheckoutCartMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_PscheckoutCartMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_pscheckout_cart"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_PscheckoutCartMgr) GetTableName() string {
	return "ps_pscheckout_cart"
}

func (obj *_PscheckoutCartMgr) Get() (result PscheckoutCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_PscheckoutCartMgr) Gets() (results []*PscheckoutCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_PscheckoutCartMgr) WithIDPscheckoutCart(idPscheckoutCart uint32) Option {
	return optionFunc(func(o *options) { o.query["id_pscheckout_cart"] = idPscheckoutCart })
}

func (obj *_PscheckoutCartMgr) WithIDCart(idCart uint32) Option {
	return optionFunc(func(o *options) { o.query["id_cart"] = idCart })
}

func (obj *_PscheckoutCartMgr) WithPaypalIntent(paypalIntent string) Option {
	return optionFunc(func(o *options) { o.query["paypal_intent"] = paypalIntent })
}

func (obj *_PscheckoutCartMgr) WithPaypalOrder(paypalOrder string) Option {
	return optionFunc(func(o *options) { o.query["paypal_order"] = paypalOrder })
}

func (obj *_PscheckoutCartMgr) WithPaypalStatus(paypalStatus string) Option {
	return optionFunc(func(o *options) { o.query["paypal_status"] = paypalStatus })
}

func (obj *_PscheckoutCartMgr) WithPaypalFunding(paypalFunding string) Option {
	return optionFunc(func(o *options) { o.query["paypal_funding"] = paypalFunding })
}

func (obj *_PscheckoutCartMgr) WithPaypalToken(paypalToken string) Option {
	return optionFunc(func(o *options) { o.query["paypal_token"] = paypalToken })
}

func (obj *_PscheckoutCartMgr) WithPaypalTokenExpire(paypalTokenExpire time.Time) Option {
	return optionFunc(func(o *options) { o.query["paypal_token_expire"] = paypalTokenExpire })
}

func (obj *_PscheckoutCartMgr) WithPaypalAuthorizationExpire(paypalAuthorizationExpire time.Time) Option {
	return optionFunc(func(o *options) { o.query["paypal_authorization_expire"] = paypalAuthorizationExpire })
}

func (obj *_PscheckoutCartMgr) WithIsExpressCheckout(isExpressCheckout bool) Option {
	return optionFunc(func(o *options) { o.query["isExpressCheckout"] = isExpressCheckout })
}

func (obj *_PscheckoutCartMgr) WithIsHostedFields(isHostedFields bool) Option {
	return optionFunc(func(o *options) { o.query["isHostedFields"] = isHostedFields })
}

func (obj *_PscheckoutCartMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

func (obj *_PscheckoutCartMgr) WithDateUpd(dateUpd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_upd"] = dateUpd })
}

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


func (obj *_PscheckoutCartMgr) GetFromIDPscheckoutCart(idPscheckoutCart uint32) (result PscheckoutCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_pscheckout_cart = ?", idPscheckoutCart).Find(&result).Error

	return
}

func (obj *_PscheckoutCartMgr) GetBatchFromIDPscheckoutCart(idPscheckoutCarts []uint32) (results []*PscheckoutCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_pscheckout_cart IN (?)", idPscheckoutCarts).Find(&results).Error

	return
}

func (obj *_PscheckoutCartMgr) GetFromIDCart(idCart uint32) (results []*PscheckoutCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart = ?", idCart).Find(&results).Error

	return
}

func (obj *_PscheckoutCartMgr) GetBatchFromIDCart(idCarts []uint32) (results []*PscheckoutCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart IN (?)", idCarts).Find(&results).Error

	return
}

func (obj *_PscheckoutCartMgr) GetFromPaypalIntent(paypalIntent string) (results []*PscheckoutCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("paypal_intent = ?", paypalIntent).Find(&results).Error

	return
}

func (obj *_PscheckoutCartMgr) GetBatchFromPaypalIntent(paypalIntents []string) (results []*PscheckoutCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("paypal_intent IN (?)", paypalIntents).Find(&results).Error

	return
}

func (obj *_PscheckoutCartMgr) GetFromPaypalOrder(paypalOrder string) (results []*PscheckoutCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("paypal_order = ?", paypalOrder).Find(&results).Error

	return
}

func (obj *_PscheckoutCartMgr) GetBatchFromPaypalOrder(paypalOrders []string) (results []*PscheckoutCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("paypal_order IN (?)", paypalOrders).Find(&results).Error

	return
}

func (obj *_PscheckoutCartMgr) GetFromPaypalStatus(paypalStatus string) (results []*PscheckoutCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("paypal_status = ?", paypalStatus).Find(&results).Error

	return
}

func (obj *_PscheckoutCartMgr) GetBatchFromPaypalStatus(paypalStatuss []string) (results []*PscheckoutCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("paypal_status IN (?)", paypalStatuss).Find(&results).Error

	return
}

func (obj *_PscheckoutCartMgr) GetFromPaypalFunding(paypalFunding string) (results []*PscheckoutCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("paypal_funding = ?", paypalFunding).Find(&results).Error

	return
}

func (obj *_PscheckoutCartMgr) GetBatchFromPaypalFunding(paypalFundings []string) (results []*PscheckoutCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("paypal_funding IN (?)", paypalFundings).Find(&results).Error

	return
}

func (obj *_PscheckoutCartMgr) GetFromPaypalToken(paypalToken string) (results []*PscheckoutCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("paypal_token = ?", paypalToken).Find(&results).Error

	return
}

func (obj *_PscheckoutCartMgr) GetBatchFromPaypalToken(paypalTokens []string) (results []*PscheckoutCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("paypal_token IN (?)", paypalTokens).Find(&results).Error

	return
}

func (obj *_PscheckoutCartMgr) GetFromPaypalTokenExpire(paypalTokenExpire time.Time) (results []*PscheckoutCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("paypal_token_expire = ?", paypalTokenExpire).Find(&results).Error

	return
}

func (obj *_PscheckoutCartMgr) GetBatchFromPaypalTokenExpire(paypalTokenExpires []time.Time) (results []*PscheckoutCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("paypal_token_expire IN (?)", paypalTokenExpires).Find(&results).Error

	return
}

func (obj *_PscheckoutCartMgr) GetFromPaypalAuthorizationExpire(paypalAuthorizationExpire time.Time) (results []*PscheckoutCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("paypal_authorization_expire = ?", paypalAuthorizationExpire).Find(&results).Error

	return
}

func (obj *_PscheckoutCartMgr) GetBatchFromPaypalAuthorizationExpire(paypalAuthorizationExpires []time.Time) (results []*PscheckoutCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("paypal_authorization_expire IN (?)", paypalAuthorizationExpires).Find(&results).Error

	return
}

func (obj *_PscheckoutCartMgr) GetFromIsExpressCheckout(isExpressCheckout bool) (results []*PscheckoutCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("isExpressCheckout = ?", isExpressCheckout).Find(&results).Error

	return
}

func (obj *_PscheckoutCartMgr) GetBatchFromIsExpressCheckout(isExpressCheckouts []bool) (results []*PscheckoutCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("isExpressCheckout IN (?)", isExpressCheckouts).Find(&results).Error

	return
}

func (obj *_PscheckoutCartMgr) GetFromIsHostedFields(isHostedFields bool) (results []*PscheckoutCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("isHostedFields = ?", isHostedFields).Find(&results).Error

	return
}

func (obj *_PscheckoutCartMgr) GetBatchFromIsHostedFields(isHostedFieldss []bool) (results []*PscheckoutCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("isHostedFields IN (?)", isHostedFieldss).Find(&results).Error

	return
}

func (obj *_PscheckoutCartMgr) GetFromDateAdd(dateAdd time.Time) (results []*PscheckoutCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error

	return
}

func (obj *_PscheckoutCartMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*PscheckoutCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error

	return
}

func (obj *_PscheckoutCartMgr) GetFromDateUpd(dateUpd time.Time) (results []*PscheckoutCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd = ?", dateUpd).Find(&results).Error

	return
}

func (obj *_PscheckoutCartMgr) GetBatchFromDateUpd(dateUpds []time.Time) (results []*PscheckoutCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd IN (?)", dateUpds).Find(&results).Error

	return
}


func (obj *_PscheckoutCartMgr) FetchByPrimaryKey(idPscheckoutCart uint32) (result PscheckoutCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_pscheckout_cart = ?", idPscheckoutCart).Find(&result).Error

	return
}
