package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgHomesliderSlidesMgr struct {
	*_BaseMgr
}

// EgHomesliderSlidesMgr open func
func EgHomesliderSlidesMgr(db *gorm.DB) *_EgHomesliderSlidesMgr {
	if db == nil {
		panic(fmt.Errorf("EgHomesliderSlidesMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgHomesliderSlidesMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_homeslider_slides"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgHomesliderSlidesMgr) GetTableName() string {
	return "eg_homeslider_slides"
}

// Get 获取
func (obj *_EgHomesliderSlidesMgr) Get() (result EgHomesliderSlides, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgHomesliderSlidesMgr) Gets() (results []*EgHomesliderSlides, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDHomesliderSlides id_homeslider_slides获取 
func (obj *_EgHomesliderSlidesMgr) WithIDHomesliderSlides(idHomesliderSlides uint32) Option {
	return optionFunc(func(o *options) { o.query["id_homeslider_slides"] = idHomesliderSlides })
}

// WithPosition position获取 
func (obj *_EgHomesliderSlidesMgr) WithPosition(position uint32) Option {
	return optionFunc(func(o *options) { o.query["position"] = position })
}

// WithActive active获取 
func (obj *_EgHomesliderSlidesMgr) WithActive(active bool) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}


// GetByOption 功能选项模式获取
func (obj *_EgHomesliderSlidesMgr) GetByOption(opts ...Option) (result EgHomesliderSlides, err error) {
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
func (obj *_EgHomesliderSlidesMgr) GetByOptions(opts ...Option) (results []*EgHomesliderSlides, err error) {
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
func (obj *_EgHomesliderSlidesMgr)  GetFromIDHomesliderSlides(idHomesliderSlides uint32) (result EgHomesliderSlides, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_homeslider_slides = ?", idHomesliderSlides).Find(&result).Error
	
	return
}

// GetBatchFromIDHomesliderSlides 批量唯一主键查找 
func (obj *_EgHomesliderSlidesMgr) GetBatchFromIDHomesliderSlides(idHomesliderSlidess []uint32) (results []*EgHomesliderSlides, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_homeslider_slides IN (?)", idHomesliderSlidess).Find(&results).Error
	
	return
}
 
// GetFromPosition 通过position获取内容  
func (obj *_EgHomesliderSlidesMgr) GetFromPosition(position uint32) (results []*EgHomesliderSlides, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("position = ?", position).Find(&results).Error
	
	return
}

// GetBatchFromPosition 批量唯一主键查找 
func (obj *_EgHomesliderSlidesMgr) GetBatchFromPosition(positions []uint32) (results []*EgHomesliderSlides, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("position IN (?)", positions).Find(&results).Error
	
	return
}
 
// GetFromActive 通过active获取内容  
func (obj *_EgHomesliderSlidesMgr) GetFromActive(active bool) (results []*EgHomesliderSlides, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active = ?", active).Find(&results).Error
	
	return
}

// GetBatchFromActive 批量唯一主键查找 
func (obj *_EgHomesliderSlidesMgr) GetBatchFromActive(actives []bool) (results []*EgHomesliderSlides, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active IN (?)", actives).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgHomesliderSlidesMgr) FetchByPrimaryKey(idHomesliderSlides uint32 ) (result EgHomesliderSlides, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_homeslider_slides = ?", idHomesliderSlides).Find(&result).Error
	
	return
}
 

 

	

