package material

import (
	"github.com/jundaychan/spotlight-mapi/model"
	"github.com/jundaychan/spotlight-mapi/util"
)

// CommentPageReqDTO 分页信息
type CommentPageReqDTO struct {
	// PageIndex 页码，默认为1
	PageIndex int `json:"pageIndex,omitempty"`
	// PageSize 每页条数，默认20，最大100
	PageSize int `json:"pageSize,omitempty"`
}

// CommentOrderClause 排序规则
type CommentOrderClause struct {
	// Column 排序字段 创建时间: create_time 更新时间: update_time
	Column string `json:"column,omitempty"`
	// Sort 升降序 正排:asc 倒排:desc
	Sort string `json:"sort,omitempty"`
}

// GetCommentsRequest 获取广告素材评论 API Request
type GetCommentsRequest struct {
	// PageReqDTO 分页信息
	PageReqDTO *CommentPageReqDTO `json:"pageReqDTO,omitempty"`
	// OrderClauseInfo 排序规则
	OrderClauseInfo []CommentOrderClause `json:"orderClauseInfo,omitempty"`
	// NoteIDs 笔记ID列表
	NoteIDs []string `json:"noteIds,omitempty"`
	// Content 评论内容关键词检索
	Content string `json:"content,omitempty"`
	// CommentLevel 评论层级 1:一级评论 2:二级评论（回复）
	CommentLevel []int `json:"commentLevel,omitempty"`
	// ReplyStatus 被回复状态 1:未被回复 2:已被回复
	ReplyStatus []int `json:"replyStatus,omitempty"`
	// CreateStartTime 笔记创建的查询范围开始时间 (yyyy-MM-dd格式)
	CreateStartTime string `json:"createStartTime,omitempty"`
	// CreateEndTime 笔记创建的查询范围结束时间 (yyyy-MM-dd格式)
	CreateEndTime string `json:"createEndTime,omitempty"`
	// IncludeSelfComment 是否包含作者评论，默认为否
	IncludeSelfComment bool `json:"includeSelfComment,omitempty"`
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiserId,omitempty"`
}

// Encode implement PostRequest interface
func (r GetCommentsRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// GetCommentsResponse 获取广告素材评论 API Response
type GetCommentsResponse struct {
	model.BaseResponse
	// Data 返回数据
	Data *GetCommentsResult `json:"data,omitempty"`
}

// CommentPageResultDTO 分页信息
type CommentPageResultDTO struct {
	// PageIndex 当前页码
	PageIndex int `json:"pageIndex,omitempty"`
	// TotalCount 评论总数 (不含作者回复)
	TotalCount int `json:"totalCount,omitempty"`
}

// GetCommentsResult 返回数据
type GetCommentsResult struct {
	// MaterialNoteCommentListDtos 评论列表
	MaterialNoteCommentListDtos []MaterialNoteComment `json:"materialNoteCommentListDtos,omitempty"`
	// CommentTotal 评论总数 (含作者回复)
	CommentTotal int64 `json:"commentTotal,omitempty"`
	// ReplyRate 回复率
	ReplyRate float64 `json:"replyRate,omitempty"`
	// PageResultDTO 分页信息
	PageResultDTO *CommentPageResultDTO `json:"pageResultDTO,omitempty"`
}

// MaterialNoteComment 评论信息
type MaterialNoteComment struct {
	// CommentID 评论ID
	CommentID string `json:"commentId,omitempty"`
	// TargetCommentID 目标评论ID（回复的评论ID，一级评论为空）
	TargetCommentID string `json:"targetCommentId,omitempty"`
	// Content 评论文字内容
	Content string `json:"content,omitempty"`
	// TagList 评论中的话题标签
	TagList []string `json:"tagList,omitempty"`
	// AtUsers 评论中@的用户列表
	AtUsers []string `json:"atUsers,omitempty"`
	// CreateTime 评论创建时间戳（毫秒）
	CreateTime int64 `json:"createTime,omitempty"`
	// CommentLevel 评论层级 1:一级评论 2:二级评论（回复）
	CommentLevel int `json:"commentLevel,omitempty"`
	// AuthorLiked 笔记作者是否点赞
	AuthorLiked bool `json:"authorLiked,omitempty"`
	// LikeInfoCount 点赞数
	LikeInfoCount int `json:"likeInfoCount,omitempty"`
	// ReplyStatus 被回复状态 1:未被回复 2:已被回复
	ReplyStatus int `json:"replyStatus,omitempty"`
	// PictureInfoDto 评论图片内容
	PictureInfoDto []CommentPictureInfo `json:"pictureInfoDto,omitempty"`
	// UserInfoDto 评论者信息
	UserInfoDto *CommentUserInfo `json:"userInfoDto,omitempty"`
	// NoteInfoDto 笔记信息
	NoteInfoDto *CommentNoteInfo `json:"noteInfoDto,omitempty"`
}

// CommentPictureInfo 评论图片内容
type CommentPictureInfo struct {
	// FileID 图片的id
	FileID string `json:"fileId,omitempty"`
	// URL 图片链接
	URL string `json:"url,omitempty"`
	// Height 高度
	Height int `json:"height,omitempty"`
	// Width 宽度
	Width int `json:"width,omitempty"`
}

// CommentUserInfo 评论者信息
type CommentUserInfo struct {
	// UserID 用户的id
	UserID string `json:"userId,omitempty"`
	// NickName 昵称
	NickName string `json:"nickName,omitempty"`
	// UserAvatar 头像链接
	UserAvatar string `json:"userAvatar,omitempty"`
}

// CommentNoteInfo 笔记信息
type CommentNoteInfo struct {
	// ID 笔记的id
	ID string `json:"id,omitempty"`
	// Title 标题
	Title string `json:"title,omitempty"`
	// Image 封面链接
	Image string `json:"image,omitempty"`
	// Type 类型
	Type int `json:"type,omitempty"`
	// CreateTime 创建时间
	CreateTime int64 `json:"createTime,omitempty"`
	// Content 内容
	Content string `json:"content,omitempty"`
}
