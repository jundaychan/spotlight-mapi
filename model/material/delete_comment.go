package material

import (
	"github.com/bububa/spotlight-mapi/util"
)

// DeleteCommentRequest 删除广告素材评论 API Request
type DeleteCommentRequest struct {
	// CommentID 评论ID
	CommentID string `json:"commentId,omitempty"`
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiserId,omitempty"`
}

// Encode implement PostRequest interface
func (r DeleteCommentRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
