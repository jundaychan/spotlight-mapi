package crowdreport

import (
	"context"

	"github.com/bububa/spotlight-mapi/core"
	"github.com/bububa/spotlight-mapi/model/crowdreport"
)

// crowdReportURL 人群包报表完整请求地址
const crowdReportURL = "https://ad.xiaohongshu.com/api/idea/group_report"

// CrowdReport 人群包报表
func CrowdReport(ctx context.Context, clt *core.SDKClient, req *crowdreport.CrowdReportRequest, accessToken string) (*crowdreport.CrowdReportResponse, error) {
	resp := new(crowdreport.CrowdReportResponse)
	if err := clt.Post(ctx, crowdReportURL, req, resp, accessToken); err != nil {
		return nil, err
	}
	return resp, nil
}
