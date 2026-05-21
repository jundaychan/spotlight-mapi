package targettemplate

import (
	"context"

	"github.com/bububa/spotlight-mapi/core"
	"github.com/bububa/spotlight-mapi/model/targettemplate"
)

// List 获取定向包列表
func List(ctx context.Context, clt *core.SDKClient, req *targettemplate.ListRequest, accessToken string) (*targettemplate.ListResult, error) {
	var resp targettemplate.ListResponse
	if err := clt.Post(ctx, "/jg/target/template/query", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
