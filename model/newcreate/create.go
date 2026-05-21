package newcreate

import (
	"github.com/bububa/spotlight-mapi/model"
	"github.com/bububa/spotlight-mapi/util"
)

// CreateRequest 新创编 API Request（级联创建计划/单元/创意）
type CreateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// CreateType 创建类型，1-创建三层、2-在已有计划下创建单元+创意、3-在已有单元下创建创意，默认1
	CreateType int `json:"create_type,omitempty"`
	// CreateCascadeInfoList 新创编数据
	CreateCascadeInfoList []CreateCascadeInfo `json:"create_cascade_info_list,omitempty"`
}

// Encode implement PostRequest interface
func (r CreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// CreateCascadeInfo 新创编级联数据
type CreateCascadeInfo struct {
	// Campaign 关联的广告计划信息
	Campaign *CreateCampaign `json:"campaign,omitempty"`
	// UnitWithCreativeList 包含广告单元及创意信息的列表
	UnitWithCreativeList []CreateUnitWithCreative `json:"unit_with_creative_list,omitempty"`
}

// CreateCampaign 计划信息
type CreateCampaign struct {
	// CampaignID 计划ID，当create_type为2/3时必传
	CampaignID uint64 `json:"campaign_id,omitempty"`
	// CampaignName 计划名称，最多50个字符
	CampaignName string `json:"campaign_name,omitempty"`
	// MarketingIndustry 行业类型，1-电商，2-游戏，3-网服工具，4-短剧小说
	MarketingIndustry int `json:"marketing_industry,omitempty"`
	// MarketingTarget 营销诉求，4-产品种草、9-客资收集、13-种草直达、16-应用唤起、20-应用下载、21-小程序推广
	MarketingTarget int `json:"marketing_target,omitempty"`
	// Placement 投放位置，1-信息流推广、2-搜索推广、4-全站、7-视频流推广
	Placement int `json:"placement,omitempty"`
	// DeliveryMode 投放模式，0-手动投放，1-自动投放，默认-1
	DeliveryMode int `json:"delivery_mode,omitempty"`
	// PromotionTarget 标的类型，1-笔记，9-落地页
	PromotionTarget int `json:"promotion_target,omitempty"`
	// SpuIDs 推广SPU
	SpuIDs []string `json:"spu_ids,omitempty"`
	// AdBizItemID 行业商品/小程序ID，营销诉求为种草直达时必填
	AdBizItemID string `json:"ad_biz_item_id,omitempty"`
	// CarrierType 投放载体类型，见枚举说明
	CarrierType int `json:"carrier_type,omitempty"`
	// PageCategory 落地页类型，1-聚光落地页、2-自研落地页、3-原生落地页
	PageCategory int `json:"page_category,omitempty"`
	// PageID 落地页ID，落地页类型为聚光落地页时必传
	PageID string `json:"page_id,omitempty"`
	// LandingPageURL 落地页URL，落地页类型为自研落地页时必传
	LandingPageURL string `json:"landing_page_url,omitempty"`
	// WechatAppID 微信小程序/小游戏ID
	WechatAppID string `json:"wechat_app_id,omitempty"`
	// JumpURLType 唤起链接类型
	JumpURLType int `json:"jump_url_type,omitempty"`
	// DeeplinkID Deeplink 直达链接 id
	DeeplinkID uint64 `json:"deeplink_id,omitempty"`
	// UniversalLinkID Universal Link 直达链接 id
	UniversalLinkID uint64 `json:"universal_link_id,omitempty"`
	// PreOrderPage APP预约下载渠道页链接
	PreOrderPage string `json:"pre_order_page,omitempty"`
	// PlayletAppID 小程序AppID
	PlayletAppID string `json:"playlet_app_id,omitempty"`
	// PlayletPath 小程序Path路径
	PlayletPath string `json:"playlet_path,omitempty"`
	// PlayletAppUserID 小程序对应的userId
	PlayletAppUserID string `json:"playlet_app_user_id,omitempty"`
	// OptimizeObjective 优化目标，见枚举说明
	OptimizeObjective int `json:"optimize_objective,omitempty"`
	// DeepOptimizeObjective 深度优化目标，不需要时传-1
	DeepOptimizeObjective int `json:"deep_optimize_objective,omitempty"`
	// ThirdPlatform 转化平台，1-小红星，2-小红盟
	ThirdPlatform int `json:"third_platform,omitempty"`
	// EventAssetID 事件资产-资产ID
	EventAssetID uint64 `json:"event_asset_id,omitempty"`
	// AssetEventID 事件资产-事件ID
	AssetEventID uint64 `json:"asset_event_id,omitempty"`
	// TimeType 投放时间类型，0-长期投放，1-设置起止日期
	TimeType int `json:"time_type,omitempty"`
	// StartTime 投放开始日期，示例 2023-09-20
	StartTime string `json:"start_time,omitempty"`
	// ExpireTime 投放结束日期，示例 2023-09-21
	ExpireTime string `json:"expire_time,omitempty"`
	// TimePeriodType 投放时段类型，0-不限，1-指定时段
	TimePeriodType int `json:"time_period_type,omitempty"`
	// TimePeriod 投放时段配置
	TimePeriod *TimePeriodDTO `json:"time_period,omitempty"`
	// BiddingStrategy 竞价策略，2-手动出价、3-最大转化、7-稳定成本
	BiddingStrategy int `json:"bidding_strategy,omitempty"`
	// LimitDayBudget 预算类型，0-不限、1-指定预算
	LimitDayBudget int `json:"limit_day_budget,omitempty"`
	// OriginCampaignDayBudget 预算金额，单位分
	OriginCampaignDayBudget int64 `json:"origin_campaign_day_budget,omitempty"`
	// SmartSwitch 节假日预算上浮，0-关闭、1-开启
	SmartSwitch int `json:"smart_switch,omitempty"`
	// ExploreState 一键起量开关，0-关闭、1-开启
	ExploreState int `json:"explore_state,omitempty"`
	// ExploreConfig 一键起量配置
	ExploreConfig *ExploreConfig `json:"explore_config,omitempty"`
	// ExploreStatus 一键起量状态
	ExploreStatus int `json:"explore_status,omitempty"`
	// PacingMode 消耗速率，1-匀速消耗、2-加速消耗
	PacingMode int `json:"pacing_mode,omitempty"`
	// SearchFlag 搜索快投开关，-1-默认、0-关闭、1-开启
	SearchFlag int `json:"search_flag,omitempty"`
	// HorseRace 笔记赛马状态，0-从未开启、1-开启、2-开启后关闭
	HorseRace int `json:"horse_race,omitempty"`
	// DetectURLLink 监测链接URL
	DetectURLLink string `json:"detect_url_link,omitempty"`
	// UgLegalAgreementURL 法务协议URL
	UgLegalAgreementURL string `json:"ug_legal_agreement_url,omitempty"`
	// CampaignGroupID 广告组ID
	CampaignGroupID uint64 `json:"campaign_group_id,omitempty"`
	// CampaignGroupName 广告组名称（不用传）
	CampaignGroupName string `json:"campaign_group_name,omitempty"`
	// AppDownloadType App下载方式，app_store/apk/channelApk
	AppDownloadType string `json:"app_download_type,omitempty"`
	// AppUniqueID App主键ID
	AppUniqueID uint64 `json:"app_unique_id,omitempty"`
	// AppID 应用id
	AppID string `json:"app_id,omitempty"`
}

// CreateUnitWithCreative 单元及创意信息
type CreateUnitWithCreative struct {
	// Unit 单元信息
	Unit *CreateUnit `json:"unit,omitempty"`
	// CreativityList 创意信息，一次最多新建20个创意
	CreativityList []CreateCreativity `json:"creativity_list,omitempty"`
}

// CreateUnit 单元信息
type CreateUnit struct {
	// UnitID 单元ID，当create_type为3时必传
	UnitID uint64 `json:"unit_id,omitempty"`
	// UnitName 单元名称，最多50个字符
	UnitName string `json:"unit_name,omitempty"`
	// PhraseMatchTypeUpgrade 关键词匹配方式升级开关，0-关闭、1-开启，默认-1
	PhraseMatchTypeUpgrade int `json:"phrase_match_type_upgrade,omitempty"`
	// KeywordWithBid 关键词出价信息
	KeywordWithBid []KeywordWithBidDTO `json:"keyword_with_bid,omitempty"`
	// TopicWithTaxnomy 高潜主题
	TopicWithTaxnomy []TopicWithTaxnomyDTO `json:"topic_with_taxnomy,omitempty"`
	// FakeUnitID 推词埋点用伪单元ID
	FakeUnitID uint64 `json:"fake_unit_id,omitempty"`
	// KeywordGenType 智能拓词类型，-1-无意义、0-手动选词、1-智能拓词、2-手动+智能
	KeywordGenType int `json:"keyword_gen_type,omitempty"`
	// CtrConstraint 智能拓词-点击率约束
	CtrConstraint float64 `json:"ctr_constraint,omitempty"`
	// AcpConstraint 智能拓词-点击成本约束
	AcpConstraint int64 `json:"acp_constraint,omitempty"`
	// TargetPosition 目标位次，0-不限位置、1-首位、3-前三位
	TargetPosition int `json:"target_position,omitempty"`
	// TargetTemplateID 定向包模板ID
	TargetTemplateID uint64 `json:"target_template_id,omitempty"`
	// TargetType 定向类型，0-默认值、2-智能定向、3-高级定向
	TargetType int `json:"target_type,omitempty"`
	// TargetInfo 定向信息
	TargetInfo *CreateTargetInfo `json:"target_info,omitempty"`
	// EventBid 出价，非roi单位为分，roi单位为%
	EventBid int64 `json:"event_bid,omitempty"`
	// SearchBidRatio 搜索出价系数，默认1.0
	SearchBidRatio float64 `json:"search_bid_ratio,omitempty"`
	// AigcNoteBlackRec aigc 智能笔记，0-关闭、1-开启
	AigcNoteBlackRec int `json:"aigc_note_black_rec,omitempty"`
}

// CreateCreativity 创意信息
type CreateCreativity struct {
	// CreativityName 创意名称，最多50个字符
	CreativityName string `json:"creativity_name,omitempty"`
	// NoteSourceType 笔记类型，4-ark代理笔记
	NoteSourceType int `json:"note_source_type,omitempty"`
	// NoteRecType 笔记优选类型，0-默认自提笔记，1-spu推荐等
	NoteRecType int `json:"note_rec_type,omitempty"`
	// SubstitutedUserID 代投标识
	SubstitutedUserID string `json:"substituted_user_id,omitempty"`
	// KosMsgType 员工笔记私信承接方，0-tok，1-tob
	KosMsgType int `json:"kos_msg_type,omitempty"`
	// KobMsgType 授权笔记私信承接方，1-B承接，2-b承接
	KobMsgType int `json:"kob_msg_type,omitempty"`
	// NoteID 笔记ID
	NoteID string `json:"note_id,omitempty"`
	// AdBizSpuID 结构化投放spu id
	AdBizSpuID string `json:"ad_biz_spu_id,omitempty"`
	// AdBizSeriesID 结构化投放系列 id
	AdBizSeriesID string `json:"ad_biz_series_id,omitempty"`
	// AdBizItemID 推广商品/小程序
	AdBizItemID string `json:"ad_biz_item_id,omitempty"`
	// PageID 落地页ID
	PageID string `json:"page_id,omitempty"`
	// H5Infos 落地页前链信息
	H5Infos []H5Info `json:"h5_infos,omitempty"`
	// ConversionType 组件类型，见枚举说明
	ConversionType int `json:"conversion_type,omitempty"`
	// ConversionComponentTypes 组件位置，0-默认位置、1-置顶评论
	ConversionComponentTypes []int `json:"conversion_component_types,omitempty"`
	// SecondJumpPattern 组件样式，twinAdsNoteBar-双行样式、engagebar-胶囊样式
	SecondJumpPattern string `json:"second_jump_pattern,omitempty"`
	// ComponentConvNumIsShow 组件展示已转换人数
	ComponentConvNumIsShow bool `json:"component_conv_num_is_show,omitempty"`
	// BarContent 自定义引导文案，最多6个字符
	BarContent string `json:"bar_content,omitempty"`
	// BarContentUserList 引导文案
	BarContentUserList []string `json:"bar_content_user_list,omitempty"`
	// Comment 单选评论区组件评论内容
	Comment string `json:"comment,omitempty"`
	// CommentUserList 自提置顶评论文案，最多5条
	CommentUserList []string `json:"comment_user_list,omitempty"`
	// AppCompIcon 主图，传图片Base64编码
	AppCompIcon string `json:"app_comp_icon,omitempty"`
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
	// ExternalGoodsCover 商品主图，传图片Base64编码
	ExternalGoodsCover string `json:"external_goods_cover,omitempty"`
	// DataPostURL 数据推送url
	DataPostURL string `json:"data_post_url,omitempty"`
	// NeedTextMessage 需要短信验证
	NeedTextMessage bool `json:"need_text_message,omitempty"`
	// MiniProgramPath 小程序链接
	MiniProgramPath string `json:"mini_program_path,omitempty"`
	// ExtensionContent 红书小程序主标题
	ExtensionContent string `json:"extension_content,omitempty"`
	// FallBackJumpURL 兜底链接
	FallBackJumpURL string `json:"fall_back_jump_url,omitempty"`
	// PrimaryTitle 主标题，最多11个字符
	PrimaryTitle string `json:"primary_title,omitempty"`
	// ActionButtonContent 唤端组件-按钮文案（不用传）
	ActionButtonContent string `json:"action_button_content,omitempty"`
	// IOSDownloadLink 唤端组件-ios下载链接
	IOSDownloadLink string `json:"ios_download_link,omitempty"`
	// AndroidDownloadLink 唤端组件-android下载链接
	AndroidDownloadLink string `json:"android_download_link,omitempty"`
	// MessageTimePeriod 私信表单同投-接线时段(24bit格式)
	MessageTimePeriod string `json:"message_time_period,omitempty"`
}

// CreateResponse 新创编 API Response
type CreateResponse struct {
	model.BaseResponse
	// Data 级联创建结果
	Data *CreateResponseData `json:"data,omitempty"`
}

// CreateResponseData 级联创建结果数据
type CreateResponseData struct {
	// InfoList 级联创建结果数据列表
	InfoList []CascadeInfoResult `json:"info_list,omitempty"`
}

// CascadeInfoResult 级联创建结果
type CascadeInfoResult struct {
	// Campaign 计划结果
	Campaign *CampaignResult `json:"campaign,omitempty"`
	// UnitWithCreativeList 单元与创意结果列表
	UnitWithCreativeList []UnitWithCreativeResult `json:"unit_with_creative_list,omitempty"`
}

// CampaignResult 计划结果
type CampaignResult struct {
	// CampaignID 计划ID
	CampaignID uint64 `json:"campaign_id,omitempty"`
}

// UnitWithCreativeResult 单元与创意结果
type UnitWithCreativeResult struct {
	// Unit 单元结果
	Unit *UnitResult `json:"unit,omitempty"`
	// CreativityList 创意结果列表
	CreativityList []CreativityResult `json:"creativity_list,omitempty"`
}

// UnitResult 单元结果
type UnitResult struct {
	// UnitID 单元ID
	UnitID uint64 `json:"unit_id,omitempty"`
}

// CreativityResult 创意结果
type CreativityResult struct {
	// CreativityID 创意ID
	CreativityID uint64 `json:"creativity_id,omitempty"`
}
