package	model	
import (	
"context"	
"gorm.io/gorm"	
"fmt"	
)	

type _EgImageMgr struct {
	*_BaseMgr
}

// EgImageMgr open func
func EgImageMgr(db *gorm.DB) *_EgImageMgr {
	if db == nil {
		panic(fmt.Errorf("EgImageMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgImageMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_image"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgImageMgr) GetTableName() string {
	return "eg_image"
}

// Get 获取
func (obj *_EgImageMgr) Get() (result EgImage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgImageMgr) Gets() (results []*EgImage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDImage id_image获取 
func (obj *_EgImageMgr) WithIDImage(idImage uint32) Option {
	return optionFunc(func(o *options) { o.query["id_image"] = idImage })
}

// WithIDProduct id_product获取 
func (obj *_EgImageMgr) WithIDProduct(idProduct uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product"] = idProduct })
}

// WithPosition position获取 
func (obj *_EgImageMgr) WithPosition(position uint16) Option {
	return optionFunc(func(o *options) { o.query["position"] = position })
}

// WithCover cover获取 
func (obj *_EgImageMgr) WithCover(cover bool) Option {
	return optionFunc(func(o *options) { o.query["cover"] = cover })
}


// GetByOption 功能选项模式获取
func (obj *_EgImageMgr) GetByOption(opts ...Option) (result EgImage, err error) {
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
func (obj *_EgImageMgr) GetByOptions(opts ...Option) (results []*EgImage, err error) {
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


// GetFromIDImage 通过id_image获取内容  
func (obj *_EgImageMgr)  GetFromIDImage(idImage uint32) (result EgImage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_image = ?", idImage).Find(&result).Error
	
	return
}

// GetBatchFromIDImage 批量唯一主键查找 
func (obj *_EgImageMgr) GetBatchFromIDImage(idImages []uint32) (results []*EgImage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_image IN (?)", idImages).Find(&results).Error
	
	return
}
 
// GetFromIDProduct 通过id_product获取内容  
func (obj *_EgImageMgr) GetFromIDProduct(idProduct uint32) (results []*EgImage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ?", idProduct).Find(&results).Error
	
	return
}

// GetBatchFromIDProduct 批量唯一主键查找 
func (obj *_EgImageMgr) GetBatchFromIDProduct(idProducts []uint32) (results []*EgImage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product IN (?)", idProducts).Find(&results).Error
	
	return
}
 
// GetFromPosition 通过position获取内容  
func (obj *_EgImageMgr) GetFromPosition(position uint16) (results []*EgImage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("position = ?", position).Find(&results).Error
	
	return
}

// GetBatchFromPosition 批量唯一主键查找 
func (obj *_EgImageMgr) GetBatchFromPosition(positions []uint16) (results []*EgImage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("position IN (?)", positions).Find(&results).Error
	
	return
}
 
// GetFromCover 通过cover获取内容  
func (obj *_EgImageMgr) GetFromCover(cover bool) (results []*EgImage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cover = ?", cover).Find(&results).Error
	
	return
}

// GetBatchFromCover 批量唯一主键查找 
func (obj *_EgImageMgr) GetBatchFromCover(covers []bool) (results []*EgImage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cover IN (?)", covers).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgImageMgr) FetchByPrimaryKey(idImage uint32 ) (result EgImage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_image = ?", idImage).Find(&result).Error
	
	return
}
 
 // FetchUniqueIndexByIDxProductImage primay or index 获取唯一内容
 func (obj *_EgImageMgr) FetchUniqueIndexByIDxProductImage(idImage uint32 ,idProduct uint32 ,cover bool ) (result EgImage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_image = ? AND id_product = ? AND cover = ?", idImage , idProduct , cover).Find(&result).Error
	
	return
}
 
 // FetchUniqueIndexByIDProductCover primay or index 获取唯一内容
 func (obj *_EgImageMgr) FetchUniqueIndexByIDProductCover(idProduct uint32 ,cover bool ) (result EgImage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ? AND cover = ?", idProduct , cover).Find(&result).Error
	
	return
}
 

 
 // FetchIndexByImageProduct  获取多个内容
 func (obj *_EgImageMgr) FetchIndexByImageProduct(idProduct uint32 ) (results []*EgImage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ?", idProduct).Find(&results).Error
	
	return
}
 

	

