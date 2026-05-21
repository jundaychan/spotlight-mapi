package simpledelivery

import (
	"context"

	"github.com/bububa/spotlight-mapi/core"
	"github.com/bububa/spotlight-mapi/model/simpledelivery"
)

// KeywordRealtime 简单投关键词层级实时数据
func KeywordRealtime(ctx context.Context, clt *core.SDKClient, req *simpledelivery.KeywordRealtimeRequest, accessToken string) (*simpledelivery.KeywordRealtimeData, error) {
	var resp simpledelivery.KeywordRealtimeResponse
	if err := clt.Post(ctx, "/jg/data/report/realtime/ube/keyword", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
