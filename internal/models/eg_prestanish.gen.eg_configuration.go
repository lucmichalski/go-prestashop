package	model	
import (	
"time"	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgConfigurationMgr struct {
	*_BaseMgr
}

// EgConfigurationMgr open func
func EgConfigurationMgr(db *gorm.DB) *_EgConfigurationMgr {
	if db == nil {
		panic(fmt.Errorf("EgConfigurationMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgConfigurationMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_configuration"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgConfigurationMgr) GetTableName() string {
	return "eg_configuration"
}

// Get 获取
func (obj *_EgConfigurationMgr) Get() (result EgConfiguration, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgConfigurationMgr) Gets() (results []*EgConfiguration, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDConfiguration id_configuration获取 
func (obj *_EgConfigurationMgr) WithIDConfiguration(idConfiguration uint32) Option {
	return optionFunc(func(o *options) { o.query["id_configuration"] = idConfiguration })
}

// WithIDShopGroup id_shop_group获取 
func (obj *_EgConfigurationMgr) WithIDShopGroup(idShopGroup uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop_group"] = idShopGroup })
}

// WithIDShop id_shop获取 
func (obj *_EgConfigurationMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

// WithName name获取 
func (obj *_EgConfigurationMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithValue value获取 
func (obj *_EgConfigurationMgr) WithValue(value string) Option {
	return optionFunc(func(o *options) { o.query["value"] = value })
}

// WithDateAdd date_add获取 
func (obj *_EgConfigurationMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

// WithDateUpd date_upd获取 
func (obj *_EgConfigurationMgr) WithDateUpd(dateUpd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_upd"] = dateUpd })
}


// GetByOption 功能选项模式获取
func (obj *_EgConfigurationMgr) GetByOption(opts ...Option) (result EgConfiguration, err error) {
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
func (obj *_EgConfigurationMgr) GetByOptions(opts ...Option) (results []*EgConfiguration, err error) {
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


// GetFromIDConfiguration 通过id_configuration获取内容  
func (obj *_EgConfigurationMgr)  GetFromIDConfiguration(idConfiguration uint32) (result EgConfiguration, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_configuration = ?", idConfiguration).Find(&result).Error
	
	return
}

// GetBatchFromIDConfiguration 批量唯一主键查找 
func (obj *_EgConfigurationMgr) GetBatchFromIDConfiguration(idConfigurations []uint32) (results []*EgConfiguration, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_configuration IN (?)", idConfigurations).Find(&results).Error
	
	return
}
 
// GetFromIDShopGroup 通过id_shop_group获取内容  
func (obj *_EgConfigurationMgr) GetFromIDShopGroup(idShopGroup uint32) (results []*EgConfiguration, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_group = ?", idShopGroup).Find(&results).Error
	
	return
}

// GetBatchFromIDShopGroup 批量唯一主键查找 
func (obj *_EgConfigurationMgr) GetBatchFromIDShopGroup(idShopGroups []uint32) (results []*EgConfiguration, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_group IN (?)", idShopGroups).Find(&results).Error
	
	return
}
 
// GetFromIDShop 通过id_shop获取内容  
func (obj *_EgConfigurationMgr) GetFromIDShop(idShop uint32) (results []*EgConfiguration, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error
	
	return
}

// GetBatchFromIDShop 批量唯一主键查找 
func (obj *_EgConfigurationMgr) GetBatchFromIDShop(idShops []uint32) (results []*EgConfiguration, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error
	
	return
}
 
// GetFromName 通过name获取内容  
func (obj *_EgConfigurationMgr) GetFromName(name string) (results []*EgConfiguration, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error
	
	return
}

// GetBatchFromName 批量唯一主键查找 
func (obj *_EgConfigurationMgr) GetBatchFromName(names []string) (results []*EgConfiguration, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error
	
	return
}
 
// GetFromValue 通过value获取内容  
func (obj *_EgConfigurationMgr) GetFromValue(value string) (results []*EgConfiguration, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("value = ?", value).Find(&results).Error
	
	return
}

// GetBatchFromValue 批量唯一主键查找 
func (obj *_EgConfigurationMgr) GetBatchFromValue(values []string) (results []*EgConfiguration, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("value IN (?)", values).Find(&results).Error
	
	return
}
 
// GetFromDateAdd 通过date_add获取内容  
func (obj *_EgConfigurationMgr) GetFromDateAdd(dateAdd time.Time) (results []*EgConfiguration, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error
	
	return
}

// GetBatchFromDateAdd 批量唯一主键查找 
func (obj *_EgConfigurationMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*EgConfiguration, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error
	
	return
}
 
// GetFromDateUpd 通过date_upd获取内容  
func (obj *_EgConfigurationMgr) GetFromDateUpd(dateUpd time.Time) (results []*EgConfiguration, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd = ?", dateUpd).Find(&results).Error
	
	return
}

// GetBatchFromDateUpd 批量唯一主键查找 
func (obj *_EgConfigurationMgr) GetBatchFromDateUpd(dateUpds []time.Time) (results []*EgConfiguration, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd IN (?)", dateUpds).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgConfigurationMgr) FetchByPrimaryKey(idConfiguration uint32 ) (result EgConfiguration, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_configuration = ?", idConfiguration).Find(&result).Error
	
	return
}
 

 
 // FetchIndexByIDShopGroup  获取多个内容
 func (obj *_EgConfigurationMgr) FetchIndexByIDShopGroup(idShopGroup uint32 ) (results []*EgConfiguration, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_group = ?", idShopGroup).Find(&results).Error
	
	return
}
 
 // FetchIndexByIDShop  获取多个内容
 func (obj *_EgConfigurationMgr) FetchIndexByIDShop(idShop uint32 ) (results []*EgConfiguration, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error
	
	return
}
 
 // FetchIndexByName  获取多个内容
 func (obj *_EgConfigurationMgr) FetchIndexByName(name string ) (results []*EgConfiguration, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error
	
	return
}
 

	

