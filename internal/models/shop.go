package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgShopMgr struct {
	*_BaseMgr
}

// EgShopMgr open func
func EgShopMgr(db *gorm.DB) *_EgShopMgr {
	if db == nil {
		panic(fmt.Errorf("EgShopMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgShopMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_shop"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgShopMgr) GetTableName() string {
	return "eg_shop"
}

// Get 获取
func (obj *_EgShopMgr) Get() (result EgShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgShopMgr) Gets() (results []*EgShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDShop id_shop获取 
func (obj *_EgShopMgr) WithIDShop(idShop int) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

// WithIDShopGroup id_shop_group获取 
func (obj *_EgShopMgr) WithIDShopGroup(idShopGroup int) Option {
	return optionFunc(func(o *options) { o.query["id_shop_group"] = idShopGroup })
}

// WithName name获取 
func (obj *_EgShopMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithIDCategory id_category获取 
func (obj *_EgShopMgr) WithIDCategory(idCategory int) Option {
	return optionFunc(func(o *options) { o.query["id_category"] = idCategory })
}

// WithThemeName theme_name获取 
func (obj *_EgShopMgr) WithThemeName(themeName string) Option {
	return optionFunc(func(o *options) { o.query["theme_name"] = themeName })
}

// WithActive active获取 
func (obj *_EgShopMgr) WithActive(active bool) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}

// WithDeleted deleted获取 
func (obj *_EgShopMgr) WithDeleted(deleted bool) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}


// GetByOption 功能选项模式获取
func (obj *_EgShopMgr) GetByOption(opts ...Option) (result EgShop, err error) {
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
func (obj *_EgShopMgr) GetByOptions(opts ...Option) (results []*EgShop, err error) {
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


// GetFromIDShop 通过id_shop获取内容  
func (obj *_EgShopMgr)  GetFromIDShop(idShop int) (result EgShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&result).Error
	
	return
}

// GetBatchFromIDShop 批量唯一主键查找 
func (obj *_EgShopMgr) GetBatchFromIDShop(idShops []int) (results []*EgShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error
	
	return
}
 
// GetFromIDShopGroup 通过id_shop_group获取内容  
func (obj *_EgShopMgr) GetFromIDShopGroup(idShopGroup int) (results []*EgShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_group = ?", idShopGroup).Find(&results).Error
	
	return
}

// GetBatchFromIDShopGroup 批量唯一主键查找 
func (obj *_EgShopMgr) GetBatchFromIDShopGroup(idShopGroups []int) (results []*EgShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_group IN (?)", idShopGroups).Find(&results).Error
	
	return
}
 
// GetFromName 通过name获取内容  
func (obj *_EgShopMgr) GetFromName(name string) (results []*EgShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error
	
	return
}

// GetBatchFromName 批量唯一主键查找 
func (obj *_EgShopMgr) GetBatchFromName(names []string) (results []*EgShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error
	
	return
}
 
// GetFromIDCategory 通过id_category获取内容  
func (obj *_EgShopMgr) GetFromIDCategory(idCategory int) (results []*EgShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_category = ?", idCategory).Find(&results).Error
	
	return
}

// GetBatchFromIDCategory 批量唯一主键查找 
func (obj *_EgShopMgr) GetBatchFromIDCategory(idCategorys []int) (results []*EgShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_category IN (?)", idCategorys).Find(&results).Error
	
	return
}
 
// GetFromThemeName 通过theme_name获取内容  
func (obj *_EgShopMgr) GetFromThemeName(themeName string) (results []*EgShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("theme_name = ?", themeName).Find(&results).Error
	
	return
}

// GetBatchFromThemeName 批量唯一主键查找 
func (obj *_EgShopMgr) GetBatchFromThemeName(themeNames []string) (results []*EgShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("theme_name IN (?)", themeNames).Find(&results).Error
	
	return
}
 
// GetFromActive 通过active获取内容  
func (obj *_EgShopMgr) GetFromActive(active bool) (results []*EgShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active = ?", active).Find(&results).Error
	
	return
}

// GetBatchFromActive 批量唯一主键查找 
func (obj *_EgShopMgr) GetBatchFromActive(actives []bool) (results []*EgShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active IN (?)", actives).Find(&results).Error
	
	return
}
 
// GetFromDeleted 通过deleted获取内容  
func (obj *_EgShopMgr) GetFromDeleted(deleted bool) (results []*EgShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("deleted = ?", deleted).Find(&results).Error
	
	return
}

// GetBatchFromDeleted 批量唯一主键查找 
func (obj *_EgShopMgr) GetBatchFromDeleted(deleteds []bool) (results []*EgShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("deleted IN (?)", deleteds).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgShopMgr) FetchByPrimaryKey(idShop int ) (result EgShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&result).Error
	
	return
}
 

 
 // FetchIndexByIDX667E487AF5C9E40  获取多个内容
 func (obj *_EgShopMgr) FetchIndexByIDX667E487AF5C9E40(idShopGroup int ) (results []*EgShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_group = ?", idShopGroup).Find(&results).Error
	
	return
}
 

	

