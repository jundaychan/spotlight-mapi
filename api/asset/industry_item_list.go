package asset

import (
	"context"

	"github.com/jundaychan/spotlight-mapi/core"
	"github.com/jundaychan/spotlight-mapi/model/asset"
)

// IndustryItemList 获取行业商品列表
func IndustryItemList(ctx context.Context, clt *core.SDKClient, req *asset.IndustryItemListRequest, accessToken string) (*asset.IndustryItemListResult, error) {
	var resp asset.IndustryItemListResponse
	if err := clt.Post(ctx, "/jg/data/product/search", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
