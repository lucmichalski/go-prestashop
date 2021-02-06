package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgCarrierGroupMgr struct {
	*_BaseMgr
}

// EgCarrierGroupMgr open func
func EgCarrierGroupMgr(db *gorm.DB) *_EgCarrierGroupMgr {
	if db == nil {
		panic(fmt.Errorf("EgCarrierGroupMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgCarrierGroupMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_carrier_group"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgCarrierGroupMgr) GetTableName() string {
	return "eg_carrier_group"
}

// Get 获取
func (obj *_EgCarrierGroupMgr) Get() (result EgCarrierGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgCarrierGroupMgr) Gets() (results []*EgCarrierGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDCarrier id_carrier获取 
func (obj *_EgCarrierGroupMgr) WithIDCarrier(idCarrier uint32) Option {
	return optionFunc(func(o *options) { o.query["id_carrier"] = idCarrier })
}

// WithIDGroup id_group获取 
func (obj *_EgCarrierGroupMgr) WithIDGroup(idGroup uint32) Option {
	return optionFunc(func(o *options) { o.query["id_group"] = idGroup })
}


// GetByOption 功能选项模式获取
func (obj *_EgCarrierGroupMgr) GetByOption(opts ...Option) (result EgCarrierGroup, err error) {
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
func (obj *_EgCarrierGroupMgr) GetByOptions(opts ...Option) (results []*EgCarrierGroup, err error) {
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
func (obj *_EgCarrierGroupMgr) GetFromIDCarrier(idCarrier uint32) (results []*EgCarrierGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_carrier = ?", idCarrier).Find(&results).Error
	
	return
}

// GetBatchFromIDCarrier 批量唯一主键查找 
func (obj *_EgCarrierGroupMgr) GetBatchFromIDCarrier(idCarriers []uint32) (results []*EgCarrierGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_carrier IN (?)", idCarriers).Find(&results).Error
	
	return
}
 
// GetFromIDGroup 通过id_group获取内容  
func (obj *_EgCarrierGroupMgr) GetFromIDGroup(idGroup uint32) (results []*EgCarrierGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_group = ?", idGroup).Find(&results).Error
	
	return
}

// GetBatchFromIDGroup 批量唯一主键查找 
func (obj *_EgCarrierGroupMgr) GetBatchFromIDGroup(idGroups []uint32) (results []*EgCarrierGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_group IN (?)", idGroups).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgCarrierGroupMgr) FetchByPrimaryKey(idCarrier uint32 ,idGroup uint32 ) (result EgCarrierGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_carrier = ? AND id_group = ?", idCarrier , idGroup).Find(&result).Error
	
	return
}
 

 

	

