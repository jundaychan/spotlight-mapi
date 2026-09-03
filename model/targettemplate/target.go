package targettemplate

import (
	"github.com/jundaychan/spotlight-mapi/model"
)

// MarketingTarget 营销诉求
// 3-商品销量 4-产品种草 8-直播推广 9-客资收集 10-抢占关键词 13-种草直达 14-直播预热 15-店铺拉新 16-应用推广
type MarketingTarget int

// Placement 广告类型
// 1-信息流 2-搜索推广 4-全站智投 7-视频内流
type Placement int

// TargetTemplateType 定向类型
// 1-通投 2-智能定向 3-高级定向
type TargetTemplateType int

// DeliveryMode 投放模式
// 0-手动投放 1-自动投放
type DeliveryMode int

// ApplyAction 关联操作
// 1-关联 2-取消关联
type ApplyAction int

// KeywordTargetAction 关键词行为类型
// 1-搜索 2-互动 3-阅读
type KeywordTargetAction int

// TargetConfig 定向配置
type TargetConfig struct {
	// TargetGender 性别 不限: all，男: 0， 女: 1
	TargetGender string `json:"target_gender,omitempty"`
	// TargetAge 年龄，不限：all, 细分年龄段如 18-22#23-27
	TargetAge string `json:"target_age,omitempty"`
	// TargetDevice 设备，不限: all, 苹果: ios, 安卓: android
	TargetDevice string `json:"target_device,omitempty"`
	// TargetDevicePrice 设备价格，不限: all, 0-1999/2000-3999/4000-5999/6000-7999/8000-100000
	TargetDevicePrice string `json:"target_device_price,omitempty"`
	// TargetCity 城市定向，例如"北京#天津#衡阳"，all 代表全部省市
	TargetCity string `json:"target_city,omitempty"`
	// TargetAreaCode 地域定向，不限传 -1，多个地域用#分隔，如 615#616
	TargetAreaCode string `json:"target_area_code,omitempty"`
	// IndustryInterestTarget 行业兴趣定向
	IndustryInterestTarget *IndustryInterestTarget `json:"industry_interest_target,omitempty"`
	// CrowdTarget dmp人群包定向
	CrowdTarget *CrowdTarget `json:"crowd_target,omitempty"`
	// InterestKeywords 关键词兴趣
	InterestKeywords []string `json:"interest_keywords,omitempty"`
	// Keywords 关键词行为
	Keywords []string `json:"keywords,omitempty"`
	// IntelligentExpansion 智能扩量，0: 否，1: 是
	IntelligentExpansion *int `json:"intelligent_expansion,omitempty"`
	// SearchTargetCityIntent 搜索地域意图定向功能 0-关闭，1-开启
	SearchTargetCityIntent string `json:"search_target_city_intent,omitempty"`
	// ReverseTargetCrowd 排除特定人群
	ReverseTargetCrowd []string `json:"reverse_target_crowd,omitempty"`
	// KeywordTargetPeriod 关键词时间周期，单位天，枚举包括 3，7，15，30
	KeywordTargetPeriod int `json:"keyword_target_period,omitempty"`
	// KeywordTargetAction 关键词行为类型，1: 搜索，2: 互动，3: 阅读
	KeywordTargetAction []KeywordTargetAction `json:"keyword_target_action,omitempty"`
}

// IndustryInterestTarget 行业兴趣定向
type IndustryInterestTarget struct {
	// ContentInterests 行业内容兴趣
	ContentInterests []model.CodeNamePair `json:"content_interests,omitempty"`
	// ShoppingInterests 行业购物兴趣（需要填写至二级）
	ShoppingInterests []model.CodeNamePair `json:"shopping_interests,omitempty"`
}

// CrowdTarget dmp人群包定向
type CrowdTarget struct {
	// CrowdPkg 人群包列表
	CrowdPkg []CrowdPackage `json:"crowd_pkg,omitempty"`
	// DmpPermission dmp权限
	DmpPermission bool `json:"dmp_permission,omitempty"`
}

// CrowdPackage 人群包
type CrowdPackage struct {
	// GroupID 人群包真实ID
	GroupID string `json:"group_id,omitempty"`
	// Value 人群包ID
	Value string `json:"value,omitempty"`
	// Name 人群包名称
	Name string `json:"name,omitempty"`
	// Tag 人群包类型，平台精选场景人群：common 节促人群：timeliness 行业策略人群：industry_strategy 行业垂类人群：industry_category
	Tag string `json:"tag,omitempty"`
	// Type 人群包类型
	Type string `json:"type,omitempty"`
	// SyncStatus 同步状态,0:未同步,1:已同步
	SyncStatus int `json:"sync_status,omitempty"`
	// Status 删除可用状态,2:成功,3:计算失败,4:停止更新
	Status int `json:"status,omitempty"`
}
