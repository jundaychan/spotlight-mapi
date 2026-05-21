package negativekeyword

import (
	"context"

	"github.com/jundaychan/spotlight-mapi/core"
	"github.com/jundaychan/spotlight-mapi/model/negativekeyword"
)

// List 查询否定词列表
func List(ctx context.Context, clt *core.SDKClient, req *negativekeyword.ListRequest, accessToken string) (*negativekeyword.ListResult, error) {
	var resp negativekeyword.ListResponse
	if err := clt.Post(ctx, "/jg/negative/keyword/list", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
