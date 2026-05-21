package oauth

import (
	"context"

	"github.com/jundaychan/spotlight-mapi/core"
	"github.com/jundaychan/spotlight-mapi/model/oauth"
	"github.com/jundaychan/spotlight-mapi/util"
)

// URL 生成oauth授权链接
func URL(ctx context.Context, clt *core.SDKClient, req *oauth.URLRequest) string {
	req.AppID = clt.AppID()
	return util.StringsJoin(core.OAUTH_URL, "?", req.Encode())
}
