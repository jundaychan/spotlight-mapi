package campaigngroup

import (
	"github.com/jundaychan/spotlight-mapi/model"
	"github.com/jundaychan/spotlight-mapi/util"
)

// UpdateRequest 更新广告组 API Request
type UpdateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Groups 广告组列表
	Groups []UpdateGroupDTO `json:"groups,omitempty"`
}

// UpdateGroupDTO 更新广告组项
type UpdateGroupDTO struct {
	// CampaignGroupID 广告组ID
	CampaignGroupID uint64 `json:"campaign_group_id,omitempty"`
	// Enable 启停状态，0-下线、1-上线
	Enable *int `json:"enable,omitempty"`
	// CampaignGroupName 广告组名称，最多50个字符
	CampaignGroupName string `json:"campaign_group_name,omitempty"`
	// LimitDayBudget 是否限制预算，0-不限制，1-限制
	LimitDayBudget *int `json:"limit_day_budget,omitempty"`
	// OriginGroupDayBudget 单位分，最小值10000，最大值99999900
	OriginGroupDayBudget int64 `json:"origin_group_day_budget,omitempty"`
}

// Encode implement PostRequest interface
func (r UpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// UpdateResponse 更新广告组 API Response
type UpdateResponse struct {
	model.BaseResponse
}
