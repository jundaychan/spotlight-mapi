package realtime

import (
	"context"

	"github.com/jundaychan/spotlight-mapi/core"
	"github.com/jundaychan/spotlight-mapi/model/report"
	"github.com/jundaychan/spotlight-mapi/model/report/realtime"
)

// Campaign 计划层级实时数据
func Campaign(ctx context.Context, clt *core.SDKClient, req *realtime.CampaignRequest, accessToken string) (*realtime.CampaignResponse, error) {
	resp := new(realtime.CampaignResponse)
	r := *req
	if len(r.Columns) == 0 {
		r.Columns = report.Columns("/jg/data/report/realtime/campaign")
	}
	if err := clt.Post(ctx, "/jg/data/report/realtime/campaign", &r, resp, accessToken); err != nil {
		return nil, err
	}
	return resp, nil
}
