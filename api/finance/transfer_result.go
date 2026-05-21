package finance

import (
	"context"

	"github.com/jundaychan/spotlight-mapi/core"
	"github.com/jundaychan/spotlight-mapi/model/finance"
)

// TransferResult 查询转账结果接口
func TransferResult(ctx context.Context, clt *core.SDKClient, req *finance.TransferResultRequest, accessToken string) (*finance.TransferResultData, error) {
	var resp finance.TransferResultResponse
	if err := clt.Post(ctx, "/finance/transfer/result/query", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
