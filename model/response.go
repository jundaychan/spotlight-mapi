package model

import (
	"strconv"

	"github.com/jundaychan/spotlight-mapi/util"
)

// Response api response interface
type Response interface {
	// IsError 是否返回错误
	IsError() bool
	// Error implement error interface
	Error() string
}

// BaseResponse shared api response data fields
type BaseResponse struct {
	// Success 是否成功
	Success bool `json:"success,omitempty"`
	// Code 返回码
	Code int `json:"code,omitempty"`
	// Message 返回信息（部分老接口用 message 键）
	Message string `json:"message,omitempty"`
	// Msg 返回信息——聚光报表/创编等大多数接口实际用的是 **msg** 键。
	// 以前只映射了 message，于是所有这些接口的错误说明全被丢掉，Error() 只剩「1007003:」
	// 「-1:」这种孤零零的码（美商实测：-1 的 msg 其实是「request for method:
	// noteOfflineReportApi timeout, rpc server: …」，一句话就能定性，却被吞了几个月）。
	Msg string `json:"msg,omitempty"`
	// ErrorCode 返回码
	ErrorCode int `json:"errorCode"`
	// ErrorMsg 返回信息
	ErrorMsg string `json:"errorMsg"`
	// RequestID 请求的日志id，唯一标识一个请求
	RequestID string `json:"request_id,omitempty"`
}

// IsError implement Response interface
func (r BaseResponse) IsError() bool {
	return !r.Success
}

// Error implement Response interface
func (r BaseResponse) Error() string {
	if r.ErrorCode != 0 {
		return util.StringsJoin(strconv.Itoa(r.ErrorCode), ":", r.ErrorMsg)
	}
	msg := r.Message
	if msg == "" {
		msg = r.Msg
	}
	return util.StringsJoin(strconv.Itoa(r.Code), ":", msg)
}
