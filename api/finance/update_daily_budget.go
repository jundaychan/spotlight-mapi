package finance

import (
	"context"

	"github.com/bububa/spotlight-mapi/core"
	"github.com/bububa/spotlight-mapi/model/finance"
)

// UpdateDailyBudget 修改账户日预算
func UpdateDailyBudget(ctx context.Context, clt *core.SDKClient, req *finance.UpdateDailyBudgetRequest, accessToken string) error {
	return clt.Post(ctx, "/jg/account/budget/update", req, nil, accessToken)
}
