package series

import (
	"context"

	"github.com/jundaychan/spotlight-mapi/core"
	"github.com/jundaychan/spotlight-mapi/model/series"
)

// List 可投剧集列表
func List(ctx context.Context, clt *core.SDKClient, req *series.ListRequest, accessToken string) (*series.ListResult, error) {
	var resp series.ListResponse
	if err := clt.Post(ctx, "/jg/data/episodes/list", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
