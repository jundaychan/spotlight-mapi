package simpledelivery

import (
	"github.com/jundaychan/spotlight-mapi/model"
	"github.com/jundaychan/spotlight-mapi/util"
)

// CreateRequest 新建简单投 API Request
// 当前只支持应用唤起、应用下载、小程序推广营销诉求
type CreateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// BaseConfig 计划
	BaseConfig *UbeSemiBaseConfigDTO `json:"base_config,omitempty"`
	// TargetConfig 定向
	TargetConfig *UbeSemiTargetConfigDTO `json:"target_config,omitempty"`
	// CreativityConfig 创意
	CreativityConfig *UbeSemiCreativityConfigDTO `json:"creativity_config,omitempty"`
}

// Encode implement PostRequest interface
func (r CreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// UbeSemiBaseConfigDTO 简单投计划配置
type UbeSemiBaseConfigDTO struct {
	// CampaignName 计划名称，最多50个字符
	CampaignName string `json:"campaign_name,omitempty"`
	// MarketingIndustry 所属行业，1-电商及交易平台推广，4-短剧小说推广
	MarketingIndustry int `json:"marketing_industry,omitempty"`
	// MarketingTarget 营销诉求，16-应用唤起，20-应用下载，21-小程序推广
	MarketingTarget int `json:"marketing_target,omitempty"`
	// DeliveryMode 投放模式，1-自动投放
	DeliveryMode int `json:"delivery_mode,omitempty"`
	// PromotionTarget 标的类型，1-笔记
	PromotionTarget int `json:"promotion_target,omitempty"`
	// CarrierType 投放载体类型，应用唤起：4-APP，5-小程序，13-不挂链；应用下载：6-IOS，7-Android；小程序推广：8-红书小程序，9-微信小游戏，10-微信小程序
	CarrierType int `json:"carrier_type,omitempty"`
	// WechatAppID 微信小程序/小游戏ID
	WechatAppID string `json:"wechat_app_id,omitempty"`
	// DeeplinkID Deeplink 直达链接 id
	DeeplinkID uint64 `json:"deeplink_id,omitempty"`
	// UniversalLinkID Universal Link 直达链接 id
	UniversalLinkID uint64 `json:"universal_link_id,omitempty"`
	// OptimizeObjective 优化目标，见枚举说明
	OptimizeObjective int `json:"optimize_objective,omitempty"`
	// DeepOptimizeObjective 深度优化目标，不需要时传-1
	DeepOptimizeObjective int `json:"deep_optimize_objective,omitempty"`
	// EventAssetID 事件资产-资产ID
	EventAssetID uint64 `json:"event_asset_id,omitempty"`
	// AssetEventID 事件资产-事件ID
	AssetEventID uint64 `json:"asset_event_id,omitempty"`
	// TimeType 投放时间类型，0-长期投放
	TimeType *int `json:"time_type,omitempty"`
	// TimePeriodType 投放时段类型，0-不限，1-指定时段
	TimePeriodType *int `json:"time_period_type,omitempty"`
	// TimePeriod 投放时段配置
	TimePeriod *TimePeriodDTO `json:"time_period,omitempty"`
	// BiddingStrategy 竞价策略，7-稳定成本
	BiddingStrategy int `json:"bidding_strategy,omitempty"`
	// LimitDayBudget 预算类型，0-不限、1-指定预算
	LimitDayBudget *int `json:"limit_day_budget,omitempty"`
	// OriginCampaignDayBudget 预算金额，单位分
	OriginCampaignDayBudget int64 `json:"origin_campaign_day_budget,omitempty"`
	// DetectURLLink 监测链接URL
	DetectURLLink string `json:"detect_url_link,omitempty"`
	// ConstraintValue 出价，范围在10-499999之间，单位分
	ConstraintValue int64 `json:"constraint_value,omitempty"`
	// PhraseMatchTypeUpgrade 关键词匹配方式升级开关，0-关闭、1-开启，默认-1
	PhraseMatchTypeUpgrade *int `json:"phrase_match_type_upgrade,omitempty"`
	// KeywordWithBid 关键词出价信息
	KeywordWithBid []KeywordWithBidDTO `json:"keyword_with_bid,omitempty"`
	// KeywordGenType 智能拓词类型，1-智能拓词、2-手动+智能关键词定向
	KeywordGenType int `json:"keyword_gen_type,omitempty"`
	// SimpleCreationNoteType 简单投笔记类型，0-全部笔记，1-在投笔记
	SimpleCreationNoteType *int `json:"simple_creation_note_type,omitempty"`
	// AppUniqueID 应用唯一ID，应用下载必传
	AppUniqueID uint64 `json:"app_unique_id,omitempty"`
	// AppDownloadType 应用下载方式，应用下载必传
	AppDownloadType string `json:"app_download_type,omitempty"`
	// PreOrderPage 渠道页
	PreOrderPage string `json:"pre_order_page,omitempty"`
	// PlayletAppID 小程序ID（短剧小程序appid）
	PlayletAppID string `json:"playlet_app_id,omitempty"`
	// PlayletPath 小程序path路径
	PlayletPath string `json:"playlet_path,omitempty"`
	// PlayletAppUserID 小程序userId
	PlayletAppUserID string `json:"playlet_app_user_id,omitempty"`
	// SearchBidRatio 搜索溢价系数
	SearchBidRatio float64 `json:"search_bid_ratio,omitempty"`
	// UbeSubjectType ube标的类型，1-adv账号，2-SPU，3-行业商品，4-小程序/小游戏，5-APP，6-剧集
	UbeSubjectType int `json:"ube_subject_type,omitempty"`
	// UbeSubjectID ube标的id，APP时为app_unique_id，剧集时为剧集ID
	UbeSubjectID string `json:"ube_subject_id,omitempty"`
}

// UbeSemiTargetConfigDTO 简单投定向配置
type UbeSemiTargetConfigDTO struct {
	// TargetType 定向类型，2-智能定向
	TargetType int `json:"target_type,omitempty"`
	// TargetAreaCode 地域定向编码，1表示中国，多个用#分隔
	TargetAreaCode string `json:"target_area_code,omitempty"`
	// SearchTargetCityIntent 地域意图定向，2-所有用户
	SearchTargetCityIntent int `json:"search_target_city_intent,omitempty"`
	// TargetGender 性别定向，默认all，0-男，1-女，all-全部
	TargetGender string `json:"target_gender,omitempty"`
	// TargetAge 年龄定向，默认all
	TargetAge string `json:"target_age,omitempty"`
	// TargetGeneralizationSwitch 定向拓展，1-开启
	TargetGeneralizationSwitch int `json:"target_generalization_switch,omitempty"`
}

// UbeSemiCreativityConfigDTO 简单投创意配置
type UbeSemiCreativityConfigDTO struct {
	// ConversionType 组件类型，0-无组件、6-商品/小程序组件、9-小程序组件、11-唤端组件、13-下载组件、15-红书小程序组件、16-微信小程序组件
	ConversionType *int `json:"conversion_type,omitempty"`
	// ConversionComponentTypes 组件位置，0-默认位置
	ConversionComponentTypes []int `json:"conversion_component_types,omitempty"`
	// BarContentUserList 引导文案
	BarContentUserList []string `json:"bar_content_user_list,omitempty"`
	// AppCompIcon 主图，传图片Base64编码
	AppCompIcon string `json:"app_comp_icon,omitempty"`
	// MaskGen 自动优化封面，0-关闭、1-开启、2-自动优化
	MaskGen *int `json:"mask_gen,omitempty"`
	// TitleGen 自动优化标题，0-关闭、1-开启、2-自动优化
	TitleGen *int `json:"title_gen,omitempty"`
	// FallBackJumpURL 兜底链接
	FallBackJumpURL string `json:"fall_back_jump_url,omitempty"`
	// PrimaryTitle 主标题，最多11个字符
	PrimaryTitle string `json:"primary_title,omitempty"`
	// IOSDownloadLink 唤端组件-ios下载链接
	IOSDownloadLink string `json:"ios_download_link,omitempty"`
	// AndroidDownloadLink 唤端组件-android下载链接
	AndroidDownloadLink string `json:"android_download_link,omitempty"`
	// RemoveNotes 剔除笔记
	RemoveNotes []string `json:"remove_notes,omitempty"`
	// NoteRecState 笔记拓展，0-关闭、1-开启，默认0
	NoteRecState *int `json:"note_rec_state,omitempty"`
	// UbeWhiteBoxList 自选笔记，最多100个
	UbeWhiteBoxList []UbeWhiteBoxDTO `json:"ube_white_box_list,omitempty"`
	// ExcludeNoteTypeList 剔除笔记类型，见 NoteSourceType 枚举
	ExcludeNoteTypeList []int `json:"exclude_note_type_list,omitempty"`
	// NoteTagList 笔记标签ID列表
	NoteTagList []int64 `json:"note_tag_list,omitempty"`
	// ExcludePgyTaskType 剔除的蒲公英笔记任务类型，1-剔除任务期内笔记，2-剔除非任务期内笔记
	ExcludePgyTaskType int `json:"exclude_pgy_task_type,omitempty"`
	// ExcludeKosAuthorList 反选的KOS笔记作者userId列表
	ExcludeKosAuthorList []string `json:"exclude_kos_author_list,omitempty"`
	// ExcludeGrantAuthorList 反选的授权笔记作者userId列表
	ExcludeGrantAuthorList []string `json:"exclude_grant_author_list,omitempty"`
}

// UbeWhiteBoxDTO 自选笔记
type UbeWhiteBoxDTO struct {
	// NoteID 笔记ID
	NoteID string `json:"note_id,omitempty"`
	// NoteDeeplinkID 直达链接-deeplink，投放载体为app时必传
	NoteDeeplinkID uint64 `json:"note_deeplink_id,omitempty"`
	// NoteUniversalLinkID 直达链接-universal link
	NoteUniversalLinkID uint64 `json:"note_universal_link_id,omitempty"`
	// NoteMiniProgramID 微信小程序ID，投放载体为微信小程序时必传
	NoteMiniProgramID string `json:"note_mini_program_id,omitempty"`
}

// CreateResponse 新建简单投 API Response
type CreateResponse struct {
	model.BaseResponse
	// Data 新建简单投返回信息
	Data *UbeSemiCreateAdResponse `json:"data,omitempty"`
}

// UbeSemiCreateAdResponse 新建简单投返回信息
type UbeSemiCreateAdResponse struct {
	// CampaignGroupID 广告组ID
	CampaignGroupID uint64 `json:"campaign_group_id,omitempty"`
	// CampaignID 计划ID
	CampaignID uint64 `json:"campaign_id,omitempty"`
}
