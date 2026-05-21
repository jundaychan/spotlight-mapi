package directlink

import (
	"context"

	"github.com/jundaychan/spotlight-mapi/core"
	"github.com/jundaychan/spotlight-mapi/model/directlink"
)

// Update 编辑直达链接
func Update(ctx context.Context, clt *core.SDKClient, req *directlink.UpdateRequest, accessToken string) error {
	return clt.Post(ctx, "/jg/direct-link/update", req, nil, accessToken)
}
