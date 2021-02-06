package	model	
import (	
"context"	
"gorm.io/gorm"	
"time"	
"fmt"	
)	

type _EgDateRangeMgr struct {
	*_BaseMgr
}

// EgDateRangeMgr open func
func EgDateRangeMgr(db *gorm.DB) *_EgDateRangeMgr {
	if db == nil {
		panic(fmt.Errorf("EgDateRangeMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgDateRangeMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_date_range"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgDateRangeMgr) GetTableName() string {
	return "eg_date_range"
}

// Get 获取
func (obj *_EgDateRangeMgr) Get() (result EgDateRange, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgDateRangeMgr) Gets() (results []*EgDateRange, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDDateRange id_date_range获取 
func (obj *_EgDateRangeMgr) WithIDDateRange(idDateRange uint32) Option {
	return optionFunc(func(o *options) { o.query["id_date_range"] = idDateRange })
}

// WithTimeStart time_start获取 
func (obj *_EgDateRangeMgr) WithTimeStart(timeStart time.Time) Option {
	return optionFunc(func(o *options) { o.query["time_start"] = timeStart })
}

// WithTimeEnd time_end获取 
func (obj *_EgDateRangeMgr) WithTimeEnd(timeEnd time.Time) Option {
	return optionFunc(func(o *options) { o.query["time_end"] = timeEnd })
}


// GetByOption 功能选项模式获取
func (obj *_EgDateRangeMgr) GetByOption(opts ...Option) (result EgDateRange, err error) {
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
func (obj *_EgDateRangeMgr) GetByOptions(opts ...Option) (results []*EgDateRange, err error) {
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


// GetFromIDDateRange 通过id_date_range获取内容  
func (obj *_EgDateRangeMgr)  GetFromIDDateRange(idDateRange uint32) (result EgDateRange, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_date_range = ?", idDateRange).Find(&result).Error
	
	return
}

// GetBatchFromIDDateRange 批量唯一主键查找 
func (obj *_EgDateRangeMgr) GetBatchFromIDDateRange(idDateRanges []uint32) (results []*EgDateRange, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_date_range IN (?)", idDateRanges).Find(&results).Error
	
	return
}
 
// GetFromTimeStart 通过time_start获取内容  
func (obj *_EgDateRangeMgr) GetFromTimeStart(timeStart time.Time) (results []*EgDateRange, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("time_start = ?", timeStart).Find(&results).Error
	
	return
}

// GetBatchFromTimeStart 批量唯一主键查找 
func (obj *_EgDateRangeMgr) GetBatchFromTimeStart(timeStarts []time.Time) (results []*EgDateRange, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("time_start IN (?)", timeStarts).Find(&results).Error
	
	return
}
 
// GetFromTimeEnd 通过time_end获取内容  
func (obj *_EgDateRangeMgr) GetFromTimeEnd(timeEnd time.Time) (results []*EgDateRange, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("time_end = ?", timeEnd).Find(&results).Error
	
	return
}

// GetBatchFromTimeEnd 批量唯一主键查找 
func (obj *_EgDateRangeMgr) GetBatchFromTimeEnd(timeEnds []time.Time) (results []*EgDateRange, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("time_end IN (?)", timeEnds).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgDateRangeMgr) FetchByPrimaryKey(idDateRange uint32 ) (result EgDateRange, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_date_range = ?", idDateRange).Find(&result).Error
	
	return
}
 

 

	

