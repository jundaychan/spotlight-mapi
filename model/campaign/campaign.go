package campaign

// Campaign 计划
type Campaign struct {
	// CampaignID 	计划id
	CampaignID uint64 `json:"campaign_id"`
	// CampaignName 	计划名称
	CampaignName string `json:"campaign_name,omitempty"`
	// CampaignFilterState 计划状态	1-有效，2-暂停，3-已删除，4-计划预算不足，5-现金余额不足，6-所有未删除状态，7-账户日预算不足
	CampaignFilterState int `json:"campaign_filter_state"`
	// CreationType 创建方式：0-标准投 1-简单投全自动 2-留资快投 4-简单投半自动
	CreationType int `json:"creation_type"`
	// CampaignCreateTime 计划创建时间
	CampaignCreateTime string `json:"campaign_create_time,omitempty"`
	// CampaignEnable 计划是否可用
	CampaignEnable int `json:"campaign_enable"`
	// MarketingTarget 营销目标	3-商品销量4-产品种草8-直播推广9-客资收集10-抢占关键词13-种草直达14-直播预热15-店铺拉新 16-应用推广
	MarketingTarget int `json:"marketing_target"`
	// Placement 广告类型	1-信息流2-搜索推广4-全站智投7-视频内流
	Placement int `json:"placement"`
	// OptimizeTarget 优化目标
	OptimizeTarget int `json:"optimize_target"`
	// PromotionTarget 投放标的
	PromotionTarget int `json:"promotion_target"`
	// BiddingStrategy 出价策略
	BiddingStrategy int `json:"bidding_strategy"`
	// ConstraintType 成本控制类型
	ConstraintType int `json:"constraint_type"`
	// ConstraintValue 成本值
	ConstraintValue int `json:"constraint_value"`
	// LimitDayBudget 预算类型	0-不限预算，1-指定预算
	LimitDayBudget int `json:"limit_day_budget"`
	// CampaignDayBudget 指定预算
	CampaignDayBudget int64 `json:"campaign_day_budget"`
	// BudgetState 推广计划日预算是否充足，	0-不足，1-充足
	BudgetState int `json:"budget_state"`
	// SmartSwitch 智能开关
	SmartSwitch int `json:"smart_switch"`
	// Platform 创建来源
	Platform int `json:"platform"`
	// PacingMode 投放速率	1-匀速2-加速
	PacingMode int `json:"pacing_mode"`
	// StartTime 推广开始时间
	StartTime string `json:"start_time,omitempty"`
	// ExpireTime 推广结束时间
	ExpireTime string `json:"expire_time,omitempty"`
	// TimePeriod 推广时间的bitmap
	TimePeriod string `json:"time_period,omitempty"`
	// TimePeroidType 推广时间段类型	0-全时段1-自定义时间段
	TimePeroidType int `json:"time_peroid_type"`
	// FeedFlag 是否开启搜索追投,	0-关1-开
	FeedFlat int `json:"feed_flag"`
	// BuildType 构建类型
	BuildType int `json:"build_type"`
	// CreativityState 总预算达成时间	创意聚合状态，ark电商推广场景使用
	CreativityState int `json:"creativity_state"`
	// EventAssetID 事件资产
	EventAssetID uint64 `json:"event_asset_id"`
	// AssetEvent 资产事件
	AssetEvent int `json:"asset_event"`
	// AssetEventID 资产事件ID
	AssetEventID uint64 `json:"asset_event_id"`
	// PageCategory 页面类目
	PageCategory int `json:"page_category"`
	// SearchFlag 搜索快投开关
	SearchFlat int `json:"search_flag"`
	// SearchBidRatio 定向拓展
	SearchBidRatio float64 `json:"search_bid_ratio"`
	// DeeplinkID deeplink ID
	DeeplinkID uint64 `json:"deeplink_id"`
	// UniversalLinkID universal link ID
	UniversalLinkID uint64 `json:"universal_link_id"`
	// DetectURLLink detect url link
	DetectURLLink string `json:"detect_url_link,omitempty"`
}
