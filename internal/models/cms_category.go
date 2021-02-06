package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
"time"	
)	

type _EgCmsCategoryMgr struct {
	*_BaseMgr
}

// EgCmsCategoryMgr open func
func EgCmsCategoryMgr(db *gorm.DB) *_EgCmsCategoryMgr {
	if db == nil {
		panic(fmt.Errorf("EgCmsCategoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgCmsCategoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_cms_category"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgCmsCategoryMgr) GetTableName() string {
	return "eg_cms_category"
}

// Get 获取
func (obj *_EgCmsCategoryMgr) Get() (result EgCmsCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgCmsCategoryMgr) Gets() (results []*EgCmsCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDCmsCategory id_cms_category获取 
func (obj *_EgCmsCategoryMgr) WithIDCmsCategory(idCmsCategory uint32) Option {
	return optionFunc(func(o *options) { o.query["id_cms_category"] = idCmsCategory })
}

// WithIDParent id_parent获取 
func (obj *_EgCmsCategoryMgr) WithIDParent(idParent uint32) Option {
	return optionFunc(func(o *options) { o.query["id_parent"] = idParent })
}

// WithLevelDepth level_depth获取 
func (obj *_EgCmsCategoryMgr) WithLevelDepth(levelDepth uint8) Option {
	return optionFunc(func(o *options) { o.query["level_depth"] = levelDepth })
}

// WithActive active获取 
func (obj *_EgCmsCategoryMgr) WithActive(active bool) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}

// WithDateAdd date_add获取 
func (obj *_EgCmsCategoryMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

// WithDateUpd date_upd获取 
func (obj *_EgCmsCategoryMgr) WithDateUpd(dateUpd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_upd"] = dateUpd })
}

// WithPosition position获取 
func (obj *_EgCmsCategoryMgr) WithPosition(position uint32) Option {
	return optionFunc(func(o *options) { o.query["position"] = position })
}


// GetByOption 功能选项模式获取
func (obj *_EgCmsCategoryMgr) GetByOption(opts ...Option) (result EgCmsCategory, err error) {
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
func (obj *_EgCmsCategoryMgr) GetByOptions(opts ...Option) (results []*EgCmsCategory, err error) {
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


// GetFromIDCmsCategory 通过id_cms_category获取内容  
func (obj *_EgCmsCategoryMgr)  GetFromIDCmsCategory(idCmsCategory uint32) (result EgCmsCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cms_category = ?", idCmsCategory).Find(&result).Error
	
	return
}

// GetBatchFromIDCmsCategory 批量唯一主键查找 
func (obj *_EgCmsCategoryMgr) GetBatchFromIDCmsCategory(idCmsCategorys []uint32) (results []*EgCmsCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cms_category IN (?)", idCmsCategorys).Find(&results).Error
	
	return
}
 
// GetFromIDParent 通过id_parent获取内容  
func (obj *_EgCmsCategoryMgr) GetFromIDParent(idParent uint32) (results []*EgCmsCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_parent = ?", idParent).Find(&results).Error
	
	return
}

// GetBatchFromIDParent 批量唯一主键查找 
func (obj *_EgCmsCategoryMgr) GetBatchFromIDParent(idParents []uint32) (results []*EgCmsCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_parent IN (?)", idParents).Find(&results).Error
	
	return
}
 
// GetFromLevelDepth 通过level_depth获取内容  
func (obj *_EgCmsCategoryMgr) GetFromLevelDepth(levelDepth uint8) (results []*EgCmsCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("level_depth = ?", levelDepth).Find(&results).Error
	
	return
}

// GetBatchFromLevelDepth 批量唯一主键查找 
func (obj *_EgCmsCategoryMgr) GetBatchFromLevelDepth(levelDepths []uint8) (results []*EgCmsCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("level_depth IN (?)", levelDepths).Find(&results).Error
	
	return
}
 
// GetFromActive 通过active获取内容  
func (obj *_EgCmsCategoryMgr) GetFromActive(active bool) (results []*EgCmsCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active = ?", active).Find(&results).Error
	
	return
}

// GetBatchFromActive 批量唯一主键查找 
func (obj *_EgCmsCategoryMgr) GetBatchFromActive(actives []bool) (results []*EgCmsCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active IN (?)", actives).Find(&results).Error
	
	return
}
 
// GetFromDateAdd 通过date_add获取内容  
func (obj *_EgCmsCategoryMgr) GetFromDateAdd(dateAdd time.Time) (results []*EgCmsCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error
	
	return
}

// GetBatchFromDateAdd 批量唯一主键查找 
func (obj *_EgCmsCategoryMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*EgCmsCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error
	
	return
}
 
// GetFromDateUpd 通过date_upd获取内容  
func (obj *_EgCmsCategoryMgr) GetFromDateUpd(dateUpd time.Time) (results []*EgCmsCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd = ?", dateUpd).Find(&results).Error
	
	return
}

// GetBatchFromDateUpd 批量唯一主键查找 
func (obj *_EgCmsCategoryMgr) GetBatchFromDateUpd(dateUpds []time.Time) (results []*EgCmsCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd IN (?)", dateUpds).Find(&results).Error
	
	return
}
 
// GetFromPosition 通过position获取内容  
func (obj *_EgCmsCategoryMgr) GetFromPosition(position uint32) (results []*EgCmsCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("position = ?", position).Find(&results).Error
	
	return
}

// GetBatchFromPosition 批量唯一主键查找 
func (obj *_EgCmsCategoryMgr) GetBatchFromPosition(positions []uint32) (results []*EgCmsCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("position IN (?)", positions).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgCmsCategoryMgr) FetchByPrimaryKey(idCmsCategory uint32 ) (result EgCmsCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cms_category = ?", idCmsCategory).Find(&result).Error
	
	return
}
 

 
 // FetchIndexByCategoryParent  获取多个内容
 func (obj *_EgCmsCategoryMgr) FetchIndexByCategoryParent(idParent uint32 ) (results []*EgCmsCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_parent = ?", idParent).Find(&results).Error
	
	return
}
 

	

