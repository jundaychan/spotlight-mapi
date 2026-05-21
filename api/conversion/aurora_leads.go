package conversion

import (
	"context"

	"github.com/jundaychan/spotlight-mapi/core"
	"github.com/jundaychan/spotlight-mapi/model/conversion"
)

// AuroraLeads 线索转化数据回传
func AuroraLeads(ctx context.Context, clt *core.SDKClient, req *conversion.AuroraLeadsRequest) error {
	return clt.Post(ctx, "/common", req, nil, "")
}
