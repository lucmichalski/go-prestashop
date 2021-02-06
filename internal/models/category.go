package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
"time"	
)	

type _EgCategoryMgr struct {
	*_BaseMgr
}

// EgCategoryMgr open func
func EgCategoryMgr(db *gorm.DB) *_EgCategoryMgr {
	if db == nil {
		panic(fmt.Errorf("EgCategoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgCategoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_category"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgCategoryMgr) GetTableName() string {
	return "eg_category"
}

// Get 获取
func (obj *_EgCategoryMgr) Get() (result EgCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgCategoryMgr) Gets() (results []*EgCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDCategory id_category获取 
func (obj *_EgCategoryMgr) WithIDCategory(idCategory uint32) Option {
	return optionFunc(func(o *options) { o.query["id_category"] = idCategory })
}

// WithIDParent id_parent获取 
func (obj *_EgCategoryMgr) WithIDParent(idParent uint32) Option {
	return optionFunc(func(o *options) { o.query["id_parent"] = idParent })
}

// WithIDShopDefault id_shop_default获取 
func (obj *_EgCategoryMgr) WithIDShopDefault(idShopDefault uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop_default"] = idShopDefault })
}

// WithLevelDepth level_depth获取 
func (obj *_EgCategoryMgr) WithLevelDepth(levelDepth uint8) Option {
	return optionFunc(func(o *options) { o.query["level_depth"] = levelDepth })
}

// WithNleft nleft获取 
func (obj *_EgCategoryMgr) WithNleft(nleft uint32) Option {
	return optionFunc(func(o *options) { o.query["nleft"] = nleft })
}

// WithNright nright获取 
func (obj *_EgCategoryMgr) WithNright(nright uint32) Option {
	return optionFunc(func(o *options) { o.query["nright"] = nright })
}

// WithActive active获取 
func (obj *_EgCategoryMgr) WithActive(active bool) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}

// WithDateAdd date_add获取 
func (obj *_EgCategoryMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

// WithDateUpd date_upd获取 
func (obj *_EgCategoryMgr) WithDateUpd(dateUpd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_upd"] = dateUpd })
}

// WithPosition position获取 
func (obj *_EgCategoryMgr) WithPosition(position uint32) Option {
	return optionFunc(func(o *options) { o.query["position"] = position })
}

// WithIsRootCategory is_root_category获取 
func (obj *_EgCategoryMgr) WithIsRootCategory(isRootCategory bool) Option {
	return optionFunc(func(o *options) { o.query["is_root_category"] = isRootCategory })
}


// GetByOption 功能选项模式获取
func (obj *_EgCategoryMgr) GetByOption(opts ...Option) (result EgCategory, err error) {
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
func (obj *_EgCategoryMgr) GetByOptions(opts ...Option) (results []*EgCategory, err error) {
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


// GetFromIDCategory 通过id_category获取内容  
func (obj *_EgCategoryMgr)  GetFromIDCategory(idCategory uint32) (result EgCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_category = ?", idCategory).Find(&result).Error
	
	return
}

// GetBatchFromIDCategory 批量唯一主键查找 
func (obj *_EgCategoryMgr) GetBatchFromIDCategory(idCategorys []uint32) (results []*EgCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_category IN (?)", idCategorys).Find(&results).Error
	
	return
}
 
// GetFromIDParent 通过id_parent获取内容  
func (obj *_EgCategoryMgr) GetFromIDParent(idParent uint32) (results []*EgCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_parent = ?", idParent).Find(&results).Error
	
	return
}

// GetBatchFromIDParent 批量唯一主键查找 
func (obj *_EgCategoryMgr) GetBatchFromIDParent(idParents []uint32) (results []*EgCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_parent IN (?)", idParents).Find(&results).Error
	
	return
}
 
// GetFromIDShopDefault 通过id_shop_default获取内容  
func (obj *_EgCategoryMgr) GetFromIDShopDefault(idShopDefault uint32) (results []*EgCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_default = ?", idShopDefault).Find(&results).Error
	
	return
}

// GetBatchFromIDShopDefault 批量唯一主键查找 
func (obj *_EgCategoryMgr) GetBatchFromIDShopDefault(idShopDefaults []uint32) (results []*EgCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_default IN (?)", idShopDefaults).Find(&results).Error
	
	return
}
 
// GetFromLevelDepth 通过level_depth获取内容  
func (obj *_EgCategoryMgr) GetFromLevelDepth(levelDepth uint8) (results []*EgCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("level_depth = ?", levelDepth).Find(&results).Error
	
	return
}

// GetBatchFromLevelDepth 批量唯一主键查找 
func (obj *_EgCategoryMgr) GetBatchFromLevelDepth(levelDepths []uint8) (results []*EgCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("level_depth IN (?)", levelDepths).Find(&results).Error
	
	return
}
 
// GetFromNleft 通过nleft获取内容  
func (obj *_EgCategoryMgr) GetFromNleft(nleft uint32) (results []*EgCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("nleft = ?", nleft).Find(&results).Error
	
	return
}

// GetBatchFromNleft 批量唯一主键查找 
func (obj *_EgCategoryMgr) GetBatchFromNleft(nlefts []uint32) (results []*EgCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("nleft IN (?)", nlefts).Find(&results).Error
	
	return
}
 
// GetFromNright 通过nright获取内容  
func (obj *_EgCategoryMgr) GetFromNright(nright uint32) (results []*EgCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("nright = ?", nright).Find(&results).Error
	
	return
}

// GetBatchFromNright 批量唯一主键查找 
func (obj *_EgCategoryMgr) GetBatchFromNright(nrights []uint32) (results []*EgCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("nright IN (?)", nrights).Find(&results).Error
	
	return
}
 
// GetFromActive 通过active获取内容  
func (obj *_EgCategoryMgr) GetFromActive(active bool) (results []*EgCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active = ?", active).Find(&results).Error
	
	return
}

// GetBatchFromActive 批量唯一主键查找 
func (obj *_EgCategoryMgr) GetBatchFromActive(actives []bool) (results []*EgCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active IN (?)", actives).Find(&results).Error
	
	return
}
 
// GetFromDateAdd 通过date_add获取内容  
func (obj *_EgCategoryMgr) GetFromDateAdd(dateAdd time.Time) (results []*EgCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error
	
	return
}

// GetBatchFromDateAdd 批量唯一主键查找 
func (obj *_EgCategoryMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*EgCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error
	
	return
}
 
// GetFromDateUpd 通过date_upd获取内容  
func (obj *_EgCategoryMgr) GetFromDateUpd(dateUpd time.Time) (results []*EgCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd = ?", dateUpd).Find(&results).Error
	
	return
}

// GetBatchFromDateUpd 批量唯一主键查找 
func (obj *_EgCategoryMgr) GetBatchFromDateUpd(dateUpds []time.Time) (results []*EgCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd IN (?)", dateUpds).Find(&results).Error
	
	return
}
 
// GetFromPosition 通过position获取内容  
func (obj *_EgCategoryMgr) GetFromPosition(position uint32) (results []*EgCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("position = ?", position).Find(&results).Error
	
	return
}

// GetBatchFromPosition 批量唯一主键查找 
func (obj *_EgCategoryMgr) GetBatchFromPosition(positions []uint32) (results []*EgCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("position IN (?)", positions).Find(&results).Error
	
	return
}
 
// GetFromIsRootCategory 通过is_root_category获取内容  
func (obj *_EgCategoryMgr) GetFromIsRootCategory(isRootCategory bool) (results []*EgCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_root_category = ?", isRootCategory).Find(&results).Error
	
	return
}

// GetBatchFromIsRootCategory 批量唯一主键查找 
func (obj *_EgCategoryMgr) GetBatchFromIsRootCategory(isRootCategorys []bool) (results []*EgCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_root_category IN (?)", isRootCategorys).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgCategoryMgr) FetchByPrimaryKey(idCategory uint32 ) (result EgCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_category = ?", idCategory).Find(&result).Error
	
	return
}
 

 
 // FetchIndexByCategoryParent  获取多个内容
 func (obj *_EgCategoryMgr) FetchIndexByCategoryParent(idParent uint32 ) (results []*EgCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_parent = ?", idParent).Find(&results).Error
	
	return
}
 
 // FetchIndexByLevelDepth  获取多个内容
 func (obj *_EgCategoryMgr) FetchIndexByLevelDepth(levelDepth uint8 ) (results []*EgCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("level_depth = ?", levelDepth).Find(&results).Error
	
	return
}
 
 // FetchIndexByNleftrightactive  获取多个内容
 func (obj *_EgCategoryMgr) FetchIndexByNleftrightactive(nleft uint32 ,nright uint32 ,active bool ) (results []*EgCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("nleft = ? AND nright = ? AND active = ?", nleft , nright , active).Find(&results).Error
	
	return
}
 
 // FetchIndexByActivenleft  获取多个内容
 func (obj *_EgCategoryMgr) FetchIndexByActivenleft(nleft uint32 ,active bool ) (results []*EgCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("nleft = ? AND active = ?", nleft , active).Find(&results).Error
	
	return
}
 
 // FetchIndexByNright  获取多个内容
 func (obj *_EgCategoryMgr) FetchIndexByNright(nright uint32 ) (results []*EgCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("nright = ?", nright).Find(&results).Error
	
	return
}
 
 // FetchIndexByActivenright  获取多个内容
 func (obj *_EgCategoryMgr) FetchIndexByActivenright(nright uint32 ,active bool ) (results []*EgCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("nright = ? AND active = ?", nright , active).Find(&results).Error
	
	return
}
 

	

