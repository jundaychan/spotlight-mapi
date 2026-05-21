package asset

import (
	"context"

	"github.com/jundaychan/spotlight-mapi/core"
	"github.com/jundaychan/spotlight-mapi/model/asset"
)

// WechatMiniAppList 获取微信小程序/小游戏列表
func WechatMiniAppList(ctx context.Context, clt *core.SDKClient, req *asset.WechatMiniAppListRequest, accessToken string) (*asset.WechatMiniAppListResult, error) {
	var resp asset.WechatMiniAppListResponse
	if err := clt.Post(ctx, "/jg/app/industry_item/list", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
