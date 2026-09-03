package reportoffline2

import (
	"github.com/jundaychan/spotlight-mapi/enum"
	"github.com/jundaychan/spotlight-mapi/model"
	"github.com/jundaychan/spotlight-mapi/model/report"
	"github.com/jundaychan/spotlight-mapi/util"
)

// SearchWordRequest 搜索词层级离线报表数据 API Request
type SearchWordRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id"`
	// StartDate 开始时间，格式 yyyy-MM-dd
	StartDate string `json:"start_date,omitempty"`
	// EndDate 结束时间，格式 yyyy-MM-dd
	EndDate string `json:"end_date,omitempty"`
	// TimeUnit 时间维度："DAY"：分天 "HOUR"：分时 "SUMMARY"：汇总,默认分天
	TimeUnit enum.TimeUnit `json:"time_unit,omitempty"`
	// MarketingTarget 营销目标过滤条件
	MarketingTarget []int `json:"marketing_target,omitempty"`
	// BiddingStrategy 出价方式过滤条件：2：手动出价 3：自动出价
	BiddingStrategy []int `json:"bidding_strategy,omitempty"`
	// OptimizeTarget 推广目标过滤条件
	OptimizeTarget []int `json:"optimize_target,omitempty"`
	// Placement 广告类型过滤条件 1：信息流推广 2：搜索推广 4：全站智投 7：视频内流
	Placement []int `json:"placement,omitempty"`
	// PromotionTarget 推广标的类型过滤条件 1：笔记 2：商品 9：落地页 18：直播间
	PromotionTarget []int `json:"promotion_target,omitempty"`
	// Programmatic 创意组合方式过滤条件 0：自定义创意 1：程序化创意
	Programmatic []int `json:"programmatic,omitempty"`
	// BuildType 搭建方式过滤条件 0：标准搭建 1：省心智投
	BuildType []int `json:"build_type,omitempty"`
	// SortColumn 排序字段见附录column字段
	SortColumn string `json:"sort_column,omitempty"`
	// Sort 升降序asc：升序desc：降序
	Sort enum.SortType `json:"sort,omitempty"`
	// PageNum 页数，默认1
	PageNum int64 `json:"page_num,omitempty"`
	// PageSize 页大小，默认20,最大500
	PageSize int64 `json:"page_size,omitempty"`
	// DataCaliber 数据指标归因时间类型 0-计费时间 1-转化时间
	DataCaliber *int `json:"data_caliber,omitempty"`
}

// Encode implement PostRequest interface
func (r SearchWordRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// SearchWordResponse 搜索词层级离线报表数据 API Response
type SearchWordResponse struct {
	model.BaseResponse
	// Data 数据
	Data *SearchWordReportList `json:"data,omitempty"`
}

// SearchWordReportList 搜索词层级离线数据列表
type SearchWordReportList struct {
	// Page 分页信息
	Page *model.PageRespDTO `json:"page,omitempty"`
	// List 详细数据
	List []SearchWordReport `json:"data_list,omitempty"`
	// AggregationData 汇总数据
	AggregationData *SearchWordReport `json:"aggregation_data,omitempty"`
}

// SearchWordReport 搜索词层级离线数据报表
type SearchWordReport struct {
	// SearchWord 搜索词
	SearchWord string `json:"search_word,omitempty"`
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
	// NoteID 笔记id
	NoteID string `json:"note_id,omitempty"`
	// Time 时间
	Time string `json:"time,omitempty"`
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
	report.DataReportDTO
}

// UnmarshalJSON 必须自己实现：本结构体匿名内嵌了 report.DataReportDTO，
// 而它自带 UnmarshalJSON（兼容 ube/* 的 camelCase 指标键）。那个方法会被提升到
// 本类型上，导致标准 json.Unmarshal 只走它、**本结构体自己声明的字段全被静默留空**。
// report.UnmarshalEmbedded 会把两边都解出来，见该函数注释。
func (r *SearchWordReport) UnmarshalJSON(b []byte) error {
	return report.UnmarshalEmbedded(b, r)
}
