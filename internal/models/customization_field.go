package	model	
import (	
"context"	
"gorm.io/gorm"	
"fmt"	
)	

type _EgCustomizationFieldMgr struct {
	*_BaseMgr
}

// EgCustomizationFieldMgr open func
func EgCustomizationFieldMgr(db *gorm.DB) *_EgCustomizationFieldMgr {
	if db == nil {
		panic(fmt.Errorf("EgCustomizationFieldMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgCustomizationFieldMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_customization_field"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgCustomizationFieldMgr) GetTableName() string {
	return "eg_customization_field"
}

// Get 获取
func (obj *_EgCustomizationFieldMgr) Get() (result EgCustomizationField, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgCustomizationFieldMgr) Gets() (results []*EgCustomizationField, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDCustomizationField id_customization_field获取 
func (obj *_EgCustomizationFieldMgr) WithIDCustomizationField(idCustomizationField uint32) Option {
	return optionFunc(func(o *options) { o.query["id_customization_field"] = idCustomizationField })
}

// WithIDProduct id_product获取 
func (obj *_EgCustomizationFieldMgr) WithIDProduct(idProduct uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product"] = idProduct })
}

// WithType type获取 
func (obj *_EgCustomizationFieldMgr) WithType(_type bool) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithRequired required获取 
func (obj *_EgCustomizationFieldMgr) WithRequired(required bool) Option {
	return optionFunc(func(o *options) { o.query["required"] = required })
}

// WithIsModule is_module获取 
func (obj *_EgCustomizationFieldMgr) WithIsModule(isModule bool) Option {
	return optionFunc(func(o *options) { o.query["is_module"] = isModule })
}

// WithIsDeleted is_deleted获取 
func (obj *_EgCustomizationFieldMgr) WithIsDeleted(isDeleted bool) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = isDeleted })
}


// GetByOption 功能选项模式获取
func (obj *_EgCustomizationFieldMgr) GetByOption(opts ...Option) (result EgCustomizationField, err error) {
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
func (obj *_EgCustomizationFieldMgr) GetByOptions(opts ...Option) (results []*EgCustomizationField, err error) {
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


// GetFromIDCustomizationField 通过id_customization_field获取内容  
func (obj *_EgCustomizationFieldMgr)  GetFromIDCustomizationField(idCustomizationField uint32) (result EgCustomizationField, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customization_field = ?", idCustomizationField).Find(&result).Error
	
	return
}

// GetBatchFromIDCustomizationField 批量唯一主键查找 
func (obj *_EgCustomizationFieldMgr) GetBatchFromIDCustomizationField(idCustomizationFields []uint32) (results []*EgCustomizationField, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customization_field IN (?)", idCustomizationFields).Find(&results).Error
	
	return
}
 
// GetFromIDProduct 通过id_product获取内容  
func (obj *_EgCustomizationFieldMgr) GetFromIDProduct(idProduct uint32) (results []*EgCustomizationField, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ?", idProduct).Find(&results).Error
	
	return
}

// GetBatchFromIDProduct 批量唯一主键查找 
func (obj *_EgCustomizationFieldMgr) GetBatchFromIDProduct(idProducts []uint32) (results []*EgCustomizationField, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product IN (?)", idProducts).Find(&results).Error
	
	return
}
 
// GetFromType 通过type获取内容  
func (obj *_EgCustomizationFieldMgr) GetFromType(_type bool) (results []*EgCustomizationField, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type = ?", _type).Find(&results).Error
	
	return
}

// GetBatchFromType 批量唯一主键查找 
func (obj *_EgCustomizationFieldMgr) GetBatchFromType(_types []bool) (results []*EgCustomizationField, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type IN (?)", _types).Find(&results).Error
	
	return
}
 
// GetFromRequired 通过required获取内容  
func (obj *_EgCustomizationFieldMgr) GetFromRequired(required bool) (results []*EgCustomizationField, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("required = ?", required).Find(&results).Error
	
	return
}

// GetBatchFromRequired 批量唯一主键查找 
func (obj *_EgCustomizationFieldMgr) GetBatchFromRequired(requireds []bool) (results []*EgCustomizationField, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("required IN (?)", requireds).Find(&results).Error
	
	return
}
 
// GetFromIsModule 通过is_module获取内容  
func (obj *_EgCustomizationFieldMgr) GetFromIsModule(isModule bool) (results []*EgCustomizationField, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_module = ?", isModule).Find(&results).Error
	
	return
}

// GetBatchFromIsModule 批量唯一主键查找 
func (obj *_EgCustomizationFieldMgr) GetBatchFromIsModule(isModules []bool) (results []*EgCustomizationField, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_module IN (?)", isModules).Find(&results).Error
	
	return
}
 
// GetFromIsDeleted 通过is_deleted获取内容  
func (obj *_EgCustomizationFieldMgr) GetFromIsDeleted(isDeleted bool) (results []*EgCustomizationField, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_deleted = ?", isDeleted).Find(&results).Error
	
	return
}

// GetBatchFromIsDeleted 批量唯一主键查找 
func (obj *_EgCustomizationFieldMgr) GetBatchFromIsDeleted(isDeleteds []bool) (results []*EgCustomizationField, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_deleted IN (?)", isDeleteds).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgCustomizationFieldMgr) FetchByPrimaryKey(idCustomizationField uint32 ) (result EgCustomizationField, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customization_field = ?", idCustomizationField).Find(&result).Error
	
	return
}
 

 
 // FetchIndexByIDProduct  获取多个内容
 func (obj *_EgCustomizationFieldMgr) FetchIndexByIDProduct(idProduct uint32 ) (results []*EgCustomizationField, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ?", idProduct).Find(&results).Error
	
	return
}
 

	

