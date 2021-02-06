package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgLinksmenutopMgr struct {
	*_BaseMgr
}

// EgLinksmenutopMgr open func
func EgLinksmenutopMgr(db *gorm.DB) *_EgLinksmenutopMgr {
	if db == nil {
		panic(fmt.Errorf("EgLinksmenutopMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgLinksmenutopMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_linksmenutop"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgLinksmenutopMgr) GetTableName() string {
	return "eg_linksmenutop"
}

// Get 获取
func (obj *_EgLinksmenutopMgr) Get() (result EgLinksmenutop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgLinksmenutopMgr) Gets() (results []*EgLinksmenutop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDLinksmenutop id_linksmenutop获取 
func (obj *_EgLinksmenutopMgr) WithIDLinksmenutop(idLinksmenutop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_linksmenutop"] = idLinksmenutop })
}

// WithIDShop id_shop获取 
func (obj *_EgLinksmenutopMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

// WithNewWindow new_window获取 
func (obj *_EgLinksmenutopMgr) WithNewWindow(newWindow bool) Option {
	return optionFunc(func(o *options) { o.query["new_window"] = newWindow })
}


// GetByOption 功能选项模式获取
func (obj *_EgLinksmenutopMgr) GetByOption(opts ...Option) (result EgLinksmenutop, err error) {
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
func (obj *_EgLinksmenutopMgr) GetByOptions(opts ...Option) (results []*EgLinksmenutop, err error) {
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


// GetFromIDLinksmenutop 通过id_linksmenutop获取内容  
func (obj *_EgLinksmenutopMgr)  GetFromIDLinksmenutop(idLinksmenutop uint32) (result EgLinksmenutop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_linksmenutop = ?", idLinksmenutop).Find(&result).Error
	
	return
}

// GetBatchFromIDLinksmenutop 批量唯一主键查找 
func (obj *_EgLinksmenutopMgr) GetBatchFromIDLinksmenutop(idLinksmenutops []uint32) (results []*EgLinksmenutop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_linksmenutop IN (?)", idLinksmenutops).Find(&results).Error
	
	return
}
 
// GetFromIDShop 通过id_shop获取内容  
func (obj *_EgLinksmenutopMgr) GetFromIDShop(idShop uint32) (results []*EgLinksmenutop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error
	
	return
}

// GetBatchFromIDShop 批量唯一主键查找 
func (obj *_EgLinksmenutopMgr) GetBatchFromIDShop(idShops []uint32) (results []*EgLinksmenutop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error
	
	return
}
 
// GetFromNewWindow 通过new_window获取内容  
func (obj *_EgLinksmenutopMgr) GetFromNewWindow(newWindow bool) (results []*EgLinksmenutop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("new_window = ?", newWindow).Find(&results).Error
	
	return
}

// GetBatchFromNewWindow 批量唯一主键查找 
func (obj *_EgLinksmenutopMgr) GetBatchFromNewWindow(newWindows []bool) (results []*EgLinksmenutop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("new_window IN (?)", newWindows).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgLinksmenutopMgr) FetchByPrimaryKey(idLinksmenutop uint32 ) (result EgLinksmenutop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_linksmenutop = ?", idLinksmenutop).Find(&result).Error
	
	return
}
 

 
 // FetchIndexByIDShop  获取多个内容
 func (obj *_EgLinksmenutopMgr) FetchIndexByIDShop(idShop uint32 ) (results []*EgLinksmenutop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error
	
	return
}
 

	

