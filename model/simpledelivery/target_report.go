package simpledelivery

import (
	"github.com/jundaychan/spotlight-mapi/enum"
	"github.com/jundaychan/spotlight-mapi/model"
	"github.com/jundaychan/spotlight-mapi/model/report"
	"github.com/jundaychan/spotlight-mapi/util"
)

// TargetReportRequest 简单投标的报表 API Request（离线）
type TargetReportRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// StartDate 开始时间，格式 yyyy-MM-dd
	StartDate string `json:"start_date,omitempty"`
	// EndDate 结束时间，格式 yyyy-MM-dd
	EndDate string `json:"end_date,omitempty"`
	// TimeUnit 时间维度，DAY/HOUR/WEEK/MONTH/SUMMARY，默认分天
	TimeUnit enum.TimeUnit `json:"time_unit,omitempty"`
	// CreationType 投放模式，1-全自动、4-半自动
	CreationType []int `json:"creation_type,omitempty"`
	// MarketingScene 营销场景过滤条件，1-产品种草、2-客资收集、3-种草直达-CID、4-种草直达-UD、5-应用推广
	MarketingScene []int `json:"marketing_scene,omitempty"`
	// OptimizeObjective 优化目标过滤条件，当 marketing_scene 不为空时必须为空
	OptimizeObjective []int `json:"optimize_objective,omitempty"`
	// SplitColumns 细分条件(group by)，noteId/campaignId/countryName/city/gender/age/keyword
	SplitColumns []string `json:"split_columns,omitempty"`
	// PageNum 页数，默认1
	PageNum int64 `json:"page_num,omitempty"`
	// PageSize 页大小，默认20，最大500
	PageSize int64 `json:"page_size,omitempty"`
	// DataCaliber 数据指标归因时间类型，0-计费时间、1-转化时间，默认计费时间
	DataCaliber *int `json:"data_caliber,omitempty"`
}

// Encode implement PostRequest interface
func (r TargetReportRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// TargetReportResponse 简单投标的报表 API Response
type TargetReportResponse struct {
	model.BaseResponse
	// Data 报表数据
	Data *TargetReportData `json:"data,omitempty"`
}

// TargetReportData 简单投标的报表数据
type TargetReportData struct {
	// Page 分页数据
	Page *model.PageRespDTO `json:"page,omitempty"`
	// List 明细数据
	List []TargetReport `json:"data_list,omitempty"`
	// AggregationData 汇总数据
	AggregationData *TargetReport `json:"aggregation_data,omitempty"`
	// UnsupportedColumns 不支持返回的指标名
	UnsupportedColumns []string `json:"unsupported_columns,omitempty"`
}

// TargetReport 简单投标的报表明细
type TargetReport struct {
	TargetReportDimension
	report.DataReportDTO
}

// TargetReportDimension 简单投标的报表维度字段
type TargetReportDimension struct {
	// Time 时间
	Time string `json:"time,omitempty"`
	// UbeViewID 标的ID
	UbeViewID string `json:"ube_view_id,omitempty"`
	// CampaignGroupName 标的名称
	CampaignGroupName string `json:"campaign_group_name,omitempty"`
	// NoteID 笔记ID，当细分条件选择笔记时有该字段
	NoteID string `json:"note_id,omitempty"`
	// NoteTitle 笔记标题，当细分条件选择笔记时有该字段
	NoteTitle string `json:"note_title,omitempty"`
	// CampaignID 计划ID，当细分条件选择计划时有该字段
	CampaignID model.Uint64 `json:"campaign_id,omitempty"`
}

// UnmarshalJSON 必须自己实现：本结构体匿名内嵌了 report.DataReportDTO，
// 而它自带 UnmarshalJSON（兼容 ube/* 的 camelCase 指标键）。那个方法会被提升到
// 本类型上，导致标准 json.Unmarshal 只走它、**本结构体自己声明的字段全被静默留空**。
// report.UnmarshalEmbedded 会把两边都解出来，见该函数注释。
func (r *TargetReport) UnmarshalJSON(b []byte) error {
	return report.UnmarshalEmbedded(b, r)
}
