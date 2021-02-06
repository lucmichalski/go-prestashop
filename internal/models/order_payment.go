package	model	
import (	
"context"	
"gorm.io/gorm"	
"time"	
"fmt"	
)	

type _EgOrderPaymentMgr struct {
	*_BaseMgr
}

// EgOrderPaymentMgr open func
func EgOrderPaymentMgr(db *gorm.DB) *_EgOrderPaymentMgr {
	if db == nil {
		panic(fmt.Errorf("EgOrderPaymentMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgOrderPaymentMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_order_payment"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgOrderPaymentMgr) GetTableName() string {
	return "eg_order_payment"
}

// Get 获取
func (obj *_EgOrderPaymentMgr) Get() (result EgOrderPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgOrderPaymentMgr) Gets() (results []*EgOrderPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDOrderPayment id_order_payment获取 
func (obj *_EgOrderPaymentMgr) WithIDOrderPayment(idOrderPayment int) Option {
	return optionFunc(func(o *options) { o.query["id_order_payment"] = idOrderPayment })
}

// WithOrderReference order_reference获取 
func (obj *_EgOrderPaymentMgr) WithOrderReference(orderReference string) Option {
	return optionFunc(func(o *options) { o.query["order_reference"] = orderReference })
}

// WithIDCurrency id_currency获取 
func (obj *_EgOrderPaymentMgr) WithIDCurrency(idCurrency uint32) Option {
	return optionFunc(func(o *options) { o.query["id_currency"] = idCurrency })
}

// WithAmount amount获取 
func (obj *_EgOrderPaymentMgr) WithAmount(amount float64) Option {
	return optionFunc(func(o *options) { o.query["amount"] = amount })
}

// WithPaymentMethod payment_method获取 
func (obj *_EgOrderPaymentMgr) WithPaymentMethod(paymentMethod string) Option {
	return optionFunc(func(o *options) { o.query["payment_method"] = paymentMethod })
}

// WithConversionRate conversion_rate获取 
func (obj *_EgOrderPaymentMgr) WithConversionRate(conversionRate float64) Option {
	return optionFunc(func(o *options) { o.query["conversion_rate"] = conversionRate })
}

// WithTransactionID transaction_id获取 
func (obj *_EgOrderPaymentMgr) WithTransactionID(transactionID string) Option {
	return optionFunc(func(o *options) { o.query["transaction_id"] = transactionID })
}

// WithCardNumber card_number获取 
func (obj *_EgOrderPaymentMgr) WithCardNumber(cardNumber string) Option {
	return optionFunc(func(o *options) { o.query["card_number"] = cardNumber })
}

// WithCardBrand card_brand获取 
func (obj *_EgOrderPaymentMgr) WithCardBrand(cardBrand string) Option {
	return optionFunc(func(o *options) { o.query["card_brand"] = cardBrand })
}

// WithCardExpiration card_expiration获取 
func (obj *_EgOrderPaymentMgr) WithCardExpiration(cardExpiration string) Option {
	return optionFunc(func(o *options) { o.query["card_expiration"] = cardExpiration })
}

// WithCardHolder card_holder获取 
func (obj *_EgOrderPaymentMgr) WithCardHolder(cardHolder string) Option {
	return optionFunc(func(o *options) { o.query["card_holder"] = cardHolder })
}

// WithDateAdd date_add获取 
func (obj *_EgOrderPaymentMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}


// GetByOption 功能选项模式获取
func (obj *_EgOrderPaymentMgr) GetByOption(opts ...Option) (result EgOrderPayment, err error) {
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
func (obj *_EgOrderPaymentMgr) GetByOptions(opts ...Option) (results []*EgOrderPayment, err error) {
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


// GetFromIDOrderPayment 通过id_order_payment获取内容  
func (obj *_EgOrderPaymentMgr)  GetFromIDOrderPayment(idOrderPayment int) (result EgOrderPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_payment = ?", idOrderPayment).Find(&result).Error
	
	return
}

// GetBatchFromIDOrderPayment 批量唯一主键查找 
func (obj *_EgOrderPaymentMgr) GetBatchFromIDOrderPayment(idOrderPayments []int) (results []*EgOrderPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_payment IN (?)", idOrderPayments).Find(&results).Error
	
	return
}
 
// GetFromOrderReference 通过order_reference获取内容  
func (obj *_EgOrderPaymentMgr) GetFromOrderReference(orderReference string) (results []*EgOrderPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("order_reference = ?", orderReference).Find(&results).Error
	
	return
}

// GetBatchFromOrderReference 批量唯一主键查找 
func (obj *_EgOrderPaymentMgr) GetBatchFromOrderReference(orderReferences []string) (results []*EgOrderPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("order_reference IN (?)", orderReferences).Find(&results).Error
	
	return
}
 
// GetFromIDCurrency 通过id_currency获取内容  
func (obj *_EgOrderPaymentMgr) GetFromIDCurrency(idCurrency uint32) (results []*EgOrderPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_currency = ?", idCurrency).Find(&results).Error
	
	return
}

// GetBatchFromIDCurrency 批量唯一主键查找 
func (obj *_EgOrderPaymentMgr) GetBatchFromIDCurrency(idCurrencys []uint32) (results []*EgOrderPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_currency IN (?)", idCurrencys).Find(&results).Error
	
	return
}
 
// GetFromAmount 通过amount获取内容  
func (obj *_EgOrderPaymentMgr) GetFromAmount(amount float64) (results []*EgOrderPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("amount = ?", amount).Find(&results).Error
	
	return
}

// GetBatchFromAmount 批量唯一主键查找 
func (obj *_EgOrderPaymentMgr) GetBatchFromAmount(amounts []float64) (results []*EgOrderPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("amount IN (?)", amounts).Find(&results).Error
	
	return
}
 
// GetFromPaymentMethod 通过payment_method获取内容  
func (obj *_EgOrderPaymentMgr) GetFromPaymentMethod(paymentMethod string) (results []*EgOrderPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("payment_method = ?", paymentMethod).Find(&results).Error
	
	return
}

// GetBatchFromPaymentMethod 批量唯一主键查找 
func (obj *_EgOrderPaymentMgr) GetBatchFromPaymentMethod(paymentMethods []string) (results []*EgOrderPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("payment_method IN (?)", paymentMethods).Find(&results).Error
	
	return
}
 
// GetFromConversionRate 通过conversion_rate获取内容  
func (obj *_EgOrderPaymentMgr) GetFromConversionRate(conversionRate float64) (results []*EgOrderPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("conversion_rate = ?", conversionRate).Find(&results).Error
	
	return
}

// GetBatchFromConversionRate 批量唯一主键查找 
func (obj *_EgOrderPaymentMgr) GetBatchFromConversionRate(conversionRates []float64) (results []*EgOrderPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("conversion_rate IN (?)", conversionRates).Find(&results).Error
	
	return
}
 
// GetFromTransactionID 通过transaction_id获取内容  
func (obj *_EgOrderPaymentMgr) GetFromTransactionID(transactionID string) (results []*EgOrderPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("transaction_id = ?", transactionID).Find(&results).Error
	
	return
}

// GetBatchFromTransactionID 批量唯一主键查找 
func (obj *_EgOrderPaymentMgr) GetBatchFromTransactionID(transactionIDs []string) (results []*EgOrderPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("transaction_id IN (?)", transactionIDs).Find(&results).Error
	
	return
}
 
// GetFromCardNumber 通过card_number获取内容  
func (obj *_EgOrderPaymentMgr) GetFromCardNumber(cardNumber string) (results []*EgOrderPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("card_number = ?", cardNumber).Find(&results).Error
	
	return
}

// GetBatchFromCardNumber 批量唯一主键查找 
func (obj *_EgOrderPaymentMgr) GetBatchFromCardNumber(cardNumbers []string) (results []*EgOrderPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("card_number IN (?)", cardNumbers).Find(&results).Error
	
	return
}
 
// GetFromCardBrand 通过card_brand获取内容  
func (obj *_EgOrderPaymentMgr) GetFromCardBrand(cardBrand string) (results []*EgOrderPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("card_brand = ?", cardBrand).Find(&results).Error
	
	return
}

// GetBatchFromCardBrand 批量唯一主键查找 
func (obj *_EgOrderPaymentMgr) GetBatchFromCardBrand(cardBrands []string) (results []*EgOrderPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("card_brand IN (?)", cardBrands).Find(&results).Error
	
	return
}
 
// GetFromCardExpiration 通过card_expiration获取内容  
func (obj *_EgOrderPaymentMgr) GetFromCardExpiration(cardExpiration string) (results []*EgOrderPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("card_expiration = ?", cardExpiration).Find(&results).Error
	
	return
}

// GetBatchFromCardExpiration 批量唯一主键查找 
func (obj *_EgOrderPaymentMgr) GetBatchFromCardExpiration(cardExpirations []string) (results []*EgOrderPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("card_expiration IN (?)", cardExpirations).Find(&results).Error
	
	return
}
 
// GetFromCardHolder 通过card_holder获取内容  
func (obj *_EgOrderPaymentMgr) GetFromCardHolder(cardHolder string) (results []*EgOrderPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("card_holder = ?", cardHolder).Find(&results).Error
	
	return
}

// GetBatchFromCardHolder 批量唯一主键查找 
func (obj *_EgOrderPaymentMgr) GetBatchFromCardHolder(cardHolders []string) (results []*EgOrderPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("card_holder IN (?)", cardHolders).Find(&results).Error
	
	return
}
 
// GetFromDateAdd 通过date_add获取内容  
func (obj *_EgOrderPaymentMgr) GetFromDateAdd(dateAdd time.Time) (results []*EgOrderPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error
	
	return
}

// GetBatchFromDateAdd 批量唯一主键查找 
func (obj *_EgOrderPaymentMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*EgOrderPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgOrderPaymentMgr) FetchByPrimaryKey(idOrderPayment int ) (result EgOrderPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_payment = ?", idOrderPayment).Find(&result).Error
	
	return
}
 

 
 // FetchIndexByOrderReference  获取多个内容
 func (obj *_EgOrderPaymentMgr) FetchIndexByOrderReference(orderReference string ) (results []*EgOrderPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("order_reference = ?", orderReference).Find(&results).Error
	
	return
}
 

	

