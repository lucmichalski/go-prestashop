package	model	
import (	
"context"	
"gorm.io/gorm"	
"fmt"	
)	

type _EgAttachmentMgr struct {
	*_BaseMgr
}

// EgAttachmentMgr open func
func EgAttachmentMgr(db *gorm.DB) *_EgAttachmentMgr {
	if db == nil {
		panic(fmt.Errorf("EgAttachmentMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgAttachmentMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_attachment"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgAttachmentMgr) GetTableName() string {
	return "eg_attachment"
}

// Get 获取
func (obj *_EgAttachmentMgr) Get() (result EgAttachment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgAttachmentMgr) Gets() (results []*EgAttachment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDAttachment id_attachment获取 
func (obj *_EgAttachmentMgr) WithIDAttachment(idAttachment uint32) Option {
	return optionFunc(func(o *options) { o.query["id_attachment"] = idAttachment })
}

// WithFile file获取 
func (obj *_EgAttachmentMgr) WithFile(file string) Option {
	return optionFunc(func(o *options) { o.query["file"] = file })
}

// WithFileName file_name获取 
func (obj *_EgAttachmentMgr) WithFileName(fileName string) Option {
	return optionFunc(func(o *options) { o.query["file_name"] = fileName })
}

// WithFileSize file_size获取 
func (obj *_EgAttachmentMgr) WithFileSize(fileSize uint64) Option {
	return optionFunc(func(o *options) { o.query["file_size"] = fileSize })
}

// WithMime mime获取 
func (obj *_EgAttachmentMgr) WithMime(mime string) Option {
	return optionFunc(func(o *options) { o.query["mime"] = mime })
}


// GetByOption 功能选项模式获取
func (obj *_EgAttachmentMgr) GetByOption(opts ...Option) (result EgAttachment, err error) {
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
func (obj *_EgAttachmentMgr) GetByOptions(opts ...Option) (results []*EgAttachment, err error) {
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
func (obj *_EgAttachmentMgr)  GetFromIDAttachment(idAttachment uint32) (result EgAttachment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attachment = ?", idAttachment).Find(&result).Error
	
	return
}

// GetBatchFromIDAttachment 批量唯一主键查找 
func (obj *_EgAttachmentMgr) GetBatchFromIDAttachment(idAttachments []uint32) (results []*EgAttachment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attachment IN (?)", idAttachments).Find(&results).Error
	
	return
}
 
// GetFromFile 通过file获取内容  
func (obj *_EgAttachmentMgr) GetFromFile(file string) (results []*EgAttachment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("file = ?", file).Find(&results).Error
	
	return
}

// GetBatchFromFile 批量唯一主键查找 
func (obj *_EgAttachmentMgr) GetBatchFromFile(files []string) (results []*EgAttachment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("file IN (?)", files).Find(&results).Error
	
	return
}
 
// GetFromFileName 通过file_name获取内容  
func (obj *_EgAttachmentMgr) GetFromFileName(fileName string) (results []*EgAttachment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("file_name = ?", fileName).Find(&results).Error
	
	return
}

// GetBatchFromFileName 批量唯一主键查找 
func (obj *_EgAttachmentMgr) GetBatchFromFileName(fileNames []string) (results []*EgAttachment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("file_name IN (?)", fileNames).Find(&results).Error
	
	return
}
 
// GetFromFileSize 通过file_size获取内容  
func (obj *_EgAttachmentMgr) GetFromFileSize(fileSize uint64) (results []*EgAttachment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("file_size = ?", fileSize).Find(&results).Error
	
	return
}

// GetBatchFromFileSize 批量唯一主键查找 
func (obj *_EgAttachmentMgr) GetBatchFromFileSize(fileSizes []uint64) (results []*EgAttachment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("file_size IN (?)", fileSizes).Find(&results).Error
	
	return
}
 
// GetFromMime 通过mime获取内容  
func (obj *_EgAttachmentMgr) GetFromMime(mime string) (results []*EgAttachment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("mime = ?", mime).Find(&results).Error
	
	return
}

// GetBatchFromMime 批量唯一主键查找 
func (obj *_EgAttachmentMgr) GetBatchFromMime(mimes []string) (results []*EgAttachment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("mime IN (?)", mimes).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgAttachmentMgr) FetchByPrimaryKey(idAttachment uint32 ) (result EgAttachment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attachment = ?", idAttachment).Find(&result).Error
	
	return
}
 

 

	

