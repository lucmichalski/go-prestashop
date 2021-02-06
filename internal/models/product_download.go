package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
"time"	
)	

type _EgProductDownloadMgr struct {
	*_BaseMgr
}

// EgProductDownloadMgr open func
func EgProductDownloadMgr(db *gorm.DB) *_EgProductDownloadMgr {
	if db == nil {
		panic(fmt.Errorf("EgProductDownloadMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgProductDownloadMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_product_download"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgProductDownloadMgr) GetTableName() string {
	return "eg_product_download"
}

// Get 获取
func (obj *_EgProductDownloadMgr) Get() (result EgProductDownload, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgProductDownloadMgr) Gets() (results []*EgProductDownload, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDProductDownload id_product_download获取 
func (obj *_EgProductDownloadMgr) WithIDProductDownload(idProductDownload uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product_download"] = idProductDownload })
}

// WithIDProduct id_product获取 
func (obj *_EgProductDownloadMgr) WithIDProduct(idProduct uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product"] = idProduct })
}

// WithDisplayFilename display_filename获取 
func (obj *_EgProductDownloadMgr) WithDisplayFilename(displayFilename string) Option {
	return optionFunc(func(o *options) { o.query["display_filename"] = displayFilename })
}

// WithFilename filename获取 
func (obj *_EgProductDownloadMgr) WithFilename(filename string) Option {
	return optionFunc(func(o *options) { o.query["filename"] = filename })
}

// WithDateAdd date_add获取 
func (obj *_EgProductDownloadMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

// WithDateExpiration date_expiration获取 
func (obj *_EgProductDownloadMgr) WithDateExpiration(dateExpiration time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_expiration"] = dateExpiration })
}

// WithNbDaysAccessible nb_days_accessible获取 
func (obj *_EgProductDownloadMgr) WithNbDaysAccessible(nbDaysAccessible uint32) Option {
	return optionFunc(func(o *options) { o.query["nb_days_accessible"] = nbDaysAccessible })
}

// WithNbDownloadable nb_downloadable获取 
func (obj *_EgProductDownloadMgr) WithNbDownloadable(nbDownloadable uint32) Option {
	return optionFunc(func(o *options) { o.query["nb_downloadable"] = nbDownloadable })
}

// WithActive active获取 
func (obj *_EgProductDownloadMgr) WithActive(active bool) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}

// WithIsShareable is_shareable获取 
func (obj *_EgProductDownloadMgr) WithIsShareable(isShareable bool) Option {
	return optionFunc(func(o *options) { o.query["is_shareable"] = isShareable })
}


// GetByOption 功能选项模式获取
func (obj *_EgProductDownloadMgr) GetByOption(opts ...Option) (result EgProductDownload, err error) {
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
func (obj *_EgProductDownloadMgr) GetByOptions(opts ...Option) (results []*EgProductDownload, err error) {
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


// GetFromIDProductDownload 通过id_product_download获取内容  
func (obj *_EgProductDownloadMgr)  GetFromIDProductDownload(idProductDownload uint32) (result EgProductDownload, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_download = ?", idProductDownload).Find(&result).Error
	
	return
}

// GetBatchFromIDProductDownload 批量唯一主键查找 
func (obj *_EgProductDownloadMgr) GetBatchFromIDProductDownload(idProductDownloads []uint32) (results []*EgProductDownload, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_download IN (?)", idProductDownloads).Find(&results).Error
	
	return
}
 
// GetFromIDProduct 通过id_product获取内容  
func (obj *_EgProductDownloadMgr) GetFromIDProduct(idProduct uint32) (results []*EgProductDownload, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ?", idProduct).Find(&results).Error
	
	return
}

// GetBatchFromIDProduct 批量唯一主键查找 
func (obj *_EgProductDownloadMgr) GetBatchFromIDProduct(idProducts []uint32) (results []*EgProductDownload, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product IN (?)", idProducts).Find(&results).Error
	
	return
}
 
// GetFromDisplayFilename 通过display_filename获取内容  
func (obj *_EgProductDownloadMgr) GetFromDisplayFilename(displayFilename string) (results []*EgProductDownload, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("display_filename = ?", displayFilename).Find(&results).Error
	
	return
}

// GetBatchFromDisplayFilename 批量唯一主键查找 
func (obj *_EgProductDownloadMgr) GetBatchFromDisplayFilename(displayFilenames []string) (results []*EgProductDownload, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("display_filename IN (?)", displayFilenames).Find(&results).Error
	
	return
}
 
// GetFromFilename 通过filename获取内容  
func (obj *_EgProductDownloadMgr) GetFromFilename(filename string) (results []*EgProductDownload, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("filename = ?", filename).Find(&results).Error
	
	return
}

// GetBatchFromFilename 批量唯一主键查找 
func (obj *_EgProductDownloadMgr) GetBatchFromFilename(filenames []string) (results []*EgProductDownload, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("filename IN (?)", filenames).Find(&results).Error
	
	return
}
 
// GetFromDateAdd 通过date_add获取内容  
func (obj *_EgProductDownloadMgr) GetFromDateAdd(dateAdd time.Time) (results []*EgProductDownload, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error
	
	return
}

// GetBatchFromDateAdd 批量唯一主键查找 
func (obj *_EgProductDownloadMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*EgProductDownload, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error
	
	return
}
 
// GetFromDateExpiration 通过date_expiration获取内容  
func (obj *_EgProductDownloadMgr) GetFromDateExpiration(dateExpiration time.Time) (results []*EgProductDownload, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_expiration = ?", dateExpiration).Find(&results).Error
	
	return
}

// GetBatchFromDateExpiration 批量唯一主键查找 
func (obj *_EgProductDownloadMgr) GetBatchFromDateExpiration(dateExpirations []time.Time) (results []*EgProductDownload, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_expiration IN (?)", dateExpirations).Find(&results).Error
	
	return
}
 
// GetFromNbDaysAccessible 通过nb_days_accessible获取内容  
func (obj *_EgProductDownloadMgr) GetFromNbDaysAccessible(nbDaysAccessible uint32) (results []*EgProductDownload, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("nb_days_accessible = ?", nbDaysAccessible).Find(&results).Error
	
	return
}

// GetBatchFromNbDaysAccessible 批量唯一主键查找 
func (obj *_EgProductDownloadMgr) GetBatchFromNbDaysAccessible(nbDaysAccessibles []uint32) (results []*EgProductDownload, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("nb_days_accessible IN (?)", nbDaysAccessibles).Find(&results).Error
	
	return
}
 
// GetFromNbDownloadable 通过nb_downloadable获取内容  
func (obj *_EgProductDownloadMgr) GetFromNbDownloadable(nbDownloadable uint32) (results []*EgProductDownload, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("nb_downloadable = ?", nbDownloadable).Find(&results).Error
	
	return
}

// GetBatchFromNbDownloadable 批量唯一主键查找 
func (obj *_EgProductDownloadMgr) GetBatchFromNbDownloadable(nbDownloadables []uint32) (results []*EgProductDownload, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("nb_downloadable IN (?)", nbDownloadables).Find(&results).Error
	
	return
}
 
// GetFromActive 通过active获取内容  
func (obj *_EgProductDownloadMgr) GetFromActive(active bool) (results []*EgProductDownload, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active = ?", active).Find(&results).Error
	
	return
}

// GetBatchFromActive 批量唯一主键查找 
func (obj *_EgProductDownloadMgr) GetBatchFromActive(actives []bool) (results []*EgProductDownload, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active IN (?)", actives).Find(&results).Error
	
	return
}
 
// GetFromIsShareable 通过is_shareable获取内容  
func (obj *_EgProductDownloadMgr) GetFromIsShareable(isShareable bool) (results []*EgProductDownload, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_shareable = ?", isShareable).Find(&results).Error
	
	return
}

// GetBatchFromIsShareable 批量唯一主键查找 
func (obj *_EgProductDownloadMgr) GetBatchFromIsShareable(isShareables []bool) (results []*EgProductDownload, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_shareable IN (?)", isShareables).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgProductDownloadMgr) FetchByPrimaryKey(idProductDownload uint32 ) (result EgProductDownload, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_download = ?", idProductDownload).Find(&result).Error
	
	return
}
 

 

	

