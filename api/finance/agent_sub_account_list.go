package finance

import (
	"context"

	"github.com/jundaychan/spotlight-mapi/core"
	"github.com/jundaychan/spotlight-mapi/model/finance"
)

// AgentSubAccountList 获取代理商子账号列表
func AgentSubAccountList(ctx context.Context, clt *core.SDKClient, req *finance.AgentSubAccountListRequest, accessToken string) (*finance.AgentSubAccountListResult, error) {
	var resp finance.AgentSubAccountListResponse
	if err := clt.Post(ctx, "/jg/account/sub/page", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
