package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgOrderReturnStateMgr struct {
	*_BaseMgr
}

// EgOrderReturnStateMgr open func
func EgOrderReturnStateMgr(db *gorm.DB) *_EgOrderReturnStateMgr {
	if db == nil {
		panic(fmt.Errorf("EgOrderReturnStateMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgOrderReturnStateMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_order_return_state"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgOrderReturnStateMgr) GetTableName() string {
	return "eg_order_return_state"
}

// Get 获取
func (obj *_EgOrderReturnStateMgr) Get() (result EgOrderReturnState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgOrderReturnStateMgr) Gets() (results []*EgOrderReturnState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDOrderReturnState id_order_return_state获取 
func (obj *_EgOrderReturnStateMgr) WithIDOrderReturnState(idOrderReturnState uint32) Option {
	return optionFunc(func(o *options) { o.query["id_order_return_state"] = idOrderReturnState })
}

// WithColor color获取 
func (obj *_EgOrderReturnStateMgr) WithColor(color string) Option {
	return optionFunc(func(o *options) { o.query["color"] = color })
}


// GetByOption 功能选项模式获取
func (obj *_EgOrderReturnStateMgr) GetByOption(opts ...Option) (result EgOrderReturnState, err error) {
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
func (obj *_EgOrderReturnStateMgr) GetByOptions(opts ...Option) (results []*EgOrderReturnState, err error) {
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


// GetFromIDOrderReturnState 通过id_order_return_state获取内容  
func (obj *_EgOrderReturnStateMgr)  GetFromIDOrderReturnState(idOrderReturnState uint32) (result EgOrderReturnState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_return_state = ?", idOrderReturnState).Find(&result).Error
	
	return
}

// GetBatchFromIDOrderReturnState 批量唯一主键查找 
func (obj *_EgOrderReturnStateMgr) GetBatchFromIDOrderReturnState(idOrderReturnStates []uint32) (results []*EgOrderReturnState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_return_state IN (?)", idOrderReturnStates).Find(&results).Error
	
	return
}
 
// GetFromColor 通过color获取内容  
func (obj *_EgOrderReturnStateMgr) GetFromColor(color string) (results []*EgOrderReturnState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("color = ?", color).Find(&results).Error
	
	return
}

// GetBatchFromColor 批量唯一主键查找 
func (obj *_EgOrderReturnStateMgr) GetBatchFromColor(colors []string) (results []*EgOrderReturnState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("color IN (?)", colors).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgOrderReturnStateMgr) FetchByPrimaryKey(idOrderReturnState uint32 ) (result EgOrderReturnState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_return_state = ?", idOrderReturnState).Find(&result).Error
	
	return
}
 

 

	

