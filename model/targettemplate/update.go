package targettemplate

import (
	"github.com/bububa/spotlight-mapi/model"
	"github.com/bububa/spotlight-mapi/util"
)

// UpdateRequest 更新定向包 API Request
type UpdateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// TemplateList 定向包列表，目前只支持传一个
	TemplateList []UpdateTemplate `json:"template_list,omitempty"`
}

// UpdateTemplate 待更新的定向包
type UpdateTemplate struct {
	// TargetTemplateID 定向包ID
	TargetTemplateID uint64 `json:"target_template_id,omitempty"`
	// Name 定向包名称，最多20个字符
	Name string `json:"name,omitempty"`
	// Desc 定向包描述，最多50个字符
	Desc string `json:"desc,omitempty"`
	// MarketingTarget 营销诉求
	MarketingTarget MarketingTarget `json:"marketing_target,omitempty"`
	// Placement 广告类型 1-信息流2-搜索推广4-全站智投7-视频内流
	Placement Placement `json:"placement,omitempty"`
	// TargetType 定向类型，1-通投,2-智能定向, 3-高级定向
	TargetType TargetTemplateType `json:"target_type,omitempty"`
	// TargetConfig 定向配置
	TargetConfig *TargetConfig `json:"target_config,omitempty"`
}

// Encode implement PostRequest interface
func (r UpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// UpdateResponse 更新定向包 API Response
type UpdateResponse struct {
	model.BaseResponse
	Data *UpdateResult `json:"data,omitempty"`
}

// UpdateResult 更新定向包结果
type UpdateResult struct {
	// TemplateID 定向包id
	TemplateID uint64 `json:"template_id,omitempty"`
}
