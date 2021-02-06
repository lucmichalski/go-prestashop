package	model	
import (	
"context"	
"gorm.io/gorm"	
"time"	
"fmt"	
)	

type _EgMailMgr struct {
	*_BaseMgr
}

// EgMailMgr open func
func EgMailMgr(db *gorm.DB) *_EgMailMgr {
	if db == nil {
		panic(fmt.Errorf("EgMailMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgMailMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_mail"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgMailMgr) GetTableName() string {
	return "eg_mail"
}

// Get 获取
func (obj *_EgMailMgr) Get() (result EgMail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgMailMgr) Gets() (results []*EgMail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDMail id_mail获取 
func (obj *_EgMailMgr) WithIDMail(idMail uint32) Option {
	return optionFunc(func(o *options) { o.query["id_mail"] = idMail })
}

// WithRecipient recipient获取 
func (obj *_EgMailMgr) WithRecipient(recipient string) Option {
	return optionFunc(func(o *options) { o.query["recipient"] = recipient })
}

// WithTemplate template获取 
func (obj *_EgMailMgr) WithTemplate(template string) Option {
	return optionFunc(func(o *options) { o.query["template"] = template })
}

// WithSubject subject获取 
func (obj *_EgMailMgr) WithSubject(subject string) Option {
	return optionFunc(func(o *options) { o.query["subject"] = subject })
}

// WithIDLang id_lang获取 
func (obj *_EgMailMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

// WithDateAdd date_add获取 
func (obj *_EgMailMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}


// GetByOption 功能选项模式获取
func (obj *_EgMailMgr) GetByOption(opts ...Option) (result EgMail, err error) {
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
func (obj *_EgMailMgr) GetByOptions(opts ...Option) (results []*EgMail, err error) {
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


// GetFromIDMail 通过id_mail获取内容  
func (obj *_EgMailMgr)  GetFromIDMail(idMail uint32) (result EgMail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_mail = ?", idMail).Find(&result).Error
	
	return
}

// GetBatchFromIDMail 批量唯一主键查找 
func (obj *_EgMailMgr) GetBatchFromIDMail(idMails []uint32) (results []*EgMail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_mail IN (?)", idMails).Find(&results).Error
	
	return
}
 
// GetFromRecipient 通过recipient获取内容  
func (obj *_EgMailMgr) GetFromRecipient(recipient string) (results []*EgMail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("recipient = ?", recipient).Find(&results).Error
	
	return
}

// GetBatchFromRecipient 批量唯一主键查找 
func (obj *_EgMailMgr) GetBatchFromRecipient(recipients []string) (results []*EgMail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("recipient IN (?)", recipients).Find(&results).Error
	
	return
}
 
// GetFromTemplate 通过template获取内容  
func (obj *_EgMailMgr) GetFromTemplate(template string) (results []*EgMail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("template = ?", template).Find(&results).Error
	
	return
}

// GetBatchFromTemplate 批量唯一主键查找 
func (obj *_EgMailMgr) GetBatchFromTemplate(templates []string) (results []*EgMail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("template IN (?)", templates).Find(&results).Error
	
	return
}
 
// GetFromSubject 通过subject获取内容  
func (obj *_EgMailMgr) GetFromSubject(subject string) (results []*EgMail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("subject = ?", subject).Find(&results).Error
	
	return
}

// GetBatchFromSubject 批量唯一主键查找 
func (obj *_EgMailMgr) GetBatchFromSubject(subjects []string) (results []*EgMail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("subject IN (?)", subjects).Find(&results).Error
	
	return
}
 
// GetFromIDLang 通过id_lang获取内容  
func (obj *_EgMailMgr) GetFromIDLang(idLang uint32) (results []*EgMail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error
	
	return
}

// GetBatchFromIDLang 批量唯一主键查找 
func (obj *_EgMailMgr) GetBatchFromIDLang(idLangs []uint32) (results []*EgMail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error
	
	return
}
 
// GetFromDateAdd 通过date_add获取内容  
func (obj *_EgMailMgr) GetFromDateAdd(dateAdd time.Time) (results []*EgMail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error
	
	return
}

// GetBatchFromDateAdd 批量唯一主键查找 
func (obj *_EgMailMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*EgMail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgMailMgr) FetchByPrimaryKey(idMail uint32 ) (result EgMail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_mail = ?", idMail).Find(&result).Error
	
	return
}
 

 
 // FetchIndexByRecipient  获取多个内容
 func (obj *_EgMailMgr) FetchIndexByRecipient(recipient string ) (results []*EgMail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("recipient = ?", recipient).Find(&results).Error
	
	return
}
 

	

