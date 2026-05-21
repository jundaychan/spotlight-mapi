package newcreate

import "github.com/jundaychan/spotlight-mapi/model"

// TimePeriodDTO 投放时段配置
// 每天24个字符，每个小时用0和1表示，0表示不投，1表示投放
type TimePeriodDTO struct {
	// Mon 星期一，默认24个1
	Mon string `json:"mon,omitempty"`
	// Tues 星期二
	Tues string `json:"tues,omitempty"`
	// Wed 星期三
	Wed string `json:"wed,omitempty"`
	// Thur 星期四
	Thur string `json:"thur,omitempty"`
	// Fri 星期五
	Fri string `json:"fri,omitempty"`
	// Sat 星期六
	Sat string `json:"sat,omitempty"`
	// Sun 星期日
	Sun string `json:"sun,omitempty"`
}

// ExploreConfig 一键起量配置
type ExploreConfig struct {
	// CampaignDayBudget 起量预算，单位分
	CampaignDayBudget int64 `json:"campaign_day_budget,omitempty"`
	// TimePeriod 起量时段配置
	TimePeriod *TimePeriodDTO `json:"time_period,omitempty"`
	// TimePeriodType 起量时段类型，0-全时段，1-自定义
	TimePeriodType int `json:"time_period_type,omitempty"`
	// StartTime 起量开始时间，毫秒时间戳
	StartTime int64 `json:"start_time,omitempty"`
	// ExpireHour 起量持续时间，单位小时，1-6
	ExpireHour int64 `json:"expire_hour,omitempty"`
}

// KeywordWithBidDTO 关键词出价信息
type KeywordWithBidDTO struct {
	// Keyword 关键词
	Keyword string `json:"keyword,omitempty"`
	// Bid 出价
	Bid int64 `json:"bid,omitempty"`
	// KeywordSource 关键词来源
	KeywordSource int `json:"keyword_source,omitempty"`
	// PhraseMatchType 关键词匹配方式，0-精确匹配、1-短语匹配
	PhraseMatchType int `json:"phrase_match_type,omitempty"`
	// FeedBid 追投出价
	FeedBid int64 `json:"feed_bid,omitempty"`
}

// TopicDTO 内容主题
type TopicDTO struct {
	// Name 主题名称
	Name string `json:"name,omitempty"`
}

// TopicWithTaxnomyDTO 高潜主题
type TopicWithTaxnomyDTO struct {
	// FirstLevelName 所属类别
	FirstLevelName string `json:"first_level_name,omitempty"`
	// SecondLevelName 高潜主题
	SecondLevelName string `json:"second_level_name,omitempty"`
	// SecondLevelBid 高潜主题出价，单位分
	SecondLevelBid int64 `json:"second_level_bid,omitempty"`
	// SecondLevelPv 预计触达流量
	SecondLevelPv int64 `json:"second_level_pv,omitempty"`
	// SecondLevelRecTag 推荐理由
	SecondLevelRecTag []int `json:"second_level_rec_tag,omitempty"`
	// TopicList 内容主题列表
	TopicList []TopicDTO `json:"topic_list,omitempty"`
}

// CreateTargetInfo 定向信息
type CreateTargetInfo struct {
	// TargetGender 性别定向，0-男，1-女，all-全部，默认all
	TargetGender string `json:"target_gender,omitempty"`
	// TargetCityType 城市定向类型，0-默认，1-线级城市
	TargetCityType int `json:"target_city_type,omitempty"`
	// TargetCity 城市定向
	TargetCity string `json:"target_city,omitempty"`
	// TargetAreaCode 地域编码，-1表示全部，多个用#分隔
	TargetAreaCode string `json:"target_area_code,omitempty"`
	// TargetAge 年龄定向，默认all
	TargetAge string `json:"target_age,omitempty"`
	// TargetDevice 平台定向，ios/android/all，默认all
	TargetDevice string `json:"target_device,omitempty"`
	// TargetDevicePrice 手机价格定向，多个用#分隔，默认all
	TargetDevicePrice string `json:"target_device_price,omitempty"`
	// ReverseTargetCrowd 排除人群（一方上传人群ID列表）
	ReverseTargetCrowd []string `json:"reverse_target_crowd,omitempty"`
	// ReverseConversionType 排除已转化用户类型，1-广告主账户，2-品牌账户，3-APP
	ReverseConversionType int `json:"reverse_conversion_type,omitempty"`
	// ReverseConversionAction 排除已转化用户事件类型
	ReverseConversionAction int `json:"reverse_conversion_action,omitempty"`
	// ReverseConversionDuration 排除已转化用户时间范围，1-当天，2-7天，3-1个月，4-3个月
	ReverseConversionDuration int `json:"reverse_conversion_duration,omitempty"`
	// ReverseConversionAppID 排除已转化用户APP包名
	ReverseConversionAppID string `json:"reverse_conversion_app_id,omitempty"`
	// PremiumTargetCrowd 人群优投（原搜索人群溢价配置）
	PremiumTargetCrowd []PremiumTargetCrowd `json:"premium_target_crowd,omitempty"`
	// CrowdTarget 人群定向（平台精选/行业人群/人群包）
	CrowdTarget *CrowdTarget `json:"crowd_target,omitempty"`
	// BrandInterestGroup 本品牌种草人群定向，为null表示未勾选
	BrandInterestGroup *PlaceholderGroup `json:"brand_interest_group,omitempty"`
	// BrandRecognitionGroup 本品牌认知人群定向，为null表示未勾选
	BrandRecognitionGroup *PlaceholderGroup `json:"brand_recognition_group,omitempty"`
	// CategoryInterestGroup 行业种草人群定向，为null表示未勾选
	CategoryInterestGroup *CategoryInterestGroup `json:"category_interest_group,omitempty"`
	// IndustryInterestTarget 行业兴趣定向
	IndustryInterestTarget *IndustryInterestTarget `json:"industry_interest_target,omitempty"`
	// InterestKeywords 关键词兴趣列表
	InterestKeywords []string `json:"interest_keywords,omitempty"`
	// Keywords 关键词行为列表
	Keywords []string `json:"keywords,omitempty"`
	// KeywordTargetPeriod 关键词行为周期
	KeywordTargetPeriod int `json:"keyword_target_period,omitempty"`
	// KeywordTargetAction 关键词行为类型，1-搜索，2-互动
	KeywordTargetAction []int `json:"keyword_target_action,omitempty"`
	// DandelionCrowd 蒲公英人群定向
	DandelionCrowd *DandelionCrowd `json:"dandelion_crowd,omitempty"`
	// IntelligentExpansion 智能扩量配置
	IntelligentExpansion int `json:"intelligent_expansion,omitempty"`
	// SearchTargetCityIntent 地域意图定向，0-居住地用户，1-居住/意图用户，2-所有用户
	SearchTargetCityIntent int `json:"search_target_city_intent,omitempty"`
	// TargetGeneralizationSwitch 定向拓宽开关
	TargetGeneralizationSwitch int `json:"target_generalization_switch,omitempty"`
}

// PremiumTargetCrowd 人群优投
type PremiumTargetCrowd struct {
	// ID 人群ID
	ID string `json:"id,omitempty"`
	// Ratio 溢价比例
	Ratio string `json:"ratio,omitempty"`
	// Name 人群名称
	Name string `json:"name,omitempty"`
}

// CrowdTarget 人群定向
type CrowdTarget struct {
	// CrowdPkg 人群包列表
	CrowdPkg []CrowdPackage `json:"crowd_pkg,omitempty"`
	// DmpPermission DMP权限状态
	DmpPermission bool `json:"dmp_permission,omitempty"`
}

// CrowdPackage 人群包
type CrowdPackage struct {
	// Value 人群包值
	Value string `json:"value,omitempty"`
	// Name 人群包名称
	Name string `json:"name,omitempty"`
	// GroupID 分组ID
	GroupID string `json:"group_id,omitempty"`
	// SyncStatus 同步状态
	SyncStatus int `json:"sync_status,omitempty"`
	// Status 当前状态
	Status int `json:"status,omitempty"`
	// Type 人群包类型
	Type string `json:"type,omitempty"`
	// Tag 标签
	Tag string `json:"tag,omitempty"`
	// Desc 描述信息
	Desc string `json:"desc,omitempty"`
}

// PlaceholderGroup 占位人群（本品牌种草/认知人群）
type PlaceholderGroup struct {
	// Raw 占位符，无实际意义，传空字符串即可
	Raw string `json:"raw,omitempty"`
}

// CategoryInterestGroup 行业种草人群定向
type CategoryInterestGroup struct {
	// ItemIDs 商品id集合
	ItemIDs []string `json:"itemIds,omitempty"`
}

// IndustryInterestTarget 行业兴趣定向
type IndustryInterestTarget struct {
	// ContentInterests 内容兴趣列表
	ContentInterests []model.CodeNamePair `json:"content_interests,omitempty"`
	// ShoppingInterests 购物兴趣列表
	ShoppingInterests []model.CodeNamePair `json:"shopping_interests,omitempty"`
}

// DandelionCrowd 蒲公英人群定向
type DandelionCrowd struct {
	// NormalDandelionCrowdList 通用蒲公英人群列表
	NormalDandelionCrowdList []NormalDandelionCrowd `json:"normal_dandelion_crowd_list,omitempty"`
	// CustomizedDandelionCrowdList 定制蒲公英人群列表
	CustomizedDandelionCrowdList []CustomizedDandelionCrowd `json:"customized_dandelion_crowd_list,omitempty"`
}

// NormalDandelionCrowd 通用蒲公英人群
type NormalDandelionCrowd struct {
	// ActionType 行为类型，imp/read
	ActionType string `json:"action_type,omitempty"`
	// TimePeriod 时间周期，天数
	TimePeriod int `json:"time_period,omitempty"`
	// BrandUserID 品牌用户ID
	BrandUserID string `json:"brand_user_id,omitempty"`
}

// CustomizedDandelionCrowd 定制蒲公英人群
type CustomizedDandelionCrowd struct {
	// CrowdName 人群包名称
	CrowdName string `json:"crowd_name,omitempty"`
	// NoteIDList 笔记ID集合
	NoteIDList []string `json:"note_id_list,omitempty"`
	// ActionType 行为类型，imp/read
	ActionType string `json:"action_type,omitempty"`
	// Channels 流量渠道，ad/nature
	Channels []string `json:"channels,omitempty"`
	// TimePeriod 时间周期，天数
	TimePeriod int `json:"time_period,omitempty"`
	// CrowdID 人群包ID
	CrowdID uint64 `json:"crowd_id,omitempty"`
}

// H5Info 落地页前链信息
type H5Info struct {
	// PhotoURL 图片链接
	PhotoURL string `json:"photo_url,omitempty"`
	// Width 图片宽度，像素
	Width int `json:"width,omitempty"`
	// Height 图片高度，像素
	Height int `json:"height,omitempty"`
	// Content 文案内容
	Content string `json:"content,omitempty"`
	// ClickURLs 点击监测链接列表
	ClickURLs []string `json:"click_urls,omitempty"`
	// ExpoURLs 曝光监测链接列表
	ExpoURLs []string `json:"expo_urls,omitempty"`
	// MonitorCompany 监测公司名称
	MonitorCompany string `json:"monitor_company,omitempty"`
	// MonitorParams 监测参数配置
	MonitorParams string `json:"monitor_params,omitempty"`
	// Desc 落地页封面描述信息
	Desc string `json:"desc,omitempty"`
}

// TitleFill 标题填充信息
type TitleFill struct {
	// Origin 是否为原标题（是否和创意标题一致）
	Origin bool `json:"origin,omitempty"`
	// Title 标题文本内容
	Title string `json:"title,omitempty"`
	// TitleSource 标题来源，1-自提标题，2-信息流·标题，3-搜索推荐标题，4-搜索&信息流推荐标题
	TitleSource int `json:"title_source,omitempty"`
	// Version 版本标识，只要不是原标题都是v1
	Version string `json:"version,omitempty"`
}

// MaskFill 封面图填充信息
type MaskFill struct {
	// Origin 是否为原图
	Origin bool `json:"origin,omitempty"`
	// FileID 图标唯一标识
	FileID string `json:"file_id,omitempty"`
	// FileKey 上传文件存储键（不要传）
	FileKey string `json:"file_key,omitempty"`
	// ImageSource 图片来源（枚举值）
	ImageSource int `json:"image_source,omitempty"`
}
