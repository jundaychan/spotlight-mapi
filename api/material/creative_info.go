package material

import (
	"context"

	"github.com/jundaychan/spotlight-mapi/core"
	"github.com/jundaychan/spotlight-mapi/model/material"
)

// CreativeInfo 创意标题和图片信息
func CreativeInfo(ctx context.Context, clt *core.SDKClient, req *material.CreativeInfoRequest, accessToken string) (*material.CreativeInfoResult, error) {
	var resp material.CreativeInfoResponse
	if err := clt.Post(ctx, "/jg/material/prefer/info", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
