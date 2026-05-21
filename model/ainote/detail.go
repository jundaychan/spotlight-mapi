package ainote

import (
	"github.com/jundaychan/spotlight-mapi/model"
	"github.com/jundaychan/spotlight-mapi/util"
)

// FilterClause AI 智能笔记详情过滤条件
type FilterClause struct {
	// Column 过滤字段，目前支持 unitId
	Column string `json:"column,omitempty"`
	// Operator 运算符，目前支持 in
	Operator string `json:"operator,omitempty"`
	// Values 过滤值列表
	Values []string `json:"values,omitempty"`
}

// DetailRequest AI 智能笔记详情接口 API Request
type DetailRequest struct {
	// AdvertiserID 广告主 ID
	AdvertiserID uint64 `json:"advertiser_id"`
	// StartDate 查询开始日期，格式 YYYY-MM-DD
	StartDate string `json:"start_date,omitempty"`
	// EndDate 查询结束日期，格式 YYYY-MM-DD
	EndDate string `json:"end_date,omitempty"`
	// Columns 需要返回的指标字段列表
	Columns []string `json:"columns,omitempty"`
	// Filters 过滤条件列表
	Filters []FilterClause `json:"filters,omitempty"`
}

// Encode implement PostRequest interface
func (r DetailRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// DetailResponse AI 智能笔记详情接口 API Response
type DetailResponse struct {
	model.BaseResponse
	// Data 数据
	Data *DetailData `json:"data,omitempty"`
}

// DetailData AI 智能笔记详情数据
type DetailData struct {
	// Page 分页信息
	Page *Page `json:"page,omitempty"`
	// TotalData 汇总数据，key 为请求的指标字段名，value 为字符串；额外含 time 字段
	TotalData map[string]string `json:"total_data,omitempty"`
	// DataList 明细数据，key 为请求的指标字段名，value 为字符串
	DataList []map[string]string `json:"data_list,omitempty"`
}

// Page 分页信息
type Page struct {
	// PageIndex 当前页码
	PageIndex int64 `json:"page_index,omitempty"`
	// TotalCount 数据总条数
	TotalCount int64 `json:"total_count,omitempty"`
	// PageSize 每页条数
	PageSize int64 `json:"page_size,omitempty"`
	// TotalPage 总页数
	TotalPage int64 `json:"total_page,omitempty"`
}
