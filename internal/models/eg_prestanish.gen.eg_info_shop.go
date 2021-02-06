package	model	
import (	
"gorm.io/gorm"	
"fmt"	
"context"	
)	

type _EgInfoShopMgr struct {
	*_BaseMgr
}

// EgInfoShopMgr open func
func EgInfoShopMgr(db *gorm.DB) *_EgInfoShopMgr {
	if db == nil {
		panic(fmt.Errorf("EgInfoShopMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgInfoShopMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_info_shop"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgInfoShopMgr) GetTableName() string {
	return "eg_info_shop"
}

// Get 获取
func (obj *_EgInfoShopMgr) Get() (result EgInfoShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgInfoShopMgr) Gets() (results []*EgInfoShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDInfo id_info获取 
func (obj *_EgInfoShopMgr) WithIDInfo(idInfo uint32) Option {
	return optionFunc(func(o *options) { o.query["id_info"] = idInfo })
}

// WithIDShop id_shop获取 
func (obj *_EgInfoShopMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}


// GetByOption 功能选项模式获取
func (obj *_EgInfoShopMgr) GetByOption(opts ...Option) (result EgInfoShop, err error) {
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
func (obj *_EgInfoShopMgr) GetByOptions(opts ...Option) (results []*EgInfoShop, err error) {
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


// GetFromIDInfo 通过id_info获取内容  
func (obj *_EgInfoShopMgr) GetFromIDInfo(idInfo uint32) (results []*EgInfoShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_info = ?", idInfo).Find(&results).Error
	
	return
}

// GetBatchFromIDInfo 批量唯一主键查找 
func (obj *_EgInfoShopMgr) GetBatchFromIDInfo(idInfos []uint32) (results []*EgInfoShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_info IN (?)", idInfos).Find(&results).Error
	
	return
}
 
// GetFromIDShop 通过id_shop获取内容  
func (obj *_EgInfoShopMgr) GetFromIDShop(idShop uint32) (results []*EgInfoShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error
	
	return
}

// GetBatchFromIDShop 批量唯一主键查找 
func (obj *_EgInfoShopMgr) GetBatchFromIDShop(idShops []uint32) (results []*EgInfoShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgInfoShopMgr) FetchByPrimaryKey(idInfo uint32 ,idShop uint32 ) (result EgInfoShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_info = ? AND id_shop = ?", idInfo , idShop).Find(&result).Error
	
	return
}
 

 

	

