package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgImageShopMgr struct {
	*_BaseMgr
}

// EgImageShopMgr open func
func EgImageShopMgr(db *gorm.DB) *_EgImageShopMgr {
	if db == nil {
		panic(fmt.Errorf("EgImageShopMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgImageShopMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_image_shop"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgImageShopMgr) GetTableName() string {
	return "eg_image_shop"
}

// Get 获取
func (obj *_EgImageShopMgr) Get() (result EgImageShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgImageShopMgr) Gets() (results []*EgImageShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDProduct id_product获取 
func (obj *_EgImageShopMgr) WithIDProduct(idProduct uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product"] = idProduct })
}

// WithIDImage id_image获取 
func (obj *_EgImageShopMgr) WithIDImage(idImage uint32) Option {
	return optionFunc(func(o *options) { o.query["id_image"] = idImage })
}

// WithIDShop id_shop获取 
func (obj *_EgImageShopMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

// WithCover cover获取 
func (obj *_EgImageShopMgr) WithCover(cover bool) Option {
	return optionFunc(func(o *options) { o.query["cover"] = cover })
}


// GetByOption 功能选项模式获取
func (obj *_EgImageShopMgr) GetByOption(opts ...Option) (result EgImageShop, err error) {
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
func (obj *_EgImageShopMgr) GetByOptions(opts ...Option) (results []*EgImageShop, err error) {
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


// GetFromIDProduct 通过id_product获取内容  
func (obj *_EgImageShopMgr) GetFromIDProduct(idProduct uint32) (results []*EgImageShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ?", idProduct).Find(&results).Error
	
	return
}

// GetBatchFromIDProduct 批量唯一主键查找 
func (obj *_EgImageShopMgr) GetBatchFromIDProduct(idProducts []uint32) (results []*EgImageShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product IN (?)", idProducts).Find(&results).Error
	
	return
}
 
// GetFromIDImage 通过id_image获取内容  
func (obj *_EgImageShopMgr) GetFromIDImage(idImage uint32) (results []*EgImageShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_image = ?", idImage).Find(&results).Error
	
	return
}

// GetBatchFromIDImage 批量唯一主键查找 
func (obj *_EgImageShopMgr) GetBatchFromIDImage(idImages []uint32) (results []*EgImageShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_image IN (?)", idImages).Find(&results).Error
	
	return
}
 
// GetFromIDShop 通过id_shop获取内容  
func (obj *_EgImageShopMgr) GetFromIDShop(idShop uint32) (results []*EgImageShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error
	
	return
}

// GetBatchFromIDShop 批量唯一主键查找 
func (obj *_EgImageShopMgr) GetBatchFromIDShop(idShops []uint32) (results []*EgImageShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error
	
	return
}
 
// GetFromCover 通过cover获取内容  
func (obj *_EgImageShopMgr) GetFromCover(cover bool) (results []*EgImageShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cover = ?", cover).Find(&results).Error
	
	return
}

// GetBatchFromCover 批量唯一主键查找 
func (obj *_EgImageShopMgr) GetBatchFromCover(covers []bool) (results []*EgImageShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cover IN (?)", covers).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgImageShopMgr) FetchByPrimaryKey(idImage uint32 ,idShop uint32 ) (result EgImageShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_image = ? AND id_shop = ?", idImage , idShop).Find(&result).Error
	
	return
}
 
 // FetchUniqueIndexByIDProduct primay or index 获取唯一内容
 func (obj *_EgImageShopMgr) FetchUniqueIndexByIDProduct(idProduct uint32 ,idShop uint32 ,cover bool ) (result EgImageShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ? AND id_shop = ? AND cover = ?", idProduct , idShop , cover).Find(&result).Error
	
	return
}
 

 
 // FetchIndexByIDShop  获取多个内容
 func (obj *_EgImageShopMgr) FetchIndexByIDShop(idShop uint32 ) (results []*EgImageShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error
	
	return
}
 

	

