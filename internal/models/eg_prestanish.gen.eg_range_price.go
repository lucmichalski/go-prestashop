package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgRangePriceMgr struct {
	*_BaseMgr
}

// EgRangePriceMgr open func
func EgRangePriceMgr(db *gorm.DB) *_EgRangePriceMgr {
	if db == nil {
		panic(fmt.Errorf("EgRangePriceMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgRangePriceMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_range_price"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgRangePriceMgr) GetTableName() string {
	return "eg_range_price"
}

// Get 获取
func (obj *_EgRangePriceMgr) Get() (result EgRangePrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgRangePriceMgr) Gets() (results []*EgRangePrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDRangePrice id_range_price获取 
func (obj *_EgRangePriceMgr) WithIDRangePrice(idRangePrice uint32) Option {
	return optionFunc(func(o *options) { o.query["id_range_price"] = idRangePrice })
}

// WithIDCarrier id_carrier获取 
func (obj *_EgRangePriceMgr) WithIDCarrier(idCarrier uint32) Option {
	return optionFunc(func(o *options) { o.query["id_carrier"] = idCarrier })
}

// WithDelimiter1 delimiter1获取 
func (obj *_EgRangePriceMgr) WithDelimiter1(delimiter1 float64) Option {
	return optionFunc(func(o *options) { o.query["delimiter1"] = delimiter1 })
}

// WithDelimiter2 delimiter2获取 
func (obj *_EgRangePriceMgr) WithDelimiter2(delimiter2 float64) Option {
	return optionFunc(func(o *options) { o.query["delimiter2"] = delimiter2 })
}


// GetByOption 功能选项模式获取
func (obj *_EgRangePriceMgr) GetByOption(opts ...Option) (result EgRangePrice, err error) {
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
func (obj *_EgRangePriceMgr) GetByOptions(opts ...Option) (results []*EgRangePrice, err error) {
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


// GetFromIDRangePrice 通过id_range_price获取内容  
func (obj *_EgRangePriceMgr)  GetFromIDRangePrice(idRangePrice uint32) (result EgRangePrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_range_price = ?", idRangePrice).Find(&result).Error
	
	return
}

// GetBatchFromIDRangePrice 批量唯一主键查找 
func (obj *_EgRangePriceMgr) GetBatchFromIDRangePrice(idRangePrices []uint32) (results []*EgRangePrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_range_price IN (?)", idRangePrices).Find(&results).Error
	
	return
}
 
// GetFromIDCarrier 通过id_carrier获取内容  
func (obj *_EgRangePriceMgr) GetFromIDCarrier(idCarrier uint32) (results []*EgRangePrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_carrier = ?", idCarrier).Find(&results).Error
	
	return
}

// GetBatchFromIDCarrier 批量唯一主键查找 
func (obj *_EgRangePriceMgr) GetBatchFromIDCarrier(idCarriers []uint32) (results []*EgRangePrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_carrier IN (?)", idCarriers).Find(&results).Error
	
	return
}
 
// GetFromDelimiter1 通过delimiter1获取内容  
func (obj *_EgRangePriceMgr) GetFromDelimiter1(delimiter1 float64) (results []*EgRangePrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("delimiter1 = ?", delimiter1).Find(&results).Error
	
	return
}

// GetBatchFromDelimiter1 批量唯一主键查找 
func (obj *_EgRangePriceMgr) GetBatchFromDelimiter1(delimiter1s []float64) (results []*EgRangePrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("delimiter1 IN (?)", delimiter1s).Find(&results).Error
	
	return
}
 
// GetFromDelimiter2 通过delimiter2获取内容  
func (obj *_EgRangePriceMgr) GetFromDelimiter2(delimiter2 float64) (results []*EgRangePrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("delimiter2 = ?", delimiter2).Find(&results).Error
	
	return
}

// GetBatchFromDelimiter2 批量唯一主键查找 
func (obj *_EgRangePriceMgr) GetBatchFromDelimiter2(delimiter2s []float64) (results []*EgRangePrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("delimiter2 IN (?)", delimiter2s).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgRangePriceMgr) FetchByPrimaryKey(idRangePrice uint32 ) (result EgRangePrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_range_price = ?", idRangePrice).Find(&result).Error
	
	return
}
 
 // FetchUniqueIndexByIDCarrier primay or index 获取唯一内容
 func (obj *_EgRangePriceMgr) FetchUniqueIndexByIDCarrier(idCarrier uint32 ,delimiter1 float64 ,delimiter2 float64 ) (result EgRangePrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_carrier = ? AND delimiter1 = ? AND delimiter2 = ?", idCarrier , delimiter1 , delimiter2).Find(&result).Error
	
	return
}
 

 

	

