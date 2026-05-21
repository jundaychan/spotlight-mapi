package campaign

import (
	"context"

	"github.com/jundaychan/spotlight-mapi/core"
	"github.com/jundaychan/spotlight-mapi/model/campaign"
)

// List 查询计划
func List(ctx context.Context, clt *core.SDKClient, req *campaign.ListRequest, accessToken string) (*campaign.ListResult, error) {
	var resp campaign.ListResponse
	if err := clt.Post(ctx, "/jg/campaign/list", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
