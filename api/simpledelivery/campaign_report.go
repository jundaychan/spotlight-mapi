package simpledelivery

import (
	"context"

	"github.com/jundaychan/spotlight-mapi/core"
	"github.com/jundaychan/spotlight-mapi/model/simpledelivery"
)

// CampaignReport 简单投计划报表（离线）
func CampaignReport(ctx context.Context, clt *core.SDKClient, req *simpledelivery.CampaignReportRequest, accessToken string) (*simpledelivery.CampaignReportData, error) {
	var resp simpledelivery.CampaignReportResponse
	if err := clt.Post(ctx, "/jg/data/report/offline/easy/promotion/base", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
