package finance

import (
	"github.com/bububa/spotlight-mapi/model"
	"github.com/bububa/spotlight-mapi/util"
)

// AgentSubAccountListRequest 获取代理商子账号列表 API Request
type AgentSubAccountListRequest struct {
	// UserID 代理商主账号ID
	UserID string `json:"user_id,omitempty"`
	// Page 第几页
	Page int `json:"page,omitempty"`
	// PageSize 页大小，上限500
	PageSize int `json:"page_size,omitempty"`
	// VirtualSellerID 子账户ID
	VirtualSellerID string `json:"virtual_seller_id,omitempty"`
	// CreateTimeStart 子账户开户开始时间，时间戳（ms）
	CreateTimeStart int64 `json:"create_time_start,omitempty"`
	// CreateTimeEnd 子账户开户结束时间，时间戳（ms）
	CreateTimeEnd int64 `json:"create_time_end,omitempty"`
}

// Encode implement PostRequest interface
func (r AgentSubAccountListRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// AgentSubAccountListResponse 获取代理商子账号列表 API Response
type AgentSubAccountListResponse struct {
	model.BaseResponse
	Data *AgentSubAccountListResult `json:"data,omitempty"`
}

// AgentSubAccountListResult 代理商子账号列表数据
type AgentSubAccountListResult struct {
	// Total 总记录数
	Total int64 `json:"total,omitempty"`
	// SubAccounts 子账号列表
	SubAccounts []AgentSubAccount `json:"sub_accounts,omitempty"`
}

// AgentSubAccount 子账号信息
type AgentSubAccount struct {
	// VirtualSellerID 子账户ID
	VirtualSellerID string `json:"virtual_seller_id,omitempty"`
	// VirtualSellerName 子账户名称
	VirtualSellerName string `json:"virtual_seller_name,omitempty"`
	// VirtualSellerStatus 子账号状态
	VirtualSellerStatus string `json:"virtual_seller_status,omitempty"`
	// BrandUserID 广告主userId
	BrandUserID string `json:"brand_user_id,omitempty"`
	// BrandUserName 广告主账户名称，K身份不返回本字段
	BrandUserName string `json:"brand_user_name,omitempty"`
	// BrandType 广告主账号类型，brand、kol
	BrandType string `json:"brand_type,omitempty"`
	// BrandStatus 广告主账号状态
	BrandStatus string `json:"brand_status,omitempty"`
	// CompanyName 广告主主体名称，K身份不返回本字段
	CompanyName string `json:"company_name,omitempty"`
	// AdvertiserID 投放账号ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// CreateTime 子账户创建时间，时间戳（ms）
	CreateTime int64 `json:"create_time,omitempty"`
}
