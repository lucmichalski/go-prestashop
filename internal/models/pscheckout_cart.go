package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
"time"	
)	

type _EgPscheckoutCartMgr struct {
	*_BaseMgr
}

// EgPscheckoutCartMgr open func
func EgPscheckoutCartMgr(db *gorm.DB) *_EgPscheckoutCartMgr {
	if db == nil {
		panic(fmt.Errorf("EgPscheckoutCartMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgPscheckoutCartMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_pscheckout_cart"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgPscheckoutCartMgr) GetTableName() string {
	return "eg_pscheckout_cart"
}

// Get 获取
func (obj *_EgPscheckoutCartMgr) Get() (result EgPscheckoutCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgPscheckoutCartMgr) Gets() (results []*EgPscheckoutCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDPscheckoutCart id_pscheckout_cart获取 
func (obj *_EgPscheckoutCartMgr) WithIDPscheckoutCart(idPscheckoutCart uint32) Option {
	return optionFunc(func(o *options) { o.query["id_pscheckout_cart"] = idPscheckoutCart })
}

// WithIDCart id_cart获取 
func (obj *_EgPscheckoutCartMgr) WithIDCart(idCart uint32) Option {
	return optionFunc(func(o *options) { o.query["id_cart"] = idCart })
}

// WithPaypalIntent paypal_intent获取 
func (obj *_EgPscheckoutCartMgr) WithPaypalIntent(paypalIntent string) Option {
	return optionFunc(func(o *options) { o.query["paypal_intent"] = paypalIntent })
}

// WithPaypalOrder paypal_order获取 
func (obj *_EgPscheckoutCartMgr) WithPaypalOrder(paypalOrder string) Option {
	return optionFunc(func(o *options) { o.query["paypal_order"] = paypalOrder })
}

// WithPaypalStatus paypal_status获取 
func (obj *_EgPscheckoutCartMgr) WithPaypalStatus(paypalStatus string) Option {
	return optionFunc(func(o *options) { o.query["paypal_status"] = paypalStatus })
}

// WithPaypalFunding paypal_funding获取 
func (obj *_EgPscheckoutCartMgr) WithPaypalFunding(paypalFunding string) Option {
	return optionFunc(func(o *options) { o.query["paypal_funding"] = paypalFunding })
}

// WithPaypalToken paypal_token获取 
func (obj *_EgPscheckoutCartMgr) WithPaypalToken(paypalToken string) Option {
	return optionFunc(func(o *options) { o.query["paypal_token"] = paypalToken })
}

// WithPaypalTokenExpire paypal_token_expire获取 
func (obj *_EgPscheckoutCartMgr) WithPaypalTokenExpire(paypalTokenExpire time.Time) Option {
	return optionFunc(func(o *options) { o.query["paypal_token_expire"] = paypalTokenExpire })
}

// WithPaypalAuthorizationExpire paypal_authorization_expire获取 
func (obj *_EgPscheckoutCartMgr) WithPaypalAuthorizationExpire(paypalAuthorizationExpire time.Time) Option {
	return optionFunc(func(o *options) { o.query["paypal_authorization_expire"] = paypalAuthorizationExpire })
}

// WithIsExpressCheckout isExpressCheckout获取 
func (obj *_EgPscheckoutCartMgr) WithIsExpressCheckout(isExpressCheckout bool) Option {
	return optionFunc(func(o *options) { o.query["isExpressCheckout"] = isExpressCheckout })
}

// WithIsHostedFields isHostedFields获取 
func (obj *_EgPscheckoutCartMgr) WithIsHostedFields(isHostedFields bool) Option {
	return optionFunc(func(o *options) { o.query["isHostedFields"] = isHostedFields })
}

// WithDateAdd date_add获取 
func (obj *_EgPscheckoutCartMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

// WithDateUpd date_upd获取 
func (obj *_EgPscheckoutCartMgr) WithDateUpd(dateUpd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_upd"] = dateUpd })
}


// GetByOption 功能选项模式获取
func (obj *_EgPscheckoutCartMgr) GetByOption(opts ...Option) (result EgPscheckoutCart, err error) {
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
func (obj *_EgPscheckoutCartMgr) GetByOptions(opts ...Option) (results []*EgPscheckoutCart, err error) {
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
func (obj *_EgPscheckoutCartMgr)  GetFromIDPscheckoutCart(idPscheckoutCart uint32) (result EgPscheckoutCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_pscheckout_cart = ?", idPscheckoutCart).Find(&result).Error
	
	return
}

// GetBatchFromIDPscheckoutCart 批量唯一主键查找 
func (obj *_EgPscheckoutCartMgr) GetBatchFromIDPscheckoutCart(idPscheckoutCarts []uint32) (results []*EgPscheckoutCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_pscheckout_cart IN (?)", idPscheckoutCarts).Find(&results).Error
	
	return
}
 
// GetFromIDCart 通过id_cart获取内容  
func (obj *_EgPscheckoutCartMgr) GetFromIDCart(idCart uint32) (results []*EgPscheckoutCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart = ?", idCart).Find(&results).Error
	
	return
}

// GetBatchFromIDCart 批量唯一主键查找 
func (obj *_EgPscheckoutCartMgr) GetBatchFromIDCart(idCarts []uint32) (results []*EgPscheckoutCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart IN (?)", idCarts).Find(&results).Error
	
	return
}
 
// GetFromPaypalIntent 通过paypal_intent获取内容  
func (obj *_EgPscheckoutCartMgr) GetFromPaypalIntent(paypalIntent string) (results []*EgPscheckoutCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("paypal_intent = ?", paypalIntent).Find(&results).Error
	
	return
}

// GetBatchFromPaypalIntent 批量唯一主键查找 
func (obj *_EgPscheckoutCartMgr) GetBatchFromPaypalIntent(paypalIntents []string) (results []*EgPscheckoutCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("paypal_intent IN (?)", paypalIntents).Find(&results).Error
	
	return
}
 
// GetFromPaypalOrder 通过paypal_order获取内容  
func (obj *_EgPscheckoutCartMgr) GetFromPaypalOrder(paypalOrder string) (results []*EgPscheckoutCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("paypal_order = ?", paypalOrder).Find(&results).Error
	
	return
}

// GetBatchFromPaypalOrder 批量唯一主键查找 
func (obj *_EgPscheckoutCartMgr) GetBatchFromPaypalOrder(paypalOrders []string) (results []*EgPscheckoutCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("paypal_order IN (?)", paypalOrders).Find(&results).Error
	
	return
}
 
// GetFromPaypalStatus 通过paypal_status获取内容  
func (obj *_EgPscheckoutCartMgr) GetFromPaypalStatus(paypalStatus string) (results []*EgPscheckoutCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("paypal_status = ?", paypalStatus).Find(&results).Error
	
	return
}

// GetBatchFromPaypalStatus 批量唯一主键查找 
func (obj *_EgPscheckoutCartMgr) GetBatchFromPaypalStatus(paypalStatuss []string) (results []*EgPscheckoutCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("paypal_status IN (?)", paypalStatuss).Find(&results).Error
	
	return
}
 
// GetFromPaypalFunding 通过paypal_funding获取内容  
func (obj *_EgPscheckoutCartMgr) GetFromPaypalFunding(paypalFunding string) (results []*EgPscheckoutCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("paypal_funding = ?", paypalFunding).Find(&results).Error
	
	return
}

// GetBatchFromPaypalFunding 批量唯一主键查找 
func (obj *_EgPscheckoutCartMgr) GetBatchFromPaypalFunding(paypalFundings []string) (results []*EgPscheckoutCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("paypal_funding IN (?)", paypalFundings).Find(&results).Error
	
	return
}
 
// GetFromPaypalToken 通过paypal_token获取内容  
func (obj *_EgPscheckoutCartMgr) GetFromPaypalToken(paypalToken string) (results []*EgPscheckoutCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("paypal_token = ?", paypalToken).Find(&results).Error
	
	return
}

// GetBatchFromPaypalToken 批量唯一主键查找 
func (obj *_EgPscheckoutCartMgr) GetBatchFromPaypalToken(paypalTokens []string) (results []*EgPscheckoutCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("paypal_token IN (?)", paypalTokens).Find(&results).Error
	
	return
}
 
// GetFromPaypalTokenExpire 通过paypal_token_expire获取内容  
func (obj *_EgPscheckoutCartMgr) GetFromPaypalTokenExpire(paypalTokenExpire time.Time) (results []*EgPscheckoutCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("paypal_token_expire = ?", paypalTokenExpire).Find(&results).Error
	
	return
}

// GetBatchFromPaypalTokenExpire 批量唯一主键查找 
func (obj *_EgPscheckoutCartMgr) GetBatchFromPaypalTokenExpire(paypalTokenExpires []time.Time) (results []*EgPscheckoutCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("paypal_token_expire IN (?)", paypalTokenExpires).Find(&results).Error
	
	return
}
 
// GetFromPaypalAuthorizationExpire 通过paypal_authorization_expire获取内容  
func (obj *_EgPscheckoutCartMgr) GetFromPaypalAuthorizationExpire(paypalAuthorizationExpire time.Time) (results []*EgPscheckoutCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("paypal_authorization_expire = ?", paypalAuthorizationExpire).Find(&results).Error
	
	return
}

// GetBatchFromPaypalAuthorizationExpire 批量唯一主键查找 
func (obj *_EgPscheckoutCartMgr) GetBatchFromPaypalAuthorizationExpire(paypalAuthorizationExpires []time.Time) (results []*EgPscheckoutCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("paypal_authorization_expire IN (?)", paypalAuthorizationExpires).Find(&results).Error
	
	return
}
 
// GetFromIsExpressCheckout 通过isExpressCheckout获取内容  
func (obj *_EgPscheckoutCartMgr) GetFromIsExpressCheckout(isExpressCheckout bool) (results []*EgPscheckoutCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("isExpressCheckout = ?", isExpressCheckout).Find(&results).Error
	
	return
}

// GetBatchFromIsExpressCheckout 批量唯一主键查找 
func (obj *_EgPscheckoutCartMgr) GetBatchFromIsExpressCheckout(isExpressCheckouts []bool) (results []*EgPscheckoutCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("isExpressCheckout IN (?)", isExpressCheckouts).Find(&results).Error
	
	return
}
 
// GetFromIsHostedFields 通过isHostedFields获取内容  
func (obj *_EgPscheckoutCartMgr) GetFromIsHostedFields(isHostedFields bool) (results []*EgPscheckoutCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("isHostedFields = ?", isHostedFields).Find(&results).Error
	
	return
}

// GetBatchFromIsHostedFields 批量唯一主键查找 
func (obj *_EgPscheckoutCartMgr) GetBatchFromIsHostedFields(isHostedFieldss []bool) (results []*EgPscheckoutCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("isHostedFields IN (?)", isHostedFieldss).Find(&results).Error
	
	return
}
 
// GetFromDateAdd 通过date_add获取内容  
func (obj *_EgPscheckoutCartMgr) GetFromDateAdd(dateAdd time.Time) (results []*EgPscheckoutCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error
	
	return
}

// GetBatchFromDateAdd 批量唯一主键查找 
func (obj *_EgPscheckoutCartMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*EgPscheckoutCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error
	
	return
}
 
// GetFromDateUpd 通过date_upd获取内容  
func (obj *_EgPscheckoutCartMgr) GetFromDateUpd(dateUpd time.Time) (results []*EgPscheckoutCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd = ?", dateUpd).Find(&results).Error
	
	return
}

// GetBatchFromDateUpd 批量唯一主键查找 
func (obj *_EgPscheckoutCartMgr) GetBatchFromDateUpd(dateUpds []time.Time) (results []*EgPscheckoutCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd IN (?)", dateUpds).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgPscheckoutCartMgr) FetchByPrimaryKey(idPscheckoutCart uint32 ) (result EgPscheckoutCart, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_pscheckout_cart = ?", idPscheckoutCart).Find(&result).Error
	
	return
}
 

 

	

