package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgProductCommentCriterionCategoryMgr struct {
	*_BaseMgr
}

// EgProductCommentCriterionCategoryMgr open func
func EgProductCommentCriterionCategoryMgr(db *gorm.DB) *_EgProductCommentCriterionCategoryMgr {
	if db == nil {
		panic(fmt.Errorf("EgProductCommentCriterionCategoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgProductCommentCriterionCategoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_product_comment_criterion_category"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgProductCommentCriterionCategoryMgr) GetTableName() string {
	return "eg_product_comment_criterion_category"
}

// Get 获取
func (obj *_EgProductCommentCriterionCategoryMgr) Get() (result EgProductCommentCriterionCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgProductCommentCriterionCategoryMgr) Gets() (results []*EgProductCommentCriterionCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDProductCommentCriterion id_product_comment_criterion获取 
func (obj *_EgProductCommentCriterionCategoryMgr) WithIDProductCommentCriterion(idProductCommentCriterion uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product_comment_criterion"] = idProductCommentCriterion })
}

// WithIDCategory id_category获取 
func (obj *_EgProductCommentCriterionCategoryMgr) WithIDCategory(idCategory uint32) Option {
	return optionFunc(func(o *options) { o.query["id_category"] = idCategory })
}


// GetByOption 功能选项模式获取
func (obj *_EgProductCommentCriterionCategoryMgr) GetByOption(opts ...Option) (result EgProductCommentCriterionCategory, err error) {
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
func (obj *_EgProductCommentCriterionCategoryMgr) GetByOptions(opts ...Option) (results []*EgProductCommentCriterionCategory, err error) {
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
func (obj *_EgProductCommentCriterionCategoryMgr) GetFromIDProductCommentCriterion(idProductCommentCriterion uint32) (results []*EgProductCommentCriterionCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_comment_criterion = ?", idProductCommentCriterion).Find(&results).Error
	
	return
}

// GetBatchFromIDProductCommentCriterion 批量唯一主键查找 
func (obj *_EgProductCommentCriterionCategoryMgr) GetBatchFromIDProductCommentCriterion(idProductCommentCriterions []uint32) (results []*EgProductCommentCriterionCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_comment_criterion IN (?)", idProductCommentCriterions).Find(&results).Error
	
	return
}
 
// GetFromIDCategory 通过id_category获取内容  
func (obj *_EgProductCommentCriterionCategoryMgr) GetFromIDCategory(idCategory uint32) (results []*EgProductCommentCriterionCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_category = ?", idCategory).Find(&results).Error
	
	return
}

// GetBatchFromIDCategory 批量唯一主键查找 
func (obj *_EgProductCommentCriterionCategoryMgr) GetBatchFromIDCategory(idCategorys []uint32) (results []*EgProductCommentCriterionCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_category IN (?)", idCategorys).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgProductCommentCriterionCategoryMgr) FetchByPrimaryKey(idProductCommentCriterion uint32 ,idCategory uint32 ) (result EgProductCommentCriterionCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_comment_criterion = ? AND id_category = ?", idProductCommentCriterion , idCategory).Find(&result).Error
	
	return
}
 

 
 // FetchIndexByIDCategory  获取多个内容
 func (obj *_EgProductCommentCriterionCategoryMgr) FetchIndexByIDCategory(idCategory uint32 ) (results []*EgProductCommentCriterionCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_category = ?", idCategory).Find(&results).Error
	
	return
}
 

	

