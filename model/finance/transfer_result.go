package finance

import (
	"github.com/bububa/spotlight-mapi/model"
	"github.com/bububa/spotlight-mapi/util"
)

// TransferStatus 转账处理状态
type TransferStatus string

const (
	// TransferStatusProcessing 处理中
	TransferStatusProcessing TransferStatus = "PROCESSING"
	// TransferStatusSuccess 成功
	TransferStatusSuccess TransferStatus = "SUCCESS"
	// TransferStatusFail 失败
	TransferStatusFail TransferStatus = "FAIL"
)

// TransferResultRequest 查询转账结果 API Request
type TransferResultRequest struct {
	// UserID 代理商主账号ID
	UserID string `json:"user_id,omitempty"`
	// VirtualSellerID 转出账户的子账号ID/代理商vsellerId
	VirtualSellerID string `json:"virtual_seller_id,omitempty"`
	// RequestNo 调用转账接口传入的request_no
	RequestNo string `json:"request_no,omitempty"`
}

// Encode implement PostRequest interface
func (r TransferResultRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// TransferResultResponse 查询转账结果 API Response
type TransferResultResponse struct {
	model.BaseResponse
	Data *TransferResultData `json:"data,omitempty"`
}

// TransferResultData 转账结果查询数据
type TransferResultData struct {
	// VirtualSellerID 转出账户的子账号ID/代理商vsellerId
	VirtualSellerID string `json:"virtual_seller_id,omitempty"`
	// RequestNo 调用转账接口传入的request_no
	RequestNo string `json:"request_no,omitempty"`
	// Status 处理中：PROCESSING 成功：SUCCESS 失败：FAIL
	Status TransferStatus `json:"status,omitempty"`
}
