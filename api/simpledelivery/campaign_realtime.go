package simpledelivery

import (
	"context"

	"github.com/bububa/spotlight-mapi/core"
	"github.com/bububa/spotlight-mapi/model/simpledelivery"
)

// CampaignRealtime 简单投计划层级实时数据
func CampaignRealtime(ctx context.Context, clt *core.SDKClient, req *simpledelivery.CampaignRealtimeRequest, accessToken string) (*simpledelivery.CampaignRealtimeData, error) {
	var resp simpledelivery.CampaignRealtimeResponse
	if err := clt.Post(ctx, "/jg/data/report/realtime/ube/campaign", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
