package finance

import (
	"context"

	"github.com/bububa/spotlight-mapi/core"
	"github.com/bububa/spotlight-mapi/model/finance"
)

// Transfer 转账接口
func Transfer(ctx context.Context, clt *core.SDKClient, req *finance.TransferRequest, accessToken string) (*finance.TransferResult, error) {
	var resp finance.TransferResponse
	if err := clt.Post(ctx, "/finance/agent/transfer", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
