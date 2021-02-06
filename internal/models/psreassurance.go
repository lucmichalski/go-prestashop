package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
"time"	
)	

type _EgPsreassuranceMgr struct {
	*_BaseMgr
}

// EgPsreassuranceMgr open func
func EgPsreassuranceMgr(db *gorm.DB) *_EgPsreassuranceMgr {
	if db == nil {
		panic(fmt.Errorf("EgPsreassuranceMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgPsreassuranceMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_psreassurance"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgPsreassuranceMgr) GetTableName() string {
	return "eg_psreassurance"
}

// Get 获取
func (obj *_EgPsreassuranceMgr) Get() (result EgPsreassurance, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgPsreassuranceMgr) Gets() (results []*EgPsreassurance, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDPsreassurance id_psreassurance获取 
func (obj *_EgPsreassuranceMgr) WithIDPsreassurance(idPsreassurance uint32) Option {
	return optionFunc(func(o *options) { o.query["id_psreassurance"] = idPsreassurance })
}

// WithIcon icon获取 
func (obj *_EgPsreassuranceMgr) WithIcon(icon string) Option {
	return optionFunc(func(o *options) { o.query["icon"] = icon })
}

// WithCustomIcon custom_icon获取 
func (obj *_EgPsreassuranceMgr) WithCustomIcon(customIcon string) Option {
	return optionFunc(func(o *options) { o.query["custom_icon"] = customIcon })
}

// WithStatus status获取 
func (obj *_EgPsreassuranceMgr) WithStatus(status uint32) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithPosition position获取 
func (obj *_EgPsreassuranceMgr) WithPosition(position uint32) Option {
	return optionFunc(func(o *options) { o.query["position"] = position })
}

// WithIDShop id_shop获取 
func (obj *_EgPsreassuranceMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

// WithTypeLink type_link获取 
func (obj *_EgPsreassuranceMgr) WithTypeLink(typeLink uint32) Option {
	return optionFunc(func(o *options) { o.query["type_link"] = typeLink })
}

// WithIDCms id_cms获取 
func (obj *_EgPsreassuranceMgr) WithIDCms(idCms uint32) Option {
	return optionFunc(func(o *options) { o.query["id_cms"] = idCms })
}

// WithDateAdd date_add获取 
func (obj *_EgPsreassuranceMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

// WithDateUpd date_upd获取 
func (obj *_EgPsreassuranceMgr) WithDateUpd(dateUpd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_upd"] = dateUpd })
}


// GetByOption 功能选项模式获取
func (obj *_EgPsreassuranceMgr) GetByOption(opts ...Option) (result EgPsreassurance, err error) {
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
func (obj *_EgPsreassuranceMgr) GetByOptions(opts ...Option) (results []*EgPsreassurance, err error) {
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


// GetFromIDPsreassurance 通过id_psreassurance获取内容  
func (obj *_EgPsreassuranceMgr)  GetFromIDPsreassurance(idPsreassurance uint32) (result EgPsreassurance, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_psreassurance = ?", idPsreassurance).Find(&result).Error
	
	return
}

// GetBatchFromIDPsreassurance 批量唯一主键查找 
func (obj *_EgPsreassuranceMgr) GetBatchFromIDPsreassurance(idPsreassurances []uint32) (results []*EgPsreassurance, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_psreassurance IN (?)", idPsreassurances).Find(&results).Error
	
	return
}
 
// GetFromIcon 通过icon获取内容  
func (obj *_EgPsreassuranceMgr) GetFromIcon(icon string) (results []*EgPsreassurance, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("icon = ?", icon).Find(&results).Error
	
	return
}

// GetBatchFromIcon 批量唯一主键查找 
func (obj *_EgPsreassuranceMgr) GetBatchFromIcon(icons []string) (results []*EgPsreassurance, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("icon IN (?)", icons).Find(&results).Error
	
	return
}
 
// GetFromCustomIcon 通过custom_icon获取内容  
func (obj *_EgPsreassuranceMgr) GetFromCustomIcon(customIcon string) (results []*EgPsreassurance, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("custom_icon = ?", customIcon).Find(&results).Error
	
	return
}

// GetBatchFromCustomIcon 批量唯一主键查找 
func (obj *_EgPsreassuranceMgr) GetBatchFromCustomIcon(customIcons []string) (results []*EgPsreassurance, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("custom_icon IN (?)", customIcons).Find(&results).Error
	
	return
}
 
// GetFromStatus 通过status获取内容  
func (obj *_EgPsreassuranceMgr) GetFromStatus(status uint32) (results []*EgPsreassurance, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error
	
	return
}

// GetBatchFromStatus 批量唯一主键查找 
func (obj *_EgPsreassuranceMgr) GetBatchFromStatus(statuss []uint32) (results []*EgPsreassurance, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status IN (?)", statuss).Find(&results).Error
	
	return
}
 
// GetFromPosition 通过position获取内容  
func (obj *_EgPsreassuranceMgr) GetFromPosition(position uint32) (results []*EgPsreassurance, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("position = ?", position).Find(&results).Error
	
	return
}

// GetBatchFromPosition 批量唯一主键查找 
func (obj *_EgPsreassuranceMgr) GetBatchFromPosition(positions []uint32) (results []*EgPsreassurance, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("position IN (?)", positions).Find(&results).Error
	
	return
}
 
// GetFromIDShop 通过id_shop获取内容  
func (obj *_EgPsreassuranceMgr) GetFromIDShop(idShop uint32) (results []*EgPsreassurance, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error
	
	return
}

// GetBatchFromIDShop 批量唯一主键查找 
func (obj *_EgPsreassuranceMgr) GetBatchFromIDShop(idShops []uint32) (results []*EgPsreassurance, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error
	
	return
}
 
// GetFromTypeLink 通过type_link获取内容  
func (obj *_EgPsreassuranceMgr) GetFromTypeLink(typeLink uint32) (results []*EgPsreassurance, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type_link = ?", typeLink).Find(&results).Error
	
	return
}

// GetBatchFromTypeLink 批量唯一主键查找 
func (obj *_EgPsreassuranceMgr) GetBatchFromTypeLink(typeLinks []uint32) (results []*EgPsreassurance, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type_link IN (?)", typeLinks).Find(&results).Error
	
	return
}
 
// GetFromIDCms 通过id_cms获取内容  
func (obj *_EgPsreassuranceMgr) GetFromIDCms(idCms uint32) (results []*EgPsreassurance, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cms = ?", idCms).Find(&results).Error
	
	return
}

// GetBatchFromIDCms 批量唯一主键查找 
func (obj *_EgPsreassuranceMgr) GetBatchFromIDCms(idCmss []uint32) (results []*EgPsreassurance, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cms IN (?)", idCmss).Find(&results).Error
	
	return
}
 
// GetFromDateAdd 通过date_add获取内容  
func (obj *_EgPsreassuranceMgr) GetFromDateAdd(dateAdd time.Time) (results []*EgPsreassurance, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error
	
	return
}

// GetBatchFromDateAdd 批量唯一主键查找 
func (obj *_EgPsreassuranceMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*EgPsreassurance, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error
	
	return
}
 
// GetFromDateUpd 通过date_upd获取内容  
func (obj *_EgPsreassuranceMgr) GetFromDateUpd(dateUpd time.Time) (results []*EgPsreassurance, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd = ?", dateUpd).Find(&results).Error
	
	return
}

// GetBatchFromDateUpd 批量唯一主键查找 
func (obj *_EgPsreassuranceMgr) GetBatchFromDateUpd(dateUpds []time.Time) (results []*EgPsreassurance, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd IN (?)", dateUpds).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgPsreassuranceMgr) FetchByPrimaryKey(idPsreassurance uint32 ) (result EgPsreassurance, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_psreassurance = ?", idPsreassurance).Find(&result).Error
	
	return
}
 

 

	

