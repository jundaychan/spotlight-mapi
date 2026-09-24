package offline

import (
	"context"

	"github.com/jundaychan/spotlight-mapi/core"
	"github.com/jundaychan/spotlight-mapi/model/report"
	"github.com/jundaychan/spotlight-mapi/model/report/offline"
)

// Advertiser 账户层级离线数据
func Advertiser(ctx context.Context, clt *core.SDKClient, req *offline.Request, accessToken string) (*offline.ReportList, error) {
	var resp offline.Response
	r := *req
	if len(r.Columns) == 0 {
		r.Columns = report.Columns("/jg/data/report/offline/account")
	}
	if err := clt.Post(ctx, "/jg/data/report/offline/account", &r, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
