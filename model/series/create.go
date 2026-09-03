package series

import (
	"github.com/jundaychan/spotlight-mapi/model"
	"github.com/jundaychan/spotlight-mapi/util"
)

// EpisodeType 剧集形态
type EpisodeType int

const (
	// EpisodeTypeNovel 小说
	EpisodeTypeNovel EpisodeType = 1
	// EpisodeTypeShortDrama 短剧
	EpisodeTypeShortDrama EpisodeType = 2
	// EpisodeTypeComicDrama 漫剧
	EpisodeTypeComicDrama EpisodeType = 3
)

// GenderType 男女频
type GenderType int

const (
	// GenderTypeMale 男
	GenderTypeMale GenderType = 1
	// GenderTypeFemale 女
	GenderTypeFemale GenderType = 2
	// GenderTypeOther 其他
	GenderTypeOther GenderType = 3
)

// MonetizationType 变现模式
type MonetizationType int

const (
	// MonetizationTypePaid 付费
	MonetizationTypePaid MonetizationType = 1
	// MonetizationTypeMixed 混合
	MonetizationTypeMixed MonetizationType = 2
	// MonetizationTypeFree 免费
	MonetizationTypeFree MonetizationType = 3
)

// MemberType 会员类型
type MemberType int

const (
	// MemberTypeLifetime 终身
	MemberTypeLifetime MemberType = 1
	// MemberTypeYearly 年度
	MemberTypeYearly MemberType = 2
	// MemberTypeMonthly 月度
	MemberTypeMonthly MemberType = 3
	// MemberTypeWeeklyOrDaily 周/天
	MemberTypeWeeklyOrDaily MemberType = 4
	// MemberTypeNone 无
	MemberTypeNone MemberType = 5
)

// CreateEpisode 剧集创建信息
type CreateEpisode struct {
	// Type 剧集形态 1=小说, 2=短剧, 3=漫剧
	Type EpisodeType `json:"type,omitempty"`
	// CategoryID 类目ID
	CategoryID string `json:"category_id,omitempty"`
	// AttributeID 属性ID
	AttributeID string `json:"attribete_id,omitempty"`
	// Name 剧集名称，企业号下唯一，最多50字，不支持《》
	Name string `json:"name,omitempty"`
	// Desc 简介，最多500字
	Desc string `json:"desc,omitempty"`
	// PreviewContent 前N章内容，最多1w字
	PreviewContent string `json:"preview_content,omitempty"`
	// CopyrightOwner 版权方
	CopyrightOwner string `json:"copyright_owner,omitempty"`
	// WordRange 全篇字数 1=2w以下 2=2w-5w 3=5w-10w 4=10w-20w 5=20w以上
	WordRange int `json:"word_range,omitempty"`
	// GenderType 男女频 1=男 2=女 3=其他
	GenderType GenderType `json:"gender_type,omitempty"`
	// AuthorName 书籍作者
	AuthorName string `json:"author_name,omitempty"`
	// EpisodeCount 章节/集数，上限<100000
	EpisodeCount int `json:"episode_count,omitempty"`
	// EpisodeDuration 单集时长（分钟），仅短剧，上限<100000
	EpisodeDuration int `json:"episode_duration,omitempty"`
	// MonetizationType 变现模式 1=付费 2=混合 3=免费
	MonetizationType MonetizationType `json:"monetization_type,omitempty"`
	// HasNetworkContent 含网赚内容 0=否 1=是
	HasNetworkContent *int `json:"has_network_content,omitempty"`
	// AdUnlockStart 起始广告解锁集数，仅免费/混合，上限<100000
	AdUnlockStart int `json:"ad_unlock_start,omitempty"`
	// PayUnlockStart 起始付费解锁集数，仅付费/混合，上限<100000
	PayUnlockStart int `json:"pay_unlock_start,omitempty"`
	// UnitPrice 单集价格（分），仅整数，上限<100000
	UnitPrice int64 `json:"unit_price,omitempty"`
	// MemberType 会员类型 1=终身 2=年度 3=月度 4=周/天 5=无
	MemberType MemberType `json:"member_type,omitempty"`
	// MaxRecharge 单次最高充值档位（分）
	MaxRecharge int64 `json:"max_recharge,omitempty"`
	// MinRecharge 单次最低充值档位（分）
	MinRecharge int64 `json:"min_recharge,omitempty"`
	// RecommendRecharge 推荐充值档位（分）
	RecommendRecharge int64 `json:"recommend_recharge,omitempty"`
	// Icon 剧集icon URL，180*180
	Icon string `json:"icon,omitempty"`
	// PrimaryTitle 转化组件文案，11字以内
	PrimaryTitle string `json:"primary_title,omitempty"`
	// DeeplinkURL 应用唤起深度链接，2000字以内
	DeeplinkURL string `json:"deeplink_url,omitempty"`
	// UniversalLink iOS通用链接，2000字以内
	UniversalLink string `json:"universal_link,omitempty"`
	// WxAppID 微信小程序原始ID
	WxAppID string `json:"wx_appid,omitempty"`
	// WxPath 微信小程序页面路径
	WxPath string `json:"wx_path,omitempty"`
	// DetectURLApp 应用推广监测链接，3000字以内
	DetectURLApp string `json:"detect_url_app,omitempty"`
	// DetectURLMp 小程序推广监测链接，3000字以内
	DetectURLMp string `json:"detect_url_mp,omitempty"`
}

// CreateRequest 剧集创建 API Request
type CreateRequest struct {
	// AdvertiserID 广告主/企业账户ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Episodes 剧集信息数组，可批量创建
	Episodes []CreateEpisode `json:"episodes,omitempty"`
}

// Encode implement PostRequest interface
func (r CreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// CreateResponse 剧集创建 API Response
type CreateResponse struct {
	model.BaseResponse
	// Data 创建结果
	Data *CreateResult `json:"data,omitempty"`
}

// CreateResult 创建结果
type CreateResult struct {
	// List 成功创建的剧集ID列表
	List []EpisodeID `json:"list,omitempty"`
}

// EpisodeID 剧集ID
type EpisodeID struct {
	// ID 剧集ID
	ID uint64 `json:"id,omitempty"`
}
