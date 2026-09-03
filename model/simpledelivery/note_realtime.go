package simpledelivery

import (
	"github.com/jundaychan/spotlight-mapi/model"
	"github.com/jundaychan/spotlight-mapi/model/report"
	"github.com/jundaychan/spotlight-mapi/util"
)

// NoteRealtimeRequest 简单投笔记层级实时数据 API Request
type NoteRealtimeRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// CampaignGroupIDList 广告组ID，必传且只能传一个
	CampaignGroupIDList []uint64 `json:"campaign_group_id_list,omitempty"`
	// NoteID 笔记ID，根据笔记ID查询时其它筛选条件无效
	NoteID string `json:"note_id,omitempty"`
	// PageNum 页数，默认1
	PageNum int64 `json:"page_num,omitempty"`
	// PageSize 页大小，默认20，最大100
	PageSize int64 `json:"page_size,omitempty"`
	// StartDate 开始时间，格式 yyyy-MM-dd
	StartDate string `json:"start_date,omitempty"`
	// EndDate 结束时间，格式 yyyy-MM-dd
	EndDate string `json:"end_date,omitempty"`
	// DataCaliber 数据指标归因时间类型，0-计费时间、1-转化时间
	DataCaliber *int `json:"data_caliber,omitempty"`
	// CreativityFilterState 笔记状态，见枚举说明
	CreativityFilterState int `json:"creativity_filter_state,omitempty"`
	// CreativityAuditState 笔记审核状态，1-审核拒绝、2-审核中、3-审核通过、4-审核通过（私密）、99-不满足审核条件
	CreativityAuditState int `json:"creativity_audit_state,omitempty"`
}

// Encode implement PostRequest interface
func (r NoteRealtimeRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// NoteRealtimeResponse 简单投笔记层级实时数据 API Response
type NoteRealtimeResponse struct {
	model.BaseResponse
	// Data 分页数据
	Data *NoteRealtimeData `json:"data,omitempty"`
}

// NoteRealtimeData 简单投笔记层级实时数据
type NoteRealtimeData struct {
	// Page 分页信息
	Page *PageDTO `json:"page,omitempty"`
	// TotalData 总计数据
	TotalData *report.DataReportDTO `json:"total_data,omitempty"`
	// List 当前页数据
	List []NoteDTO `json:"note_dtos,omitempty"`
}

// NoteDTO 笔记数据
type NoteDTO struct {
	// BaseNoteDTO 笔记实体对象
	BaseNoteDTO *BaseNoteDTO `json:"base_note_dto,omitempty"`
	// Data 数据报表
	Data *report.DataReportDTO `json:"data,omitempty"`
}

// BaseNoteDTO 笔记实体对象
type BaseNoteDTO struct {
	// NoteID 笔记ID
	NoteID string `json:"note_id,omitempty"`
	// CreativityID 创意ID
	CreativityID uint64 `json:"creativity_id,omitempty"`
	// CreativityName 创意名称
	CreativityName string `json:"creativity_name,omitempty"`
	// CreativityEnable 创意上线状态，0-下线、1-上线
	CreativityEnable int `json:"creativity_enable,omitempty"`
	// CreativityFilterState 创意状态，见枚举说明
	CreativityFilterState int `json:"creativity_filter_state,omitempty"`
	// CreativityCreateTime 创建时间，格式 yyyy-MM-dd HH:mm:ss
	CreativityCreateTime string `json:"creativity_create_time,omitempty"`
	// CreativityAuditState 审核状态，1-审核拒绝、2-审核中、3-审核通过、4-审核通过（私密）、99-不满足审核条件
	CreativityAuditState int `json:"creativity_audit_state,omitempty"`
	// AuditComment 拒审理由
	AuditComment map[int]string `json:"audit_comment,omitempty"`
}
