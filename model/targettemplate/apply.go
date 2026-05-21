package targettemplate

import (
	"github.com/jundaychan/spotlight-mapi/model"
	"github.com/jundaychan/spotlight-mapi/util"
)

// ApplyRequest 定向包关联单元 API Request
type ApplyRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// TargetTemplateID 定向包id
	TargetTemplateID uint64 `json:"target_template_id,omitempty"`
	// UnitIDs 单元 id 列表
	UnitIDs []uint64 `json:"unit_ids,omitempty"`
	// Action 关联：1 取消关联：2
	Action ApplyAction `json:"action,omitempty"`
}

// Encode implement PostRequest interface
func (r ApplyRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// ApplyResponse 定向包关联单元 API Response
type ApplyResponse struct {
	model.BaseResponse
	Data *ApplyResult `json:"data,omitempty"`
}

// ApplyResult 定向包关联单元结果
type ApplyResult struct {
	// FailUnitIDs 失败的单元id
	FailUnitIDs []uint64 `json:"fail_unit_ids,omitempty"`
}
