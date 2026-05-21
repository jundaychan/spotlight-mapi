package reportoffline2

import (
	"context"

	"github.com/jundaychan/spotlight-mapi/core"
	"github.com/jundaychan/spotlight-mapi/model/reportoffline2"
)

// SPU SPU层级离线报表数据
func SPU(ctx context.Context, clt *core.SDKClient, req *reportoffline2.SPURequest, accessToken string) (*reportoffline2.SPUReportList, error) {
	var resp reportoffline2.SPUResponse
	if err := clt.Post(ctx, "/jg/data/report/offline/spu", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
