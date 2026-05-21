package material

import (
	"github.com/jundaychan/spotlight-mapi/model"
	"github.com/jundaychan/spotlight-mapi/util"
)

// NoteType 笔记类型
type NoteType int

const (
	// NoteTypeImage 图文
	NoteTypeImage NoteType = 1
	// NoteTypeVideo 视频
	NoteTypeVideo NoteType = 2
)

// UserType 笔记作者类型
type UserType int

const (
	// UserTypeBrand 品牌账号
	UserTypeBrand UserType = 1
	// UserTypeKOS KOS账号
	UserTypeKOS UserType = 2
)

// ImageData 图片信息
type ImageData struct {
	// URL 图片的链接
	URL string `json:"url,omitempty"`
	// Height 高度
	Height int `json:"height,omitempty"`
	// Width 宽度
	Width int `json:"width,omitempty"`
}

// CoverData 视频封面信息
type CoverData struct {
	// URL 视频封面的链接
	URL string `json:"url,omitempty"`
	// Height 高度
	Height int `json:"height,omitempty"`
	// Width 宽度
	Width int `json:"width,omitempty"`
}

// VideoData 视频信息
type VideoData struct {
	// URL 视频的链接
	URL string `json:"url,omitempty"`
	// Height 高度
	Height int `json:"height,omitempty"`
	// Width 宽度
	Width int `json:"width,omitempty"`
}

// CommonData 笔记的基本信息
type CommonData struct {
	// Title 笔记的标题
	Title string `json:"title,omitempty"`
	// Desc 笔记的正文
	Desc string `json:"desc,omitempty"`
}

// PublishRequest 发布广告素材 API Request
type PublishRequest struct {
	// Type 笔记类型：1图文，2视频
	Type NoteType `json:"type,omitempty"`
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiserId,omitempty"`
	// ImageList 图文笔记的图片列表（图文笔记必传）
	ImageList []ImageData `json:"imageList,omitempty"`
	// CoverImage 视频笔记的封面信息（视频笔记必传）
	CoverImage *CoverData `json:"coverImage,omitempty"`
	// Video 视频笔记的视频信息（视频笔记必传）
	Video *VideoData `json:"video,omitempty"`
	// Content 笔记的基本信息
	Content *CommonData `json:"content,omitempty"`
	// TagList 笔记的话题列表
	TagList []string `json:"tagList,omitempty"`
	// UserID 笔记作者UID(注意不是小红书号), 不传默认是品牌账号
	UserID string `json:"userId,omitempty"`
	// UserType 笔记作者类型, 1:品牌账号 2:KOS账号, 默认是品牌账号
	UserType UserType `json:"userType,omitempty"`
}

// Encode implement PostRequest interface
func (r PublishRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// PublishResponse 发布广告素材 API Response
type PublishResponse struct {
	model.BaseResponse
	// Data 临时笔记信息
	Data *PublishResult `json:"data,omitempty"`
}

// PublishResult 临时笔记信息
type PublishResult struct {
	// TemporaryInfo 临时笔记信息
	TemporaryInfo *TemporaryInfo `json:"temporaryInfo,omitempty"`
}

// TemporaryInfo 临时笔记信息
type TemporaryInfo struct {
	// TemporaryID 临时的笔记ID
	TemporaryID string `json:"temporaryId,omitempty"`
}
