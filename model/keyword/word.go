package keyword

type Word struct {
	// Keyword 词名词
	Keyword string `json:"keyword,omitempty"`
	// CompetitionLevel 竞争指数,示例：高
	CompetitionLevel string `json:"competition_level,omitempty"`
	// RecommendReason 推荐理由
	RecommendReason []string `json:"recommend_reason,omitempty"`
	// Bid 市场出价，单位：分
	Bid int64 `json:"bid"`
	// MonthPV 月均搜索指数(月pv)
	MonthPV int64 `json:"monthpv"`
	// Source 	词来源
	Source int `json:"source"`
}

// WordBag 词包信息
type WordBag struct {
	// Name 词包名称
	Name string `json:"name,omitempty"`
	// CreateAudit 创建账号名称
	CreateAudit string `json:"create_audit,omitempty"`
	// CreateTime 创建时间，示例："2023-12-26 15:50:43"
	CreateTime string `json:"create_time,omitempty"`
	// WordList   所有词信息
	WordList []Word `json:"word_list,omitempty"`
	// Source 词包来源1：自建2：平台
	Source int `json:"source"`
	// KeywordSource 词来源
	KeywordSource int `json:"keyword_source"`
}
