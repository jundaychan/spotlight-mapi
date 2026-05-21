package history

import (
	"context"

	"github.com/jundaychan/spotlight-mapi/core"
	"github.com/jundaychan/spotlight-mapi/model/history"
)

// List 历史操作记录
func List(ctx context.Context, clt *core.SDKClient, req *history.ListRequest, accessToken string) (*history.ListResult, error) {
	var resp history.ListResponse
	if err := clt.Post(ctx, "/jg/history/list", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
