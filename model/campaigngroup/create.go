package campaigngroup

import (
	"github.com/jundaychan/spotlight-mapi/model"
	"github.com/jundaychan/spotlight-mapi/util"
)

// CreateRequest 创建广告组 API Request
type CreateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// CampaignGroupName 广告组名称，最多50个字符
	CampaignGroupName string `json:"campaign_group_name,omitempty"`
	// LimitDayBudget 是否限制预算，0-不限制，1-限制
	LimitDayBudget int `json:"limit_day_budget,omitempty"`
	// OriginGroupDayBudget 单位分，最小值10000，最大值99999900
	OriginGroupDayBudget int64 `json:"origin_group_day_budget,omitempty"`
}

// Encode implement PostRequest interface
func (r CreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// CreateResponse 创建广告组 API Response
type CreateResponse struct {
	model.BaseResponse
	Data struct {
		// ID 广告组ID
		ID uint64 `json:"id,omitempty"`
	} `json:"data,omitempty"`
}
