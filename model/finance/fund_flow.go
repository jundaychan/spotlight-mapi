package finance

import (
	"github.com/bububa/spotlight-mapi/model"
	"github.com/bububa/spotlight-mapi/util"
)

// FundDateKey 拉取日期口径
type FundDateKey int

const (
	// FundDateKeyLaunch 按投放日期拉取
	FundDateKeyLaunch FundDateKey = 1
	// FundDateKeyCharge 按扣费日期拉取
	FundDateKeyCharge FundDateKey = 2
)

// TradeType 交易类型
type TradeType int

const (
	// TradeTypeRecharge 充值
	TradeTypeRecharge TradeType = 1
	// TradeTypeEntry 入账
	TradeTypeEntry TradeType = 2
	// TradeTypeConsume 消费扣款
	TradeTypeConsume TradeType = 3
	// TradeTypeRefund 退款
	TradeTypeRefund TradeType = 4
	// TradeTypeTransfer 转账
	TradeTypeTransfer TradeType = 5
	// TradeTypeCharge 出款
	TradeTypeCharge TradeType = 6
	// TradeTypeOrderPay 订单支付
	TradeTypeOrderPay TradeType = 11
	// TradeTypeOrderRefund 订单退款
	TradeTypeOrderRefund TradeType = 12
)

// FundFlowRequest 资金流水查询 API Request
type FundFlowRequest struct {
	// UserID 代理商主账号ID
	UserID string `json:"user_id,omitempty"`
	// VirtualSellerID 代理商/子账号vsellerId
	VirtualSellerID string `json:"virtual_seller_id,omitempty"`
	// StartDate 起始日期，yyyy-MM-dd
	StartDate string `json:"start_date,omitempty"`
	// EndDate 结束日期，yyyy-MM-dd
	EndDate string `json:"end_date,omitempty"`
	// PageIndex 页码
	PageIndex int `json:"page_index,omitempty"`
	// PageSize 每页条数
	PageSize int `json:"page_size,omitempty"`
	// DateKey 1: 按投放日期拉取 2: 按扣费日期拉取
	DateKey FundDateKey `json:"date_key,omitempty"`
	// WalletTypeList 资金类型，CASH = 0 // 现金；CREDIT = 1 // 授信；COMPENSATION_RETURN = 4 // 赔付返货；
	WalletTypeList []WalletType `json:"wallet_type_list,omitempty"`
	// TradeTypeList 交易类型
	TradeTypeList []TradeType `json:"trade_type_list,omitempty"`
}

// Encode implement PostRequest interface
func (r FundFlowRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// FundFlowResponse 资金流水查询 API Response
type FundFlowResponse struct {
	model.BaseResponse
	Data *FundFlowResult `json:"data,omitempty"`
}

// FundFlowResult 资金流水查询数据
type FundFlowResult struct {
	// Total 总条数
	Total int64 `json:"total,omitempty"`
	// TotalIncomeAmount 总收入金额，单位元，保留两位小数
	TotalIncomeAmount string `json:"total_income_amount,omitempty"`
	// TotalOutcomeAmount 总支出金额，单位元，保留两位小数
	TotalOutcomeAmount string `json:"total_outcome_amount,omitempty"`
	// WalletTradeAggList 分账户聚合流水信息
	WalletTradeAggList []WalletTradeAgg `json:"wallet_trade_agg_list,omitempty"`
	// RecordList 流水明细
	RecordList []FundFlowRecord `json:"record_list,omitempty"`
}

// WalletTradeAgg 分账户聚合流水信息
type WalletTradeAgg struct {
	// VirtualSellerID 代理商/子账号vsellerId
	VirtualSellerID string `json:"virtual_seller_id,omitempty"`
	// TradeAmountAggList 聚合数据
	TradeAmountAggList []TradeAmountAgg `json:"trade_amount_agg_list,omitempty"`
	// TotalIncomeAmount 总收入金额，单位元，保留两位小数
	TotalIncomeAmount string `json:"total_income_amount,omitempty"`
	// TotalOutcomeAmount 总支出金额，单位元，保留两位小数
	TotalOutcomeAmount string `json:"total_outcome_amount,omitempty"`
}

// TradeAmountAgg 资金类型聚合数据
type TradeAmountAgg struct {
	// WalletType 资金类型
	WalletType WalletType `json:"wallet_type,omitempty"`
	// IncomeTotalAmount 总收入金额，单位元，保留两位小数
	IncomeTotalAmount string `json:"income_total_amount,omitempty"`
	// OutcomeTotalAmount 总支出金额，单位元，保留两位小数
	OutcomeTotalAmount string `json:"outcome_total_amount,omitempty"`
}

// FundFlowRecord 流水明细
type FundFlowRecord struct {
	// TradeNo 交易流水号
	TradeNo string `json:"trade_no,omitempty"`
	// TradeTime 交易时间，毫秒时间戳
	TradeTime int64 `json:"trade_time,omitempty"`
	// BusinessTradeTime 业务时间，毫秒时间戳
	BusinessTradeTime int64 `json:"business_trade_time,omitempty"`
	// BusinessTradeNo 业务交易单号
	BusinessTradeNo string `json:"business_trade_no,omitempty"`
	// BusinessNo 外部业务单号
	BusinessNo string `json:"business_no,omitempty"`
	// VirtualSellerID 子账号ID/代理商vsellerId
	VirtualSellerID string `json:"virtual_seller_id,omitempty"`
	// WalletType 资金类型
	WalletType WalletType `json:"wallet_type,omitempty"`
	// ParentVirtualSellerID 子账号对应的代理商vsellerId
	ParentVirtualSellerID string `json:"parent_virtual_seller_id,omitempty"`
	// TargetVirtualSellerID 转账对手方vsellerId
	TargetVirtualSellerID string `json:"target_virtual_seller_id,omitempty"`
	// AccountName 账户名称
	AccountName string `json:"account_name,omitempty"`
	// TargetAccountName 转账对手方账户名称
	TargetAccountName string `json:"target_account_name,omitempty"`
	// Direction 动帐方向：1收入，-1支出
	Direction model.Int `json:"direction,omitempty"`
	// TradeType 交易类型枚举名称
	TradeType string `json:"trade_type,omitempty"`
	// Amount 交易金额，单位元，保留两位小数
	Amount string `json:"amount,omitempty"`
	// BalanceBefore 变更前余额，单位元，保留两位小数（不含冻结金额）
	BalanceBefore string `json:"balance_before,omitempty"`
	// BalanceAfter 变更后余额，单位元，保留两位小数（不含冻结金额）
	BalanceAfter string `json:"balance_after,omitempty"`
	// Remark 备注
	Remark string `json:"remark,omitempty"`
}
