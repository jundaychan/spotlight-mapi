package realtime

import (
	"context"

	"github.com/jundaychan/spotlight-mapi/core"
	"github.com/jundaychan/spotlight-mapi/model/report"
	"github.com/jundaychan/spotlight-mapi/model/report/realtime"
)

// Keyword 关键词层级实时数据
func Keyword(ctx context.Context, clt *core.SDKClient, req *realtime.KeywordRequest, accessToken string) (*realtime.KeywordResponse, error) {
	resp := new(realtime.KeywordResponse)
	r := *req
	if len(r.Columns) == 0 {
		r.Columns = report.Columns("/jg/data/report/realtime/keyword")
	}
	if err := clt.Post(ctx, "/jg/data/report/realtime/keyword", &r, resp, accessToken); err != nil {
		return nil, err
	}
	return resp, nil
}
