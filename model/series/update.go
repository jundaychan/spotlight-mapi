package series

import (
	"github.com/bububa/spotlight-mapi/model"
	"github.com/bububa/spotlight-mapi/util"
)

// UpdateEpisode 剧集修改信息（不修改的字段可不传）
type UpdateEpisode struct {
	// ID 要更新的剧集ID（唯一标识）
	ID uint64 `json:"id,omitempty"`
	// Type 剧集形态（不可修改，原值传）1=短剧 2=小说 3=漫剧
	Type EpisodeType `json:"type,omitempty"`
	// CategoryID 类目ID（接口不生效，仅占位，不可修改）
	CategoryID string `json:"category_id,omitempty"`
	// Name 剧集名称（接口不生效，仅占位，不可修改）
	Name string `json:"name,omitempty"`
	// Desc 简介（可变更），最多500字
	Desc string `json:"desc,omitempty"`
	// PreviewContent 前N章内容（可变更），最多1w字
	PreviewContent string `json:"preview_content,omitempty"`
	// CopyrightOwner 版权方（可变更）
	CopyrightOwner string `json:"copyright_owner,omitempty"`
	// WordRange 全篇字数范围（1=1w~10w 2=10w~20w 3=20w+）
	WordRange int `json:"word_range,omitempty"`
	// GenderType 男女频（1=男 2=女 3=其他）
	GenderType GenderType `json:"gender_type,omitempty"`
	// AuthorName 书籍作者（可变更）
	AuthorName string `json:"author_name,omitempty"`
	// EpisodeCount 章节/集数（可变更）
	EpisodeCount int `json:"episode_count,omitempty"`
	// EpisodeDuration 单集时长/分钟（可变更）
	EpisodeDuration int `json:"episode_duration,omitempty"`
	// MonetizationType 变现模式（1=付费 2=混合 3=免费）
	MonetizationType MonetizationType `json:"monetization_type,omitempty"`
	// HasNetworkContent 含网赚内容（0=否 1=是）
	HasNetworkContent int `json:"has_network_content,omitempty"`
	// AdUnlockStart 起始广告解锁集数（可变更）
	AdUnlockStart int `json:"ad_unlock_start,omitempty"`
	// PayUnlockStart 起始付费解锁集数（可变更）
	PayUnlockStart int `json:"pay_unlock_start,omitempty"`
	// UnitPrice 单集价格/分（可变更）
	UnitPrice int64 `json:"unit_price,omitempty"`
	// MemberType 会员类型（1=终身 2=年度 3=月度 4=周/天 5=无）
	MemberType MemberType `json:"member_type,omitempty"`
	// MaxRecharge 单次最高充值档位/分（可变更）
	MaxRecharge int64 `json:"max_recharge,omitempty"`
	// MinRecharge 单次最低充值档位/分（可变更）
	MinRecharge int64 `json:"min_recharge,omitempty"`
	// RecommendRecharge 推荐充值档位/分（可变更）
	RecommendRecharge int64 `json:"recommend_recharge,omitempty"`
	// Icon 剧集icon URL（可变更）
	Icon string `json:"icon,omitempty"`
	// PrimaryTitle 转化组件文案（可变更），11字
	PrimaryTitle string `json:"primary_title,omitempty"`
	// DeeplinkURL 应用唤起深度链接（可变更），2000字
	DeeplinkURL string `json:"deeplink_url,omitempty"`
	// UniversalLink iOS通用链接（可变更），2000字
	UniversalLink string `json:"universal_link,omitempty"`
	// WxAppID 微信小程序原始ID（可变更）
	WxAppID string `json:"wx_appid,omitempty"`
	// WxPath 微信小程序页面路径（可变更）
	WxPath string `json:"wx_path,omitempty"`
	// DetectURLApp 应用推广监测链接（可变更），3000字
	DetectURLApp string `json:"detect_url_app,omitempty"`
	// DetectURLMp 小程序推广监测链接（可变更），3000字
	DetectURLMp string `json:"detect_url_mp,omitempty"`
}

// UpdateRequest 剧集修改 API Request（支持批量更新，单次不超过10个剧集）
type UpdateRequest struct {
	// AdvertiserID 广告主/企业账户ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Episodes 待更新的剧集信息数组
	Episodes []UpdateEpisode `json:"episodes,omitempty"`
}

// Encode implement PostRequest interface
func (r UpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// UpdateResponse 剧集修改 API Response
type UpdateResponse struct {
	model.BaseResponse
	// Data 更新结果
	Data *UpdateResult `json:"data,omitempty"`
}

// UpdateResult 更新结果
type UpdateResult struct {
	// List 成功更新的剧集ID列表
	List []EpisodeID `json:"list,omitempty"`
}
