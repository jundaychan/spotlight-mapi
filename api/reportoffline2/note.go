package reportoffline2

import (
	"context"

	"github.com/bububa/spotlight-mapi/core"
	"github.com/bububa/spotlight-mapi/model/reportoffline2"
)

// Note 笔记层级离线报表数据
func Note(ctx context.Context, clt *core.SDKClient, req *reportoffline2.NoteRequest, accessToken string) (*reportoffline2.NoteReportList, error) {
	var resp reportoffline2.NoteResponse
	if err := clt.Post(ctx, "/jg/data/report/offline/note", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
