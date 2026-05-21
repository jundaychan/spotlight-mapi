package finance

import (
	"github.com/bububa/spotlight-mapi/model"
	"github.com/bububa/spotlight-mapi/util"
)

// LimitDayBudget 是否限制预算
type LimitDayBudget int

const (
	// LimitDayBudgetUnlimited 不限预算
	LimitDayBudgetUnlimited LimitDayBudget = 0
	// LimitDayBudgetLimited 限制预算
	LimitDayBudgetLimited LimitDayBudget = 1
)

// DailyBudgetBalanceRequest 获取账户日预算余额 API Request
type DailyBudgetBalanceRequest struct {
	// AdvertiserID 品牌方ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
}

// Encode implement PostRequest interface
func (r DailyBudgetBalanceRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// DailyBudgetBalanceResponse 获取账户日预算余额 API Response
type DailyBudgetBalanceResponse struct {
	model.BaseResponse
	Data *DailyBudgetBalance `json:"data,omitempty"`
}

// DailyBudgetBalance 账户日预算余额数据
type DailyBudgetBalance struct {
	// TotalBalance 账户余额
	TotalBalance int64 `json:"total_balance,omitempty"`
	// CashBalance 现金余额
	CashBalance int64 `json:"cash_balance,omitempty"`
	// ReturnBalance 常规返货余额
	ReturnBalance int64 `json:"return_balance,omitempty"`
	// CreditBalance 授信金额
	CreditBalance int64 `json:"credit_balance,omitempty"`
	// FreezeBalance 冻结余额
	FreezeBalance int64 `json:"freeze_balance,omitempty"`
	// AvailableBalance 可用金额
	AvailableBalance int64 `json:"available_balance,omitempty"`
	// TodaySpend 今日花费
	TodaySpend int64 `json:"today_spend,omitempty"`
	// CompensateReturnBalance 赔付返货余额
	CompensateReturnBalance int64 `json:"compensate_return_balance,omitempty"`
	// AccountBudget 账户日预算
	AccountBudget int `json:"account_budget,omitempty"`
	// LimitDayBudget 是否限制预算 0-不限预算，1-限制预算
	LimitDayBudget LimitDayBudget `json:"limit_day_budget,omitempty"`
}
