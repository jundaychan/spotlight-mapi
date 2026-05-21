package note

import (
	"context"

	"github.com/bububa/spotlight-mapi/core"
	"github.com/bububa/spotlight-mapi/model/note"
)

// QueryID 获取笔记ID
func QueryID(ctx context.Context, clt *core.SDKClient, req *note.QueryIDRequest, accessToken string) (*note.QueryIDResult, error) {
	var resp note.QueryIDResponse
	if err := clt.Post(ctx, "/jg/noteid/query", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
