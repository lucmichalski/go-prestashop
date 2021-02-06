package	model	
import (	
"context"	
"gorm.io/gorm"	
"fmt"	
)	

type _EgCategoryProductMgr struct {
	*_BaseMgr
}

// EgCategoryProductMgr open func
func EgCategoryProductMgr(db *gorm.DB) *_EgCategoryProductMgr {
	if db == nil {
		panic(fmt.Errorf("EgCategoryProductMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgCategoryProductMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_category_product"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgCategoryProductMgr) GetTableName() string {
	return "eg_category_product"
}

// Get 获取
func (obj *_EgCategoryProductMgr) Get() (result EgCategoryProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgCategoryProductMgr) Gets() (results []*EgCategoryProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDCategory id_category获取 
func (obj *_EgCategoryProductMgr) WithIDCategory(idCategory uint32) Option {
	return optionFunc(func(o *options) { o.query["id_category"] = idCategory })
}

// WithIDProduct id_product获取 
func (obj *_EgCategoryProductMgr) WithIDProduct(idProduct uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product"] = idProduct })
}

// WithPosition position获取 
func (obj *_EgCategoryProductMgr) WithPosition(position uint32) Option {
	return optionFunc(func(o *options) { o.query["position"] = position })
}


// GetByOption 功能选项模式获取
func (obj *_EgCategoryProductMgr) GetByOption(opts ...Option) (result EgCategoryProduct, err error) {
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
func (obj *_EgCategoryProductMgr) GetByOptions(opts ...Option) (results []*EgCategoryProduct, err error) {
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


// GetFromIDCategory 通过id_category获取内容  
func (obj *_EgCategoryProductMgr) GetFromIDCategory(idCategory uint32) (results []*EgCategoryProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_category = ?", idCategory).Find(&results).Error
	
	return
}

// GetBatchFromIDCategory 批量唯一主键查找 
func (obj *_EgCategoryProductMgr) GetBatchFromIDCategory(idCategorys []uint32) (results []*EgCategoryProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_category IN (?)", idCategorys).Find(&results).Error
	
	return
}
 
// GetFromIDProduct 通过id_product获取内容  
func (obj *_EgCategoryProductMgr) GetFromIDProduct(idProduct uint32) (results []*EgCategoryProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ?", idProduct).Find(&results).Error
	
	return
}

// GetBatchFromIDProduct 批量唯一主键查找 
func (obj *_EgCategoryProductMgr) GetBatchFromIDProduct(idProducts []uint32) (results []*EgCategoryProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product IN (?)", idProducts).Find(&results).Error
	
	return
}
 
// GetFromPosition 通过position获取内容  
func (obj *_EgCategoryProductMgr) GetFromPosition(position uint32) (results []*EgCategoryProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("position = ?", position).Find(&results).Error
	
	return
}

// GetBatchFromPosition 批量唯一主键查找 
func (obj *_EgCategoryProductMgr) GetBatchFromPosition(positions []uint32) (results []*EgCategoryProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("position IN (?)", positions).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgCategoryProductMgr) FetchByPrimaryKey(idCategory uint32 ,idProduct uint32 ) (result EgCategoryProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_category = ? AND id_product = ?", idCategory , idProduct).Find(&result).Error
	
	return
}
 

 
 // FetchIndexByIDCategory  获取多个内容
 func (obj *_EgCategoryProductMgr) FetchIndexByIDCategory(idCategory uint32 ,position uint32 ) (results []*EgCategoryProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_category = ? AND position = ?", idCategory , position).Find(&results).Error
	
	return
}
 
 // FetchIndexByIDProduct  获取多个内容
 func (obj *_EgCategoryProductMgr) FetchIndexByIDProduct(idProduct uint32 ) (results []*EgCategoryProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ?", idProduct).Find(&results).Error
	
	return
}
 

	

