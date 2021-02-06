package	model	
import (	
"gorm.io/gorm"	
"fmt"	
"context"	
)	

type _EgRangeWeightMgr struct {
	*_BaseMgr
}

// EgRangeWeightMgr open func
func EgRangeWeightMgr(db *gorm.DB) *_EgRangeWeightMgr {
	if db == nil {
		panic(fmt.Errorf("EgRangeWeightMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgRangeWeightMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_range_weight"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgRangeWeightMgr) GetTableName() string {
	return "eg_range_weight"
}

// Get 获取
func (obj *_EgRangeWeightMgr) Get() (result EgRangeWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgRangeWeightMgr) Gets() (results []*EgRangeWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDRangeWeight id_range_weight获取 
func (obj *_EgRangeWeightMgr) WithIDRangeWeight(idRangeWeight uint32) Option {
	return optionFunc(func(o *options) { o.query["id_range_weight"] = idRangeWeight })
}

// WithIDCarrier id_carrier获取 
func (obj *_EgRangeWeightMgr) WithIDCarrier(idCarrier uint32) Option {
	return optionFunc(func(o *options) { o.query["id_carrier"] = idCarrier })
}

// WithDelimiter1 delimiter1获取 
func (obj *_EgRangeWeightMgr) WithDelimiter1(delimiter1 float64) Option {
	return optionFunc(func(o *options) { o.query["delimiter1"] = delimiter1 })
}

// WithDelimiter2 delimiter2获取 
func (obj *_EgRangeWeightMgr) WithDelimiter2(delimiter2 float64) Option {
	return optionFunc(func(o *options) { o.query["delimiter2"] = delimiter2 })
}


// GetByOption 功能选项模式获取
func (obj *_EgRangeWeightMgr) GetByOption(opts ...Option) (result EgRangeWeight, err error) {
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
func (obj *_EgRangeWeightMgr) GetByOptions(opts ...Option) (results []*EgRangeWeight, err error) {
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


// GetFromIDRangeWeight 通过id_range_weight获取内容  
func (obj *_EgRangeWeightMgr)  GetFromIDRangeWeight(idRangeWeight uint32) (result EgRangeWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_range_weight = ?", idRangeWeight).Find(&result).Error
	
	return
}

// GetBatchFromIDRangeWeight 批量唯一主键查找 
func (obj *_EgRangeWeightMgr) GetBatchFromIDRangeWeight(idRangeWeights []uint32) (results []*EgRangeWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_range_weight IN (?)", idRangeWeights).Find(&results).Error
	
	return
}
 
// GetFromIDCarrier 通过id_carrier获取内容  
func (obj *_EgRangeWeightMgr) GetFromIDCarrier(idCarrier uint32) (results []*EgRangeWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_carrier = ?", idCarrier).Find(&results).Error
	
	return
}

// GetBatchFromIDCarrier 批量唯一主键查找 
func (obj *_EgRangeWeightMgr) GetBatchFromIDCarrier(idCarriers []uint32) (results []*EgRangeWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_carrier IN (?)", idCarriers).Find(&results).Error
	
	return
}
 
// GetFromDelimiter1 通过delimiter1获取内容  
func (obj *_EgRangeWeightMgr) GetFromDelimiter1(delimiter1 float64) (results []*EgRangeWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("delimiter1 = ?", delimiter1).Find(&results).Error
	
	return
}

// GetBatchFromDelimiter1 批量唯一主键查找 
func (obj *_EgRangeWeightMgr) GetBatchFromDelimiter1(delimiter1s []float64) (results []*EgRangeWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("delimiter1 IN (?)", delimiter1s).Find(&results).Error
	
	return
}
 
// GetFromDelimiter2 通过delimiter2获取内容  
func (obj *_EgRangeWeightMgr) GetFromDelimiter2(delimiter2 float64) (results []*EgRangeWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("delimiter2 = ?", delimiter2).Find(&results).Error
	
	return
}

// GetBatchFromDelimiter2 批量唯一主键查找 
func (obj *_EgRangeWeightMgr) GetBatchFromDelimiter2(delimiter2s []float64) (results []*EgRangeWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("delimiter2 IN (?)", delimiter2s).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgRangeWeightMgr) FetchByPrimaryKey(idRangeWeight uint32 ) (result EgRangeWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_range_weight = ?", idRangeWeight).Find(&result).Error
	
	return
}
 
 // FetchUniqueIndexByIDCarrier primay or index 获取唯一内容
 func (obj *_EgRangeWeightMgr) FetchUniqueIndexByIDCarrier(idCarrier uint32 ,delimiter1 float64 ,delimiter2 float64 ) (result EgRangeWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_carrier = ? AND delimiter1 = ? AND delimiter2 = ?", idCarrier , delimiter1 , delimiter2).Find(&result).Error
	
	return
}
 

 

	

