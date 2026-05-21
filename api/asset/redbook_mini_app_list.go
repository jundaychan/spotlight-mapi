package asset

import (
	"context"

	"github.com/bububa/spotlight-mapi/core"
	"github.com/bububa/spotlight-mapi/model/asset"
)

// RedbookMiniAppList 获取红书小程序列表
func RedbookMiniAppList(ctx context.Context, clt *core.SDKClient, req *asset.RedbookMiniAppListRequest, accessToken string) (*asset.RedbookMiniAppListResult, error) {
	var resp asset.RedbookMiniAppListResponse
	if err := clt.Post(ctx, "/jg/query/mini_program", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
