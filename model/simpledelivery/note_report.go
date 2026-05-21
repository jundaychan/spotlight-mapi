package simpledelivery

import (
	"github.com/jundaychan/spotlight-mapi/enum"
	"github.com/jundaychan/spotlight-mapi/model"
	"github.com/jundaychan/spotlight-mapi/model/report"
	"github.com/jundaychan/spotlight-mapi/util"
)

// NoteReportRequest 简单投笔记报表 API Request（离线）
type NoteReportRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// StartDate 开始时间，格式 yyyy-MM-dd
	StartDate string `json:"start_date,omitempty"`
	// EndDate 结束时间，格式 yyyy-MM-dd
	EndDate string `json:"end_date,omitempty"`
	// TimeUnit 时间维度，DAY/HOUR/WEEK/MONTH/SUMMARY，默认分天
	TimeUnit enum.TimeUnit `json:"time_unit,omitempty"`
	// NoteIDs 笔记ID列表，最多100个
	NoteIDs []string `json:"note_ids,omitempty"`
	// SplitColumns 细分条件(group by)，campaignId/keyword/campaignGroupId
	SplitColumns []string `json:"split_columns,omitempty"`
	// PageNum 页数，默认1
	PageNum int64 `json:"page_num,omitempty"`
	// PageSize 页大小，默认20，最大500
	PageSize int64 `json:"page_size,omitempty"`
	// DataCaliber 数据指标归因时间类型，0-计费时间、1-转化时间，默认计费时间
	DataCaliber int `json:"data_caliber,omitempty"`
	// CreationType 托管模式，1-全自动、4-半自动，默认查全部
	CreationType []int `json:"creation_type,omitempty"`
}

// Encode implement PostRequest interface
func (r NoteReportRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// NoteReportResponse 简单投笔记报表 API Response
type NoteReportResponse struct {
	model.BaseResponse
	// Data 报表数据
	Data *NoteReportData `json:"data,omitempty"`
}

// NoteReportData 简单投笔记报表数据
type NoteReportData struct {
	// Page 分页数据
	Page *model.PageRespDTO `json:"page,omitempty"`
	// List 明细数据
	List []NoteReport `json:"data_list,omitempty"`
	// AggregationData 汇总数据
	AggregationData *NoteReport `json:"aggregation_data,omitempty"`
	// UnsupportedColumns 不支持返回的指标名
	UnsupportedColumns []string `json:"unsupported_columns,omitempty"`
}

// NoteReport 简单投笔记报表明细
type NoteReport struct {
	NoteReportDimension
	report.DataReportDTO
}

// NoteReportDimension 简单投笔记报表维度字段
type NoteReportDimension struct {
	// Time 时间
	Time string `json:"time,omitempty"`
	// CampaignID 计划ID，当细分条件选择标的时有该字段
	CampaignID model.Uint64 `json:"campaign_id,omitempty"`
	// CampaignName 计划名称，当细分条件选择标的时有该字段
	CampaignName string `json:"campaign_name,omitempty"`
	// CreativityID 创意ID，当细分条件选择计划时该字段值有效
	CreativityID model.Uint64 `json:"creativity_id,omitempty"`
	// CreativityName 创意名称，当细分条件选择计划时该字段值有效
	CreativityName string `json:"creativity_name,omitempty"`
	// NoteID 笔记ID
	NoteID string `json:"note_id,omitempty"`
	// NoteTitle 笔记标题
	NoteTitle string `json:"note_title,omitempty"`
}
