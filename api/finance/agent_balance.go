package finance

import (
	"context"

	"github.com/jundaychan/spotlight-mapi/core"
	"github.com/jundaychan/spotlight-mapi/model/finance"
)

// AgentBalance 代理商主子账号余额查询
func AgentBalance(ctx context.Context, clt *core.SDKClient, req *finance.AgentBalanceRequest, accessToken string) (*finance.AgentBalanceResult, error) {
	var resp finance.AgentBalanceResponse
	if err := clt.Post(ctx, "/finance/balance/query", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
