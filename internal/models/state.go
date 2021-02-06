package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgStateMgr struct {
	*_BaseMgr
}

// EgStateMgr open func
func EgStateMgr(db *gorm.DB) *_EgStateMgr {
	if db == nil {
		panic(fmt.Errorf("EgStateMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgStateMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_state"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgStateMgr) GetTableName() string {
	return "eg_state"
}

// Get 获取
func (obj *_EgStateMgr) Get() (result EgState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgStateMgr) Gets() (results []*EgState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDState id_state获取 
func (obj *_EgStateMgr) WithIDState(idState uint32) Option {
	return optionFunc(func(o *options) { o.query["id_state"] = idState })
}

// WithIDCountry id_country获取 
func (obj *_EgStateMgr) WithIDCountry(idCountry uint32) Option {
	return optionFunc(func(o *options) { o.query["id_country"] = idCountry })
}

// WithIDZone id_zone获取 
func (obj *_EgStateMgr) WithIDZone(idZone uint32) Option {
	return optionFunc(func(o *options) { o.query["id_zone"] = idZone })
}

// WithName name获取 
func (obj *_EgStateMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithIsoCode iso_code获取 
func (obj *_EgStateMgr) WithIsoCode(isoCode string) Option {
	return optionFunc(func(o *options) { o.query["iso_code"] = isoCode })
}

// WithTaxBehavior tax_behavior获取 
func (obj *_EgStateMgr) WithTaxBehavior(taxBehavior int16) Option {
	return optionFunc(func(o *options) { o.query["tax_behavior"] = taxBehavior })
}

// WithActive active获取 
func (obj *_EgStateMgr) WithActive(active bool) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}


// GetByOption 功能选项模式获取
func (obj *_EgStateMgr) GetByOption(opts ...Option) (result EgState, err error) {
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
func (obj *_EgStateMgr) GetByOptions(opts ...Option) (results []*EgState, err error) {
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


// GetFromIDState 通过id_state获取内容  
func (obj *_EgStateMgr)  GetFromIDState(idState uint32) (result EgState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_state = ?", idState).Find(&result).Error
	
	return
}

// GetBatchFromIDState 批量唯一主键查找 
func (obj *_EgStateMgr) GetBatchFromIDState(idStates []uint32) (results []*EgState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_state IN (?)", idStates).Find(&results).Error
	
	return
}
 
// GetFromIDCountry 通过id_country获取内容  
func (obj *_EgStateMgr) GetFromIDCountry(idCountry uint32) (results []*EgState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_country = ?", idCountry).Find(&results).Error
	
	return
}

// GetBatchFromIDCountry 批量唯一主键查找 
func (obj *_EgStateMgr) GetBatchFromIDCountry(idCountrys []uint32) (results []*EgState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_country IN (?)", idCountrys).Find(&results).Error
	
	return
}
 
// GetFromIDZone 通过id_zone获取内容  
func (obj *_EgStateMgr) GetFromIDZone(idZone uint32) (results []*EgState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_zone = ?", idZone).Find(&results).Error
	
	return
}

// GetBatchFromIDZone 批量唯一主键查找 
func (obj *_EgStateMgr) GetBatchFromIDZone(idZones []uint32) (results []*EgState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_zone IN (?)", idZones).Find(&results).Error
	
	return
}
 
// GetFromName 通过name获取内容  
func (obj *_EgStateMgr) GetFromName(name string) (results []*EgState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error
	
	return
}

// GetBatchFromName 批量唯一主键查找 
func (obj *_EgStateMgr) GetBatchFromName(names []string) (results []*EgState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error
	
	return
}
 
// GetFromIsoCode 通过iso_code获取内容  
func (obj *_EgStateMgr) GetFromIsoCode(isoCode string) (results []*EgState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("iso_code = ?", isoCode).Find(&results).Error
	
	return
}

// GetBatchFromIsoCode 批量唯一主键查找 
func (obj *_EgStateMgr) GetBatchFromIsoCode(isoCodes []string) (results []*EgState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("iso_code IN (?)", isoCodes).Find(&results).Error
	
	return
}
 
// GetFromTaxBehavior 通过tax_behavior获取内容  
func (obj *_EgStateMgr) GetFromTaxBehavior(taxBehavior int16) (results []*EgState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("tax_behavior = ?", taxBehavior).Find(&results).Error
	
	return
}

// GetBatchFromTaxBehavior 批量唯一主键查找 
func (obj *_EgStateMgr) GetBatchFromTaxBehavior(taxBehaviors []int16) (results []*EgState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("tax_behavior IN (?)", taxBehaviors).Find(&results).Error
	
	return
}
 
// GetFromActive 通过active获取内容  
func (obj *_EgStateMgr) GetFromActive(active bool) (results []*EgState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active = ?", active).Find(&results).Error
	
	return
}

// GetBatchFromActive 批量唯一主键查找 
func (obj *_EgStateMgr) GetBatchFromActive(actives []bool) (results []*EgState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active IN (?)", actives).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgStateMgr) FetchByPrimaryKey(idState uint32 ) (result EgState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_state = ?", idState).Find(&result).Error
	
	return
}
 

 
 // FetchIndexByIDCountry  获取多个内容
 func (obj *_EgStateMgr) FetchIndexByIDCountry(idCountry uint32 ) (results []*EgState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_country = ?", idCountry).Find(&results).Error
	
	return
}
 
 // FetchIndexByIDZone  获取多个内容
 func (obj *_EgStateMgr) FetchIndexByIDZone(idZone uint32 ) (results []*EgState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_zone = ?", idZone).Find(&results).Error
	
	return
}
 
 // FetchIndexByName  获取多个内容
 func (obj *_EgStateMgr) FetchIndexByName(name string ) (results []*EgState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error
	
	return
}
 

	

