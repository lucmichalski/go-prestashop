package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgOperatingSystemMgr struct {
	*_BaseMgr
}

// EgOperatingSystemMgr open func
func EgOperatingSystemMgr(db *gorm.DB) *_EgOperatingSystemMgr {
	if db == nil {
		panic(fmt.Errorf("EgOperatingSystemMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgOperatingSystemMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_operating_system"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgOperatingSystemMgr) GetTableName() string {
	return "eg_operating_system"
}

// Get 获取
func (obj *_EgOperatingSystemMgr) Get() (result EgOperatingSystem, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgOperatingSystemMgr) Gets() (results []*EgOperatingSystem, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDOperatingSystem id_operating_system获取 
func (obj *_EgOperatingSystemMgr) WithIDOperatingSystem(idOperatingSystem uint32) Option {
	return optionFunc(func(o *options) { o.query["id_operating_system"] = idOperatingSystem })
}

// WithName name获取 
func (obj *_EgOperatingSystemMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}


// GetByOption 功能选项模式获取
func (obj *_EgOperatingSystemMgr) GetByOption(opts ...Option) (result EgOperatingSystem, err error) {
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
func (obj *_EgOperatingSystemMgr) GetByOptions(opts ...Option) (results []*EgOperatingSystem, err error) {
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


// GetFromIDOperatingSystem 通过id_operating_system获取内容  
func (obj *_EgOperatingSystemMgr)  GetFromIDOperatingSystem(idOperatingSystem uint32) (result EgOperatingSystem, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_operating_system = ?", idOperatingSystem).Find(&result).Error
	
	return
}

// GetBatchFromIDOperatingSystem 批量唯一主键查找 
func (obj *_EgOperatingSystemMgr) GetBatchFromIDOperatingSystem(idOperatingSystems []uint32) (results []*EgOperatingSystem, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_operating_system IN (?)", idOperatingSystems).Find(&results).Error
	
	return
}
 
// GetFromName 通过name获取内容  
func (obj *_EgOperatingSystemMgr) GetFromName(name string) (results []*EgOperatingSystem, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error
	
	return
}

// GetBatchFromName 批量唯一主键查找 
func (obj *_EgOperatingSystemMgr) GetBatchFromName(names []string) (results []*EgOperatingSystem, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgOperatingSystemMgr) FetchByPrimaryKey(idOperatingSystem uint32 ) (result EgOperatingSystem, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_operating_system = ?", idOperatingSystem).Find(&result).Error
	
	return
}
 

 

	

