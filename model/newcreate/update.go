package newcreate

import (
	"github.com/jundaychan/spotlight-mapi/model"
	"github.com/jundaychan/spotlight-mapi/util"
)

// UpdateRequest 新创编编辑 API Request（目前只支持产品种草、客资收集营销诉求；文档 4752）
type UpdateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// ModifyType 编辑类型，1-编辑计划、2-编辑单元、3-编辑创意
	ModifyType int `json:"modify_type,omitempty"`
	// ModCascadeInfoList 新创编数据
	ModCascadeInfoList []ModCascadeInfo `json:"mod_cascade_info_list,omitempty"`
}

// Encode implement PostRequest interface
func (r UpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// ModCascadeInfo 新创编编辑级联数据
type ModCascadeInfo struct {
	// Campaign 关联的广告计划信息
	Campaign *ModifyCampaign `json:"campaign,omitempty"`
	// UnitWithCreativeList 包含广告单元及创意信息的列表
	UnitWithCreativeList []ModifyUnitWithCreative `json:"unit_with_creative_list,omitempty"`
}

// ModifyCampaign 编辑计划信息
type ModifyCampaign struct {
	// CampaignID 计划ID，当modify_type为1时必传
	CampaignID uint64 `json:"campaign_id,omitempty"`
	// CampaignName 计划名称，最多50个字符
	CampaignName string `json:"campaign_name,omitempty"`
	// TimeType 投放时间类型，0-长期投放，1-设置起止日期
	TimeType *int `json:"time_type,omitempty"`
	// StartTime 投放开始日期，示例 2023-09-20
	StartTime string `json:"start_time,omitempty"`
	// ExpireTime 投放结束日期，示例 2023-09-21
	ExpireTime string `json:"expire_time,omitempty"`
	// TimePeriodType 投放时段类型，0-不限，1-指定时段
	TimePeriodType *int `json:"time_period_type,omitempty"`
	// TimePeriod 投放时段配置
	TimePeriod *TimePeriodDTO `json:"time_period,omitempty"`
	// LimitDayBudget 预算类型，0-不限、1-指定预算
	LimitDayBudget *int `json:"limit_day_budget,omitempty"`
	// OriginCampaignDayBudget 预算金额，单位分
	OriginCampaignDayBudget int64 `json:"origin_campaign_day_budget,omitempty"`
	// SmartSwitch 节假日预算上浮，0-关闭、1-开启
	SmartSwitch *int `json:"smart_switch,omitempty"`
	// ExploreState 一键起量开关，0-关闭、1-开启
	ExploreState *int `json:"explore_state,omitempty"`
	// ExploreConfig 一键起量配置，一键起量关闭时不用传
	ExploreConfig *ExploreConfig `json:"explore_config,omitempty"`
	// SearchFlag 搜索快投开关，0-关闭、1-开启
	SearchFlag *int `json:"search_flag,omitempty"`
	// UpdateFields Field Mask：本次真正要更新的字段名集合（下划线命名）。2026-08-17 起建议必传：
	// 服务端以它为唯一依据判定更新哪些字段，能表达「改回默认值 / 清空」；不传回退到启发式合并。
	// ⚠️ 字符串/切片字段仍是裸类型 + omitempty，空值不进报文；靠 mask 声明为「要更新」而报文里
	// 又没有的字段，服务端按缺省处理——这正是 mask 的语义，但调用方要清楚自己发的是"清空"。
	// 零值有语义的枚举字段（time_type / limit_day_budget / conversion_type…）已改成指针，
	// 传 &0 就能真正发出 0，不再被 omitempty 吃掉。
	UpdateFields []string `json:"update_fields,omitempty"`
}

// ModifyUnitWithCreative 编辑单元及创意信息
type ModifyUnitWithCreative struct {
	// Unit 单元信息
	Unit *ModifyUnit `json:"unit,omitempty"`
	// CreativityList 创意信息
	CreativityList []ModifyCreativity `json:"creativity_list,omitempty"`
}

// ModifyUnit 编辑单元信息
type ModifyUnit struct {
	// UnitID 单元ID，当modify_type为2时必传
	UnitID uint64 `json:"unit_id,omitempty"`
	// UnitName 单元名称，最多50个字符
	UnitName string `json:"unit_name,omitempty"`
	// KeywordGenType 智能拓词类型，0-手动选词、1-智能拓词、2-手动+智能关键词定向
	KeywordGenType *int `json:"keyword_gen_type,omitempty"`
	// CtrConstraint 智能拓词-点击率约束
	CtrConstraint float64 `json:"ctr_constraint,omitempty"`
	// AcpConstraint 智能拓词-点击成本约束
	AcpConstraint int64 `json:"acp_constraint,omitempty"`
	// TargetPosition 目标位次，0-不限位置、1-首位、3-前三位
	TargetPosition *int `json:"target_position,omitempty"`
	// TargetType 定向类型，0-默认值（搜索位置传 0）、2-智能定向、3-高级定向
	TargetType *int `json:"target_type,omitempty"`
	// TargetInfo 定向信息
	TargetInfo *CreateTargetInfo `json:"target_info,omitempty"`
	// SearchBidRatio 搜索出价系数，默认1.0
	SearchBidRatio float64 `json:"search_bid_ratio,omitempty"`
	// UpdateFields Field Mask：本次真正要更新的字段名集合（下划线命名）。2026-08-17 起建议必传：
	// 服务端以它为唯一依据判定更新哪些字段，能表达「改回默认值 / 清空」；不传回退到启发式合并。
	// ⚠️ 字符串/切片字段仍是裸类型 + omitempty，空值不进报文；靠 mask 声明为「要更新」而报文里
	// 又没有的字段，服务端按缺省处理——这正是 mask 的语义，但调用方要清楚自己发的是"清空"。
	// 零值有语义的枚举字段（time_type / limit_day_budget / conversion_type…）已改成指针，
	// 传 &0 就能真正发出 0，不再被 omitempty 吃掉。
	UpdateFields []string `json:"update_fields,omitempty"`
}

// ModifyCreativity 编辑创意信息
type ModifyCreativity struct {
	// CreativityID 创意ID，当modify_type为3时必传
	CreativityID uint64 `json:"creativity_id,omitempty"`
	// CreativityName 创意名称，最多50个字符
	CreativityName string `json:"creativity_name,omitempty"`
	// ConversionType 组件类型，见枚举说明。0-无组件 是合法值，必须显式传
	ConversionType *int `json:"conversion_type,omitempty"`
	// ConversionComponentTypes 组件位置，0-默认位置、1-置顶评论
	ConversionComponentTypes []int `json:"conversion_component_types,omitempty"`
	// ComponentConvNumIsShow 组件展示已转换人数，false-不展示是合法值，必须显式传
	ComponentConvNumIsShow *bool `json:"component_conv_num_is_show,omitempty"`
	// BarContent 搜索组件搜索词，最多6个字符
	BarContent string `json:"bar_content,omitempty"`
	// BarContentUserList 引导文案
	BarContentUserList []string `json:"bar_content_user_list,omitempty"`
	// Comment 单选评论区组件评论内容
	Comment string `json:"comment,omitempty"`
	// CommentUserList 自提置顶评论文案，最多5条
	CommentUserList []string `json:"comment_user_list,omitempty"`
	// ClickURLs 点击监测链接
	ClickURLs []string `json:"click_urls,omitempty"`
	// ExpoURLs 曝光监测链接
	ExpoURLs []string `json:"expo_urls,omitempty"`
	// MaskGen 自动优化封面，1-开启、2-不开启
	MaskGen int `json:"mask_gen,omitempty"`
	// TitleGen 自动优化标题，1-开启、2-不开启
	TitleGen int `json:"title_gen,omitempty"`
	// TitleFills 标题信息
	TitleFills []TitleFill `json:"title_fills,omitempty"`
	// MaskFills 图片信息
	MaskFills []MaskFill `json:"mask_fills,omitempty"`
	// LandingPageType 落地页组件类型，1-表单、2-外跳链接
	LandingPageType int `json:"landing_page_type,omitempty"`
	// JumpURL 外跳链接
	JumpURL string `json:"jump_url,omitempty"`
	// PoiJumpType poi组件-跳转类型，poi-门店详情、poi_list-门店列表
	PoiJumpType string `json:"poi_jump_type,omitempty"`
	// PoiID poi组件-id
	PoiID string `json:"poi_id,omitempty"`
	// UpdateFields Field Mask：本次真正要更新的字段名集合（下划线命名）。2026-08-17 起建议必传：
	// 服务端以它为唯一依据判定更新哪些字段，能表达「改回默认值 / 清空」；不传回退到启发式合并。
	// ⚠️ 字符串/切片字段仍是裸类型 + omitempty，空值不进报文；靠 mask 声明为「要更新」而报文里
	// 又没有的字段，服务端按缺省处理——这正是 mask 的语义，但调用方要清楚自己发的是"清空"。
	// 零值有语义的枚举字段（time_type / limit_day_budget / conversion_type…）已改成指针，
	// 传 &0 就能真正发出 0，不再被 omitempty 吃掉。
	UpdateFields []string `json:"update_fields,omitempty"`
}

// UpdateResponse 新创编编辑 API Response（data为null）
type UpdateResponse struct {
	model.BaseResponse
}
