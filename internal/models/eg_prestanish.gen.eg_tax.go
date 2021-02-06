package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgTaxMgr struct {
	*_BaseMgr
}

// EgTaxMgr open func
func EgTaxMgr(db *gorm.DB) *_EgTaxMgr {
	if db == nil {
		panic(fmt.Errorf("EgTaxMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgTaxMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_tax"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgTaxMgr) GetTableName() string {
	return "eg_tax"
}

// Get 获取
func (obj *_EgTaxMgr) Get() (result EgTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgTaxMgr) Gets() (results []*EgTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDTax id_tax获取 
func (obj *_EgTaxMgr) WithIDTax(idTax uint32) Option {
	return optionFunc(func(o *options) { o.query["id_tax"] = idTax })
}

// WithRate rate获取 
func (obj *_EgTaxMgr) WithRate(rate float64) Option {
	return optionFunc(func(o *options) { o.query["rate"] = rate })
}

// WithActive active获取 
func (obj *_EgTaxMgr) WithActive(active bool) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}

// WithDeleted deleted获取 
func (obj *_EgTaxMgr) WithDeleted(deleted bool) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}


// GetByOption 功能选项模式获取
func (obj *_EgTaxMgr) GetByOption(opts ...Option) (result EgTax, err error) {
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
func (obj *_EgTaxMgr) GetByOptions(opts ...Option) (results []*EgTax, err error) {
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


// GetFromIDTax 通过id_tax获取内容  
func (obj *_EgTaxMgr)  GetFromIDTax(idTax uint32) (result EgTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tax = ?", idTax).Find(&result).Error
	
	return
}

// GetBatchFromIDTax 批量唯一主键查找 
func (obj *_EgTaxMgr) GetBatchFromIDTax(idTaxs []uint32) (results []*EgTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tax IN (?)", idTaxs).Find(&results).Error
	
	return
}
 
// GetFromRate 通过rate获取内容  
func (obj *_EgTaxMgr) GetFromRate(rate float64) (results []*EgTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("rate = ?", rate).Find(&results).Error
	
	return
}

// GetBatchFromRate 批量唯一主键查找 
func (obj *_EgTaxMgr) GetBatchFromRate(rates []float64) (results []*EgTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("rate IN (?)", rates).Find(&results).Error
	
	return
}
 
// GetFromActive 通过active获取内容  
func (obj *_EgTaxMgr) GetFromActive(active bool) (results []*EgTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active = ?", active).Find(&results).Error
	
	return
}

// GetBatchFromActive 批量唯一主键查找 
func (obj *_EgTaxMgr) GetBatchFromActive(actives []bool) (results []*EgTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active IN (?)", actives).Find(&results).Error
	
	return
}
 
// GetFromDeleted 通过deleted获取内容  
func (obj *_EgTaxMgr) GetFromDeleted(deleted bool) (results []*EgTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("deleted = ?", deleted).Find(&results).Error
	
	return
}

// GetBatchFromDeleted 批量唯一主键查找 
func (obj *_EgTaxMgr) GetBatchFromDeleted(deleteds []bool) (results []*EgTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("deleted IN (?)", deleteds).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgTaxMgr) FetchByPrimaryKey(idTax uint32 ) (result EgTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tax = ?", idTax).Find(&result).Error
	
	return
}
 

 

	

