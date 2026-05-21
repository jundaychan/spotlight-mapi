package data

import (
	"context"

	"github.com/jundaychan/spotlight-mapi/core"
	"github.com/jundaychan/spotlight-mapi/model/data"
)

// EventAssetInfo 资产事件获取
func EventAssetInfo(ctx context.Context, clt *core.SDKClient, req *data.EventAssetInfoRequest, accessToken string) (*data.EventAssetInfoResult, error) {
	var resp data.EventAssetInfoResponse
	if err := clt.Post(ctx, "/jg/data/event/asset/info", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
