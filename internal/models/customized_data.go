package	model	
import (	
"context"	
"gorm.io/gorm"	
"fmt"	
)	

type _EgCustomizedDataMgr struct {
	*_BaseMgr
}

// EgCustomizedDataMgr open func
func EgCustomizedDataMgr(db *gorm.DB) *_EgCustomizedDataMgr {
	if db == nil {
		panic(fmt.Errorf("EgCustomizedDataMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgCustomizedDataMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_customized_data"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgCustomizedDataMgr) GetTableName() string {
	return "eg_customized_data"
}

// Get 获取
func (obj *_EgCustomizedDataMgr) Get() (result EgCustomizedData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgCustomizedDataMgr) Gets() (results []*EgCustomizedData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDCustomization id_customization获取 
func (obj *_EgCustomizedDataMgr) WithIDCustomization(idCustomization uint32) Option {
	return optionFunc(func(o *options) { o.query["id_customization"] = idCustomization })
}

// WithType type获取 
func (obj *_EgCustomizedDataMgr) WithType(_type bool) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithIndex index获取 
func (obj *_EgCustomizedDataMgr) WithIndex(index int) Option {
	return optionFunc(func(o *options) { o.query["index"] = index })
}

// WithValue value获取 
func (obj *_EgCustomizedDataMgr) WithValue(value string) Option {
	return optionFunc(func(o *options) { o.query["value"] = value })
}

// WithIDModule id_module获取 
func (obj *_EgCustomizedDataMgr) WithIDModule(idModule int) Option {
	return optionFunc(func(o *options) { o.query["id_module"] = idModule })
}

// WithPrice price获取 
func (obj *_EgCustomizedDataMgr) WithPrice(price float64) Option {
	return optionFunc(func(o *options) { o.query["price"] = price })
}

// WithWeight weight获取 
func (obj *_EgCustomizedDataMgr) WithWeight(weight float64) Option {
	return optionFunc(func(o *options) { o.query["weight"] = weight })
}


// GetByOption 功能选项模式获取
func (obj *_EgCustomizedDataMgr) GetByOption(opts ...Option) (result EgCustomizedData, err error) {
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
func (obj *_EgCustomizedDataMgr) GetByOptions(opts ...Option) (results []*EgCustomizedData, err error) {
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


// GetFromIDCustomization 通过id_customization获取内容  
func (obj *_EgCustomizedDataMgr) GetFromIDCustomization(idCustomization uint32) (results []*EgCustomizedData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customization = ?", idCustomization).Find(&results).Error
	
	return
}

// GetBatchFromIDCustomization 批量唯一主键查找 
func (obj *_EgCustomizedDataMgr) GetBatchFromIDCustomization(idCustomizations []uint32) (results []*EgCustomizedData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customization IN (?)", idCustomizations).Find(&results).Error
	
	return
}
 
// GetFromType 通过type获取内容  
func (obj *_EgCustomizedDataMgr) GetFromType(_type bool) (results []*EgCustomizedData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type = ?", _type).Find(&results).Error
	
	return
}

// GetBatchFromType 批量唯一主键查找 
func (obj *_EgCustomizedDataMgr) GetBatchFromType(_types []bool) (results []*EgCustomizedData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type IN (?)", _types).Find(&results).Error
	
	return
}
 
// GetFromIndex 通过index获取内容  
func (obj *_EgCustomizedDataMgr) GetFromIndex(index int) (results []*EgCustomizedData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("index = ?", index).Find(&results).Error
	
	return
}

// GetBatchFromIndex 批量唯一主键查找 
func (obj *_EgCustomizedDataMgr) GetBatchFromIndex(indexs []int) (results []*EgCustomizedData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("index IN (?)", indexs).Find(&results).Error
	
	return
}
 
// GetFromValue 通过value获取内容  
func (obj *_EgCustomizedDataMgr) GetFromValue(value string) (results []*EgCustomizedData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("value = ?", value).Find(&results).Error
	
	return
}

// GetBatchFromValue 批量唯一主键查找 
func (obj *_EgCustomizedDataMgr) GetBatchFromValue(values []string) (results []*EgCustomizedData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("value IN (?)", values).Find(&results).Error
	
	return
}
 
// GetFromIDModule 通过id_module获取内容  
func (obj *_EgCustomizedDataMgr) GetFromIDModule(idModule int) (results []*EgCustomizedData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_module = ?", idModule).Find(&results).Error
	
	return
}

// GetBatchFromIDModule 批量唯一主键查找 
func (obj *_EgCustomizedDataMgr) GetBatchFromIDModule(idModules []int) (results []*EgCustomizedData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_module IN (?)", idModules).Find(&results).Error
	
	return
}
 
// GetFromPrice 通过price获取内容  
func (obj *_EgCustomizedDataMgr) GetFromPrice(price float64) (results []*EgCustomizedData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("price = ?", price).Find(&results).Error
	
	return
}

// GetBatchFromPrice 批量唯一主键查找 
func (obj *_EgCustomizedDataMgr) GetBatchFromPrice(prices []float64) (results []*EgCustomizedData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("price IN (?)", prices).Find(&results).Error
	
	return
}
 
// GetFromWeight 通过weight获取内容  
func (obj *_EgCustomizedDataMgr) GetFromWeight(weight float64) (results []*EgCustomizedData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("weight = ?", weight).Find(&results).Error
	
	return
}

// GetBatchFromWeight 批量唯一主键查找 
func (obj *_EgCustomizedDataMgr) GetBatchFromWeight(weights []float64) (results []*EgCustomizedData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("weight IN (?)", weights).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgCustomizedDataMgr) FetchByPrimaryKey(idCustomization uint32 ,_type bool ,index int ) (result EgCustomizedData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customization = ? AND type = ? AND index = ?", idCustomization , _type , index).Find(&result).Error
	
	return
}
 

 

	

