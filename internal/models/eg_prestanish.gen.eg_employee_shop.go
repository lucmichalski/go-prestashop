package	model	
import (	
"gorm.io/gorm"	
"fmt"	
"context"	
)	

type _EgEmployeeShopMgr struct {
	*_BaseMgr
}

// EgEmployeeShopMgr open func
func EgEmployeeShopMgr(db *gorm.DB) *_EgEmployeeShopMgr {
	if db == nil {
		panic(fmt.Errorf("EgEmployeeShopMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgEmployeeShopMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_employee_shop"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgEmployeeShopMgr) GetTableName() string {
	return "eg_employee_shop"
}

// Get 获取
func (obj *_EgEmployeeShopMgr) Get() (result EgEmployeeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgEmployeeShopMgr) Gets() (results []*EgEmployeeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDEmployee id_employee获取 
func (obj *_EgEmployeeShopMgr) WithIDEmployee(idEmployee uint32) Option {
	return optionFunc(func(o *options) { o.query["id_employee"] = idEmployee })
}

// WithIDShop id_shop获取 
func (obj *_EgEmployeeShopMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}


// GetByOption 功能选项模式获取
func (obj *_EgEmployeeShopMgr) GetByOption(opts ...Option) (result EgEmployeeShop, err error) {
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
func (obj *_EgEmployeeShopMgr) GetByOptions(opts ...Option) (results []*EgEmployeeShop, err error) {
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


// GetFromIDEmployee 通过id_employee获取内容  
func (obj *_EgEmployeeShopMgr) GetFromIDEmployee(idEmployee uint32) (results []*EgEmployeeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee = ?", idEmployee).Find(&results).Error
	
	return
}

// GetBatchFromIDEmployee 批量唯一主键查找 
func (obj *_EgEmployeeShopMgr) GetBatchFromIDEmployee(idEmployees []uint32) (results []*EgEmployeeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee IN (?)", idEmployees).Find(&results).Error
	
	return
}
 
// GetFromIDShop 通过id_shop获取内容  
func (obj *_EgEmployeeShopMgr) GetFromIDShop(idShop uint32) (results []*EgEmployeeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error
	
	return
}

// GetBatchFromIDShop 批量唯一主键查找 
func (obj *_EgEmployeeShopMgr) GetBatchFromIDShop(idShops []uint32) (results []*EgEmployeeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgEmployeeShopMgr) FetchByPrimaryKey(idEmployee uint32 ,idShop uint32 ) (result EgEmployeeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee = ? AND id_shop = ?", idEmployee , idShop).Find(&result).Error
	
	return
}
 

 
 // FetchIndexByIDShop  获取多个内容
 func (obj *_EgEmployeeShopMgr) FetchIndexByIDShop(idShop uint32 ) (results []*EgEmployeeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error
	
	return
}
 

	

