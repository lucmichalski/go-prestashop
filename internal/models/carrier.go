package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgCarrierMgr struct {
	*_BaseMgr
}

// EgCarrierMgr open func
func EgCarrierMgr(db *gorm.DB) *_EgCarrierMgr {
	if db == nil {
		panic(fmt.Errorf("EgCarrierMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgCarrierMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_carrier"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgCarrierMgr) GetTableName() string {
	return "eg_carrier"
}

// Get 获取
func (obj *_EgCarrierMgr) Get() (result EgCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgCarrierMgr) Gets() (results []*EgCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDCarrier id_carrier获取 
func (obj *_EgCarrierMgr) WithIDCarrier(idCarrier uint32) Option {
	return optionFunc(func(o *options) { o.query["id_carrier"] = idCarrier })
}

// WithIDReference id_reference获取 
func (obj *_EgCarrierMgr) WithIDReference(idReference uint32) Option {
	return optionFunc(func(o *options) { o.query["id_reference"] = idReference })
}

// WithIDTaxRulesGroup id_tax_rules_group获取 
func (obj *_EgCarrierMgr) WithIDTaxRulesGroup(idTaxRulesGroup uint32) Option {
	return optionFunc(func(o *options) { o.query["id_tax_rules_group"] = idTaxRulesGroup })
}

// WithName name获取 
func (obj *_EgCarrierMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithURL url获取 
func (obj *_EgCarrierMgr) WithURL(url string) Option {
	return optionFunc(func(o *options) { o.query["url"] = url })
}

// WithActive active获取 
func (obj *_EgCarrierMgr) WithActive(active bool) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}

// WithDeleted deleted获取 
func (obj *_EgCarrierMgr) WithDeleted(deleted bool) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// WithShippingHandling shipping_handling获取 
func (obj *_EgCarrierMgr) WithShippingHandling(shippingHandling bool) Option {
	return optionFunc(func(o *options) { o.query["shipping_handling"] = shippingHandling })
}

// WithRangeBehavior range_behavior获取 
func (obj *_EgCarrierMgr) WithRangeBehavior(rangeBehavior bool) Option {
	return optionFunc(func(o *options) { o.query["range_behavior"] = rangeBehavior })
}

// WithIsModule is_module获取 
func (obj *_EgCarrierMgr) WithIsModule(isModule bool) Option {
	return optionFunc(func(o *options) { o.query["is_module"] = isModule })
}

// WithIsFree is_free获取 
func (obj *_EgCarrierMgr) WithIsFree(isFree bool) Option {
	return optionFunc(func(o *options) { o.query["is_free"] = isFree })
}

// WithShippingExternal shipping_external获取 
func (obj *_EgCarrierMgr) WithShippingExternal(shippingExternal bool) Option {
	return optionFunc(func(o *options) { o.query["shipping_external"] = shippingExternal })
}

// WithNeedRange need_range获取 
func (obj *_EgCarrierMgr) WithNeedRange(needRange bool) Option {
	return optionFunc(func(o *options) { o.query["need_range"] = needRange })
}

// WithExternalModuleName external_module_name获取 
func (obj *_EgCarrierMgr) WithExternalModuleName(externalModuleName string) Option {
	return optionFunc(func(o *options) { o.query["external_module_name"] = externalModuleName })
}

// WithShippingMethod shipping_method获取 
func (obj *_EgCarrierMgr) WithShippingMethod(shippingMethod int) Option {
	return optionFunc(func(o *options) { o.query["shipping_method"] = shippingMethod })
}

// WithPosition position获取 
func (obj *_EgCarrierMgr) WithPosition(position uint32) Option {
	return optionFunc(func(o *options) { o.query["position"] = position })
}

// WithMaxWidth max_width获取 
func (obj *_EgCarrierMgr) WithMaxWidth(maxWidth int) Option {
	return optionFunc(func(o *options) { o.query["max_width"] = maxWidth })
}

// WithMaxHeight max_height获取 
func (obj *_EgCarrierMgr) WithMaxHeight(maxHeight int) Option {
	return optionFunc(func(o *options) { o.query["max_height"] = maxHeight })
}

// WithMaxDepth max_depth获取 
func (obj *_EgCarrierMgr) WithMaxDepth(maxDepth int) Option {
	return optionFunc(func(o *options) { o.query["max_depth"] = maxDepth })
}

// WithMaxWeight max_weight获取 
func (obj *_EgCarrierMgr) WithMaxWeight(maxWeight float64) Option {
	return optionFunc(func(o *options) { o.query["max_weight"] = maxWeight })
}

// WithGrade grade获取 
func (obj *_EgCarrierMgr) WithGrade(grade int) Option {
	return optionFunc(func(o *options) { o.query["grade"] = grade })
}


// GetByOption 功能选项模式获取
func (obj *_EgCarrierMgr) GetByOption(opts ...Option) (result EgCarrier, err error) {
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
func (obj *_EgCarrierMgr) GetByOptions(opts ...Option) (results []*EgCarrier, err error) {
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


// GetFromIDCarrier 通过id_carrier获取内容  
func (obj *_EgCarrierMgr)  GetFromIDCarrier(idCarrier uint32) (result EgCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_carrier = ?", idCarrier).Find(&result).Error
	
	return
}

// GetBatchFromIDCarrier 批量唯一主键查找 
func (obj *_EgCarrierMgr) GetBatchFromIDCarrier(idCarriers []uint32) (results []*EgCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_carrier IN (?)", idCarriers).Find(&results).Error
	
	return
}
 
// GetFromIDReference 通过id_reference获取内容  
func (obj *_EgCarrierMgr) GetFromIDReference(idReference uint32) (results []*EgCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_reference = ?", idReference).Find(&results).Error
	
	return
}

// GetBatchFromIDReference 批量唯一主键查找 
func (obj *_EgCarrierMgr) GetBatchFromIDReference(idReferences []uint32) (results []*EgCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_reference IN (?)", idReferences).Find(&results).Error
	
	return
}
 
// GetFromIDTaxRulesGroup 通过id_tax_rules_group获取内容  
func (obj *_EgCarrierMgr) GetFromIDTaxRulesGroup(idTaxRulesGroup uint32) (results []*EgCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tax_rules_group = ?", idTaxRulesGroup).Find(&results).Error
	
	return
}

// GetBatchFromIDTaxRulesGroup 批量唯一主键查找 
func (obj *_EgCarrierMgr) GetBatchFromIDTaxRulesGroup(idTaxRulesGroups []uint32) (results []*EgCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tax_rules_group IN (?)", idTaxRulesGroups).Find(&results).Error
	
	return
}
 
// GetFromName 通过name获取内容  
func (obj *_EgCarrierMgr) GetFromName(name string) (results []*EgCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error
	
	return
}

// GetBatchFromName 批量唯一主键查找 
func (obj *_EgCarrierMgr) GetBatchFromName(names []string) (results []*EgCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error
	
	return
}
 
// GetFromURL 通过url获取内容  
func (obj *_EgCarrierMgr) GetFromURL(url string) (results []*EgCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("url = ?", url).Find(&results).Error
	
	return
}

// GetBatchFromURL 批量唯一主键查找 
func (obj *_EgCarrierMgr) GetBatchFromURL(urls []string) (results []*EgCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("url IN (?)", urls).Find(&results).Error
	
	return
}
 
// GetFromActive 通过active获取内容  
func (obj *_EgCarrierMgr) GetFromActive(active bool) (results []*EgCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active = ?", active).Find(&results).Error
	
	return
}

// GetBatchFromActive 批量唯一主键查找 
func (obj *_EgCarrierMgr) GetBatchFromActive(actives []bool) (results []*EgCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active IN (?)", actives).Find(&results).Error
	
	return
}
 
// GetFromDeleted 通过deleted获取内容  
func (obj *_EgCarrierMgr) GetFromDeleted(deleted bool) (results []*EgCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("deleted = ?", deleted).Find(&results).Error
	
	return
}

// GetBatchFromDeleted 批量唯一主键查找 
func (obj *_EgCarrierMgr) GetBatchFromDeleted(deleteds []bool) (results []*EgCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("deleted IN (?)", deleteds).Find(&results).Error
	
	return
}
 
// GetFromShippingHandling 通过shipping_handling获取内容  
func (obj *_EgCarrierMgr) GetFromShippingHandling(shippingHandling bool) (results []*EgCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("shipping_handling = ?", shippingHandling).Find(&results).Error
	
	return
}

// GetBatchFromShippingHandling 批量唯一主键查找 
func (obj *_EgCarrierMgr) GetBatchFromShippingHandling(shippingHandlings []bool) (results []*EgCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("shipping_handling IN (?)", shippingHandlings).Find(&results).Error
	
	return
}
 
// GetFromRangeBehavior 通过range_behavior获取内容  
func (obj *_EgCarrierMgr) GetFromRangeBehavior(rangeBehavior bool) (results []*EgCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("range_behavior = ?", rangeBehavior).Find(&results).Error
	
	return
}

// GetBatchFromRangeBehavior 批量唯一主键查找 
func (obj *_EgCarrierMgr) GetBatchFromRangeBehavior(rangeBehaviors []bool) (results []*EgCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("range_behavior IN (?)", rangeBehaviors).Find(&results).Error
	
	return
}
 
// GetFromIsModule 通过is_module获取内容  
func (obj *_EgCarrierMgr) GetFromIsModule(isModule bool) (results []*EgCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_module = ?", isModule).Find(&results).Error
	
	return
}

// GetBatchFromIsModule 批量唯一主键查找 
func (obj *_EgCarrierMgr) GetBatchFromIsModule(isModules []bool) (results []*EgCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_module IN (?)", isModules).Find(&results).Error
	
	return
}
 
// GetFromIsFree 通过is_free获取内容  
func (obj *_EgCarrierMgr) GetFromIsFree(isFree bool) (results []*EgCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_free = ?", isFree).Find(&results).Error
	
	return
}

// GetBatchFromIsFree 批量唯一主键查找 
func (obj *_EgCarrierMgr) GetBatchFromIsFree(isFrees []bool) (results []*EgCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_free IN (?)", isFrees).Find(&results).Error
	
	return
}
 
// GetFromShippingExternal 通过shipping_external获取内容  
func (obj *_EgCarrierMgr) GetFromShippingExternal(shippingExternal bool) (results []*EgCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("shipping_external = ?", shippingExternal).Find(&results).Error
	
	return
}

// GetBatchFromShippingExternal 批量唯一主键查找 
func (obj *_EgCarrierMgr) GetBatchFromShippingExternal(shippingExternals []bool) (results []*EgCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("shipping_external IN (?)", shippingExternals).Find(&results).Error
	
	return
}
 
// GetFromNeedRange 通过need_range获取内容  
func (obj *_EgCarrierMgr) GetFromNeedRange(needRange bool) (results []*EgCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("need_range = ?", needRange).Find(&results).Error
	
	return
}

// GetBatchFromNeedRange 批量唯一主键查找 
func (obj *_EgCarrierMgr) GetBatchFromNeedRange(needRanges []bool) (results []*EgCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("need_range IN (?)", needRanges).Find(&results).Error
	
	return
}
 
// GetFromExternalModuleName 通过external_module_name获取内容  
func (obj *_EgCarrierMgr) GetFromExternalModuleName(externalModuleName string) (results []*EgCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("external_module_name = ?", externalModuleName).Find(&results).Error
	
	return
}

// GetBatchFromExternalModuleName 批量唯一主键查找 
func (obj *_EgCarrierMgr) GetBatchFromExternalModuleName(externalModuleNames []string) (results []*EgCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("external_module_name IN (?)", externalModuleNames).Find(&results).Error
	
	return
}
 
// GetFromShippingMethod 通过shipping_method获取内容  
func (obj *_EgCarrierMgr) GetFromShippingMethod(shippingMethod int) (results []*EgCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("shipping_method = ?", shippingMethod).Find(&results).Error
	
	return
}

// GetBatchFromShippingMethod 批量唯一主键查找 
func (obj *_EgCarrierMgr) GetBatchFromShippingMethod(shippingMethods []int) (results []*EgCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("shipping_method IN (?)", shippingMethods).Find(&results).Error
	
	return
}
 
// GetFromPosition 通过position获取内容  
func (obj *_EgCarrierMgr) GetFromPosition(position uint32) (results []*EgCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("position = ?", position).Find(&results).Error
	
	return
}

// GetBatchFromPosition 批量唯一主键查找 
func (obj *_EgCarrierMgr) GetBatchFromPosition(positions []uint32) (results []*EgCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("position IN (?)", positions).Find(&results).Error
	
	return
}
 
// GetFromMaxWidth 通过max_width获取内容  
func (obj *_EgCarrierMgr) GetFromMaxWidth(maxWidth int) (results []*EgCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("max_width = ?", maxWidth).Find(&results).Error
	
	return
}

// GetBatchFromMaxWidth 批量唯一主键查找 
func (obj *_EgCarrierMgr) GetBatchFromMaxWidth(maxWidths []int) (results []*EgCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("max_width IN (?)", maxWidths).Find(&results).Error
	
	return
}
 
// GetFromMaxHeight 通过max_height获取内容  
func (obj *_EgCarrierMgr) GetFromMaxHeight(maxHeight int) (results []*EgCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("max_height = ?", maxHeight).Find(&results).Error
	
	return
}

// GetBatchFromMaxHeight 批量唯一主键查找 
func (obj *_EgCarrierMgr) GetBatchFromMaxHeight(maxHeights []int) (results []*EgCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("max_height IN (?)", maxHeights).Find(&results).Error
	
	return
}
 
// GetFromMaxDepth 通过max_depth获取内容  
func (obj *_EgCarrierMgr) GetFromMaxDepth(maxDepth int) (results []*EgCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("max_depth = ?", maxDepth).Find(&results).Error
	
	return
}

// GetBatchFromMaxDepth 批量唯一主键查找 
func (obj *_EgCarrierMgr) GetBatchFromMaxDepth(maxDepths []int) (results []*EgCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("max_depth IN (?)", maxDepths).Find(&results).Error
	
	return
}
 
// GetFromMaxWeight 通过max_weight获取内容  
func (obj *_EgCarrierMgr) GetFromMaxWeight(maxWeight float64) (results []*EgCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("max_weight = ?", maxWeight).Find(&results).Error
	
	return
}

// GetBatchFromMaxWeight 批量唯一主键查找 
func (obj *_EgCarrierMgr) GetBatchFromMaxWeight(maxWeights []float64) (results []*EgCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("max_weight IN (?)", maxWeights).Find(&results).Error
	
	return
}
 
// GetFromGrade 通过grade获取内容  
func (obj *_EgCarrierMgr) GetFromGrade(grade int) (results []*EgCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("grade = ?", grade).Find(&results).Error
	
	return
}

// GetBatchFromGrade 批量唯一主键查找 
func (obj *_EgCarrierMgr) GetBatchFromGrade(grades []int) (results []*EgCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("grade IN (?)", grades).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgCarrierMgr) FetchByPrimaryKey(idCarrier uint32 ) (result EgCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_carrier = ?", idCarrier).Find(&result).Error
	
	return
}
 

 
 // FetchIndexByReference  获取多个内容
 func (obj *_EgCarrierMgr) FetchIndexByReference(idReference uint32 ,active bool ,deleted bool ) (results []*EgCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_reference = ? AND active = ? AND deleted = ?", idReference , active , deleted).Find(&results).Error
	
	return
}
 
 // FetchIndexByIDTaxRulesGroup  获取多个内容
 func (obj *_EgCarrierMgr) FetchIndexByIDTaxRulesGroup(idTaxRulesGroup uint32 ) (results []*EgCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tax_rules_group = ?", idTaxRulesGroup).Find(&results).Error
	
	return
}
 
 // FetchIndexByDeleted  获取多个内容
 func (obj *_EgCarrierMgr) FetchIndexByDeleted(active bool ,deleted bool ) (results []*EgCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active = ? AND deleted = ?", active , deleted).Find(&results).Error
	
	return
}
 

	

