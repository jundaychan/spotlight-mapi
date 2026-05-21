package finance

import (
	"context"

	"github.com/jundaychan/spotlight-mapi/core"
	"github.com/jundaychan/spotlight-mapi/model/finance"
)

// CampaignFlow 获取计划流水接口
func CampaignFlow(ctx context.Context, clt *core.SDKClient, req *finance.CampaignFlowRequest, accessToken string) (*finance.CampaignFlowResult, error) {
	var resp finance.CampaignFlowResponse
	if err := clt.Post(ctx, "/jg/account/ad/order/info", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
