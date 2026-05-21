package targettemplate

import (
	"context"

	"github.com/bububa/spotlight-mapi/core"
	"github.com/bububa/spotlight-mapi/model/targettemplate"
)

// Apply 定向包关联单元
func Apply(ctx context.Context, clt *core.SDKClient, req *targettemplate.ApplyRequest, accessToken string) (*targettemplate.ApplyResult, error) {
	var resp targettemplate.ApplyResponse
	if err := clt.Post(ctx, "/jg/target/template/apply", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
