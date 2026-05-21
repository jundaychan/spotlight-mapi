package data

import (
	"context"

	"github.com/jundaychan/spotlight-mapi/core"
	"github.com/jundaychan/spotlight-mapi/model/data"
)

// CheckNameDup 计划单元名称重复性校验
func CheckNameDup(ctx context.Context, clt *core.SDKClient, req *data.CheckNameDupRequest, accessToken string) (*data.CheckNameDupResult, error) {
	var resp data.CheckNameDupResponse
	if err := clt.Post(ctx, "/jg/data/check/name/dup", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
