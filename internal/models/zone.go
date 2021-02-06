package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgZoneMgr struct {
	*_BaseMgr
}

// EgZoneMgr open func
func EgZoneMgr(db *gorm.DB) *_EgZoneMgr {
	if db == nil {
		panic(fmt.Errorf("EgZoneMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgZoneMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_zone"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgZoneMgr) GetTableName() string {
	return "eg_zone"
}

// Get 获取
func (obj *_EgZoneMgr) Get() (result EgZone, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgZoneMgr) Gets() (results []*EgZone, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDZone id_zone获取 
func (obj *_EgZoneMgr) WithIDZone(idZone uint32) Option {
	return optionFunc(func(o *options) { o.query["id_zone"] = idZone })
}

// WithName name获取 
func (obj *_EgZoneMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithActive active获取 
func (obj *_EgZoneMgr) WithActive(active bool) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}


// GetByOption 功能选项模式获取
func (obj *_EgZoneMgr) GetByOption(opts ...Option) (result EgZone, err error) {
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
func (obj *_EgZoneMgr) GetByOptions(opts ...Option) (results []*EgZone, err error) {
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


// GetFromIDZone 通过id_zone获取内容  
func (obj *_EgZoneMgr)  GetFromIDZone(idZone uint32) (result EgZone, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_zone = ?", idZone).Find(&result).Error
	
	return
}

// GetBatchFromIDZone 批量唯一主键查找 
func (obj *_EgZoneMgr) GetBatchFromIDZone(idZones []uint32) (results []*EgZone, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_zone IN (?)", idZones).Find(&results).Error
	
	return
}
 
// GetFromName 通过name获取内容  
func (obj *_EgZoneMgr) GetFromName(name string) (results []*EgZone, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error
	
	return
}

// GetBatchFromName 批量唯一主键查找 
func (obj *_EgZoneMgr) GetBatchFromName(names []string) (results []*EgZone, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error
	
	return
}
 
// GetFromActive 通过active获取内容  
func (obj *_EgZoneMgr) GetFromActive(active bool) (results []*EgZone, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active = ?", active).Find(&results).Error
	
	return
}

// GetBatchFromActive 批量唯一主键查找 
func (obj *_EgZoneMgr) GetBatchFromActive(actives []bool) (results []*EgZone, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active IN (?)", actives).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgZoneMgr) FetchByPrimaryKey(idZone uint32 ) (result EgZone, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_zone = ?", idZone).Find(&result).Error
	
	return
}
 

 

	

