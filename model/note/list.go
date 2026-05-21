package note

import (
	"github.com/bububa/spotlight-mapi/model"
	"github.com/bububa/spotlight-mapi/util"
)

// ListRequest 获取笔记列表 API Request
type ListRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// NoteType 笔记类型 1：我的笔记 2：合作笔记 4：主理人笔记 6:员工笔记 11：授权笔记 12:素材笔记 13:合作码笔记（参数已传13系统自动转换为14，后续合作码笔记传14） 14:合作码笔记
	NoteType int `json:"note_type,omitempty"`
	// Keyword 搜索关键词 笔记id或者笔记名称，传多个笔记id时以','分隔，最大数量不超过100
	Keyword string `json:"keyword,omitempty"`
	// OrderField 排序字段，可选 阅读数:read_count 互动数:interact_count 阅读率:read_rate 互动率:interact_rate 创建时间:create_time
	OrderField string `json:"order_field,omitempty"`
	// OrderType 升降序 正排:asc 倒排:desc 默认desc
	OrderType string `json:"order_type,omitempty"`
	// NoteContentType 1：图文笔记 2：视频笔记
	NoteContentType int `json:"note_content_type,omitempty"`
	// PlacementType 推广场景：1：信息流推广 2：搜索推广 3：开屏推广 4：全站智投 7：视频内流
	PlacementType int `json:"placement_type,omitempty"`
	// SpuID spu_id（白名单支持）
	SpuID string `json:"spu_id,omitempty"`
	// FilterTaobao 是否只展示小红星笔记 0:展示所有 1:仅小红星
	FilterTaobao int `json:"filter_taobao,omitempty"`
	// MarketTarget 营销诉求 4：产品种草 9：客资收集 16：应用换端
	MarketTarget int `json:"market_target,omitempty"`
	// SpuType 1：spu 2：非标品 3：品牌推广（白名单支持）
	SpuType int `json:"spu_type,omitempty"`
	// Page 请求的页码，默认为1
	Page int `json:"page,omitempty"`
	// PageSize 每页行数，默认为20，最大100
	PageSize int `json:"page_size,omitempty"`
	// BaseOnly 是否只取笔记基本信息，开启后只拉取笔记基本信息不拉取指标信息，默认为false
	BaseOnly bool `json:"base_only,omitempty"`
	// CreateStartTime 笔记创建的查询范围开始时间 (yyyy-MM-dd格式，例：2025-10-20)
	CreateStartTime string `json:"create_start_time,omitempty"`
	// CreateEndTime 笔记创建的查询范围结束时间 (yyyy-MM-dd格式，例：2025-10-21)
	CreateEndTime string `json:"create_end_time,omitempty"`
	// UpdateStartTime 笔记更新的查询范围开始时间 (我的笔记、合作码和广告素材不支持更新时间字段筛选)
	UpdateStartTime string `json:"update_start_time,omitempty"`
	// UpdateEndTime 笔记更新的查询范围结束时间 (我的笔记、合作码和广告素材不支持更新时间字段筛选)
	UpdateEndTime string `json:"update_end_time,omitempty"`
}

// Encode implement PostRequest interface
func (r ListRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// ListResponse 获取笔记列表 API Response
type ListResponse struct {
	model.BaseResponse
	Data *ListResult `json:"data,omitempty"`
}

// ListResult 获取笔记列表结果
type ListResult struct {
	// Total 总数
	Total int `json:"total,omitempty"`
	// Notes 笔记信息
	Notes []BaseNoteItemDTO `json:"notes,omitempty"`
}

// BaseNoteItemDTO 笔记信息
type BaseNoteItemDTO struct {
	// NoteID 笔记id
	NoteID string `json:"note_id,omitempty"`
	// Image 图片
	Image string `json:"image,omitempty"`
	// Desc 笔记内容
	Desc string `json:"desc,omitempty"`
	// CreateTime 创建时间
	CreateTime int64 `json:"create_time,omitempty"`
	// Author 笔记作者
	Author string `json:"author,omitempty"`
	// AuthorImage 作者头像
	AuthorImage string `json:"author_image,omitempty"`
	// Status 笔记状态 0：正常 1：不匹配 2：非法 3：已删除 4：反作弊处罚等级一 5：反作弊处罚等级二 6：反作弊处罚等级三 8：私信笔记 9：淘宝笔记 10：设置隐私 11：抽奖笔记 15：反作弊处罚等级 20：笔记需要强绑spu 26：反作弊处罚等级 37：高展拒绝（0，4，5，6表示笔记可选）
	Status int `json:"status,omitempty"`
	// NoteContentType 笔记类型 1:图文笔记 2：视频笔记
	NoteContentType int `json:"note_content_type,omitempty"`
	// CooperateState 是否合作笔记
	CooperateState bool `json:"cooperate_state,omitempty"`
	// CooperateStateAuthorization 合作笔记是否启用智能标题 0-不启用、1-启用（base_only为true时不返回）
	CooperateStateAuthorization int `json:"cooperate_state_authorization,omitempty"`
	// IntelligentCreativeLink 合作笔记是否为笔记授权老链路 0-老链路、1-新链路（base_only为true时不返回）
	IntelligentCreativeLink int `json:"intelligent_creative_link,omitempty"`
	// Title 标题
	Title string `json:"title,omitempty"`
	// ImageList 图片，仅返回笔记封面图片
	ImageList []string `json:"image_list,omitempty"`
	// CooperateComponentType 合作组件类型 1：评论区商品组件 2：评论区店铺组件 3：评论区搜索组件 4：评论区私信组件 6：评论区POI组件（笔记类型为合作笔记且笔记有评论区组件时有值，0为默认值无意义，base_only为true时不返回）
	CooperateComponentType int `json:"cooperate_component_type,omitempty"`
	// CrowdCreationNote 是否为共创笔记 0：否 1：是（base_only为true时不返回）
	CrowdCreationNote bool `json:"crowd_creation_note,omitempty"`
	// ItemID 商品id（base_only为true时不返回）
	ItemID string `json:"item_id,omitempty"`
	// ReadCount 阅读数
	ReadCount int `json:"read_count,omitempty"`
	// ReadRate 阅读率
	ReadRate string `json:"read_rate,omitempty"`
	// InteractCount 互动数
	InteractCount int `json:"interact_count,omitempty"`
	// InteractRate 互动率
	InteractRate string `json:"interact_rate,omitempty"`
	// HighQuality 优质笔记(品合) 0：否 1：是（base_only为true时不返回）
	HighQuality int `json:"high_quality,omitempty"`
	// HighPotential 高潜笔记(品合) 0：否 1：是（base_only为true时不返回）
	HighPotential int `json:"high_potential,omitempty"`
	// ItemIDs 笔记挂接的商品，对于商卡笔记为社区绑定的商品集合，否则为聚光绑定的商品id集合（base_only为true时不返回）
	ItemIDs []string `json:"item_ids,omitempty"`
	// NoteSpuInfo 笔记绑定spu信息，为null表示当前笔记未绑定spu（base_only为true时不返回，废弃）
	NoteSpuInfo interface{} `json:"note_spu_info,omitempty"`
	// NoteMultiSpuInfo 笔记绑定spu信息，笔记支持绑定多spu（base_only为true时不返回，替换原来的note_spu_info）
	NoteMultiSpuInfo []NoteMultiSpuInfo `json:"note_multi_spu_info,omitempty"`
	// Taxonomy1 一级类目信息（当笔记没有一级类目信息时，展示"类目未知"）
	Taxonomy1 string `json:"taxonomy1,omitempty"`
	// Taxonomy2 二级类目信息（当笔记没有二级类目信息时，展示"类目未知"）
	Taxonomy2 string `json:"taxonomy2,omitempty"`
	// Taxonomy3 三级类目信息（当笔记没有三级类目信息时，展示"类目未知"）
	Taxonomy3 string `json:"taxonomy3,omitempty"`
	// IsHitStrategy 是否命中策略 0:未命中，1：命中
	IsHitStrategy int `json:"is_hit_strategy,omitempty"`
	// HitStrategyContent 命中策略内容，多个策略之间用逗号分隔
	HitStrategyContent string `json:"hit_strategy_content,omitempty"`
	// WinHorseNote 是否为优胜笔记
	WinHorseNote bool `json:"win_horse_note,omitempty"`
	// StaffTag 员工标签，仅员工笔记有值
	StaffTag string `json:"staff_tag,omitempty"`
	// StaffArea 地域，仅员工笔记有值
	StaffArea string `json:"staff_area,omitempty"`
	// NoteURL 笔记链接，有效期两个月（base_only为true时不返回）
	NoteURL string `json:"note_url,omitempty"`
	// HasShopCard 是否为商品笔记（base_only为true时不返回）
	HasShopCard bool `json:"has_shop_card,omitempty"`
}

// NoteMultiSpuInfo 笔记绑定spu信息
type NoteMultiSpuInfo struct {
	// BindID 绑定id
	BindID uint64 `json:"bind_id,omitempty"`
	// SpuID spu_id
	SpuID string `json:"spu_id,omitempty"`
	// SpuName spu名称
	SpuName string `json:"spu_name,omitempty"`
	// ExceedModifyLimitIn30Day 是否超出30内修改绑定三次限制 true：是 false：否
	ExceedModifyLimitIn30Day bool `json:"exceed_modify_limit_in_30_day,omitempty"`
	// ExceedModifyLimitToday 是否超出一天内修改绑定一次限制 true：是 false：否
	ExceedModifyLimitToday bool `json:"exceed_modify_limit_today,omitempty"`
	// BindByCurAccount 是否为当前账户绑定
	BindByCurAccount bool `json:"bind_by_cur_account,omitempty"`
	// BindAuditStatus 绑定关系审核状态 0：未绑定 1：申请中 2：通过 3：拒绝 4：等待审核
	BindAuditStatus int `json:"bind_audit_status,omitempty"`
	// BindAuditReason 绑定关系审核拒绝原因
	BindAuditReason string `json:"bind_audit_reason,omitempty"`
	// SpuType spu类型 1：标品 2：非标品 3：品牌 4:系列
	SpuType int `json:"spu_type,omitempty"`
	// SpuSubName 非标、品牌名称
	SpuSubName string `json:"spu_sub_name,omitempty"`
	// SeriesID 系列ID
	SeriesID string `json:"series_id,omitempty"`
}
