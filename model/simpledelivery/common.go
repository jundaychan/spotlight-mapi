package simpledelivery

// TimePeriodDTO 投放时段配置
// 每天24个字符，每个小时用0和1表示，0表示不投，1表示投放
type TimePeriodDTO struct {
	// Mon 星期一，默认24个1
	Mon string `json:"mon,omitempty"`
	// Tues 星期二
	Tues string `json:"tues,omitempty"`
	// Wed 星期三
	Wed string `json:"wed,omitempty"`
	// Thur 星期四
	Thur string `json:"thur,omitempty"`
	// Fri 星期五
	Fri string `json:"fri,omitempty"`
	// Sat 星期六
	Sat string `json:"sat,omitempty"`
	// Sun 星期日
	Sun string `json:"sun,omitempty"`
}

// KeywordWithBidDTO 关键词出价信息
type KeywordWithBidDTO struct {
	// Keyword 关键词
	Keyword string `json:"keyword,omitempty"`
	// Bid 出价，单位分
	Bid int64 `json:"bid,omitempty"`
	// PhraseMatchType 关键词匹配方式，0-精确匹配、1-短语匹配
	PhraseMatchType int `json:"phrase_match_type,omitempty"`
}

// PageDTO 实时数据分页信息
type PageDTO struct {
	// PageNum 页码
	PageNum int64 `json:"page_num,omitempty"`
	// PageSize 页大小
	PageSize int64 `json:"page_size,omitempty"`
	// TotalPage 总页数
	TotalPage int64 `json:"total_page,omitempty"`
	// TotalCount 总记录数
	TotalCount int64 `json:"total_count,omitempty"`
}
