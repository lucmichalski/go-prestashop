package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgConditionBadgeMgr struct {
	*_BaseMgr
}

// EgConditionBadgeMgr open func
func EgConditionBadgeMgr(db *gorm.DB) *_EgConditionBadgeMgr {
	if db == nil {
		panic(fmt.Errorf("EgConditionBadgeMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgConditionBadgeMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_condition_badge"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgConditionBadgeMgr) GetTableName() string {
	return "eg_condition_badge"
}

// Get 获取
func (obj *_EgConditionBadgeMgr) Get() (result EgConditionBadge, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgConditionBadgeMgr) Gets() (results []*EgConditionBadge, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDCondition id_condition获取 
func (obj *_EgConditionBadgeMgr) WithIDCondition(idCondition int) Option {
	return optionFunc(func(o *options) { o.query["id_condition"] = idCondition })
}

// WithIDBadge id_badge获取 
func (obj *_EgConditionBadgeMgr) WithIDBadge(idBadge int) Option {
	return optionFunc(func(o *options) { o.query["id_badge"] = idBadge })
}


// GetByOption 功能选项模式获取
func (obj *_EgConditionBadgeMgr) GetByOption(opts ...Option) (result EgConditionBadge, err error) {
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
func (obj *_EgConditionBadgeMgr) GetByOptions(opts ...Option) (results []*EgConditionBadge, err error) {
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


// GetFromIDCondition 通过id_condition获取内容  
func (obj *_EgConditionBadgeMgr) GetFromIDCondition(idCondition int) (results []*EgConditionBadge, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_condition = ?", idCondition).Find(&results).Error
	
	return
}

// GetBatchFromIDCondition 批量唯一主键查找 
func (obj *_EgConditionBadgeMgr) GetBatchFromIDCondition(idConditions []int) (results []*EgConditionBadge, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_condition IN (?)", idConditions).Find(&results).Error
	
	return
}
 
// GetFromIDBadge 通过id_badge获取内容  
func (obj *_EgConditionBadgeMgr) GetFromIDBadge(idBadge int) (results []*EgConditionBadge, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_badge = ?", idBadge).Find(&results).Error
	
	return
}

// GetBatchFromIDBadge 批量唯一主键查找 
func (obj *_EgConditionBadgeMgr) GetBatchFromIDBadge(idBadges []int) (results []*EgConditionBadge, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_badge IN (?)", idBadges).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgConditionBadgeMgr) FetchByPrimaryKey(idCondition int ,idBadge int ) (result EgConditionBadge, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_condition = ? AND id_badge = ?", idCondition , idBadge).Find(&result).Error
	
	return
}
 

 

	

