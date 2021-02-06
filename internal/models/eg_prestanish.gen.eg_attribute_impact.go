package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgAttributeImpactMgr struct {
	*_BaseMgr
}

// EgAttributeImpactMgr open func
func EgAttributeImpactMgr(db *gorm.DB) *_EgAttributeImpactMgr {
	if db == nil {
		panic(fmt.Errorf("EgAttributeImpactMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgAttributeImpactMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_attribute_impact"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgAttributeImpactMgr) GetTableName() string {
	return "eg_attribute_impact"
}

// Get 获取
func (obj *_EgAttributeImpactMgr) Get() (result EgAttributeImpact, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgAttributeImpactMgr) Gets() (results []*EgAttributeImpact, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDAttributeImpact id_attribute_impact获取 
func (obj *_EgAttributeImpactMgr) WithIDAttributeImpact(idAttributeImpact uint32) Option {
	return optionFunc(func(o *options) { o.query["id_attribute_impact"] = idAttributeImpact })
}

// WithIDProduct id_product获取 
func (obj *_EgAttributeImpactMgr) WithIDProduct(idProduct uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product"] = idProduct })
}

// WithIDAttribute id_attribute获取 
func (obj *_EgAttributeImpactMgr) WithIDAttribute(idAttribute uint32) Option {
	return optionFunc(func(o *options) { o.query["id_attribute"] = idAttribute })
}

// WithWeight weight获取 
func (obj *_EgAttributeImpactMgr) WithWeight(weight float64) Option {
	return optionFunc(func(o *options) { o.query["weight"] = weight })
}

// WithPrice price获取 
func (obj *_EgAttributeImpactMgr) WithPrice(price float64) Option {
	return optionFunc(func(o *options) { o.query["price"] = price })
}


// GetByOption 功能选项模式获取
func (obj *_EgAttributeImpactMgr) GetByOption(opts ...Option) (result EgAttributeImpact, err error) {
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
func (obj *_EgAttributeImpactMgr) GetByOptions(opts ...Option) (results []*EgAttributeImpact, err error) {
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


// GetFromIDAttributeImpact 通过id_attribute_impact获取内容  
func (obj *_EgAttributeImpactMgr)  GetFromIDAttributeImpact(idAttributeImpact uint32) (result EgAttributeImpact, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attribute_impact = ?", idAttributeImpact).Find(&result).Error
	
	return
}

// GetBatchFromIDAttributeImpact 批量唯一主键查找 
func (obj *_EgAttributeImpactMgr) GetBatchFromIDAttributeImpact(idAttributeImpacts []uint32) (results []*EgAttributeImpact, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attribute_impact IN (?)", idAttributeImpacts).Find(&results).Error
	
	return
}
 
// GetFromIDProduct 通过id_product获取内容  
func (obj *_EgAttributeImpactMgr) GetFromIDProduct(idProduct uint32) (results []*EgAttributeImpact, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ?", idProduct).Find(&results).Error
	
	return
}

// GetBatchFromIDProduct 批量唯一主键查找 
func (obj *_EgAttributeImpactMgr) GetBatchFromIDProduct(idProducts []uint32) (results []*EgAttributeImpact, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product IN (?)", idProducts).Find(&results).Error
	
	return
}
 
// GetFromIDAttribute 通过id_attribute获取内容  
func (obj *_EgAttributeImpactMgr) GetFromIDAttribute(idAttribute uint32) (results []*EgAttributeImpact, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attribute = ?", idAttribute).Find(&results).Error
	
	return
}

// GetBatchFromIDAttribute 批量唯一主键查找 
func (obj *_EgAttributeImpactMgr) GetBatchFromIDAttribute(idAttributes []uint32) (results []*EgAttributeImpact, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attribute IN (?)", idAttributes).Find(&results).Error
	
	return
}
 
// GetFromWeight 通过weight获取内容  
func (obj *_EgAttributeImpactMgr) GetFromWeight(weight float64) (results []*EgAttributeImpact, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("weight = ?", weight).Find(&results).Error
	
	return
}

// GetBatchFromWeight 批量唯一主键查找 
func (obj *_EgAttributeImpactMgr) GetBatchFromWeight(weights []float64) (results []*EgAttributeImpact, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("weight IN (?)", weights).Find(&results).Error
	
	return
}
 
// GetFromPrice 通过price获取内容  
func (obj *_EgAttributeImpactMgr) GetFromPrice(price float64) (results []*EgAttributeImpact, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("price = ?", price).Find(&results).Error
	
	return
}

// GetBatchFromPrice 批量唯一主键查找 
func (obj *_EgAttributeImpactMgr) GetBatchFromPrice(prices []float64) (results []*EgAttributeImpact, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("price IN (?)", prices).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgAttributeImpactMgr) FetchByPrimaryKey(idAttributeImpact uint32 ) (result EgAttributeImpact, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attribute_impact = ?", idAttributeImpact).Find(&result).Error
	
	return
}
 
 // FetchUniqueIndexByIDProduct primay or index 获取唯一内容
 func (obj *_EgAttributeImpactMgr) FetchUniqueIndexByIDProduct(idProduct uint32 ,idAttribute uint32 ) (result EgAttributeImpact, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ? AND id_attribute = ?", idProduct , idAttribute).Find(&result).Error
	
	return
}
 

 

	

