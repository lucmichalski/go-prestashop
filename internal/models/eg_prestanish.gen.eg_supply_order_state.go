package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgSupplyOrderStateMgr struct {
	*_BaseMgr
}

// EgSupplyOrderStateMgr open func
func EgSupplyOrderStateMgr(db *gorm.DB) *_EgSupplyOrderStateMgr {
	if db == nil {
		panic(fmt.Errorf("EgSupplyOrderStateMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgSupplyOrderStateMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_supply_order_state"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgSupplyOrderStateMgr) GetTableName() string {
	return "eg_supply_order_state"
}

// Get 获取
func (obj *_EgSupplyOrderStateMgr) Get() (result EgSupplyOrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgSupplyOrderStateMgr) Gets() (results []*EgSupplyOrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDSupplyOrderState id_supply_order_state获取 
func (obj *_EgSupplyOrderStateMgr) WithIDSupplyOrderState(idSupplyOrderState uint32) Option {
	return optionFunc(func(o *options) { o.query["id_supply_order_state"] = idSupplyOrderState })
}

// WithDeliveryNote delivery_note获取 
func (obj *_EgSupplyOrderStateMgr) WithDeliveryNote(deliveryNote bool) Option {
	return optionFunc(func(o *options) { o.query["delivery_note"] = deliveryNote })
}

// WithEditable editable获取 
func (obj *_EgSupplyOrderStateMgr) WithEditable(editable bool) Option {
	return optionFunc(func(o *options) { o.query["editable"] = editable })
}

// WithReceiptState receipt_state获取 
func (obj *_EgSupplyOrderStateMgr) WithReceiptState(receiptState bool) Option {
	return optionFunc(func(o *options) { o.query["receipt_state"] = receiptState })
}

// WithPendingReceipt pending_receipt获取 
func (obj *_EgSupplyOrderStateMgr) WithPendingReceipt(pendingReceipt bool) Option {
	return optionFunc(func(o *options) { o.query["pending_receipt"] = pendingReceipt })
}

// WithEnclosed enclosed获取 
func (obj *_EgSupplyOrderStateMgr) WithEnclosed(enclosed bool) Option {
	return optionFunc(func(o *options) { o.query["enclosed"] = enclosed })
}

// WithColor color获取 
func (obj *_EgSupplyOrderStateMgr) WithColor(color string) Option {
	return optionFunc(func(o *options) { o.query["color"] = color })
}


// GetByOption 功能选项模式获取
func (obj *_EgSupplyOrderStateMgr) GetByOption(opts ...Option) (result EgSupplyOrderState, err error) {
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
func (obj *_EgSupplyOrderStateMgr) GetByOptions(opts ...Option) (results []*EgSupplyOrderState, err error) {
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


// GetFromIDSupplyOrderState 通过id_supply_order_state获取内容  
func (obj *_EgSupplyOrderStateMgr)  GetFromIDSupplyOrderState(idSupplyOrderState uint32) (result EgSupplyOrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supply_order_state = ?", idSupplyOrderState).Find(&result).Error
	
	return
}

// GetBatchFromIDSupplyOrderState 批量唯一主键查找 
func (obj *_EgSupplyOrderStateMgr) GetBatchFromIDSupplyOrderState(idSupplyOrderStates []uint32) (results []*EgSupplyOrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supply_order_state IN (?)", idSupplyOrderStates).Find(&results).Error
	
	return
}
 
// GetFromDeliveryNote 通过delivery_note获取内容  
func (obj *_EgSupplyOrderStateMgr) GetFromDeliveryNote(deliveryNote bool) (results []*EgSupplyOrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("delivery_note = ?", deliveryNote).Find(&results).Error
	
	return
}

// GetBatchFromDeliveryNote 批量唯一主键查找 
func (obj *_EgSupplyOrderStateMgr) GetBatchFromDeliveryNote(deliveryNotes []bool) (results []*EgSupplyOrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("delivery_note IN (?)", deliveryNotes).Find(&results).Error
	
	return
}
 
// GetFromEditable 通过editable获取内容  
func (obj *_EgSupplyOrderStateMgr) GetFromEditable(editable bool) (results []*EgSupplyOrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("editable = ?", editable).Find(&results).Error
	
	return
}

// GetBatchFromEditable 批量唯一主键查找 
func (obj *_EgSupplyOrderStateMgr) GetBatchFromEditable(editables []bool) (results []*EgSupplyOrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("editable IN (?)", editables).Find(&results).Error
	
	return
}
 
// GetFromReceiptState 通过receipt_state获取内容  
func (obj *_EgSupplyOrderStateMgr) GetFromReceiptState(receiptState bool) (results []*EgSupplyOrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("receipt_state = ?", receiptState).Find(&results).Error
	
	return
}

// GetBatchFromReceiptState 批量唯一主键查找 
func (obj *_EgSupplyOrderStateMgr) GetBatchFromReceiptState(receiptStates []bool) (results []*EgSupplyOrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("receipt_state IN (?)", receiptStates).Find(&results).Error
	
	return
}
 
// GetFromPendingReceipt 通过pending_receipt获取内容  
func (obj *_EgSupplyOrderStateMgr) GetFromPendingReceipt(pendingReceipt bool) (results []*EgSupplyOrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("pending_receipt = ?", pendingReceipt).Find(&results).Error
	
	return
}

// GetBatchFromPendingReceipt 批量唯一主键查找 
func (obj *_EgSupplyOrderStateMgr) GetBatchFromPendingReceipt(pendingReceipts []bool) (results []*EgSupplyOrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("pending_receipt IN (?)", pendingReceipts).Find(&results).Error
	
	return
}
 
// GetFromEnclosed 通过enclosed获取内容  
func (obj *_EgSupplyOrderStateMgr) GetFromEnclosed(enclosed bool) (results []*EgSupplyOrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("enclosed = ?", enclosed).Find(&results).Error
	
	return
}

// GetBatchFromEnclosed 批量唯一主键查找 
func (obj *_EgSupplyOrderStateMgr) GetBatchFromEnclosed(encloseds []bool) (results []*EgSupplyOrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("enclosed IN (?)", encloseds).Find(&results).Error
	
	return
}
 
// GetFromColor 通过color获取内容  
func (obj *_EgSupplyOrderStateMgr) GetFromColor(color string) (results []*EgSupplyOrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("color = ?", color).Find(&results).Error
	
	return
}

// GetBatchFromColor 批量唯一主键查找 
func (obj *_EgSupplyOrderStateMgr) GetBatchFromColor(colors []string) (results []*EgSupplyOrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("color IN (?)", colors).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgSupplyOrderStateMgr) FetchByPrimaryKey(idSupplyOrderState uint32 ) (result EgSupplyOrderState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supply_order_state = ?", idSupplyOrderState).Find(&result).Error
	
	return
}
 

 

	

