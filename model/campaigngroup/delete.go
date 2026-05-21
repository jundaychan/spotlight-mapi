package campaigngroup

import (
	"github.com/jundaychan/spotlight-mapi/model"
	"github.com/jundaychan/spotlight-mapi/util"
)

// DeleteRequest 删除广告组 API Request
type DeleteRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Groups 广告组列表
	Groups []DeleteGroupDTO `json:"groups,omitempty"`
}

// DeleteGroupDTO 删除广告组项
type DeleteGroupDTO struct {
	// CampaignGroupID 广告组ID
	CampaignGroupID uint64 `json:"campaign_group_id,omitempty"`
	// State 传0
	State int `json:"state"`
}

// Encode implement PostRequest interface
func (r DeleteRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// DeleteResponse 删除广告组 API Response
type DeleteResponse struct {
	model.BaseResponse
}
