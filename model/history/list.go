package history

import (
	"github.com/jundaychan/spotlight-mapi/model"
	"github.com/jundaychan/spotlight-mapi/util"
)

// ListRequest 历史操作记录 API Request
type ListRequest struct {
	// AdvertiserID 品牌方ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// StartTime 开始时间，格式需为yyyy-mm-dd，最早为90天前
	StartTime string `json:"start_time,omitempty"`
	// EndTime 结束时间，格式需为yyyy-mm-dd
	EndTime string `json:"end_time,omitempty"`
	// Page 第几页，默认为1
	Page int `json:"page,omitempty"`
	// PageSize 页大小，默认为20，最大50
	PageSize int `json:"page_size,omitempty"`
	// IDs 操作对象ID，仅当操作层级为计划、单元、创意、关键词时可传值 最多50个
	IDs []uint64 `json:"ids,omitempty"`
	// Name 操作对象名称
	Name string `json:"name,omitempty"`
	// OptType 操作内容，-1:全部，详见接口文档枚举说明，默认为-1
	OptType int `json:"opt_type,omitempty"`
	// Level 操作层级，-1:全部 0:账户 1:计划 2:单元 3:创意 4:关键词 5:定向 6:工具 7:报表 8:广告组 默认为-1
	Level int `json:"level,omitempty"`
}

// Encode implement PostRequest interface
func (r ListRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// ListResponse 历史操作记录 API Response
type ListResponse struct {
	model.BaseResponse
	Data *ListResult `json:"data,omitempty"`
}

// ListResult 历史操作记录结果
type ListResult struct {
	// Total 总记录数
	Total int64 `json:"total,omitempty"`
	// TotalPage 总页数
	TotalPage int64 `json:"total_page,omitempty"`
	// HistoryList 操作记录
	HistoryList []OperateRecord `json:"history_list,omitempty"`
}

// OperateRecord 操作记录
type OperateRecord struct {
	// ID 操作记录ID
	ID uint64 `json:"id,omitempty"`
	// OptTime 操作时间，示例：2025-01-01 00:00:00
	OptTime string `json:"opt_time,omitempty"`
	// OptAccountName 操作人
	OptAccountName string `json:"opt_account_name,omitempty"`
	// OptIP ip地址
	OptIP string `json:"opt_ip,omitempty"`
	// OptLevelName 操作层级
	OptLevelName string `json:"opt_level_name,omitempty"`
	// OptObject 操作对象
	OptObject string `json:"opt_object,omitempty"`
	// OptObjectID 操作对象ID
	OptObjectID model.Uint64 `json:"opt_object_id,omitempty"`
	// OptTypeName 操作内容
	OptTypeName string `json:"opt_type_name,omitempty"`
	// OldValue 操作前
	OldValue string `json:"old_value,omitempty"`
	// NewValue 操作后
	NewValue string `json:"new_value,omitempty"`
}
