package material

import (
	"context"

	"github.com/jundaychan/spotlight-mapi/core"
	"github.com/jundaychan/spotlight-mapi/model/material"
)

// Publish 发布广告素材
func Publish(ctx context.Context, clt *core.SDKClient, req *material.PublishRequest, accessToken string) (*material.PublishResult, error) {
	var resp material.PublishResponse
	if err := clt.Post(ctx, "/ad_material_note/save", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
