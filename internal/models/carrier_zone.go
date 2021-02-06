package	model	
import (	
"gorm.io/gorm"	
"fmt"	
"context"	
)	

type _EgCarrierZoneMgr struct {
	*_BaseMgr
}

// EgCarrierZoneMgr open func
func EgCarrierZoneMgr(db *gorm.DB) *_EgCarrierZoneMgr {
	if db == nil {
		panic(fmt.Errorf("EgCarrierZoneMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgCarrierZoneMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_carrier_zone"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgCarrierZoneMgr) GetTableName() string {
	return "eg_carrier_zone"
}

// Get 获取
func (obj *_EgCarrierZoneMgr) Get() (result EgCarrierZone, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgCarrierZoneMgr) Gets() (results []*EgCarrierZone, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDCarrier id_carrier获取 
func (obj *_EgCarrierZoneMgr) WithIDCarrier(idCarrier uint32) Option {
	return optionFunc(func(o *options) { o.query["id_carrier"] = idCarrier })
}

// WithIDZone id_zone获取 
func (obj *_EgCarrierZoneMgr) WithIDZone(idZone uint32) Option {
	return optionFunc(func(o *options) { o.query["id_zone"] = idZone })
}


// GetByOption 功能选项模式获取
func (obj *_EgCarrierZoneMgr) GetByOption(opts ...Option) (result EgCarrierZone, err error) {
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
func (obj *_EgCarrierZoneMgr) GetByOptions(opts ...Option) (results []*EgCarrierZone, err error) {
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


// GetFromIDCarrier 通过id_carrier获取内容  
func (obj *_EgCarrierZoneMgr) GetFromIDCarrier(idCarrier uint32) (results []*EgCarrierZone, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_carrier = ?", idCarrier).Find(&results).Error
	
	return
}

// GetBatchFromIDCarrier 批量唯一主键查找 
func (obj *_EgCarrierZoneMgr) GetBatchFromIDCarrier(idCarriers []uint32) (results []*EgCarrierZone, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_carrier IN (?)", idCarriers).Find(&results).Error
	
	return
}
 
// GetFromIDZone 通过id_zone获取内容  
func (obj *_EgCarrierZoneMgr) GetFromIDZone(idZone uint32) (results []*EgCarrierZone, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_zone = ?", idZone).Find(&results).Error
	
	return
}

// GetBatchFromIDZone 批量唯一主键查找 
func (obj *_EgCarrierZoneMgr) GetBatchFromIDZone(idZones []uint32) (results []*EgCarrierZone, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_zone IN (?)", idZones).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgCarrierZoneMgr) FetchByPrimaryKey(idCarrier uint32 ,idZone uint32 ) (result EgCarrierZone, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_carrier = ? AND id_zone = ?", idCarrier , idZone).Find(&result).Error
	
	return
}
 

 

	

