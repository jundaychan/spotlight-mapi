package reportrealtime2

import (
	"context"

	"github.com/jundaychan/spotlight-mapi/core"
	"github.com/jundaychan/spotlight-mapi/model/reportrealtime2"
)

// Target 定向层级实时数据
func Target(ctx context.Context, clt *core.SDKClient, req *reportrealtime2.TargetRequest, accessToken string) (*reportrealtime2.TargetResponse, error) {
	resp := new(reportrealtime2.TargetResponse)
	if err := clt.Post(ctx, "/jg/data/report/realtime/target", req, resp, accessToken); err != nil {
		return nil, err
	}
	return resp, nil
}
