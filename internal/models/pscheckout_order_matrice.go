package	model	
import (	
"gorm.io/gorm"	
"fmt"	
"context"	
)	

type _EgPscheckoutOrderMatriceMgr struct {
	*_BaseMgr
}

// EgPscheckoutOrderMatriceMgr open func
func EgPscheckoutOrderMatriceMgr(db *gorm.DB) *_EgPscheckoutOrderMatriceMgr {
	if db == nil {
		panic(fmt.Errorf("EgPscheckoutOrderMatriceMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgPscheckoutOrderMatriceMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_pscheckout_order_matrice"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgPscheckoutOrderMatriceMgr) GetTableName() string {
	return "eg_pscheckout_order_matrice"
}

// Get 获取
func (obj *_EgPscheckoutOrderMatriceMgr) Get() (result EgPscheckoutOrderMatrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgPscheckoutOrderMatriceMgr) Gets() (results []*EgPscheckoutOrderMatrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDOrderMatrice id_order_matrice获取 
func (obj *_EgPscheckoutOrderMatriceMgr) WithIDOrderMatrice(idOrderMatrice uint32) Option {
	return optionFunc(func(o *options) { o.query["id_order_matrice"] = idOrderMatrice })
}

// WithIDOrderPrestashop id_order_prestashop获取 
func (obj *_EgPscheckoutOrderMatriceMgr) WithIDOrderPrestashop(idOrderPrestashop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_order_prestashop"] = idOrderPrestashop })
}

// WithIDOrderPaypal id_order_paypal获取 
func (obj *_EgPscheckoutOrderMatriceMgr) WithIDOrderPaypal(idOrderPaypal string) Option {
	return optionFunc(func(o *options) { o.query["id_order_paypal"] = idOrderPaypal })
}


// GetByOption 功能选项模式获取
func (obj *_EgPscheckoutOrderMatriceMgr) GetByOption(opts ...Option) (result EgPscheckoutOrderMatrice, err error) {
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
func (obj *_EgPscheckoutOrderMatriceMgr) GetByOptions(opts ...Option) (results []*EgPscheckoutOrderMatrice, err error) {
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


// GetFromIDOrderMatrice 通过id_order_matrice获取内容  
func (obj *_EgPscheckoutOrderMatriceMgr)  GetFromIDOrderMatrice(idOrderMatrice uint32) (result EgPscheckoutOrderMatrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_matrice = ?", idOrderMatrice).Find(&result).Error
	
	return
}

// GetBatchFromIDOrderMatrice 批量唯一主键查找 
func (obj *_EgPscheckoutOrderMatriceMgr) GetBatchFromIDOrderMatrice(idOrderMatrices []uint32) (results []*EgPscheckoutOrderMatrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_matrice IN (?)", idOrderMatrices).Find(&results).Error
	
	return
}
 
// GetFromIDOrderPrestashop 通过id_order_prestashop获取内容  
func (obj *_EgPscheckoutOrderMatriceMgr) GetFromIDOrderPrestashop(idOrderPrestashop uint32) (results []*EgPscheckoutOrderMatrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_prestashop = ?", idOrderPrestashop).Find(&results).Error
	
	return
}

// GetBatchFromIDOrderPrestashop 批量唯一主键查找 
func (obj *_EgPscheckoutOrderMatriceMgr) GetBatchFromIDOrderPrestashop(idOrderPrestashops []uint32) (results []*EgPscheckoutOrderMatrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_prestashop IN (?)", idOrderPrestashops).Find(&results).Error
	
	return
}
 
// GetFromIDOrderPaypal 通过id_order_paypal获取内容  
func (obj *_EgPscheckoutOrderMatriceMgr) GetFromIDOrderPaypal(idOrderPaypal string) (results []*EgPscheckoutOrderMatrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_paypal = ?", idOrderPaypal).Find(&results).Error
	
	return
}

// GetBatchFromIDOrderPaypal 批量唯一主键查找 
func (obj *_EgPscheckoutOrderMatriceMgr) GetBatchFromIDOrderPaypal(idOrderPaypals []string) (results []*EgPscheckoutOrderMatrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_paypal IN (?)", idOrderPaypals).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgPscheckoutOrderMatriceMgr) FetchByPrimaryKey(idOrderMatrice uint32 ) (result EgPscheckoutOrderMatrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_matrice = ?", idOrderMatrice).Find(&result).Error
	
	return
}
 

 

	

