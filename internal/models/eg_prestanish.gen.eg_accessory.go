package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgAccessoryMgr struct {
	*_BaseMgr
}

// EgAccessoryMgr open func
func EgAccessoryMgr(db *gorm.DB) *_EgAccessoryMgr {
	if db == nil {
		panic(fmt.Errorf("EgAccessoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgAccessoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_accessory"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgAccessoryMgr) GetTableName() string {
	return "eg_accessory"
}

// Get 获取
func (obj *_EgAccessoryMgr) Get() (result EgAccessory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgAccessoryMgr) Gets() (results []*EgAccessory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDProduct1 id_product_1获取 
func (obj *_EgAccessoryMgr) WithIDProduct1(idProduct1 uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product_1"] = idProduct1 })
}

// WithIDProduct2 id_product_2获取 
func (obj *_EgAccessoryMgr) WithIDProduct2(idProduct2 uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product_2"] = idProduct2 })
}


// GetByOption 功能选项模式获取
func (obj *_EgAccessoryMgr) GetByOption(opts ...Option) (result EgAccessory, err error) {
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
func (obj *_EgAccessoryMgr) GetByOptions(opts ...Option) (results []*EgAccessory, err error) {
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


// GetFromIDProduct1 通过id_product_1获取内容  
func (obj *_EgAccessoryMgr) GetFromIDProduct1(idProduct1 uint32) (results []*EgAccessory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_1 = ?", idProduct1).Find(&results).Error
	
	return
}

// GetBatchFromIDProduct1 批量唯一主键查找 
func (obj *_EgAccessoryMgr) GetBatchFromIDProduct1(idProduct1s []uint32) (results []*EgAccessory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_1 IN (?)", idProduct1s).Find(&results).Error
	
	return
}
 
// GetFromIDProduct2 通过id_product_2获取内容  
func (obj *_EgAccessoryMgr) GetFromIDProduct2(idProduct2 uint32) (results []*EgAccessory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_2 = ?", idProduct2).Find(&results).Error
	
	return
}

// GetBatchFromIDProduct2 批量唯一主键查找 
func (obj *_EgAccessoryMgr) GetBatchFromIDProduct2(idProduct2s []uint32) (results []*EgAccessory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_2 IN (?)", idProduct2s).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 

 
 // FetchIndexByAccessoryProduct  获取多个内容
 func (obj *_EgAccessoryMgr) FetchIndexByAccessoryProduct(idProduct1 uint32 ,idProduct2 uint32 ) (results []*EgAccessory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_1 = ? AND id_product_2 = ?", idProduct1 , idProduct2).Find(&results).Error
	
	return
}
 

	

