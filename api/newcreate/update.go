package newcreate

import (
	"context"

	"github.com/bububa/spotlight-mapi/core"
	"github.com/bububa/spotlight-mapi/model/newcreate"
)

// Update 新创编编辑（目前只支持产品种草营销诉求）
func Update(ctx context.Context, clt *core.SDKClient, req *newcreate.UpdateRequest, accessToken string) error {
	return clt.Post(ctx, "/jg/cascade/modify", req, nil, accessToken)
}
