package note

import (
	"context"

	"github.com/jundaychan/spotlight-mapi/core"
	"github.com/jundaychan/spotlight-mapi/model/note"
)

// List 获取笔记列表
func List(ctx context.Context, clt *core.SDKClient, req *note.ListRequest, accessToken string) (*note.ListResult, error) {
	var resp note.ListResponse
	if err := clt.Post(ctx, "/jg/note/list", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
