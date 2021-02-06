package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
"time"	
)	

type _EgGroupMgr struct {
	*_BaseMgr
}

// EgGroupMgr open func
func EgGroupMgr(db *gorm.DB) *_EgGroupMgr {
	if db == nil {
		panic(fmt.Errorf("EgGroupMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgGroupMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_group"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgGroupMgr) GetTableName() string {
	return "eg_group"
}

// Get 获取
func (obj *_EgGroupMgr) Get() (result EgGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgGroupMgr) Gets() (results []*EgGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDGroup id_group获取 
func (obj *_EgGroupMgr) WithIDGroup(idGroup uint32) Option {
	return optionFunc(func(o *options) { o.query["id_group"] = idGroup })
}

// WithReduction reduction获取 
func (obj *_EgGroupMgr) WithReduction(reduction float64) Option {
	return optionFunc(func(o *options) { o.query["reduction"] = reduction })
}

// WithPriceDisplayMethod price_display_method获取 
func (obj *_EgGroupMgr) WithPriceDisplayMethod(priceDisplayMethod int8) Option {
	return optionFunc(func(o *options) { o.query["price_display_method"] = priceDisplayMethod })
}

// WithShowPrices show_prices获取 
func (obj *_EgGroupMgr) WithShowPrices(showPrices bool) Option {
	return optionFunc(func(o *options) { o.query["show_prices"] = showPrices })
}

// WithDateAdd date_add获取 
func (obj *_EgGroupMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

// WithDateUpd date_upd获取 
func (obj *_EgGroupMgr) WithDateUpd(dateUpd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_upd"] = dateUpd })
}


// GetByOption 功能选项模式获取
func (obj *_EgGroupMgr) GetByOption(opts ...Option) (result EgGroup, err error) {
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
func (obj *_EgGroupMgr) GetByOptions(opts ...Option) (results []*EgGroup, err error) {
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


// GetFromIDGroup 通过id_group获取内容  
func (obj *_EgGroupMgr)  GetFromIDGroup(idGroup uint32) (result EgGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_group = ?", idGroup).Find(&result).Error
	
	return
}

// GetBatchFromIDGroup 批量唯一主键查找 
func (obj *_EgGroupMgr) GetBatchFromIDGroup(idGroups []uint32) (results []*EgGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_group IN (?)", idGroups).Find(&results).Error
	
	return
}
 
// GetFromReduction 通过reduction获取内容  
func (obj *_EgGroupMgr) GetFromReduction(reduction float64) (results []*EgGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reduction = ?", reduction).Find(&results).Error
	
	return
}

// GetBatchFromReduction 批量唯一主键查找 
func (obj *_EgGroupMgr) GetBatchFromReduction(reductions []float64) (results []*EgGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reduction IN (?)", reductions).Find(&results).Error
	
	return
}
 
// GetFromPriceDisplayMethod 通过price_display_method获取内容  
func (obj *_EgGroupMgr) GetFromPriceDisplayMethod(priceDisplayMethod int8) (results []*EgGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("price_display_method = ?", priceDisplayMethod).Find(&results).Error
	
	return
}

// GetBatchFromPriceDisplayMethod 批量唯一主键查找 
func (obj *_EgGroupMgr) GetBatchFromPriceDisplayMethod(priceDisplayMethods []int8) (results []*EgGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("price_display_method IN (?)", priceDisplayMethods).Find(&results).Error
	
	return
}
 
// GetFromShowPrices 通过show_prices获取内容  
func (obj *_EgGroupMgr) GetFromShowPrices(showPrices bool) (results []*EgGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("show_prices = ?", showPrices).Find(&results).Error
	
	return
}

// GetBatchFromShowPrices 批量唯一主键查找 
func (obj *_EgGroupMgr) GetBatchFromShowPrices(showPricess []bool) (results []*EgGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("show_prices IN (?)", showPricess).Find(&results).Error
	
	return
}
 
// GetFromDateAdd 通过date_add获取内容  
func (obj *_EgGroupMgr) GetFromDateAdd(dateAdd time.Time) (results []*EgGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error
	
	return
}

// GetBatchFromDateAdd 批量唯一主键查找 
func (obj *_EgGroupMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*EgGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error
	
	return
}
 
// GetFromDateUpd 通过date_upd获取内容  
func (obj *_EgGroupMgr) GetFromDateUpd(dateUpd time.Time) (results []*EgGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd = ?", dateUpd).Find(&results).Error
	
	return
}

// GetBatchFromDateUpd 批量唯一主键查找 
func (obj *_EgGroupMgr) GetBatchFromDateUpd(dateUpds []time.Time) (results []*EgGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd IN (?)", dateUpds).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgGroupMgr) FetchByPrimaryKey(idGroup uint32 ) (result EgGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_group = ?", idGroup).Find(&result).Error
	
	return
}
 

 

	

