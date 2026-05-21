package series

import (
	"context"

	"github.com/jundaychan/spotlight-mapi/core"
	"github.com/jundaychan/spotlight-mapi/model/series"
)

// Update 剧集修改
func Update(ctx context.Context, clt *core.SDKClient, req *series.UpdateRequest, accessToken string) (*series.UpdateResult, error) {
	var resp series.UpdateResponse
	if err := clt.Post(ctx, "/jg/data/episodes/update", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
