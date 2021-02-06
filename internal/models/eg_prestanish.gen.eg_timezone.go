package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgTimezoneMgr struct {
	*_BaseMgr
}

// EgTimezoneMgr open func
func EgTimezoneMgr(db *gorm.DB) *_EgTimezoneMgr {
	if db == nil {
		panic(fmt.Errorf("EgTimezoneMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgTimezoneMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_timezone"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgTimezoneMgr) GetTableName() string {
	return "eg_timezone"
}

// Get 获取
func (obj *_EgTimezoneMgr) Get() (result EgTimezone, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgTimezoneMgr) Gets() (results []*EgTimezone, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDTimezone id_timezone获取 
func (obj *_EgTimezoneMgr) WithIDTimezone(idTimezone uint32) Option {
	return optionFunc(func(o *options) { o.query["id_timezone"] = idTimezone })
}

// WithName name获取 
func (obj *_EgTimezoneMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}


// GetByOption 功能选项模式获取
func (obj *_EgTimezoneMgr) GetByOption(opts ...Option) (result EgTimezone, err error) {
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
func (obj *_EgTimezoneMgr) GetByOptions(opts ...Option) (results []*EgTimezone, err error) {
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


// GetFromIDTimezone 通过id_timezone获取内容  
func (obj *_EgTimezoneMgr)  GetFromIDTimezone(idTimezone uint32) (result EgTimezone, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_timezone = ?", idTimezone).Find(&result).Error
	
	return
}

// GetBatchFromIDTimezone 批量唯一主键查找 
func (obj *_EgTimezoneMgr) GetBatchFromIDTimezone(idTimezones []uint32) (results []*EgTimezone, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_timezone IN (?)", idTimezones).Find(&results).Error
	
	return
}
 
// GetFromName 通过name获取内容  
func (obj *_EgTimezoneMgr) GetFromName(name string) (results []*EgTimezone, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error
	
	return
}

// GetBatchFromName 批量唯一主键查找 
func (obj *_EgTimezoneMgr) GetBatchFromName(names []string) (results []*EgTimezone, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgTimezoneMgr) FetchByPrimaryKey(idTimezone uint32 ) (result EgTimezone, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_timezone = ?", idTimezone).Find(&result).Error
	
	return
}
 

 

	

