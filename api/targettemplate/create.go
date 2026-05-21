package targettemplate

import (
	"context"

	"github.com/jundaychan/spotlight-mapi/core"
	"github.com/jundaychan/spotlight-mapi/model/targettemplate"
)

// Create 创建定向包
func Create(ctx context.Context, clt *core.SDKClient, req *targettemplate.CreateRequest, accessToken string) (uint64, error) {
	var resp targettemplate.CreateResponse
	if err := clt.Post(ctx, "/jg/target/template/create", req, &resp, accessToken); err != nil {
		return 0, err
	}
	if resp.Data == nil {
		return 0, nil
	}
	return resp.Data.TargetTemplateID, nil
}
