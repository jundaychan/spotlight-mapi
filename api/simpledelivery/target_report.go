package simpledelivery

import (
	"context"

	"github.com/bububa/spotlight-mapi/core"
	"github.com/bububa/spotlight-mapi/model/simpledelivery"
)

// TargetReport 简单投标的报表（离线）
func TargetReport(ctx context.Context, clt *core.SDKClient, req *simpledelivery.TargetReportRequest, accessToken string) (*simpledelivery.TargetReportData, error) {
	var resp simpledelivery.TargetReportResponse
	if err := clt.Post(ctx, "/jg/data/report/offline/easy/promotion/group", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
