package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
"time"	
)	

type _EgManufacturerMgr struct {
	*_BaseMgr
}

// EgManufacturerMgr open func
func EgManufacturerMgr(db *gorm.DB) *_EgManufacturerMgr {
	if db == nil {
		panic(fmt.Errorf("EgManufacturerMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgManufacturerMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_manufacturer"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgManufacturerMgr) GetTableName() string {
	return "eg_manufacturer"
}

// Get 获取
func (obj *_EgManufacturerMgr) Get() (result EgManufacturer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgManufacturerMgr) Gets() (results []*EgManufacturer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDManufacturer id_manufacturer获取 
func (obj *_EgManufacturerMgr) WithIDManufacturer(idManufacturer uint32) Option {
	return optionFunc(func(o *options) { o.query["id_manufacturer"] = idManufacturer })
}

// WithName name获取 
func (obj *_EgManufacturerMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithDateAdd date_add获取 
func (obj *_EgManufacturerMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

// WithDateUpd date_upd获取 
func (obj *_EgManufacturerMgr) WithDateUpd(dateUpd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_upd"] = dateUpd })
}

// WithActive active获取 
func (obj *_EgManufacturerMgr) WithActive(active bool) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}


// GetByOption 功能选项模式获取
func (obj *_EgManufacturerMgr) GetByOption(opts ...Option) (result EgManufacturer, err error) {
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
func (obj *_EgManufacturerMgr) GetByOptions(opts ...Option) (results []*EgManufacturer, err error) {
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


// GetFromIDManufacturer 通过id_manufacturer获取内容  
func (obj *_EgManufacturerMgr)  GetFromIDManufacturer(idManufacturer uint32) (result EgManufacturer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_manufacturer = ?", idManufacturer).Find(&result).Error
	
	return
}

// GetBatchFromIDManufacturer 批量唯一主键查找 
func (obj *_EgManufacturerMgr) GetBatchFromIDManufacturer(idManufacturers []uint32) (results []*EgManufacturer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_manufacturer IN (?)", idManufacturers).Find(&results).Error
	
	return
}
 
// GetFromName 通过name获取内容  
func (obj *_EgManufacturerMgr) GetFromName(name string) (results []*EgManufacturer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error
	
	return
}

// GetBatchFromName 批量唯一主键查找 
func (obj *_EgManufacturerMgr) GetBatchFromName(names []string) (results []*EgManufacturer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error
	
	return
}
 
// GetFromDateAdd 通过date_add获取内容  
func (obj *_EgManufacturerMgr) GetFromDateAdd(dateAdd time.Time) (results []*EgManufacturer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error
	
	return
}

// GetBatchFromDateAdd 批量唯一主键查找 
func (obj *_EgManufacturerMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*EgManufacturer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error
	
	return
}
 
// GetFromDateUpd 通过date_upd获取内容  
func (obj *_EgManufacturerMgr) GetFromDateUpd(dateUpd time.Time) (results []*EgManufacturer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd = ?", dateUpd).Find(&results).Error
	
	return
}

// GetBatchFromDateUpd 批量唯一主键查找 
func (obj *_EgManufacturerMgr) GetBatchFromDateUpd(dateUpds []time.Time) (results []*EgManufacturer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd IN (?)", dateUpds).Find(&results).Error
	
	return
}
 
// GetFromActive 通过active获取内容  
func (obj *_EgManufacturerMgr) GetFromActive(active bool) (results []*EgManufacturer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active = ?", active).Find(&results).Error
	
	return
}

// GetBatchFromActive 批量唯一主键查找 
func (obj *_EgManufacturerMgr) GetBatchFromActive(actives []bool) (results []*EgManufacturer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active IN (?)", actives).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgManufacturerMgr) FetchByPrimaryKey(idManufacturer uint32 ) (result EgManufacturer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_manufacturer = ?", idManufacturer).Find(&result).Error
	
	return
}
 

 

	
