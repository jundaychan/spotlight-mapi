package material

import (
	"github.com/jundaychan/spotlight-mapi/model"
	"github.com/jundaychan/spotlight-mapi/util"
)

// CreativeInfoMaterial 查询物料信息
type CreativeInfoMaterial struct {
	// NoteID 笔记id（新建和编辑时候必传）
	NoteID string `json:"note_id,omitempty"`
	// CreativityID 创意id（在编辑场景下必传）
	CreativityID int64 `json:"creativity_id,omitempty"`
}

// CreativeInfoRequest 创意标题和图片信息 API Request
type CreativeInfoRequest struct {
	// MaterialInfoDtos 查询物料信息
	MaterialInfoDtos []CreativeInfoMaterial `json:"material_info_dtos,omitempty"`
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
}

// Encode implement PostRequest interface
func (r CreativeInfoRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// CreativeInfoResponse 创意标题和图片信息 API Response
type CreativeInfoResponse struct {
	model.BaseResponse
	// Data 封面&标题
	Data *CreativeInfoResult `json:"data,omitempty"`
}

// CreativeInfoResult 封面&标题
type CreativeInfoResult struct {
	// MaterialInfoList 优选候选集
	MaterialInfoList []MaterialInfo `json:"material_info_list,omitempty"`
}

// MaterialInfo 物料信息
type MaterialInfo struct {
	// TitleInfoList 标题列表
	TitleInfoList []TitleInfo `json:"title_info_list,omitempty"`
	// PhotoInfoList 封面列表
	PhotoInfoList []PhotoInfo `json:"photo_info_list,omitempty"`
	// NoteID 笔记 id
	NoteID string `json:"note_id,omitempty"`
	// CreativityID 创意 id
	CreativityID int64 `json:"creativity_id,omitempty"`
}

// TitleInfo 标题信息
type TitleInfo struct {
	// Title 标题
	Title string `json:"title,omitempty"`
	// Origin 是否原素材
	Origin bool `json:"origin,omitempty"`
	// Selected 是否选择
	Selected bool `json:"selected,omitempty"`
	// AuditStatus 审核状态
	AuditStatus int32 `json:"audit_status,omitempty"`
	// AuditReason 审核原因
	AuditReason string `json:"audit_reason,omitempty"`
	// TitleSource 标题来源
	TitleSource int32 `json:"title_source,omitempty"`
	// TitleTab 标题tab
	TitleTab int32 `json:"title_tab,omitempty"`
	// Version 版本
	Version string `json:"version,omitempty"`
}

// PhotoInfo 封面信息
type PhotoInfo struct {
	// FileID 文件id
	FileID string `json:"file_id,omitempty"`
	// FileURL 图片url
	FileURL string `json:"file_url,omitempty"`
	// Origin 是否原始素材
	Origin bool `json:"origin,omitempty"`
	// Selected 是否选择
	Selected bool `json:"selected,omitempty"`
	// AuditStatus 审核状态
	AuditStatus int32 `json:"audit_status,omitempty"`
	// AuditReason 审核原因
	AuditReason string `json:"audit_reason,omitempty"`
	// ImageSource 图片来源
	ImageSource int32 `json:"image_source,omitempty"`
}
