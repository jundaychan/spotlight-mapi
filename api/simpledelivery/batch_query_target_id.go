package simpledelivery

import (
	"context"

	"github.com/bububa/spotlight-mapi/core"
	"github.com/bububa/spotlight-mapi/model/simpledelivery"
)

// BatchQueryTargetID 批量查询简单投标的ID
func BatchQueryTargetID(ctx context.Context, clt *core.SDKClient, req *simpledelivery.BatchQueryTargetIDRequest, accessToken string) ([]simpledelivery.ExtraDTO, error) {
	var resp simpledelivery.BatchQueryTargetIDResponse
	if err := clt.Post(ctx, "/jg/ube/extra/query", req, &resp, accessToken); err != nil {
		return nil, err
	}
	if resp.Data == nil {
		return nil, nil
	}
	return resp.Data.ExtraDTO, nil
}
