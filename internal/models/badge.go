package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgBadgeMgr struct {
	*_BaseMgr
}

// EgBadgeMgr open func
func EgBadgeMgr(db *gorm.DB) *_EgBadgeMgr {
	if db == nil {
		panic(fmt.Errorf("EgBadgeMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgBadgeMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_badge"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgBadgeMgr) GetTableName() string {
	return "eg_badge"
}

// Get 获取
func (obj *_EgBadgeMgr) Get() (result EgBadge, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgBadgeMgr) Gets() (results []*EgBadge, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDBadge id_badge获取 
func (obj *_EgBadgeMgr) WithIDBadge(idBadge int) Option {
	return optionFunc(func(o *options) { o.query["id_badge"] = idBadge })
}

// WithIDPsBadge id_ps_badge获取 
func (obj *_EgBadgeMgr) WithIDPsBadge(idPsBadge int) Option {
	return optionFunc(func(o *options) { o.query["id_ps_badge"] = idPsBadge })
}

// WithType type获取 
func (obj *_EgBadgeMgr) WithType(_type string) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithIDGroup id_group获取 
func (obj *_EgBadgeMgr) WithIDGroup(idGroup int) Option {
	return optionFunc(func(o *options) { o.query["id_group"] = idGroup })
}

// WithGroupPosition group_position获取 
func (obj *_EgBadgeMgr) WithGroupPosition(groupPosition int) Option {
	return optionFunc(func(o *options) { o.query["group_position"] = groupPosition })
}

// WithScoring scoring获取 
func (obj *_EgBadgeMgr) WithScoring(scoring int) Option {
	return optionFunc(func(o *options) { o.query["scoring"] = scoring })
}

// WithAwb awb获取 
func (obj *_EgBadgeMgr) WithAwb(awb int) Option {
	return optionFunc(func(o *options) { o.query["awb"] = awb })
}

// WithValidated validated获取 
func (obj *_EgBadgeMgr) WithValidated(validated bool) Option {
	return optionFunc(func(o *options) { o.query["validated"] = validated })
}


// GetByOption 功能选项模式获取
func (obj *_EgBadgeMgr) GetByOption(opts ...Option) (result EgBadge, err error) {
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
func (obj *_EgBadgeMgr) GetByOptions(opts ...Option) (results []*EgBadge, err error) {
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


// GetFromIDBadge 通过id_badge获取内容  
func (obj *_EgBadgeMgr)  GetFromIDBadge(idBadge int) (result EgBadge, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_badge = ?", idBadge).Find(&result).Error
	
	return
}

// GetBatchFromIDBadge 批量唯一主键查找 
func (obj *_EgBadgeMgr) GetBatchFromIDBadge(idBadges []int) (results []*EgBadge, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_badge IN (?)", idBadges).Find(&results).Error
	
	return
}
 
// GetFromIDPsBadge 通过id_ps_badge获取内容  
func (obj *_EgBadgeMgr) GetFromIDPsBadge(idPsBadge int) (results []*EgBadge, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_ps_badge = ?", idPsBadge).Find(&results).Error
	
	return
}

// GetBatchFromIDPsBadge 批量唯一主键查找 
func (obj *_EgBadgeMgr) GetBatchFromIDPsBadge(idPsBadges []int) (results []*EgBadge, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_ps_badge IN (?)", idPsBadges).Find(&results).Error
	
	return
}
 
// GetFromType 通过type获取内容  
func (obj *_EgBadgeMgr) GetFromType(_type string) (results []*EgBadge, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type = ?", _type).Find(&results).Error
	
	return
}

// GetBatchFromType 批量唯一主键查找 
func (obj *_EgBadgeMgr) GetBatchFromType(_types []string) (results []*EgBadge, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type IN (?)", _types).Find(&results).Error
	
	return
}
 
// GetFromIDGroup 通过id_group获取内容  
func (obj *_EgBadgeMgr) GetFromIDGroup(idGroup int) (results []*EgBadge, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_group = ?", idGroup).Find(&results).Error
	
	return
}

// GetBatchFromIDGroup 批量唯一主键查找 
func (obj *_EgBadgeMgr) GetBatchFromIDGroup(idGroups []int) (results []*EgBadge, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_group IN (?)", idGroups).Find(&results).Error
	
	return
}
 
// GetFromGroupPosition 通过group_position获取内容  
func (obj *_EgBadgeMgr) GetFromGroupPosition(groupPosition int) (results []*EgBadge, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("group_position = ?", groupPosition).Find(&results).Error
	
	return
}

// GetBatchFromGroupPosition 批量唯一主键查找 
func (obj *_EgBadgeMgr) GetBatchFromGroupPosition(groupPositions []int) (results []*EgBadge, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("group_position IN (?)", groupPositions).Find(&results).Error
	
	return
}
 
// GetFromScoring 通过scoring获取内容  
func (obj *_EgBadgeMgr) GetFromScoring(scoring int) (results []*EgBadge, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("scoring = ?", scoring).Find(&results).Error
	
	return
}

// GetBatchFromScoring 批量唯一主键查找 
func (obj *_EgBadgeMgr) GetBatchFromScoring(scorings []int) (results []*EgBadge, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("scoring IN (?)", scorings).Find(&results).Error
	
	return
}
 
// GetFromAwb 通过awb获取内容  
func (obj *_EgBadgeMgr) GetFromAwb(awb int) (results []*EgBadge, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("awb = ?", awb).Find(&results).Error
	
	return
}

// GetBatchFromAwb 批量唯一主键查找 
func (obj *_EgBadgeMgr) GetBatchFromAwb(awbs []int) (results []*EgBadge, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("awb IN (?)", awbs).Find(&results).Error
	
	return
}
 
// GetFromValidated 通过validated获取内容  
func (obj *_EgBadgeMgr) GetFromValidated(validated bool) (results []*EgBadge, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("validated = ?", validated).Find(&results).Error
	
	return
}

// GetBatchFromValidated 批量唯一主键查找 
func (obj *_EgBadgeMgr) GetBatchFromValidated(validateds []bool) (results []*EgBadge, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("validated IN (?)", validateds).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgBadgeMgr) FetchByPrimaryKey(idBadge int ) (result EgBadge, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_badge = ?", idBadge).Find(&result).Error
	
	return
}
 

 

	

