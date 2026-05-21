package negativekeyword

import (
	"context"

	"github.com/jundaychan/spotlight-mapi/core"
	"github.com/jundaychan/spotlight-mapi/model/negativekeyword"
)

// BatchDelete 批量删除否定词
func BatchDelete(ctx context.Context, clt *core.SDKClient, req *negativekeyword.BatchDeleteRequest, accessToken string) error {
	return clt.Post(ctx, "/jg/negative/keyword/batch/delete", req, nil, accessToken)
}
