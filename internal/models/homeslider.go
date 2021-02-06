package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgHomesliderMgr struct {
	*_BaseMgr
}

// EgHomesliderMgr open func
func EgHomesliderMgr(db *gorm.DB) *_EgHomesliderMgr {
	if db == nil {
		panic(fmt.Errorf("EgHomesliderMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgHomesliderMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_homeslider"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgHomesliderMgr) GetTableName() string {
	return "eg_homeslider"
}

// Get 获取
func (obj *_EgHomesliderMgr) Get() (result EgHomeslider, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgHomesliderMgr) Gets() (results []*EgHomeslider, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDHomesliderSlides id_homeslider_slides获取 
func (obj *_EgHomesliderMgr) WithIDHomesliderSlides(idHomesliderSlides uint32) Option {
	return optionFunc(func(o *options) { o.query["id_homeslider_slides"] = idHomesliderSlides })
}

// WithIDShop id_shop获取 
func (obj *_EgHomesliderMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}


// GetByOption 功能选项模式获取
func (obj *_EgHomesliderMgr) GetByOption(opts ...Option) (result EgHomeslider, err error) {
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
func (obj *_EgHomesliderMgr) GetByOptions(opts ...Option) (results []*EgHomeslider, err error) {
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


// GetFromIDHomesliderSlides 通过id_homeslider_slides获取内容  
func (obj *_EgHomesliderMgr) GetFromIDHomesliderSlides(idHomesliderSlides uint32) (results []*EgHomeslider, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_homeslider_slides = ?", idHomesliderSlides).Find(&results).Error
	
	return
}

// GetBatchFromIDHomesliderSlides 批量唯一主键查找 
func (obj *_EgHomesliderMgr) GetBatchFromIDHomesliderSlides(idHomesliderSlidess []uint32) (results []*EgHomeslider, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_homeslider_slides IN (?)", idHomesliderSlidess).Find(&results).Error
	
	return
}
 
// GetFromIDShop 通过id_shop获取内容  
func (obj *_EgHomesliderMgr) GetFromIDShop(idShop uint32) (results []*EgHomeslider, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error
	
	return
}

// GetBatchFromIDShop 批量唯一主键查找 
func (obj *_EgHomesliderMgr) GetBatchFromIDShop(idShops []uint32) (results []*EgHomeslider, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgHomesliderMgr) FetchByPrimaryKey(idHomesliderSlides uint32 ,idShop uint32 ) (result EgHomeslider, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_homeslider_slides = ? AND id_shop = ?", idHomesliderSlides , idShop).Find(&result).Error
	
	return
}
 

 

	

