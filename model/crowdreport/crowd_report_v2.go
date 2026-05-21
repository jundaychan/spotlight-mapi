package crowdreport

import (
	"github.com/jundaychan/spotlight-mapi/enum"
	"github.com/jundaychan/spotlight-mapi/model"
	"github.com/jundaychan/spotlight-mapi/util"
)

// TimeUnitV2 人群包报表 V2 时间维度
type TimeUnitV2 string

const (
	// TimeUnitV2_DAY 分天
	TimeUnitV2_DAY TimeUnitV2 = "DAY"
	// TimeUnitV2_WEEK 分周
	TimeUnitV2_WEEK TimeUnitV2 = "WEEK"
	// TimeUnitV2_SUMMARY 汇总,默认分天
	TimeUnitV2_SUMMARY TimeUnitV2 = "SUMMARY"
)

// CrowdReportV2Request 人群包报表 V2 API Request
type CrowdReportV2Request struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiserId,omitempty"`
	// PageNum 页码
	PageNum int `json:"pageNum,omitempty"`
	// PageSize 分页 size
	PageSize int `json:"pageSize,omitempty"`
	// StartDate 开始时间，格式 yyyy-MM-dd
	StartDate string `json:"startDate,omitempty"`
	// EndDate 结束时间，格式 yyyy-MM-dd
	EndDate string `json:"endDate,omitempty"`
	// TimeUnit 时间维度："DAY"：分天 "WEEK"：分周 "SUMMARY"：汇总,默认分天
	TimeUnit TimeUnitV2 `json:"timeUnit,omitempty"`
	// ReportType 人群包报表类型：USER_GROUP、CAMPAIGN、UNIT、CREATIVITY、SPU、NOTE
	ReportType ReportType `json:"reportType,omitempty"`
	// DataCaliber 0-归因时间，1-转化时间
	DataCaliber int `json:"dataCaliber,omitempty"`
	// Sorts 排序选项
	Sorts []SortClauseV2 `json:"sorts,omitempty"`
	// Columns 指标列
	Columns []string `json:"columns,omitempty"`
	// SplitColumns 细分条件
	SplitColumns []string `json:"splitColumns,omitempty"`
	// GroupID 人群包 ID（过滤条件）
	GroupID uint64 `json:"groupId,omitempty"`
	// IsPremium 是否人群优投（过滤条件）; 是 否
	IsPremium []string `json:"isPremium,omitempty"`
	// MarketingTarget 营销目标过滤条件
	MarketingTarget []int `json:"marketingTarget,omitempty"`
	// BiddingStrategy 出价方式过滤条件：2：手动出价 3：自动出价 4：MCB 7：OCPX
	BiddingStrategy []int `json:"biddingStrategy,omitempty"`
	// OptimizeTarget 推广目标过滤条件
	OptimizeTarget []int `json:"optimizeTarget,omitempty"`
	// Placement 广告类型过滤条件 1：信息流 2：搜索 4：全站智投 7：视频流
	Placement []int `json:"placement,omitempty"`
	// ConstraintType 深度优化目标过滤条件
	ConstraintType []int `json:"constraintType,omitempty"`
	// DeliveryMode 投放模式过滤条件 0：手动投放 1：自动投放
	DeliveryMode []int `json:"deliveryMode,omitempty"`
}

// SortClauseV2 人群包报表 V2 排序条件
type SortClauseV2 struct {
	// Column 排序字段
	Column string `json:"column,omitempty"`
	// Sort 排序顺序; asc 升序 desc 降序
	Sort enum.SortType `json:"sort,omitempty"`
}

// Encode implement PostRequest interface
func (r CrowdReportV2Request) Encode() []byte {
	return util.JSONMarshal(r)
}

// CrowdReportV2Response 人群包报表 V2 API Response
type CrowdReportV2Response struct {
	model.BaseResponse
	// Data 数据
	Data *CrowdReportV2Data `json:"data,omitempty"`
}

// CrowdReportV2Data 人群包报表 V2 数据
type CrowdReportV2Data struct {
	// TotalCount 总条数
	TotalCount int64 `json:"total_count,omitempty"`
	// DataList 明细数据，每行 key 为指标 ename，value 为字符串
	DataList []map[string]string `json:"data_list,omitempty"`
	// AggregationData 综合数据，key 为指标 ename，value 为字符串
	AggregationData map[string]string `json:"aggregationData,omitempty"`
}
