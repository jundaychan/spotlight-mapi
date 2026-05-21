package campaigngroup

import (
	"context"

	"github.com/jundaychan/spotlight-mapi/core"
	"github.com/jundaychan/spotlight-mapi/model/campaigngroup"
)

// Create 创建广告组
func Create(ctx context.Context, clt *core.SDKClient, req *campaigngroup.CreateRequest, accessToken string) (uint64, error) {
	var resp campaigngroup.CreateResponse
	if err := clt.Post(ctx, "/jg/campaign/group/create", req, &resp, accessToken); err != nil {
		return 0, err
	}
	return resp.Data.ID, nil
}
