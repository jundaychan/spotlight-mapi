package finance

import (
	"github.com/jundaychan/spotlight-mapi/model"
	"github.com/jundaychan/spotlight-mapi/util"
)

// AccountFundType 资金账户类型
type AccountFundType int

const (
	// AccountFundTypeCash 现金
	AccountFundTypeCash AccountFundType = 0
	// AccountFundTypeReturn 常规返货
	AccountFundTypeReturn AccountFundType = 1
	// AccountFundTypeCredit 授信
	AccountFundTypeCredit AccountFundType = 2
	// AccountFundTypeCompensateReturn 赔付返货
	AccountFundTypeCompensateReturn AccountFundType = 6
)

// AccountOperateType 账户流水消费类型
type AccountOperateType int

const (
	// AccountOperateTypeCashRecharge 现金充值
	AccountOperateTypeCashRecharge AccountOperateType = 1
	// AccountOperateTypeReturnCash 返现
	AccountOperateTypeReturnCash AccountOperateType = 2
	// AccountOperateTypeConsume 消费扣款
	AccountOperateTypeConsume AccountOperateType = 3
	// AccountOperateTypeReduction 资金补扣
	AccountOperateTypeReduction AccountOperateType = 4
	// AccountOperateTypeVirtualRecharge 虚拟充值
	AccountOperateTypeVirtualRecharge AccountOperateType = 5
	// AccountOperateTypeTransfer 转账
	AccountOperateTypeTransfer AccountOperateType = 6
	// AccountOperateTypeRefund 退款
	AccountOperateTypeRefund AccountOperateType = 7
	// AccountOperateTypeWithdraw 提现
	AccountOperateTypeWithdraw AccountOperateType = 8
	// AccountOperateTypeRecycle 回收
	AccountOperateTypeRecycle AccountOperateType = 9
	// AccountOperateTypeReturnRecover 返货追回
	AccountOperateTypeReturnRecover AccountOperateType = 10
	// AccountOperateTypePaymentSuccess 打款成功
	AccountOperateTypePaymentSuccess AccountOperateType = 11
	// AccountOperateTypePaymentFail 打款失败
	AccountOperateTypePaymentFail AccountOperateType = 12
	// AccountOperateTypeEntry 入账
	AccountOperateTypeEntry AccountOperateType = 13
	// AccountOperateTypeCharge 出款
	AccountOperateTypeCharge AccountOperateType = 14
	// AccountOperateTypeEcommerceRecharge 电商货款充值
	AccountOperateTypeEcommerceRecharge AccountOperateType = 15
	// AccountOperateTypeOrderFreeze 下单冻结
	AccountOperateTypeOrderFreeze AccountOperateType = 16
	// AccountOperateTypeOrderCancel 取消订单
	AccountOperateTypeOrderCancel AccountOperateType = 17
	// AccountOperateTypeOrderRelease 订单完结释放余额
	AccountOperateTypeOrderRelease AccountOperateType = 18
)

// AccountFlowRequest 获取账户流水 API Request
type AccountFlowRequest struct {
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
	// AccountTypes 资金账户类型 0-现金 1-常规返货 2-授信 6-赔付返货 不传，默认值为[0,1,2,6]
	AccountTypes []AccountFundType `json:"account_types,omitempty"`
	// Type 数据类型，账户为服务商时必传，传type=2
	Type int `json:"type,omitempty"`
}

// Encode implement PostRequest interface
func (r AccountFlowRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// AccountFlowResponse 获取账户流水 API Response
type AccountFlowResponse struct {
	model.BaseResponse
	Data *AccountFlowResult `json:"data,omitempty"`
}

// AccountFlowResult 账户流水数据
type AccountFlowResult struct {
	// TotalCount 总记录数
	TotalCount int64 `json:"total_count,omitempty"`
	// AggregateOrder 消费全部金额总计，单位：分
	AggregateOrder int64 `json:"aggregate_order,omitempty"`
	// AccountTradeDetail 账户流水详情
	AccountTradeDetail []AccountTradeDetail `json:"account_trade_detail,omitempty"`
}

// AccountTradeDetail 账户流水详情
type AccountTradeDetail struct {
	// LaunchDate 消费时间
	LaunchDate string `json:"launch_date,omitempty"`
	// OperateType 消费类型
	OperateType AccountOperateType `json:"operate_type,omitempty"`
	// TradeTime 操作时间
	TradeTime string `json:"trade_time,omitempty"`
	// AccountName 账户名称
	AccountName string `json:"account_name,omitempty"`
	// OrderAmount 消费金额，单位：分
	OrderAmount int64 `json:"order_amount,omitempty"`
	// Balance 余额，单位：分
	Balance int64 `json:"balance,omitempty"`
	// TransferObject 款项对象
	TransferObject string `json:"transfer_object,omitempty"`
	// Remark 备注
	Remark string `json:"remark,omitempty"`
	// AccountType 资金类型
	AccountType AccountFundType `json:"account_type,omitempty"`
	// AccountTypeName 资金类型名称
	AccountTypeName string `json:"account_type_name,omitempty"`
	// BusinessTypeName 业务类型名称
	BusinessTypeName string `json:"business_type_name,omitempty"`
}
