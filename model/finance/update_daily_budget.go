package finance

import (
	"github.com/jundaychan/spotlight-mapi/util"
)

// SmartSwitch 是否开启节假日上浮
type SmartSwitch int

const (
	// SmartSwitchOff 关闭预算上浮
	SmartSwitchOff SmartSwitch = 0
	// SmartSwitchOn 开启预算上浮
	SmartSwitchOn SmartSwitch = 1
)

// UpdateDailyBudgetRequest 修改账户日预算 API Request
type UpdateDailyBudgetRequest struct {
	// AdvertiserID 品牌方ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// LimitDayBudget 是否限制预算，不限预算：0 指定预算：1
	LimitDayBudget LimitDayBudget `json:"limit_day_budget,omitempty"`
	// AccountBudget 账户日预算，单位为分，指定预算时必填
	AccountBudget int `json:"account_budget,omitempty"`
	// SmartSwitch 是否开启节假日上浮，关闭预算上浮：0 开启预算上浮：1 指定预算时必填
	SmartSwitch SmartSwitch `json:"smart_switch,omitempty"`
}

// Encode implement PostRequest interface
func (r UpdateDailyBudgetRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
