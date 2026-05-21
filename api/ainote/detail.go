package ainote

import (
	"context"

	"github.com/jundaychan/spotlight-mapi/core"
	"github.com/jundaychan/spotlight-mapi/model/ainote"
)

// Detail AI 智能笔记详情，获取 AIGC 广告单元维度的实时报表数据
func Detail(ctx context.Context, clt *core.SDKClient, req *ainote.DetailRequest, accessToken string) (*ainote.DetailData, error) {
	var resp ainote.DetailResponse
	if err := clt.Post(ctx, "/jg/data/report/realtime/unit/aigc", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
