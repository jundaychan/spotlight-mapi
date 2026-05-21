package unit

import (
	"context"

	"github.com/bububa/spotlight-mapi/core"
	"github.com/bububa/spotlight-mapi/model/unit"
)

// UpdateBid 修改单元出价
func UpdateBid(ctx context.Context, clt *core.SDKClient, req *unit.UpdateBidRequest, accessToken string) error {
	return clt.Post(ctx, "/jg/unit/batch/update/bid", req, nil, accessToken)
}
