package reportrealtime2

import (
	"github.com/jundaychan/spotlight-mapi/enum"
	"github.com/jundaychan/spotlight-mapi/model"
	"github.com/jundaychan/spotlight-mapi/model/report"
	"github.com/jundaychan/spotlight-mapi/util"
)

// TargetRequest 定向层级实时数据 API Request
type TargetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id"`
	// PageNum 页数，默认1
	PageNum int64 `json:"page_num,omitempty"`
	// PageSize 页大小，默认20,最大100
	PageSize int64 `json:"page_size,omitempty"`
	// StartDate 开始时间，格式 yyyy-MM-dd
	StartDate string `json:"start_date,omitempty"`
	// EndDate 结束时间，格式 yyyy-MM-dd
	EndDate string `json:"end_date,omitempty"`
	// SortColumn 排序字段见附录column字段
	SortColumn string `json:"sort_column,omitempty"`
	// Sort 升降序asc：升序desc：降序
	Sort enum.SortType `json:"sort,omitempty"`
	// Name 搜索定向名称
	Name string `json:"name,omitempty"`
	// MarketingTargetList 营销目标过滤条件
	MarketingTargetList []int `json:"marketing_target_list,omitempty"`
	// NeedHourlyData 是否拉取小时数据，默认为 false; 只支持拉取今日数据
	NeedHourlyData bool `json:"need_hourly_data,omitempty"`
}

// Encode implement PostRequest interface
func (r TargetRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// TargetResponse 定向层级实时数据 API Response
type TargetResponse struct {
	model.BaseResponse
	// Page 分页信息
	Page *model.PageRespDTO `json:"page,omitempty"`
	// TargetDTOs 定向数据
	TargetDTOs []TargetDTO `json:"target_dtos,omitempty"`
	// TotalData 汇总数据
	TotalData *report.DataReportDTO `json:"total_data,omitempty"`
}

// TargetDTO 定向数据
type TargetDTO struct {
	// Data 数据指标
	Data *report.DataReportDTO `json:"data,omitempty"`
	// BaseCampaignDTO 计划属性信息
	BaseCampaignDTO *BaseCampaignDTO `json:"base_campaign_dto,omitempty"`
	// BaseUnitDTO 单元属性信息
	BaseUnitDTO *BaseUnitDTO `json:"base_unit_dto,omitempty"`
	// BaseTargetDTO 定向属性信息
	BaseTargetDTO *BaseTargetDTO `json:"base_target_dto,omitempty"`
	// HourlyData 小时数据
	HourlyData []report.DataReportDTO `json:"hourly_data,omitempty"`
}

// BaseCampaignDTO 计划属性信息
type BaseCampaignDTO struct {
	// CampaignID 计划id
	CampaignID uint64 `json:"campaign_id,omitempty"`
	// CampaignName 计划名称
	CampaignName string `json:"campaign_name,omitempty"`
	// CampaignFilterState 计划状态
	CampaignFilterState int `json:"campaign_filter_state,omitempty"`
	// CampaignCreateTime 计划创建时间: 格式yyyy-MM-dd HH:mm:ss
	CampaignCreateTime string `json:"campaign_create_time,omitempty"`
	// CampaignEnable 计划启停状态：0：暂停，1：开启
	CampaignEnable int `json:"campaign_enable,omitempty"`
	// MarketingTarget 营销诉求
	MarketingTarget int `json:"marketing_target,omitempty"`
	// Placement 广告类型:1：信息流，2：搜索，4：全站智投，7：视频内流
	Placement int `json:"placement,omitempty"`
	// OptimizeTarget 推广目标
	OptimizeTarget int `json:"optimize_target,omitempty"`
	// PromotionTarget 投放标的:1：笔记，2：商品，7：外链落地页，9：落地页，18：直播间
	PromotionTarget int `json:"promotion_target,omitempty"`
	// BiddingStrategy 出价方式：2：手动出价3：自动出价
	BiddingStrategy int `json:"bidding_strategy,omitempty"`
	// ConstraintType 成本控制方式
	ConstraintType int `json:"constraint_type,omitempty"`
	// ConstraintValue 成本控制值
	ConstraintValue int `json:"constraint_value,omitempty"`
	// LimitDayBudget 预算类型：0：不限预算，1：指定预算
	LimitDayBudget int `json:"limit_day_budget,omitempty"`
	// OriginCampaignDayBudget 计划日预算
	OriginCampaignDayBudget int `json:"origin_campaign_day_budget,omitempty"`
	// BudgetState 预算状态，0: 计划预算不足，1 计划预算充足
	BudgetState int `json:"budget_state,omitempty"`
	// SmartSwitch 是否节假日预算上调，0: 关闭，1: 开启
	SmartSwitch int `json:"smart_switch,omitempty"`
	// PacingMode 投放速率，1: 匀速投放，2: 加速投放
	PacingMode int `json:"pacing_mode,omitempty"`
	// StartTime 计划开始时间：格式yyyy-MM-dd
	StartTime string `json:"start_time,omitempty"`
	// ExpireTime 计划结束时间：格式yyyy-MM-dd
	ExpireTime string `json:"expire_time,omitempty"`
	// TimePeriod 时段: 默认168个1
	TimePeriod string `json:"time_period,omitempty"`
	// TimePeriodType 推广时段类型, 0: 全时段，1:自定义时间段
	TimePeriodType int `json:"time_period_type,omitempty"`
	// BuildType 搭建方式，0：标准搭建，1：省心智投
	BuildType int `json:"build_type,omitempty"`
	// FeedFlag 是否搜索追投信息流：0: 否，1：是
	FeedFlag int `json:"feed_flag,omitempty"`
	// SearchFlag 是否信息流快投搜索：0: 否，1：是
	SearchFlag int `json:"search_flag,omitempty"`
	// MigrationStatus 专业号平台计划迁移状态: 0：非迁移计划，2：迁移计划
	MigrationStatus int `json:"migration_status,omitempty"`
}

// BaseUnitDTO 单元属性信息
type BaseUnitDTO struct {
	// UnitID 单元id
	UnitID uint64 `json:"unit_id,omitempty"`
	// UnitName 单元名称
	UnitName string `json:"unit_name,omitempty"`
	// UnitFilterState 单元状态
	UnitFilterState int `json:"unit_filter_state,omitempty"`
	// UnitCreateTime 单元创建时间：格式 yyyy-MM-dd HH:mm:ss
	UnitCreateTime string `json:"unit_create_time,omitempty"`
	// UnitEnable 单元启停状态：0：暂停，1：开启
	UnitEnable int `json:"unit_enable,omitempty"`
	// CampaignID 计划id
	CampaignID uint64 `json:"campaign_id,omitempty"`
	// EventBid 出价
	EventBid int `json:"event_bid,omitempty"`
}

// BaseTargetDTO 定向属性信息
type BaseTargetDTO struct {
	// TargetName 定向名称
	TargetName string `json:"target_name,omitempty"`
	// TargetStatus 定向状态
	TargetStatus int `json:"target_status,omitempty"`
	// UnitID 单元id
	UnitID uint64 `json:"unit_id,omitempty"`
	// CampaignID 计划id
	CampaignID uint64 `json:"campaign_id,omitempty"`
	// TargetID 定向id
	TargetID uint64 `json:"target_id,omitempty"`
}
