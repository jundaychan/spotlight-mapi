package simpledelivery

import (
	"context"

	"github.com/jundaychan/spotlight-mapi/core"
	"github.com/jundaychan/spotlight-mapi/model/simpledelivery"
)

// NoteRealtime 简单投笔记层级实时数据
func NoteRealtime(ctx context.Context, clt *core.SDKClient, req *simpledelivery.NoteRealtimeRequest, accessToken string) (*simpledelivery.NoteRealtimeData, error) {
	var resp simpledelivery.NoteRealtimeResponse
	if err := clt.Post(ctx, "/jg/data/report/realtime/ube/note", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
