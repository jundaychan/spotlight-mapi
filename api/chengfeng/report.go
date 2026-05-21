// Package chengfeng 乘风(wind)开放平台报表接口。
//
// 网关与聚光共用 core.SDKClient(BASE_URL=adapi.xiaohongshu.com/api/open + Access-Token)，
// 仅路径前缀为 /wind/...。离线报表请求/响应直接复用 model/report/offline；
// 实时报表用 model/chengfeng 的容错模型。
package chengfeng

import (
	"context"

	mdlcf "github.com/jundaychan/spotlight-mapi/model/chengfeng"
	"github.com/jundaychan/spotlight-mapi/model/report"
	"github.com/jundaychan/spotlight-mapi/model/report/offline"

	"github.com/jundaychan/spotlight-mapi/core"
)

// ===== 离线报表 /wind/data/report/offline/{level} (复用聚光 offline 模型) =====

func postOffline(ctx context.Context, clt *core.SDKClient, gw string, req *offline.Request, accessToken string) (*offline.ReportList, error) {
	var resp offline.Response
	if err := clt.Post(ctx, gw, req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}

// OfflineAccount 账户层级离线报表
func OfflineAccount(ctx context.Context, clt *core.SDKClient, req *offline.Request, accessToken string) (*offline.ReportList, error) {
	return postOffline(ctx, clt, "/wind/data/report/offline/account", req, accessToken)
}

// OfflineCampaign 计划层级离线报表
func OfflineCampaign(ctx context.Context, clt *core.SDKClient, req *offline.Request, accessToken string) (*offline.ReportList, error) {
	return postOffline(ctx, clt, "/wind/data/report/offline/campaign", req, accessToken)
}

// OfflineCreativity 创意层级离线报表
func OfflineCreativity(ctx context.Context, clt *core.SDKClient, req *offline.Request, accessToken string) (*offline.ReportList, error) {
	return postOffline(ctx, clt, "/wind/data/report/offline/creativity", req, accessToken)
}

// OfflineNote 笔记层级离线报表
func OfflineNote(ctx context.Context, clt *core.SDKClient, req *offline.Request, accessToken string) (*offline.ReportList, error) {
	return postOffline(ctx, clt, "/wind/data/report/offline/note", req, accessToken)
}

// OfflineKeyword 关键词层级离线报表
func OfflineKeyword(ctx context.Context, clt *core.SDKClient, req *offline.Request, accessToken string) (*offline.ReportList, error) {
	return postOffline(ctx, clt, "/wind/data/report/offline/keyword", req, accessToken)
}

// OfflineSpu 商品层级离线报表
func OfflineSpu(ctx context.Context, clt *core.SDKClient, req *offline.Request, accessToken string) (*offline.ReportList, error) {
	return postOffline(ctx, clt, "/wind/data/report/offline/spu", req, accessToken)
}

// OfflineLive 直播间层级离线报表
func OfflineLive(ctx context.Context, clt *core.SDKClient, req *offline.Request, accessToken string) (*offline.ReportList, error) {
	return postOffline(ctx, clt, "/wind/data/report/offline/live", req, accessToken)
}

// ===== 实时报表 /wind/data/report/realtime/{level} =====

// RealtimeAccount 账户层级实时数据(返回单条汇总)。
func RealtimeAccount(ctx context.Context, clt *core.SDKClient, req *mdlcf.RealtimeRequest, accessToken string) (*report.DataReportDTO, error) {
	var resp mdlcf.RealtimeAccountResponse
	if err := clt.Post(ctx, "/wind/data/report/realtime/account", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}

// RealtimeCampaign 计划层级实时数据。
func RealtimeCampaign(ctx context.Context, clt *core.SDKClient, req *mdlcf.RealtimeRequest, accessToken string) (*mdlcf.RealtimeList, error) {
	var resp mdlcf.RealtimeListResponse
	if err := clt.Post(ctx, "/wind/data/report/realtime/campaign", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}

// RealtimeCreativity 创意层级实时数据。
func RealtimeCreativity(ctx context.Context, clt *core.SDKClient, req *mdlcf.RealtimeRequest, accessToken string) (*mdlcf.RealtimeList, error) {
	var resp mdlcf.RealtimeListResponse
	if err := clt.Post(ctx, "/wind/data/report/realtime/creativity", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
