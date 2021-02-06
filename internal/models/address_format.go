package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgAddressFormatMgr struct {
	*_BaseMgr
}

// EgAddressFormatMgr open func
func EgAddressFormatMgr(db *gorm.DB) *_EgAddressFormatMgr {
	if db == nil {
		panic(fmt.Errorf("EgAddressFormatMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgAddressFormatMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_address_format"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgAddressFormatMgr) GetTableName() string {
	return "eg_address_format"
}

// Get 获取
func (obj *_EgAddressFormatMgr) Get() (result EgAddressFormat, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgAddressFormatMgr) Gets() (results []*EgAddressFormat, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDCountry id_country获取 
func (obj *_EgAddressFormatMgr) WithIDCountry(idCountry uint32) Option {
	return optionFunc(func(o *options) { o.query["id_country"] = idCountry })
}

// WithFormat format获取 
func (obj *_EgAddressFormatMgr) WithFormat(format string) Option {
	return optionFunc(func(o *options) { o.query["format"] = format })
}


// GetByOption 功能选项模式获取
func (obj *_EgAddressFormatMgr) GetByOption(opts ...Option) (result EgAddressFormat, err error) {
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
func (obj *_EgAddressFormatMgr) GetByOptions(opts ...Option) (results []*EgAddressFormat, err error) {
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


// GetFromIDCountry 通过id_country获取内容  
func (obj *_EgAddressFormatMgr)  GetFromIDCountry(idCountry uint32) (result EgAddressFormat, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_country = ?", idCountry).Find(&result).Error
	
	return
}

// GetBatchFromIDCountry 批量唯一主键查找 
func (obj *_EgAddressFormatMgr) GetBatchFromIDCountry(idCountrys []uint32) (results []*EgAddressFormat, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_country IN (?)", idCountrys).Find(&results).Error
	
	return
}
 
// GetFromFormat 通过format获取内容  
func (obj *_EgAddressFormatMgr) GetFromFormat(format string) (results []*EgAddressFormat, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("format = ?", format).Find(&results).Error
	
	return
}

// GetBatchFromFormat 批量唯一主键查找 
func (obj *_EgAddressFormatMgr) GetBatchFromFormat(formats []string) (results []*EgAddressFormat, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("format IN (?)", formats).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgAddressFormatMgr) FetchByPrimaryKey(idCountry uint32 ) (result EgAddressFormat, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_country = ?", idCountry).Find(&result).Error
	
	return
}
 

 

	

