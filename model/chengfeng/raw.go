package chengfeng

import "github.com/jundaychan/spotlight-mapi/model"

// RawReportResponse 通用离线报表(原样保留返回的全部字段，不丢任何列)。
// 用于"笔记列表"等需要拿到 note_title/note_image/note_jump_url 等 SDK 强类型未覆盖字段的场景。
type RawReportResponse struct {
	model.BaseResponse
	Data *RawReportList `json:"data,omitempty"`
}

// RawReportList 原始报表数据(每行是完整 JSON map)。
type RawReportList struct {
	DataList   []map[string]any `json:"data_list,omitempty"`
	TotalData  map[string]any   `json:"total_data,omitempty"`
	TotalCount int64            `json:"total_count,omitempty"`
}
