package realtime

import (
	"context"

	"github.com/jundaychan/spotlight-mapi/core"
	"github.com/jundaychan/spotlight-mapi/model/report"
	"github.com/jundaychan/spotlight-mapi/model/report/realtime"
)

// Creativity 创意层级实时数据
func Creativity(ctx context.Context, clt *core.SDKClient, req *realtime.CreativityRequest, accessToken string) (*realtime.CreativityResponse, error) {
	resp := new(realtime.CreativityResponse)
	r := *req
	if len(r.Columns) == 0 {
		r.Columns = report.Columns("/jg/data/report/realtime/creativity")
	}
	if err := clt.Post(ctx, "/jg/data/report/realtime/creativity", &r, resp, accessToken); err != nil {
		return nil, err
	}
	return resp, nil
}
