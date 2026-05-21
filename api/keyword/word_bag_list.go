package keyword

import (
	"context"

	"github.com/jundaychan/spotlight-mapi/core"
	"github.com/jundaychan/spotlight-mapi/model/keyword"
)

// WordBagList 词包推荐
func WordBagList(ctx context.Context, clt *core.SDKClient, req *keyword.WordBagListRequest, accessToken string) (*keyword.WordBagListResult, error) {
	var resp keyword.WordBagListResponse
	if err := clt.Post(ctx, "/jg/keyword/word/bag/list", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
