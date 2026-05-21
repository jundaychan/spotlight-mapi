package targettemplate

import (
	"context"

	"github.com/bububa/spotlight-mapi/core"
	"github.com/bububa/spotlight-mapi/model/targettemplate"
)

// Delete 删除定向包
func Delete(ctx context.Context, clt *core.SDKClient, req *targettemplate.DeleteRequest, accessToken string) error {
	return clt.Post(ctx, "/jg/target/template/delete", req, nil, accessToken)
}
