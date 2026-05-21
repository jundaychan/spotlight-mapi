package targettemplate

import (
	"github.com/jundaychan/spotlight-mapi/model"
	"github.com/jundaychan/spotlight-mapi/util"
)

// CreateRequest 创建定向包 API Request
type CreateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Name 定向包名称，最多20个字符
	Name string `json:"name,omitempty"`
	// Desc 定向包描述，最多50个字符
	Desc string `json:"desc,omitempty"`
	// TargetType 定向类型，1-通投,2-智能定向, 3-高级定向
	TargetType TargetTemplateType `json:"target_type,omitempty"`
	// TargetConfig 定向配置
	TargetConfig *TargetConfig `json:"target_config,omitempty"`
	// MarketingTarget 营销诉求
	MarketingTarget MarketingTarget `json:"marketing_target,omitempty"`
	// Placement 广告类型 1-信息流2-搜索推广4-全站智投7-视频内流
	Placement Placement `json:"placement,omitempty"`
	// DeliveryMode 投放模式 0：手动投放，1：自动投放
	DeliveryMode DeliveryMode `json:"delivery_mode,omitempty"`
}

// Encode implement PostRequest interface
func (r CreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// CreateResponse 创建定向包 API Response
type CreateResponse struct {
	model.BaseResponse
	Data *CreateResult `json:"data,omitempty"`
}

// CreateResult 创建定向包结果
type CreateResult struct {
	// TargetTemplateID 定向包id
	TargetTemplateID uint64 `json:"target_template_id,omitempty"`
}
