package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgCmsMgr struct {
	*_BaseMgr
}

// EgCmsMgr open func
func EgCmsMgr(db *gorm.DB) *_EgCmsMgr {
	if db == nil {
		panic(fmt.Errorf("EgCmsMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgCmsMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_cms"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgCmsMgr) GetTableName() string {
	return "eg_cms"
}

// Get 获取
func (obj *_EgCmsMgr) Get() (result EgCms, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgCmsMgr) Gets() (results []*EgCms, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDCms id_cms获取 
func (obj *_EgCmsMgr) WithIDCms(idCms uint32) Option {
	return optionFunc(func(o *options) { o.query["id_cms"] = idCms })
}

// WithIDCmsCategory id_cms_category获取 
func (obj *_EgCmsMgr) WithIDCmsCategory(idCmsCategory uint32) Option {
	return optionFunc(func(o *options) { o.query["id_cms_category"] = idCmsCategory })
}

// WithPosition position获取 
func (obj *_EgCmsMgr) WithPosition(position uint32) Option {
	return optionFunc(func(o *options) { o.query["position"] = position })
}

// WithActive active获取 
func (obj *_EgCmsMgr) WithActive(active bool) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}

// WithIndexation indexation获取 
func (obj *_EgCmsMgr) WithIndexation(indexation bool) Option {
	return optionFunc(func(o *options) { o.query["indexation"] = indexation })
}


// GetByOption 功能选项模式获取
func (obj *_EgCmsMgr) GetByOption(opts ...Option) (result EgCms, err error) {
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
func (obj *_EgCmsMgr) GetByOptions(opts ...Option) (results []*EgCms, err error) {
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


// GetFromIDCms 通过id_cms获取内容  
func (obj *_EgCmsMgr)  GetFromIDCms(idCms uint32) (result EgCms, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cms = ?", idCms).Find(&result).Error
	
	return
}

// GetBatchFromIDCms 批量唯一主键查找 
func (obj *_EgCmsMgr) GetBatchFromIDCms(idCmss []uint32) (results []*EgCms, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cms IN (?)", idCmss).Find(&results).Error
	
	return
}
 
// GetFromIDCmsCategory 通过id_cms_category获取内容  
func (obj *_EgCmsMgr) GetFromIDCmsCategory(idCmsCategory uint32) (results []*EgCms, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cms_category = ?", idCmsCategory).Find(&results).Error
	
	return
}

// GetBatchFromIDCmsCategory 批量唯一主键查找 
func (obj *_EgCmsMgr) GetBatchFromIDCmsCategory(idCmsCategorys []uint32) (results []*EgCms, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cms_category IN (?)", idCmsCategorys).Find(&results).Error
	
	return
}
 
// GetFromPosition 通过position获取内容  
func (obj *_EgCmsMgr) GetFromPosition(position uint32) (results []*EgCms, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("position = ?", position).Find(&results).Error
	
	return
}

// GetBatchFromPosition 批量唯一主键查找 
func (obj *_EgCmsMgr) GetBatchFromPosition(positions []uint32) (results []*EgCms, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("position IN (?)", positions).Find(&results).Error
	
	return
}
 
// GetFromActive 通过active获取内容  
func (obj *_EgCmsMgr) GetFromActive(active bool) (results []*EgCms, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active = ?", active).Find(&results).Error
	
	return
}

// GetBatchFromActive 批量唯一主键查找 
func (obj *_EgCmsMgr) GetBatchFromActive(actives []bool) (results []*EgCms, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active IN (?)", actives).Find(&results).Error
	
	return
}
 
// GetFromIndexation 通过indexation获取内容  
func (obj *_EgCmsMgr) GetFromIndexation(indexation bool) (results []*EgCms, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("indexation = ?", indexation).Find(&results).Error
	
	return
}

// GetBatchFromIndexation 批量唯一主键查找 
func (obj *_EgCmsMgr) GetBatchFromIndexation(indexations []bool) (results []*EgCms, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("indexation IN (?)", indexations).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgCmsMgr) FetchByPrimaryKey(idCms uint32 ) (result EgCms, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cms = ?", idCms).Find(&result).Error
	
	return
}
 

 

	

