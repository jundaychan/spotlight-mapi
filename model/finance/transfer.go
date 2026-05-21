package finance

import (
	"github.com/bububa/spotlight-mapi/model"
	"github.com/bububa/spotlight-mapi/util"
)

// WalletType 资金类型
type WalletType int

const (
	// WalletTypeCash 现金
	WalletTypeCash WalletType = 0
	// WalletTypeCredit 授信
	WalletTypeCredit WalletType = 1
	// WalletTypeReturn 返货
	WalletTypeReturn WalletType = 2
	// WalletTypeVoucher 券
	WalletTypeVoucher WalletType = 3
	// WalletTypeCompensationReturn 赔付返货
	WalletTypeCompensationReturn WalletType = 4
)

// TransferWallet 转账账户
type TransferWallet struct {
	// VirtualSellerID 代理商/子账号vsellerId
	VirtualSellerID string `json:"virtual_seller_id,omitempty"`
	// WalletType 资金类型，CASH = 0 // 现金；CREDIT = 1 // 授信；COMPENSATION_RETURN = 4 // 赔付返货；
	WalletType WalletType `json:"wallet_type,omitempty"`
}

// TransferRequest 转账 API Request
type TransferRequest struct {
	// UserID 代理商主账号ID
	UserID string `json:"user_id,omitempty"`
	// FromWallet 转出账户
	FromWallet *TransferWallet `json:"from_wallet,omitempty"`
	// ToWallet 转入账户
	ToWallet *TransferWallet `json:"to_wallet,omitempty"`
	// BusinessNo 业务单号，64字符，同request_no
	BusinessNo string `json:"business_no,omitempty"`
	// RequestNo 请求单号，64字符，必须包含业务前缀，且每次请求重新生成，确保全局唯一
	RequestNo string `json:"request_no,omitempty"`
	// Amount 转账金额，单位元，保留两位小数
	Amount string `json:"amount,omitempty"`
	// Remark 备注
	Remark string `json:"remark,omitempty"`
	// PartTransfer 应转尽转
	PartTransfer bool `json:"part_transfer,omitempty"`
}

// Encode implement PostRequest interface
func (r TransferRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// TransferResponse 转账 API Response
type TransferResponse struct {
	model.BaseResponse
	Data *TransferResult `json:"data,omitempty"`
}

// TransferResult 转账结果
type TransferResult struct {
	// TransferNo 转账流水号，内部流水号，非业务单号
	TransferNo string `json:"transfer_no,omitempty"`
	// ApplyAmount 申请转账金额
	ApplyAmount string `json:"apply_amount,omitempty"`
	// ActualAmount 实际转账金额，单位元，保留两位小数
	ActualAmount string `json:"actual_amount,omitempty"`
}
