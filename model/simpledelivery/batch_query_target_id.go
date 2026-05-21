package simpledelivery

import (
	"github.com/bububa/spotlight-mapi/model"
	"github.com/bububa/spotlight-mapi/util"
)

// BatchQueryTargetIDRequest 批量查询简单投标的ID API Request
type BatchQueryTargetIDRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// CampaignGroupIDList 广告组ID列表
	CampaignGroupIDList []uint64 `json:"campaign_group_id_list,omitempty"`
}

// Encode implement PostRequest interface
func (r BatchQueryTargetIDRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// BatchQueryTargetIDResponse 批量查询简单投标的ID API Response
type BatchQueryTargetIDResponse struct {
	model.BaseResponse
	// Data 标的数据
	Data *BatchQueryTargetIDData `json:"data,omitempty"`
}

// BatchQueryTargetIDData 标的数据
type BatchQueryTargetIDData struct {
	// ExtraDTO 标的数据集合
	ExtraDTO []ExtraDTO `json:"extra_dto,omitempty"`
}

// ExtraDTO 标的数据
type ExtraDTO struct {
	// UbeViewID 标的ID
	UbeViewID string `json:"ube_view_id,omitempty"`
	// GroupType 标的类型，1-全自动托管、2-半自动托管
	GroupType int `json:"group_type,omitempty"`
	// CampaignGroupID 广告组ID
	CampaignGroupID model.Uint64 `json:"campaign_group_id,omitempty"`
}
