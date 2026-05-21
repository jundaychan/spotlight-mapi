package crowd

import (
	"context"

	"github.com/jundaychan/spotlight-mapi/core"
	"github.com/jundaychan/spotlight-mapi/model/crowd"
)

// Estimate 人群预估
func Estimate(ctx context.Context, clt *core.SDKClient, req *crowd.EstimateRequest, accessToken string) (*crowd.EstimateResult, error) {
	var resp crowd.EstimateResponse
	if err := clt.Post(ctx, "/jg/crowd/estimate", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
