package	model	
import (	
"context"	
"gorm.io/gorm"	
"fmt"	
)	

type _EgProfileMgr struct {
	*_BaseMgr
}

// EgProfileMgr open func
func EgProfileMgr(db *gorm.DB) *_EgProfileMgr {
	if db == nil {
		panic(fmt.Errorf("EgProfileMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgProfileMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_profile"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgProfileMgr) GetTableName() string {
	return "eg_profile"
}

// Get 获取
func (obj *_EgProfileMgr) Get() (result EgProfile, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgProfileMgr) Gets() (results []*EgProfile, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDProfile id_profile获取 
func (obj *_EgProfileMgr) WithIDProfile(idProfile uint32) Option {
	return optionFunc(func(o *options) { o.query["id_profile"] = idProfile })
}


// GetByOption 功能选项模式获取
func (obj *_EgProfileMgr) GetByOption(opts ...Option) (result EgProfile, err error) {
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
func (obj *_EgProfileMgr) GetByOptions(opts ...Option) (results []*EgProfile, err error) {
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


// GetFromIDProfile 通过id_profile获取内容  
func (obj *_EgProfileMgr)  GetFromIDProfile(idProfile uint32) (result EgProfile, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_profile = ?", idProfile).Find(&result).Error
	
	return
}

// GetBatchFromIDProfile 批量唯一主键查找 
func (obj *_EgProfileMgr) GetBatchFromIDProfile(idProfiles []uint32) (results []*EgProfile, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_profile IN (?)", idProfiles).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgProfileMgr) FetchByPrimaryKey(idProfile uint32 ) (result EgProfile, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_profile = ?", idProfile).Find(&result).Error
	
	return
}
 

 

	

