package finance

import (
	"github.com/bububa/spotlight-mapi/model"
	"github.com/bububa/spotlight-mapi/util"
)

// CampaignFlowRequest 获取计划流水 API Request
type CampaignFlowRequest struct {
	// StartTime 开始时间，格式需为yyyy-mm-dd
	StartTime string `json:"start_time,omitempty"`
	// EndTime 结束时间，格式需为yyyy-mm-dd
	EndTime string `json:"end_time,omitempty"`
	// Page 第几页
	Page int `json:"page,omitempty"`
	// PageSize 页大小
	PageSize int `json:"page_size,omitempty"`
	// AdvertiserID 品牌方ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
}

// Encode implement PostRequest interface
func (r CampaignFlowRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// CampaignFlowResponse 获取计划流水 API Response
type CampaignFlowResponse struct {
	model.BaseResponse
	Data *CampaignFlowResult `json:"data,omitempty"`
}

// CampaignFlowResult 计划流水数据
type CampaignFlowResult struct {
	// Total 总记录数
	Total int64 `json:"total,omitempty"`
	// Spend 消费返货金额总计，单位：分
	Spend int64 `json:"spend,omitempty"`
	// AdCampaignTradeDetail 计划流水详情
	AdCampaignTradeDetail []CampaignTradeDetail `json:"ad_campaign_trade_detail,omitempty"`
}

// CampaignTradeDetail 计划流水详情
type CampaignTradeDetail struct {
	// CampaignID 计划ID
	CampaignID uint64 `json:"campaign_id,omitempty"`
	// CampaignName 计划名称
	CampaignName string `json:"campaign_name,omitempty"`
	// LaunchDate 投放时间
	LaunchDate string `json:"launch_date,omitempty"`
	// PayTime 操作时间
	PayTime string `json:"pay_time,omitempty"`
	// CampaignDayBudget 日预算，单位：分
	CampaignDayBudget int64 `json:"campaign_day_budget,omitempty"`
	// OrderAmount 消耗金额，单位：分
	OrderAmount int64 `json:"order_amount,omitempty"`
}
