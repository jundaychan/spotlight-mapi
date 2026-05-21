package asset

import (
	"github.com/jundaychan/spotlight-mapi/model"
	"github.com/jundaychan/spotlight-mapi/util"
)

// MobileAppPlatform 平台类型
type MobileAppPlatform int

const (
	// MobileAppPlatformIOS iOS
	MobileAppPlatformIOS MobileAppPlatform = 1
	// MobileAppPlatformAndroid android(包含唤起商店和直接下载)
	MobileAppPlatformAndroid MobileAppPlatform = 2
)

// MobileAppListRequest 获取移动应用列表 API Request
type MobileAppListRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// PageIndex 页码
	PageIndex int32 `json:"page_index,omitempty"`
	// PageSize 页大小建议第一次不超过20
	PageSize int32 `json:"page_size,omitempty"`
	// Platform 平台类型1:ios, 2:androd(包含唤起商店和直接下载)
	Platform MobileAppPlatform `json:"platform,omitempty"`
	// AppID appid:ios是官方id android是包名
	AppID string `json:"app_id,omitempty"`
	// AppName 应用包名
	AppName string `json:"app_name,omitempty"`
}

// Encode implement PostRequest interface
func (r MobileAppListRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// MobileAppListResponse 获取移动应用列表 API Response
type MobileAppListResponse struct {
	model.BaseResponse
	// Data 应用明细信息
	Data *MobileAppListResult `json:"data,omitempty"`
}

// MobileAppListResult 应用明细信息
type MobileAppListResult struct {
	// Page 分页信息
	Page *MobileAppPage `json:"page,omitempty"`
	// AppInfoList 应用信息列表
	AppInfoList []MobileAppInfo `json:"app_info_list,omitempty"`
}

// MobileAppPage 分页信息
type MobileAppPage struct {
	// PageIndex 页码
	PageIndex int32 `json:"page_index,omitempty"`
	// TotalCount 总数
	TotalCount int32 `json:"total_count,omitempty"`
}

// MobileAppInfo 应用信息
type MobileAppInfo struct {
	// ID 唯一id，投放时app_unique_id填此值
	ID uint64 `json:"id,omitempty"`
	// AppID 应用id，投放时app_id填此值
	AppID string `json:"app_id,omitempty"`
	// AppName 应用名称
	AppName string `json:"app_name,omitempty"`
	// PackageName 包名
	PackageName string `json:"package_name,omitempty"`
	// Version 版本号
	Version string `json:"version,omitempty"`
	// PublishTime 发布时间
	PublishTime int64 `json:"publish_time,omitempty"`
	// Desc 应用描述
	Desc string `json:"desc,omitempty"`
	// Developer 开发者
	Developer string `json:"developer,omitempty"`
	// Icon 应用图标
	Icon string `json:"icon,omitempty"`
	// Platform 应用类型：1-iOS，2-android 3-android-apk
	Platform int32 `json:"platform,omitempty"`
	// ReviewStatus 审核状态，只有1通过
	ReviewStatus int32 `json:"review_status,omitempty"`
	// PublishStatus 发布状态，只有1发布
	PublishStatus int64 `json:"publish_status,omitempty"`
	// AuditFailedReason 审核失败原因（暂时没有）
	AuditFailedReason string `json:"audit_failed_reason,omitempty"`
}
