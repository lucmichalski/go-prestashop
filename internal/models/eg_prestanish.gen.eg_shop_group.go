package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgShopGroupMgr struct {
	*_BaseMgr
}

// EgShopGroupMgr open func
func EgShopGroupMgr(db *gorm.DB) *_EgShopGroupMgr {
	if db == nil {
		panic(fmt.Errorf("EgShopGroupMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgShopGroupMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_shop_group"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgShopGroupMgr) GetTableName() string {
	return "eg_shop_group"
}

// Get 获取
func (obj *_EgShopGroupMgr) Get() (result EgShopGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgShopGroupMgr) Gets() (results []*EgShopGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDShopGroup id_shop_group获取 
func (obj *_EgShopGroupMgr) WithIDShopGroup(idShopGroup int) Option {
	return optionFunc(func(o *options) { o.query["id_shop_group"] = idShopGroup })
}

// WithName name获取 
func (obj *_EgShopGroupMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithShareCustomer share_customer获取 
func (obj *_EgShopGroupMgr) WithShareCustomer(shareCustomer bool) Option {
	return optionFunc(func(o *options) { o.query["share_customer"] = shareCustomer })
}

// WithShareOrder share_order获取 
func (obj *_EgShopGroupMgr) WithShareOrder(shareOrder bool) Option {
	return optionFunc(func(o *options) { o.query["share_order"] = shareOrder })
}

// WithShareStock share_stock获取 
func (obj *_EgShopGroupMgr) WithShareStock(shareStock bool) Option {
	return optionFunc(func(o *options) { o.query["share_stock"] = shareStock })
}

// WithActive active获取 
func (obj *_EgShopGroupMgr) WithActive(active bool) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}

// WithDeleted deleted获取 
func (obj *_EgShopGroupMgr) WithDeleted(deleted bool) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}


// GetByOption 功能选项模式获取
func (obj *_EgShopGroupMgr) GetByOption(opts ...Option) (result EgShopGroup, err error) {
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
func (obj *_EgShopGroupMgr) GetByOptions(opts ...Option) (results []*EgShopGroup, err error) {
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


// GetFromIDShopGroup 通过id_shop_group获取内容  
func (obj *_EgShopGroupMgr)  GetFromIDShopGroup(idShopGroup int) (result EgShopGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_group = ?", idShopGroup).Find(&result).Error
	
	return
}

// GetBatchFromIDShopGroup 批量唯一主键查找 
func (obj *_EgShopGroupMgr) GetBatchFromIDShopGroup(idShopGroups []int) (results []*EgShopGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_group IN (?)", idShopGroups).Find(&results).Error
	
	return
}
 
// GetFromName 通过name获取内容  
func (obj *_EgShopGroupMgr) GetFromName(name string) (results []*EgShopGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error
	
	return
}

// GetBatchFromName 批量唯一主键查找 
func (obj *_EgShopGroupMgr) GetBatchFromName(names []string) (results []*EgShopGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error
	
	return
}
 
// GetFromShareCustomer 通过share_customer获取内容  
func (obj *_EgShopGroupMgr) GetFromShareCustomer(shareCustomer bool) (results []*EgShopGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("share_customer = ?", shareCustomer).Find(&results).Error
	
	return
}

// GetBatchFromShareCustomer 批量唯一主键查找 
func (obj *_EgShopGroupMgr) GetBatchFromShareCustomer(shareCustomers []bool) (results []*EgShopGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("share_customer IN (?)", shareCustomers).Find(&results).Error
	
	return
}
 
// GetFromShareOrder 通过share_order获取内容  
func (obj *_EgShopGroupMgr) GetFromShareOrder(shareOrder bool) (results []*EgShopGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("share_order = ?", shareOrder).Find(&results).Error
	
	return
}

// GetBatchFromShareOrder 批量唯一主键查找 
func (obj *_EgShopGroupMgr) GetBatchFromShareOrder(shareOrders []bool) (results []*EgShopGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("share_order IN (?)", shareOrders).Find(&results).Error
	
	return
}
 
// GetFromShareStock 通过share_stock获取内容  
func (obj *_EgShopGroupMgr) GetFromShareStock(shareStock bool) (results []*EgShopGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("share_stock = ?", shareStock).Find(&results).Error
	
	return
}

// GetBatchFromShareStock 批量唯一主键查找 
func (obj *_EgShopGroupMgr) GetBatchFromShareStock(shareStocks []bool) (results []*EgShopGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("share_stock IN (?)", shareStocks).Find(&results).Error
	
	return
}
 
// GetFromActive 通过active获取内容  
func (obj *_EgShopGroupMgr) GetFromActive(active bool) (results []*EgShopGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active = ?", active).Find(&results).Error
	
	return
}

// GetBatchFromActive 批量唯一主键查找 
func (obj *_EgShopGroupMgr) GetBatchFromActive(actives []bool) (results []*EgShopGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active IN (?)", actives).Find(&results).Error
	
	return
}
 
// GetFromDeleted 通过deleted获取内容  
func (obj *_EgShopGroupMgr) GetFromDeleted(deleted bool) (results []*EgShopGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("deleted = ?", deleted).Find(&results).Error
	
	return
}

// GetBatchFromDeleted 批量唯一主键查找 
func (obj *_EgShopGroupMgr) GetBatchFromDeleted(deleteds []bool) (results []*EgShopGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("deleted IN (?)", deleteds).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgShopGroupMgr) FetchByPrimaryKey(idShopGroup int ) (result EgShopGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_group = ?", idShopGroup).Find(&result).Error
	
	return
}
 

 

	

