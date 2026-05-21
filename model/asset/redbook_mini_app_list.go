package asset

import (
	"github.com/bububa/spotlight-mapi/model"
	"github.com/bububa/spotlight-mapi/util"
)

// PlayletType 小程序类型
type PlayletType int

const (
	// PlayletTypeRed 小红书小程序
	PlayletTypeRed PlayletType = 1
	// PlayletTypeWechat 微信小程序
	PlayletTypeWechat PlayletType = 2
)

// RedbookMiniAppSceneType 场景类型
type RedbookMiniAppSceneType int

const (
	// RedbookMiniAppSceneTypeMine 我的小程序
	RedbookMiniAppSceneTypeMine RedbookMiniAppSceneType = 0
	// RedbookMiniAppSceneTypeAuthorized 授权小程序
	RedbookMiniAppSceneTypeAuthorized RedbookMiniAppSceneType = 2
)

// RedbookMiniAppListRequest 获取红书小程序列表 API Request
type RedbookMiniAppListRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// PageNum 页码
	PageNum int32 `json:"page_num,omitempty"`
	// PageSize 页大小
	PageSize int32 `json:"page_size,omitempty"`
	// PlayletType 小程序类型：1-RED_MINI_PROGRAM（小红书小程序）2-WECHAT_MINI_PROGRAM（微信小程序）
	PlayletType PlayletType `json:"playlet_type,omitempty"`
	// SceneType 场景类型（2-授权小程序，0-我的小程序）默认 0
	SceneType RedbookMiniAppSceneType `json:"scene_type,omitempty"`
}

// Encode implement PostRequest interface
func (r RedbookMiniAppListRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// RedbookMiniAppListResponse 获取红书小程序列表 API Response
type RedbookMiniAppListResponse struct {
	model.BaseResponse
	// Data 小程序明细信息
	Data *RedbookMiniAppListResult `json:"data,omitempty"`
}

// RedbookMiniAppListResult 小程序明细信息
type RedbookMiniAppListResult struct {
	// Total 总数
	Total int32 `json:"total,omitempty"`
	// MiniProgramInfoList 小程序明细信息列表
	MiniProgramInfoList []RedbookMiniAppInfo `json:"mini_program_info_list,omitempty"`
}

// RedbookMiniAppInfo 小程序明细信息
type RedbookMiniAppInfo struct {
	// BrandUserID 所属人userId
	BrandUserID string `json:"brand_user_id,omitempty"`
	// PlayletAppID 小程序id
	PlayletAppID string `json:"playlet_app_id,omitempty"`
	// BrandUserName 所属人名称
	BrandUserName string `json:"brand_user_name,omitempty"`
	// ImageURL 所属人头像
	ImageURL string `json:"image_url,omitempty"`
}
