package	model	
import (	
"context"	
"gorm.io/gorm"	
"fmt"	
)	

type _EgLayeredFilterShopMgr struct {
	*_BaseMgr
}

// EgLayeredFilterShopMgr open func
func EgLayeredFilterShopMgr(db *gorm.DB) *_EgLayeredFilterShopMgr {
	if db == nil {
		panic(fmt.Errorf("EgLayeredFilterShopMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgLayeredFilterShopMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_layered_filter_shop"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgLayeredFilterShopMgr) GetTableName() string {
	return "eg_layered_filter_shop"
}

// Get 获取
func (obj *_EgLayeredFilterShopMgr) Get() (result EgLayeredFilterShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgLayeredFilterShopMgr) Gets() (results []*EgLayeredFilterShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDLayeredFilter id_layered_filter获取 
func (obj *_EgLayeredFilterShopMgr) WithIDLayeredFilter(idLayeredFilter uint32) Option {
	return optionFunc(func(o *options) { o.query["id_layered_filter"] = idLayeredFilter })
}

// WithIDShop id_shop获取 
func (obj *_EgLayeredFilterShopMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}


// GetByOption 功能选项模式获取
func (obj *_EgLayeredFilterShopMgr) GetByOption(opts ...Option) (result EgLayeredFilterShop, err error) {
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
func (obj *_EgLayeredFilterShopMgr) GetByOptions(opts ...Option) (results []*EgLayeredFilterShop, err error) {
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


// GetFromIDLayeredFilter 通过id_layered_filter获取内容  
func (obj *_EgLayeredFilterShopMgr) GetFromIDLayeredFilter(idLayeredFilter uint32) (results []*EgLayeredFilterShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_layered_filter = ?", idLayeredFilter).Find(&results).Error
	
	return
}

// GetBatchFromIDLayeredFilter 批量唯一主键查找 
func (obj *_EgLayeredFilterShopMgr) GetBatchFromIDLayeredFilter(idLayeredFilters []uint32) (results []*EgLayeredFilterShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_layered_filter IN (?)", idLayeredFilters).Find(&results).Error
	
	return
}
 
// GetFromIDShop 通过id_shop获取内容  
func (obj *_EgLayeredFilterShopMgr) GetFromIDShop(idShop uint32) (results []*EgLayeredFilterShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error
	
	return
}

// GetBatchFromIDShop 批量唯一主键查找 
func (obj *_EgLayeredFilterShopMgr) GetBatchFromIDShop(idShops []uint32) (results []*EgLayeredFilterShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgLayeredFilterShopMgr) FetchByPrimaryKey(idLayeredFilter uint32 ,idShop uint32 ) (result EgLayeredFilterShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_layered_filter = ? AND id_shop = ?", idLayeredFilter , idShop).Find(&result).Error
	
	return
}
 

 
 // FetchIndexByIDShop  获取多个内容
 func (obj *_EgLayeredFilterShopMgr) FetchIndexByIDShop(idShop uint32 ) (results []*EgLayeredFilterShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error
	
	return
}
 

	

