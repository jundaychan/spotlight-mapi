package simpledelivery

import (
	"github.com/jundaychan/spotlight-mapi/model"
	"github.com/jundaychan/spotlight-mapi/model/report"
	"github.com/jundaychan/spotlight-mapi/util"
)

// CampaignRealtimeRequest 简单投计划层级实时数据 API Request
type CampaignRealtimeRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// CampaignGroupIDList 广告组ID，必传且只能传一个
	CampaignGroupIDList []uint64 `json:"campaign_group_id_list,omitempty"`
	// ID 计划ID，根据计划ID查询时其它筛选条件无效
	ID uint64 `json:"id,omitempty"`
	// Name 计划名称，根据计划名查询时其它筛选条件无效
	Name string `json:"name,omitempty"`
	// PageNum 页数，默认1
	PageNum int64 `json:"page_num,omitempty"`
	// PageSize 页大小，默认20，最大100
	PageSize int64 `json:"page_size,omitempty"`
	// StartDate 开始时间，格式 yyyy-MM-dd
	StartDate string `json:"start_date,omitempty"`
	// EndDate 结束时间，格式 yyyy-MM-dd
	EndDate string `json:"end_date,omitempty"`
	// DataCaliber 数据指标归因时间类型，0-计费时间、1-转化时间
	DataCaliber int `json:"data_caliber,omitempty"`
	// CampaignFilterStateList 计划状态，见枚举说明
	CampaignFilterStateList []int `json:"campaign_filter_state_list,omitempty"`
	// CombineAuditStatus 审核状态，1-部分审核拒绝、2-部分审核中、3-审核通过、4-审核通过（私密）、99-不满足审核条件
	CombineAuditStatus int `json:"combine_audit_status,omitempty"`
}

// Encode implement PostRequest interface
func (r CampaignRealtimeRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// CampaignRealtimeResponse 简单投计划层级实时数据 API Response
type CampaignRealtimeResponse struct {
	model.BaseResponse
	// Data 分页数据
	Data *CampaignRealtimeData `json:"data,omitempty"`
}

// CampaignRealtimeData 简单投计划层级实时数据
type CampaignRealtimeData struct {
	// Page 分页信息
	Page *PageDTO `json:"page,omitempty"`
	// TotalData 总计数据
	TotalData *report.DataReportDTO `json:"total_data,omitempty"`
	// List 当前页数据
	List []CampaignDTO `json:"campaign_dtos,omitempty"`
}

// CampaignDTO 计划数据
type CampaignDTO struct {
	// BaseCampaignDTO 计划实体对象
	BaseCampaignDTO *BaseCampaignDTO `json:"base_campaign_dto,omitempty"`
	// Data 数据报表
	Data *report.DataReportDTO `json:"data,omitempty"`
}

// BaseCampaignDTO 计划实体对象
type BaseCampaignDTO struct {
	// CampaignID 计划ID
	CampaignID uint64 `json:"campaign_id,omitempty"`
	// CampaignName 计划名称
	CampaignName string `json:"campaign_name,omitempty"`
	// CampaignGroupID 广告组ID
	CampaignGroupID uint64 `json:"campaign_group_id,omitempty"`
	// CampaignGroupName 标的名称
	CampaignGroupName string `json:"campaign_group_name,omitempty"`
	// CampaignEnable 计划上线状态，0-下线、1-上线
	CampaignEnable int `json:"campaign_enable,omitempty"`
	// CampaignFilterState 计划状态，见枚举说明
	CampaignFilterState int `json:"campaign_filter_state,omitempty"`
	// CampaignCreateTime 创建时间，格式 yyyy-MM-dd HH:mm:ss
	CampaignCreateTime string `json:"campaign_create_time,omitempty"`
	// LimitDayBudget 是否限制日预算，0-不限、1-限制
	LimitDayBudget int `json:"limit_day_budget,omitempty"`
	// OriginCampaignDayBudget 计划日预算，单位分
	OriginCampaignDayBudget int64 `json:"origin_campaign_day_budget,omitempty"`
	// SmartSwitch 是否开启节假日预算上浮，0-不开启、1-开启
	SmartSwitch int `json:"smart_switch,omitempty"`
	// ConstraintValue 出价，单位分
	ConstraintValue int64 `json:"constraint_value,omitempty"`
}
