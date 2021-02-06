package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgProductTagMgr struct {
	*_BaseMgr
}

// EgProductTagMgr open func
func EgProductTagMgr(db *gorm.DB) *_EgProductTagMgr {
	if db == nil {
		panic(fmt.Errorf("EgProductTagMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgProductTagMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_product_tag"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgProductTagMgr) GetTableName() string {
	return "eg_product_tag"
}

// Get 获取
func (obj *_EgProductTagMgr) Get() (result EgProductTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgProductTagMgr) Gets() (results []*EgProductTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDProduct id_product获取 
func (obj *_EgProductTagMgr) WithIDProduct(idProduct uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product"] = idProduct })
}

// WithIDTag id_tag获取 
func (obj *_EgProductTagMgr) WithIDTag(idTag uint32) Option {
	return optionFunc(func(o *options) { o.query["id_tag"] = idTag })
}

// WithIDLang id_lang获取 
func (obj *_EgProductTagMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}


// GetByOption 功能选项模式获取
func (obj *_EgProductTagMgr) GetByOption(opts ...Option) (result EgProductTag, err error) {
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
func (obj *_EgProductTagMgr) GetByOptions(opts ...Option) (results []*EgProductTag, err error) {
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


// GetFromIDProduct 通过id_product获取内容  
func (obj *_EgProductTagMgr) GetFromIDProduct(idProduct uint32) (results []*EgProductTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ?", idProduct).Find(&results).Error
	
	return
}

// GetBatchFromIDProduct 批量唯一主键查找 
func (obj *_EgProductTagMgr) GetBatchFromIDProduct(idProducts []uint32) (results []*EgProductTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product IN (?)", idProducts).Find(&results).Error
	
	return
}
 
// GetFromIDTag 通过id_tag获取内容  
func (obj *_EgProductTagMgr) GetFromIDTag(idTag uint32) (results []*EgProductTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tag = ?", idTag).Find(&results).Error
	
	return
}

// GetBatchFromIDTag 批量唯一主键查找 
func (obj *_EgProductTagMgr) GetBatchFromIDTag(idTags []uint32) (results []*EgProductTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tag IN (?)", idTags).Find(&results).Error
	
	return
}
 
// GetFromIDLang 通过id_lang获取内容  
func (obj *_EgProductTagMgr) GetFromIDLang(idLang uint32) (results []*EgProductTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error
	
	return
}

// GetBatchFromIDLang 批量唯一主键查找 
func (obj *_EgProductTagMgr) GetBatchFromIDLang(idLangs []uint32) (results []*EgProductTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgProductTagMgr) FetchByPrimaryKey(idProduct uint32 ,idTag uint32 ) (result EgProductTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ? AND id_tag = ?", idProduct , idTag).Find(&result).Error
	
	return
}
 

 
 // FetchIndexByIDTag  获取多个内容
 func (obj *_EgProductTagMgr) FetchIndexByIDTag(idTag uint32 ) (results []*EgProductTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tag = ?", idTag).Find(&results).Error
	
	return
}
 
 // FetchIndexByIDLang  获取多个内容
 func (obj *_EgProductTagMgr) FetchIndexByIDLang(idTag uint32 ,idLang uint32 ) (results []*EgProductTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tag = ? AND id_lang = ?", idTag , idLang).Find(&results).Error
	
	return
}
 

	

