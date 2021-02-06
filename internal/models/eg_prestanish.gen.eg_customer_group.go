package	model	
import (	
"context"	
"gorm.io/gorm"	
"fmt"	
)	

type _EgCustomerGroupMgr struct {
	*_BaseMgr
}

// EgCustomerGroupMgr open func
func EgCustomerGroupMgr(db *gorm.DB) *_EgCustomerGroupMgr {
	if db == nil {
		panic(fmt.Errorf("EgCustomerGroupMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgCustomerGroupMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_customer_group"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgCustomerGroupMgr) GetTableName() string {
	return "eg_customer_group"
}

// Get 获取
func (obj *_EgCustomerGroupMgr) Get() (result EgCustomerGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgCustomerGroupMgr) Gets() (results []*EgCustomerGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDCustomer id_customer获取 
func (obj *_EgCustomerGroupMgr) WithIDCustomer(idCustomer uint32) Option {
	return optionFunc(func(o *options) { o.query["id_customer"] = idCustomer })
}

// WithIDGroup id_group获取 
func (obj *_EgCustomerGroupMgr) WithIDGroup(idGroup uint32) Option {
	return optionFunc(func(o *options) { o.query["id_group"] = idGroup })
}


// GetByOption 功能选项模式获取
func (obj *_EgCustomerGroupMgr) GetByOption(opts ...Option) (result EgCustomerGroup, err error) {
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
func (obj *_EgCustomerGroupMgr) GetByOptions(opts ...Option) (results []*EgCustomerGroup, err error) {
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


// GetFromIDCustomer 通过id_customer获取内容  
func (obj *_EgCustomerGroupMgr) GetFromIDCustomer(idCustomer uint32) (results []*EgCustomerGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer = ?", idCustomer).Find(&results).Error
	
	return
}

// GetBatchFromIDCustomer 批量唯一主键查找 
func (obj *_EgCustomerGroupMgr) GetBatchFromIDCustomer(idCustomers []uint32) (results []*EgCustomerGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer IN (?)", idCustomers).Find(&results).Error
	
	return
}
 
// GetFromIDGroup 通过id_group获取内容  
func (obj *_EgCustomerGroupMgr) GetFromIDGroup(idGroup uint32) (results []*EgCustomerGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_group = ?", idGroup).Find(&results).Error
	
	return
}

// GetBatchFromIDGroup 批量唯一主键查找 
func (obj *_EgCustomerGroupMgr) GetBatchFromIDGroup(idGroups []uint32) (results []*EgCustomerGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_group IN (?)", idGroups).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgCustomerGroupMgr) FetchByPrimaryKey(idCustomer uint32 ,idGroup uint32 ) (result EgCustomerGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer = ? AND id_group = ?", idCustomer , idGroup).Find(&result).Error
	
	return
}
 

 
 // FetchIndexByIDCustomer  获取多个内容
 func (obj *_EgCustomerGroupMgr) FetchIndexByIDCustomer(idCustomer uint32 ) (results []*EgCustomerGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer = ?", idCustomer).Find(&results).Error
	
	return
}
 
 // FetchIndexByCustomerLogin  获取多个内容
 func (obj *_EgCustomerGroupMgr) FetchIndexByCustomerLogin(idGroup uint32 ) (results []*EgCustomerGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_group = ?", idGroup).Find(&results).Error
	
	return
}
 

	

