package note

import (
	"context"

	"github.com/bububa/spotlight-mapi/core"
	"github.com/bububa/spotlight-mapi/model/note"
)

// Delete 删除笔记
func Delete(ctx context.Context, clt *core.SDKClient, req *note.DeleteRequest, accessToken string) error {
	return clt.Post(ctx, "/jg/note/delete", req, nil, accessToken)
}
