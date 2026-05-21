package simpledelivery

import (
	"context"

	"github.com/jundaychan/spotlight-mapi/core"
	"github.com/jundaychan/spotlight-mapi/model/simpledelivery"
)

// TargetRealtime 简单投标的层级实时数据
func TargetRealtime(ctx context.Context, clt *core.SDKClient, req *simpledelivery.TargetRealtimeRequest, accessToken string) (*simpledelivery.TargetRealtimeData, error) {
	var resp simpledelivery.TargetRealtimeResponse
	if err := clt.Post(ctx, "/jg/data/report/realtime/ube/group", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
