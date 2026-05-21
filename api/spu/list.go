package spu

import (
	"context"

	"github.com/jundaychan/spotlight-mapi/core"
	"github.com/jundaychan/spotlight-mapi/model/spu"
)

// List 获取spu列表
func List(ctx context.Context, clt *core.SDKClient, req *spu.ListRequest, accessToken string) (*spu.ListResult, error) {
	var resp spu.ListResponse
	if err := clt.Post(ctx, "/jg/spu/list", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
