package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
"time"	
)	

type _EgReferrerMgr struct {
	*_BaseMgr
}

// EgReferrerMgr open func
func EgReferrerMgr(db *gorm.DB) *_EgReferrerMgr {
	if db == nil {
		panic(fmt.Errorf("EgReferrerMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgReferrerMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_referrer"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgReferrerMgr) GetTableName() string {
	return "eg_referrer"
}

// Get 获取
func (obj *_EgReferrerMgr) Get() (result EgReferrer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgReferrerMgr) Gets() (results []*EgReferrer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDReferrer id_referrer获取 
func (obj *_EgReferrerMgr) WithIDReferrer(idReferrer uint32) Option {
	return optionFunc(func(o *options) { o.query["id_referrer"] = idReferrer })
}

// WithName name获取 
func (obj *_EgReferrerMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithPasswd passwd获取 
func (obj *_EgReferrerMgr) WithPasswd(passwd string) Option {
	return optionFunc(func(o *options) { o.query["passwd"] = passwd })
}

// WithHTTPRefererRegexp http_referer_regexp获取 
func (obj *_EgReferrerMgr) WithHTTPRefererRegexp(httpRefererRegexp string) Option {
	return optionFunc(func(o *options) { o.query["http_referer_regexp"] = httpRefererRegexp })
}

// WithHTTPRefererLike http_referer_like获取 
func (obj *_EgReferrerMgr) WithHTTPRefererLike(httpRefererLike string) Option {
	return optionFunc(func(o *options) { o.query["http_referer_like"] = httpRefererLike })
}

// WithRequestURIRegexp request_uri_regexp获取 
func (obj *_EgReferrerMgr) WithRequestURIRegexp(requestURIRegexp string) Option {
	return optionFunc(func(o *options) { o.query["request_uri_regexp"] = requestURIRegexp })
}

// WithRequestURILike request_uri_like获取 
func (obj *_EgReferrerMgr) WithRequestURILike(requestURILike string) Option {
	return optionFunc(func(o *options) { o.query["request_uri_like"] = requestURILike })
}

// WithHTTPRefererRegexpNot http_referer_regexp_not获取 
func (obj *_EgReferrerMgr) WithHTTPRefererRegexpNot(httpRefererRegexpNot string) Option {
	return optionFunc(func(o *options) { o.query["http_referer_regexp_not"] = httpRefererRegexpNot })
}

// WithHTTPRefererLikeNot http_referer_like_not获取 
func (obj *_EgReferrerMgr) WithHTTPRefererLikeNot(httpRefererLikeNot string) Option {
	return optionFunc(func(o *options) { o.query["http_referer_like_not"] = httpRefererLikeNot })
}

// WithRequestURIRegexpNot request_uri_regexp_not获取 
func (obj *_EgReferrerMgr) WithRequestURIRegexpNot(requestURIRegexpNot string) Option {
	return optionFunc(func(o *options) { o.query["request_uri_regexp_not"] = requestURIRegexpNot })
}

// WithRequestURILikeNot request_uri_like_not获取 
func (obj *_EgReferrerMgr) WithRequestURILikeNot(requestURILikeNot string) Option {
	return optionFunc(func(o *options) { o.query["request_uri_like_not"] = requestURILikeNot })
}

// WithBaseFee base_fee获取 
func (obj *_EgReferrerMgr) WithBaseFee(baseFee float64) Option {
	return optionFunc(func(o *options) { o.query["base_fee"] = baseFee })
}

// WithPercentFee percent_fee获取 
func (obj *_EgReferrerMgr) WithPercentFee(percentFee float64) Option {
	return optionFunc(func(o *options) { o.query["percent_fee"] = percentFee })
}

// WithClickFee click_fee获取 
func (obj *_EgReferrerMgr) WithClickFee(clickFee float64) Option {
	return optionFunc(func(o *options) { o.query["click_fee"] = clickFee })
}

// WithDateAdd date_add获取 
func (obj *_EgReferrerMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}


// GetByOption 功能选项模式获取
func (obj *_EgReferrerMgr) GetByOption(opts ...Option) (result EgReferrer, err error) {
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
func (obj *_EgReferrerMgr) GetByOptions(opts ...Option) (results []*EgReferrer, err error) {
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


// GetFromIDReferrer 通过id_referrer获取内容  
func (obj *_EgReferrerMgr)  GetFromIDReferrer(idReferrer uint32) (result EgReferrer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_referrer = ?", idReferrer).Find(&result).Error
	
	return
}

// GetBatchFromIDReferrer 批量唯一主键查找 
func (obj *_EgReferrerMgr) GetBatchFromIDReferrer(idReferrers []uint32) (results []*EgReferrer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_referrer IN (?)", idReferrers).Find(&results).Error
	
	return
}
 
// GetFromName 通过name获取内容  
func (obj *_EgReferrerMgr) GetFromName(name string) (results []*EgReferrer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error
	
	return
}

// GetBatchFromName 批量唯一主键查找 
func (obj *_EgReferrerMgr) GetBatchFromName(names []string) (results []*EgReferrer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error
	
	return
}
 
// GetFromPasswd 通过passwd获取内容  
func (obj *_EgReferrerMgr) GetFromPasswd(passwd string) (results []*EgReferrer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("passwd = ?", passwd).Find(&results).Error
	
	return
}

// GetBatchFromPasswd 批量唯一主键查找 
func (obj *_EgReferrerMgr) GetBatchFromPasswd(passwds []string) (results []*EgReferrer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("passwd IN (?)", passwds).Find(&results).Error
	
	return
}
 
// GetFromHTTPRefererRegexp 通过http_referer_regexp获取内容  
func (obj *_EgReferrerMgr) GetFromHTTPRefererRegexp(httpRefererRegexp string) (results []*EgReferrer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("http_referer_regexp = ?", httpRefererRegexp).Find(&results).Error
	
	return
}

// GetBatchFromHTTPRefererRegexp 批量唯一主键查找 
func (obj *_EgReferrerMgr) GetBatchFromHTTPRefererRegexp(httpRefererRegexps []string) (results []*EgReferrer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("http_referer_regexp IN (?)", httpRefererRegexps).Find(&results).Error
	
	return
}
 
// GetFromHTTPRefererLike 通过http_referer_like获取内容  
func (obj *_EgReferrerMgr) GetFromHTTPRefererLike(httpRefererLike string) (results []*EgReferrer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("http_referer_like = ?", httpRefererLike).Find(&results).Error
	
	return
}

// GetBatchFromHTTPRefererLike 批量唯一主键查找 
func (obj *_EgReferrerMgr) GetBatchFromHTTPRefererLike(httpRefererLikes []string) (results []*EgReferrer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("http_referer_like IN (?)", httpRefererLikes).Find(&results).Error
	
	return
}
 
// GetFromRequestURIRegexp 通过request_uri_regexp获取内容  
func (obj *_EgReferrerMgr) GetFromRequestURIRegexp(requestURIRegexp string) (results []*EgReferrer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("request_uri_regexp = ?", requestURIRegexp).Find(&results).Error
	
	return
}

// GetBatchFromRequestURIRegexp 批量唯一主键查找 
func (obj *_EgReferrerMgr) GetBatchFromRequestURIRegexp(requestURIRegexps []string) (results []*EgReferrer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("request_uri_regexp IN (?)", requestURIRegexps).Find(&results).Error
	
	return
}
 
// GetFromRequestURILike 通过request_uri_like获取内容  
func (obj *_EgReferrerMgr) GetFromRequestURILike(requestURILike string) (results []*EgReferrer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("request_uri_like = ?", requestURILike).Find(&results).Error
	
	return
}

// GetBatchFromRequestURILike 批量唯一主键查找 
func (obj *_EgReferrerMgr) GetBatchFromRequestURILike(requestURILikes []string) (results []*EgReferrer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("request_uri_like IN (?)", requestURILikes).Find(&results).Error
	
	return
}
 
// GetFromHTTPRefererRegexpNot 通过http_referer_regexp_not获取内容  
func (obj *_EgReferrerMgr) GetFromHTTPRefererRegexpNot(httpRefererRegexpNot string) (results []*EgReferrer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("http_referer_regexp_not = ?", httpRefererRegexpNot).Find(&results).Error
	
	return
}

// GetBatchFromHTTPRefererRegexpNot 批量唯一主键查找 
func (obj *_EgReferrerMgr) GetBatchFromHTTPRefererRegexpNot(httpRefererRegexpNots []string) (results []*EgReferrer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("http_referer_regexp_not IN (?)", httpRefererRegexpNots).Find(&results).Error
	
	return
}
 
// GetFromHTTPRefererLikeNot 通过http_referer_like_not获取内容  
func (obj *_EgReferrerMgr) GetFromHTTPRefererLikeNot(httpRefererLikeNot string) (results []*EgReferrer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("http_referer_like_not = ?", httpRefererLikeNot).Find(&results).Error
	
	return
}

// GetBatchFromHTTPRefererLikeNot 批量唯一主键查找 
func (obj *_EgReferrerMgr) GetBatchFromHTTPRefererLikeNot(httpRefererLikeNots []string) (results []*EgReferrer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("http_referer_like_not IN (?)", httpRefererLikeNots).Find(&results).Error
	
	return
}
 
// GetFromRequestURIRegexpNot 通过request_uri_regexp_not获取内容  
func (obj *_EgReferrerMgr) GetFromRequestURIRegexpNot(requestURIRegexpNot string) (results []*EgReferrer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("request_uri_regexp_not = ?", requestURIRegexpNot).Find(&results).Error
	
	return
}

// GetBatchFromRequestURIRegexpNot 批量唯一主键查找 
func (obj *_EgReferrerMgr) GetBatchFromRequestURIRegexpNot(requestURIRegexpNots []string) (results []*EgReferrer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("request_uri_regexp_not IN (?)", requestURIRegexpNots).Find(&results).Error
	
	return
}
 
// GetFromRequestURILikeNot 通过request_uri_like_not获取内容  
func (obj *_EgReferrerMgr) GetFromRequestURILikeNot(requestURILikeNot string) (results []*EgReferrer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("request_uri_like_not = ?", requestURILikeNot).Find(&results).Error
	
	return
}

// GetBatchFromRequestURILikeNot 批量唯一主键查找 
func (obj *_EgReferrerMgr) GetBatchFromRequestURILikeNot(requestURILikeNots []string) (results []*EgReferrer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("request_uri_like_not IN (?)", requestURILikeNots).Find(&results).Error
	
	return
}
 
// GetFromBaseFee 通过base_fee获取内容  
func (obj *_EgReferrerMgr) GetFromBaseFee(baseFee float64) (results []*EgReferrer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("base_fee = ?", baseFee).Find(&results).Error
	
	return
}

// GetBatchFromBaseFee 批量唯一主键查找 
func (obj *_EgReferrerMgr) GetBatchFromBaseFee(baseFees []float64) (results []*EgReferrer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("base_fee IN (?)", baseFees).Find(&results).Error
	
	return
}
 
// GetFromPercentFee 通过percent_fee获取内容  
func (obj *_EgReferrerMgr) GetFromPercentFee(percentFee float64) (results []*EgReferrer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("percent_fee = ?", percentFee).Find(&results).Error
	
	return
}

// GetBatchFromPercentFee 批量唯一主键查找 
func (obj *_EgReferrerMgr) GetBatchFromPercentFee(percentFees []float64) (results []*EgReferrer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("percent_fee IN (?)", percentFees).Find(&results).Error
	
	return
}
 
// GetFromClickFee 通过click_fee获取内容  
func (obj *_EgReferrerMgr) GetFromClickFee(clickFee float64) (results []*EgReferrer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("click_fee = ?", clickFee).Find(&results).Error
	
	return
}

// GetBatchFromClickFee 批量唯一主键查找 
func (obj *_EgReferrerMgr) GetBatchFromClickFee(clickFees []float64) (results []*EgReferrer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("click_fee IN (?)", clickFees).Find(&results).Error
	
	return
}
 
// GetFromDateAdd 通过date_add获取内容  
func (obj *_EgReferrerMgr) GetFromDateAdd(dateAdd time.Time) (results []*EgReferrer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error
	
	return
}

// GetBatchFromDateAdd 批量唯一主键查找 
func (obj *_EgReferrerMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*EgReferrer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgReferrerMgr) FetchByPrimaryKey(idReferrer uint32 ) (result EgReferrer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_referrer = ?", idReferrer).Find(&result).Error
	
	return
}
 

 

	
