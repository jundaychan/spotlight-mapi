package offline

import (
	"context"

	"github.com/jundaychan/spotlight-mapi/core"
	"github.com/jundaychan/spotlight-mapi/model/report"
	"github.com/jundaychan/spotlight-mapi/model/report/offline"
)

// Creativity 广告创意层级离线数据。官方路径是 .../offline/creative，旧写法 .../creativity 上游回 404。
func Creativity(ctx context.Context, clt *core.SDKClient, req *offline.Request, accessToken string) (*offline.ReportList, error) {
	var resp offline.Response
	r := *req
	if len(r.Columns) == 0 {
		r.Columns = report.Columns("/jg/data/report/offline/creative")
	}
	if err := clt.Post(ctx, "/jg/data/report/offline/creative", &r, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
