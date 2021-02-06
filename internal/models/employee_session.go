package	model	
import (	
"context"	
"gorm.io/gorm"	
"fmt"	
)	

type _EgEmployeeSessionMgr struct {
	*_BaseMgr
}

// EgEmployeeSessionMgr open func
func EgEmployeeSessionMgr(db *gorm.DB) *_EgEmployeeSessionMgr {
	if db == nil {
		panic(fmt.Errorf("EgEmployeeSessionMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgEmployeeSessionMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_employee_session"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgEmployeeSessionMgr) GetTableName() string {
	return "eg_employee_session"
}

// Get 获取
func (obj *_EgEmployeeSessionMgr) Get() (result EgEmployeeSession, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgEmployeeSessionMgr) Gets() (results []*EgEmployeeSession, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDEmployeeSession id_employee_session获取 
func (obj *_EgEmployeeSessionMgr) WithIDEmployeeSession(idEmployeeSession uint32) Option {
	return optionFunc(func(o *options) { o.query["id_employee_session"] = idEmployeeSession })
}

// WithIDEmployee id_employee获取 
func (obj *_EgEmployeeSessionMgr) WithIDEmployee(idEmployee uint32) Option {
	return optionFunc(func(o *options) { o.query["id_employee"] = idEmployee })
}

// WithToken token获取 
func (obj *_EgEmployeeSessionMgr) WithToken(token string) Option {
	return optionFunc(func(o *options) { o.query["token"] = token })
}


// GetByOption 功能选项模式获取
func (obj *_EgEmployeeSessionMgr) GetByOption(opts ...Option) (result EgEmployeeSession, err error) {
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
func (obj *_EgEmployeeSessionMgr) GetByOptions(opts ...Option) (results []*EgEmployeeSession, err error) {
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


// GetFromIDEmployeeSession 通过id_employee_session获取内容  
func (obj *_EgEmployeeSessionMgr)  GetFromIDEmployeeSession(idEmployeeSession uint32) (result EgEmployeeSession, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee_session = ?", idEmployeeSession).Find(&result).Error
	
	return
}

// GetBatchFromIDEmployeeSession 批量唯一主键查找 
func (obj *_EgEmployeeSessionMgr) GetBatchFromIDEmployeeSession(idEmployeeSessions []uint32) (results []*EgEmployeeSession, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee_session IN (?)", idEmployeeSessions).Find(&results).Error
	
	return
}
 
// GetFromIDEmployee 通过id_employee获取内容  
func (obj *_EgEmployeeSessionMgr) GetFromIDEmployee(idEmployee uint32) (results []*EgEmployeeSession, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee = ?", idEmployee).Find(&results).Error
	
	return
}

// GetBatchFromIDEmployee 批量唯一主键查找 
func (obj *_EgEmployeeSessionMgr) GetBatchFromIDEmployee(idEmployees []uint32) (results []*EgEmployeeSession, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee IN (?)", idEmployees).Find(&results).Error
	
	return
}
 
// GetFromToken 通过token获取内容  
func (obj *_EgEmployeeSessionMgr) GetFromToken(token string) (results []*EgEmployeeSession, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("token = ?", token).Find(&results).Error
	
	return
}

// GetBatchFromToken 批量唯一主键查找 
func (obj *_EgEmployeeSessionMgr) GetBatchFromToken(tokens []string) (results []*EgEmployeeSession, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("token IN (?)", tokens).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgEmployeeSessionMgr) FetchByPrimaryKey(idEmployeeSession uint32 ) (result EgEmployeeSession, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee_session = ?", idEmployeeSession).Find(&result).Error
	
	return
}
 

 

	

