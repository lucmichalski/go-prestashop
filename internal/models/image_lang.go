package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgImageLangMgr struct {
	*_BaseMgr
}

// EgImageLangMgr open func
func EgImageLangMgr(db *gorm.DB) *_EgImageLangMgr {
	if db == nil {
		panic(fmt.Errorf("EgImageLangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgImageLangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_image_lang"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgImageLangMgr) GetTableName() string {
	return "eg_image_lang"
}

// Get 获取
func (obj *_EgImageLangMgr) Get() (result EgImageLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgImageLangMgr) Gets() (results []*EgImageLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDImage id_image获取 
func (obj *_EgImageLangMgr) WithIDImage(idImage uint32) Option {
	return optionFunc(func(o *options) { o.query["id_image"] = idImage })
}

// WithIDLang id_lang获取 
func (obj *_EgImageLangMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

// WithLegend legend获取 
func (obj *_EgImageLangMgr) WithLegend(legend string) Option {
	return optionFunc(func(o *options) { o.query["legend"] = legend })
}


// GetByOption 功能选项模式获取
func (obj *_EgImageLangMgr) GetByOption(opts ...Option) (result EgImageLang, err error) {
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
func (obj *_EgImageLangMgr) GetByOptions(opts ...Option) (results []*EgImageLang, err error) {
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
func (obj *_EgImageLangMgr) GetFromIDImage(idImage uint32) (results []*EgImageLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_image = ?", idImage).Find(&results).Error
	
	return
}

// GetBatchFromIDImage 批量唯一主键查找 
func (obj *_EgImageLangMgr) GetBatchFromIDImage(idImages []uint32) (results []*EgImageLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_image IN (?)", idImages).Find(&results).Error
	
	return
}
 
// GetFromIDLang 通过id_lang获取内容  
func (obj *_EgImageLangMgr) GetFromIDLang(idLang uint32) (results []*EgImageLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error
	
	return
}

// GetBatchFromIDLang 批量唯一主键查找 
func (obj *_EgImageLangMgr) GetBatchFromIDLang(idLangs []uint32) (results []*EgImageLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error
	
	return
}
 
// GetFromLegend 通过legend获取内容  
func (obj *_EgImageLangMgr) GetFromLegend(legend string) (results []*EgImageLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("legend = ?", legend).Find(&results).Error
	
	return
}

// GetBatchFromLegend 批量唯一主键查找 
func (obj *_EgImageLangMgr) GetBatchFromLegend(legends []string) (results []*EgImageLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("legend IN (?)", legends).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgImageLangMgr) FetchByPrimaryKey(idImage uint32 ,idLang uint32 ) (result EgImageLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_image = ? AND id_lang = ?", idImage , idLang).Find(&result).Error
	
	return
}
 

 
 // FetchIndexByIDImage  获取多个内容
 func (obj *_EgImageLangMgr) FetchIndexByIDImage(idImage uint32 ) (results []*EgImageLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_image = ?", idImage).Find(&results).Error
	
	return
}
 

	

