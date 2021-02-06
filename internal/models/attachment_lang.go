package	model	
import (	
"context"	
"gorm.io/gorm"	
"fmt"	
)	

type _EgAttachmentLangMgr struct {
	*_BaseMgr
}

// EgAttachmentLangMgr open func
func EgAttachmentLangMgr(db *gorm.DB) *_EgAttachmentLangMgr {
	if db == nil {
		panic(fmt.Errorf("EgAttachmentLangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgAttachmentLangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_attachment_lang"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgAttachmentLangMgr) GetTableName() string {
	return "eg_attachment_lang"
}

// Get 获取
func (obj *_EgAttachmentLangMgr) Get() (result EgAttachmentLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgAttachmentLangMgr) Gets() (results []*EgAttachmentLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDAttachment id_attachment获取 
func (obj *_EgAttachmentLangMgr) WithIDAttachment(idAttachment uint32) Option {
	return optionFunc(func(o *options) { o.query["id_attachment"] = idAttachment })
}

// WithIDLang id_lang获取 
func (obj *_EgAttachmentLangMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

// WithName name获取 
func (obj *_EgAttachmentLangMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithDescription description获取 
func (obj *_EgAttachmentLangMgr) WithDescription(description string) Option {
	return optionFunc(func(o *options) { o.query["description"] = description })
}


// GetByOption 功能选项模式获取
func (obj *_EgAttachmentLangMgr) GetByOption(opts ...Option) (result EgAttachmentLang, err error) {
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
func (obj *_EgAttachmentLangMgr) GetByOptions(opts ...Option) (results []*EgAttachmentLang, err error) {
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


// GetFromIDAttachment 通过id_attachment获取内容  
func (obj *_EgAttachmentLangMgr) GetFromIDAttachment(idAttachment uint32) (results []*EgAttachmentLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attachment = ?", idAttachment).Find(&results).Error
	
	return
}

// GetBatchFromIDAttachment 批量唯一主键查找 
func (obj *_EgAttachmentLangMgr) GetBatchFromIDAttachment(idAttachments []uint32) (results []*EgAttachmentLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attachment IN (?)", idAttachments).Find(&results).Error
	
	return
}
 
// GetFromIDLang 通过id_lang获取内容  
func (obj *_EgAttachmentLangMgr) GetFromIDLang(idLang uint32) (results []*EgAttachmentLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error
	
	return
}

// GetBatchFromIDLang 批量唯一主键查找 
func (obj *_EgAttachmentLangMgr) GetBatchFromIDLang(idLangs []uint32) (results []*EgAttachmentLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error
	
	return
}
 
// GetFromName 通过name获取内容  
func (obj *_EgAttachmentLangMgr) GetFromName(name string) (results []*EgAttachmentLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error
	
	return
}

// GetBatchFromName 批量唯一主键查找 
func (obj *_EgAttachmentLangMgr) GetBatchFromName(names []string) (results []*EgAttachmentLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error
	
	return
}
 
// GetFromDescription 通过description获取内容  
func (obj *_EgAttachmentLangMgr) GetFromDescription(description string) (results []*EgAttachmentLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("description = ?", description).Find(&results).Error
	
	return
}

// GetBatchFromDescription 批量唯一主键查找 
func (obj *_EgAttachmentLangMgr) GetBatchFromDescription(descriptions []string) (results []*EgAttachmentLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("description IN (?)", descriptions).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgAttachmentLangMgr) FetchByPrimaryKey(idAttachment uint32 ,idLang uint32 ) (result EgAttachmentLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attachment = ? AND id_lang = ?", idAttachment , idLang).Find(&result).Error
	
	return
}
 

 

	

