package simpledelivery

import (
	"context"

	"github.com/jundaychan/spotlight-mapi/core"
	"github.com/jundaychan/spotlight-mapi/model/simpledelivery"
)

// NoteReport 简单投笔记报表（离线）
func NoteReport(ctx context.Context, clt *core.SDKClient, req *simpledelivery.NoteReportRequest, accessToken string) (*simpledelivery.NoteReportData, error) {
	var resp simpledelivery.NoteReportResponse
	if err := clt.Post(ctx, "/jg/data/report/offline/easy/promotion/note", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
