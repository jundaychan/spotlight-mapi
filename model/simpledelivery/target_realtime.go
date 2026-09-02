package simpledelivery

import (
	"github.com/jundaychan/spotlight-mapi/model"
	"github.com/jundaychan/spotlight-mapi/model/report"
	"github.com/jundaychan/spotlight-mapi/util"
)

// TargetRealtimeRequest 简单投标的层级实时数据 API Request
type TargetRealtimeRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// PageNum 页数，默认1
	PageNum int64 `json:"page_num,omitempty"`
	// PageSize 页大小，默认20，最大200
	PageSize int64 `json:"page_size,omitempty"`
	// StartDate 开始时间，格式 yyyy-MM-dd
	StartDate string `json:"start_date,omitempty"`
	// EndDate 结束时间，格式 yyyy-MM-dd
	EndDate string `json:"end_date,omitempty"`
	// DataCaliber 数据指标归因时间类型，0-点击时间、1-转化时间
	DataCaliber int `json:"data_caliber,omitempty"`
	// Columns 数据指标展示列，传驼峰式的字段名
	Columns []string `json:"columns,omitempty"`
	// Filter 筛选项
	Filter *QueryFilterDTO `json:"filter,omitempty"`
	// MarketingTargetList 营销诉求，4-产品种草、9-客资收集、13-种草直达、16-应用唤起、20-应用下载、21-小程序推广
	MarketingTargetList []int `json:"marketing_target_list,omitempty"`
	// GroupTypeList 标的类型，1-全自动托管、4-半自动托管
	GroupTypeList []int `json:"group_type_list,omitempty"`
}

// Encode implement PostRequest interface
func (r TargetRealtimeRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// QueryFilterDTO 标的筛选项
type QueryFilterDTO struct {
	// Name 标的名称
	Name string `json:"name,omitempty"`
	// GroupFilterState 标的状态，0-所有未删除状态、1-有效、2-暂停、3-广告组预算不足、4-账户预算不足、5-现金余额不足、6-已删除
	GroupFilterState int `json:"group_filter_state,omitempty"`
}

// TargetRealtimeResponse 简单投标的层级实时数据 API Response
type TargetRealtimeResponse struct {
	model.BaseResponse
	// Data 分页数据
	Data *TargetRealtimeData `json:"data,omitempty"`
}

// TargetRealtimeData 简单投标的层级实时数据
type TargetRealtimeData struct {
	// Page 分页信息
	Page *PageDTO `json:"page,omitempty"`
	// TotalData 总计数据
	TotalData *report.DataReportDTO `json:"total_data,omitempty"`
	// List 当前页数据
	List []GroupDTO `json:"data_list,omitempty"`
}

// GroupDTO 标的数据
type GroupDTO struct {
	// GroupParadigm 广告组实体对象
	GroupParadigm *GroupParadigm `json:"group_paradigm,omitempty"`
	// Data 数据报表
	Data *report.DataReportDTO `json:"data,omitempty"`
}

// GroupParadigm 广告组实体对象
type GroupParadigm struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// CampaignGroupID 广告组ID（标的ID）。ube/campaign、ube/note、ube/keyword 三个实时口必传且只能传一个，就是这里拿
	CampaignGroupID uint64 `json:"campaign_group_id,omitempty"`
	// GroupType 标的类型，1-全自动托管、4-半自动托管
	GroupType int `json:"group_type,omitempty"`
	// MarketingTarget 营销诉求，4-产品种草、9-客资收集、13-种草直达、16-应用唤起、20-应用下载、21-小程序推广
	MarketingTarget int `json:"marketing_target,omitempty"`
	// CampaignGroupName 标的名称
	CampaignGroupName string `json:"campaign_group_name,omitempty"`
	// Enable 上线状态，0-下线、1-上线
	Enable int `json:"enable,omitempty"`
	// GroupFilterState 标的状态，0-所有未删除状态、1-有效、2-暂停、3-广告组预算不足、4-账户预算不足、5-现金余额不足、6-已删除
	GroupFilterState int `json:"group_filter_state,omitempty"`
	// CreateTime 创建时间，格式 yyyy-MM-dd
	CreateTime string `json:"create_time,omitempty"`
	// ValidCampaignCount 关联计划数
	ValidCampaignCount int `json:"valid_campaign_count,omitempty"`
}
