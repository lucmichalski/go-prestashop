package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgProductCommentGradeMgr struct {
	*_BaseMgr
}

// EgProductCommentGradeMgr open func
func EgProductCommentGradeMgr(db *gorm.DB) *_EgProductCommentGradeMgr {
	if db == nil {
		panic(fmt.Errorf("EgProductCommentGradeMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgProductCommentGradeMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_product_comment_grade"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgProductCommentGradeMgr) GetTableName() string {
	return "eg_product_comment_grade"
}

// Get 获取
func (obj *_EgProductCommentGradeMgr) Get() (result EgProductCommentGrade, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgProductCommentGradeMgr) Gets() (results []*EgProductCommentGrade, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDProductComment id_product_comment获取 
func (obj *_EgProductCommentGradeMgr) WithIDProductComment(idProductComment int) Option {
	return optionFunc(func(o *options) { o.query["id_product_comment"] = idProductComment })
}

// WithIDProductCommentCriterion id_product_comment_criterion获取 
func (obj *_EgProductCommentGradeMgr) WithIDProductCommentCriterion(idProductCommentCriterion int) Option {
	return optionFunc(func(o *options) { o.query["id_product_comment_criterion"] = idProductCommentCriterion })
}

// WithGrade grade获取 
func (obj *_EgProductCommentGradeMgr) WithGrade(grade int) Option {
	return optionFunc(func(o *options) { o.query["grade"] = grade })
}


// GetByOption 功能选项模式获取
func (obj *_EgProductCommentGradeMgr) GetByOption(opts ...Option) (result EgProductCommentGrade, err error) {
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
func (obj *_EgProductCommentGradeMgr) GetByOptions(opts ...Option) (results []*EgProductCommentGrade, err error) {
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


// GetFromIDProductComment 通过id_product_comment获取内容  
func (obj *_EgProductCommentGradeMgr) GetFromIDProductComment(idProductComment int) (results []*EgProductCommentGrade, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_comment = ?", idProductComment).Find(&results).Error
	
	return
}

// GetBatchFromIDProductComment 批量唯一主键查找 
func (obj *_EgProductCommentGradeMgr) GetBatchFromIDProductComment(idProductComments []int) (results []*EgProductCommentGrade, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_comment IN (?)", idProductComments).Find(&results).Error
	
	return
}
 
// GetFromIDProductCommentCriterion 通过id_product_comment_criterion获取内容  
func (obj *_EgProductCommentGradeMgr) GetFromIDProductCommentCriterion(idProductCommentCriterion int) (results []*EgProductCommentGrade, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_comment_criterion = ?", idProductCommentCriterion).Find(&results).Error
	
	return
}

// GetBatchFromIDProductCommentCriterion 批量唯一主键查找 
func (obj *_EgProductCommentGradeMgr) GetBatchFromIDProductCommentCriterion(idProductCommentCriterions []int) (results []*EgProductCommentGrade, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_comment_criterion IN (?)", idProductCommentCriterions).Find(&results).Error
	
	return
}
 
// GetFromGrade 通过grade获取内容  
func (obj *_EgProductCommentGradeMgr) GetFromGrade(grade int) (results []*EgProductCommentGrade, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("grade = ?", grade).Find(&results).Error
	
	return
}

// GetBatchFromGrade 批量唯一主键查找 
func (obj *_EgProductCommentGradeMgr) GetBatchFromGrade(grades []int) (results []*EgProductCommentGrade, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("grade IN (?)", grades).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgProductCommentGradeMgr) FetchByPrimaryKey(idProductComment int ,idProductCommentCriterion int ) (result EgProductCommentGrade, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_comment = ? AND id_product_comment_criterion = ?", idProductComment , idProductCommentCriterion).Find(&result).Error
	
	return
}
 

 
 // FetchIndexByIDX367426EBACF38A54  获取多个内容
 func (obj *_EgProductCommentGradeMgr) FetchIndexByIDX367426EBACF38A54(idProductComment int ) (results []*EgProductCommentGrade, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_comment = ?", idProductComment).Find(&results).Error
	
	return
}
 
 // FetchIndexByIDX367426EB8375853C  获取多个内容
 func (obj *_EgProductCommentGradeMgr) FetchIndexByIDX367426EB8375853C(idProductCommentCriterion int ) (results []*EgProductCommentGrade, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_comment_criterion = ?", idProductCommentCriterion).Find(&results).Error
	
	return
}
 

	

