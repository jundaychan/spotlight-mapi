package negativekeyword

import (
	"context"

	"github.com/jundaychan/spotlight-mapi/core"
	"github.com/jundaychan/spotlight-mapi/model/negativekeyword"
)

// BatchAdd 批量添加否定词
func BatchAdd(ctx context.Context, clt *core.SDKClient, req *negativekeyword.BatchAddRequest, accessToken string) error {
	return clt.Post(ctx, "/jg/negative/keyword/batch/add", req, nil, accessToken)
}
