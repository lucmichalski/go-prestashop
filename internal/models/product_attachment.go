package	model	
import (	
"context"	
"gorm.io/gorm"	
"fmt"	
)	

type _EgProductAttachmentMgr struct {
	*_BaseMgr
}

// EgProductAttachmentMgr open func
func EgProductAttachmentMgr(db *gorm.DB) *_EgProductAttachmentMgr {
	if db == nil {
		panic(fmt.Errorf("EgProductAttachmentMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgProductAttachmentMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_product_attachment"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgProductAttachmentMgr) GetTableName() string {
	return "eg_product_attachment"
}

// Get 获取
func (obj *_EgProductAttachmentMgr) Get() (result EgProductAttachment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgProductAttachmentMgr) Gets() (results []*EgProductAttachment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDProduct id_product获取 
func (obj *_EgProductAttachmentMgr) WithIDProduct(idProduct uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product"] = idProduct })
}

// WithIDAttachment id_attachment获取 
func (obj *_EgProductAttachmentMgr) WithIDAttachment(idAttachment uint32) Option {
	return optionFunc(func(o *options) { o.query["id_attachment"] = idAttachment })
}


// GetByOption 功能选项模式获取
func (obj *_EgProductAttachmentMgr) GetByOption(opts ...Option) (result EgProductAttachment, err error) {
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
func (obj *_EgProductAttachmentMgr) GetByOptions(opts ...Option) (results []*EgProductAttachment, err error) {
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
func (obj *_EgProductAttachmentMgr) GetFromIDProduct(idProduct uint32) (results []*EgProductAttachment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ?", idProduct).Find(&results).Error
	
	return
}

// GetBatchFromIDProduct 批量唯一主键查找 
func (obj *_EgProductAttachmentMgr) GetBatchFromIDProduct(idProducts []uint32) (results []*EgProductAttachment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product IN (?)", idProducts).Find(&results).Error
	
	return
}
 
// GetFromIDAttachment 通过id_attachment获取内容  
func (obj *_EgProductAttachmentMgr) GetFromIDAttachment(idAttachment uint32) (results []*EgProductAttachment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attachment = ?", idAttachment).Find(&results).Error
	
	return
}

// GetBatchFromIDAttachment 批量唯一主键查找 
func (obj *_EgProductAttachmentMgr) GetBatchFromIDAttachment(idAttachments []uint32) (results []*EgProductAttachment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attachment IN (?)", idAttachments).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgProductAttachmentMgr) FetchByPrimaryKey(idProduct uint32 ,idAttachment uint32 ) (result EgProductAttachment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ? AND id_attachment = ?", idProduct , idAttachment).Find(&result).Error
	
	return
}
 

 

	

