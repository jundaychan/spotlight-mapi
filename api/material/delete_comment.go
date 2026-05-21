package material

import (
	"context"

	"github.com/jundaychan/spotlight-mapi/core"
	"github.com/jundaychan/spotlight-mapi/model/material"
)

// DeleteComment 删除广告素材评论
func DeleteComment(ctx context.Context, clt *core.SDKClient, req *material.DeleteCommentRequest, accessToken string) error {
	return clt.Post(ctx, "/ad_material_note/comment/del", req, nil, accessToken)
}
