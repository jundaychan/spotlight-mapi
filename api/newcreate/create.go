package newcreate

import (
	"context"

	"github.com/jundaychan/spotlight-mapi/core"
	"github.com/jundaychan/spotlight-mapi/model/newcreate"
)

// Create 新创编（级联创建计划/单元/创意）
func Create(ctx context.Context, clt *core.SDKClient, req *newcreate.CreateRequest, accessToken string) ([]newcreate.CascadeInfoResult, error) {
	var resp newcreate.CreateResponse
	if err := clt.Post(ctx, "/jg/cascade/create", req, &resp, accessToken); err != nil {
		return nil, err
	}
	if resp.Data == nil {
		return nil, nil
	}
	return resp.Data.InfoList, nil
}
