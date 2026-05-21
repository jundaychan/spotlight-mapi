package campaigngroup

import (
	"context"

	"github.com/jundaychan/spotlight-mapi/core"
	"github.com/jundaychan/spotlight-mapi/model/campaigngroup"
)

// Update 更新广告组
func Update(ctx context.Context, clt *core.SDKClient, req *campaigngroup.UpdateRequest, accessToken string) error {
	return clt.Post(ctx, "/jg/campaign/group/update", req, nil, accessToken)
}
