package material

import (
	"context"

	"github.com/bububa/spotlight-mapi/core"
	"github.com/bububa/spotlight-mapi/model/material"
)

// GetComments 获取广告素材评论
func GetComments(ctx context.Context, clt *core.SDKClient, req *material.GetCommentsRequest, accessToken string) (*material.GetCommentsResult, error) {
	var resp material.GetCommentsResponse
	if err := clt.Post(ctx, "/ad_material_note/comment/getlist", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
