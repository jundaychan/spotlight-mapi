package simpledelivery

import (
	"context"

	"github.com/bububa/spotlight-mapi/core"
	"github.com/bububa/spotlight-mapi/model/simpledelivery"
)

// Create 新建简单投
func Create(ctx context.Context, clt *core.SDKClient, req *simpledelivery.CreateRequest, accessToken string) (*simpledelivery.UbeSemiCreateAdResponse, error) {
	var resp simpledelivery.CreateResponse
	if err := clt.Post(ctx, "/jg/ube/semi/auto/create", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
