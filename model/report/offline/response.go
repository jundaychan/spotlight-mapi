package offline

import (
	"github.com/jundaychan/spotlight-mapi/model"
	"github.com/jundaychan/spotlight-mapi/model/report"
)

// Response 离线数据 API Response
type Response struct {
	model.BaseResponse
	Data *ReportList `json:"data,omitempty"`
}

// ReportList 离线数据列表
type ReportList struct {
	List       []Report `json:"data_list,omitempty"`
	TotalCount int64    `json:"total_count,omitempty"`
}

// Report 离线数据报表
type Report struct {
	Dimension
	report.DataReportDTO
}

type Dimension struct {
	// Time 时间
	Time string `json:"time,omitempty"`
	// CampaignID 计划id
	CampaignID model.Uint64 `json:"campaign_id,omitempty"`
	// CampaignName 计划名称
	CampaignName string `json:"campaign_name,omitempty"`
	// UnitID 单元id
	UnitID model.Uint64 `json:"unit_id,omitempty"`
	// UnitName 单元名称
	UnitName string `json:"unit_name,omitempty"`
	// CreativityID 创意id
	CreativityID model.Uint64 `json:"creativity_id,omitempty"`
	// CreativityName 创意名称
	CreativityName string `json:"creativity_name,omitempty"`
	// CreativityImage 创意图片
	CreativityImage string `json:"creativity_image,omitempty"`
	// Keyword 关键词
	Keyword string `json:"keyword,omitempty"`
	// Placement 广告类型
	Placement model.Int `json:"placement,omitempty"`
	// OptimizeTarget 优化目标
	OptimizeTarget model.Int `json:"optimize_target,omitempty"`
	// PromotionTarget 推广标的
	PromotionTarget model.Int `json:"promotion_target,omitempty"`
	// BiddingStrategy 出价方式
	BiddingStrategy model.Int `json:"bidding_strategy,omitempty"`
	// BuildType 搭建类型
	BuildType model.Int `json:"build_type,omitempty"`
	// MarketingTarget 营销诉求
	MarketingTarget model.Int `json:"marketing_target,omitempty"`
	// NoteID 笔记id
	NoteID string `json:"note_id,omitempty"`
	// PageID 落地页id
	PageID string `json:"page_id,omitempty"`
	// ItemID 商品id
	ItemID string `json:"item_id,omitempty"`
	// LiveRedID 直播间id
	LiveRedID string `json:"live_red_id,omitempty"`
	// KeywordID 关键词id
	KeywordID uint64 `json:"keyword_id,omitempty"`
}

// UnmarshalJSON 必须自己实现：本结构体匿名内嵌了 report.DataReportDTO，
// 而它自带 UnmarshalJSON（兼容 ube/* 的 camelCase 指标键）。那个方法会被提升到
// 本类型上，导致标准 json.Unmarshal 只走它、**本结构体自己声明的字段全被静默留空**。
// report.UnmarshalEmbedded 会把两边都解出来，见该函数注释。
func (r *Report) UnmarshalJSON(b []byte) error {
	return report.UnmarshalEmbedded(b, r)
}
