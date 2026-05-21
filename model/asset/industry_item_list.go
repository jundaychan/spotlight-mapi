package asset

import (
	"github.com/jundaychan/spotlight-mapi/model"
	"github.com/jundaychan/spotlight-mapi/util"
)

// IndustryItemListRequest 获取行业商品列表 API Request
type IndustryItemListRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// PageIndex 页码
	PageIndex int32 `json:"page_index,omitempty"`
	// PageSize 页大小建议第一次不超过20
	PageSize int32 `json:"page_size,omitempty"`
	// IndustryItemIDs 行业商品主键id
	IndustryItemIDs []string `json:"industry_item_ids,omitempty"`
	// Platforms 行业商品平台：1淘宝；2京东；3拼多多；6：淘宝ud；7京东ud。默认不传，会返回账号下所有商品
	Platforms []string `json:"platforms,omitempty"`
}

// Encode implement PostRequest interface
func (r IndustryItemListRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// IndustryItemListResponse 获取行业商品列表 API Response
type IndustryItemListResponse struct {
	model.BaseResponse
	// Data 行业商品明细信息
	Data *IndustryItemListResult `json:"data,omitempty"`
}

// IndustryItemListResult 行业商品明细信息
type IndustryItemListResult struct {
	// PageIndex 页码
	PageIndex int `json:"page_index,omitempty"`
	// TotalCount 总数
	TotalCount int `json:"total_count,omitempty"`
	// IndustryItemDtos 商品明细信息列表
	IndustryItemDtos []IndustryItem `json:"industry_item_dtos,omitempty"`
}

// IndustryItem 行业商品
type IndustryItem struct {
	// IndustryItemID 商品id：主键id，投放的时候填这个字段
	IndustryItemID string `json:"industry_item_id,omitempty"`
	// IndustryItemName 商品名称
	IndustryItemName string `json:"industry_item_name,omitempty"`
	// ImageURL 商品主图
	ImageURL string `json:"image_url,omitempty"`
	// Platform 平台：1淘宝；2京东；3拼多多；6：淘宝ud；7京东ud
	Platform int `json:"platform,omitempty"`
	// OuterItemID 站外商品id
	OuterItemID string `json:"outer_item_id,omitempty"`
	// JumpURLType 跳转链接:1直达链接 2一跳落地页 3二跳落地页
	JumpURLType int `json:"jump_url_type,omitempty"`
	// JumpURL jump_url_type是1时表示deeplink 一跳落地页的话是跳转链接
	JumpURL string `json:"jump_url,omitempty"`
	// H5Link 兜底h5链接
	H5Link string `json:"h5_link,omitempty"`
	// UniversalLink universal_link
	UniversalLink string `json:"universal_link,omitempty"`
}
