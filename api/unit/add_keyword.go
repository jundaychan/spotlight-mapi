package unit

import (
	"context"

	"github.com/jundaychan/spotlight-mapi/core"
	"github.com/jundaychan/spotlight-mapi/model/unit"
)

// AddKeyword 修改单元关键词
func AddKeyword(ctx context.Context, clt *core.SDKClient, req *unit.AddKeywordRequest, accessToken string) (int, error) {
	var resp unit.AddKeywordResponse
	if err := clt.Post(ctx, "/jg/unit/keyword/add", req, &resp, accessToken); err != nil {
		return 0, err
	}
	if resp.Data == nil {
		return 0, nil
	}
	return resp.Data.UploadKeywordNum, nil
}
