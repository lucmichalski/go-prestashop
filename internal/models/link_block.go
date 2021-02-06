package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgLinkBlockMgr struct {
	*_BaseMgr
}

// EgLinkBlockMgr open func
func EgLinkBlockMgr(db *gorm.DB) *_EgLinkBlockMgr {
	if db == nil {
		panic(fmt.Errorf("EgLinkBlockMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgLinkBlockMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_link_block"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgLinkBlockMgr) GetTableName() string {
	return "eg_link_block"
}

// Get 获取
func (obj *_EgLinkBlockMgr) Get() (result EgLinkBlock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgLinkBlockMgr) Gets() (results []*EgLinkBlock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDLinkBlock id_link_block获取 
func (obj *_EgLinkBlockMgr) WithIDLinkBlock(idLinkBlock uint32) Option {
	return optionFunc(func(o *options) { o.query["id_link_block"] = idLinkBlock })
}

// WithIDHook id_hook获取 
func (obj *_EgLinkBlockMgr) WithIDHook(idHook uint32) Option {
	return optionFunc(func(o *options) { o.query["id_hook"] = idHook })
}

// WithPosition position获取 
func (obj *_EgLinkBlockMgr) WithPosition(position uint32) Option {
	return optionFunc(func(o *options) { o.query["position"] = position })
}

// WithContent content获取 
func (obj *_EgLinkBlockMgr) WithContent(content string) Option {
	return optionFunc(func(o *options) { o.query["content"] = content })
}


// GetByOption 功能选项模式获取
func (obj *_EgLinkBlockMgr) GetByOption(opts ...Option) (result EgLinkBlock, err error) {
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
func (obj *_EgLinkBlockMgr) GetByOptions(opts ...Option) (results []*EgLinkBlock, err error) {
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


// GetFromIDLinkBlock 通过id_link_block获取内容  
func (obj *_EgLinkBlockMgr)  GetFromIDLinkBlock(idLinkBlock uint32) (result EgLinkBlock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_link_block = ?", idLinkBlock).Find(&result).Error
	
	return
}

// GetBatchFromIDLinkBlock 批量唯一主键查找 
func (obj *_EgLinkBlockMgr) GetBatchFromIDLinkBlock(idLinkBlocks []uint32) (results []*EgLinkBlock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_link_block IN (?)", idLinkBlocks).Find(&results).Error
	
	return
}
 
// GetFromIDHook 通过id_hook获取内容  
func (obj *_EgLinkBlockMgr) GetFromIDHook(idHook uint32) (results []*EgLinkBlock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_hook = ?", idHook).Find(&results).Error
	
	return
}

// GetBatchFromIDHook 批量唯一主键查找 
func (obj *_EgLinkBlockMgr) GetBatchFromIDHook(idHooks []uint32) (results []*EgLinkBlock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_hook IN (?)", idHooks).Find(&results).Error
	
	return
}
 
// GetFromPosition 通过position获取内容  
func (obj *_EgLinkBlockMgr) GetFromPosition(position uint32) (results []*EgLinkBlock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("position = ?", position).Find(&results).Error
	
	return
}

// GetBatchFromPosition 批量唯一主键查找 
func (obj *_EgLinkBlockMgr) GetBatchFromPosition(positions []uint32) (results []*EgLinkBlock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("position IN (?)", positions).Find(&results).Error
	
	return
}
 
// GetFromContent 通过content获取内容  
func (obj *_EgLinkBlockMgr) GetFromContent(content string) (results []*EgLinkBlock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("content = ?", content).Find(&results).Error
	
	return
}

// GetBatchFromContent 批量唯一主键查找 
func (obj *_EgLinkBlockMgr) GetBatchFromContent(contents []string) (results []*EgLinkBlock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("content IN (?)", contents).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgLinkBlockMgr) FetchByPrimaryKey(idLinkBlock uint32 ) (result EgLinkBlock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_link_block = ?", idLinkBlock).Find(&result).Error
	
	return
}
 

 

	

