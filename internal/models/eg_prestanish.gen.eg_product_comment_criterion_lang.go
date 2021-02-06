package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgProductCommentCriterionLangMgr struct {
	*_BaseMgr
}

// EgProductCommentCriterionLangMgr open func
func EgProductCommentCriterionLangMgr(db *gorm.DB) *_EgProductCommentCriterionLangMgr {
	if db == nil {
		panic(fmt.Errorf("EgProductCommentCriterionLangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgProductCommentCriterionLangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_product_comment_criterion_lang"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgProductCommentCriterionLangMgr) GetTableName() string {
	return "eg_product_comment_criterion_lang"
}

// Get 获取
func (obj *_EgProductCommentCriterionLangMgr) Get() (result EgProductCommentCriterionLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgProductCommentCriterionLangMgr) Gets() (results []*EgProductCommentCriterionLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDProductCommentCriterion id_product_comment_criterion获取 
func (obj *_EgProductCommentCriterionLangMgr) WithIDProductCommentCriterion(idProductCommentCriterion uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product_comment_criterion"] = idProductCommentCriterion })
}

// WithIDLang id_lang获取 
func (obj *_EgProductCommentCriterionLangMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

// WithName name获取 
func (obj *_EgProductCommentCriterionLangMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}


// GetByOption 功能选项模式获取
func (obj *_EgProductCommentCriterionLangMgr) GetByOption(opts ...Option) (result EgProductCommentCriterionLang, err error) {
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
func (obj *_EgProductCommentCriterionLangMgr) GetByOptions(opts ...Option) (results []*EgProductCommentCriterionLang, err error) {
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
func (obj *_EgProductCommentCriterionLangMgr) GetFromIDProductCommentCriterion(idProductCommentCriterion uint32) (results []*EgProductCommentCriterionLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_comment_criterion = ?", idProductCommentCriterion).Find(&results).Error
	
	return
}

// GetBatchFromIDProductCommentCriterion 批量唯一主键查找 
func (obj *_EgProductCommentCriterionLangMgr) GetBatchFromIDProductCommentCriterion(idProductCommentCriterions []uint32) (results []*EgProductCommentCriterionLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_comment_criterion IN (?)", idProductCommentCriterions).Find(&results).Error
	
	return
}
 
// GetFromIDLang 通过id_lang获取内容  
func (obj *_EgProductCommentCriterionLangMgr) GetFromIDLang(idLang uint32) (results []*EgProductCommentCriterionLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error
	
	return
}

// GetBatchFromIDLang 批量唯一主键查找 
func (obj *_EgProductCommentCriterionLangMgr) GetBatchFromIDLang(idLangs []uint32) (results []*EgProductCommentCriterionLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error
	
	return
}
 
// GetFromName 通过name获取内容  
func (obj *_EgProductCommentCriterionLangMgr) GetFromName(name string) (results []*EgProductCommentCriterionLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error
	
	return
}

// GetBatchFromName 批量唯一主键查找 
func (obj *_EgProductCommentCriterionLangMgr) GetBatchFromName(names []string) (results []*EgProductCommentCriterionLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgProductCommentCriterionLangMgr) FetchByPrimaryKey(idProductCommentCriterion uint32 ,idLang uint32 ) (result EgProductCommentCriterionLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_comment_criterion = ? AND id_lang = ?", idProductCommentCriterion , idLang).Find(&result).Error
	
	return
}
 

 

	

