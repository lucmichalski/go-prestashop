package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
"gorm.io/datatypes"	
)	

type _EgProductSaleMgr struct {
	*_BaseMgr
}

// EgProductSaleMgr open func
func EgProductSaleMgr(db *gorm.DB) *_EgProductSaleMgr {
	if db == nil {
		panic(fmt.Errorf("EgProductSaleMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgProductSaleMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_product_sale"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgProductSaleMgr) GetTableName() string {
	return "eg_product_sale"
}

// Get 获取
func (obj *_EgProductSaleMgr) Get() (result EgProductSale, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgProductSaleMgr) Gets() (results []*EgProductSale, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDProduct id_product获取 
func (obj *_EgProductSaleMgr) WithIDProduct(idProduct uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product"] = idProduct })
}

// WithQuantity quantity获取 
func (obj *_EgProductSaleMgr) WithQuantity(quantity uint32) Option {
	return optionFunc(func(o *options) { o.query["quantity"] = quantity })
}

// WithSaleNbr sale_nbr获取 
func (obj *_EgProductSaleMgr) WithSaleNbr(saleNbr uint32) Option {
	return optionFunc(func(o *options) { o.query["sale_nbr"] = saleNbr })
}

// WithDateUpd date_upd获取 
func (obj *_EgProductSaleMgr) WithDateUpd(dateUpd datatypes.Date) Option {
	return optionFunc(func(o *options) { o.query["date_upd"] = dateUpd })
}


// GetByOption 功能选项模式获取
func (obj *_EgProductSaleMgr) GetByOption(opts ...Option) (result EgProductSale, err error) {
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
func (obj *_EgProductSaleMgr) GetByOptions(opts ...Option) (results []*EgProductSale, err error) {
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


// GetFromIDProduct 通过id_product获取内容  
func (obj *_EgProductSaleMgr)  GetFromIDProduct(idProduct uint32) (result EgProductSale, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ?", idProduct).Find(&result).Error
	
	return
}

// GetBatchFromIDProduct 批量唯一主键查找 
func (obj *_EgProductSaleMgr) GetBatchFromIDProduct(idProducts []uint32) (results []*EgProductSale, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product IN (?)", idProducts).Find(&results).Error
	
	return
}
 
// GetFromQuantity 通过quantity获取内容  
func (obj *_EgProductSaleMgr) GetFromQuantity(quantity uint32) (results []*EgProductSale, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("quantity = ?", quantity).Find(&results).Error
	
	return
}

// GetBatchFromQuantity 批量唯一主键查找 
func (obj *_EgProductSaleMgr) GetBatchFromQuantity(quantitys []uint32) (results []*EgProductSale, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("quantity IN (?)", quantitys).Find(&results).Error
	
	return
}
 
// GetFromSaleNbr 通过sale_nbr获取内容  
func (obj *_EgProductSaleMgr) GetFromSaleNbr(saleNbr uint32) (results []*EgProductSale, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("sale_nbr = ?", saleNbr).Find(&results).Error
	
	return
}

// GetBatchFromSaleNbr 批量唯一主键查找 
func (obj *_EgProductSaleMgr) GetBatchFromSaleNbr(saleNbrs []uint32) (results []*EgProductSale, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("sale_nbr IN (?)", saleNbrs).Find(&results).Error
	
	return
}
 
// GetFromDateUpd 通过date_upd获取内容  
func (obj *_EgProductSaleMgr) GetFromDateUpd(dateUpd datatypes.Date) (results []*EgProductSale, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd = ?", dateUpd).Find(&results).Error
	
	return
}

// GetBatchFromDateUpd 批量唯一主键查找 
func (obj *_EgProductSaleMgr) GetBatchFromDateUpd(dateUpds []datatypes.Date) (results []*EgProductSale, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd IN (?)", dateUpds).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgProductSaleMgr) FetchByPrimaryKey(idProduct uint32 ) (result EgProductSale, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ?", idProduct).Find(&result).Error
	
	return
}
 

 
 // FetchIndexByQuantity  获取多个内容
 func (obj *_EgProductSaleMgr) FetchIndexByQuantity(quantity uint32 ) (results []*EgProductSale, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("quantity = ?", quantity).Find(&results).Error
	
	return
}
 

	

