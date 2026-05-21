package campaigngroup

import (
	"context"

	"github.com/bububa/spotlight-mapi/core"
	"github.com/bububa/spotlight-mapi/model/campaigngroup"
)

// Delete 删除广告组
func Delete(ctx context.Context, clt *core.SDKClient, req *campaigngroup.DeleteRequest, accessToken string) error {
	return clt.Post(ctx, "/jg/campaign/group/delete", req, nil, accessToken)
}
