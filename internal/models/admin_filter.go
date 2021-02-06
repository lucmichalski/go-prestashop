package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgAdminFilterMgr struct {
	*_BaseMgr
}

// EgAdminFilterMgr open func
func EgAdminFilterMgr(db *gorm.DB) *_EgAdminFilterMgr {
	if db == nil {
		panic(fmt.Errorf("EgAdminFilterMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgAdminFilterMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_admin_filter"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgAdminFilterMgr) GetTableName() string {
	return "eg_admin_filter"
}

// Get 获取
func (obj *_EgAdminFilterMgr) Get() (result EgAdminFilter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgAdminFilterMgr) Gets() (results []*EgAdminFilter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 
func (obj *_EgAdminFilterMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithEmployee employee获取 
func (obj *_EgAdminFilterMgr) WithEmployee(employee int) Option {
	return optionFunc(func(o *options) { o.query["employee"] = employee })
}

// WithShop shop获取 
func (obj *_EgAdminFilterMgr) WithShop(shop int) Option {
	return optionFunc(func(o *options) { o.query["shop"] = shop })
}

// WithController controller获取 
func (obj *_EgAdminFilterMgr) WithController(controller string) Option {
	return optionFunc(func(o *options) { o.query["controller"] = controller })
}

// WithAction action获取 
func (obj *_EgAdminFilterMgr) WithAction(action string) Option {
	return optionFunc(func(o *options) { o.query["action"] = action })
}

// WithFilter filter获取 
func (obj *_EgAdminFilterMgr) WithFilter(filter string) Option {
	return optionFunc(func(o *options) { o.query["filter"] = filter })
}

// WithFilterID filter_id获取 
func (obj *_EgAdminFilterMgr) WithFilterID(filterID string) Option {
	return optionFunc(func(o *options) { o.query["filter_id"] = filterID })
}


// GetByOption 功能选项模式获取
func (obj *_EgAdminFilterMgr) GetByOption(opts ...Option) (result EgAdminFilter, err error) {
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
func (obj *_EgAdminFilterMgr) GetByOptions(opts ...Option) (results []*EgAdminFilter, err error) {
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


// GetFromID 通过id获取内容  
func (obj *_EgAdminFilterMgr)  GetFromID(id int) (result EgAdminFilter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error
	
	return
}

// GetBatchFromID 批量唯一主键查找 
func (obj *_EgAdminFilterMgr) GetBatchFromID(ids []int) (results []*EgAdminFilter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error
	
	return
}
 
// GetFromEmployee 通过employee获取内容  
func (obj *_EgAdminFilterMgr) GetFromEmployee(employee int) (results []*EgAdminFilter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("employee = ?", employee).Find(&results).Error
	
	return
}

// GetBatchFromEmployee 批量唯一主键查找 
func (obj *_EgAdminFilterMgr) GetBatchFromEmployee(employees []int) (results []*EgAdminFilter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("employee IN (?)", employees).Find(&results).Error
	
	return
}
 
// GetFromShop 通过shop获取内容  
func (obj *_EgAdminFilterMgr) GetFromShop(shop int) (results []*EgAdminFilter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("shop = ?", shop).Find(&results).Error
	
	return
}

// GetBatchFromShop 批量唯一主键查找 
func (obj *_EgAdminFilterMgr) GetBatchFromShop(shops []int) (results []*EgAdminFilter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("shop IN (?)", shops).Find(&results).Error
	
	return
}
 
// GetFromController 通过controller获取内容  
func (obj *_EgAdminFilterMgr) GetFromController(controller string) (results []*EgAdminFilter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("controller = ?", controller).Find(&results).Error
	
	return
}

// GetBatchFromController 批量唯一主键查找 
func (obj *_EgAdminFilterMgr) GetBatchFromController(controllers []string) (results []*EgAdminFilter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("controller IN (?)", controllers).Find(&results).Error
	
	return
}
 
// GetFromAction 通过action获取内容  
func (obj *_EgAdminFilterMgr) GetFromAction(action string) (results []*EgAdminFilter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("action = ?", action).Find(&results).Error
	
	return
}

// GetBatchFromAction 批量唯一主键查找 
func (obj *_EgAdminFilterMgr) GetBatchFromAction(actions []string) (results []*EgAdminFilter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("action IN (?)", actions).Find(&results).Error
	
	return
}
 
// GetFromFilter 通过filter获取内容  
func (obj *_EgAdminFilterMgr) GetFromFilter(filter string) (results []*EgAdminFilter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("filter = ?", filter).Find(&results).Error
	
	return
}

// GetBatchFromFilter 批量唯一主键查找 
func (obj *_EgAdminFilterMgr) GetBatchFromFilter(filters []string) (results []*EgAdminFilter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("filter IN (?)", filters).Find(&results).Error
	
	return
}
 
// GetFromFilterID 通过filter_id获取内容  
func (obj *_EgAdminFilterMgr) GetFromFilterID(filterID string) (results []*EgAdminFilter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("filter_id = ?", filterID).Find(&results).Error
	
	return
}

// GetBatchFromFilterID 批量唯一主键查找 
func (obj *_EgAdminFilterMgr) GetBatchFromFilterID(filterIDs []string) (results []*EgAdminFilter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("filter_id IN (?)", filterIDs).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgAdminFilterMgr) FetchByPrimaryKey(id int ) (result EgAdminFilter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error
	
	return
}
 
 // FetchUniqueIndexByAdminFilterSearchIDIDx primay or index 获取唯一内容
 func (obj *_EgAdminFilterMgr) FetchUniqueIndexByAdminFilterSearchIDIDx(employee int ,shop int ,controller string ,action string ,filterID string ) (result EgAdminFilter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("employee = ? AND shop = ? AND controller = ? AND action = ? AND filter_id = ?", employee , shop , controller , action , filterID).Find(&result).Error
	
	return
}
 

 

	

