package reportoffline2

import (
	"context"

	"github.com/bububa/spotlight-mapi/core"
	"github.com/bububa/spotlight-mapi/model/reportoffline2"
)

// SearchWord 搜索词层级离线报表数据
func SearchWord(ctx context.Context, clt *core.SDKClient, req *reportoffline2.SearchWordRequest, accessToken string) (*reportoffline2.SearchWordReportList, error) {
	var resp reportoffline2.SearchWordResponse
	if err := clt.Post(ctx, "/jg/data/report/offline/search/word", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
