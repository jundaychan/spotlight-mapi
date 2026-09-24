package realtime

import (
	"context"

	"github.com/jundaychan/spotlight-mapi/core"
	"github.com/jundaychan/spotlight-mapi/model/report"
	"github.com/jundaychan/spotlight-mapi/model/report/realtime"
)

// Advertiser 账户层级实时数据
func Advertiser(ctx context.Context, clt *core.SDKClient, req *realtime.AdvertiserRequest, accessToken string) (*report.DataReportDTO, error) {
	var resp realtime.AdvertiserResponse
	r := *req
	if len(r.Columns) == 0 {
		r.Columns = report.Columns("/jg/data/report/realtime/account")
	}
	if err := clt.Post(ctx, "/jg/data/report/realtime/account", &r, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
