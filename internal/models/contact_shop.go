package	model	
import (	
"context"	
"gorm.io/gorm"	
"fmt"	
)	

type _EgContactShopMgr struct {
	*_BaseMgr
}

// EgContactShopMgr open func
func EgContactShopMgr(db *gorm.DB) *_EgContactShopMgr {
	if db == nil {
		panic(fmt.Errorf("EgContactShopMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgContactShopMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_contact_shop"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgContactShopMgr) GetTableName() string {
	return "eg_contact_shop"
}

// Get 获取
func (obj *_EgContactShopMgr) Get() (result EgContactShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgContactShopMgr) Gets() (results []*EgContactShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDContact id_contact获取 
func (obj *_EgContactShopMgr) WithIDContact(idContact uint32) Option {
	return optionFunc(func(o *options) { o.query["id_contact"] = idContact })
}

// WithIDShop id_shop获取 
func (obj *_EgContactShopMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}


// GetByOption 功能选项模式获取
func (obj *_EgContactShopMgr) GetByOption(opts ...Option) (result EgContactShop, err error) {
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
func (obj *_EgContactShopMgr) GetByOptions(opts ...Option) (results []*EgContactShop, err error) {
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


// GetFromIDContact 通过id_contact获取内容  
func (obj *_EgContactShopMgr) GetFromIDContact(idContact uint32) (results []*EgContactShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_contact = ?", idContact).Find(&results).Error
	
	return
}

// GetBatchFromIDContact 批量唯一主键查找 
func (obj *_EgContactShopMgr) GetBatchFromIDContact(idContacts []uint32) (results []*EgContactShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_contact IN (?)", idContacts).Find(&results).Error
	
	return
}
 
// GetFromIDShop 通过id_shop获取内容  
func (obj *_EgContactShopMgr) GetFromIDShop(idShop uint32) (results []*EgContactShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error
	
	return
}

// GetBatchFromIDShop 批量唯一主键查找 
func (obj *_EgContactShopMgr) GetBatchFromIDShop(idShops []uint32) (results []*EgContactShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgContactShopMgr) FetchByPrimaryKey(idContact uint32 ,idShop uint32 ) (result EgContactShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_contact = ? AND id_shop = ?", idContact , idShop).Find(&result).Error
	
	return
}
 

 
 // FetchIndexByIDShop  获取多个内容
 func (obj *_EgContactShopMgr) FetchIndexByIDShop(idShop uint32 ) (results []*EgContactShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error
	
	return
}
 

	

