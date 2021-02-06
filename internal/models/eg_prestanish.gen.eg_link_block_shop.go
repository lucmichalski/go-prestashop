package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgLinkBlockShopMgr struct {
	*_BaseMgr
}

// EgLinkBlockShopMgr open func
func EgLinkBlockShopMgr(db *gorm.DB) *_EgLinkBlockShopMgr {
	if db == nil {
		panic(fmt.Errorf("EgLinkBlockShopMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgLinkBlockShopMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_link_block_shop"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgLinkBlockShopMgr) GetTableName() string {
	return "eg_link_block_shop"
}

// Get 获取
func (obj *_EgLinkBlockShopMgr) Get() (result EgLinkBlockShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgLinkBlockShopMgr) Gets() (results []*EgLinkBlockShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDLinkBlock id_link_block获取 
func (obj *_EgLinkBlockShopMgr) WithIDLinkBlock(idLinkBlock uint32) Option {
	return optionFunc(func(o *options) { o.query["id_link_block"] = idLinkBlock })
}

// WithIDShop id_shop获取 
func (obj *_EgLinkBlockShopMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}


// GetByOption 功能选项模式获取
func (obj *_EgLinkBlockShopMgr) GetByOption(opts ...Option) (result EgLinkBlockShop, err error) {
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
func (obj *_EgLinkBlockShopMgr) GetByOptions(opts ...Option) (results []*EgLinkBlockShop, err error) {
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


// GetFromIDLinkBlock 通过id_link_block获取内容  
func (obj *_EgLinkBlockShopMgr) GetFromIDLinkBlock(idLinkBlock uint32) (results []*EgLinkBlockShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_link_block = ?", idLinkBlock).Find(&results).Error
	
	return
}

// GetBatchFromIDLinkBlock 批量唯一主键查找 
func (obj *_EgLinkBlockShopMgr) GetBatchFromIDLinkBlock(idLinkBlocks []uint32) (results []*EgLinkBlockShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_link_block IN (?)", idLinkBlocks).Find(&results).Error
	
	return
}
 
// GetFromIDShop 通过id_shop获取内容  
func (obj *_EgLinkBlockShopMgr) GetFromIDShop(idShop uint32) (results []*EgLinkBlockShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error
	
	return
}

// GetBatchFromIDShop 批量唯一主键查找 
func (obj *_EgLinkBlockShopMgr) GetBatchFromIDShop(idShops []uint32) (results []*EgLinkBlockShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgLinkBlockShopMgr) FetchByPrimaryKey(idLinkBlock uint32 ,idShop uint32 ) (result EgLinkBlockShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_link_block = ? AND id_shop = ?", idLinkBlock , idShop).Find(&result).Error
	
	return
}
 

 

	

