package asset

import (
	"github.com/jundaychan/spotlight-mapi/model"
	"github.com/jundaychan/spotlight-mapi/util"
)

// WechatMiniAppListRequest 获取微信小程序/小游戏列表 API Request
type WechatMiniAppListRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// PageIndex 页码
	PageIndex int32 `json:"page_index,omitempty"`
	// PageSize 页大小建议第一次不超过20
	PageSize int32 `json:"page_size,omitempty"`
	// ItemIDList 搜索 itemId（小程序的itemid）
	ItemIDList []string `json:"item_id_list,omitempty"`
	// Keyword 搜索关键词小程序名称
	Keyword string `json:"keyword,omitempty"`
}

// Encode implement PostRequest interface
func (r WechatMiniAppListRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// WechatMiniAppListResponse 获取微信小程序/小游戏列表 API Response
type WechatMiniAppListResponse struct {
	model.BaseResponse
	// Data 小程序明细信息
	Data *WechatMiniAppListResult `json:"data,omitempty"`
}

// WechatMiniAppListResult 小程序明细信息
type WechatMiniAppListResult struct {
	// PageIndex 页码
	PageIndex int `json:"page_index,omitempty"`
	// TotalCount 总数
	TotalCount int `json:"total_count,omitempty"`
	// IndustryItemDtos 小程序明细信息列表
	IndustryItemDtos []WechatMiniAppItem `json:"industry_item_dtos,omitempty"`
}

// WechatMiniAppItem 小程序明细信息
type WechatMiniAppItem struct {
	// ItemID 小程序itemid
	ItemID string `json:"item_id,omitempty"`
	// MiniProgramID 小程序id
	MiniProgramID string `json:"mini_program_id,omitempty"`
	// MiniProgramPath 小程序路径
	MiniProgramPath string `json:"mini_program_path,omitempty"`
	// MinProgramVersion 小程序版本
	MinProgramVersion string `json:"min_program_version,omitempty"`
}
