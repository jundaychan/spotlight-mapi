package targettemplate

import (
	"github.com/bububa/spotlight-mapi/model"
	"github.com/bububa/spotlight-mapi/util"
)

// ListRequest 获取定向包列表 API Request
type ListRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// TargetTemplateID 定向包id，当根据定向包ID查询时，其他入参无效
	TargetTemplateID uint64 `json:"target_template_id,omitempty"`
	// MarketingTarget 营销诉求
	MarketingTarget MarketingTarget `json:"marketing_target,omitempty"`
	// Placement 广告类型 1-信息流 2-搜索推广 4-全站智投 7-视频内流
	Placement Placement `json:"placement,omitempty"`
	// Page 请求的页码数，默认1
	Page int `json:"page,omitempty"`
	// PageSize 请求的每页行数，默认20，最大100
	PageSize int `json:"page_size,omitempty"`
}

// Encode implement PostRequest interface
func (r ListRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// ListResponse 获取定向包列表 API Response
type ListResponse struct {
	model.BaseResponse
	Data *ListResult `json:"data,omitempty"`
}

// ListResult 获取定向包列表结果
type ListResult struct {
	// DataList 定向包列表
	DataList []TargetTemplate `json:"data_list,omitempty"`
	// Page 分页信息
	Page *model.PageRespDTO `json:"page,omitempty"`
}

// TargetTemplate 定向包
type TargetTemplate struct {
	// ID 定向包ID
	ID uint64 `json:"id,omitempty"`
	// Name 定向包名称
	Name string `json:"name,omitempty"`
	// Desc 定向包描述
	Desc string `json:"desc,omitempty"`
	// TargetType 定向类型 1-通投 2-智能定向 3-高级定向
	TargetType TargetTemplateType `json:"target_type,omitempty"`
	// MarketingTarget 营销诉求
	MarketingTarget MarketingTarget `json:"marketing_target,omitempty"`
	// Placement 广告类型
	Placement Placement `json:"placement,omitempty"`
	// DeliveryMode 投放模式 0-手动投放 1-自动投放
	DeliveryMode DeliveryMode `json:"delivery_mode,omitempty"`
	// State 状态
	State int `json:"state,omitempty"`
	// UnitList 已关联该定向包的单元ID列表
	UnitList []uint64 `json:"unit_list,omitempty"`
	// TargetConfig 定向配置
	TargetConfig *TargetConfig `json:"target_config,omitempty"`
}
