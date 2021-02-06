package	model	
import (	
"context"	
"gorm.io/gorm"	
"time"	
"fmt"	
)	

type _EgOrderMessageMgr struct {
	*_BaseMgr
}

// EgOrderMessageMgr open func
func EgOrderMessageMgr(db *gorm.DB) *_EgOrderMessageMgr {
	if db == nil {
		panic(fmt.Errorf("EgOrderMessageMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgOrderMessageMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_order_message"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgOrderMessageMgr) GetTableName() string {
	return "eg_order_message"
}

// Get 获取
func (obj *_EgOrderMessageMgr) Get() (result EgOrderMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgOrderMessageMgr) Gets() (results []*EgOrderMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDOrderMessage id_order_message获取 
func (obj *_EgOrderMessageMgr) WithIDOrderMessage(idOrderMessage uint32) Option {
	return optionFunc(func(o *options) { o.query["id_order_message"] = idOrderMessage })
}

// WithDateAdd date_add获取 
func (obj *_EgOrderMessageMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}


// GetByOption 功能选项模式获取
func (obj *_EgOrderMessageMgr) GetByOption(opts ...Option) (result EgOrderMessage, err error) {
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
func (obj *_EgOrderMessageMgr) GetByOptions(opts ...Option) (results []*EgOrderMessage, err error) {
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


// GetFromIDOrderMessage 通过id_order_message获取内容  
func (obj *_EgOrderMessageMgr)  GetFromIDOrderMessage(idOrderMessage uint32) (result EgOrderMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_message = ?", idOrderMessage).Find(&result).Error
	
	return
}

// GetBatchFromIDOrderMessage 批量唯一主键查找 
func (obj *_EgOrderMessageMgr) GetBatchFromIDOrderMessage(idOrderMessages []uint32) (results []*EgOrderMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_message IN (?)", idOrderMessages).Find(&results).Error
	
	return
}
 
// GetFromDateAdd 通过date_add获取内容  
func (obj *_EgOrderMessageMgr) GetFromDateAdd(dateAdd time.Time) (results []*EgOrderMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error
	
	return
}

// GetBatchFromDateAdd 批量唯一主键查找 
func (obj *_EgOrderMessageMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*EgOrderMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgOrderMessageMgr) FetchByPrimaryKey(idOrderMessage uint32 ) (result EgOrderMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_message = ?", idOrderMessage).Find(&result).Error
	
	return
}
 

 

	

