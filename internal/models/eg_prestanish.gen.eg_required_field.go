package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgRequiredFieldMgr struct {
	*_BaseMgr
}

// EgRequiredFieldMgr open func
func EgRequiredFieldMgr(db *gorm.DB) *_EgRequiredFieldMgr {
	if db == nil {
		panic(fmt.Errorf("EgRequiredFieldMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgRequiredFieldMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_required_field"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgRequiredFieldMgr) GetTableName() string {
	return "eg_required_field"
}

// Get 获取
func (obj *_EgRequiredFieldMgr) Get() (result EgRequiredField, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgRequiredFieldMgr) Gets() (results []*EgRequiredField, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDRequiredField id_required_field获取 
func (obj *_EgRequiredFieldMgr) WithIDRequiredField(idRequiredField int) Option {
	return optionFunc(func(o *options) { o.query["id_required_field"] = idRequiredField })
}

// WithObjectName object_name获取 
func (obj *_EgRequiredFieldMgr) WithObjectName(objectName string) Option {
	return optionFunc(func(o *options) { o.query["object_name"] = objectName })
}

// WithFieldName field_name获取 
func (obj *_EgRequiredFieldMgr) WithFieldName(fieldName string) Option {
	return optionFunc(func(o *options) { o.query["field_name"] = fieldName })
}


// GetByOption 功能选项模式获取
func (obj *_EgRequiredFieldMgr) GetByOption(opts ...Option) (result EgRequiredField, err error) {
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
func (obj *_EgRequiredFieldMgr) GetByOptions(opts ...Option) (results []*EgRequiredField, err error) {
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


// GetFromIDRequiredField 通过id_required_field获取内容  
func (obj *_EgRequiredFieldMgr)  GetFromIDRequiredField(idRequiredField int) (result EgRequiredField, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_required_field = ?", idRequiredField).Find(&result).Error
	
	return
}

// GetBatchFromIDRequiredField 批量唯一主键查找 
func (obj *_EgRequiredFieldMgr) GetBatchFromIDRequiredField(idRequiredFields []int) (results []*EgRequiredField, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_required_field IN (?)", idRequiredFields).Find(&results).Error
	
	return
}
 
// GetFromObjectName 通过object_name获取内容  
func (obj *_EgRequiredFieldMgr) GetFromObjectName(objectName string) (results []*EgRequiredField, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("object_name = ?", objectName).Find(&results).Error
	
	return
}

// GetBatchFromObjectName 批量唯一主键查找 
func (obj *_EgRequiredFieldMgr) GetBatchFromObjectName(objectNames []string) (results []*EgRequiredField, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("object_name IN (?)", objectNames).Find(&results).Error
	
	return
}
 
// GetFromFieldName 通过field_name获取内容  
func (obj *_EgRequiredFieldMgr) GetFromFieldName(fieldName string) (results []*EgRequiredField, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("field_name = ?", fieldName).Find(&results).Error
	
	return
}

// GetBatchFromFieldName 批量唯一主键查找 
func (obj *_EgRequiredFieldMgr) GetBatchFromFieldName(fieldNames []string) (results []*EgRequiredField, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("field_name IN (?)", fieldNames).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgRequiredFieldMgr) FetchByPrimaryKey(idRequiredField int ) (result EgRequiredField, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_required_field = ?", idRequiredField).Find(&result).Error
	
	return
}
 

 
 // FetchIndexByObjectName  获取多个内容
 func (obj *_EgRequiredFieldMgr) FetchIndexByObjectName(objectName string ) (results []*EgRequiredField, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("object_name = ?", objectName).Find(&results).Error
	
	return
}
 

	

