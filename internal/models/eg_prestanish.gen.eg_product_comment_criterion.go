package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgProductCommentCriterionMgr struct {
	*_BaseMgr
}

// EgProductCommentCriterionMgr open func
func EgProductCommentCriterionMgr(db *gorm.DB) *_EgProductCommentCriterionMgr {
	if db == nil {
		panic(fmt.Errorf("EgProductCommentCriterionMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgProductCommentCriterionMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_product_comment_criterion"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgProductCommentCriterionMgr) GetTableName() string {
	return "eg_product_comment_criterion"
}

// Get 获取
func (obj *_EgProductCommentCriterionMgr) Get() (result EgProductCommentCriterion, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgProductCommentCriterionMgr) Gets() (results []*EgProductCommentCriterion, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDProductCommentCriterion id_product_comment_criterion获取 
func (obj *_EgProductCommentCriterionMgr) WithIDProductCommentCriterion(idProductCommentCriterion int) Option {
	return optionFunc(func(o *options) { o.query["id_product_comment_criterion"] = idProductCommentCriterion })
}

// WithIDProductCommentCriterionType id_product_comment_criterion_type获取 
func (obj *_EgProductCommentCriterionMgr) WithIDProductCommentCriterionType(idProductCommentCriterionType int) Option {
	return optionFunc(func(o *options) { o.query["id_product_comment_criterion_type"] = idProductCommentCriterionType })
}

// WithActive active获取 
func (obj *_EgProductCommentCriterionMgr) WithActive(active bool) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}


// GetByOption 功能选项模式获取
func (obj *_EgProductCommentCriterionMgr) GetByOption(opts ...Option) (result EgProductCommentCriterion, err error) {
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
func (obj *_EgProductCommentCriterionMgr) GetByOptions(opts ...Option) (results []*EgProductCommentCriterion, err error) {
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


// GetFromIDProductCommentCriterion 通过id_product_comment_criterion获取内容  
func (obj *_EgProductCommentCriterionMgr)  GetFromIDProductCommentCriterion(idProductCommentCriterion int) (result EgProductCommentCriterion, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_comment_criterion = ?", idProductCommentCriterion).Find(&result).Error
	
	return
}

// GetBatchFromIDProductCommentCriterion 批量唯一主键查找 
func (obj *_EgProductCommentCriterionMgr) GetBatchFromIDProductCommentCriterion(idProductCommentCriterions []int) (results []*EgProductCommentCriterion, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_comment_criterion IN (?)", idProductCommentCriterions).Find(&results).Error
	
	return
}
 
// GetFromIDProductCommentCriterionType 通过id_product_comment_criterion_type获取内容  
func (obj *_EgProductCommentCriterionMgr) GetFromIDProductCommentCriterionType(idProductCommentCriterionType int) (results []*EgProductCommentCriterion, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_comment_criterion_type = ?", idProductCommentCriterionType).Find(&results).Error
	
	return
}

// GetBatchFromIDProductCommentCriterionType 批量唯一主键查找 
func (obj *_EgProductCommentCriterionMgr) GetBatchFromIDProductCommentCriterionType(idProductCommentCriterionTypes []int) (results []*EgProductCommentCriterion, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_comment_criterion_type IN (?)", idProductCommentCriterionTypes).Find(&results).Error
	
	return
}
 
// GetFromActive 通过active获取内容  
func (obj *_EgProductCommentCriterionMgr) GetFromActive(active bool) (results []*EgProductCommentCriterion, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active = ?", active).Find(&results).Error
	
	return
}

// GetBatchFromActive 批量唯一主键查找 
func (obj *_EgProductCommentCriterionMgr) GetBatchFromActive(actives []bool) (results []*EgProductCommentCriterion, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active IN (?)", actives).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgProductCommentCriterionMgr) FetchByPrimaryKey(idProductCommentCriterion int ) (result EgProductCommentCriterion, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_comment_criterion = ?", idProductCommentCriterion).Find(&result).Error
	
	return
}
 

 

	

