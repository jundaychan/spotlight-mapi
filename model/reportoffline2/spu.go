package reportoffline2

import (
	"github.com/jundaychan/spotlight-mapi/enum"
	"github.com/jundaychan/spotlight-mapi/model"
	"github.com/jundaychan/spotlight-mapi/model/report"
	"github.com/jundaychan/spotlight-mapi/util"
)

// SPURequest SPU层级离线报表数据 API Request
type SPURequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id"`
	// StartDate 开始时间，格式 yyyy-MM-dd
	StartDate string `json:"start_date,omitempty"`
	// EndDate 结束时间，格式 yyyy-MM-dd
	EndDate string `json:"end_date,omitempty"`
	// TimeUnit 时间维度："DAY"：分天 "SUMMARY"：汇总,默认分天
	TimeUnit enum.TimeUnit `json:"time_unit,omitempty"`
	// SortColumn 排序字段见附录column字段
	SortColumn string `json:"sort_column,omitempty"`
	// Sort 升降序asc：升序desc：降序
	Sort enum.SortType `json:"sort,omitempty"`
	// PageNum 页数，默认1
	PageNum int64 `json:"page_num,omitempty"`
	// PageSize 页大小，默认20,最大500
	PageSize int64 `json:"page_size,omitempty"`
}

// Encode implement PostRequest interface
func (r SPURequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// SPUResponse SPU层级离线报表数据 API Response
type SPUResponse struct {
	model.BaseResponse
	// Data 数据
	Data *SPUReportList `json:"data,omitempty"`
}

// SPUReportList SPU层级离线数据列表
type SPUReportList struct {
	// TotalCount 总条数
	TotalCount int64 `json:"total_count,omitempty"`
	// List 详细数据
	List []SPUReport `json:"data_list,omitempty"`
	// AggregationData 汇总数据
	AggregationData *SPUReport `json:"aggregation_data,omitempty"`
}

// SPUReport SPU层级离线数据报表
type SPUReport struct {
	// SpuID spu id
	SpuID string `json:"spu_id,omitempty"`
	// SpuName spu名称
	SpuName string `json:"spu_name,omitempty"`
	// Time 时间
	Time string `json:"time,omitempty"`
	report.DataReportDTO
	// IUserNum 新增种草人群; 用户点击推广后转化为相应SPU的I或TI人群的总人数
	IUserNum model.Int64 `json:"i_user_num,omitempty"`
	// TiUserNum 新增深度种草人群; 用户点击推广后转化为相应SPU的TI人群的总人数
	TiUserNum model.Int64 `json:"ti_user_num,omitempty"`
	// IUserPrice 新增种草人群成本; 推广消费/新增种草人群
	IUserPrice model.Float64 `json:"i_user_price,omitempty"`
	// TiUserPrice 新增深度种草人群成本; 推广消费/新增深度种草人群
	TiUserPrice model.Float64 `json:"ti_user_price,omitempty"`
}

// UnmarshalJSON 必须自己实现：本结构体匿名内嵌了 report.DataReportDTO，
// 而它自带 UnmarshalJSON（兼容 ube/* 的 camelCase 指标键）。那个方法会被提升到
// 本类型上，导致标准 json.Unmarshal 只走它、**本结构体自己声明的字段全被静默留空**。
// report.UnmarshalEmbedded 会把两边都解出来，见该函数注释。
func (r *SPUReport) UnmarshalJSON(b []byte) error {
	return report.UnmarshalEmbedded(b, r)
}
