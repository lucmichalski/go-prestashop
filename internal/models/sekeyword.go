package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
"time"	
)	

type _EgSekeywordMgr struct {
	*_BaseMgr
}

// EgSekeywordMgr open func
func EgSekeywordMgr(db *gorm.DB) *_EgSekeywordMgr {
	if db == nil {
		panic(fmt.Errorf("EgSekeywordMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgSekeywordMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_sekeyword"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgSekeywordMgr) GetTableName() string {
	return "eg_sekeyword"
}

// Get 获取
func (obj *_EgSekeywordMgr) Get() (result EgSekeyword, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgSekeywordMgr) Gets() (results []*EgSekeyword, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDSekeyword id_sekeyword获取 
func (obj *_EgSekeywordMgr) WithIDSekeyword(idSekeyword uint32) Option {
	return optionFunc(func(o *options) { o.query["id_sekeyword"] = idSekeyword })
}

// WithIDShop id_shop获取 
func (obj *_EgSekeywordMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

// WithIDShopGroup id_shop_group获取 
func (obj *_EgSekeywordMgr) WithIDShopGroup(idShopGroup uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop_group"] = idShopGroup })
}

// WithKeyword keyword获取 
func (obj *_EgSekeywordMgr) WithKeyword(keyword string) Option {
	return optionFunc(func(o *options) { o.query["keyword"] = keyword })
}

// WithDateAdd date_add获取 
func (obj *_EgSekeywordMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}


// GetByOption 功能选项模式获取
func (obj *_EgSekeywordMgr) GetByOption(opts ...Option) (result EgSekeyword, err error) {
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
func (obj *_EgSekeywordMgr) GetByOptions(opts ...Option) (results []*EgSekeyword, err error) {
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


// GetFromIDSekeyword 通过id_sekeyword获取内容  
func (obj *_EgSekeywordMgr)  GetFromIDSekeyword(idSekeyword uint32) (result EgSekeyword, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_sekeyword = ?", idSekeyword).Find(&result).Error
	
	return
}

// GetBatchFromIDSekeyword 批量唯一主键查找 
func (obj *_EgSekeywordMgr) GetBatchFromIDSekeyword(idSekeywords []uint32) (results []*EgSekeyword, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_sekeyword IN (?)", idSekeywords).Find(&results).Error
	
	return
}
 
// GetFromIDShop 通过id_shop获取内容  
func (obj *_EgSekeywordMgr) GetFromIDShop(idShop uint32) (results []*EgSekeyword, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error
	
	return
}

// GetBatchFromIDShop 批量唯一主键查找 
func (obj *_EgSekeywordMgr) GetBatchFromIDShop(idShops []uint32) (results []*EgSekeyword, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error
	
	return
}
 
// GetFromIDShopGroup 通过id_shop_group获取内容  
func (obj *_EgSekeywordMgr) GetFromIDShopGroup(idShopGroup uint32) (results []*EgSekeyword, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_group = ?", idShopGroup).Find(&results).Error
	
	return
}

// GetBatchFromIDShopGroup 批量唯一主键查找 
func (obj *_EgSekeywordMgr) GetBatchFromIDShopGroup(idShopGroups []uint32) (results []*EgSekeyword, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_group IN (?)", idShopGroups).Find(&results).Error
	
	return
}
 
// GetFromKeyword 通过keyword获取内容  
func (obj *_EgSekeywordMgr) GetFromKeyword(keyword string) (results []*EgSekeyword, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("keyword = ?", keyword).Find(&results).Error
	
	return
}

// GetBatchFromKeyword 批量唯一主键查找 
func (obj *_EgSekeywordMgr) GetBatchFromKeyword(keywords []string) (results []*EgSekeyword, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("keyword IN (?)", keywords).Find(&results).Error
	
	return
}
 
// GetFromDateAdd 通过date_add获取内容  
func (obj *_EgSekeywordMgr) GetFromDateAdd(dateAdd time.Time) (results []*EgSekeyword, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error
	
	return
}

// GetBatchFromDateAdd 批量唯一主键查找 
func (obj *_EgSekeywordMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*EgSekeyword, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgSekeywordMgr) FetchByPrimaryKey(idSekeyword uint32 ) (result EgSekeyword, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_sekeyword = ?", idSekeyword).Find(&result).Error
	
	return
}
 

 

	

