package finance

import (
	"context"

	"github.com/jundaychan/spotlight-mapi/core"
	"github.com/jundaychan/spotlight-mapi/model/finance"
)

// AccountFlow 获取账户流水接口
func AccountFlow(ctx context.Context, clt *core.SDKClient, req *finance.AccountFlowRequest, accessToken string) (*finance.AccountFlowResult, error) {
	var resp finance.AccountFlowResponse
	if err := clt.Post(ctx, "/jg/account/order/info", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
