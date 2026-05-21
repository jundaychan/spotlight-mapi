package targettemplate

import (
	"context"

	"github.com/jundaychan/spotlight-mapi/core"
	"github.com/jundaychan/spotlight-mapi/model/targettemplate"
)

// Update 更新定向包
func Update(ctx context.Context, clt *core.SDKClient, req *targettemplate.UpdateRequest, accessToken string) (uint64, error) {
	var resp targettemplate.UpdateResponse
	if err := clt.Post(ctx, "/jg/target/template/update", req, &resp, accessToken); err != nil {
		return 0, err
	}
	if resp.Data == nil {
		return 0, nil
	}
	return resp.Data.TemplateID, nil
}
