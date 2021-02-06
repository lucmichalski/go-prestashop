package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgHomesliderSlidesLangMgr struct {
	*_BaseMgr
}

// EgHomesliderSlidesLangMgr open func
func EgHomesliderSlidesLangMgr(db *gorm.DB) *_EgHomesliderSlidesLangMgr {
	if db == nil {
		panic(fmt.Errorf("EgHomesliderSlidesLangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgHomesliderSlidesLangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_homeslider_slides_lang"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgHomesliderSlidesLangMgr) GetTableName() string {
	return "eg_homeslider_slides_lang"
}

// Get 获取
func (obj *_EgHomesliderSlidesLangMgr) Get() (result EgHomesliderSlidesLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgHomesliderSlidesLangMgr) Gets() (results []*EgHomesliderSlidesLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDHomesliderSlides id_homeslider_slides获取 
func (obj *_EgHomesliderSlidesLangMgr) WithIDHomesliderSlides(idHomesliderSlides uint32) Option {
	return optionFunc(func(o *options) { o.query["id_homeslider_slides"] = idHomesliderSlides })
}

// WithIDLang id_lang获取 
func (obj *_EgHomesliderSlidesLangMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

// WithTitle title获取 
func (obj *_EgHomesliderSlidesLangMgr) WithTitle(title string) Option {
	return optionFunc(func(o *options) { o.query["title"] = title })
}

// WithDescription description获取 
func (obj *_EgHomesliderSlidesLangMgr) WithDescription(description string) Option {
	return optionFunc(func(o *options) { o.query["description"] = description })
}

// WithLegend legend获取 
func (obj *_EgHomesliderSlidesLangMgr) WithLegend(legend string) Option {
	return optionFunc(func(o *options) { o.query["legend"] = legend })
}

// WithURL url获取 
func (obj *_EgHomesliderSlidesLangMgr) WithURL(url string) Option {
	return optionFunc(func(o *options) { o.query["url"] = url })
}

// WithImage image获取 
func (obj *_EgHomesliderSlidesLangMgr) WithImage(image string) Option {
	return optionFunc(func(o *options) { o.query["image"] = image })
}


// GetByOption 功能选项模式获取
func (obj *_EgHomesliderSlidesLangMgr) GetByOption(opts ...Option) (result EgHomesliderSlidesLang, err error) {
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
func (obj *_EgHomesliderSlidesLangMgr) GetByOptions(opts ...Option) (results []*EgHomesliderSlidesLang, err error) {
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
func (obj *_EgHomesliderSlidesLangMgr) GetFromIDHomesliderSlides(idHomesliderSlides uint32) (results []*EgHomesliderSlidesLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_homeslider_slides = ?", idHomesliderSlides).Find(&results).Error
	
	return
}

// GetBatchFromIDHomesliderSlides 批量唯一主键查找 
func (obj *_EgHomesliderSlidesLangMgr) GetBatchFromIDHomesliderSlides(idHomesliderSlidess []uint32) (results []*EgHomesliderSlidesLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_homeslider_slides IN (?)", idHomesliderSlidess).Find(&results).Error
	
	return
}
 
// GetFromIDLang 通过id_lang获取内容  
func (obj *_EgHomesliderSlidesLangMgr) GetFromIDLang(idLang uint32) (results []*EgHomesliderSlidesLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error
	
	return
}

// GetBatchFromIDLang 批量唯一主键查找 
func (obj *_EgHomesliderSlidesLangMgr) GetBatchFromIDLang(idLangs []uint32) (results []*EgHomesliderSlidesLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error
	
	return
}
 
// GetFromTitle 通过title获取内容  
func (obj *_EgHomesliderSlidesLangMgr) GetFromTitle(title string) (results []*EgHomesliderSlidesLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("title = ?", title).Find(&results).Error
	
	return
}

// GetBatchFromTitle 批量唯一主键查找 
func (obj *_EgHomesliderSlidesLangMgr) GetBatchFromTitle(titles []string) (results []*EgHomesliderSlidesLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("title IN (?)", titles).Find(&results).Error
	
	return
}
 
// GetFromDescription 通过description获取内容  
func (obj *_EgHomesliderSlidesLangMgr) GetFromDescription(description string) (results []*EgHomesliderSlidesLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("description = ?", description).Find(&results).Error
	
	return
}

// GetBatchFromDescription 批量唯一主键查找 
func (obj *_EgHomesliderSlidesLangMgr) GetBatchFromDescription(descriptions []string) (results []*EgHomesliderSlidesLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("description IN (?)", descriptions).Find(&results).Error
	
	return
}
 
// GetFromLegend 通过legend获取内容  
func (obj *_EgHomesliderSlidesLangMgr) GetFromLegend(legend string) (results []*EgHomesliderSlidesLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("legend = ?", legend).Find(&results).Error
	
	return
}

// GetBatchFromLegend 批量唯一主键查找 
func (obj *_EgHomesliderSlidesLangMgr) GetBatchFromLegend(legends []string) (results []*EgHomesliderSlidesLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("legend IN (?)", legends).Find(&results).Error
	
	return
}
 
// GetFromURL 通过url获取内容  
func (obj *_EgHomesliderSlidesLangMgr) GetFromURL(url string) (results []*EgHomesliderSlidesLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("url = ?", url).Find(&results).Error
	
	return
}

// GetBatchFromURL 批量唯一主键查找 
func (obj *_EgHomesliderSlidesLangMgr) GetBatchFromURL(urls []string) (results []*EgHomesliderSlidesLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("url IN (?)", urls).Find(&results).Error
	
	return
}
 
// GetFromImage 通过image获取内容  
func (obj *_EgHomesliderSlidesLangMgr) GetFromImage(image string) (results []*EgHomesliderSlidesLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("image = ?", image).Find(&results).Error
	
	return
}

// GetBatchFromImage 批量唯一主键查找 
func (obj *_EgHomesliderSlidesLangMgr) GetBatchFromImage(images []string) (results []*EgHomesliderSlidesLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("image IN (?)", images).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgHomesliderSlidesLangMgr) FetchByPrimaryKey(idHomesliderSlides uint32 ,idLang uint32 ) (result EgHomesliderSlidesLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_homeslider_slides = ? AND id_lang = ?", idHomesliderSlides , idLang).Find(&result).Error
	
	return
}
 

 

	

