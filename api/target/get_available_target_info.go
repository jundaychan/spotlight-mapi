package target

import (
	"context"

	"github.com/jundaychan/spotlight-mapi/core"
	"github.com/jundaychan/spotlight-mapi/model/target"
)

// GetAvailableTargetInfo 获取定向信息
func GetAvailableTargetInfo(ctx context.Context, clt *core.SDKClient, req *target.GetAvailableTargetInfoRequest, accessToken string) (*target.TargetInfo, error) {
	var resp target.GetAvailableTargetInfoResponse
	if err := clt.Post(ctx, "/jg/target/get_available_target_info", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
