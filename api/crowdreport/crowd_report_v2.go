package crowdreport

import (
	"context"

	"github.com/bububa/spotlight-mapi/core"
	"github.com/bububa/spotlight-mapi/model/crowdreport"
)

// crowdReportV2URL 人群包报表 V2 完整请求地址
const crowdReportV2URL = "https://ad.xiaohongshu.com/api/idea/group_report_v2"

// CrowdReportV2 人群包报表 V2
func CrowdReportV2(ctx context.Context, clt *core.SDKClient, req *crowdreport.CrowdReportV2Request, accessToken string) (*crowdreport.CrowdReportV2Data, error) {
	var resp crowdreport.CrowdReportV2Response
	if err := clt.Post(ctx, crowdReportV2URL, req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
