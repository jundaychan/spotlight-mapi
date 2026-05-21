package campaigngroup

import (
	"context"

	"github.com/jundaychan/spotlight-mapi/core"
	"github.com/jundaychan/spotlight-mapi/model/campaigngroup"
)

// List 查询广告组列表
func List(ctx context.Context, clt *core.SDKClient, req *campaigngroup.ListRequest, accessToken string) (*campaigngroup.ListResult, error) {
	var resp campaigngroup.ListResponse
	if err := clt.Post(ctx, "/jg/campaign/group/base/list", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
