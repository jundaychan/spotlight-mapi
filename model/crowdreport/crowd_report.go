package crowdreport

import (
	"github.com/jundaychan/spotlight-mapi/enum"
	"github.com/jundaychan/spotlight-mapi/model"
	"github.com/jundaychan/spotlight-mapi/util"
)

// ReportType 人群包报表查询维度
type ReportType string

const (
	// ReportType_USER_GROUP 人群包维度（默认）
	ReportType_USER_GROUP ReportType = "USER_GROUP"
	// ReportType_CAMPAIGN 计划维度
	ReportType_CAMPAIGN ReportType = "CAMPAIGN"
	// ReportType_UNIT 单元维度
	ReportType_UNIT ReportType = "UNIT"
	// ReportType_CREATIVITY 创意维度
	ReportType_CREATIVITY ReportType = "CREATIVITY"
	// ReportType_NOTE 笔记维度
	ReportType_NOTE ReportType = "NOTE"
	// ReportType_SPU SPU 维度
	ReportType_SPU ReportType = "SPU"
)

// FilterClause 人群包报表过滤条件
type FilterClause struct {
	// Column 需要进行筛选的列名; groupId/spuId/placement/marketingTarget
	Column string `json:"column,omitempty"`
	// Operator 对列的操作; 目前仅支持 in
	Operator string `json:"operator,omitempty"`
	// Values 列的筛选值
	Values []int `json:"values,omitempty"`
}

// SortClause 人群包报表排序条件
type SortClause struct {
	// Column 排序字段
	Column string `json:"column,omitempty"`
	// Sort 排序顺序; asc 升序 desc 降序
	Sort enum.SortType `json:"sort,omitempty"`
}

// CrowdReportRequest 人群包报表 API Request
type CrowdReportRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// StartDate 开始时间 YYYY-MM-DD
	StartDate string `json:"startDate,omitempty"`
	// EndDate 结束时间 YYYY-MM-DD
	EndDate string `json:"endDate,omitempty"`
	// TimeUnit 时间维度："DAY"：分天 "SUMMARY"：汇总,默认分天
	TimeUnit enum.TimeUnit `json:"timeUnit,omitempty"`
	// Columns 请求列，必须包含 "groupName", "groupId"
	Columns []string `json:"columns,omitempty"`
	// PageSize 页面大小; 默认20
	PageSize int `json:"pageSize,omitempty"`
	// PageNum 页面index; 默认1
	PageNum int `json:"pageNum,omitempty"`
	// Filters 过滤选项
	Filters []FilterClause `json:"filters,omitempty"`
	// Sorts 排序选项
	Sorts []SortClause `json:"sorts,omitempty"`
	// NeedTotal 是否需要综合数据
	NeedTotal bool `json:"needTotal,omitempty"`
	// ReportType 查询维度; 默认为 USER_GROUP
	ReportType ReportType `json:"reportType,omitempty"`
}

// Encode implement PostRequest interface
func (r CrowdReportRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// CrowdReportResponse 人群包报表 API Response
type CrowdReportResponse struct {
	model.BaseResponse
	// Page 页数信息
	Page *model.PageRespDTO `json:"page,omitempty"`
	// AggregationData 综合数据，key 为 ColumnName，value 为字符串
	AggregationData map[string]string `json:"aggregationData,omitempty"`
	// DataList 分时数据，每行 key 为 ColumnName，value 为字符串
	DataList []map[string]string `json:"dataList,omitempty"`
}
