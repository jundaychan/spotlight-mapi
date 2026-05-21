package campaign

import (
	"context"

	"github.com/jundaychan/spotlight-mapi/core"
	"github.com/jundaychan/spotlight-mapi/model/campaign"
)

// Update 编辑计划
func Update(ctx context.Context, clt *core.SDKClient, req *campaign.UpdateRequest, accessToken string) (uint64, error) {
	var resp campaign.UpdateResponse
	if err := clt.Post(ctx, "/jg/campaign/update", req, &resp, accessToken); err != nil {
		return 0, err
	}
	return resp.Data.CampaignID, nil
}
