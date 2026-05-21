package asset

import (
	"context"

	"github.com/bububa/spotlight-mapi/core"
	"github.com/bububa/spotlight-mapi/model/asset"
)

// MobileAppList 获取移动应用列表
func MobileAppList(ctx context.Context, clt *core.SDKClient, req *asset.MobileAppListRequest, accessToken string) (*asset.MobileAppListResult, error) {
	var resp asset.MobileAppListResponse
	if err := clt.Post(ctx, "/jg/data/app/info", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
