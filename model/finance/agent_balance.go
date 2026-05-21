package finance

import (
	"github.com/bububa/spotlight-mapi/model"
	"github.com/bububa/spotlight-mapi/util"
)

// WalletAccountStatus 钱包账户状态
type WalletAccountStatus int

const (
	// WalletAccountStatusOpen 正常
	WalletAccountStatusOpen WalletAccountStatus = 1
	// WalletAccountStatusFrozen 冻结
	WalletAccountStatusFrozen WalletAccountStatus = 2
)

// AgentBalanceRequest 代理商主子账号余额查询 API Request
type AgentBalanceRequest struct {
	// UserID 代理商主账号ID
	UserID string `json:"user_id,omitempty"`
	// VirtualSellerIDList 子账号ID/代理商vsellerId，最多传100个
	VirtualSellerIDList []string `json:"virtual_seller_id_list,omitempty"`
}

// Encode implement PostRequest interface
func (r AgentBalanceRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// AgentBalanceResponse 代理商主子账号余额查询 API Response
type AgentBalanceResponse struct {
	model.BaseResponse
	Data *AgentBalanceResult `json:"data,omitempty"`
}

// AgentBalanceResult 代理商主子账号余额数据
type AgentBalanceResult struct {
	// Total 总条数
	Total int64 `json:"total,omitempty"`
	// WalletBalanceList 分账户余额信息
	WalletBalanceList []WalletBalance `json:"wallet_balance_list,omitempty"`
}

// WalletBalance 分账户余额信息
type WalletBalance struct {
	// VirtualSellerID 子账号ID/代理商vsellerId
	VirtualSellerID string `json:"virtual_seller_id,omitempty"`
	// TotalAvailableBalance 总可用余额，单位元，保留两位小数
	TotalAvailableBalance string `json:"total_available_balance,omitempty"`
	// TotalFrozenBalance 总冻结余额，单位元，保留两位小数
	TotalFrozenBalance string `json:"total_frozen_balance,omitempty"`
	// TotalBalance 总余额（冻结+可用）
	TotalBalance string `json:"total_balance,omitempty"`
	// BalanceList 分资金类型的余额明细
	BalanceList []WalletTypeBalance `json:"balance_list,omitempty"`
	// AccountStatus 钱包账户状态，正常：OPEN=1，冻结：FROZEN=2
	AccountStatus WalletAccountStatus `json:"account_status,omitempty"`
}

// WalletTypeBalance 分资金类型的余额明细
type WalletTypeBalance struct {
	// VirtualSellerID 子账号ID/代理商vsellerId
	VirtualSellerID string `json:"virtual_seller_id,omitempty"`
	// WalletType 资金类型 0现金 1授信 2返货 3券 4赔付返货
	WalletType WalletType `json:"wallet_type,omitempty"`
	// AvailableBalance 可用余额，单位元，保留两位小数
	AvailableBalance string `json:"available_balance,omitempty"`
	// FrozenBalance 冻结余额，单位元，保留两位小数
	FrozenBalance string `json:"frozen_balance,omitempty"`
	// TotalBalance 总余额（冻结+可用）
	TotalBalance string `json:"total_balance,omitempty"`
}
