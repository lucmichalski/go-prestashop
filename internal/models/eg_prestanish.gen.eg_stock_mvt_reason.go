package	model	
import (	
"context"	
"gorm.io/gorm"	
"time"	
"fmt"	
)	

type _EgStockMvtReasonMgr struct {
	*_BaseMgr
}

// EgStockMvtReasonMgr open func
func EgStockMvtReasonMgr(db *gorm.DB) *_EgStockMvtReasonMgr {
	if db == nil {
		panic(fmt.Errorf("EgStockMvtReasonMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgStockMvtReasonMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_stock_mvt_reason"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgStockMvtReasonMgr) GetTableName() string {
	return "eg_stock_mvt_reason"
}

// Get 获取
func (obj *_EgStockMvtReasonMgr) Get() (result EgStockMvtReason, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgStockMvtReasonMgr) Gets() (results []*EgStockMvtReason, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDStockMvtReason id_stock_mvt_reason获取 
func (obj *_EgStockMvtReasonMgr) WithIDStockMvtReason(idStockMvtReason uint32) Option {
	return optionFunc(func(o *options) { o.query["id_stock_mvt_reason"] = idStockMvtReason })
}

// WithSign sign获取 
func (obj *_EgStockMvtReasonMgr) WithSign(sign bool) Option {
	return optionFunc(func(o *options) { o.query["sign"] = sign })
}

// WithDateAdd date_add获取 
func (obj *_EgStockMvtReasonMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

// WithDateUpd date_upd获取 
func (obj *_EgStockMvtReasonMgr) WithDateUpd(dateUpd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_upd"] = dateUpd })
}

// WithDeleted deleted获取 
func (obj *_EgStockMvtReasonMgr) WithDeleted(deleted bool) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}


// GetByOption 功能选项模式获取
func (obj *_EgStockMvtReasonMgr) GetByOption(opts ...Option) (result EgStockMvtReason, err error) {
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
func (obj *_EgStockMvtReasonMgr) GetByOptions(opts ...Option) (results []*EgStockMvtReason, err error) {
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


// GetFromIDStockMvtReason 通过id_stock_mvt_reason获取内容  
func (obj *_EgStockMvtReasonMgr)  GetFromIDStockMvtReason(idStockMvtReason uint32) (result EgStockMvtReason, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_stock_mvt_reason = ?", idStockMvtReason).Find(&result).Error
	
	return
}

// GetBatchFromIDStockMvtReason 批量唯一主键查找 
func (obj *_EgStockMvtReasonMgr) GetBatchFromIDStockMvtReason(idStockMvtReasons []uint32) (results []*EgStockMvtReason, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_stock_mvt_reason IN (?)", idStockMvtReasons).Find(&results).Error
	
	return
}
 
// GetFromSign 通过sign获取内容  
func (obj *_EgStockMvtReasonMgr) GetFromSign(sign bool) (results []*EgStockMvtReason, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("sign = ?", sign).Find(&results).Error
	
	return
}

// GetBatchFromSign 批量唯一主键查找 
func (obj *_EgStockMvtReasonMgr) GetBatchFromSign(signs []bool) (results []*EgStockMvtReason, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("sign IN (?)", signs).Find(&results).Error
	
	return
}
 
// GetFromDateAdd 通过date_add获取内容  
func (obj *_EgStockMvtReasonMgr) GetFromDateAdd(dateAdd time.Time) (results []*EgStockMvtReason, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error
	
	return
}

// GetBatchFromDateAdd 批量唯一主键查找 
func (obj *_EgStockMvtReasonMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*EgStockMvtReason, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error
	
	return
}
 
// GetFromDateUpd 通过date_upd获取内容  
func (obj *_EgStockMvtReasonMgr) GetFromDateUpd(dateUpd time.Time) (results []*EgStockMvtReason, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd = ?", dateUpd).Find(&results).Error
	
	return
}

// GetBatchFromDateUpd 批量唯一主键查找 
func (obj *_EgStockMvtReasonMgr) GetBatchFromDateUpd(dateUpds []time.Time) (results []*EgStockMvtReason, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd IN (?)", dateUpds).Find(&results).Error
	
	return
}
 
// GetFromDeleted 通过deleted获取内容  
func (obj *_EgStockMvtReasonMgr) GetFromDeleted(deleted bool) (results []*EgStockMvtReason, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("deleted = ?", deleted).Find(&results).Error
	
	return
}

// GetBatchFromDeleted 批量唯一主键查找 
func (obj *_EgStockMvtReasonMgr) GetBatchFromDeleted(deleteds []bool) (results []*EgStockMvtReason, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("deleted IN (?)", deleteds).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgStockMvtReasonMgr) FetchByPrimaryKey(idStockMvtReason uint32 ) (result EgStockMvtReason, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_stock_mvt_reason = ?", idStockMvtReason).Find(&result).Error
	
	return
}
 

 

	

