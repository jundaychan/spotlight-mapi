package series

import (
	"context"

	"github.com/jundaychan/spotlight-mapi/core"
	"github.com/jundaychan/spotlight-mapi/model/series"
)

// Create 剧集创建
func Create(ctx context.Context, clt *core.SDKClient, req *series.CreateRequest, accessToken string) (*series.CreateResult, error) {
	var resp series.CreateResponse
	if err := clt.Post(ctx, "/jg/data/episodes/create", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
