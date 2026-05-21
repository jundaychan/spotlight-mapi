package finance

import (
	"context"

	"github.com/jundaychan/spotlight-mapi/core"
	"github.com/jundaychan/spotlight-mapi/model/finance"
)

// FundFlow 资金流水查询
func FundFlow(ctx context.Context, clt *core.SDKClient, req *finance.FundFlowRequest, accessToken string) (*finance.FundFlowResult, error) {
	var resp finance.FundFlowResponse
	if err := clt.Post(ctx, "/finance/transaction/record/query", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
