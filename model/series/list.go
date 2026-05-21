package series

import (
	"github.com/jundaychan/spotlight-mapi/model"
	"github.com/jundaychan/spotlight-mapi/util"
)

// ListRequest 可投剧集列表 API Request
type ListRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Type 剧集形态（1=小说 2=短剧 3=漫剧）
	Type EpisodeType `json:"type,omitempty"`
	// IDs 剧集ID集合，精确查询，单次最多20个
	IDs []uint64 `json:"ids,omitempty"`
	// Name 剧集名称，支持模糊查询
	Name string `json:"name,omitempty"`
	// PageNum 页码，从1开始
	PageNum int `json:"pageNum,omitempty"`
	// PageSize 每页数量
	PageSize int `json:"pageSize,omitempty"`
}

// Encode implement PostRequest interface
func (r ListRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// ListResponse 可投剧集列表 API Response
type ListResponse struct {
	model.BaseResponse
	// Data 剧集列表数据
	Data *ListResult `json:"data,omitempty"`
}

// ListResult 剧集列表数据
type ListResult struct {
	// Episodes 剧集对象数组
	Episodes []Episode `json:"episodes,omitempty"`
}

// Episode 剧集对象
type Episode struct {
	// ID 剧集ID
	ID uint64 `json:"id,omitempty"`
	// Type 剧集形态（1=小说 2=短剧 3=漫剧）
	Type EpisodeType `json:"type,omitempty"`
	// CategoryID 所属类目ID
	CategoryID string `json:"category_id,omitempty"`
	// AttributeID 属性ID
	AttributeID string `json:"attribute_id,omitempty"`
	// Name 剧集名称
	Name string `json:"name,omitempty"`
	// Desc 简介
	Desc string `json:"desc,omitempty"`
	// PreviewContent 前N章内容
	PreviewContent string `json:"preview_content,omitempty"`
	// CopyrightOwner 版权方
	CopyrightOwner string `json:"copyright_owner,omitempty"`
	// WordRange 全篇字数范围（仅type=1小说时必填）
	WordRange int `json:"word_range,omitempty"`
	// GenderType 男女频（1=男 2=女 3=其他）
	GenderType GenderType `json:"gender_type,omitempty"`
	// AuthorName 书籍作者（仅type=1小说时必填）
	AuthorName string `json:"author_name,omitempty"`
	// EpisodeCount 总章节/集数
	EpisodeCount int `json:"episode_count,omitempty"`
	// EpisodeDuration 单集时长/分钟（仅type=2短剧时必填）
	EpisodeDuration int `json:"episode_duration,omitempty"`
	// MonetizationType 变现模式（1=付费 2=混合 3=免费）
	MonetizationType MonetizationType `json:"monetization_type,omitempty"`
	// HasNetworkContent 含网赚内容（0=否 1=是）
	HasNetworkContent int `json:"has_network_content,omitempty"`
	// AdUnlockStart 起始广告解锁集数
	AdUnlockStart int `json:"ad_unlock_start,omitempty"`
	// PayUnlockStart 起始付费解锁集数
	PayUnlockStart int `json:"pay_unlock_start,omitempty"`
	// UnitPrice 单集/章节价格（元）
	UnitPrice int64 `json:"unit_price,omitempty"`
	// MemberType 会员类型
	MemberType MemberType `json:"member_type,omitempty"`
	// MaxRecharge 单次最高充值档位（分）
	MaxRecharge int64 `json:"max_recharge,omitempty"`
	// MinRecharge 单次最低充值档位（分）
	MinRecharge int64 `json:"min_recharge,omitempty"`
	// RecommendRecharge 推荐充值档位（分）
	RecommendRecharge int64 `json:"recommend_recharge,omitempty"`
	// Icon 剧集icon图片URL
	Icon string `json:"icon,omitempty"`
	// PrimaryTitle 投放按钮主标题文案
	PrimaryTitle string `json:"primary_title,omitempty"`
	// DeeplinkURL 应用唤起深度链接
	DeeplinkURL string `json:"deeplink_url,omitempty"`
	// UniversalLink iOS通用链接
	UniversalLink string `json:"universal_link,omitempty"`
	// WxAppID 微信小程序原始ID
	WxAppID string `json:"wx_app_id,omitempty"`
	// WxPath 微信小程序页面路径
	WxPath string `json:"wx_path,omitempty"`
	// DetectURLApp 应用推广监测链接
	DetectURLApp string `json:"detect_url_app,omitempty"`
	// DetectURLMp 小程序推广监测链接
	DetectURLMp string `json:"detect_url_mp,omitempty"`
	// AuditStatus 审核状态
	AuditStatus int `json:"audit_status,omitempty"`
	// AuditComment 审核备注（如审核失败原因）
	AuditComment string `json:"audit_comment,omitempty"`
}
