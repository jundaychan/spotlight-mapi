package finance

import (
	"context"

	"github.com/bububa/spotlight-mapi/core"
	"github.com/bububa/spotlight-mapi/model/finance"
)

// DailyBudgetBalance 获取账户日预算余额接口
func DailyBudgetBalance(ctx context.Context, clt *core.SDKClient, req *finance.DailyBudgetBalanceRequest, accessToken string) (*finance.DailyBudgetBalance, error) {
	var resp finance.DailyBudgetBalanceResponse
	if err := clt.Post(ctx, "/jg/account/budget/info", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
